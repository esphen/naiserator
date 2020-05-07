// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	naisiov1 "github.com/nais/naiserator/pkg/apis/nais.io/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeJwkers implements JwkerInterface
type FakeJwkers struct {
	Fake *FakeNaiseratorV1
	ns   string
}

var jwkersResource = schema.GroupVersionResource{Group: "naiserator.nais.io", Version: "v1", Resource: "jwkers"}

var jwkersKind = schema.GroupVersionKind{Group: "naiserator.nais.io", Version: "v1", Kind: "Jwker"}

// Get takes name of the jwker, and returns the corresponding jwker object, and an error if there is any.
func (c *FakeJwkers) Get(name string, options v1.GetOptions) (result *naisiov1.Jwker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(jwkersResource, c.ns, name), &naisiov1.Jwker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*naisiov1.Jwker), err
}

// List takes label and field selectors, and returns the list of Jwkers that match those selectors.
func (c *FakeJwkers) List(opts v1.ListOptions) (result *naisiov1.JwkerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(jwkersResource, jwkersKind, c.ns, opts), &naisiov1.JwkerList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &naisiov1.JwkerList{ListMeta: obj.(*naisiov1.JwkerList).ListMeta}
	for _, item := range obj.(*naisiov1.JwkerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested jwkers.
func (c *FakeJwkers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(jwkersResource, c.ns, opts))

}

// Create takes the representation of a jwker and creates it.  Returns the server's representation of the jwker, and an error, if there is any.
func (c *FakeJwkers) Create(jwker *naisiov1.Jwker) (result *naisiov1.Jwker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(jwkersResource, c.ns, jwker), &naisiov1.Jwker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*naisiov1.Jwker), err
}

// Update takes the representation of a jwker and updates it. Returns the server's representation of the jwker, and an error, if there is any.
func (c *FakeJwkers) Update(jwker *naisiov1.Jwker) (result *naisiov1.Jwker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(jwkersResource, c.ns, jwker), &naisiov1.Jwker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*naisiov1.Jwker), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeJwkers) UpdateStatus(jwker *naisiov1.Jwker) (*naisiov1.Jwker, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(jwkersResource, "status", c.ns, jwker), &naisiov1.Jwker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*naisiov1.Jwker), err
}

// Delete takes name of the jwker and deletes it. Returns an error if one occurs.
func (c *FakeJwkers) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(jwkersResource, c.ns, name), &naisiov1.Jwker{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeJwkers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(jwkersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &naisiov1.JwkerList{})
	return err
}

// Patch applies the patch and returns the patched jwker.
func (c *FakeJwkers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *naisiov1.Jwker, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(jwkersResource, c.ns, name, pt, data, subresources...), &naisiov1.Jwker{})

	if obj == nil {
		return nil, err
	}
	return obj.(*naisiov1.Jwker), err
}
