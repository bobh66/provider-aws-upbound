// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeploymentTargetsInitParameters struct {

	// Limit deployment targets to individual accounts or include additional accounts with provided OUs. Valid values: INTERSECTION, DIFFERENCE, UNION, NONE.
	AccountFilterType *string `json:"accountFilterType,omitempty" tf:"account_filter_type,omitempty"`

	// List of accounts to deploy stack set updates.
	// +listType=set
	Accounts []*string `json:"accounts,omitempty" tf:"accounts,omitempty"`

	// S3 URL of the file containing the list of accounts.
	AccountsURL *string `json:"accountsUrl,omitempty" tf:"accounts_url,omitempty"`

	// Organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	// +listType=set
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" tf:"organizational_unit_ids,omitempty"`
}

type DeploymentTargetsObservation struct {

	// Limit deployment targets to individual accounts or include additional accounts with provided OUs. Valid values: INTERSECTION, DIFFERENCE, UNION, NONE.
	AccountFilterType *string `json:"accountFilterType,omitempty" tf:"account_filter_type,omitempty"`

	// List of accounts to deploy stack set updates.
	// +listType=set
	Accounts []*string `json:"accounts,omitempty" tf:"accounts,omitempty"`

	// S3 URL of the file containing the list of accounts.
	AccountsURL *string `json:"accountsUrl,omitempty" tf:"accounts_url,omitempty"`

	// Organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	// +listType=set
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" tf:"organizational_unit_ids,omitempty"`
}

type DeploymentTargetsParameters struct {

	// Limit deployment targets to individual accounts or include additional accounts with provided OUs. Valid values: INTERSECTION, DIFFERENCE, UNION, NONE.
	// +kubebuilder:validation:Optional
	AccountFilterType *string `json:"accountFilterType,omitempty" tf:"account_filter_type,omitempty"`

	// List of accounts to deploy stack set updates.
	// +kubebuilder:validation:Optional
	// +listType=set
	Accounts []*string `json:"accounts,omitempty" tf:"accounts,omitempty"`

	// S3 URL of the file containing the list of accounts.
	// +kubebuilder:validation:Optional
	AccountsURL *string `json:"accountsUrl,omitempty" tf:"accounts_url,omitempty"`

	// Organization root ID or organizational unit (OU) IDs to which StackSets deploys.
	// +kubebuilder:validation:Optional
	// +listType=set
	OrganizationalUnitIds []*string `json:"organizationalUnitIds,omitempty" tf:"organizational_unit_ids,omitempty"`
}

type StackInstanceSummariesInitParameters struct {
}

type StackInstanceSummariesObservation struct {

	// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Organization root ID or organizational unit (OU) ID in which the stack is deployed.
	OrganizationalUnitID *string `json:"organizationalUnitId,omitempty" tf:"organizational_unit_id,omitempty"`

	// Stack identifier.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`
}

type StackInstanceSummariesParameters struct {
}

type StackSetInstanceInitParameters struct {

	// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account. Valid values: SELF (default), DELEGATED_ADMIN.
	CallAs *string `json:"callAs,omitempty" tf:"call_as,omitempty"`

	// AWS Organizations accounts to which StackSets deploys. StackSets doesn't deploy stack instances to the organization management account, even if the organization management account is in your organization or in an OU in your organization. Drift detection is not possible for this argument. See deployment_targets below.
	DeploymentTargets *DeploymentTargetsInitParameters `json:"deploymentTargets,omitempty" tf:"deployment_targets,omitempty"`

	// Preferences for how AWS CloudFormation performs a stack set operation.
	OperationPreferences *StackSetInstanceOperationPreferencesInitParameters `json:"operationPreferences,omitempty" tf:"operation_preferences,omitempty"`

	// Key-value map of input parameters to override from the StackSet for this Instance.
	// +mapType=granular
	ParameterOverrides map[string]*string `json:"parameterOverrides,omitempty" tf:"parameter_overrides,omitempty"`

	// You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to false.
	RetainStack *bool `json:"retainStack,omitempty" tf:"retain_stack,omitempty"`

	// Name of the StackSet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudformation/v1beta2.StackSet
	StackSetName *string `json:"stackSetName,omitempty" tf:"stack_set_name,omitempty"`

	// Reference to a StackSet in cloudformation to populate stackSetName.
	// +kubebuilder:validation:Optional
	StackSetNameRef *v1.Reference `json:"stackSetNameRef,omitempty" tf:"-"`

	// Selector for a StackSet in cloudformation to populate stackSetName.
	// +kubebuilder:validation:Optional
	StackSetNameSelector *v1.Selector `json:"stackSetNameSelector,omitempty" tf:"-"`
}

