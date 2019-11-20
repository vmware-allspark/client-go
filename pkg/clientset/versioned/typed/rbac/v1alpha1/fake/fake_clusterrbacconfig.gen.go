// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "istio.io/client-go/pkg/apis/rbac/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterRbacConfigs implements ClusterRbacConfigInterface
type FakeClusterRbacConfigs struct {
	Fake *FakeRbacV1alpha1
}

var clusterrbacconfigsResource = schema.GroupVersionResource{Group: "rbac.istio.io", Version: "v1alpha1", Resource: "clusterrbacconfigs"}

var clusterrbacconfigsKind = schema.GroupVersionKind{Group: "rbac.istio.io", Version: "v1alpha1", Kind: "ClusterRbacConfig"}

// Get takes name of the clusterRbacConfig, and returns the corresponding clusterRbacConfig object, and an error if there is any.
func (c *FakeClusterRbacConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterRbacConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterrbacconfigsResource, name), &v1alpha1.ClusterRbacConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRbacConfig), err
}

// List takes label and field selectors, and returns the list of ClusterRbacConfigs that match those selectors.
func (c *FakeClusterRbacConfigs) List(opts v1.ListOptions) (result *v1alpha1.ClusterRbacConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterrbacconfigsResource, clusterrbacconfigsKind, opts), &v1alpha1.ClusterRbacConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterRbacConfigList{ListMeta: obj.(*v1alpha1.ClusterRbacConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterRbacConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterRbacConfigs.
func (c *FakeClusterRbacConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterrbacconfigsResource, opts))
}

// Create takes the representation of a clusterRbacConfig and creates it.  Returns the server's representation of the clusterRbacConfig, and an error, if there is any.
func (c *FakeClusterRbacConfigs) Create(clusterRbacConfig *v1alpha1.ClusterRbacConfig) (result *v1alpha1.ClusterRbacConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterrbacconfigsResource, clusterRbacConfig), &v1alpha1.ClusterRbacConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRbacConfig), err
}

// Update takes the representation of a clusterRbacConfig and updates it. Returns the server's representation of the clusterRbacConfig, and an error, if there is any.
func (c *FakeClusterRbacConfigs) Update(clusterRbacConfig *v1alpha1.ClusterRbacConfig) (result *v1alpha1.ClusterRbacConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterrbacconfigsResource, clusterRbacConfig), &v1alpha1.ClusterRbacConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRbacConfig), err
}

// Delete takes name of the clusterRbacConfig and deletes it. Returns an error if one occurs.
func (c *FakeClusterRbacConfigs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterrbacconfigsResource, name), &v1alpha1.ClusterRbacConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterRbacConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterrbacconfigsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterRbacConfigList{})
	return err
}

// Patch applies the patch and returns the patched clusterRbacConfig.
func (c *FakeClusterRbacConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterRbacConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterrbacconfigsResource, name, pt, data, subresources...), &v1alpha1.ClusterRbacConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterRbacConfig), err
}
