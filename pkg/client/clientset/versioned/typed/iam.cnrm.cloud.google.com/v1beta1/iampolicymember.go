// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/nais/naiserator/pkg/apis/iam.cnrm.cloud.google.com/v1beta1"
	scheme "github.com/nais/naiserator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IAMPolicyMembersGetter has a method to return a IAMPolicyMemberInterface.
// A group's client should implement this interface.
type IAMPolicyMembersGetter interface {
	IAMPolicyMembers(namespace string) IAMPolicyMemberInterface
}

// IAMPolicyMemberInterface has methods to work with IAMPolicyMember resources.
type IAMPolicyMemberInterface interface {
	Create(*v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error)
	Update(*v1beta1.IAMPolicyMember) (*v1beta1.IAMPolicyMember, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.IAMPolicyMember, error)
	List(opts v1.ListOptions) (*v1beta1.IAMPolicyMemberList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.IAMPolicyMember, err error)
	IAMPolicyMemberExpansion
}

// iAMPolicyMembers implements IAMPolicyMemberInterface
type iAMPolicyMembers struct {
	client rest.Interface
	ns     string
}

// newIAMPolicyMembers returns a IAMPolicyMembers
func newIAMPolicyMembers(c *IamV1beta1Client, namespace string) *iAMPolicyMembers {
	return &iAMPolicyMembers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iAMPolicyMember, and returns the corresponding iAMPolicyMember object, and an error if there is any.
func (c *iAMPolicyMembers) Get(name string, options v1.GetOptions) (result *v1beta1.IAMPolicyMember, err error) {
	result = &v1beta1.IAMPolicyMember{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iampolicymembers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IAMPolicyMembers that match those selectors.
func (c *iAMPolicyMembers) List(opts v1.ListOptions) (result *v1beta1.IAMPolicyMemberList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.IAMPolicyMemberList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("iampolicymembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iAMPolicyMembers.
func (c *iAMPolicyMembers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("iampolicymembers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a iAMPolicyMember and creates it.  Returns the server's representation of the iAMPolicyMember, and an error, if there is any.
func (c *iAMPolicyMembers) Create(iAMPolicyMember *v1beta1.IAMPolicyMember) (result *v1beta1.IAMPolicyMember, err error) {
	result = &v1beta1.IAMPolicyMember{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("iampolicymembers").
		Body(iAMPolicyMember).
		Do().
		Into(result)
	return
}

// Update takes the representation of a iAMPolicyMember and updates it. Returns the server's representation of the iAMPolicyMember, and an error, if there is any.
func (c *iAMPolicyMembers) Update(iAMPolicyMember *v1beta1.IAMPolicyMember) (result *v1beta1.IAMPolicyMember, err error) {
	result = &v1beta1.IAMPolicyMember{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("iampolicymembers").
		Name(iAMPolicyMember.Name).
		Body(iAMPolicyMember).
		Do().
		Into(result)
	return
}

// Delete takes name of the iAMPolicyMember and deletes it. Returns an error if one occurs.
func (c *iAMPolicyMembers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iampolicymembers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iAMPolicyMembers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("iampolicymembers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched iAMPolicyMember.
func (c *iAMPolicyMembers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.IAMPolicyMember, err error) {
	result = &v1beta1.IAMPolicyMember{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("iampolicymembers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}