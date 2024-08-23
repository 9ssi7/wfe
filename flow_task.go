package wfe

import (
	"context"
)

type taskFlow[P any] struct {
	name    string
	nodes   []Node
	actions map[string]Action[P]
}

func (f *taskFlow[P]) Name() string {
	return f.name
}

func (f *taskFlow[P]) Kind() Kind {
	return KindTask
}

func (f *taskFlow[P]) AddNode(n ...Node) {
	f.nodes = append(f.nodes, n...)
}

func (f *taskFlow[P]) AddAction(ref string, fn ActionRunner[P], errRef ...string) {
	if f.actions == nil {
		f.actions = make(map[string]Action[P])
	}
	f.actions[ref] = NewAction(ref, fn, errRef...)
}

func (f *taskFlow[P]) GetAction(ref string) (Action[P], bool) {
	a, ok := f.actions[ref]
	return a, ok
}

func (f *taskFlow[P]) Run(ctx context.Context, p P) error {
	return run(ctx, p, f, f.nodes)
}

func (f *taskFlow[P]) Cancel(ctx context.Context) error {
	return nil
}
