//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcemover

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
	"strings"
)

// MoveResourcesClient contains the methods for the MoveResources group.
// Don't use this type directly, use NewMoveResourcesClient() instead.
type MoveResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewMoveResourcesClient creates a new instance of MoveResourcesClient with the specified values.
// subscriptionID - The Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewMoveResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *MoveResourcesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &MoveResourcesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreate - Creates or updates a Move Resource in the move collection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The Resource Group Name.
// moveCollectionName - The Move Collection Name.
// moveResourceName - The Move Resource Name.
// options - MoveResourcesClientBeginCreateOptions contains the optional parameters for the MoveResourcesClient.BeginCreate
// method.
func (client *MoveResourcesClient) BeginCreate(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, options *MoveResourcesClientBeginCreateOptions) (MoveResourcesClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, moveCollectionName, moveResourceName, options)
	if err != nil {
		return MoveResourcesClientCreatePollerResponse{}, err
	}
	result := MoveResourcesClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MoveResourcesClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return MoveResourcesClientCreatePollerResponse{}, err
	}
	result.Poller = &MoveResourcesClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Creates or updates a Move Resource in the move collection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MoveResourcesClient) create(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, options *MoveResourcesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, moveCollectionName, moveResourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *MoveResourcesClient) createCreateRequest(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, options *MoveResourcesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources/{moveResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if moveCollectionName == "" {
		return nil, errors.New("parameter moveCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveCollectionName}", url.PathEscape(moveCollectionName))
	if moveResourceName == "" {
		return nil, errors.New("parameter moveResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveResourceName}", url.PathEscape(moveResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Body != nil {
		return req, runtime.MarshalAsJSON(req, *options.Body)
	}
	return req, nil
}

// BeginDelete - Deletes a Move Resource from the move collection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The Resource Group Name.
// moveCollectionName - The Move Collection Name.
// moveResourceName - The Move Resource Name.
// options - MoveResourcesClientBeginDeleteOptions contains the optional parameters for the MoveResourcesClient.BeginDelete
// method.
func (client *MoveResourcesClient) BeginDelete(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, options *MoveResourcesClientBeginDeleteOptions) (MoveResourcesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, moveCollectionName, moveResourceName, options)
	if err != nil {
		return MoveResourcesClientDeletePollerResponse{}, err
	}
	result := MoveResourcesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("MoveResourcesClient.Delete", "azure-async-operation", resp, client.pl)
	if err != nil {
		return MoveResourcesClientDeletePollerResponse{}, err
	}
	result.Poller = &MoveResourcesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a Move Resource from the move collection.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *MoveResourcesClient) deleteOperation(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, options *MoveResourcesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, moveCollectionName, moveResourceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MoveResourcesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, options *MoveResourcesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources/{moveResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if moveCollectionName == "" {
		return nil, errors.New("parameter moveCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveCollectionName}", url.PathEscape(moveCollectionName))
	if moveResourceName == "" {
		return nil, errors.New("parameter moveResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveResourceName}", url.PathEscape(moveResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the Move Resource.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The Resource Group Name.
// moveCollectionName - The Move Collection Name.
// moveResourceName - The Move Resource Name.
// options - MoveResourcesClientGetOptions contains the optional parameters for the MoveResourcesClient.Get method.
func (client *MoveResourcesClient) Get(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, options *MoveResourcesClientGetOptions) (MoveResourcesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, moveCollectionName, moveResourceName, options)
	if err != nil {
		return MoveResourcesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return MoveResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MoveResourcesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *MoveResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, moveCollectionName string, moveResourceName string, options *MoveResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources/{moveResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if moveCollectionName == "" {
		return nil, errors.New("parameter moveCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveCollectionName}", url.PathEscape(moveCollectionName))
	if moveResourceName == "" {
		return nil, errors.New("parameter moveResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveResourceName}", url.PathEscape(moveResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MoveResourcesClient) getHandleResponse(resp *http.Response) (MoveResourcesClientGetResponse, error) {
	result := MoveResourcesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MoveResource); err != nil {
		return MoveResourcesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists the Move Resources in the move collection.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The Resource Group Name.
// moveCollectionName - The Move Collection Name.
// options - MoveResourcesClientListOptions contains the optional parameters for the MoveResourcesClient.List method.
func (client *MoveResourcesClient) List(resourceGroupName string, moveCollectionName string, options *MoveResourcesClientListOptions) *MoveResourcesClientListPager {
	return &MoveResourcesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, moveCollectionName, options)
		},
		advancer: func(ctx context.Context, resp MoveResourcesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.MoveResourceCollection.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *MoveResourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, moveCollectionName string, options *MoveResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Migrate/moveCollections/{moveCollectionName}/moveResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if moveCollectionName == "" {
		return nil, errors.New("parameter moveCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{moveCollectionName}", url.PathEscape(moveCollectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-08-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MoveResourcesClient) listHandleResponse(resp *http.Response) (MoveResourcesClientListResponse, error) {
	result := MoveResourcesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.MoveResourceCollection); err != nil {
		return MoveResourcesClientListResponse{}, err
	}
	return result, nil
}