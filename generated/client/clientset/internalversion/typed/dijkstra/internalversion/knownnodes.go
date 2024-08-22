/*
Copyright (C) 2024 JinLi Co.,Ltd. All rights reserved.

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

package internalversion

import (
	"context"
	"time"

	scheme "jinli.io/shortestpath/generated/client/clientset/internalversion/scheme"
	dijkstra "jinli.io/shortestpath/pkg/apis/dijkstra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KnownNodesesGetter has a method to return a KnownNodesInterface.
// A group's client should implement this interface.
type KnownNodesesGetter interface {
	KnownNodeses(namespace string) KnownNodesInterface
}

// KnownNodesInterface has methods to work with KnownNodes resources.
type KnownNodesInterface interface {
	Create(ctx context.Context, knownNodes *dijkstra.KnownNodes, opts v1.CreateOptions) (*dijkstra.KnownNodes, error)
	Update(ctx context.Context, knownNodes *dijkstra.KnownNodes, opts v1.UpdateOptions) (*dijkstra.KnownNodes, error)
	UpdateStatus(ctx context.Context, knownNodes *dijkstra.KnownNodes, opts v1.UpdateOptions) (*dijkstra.KnownNodes, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*dijkstra.KnownNodes, error)
	List(ctx context.Context, opts v1.ListOptions) (*dijkstra.KnownNodesList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *dijkstra.KnownNodes, err error)
	KnownNodesExpansion
}

// knownNodeses implements KnownNodesInterface
type knownNodeses struct {
	client rest.Interface
	ns     string
}

// newKnownNodeses returns a KnownNodeses
func newKnownNodeses(c *DijkstraClient, namespace string) *knownNodeses {
	return &knownNodeses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the knownNodes, and returns the corresponding knownNodes object, and an error if there is any.
func (c *knownNodeses) Get(ctx context.Context, name string, options v1.GetOptions) (result *dijkstra.KnownNodes, err error) {
	result = &dijkstra.KnownNodes{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("knownnodeses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KnownNodeses that match those selectors.
func (c *knownNodeses) List(ctx context.Context, opts v1.ListOptions) (result *dijkstra.KnownNodesList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &dijkstra.KnownNodesList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("knownnodeses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested knownNodeses.
func (c *knownNodeses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("knownnodeses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a knownNodes and creates it.  Returns the server's representation of the knownNodes, and an error, if there is any.
func (c *knownNodeses) Create(ctx context.Context, knownNodes *dijkstra.KnownNodes, opts v1.CreateOptions) (result *dijkstra.KnownNodes, err error) {
	result = &dijkstra.KnownNodes{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("knownnodeses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(knownNodes).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a knownNodes and updates it. Returns the server's representation of the knownNodes, and an error, if there is any.
func (c *knownNodeses) Update(ctx context.Context, knownNodes *dijkstra.KnownNodes, opts v1.UpdateOptions) (result *dijkstra.KnownNodes, err error) {
	result = &dijkstra.KnownNodes{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("knownnodeses").
		Name(knownNodes.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(knownNodes).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *knownNodeses) UpdateStatus(ctx context.Context, knownNodes *dijkstra.KnownNodes, opts v1.UpdateOptions) (result *dijkstra.KnownNodes, err error) {
	result = &dijkstra.KnownNodes{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("knownnodeses").
		Name(knownNodes.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(knownNodes).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the knownNodes and deletes it. Returns an error if one occurs.
func (c *knownNodeses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("knownnodeses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *knownNodeses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("knownnodeses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched knownNodes.
func (c *knownNodeses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *dijkstra.KnownNodes, err error) {
	result = &dijkstra.KnownNodes{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("knownnodeses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
