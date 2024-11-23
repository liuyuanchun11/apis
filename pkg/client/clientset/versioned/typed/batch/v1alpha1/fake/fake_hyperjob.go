/*
Copyright The Volcano Authors.

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
	json "encoding/json"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "volcano.sh/apis/pkg/apis/batch/v1alpha1"
	batchv1alpha1 "volcano.sh/apis/pkg/client/applyconfiguration/batch/v1alpha1"
)

// FakeHyperJobs implements HyperJobInterface
type FakeHyperJobs struct {
	Fake *FakeBatchV1alpha1
	ns   string
}

var hyperjobsResource = v1alpha1.SchemeGroupVersion.WithResource("hyperjobs")

var hyperjobsKind = v1alpha1.SchemeGroupVersion.WithKind("HyperJob")

// Get takes name of the hyperJob, and returns the corresponding hyperJob object, and an error if there is any.
func (c *FakeHyperJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HyperJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(hyperjobsResource, c.ns, name), &v1alpha1.HyperJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperJob), err
}

// List takes label and field selectors, and returns the list of HyperJobs that match those selectors.
func (c *FakeHyperJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HyperJobList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(hyperjobsResource, hyperjobsKind, c.ns, opts), &v1alpha1.HyperJobList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HyperJobList{ListMeta: obj.(*v1alpha1.HyperJobList).ListMeta}
	for _, item := range obj.(*v1alpha1.HyperJobList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hyperJobs.
func (c *FakeHyperJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(hyperjobsResource, c.ns, opts))

}

// Create takes the representation of a hyperJob and creates it.  Returns the server's representation of the hyperJob, and an error, if there is any.
func (c *FakeHyperJobs) Create(ctx context.Context, hyperJob *v1alpha1.HyperJob, opts v1.CreateOptions) (result *v1alpha1.HyperJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(hyperjobsResource, c.ns, hyperJob), &v1alpha1.HyperJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperJob), err
}

// Update takes the representation of a hyperJob and updates it. Returns the server's representation of the hyperJob, and an error, if there is any.
func (c *FakeHyperJobs) Update(ctx context.Context, hyperJob *v1alpha1.HyperJob, opts v1.UpdateOptions) (result *v1alpha1.HyperJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(hyperjobsResource, c.ns, hyperJob), &v1alpha1.HyperJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperJob), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeHyperJobs) UpdateStatus(ctx context.Context, hyperJob *v1alpha1.HyperJob, opts v1.UpdateOptions) (*v1alpha1.HyperJob, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(hyperjobsResource, "status", c.ns, hyperJob), &v1alpha1.HyperJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperJob), err
}

// Delete takes name of the hyperJob and deletes it. Returns an error if one occurs.
func (c *FakeHyperJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(hyperjobsResource, c.ns, name, opts), &v1alpha1.HyperJob{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHyperJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(hyperjobsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.HyperJobList{})
	return err
}

// Patch applies the patch and returns the patched hyperJob.
func (c *FakeHyperJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HyperJob, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hyperjobsResource, c.ns, name, pt, data, subresources...), &v1alpha1.HyperJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperJob), err
}

// Apply takes the given apply declarative configuration, applies it and returns the applied hyperJob.
func (c *FakeHyperJobs) Apply(ctx context.Context, hyperJob *batchv1alpha1.HyperJobApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.HyperJob, err error) {
	if hyperJob == nil {
		return nil, fmt.Errorf("hyperJob provided to Apply must not be nil")
	}
	data, err := json.Marshal(hyperJob)
	if err != nil {
		return nil, err
	}
	name := hyperJob.Name
	if name == nil {
		return nil, fmt.Errorf("hyperJob.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hyperjobsResource, c.ns, *name, types.ApplyPatchType, data), &v1alpha1.HyperJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperJob), err
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *FakeHyperJobs) ApplyStatus(ctx context.Context, hyperJob *batchv1alpha1.HyperJobApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha1.HyperJob, err error) {
	if hyperJob == nil {
		return nil, fmt.Errorf("hyperJob provided to Apply must not be nil")
	}
	data, err := json.Marshal(hyperJob)
	if err != nil {
		return nil, err
	}
	name := hyperJob.Name
	if name == nil {
		return nil, fmt.Errorf("hyperJob.Name must be provided to Apply")
	}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(hyperjobsResource, c.ns, *name, types.ApplyPatchType, data, "status"), &v1alpha1.HyperJob{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HyperJob), err
}
