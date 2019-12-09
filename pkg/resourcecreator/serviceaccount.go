package resourcecreator

import (
	"fmt"

	nais "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ServiceAccount(app *nais.Application, googleCluster bool, projectId string) *corev1.ServiceAccount {
	objectMeta := app.CreateObjectMeta()
	gcpSvcAcc := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", app.CreateAppNamespaceHash(), projectId)
	if googleCluster {
		objectMeta.Annotations["iam.gke.io/gcp-service-account"] = gcpSvcAcc
	}

	return &corev1.ServiceAccount{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ServiceAccount",
			APIVersion: "v1",
		},
		ObjectMeta: objectMeta,
	}
}