type StackSetInstanceObservation struct {

	// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account. Valid values: SELF (default), DELEGATED_ADMIN.
	CallAs *string `json:"callAs,omitempty" tf:"call_as,omitempty"`

	// AWS Organizations accounts to which StackSets deploys. StackSets doesn't deploy stack instances to the organization management account, even if the organization management account is in your organization or in an OU in your organization. Drift detection is not possible for this argument. See deployment_targets below.
	DeploymentTargets *DeploymentTargetsObservation `json:"deploymentTargets,omitempty" tf:"deployment_targets,omitempty"`

	// Unique identifier for the resource. If deployment_targets is set, this is a comma-delimited string combining stack set name, organizational unit IDs (/-delimited), and region (ie. mystack,ou-123/ou-456,us-east-1). Otherwise, this is a comma-delimited string combining stack set name, AWS account ID, and region (ie. mystack,123456789012,us-east-1).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Preferences for how AWS CloudFormation performs a stack set operation.
	OperationPreferences *StackSetInstanceOperationPreferencesObservation `json:"operationPreferences,omitempty" tf:"operation_preferences,omitempty"`

	// Organization root ID or organizational unit (OU) ID in which the stack is deployed.
	OrganizationalUnitID *string `json:"organizationalUnitId,omitempty" tf:"organizational_unit_id,omitempty"`

	// Key-value map of input parameters to override from the StackSet for this Instance.
	// +mapType=granular
	ParameterOverrides map[string]*string `json:"parameterOverrides,omitempty" tf:"parameter_overrides,omitempty"`

	// You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to false.
	RetainStack *bool `json:"retainStack,omitempty" tf:"retain_stack,omitempty"`

	// Stack identifier.
	StackID *string `json:"stackId,omitempty" tf:"stack_id,omitempty"`

	// List of stack instances created from an organizational unit deployment target. This will only be populated when deployment_targets is set. See stack_instance_summaries.
	// List of stack instances created from an organizational unit deployment target. This will only be populated when `deployment_targets` is set.
	StackInstanceSummaries []StackInstanceSummariesObservation `json:"stackInstanceSummaries,omitempty" tf:"stack_instance_summaries,omitempty"`

	// Name of the StackSet.
	StackSetName *string `json:"stackSetName,omitempty" tf:"stack_set_name,omitempty"`
}

type StackSetInstanceOperationPreferencesInitParameters struct {

	// Number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureToleranceCount *float64 `json:"failureToleranceCount,omitempty" tf:"failure_tolerance_count,omitempty"`

	// Percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureTolerancePercentage *float64 `json:"failureTolerancePercentage,omitempty" tf:"failure_tolerance_percentage,omitempty"`

	// Maximum number of accounts in which to perform this operation at one time.
	MaxConcurrentCount *float64 `json:"maxConcurrentCount,omitempty" tf:"max_concurrent_count,omitempty"`

	// Maximum percentage of accounts in which to perform this operation at one time.
	MaxConcurrentPercentage *float64 `json:"maxConcurrentPercentage,omitempty" tf:"max_concurrent_percentage,omitempty"`

	// Concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time. Valid values are SEQUENTIAL and PARALLEL.
	RegionConcurrencyType *string `json:"regionConcurrencyType,omitempty" tf:"region_concurrency_type,omitempty"`

	// Order of the Regions in where you want to perform the stack operation.
	RegionOrder []*string `json:"regionOrder,omitempty" tf:"region_order,omitempty"`
}

type StackSetInstanceOperationPreferencesObservation struct {

	// Number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureToleranceCount *float64 `json:"failureToleranceCount,omitempty" tf:"failure_tolerance_count,omitempty"`

	// Percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureTolerancePercentage *float64 `json:"failureTolerancePercentage,omitempty" tf:"failure_tolerance_percentage,omitempty"`

	// Maximum number of accounts in which to perform this operation at one time.
	MaxConcurrentCount *float64 `json:"maxConcurrentCount,omitempty" tf:"max_concurrent_count,omitempty"`

	// Maximum percentage of accounts in which to perform this operation at one time.
	MaxConcurrentPercentage *float64 `json:"maxConcurrentPercentage,omitempty" tf:"max_concurrent_percentage,omitempty"`

	// Concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time. Valid values are SEQUENTIAL and PARALLEL.
	RegionConcurrencyType *string `json:"regionConcurrencyType,omitempty" tf:"region_concurrency_type,omitempty"`

	// Order of the Regions in where you want to perform the stack operation.
	RegionOrder []*string `json:"regionOrder,omitempty" tf:"region_order,omitempty"`
}

