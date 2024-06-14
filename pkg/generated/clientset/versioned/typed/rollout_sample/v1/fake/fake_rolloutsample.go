/*
Copyright The Kubernetes Authors.

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
	v1 "k8s-controller-sample/api/rollout_sample/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRolloutSamples implements RolloutSampleInterface
type FakeRolloutSamples struct {
	Fake *FakeZlwV1
	ns   string
}

var rolloutsamplesResource = v1.SchemeGroupVersion.WithResource("rolloutsamples")

var rolloutsamplesKind = v1.SchemeGroupVersion.WithKind("RolloutSample")

// Get takes name of the rolloutSample, and returns the corresponding rolloutSample object, and an error if there is any.
func (c *FakeRolloutSamples) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.RolloutSample, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rolloutsamplesResource, c.ns, name), &v1.RolloutSample{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RolloutSample), err
}

// List takes label and field selectors, and returns the list of RolloutSamples that match those selectors.
func (c *FakeRolloutSamples) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RolloutSampleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rolloutsamplesResource, rolloutsamplesKind, c.ns, opts), &v1.RolloutSampleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RolloutSampleList{ListMeta: obj.(*v1.RolloutSampleList).ListMeta}
	for _, item := range obj.(*v1.RolloutSampleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rolloutSamples.
func (c *FakeRolloutSamples) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rolloutsamplesResource, c.ns, opts))

}

// Create takes the representation of a rolloutSample and creates it.  Returns the server's representation of the rolloutSample, and an error, if there is any.
func (c *FakeRolloutSamples) Create(ctx context.Context, rolloutSample *v1.RolloutSample, opts metav1.CreateOptions) (result *v1.RolloutSample, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rolloutsamplesResource, c.ns, rolloutSample), &v1.RolloutSample{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RolloutSample), err
}

// Update takes the representation of a rolloutSample and updates it. Returns the server's representation of the rolloutSample, and an error, if there is any.
func (c *FakeRolloutSamples) Update(ctx context.Context, rolloutSample *v1.RolloutSample, opts metav1.UpdateOptions) (result *v1.RolloutSample, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rolloutsamplesResource, c.ns, rolloutSample), &v1.RolloutSample{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RolloutSample), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRolloutSamples) UpdateStatus(ctx context.Context, rolloutSample *v1.RolloutSample, opts metav1.UpdateOptions) (*v1.RolloutSample, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rolloutsamplesResource, "status", c.ns, rolloutSample), &v1.RolloutSample{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RolloutSample), err
}

// Delete takes name of the rolloutSample and deletes it. Returns an error if one occurs.
func (c *FakeRolloutSamples) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rolloutsamplesResource, c.ns, name, opts), &v1.RolloutSample{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRolloutSamples) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rolloutsamplesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RolloutSampleList{})
	return err
}

// Patch applies the patch and returns the patched rolloutSample.
func (c *FakeRolloutSamples) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.RolloutSample, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rolloutsamplesResource, c.ns, name, pt, data, subresources...), &v1.RolloutSample{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.RolloutSample), err
}