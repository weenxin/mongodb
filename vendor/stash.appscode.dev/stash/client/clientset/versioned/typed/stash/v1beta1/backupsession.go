/*
Copyright 2019 The Stash Authors.

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

	v1beta1 "stash.appscode.dev/stash/apis/stash/v1beta1"
	scheme "stash.appscode.dev/stash/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BackupSessionsGetter has a method to return a BackupSessionInterface.
// A group's client should implement this interface.
type BackupSessionsGetter interface {
	BackupSessions(namespace string) BackupSessionInterface
}

// BackupSessionInterface has methods to work with BackupSession resources.
type BackupSessionInterface interface {
	Create(*v1beta1.BackupSession) (*v1beta1.BackupSession, error)
	Update(*v1beta1.BackupSession) (*v1beta1.BackupSession, error)
	UpdateStatus(*v1beta1.BackupSession) (*v1beta1.BackupSession, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1beta1.BackupSession, error)
	List(opts v1.ListOptions) (*v1beta1.BackupSessionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackupSession, err error)
	BackupSessionExpansion
}

// backupSessions implements BackupSessionInterface
type backupSessions struct {
	client rest.Interface
	ns     string
}

// newBackupSessions returns a BackupSessions
func newBackupSessions(c *StashV1beta1Client, namespace string) *backupSessions {
	return &backupSessions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the backupSession, and returns the corresponding backupSession object, and an error if there is any.
func (c *backupSessions) Get(name string, options v1.GetOptions) (result *v1beta1.BackupSession, err error) {
	result = &v1beta1.BackupSession{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupsessions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of BackupSessions that match those selectors.
func (c *backupSessions) List(opts v1.ListOptions) (result *v1beta1.BackupSessionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.BackupSessionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("backupsessions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested backupSessions.
func (c *backupSessions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("backupsessions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a backupSession and creates it.  Returns the server's representation of the backupSession, and an error, if there is any.
func (c *backupSessions) Create(backupSession *v1beta1.BackupSession) (result *v1beta1.BackupSession, err error) {
	result = &v1beta1.BackupSession{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("backupsessions").
		Body(backupSession).
		Do().
		Into(result)
	return
}

// Update takes the representation of a backupSession and updates it. Returns the server's representation of the backupSession, and an error, if there is any.
func (c *backupSessions) Update(backupSession *v1beta1.BackupSession) (result *v1beta1.BackupSession, err error) {
	result = &v1beta1.BackupSession{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupsessions").
		Name(backupSession.Name).
		Body(backupSession).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *backupSessions) UpdateStatus(backupSession *v1beta1.BackupSession) (result *v1beta1.BackupSession, err error) {
	result = &v1beta1.BackupSession{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("backupsessions").
		Name(backupSession.Name).
		SubResource("status").
		Body(backupSession).
		Do().
		Into(result)
	return
}

// Delete takes name of the backupSession and deletes it. Returns an error if one occurs.
func (c *backupSessions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupsessions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *backupSessions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("backupsessions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched backupSession.
func (c *backupSessions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.BackupSession, err error) {
	result = &v1beta1.BackupSession{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("backupsessions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