type StackSetInstanceOperationPreferencesParameters struct {

	// Number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	// +kubebuilder:validation:Optional
	FailureToleranceCount *float64 `json:"failureToleranceCount,omitempty" tf:"failure_tolerance_count,omitempty"`

	// Percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	// +kubebuilder:validation:Optional
	FailureTolerancePercentage *float64 `json:"failureTolerancePercentage,omitempty" tf:"failure_tolerance_percentage,omitempty"`

	// Maximum number of accounts in which to perform this operation at one time.
	// +kubebuilder:validation:Optional
	MaxConcurrentCount *float64 `json:"maxConcurrentCount,omitempty" tf:"max_concurrent_count,omitempty"`

	// Maximum percentage of accounts in which to perform this operation at one time.
	// +kubebuilder:validation:Optional
	MaxConcurrentPercentage *float64 `json:"maxConcurrentPercentage,omitempty" tf:"max_concurrent_percentage,omitempty"`

	// Concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time. Valid values are SEQUENTIAL and PARALLEL.
	// +kubebuilder:validation:Optional
	RegionConcurrencyType *string `json:"regionConcurrencyType,omitempty" tf:"region_concurrency_type,omitempty"`

	// Order of the Regions in where you want to perform the stack operation.
	// +kubebuilder:validation:Optional
	RegionOrder []*string `json:"regionOrder,omitempty" tf:"region_order,omitempty"`
}

type StackSetInstanceParameters struct {

	// Target AWS Account ID to create a Stack based on the StackSet. Defaults to current account.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// Specifies whether you are acting as an account administrator in the organization's management account or as a delegated administrator in a member account. Valid values: SELF (default), DELEGATED_ADMIN.
	// +kubebuilder:validation:Optional
	CallAs *string `json:"callAs,omitempty" tf:"call_as,omitempty"`

	// AWS Organizations accounts to which StackSets deploys. StackSets doesn't deploy stack instances to the organization management account, even if the organization management account is in your organization or in an OU in your organization. Drift detection is not possible for this argument. See deployment_targets below.
	// +kubebuilder:validation:Optional
	DeploymentTargets *DeploymentTargetsParameters `json:"deploymentTargets,omitempty" tf:"deployment_targets,omitempty"`

	// Preferences for how AWS CloudFormation performs a stack set operation.
	// +kubebuilder:validation:Optional
	OperationPreferences *StackSetInstanceOperationPreferencesParameters `json:"operationPreferences,omitempty" tf:"operation_preferences,omitempty"`

	// Key-value map of input parameters to override from the StackSet for this Instance.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	ParameterOverrides map[string]*string `json:"parameterOverrides,omitempty" tf:"parameter_overrides,omitempty"`

	// Target AWS Region to create a Stack based on the StackSet. Defaults to current region.
	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// You cannot reassociate a retained Stack or add an existing, saved Stack to a new StackSet. Defaults to false.
	// +kubebuilder:validation:Optional
	RetainStack *bool `json:"retainStack,omitempty" tf:"retain_stack,omitempty"`

	// Name of the StackSet.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/cloudformation/v1beta2.StackSet
	// +kubebuilder:validation:Optional
	StackSetName *string `json:"stackSetName,omitempty" tf:"stack_set_name,omitempty"`

	// Reference to a StackSet in cloudformation to populate stackSetName.
	// +kubebuilder:validation:Optional
	StackSetNameRef *v1.Reference `json:"stackSetNameRef,omitempty" tf:"-"`

	// Selector for a StackSet in cloudformation to populate stackSetName.
	// +kubebuilder:validation:Optional
	StackSetNameSelector *v1.Selector `json:"stackSetNameSelector,omitempty" tf:"-"`
}

// StackSetInstanceSpec defines the desired state of StackSetInstance
type StackSetInstanceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     StackSetInstanceParameters `json:"forProvider"`
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
	InitProvider StackSetInstanceInitParameters `json:"initProvider,omitempty"`
}

// StackSetInstanceStatus defines the observed state of StackSetInstance.
type StackSetInstanceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        StackSetInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// StackSetInstance is the Schema for the StackSetInstances API. Manages a CloudFormation StackSet Instance.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type StackSetInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StackSetInstanceSpec   `json:"spec"`
	Status            StackSetInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// StackSetInstanceList contains a list of StackSetInstances
type StackSetInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StackSetInstance `json:"items"`
}

// Repository type metadata.
var (
	StackSetInstance_Kind             = "StackSetInstance"
	StackSetInstance_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: StackSetInstance_Kind}.String()
	StackSetInstance_KindAPIVersion   = StackSetInstance_Kind + "." + CRDGroupVersion.String()
	StackSetInstance_GroupVersionKind = CRDGroupVersion.WithKind(StackSetInstance_Kind)
)

func init() {
	SchemeBuilder.Register(&StackSetInstance{}, &StackSetInstanceList{})
}
