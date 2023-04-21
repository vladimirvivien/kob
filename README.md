# kob 🦌 
Kubernetes Object Builder

Project `kob` uses a builder pattern to simplify the construction and composition of Kubernetes API object graphs. It provides helper constructors to reduce typing and insert sensible defaults where needed.

### Example

The following:

```go
deployment.Object(objmeta.Name("simple-dep")).Replicas(3).Strategy(StrategyDefault).PodSpec(container.Name("simple-container"))
```
Produces an Deployment object equivalent to the following:

```go
reps := int32(3)
appsV1.Deployment{
	ObjectMeta: metaV1.ObjectMeta{Name: "simple-dep"},
	Spec: appsV1.DeploymentSpec{
		Replicas: &reps,
		Strategy: appsV1.DeploymentStrategy{
			Type: appsV1.RecreateDeploymentStrategyType,
		},
		Template: coreV1.PodTemplateSpec{
			Spec: coreV1.PodSpec{
				Containers: []coreV1.Container{
					{Name: "simple-container"},
				},
			},
		},
	},
}
```
