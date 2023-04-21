package container

import (
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Builder map[string]interface{}

// Name returns a container Builder
func Name(name string) Builder {
	return Builder{
		"name": name,
	}
}

func WithNameAndImage(name, image string) Builder {
	return Builder{
		"name":  name,
		"image": image,
	}
}

func (b Builder) U() map[string]interface{} {
	return b
}

func (b Builder) T() (coreV1.Container, error) {
	var container coreV1.Container
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(b, &container); err != nil {
		return coreV1.Container{}, err
	}
	return container, nil
}

func (b Builder) Image(img string) Builder {
	b["image"] = img
	return b
}

// Args sets the command arguments for the container
func (b Builder) Args(args ...string) Builder {
	b["args"] = args
	return b
}

// Commands sets the container's entry point command
func (b Builder) Commands(cmds ...string) Builder {
	b["command"] = cmds
	return b
}

func (b Builder) WorkingDir(dir string) Builder {
	b["workingDir"] = dir
	return b
}

// func (b ContainerBuilder) Ports(ports ...coreV1.ContainerPort) ContainerBuilder {
// 	b["ports"] = ports
// 	return b
// }

// func (b *ContainerBuilder) EnvSources(sources ...coreV1.EnvFromSource) *ContainerBuilder {
// 	b.container.EnvFrom = sources
// 	return b
// }

// func (b *ContainerBuilder) AddEnvSource(source coreV1.EnvFromSource) *ContainerBuilder {
// 	b.container.EnvFrom = append(b.container.EnvFrom, source)
// 	return b
// }

// func (b *ContainerBuilder) EnvVars(envs ...coreV1.EnvVar) *ContainerBuilder {
// 	b.container.Env = envs
// 	return b
// }

// func (b *ContainerBuilder) AddEnvVar(env coreV1.EnvVar) *ContainerBuilder {
// 	b.container.Env = append(b.container.Env, env)
// 	return b
// }

// func (b *ContainerBuilder) ResourceLimits(limits coreV1.ResourceList) *ContainerBuilder {
// 	b.container.Resources.Limits = limits
// 	return b
// }

// func (b *ContainerBuilder) AddResourceLimit(name coreV1.ResourceName, qty resource.Quantity) *ContainerBuilder {
// 	if b.container.Resources.Limits == nil {
// 		b.container.Resources.Limits = make(coreV1.ResourceList)
// 	}
// 	b.container.Resources.Limits[name] = qty
// 	return b
// }

// func (b *ContainerBuilder) ResourceRequests(reqs coreV1.ResourceList) *ContainerBuilder {
// 	b.container.Resources.Requests = reqs
// 	return b
// }

// func (b *ContainerBuilder) AddResourceRequest(name coreV1.ResourceName, qty resource.Quantity) *ContainerBuilder {
// 	if b.container.Resources.Requests == nil {
// 		b.container.Resources.Requests = make(coreV1.ResourceList)
// 	}
// 	b.container.Resources.Requests[name] = qty
// 	return b
// }

// func (b *ContainerBuilder) VolumeMounts(mounts ...coreV1.VolumeMount) *ContainerBuilder {
// 	b.container.VolumeMounts = mounts
// 	return b
// }

// func (b *ContainerBuilder) AddVolumeMount(mount coreV1.VolumeMount) *ContainerBuilder {
// 	b.container.VolumeMounts = append(b.container.VolumeMounts, mount)
// 	return b
// }

// func (b *ContainerBuilder) VolumeDevices(devices ...coreV1.VolumeDevice) *ContainerBuilder {
// 	b.container.VolumeDevices = devices
// 	return b
// }

// func (b *ContainerBuilder) AddVolumeDevice(device coreV1.VolumeDevice) *ContainerBuilder {
// 	b.container.VolumeDevices = append(b.container.VolumeDevices, device)
// 	return b
// }

// func (b *ContainerBuilder) ImagePullPolicy(policy coreV1.PullPolicy) *ContainerBuilder {
// 	b.container.ImagePullPolicy = policy
// 	return b
// }

// func (b *ContainerBuilder) SecurityContext(ctx *coreV1.SecurityContext) *ContainerBuilder {
// 	b.container.SecurityContext = ctx
// 	return b
// }

// // Do finalizes the build sequence and returns the *coreV1.Container
// func (b *ContainerBuilder) Do() coreV1.Container {
// 	return b.container
// }

// func Simple(name, image string) coreV1.Container {
// 	return coreV1.Container{Name: name, Image: image}
// }

// func BuildEnvsFromConfigMap(configMapName string) coreV1.EnvFromSource {
// 	return coreV1.EnvFromSource{
// 		ConfigMapRef: &coreV1.ConfigMapEnvSource{LocalObjectReference: coreV1.LocalObjectReference{Name: configMapName}},
// 	}
// }

// func BuildEnvsFromSecret(secretName string) coreV1.EnvFromSource {
// 	return coreV1.EnvFromSource{
// 		SecretRef: &coreV1.SecretEnvSource{LocalObjectReference: coreV1.LocalObjectReference{Name: secretName}},
// 	}
// }
