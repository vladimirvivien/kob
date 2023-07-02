package container

import (
	"strings"

	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/yaml"
)

type Builder struct {
	obj coreV1.Container
}

// From creates a new builder using the provided object
func From(obj coreV1.Container) Builder {
	return Builder{obj: obj}
}

// FromUnstructured creates a builder from an unstructured value
func FromUnstructured(unstruct map[string]any) (Builder, error) {
	var obj coreV1.Container
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(unstruct, &obj); err != nil {
		return Builder{}, err
	}
	return Builder{obj: obj}, nil
}

// FromString creates a builder from a valid YAML or JSON fragment
func FromString(str string) (Builder, error) {
	var obj coreV1.Container
	if err := yaml.NewYAMLOrJSONDecoder(strings.NewReader(str), 1024).Decode(&obj); err != nil {
		return Builder{}, err
	}
	return Builder{obj: obj}, nil
}

// Name creates a new builder starting with container's name
func Name(name string) Builder {
	return Builder{obj: coreV1.Container{Name: name}}
}

// WithNameAndImage creates a new builder with container name and image
func WithNameAndImage(name, image string) Builder {
	return Builder{obj: coreV1.Container{Name: name, Image: image}}
}

// U returns an unstructured value of builder's object
func (b Builder) U() (map[string]any, error) {
	unstruct, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&b.obj)
	if err != nil {
		return nil, err
	}
	return unstruct, nil
}

// T returns a typed value of builder's object
func (b Builder) T() coreV1.Container {
	return b.obj
}

func (b Builder) Image(img string) Builder {
	b.obj.Image = img
	return b
}

// Args sets the command arguments for the container
func (b Builder) Args(args ...string) Builder {
	b.obj.Args = args
	return b
}

// Commands sets the container's entry point command
func (b Builder) Commands(cmds ...string) Builder {
	b.obj.Command = cmds
	return b
}

// WorkingDir sets container WorkingDir value
func (b Builder) WorkingDir(dir string) Builder {
	b.obj.WorkingDir = dir
	return b
}

// Ports sets container port values
func (b Builder) Ports(ports ...coreV1.ContainerPort) Builder {
	b.obj.Ports = ports
	return b
}

// EnvFromSources sets environment value from provided sources
func (b Builder) EnvFromSources(sources ...coreV1.EnvFromSource) Builder {
	b.obj.EnvFrom = sources
	return b
}

// AddEnvFromConfigMapSource adds environment values from specified config map name
func (b Builder) AddEnvFromConfigMapSource(name string) Builder {
	source := coreV1.EnvFromSource{ConfigMapRef: &coreV1.ConfigMapEnvSource{LocalObjectReference: coreV1.LocalObjectReference{Name: name}}}
	b.obj.EnvFrom = append(b.obj.EnvFrom, source)
	return b
}

// AddEnvFromSecretSource adds secret environment values from specified secret
func (b Builder) AddEnvFromSecretSource(name string) Builder {
	source := coreV1.EnvFromSource{SecretRef: &coreV1.SecretEnvSource{LocalObjectReference: coreV1.LocalObjectReference{Name: name}}}
	b.obj.EnvFrom = append(b.obj.EnvFrom, source)
	return b
}

// EnvVars sets environment variable name/value pair
func (b Builder) EnvVars(vars ...coreV1.EnvVar) Builder {
	b.obj.Env = vars
	return b
}

// AddEnv adds a name/value pair environment variable for container
func (b Builder) AddEnv(name, value string) Builder {
	b.obj.Env = append(b.obj.Env, coreV1.EnvVar{Name: name, Value: value})
	return b
}

// ResourceLimits sets container's resource limits
func (b Builder) ResourceLimits(limits coreV1.ResourceList) Builder {
	b.obj.Resources.Limits = limits
	return b
}

func (b Builder) AddResourceLimit(name coreV1.ResourceName, qty resource.Quantity) Builder {
	if b.obj.Resources.Limits == nil {
		b.obj.Resources.Limits = make(coreV1.ResourceList)
	}
	b.obj.Resources.Limits[name] = qty
	return b
}

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
