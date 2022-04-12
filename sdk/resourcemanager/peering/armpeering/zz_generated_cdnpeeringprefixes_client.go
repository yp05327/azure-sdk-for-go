//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CdnPeeringPrefixesClient contains the methods for the CdnPeeringPrefixes group.
// Don't use this type directly, use NewCdnPeeringPrefixesClient() instead.
type CdnPeeringPrefixesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewCdnPeeringPrefixesClient creates a new instance of CdnPeeringPrefixesClient with the specified values.
// subscriptionID - The Azure subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewCdnPeeringPrefixesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CdnPeeringPrefixesClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &CdnPeeringPrefixesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Lists all of the advertised prefixes for the specified peering location
// If the operation fails it returns an *azcore.ResponseError type.
// peeringLocation - The peering location.
// options - CdnPeeringPrefixesClientListOptions contains the optional parameters for the CdnPeeringPrefixesClient.List method.
func (client *CdnPeeringPrefixesClient) List(peeringLocation string, options *CdnPeeringPrefixesClientListOptions) *runtime.Pager[CdnPeeringPrefixesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[CdnPeeringPrefixesClientListResponse]{
		More: func(page CdnPeeringPrefixesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CdnPeeringPrefixesClientListResponse) (CdnPeeringPrefixesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, peeringLocation, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CdnPeeringPrefixesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return CdnPeeringPrefixesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CdnPeeringPrefixesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CdnPeeringPrefixesClient) listCreateRequest(ctx context.Context, peeringLocation string, options *CdnPeeringPrefixesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/cdnPeeringPrefixes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("peeringLocation", peeringLocation)
	reqQP.Set("api-version", "2022-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CdnPeeringPrefixesClient) listHandleResponse(resp *http.Response) (CdnPeeringPrefixesClientListResponse, error) {
	result := CdnPeeringPrefixesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CdnPeeringPrefixListResult); err != nil {
		return CdnPeeringPrefixesClientListResponse{}, err
	}
	return result, nil
}
