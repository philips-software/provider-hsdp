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

type ObjectStoreInitParameters struct {

	// The base config URL of the DICOM Object store instance
	ConfigURL *string `json:"configUrl,omitempty" tf:"config_url,omitempty"`

	// Description of the object store
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// By default object stores will not be deleted by the provider (soft-delete).
	// By setting this value to true the provider removes the object store. We strongly suggest enabling this only for ephemeral deployments.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// the IAM organization ID to use for authorization
	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Selector for a Organization in iam to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Reference to a Organization in iam to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// the FHIR store configuration
	S3CredsAccess []S3CredsAccessInitParameters `json:"s3credsAccess,omitempty" tf:"s3creds_access,omitempty"`

	// Details of the CDR service account
	StaticAccess []StaticAccessInitParameters `json:"staticAccess,omitempty" tf:"static_access,omitempty"`
}

type ObjectStoreObservation struct {

	// The access type for this object store
	AccessType *string `json:"accessType,omitempty" tf:"access_type,omitempty"`

	// The base config URL of the DICOM Object store instance
	ConfigURL *string `json:"configUrl,omitempty" tf:"config_url,omitempty"`

	// Description of the object store
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// By default object stores will not be deleted by the provider (soft-delete).
	// By setting this value to true the provider removes the object store. We strongly suggest enabling this only for ephemeral deployments.
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// the IAM organization ID to use for authorization
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// the FHIR store configuration
	S3CredsAccess []S3CredsAccessObservation `json:"s3credsAccess,omitempty" tf:"s3creds_access,omitempty"`

	// Details of the CDR service account
	StaticAccess []StaticAccessObservation `json:"staticAccess,omitempty" tf:"static_access,omitempty"`
}

type ObjectStoreParameters struct {

	// The base config URL of the DICOM Object store instance
	// +kubebuilder:validation:Optional
	ConfigURL *string `json:"configUrl,omitempty" tf:"config_url,omitempty"`

	// Description of the object store
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// By default object stores will not be deleted by the provider (soft-delete).
	// By setting this value to true the provider removes the object store. We strongly suggest enabling this only for ephemeral deployments.
	// +kubebuilder:validation:Optional
	ForceDelete *bool `json:"forceDelete,omitempty" tf:"force_delete,omitempty"`

	// the IAM organization ID to use for authorization
	// +crossplane:generate:reference:type=github.com/philips-software/provider-hsdp/apis/iam/v1alpha1.Organization
	// +crossplane:generate:reference:refFieldName=OrganizationRef
	// +kubebuilder:validation:Optional
	OrganizationID *string `json:"organizationId,omitempty" tf:"organization_id,omitempty"`

	// Selector for a Organization in iam to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationIDSelector *v1.Selector `json:"organizationIdSelector,omitempty" tf:"-"`

	// Reference to a Organization in iam to populate organizationId.
	// +kubebuilder:validation:Optional
	OrganizationRef *v1.Reference `json:"organizationRef,omitempty" tf:"-"`

	// the FHIR store configuration
	// +kubebuilder:validation:Optional
	S3CredsAccess []S3CredsAccessParameters `json:"s3credsAccess,omitempty" tf:"s3creds_access,omitempty"`

	// Details of the CDR service account
	// +kubebuilder:validation:Optional
	StaticAccess []StaticAccessParameters `json:"staticAccess,omitempty" tf:"static_access,omitempty"`
}

type S3CredsAccessInitParameters struct {

	// The S3 bucket name
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The S3 bucket endpoint
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The S3Creds folder path to use
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// The S3Creds product key
	ProductKeySecretRef v1.SecretKeySelector `json:"productKeySecretRef" tf:"-"`

	// The IAM service account to use
	ServiceAccount []ServiceAccountInitParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
}

