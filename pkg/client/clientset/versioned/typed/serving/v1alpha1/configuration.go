/*
Copyright 2020 The Knative Authors

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "knative.dev/serving/pkg/apis/serving/v1alpha1"
	scheme "knative.dev/serving/pkg/client/clientset/versioned/scheme"
)

// ConfigurationsGetter has a method to return a ConfigurationInterface.
// A group's client should implement this interface.
type ConfigurationsGetter interface {
	Configurations(namespace string) ConfigurationInterface
}

// ConfigurationInterface has methods to work with Configuration resources.
type ConfigurationInterface interface {
	Create(ctx context.Context, *v1alpha1.Configuration) (*v1alpha1.Configuration, error)
	Update(ctx context.Context, *v1alpha1.Configuration) (*v1alpha1.Configuration, error)
	UpdateStatus(ctx context.Context, *v1alpha1.Configuration) (*v1alpha1.Configuration, error)
	Delete(ctx context.Context, name string, options *v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(ctx context.Context, name string, options v1.GetOptions) (*v1alpha1.Configuration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Configuration, err error)
	ConfigurationExpansion
}

// configurations implements ConfigurationInterface
type configurations struct {
	client rest.Interface
	ns     string
}

// newConfigurations returns a Configurations
func newConfigurations(c *ServingV1alpha1Client, namespace string) *configurations {
	return &configurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the configuration, and returns the corresponding configuration object, and an error if there is any.
func (c *configurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Configuration, err error) {
	result = &v1alpha1.Configuration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Configurations that match those selectors.
func (c *configurations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("configurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested configurations.
func (c *configurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("configurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a configuration and creates it.  Returns the server's representation of the configuration, and an error, if there is any.
func (c *configurations) Create(ctx context.Context, configuration *v1alpha1.Configuration) (result *v1alpha1.Configuration, err error) {
	result = &v1alpha1.Configuration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("configurations").
		Body(configuration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a configuration and updates it. Returns the server's representation of the configuration, and an error, if there is any.
func (c *configurations) Update(ctx context.Context, configuration *v1alpha1.Configuration) (result *v1alpha1.Configuration, err error) {
	result = &v1alpha1.Configuration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configurations").
		Name(configuration.Name).
		Body(configuration).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *configurations) UpdateStatus(ctx context.Context, configuration *v1alpha1.Configuration) (result *v1alpha1.Configuration, err error) {
	result = &v1alpha1.Configuration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("configurations").
		Name(configuration.Name).
		SubResource("status").
		Body(configuration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the configuration and deletes it. Returns an error if one occurs.
func (c *configurations) Delete(ctx context.Context, name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configurations").
		Name(name).
		Body(options).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *configurations) DeleteCollection(ctx context.Context, options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("configurations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched configuration.
func (c *configurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Configuration, err error) {
	result = &v1alpha1.Configuration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("configurations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
