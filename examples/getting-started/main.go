package main

import (
	"context"

	"github.com/9ssi7/wfe"
)

type Todo struct {
	Title string
}

func main() {
	flow := wfe.New[Todo]("todo")
	flow.AddAction("print", func(ctx context.Context, p Todo) error {
		println(p.Title)
		return nil
	})
	flow.AddNode(wfe.NewNode("print"))
	flow.Run(context.Background(), Todo{Title: "Hello, World!"})
}
