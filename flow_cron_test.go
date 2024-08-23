package wfe

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestCronFlow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		flowName  string
		trigger   string
		nodes     []Node
		actions   map[string]Action[int]
		wantName  string
		wantKind  Kind
		getAction string
		wantOk    bool
		wantErr   bool
	}{
		{
			name:     "Basic Cron Flow",
			flowName: "testCronFlow",
			trigger:  "@every 1s",
			nodes:    []Node{},
			actions: map[string]Action[int]{
				"testAction": NewAction("testAction", func(ctx context.Context, p int) error {
					return nil
				}),
			},
			wantName:  "testCronFlow",
			wantKind:  KindCron,
			getAction: "testAction",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "No Actions",
			flowName:  "noActionsCronFlow",
			trigger:   "@every 1s",
			nodes:     []Node{},
			actions:   nil,
			wantName:  "noActionsCronFlow",
			wantKind:  KindCron,
			getAction: "nonExistentAction",
			wantOk:    false,
			wantErr:   false,
		},
		{
			name:     "Error in Run",
			flowName: "errorCronFlow",
			trigger:  "@every 1s",
			nodes: []Node{
				{
					ActionRef: "errorAction",
				},
			},
			actions: map[string]Action[int]{
				"errorAction": NewAction("errorAction", func(ctx context.Context, p int) error {
					return errors.New("some error")
				}),
			},
			wantName:  "errorCronFlow",
			wantKind:  KindCron,
			getAction: "errorAction",
			wantOk:    true,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flow := &cronFlow[int]{
				name:    tt.flowName,
				trigger: tt.trigger,
				nodes:   tt.nodes,
				actions: tt.actions,
			}

			if got := flow.Name(); got != tt.wantName {
				t.Errorf("Name() = %v, want %v", got, tt.wantName)
			}

			if got := flow.Kind(); got != tt.wantKind {
				t.Errorf("Kind() = %v, want %v", got, tt.wantKind)
			}

			action, ok := flow.GetAction(tt.getAction)
			if ok != tt.wantOk {
				t.Errorf("GetAction(%v) ok = %v, want %v", tt.getAction, ok, tt.wantOk)
			}
			if tt.wantOk && action.Reference() != tt.getAction {
				t.Errorf("GetAction(%v) action ref = %v, want %v", tt.getAction, action.Reference(), tt.getAction)
			}

			go func() {
				time.Sleep(2 * time.Second)
				flow.Cancel(context.Background())
			}()

			err := flow.Run(context.Background(), 10)

			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}

			// Cancel fonksiyonunu test ediyoruz
			if err := flow.Cancel(context.Background()); err != nil {
				t.Errorf("Cancel() error = %v, want nil", err)
			}
		})
	}

	t.Run("AddNode", func(t *testing.T) {
		flow := &cronFlow[int]{
			name:    "testFlow",
			trigger: "@every 1s",
			nodes:   []Node{},
			actions: nil,
		}

		flow.AddNode(Node{
			ActionRef: "testAction",
		})

		if len(flow.nodes) != 1 {
			t.Errorf("AddNode() len(nodes) = %v, want 1", len(flow.nodes))
		}
	})

	t.Run("AddAction", func(t *testing.T) {
		flow := &cronFlow[int]{
			name:    "testFlow",
			trigger: "@every 1s",
			nodes:   []Node{},
			actions: nil,
		}

		flow.AddAction("testAction", func(ctx context.Context, p int) error {
			return nil
		})

		if len(flow.actions) != 1 {
			t.Errorf("AddAction() len(actions) = %v, want 1", len(flow.actions))
		}
	})
}
