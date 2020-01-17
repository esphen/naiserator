// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "github.com/nais/naiserator/pkg/client/clientset/versioned"
	iamv1beta1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/iam.cnrm.cloud.google.com/v1beta1"
	fakeiamv1beta1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/iam.cnrm.cloud.google.com/v1beta1/fake"
	naiseratorv1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/nais.io/v1alpha1"
	fakenaiseratorv1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/nais.io/v1alpha1/fake"
	networkingv1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/networking.istio.io/v1alpha3"
	fakenetworkingv1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/networking.istio.io/v1alpha3/fake"
	rbacv1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/rbac.istio.io/v1alpha1"
	fakerbacv1alpha1 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/rbac.istio.io/v1alpha1/fake"
	sqlv1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/sql.cnrm.cloud.google.com/v1alpha3"
	fakesqlv1alpha3 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/sql.cnrm.cloud.google.com/v1alpha3/fake"
	storagev1alpha2 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/storage.cnrm.cloud.google.com/v1alpha2"
	fakestoragev1alpha2 "github.com/nais/naiserator/pkg/client/clientset/versioned/typed/storage.cnrm.cloud.google.com/v1alpha2/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// IamV1beta1 retrieves the IamV1beta1Client
func (c *Clientset) IamV1beta1() iamv1beta1.IamV1beta1Interface {
	return &fakeiamv1beta1.FakeIamV1beta1{Fake: &c.Fake}
}

// NaiseratorV1alpha1 retrieves the NaiseratorV1alpha1Client
func (c *Clientset) NaiseratorV1alpha1() naiseratorv1alpha1.NaiseratorV1alpha1Interface {
	return &fakenaiseratorv1alpha1.FakeNaiseratorV1alpha1{Fake: &c.Fake}
}

// NetworkingV1alpha3 retrieves the NetworkingV1alpha3Client
func (c *Clientset) NetworkingV1alpha3() networkingv1alpha3.NetworkingV1alpha3Interface {
	return &fakenetworkingv1alpha3.FakeNetworkingV1alpha3{Fake: &c.Fake}
}

// RbacV1alpha1 retrieves the RbacV1alpha1Client
func (c *Clientset) RbacV1alpha1() rbacv1alpha1.RbacV1alpha1Interface {
	return &fakerbacv1alpha1.FakeRbacV1alpha1{Fake: &c.Fake}
}

// SqlV1alpha3 retrieves the SqlV1alpha3Client
func (c *Clientset) SqlV1alpha3() sqlv1alpha3.SqlV1alpha3Interface {
	return &fakesqlv1alpha3.FakeSqlV1alpha3{Fake: &c.Fake}
}

// StorageV1alpha2 retrieves the StorageV1alpha2Client
func (c *Clientset) StorageV1alpha2() storagev1alpha2.StorageV1alpha2Interface {
	return &fakestoragev1alpha2.FakeStorageV1alpha2{Fake: &c.Fake}
}
