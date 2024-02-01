// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type LBListenerCertificateInitParameters struct {

	// The ARN of the certificate to attach to the listener.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/acm/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Reference to a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnRef *v1.Reference `json:"certificateArnRef,omitempty" tf:"-"`

	// Selector for a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnSelector *v1.Selector `json:"certificateArnSelector,omitempty" tf:"-"`

	// The ARN of the listener to which to attach the certificate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBListener
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	ListenerArn *string `json:"listenerArn,omitempty" tf:"listener_arn,omitempty"`

	// Reference to a LBListener in elbv2 to populate listenerArn.
	// +kubebuilder:validation:Optional
	ListenerArnRef *v1.Reference `json:"listenerArnRef,omitempty" tf:"-"`

	// Selector for a LBListener in elbv2 to populate listenerArn.
	// +kubebuilder:validation:Optional
	ListenerArnSelector *v1.Selector `json:"listenerArnSelector,omitempty" tf:"-"`
}

type LBListenerCertificateObservation struct {

	// The ARN of the certificate to attach to the listener.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// The listener_arn and certificate_arn separated by a _.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ARN of the listener to which to attach the certificate.
	ListenerArn *string `json:"listenerArn,omitempty" tf:"listener_arn,omitempty"`
}

type LBListenerCertificateParameters struct {

	// The ARN of the certificate to attach to the listener.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/acm/v1beta1.Certificate
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Reference to a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnRef *v1.Reference `json:"certificateArnRef,omitempty" tf:"-"`

	// Selector for a Certificate in acm to populate certificateArn.
	// +kubebuilder:validation:Optional
	CertificateArnSelector *v1.Selector `json:"certificateArnSelector,omitempty" tf:"-"`

	// The ARN of the listener to which to attach the certificate.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/elbv2/v1beta1.LBListener
	// +crossplane:generate:reference:extractor=github.com/crossplane/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ListenerArn *string `json:"listenerArn,omitempty" tf:"listener_arn,omitempty"`

	// Reference to a LBListener in elbv2 to populate listenerArn.
	// +kubebuilder:validation:Optional
	ListenerArnRef *v1.Reference `json:"listenerArnRef,omitempty" tf:"-"`

	// Selector for a LBListener in elbv2 to populate listenerArn.
	// +kubebuilder:validation:Optional
	ListenerArnSelector *v1.Selector `json:"listenerArnSelector,omitempty" tf:"-"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`
}

// LBListenerCertificateSpec defines the desired state of LBListenerCertificate
type LBListenerCertificateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LBListenerCertificateParameters `json:"forProvider"`
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
	InitProvider LBListenerCertificateInitParameters `json:"initProvider,omitempty"`
}

// LBListenerCertificateStatus defines the observed state of LBListenerCertificate.
type LBListenerCertificateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LBListenerCertificateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// LBListenerCertificate is the Schema for the LBListenerCertificates API. Provides a Load Balancer Listener Certificate resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type LBListenerCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LBListenerCertificateSpec   `json:"spec"`
	Status            LBListenerCertificateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LBListenerCertificateList contains a list of LBListenerCertificates
type LBListenerCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LBListenerCertificate `json:"items"`
}

// Repository type metadata.
var (
	LBListenerCertificate_Kind             = "LBListenerCertificate"
	LBListenerCertificate_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: LBListenerCertificate_Kind}.String()
	LBListenerCertificate_KindAPIVersion   = LBListenerCertificate_Kind + "." + CRDGroupVersion.String()
	LBListenerCertificate_GroupVersionKind = CRDGroupVersion.WithKind(LBListenerCertificate_Kind)
)

func init() {
	SchemeBuilder.Register(&LBListenerCertificate{}, &LBListenerCertificateList{})
}