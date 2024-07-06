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

type DataTypeInitParameters struct {

	// A short description of the device group
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the device group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID reference of the service action (format: DataType/${GUID})
	// +crossplane:generate:reference:type=Proposition
	// +crossplane:generate:reference:refFieldName=PropositionRef
	// +crossplane:generate:reference:selectorFieldName=PropositionSelector
	PropositionID *string `json:"propositionId,omitempty" tf:"proposition_id,omitempty"`

	// Reference to a Proposition to populate propositionId.
	// +kubebuilder:validation:Optional
	PropositionRef *v1.Reference `json:"propositionRef,omitempty" tf:"-"`

	// Selector for a Proposition to populate propositionId.
	// +kubebuilder:validation:Optional
	PropositionSelector *v1.Selector `json:"propositionSelector,omitempty" tf:"-"`

	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DataTypeObservation struct {

	// A short description of the device group
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The GUID of the service action
	GUID *string `json:"guid,omitempty" tf:"guid,omitempty"`

	// The ID reference of the service action (format: DataType/${GUID})
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the device group
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID reference of the service action (format: DataType/${GUID})
	PropositionID *string `json:"propositionId,omitempty" tf:"proposition_id,omitempty"`

	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// The ID reference of the service action (format: DataType/${GUID})
	VersionID *string `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type DataTypeParameters struct {

	// A short description of the device group
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the device group
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The ID reference of the service action (format: DataType/${GUID})
	// +crossplane:generate:reference:type=Proposition
	// +crossplane:generate:reference:refFieldName=PropositionRef
	// +crossplane:generate:reference:selectorFieldName=PropositionSelector
	// +kubebuilder:validation:Optional
	PropositionID *string `json:"propositionId,omitempty" tf:"proposition_id,omitempty"`

	// Reference to a Proposition to populate propositionId.
	// +kubebuilder:validation:Optional
	PropositionRef *v1.Reference `json:"propositionRef,omitempty" tf:"-"`

	// Selector for a Proposition to populate propositionId.
	// +kubebuilder:validation:Optional
	PropositionSelector *v1.Selector `json:"propositionSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	// +listType=set
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// DataTypeSpec defines the desired state of DataType
type DataTypeSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DataTypeParameters `json:"forProvider"`
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
	InitProvider DataTypeInitParameters `json:"initProvider,omitempty"`
}

// DataTypeStatus defines the observed state of DataType.
type DataTypeStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DataTypeObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// DataType is the Schema for the DataTypes API. Manages HSDP Connect MDM Data types
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type DataType struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   DataTypeSpec   `json:"spec"`
	Status DataTypeStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DataTypeList contains a list of DataTypes
type DataTypeList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DataType `json:"items"`
}

// Repository type metadata.
var (
	DataType_Kind             = "DataType"
	DataType_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DataType_Kind}.String()
	DataType_KindAPIVersion   = DataType_Kind + "." + CRDGroupVersion.String()
	DataType_GroupVersionKind = CRDGroupVersion.WithKind(DataType_Kind)
)

func init() {
	SchemeBuilder.Register(&DataType{}, &DataTypeList{})
}
