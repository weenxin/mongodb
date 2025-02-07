/*
Copyright 2019 The KubeDB Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "kubedb.dev/apimachinery/apis/catalog/v1alpha1"
	scheme "kubedb.dev/apimachinery/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RedisVersionsGetter has a method to return a RedisVersionInterface.
// A group's client should implement this interface.
type RedisVersionsGetter interface {
	RedisVersions() RedisVersionInterface
}

// RedisVersionInterface has methods to work with RedisVersion resources.
type RedisVersionInterface interface {
	Create(*v1alpha1.RedisVersion) (*v1alpha1.RedisVersion, error)
	Update(*v1alpha1.RedisVersion) (*v1alpha1.RedisVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RedisVersion, error)
	List(opts v1.ListOptions) (*v1alpha1.RedisVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisVersion, err error)
	RedisVersionExpansion
}

// redisVersions implements RedisVersionInterface
type redisVersions struct {
	client rest.Interface
}

// newRedisVersions returns a RedisVersions
func newRedisVersions(c *CatalogV1alpha1Client) *redisVersions {
	return &redisVersions{
		client: c.RESTClient(),
	}
}

// Get takes name of the redisVersion, and returns the corresponding redisVersion object, and an error if there is any.
func (c *redisVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.RedisVersion, err error) {
	result = &v1alpha1.RedisVersion{}
	err = c.client.Get().
		Resource("redisversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RedisVersions that match those selectors.
func (c *redisVersions) List(opts v1.ListOptions) (result *v1alpha1.RedisVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RedisVersionList{}
	err = c.client.Get().
		Resource("redisversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested redisVersions.
func (c *redisVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("redisversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a redisVersion and creates it.  Returns the server's representation of the redisVersion, and an error, if there is any.
func (c *redisVersions) Create(redisVersion *v1alpha1.RedisVersion) (result *v1alpha1.RedisVersion, err error) {
	result = &v1alpha1.RedisVersion{}
	err = c.client.Post().
		Resource("redisversions").
		Body(redisVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a redisVersion and updates it. Returns the server's representation of the redisVersion, and an error, if there is any.
func (c *redisVersions) Update(redisVersion *v1alpha1.RedisVersion) (result *v1alpha1.RedisVersion, err error) {
	result = &v1alpha1.RedisVersion{}
	err = c.client.Put().
		Resource("redisversions").
		Name(redisVersion.Name).
		Body(redisVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the redisVersion and deletes it. Returns an error if one occurs.
func (c *redisVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("redisversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *redisVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("redisversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched redisVersion.
func (c *redisVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RedisVersion, err error) {
	result = &v1alpha1.RedisVersion{}
	err = c.client.Patch(pt).
		Resource("redisversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
