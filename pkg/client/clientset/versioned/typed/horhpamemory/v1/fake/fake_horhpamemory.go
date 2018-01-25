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
package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	horhpamemory_v1 "kube-hpamemory/pkg/apis/horhpamemory/v1"
)

// FakeHORHPAMemories implements HORHPAMemoryInterface
type FakeHORHPAMemories struct {
	Fake *FakeHorhpamemoryV1
	ns   string
}

var horhpamemoriesResource = schema.GroupVersionResource{Group: "horhpamemory", Version: "v1", Resource: "horhpamemories"}

var horhpamemoriesKind = schema.GroupVersionKind{Group: "horhpamemory", Version: "v1", Kind: "HORHPAMemory"}

// Get takes name of the hORHPAMemory, and returns the corresponding hORHPAMemory object, and an error if there is any.
func (c *FakeHORHPAMemories) Get(name string, options v1.GetOptions) (result *horhpamemory_v1.HORHPAMemory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(horhpamemoriesResource, c.ns, name), &horhpamemory_v1.HORHPAMemory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*horhpamemory_v1.HORHPAMemory), err
}

// List takes label and field selectors, and returns the list of HORHPAMemories that match those selectors.
func (c *FakeHORHPAMemories) List(opts v1.ListOptions) (result *horhpamemory_v1.HORHPAMemoryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(horhpamemoriesResource, horhpamemoriesKind, c.ns, opts), &horhpamemory_v1.HORHPAMemoryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &horhpamemory_v1.HORHPAMemoryList{}
	for _, item := range obj.(*horhpamemory_v1.HORHPAMemoryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hORHPAMemories.
func (c *FakeHORHPAMemories) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(horhpamemoriesResource, c.ns, opts))

}

// Create takes the representation of a hORHPAMemory and creates it.  Returns the server's representation of the hORHPAMemory, and an error, if there is any.
func (c *FakeHORHPAMemories) Create(hORHPAMemory *horhpamemory_v1.HORHPAMemory) (result *horhpamemory_v1.HORHPAMemory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(horhpamemoriesResource, c.ns, hORHPAMemory), &horhpamemory_v1.HORHPAMemory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*horhpamemory_v1.HORHPAMemory), err
}

// Update takes the representation of a hORHPAMemory and updates it. Returns the server's representation of the hORHPAMemory, and an error, if there is any.
func (c *FakeHORHPAMemories) Update(hORHPAMemory *horhpamemory_v1.HORHPAMemory) (result *horhpamemory_v1.HORHPAMemory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(horhpamemoriesResource, c.ns, hORHPAMemory), &horhpamemory_v1.HORHPAMemory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*horhpamemory_v1.HORHPAMemory), err
}

// Delete takes name of the hORHPAMemory and deletes it. Returns an error if one occurs.
func (c *FakeHORHPAMemories) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(horhpamemoriesResource, c.ns, name), &horhpamemory_v1.HORHPAMemory{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHORHPAMemories) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(horhpamemoriesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &horhpamemory_v1.HORHPAMemoryList{})
	return err
}

// Patch applies the patch and returns the patched hORHPAMemory.
func (c *FakeHORHPAMemories) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *horhpamemory_v1.HORHPAMemory, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(horhpamemoriesResource, c.ns, name, data, subresources...), &horhpamemory_v1.HORHPAMemory{})

	if obj == nil {
		return nil, err
	}
	return obj.(*horhpamemory_v1.HORHPAMemory), err
}
