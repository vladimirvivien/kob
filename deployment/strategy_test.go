package deployment

import (
	"reflect"
	"testing"

	appsV1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

func TestStrategyUnstructured(t *testing.T) {
	tests := map[string]struct {
		builder  StrategyBuilder
		expected map[string]interface{}
	}{
		"empty": {
			builder:  StrategyBuilder{},
			expected: map[string]interface{}{},
		},
		"default strategy": {
			builder:  StrategyDefault,
			expected: map[string]interface{}{"type": appsV1.RecreateDeploymentStrategyType},
		},
		"recreate strategy": {
			builder:  StrategyRecreate,
			expected: map[string]interface{}{"type": appsV1.RecreateDeploymentStrategyType},
		},
		"rolling update strategy": {
			builder: RollingUpdate("0", "0"),
			expected: map[string]interface{}{
				"type": appsV1.RollingUpdateDeploymentStrategyType,
				"rollingUpdate": map[string]interface{}{
					"maxUnavailable": func() *intstr.IntOrString {
						max := intstr.FromString("0")
						return &max
					}(),
					"maxSurge": func() *intstr.IntOrString {
						max := intstr.FromString("0")
						return &max
					}(),
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			spec := test.builder
			if !reflect.DeepEqual(spec.U(), test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", spec.U(), test.expected)
			}
		})
	}
}

func TestStrategyStructured(t *testing.T) {
	tests := map[string]struct {
		builder  StrategyBuilder
		expected appsV1.DeploymentStrategy
	}{
		"empty": {
			builder:  StrategyBuilder{},
			expected: appsV1.DeploymentStrategy{},
		},
		"default strategy": {
			builder:  StrategyDefault,
			expected: appsV1.DeploymentStrategy{Type: appsV1.RecreateDeploymentStrategyType},
		},
		"rolling update strategy": {
			builder: RollingUpdate("0", "0"),
			expected: func() appsV1.DeploymentStrategy {
				unavailParsed := intstr.FromString("0")
				surgeParsed := intstr.FromString("0")
				return appsV1.DeploymentStrategy{
					Type: appsV1.RollingUpdateDeploymentStrategyType,
					RollingUpdate: &appsV1.RollingUpdateDeployment{
						MaxUnavailable: &unavailParsed,
						MaxSurge:       &surgeParsed,
					},
				}
			}(),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			pod, err := test.builder.T()
			if err != nil {
				t.Fatalf("failed to convert to typed value: %s", err)
			}
			if !reflect.DeepEqual(pod, test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", pod, test.expected)
			}
		})
	}
}
