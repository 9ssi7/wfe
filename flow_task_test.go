package wfe

import (
	"context"
	"errors"
	"testing"
)

func TestTaskFlow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		flowName  string
		nodes     []Node
		actions   map[string]Action[int]
		wantName  string
		wantKind  Kind
		getAction string
		wantOk    bool
		wantErr   bool
	}{
		{
			name:     "Basic Flow",
			flowName: "testFlow",
			nodes:    []Node{},
			actions: map[string]Action[int]{
				"testAction": NewAction("testAction", func(ctx context.Context, p int) error {
					return nil
				}),
			},
			wantName:  "testFlow",
			wantKind:  KindTask,
			getAction: "testAction",
			wantOk:    true,
			wantErr:   false,
		},
		{
			name:      "No Actions",
			flowName:  "noActionsFlow",
			nodes:     []Node{},
			actions:   nil,
			wantName:  "noActionsFlow",
			wantKind:  KindTask,
			getAction: "nonExistentAction",
			wantOk:    false,
			wantErr:   false,
		},
		{
			name:     "Error in Run",
			flowName: "errorFlow",
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
			wantName:  "errorFlow",
			wantKind:  KindTask,
			getAction: "errorAction",
			wantOk:    true,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flow := &taskFlow[int]{
				name:    tt.flowName,
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

			err := flow.Run(context.Background(), 10)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := flow.Cancel(context.Background()); err != nil {
				t.Errorf("Cancel() error = %v, want nil", err)
			}
		})
	}

	t.Run("AddAction", func(t *testing.T) {
		flow := &taskFlow[int]{
			name:  "testFlow",
			nodes: []Node{},
		}

		flow.AddAction("testAction", func(ctx context.Context, p int) error {
			return nil
		})

		action, ok := flow.GetAction("testAction")
		if !ok {
			t.Errorf("GetAction() ok = %v, want true", ok)
		}
		if action.Reference() != "testAction" {
			t.Errorf("GetAction() action ref = %v, want testAction", action.Reference())
		}
	})

	t.Run("AddNode", func(t *testing.T) {
		flow := &taskFlow[int]{
			name:  "testFlow",
			nodes: []Node{},
		}

		flow.AddNode(Node{
			ActionRef: "testAction",
		})

		if len(flow.nodes) != 1 {
			t.Errorf("AddNode() len(nodes) = %v, want 1", len(flow.nodes))
		}
		if flow.nodes[0].ActionRef != "testAction" {
			t.Errorf("AddNode() node action ref = %v, want testAction", flow.nodes[0].ActionRef)
		}
	})
}
