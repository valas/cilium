// Copyright 2017-2019 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v2

import (
	"time"

	v2 "github.com/cilium/cilium/pkg/k8s/apis/cilium.io/v2"
	scheme "github.com/cilium/cilium/pkg/k8s/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CiliumNodesGetter has a method to return a CiliumNodeInterface.
// A group's client should implement this interface.
type CiliumNodesGetter interface {
	CiliumNodes() CiliumNodeInterface
}

// CiliumNodeInterface has methods to work with CiliumNode resources.
type CiliumNodeInterface interface {
	Create(*v2.CiliumNode) (*v2.CiliumNode, error)
	Update(*v2.CiliumNode) (*v2.CiliumNode, error)
	UpdateStatus(*v2.CiliumNode) (*v2.CiliumNode, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v2.CiliumNode, error)
	List(opts v1.ListOptions) (*v2.CiliumNodeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.CiliumNode, err error)
	CiliumNodeExpansion
}

// ciliumNodes implements CiliumNodeInterface
type ciliumNodes struct {
	client rest.Interface
}

// newCiliumNodes returns a CiliumNodes
func newCiliumNodes(c *CiliumV2Client) *ciliumNodes {
	return &ciliumNodes{
		client: c.RESTClient(),
	}
}

// Get takes name of the ciliumNode, and returns the corresponding ciliumNode object, and an error if there is any.
func (c *ciliumNodes) Get(name string, options v1.GetOptions) (result *v2.CiliumNode, err error) {
	result = &v2.CiliumNode{}
	err = c.client.Get().
		Resource("ciliumnodes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CiliumNodes that match those selectors.
func (c *ciliumNodes) List(opts v1.ListOptions) (result *v2.CiliumNodeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.CiliumNodeList{}
	err = c.client.Get().
		Resource("ciliumnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested ciliumNodes.
func (c *ciliumNodes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ciliumnodes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a ciliumNode and creates it.  Returns the server's representation of the ciliumNode, and an error, if there is any.
func (c *ciliumNodes) Create(ciliumNode *v2.CiliumNode) (result *v2.CiliumNode, err error) {
	result = &v2.CiliumNode{}
	err = c.client.Post().
		Resource("ciliumnodes").
		Body(ciliumNode).
		Do().
		Into(result)
	return
}

// Update takes the representation of a ciliumNode and updates it. Returns the server's representation of the ciliumNode, and an error, if there is any.
func (c *ciliumNodes) Update(ciliumNode *v2.CiliumNode) (result *v2.CiliumNode, err error) {
	result = &v2.CiliumNode{}
	err = c.client.Put().
		Resource("ciliumnodes").
		Name(ciliumNode.Name).
		Body(ciliumNode).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *ciliumNodes) UpdateStatus(ciliumNode *v2.CiliumNode) (result *v2.CiliumNode, err error) {
	result = &v2.CiliumNode{}
	err = c.client.Put().
		Resource("ciliumnodes").
		Name(ciliumNode.Name).
		SubResource("status").
		Body(ciliumNode).
		Do().
		Into(result)
	return
}

// Delete takes name of the ciliumNode and deletes it. Returns an error if one occurs.
func (c *ciliumNodes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ciliumnodes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *ciliumNodes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ciliumnodes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched ciliumNode.
func (c *ciliumNodes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.CiliumNode, err error) {
	result = &v2.CiliumNode{}
	err = c.client.Patch(pt).
		Resource("ciliumnodes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
