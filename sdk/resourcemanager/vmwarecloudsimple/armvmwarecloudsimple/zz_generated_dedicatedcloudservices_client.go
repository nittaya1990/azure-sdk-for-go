//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

import (
	"context"
	"errors"
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

// DedicatedCloudServicesClient contains the methods for the DedicatedCloudServices group.
// Don't use this type directly, use NewDedicatedCloudServicesClient() instead.
type DedicatedCloudServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewDedicatedCloudServicesClient creates a new instance of DedicatedCloudServicesClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewDedicatedCloudServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *DedicatedCloudServicesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &DedicatedCloudServicesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Create dedicate cloud service
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group
// dedicatedCloudServiceName - dedicated cloud Service name
// dedicatedCloudServiceRequest - Create Dedicated Cloud Service request
// options - DedicatedCloudServicesClientCreateOrUpdateOptions contains the optional parameters for the DedicatedCloudServicesClient.CreateOrUpdate
// method.
func (client *DedicatedCloudServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest DedicatedCloudService, options *DedicatedCloudServicesClientCreateOrUpdateOptions) (DedicatedCloudServicesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, dedicatedCloudServiceName, dedicatedCloudServiceRequest, options)
	if err != nil {
		return DedicatedCloudServicesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedCloudServicesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedCloudServicesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DedicatedCloudServicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest DedicatedCloudService, options *DedicatedCloudServicesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCloudServiceName == "" {
		return nil, errors.New("parameter dedicatedCloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCloudServiceName}", url.PathEscape(dedicatedCloudServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dedicatedCloudServiceRequest)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DedicatedCloudServicesClient) createOrUpdateHandleResponse(resp *http.Response) (DedicatedCloudServicesClientCreateOrUpdateResponse, error) {
	result := DedicatedCloudServicesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudService); err != nil {
		return DedicatedCloudServicesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Delete dedicate cloud service
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group
// dedicatedCloudServiceName - dedicated cloud service name
// options - DedicatedCloudServicesClientBeginDeleteOptions contains the optional parameters for the DedicatedCloudServicesClient.BeginDelete
// method.
func (client *DedicatedCloudServicesClient) BeginDelete(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, options *DedicatedCloudServicesClientBeginDeleteOptions) (DedicatedCloudServicesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, dedicatedCloudServiceName, options)
	if err != nil {
		return DedicatedCloudServicesClientDeletePollerResponse{}, err
	}
	result := DedicatedCloudServicesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("DedicatedCloudServicesClient.Delete", "", resp, client.pl)
	if err != nil {
		return DedicatedCloudServicesClientDeletePollerResponse{}, err
	}
	result.Poller = &DedicatedCloudServicesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Delete dedicate cloud service
// If the operation fails it returns an *azcore.ResponseError type.
func (client *DedicatedCloudServicesClient) deleteOperation(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, options *DedicatedCloudServicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dedicatedCloudServiceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DedicatedCloudServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, options *DedicatedCloudServicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCloudServiceName == "" {
		return nil, errors.New("parameter dedicatedCloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCloudServiceName}", url.PathEscape(dedicatedCloudServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns Dedicate Cloud Service
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group
// dedicatedCloudServiceName - dedicated cloud Service name
// options - DedicatedCloudServicesClientGetOptions contains the optional parameters for the DedicatedCloudServicesClient.Get
// method.
func (client *DedicatedCloudServicesClient) Get(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, options *DedicatedCloudServicesClientGetOptions) (DedicatedCloudServicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, dedicatedCloudServiceName, options)
	if err != nil {
		return DedicatedCloudServicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedCloudServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedCloudServicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DedicatedCloudServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, options *DedicatedCloudServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCloudServiceName == "" {
		return nil, errors.New("parameter dedicatedCloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCloudServiceName}", url.PathEscape(dedicatedCloudServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DedicatedCloudServicesClient) getHandleResponse(resp *http.Response) (DedicatedCloudServicesClientGetResponse, error) {
	result := DedicatedCloudServicesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudService); err != nil {
		return DedicatedCloudServicesClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Returns list of dedicated cloud services within a resource group
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group
// options - DedicatedCloudServicesClientListByResourceGroupOptions contains the optional parameters for the DedicatedCloudServicesClient.ListByResourceGroup
// method.
func (client *DedicatedCloudServicesClient) ListByResourceGroup(resourceGroupName string, options *DedicatedCloudServicesClientListByResourceGroupOptions) *DedicatedCloudServicesClientListByResourceGroupPager {
	return &DedicatedCloudServicesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp DedicatedCloudServicesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DedicatedCloudServiceListResponse.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *DedicatedCloudServicesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *DedicatedCloudServicesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *DedicatedCloudServicesClient) listByResourceGroupHandleResponse(resp *http.Response) (DedicatedCloudServicesClientListByResourceGroupResponse, error) {
	result := DedicatedCloudServicesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudServiceListResponse); err != nil {
		return DedicatedCloudServicesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Returns list of dedicated cloud services within a subscription
// If the operation fails it returns an *azcore.ResponseError type.
// options - DedicatedCloudServicesClientListBySubscriptionOptions contains the optional parameters for the DedicatedCloudServicesClient.ListBySubscription
// method.
func (client *DedicatedCloudServicesClient) ListBySubscription(options *DedicatedCloudServicesClientListBySubscriptionOptions) *DedicatedCloudServicesClientListBySubscriptionPager {
	return &DedicatedCloudServicesClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp DedicatedCloudServicesClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.DedicatedCloudServiceListResponse.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DedicatedCloudServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *DedicatedCloudServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DedicatedCloudServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (DedicatedCloudServicesClientListBySubscriptionResponse, error) {
	result := DedicatedCloudServicesClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudServiceListResponse); err != nil {
		return DedicatedCloudServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Patch dedicated cloud service's properties
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group
// dedicatedCloudServiceName - dedicated cloud service name
// dedicatedCloudServiceRequest - Patch Dedicated Cloud Service request
// options - DedicatedCloudServicesClientUpdateOptions contains the optional parameters for the DedicatedCloudServicesClient.Update
// method.
func (client *DedicatedCloudServicesClient) Update(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest PatchPayload, options *DedicatedCloudServicesClientUpdateOptions) (DedicatedCloudServicesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dedicatedCloudServiceName, dedicatedCloudServiceRequest, options)
	if err != nil {
		return DedicatedCloudServicesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return DedicatedCloudServicesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DedicatedCloudServicesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *DedicatedCloudServicesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCloudServiceName string, dedicatedCloudServiceRequest PatchPayload, options *DedicatedCloudServicesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.VMwareCloudSimple/dedicatedCloudServices/{dedicatedCloudServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCloudServiceName == "" {
		return nil, errors.New("parameter dedicatedCloudServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCloudServiceName}", url.PathEscape(dedicatedCloudServiceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, dedicatedCloudServiceRequest)
}

// updateHandleResponse handles the Update response.
func (client *DedicatedCloudServicesClient) updateHandleResponse(resp *http.Response) (DedicatedCloudServicesClientUpdateResponse, error) {
	result := DedicatedCloudServicesClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCloudService); err != nil {
		return DedicatedCloudServicesClientUpdateResponse{}, err
	}
	return result, nil
}