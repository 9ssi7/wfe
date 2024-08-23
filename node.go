package wfe

// Node represents a node in a workflow.
type Node struct {
	// Name of the node.
	Name string
	// ActionRef is the reference to the action associated with the node.
	ActionRef string
	// Env is a map of environment variables for the node.
	Env map[string]any
}

// NewNode creates a new Node with the given name and optional action reference.
// If actionRef is provided, it will be used as the ActionRef; otherwise, the name will be used.
func NewNode(name string, actionRef ...string) Node {
	ref := name
	if len(actionRef) > 0 {
		ref = actionRef[0]
	}
	return Node{
		Name:      name,
		ActionRef: ref,
	}
}
