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

type ApplicationInitParameters struct {

	// The description of the application
	// The description of the application.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Reference identifier defined by the provisioning user.
	// Reference identifier defined by the provisioning user.
	GlobalReferenceID *string `json:"globalReferenceId,omitempty" tf:"global_reference_id,omitempty"`

	// The name of the application
	// The name of the application.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Blocks until the application delete has completed. Default: false.
	// The application delete process can take some time as all its associated resources like
	// services and clients are removed recursively. This option is useful for ephemeral environments
	// where the same application might be recreated shortly after a destroy operation.
	// Blocks until the application delete has completed. Default: false. The application delete process can take some time as all its associated resources like services and clients are removed recursively. This option is useful for ephemeral environments where the same application might be recreated shortly after a destroy operation.
	WaitForDelete *bool `json:"waitForDelete,omitempty" tf:"wait_for_delete,omitempty"`
}

type ApplicationObservation struct {

	// The description of the application
	// The description of the application.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Reference identifier defined by the provisioning user.
	// Reference identifier defined by the provisioning user.
	GlobalReferenceID *string `json:"globalReferenceId,omitempty" tf:"global_reference_id,omitempty"`

	// The GUID of the application
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the application
	// The name of the application.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// the proposition ID (GUID) to attach this a application to
	// The proposition ID (GUID) to attach this a application to.
	PropositionID *string `json:"propositionId,omitempty" tf:"proposition_id,omitempty"`

	// Blocks until the application delete has completed. Default: false.
	// The application delete process can take some time as all its associated resources like
	// services and clients are removed recursively. This option is useful for ephemeral environments
	// where the same application might be recreated shortly after a destroy operation.
	// Blocks until the application delete has completed. Default: false. The application delete process can take some time as all its associated resources like services and clients are removed recursively. This option is useful for ephemeral environments where the same application might be recreated shortly after a destroy operation.
	WaitForDelete *bool `json:"waitForDelete,omitempty" tf:"wait_for_delete,omitempty"`
}

type ApplicationParameters struct {

	// The description of the application
	// The description of the application.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Reference identifier defined by the provisioning user.
	// Reference identifier defined by the provisioning user.
	// +kubebuilder:validation:Optional
	GlobalReferenceID *string `json:"globalReferenceId,omitempty" tf:"global_reference_id,omitempty"`

	// The name of the application
	// The name of the application.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// the proposition ID (GUID) to attach this a application to
	// The proposition ID (GUID) to attach this a application to.
	// +crossplane:generate:reference:type=Proposition
	// +crossplane:generate:reference:refFieldName=PropositionRef
	// +kubebuilder:validation:Optional
	PropositionID *string `json:"propositionId,omitempty" tf:"proposition_id,omitempty"`

	// Selector for a Proposition to populate propositionId.
	// +kubebuilder:validation:Optional
	PropositionIDSelector *v1.Selector `json:"propositionIdSelector,omitempty" tf:"-"`

	// Reference to a Proposition to populate propositionId.
	// +kubebuilder:validation:Optional
	PropositionRef *v1.Reference `json:"propositionRef,omitempty" tf:"-"`

	// Blocks until the application delete has completed. Default: false.
	// The application delete process can take some time as all its associated resources like
	// services and clients are removed recursively. This option is useful for ephemeral environments
	// where the same application might be recreated shortly after a destroy operation.
	// Blocks until the application delete has completed. Default: false. The application delete process can take some time as all its associated resources like services and clients are removed recursively. This option is useful for ephemeral environments where the same application might be recreated shortly after a destroy operation.
	// +kubebuilder:validation:Optional
	WaitForDelete *bool `json:"waitForDelete,omitempty" tf:"wait_for_delete,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
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
	InitProvider ApplicationInitParameters `json:"initProvider,omitempty"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. Manages HSDP IAM Application resources
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   ApplicationSpec   `json:"spec"`
	Status ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
