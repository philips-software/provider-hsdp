// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type PropositionInitParameters struct {

	// The description of the application
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Reference identifier defined by the provisioning user.
	GlobalReferenceID *string `json:"globalReferenceId,omitempty" tf:"global_reference_id,omitempty"`

	// The name of the application
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type PropositionObservation struct {

	// The description of the application
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Reference identifier defined by the provisioning user.
	GlobalReferenceID *string `json:"globalReferenceId,omitempty" tf:"global_reference_id,omitempty"`

	// The GUID of the proposition
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the application
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// the organization ID (GUID) to attach this a proposition to
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`
}

type PropositionParameters struct {

	// The description of the application
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Reference identifier defined by the provisioning user.
	// +kubebuilder:validation:Optional
	GlobalReferenceID *string `json:"globalReferenceId,omitempty" tf:"global_reference_id,omitempty"`

	// The name of the application
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// the organization ID (GUID) to attach this a proposition to
	// +crossplane:generate:reference:type=Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Selector for a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Reference to a Organization to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`
}

// PropositionSpec defines the desired state of Proposition
type PropositionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     PropositionParameters `json:"forProvider"`
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
	InitProvider PropositionInitParameters `json:"initProvider,omitempty"`
}

// PropositionStatus defines the observed state of Proposition.
type PropositionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        PropositionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Proposition is the Schema for the Propositions API. Manages HSDP IAM Proposition resources
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type Proposition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   PropositionSpec   `json:"spec"`
	Status PropositionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PropositionList contains a list of Propositions
type PropositionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Proposition `json:"items"`
}

// Repository type metadata.
var (
	Proposition_Kind             = "Proposition"
	Proposition_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Proposition_Kind}.String()
	Proposition_KindAPIVersion   = Proposition_Kind + "." + CRDGroupVersion.String()
	Proposition_GroupVersionKind = CRDGroupVersion.WithKind(Proposition_Kind)
)

func init() {
	SchemeBuilder.Register(&Proposition{}, &PropositionList{})
}
