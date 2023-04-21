package deployment

import (
	"reflect"
	"testing"

	"github.com/vladimirvivien/kob/container"
	"github.com/vladimirvivien/kob/objmeta"
	appsV1 "k8s.io/api/apps/v1"
	coreV1 "k8s.io/api/core/v1"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestDeploymentUnstructured(t *testing.T) {
	tests := map[string]struct {
		builder  Builder
		expected map[string]interface{}
	}{
		"empty spec": {
			builder:  Builder{},
			expected: map[string]interface{}{},
		},
		"deployment with object meta only": {
			builder:  Object(objmeta.Name("simple-dep").Namespace("default")),
			expected: map[string]interface{}{"metadata": map[string]interface{}{"name": "simple-dep", "namespace": "default"}},
		},
		"deployment with replicas": {
			builder: Object(objmeta.Name("simple-dep")).Replicas(3),
			expected: map[string]interface{}{
				"metadata": map[string]interface{}{"name": "simple-dep"},
				"spec": map[string]interface{}{
					"replicas": replicas(3),
				},
			},
		},
		"deployment with strategy": {
			builder: Object(objmeta.Name("simple-dep")).Replicas(3).Strategy(StrategyDefault),
			expected: map[string]interface{}{
				"metadata": map[string]interface{}{"name": "simple-dep"},
				"spec": map[string]interface{}{
					"replicas": replicas(3),
					"strategy": map[string]interface{}{
						"type": appsV1.RecreateDeploymentStrategyType,
					},
				},
			},
		},
		"deployment with podspec": {
			builder: Object(objmeta.Name("simple-dep")).Replicas(3).Strategy(StrategyDefault).PodSpec(container.Name("simple-container")),
			expected: map[string]interface{}{
				"metadata": map[string]interface{}{"name": "simple-dep"},
				"spec": map[string]interface{}{
					"replicas": replicas(3),
					"strategy": map[string]interface{}{
						"type": appsV1.RecreateDeploymentStrategyType,
					},
					"template": map[string]interface{}{
						"spec": map[string]interface{}{
							"containers": []interface{}{
								map[string]interface{}{"name": "simple-container"},
							},
						},
					},
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

func TestDeploymentTyped(t *testing.T) {
	tests := map[string]struct {
		builder  Builder
		expected appsV1.Deployment
	}{
		"empty": {
			builder:  Builder{},
			expected: appsV1.Deployment{},
		},
		"object meta only": {
			builder: Object(objmeta.Name("simple-dep").Namespace("default")),
			expected: appsV1.Deployment{
				ObjectMeta: metaV1.ObjectMeta{Name: "simple-dep", Namespace: "default"},
			},
		},
		"with replicas": {
			builder: Object(objmeta.Name("simple-dep")).Replicas(3),
			expected: appsV1.Deployment{
				ObjectMeta: metaV1.ObjectMeta{Name: "simple-dep"},
				Spec: appsV1.DeploymentSpec{
					Replicas: replicas(3),
				},
			},
		},
		"with strategy": {
			builder: Object(objmeta.Name("simple-dep")).Replicas(3).Strategy(StrategyDefault),
			expected: appsV1.Deployment{
				ObjectMeta: metaV1.ObjectMeta{Name: "simple-dep"},
				Spec: appsV1.DeploymentSpec{
					Replicas: replicas(3),
					Strategy: appsV1.DeploymentStrategy{
						Type: appsV1.RecreateDeploymentStrategyType,
					},
				},
			},
		},
		"with podspec no metadata": {
			builder: Object(objmeta.Name("simple-dep")).Replicas(3).Strategy(StrategyDefault).PodSpec(container.Name("simple-container")),
			expected: appsV1.Deployment{
				ObjectMeta: metaV1.ObjectMeta{Name: "simple-dep"},
				Spec: appsV1.DeploymentSpec{
					Replicas: replicas(3),
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
			},
		},
		"with podspec and pod metadata": {
			builder: Object(objmeta.Name("simple-dep")).Replicas(3).Strategy(StrategyDefault).PodSpecWithMetadata(objmeta.Name("dep-pods"), container.Name("simple-container")),
			expected: appsV1.Deployment{
				ObjectMeta: metaV1.ObjectMeta{Name: "simple-dep"},
				Spec: appsV1.DeploymentSpec{
					Replicas: replicas(3),
					Strategy: appsV1.DeploymentStrategy{
						Type: appsV1.RecreateDeploymentStrategyType,
					},
					Template: coreV1.PodTemplateSpec{
						ObjectMeta: metaV1.ObjectMeta{Name: "dep-pods"},
						Spec: coreV1.PodSpec{
							Containers: []coreV1.Container{
								{Name: "simple-container"},
							},
						},
					},
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			dep, err := test.builder.T()
			if err != nil {
				t.Fatalf("failed to convert to typed value: %s", err)
			}
			if !reflect.DeepEqual(dep, test.expected) {
				t.Errorf("object not equal \n\n Constructor: %#v \n\n Expected: %#v", dep, test.expected)
			}
		})
	}
}
