/*
Copyright The Kubernetes Authors.

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

package fake

import (
	"istio.io/istio/pkg/servicemesh/apis/servicemesh/v1alpha3"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/testing"
)

// FakeServiceMeshMemberRolls implements ServiceMeshMemberRollInterface
type FakeServiceMeshMemberRolls struct {
	Fake *FakeIstioV1alpha3
	ns   string
}

var servicemeshmemberrollsResource = schema.GroupVersionResource{Group: "istio.openshift.com", Version: "v1alpha3", Resource: "servicemeshmemberrolls"}

var servicemeshmemberrollsKind = schema.GroupVersionKind{Group: "istio.openshift.com", Version: "v1alpha3", Kind: "ServiceMeshMemberRoll"}

// Get takes name of the serviceMeshMemberRoll, and returns the corresponding serviceMeshMemberRoll object, and an error if there is any.
func (c *FakeServiceMeshMemberRolls) Get(name string, options v1.GetOptions) (result *v1alpha3.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicemeshmemberrollsResource, c.ns, name), &v1alpha3.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceMeshMemberRoll), err
}

// List takes label and field selectors, and returns the list of ServiceMeshMemberRolls that match those selectors.
func (c *FakeServiceMeshMemberRolls) List(opts v1.ListOptions) (result *v1alpha3.ServiceMeshMemberRollList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicemeshmemberrollsResource, servicemeshmemberrollsKind, c.ns, opts), &v1alpha3.ServiceMeshMemberRollList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.ServiceMeshMemberRollList{ListMeta: obj.(*v1alpha3.ServiceMeshMemberRollList).ListMeta}
	for _, item := range obj.(*v1alpha3.ServiceMeshMemberRollList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceMeshMemberRolls.
func (c *FakeServiceMeshMemberRolls) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicemeshmemberrollsResource, c.ns, opts))

}

// Create takes the representation of a serviceMeshMemberRoll and creates it.  Returns the server's representation of the serviceMeshMemberRoll, and an error, if there is any.
func (c *FakeServiceMeshMemberRolls) Create(serviceMeshMemberRoll *v1alpha3.ServiceMeshMemberRoll) (result *v1alpha3.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicemeshmemberrollsResource, c.ns, serviceMeshMemberRoll), &v1alpha3.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceMeshMemberRoll), err
}

// Update takes the representation of a serviceMeshMemberRoll and updates it. Returns the server's representation of the serviceMeshMemberRoll, and an error, if there is any.
func (c *FakeServiceMeshMemberRolls) Update(serviceMeshMemberRoll *v1alpha3.ServiceMeshMemberRoll) (result *v1alpha3.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicemeshmemberrollsResource, c.ns, serviceMeshMemberRoll), &v1alpha3.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceMeshMemberRoll), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceMeshMemberRolls) UpdateStatus(serviceMeshMemberRoll *v1alpha3.ServiceMeshMemberRoll) (*v1alpha3.ServiceMeshMemberRoll, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicemeshmemberrollsResource, "status", c.ns, serviceMeshMemberRoll), &v1alpha3.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceMeshMemberRoll), err
}

// Delete takes name of the serviceMeshMemberRoll and deletes it. Returns an error if one occurs.
func (c *FakeServiceMeshMemberRolls) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(servicemeshmemberrollsResource, c.ns, name), &v1alpha3.ServiceMeshMemberRoll{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceMeshMemberRolls) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicemeshmemberrollsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha3.ServiceMeshMemberRollList{})
	return err
}

// Patch applies the patch and returns the patched serviceMeshMemberRoll.
func (c *FakeServiceMeshMemberRolls) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha3.ServiceMeshMemberRoll, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicemeshmemberrollsResource, c.ns, name, data, subresources...), &v1alpha3.ServiceMeshMemberRoll{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.ServiceMeshMemberRoll), err
}
