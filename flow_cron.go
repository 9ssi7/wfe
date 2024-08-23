package wfe

import (
	"context"
	"sync"

	"github.com/robfig/cron/v3"
)

type cronFlow[P any] struct {
	name    string
	trigger string
	nodes   []Node
	actions map[string]Action[P]

	crons []*cron.Cron
	wgs   []*sync.WaitGroup
}

func (f *cronFlow[P]) Name() string {
	return f.name
}

func (f *cronFlow[P]) Kind() Kind {
	return KindCron
}

func (f *cronFlow[P]) AddNode(n ...Node) {
	f.nodes = append(f.nodes, n...)
}

func (f *cronFlow[P]) AddAction(ref string, fn ActionRunner[P], errRef ...string) {
	if f.actions == nil {
		f.actions = make(map[string]Action[P])
	}
	f.actions[ref] = NewAction(ref, fn, errRef...)
}

func (f *cronFlow[P]) GetAction(ref string) (Action[P], bool) {
	act, ok := f.actions[ref]
	return act, ok
}

func (f *cronFlow[P]) Run(ctx context.Context, p P) error {
	c := cron.New()
	wg := sync.WaitGroup{}
	f.crons = append(f.crons, c)
	f.wgs = append(f.wgs, &wg)
	var err error
	wg.Add(1)
	c.AddFunc(f.trigger, func() {
		defer wg.Done()
		err = run(ctx, p, f, f.nodes)
	})
	c.Start()
	wg.Wait()
	return err
}

func (f *cronFlow[P]) Cancel(ctx context.Context) error {
	for _, c := range f.crons {
		c.Stop()
	}
	return nil
}
