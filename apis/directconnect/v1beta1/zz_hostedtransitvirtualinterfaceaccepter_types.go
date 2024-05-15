// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type HostedTransitVirtualInterfaceAccepterInitParameters struct {

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.Gateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// Reference to a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDRef *v1.Reference `json:"dxGatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDSelector *v1.Selector `json:"dxGatewayIdSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Direct Connect virtual interface to accept.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.HostedTransitVirtualInterface
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`

	// Reference to a HostedTransitVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDRef *v1.Reference `json:"virtualInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a HostedTransitVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDSelector *v1.Selector `json:"virtualInterfaceIdSelector,omitempty" tf:"-"`
}

type HostedTransitVirtualInterfaceAccepterObservation struct {

	// The ARN of the virtual interface.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// The ID of the virtual interface.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Key-value map of resource tags.
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	// +mapType=granular
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The ID of the Direct Connect virtual interface to accept.
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`
}

type HostedTransitVirtualInterfaceAccepterParameters struct {

	// The ID of the Direct Connect gateway to which to connect the virtual interface.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.Gateway
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	DxGatewayID *string `json:"dxGatewayId,omitempty" tf:"dx_gateway_id,omitempty"`

	// Reference to a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDRef *v1.Reference `json:"dxGatewayIdRef,omitempty" tf:"-"`

	// Selector for a Gateway in directconnect to populate dxGatewayId.
	// +kubebuilder:validation:Optional
	DxGatewayIDSelector *v1.Selector `json:"dxGatewayIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID of the Direct Connect virtual interface to accept.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/directconnect/v1beta1.HostedTransitVirtualInterface
	// +kubebuilder:validation:Optional
	VirtualInterfaceID *string `json:"virtualInterfaceId,omitempty" tf:"virtual_interface_id,omitempty"`

	// Reference to a HostedTransitVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDRef *v1.Reference `json:"virtualInterfaceIdRef,omitempty" tf:"-"`

	// Selector for a HostedTransitVirtualInterface in directconnect to populate virtualInterfaceId.
	// +kubebuilder:validation:Optional
	VirtualInterfaceIDSelector *v1.Selector `json:"virtualInterfaceIdSelector,omitempty" tf:"-"`
}

// HostedTransitVirtualInterfaceAccepterSpec defines the desired state of HostedTransitVirtualInterfaceAccepter
type HostedTransitVirtualInterfaceAccepterSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     HostedTransitVirtualInterfaceAccepterParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider HostedTransitVirtualInterfaceAccepterInitParameters `json:"initProvider,omitempty"`
}

// HostedTransitVirtualInterfaceAccepterStatus defines the observed state of HostedTransitVirtualInterfaceAccepter.
type HostedTransitVirtualInterfaceAccepterStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        HostedTransitVirtualInterfaceAccepterObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// HostedTransitVirtualInterfaceAccepter is the Schema for the HostedTransitVirtualInterfaceAccepters API. Provides a resource to manage the accepter's side of a Direct Connect hosted transit virtual interface.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type HostedTransitVirtualInterfaceAccepter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HostedTransitVirtualInterfaceAccepterSpec   `json:"spec"`
	Status            HostedTransitVirtualInterfaceAccepterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HostedTransitVirtualInterfaceAccepterList contains a list of HostedTransitVirtualInterfaceAccepters
type HostedTransitVirtualInterfaceAccepterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HostedTransitVirtualInterfaceAccepter `json:"items"`
}

// Repository type metadata.
var (
	HostedTransitVirtualInterfaceAccepter_Kind             = "HostedTransitVirtualInterfaceAccepter"
	HostedTransitVirtualInterfaceAccepter_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: HostedTransitVirtualInterfaceAccepter_Kind}.String()
	HostedTransitVirtualInterfaceAccepter_KindAPIVersion   = HostedTransitVirtualInterfaceAccepter_Kind + "." + CRDGroupVersion.String()
	HostedTransitVirtualInterfaceAccepter_GroupVersionKind = CRDGroupVersion.WithKind(HostedTransitVirtualInterfaceAccepter_Kind)
)

func init() {
	SchemeBuilder.Register(&HostedTransitVirtualInterfaceAccepter{}, &HostedTransitVirtualInterfaceAccepterList{})
}
