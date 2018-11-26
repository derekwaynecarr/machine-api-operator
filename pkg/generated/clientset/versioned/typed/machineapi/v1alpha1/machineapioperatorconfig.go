// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/machine-api-operator/pkg/apis/machineapi/v1alpha1"
	scheme "github.com/openshift/machine-api-operator/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineAPIOperatorConfigsGetter has a method to return a MachineAPIOperatorConfigInterface.
// A group's client should implement this interface.
type MachineAPIOperatorConfigsGetter interface {
	MachineAPIOperatorConfigs() MachineAPIOperatorConfigInterface
}

// MachineAPIOperatorConfigInterface has methods to work with MachineAPIOperatorConfig resources.
type MachineAPIOperatorConfigInterface interface {
	Create(*v1alpha1.MachineAPIOperatorConfig) (*v1alpha1.MachineAPIOperatorConfig, error)
	Update(*v1alpha1.MachineAPIOperatorConfig) (*v1alpha1.MachineAPIOperatorConfig, error)
	UpdateStatus(*v1alpha1.MachineAPIOperatorConfig) (*v1alpha1.MachineAPIOperatorConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.MachineAPIOperatorConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.MachineAPIOperatorConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineAPIOperatorConfig, err error)
	MachineAPIOperatorConfigExpansion
}

// machineAPIOperatorConfigs implements MachineAPIOperatorConfigInterface
type machineAPIOperatorConfigs struct {
	client rest.Interface
}

// newMachineAPIOperatorConfigs returns a MachineAPIOperatorConfigs
func newMachineAPIOperatorConfigs(c *MachineapiV1alpha1Client) *machineAPIOperatorConfigs {
	return &machineAPIOperatorConfigs{
		client: c.RESTClient(),
	}
}

// Get takes name of the machineAPIOperatorConfig, and returns the corresponding machineAPIOperatorConfig object, and an error if there is any.
func (c *machineAPIOperatorConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.MachineAPIOperatorConfig, err error) {
	result = &v1alpha1.MachineAPIOperatorConfig{}
	err = c.client.Get().
		Resource("machineapioperatorconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineAPIOperatorConfigs that match those selectors.
func (c *machineAPIOperatorConfigs) List(opts v1.ListOptions) (result *v1alpha1.MachineAPIOperatorConfigList, err error) {
	result = &v1alpha1.MachineAPIOperatorConfigList{}
	err = c.client.Get().
		Resource("machineapioperatorconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineAPIOperatorConfigs.
func (c *machineAPIOperatorConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("machineapioperatorconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a machineAPIOperatorConfig and creates it.  Returns the server's representation of the machineAPIOperatorConfig, and an error, if there is any.
func (c *machineAPIOperatorConfigs) Create(machineAPIOperatorConfig *v1alpha1.MachineAPIOperatorConfig) (result *v1alpha1.MachineAPIOperatorConfig, err error) {
	result = &v1alpha1.MachineAPIOperatorConfig{}
	err = c.client.Post().
		Resource("machineapioperatorconfigs").
		Body(machineAPIOperatorConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machineAPIOperatorConfig and updates it. Returns the server's representation of the machineAPIOperatorConfig, and an error, if there is any.
func (c *machineAPIOperatorConfigs) Update(machineAPIOperatorConfig *v1alpha1.MachineAPIOperatorConfig) (result *v1alpha1.MachineAPIOperatorConfig, err error) {
	result = &v1alpha1.MachineAPIOperatorConfig{}
	err = c.client.Put().
		Resource("machineapioperatorconfigs").
		Name(machineAPIOperatorConfig.Name).
		Body(machineAPIOperatorConfig).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *machineAPIOperatorConfigs) UpdateStatus(machineAPIOperatorConfig *v1alpha1.MachineAPIOperatorConfig) (result *v1alpha1.MachineAPIOperatorConfig, err error) {
	result = &v1alpha1.MachineAPIOperatorConfig{}
	err = c.client.Put().
		Resource("machineapioperatorconfigs").
		Name(machineAPIOperatorConfig.Name).
		SubResource("status").
		Body(machineAPIOperatorConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the machineAPIOperatorConfig and deletes it. Returns an error if one occurs.
func (c *machineAPIOperatorConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("machineapioperatorconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineAPIOperatorConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("machineapioperatorconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machineAPIOperatorConfig.
func (c *machineAPIOperatorConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MachineAPIOperatorConfig, err error) {
	result = &v1alpha1.MachineAPIOperatorConfig{}
	err = c.client.Patch(pt).
		Resource("machineapioperatorconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}