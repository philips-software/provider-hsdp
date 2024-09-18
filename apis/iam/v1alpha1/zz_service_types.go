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

type ServiceInitParameters struct {

	// the application ID (GUID) to attach this service to
	// The application ID this service falls under.
	// +crossplane:generate:reference:type=Application
	// +crossplane:generate:reference:refFieldName=ApplicationRef
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Selector for a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// Reference to a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationRef *v1.Reference `json:"applicationRef,omitempty" tf:"-"`

	// Array. Default scopes. You do not have to specify these explicitly when requesting a token. Minimum: ["openid"]
	// Default scopes. You do not have to specify these explicitly when requesting a token.
	// +listType=set
	DefaultScopes []*string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	// The description of the service
	// The service description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the service
	// The service name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Array. List of supported scopes for this service. Minimum: ["openid"]
	// List of supported scopes for this service.
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// X509 Certificate in PEM format. When provided, overrides the generated certificate / private key combination of the IAM service. This gives you full control over the credentials. When not specified, a private key will be generated by IAM.
	// Mutually exclusive with `self_managed_private_key`
	SelfManagedCertificateSecretRef *v1.SecretKeySelector `json:"selfManagedCertificateSecretRef,omitempty" tf:"-"`

	// Sets the certificate validity. When not specified, the certificate will have a validity of 5 years.
	// Sets the certificate validity. When not specified, the certificate will have a validity of 5 years.
	// Only applicable when `self_managed_private_key` is used
	SelfManagedExpiresOn *string `json:"selfManagedExpiresOn,omitempty" tf:"self_managed_expires_on,omitempty"`

	// RSA private key in PEM format. When provided, overrides the generated certificate / private key combination of the
	// IAM service. This gives you full control over the credentials. When not specified, a private key will be generated by IAM
	// RSA private key in PEM format. When provided, overrides the generated certificate / private key combination of the IAM service. This gives you full control over the credentials. When not specified, a private key will be generated by IAM.
	// Mutually exclusive with `self_managed_certificate`
	SelfManagedPrivateKeySecretRef *v1.SecretKeySelector `json:"selfManagedPrivateKeySecretRef,omitempty" tf:"-"`

	// Integer. Access Token Lifetime (in seconds). Default: 1800 (30 minutes), Maximum: 2592000 (30 days)
	// Access Token Lifetime (in seconds). Default: 1800 (30 minutes), Maximum: 2592000 (30 days).
	TokenValidity *int64 `json:"tokenValidity,omitempty" tf:"token_validity,omitempty"`

	// Integer. Validity of service (in months). Minimum: 1, Maximum: 600 (5 years), Default: 12
	// The validity of the service credentials in months.
	Validity *int64 `json:"validity,omitempty" tf:"validity,omitempty"`
}

type ServiceObservation struct {

	// the application ID (GUID) to attach this service to
	// The application ID this service falls under.
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Array. Default scopes. You do not have to specify these explicitly when requesting a token. Minimum: ["openid"]
	// Default scopes. You do not have to specify these explicitly when requesting a token.
	// +listType=set
	DefaultScopes []*string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	// The description of the service
	// The service description.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Generated) Sets the certificate validity. When not specified, the certificate will have a validity of 5 years.
	// The expiration date of the service credentials.
	ExpiresOn *string `json:"expiresOn,omitempty" tf:"expires_on,omitempty"`

	// The GUID of the client
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The name of the service
	// The service name.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The organization ID this service belongs to (via application and proposition)
	// The organization this service falls under. Relationship established through application.
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Array. List of supported scopes for this service. Minimum: ["openid"]
	// List of supported scopes for this service.
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// Sets the certificate validity. When not specified, the certificate will have a validity of 5 years.
	// Sets the certificate validity. When not specified, the certificate will have a validity of 5 years.
	// Only applicable when `self_managed_private_key` is used
	SelfManagedExpiresOn *string `json:"selfManagedExpiresOn,omitempty" tf:"self_managed_expires_on,omitempty"`

	// (Generated) The service id
	// The service ID
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// Integer. Access Token Lifetime (in seconds). Default: 1800 (30 minutes), Maximum: 2592000 (30 days)
	// Access Token Lifetime (in seconds). Default: 1800 (30 minutes), Maximum: 2592000 (30 days).
	TokenValidity *int64 `json:"tokenValidity,omitempty" tf:"token_validity,omitempty"`

	// Integer. Validity of service (in months). Minimum: 1, Maximum: 600 (5 years), Default: 12
	// The validity of the service credentials in months.
	Validity *int64 `json:"validity,omitempty" tf:"validity,omitempty"`
}

