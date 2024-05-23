/*
   zlw
*/
// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	scheme "k8s-controller-sample/api/generated/clientset/versioned/scheme"
	v1 "k8s-controller-sample/api/sample/v1"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SamplesGetter has a method to return a SampleInterface.
// A group's client should implement this interface.
type SamplesGetter interface {
	Samples(namespace string) SampleInterface
}

// SampleInterface has methods to work with Sample resources.
type SampleInterface interface {
	Create(ctx context.Context, sample *v1.Sample, opts metav1.CreateOptions) (*v1.Sample, error)
	Update(ctx context.Context, sample *v1.Sample, opts metav1.UpdateOptions) (*v1.Sample, error)
	UpdateStatus(ctx context.Context, sample *v1.Sample, opts metav1.UpdateOptions) (*v1.Sample, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Sample, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SampleList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Sample, err error)
	SampleExpansion
}

// samples implements SampleInterface
type samples struct {
	client rest.Interface
	ns     string
}

// newSamples returns a Samples
func newSamples(c *SampleV1Client, namespace string) *samples {
	return &samples{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sample, and returns the corresponding sample object, and an error if there is any.
func (c *samples) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Sample, err error) {
	result = &v1.Sample{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("samples").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Samples that match those selectors.
func (c *samples) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SampleList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SampleList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("samples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested samples.
func (c *samples) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("samples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sample and creates it.  Returns the server's representation of the sample, and an error, if there is any.
func (c *samples) Create(ctx context.Context, sample *v1.Sample, opts metav1.CreateOptions) (result *v1.Sample, err error) {
	result = &v1.Sample{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("samples").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sample).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sample and updates it. Returns the server's representation of the sample, and an error, if there is any.
func (c *samples) Update(ctx context.Context, sample *v1.Sample, opts metav1.UpdateOptions) (result *v1.Sample, err error) {
	result = &v1.Sample{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("samples").
		Name(sample.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sample).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *samples) UpdateStatus(ctx context.Context, sample *v1.Sample, opts metav1.UpdateOptions) (result *v1.Sample, err error) {
	result = &v1.Sample{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("samples").
		Name(sample.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sample).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sample and deletes it. Returns an error if one occurs.
func (c *samples) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("samples").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *samples) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("samples").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sample.
func (c *samples) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Sample, err error) {
	result = &v1.Sample{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("samples").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
