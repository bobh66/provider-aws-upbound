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

type VPCAssociationAuthorizationInitParameters struct {

	// The VPC to authorize for association with the private hosted zone.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The VPC's region. Defaults to the region of the AWS provider.
	VPCRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`

	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53/v1beta1.Zone
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

type VPCAssociationAuthorizationObservation struct {

	// The calculated unique identifier for the association.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The VPC to authorize for association with the private hosted zone.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// The VPC's region. Defaults to the region of the AWS provider.
	VPCRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`

	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type VPCAssociationAuthorizationParameters struct {

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The VPC to authorize for association with the private hosted zone.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/ec2/v1beta1.VPC
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// Reference to a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// Selector for a VPC in ec2 to populate vpcId.
	// +kubebuilder:validation:Optional
	VPCIDSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`

	// The VPC's region. Defaults to the region of the AWS provider.
	// +kubebuilder:validation:Optional
	VPCRegion *string `json:"vpcRegion,omitempty" tf:"vpc_region,omitempty"`

	// The ID of the private hosted zone that you want to authorize associating a VPC with.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/route53/v1beta1.Zone
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`

	// Reference to a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDRef *v1.Reference `json:"zoneIdRef,omitempty" tf:"-"`

	// Selector for a Zone in route53 to populate zoneId.
	// +kubebuilder:validation:Optional
	ZoneIDSelector *v1.Selector `json:"zoneIdSelector,omitempty" tf:"-"`
}

// VPCAssociationAuthorizationSpec defines the desired state of VPCAssociationAuthorization
type VPCAssociationAuthorizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     VPCAssociationAuthorizationParameters `json:"forProvider"`
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
	InitProvider VPCAssociationAuthorizationInitParameters `json:"initProvider,omitempty"`
}

// VPCAssociationAuthorizationStatus defines the observed state of VPCAssociationAuthorization.
type VPCAssociationAuthorizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        VPCAssociationAuthorizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// VPCAssociationAuthorization is the Schema for the VPCAssociationAuthorizations API. Authorizes a VPC in a different account to be associated with a local Route53 Hosted Zone
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type VPCAssociationAuthorization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VPCAssociationAuthorizationSpec   `json:"spec"`
	Status            VPCAssociationAuthorizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VPCAssociationAuthorizationList contains a list of VPCAssociationAuthorizations
type VPCAssociationAuthorizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VPCAssociationAuthorization `json:"items"`
}

// Repository type metadata.
var (
	VPCAssociationAuthorization_Kind             = "VPCAssociationAuthorization"
	VPCAssociationAuthorization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: VPCAssociationAuthorization_Kind}.String()
	VPCAssociationAuthorization_KindAPIVersion   = VPCAssociationAuthorization_Kind + "." + CRDGroupVersion.String()
	VPCAssociationAuthorization_GroupVersionKind = CRDGroupVersion.WithKind(VPCAssociationAuthorization_Kind)
)

func init() {
	SchemeBuilder.Register(&VPCAssociationAuthorization{}, &VPCAssociationAuthorizationList{})
}
