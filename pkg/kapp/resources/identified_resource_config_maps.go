package resources

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
)

func (r IdentifiedResources) ConfigMapResources(labelSelector labels.Selector) ([]corev1.ConfigMap, error) {
	listOpts := metav1.ListOptions{LabelSelector: labelSelector.String()}

	mapList, err := r.coreClient.CoreV1().ConfigMaps("").List(listOpts)
	if err != nil {
		return nil, err
	}

	var maps []corev1.ConfigMap

	for _, m := range mapList.Items {
		maps = append(maps, m)
	}

	return maps, nil
}
