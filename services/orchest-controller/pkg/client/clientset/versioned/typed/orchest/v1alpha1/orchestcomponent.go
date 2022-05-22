/*
Copyright 2022 The orchest Authors.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/orchest/orchest/services/orchest-controller/pkg/apis/orchest/v1alpha1"
	scheme "github.com/orchest/orchest/services/orchest-controller/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OrchestComponentsGetter has a method to return a OrchestComponentInterface.
// A group's client should implement this interface.
type OrchestComponentsGetter interface {
	OrchestComponents(namespace string) OrchestComponentInterface
}

// OrchestComponentInterface has methods to work with OrchestComponent resources.
type OrchestComponentInterface interface {
	Create(ctx context.Context, orchestComponent *v1alpha1.OrchestComponent, opts v1.CreateOptions) (*v1alpha1.OrchestComponent, error)
	Update(ctx context.Context, orchestComponent *v1alpha1.OrchestComponent, opts v1.UpdateOptions) (*v1alpha1.OrchestComponent, error)
	UpdateStatus(ctx context.Context, orchestComponent *v1alpha1.OrchestComponent, opts v1.UpdateOptions) (*v1alpha1.OrchestComponent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OrchestComponent, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OrchestComponentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrchestComponent, err error)
	OrchestComponentExpansion
}

// orchestComponents implements OrchestComponentInterface
type orchestComponents struct {
	client rest.Interface
	ns     string
}

// newOrchestComponents returns a OrchestComponents
func newOrchestComponents(c *OrchestV1alpha1Client, namespace string) *orchestComponents {
	return &orchestComponents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the orchestComponent, and returns the corresponding orchestComponent object, and an error if there is any.
func (c *orchestComponents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OrchestComponent, err error) {
	result = &v1alpha1.OrchestComponent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("orchestcomponents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OrchestComponents that match those selectors.
func (c *orchestComponents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrchestComponentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OrchestComponentList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("orchestcomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested orchestComponents.
func (c *orchestComponents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("orchestcomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a orchestComponent and creates it.  Returns the server's representation of the orchestComponent, and an error, if there is any.
func (c *orchestComponents) Create(ctx context.Context, orchestComponent *v1alpha1.OrchestComponent, opts v1.CreateOptions) (result *v1alpha1.OrchestComponent, err error) {
	result = &v1alpha1.OrchestComponent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("orchestcomponents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orchestComponent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a orchestComponent and updates it. Returns the server's representation of the orchestComponent, and an error, if there is any.
func (c *orchestComponents) Update(ctx context.Context, orchestComponent *v1alpha1.OrchestComponent, opts v1.UpdateOptions) (result *v1alpha1.OrchestComponent, err error) {
	result = &v1alpha1.OrchestComponent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("orchestcomponents").
		Name(orchestComponent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orchestComponent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *orchestComponents) UpdateStatus(ctx context.Context, orchestComponent *v1alpha1.OrchestComponent, opts v1.UpdateOptions) (result *v1alpha1.OrchestComponent, err error) {
	result = &v1alpha1.OrchestComponent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("orchestcomponents").
		Name(orchestComponent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orchestComponent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the orchestComponent and deletes it. Returns an error if one occurs.
func (c *orchestComponents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("orchestcomponents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *orchestComponents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("orchestcomponents").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched orchestComponent.
func (c *orchestComponents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrchestComponent, err error) {
	result = &v1alpha1.OrchestComponent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("orchestcomponents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
