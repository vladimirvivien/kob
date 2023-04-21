package deployment

import (
	"github.com/vladimirvivien/kob/container"
	"github.com/vladimirvivien/kob/objmeta"
	"github.com/vladimirvivien/kob/pod"
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Builder map[string]interface{}

func Object(metadata objmeta.Builder) Builder {
	return Builder{"metadata": metadata.U()}
}

func (b Builder) U() map[string]interface{} {
	return b
}

func (b Builder) T() (appsV1.Deployment, error) {
	var dep appsV1.Deployment
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(b, &dep); err != nil {
		return appsV1.Deployment{}, err
	}
	return dep, nil
}

func (b Builder) Replicas(r int) Builder {
	spec := b.getDeploymentSpec()
	spec["replicas"] = replicas(r)
	b["spec"] = spec
	return b
}

func (b Builder) Strategy(strat StrategyBuilder) Builder {
	spec := b.getDeploymentSpec()
	spec["strategy"] = strat.U()
	b["spec"] = spec
	return b
}

func (b Builder) PodSpec(containers ...container.Builder) Builder {
	spec := b.getDeploymentSpec()
	spec["template"] = map[string]interface{}{
		"spec": pod.Spec(containers...).U(),
	}
	b["spec"] = spec
	return b
}

func (b Builder) PodSpecWithMetadata(metadata objmeta.Builder, containers ...container.Builder) Builder {
	spec := b.getDeploymentSpec()
	spec["template"] = map[string]interface{}{
		"metadata": metadata.U(),
		"spec":     pod.Spec(containers...).U(),
	}
	b["spec"] = spec
	return b
}

func replicas(r int) *int32 {
	rep := int32(r)
	return &rep
}

func (b Builder) getDeploymentSpec() map[string]interface{} {
	specIface, ok := b["spec"]
	if !ok {
		specIface = map[string]interface{}{}
	}
	return specIface.(map[string]interface{})
}
