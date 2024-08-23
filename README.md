## wfe - Go Workflow Engine

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![GoDoc](https://godoc.org/github.com/9ssi7/wfe?status.svg)](https://pkg.go.dev/github.com/9ssi7/wfe)
![Project status](https://img.shields.io/badge/version-0.0.1-green.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/9ssi7/wfe)](https://goreportcard.com/report/github.com/9ssi7/wfe)

`wfe` is a lightweight and flexible engine that allows you to define and manage workflows in Go. It enables you to easily model and execute a variety of workflows, from simple tasks to complex processes.

### Features

- **Node-based structure:** Structure your workflows as a series of steps (nodes).
- **Actions:** Define actions to be executed at each step.
- **Type safety:** Ensure type safety in your workflows using generics.
- **Easy integration:** Easily integrate into your existing Go projects.
- **Flexibility:** Customizable to support different workflow scenarios.

### Installation

```bash
go get github.com/9ssi7/wfe
```

### Simple Example

```go
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
```

### Kanban Application Example

For a more complex example, see the [Kanban application example](./examples/kanban/main.go).

### Recommended Use Cases

- **Business Process Automation:** Model and automate complex business processes.
- **Event-driven Systems:** Handle events and trigger corresponding workflows.
- **User Interface Interactions:** Orchestrate UI interactions and backend logic.
- **Algorithm Generation:** Allow users to create algorithms through a visual interface, defining the steps and actions within the workflow.

### Contributing

Contributions are welcome! Feel free to submit bug reports, feature requests, or pull requests.

### License

This project is licensed under the Apache License 2.0. See the `LICENSE` file for the full license text. 
