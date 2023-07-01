// Package obj contains builder types to build values of type coreV1.ObjectMeta
package objmeta

import (
	"strings"

	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
)

var (
	EmptyName        = ""
	EmptyLabels      = map[string]string{}
	DefaultNamespace = "default"
	ObjectMetaNone   = metaV1.ObjectMeta{}
)

// Builder provides a way to build values of type coreV1.ObjectMeta
type Builder struct {
	obj metaV1.ObjectMeta
}

// From creates a new builder using the provided metaV1.ObjectMeta as its base
func From(obj metaV1.ObjectMeta) Builder {
	return Builder{obj: obj}
}

// FromUnstructured attempts to convert an unstructured value into metaV1.ObjectMeta
// and uses it as the basis for a Builder
func FromUnstructured(unstruct map[string]any) (Builder, error) {
	var obj metaV1.ObjectMeta
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstruct, &obj); err != nil {
		return Builder{}, err
	}
	return Builder{obj: obj}, nil
}

// FromString attempts to convert the provided YAML or JSON string fragment
// into a valid metaV1.ObjectMeta and uses it as the basis for the builder
func FromString(str string) (Builder, error) {
	var obj metaV1.ObjectMeta
	if err := yaml.NewYAMLOrJSONDecoder(strings.NewReader(str), 1024).Decode(&obj); err != nil {
		return Builder{}, err
	}
	return Builder{obj: obj}, nil
}

// Name starts a new builer by setting the Objectmeta name of the object
func Name(name string) Builder {
	return Builder{obj: metaV1.ObjectMeta{Name: name}}
}

// U converts the value of the builder to an unstructured map[string]any value
func (b Builder) U() (map[string]any, error) {
	unstruct, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&b.obj)
	if err != nil {
		return nil, err
	}
	return unstruct, nil
}

// T returns the typed value of the builder of metaV1.ObjectMeta
func (b Builder) T() metaV1.ObjectMeta {
	return b.obj
}

// Namespace value setter
func (b Builder) Namespace(ns string) Builder {
	b.obj.Namespace = ns
	return b
}

// Labels value setter
func (b Builder) Labels(labels map[string]string) Builder {
	b.obj.Labels = labels
	return b
}

// Annotations setter
func (b Builder) Annotations(annotations map[string]string) Builder {
	b.obj.Annotations = annotations
	return b
}
