/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TransitGatewayMulticastDomainObservation struct {

	// EC2 Transit Gateway Multicast Domain Amazon Resource Name (ARN).
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// EC2 Transit Gateway Multicast Domain identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Identifier of the AWS account that owns the EC2 Transit Gateway Multicast Domain.
	OwnerID *string `json:"ownerId,omitempty" tf:"owner_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type TransitGatewayMulticastDomainParameters struct {

	// Whether to automatically accept cross-account subnet associations that are associated with the EC2 Transit Gateway Multicast Domain. Valid values: disable, enable. Default value: disable.
	// +kubebuilder:validation:Optional
	AutoAcceptSharedAssociations *string `json:"autoAcceptSharedAssociations,omitempty" tf:"auto_accept_shared_associations,omitempty"`

	// Whether to enable Internet Group Management Protocol (IGMP) version 2 for the EC2 Transit Gateway Multicast Domain. Valid values: disable, enable. Default value: disable.
	// +kubebuilder:validation:Optional
	Igmpv2Support *string `json:"igmpv2Support,omitempty" tf:"igmpv2_support,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Whether to enable support for statically configuring multicast group sources for the EC2 Transit Gateway Multicast Domain. Valid values: disable, enable. Default value: disable.
	// +kubebuilder:validation:Optional
	StaticSourcesSupport *string `json:"staticSourcesSupport,omitempty" tf:"static_sources_support,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// EC2 Transit Gateway identifier. The EC2 Transit Gateway must have multicast_support enabled.
	// +crossplane:generate:reference:type=TransitGateway
	// +kubebuilder:validation:Optional
	TransitGatewayID *string `json:"transitGatewayId,omitempty" tf:"transit_gateway_id,omitempty"`

	// Reference to a TransitGateway to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDRef *v1.Reference `json:"transitGatewayIdRef,omitempty" tf:"-"`

	// Selector for a TransitGateway to populate transitGatewayId.
	// +kubebuilder:validation:Optional
	TransitGatewayIDSelector *v1.Selector `json:"transitGatewayIdSelector,omitempty" tf:"-"`
}

// TransitGatewayMulticastDomainSpec defines the desired state of TransitGatewayMulticastDomain
type TransitGatewayMulticastDomainSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitGatewayMulticastDomainParameters `json:"forProvider"`
}

// TransitGatewayMulticastDomainStatus defines the observed state of TransitGatewayMulticastDomain.
type TransitGatewayMulticastDomainStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitGatewayMulticastDomainObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayMulticastDomain is the Schema for the TransitGatewayMulticastDomains API. Manages an EC2 Transit Gateway Multicast Domain
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type TransitGatewayMulticastDomain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransitGatewayMulticastDomainSpec   `json:"spec"`
	Status            TransitGatewayMulticastDomainStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitGatewayMulticastDomainList contains a list of TransitGatewayMulticastDomains
type TransitGatewayMulticastDomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitGatewayMulticastDomain `json:"items"`
}

// Repository type metadata.
var (
	TransitGatewayMulticastDomain_Kind             = "TransitGatewayMulticastDomain"
	TransitGatewayMulticastDomain_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitGatewayMulticastDomain_Kind}.String()
	TransitGatewayMulticastDomain_KindAPIVersion   = TransitGatewayMulticastDomain_Kind + "." + CRDGroupVersion.String()
	TransitGatewayMulticastDomain_GroupVersionKind = CRDGroupVersion.WithKind(TransitGatewayMulticastDomain_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitGatewayMulticastDomain{}, &TransitGatewayMulticastDomainList{})
}
