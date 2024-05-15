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

type LinkAssociationInitParameters struct {
}

type LinkAssociationObservation struct {

	// The ID of the device.
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// The ID of the global network.
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the link.
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`
}

type LinkAssociationParameters struct {

	// The ID of the device.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.Device
	// +kubebuilder:validation:Optional
	DeviceID *string `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// Reference to a Device in networkmanager to populate deviceId.
	// +kubebuilder:validation:Optional
	DeviceIDRef *v1.Reference `json:"deviceIdRef,omitempty" tf:"-"`

	// Selector for a Device in networkmanager to populate deviceId.
	// +kubebuilder:validation:Optional
	DeviceIDSelector *v1.Selector `json:"deviceIdSelector,omitempty" tf:"-"`

	// The ID of the global network.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.GlobalNetwork
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GlobalNetworkID *string `json:"globalNetworkId,omitempty" tf:"global_network_id,omitempty"`

	// Reference to a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDRef *v1.Reference `json:"globalNetworkIdRef,omitempty" tf:"-"`

	// Selector for a GlobalNetwork in networkmanager to populate globalNetworkId.
	// +kubebuilder:validation:Optional
	GlobalNetworkIDSelector *v1.Selector `json:"globalNetworkIdSelector,omitempty" tf:"-"`

	// The ID of the link.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/networkmanager/v1beta1.Link
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	LinkID *string `json:"linkId,omitempty" tf:"link_id,omitempty"`

	// Reference to a Link in networkmanager to populate linkId.
	// +kubebuilder:validation:Optional
	LinkIDRef *v1.Reference `json:"linkIdRef,omitempty" tf:"-"`

	// Selector for a Link in networkmanager to populate linkId.
	// +kubebuilder:validation:Optional
	LinkIDSelector *v1.Selector `json:"linkIdSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LinkAssociationSpec defines the desired state of LinkAssociation
type LinkAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LinkAssociationParameters `json:"forProvider"`
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
	InitProvider LinkAssociationInitParameters `json:"initProvider,omitempty"`
}

// LinkAssociationStatus defines the observed state of LinkAssociation.
type LinkAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LinkAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// LinkAssociation is the Schema for the LinkAssociations API. Associates a link to a device.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LinkAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LinkAssociationSpec   `json:"spec"`
	Status            LinkAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinkAssociationList contains a list of LinkAssociations
type LinkAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinkAssociation `json:"items"`
}

// Repository type metadata.
var (
	LinkAssociation_Kind             = "LinkAssociation"
	LinkAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LinkAssociation_Kind}.String()
	LinkAssociation_KindAPIVersion   = LinkAssociation_Kind + "." + CRDGroupVersion.String()
	LinkAssociation_GroupVersionKind = CRDGroupVersion.WithKind(LinkAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&LinkAssociation{}, &LinkAssociationList{})
}
