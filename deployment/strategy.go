package deployment

import (
	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/intstr"
)

var (
	StrategyDefault  = StrategyBuilder{"type": appsV1.RecreateDeploymentStrategyType}
	StrategyRecreate = StrategyDefault
)

type StrategyBuilder map[string]interface{}

func RollingUpdate(maxUnavailable, maxSurge string) StrategyBuilder {
	unavailParsed := intstr.FromString(maxUnavailable)
	surgeParsed := intstr.FromString(maxSurge)
	return StrategyBuilder{
		"type": appsV1.RollingUpdateDeploymentStrategyType,
		"rollingUpdate": map[string]interface{}{
			"maxUnavailable": &unavailParsed,
			"maxSurge":       &surgeParsed,
		},
	}
}

func (b StrategyBuilder) U() map[string]interface{} {
	return b
}

func (b StrategyBuilder) T() (appsV1.DeploymentStrategy, error) {
	var strat appsV1.DeploymentStrategy
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(b, &strat); err != nil {
		return appsV1.DeploymentStrategy{}, err
	}
	return strat, nil
}
