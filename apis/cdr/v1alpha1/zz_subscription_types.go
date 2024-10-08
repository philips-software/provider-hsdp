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

type SubscriptionInitParameters struct {

	// On which resource to notify
	Criteria *string `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The REST endpoint to call for DELETE operations. Must use https:// schema
	DeleteEndpoint *string `json:"deleteEndpoint,omitempty" tf:"delete_endpoint,omitempty"`

	// RFC3339 formatted timestamp when to end notifications
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The REST endpoint to call. Must use https://  schema
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The CDR FHIR store endpoint to use
	FHIRStore *string `json:"fhirStore,omitempty" tf:"fhir_store,omitempty"`

	// List of headers to add to the REST call
	// +listType=set
	Headers []*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Reason for creating the subscription
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// The FHIR version to use. Options [ stu3 | r4 ]. Default is stu3
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SubscriptionObservation struct {

	// On which resource to notify
	Criteria *string `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The REST endpoint to call for DELETE operations. Must use https:// schema
	DeleteEndpoint *string `json:"deleteEndpoint,omitempty" tf:"delete_endpoint,omitempty"`

	// RFC3339 formatted timestamp when to end notifications
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The REST endpoint to call. Must use https://  schema
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The CDR FHIR store endpoint to use
	FHIRStore *string `json:"fhirStore,omitempty" tf:"fhir_store,omitempty"`

	// List of headers to add to the REST call
	// +listType=set
	Headers []*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// The ID of the CDR subscription
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Reason for creating the subscription
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// The status of the subscription (requested | active | error  | off)
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The FHIR version to use. Options [ stu3 | r4 ]. Default is stu3
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

type SubscriptionParameters struct {

	// On which resource to notify
	// +kubebuilder:validation:Optional
	Criteria *string `json:"criteria,omitempty" tf:"criteria,omitempty"`

	// The REST endpoint to call for DELETE operations. Must use https:// schema
	// +kubebuilder:validation:Optional
	DeleteEndpoint *string `json:"deleteEndpoint,omitempty" tf:"delete_endpoint,omitempty"`

	// RFC3339 formatted timestamp when to end notifications
	// +kubebuilder:validation:Optional
	End *string `json:"end,omitempty" tf:"end,omitempty"`

	// The REST endpoint to call. Must use https://  schema
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The CDR FHIR store endpoint to use
	// +kubebuilder:validation:Optional
	FHIRStore *string `json:"fhirStore,omitempty" tf:"fhir_store,omitempty"`

	// List of headers to add to the REST call
	// +kubebuilder:validation:Optional
	// +listType=set
	Headers []*string `json:"headers,omitempty" tf:"headers,omitempty"`

	// Reason for creating the subscription
	// +kubebuilder:validation:Optional
	Reason *string `json:"reason,omitempty" tf:"reason,omitempty"`

	// The FHIR version to use. Options [ stu3 | r4 ]. Default is stu3
	// +kubebuilder:validation:Optional
	Version *string `json:"version,omitempty" tf:"version,omitempty"`
}

// SubscriptionSpec defines the desired state of Subscription
type SubscriptionSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     SubscriptionParameters `json:"forProvider"`
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
	InitProvider SubscriptionInitParameters `json:"initProvider,omitempty"`
}

// SubscriptionStatus defines the observed state of Subscription.
type SubscriptionStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        SubscriptionObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Subscription is the Schema for the Subscriptions API. Manages HSDP CDR Subscription resources
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type Subscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.criteria) || (has(self.initProvider) && has(self.initProvider.criteria))",message="spec.forProvider.criteria is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.end) || (has(self.initProvider) && has(self.initProvider.end))",message="spec.forProvider.end is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.fhirStore) || (has(self.initProvider) && has(self.initProvider.fhirStore))",message="spec.forProvider.fhirStore is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.reason) || (has(self.initProvider) && has(self.initProvider.reason))",message="spec.forProvider.reason is a required parameter"
	Spec   SubscriptionSpec   `json:"spec"`
	Status SubscriptionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SubscriptionList contains a list of Subscriptions
type SubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Subscription `json:"items"`
}

// Repository type metadata.
var (
	Subscription_Kind             = "Subscription"
	Subscription_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Subscription_Kind}.String()
	Subscription_KindAPIVersion   = Subscription_Kind + "." + CRDGroupVersion.String()
	Subscription_GroupVersionKind = CRDGroupVersion.WithKind(Subscription_Kind)
)

func init() {
	SchemeBuilder.Register(&Subscription{}, &SubscriptionList{})
}
