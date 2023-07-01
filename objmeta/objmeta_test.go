package objmeta

import (
	"reflect"
	"testing"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestObjectMeta(t *testing.T) {
	tests := map[string]struct {
		builder  Builder
		expected metaV1.ObjectMeta
	}{
		"empty object": {
			builder:  Builder{},
			expected: metaV1.ObjectMeta{},
		},
		"name only": {
			builder:  Name("simple-name"),
			expected: metaV1.ObjectMeta{Name: "simple-name"},
		},
		"name and namespace": {
			builder:  Name("simple-name").Namespace("my-namespace"),
			expected: metaV1.ObjectMeta{Name: "simple-name", Namespace: "my-namespace"},
		},
		"name and namespace and labels": {
			builder:  Name("simple-name").Namespace("my-namespace").Labels(map[string]string{"tier": "web"}),
			expected: metaV1.ObjectMeta{Name: "simple-name", Namespace: "my-namespace", Labels: map[string]string{"tier": "web"}},
		},
		"all fields": {
			builder:  Name("simple-name").Namespace("my-namespace").Labels(map[string]string{"tier": "web"}).Annotations(map[string]string{"status": "ready"}),
			expected: metaV1.ObjectMeta{Name: "simple-name", Namespace: "my-namespace", Labels: map[string]string{"tier": "web"}, Annotations: map[string]string{"status": "ready"}},
		},
		"from unstructured": {
			builder: func() Builder {
				b, _ := FromUnstructured(map[string]any{"name": "simple-name", "namespace": "my-namespace"})
				return b
			}(),
			expected: metaV1.ObjectMeta{Name: "simple-name", Namespace: "my-namespace"},
		},
		"from string": {
			builder: func() Builder {
				b, _ := FromString(`{"name":"simple-name", "namespace":"my-namespace"}`)
				return b
			}(),
			expected: metaV1.ObjectMeta{Name: "simple-name", Namespace: "my-namespace"},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			objMeta := test.builder.T()

			if !reflect.DeepEqual(objMeta, test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", objMeta, test.expected)
			}
		})
	}
}
