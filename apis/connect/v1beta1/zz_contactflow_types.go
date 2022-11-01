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

type ContactFlowObservation struct {

	// The Amazon Resource Name (ARN) of the Contact Flow.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The identifier of the Contact Flow.
	ContactFlowID *string `json:"contactFlowId,omitempty" tf:"contact_flow_id,omitempty"`

	// The identifier of the hosting Amazon Connect Instance and identifier of the Contact Flow separated by a colon (:).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type ContactFlowParameters struct {

	// Specifies the content of the Contact Flow, provided as a JSON string, written in Amazon Connect Contact Flow Language. If defined, the filename argument cannot be used.
	// +kubebuilder:validation:Optional
	Content *string `json:"content,omitempty" tf:"content,omitempty"`

	// Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the Contact Flow source specified with filename. The usual way to set this is filebase64sha256("mycontact_flow.11.12 and later) or base64sha256(file("mycontact_flow.11.11 and earlier), where "mycontact_flow.json" is the local filename of the Contact Flow source.
	// +kubebuilder:validation:Optional
	ContentHash *string `json:"contentHash,omitempty" tf:"content_hash,omitempty"`

	// Specifies the description of the Contact Flow.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The path to the Contact Flow source within the local filesystem. Conflicts with content.
	// +kubebuilder:validation:Optional
	Filename *string `json:"filename,omitempty" tf:"filename,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/connect/v1beta1.Instance
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	InstanceID *string `json:"instanceId,omitempty" tf:"instance_id,omitempty"`

	// Reference to a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDRef *v1.Reference `json:"instanceIdRef,omitempty" tf:"-"`

	// Selector for a Instance in connect to populate instanceId.
	// +kubebuilder:validation:Optional
	InstanceIDSelector *v1.Selector `json:"instanceIdSelector,omitempty" tf:"-"`

	// Specifies the name of the Contact Flow.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// Tags to apply to the Contact Flow. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Specifies the type of the Contact Flow. Defaults to CONTACT_FLOW. Allowed Values are: CONTACT_FLOW, CUSTOMER_QUEUE, CUSTOMER_HOLD, CUSTOMER_WHISPER, AGENT_HOLD, AGENT_WHISPER, OUTBOUND_WHISPER, AGENT_TRANSFER, QUEUE_TRANSFER.
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// ContactFlowSpec defines the desired state of ContactFlow
type ContactFlowSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ContactFlowParameters `json:"forProvider"`
}

// ContactFlowStatus defines the observed state of ContactFlow.
type ContactFlowStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ContactFlowObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ContactFlow is the Schema for the ContactFlows API. Provides details about a specific Amazon Connect Contact Flow.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type ContactFlow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ContactFlowSpec   `json:"spec"`
	Status            ContactFlowStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContactFlowList contains a list of ContactFlows
type ContactFlowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContactFlow `json:"items"`
}

// Repository type metadata.
var (
	ContactFlow_Kind             = "ContactFlow"
	ContactFlow_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ContactFlow_Kind}.String()
	ContactFlow_KindAPIVersion   = ContactFlow_Kind + "." + CRDGroupVersion.String()
	ContactFlow_GroupVersionKind = CRDGroupVersion.WithKind(ContactFlow_Kind)
)

func init() {
	SchemeBuilder.Register(&ContactFlow{}, &ContactFlowList{})
}
