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

package v1alpha2

import (
	"time"

	v1alpha2 "istio.io/client-go/pkg/apis/config/v1alpha2"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// QuotaSpecsGetter has a method to return a QuotaSpecInterface.
// A group's client should implement this interface.
type QuotaSpecsGetter interface {
	QuotaSpecs(namespace string) QuotaSpecInterface
}

// QuotaSpecInterface has methods to work with QuotaSpec resources.
type QuotaSpecInterface interface {
	Create(*v1alpha2.QuotaSpec) (*v1alpha2.QuotaSpec, error)
	Update(*v1alpha2.QuotaSpec) (*v1alpha2.QuotaSpec, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha2.QuotaSpec, error)
	List(opts v1.ListOptions) (*v1alpha2.QuotaSpecList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.QuotaSpec, err error)
	QuotaSpecExpansion
}

// quotaSpecs implements QuotaSpecInterface
type quotaSpecs struct {
	client rest.Interface
	ns     string
}

// newQuotaSpecs returns a QuotaSpecs
func newQuotaSpecs(c *ConfigV1alpha2Client, namespace string) *quotaSpecs {
	return &quotaSpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the quotaSpec, and returns the corresponding quotaSpec object, and an error if there is any.
func (c *quotaSpecs) Get(name string, options v1.GetOptions) (result *v1alpha2.QuotaSpec, err error) {
	result = &v1alpha2.QuotaSpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of QuotaSpecs that match those selectors.
func (c *quotaSpecs) List(opts v1.ListOptions) (result *v1alpha2.QuotaSpecList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.QuotaSpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested quotaSpecs.
func (c *quotaSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("quotaspecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a quotaSpec and creates it.  Returns the server's representation of the quotaSpec, and an error, if there is any.
func (c *quotaSpecs) Create(quotaSpec *v1alpha2.QuotaSpec) (result *v1alpha2.QuotaSpec, err error) {
	result = &v1alpha2.QuotaSpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("quotaspecs").
		Body(quotaSpec).
		Do().
		Into(result)
	return
}

// Update takes the representation of a quotaSpec and updates it. Returns the server's representation of the quotaSpec, and an error, if there is any.
func (c *quotaSpecs) Update(quotaSpec *v1alpha2.QuotaSpec) (result *v1alpha2.QuotaSpec, err error) {
	result = &v1alpha2.QuotaSpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("quotaspecs").
		Name(quotaSpec.Name).
		Body(quotaSpec).
		Do().
		Into(result)
	return
}

// Delete takes name of the quotaSpec and deletes it. Returns an error if one occurs.
func (c *quotaSpecs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quotaspecs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *quotaSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("quotaspecs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched quotaSpec.
func (c *quotaSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.QuotaSpec, err error) {
	result = &v1alpha2.QuotaSpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("quotaspecs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
