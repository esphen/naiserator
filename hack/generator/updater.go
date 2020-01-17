package main

import (
	"os"
	"text/template"
)

//go:generate go run updater.go

type Service struct {
	Name          string // Resource type.
	Interface     string // Kubernetes client interface for using REST.
	Type          string // Real type of the object being modified.
	TransformFunc string // Optional name of function that will be called to copy the old object into the new.
	Client        string // Function to call on the clientset to get a REST client
}

var i = []Service{
	{
		Name:          "service",
		Interface:     "typed_core_v1.ServiceInterface",
		Type:          "*corev1.Service",
		TransformFunc: "CopyService",
		Client:        "clientSet.CoreV1().Services",
	},
	{
		Name:      "secret",
		Interface: "typed_core_v1.SecretInterface",
		Type:      "*corev1.Secret",
		Client:    "clientSet.CoreV1().Secrets",
	},
	{
		Name:      "serviceAccount",
		Interface: "typed_core_v1.ServiceAccountInterface",
		Type:      "*corev1.ServiceAccount",
		Client:    "clientSet.CoreV1().ServiceAccounts",
	},
	{
		Name:      "deployment",
		Interface: "typed_apps_v1.DeploymentInterface",
		Type:      "*appsv1.Deployment",
		Client:    "clientSet.AppsV1().Deployments",
	},
	{
		Name:      "ingress",
		Interface: "typed_extensions_v1beta1.IngressInterface",
		Type:      "*extensionsv1beta1.Ingress",
		Client:    "clientSet.ExtensionsV1beta1().Ingresses",
	},
	{
		Name:      "horizontalPodAutoscaler",
		Interface: "typed_autoscaling_v1.HorizontalPodAutoscalerInterface",
		Type:      "*autoscalingv1.HorizontalPodAutoscaler",
		Client:    "clientSet.AutoscalingV1().HorizontalPodAutoscalers",
	},
	{
		Name:      "networkPolicy",
		Interface: "typed_networking_v1.NetworkPolicyInterface",
		Type:      "*networkingv1.NetworkPolicy",
		Client:    "clientSet.NetworkingV1().NetworkPolicies",
	},
	{
		Name:      "serviceRole",
		Interface: "istio_v1alpha1.ServiceRoleInterface",
		Type:      "*v1alpha1.ServiceRole",
		Client:    "customClient.RbacV1alpha1().ServiceRoles",
	},
	{
		Name:      "serviceRoleBinding",
		Interface: "istio_v1alpha1.ServiceRoleBindingInterface",
		Type:      "*v1alpha1.ServiceRoleBinding",
		Client:    "customClient.RbacV1alpha1().ServiceRoleBindings",
	},
	{
		Name:      "virtualService",
		Interface: "typed_networking_istio_io_v1alpha3.VirtualServiceInterface",
		Type:      "*networking_istio_io_v1alpha3.VirtualService",
		Client:    "customClient.NetworkingV1alpha3().VirtualServices",
	},
	{
		Name:      "ServiceEntry",
		Interface: "typed_networking_istio_io_v1alpha3.ServiceEntryInterface",
		Type:      "*networking_istio_io_v1alpha3.ServiceEntry",
		Client:    "customClient.NetworkingV1alpha3().ServiceEntries",
	},
	{
		Name:      "Role",
		Interface: "typed_rbac_v1.RoleInterface",
		Type:      "*rbacv1.Role",
		Client:    "clientSet.RbacV1().Roles",
	},
	{
		Name:      "RoleBinding",
		Interface: "typed_rbac_v1.RoleBindingInterface",
		Type:      "*rbacv1.RoleBinding",
		Client:    "clientSet.RbacV1().RoleBindings",
	},
	{
		Name:      "iamServiceAccount",
		Interface: "typed_iam_cnrm_cloud_google_com_v1beta1.IAMServiceAccountInterface",
		Type:      "*iam_cnrm_cloud_google_com_v1beta1.IAMServiceAccount",
		Client:    "customClient.IamV1beta1().IAMServiceAccounts",
	},
	{
		Name:      "iamPolicy",
		Interface: "typed_iam_cnrm_cloud_google_com_v1beta1.IAMPolicyInterface",
		Type:      "*iam_cnrm_cloud_google_com_v1beta1.IAMPolicy",
		Client:    "customClient.IamV1beta1().IAMPolicies",
	},
	{
		Name:      "iamPolicyMember",
		Interface: "typed_iam_cnrm_cloud_google_com_v1beta1.IAMPolicyMemberInterface",
		Type:      "*iam_cnrm_cloud_google_com_v1beta1.IAMPolicyMember",
		Client:    "customClient.IamV1beta1().IAMPolicyMembers",
	},
	{
		Name:      "googleStorageBucket",
		Interface: "typed_storage_cnrm_cloud_google_com_v1alpha2.StorageBucketInterface",
		Type:      "*storage_cnrm_cloud_google_com_v1alpha2.StorageBucket",
		Client:    "customClient.StorageV1alpha2().StorageBuckets",
	},
	{
		Name:      "googleStorageBucketAccessControl",
		Interface: "typed_storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControlInterface",
		Type:      "*storage_cnrm_cloud_google_com_v1alpha2.StorageBucketAccessControl",
		Client:    "customClient.StorageV1alpha2().StorageBucketAccessControls",
	},
	{
		Name:      "sqlInstance",
		Interface: "typed_sql_cnrm_cloud_google_com_v1alpha3.SQLInstanceInterface",
		Type:      "*sql_cnrm_cloud_google_com_v1alpha3.SQLInstance",
		Client:    "customClient.SqlV1alpha3().SQLInstances",
	},
	{
		Name:      "sqlDatabase",
		Interface: "typed_sql_cnrm_cloud_google_com_v1alpha3.SQLDatabaseInterface",
		Type:      "*sql_cnrm_cloud_google_com_v1alpha3.SQLDatabase",
		Client:    "customClient.SqlV1alpha3().SQLDatabases",
	},
	{
		Name:      "sqlUser",
		Interface: "typed_sql_cnrm_cloud_google_com_v1alpha3.SQLUserInterface",
		Type:      "*sql_cnrm_cloud_google_com_v1alpha3.SQLUser",
		Client:    "customClient.SqlV1alpha3().SQLUsers",
	},
}

func main() {
	t, err := template.ParseFiles("updater.go.tpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, i)
	if err != nil {
		panic(err)
	}
}
