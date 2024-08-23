package wfe_test

import (
	"testing"

	"github.com/9ssi7/wfe"
)

func TestNewNode(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name      string
		inputName string
		inputRef  []string
		expected  wfe.Node
	}{
		{
			name:      "No actionRef provided",
			inputName: "Node1",
			expected: wfe.Node{
				Name:      "Node1",
				ActionRef: "Node1",
			},
		},
		{
			name:      "actionRef provided",
			inputName: "Node2",
			inputRef:  []string{"ActionRef2"},
			expected: wfe.Node{
				Name:      "Node2",
				ActionRef: "ActionRef2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := wfe.NewNode(test.inputName, test.inputRef...)
			if result.Name != test.expected.Name || result.ActionRef != test.expected.ActionRef {
				t.Errorf("Expected %+v, but got %+v", test.expected, result)
			}
		})
	}
}
