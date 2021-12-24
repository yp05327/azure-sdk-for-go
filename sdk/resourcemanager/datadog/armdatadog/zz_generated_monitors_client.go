//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// MonitorsClient contains the methods for the Monitors group.
// Don't use this type directly, use NewMonitorsClient() instead.
type MonitorsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewMonitorsClient creates a new instance of MonitorsClient with the specified values.
func NewMonitorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MonitorsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &MonitorsClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// BeginCreate - Create a monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) BeginCreate(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginCreateOptions) (MonitorsCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsCreatePollerResponse{}, err
	}
	result := MonitorsCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitorsClient.Create", "azure-async-operation", resp, client.pl, client.createHandleError)
	if err != nil {
		return MonitorsCreatePollerResponse{}, err
	}
	result.Poller = &MonitorsCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create a monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) create(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *MonitorsClient) createCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// createHandleError handles the Create error response.
func (client *MonitorsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Delete a monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) BeginDelete(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginDeleteOptions) (MonitorsDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsDeletePollerResponse{}, err
	}
	result := MonitorsDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitorsClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return MonitorsDeletePollerResponse{}, err
	}
	result.Poller = &MonitorsDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete a monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) deleteOperation(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MonitorsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *MonitorsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Get the properties of a specific monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) Get(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsGetOptions) (MonitorsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitorsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitorsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MonitorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MonitorsClient) getHandleResponse(resp *http.Response) (MonitorsGetResponse, error) {
	result := MonitorsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatadogMonitorResource); err != nil {
		return MonitorsGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *MonitorsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetDefaultKey - Get the default api key.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) GetDefaultKey(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsGetDefaultKeyOptions) (MonitorsGetDefaultKeyResponse, error) {
	req, err := client.getDefaultKeyCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsGetDefaultKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitorsGetDefaultKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitorsGetDefaultKeyResponse{}, client.getDefaultKeyHandleError(resp)
	}
	return client.getDefaultKeyHandleResponse(resp)
}

// getDefaultKeyCreateRequest creates the GetDefaultKey request.
func (client *MonitorsClient) getDefaultKeyCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsGetDefaultKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/getDefaultKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDefaultKeyHandleResponse handles the GetDefaultKey response.
func (client *MonitorsClient) getDefaultKeyHandleResponse(resp *http.Response) (MonitorsGetDefaultKeyResponse, error) {
	result := MonitorsGetDefaultKeyResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatadogAPIKey); err != nil {
		return MonitorsGetDefaultKeyResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getDefaultKeyHandleError handles the GetDefaultKey error response.
func (client *MonitorsClient) getDefaultKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - List all monitors under the specified subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) List(options *MonitorsListOptions) *MonitorsListPager {
	return &MonitorsListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp MonitorsListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatadogMonitorResourceListResponse.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *MonitorsClient) listCreateRequest(ctx context.Context, options *MonitorsListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Datadog/monitors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MonitorsClient) listHandleResponse(resp *http.Response) (MonitorsListResponse, error) {
	result := MonitorsListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatadogMonitorResourceListResponse); err != nil {
		return MonitorsListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *MonitorsClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListAPIKeys - List the api keys for a given monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) ListAPIKeys(resourceGroupName string, monitorName string, options *MonitorsListAPIKeysOptions) *MonitorsListAPIKeysPager {
	return &MonitorsListAPIKeysPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listAPIKeysCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp MonitorsListAPIKeysResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatadogAPIKeyListResponse.NextLink)
		},
	}
}

// listAPIKeysCreateRequest creates the ListAPIKeys request.
func (client *MonitorsClient) listAPIKeysCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsListAPIKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/listApiKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAPIKeysHandleResponse handles the ListAPIKeys response.
func (client *MonitorsClient) listAPIKeysHandleResponse(resp *http.Response) (MonitorsListAPIKeysResponse, error) {
	result := MonitorsListAPIKeysResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatadogAPIKeyListResponse); err != nil {
		return MonitorsListAPIKeysResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listAPIKeysHandleError handles the ListAPIKeys error response.
func (client *MonitorsClient) listAPIKeysHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - List all monitors under the specified resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) ListByResourceGroup(resourceGroupName string, options *MonitorsListByResourceGroupOptions) *MonitorsListByResourceGroupPager {
	return &MonitorsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp MonitorsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatadogMonitorResourceListResponse.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *MonitorsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *MonitorsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *MonitorsClient) listByResourceGroupHandleResponse(resp *http.Response) (MonitorsListByResourceGroupResponse, error) {
	result := MonitorsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatadogMonitorResourceListResponse); err != nil {
		return MonitorsListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *MonitorsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListHosts - List the hosts for a given monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) ListHosts(resourceGroupName string, monitorName string, options *MonitorsListHostsOptions) *MonitorsListHostsPager {
	return &MonitorsListHostsPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listHostsCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp MonitorsListHostsResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DatadogHostListResponse.NextLink)
		},
	}
}

