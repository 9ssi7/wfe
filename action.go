package wfe

import "context"

// ActionRunner is a function that runs an action with the given context and payload.
type ActionRunner[P any] func(ctx context.Context, p P) error

// Action is an interface representing an action in a workflow.
type Action[P any] interface {
	// Run executes the action with the given context and payload.
	Run(ctx context.Context, p P) error
	Reference() string
	// ErrorRef returns the error reference of the action, if any.
	ErrorRef() *string
}

type action[P any] struct {
	ref    string
	fn     ActionRunner[P]
	errRef []string
}

func NewAction[P any](ref string, fn ActionRunner[P], errRef ...string) Action[P] {
	return &action[P]{ref, fn, errRef}
}

func (a *action[P]) Run(ctx context.Context, p P) error {
	return a.fn(ctx, p)
}

func (a *action[P]) Reference() string {
	return a.ref
}

func (a *action[P]) ErrorRef() *string {
	if len(a.errRef) == 0 {
		return nil
	}
	return &a.errRef[0]
}
