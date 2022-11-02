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

type ControlObservation struct {
}

type ControlParameters struct {

	// One or more input parameter blocks. An example of a control with two parameters is: "backup plan frequency is at least daily and the retention period is at least 1 year". The first parameter is daily. The second parameter is 1 year. Detailed below.
	// +kubebuilder:validation:Optional
	InputParameter []InputParameterParameters `json:"inputParameter,omitempty" tf:"input_parameter,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The scope of a control. The control scope defines what the control will evaluate. Three examples of control scopes are: a specific backup plan, all backup plans with a specific tag, or all backup plans. Detailed below.
	// +kubebuilder:validation:Optional
	Scope []ScopeParameters `json:"scope,omitempty" tf:"scope,omitempty"`
}

type FrameworkObservation struct {

	// The ARN of the backup framework.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// The date and time that a framework is created, in Unix format and Coordinated Universal Time (UTC).
	CreationTime *string `json:"creationTime,omitempty" tf:"creation_time,omitempty"`

	// The deployment status of a framework. The statuses are: CREATE_IN_PROGRESS | UPDATE_IN_PROGRESS | DELETE_IN_PROGRESS | COMPLETED | FAILED.
	DeploymentStatus *string `json:"deploymentStatus,omitempty" tf:"deployment_status,omitempty"`

	// The id of the backup framework.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A framework consists of one or more controls. Each control governs a resource, such as backup plans, backup selections, backup vaults, or recovery points. You can also turn AWS Config recording on or off for each resource. For more information refer to the AWS documentation for Framework Status
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FrameworkParameters struct {

	// One or more control blocks that make up the framework. Each control in the list has a name, input parameters, and scope. Detailed below.
	// +kubebuilder:validation:Required
	Control []ControlParameters `json:"control" tf:"control,omitempty"`

	// The description of the framework with a maximum of 1,024 characters
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
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

type InputParameterObservation struct {
}

type InputParameterParameters struct {

	// The unique name of the framework. The name must be between 1 and 256 characters, starting with a letter, and consisting of letters, numbers, and underscores.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The value of parameter, for example, hourly.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`
}

type ScopeObservation struct {
}

type ScopeParameters struct {

	// The ID of the only AWS resource that you want your control scope to contain. Minimum number of 1 item. Maximum number of 100 items.
	// +kubebuilder:validation:Optional
	ComplianceResourceIds []*string `json:"complianceResourceIds,omitempty" tf:"compliance_resource_ids,omitempty"`

	// Describes whether the control scope includes one or more types of resources, such as EFS or RDS.
	// +kubebuilder:validation:Optional
	ComplianceResourceTypes []*string `json:"complianceResourceTypes,omitempty" tf:"compliance_resource_types,omitempty"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// FrameworkSpec defines the desired state of Framework
type FrameworkSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FrameworkParameters `json:"forProvider"`
}

// FrameworkStatus defines the observed state of Framework.
type FrameworkStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FrameworkObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Framework is the Schema for the Frameworks API. Provides an AWS Backup Framework resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Framework struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrameworkSpec   `json:"spec"`
	Status            FrameworkStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FrameworkList contains a list of Frameworks
type FrameworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Framework `json:"items"`
}

// Repository type metadata.
var (
	Framework_Kind             = "Framework"
	Framework_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Framework_Kind}.String()
	Framework_KindAPIVersion   = Framework_Kind + "." + CRDGroupVersion.String()
	Framework_GroupVersionKind = CRDGroupVersion.WithKind(Framework_Kind)
)

func init() {
	SchemeBuilder.Register(&Framework{}, &FrameworkList{})
}
