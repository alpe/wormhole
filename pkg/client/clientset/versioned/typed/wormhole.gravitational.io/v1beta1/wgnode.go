/*
Copyright 2019 Gravitational, Inc.

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

package v1beta1

import (
	"time"

	v1beta1 "github.com/gravitational/wormhole/pkg/apis/wormhole.gravitational.io/v1beta1"
	scheme "github.com/gravitational/wormhole/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// WgnodesGetter has a method to return a WgnodeInterface.
// A group's client should implement this interface.
type WgnodesGetter interface {
	Wgnodes(namespace string) WgnodeInterface
}

// WgnodeInterface has methods to work with Wgnode resources.
type WgnodeInterface interface {
	Create(*v1beta1.Wgnode) (*v1beta1.Wgnode, error)
	Update(*v1beta1.Wgnode) (*v1beta1.Wgnode, error)
	UpdateStatus(*v1beta1.Wgnode) (*v1beta1.Wgnode, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.Wgnode, error)
	List(opts v1.ListOptions) (*v1beta1.WgnodeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Wgnode, err error)
	WgnodeExpansion
}

// wgnodes implements WgnodeInterface
type wgnodes struct {
	client rest.Interface
	ns     string
}

// newWgnodes returns a Wgnodes
func newWgnodes(c *WormholeV1beta1Client, namespace string) *wgnodes {
	return &wgnodes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the wgnode, and returns the corresponding wgnode object, and an error if there is any.
func (c *wgnodes) Get(name string, options v1.GetOptions) (result *v1beta1.Wgnode, err error) {
	result = &v1beta1.Wgnode{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wgnodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Wgnodes that match those selectors.
func (c *wgnodes) List(opts v1.ListOptions) (result *v1beta1.WgnodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.WgnodeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("wgnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested wgnodes.
func (c *wgnodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("wgnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a wgnode and creates it.  Returns the server's representation of the wgnode, and an error, if there is any.
func (c *wgnodes) Create(wgnode *v1beta1.Wgnode) (result *v1beta1.Wgnode, err error) {
	result = &v1beta1.Wgnode{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("wgnodes").
		Body(wgnode).
		Do().
		Into(result)
	return
}

// Update takes the representation of a wgnode and updates it. Returns the server's representation of the wgnode, and an error, if there is any.
func (c *wgnodes) Update(wgnode *v1beta1.Wgnode) (result *v1beta1.Wgnode, err error) {
	result = &v1beta1.Wgnode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wgnodes").
		Name(wgnode.Name).
		Body(wgnode).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *wgnodes) UpdateStatus(wgnode *v1beta1.Wgnode) (result *v1beta1.Wgnode, err error) {
	result = &v1beta1.Wgnode{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("wgnodes").
		Name(wgnode.Name).
		SubResource("status").
		Body(wgnode).
		Do().
		Into(result)
	return
}

// Delete takes name of the wgnode and deletes it. Returns an error if one occurs.
func (c *wgnodes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wgnodes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *wgnodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("wgnodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched wgnode.
func (c *wgnodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.Wgnode, err error) {
	result = &v1beta1.Wgnode{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("wgnodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}