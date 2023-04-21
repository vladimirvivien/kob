package pod

import (
	"reflect"
	"testing"

	"github.com/vladimirvivien/kob/container"
	"github.com/vladimirvivien/kob/objmeta"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPodUnstructured(t *testing.T) {
	tests := map[string]struct {
		builder  Builder
		expected map[string]interface{}
	}{
		"empty spec": {
			builder:  Builder{},
			expected: map[string]interface{}{},
		},
		"pod with object meta only": {
			builder:  Object(objmeta.Name("simple-pod").Namespace("default")),
			expected: map[string]interface{}{"metadata": map[string]interface{}{"name": "simple-pod", "namespace": "default"}},
		},
		"pod one container": {
			builder:  Object(objmeta.Name("simple-pod")).Spec(container.Name("simple-container")),
			expected: map[string]interface{}{"metadata": map[string]interface{}{"name": "simple-pod"}, "spec": map[string]interface{}{"containers": []interface{}{map[string]interface{}{"name": "simple-container"}}}},
		},
		"pod with two containers": {
			builder:  Object(objmeta.Name("simple-pod")).Spec(container.Name("simple-container-1"), container.Name("simple-container-2")),
			expected: map[string]interface{}{"metadata": map[string]interface{}{"name": "simple-pod"}, "spec": map[string]interface{}{"containers": []interface{}{map[string]interface{}{"name": "simple-container-1"}, map[string]interface{}{"name": "simple-container-2"}}}},
		},
		"pod with container and image": {
			builder:  Object(objmeta.Name("simple-pod")).Spec(container.WithNameAndImage("simple-container", "simple-image")),
			expected: map[string]interface{}{"metadata": map[string]interface{}{"name": "simple-pod"}, "spec": map[string]interface{}{"containers": []interface{}{map[string]interface{}{"name": "simple-container", "image": "simple-image"}}}},
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

func TestPodStructured(t *testing.T) {
	tests := map[string]struct {
		builder  Builder
		expected coreV1.Pod
	}{
		"empty spec": {
			builder:  Builder{},
			expected: coreV1.Pod{},
		},
		"pod with no spec": {
			builder:  Object(objmeta.Name("simple-pod")),
			expected: coreV1.Pod{ObjectMeta: metaV1.ObjectMeta{Name: "simple-pod"}},
		},
		"pod with one container": {
			builder:  Object(objmeta.Name("simple-pod")).Spec(container.Name("simple-container")),
			expected: coreV1.Pod{ObjectMeta: metaV1.ObjectMeta{Name: "simple-pod"}, Spec: coreV1.PodSpec{Containers: []coreV1.Container{{Name: "simple-container"}}}},
		},
		"pod with two containers": {
			builder:  Object(objmeta.Name("simple-pod")).Spec(container.Name("simple-container-1"), container.Name("simple-container-2")),
			expected: coreV1.Pod{ObjectMeta: metaV1.ObjectMeta{Name: "simple-pod"}, Spec: coreV1.PodSpec{Containers: []coreV1.Container{{Name: "simple-container-1"}, {Name: "simple-container-2"}}}},
		},
		"pod with container and image": {
			builder:  Object(objmeta.Name("simple-pod")).Spec(container.Name("simple-container").Image("simple-image")),
			expected: coreV1.Pod{ObjectMeta: metaV1.ObjectMeta{Name: "simple-pod"}, Spec: coreV1.PodSpec{Containers: []coreV1.Container{{Name: "simple-container", Image: "simple-image"}}}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pod, err := test.builder.T()
			if err != nil {
				t.Fatalf("failed to convert to typed value: %s", err)
			}
			if !reflect.DeepEqual(pod, test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", pod, test.expected)
			}
		})
	}
}
