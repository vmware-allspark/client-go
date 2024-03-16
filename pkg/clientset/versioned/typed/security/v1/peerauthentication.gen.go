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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"
	"time"

	v1 "istio.io/client-go/pkg/apis/security/v1"
	securityv1 "istio.io/client-go/pkg/applyconfiguration/security/v1"
	scheme "istio.io/client-go/pkg/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PeerAuthenticationsGetter has a method to return a PeerAuthenticationInterface.
// A group's client should implement this interface.
type PeerAuthenticationsGetter interface {
	PeerAuthentications(namespace string) PeerAuthenticationInterface
}

// PeerAuthenticationInterface has methods to work with PeerAuthentication resources.
type PeerAuthenticationInterface interface {
	Create(ctx context.Context, peerAuthentication *v1.PeerAuthentication, opts metav1.CreateOptions) (*v1.PeerAuthentication, error)
	Update(ctx context.Context, peerAuthentication *v1.PeerAuthentication, opts metav1.UpdateOptions) (*v1.PeerAuthentication, error)
	UpdateStatus(ctx context.Context, peerAuthentication *v1.PeerAuthentication, opts metav1.UpdateOptions) (*v1.PeerAuthentication, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.PeerAuthentication, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PeerAuthenticationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PeerAuthentication, err error)
	Apply(ctx context.Context, peerAuthentication *securityv1.PeerAuthenticationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PeerAuthentication, err error)
	ApplyStatus(ctx context.Context, peerAuthentication *securityv1.PeerAuthenticationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PeerAuthentication, err error)
	PeerAuthenticationExpansion
}

// peerAuthentications implements PeerAuthenticationInterface
type peerAuthentications struct {
	client rest.Interface
	ns     string
}

// newPeerAuthentications returns a PeerAuthentications
func newPeerAuthentications(c *SecurityV1Client, namespace string) *peerAuthentications {
	return &peerAuthentications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the peerAuthentication, and returns the corresponding peerAuthentication object, and an error if there is any.
func (c *peerAuthentications) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.PeerAuthentication, err error) {
	result = &v1.PeerAuthentication{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("peerauthentications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PeerAuthentications that match those selectors.
func (c *peerAuthentications) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PeerAuthenticationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PeerAuthenticationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("peerauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested peerAuthentications.
func (c *peerAuthentications) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("peerauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a peerAuthentication and creates it.  Returns the server's representation of the peerAuthentication, and an error, if there is any.
func (c *peerAuthentications) Create(ctx context.Context, peerAuthentication *v1.PeerAuthentication, opts metav1.CreateOptions) (result *v1.PeerAuthentication, err error) {
	result = &v1.PeerAuthentication{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("peerauthentications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(peerAuthentication).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a peerAuthentication and updates it. Returns the server's representation of the peerAuthentication, and an error, if there is any.
func (c *peerAuthentications) Update(ctx context.Context, peerAuthentication *v1.PeerAuthentication, opts metav1.UpdateOptions) (result *v1.PeerAuthentication, err error) {
	result = &v1.PeerAuthentication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("peerauthentications").
		Name(peerAuthentication.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(peerAuthentication).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *peerAuthentications) UpdateStatus(ctx context.Context, peerAuthentication *v1.PeerAuthentication, opts metav1.UpdateOptions) (result *v1.PeerAuthentication, err error) {
	result = &v1.PeerAuthentication{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("peerauthentications").
		Name(peerAuthentication.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(peerAuthentication).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the peerAuthentication and deletes it. Returns an error if one occurs.
func (c *peerAuthentications) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("peerauthentications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *peerAuthentications) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("peerauthentications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched peerAuthentication.
func (c *peerAuthentications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.PeerAuthentication, err error) {
	result = &v1.PeerAuthentication{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("peerauthentications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// Apply takes the given apply declarative configuration, applies it and returns the applied peerAuthentication.
func (c *peerAuthentications) Apply(ctx context.Context, peerAuthentication *securityv1.PeerAuthenticationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PeerAuthentication, err error) {
	if peerAuthentication == nil {
		return nil, fmt.Errorf("peerAuthentication provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(peerAuthentication)
	if err != nil {
		return nil, err
	}
	name := peerAuthentication.Name
	if name == nil {
		return nil, fmt.Errorf("peerAuthentication.Name must be provided to Apply")
	}
	result = &v1.PeerAuthentication{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("peerauthentications").
		Name(*name).
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// ApplyStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
func (c *peerAuthentications) ApplyStatus(ctx context.Context, peerAuthentication *securityv1.PeerAuthenticationApplyConfiguration, opts metav1.ApplyOptions) (result *v1.PeerAuthentication, err error) {
	if peerAuthentication == nil {
		return nil, fmt.Errorf("peerAuthentication provided to Apply must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(peerAuthentication)
	if err != nil {
		return nil, err
	}

	name := peerAuthentication.Name
	if name == nil {
		return nil, fmt.Errorf("peerAuthentication.Name must be provided to Apply")
	}

	result = &v1.PeerAuthentication{}
	err = c.client.Patch(types.ApplyPatchType).
		Namespace(c.ns).
		Resource("peerauthentications").
		Name(*name).
		SubResource("status").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}