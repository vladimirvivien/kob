package pod

import (
	"github.com/vladimirvivien/kob/container"
	"github.com/vladimirvivien/kob/objmeta"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type Builder map[string]interface{}

func Object(metadata objmeta.Builder) Builder {
	return Builder{"metadata": metadata.U()}
}

func (b Builder) U() map[string]interface{} {
	return b
}

func (b Builder) T() (coreV1.Pod, error) {
	var pod coreV1.Pod
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(b, &pod); err != nil {
		return coreV1.Pod{}, err
	}
	return pod, nil
}

func (b Builder) Spec(containers ...container.Builder) Builder {
	b["spec"] = Spec(containers...).U()
	return b
}
