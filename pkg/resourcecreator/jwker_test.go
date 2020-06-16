package resourcecreator_test

import (
	"testing"

	nais "github.com/nais/naiserator/pkg/apis/nais.io/v1alpha1"
	"github.com/nais/naiserator/pkg/resourcecreator"
	"github.com/nais/naiserator/pkg/test/fixtures"
	"github.com/stretchr/testify/assert"
)

func TestJwker(t *testing.T) {
	otherApplication := "a"
	clusterName := "myCluster"
	otherNamespace := "othernamespace"
	otherCluster := "otherCluster"
	otherApplication2 := "b"
	otherNamespace2 := "othernamespace2"
	otherApplication3 := "c"

	t.Run("no jwker for app with no access policy", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		jwker := resourcecreator.Jwker(app, clusterName)
		assert.Empty(t, jwker)
	})

	t.Run("one inbound without cluster/namespace and no outbound", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.AccessPolicy.Inbound.Rules = []nais.AccessPolicyRule{{otherApplication, "", ""}}
		jwker := resourcecreator.Jwker(app, clusterName)
		assert.Len(t, jwker.Spec.AccessPolicy.Inbound.Rules, 1)
		assert.NotEmpty(t, jwker.Spec.SecretName)
		assert.Equal(t, otherApplication, jwker.Spec.AccessPolicy.Inbound.Rules[0].Application)
		assert.Equal(t, fixtures.ApplicationNamespace, jwker.Spec.AccessPolicy.Inbound.Rules[0].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Inbound.Rules[0].Cluster)
		assert.Len(t, jwker.Spec.AccessPolicy.Outbound.Rules, 0)
	})

	t.Run("one inbound with cluster/namespace and no outbound", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.AccessPolicy.Inbound.Rules = []nais.AccessPolicyRule{{otherApplication, otherNamespace, otherCluster}}
		jwker := resourcecreator.Jwker(app, clusterName)
		assert.Len(t, jwker.Spec.AccessPolicy.Inbound.Rules, 1)
		assert.NotEmpty(t, jwker.Spec.SecretName)
		assert.Equal(t, otherApplication, jwker.Spec.AccessPolicy.Inbound.Rules[0].Application)
		assert.Equal(t, otherNamespace, jwker.Spec.AccessPolicy.Inbound.Rules[0].Namespace)
		assert.Equal(t, otherCluster, jwker.Spec.AccessPolicy.Inbound.Rules[0].Cluster)
		assert.Len(t, jwker.Spec.AccessPolicy.Outbound.Rules, 0)
	})

	t.Run("one outbound and no inbound", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.AccessPolicy.Outbound.Rules = []nais.AccessPolicyRule{{otherApplication, otherNamespace, otherCluster}}
		jwker := resourcecreator.Jwker(app, clusterName)
		assert.Len(t, jwker.Spec.AccessPolicy.Outbound.Rules, 1)
		assert.NotEmpty(t, jwker.Spec.SecretName)
		assert.Equal(t, otherApplication, jwker.Spec.AccessPolicy.Outbound.Rules[0].Application)
		assert.Equal(t, otherNamespace, jwker.Spec.AccessPolicy.Outbound.Rules[0].Namespace)
		assert.Equal(t, otherCluster, jwker.Spec.AccessPolicy.Outbound.Rules[0].Cluster)
		assert.Len(t, jwker.Spec.AccessPolicy.Inbound.Rules, 0)
	})

	t.Run("multiple inbound and no outbound", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.AccessPolicy.Inbound.Rules = []nais.AccessPolicyRule{
			{
				otherApplication, otherNamespace, otherCluster,
			},
			{
				otherApplication2, otherNamespace2, "",
			},
			{
				otherApplication3, "", "",
			},
		}
		jwker := resourcecreator.Jwker(app, clusterName)
		assert.Len(t, jwker.Spec.AccessPolicy.Inbound.Rules, 3)
		assert.Len(t, jwker.Spec.AccessPolicy.Outbound.Rules, 0)
		assert.NotEmpty(t, jwker.Spec.SecretName)
		assert.Equal(t, otherApplication, jwker.Spec.AccessPolicy.Inbound.Rules[0].Application)
		assert.Equal(t, otherNamespace, jwker.Spec.AccessPolicy.Inbound.Rules[0].Namespace)
		assert.Equal(t, otherCluster, jwker.Spec.AccessPolicy.Inbound.Rules[0].Cluster)
		assert.Equal(t, otherApplication2, jwker.Spec.AccessPolicy.Inbound.Rules[1].Application)
		assert.Equal(t, otherNamespace2, jwker.Spec.AccessPolicy.Inbound.Rules[1].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Inbound.Rules[1].Cluster)
		assert.Equal(t, otherApplication3, jwker.Spec.AccessPolicy.Inbound.Rules[2].Application)
		assert.Equal(t, fixtures.ApplicationNamespace, jwker.Spec.AccessPolicy.Inbound.Rules[2].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Inbound.Rules[2].Cluster)
	})

	t.Run("multiple outbound and no inbound", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.AccessPolicy.Outbound.Rules = []nais.AccessPolicyRule{
			{
				otherApplication, otherNamespace, otherCluster,
			},
			{
				otherApplication2, otherNamespace2, "",
			},
			{
				otherApplication3, "", "",
			},
		}
		jwker := resourcecreator.Jwker(app, clusterName)
		assert.Len(t, jwker.Spec.AccessPolicy.Outbound.Rules, 3)
		assert.Len(t, jwker.Spec.AccessPolicy.Inbound.Rules, 0)
		assert.NotEmpty(t, jwker.Spec.SecretName)
		assert.Equal(t, otherApplication, jwker.Spec.AccessPolicy.Outbound.Rules[0].Application)
		assert.Equal(t, otherNamespace, jwker.Spec.AccessPolicy.Outbound.Rules[0].Namespace)
		assert.Equal(t, otherCluster, jwker.Spec.AccessPolicy.Outbound.Rules[0].Cluster)
		assert.Equal(t, otherApplication2, jwker.Spec.AccessPolicy.Outbound.Rules[1].Application)
		assert.Equal(t, otherNamespace2, jwker.Spec.AccessPolicy.Outbound.Rules[1].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Outbound.Rules[1].Cluster)
		assert.Equal(t, otherApplication3, jwker.Spec.AccessPolicy.Outbound.Rules[2].Application)
		assert.Equal(t, fixtures.ApplicationNamespace, jwker.Spec.AccessPolicy.Outbound.Rules[2].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Outbound.Rules[2].Cluster)
	})
	//
	t.Run("multiple inbound and multiple outbound", func(t *testing.T) {
		app := fixtures.MinimalApplication()
		app.Spec.AccessPolicy.Inbound.Rules = []nais.AccessPolicyRule{
			{
				otherApplication, otherNamespace, otherCluster,
			},
			{
				otherApplication2, otherNamespace2, "",
			},
			{
				otherApplication3, "", "",
			},
		}

		app.Spec.AccessPolicy.Outbound.Rules = []nais.AccessPolicyRule{
			{
				otherApplication, otherNamespace, otherCluster,
			},
			{
				otherApplication2, otherNamespace2, "",
			},
			{
				otherApplication3, "", "",
			},
		}
		jwker := resourcecreator.Jwker(app, clusterName)
		assert.Len(t, jwker.Spec.AccessPolicy.Inbound.Rules, 3)
		assert.Len(t, jwker.Spec.AccessPolicy.Outbound.Rules, 3)
		assert.NotEmpty(t, jwker.Spec.SecretName)
		assert.Equal(t, otherApplication, jwker.Spec.AccessPolicy.Inbound.Rules[0].Application)
		assert.Equal(t, otherNamespace, jwker.Spec.AccessPolicy.Inbound.Rules[0].Namespace)
		assert.Equal(t, otherCluster, jwker.Spec.AccessPolicy.Inbound.Rules[0].Cluster)
		assert.Equal(t, otherApplication2, jwker.Spec.AccessPolicy.Inbound.Rules[1].Application)
		assert.Equal(t, otherNamespace2, jwker.Spec.AccessPolicy.Inbound.Rules[1].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Inbound.Rules[1].Cluster)
		assert.Equal(t, otherApplication3, jwker.Spec.AccessPolicy.Inbound.Rules[2].Application)
		assert.Equal(t, fixtures.ApplicationNamespace, jwker.Spec.AccessPolicy.Inbound.Rules[2].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Inbound.Rules[2].Cluster)
		assert.Equal(t, otherApplication, jwker.Spec.AccessPolicy.Outbound.Rules[0].Application)
		assert.Equal(t, otherNamespace, jwker.Spec.AccessPolicy.Outbound.Rules[0].Namespace)
		assert.Equal(t, otherCluster, jwker.Spec.AccessPolicy.Outbound.Rules[0].Cluster)
		assert.Equal(t, otherApplication2, jwker.Spec.AccessPolicy.Outbound.Rules[1].Application)
		assert.Equal(t, otherNamespace2, jwker.Spec.AccessPolicy.Outbound.Rules[1].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Outbound.Rules[1].Cluster)
		assert.Equal(t, otherApplication3, jwker.Spec.AccessPolicy.Outbound.Rules[2].Application)
		assert.Equal(t, fixtures.ApplicationNamespace, jwker.Spec.AccessPolicy.Outbound.Rules[2].Namespace)
		assert.Equal(t, clusterName, jwker.Spec.AccessPolicy.Outbound.Rules[2].Cluster)
	})
}