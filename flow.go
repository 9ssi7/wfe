package wfe

import (
	"context"
)

// Kind represents the type of a flow.
type Kind int

const (
	KindTask Kind = iota
	KindEvent
	KindCron
)

// Flow is an interface representing a workflow.
type Flow[P any] interface {
	Name() string
	Kind() Kind

	// AddNode adds one or more nodes to the flow.
	AddNode(n ...Node)

	// AddAction adds an action to the flow with the given reference, runner function, and optional error reference.
	AddAction(ref string, fn ActionRunner[P], errRef ...string)

	// GetAction retrieves an action from the flow by its reference.
	GetAction(ref string) (Action[P], bool)

	// Run executes the flow with the given context and payload.
	Run(ctx context.Context, p P) error

	// Cancel cancels the flow execution with the given context.
	Cancel(ctx context.Context) error
}

// NewWithCron creates a new cron flow with the given name and trigger.
func NewWithCron[P any](name, trigger string) Flow[P] {
	return &cronFlow[P]{
		name:    name,
		trigger: trigger,
	}
}

// New creates a new task flow with the given name.
func New[P any](name string) Flow[P] {
	return &taskFlow[P]{
		name: name,
	}
}
