package container

import (
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type PortBuilder struct {
	obj coreV1.ContainerPort
}

func HostPort(p int32) PortBuilder {
	return PortBuilder{obj: coreV1.ContainerPort{HostPort: p}}
}

func ContainerPort(p int32) PortBuilder {
	return PortBuilder{obj: coreV1.ContainerPort{ContainerPort: p}}
}

func (b PortBuilder) Name(name string) PortBuilder {
	b.obj.Name = name
	return b
}

func (b PortBuilder) Protocol(proto string) PortBuilder {
	b.obj.Protocol = coreV1.Protocol(proto)
	return b
}

func (b PortBuilder) HostPort(p int32) PortBuilder {
	b.obj.HostPort = p
	return b
}

func (b PortBuilder) ContainerPort(p int32) PortBuilder {
	b.obj.ContainerPort = p
	return b
}

func (b PortBuilder) HostIP(ip string) PortBuilder {
	b.obj.HostIP = ip
	return b
}

// U returns an unstructured value of builder's object
func (b PortBuilder) U() (map[string]any, error) {
	unstruct, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&b.obj)
	if err != nil {
		return nil, err
	}
	return unstruct, nil
}

// T returns a typed value of builder's object
func (b PortBuilder) T() coreV1.ContainerPort {
	return b.obj
}
