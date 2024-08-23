// Package wfe provides a workflow engine.
//
// This package allows you to define and execute workflows consisting of nodes and actions.
// You can create different types of flows, such as task flows and cron flows, and add nodes and
// actions to them. Actions can be defined using `ActionRunner` functions, and flows can be
// executed with a context and payload.
//
// Example usage:
//
//	package main
//
//	import (
//	    "context"
//	    "fmt"
//	    "wfe"
//	)
//
//	func main() {
//	    flow := wfe.New[string]("myFlow")
//
//	    flow.AddAction("hello", func(ctx context.Context, name string) error {
//	        fmt.Println("Hello,", name)
//	        return nil
//	    })
//
//	    flow.AddNode(wfe.NewNode("helloNode", "hello"))
//
//	    if err := flow.Run(context.Background(), "World"); err != nil {
//	        fmt.Println("Error:", err)
//	    }
//	}
package wfe
