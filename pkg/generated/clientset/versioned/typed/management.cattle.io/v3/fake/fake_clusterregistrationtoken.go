/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRegistrationTokens implements ClusterRegistrationTokenInterface
type FakeClusterRegistrationTokens struct {
	Fake *FakeManagementV3
	ns   string
}

var clusterregistrationtokensResource = v3.SchemeGroupVersion.WithResource("clusterregistrationtokens")

var clusterregistrationtokensKind = v3.SchemeGroupVersion.WithKind("ClusterRegistrationToken")

// Get takes name of the clusterRegistrationToken, and returns the corresponding clusterRegistrationToken object, and an error if there is any.
func (c *FakeClusterRegistrationTokens) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.ClusterRegistrationToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clusterregistrationtokensResource, c.ns, name), &v3.ClusterRegistrationToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterRegistrationToken), err
}

// List takes label and field selectors, and returns the list of ClusterRegistrationTokens that match those selectors.
func (c *FakeClusterRegistrationTokens) List(ctx context.Context, opts v1.ListOptions) (result *v3.ClusterRegistrationTokenList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clusterregistrationtokensResource, clusterregistrationtokensKind, c.ns, opts), &v3.ClusterRegistrationTokenList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.ClusterRegistrationTokenList{ListMeta: obj.(*v3.ClusterRegistrationTokenList).ListMeta}
	for _, item := range obj.(*v3.ClusterRegistrationTokenList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRegistrationTokens.
func (c *FakeClusterRegistrationTokens) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clusterregistrationtokensResource, c.ns, opts))

}

// Create takes the representation of a clusterRegistrationToken and creates it.  Returns the server's representation of the clusterRegistrationToken, and an error, if there is any.
func (c *FakeClusterRegistrationTokens) Create(ctx context.Context, clusterRegistrationToken *v3.ClusterRegistrationToken, opts v1.CreateOptions) (result *v3.ClusterRegistrationToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clusterregistrationtokensResource, c.ns, clusterRegistrationToken), &v3.ClusterRegistrationToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterRegistrationToken), err
}

// Update takes the representation of a clusterRegistrationToken and updates it. Returns the server's representation of the clusterRegistrationToken, and an error, if there is any.
func (c *FakeClusterRegistrationTokens) Update(ctx context.Context, clusterRegistrationToken *v3.ClusterRegistrationToken, opts v1.UpdateOptions) (result *v3.ClusterRegistrationToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clusterregistrationtokensResource, c.ns, clusterRegistrationToken), &v3.ClusterRegistrationToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterRegistrationToken), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterRegistrationTokens) UpdateStatus(ctx context.Context, clusterRegistrationToken *v3.ClusterRegistrationToken, opts v1.UpdateOptions) (*v3.ClusterRegistrationToken, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(clusterregistrationtokensResource, "status", c.ns, clusterRegistrationToken), &v3.ClusterRegistrationToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterRegistrationToken), err
}

// Delete takes name of the clusterRegistrationToken and deletes it. Returns an error if one occurs.
func (c *FakeClusterRegistrationTokens) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clusterregistrationtokensResource, c.ns, name, opts), &v3.ClusterRegistrationToken{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRegistrationTokens) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clusterregistrationtokensResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.ClusterRegistrationTokenList{})
	return err
}

// Patch applies the patch and returns the patched clusterRegistrationToken.
func (c *FakeClusterRegistrationTokens) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.ClusterRegistrationToken, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clusterregistrationtokensResource, c.ns, name, pt, data, subresources...), &v3.ClusterRegistrationToken{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.ClusterRegistrationToken), err
}
