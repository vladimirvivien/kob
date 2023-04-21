package pod

import (
	"reflect"
	"testing"

	"github.com/vladimirvivien/kob/container"
	coreV1 "k8s.io/api/core/v1"
)

func TestPodSpecUnstructured(t *testing.T) {
	tests := map[string]struct {
		builder  SpecBuilder
		expected map[string]interface{}
	}{
		"empty spec": {
			builder:  SpecBuilder{},
			expected: map[string]interface{}{},
		},
		"spec with one container": {
			builder:  Spec(container.Name("simple-container")),
			expected: map[string]interface{}{"containers": []interface{}{map[string]interface{}{"name": "simple-container"}}},
		},
		"spec with two containers": {
			builder:  Spec(container.Name("simple-container-1"), container.Name("simple-container-2")),
			expected: map[string]interface{}{"containers": []interface{}{map[string]interface{}{"name": "simple-container-1"}, map[string]interface{}{"name": "simple-container-2"}}},
		},
		"spec with container name and image": {
			builder:  Spec(container.Name("simple-name").Image("simple-image")),
			expected: map[string]interface{}{"containers": []interface{}{map[string]interface{}{"name": "simple-name", "image": "simple-image"}}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			spec := test.builder
			if !reflect.DeepEqual(spec.U(), test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", spec.U(), test.expected)
			}
		})
	}
}

func TestPodSpecStructured(t *testing.T) {
	tests := map[string]struct {
		builder  SpecBuilder
		expected coreV1.PodSpec
	}{
		"empty spec": {
			builder:  SpecBuilder{},
			expected: coreV1.PodSpec{},
		},
		"spec with one container": {
			builder:  Spec(container.Name("simple-container")),
			expected: coreV1.PodSpec{Containers: []coreV1.Container{{Name: "simple-container"}}},
		},
		"spec with two containers": {
			builder:  Spec(container.Name("container-1"), container.Name("container-2")),
			expected: coreV1.PodSpec{Containers: []coreV1.Container{{Name: "container-1"}, {Name: "container-2"}}},
		},
		"spec with container and image": {
			builder:  Spec(container.Name("container-name").Image("image-name")),
			expected: coreV1.PodSpec{Containers: []coreV1.Container{{Name: "container-name", Image: "image-name"}}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			spec, err := test.builder.T()
			if err != nil {
				t.Fatalf("failed to create typed value: %s", err)
			}
			if !reflect.DeepEqual(spec, test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", spec, test.expected)
			}
		})
	}
}
