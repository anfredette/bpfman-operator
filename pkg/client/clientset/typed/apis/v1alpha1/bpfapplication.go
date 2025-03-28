/*
Copyright 2025 The bpfman Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/bpfman/bpfman-operator/apis/v1alpha1"
	scheme "github.com/bpfman/bpfman-operator/pkg/client/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BpfApplicationsGetter has a method to return a BpfApplicationInterface.
// A group's client should implement this interface.
type BpfApplicationsGetter interface {
	BpfApplications(namespace string) BpfApplicationInterface
}

// BpfApplicationInterface has methods to work with BpfApplication resources.
type BpfApplicationInterface interface {
	Create(ctx context.Context, bpfApplication *v1alpha1.BpfApplication, opts v1.CreateOptions) (*v1alpha1.BpfApplication, error)
	Update(ctx context.Context, bpfApplication *v1alpha1.BpfApplication, opts v1.UpdateOptions) (*v1alpha1.BpfApplication, error)
	UpdateStatus(ctx context.Context, bpfApplication *v1alpha1.BpfApplication, opts v1.UpdateOptions) (*v1alpha1.BpfApplication, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.BpfApplication, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.BpfApplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BpfApplication, err error)
	BpfApplicationExpansion
}

// bpfApplications implements BpfApplicationInterface
type bpfApplications struct {
	client rest.Interface
	ns     string
}

// newBpfApplications returns a BpfApplications
func newBpfApplications(c *BpfmanV1alpha1Client, namespace string) *bpfApplications {
	return &bpfApplications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the bpfApplication, and returns the corresponding bpfApplication object, and an error if there is any.
func (c *bpfApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BpfApplication, err error) {
	result = &v1alpha1.BpfApplication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bpfapplications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BpfApplications that match those selectors.
func (c *bpfApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BpfApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.BpfApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("bpfapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bpfApplications.
func (c *bpfApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("bpfapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bpfApplication and creates it.  Returns the server's representation of the bpfApplication, and an error, if there is any.
func (c *bpfApplications) Create(ctx context.Context, bpfApplication *v1alpha1.BpfApplication, opts v1.CreateOptions) (result *v1alpha1.BpfApplication, err error) {
	result = &v1alpha1.BpfApplication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("bpfapplications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bpfApplication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bpfApplication and updates it. Returns the server's representation of the bpfApplication, and an error, if there is any.
func (c *bpfApplications) Update(ctx context.Context, bpfApplication *v1alpha1.BpfApplication, opts v1.UpdateOptions) (result *v1alpha1.BpfApplication, err error) {
	result = &v1alpha1.BpfApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bpfapplications").
		Name(bpfApplication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bpfApplication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *bpfApplications) UpdateStatus(ctx context.Context, bpfApplication *v1alpha1.BpfApplication, opts v1.UpdateOptions) (result *v1alpha1.BpfApplication, err error) {
	result = &v1alpha1.BpfApplication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("bpfapplications").
		Name(bpfApplication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bpfApplication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bpfApplication and deletes it. Returns an error if one occurs.
func (c *bpfApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bpfapplications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bpfApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("bpfapplications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bpfApplication.
func (c *bpfApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BpfApplication, err error) {
	result = &v1alpha1.BpfApplication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("bpfapplications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
