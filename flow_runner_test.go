package wfe

import (
	"context"
	"errors"
	"testing"
)

func TestRun(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		flow := New[any]("Flow1")

		flow.AddAction("ActionRef1", func(ctx context.Context, p any) error {
			return nil
		})
		nodes := []Node{
			{
				Name:      "Node1",
				ActionRef: "ActionRef1",
			},
			{
				Name:      "Node2",
				ActionRef: "ActionRef2",
			},
		}

		err := run(ctx, nil, flow, nodes)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
	})

	t.Run("error", func(t *testing.T) {
		flow := New[any]("Flow1")

		flow.AddAction("ActionRef1", func(ctx context.Context, p any) error {
			return nil
		})

		flow.AddAction("ActionRef2", func(ctx context.Context, p any) error {
			return errors.New("ActionRef2 failed")
		})

		nodes := []Node{
			{
				Name:      "Node1",
				ActionRef: "ActionRef1",
			},
			{
				Name:      "Node2",
				ActionRef: "ActionRef2",
			},
		}

		err := run(ctx, nil, flow, nodes)
		if err == nil {
			t.Errorf("Expected an error, but got nil")
		}
	})
}
