/*
Copyright 2023 The bpfman Authors.

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

package fake

import (
	"context"

	v1alpha1 "github.com/bpfman/bpfman-operator/apis/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBpfNsApplications implements BpfNsApplicationInterface
type FakeBpfNsApplications struct {
	Fake *FakeBpfmanV1alpha1
	ns   string
}

var bpfnsapplicationsResource = v1alpha1.SchemeGroupVersion.WithResource("bpfnsapplications")

var bpfnsapplicationsKind = v1alpha1.SchemeGroupVersion.WithKind("BpfNsApplication")

// Get takes name of the bpfNsApplication, and returns the corresponding bpfNsApplication object, and an error if there is any.
func (c *FakeBpfNsApplications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BpfNsApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bpfnsapplicationsResource, c.ns, name), &v1alpha1.BpfNsApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BpfNsApplication), err
}

// List takes label and field selectors, and returns the list of BpfNsApplications that match those selectors.
func (c *FakeBpfNsApplications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BpfNsApplicationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bpfnsapplicationsResource, bpfnsapplicationsKind, c.ns, opts), &v1alpha1.BpfNsApplicationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BpfNsApplicationList{ListMeta: obj.(*v1alpha1.BpfNsApplicationList).ListMeta}
	for _, item := range obj.(*v1alpha1.BpfNsApplicationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bpfNsApplications.
func (c *FakeBpfNsApplications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bpfnsapplicationsResource, c.ns, opts))

}

// Create takes the representation of a bpfNsApplication and creates it.  Returns the server's representation of the bpfNsApplication, and an error, if there is any.
func (c *FakeBpfNsApplications) Create(ctx context.Context, bpfNsApplication *v1alpha1.BpfNsApplication, opts v1.CreateOptions) (result *v1alpha1.BpfNsApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bpfnsapplicationsResource, c.ns, bpfNsApplication), &v1alpha1.BpfNsApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BpfNsApplication), err
}

// Update takes the representation of a bpfNsApplication and updates it. Returns the server's representation of the bpfNsApplication, and an error, if there is any.
func (c *FakeBpfNsApplications) Update(ctx context.Context, bpfNsApplication *v1alpha1.BpfNsApplication, opts v1.UpdateOptions) (result *v1alpha1.BpfNsApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bpfnsapplicationsResource, c.ns, bpfNsApplication), &v1alpha1.BpfNsApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BpfNsApplication), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBpfNsApplications) UpdateStatus(ctx context.Context, bpfNsApplication *v1alpha1.BpfNsApplication, opts v1.UpdateOptions) (*v1alpha1.BpfNsApplication, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bpfnsapplicationsResource, "status", c.ns, bpfNsApplication), &v1alpha1.BpfNsApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BpfNsApplication), err
}

// Delete takes name of the bpfNsApplication and deletes it. Returns an error if one occurs.
func (c *FakeBpfNsApplications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(bpfnsapplicationsResource, c.ns, name, opts), &v1alpha1.BpfNsApplication{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBpfNsApplications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bpfnsapplicationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BpfNsApplicationList{})
	return err
}

// Patch applies the patch and returns the patched bpfNsApplication.
func (c *FakeBpfNsApplications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BpfNsApplication, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bpfnsapplicationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BpfNsApplication{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BpfNsApplication), err
}
