package objmeta

import (
	"reflect"
	"testing"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestObjectMetaUnstructuredValues(t *testing.T) {
	tests := map[string]struct {
		builder  Builder
		expected map[string]interface{}
	}{
		"empty object": {
			builder:  Builder{},
			expected: map[string]interface{}{},
		},
		"name only": {
			builder:  Name("simple-name"),
			expected: map[string]interface{}{"name": "simple-name"},
		},
		"name and namespace": {
			builder:  Name("simple-name").Namespace("my-namespace"),
			expected: map[string]interface{}{"name": "simple-name", "namespace": "my-namespace"},
		},
		"name and namespace and labels": {
			builder:  Name("simple-name").Namespace("my-namespace").Labels(map[string]string{"tier": "web"}),
			expected: map[string]interface{}{"name": "simple-name", "namespace": "my-namespace", "labels": map[string]string{"tier": "web"}},
		},
		"all fields": {
			builder:  Name("simple-name").Namespace("my-namespace").Labels(map[string]string{"tier": "web"}).Annotations(map[string]string{"status": "ready"}),
			expected: map[string]interface{}{"name": "simple-name", "namespace": "my-namespace", "labels": map[string]string{"tier": "web"}, "annotations": map[string]string{"status": "ready"}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			objMeta := test.builder.U()
			if !reflect.DeepEqual(objMeta, test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", objMeta, test.expected)
			}
		})
	}
}

func TestObjectMetaTypedValues(t *testing.T) {
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
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			objMeta, err := test.builder.T()
			if err != nil {
				t.Fatalf("typed conversion failed: %s", err)
			}
			if !reflect.DeepEqual(objMeta, test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", objMeta, test.expected)
			}
		})
	}
}
