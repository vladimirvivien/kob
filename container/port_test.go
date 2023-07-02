package container

import (
	"reflect"
	"testing"

	coreV1 "k8s.io/api/core/v1"
)

func TestContainerPort(t *testing.T) {
	tests := map[string]struct {
		builder  PortBuilder
		expected coreV1.ContainerPort
	}{
		"empty container": {
			builder:  PortBuilder{},
			expected: coreV1.ContainerPort{},
		},
		"host port": {
			builder:  HostPort(12),
			expected: coreV1.ContainerPort{HostPort: 12},
		},
		"container port": {
			builder:  ContainerPort(12),
			expected: coreV1.ContainerPort{ContainerPort: 12},
		},
		"port with name": {
			builder:  HostPort(12).Name("standard"),
			expected: coreV1.ContainerPort{HostPort: 12, Name: "standard"},
		},
		"all": {
			builder:  HostPort(12).ContainerPort(13).Name("standard").Protocol("basic").HostIP("unknown"),
			expected: coreV1.ContainerPort{HostPort: 12, Name: "standard", ContainerPort: 13, Protocol: "basic", HostIP: "unknown"},
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
