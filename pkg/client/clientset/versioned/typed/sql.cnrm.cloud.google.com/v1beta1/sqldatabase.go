// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"time"

	v1beta1 "github.com/nais/naiserator/pkg/apis/sql.cnrm.cloud.google.com/v1beta1"
	scheme "github.com/nais/naiserator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SQLDatabasesGetter has a method to return a SQLDatabaseInterface.
// A group's client should implement this interface.
type SQLDatabasesGetter interface {
	SQLDatabases(namespace string) SQLDatabaseInterface
}

// SQLDatabaseInterface has methods to work with SQLDatabase resources.
type SQLDatabaseInterface interface {
	Create(*v1beta1.SQLDatabase) (*v1beta1.SQLDatabase, error)
	Update(*v1beta1.SQLDatabase) (*v1beta1.SQLDatabase, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.SQLDatabase, error)
	List(opts v1.ListOptions) (*v1beta1.SQLDatabaseList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.SQLDatabase, err error)
	SQLDatabaseExpansion
}

// sQLDatabases implements SQLDatabaseInterface
type sQLDatabases struct {
	client rest.Interface
	ns     string
}

// newSQLDatabases returns a SQLDatabases
func newSQLDatabases(c *SqlV1beta1Client, namespace string) *sQLDatabases {
	return &sQLDatabases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sQLDatabase, and returns the corresponding sQLDatabase object, and an error if there is any.
func (c *sQLDatabases) Get(name string, options v1.GetOptions) (result *v1beta1.SQLDatabase, err error) {
	result = &v1beta1.SQLDatabase{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqldatabases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SQLDatabases that match those selectors.
func (c *sQLDatabases) List(opts v1.ListOptions) (result *v1beta1.SQLDatabaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.SQLDatabaseList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sQLDatabases.
func (c *sQLDatabases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sqldatabases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a sQLDatabase and creates it.  Returns the server's representation of the sQLDatabase, and an error, if there is any.
func (c *sQLDatabases) Create(sQLDatabase *v1beta1.SQLDatabase) (result *v1beta1.SQLDatabase, err error) {
	result = &v1beta1.SQLDatabase{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sqldatabases").
		Body(sQLDatabase).
		Do().
		Into(result)
	return
}

// Update takes the representation of a sQLDatabase and updates it. Returns the server's representation of the sQLDatabase, and an error, if there is any.
func (c *sQLDatabases) Update(sQLDatabase *v1beta1.SQLDatabase) (result *v1beta1.SQLDatabase, err error) {
	result = &v1beta1.SQLDatabase{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sqldatabases").
		Name(sQLDatabase.Name).
		Body(sQLDatabase).
		Do().
		Into(result)
	return
}

// Delete takes name of the sQLDatabase and deletes it. Returns an error if one occurs.
func (c *sQLDatabases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqldatabases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sQLDatabases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sqldatabases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched sQLDatabase.
func (c *sQLDatabases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.SQLDatabase, err error) {
	result = &v1beta1.SQLDatabase{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sqldatabases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}