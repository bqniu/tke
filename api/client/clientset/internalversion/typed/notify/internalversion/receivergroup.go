/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	notify "tkestack.io/tke/api/notify"
)

// ReceiverGroupsGetter has a method to return a ReceiverGroupInterface.
// A group's client should implement this interface.
type ReceiverGroupsGetter interface {
	ReceiverGroups() ReceiverGroupInterface
}

// ReceiverGroupInterface has methods to work with ReceiverGroup resources.
type ReceiverGroupInterface interface {
	Create(*notify.ReceiverGroup) (*notify.ReceiverGroup, error)
	Update(*notify.ReceiverGroup) (*notify.ReceiverGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	Get(name string, options v1.GetOptions) (*notify.ReceiverGroup, error)
	List(opts v1.ListOptions) (*notify.ReceiverGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.ReceiverGroup, err error)
	ReceiverGroupExpansion
}

// receiverGroups implements ReceiverGroupInterface
type receiverGroups struct {
	client rest.Interface
}

// newReceiverGroups returns a ReceiverGroups
func newReceiverGroups(c *NotifyClient) *receiverGroups {
	return &receiverGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the receiverGroup, and returns the corresponding receiverGroup object, and an error if there is any.
func (c *receiverGroups) Get(name string, options v1.GetOptions) (result *notify.ReceiverGroup, err error) {
	result = &notify.ReceiverGroup{}
	err = c.client.Get().
		Resource("receivergroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ReceiverGroups that match those selectors.
func (c *receiverGroups) List(opts v1.ListOptions) (result *notify.ReceiverGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &notify.ReceiverGroupList{}
	err = c.client.Get().
		Resource("receivergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested receiverGroups.
func (c *receiverGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("receivergroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a receiverGroup and creates it.  Returns the server's representation of the receiverGroup, and an error, if there is any.
func (c *receiverGroups) Create(receiverGroup *notify.ReceiverGroup) (result *notify.ReceiverGroup, err error) {
	result = &notify.ReceiverGroup{}
	err = c.client.Post().
		Resource("receivergroups").
		Body(receiverGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a receiverGroup and updates it. Returns the server's representation of the receiverGroup, and an error, if there is any.
func (c *receiverGroups) Update(receiverGroup *notify.ReceiverGroup) (result *notify.ReceiverGroup, err error) {
	result = &notify.ReceiverGroup{}
	err = c.client.Put().
		Resource("receivergroups").
		Name(receiverGroup.Name).
		Body(receiverGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the receiverGroup and deletes it. Returns an error if one occurs.
func (c *receiverGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("receivergroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched receiverGroup.
func (c *receiverGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *notify.ReceiverGroup, err error) {
	result = &notify.ReceiverGroup{}
	err = c.client.Patch(pt).
		Resource("receivergroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}