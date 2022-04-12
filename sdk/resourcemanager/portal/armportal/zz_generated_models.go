//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armportal

// Configuration - Tenant configuration.
type Configuration struct {
	// Tenant configuration properties.
	Properties *ConfigurationProperties `json:"properties,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ConfigurationList - List of tenant configurations.
type ConfigurationList struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// The array of tenant configurations.
	Value []*Configuration `json:"value,omitempty"`
}

// ConfigurationProperties - Tenant configuration properties.
type ConfigurationProperties struct {
	// When flag is set to true Markdown tile will require external storage configuration (URI). The inline content configuration
	// will be prohibited.
	EnforcePrivateMarkdownStorage *bool `json:"enforcePrivateMarkdownStorage,omitempty"`
}

// Dashboard - The shared dashboard resource definition.
type Dashboard struct {
	// REQUIRED; Resource location
	Location *string `json:"location,omitempty"`

	// The shared dashboard properties.
	Properties *DashboardProperties `json:"properties,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Resource Id
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Resource name
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Resource type
	Type *string `json:"type,omitempty" azure:"ro"`
}

// DashboardLens - A dashboard lens.
type DashboardLens struct {
	// REQUIRED; The lens order.
	Order *int32 `json:"order,omitempty"`

	// REQUIRED; The dashboard parts.
	Parts []*DashboardParts `json:"parts,omitempty"`

	// The dashboard len's metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// DashboardListResult - List of dashboards.
type DashboardListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// The array of custom resource provider manifests.
	Value []*Dashboard `json:"value,omitempty"`
}

// DashboardPartMetadataClassification provides polymorphic access to related types.
// Call the interface's GetDashboardPartMetadata() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DashboardPartMetadata, *MarkdownPartMetadata
type DashboardPartMetadataClassification interface {
	// GetDashboardPartMetadata returns the DashboardPartMetadata content of the underlying type.
	GetDashboardPartMetadata() *DashboardPartMetadata
}

// DashboardPartMetadata - A dashboard part metadata.
type DashboardPartMetadata struct {
	// REQUIRED; The type of dashboard part.
	Type *string `json:"type,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]interface{}
}

// DashboardParts - A dashboard part.
type DashboardParts struct {
	// REQUIRED; The dashboard's part position.
	Position *DashboardPartsPosition `json:"position,omitempty"`

	// The dashboard part's metadata.
	Metadata DashboardPartMetadataClassification `json:"metadata,omitempty"`
}

// DashboardPartsPosition - The dashboard's part position.
type DashboardPartsPosition struct {
	// REQUIRED; The dashboard's part column span.
	ColSpan *int32 `json:"colSpan,omitempty"`

	// REQUIRED; The dashboard's part row span.
	RowSpan *int32 `json:"rowSpan,omitempty"`

	// REQUIRED; The dashboard's part x coordinate.
	X *int32 `json:"x,omitempty"`

	// REQUIRED; The dashboard's part y coordinate.
	Y *int32 `json:"y,omitempty"`

	// The dashboard part's metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// DashboardProperties - The shared dashboard properties.
type DashboardProperties struct {
	// The dashboard lenses.
	Lenses []*DashboardLens `json:"lenses,omitempty"`

	// The dashboard metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// DashboardsClientCreateOrUpdateOptions contains the optional parameters for the DashboardsClient.CreateOrUpdate method.
type DashboardsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// DashboardsClientDeleteOptions contains the optional parameters for the DashboardsClient.Delete method.
type DashboardsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// DashboardsClientGetOptions contains the optional parameters for the DashboardsClient.Get method.
type DashboardsClientGetOptions struct {
	// placeholder for future optional parameters
}

// DashboardsClientListByResourceGroupOptions contains the optional parameters for the DashboardsClient.ListByResourceGroup
// method.
type DashboardsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// DashboardsClientListBySubscriptionOptions contains the optional parameters for the DashboardsClient.ListBySubscription
// method.
type DashboardsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// DashboardsClientUpdateOptions contains the optional parameters for the DashboardsClient.Update method.
type DashboardsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ErrorDefinition - Error definition.
type ErrorDefinition struct {
	// READ-ONLY; Service specific error code which serves as the substatus for the HTTP error code.
	Code *int32 `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; Internal error details.
	Details []*ErrorDefinition `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; Description of the error.
	Message *string `json:"message,omitempty" azure:"ro"`
}

// ErrorResponse - Error response.
type ErrorResponse struct {
	// The error details.
	Error *ErrorDefinition `json:"error,omitempty"`
}

// ListTenantConfigurationViolationsClientListOptions contains the optional parameters for the ListTenantConfigurationViolationsClient.List
// method.
type ListTenantConfigurationViolationsClientListOptions struct {
	// placeholder for future optional parameters
}

// MarkdownPartMetadata - Markdown part metadata.
type MarkdownPartMetadata struct {
	// REQUIRED; The type of dashboard part.
	Type *string `json:"type,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]interface{}

	// Input to dashboard part.
	Inputs []interface{} `json:"inputs,omitempty"`

	// Markdown part settings.
	Settings *MarkdownPartMetadataSettings `json:"settings,omitempty"`
}

