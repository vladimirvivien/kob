// Package obj contains builder types to build values of type coreV1.ObjectMeta
package objmeta

import (
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	EmptyName        = ""
	EmptyLabels      = map[string]string{}
	DefaultNamespace = "default"
	ObjectMetaNone   = metaV1.ObjectMeta{}
)

// Builder provides a way to build values of type coreV1.ObjectMeta
type Builder map[string]interface{}

// Name sets the name of the object
func Name(name string) Builder {
	return Builder{"name": name}
}

func (b Builder) T() (metaV1.ObjectMeta, error) {
	var obj metaV1.ObjectMeta
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(b, &obj); err != nil {
		return ObjectMetaNone, err
	}
	return obj, nil
}

func (b Builder) U() map[string]interface{} {
	return b
}

// Namespace value setter
func (b Builder) Namespace(ns string) Builder {
	b["namespace"] = ns
	return b
}

// Labels value setter
func (b Builder) Labels(labels map[string]string) Builder {
	b["labels"] = labels
	return b
}

// Annotations setter
func (b Builder) Annotations(annotations map[string]string) Builder {
	b["annotations"] = annotations
	return b
}