type ServiceParameters struct {

	// the application ID (GUID) to attach this service to
	// The application ID this service falls under.
	// +crossplane:generate:reference:type=Application
	// +crossplane:generate:reference:refFieldName=ApplicationRef
	// +kubebuilder:validation:Optional
	ApplicationID *string `json:"applicationId,omitempty" tf:"application_id,omitempty"`

	// Selector for a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationIDSelector *v1.Selector `json:"applicationIdSelector,omitempty" tf:"-"`

	// Reference to a Application to populate applicationId.
	// +kubebuilder:validation:Optional
	ApplicationRef *v1.Reference `json:"applicationRef,omitempty" tf:"-"`

	// Array. Default scopes. You do not have to specify these explicitly when requesting a token. Minimum: ["openid"]
	// Default scopes. You do not have to specify these explicitly when requesting a token.
	// +kubebuilder:validation:Optional
	// +listType=set
	DefaultScopes []*string `json:"defaultScopes,omitempty" tf:"default_scopes,omitempty"`

	// The description of the service
	// The service description.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// The name of the service
	// The service name.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Array. List of supported scopes for this service. Minimum: ["openid"]
	// List of supported scopes for this service.
	// +kubebuilder:validation:Optional
	// +listType=set
	Scopes []*string `json:"scopes,omitempty" tf:"scopes,omitempty"`

	// X509 Certificate in PEM format. When provided, overrides the generated certificate / private key combination of the IAM service. This gives you full control over the credentials. When not specified, a private key will be generated by IAM.
	// Mutually exclusive with `self_managed_private_key`
	// +kubebuilder:validation:Optional
	SelfManagedCertificateSecretRef *v1.SecretKeySelector `json:"selfManagedCertificateSecretRef,omitempty" tf:"-"`

	// Sets the certificate validity. When not specified, the certificate will have a validity of 5 years.
	// Sets the certificate validity. When not specified, the certificate will have a validity of 5 years.
	// Only applicable when `self_managed_private_key` is used
	// +kubebuilder:validation:Optional
	SelfManagedExpiresOn *string `json:"selfManagedExpiresOn,omitempty" tf:"self_managed_expires_on,omitempty"`

	// RSA private key in PEM format. When provided, overrides the generated certificate / private key combination of the
	// IAM service. This gives you full control over the credentials. When not specified, a private key will be generated by IAM
	// RSA private key in PEM format. When provided, overrides the generated certificate / private key combination of the IAM service. This gives you full control over the credentials. When not specified, a private key will be generated by IAM.
	// Mutually exclusive with `self_managed_certificate`
	// +kubebuilder:validation:Optional
	SelfManagedPrivateKeySecretRef *v1.SecretKeySelector `json:"selfManagedPrivateKeySecretRef,omitempty" tf:"-"`

	// Integer. Access Token Lifetime (in seconds). Default: 1800 (30 minutes), Maximum: 2592000 (30 days)
	// Access Token Lifetime (in seconds). Default: 1800 (30 minutes), Maximum: 2592000 (30 days).
	// +kubebuilder:validation:Optional
	TokenValidity *int64 `json:"tokenValidity,omitempty" tf:"token_validity,omitempty"`

	// Integer. Validity of service (in months). Minimum: 1, Maximum: 600 (5 years), Default: 12
	// The validity of the service credentials in months.
	// +kubebuilder:validation:Optional
	Validity *int64 `json:"validity,omitempty" tf:"validity,omitempty"`
}

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServiceParameters `json:"forProvider"`
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
	InitProvider ServiceInitParameters `json:"initProvider,omitempty"`
}

// ServiceStatus defines the observed state of Service.
type ServiceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServiceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Service is the Schema for the Services API. Manages HSDP IAM Service resources
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.defaultScopes) || (has(self.initProvider) && has(self.initProvider.defaultScopes))",message="spec.forProvider.defaultScopes is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.description) || (has(self.initProvider) && has(self.initProvider.description))",message="spec.forProvider.description is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scopes) || (has(self.initProvider) && has(self.initProvider.scopes))",message="spec.forProvider.scopes is a required parameter"
	Spec   ServiceSpec   `json:"spec"`
	Status ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Services
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

// Repository type metadata.
var (
	Service_Kind             = "Service"
	Service_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Service_Kind}.String()
	Service_KindAPIVersion   = Service_Kind + "." + CRDGroupVersion.String()
	Service_GroupVersionKind = CRDGroupVersion.WithKind(Service_Kind)
)

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