// MarkdownPartMetadataSettings - Markdown part settings.
type MarkdownPartMetadataSettings struct {
	// The content of markdown part.
	Content *MarkdownPartMetadataSettingsContent `json:"content,omitempty"`
}

// MarkdownPartMetadataSettingsContent - The content of markdown part.
type MarkdownPartMetadataSettingsContent struct {
	// The setting of the content of markdown part.
	Settings *MarkdownPartMetadataSettingsContentSettings `json:"settings,omitempty"`
}

// MarkdownPartMetadataSettingsContentSettings - The setting of the content of markdown part.
type MarkdownPartMetadataSettingsContentSettings struct {
	// The content of the markdown part.
	Content *string `json:"content,omitempty"`

	// The source of the content of the markdown part.
	MarkdownSource *int32 `json:"markdownSource,omitempty"`

	// The uri of markdown content.
	MarkdownURI *string `json:"markdownUri,omitempty"`

	// The subtitle of the markdown part.
	Subtitle *string `json:"subtitle,omitempty"`

	// The title of the markdown part.
	Title *string `json:"title,omitempty"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PatchableDashboard - The shared dashboard resource definition.
type PatchableDashboard struct {
	// The shared dashboard properties.
	Properties *DashboardProperties `json:"properties,omitempty"`

	// Resource tags
	Tags map[string]*string `json:"tags,omitempty"`
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ResourceProviderOperation - Supported operations of this resource provider.
type ResourceProviderOperation struct {
	// Display metadata associated with the operation.
	Display *ResourceProviderOperationDisplay `json:"display,omitempty"`

	// Indicates whether the operation applies to data-plane.
	IsDataAction *string `json:"isDataAction,omitempty"`

	// Operation name, in format of {provider}/{resource}/{operation}
	Name *string `json:"name,omitempty"`
}

// ResourceProviderOperationDisplay - Display metadata associated with the operation.
type ResourceProviderOperationDisplay struct {
	// Description of this operation.
	Description *string `json:"description,omitempty"`

	// Type of operation: get, read, delete, etc.
	Operation *string `json:"operation,omitempty"`

	// Resource provider: Microsoft Custom Providers.
	Provider *string `json:"provider,omitempty"`

	// Resource on which the operation is performed.
	Resource *string `json:"resource,omitempty"`
}

// ResourceProviderOperationList - Results of the request to list operations.
type ResourceProviderOperationList struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// List of operations supported by this resource provider.
	Value []*ResourceProviderOperation `json:"value,omitempty"`
}

// TenantConfigurationsClientCreateOptions contains the optional parameters for the TenantConfigurationsClient.Create method.
type TenantConfigurationsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// TenantConfigurationsClientDeleteOptions contains the optional parameters for the TenantConfigurationsClient.Delete method.
type TenantConfigurationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// TenantConfigurationsClientGetOptions contains the optional parameters for the TenantConfigurationsClient.Get method.
type TenantConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TenantConfigurationsClientListOptions contains the optional parameters for the TenantConfigurationsClient.List method.
type TenantConfigurationsClientListOptions struct {
	// placeholder for future optional parameters
}

// Violation information.
type Violation struct {
	// READ-ONLY; Error message.
	ErrorMessage *string `json:"errorMessage,omitempty" azure:"ro"`

	// READ-ONLY; Id of the item that violates tenant configuration.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Id of the user who owns violated item.
	UserID *string `json:"userId,omitempty" azure:"ro"`
}

// ViolationsList - List of list of items that violate tenant's configuration.
type ViolationsList struct {
	// The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty"`

	// The array of violations.
	Value []*Violation `json:"value,omitempty"`
}
