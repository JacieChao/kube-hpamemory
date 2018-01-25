/*
Copyright 2018 The Openshift Evangelists

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
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "kube-hpamemory/pkg/apis/horhpamemory/v1"
	scheme "kube-hpamemory/pkg/client/clientset/versioned/scheme"
)

// HORHPAMemoriesGetter has a method to return a HORHPAMemoryInterface.
// A group's client should implement this interface.
type HORHPAMemoriesGetter interface {
	HORHPAMemories(namespace string) HORHPAMemoryInterface
}

// HORHPAMemoryInterface has methods to work with HORHPAMemory resources.
type HORHPAMemoryInterface interface {
	Create(*v1.HORHPAMemory) (*v1.HORHPAMemory, error)
	Update(*v1.HORHPAMemory) (*v1.HORHPAMemory, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.HORHPAMemory, error)
	List(opts meta_v1.ListOptions) (*v1.HORHPAMemoryList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HORHPAMemory, err error)
	HORHPAMemoryExpansion
}

// hORHPAMemories implements HORHPAMemoryInterface
type hORHPAMemories struct {
	client rest.Interface
	ns     string
}

// newHORHPAMemories returns a HORHPAMemories
func newHORHPAMemories(c *HorhpamemoryV1Client, namespace string) *hORHPAMemories {
	return &hORHPAMemories{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hORHPAMemory, and returns the corresponding hORHPAMemory object, and an error if there is any.
func (c *hORHPAMemories) Get(name string, options meta_v1.GetOptions) (result *v1.HORHPAMemory, err error) {
	result = &v1.HORHPAMemory{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("horhpamemories").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HORHPAMemories that match those selectors.
func (c *hORHPAMemories) List(opts meta_v1.ListOptions) (result *v1.HORHPAMemoryList, err error) {
	result = &v1.HORHPAMemoryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("horhpamemories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hORHPAMemories.
func (c *hORHPAMemories) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("horhpamemories").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hORHPAMemory and creates it.  Returns the server's representation of the hORHPAMemory, and an error, if there is any.
func (c *hORHPAMemories) Create(hORHPAMemory *v1.HORHPAMemory) (result *v1.HORHPAMemory, err error) {
	result = &v1.HORHPAMemory{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("horhpamemories").
		Body(hORHPAMemory).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hORHPAMemory and updates it. Returns the server's representation of the hORHPAMemory, and an error, if there is any.
func (c *hORHPAMemories) Update(hORHPAMemory *v1.HORHPAMemory) (result *v1.HORHPAMemory, err error) {
	result = &v1.HORHPAMemory{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("horhpamemories").
		Name(hORHPAMemory.Name).
		Body(hORHPAMemory).
		Do().
		Into(result)
	return
}

// Delete takes name of the hORHPAMemory and deletes it. Returns an error if one occurs.
func (c *hORHPAMemories) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("horhpamemories").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hORHPAMemories) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("horhpamemories").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hORHPAMemory.
func (c *hORHPAMemories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.HORHPAMemory, err error) {
	result = &v1.HORHPAMemory{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("horhpamemories").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
