//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatafactory

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
	"strconv"
	"strings"
)

// PipelinesClient contains the methods for the Pipelines group.
// Don't use this type directly, use NewPipelinesClient() instead.
type PipelinesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewPipelinesClient creates a new instance of PipelinesClient with the specified values.
func NewPipelinesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PipelinesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &PipelinesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CreateOrUpdate - Creates or updates a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *PipelinesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline PipelineResource, options *PipelinesCreateOrUpdateOptions) (PipelinesCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, pipeline, options)
	if err != nil {
		return PipelinesCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelinesCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelinesCreateOrUpdateResponse{}, client.createOrUpdateHandleError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PipelinesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, pipeline PipelineResource, options *PipelinesCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header.Set("If-Match", *options.IfMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, pipeline)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PipelinesClient) createOrUpdateHandleResponse(resp *http.Response) (PipelinesCreateOrUpdateResponse, error) {
	result := PipelinesCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineResource); err != nil {
		return PipelinesCreateOrUpdateResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *PipelinesClient) createOrUpdateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// CreateRun - Creates a run of a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *PipelinesClient) CreateRun(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesCreateRunOptions) (PipelinesCreateRunResponse, error) {
	req, err := client.createRunCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesCreateRunResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelinesCreateRunResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PipelinesCreateRunResponse{}, client.createRunHandleError(resp)
	}
	return client.createRunHandleResponse(resp)
}

// createRunCreateRequest creates the CreateRun request.
func (client *PipelinesClient) createRunCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesCreateRunOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}/createRun"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	if options != nil && options.ReferencePipelineRunID != nil {
		reqQP.Set("referencePipelineRunId", *options.ReferencePipelineRunID)
	}
	if options != nil && options.IsRecovery != nil {
		reqQP.Set("isRecovery", strconv.FormatBool(*options.IsRecovery))
	}
	if options != nil && options.StartActivityName != nil {
		reqQP.Set("startActivityName", *options.StartActivityName)
	}
	if options != nil && options.StartFromFailure != nil {
		reqQP.Set("startFromFailure", strconv.FormatBool(*options.StartFromFailure))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Parameters != nil {
		return req, runtime.MarshalAsJSON(req, options.Parameters)
	}
	return req, nil
}

// createRunHandleResponse handles the CreateRun response.
func (client *PipelinesClient) createRunHandleResponse(resp *http.Response) (PipelinesCreateRunResponse, error) {
	result := PipelinesCreateRunResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CreateRunResponse); err != nil {
		return PipelinesCreateRunResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// createRunHandleError handles the CreateRun error response.
func (client *PipelinesClient) createRunHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Deletes a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *PipelinesClient) Delete(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesDeleteOptions) (PipelinesDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelinesDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return PipelinesDeleteResponse{}, client.deleteHandleError(resp)
	}
	return PipelinesDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PipelinesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *PipelinesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a pipeline.
// If the operation fails it returns the *CloudError error type.
func (client *PipelinesClient) Get(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesGetOptions) (PipelinesGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, factoryName, pipelineName, options)
	if err != nil {
		return PipelinesGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PipelinesGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNotModified) {
		return PipelinesGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PipelinesClient) getCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, pipelineName string, options *PipelinesGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines/{pipelineName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	if pipelineName == "" {
		return nil, errors.New("parameter pipelineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pipelineName}", url.PathEscape(pipelineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfNoneMatch != nil {
		req.Raw().Header.Set("If-None-Match", *options.IfNoneMatch)
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PipelinesClient) getHandleResponse(resp *http.Response) (PipelinesGetResponse, error) {
	result := PipelinesGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineResource); err != nil {
		return PipelinesGetResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *PipelinesClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByFactory - Lists pipelines.
// If the operation fails it returns the *CloudError error type.
func (client *PipelinesClient) ListByFactory(resourceGroupName string, factoryName string, options *PipelinesListByFactoryOptions) *PipelinesListByFactoryPager {
	return &PipelinesListByFactoryPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByFactoryCreateRequest(ctx, resourceGroupName, factoryName, options)
		},
		advancer: func(ctx context.Context, resp PipelinesListByFactoryResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.PipelineListResponse.NextLink)
		},
	}
}

// listByFactoryCreateRequest creates the ListByFactory request.
func (client *PipelinesClient) listByFactoryCreateRequest(ctx context.Context, resourceGroupName string, factoryName string, options *PipelinesListByFactoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataFactory/factories/{factoryName}/pipelines"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if factoryName == "" {
		return nil, errors.New("parameter factoryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{factoryName}", url.PathEscape(factoryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByFactoryHandleResponse handles the ListByFactory response.
func (client *PipelinesClient) listByFactoryHandleResponse(resp *http.Response) (PipelinesListByFactoryResponse, error) {
	result := PipelinesListByFactoryResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PipelineListResponse); err != nil {
		return PipelinesListByFactoryResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByFactoryHandleError handles the ListByFactory error response.
func (client *PipelinesClient) listByFactoryHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}