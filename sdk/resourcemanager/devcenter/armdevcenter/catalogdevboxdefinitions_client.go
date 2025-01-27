//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevcenter

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// CatalogDevBoxDefinitionsClient contains the methods for the CatalogDevBoxDefinitions group.
// Don't use this type directly, use NewCatalogDevBoxDefinitionsClient() instead.
type CatalogDevBoxDefinitionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCatalogDevBoxDefinitionsClient creates a new instance of CatalogDevBoxDefinitionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCatalogDevBoxDefinitionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CatalogDevBoxDefinitionsClient, error) {
	cl, err := arm.NewClient(moduleName+".CatalogDevBoxDefinitionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CatalogDevBoxDefinitionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a Dev Box definition from the catalog
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - devBoxDefinitionName - The name of the Dev Box definition.
//   - options - CatalogDevBoxDefinitionsClientGetOptions contains the optional parameters for the CatalogDevBoxDefinitionsClient.Get
//     method.
func (client *CatalogDevBoxDefinitionsClient) Get(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, devBoxDefinitionName string, options *CatalogDevBoxDefinitionsClientGetOptions) (CatalogDevBoxDefinitionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, devBoxDefinitionName, options)
	if err != nil {
		return CatalogDevBoxDefinitionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CatalogDevBoxDefinitionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CatalogDevBoxDefinitionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CatalogDevBoxDefinitionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, devBoxDefinitionName string, options *CatalogDevBoxDefinitionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/devboxdefinitions/{devBoxDefinitionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if devBoxDefinitionName == "" {
		return nil, errors.New("parameter devBoxDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devBoxDefinitionName}", url.PathEscape(devBoxDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CatalogDevBoxDefinitionsClient) getHandleResponse(resp *http.Response) (CatalogDevBoxDefinitionsClientGetResponse, error) {
	result := CatalogDevBoxDefinitionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevBoxDefinition); err != nil {
		return CatalogDevBoxDefinitionsClientGetResponse{}, err
	}
	return result, nil
}

// GetErrorDetails - Gets Catalog Devbox Definition error details
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - devBoxDefinitionName - The name of the Dev Box definition.
//   - options - CatalogDevBoxDefinitionsClientGetErrorDetailsOptions contains the optional parameters for the CatalogDevBoxDefinitionsClient.GetErrorDetails
//     method.
func (client *CatalogDevBoxDefinitionsClient) GetErrorDetails(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, devBoxDefinitionName string, options *CatalogDevBoxDefinitionsClientGetErrorDetailsOptions) (CatalogDevBoxDefinitionsClientGetErrorDetailsResponse, error) {
	var err error
	req, err := client.getErrorDetailsCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, devBoxDefinitionName, options)
	if err != nil {
		return CatalogDevBoxDefinitionsClientGetErrorDetailsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CatalogDevBoxDefinitionsClientGetErrorDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CatalogDevBoxDefinitionsClientGetErrorDetailsResponse{}, err
	}
	resp, err := client.getErrorDetailsHandleResponse(httpResp)
	return resp, err
}

// getErrorDetailsCreateRequest creates the GetErrorDetails request.
func (client *CatalogDevBoxDefinitionsClient) getErrorDetailsCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, devBoxDefinitionName string, options *CatalogDevBoxDefinitionsClientGetErrorDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/devboxdefinitions/{devBoxDefinitionName}/getErrorDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	if devBoxDefinitionName == "" {
		return nil, errors.New("parameter devBoxDefinitionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devBoxDefinitionName}", url.PathEscape(devBoxDefinitionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getErrorDetailsHandleResponse handles the GetErrorDetails response.
func (client *CatalogDevBoxDefinitionsClient) getErrorDetailsHandleResponse(resp *http.Response) (CatalogDevBoxDefinitionsClientGetErrorDetailsResponse, error) {
	result := CatalogDevBoxDefinitionsClientGetErrorDetailsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CatalogResourceValidationErrorDetails); err != nil {
		return CatalogDevBoxDefinitionsClientGetErrorDetailsResponse{}, err
	}
	return result, nil
}

// NewListByCatalogPager - List Dev Box definitions in the catalog.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - devCenterName - The name of the devcenter.
//   - catalogName - The name of the Catalog.
//   - options - CatalogDevBoxDefinitionsClientListByCatalogOptions contains the optional parameters for the CatalogDevBoxDefinitionsClient.NewListByCatalogPager
//     method.
func (client *CatalogDevBoxDefinitionsClient) NewListByCatalogPager(resourceGroupName string, devCenterName string, catalogName string, options *CatalogDevBoxDefinitionsClientListByCatalogOptions) *runtime.Pager[CatalogDevBoxDefinitionsClientListByCatalogResponse] {
	return runtime.NewPager(runtime.PagingHandler[CatalogDevBoxDefinitionsClientListByCatalogResponse]{
		More: func(page CatalogDevBoxDefinitionsClientListByCatalogResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CatalogDevBoxDefinitionsClientListByCatalogResponse) (CatalogDevBoxDefinitionsClientListByCatalogResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByCatalogCreateRequest(ctx, resourceGroupName, devCenterName, catalogName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CatalogDevBoxDefinitionsClientListByCatalogResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CatalogDevBoxDefinitionsClientListByCatalogResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CatalogDevBoxDefinitionsClientListByCatalogResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByCatalogHandleResponse(resp)
		},
	})
}

// listByCatalogCreateRequest creates the ListByCatalog request.
func (client *CatalogDevBoxDefinitionsClient) listByCatalogCreateRequest(ctx context.Context, resourceGroupName string, devCenterName string, catalogName string, options *CatalogDevBoxDefinitionsClientListByCatalogOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevCenter/devcenters/{devCenterName}/catalogs/{catalogName}/devboxdefinitions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if devCenterName == "" {
		return nil, errors.New("parameter devCenterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devCenterName}", url.PathEscape(devCenterName))
	if catalogName == "" {
		return nil, errors.New("parameter catalogName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{catalogName}", url.PathEscape(catalogName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCatalogHandleResponse handles the ListByCatalog response.
func (client *CatalogDevBoxDefinitionsClient) listByCatalogHandleResponse(resp *http.Response) (CatalogDevBoxDefinitionsClientListByCatalogResponse, error) {
	result := CatalogDevBoxDefinitionsClientListByCatalogResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DevBoxDefinitionListResult); err != nil {
		return CatalogDevBoxDefinitionsClientListByCatalogResponse{}, err
	}
	return result, nil
}
