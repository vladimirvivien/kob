package container

import (
	"reflect"
	"testing"

	coreV1 "k8s.io/api/core/v1"
)

func TestContainerStructured(t *testing.T) {
	tests := map[string]struct {
		builder  Builder
		expected coreV1.Container
	}{
		"empty container": {
			builder:  Builder{},
			expected: coreV1.Container{},
		},
		"name only": {
			builder:  Name("simple-name"),
			expected: coreV1.Container{Name: "simple-name"},
		},
		"name and image": {
			builder:  Name("simple-name").Image("simple-container"),
			expected: coreV1.Container{Name: "simple-name", Image: "simple-container"},
		},
		"with name and image": {
			builder:  WithNameAndImage("simple-name", "simple-container"),
			expected: coreV1.Container{Name: "simple-name", Image: "simple-container"},
		},
		"with args": {
			builder:  Name("simple-name").Image("simple-container").Args("arg1", "arg2"),
			expected: coreV1.Container{Name: "simple-name", Image: "simple-container", Args: []string{"arg1", "arg2"}},
		},
		"with commands": {
			builder:  Name("simple-name").Image("simple-container").Commands("cmd1", "cmd2"),
			expected: coreV1.Container{Name: "simple-name", Image: "simple-container", Command: []string{"cmd1", "cmd2"}},
		},
		"working dir": {
			builder:  Name("simple-name").Image("simple-container").WorkingDir("workdir"),
			expected: coreV1.Container{Name: "simple-name", Image: "simple-container", WorkingDir: "workdir"},
		},
		"from unstructured": {
			builder: func() Builder {
				b, _ := FromUnstructured(map[string]any{"name": "simple-name", "image": "simple-container"})
				return b
			}(),
			expected: coreV1.Container{Name: "simple-name", Image: "simple-container"},
		},
		"from string": {
			builder: func() Builder {
				b, _ := FromString(`{"name":"simple-name","image":"simple-container"}`)
				return b
			}(),
			expected: coreV1.Container{Name: "simple-name", Image: "simple-container"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			container := test.builder.T()
			if !reflect.DeepEqual(container, test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", container, test.expected)
			}
		})
	}
}
