package util

import "k8s.io/apimachinery/pkg/apis/meta/v1"

func SetAnnotation(resource v1.ObjectMetaAccessor, key, value string) {
	m := resource.GetObjectMeta().GetAnnotations()
	m[key] = value
	resource.GetObjectMeta().SetAnnotations(m)
}