// listHostsCreateRequest creates the ListHosts request.
func (client *MonitorsClient) listHostsCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsListHostsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/listHosts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHostsHandleResponse handles the ListHosts response.
func (client *MonitorsClient) listHostsHandleResponse(resp *http.Response) (MonitorsListHostsResponse, error) {
	result := MonitorsListHostsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatadogHostListResponse); err != nil {
		return MonitorsListHostsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHostsHandleError handles the ListHosts error response.
func (client *MonitorsClient) listHostsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListLinkedResources - List all Azure resources associated to the same Datadog organization as the target resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) ListLinkedResources(resourceGroupName string, monitorName string, options *MonitorsListLinkedResourcesOptions) *MonitorsListLinkedResourcesPager {
	return &MonitorsListLinkedResourcesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listLinkedResourcesCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp MonitorsListLinkedResourcesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.LinkedResourceListResponse.NextLink)
		},
	}
}

// listLinkedResourcesCreateRequest creates the ListLinkedResources request.
func (client *MonitorsClient) listLinkedResourcesCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsListLinkedResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/listLinkedResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listLinkedResourcesHandleResponse handles the ListLinkedResources response.
func (client *MonitorsClient) listLinkedResourcesHandleResponse(resp *http.Response) (MonitorsListLinkedResourcesResponse, error) {
	result := MonitorsListLinkedResourcesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedResourceListResponse); err != nil {
		return MonitorsListLinkedResourcesResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listLinkedResourcesHandleError handles the ListLinkedResources error response.
func (client *MonitorsClient) listLinkedResourcesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListMonitoredResources - List the resources currently being monitored by the Datadog monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) ListMonitoredResources(resourceGroupName string, monitorName string, options *MonitorsListMonitoredResourcesOptions) *MonitorsListMonitoredResourcesPager {
	return &MonitorsListMonitoredResourcesPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listMonitoredResourcesCreateRequest(ctx, resourceGroupName, monitorName, options)
		},
		advancer: func(ctx context.Context, resp MonitorsListMonitoredResourcesResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MonitoredResourceListResponse.NextLink)
		},
	}
}

// listMonitoredResourcesCreateRequest creates the ListMonitoredResources request.
func (client *MonitorsClient) listMonitoredResourcesCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsListMonitoredResourcesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/listMonitoredResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listMonitoredResourcesHandleResponse handles the ListMonitoredResources response.
func (client *MonitorsClient) listMonitoredResourcesHandleResponse(resp *http.Response) (MonitorsListMonitoredResourcesResponse, error) {
	result := MonitorsListMonitoredResourcesResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoredResourceListResponse); err != nil {
		return MonitorsListMonitoredResourcesResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listMonitoredResourcesHandleError handles the ListMonitoredResources error response.
func (client *MonitorsClient) listMonitoredResourcesHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// RefreshSetPasswordLink - Refresh the set password link and return a latest one.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) RefreshSetPasswordLink(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsRefreshSetPasswordLinkOptions) (MonitorsRefreshSetPasswordLinkResponse, error) {
	req, err := client.refreshSetPasswordLinkCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsRefreshSetPasswordLinkResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitorsRefreshSetPasswordLinkResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitorsRefreshSetPasswordLinkResponse{}, client.refreshSetPasswordLinkHandleError(resp)
	}
	return client.refreshSetPasswordLinkHandleResponse(resp)
}

// refreshSetPasswordLinkCreateRequest creates the RefreshSetPasswordLink request.
func (client *MonitorsClient) refreshSetPasswordLinkCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsRefreshSetPasswordLinkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/refreshSetPasswordLink"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// refreshSetPasswordLinkHandleResponse handles the RefreshSetPasswordLink response.
func (client *MonitorsClient) refreshSetPasswordLinkHandleResponse(resp *http.Response) (MonitorsRefreshSetPasswordLinkResponse, error) {
	result := MonitorsRefreshSetPasswordLinkResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatadogSetPasswordLink); err != nil {
		return MonitorsRefreshSetPasswordLinkResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// refreshSetPasswordLinkHandleError handles the RefreshSetPasswordLink error response.
func (client *MonitorsClient) refreshSetPasswordLinkHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// SetDefaultKey - Set the default api key.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) SetDefaultKey(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsSetDefaultKeyOptions) (MonitorsSetDefaultKeyResponse, error) {
	req, err := client.setDefaultKeyCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsSetDefaultKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MonitorsSetDefaultKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MonitorsSetDefaultKeyResponse{}, client.setDefaultKeyHandleError(resp)
	}
	return MonitorsSetDefaultKeyResponse{RawResponse: resp}, nil
}

// setDefaultKeyCreateRequest creates the SetDefaultKey request.
func (client *MonitorsClient) setDefaultKeyCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsSetDefaultKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/setDefaultKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// setDefaultKeyHandleError handles the SetDefaultKey error response.
func (client *MonitorsClient) setDefaultKeyHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Update a monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) BeginUpdate(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginUpdateOptions) (MonitorsUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return MonitorsUpdatePollerResponse{}, err
	}
	result := MonitorsUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MonitorsClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return MonitorsUpdatePollerResponse{}, err
	}
	result.Poller = &MonitorsUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Update a monitor resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *MonitorsClient) update(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, monitorName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *MonitorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *MonitorsBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// updateHandleError handles the Update error response.
func (client *MonitorsClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}