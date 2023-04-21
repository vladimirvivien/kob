package pod

import (
	"github.com/vladimirvivien/kob/container"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type SpecBuilder map[string]interface{}

func Spec(containers ...container.Builder) SpecBuilder {
	var slice []interface{}
	for _, c := range containers {
		slice = append(slice, c.U())
	}
	return SpecBuilder{
		"containers": slice,
	}
}

func (b SpecBuilder) U() map[string]interface{} {
	return map[string]interface{}(b)
}

func (b SpecBuilder) T() (coreV1.PodSpec, error) {
	var spec coreV1.PodSpec
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(b, &spec); err != nil {
		return coreV1.PodSpec{}, err
	}
	return spec, nil
}

func (b SpecBuilder) InitContainers(containers ...container.Builder) SpecBuilder {
	var slice []interface{}
	for _, c := range containers {
		slice = append(slice, c.U())
	}
	return map[string]interface{}{
		"initContainers": slice,
	}
}

// func (b *PodSpecBuilder) Volumes(vols ...coreV1.Volume) *PodSpecBuilder {
// 	b.spec.Volumes = vols
// 	return b
// }

// func (b *PodSpecBuilder) AddVolume(vol coreV1.Volume) *PodSpecBuilder {
// 	b.spec.Volumes = append(b.spec.Volumes, vol)
// 	return b
// }

// func (b *PodSpecBuilder) InitContainers(containers ...coreV1.Container) *PodSpecBuilder {
// 	b.spec.InitContainers = containers
// 	return b
// }

// func (b *PodSpecBuilder) AddInitContainer(container coreV1.Container) *PodSpecBuilder {
// 	b.spec.InitContainers = append(b.spec.InitContainers, container)
// 	return b
// }

// func (b *PodSpecBuilder) Containers(containers ...coreV1.Container) *PodSpecBuilder {
// 	b.spec.Containers = containers
// 	return b
// }

// func (b *PodSpecBuilder) AddContainer(container coreV1.Container) *PodSpecBuilder {
// 	b.spec.Containers = append(b.spec.Containers, container)
// 	return b
// }

// func (b *PodSpecBuilder) RestartPolicy(pol coreV1.RestartPolicy) *PodSpecBuilder {
// 	b.spec.RestartPolicy = pol
// 	return b
// }

// func (b *PodSpecBuilder) DNSPolicy(pol coreV1.DNSPolicy) *PodSpecBuilder {
// 	b.spec.DNSPolicy = pol
// 	return b
// }

// // Build is the finalizer method that returns a value of type coreV1.PodSpec
// func (c *PodSpecBuilder) Do() coreV1.PodSpec {
// 	return c.spec
// }
