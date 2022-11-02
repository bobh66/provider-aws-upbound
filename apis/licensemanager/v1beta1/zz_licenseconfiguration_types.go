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

type LicenseConfigurationObservation struct {

	// The license configuration ARN.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The license configuration ARN.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Account ID of the owner of the license configuration.
	OwnerAccountID *string `json:"ownerAccountId,omitempty" tf:"owner_account_id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type LicenseConfigurationParameters struct {

	// Description of the license configuration.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Number of licenses managed by the license configuration.
	// +kubebuilder:validation:Optional
	LicenseCount *float64 `json:"licenseCount,omitempty" tf:"license_count,omitempty"`

	// Sets the number of available licenses as a hard limit.
	// +kubebuilder:validation:Optional
	LicenseCountHardLimit *bool `json:"licenseCountHardLimit,omitempty" tf:"license_count_hard_limit,omitempty"`

	// Dimension to use to track license inventory. Specify either vCPU, Instance, Core or Socket.
	// +kubebuilder:validation:Required
	LicenseCountingType *string `json:"licenseCountingType" tf:"license_counting_type,omitempty"`

	// Array of configured License Manager rules.
	// +kubebuilder:validation:Optional
	LicenseRules []*string `json:"licenseRules,omitempty" tf:"license_rules,omitempty"`

	// Name of the license configuration.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// LicenseConfigurationSpec defines the desired state of LicenseConfiguration
type LicenseConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LicenseConfigurationParameters `json:"forProvider"`
}

// LicenseConfigurationStatus defines the observed state of LicenseConfiguration.
type LicenseConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LicenseConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseConfiguration is the Schema for the LicenseConfigurations API. Provides a License Manager license configuration resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LicenseConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LicenseConfigurationSpec   `json:"spec"`
	Status            LicenseConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LicenseConfigurationList contains a list of LicenseConfigurations
type LicenseConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LicenseConfiguration `json:"items"`
}

// Repository type metadata.
var (
	LicenseConfiguration_Kind             = "LicenseConfiguration"
	LicenseConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LicenseConfiguration_Kind}.String()
	LicenseConfiguration_KindAPIVersion   = LicenseConfiguration_Kind + "." + CRDGroupVersion.String()
	LicenseConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(LicenseConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&LicenseConfiguration{}, &LicenseConfigurationList{})
}