type S3CredsAccessObservation struct {

	// The S3 bucket name
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The S3 bucket endpoint
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The S3Creds folder path to use
	FolderPath *string `json:"folderPath,omitempty" tf:"folder_path,omitempty"`

	// The IAM service account to use
	ServiceAccount []ServiceAccountObservation `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
}

type S3CredsAccessParameters struct {

	// The S3 bucket name
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// The S3 bucket endpoint
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// The S3Creds folder path to use
	// +kubebuilder:validation:Optional
	FolderPath *string `json:"folderPath" tf:"folder_path,omitempty"`

	// The S3Creds product key
	// +kubebuilder:validation:Optional
	ProductKeySecretRef v1.SecretKeySelector `json:"productKeySecretRef" tf:"-"`

	// The IAM service account to use
	// +kubebuilder:validation:Optional
	ServiceAccount []ServiceAccountParameters `json:"serviceAccount,omitempty" tf:"service_account,omitempty"`
}

type ServiceAccountInitParameters struct {

	// The IAM access token endpoint
	AccessTokenEndpoint *string `json:"accessTokenEndpoint,omitempty" tf:"access_token_endpoint,omitempty"`

	// Name of the service
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The IAM service private key
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// The IAM service id
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// The IAM token endpoint
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`
}

type ServiceAccountObservation struct {

	// The IAM access token endpoint
	AccessTokenEndpoint *string `json:"accessTokenEndpoint,omitempty" tf:"access_token_endpoint,omitempty"`

	// Name of the service
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The IAM service id
	ServiceID *string `json:"serviceId,omitempty" tf:"service_id,omitempty"`

	// The IAM token endpoint
	TokenEndpoint *string `json:"tokenEndpoint,omitempty" tf:"token_endpoint,omitempty"`
}

type ServiceAccountParameters struct {

	// The IAM access token endpoint
	// +kubebuilder:validation:Optional
	AccessTokenEndpoint *string `json:"accessTokenEndpoint" tf:"access_token_endpoint,omitempty"`

	// Name of the service
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// The IAM service private key
	// +kubebuilder:validation:Optional
	PrivateKeySecretRef v1.SecretKeySelector `json:"privateKeySecretRef" tf:"-"`

	// The IAM service id
	// +kubebuilder:validation:Optional
	ServiceID *string `json:"serviceId" tf:"service_id,omitempty"`

	// The IAM token endpoint
	// +kubebuilder:validation:Optional
	TokenEndpoint *string `json:"tokenEndpoint" tf:"token_endpoint,omitempty"`
}

type StaticAccessInitParameters struct {

	// The S3 access key
	AccessKeySecretRef v1.SecretKeySelector `json:"accessKeySecretRef" tf:"-"`

	// The S3 bucket name
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The S3 bucket endpoint
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`

	// The S3 secret key
	SecretKeySecretRef v1.SecretKeySelector `json:"secretKeySecretRef" tf:"-"`
}

type StaticAccessObservation struct {

	// The S3 bucket name
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// The S3 bucket endpoint
	Endpoint *string `json:"endpoint,omitempty" tf:"endpoint,omitempty"`
}

type StaticAccessParameters struct {

	// The S3 access key
	// +kubebuilder:validation:Optional
	AccessKeySecretRef v1.SecretKeySelector `json:"accessKeySecretRef" tf:"-"`

	// The S3 bucket name
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName" tf:"bucket_name,omitempty"`

	// The S3 bucket endpoint
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint" tf:"endpoint,omitempty"`

	// The S3 secret key
	// +kubebuilder:validation:Optional
	SecretKeySecretRef v1.SecretKeySelector `json:"secretKeySecretRef" tf:"-"`
}

// ObjectStoreSpec defines the desired state of ObjectStore
type ObjectStoreSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectStoreParameters `json:"forProvider"`
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
	InitProvider ObjectStoreInitParameters `json:"initProvider,omitempty"`
}

// ObjectStoreStatus defines the observed state of ObjectStore.
type ObjectStoreStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectStoreObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// ObjectStore is the Schema for the ObjectStores API. Manages HSDP DICOM Object Stores
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,hsdp}
type ObjectStore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.configUrl) || (has(self.initProvider) && has(self.initProvider.configUrl))",message="spec.forProvider.configUrl is a required parameter"
	Spec   ObjectStoreSpec   `json:"spec"`
	Status ObjectStoreStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectStoreList contains a list of ObjectStores
type ObjectStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ObjectStore `json:"items"`
}

// Repository type metadata.
var (
	ObjectStore_Kind             = "ObjectStore"
	ObjectStore_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ObjectStore_Kind}.String()
	ObjectStore_KindAPIVersion   = ObjectStore_Kind + "." + CRDGroupVersion.String()
	ObjectStore_GroupVersionKind = CRDGroupVersion.WithKind(ObjectStore_Kind)
)

func init() {
	SchemeBuilder.Register(&ObjectStore{}, &ObjectStoreList{})
}
