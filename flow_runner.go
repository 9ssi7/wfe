package wfe

import "context"

func run[P any](ctx context.Context, p P, f Flow[P], nodes []Node) error {
	for _, n := range nodes {
		a, exists := f.GetAction(n.ActionRef)
		if exists {
			err := a.Run(ctx, p)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
