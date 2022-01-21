//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// WebTestsClient contains the methods for the WebTests group.
// Don't use this type directly, use NewWebTestsClient() instead.
type WebTestsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWebTestsClient creates a new instance of WebTestsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWebTestsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WebTestsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &WebTestsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Creates or updates an Application Insights web test definition.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// webTestName - The name of the Application Insights webtest resource.
// webTestDefinition - Properties that need to be specified to create or update an Application Insights web test definition.
// options - WebTestsClientCreateOrUpdateOptions contains the optional parameters for the WebTestsClient.CreateOrUpdate method.
func (client *WebTestsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, webTestName string, webTestDefinition WebTest, options *WebTestsClientCreateOrUpdateOptions) (WebTestsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, webTestName, webTestDefinition, options)
	if err != nil {
		return WebTestsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebTestsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WebTestsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, webTestDefinition WebTest, options *WebTestsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, webTestDefinition)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WebTestsClient) createOrUpdateHandleResponse(resp *http.Response) (WebTestsClientCreateOrUpdateResponse, error) {
	result := WebTestsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Application Insights web test.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// webTestName - The name of the Application Insights webtest resource.
// options - WebTestsClientDeleteOptions contains the optional parameters for the WebTestsClient.Delete method.
func (client *WebTestsClient) Delete(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsClientDeleteOptions) (WebTestsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, webTestName, options)
	if err != nil {
		return WebTestsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WebTestsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WebTestsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebTestsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get a specific Application Insights web test definition.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// webTestName - The name of the Application Insights webtest resource.
// options - WebTestsClientGetOptions contains the optional parameters for the WebTestsClient.Get method.
func (client *WebTestsClient) Get(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsClientGetOptions) (WebTestsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, webTestName, options)
	if err != nil {
		return WebTestsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebTestsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WebTestsClient) getCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, options *WebTestsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebTestsClient) getHandleResponse(resp *http.Response) (WebTestsClientGetResponse, error) {
	result := WebTestsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsClientGetResponse{}, err
	}
	return result, nil
}

// List - Get all Application Insights web test alerts definitions within a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - WebTestsClientListOptions contains the optional parameters for the WebTestsClient.List method.
func (client *WebTestsClient) List(options *WebTestsClientListOptions) *WebTestsClientListPager {
	return &WebTestsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp WebTestsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebTestListResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *WebTestsClient) listCreateRequest(ctx context.Context, options *WebTestsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Insights/webtests"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WebTestsClient) listHandleResponse(resp *http.Response) (WebTestsClientListResponse, error) {
	result := WebTestsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsClientListResponse{}, err
	}
	return result, nil
}

// ListByComponent - Get all Application Insights web tests defined for the specified component.
// If the operation fails it returns an *azcore.ResponseError type.
// componentName - The name of the Application Insights component resource.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - WebTestsClientListByComponentOptions contains the optional parameters for the WebTestsClient.ListByComponent
// method.
func (client *WebTestsClient) ListByComponent(componentName string, resourceGroupName string, options *WebTestsClientListByComponentOptions) *WebTestsClientListByComponentPager {
	return &WebTestsClientListByComponentPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByComponentCreateRequest(ctx, componentName, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp WebTestsClientListByComponentResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebTestListResult.NextLink)
		},
	}
}

// listByComponentCreateRequest creates the ListByComponent request.
func (client *WebTestsClient) listByComponentCreateRequest(ctx context.Context, componentName string, resourceGroupName string, options *WebTestsClientListByComponentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{componentName}/webtests"
	if componentName == "" {
		return nil, errors.New("parameter componentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{componentName}", url.PathEscape(componentName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByComponentHandleResponse handles the ListByComponent response.
func (client *WebTestsClient) listByComponentHandleResponse(resp *http.Response) (WebTestsClientListByComponentResponse, error) {
	result := WebTestsClientListByComponentResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsClientListByComponentResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Get all Application Insights web tests defined within a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// options - WebTestsClientListByResourceGroupOptions contains the optional parameters for the WebTestsClient.ListByResourceGroup
// method.
func (client *WebTestsClient) ListByResourceGroup(resourceGroupName string, options *WebTestsClientListByResourceGroupOptions) *WebTestsClientListByResourceGroupPager {
	return &WebTestsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp WebTestsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WebTestListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WebTestsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WebTestsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WebTestsClient) listByResourceGroupHandleResponse(resp *http.Response) (WebTestsClientListByResourceGroupResponse, error) {
	result := WebTestsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTestListResult); err != nil {
		return WebTestsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Creates or updates an Application Insights web test definition.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// webTestName - The name of the Application Insights webtest resource.
// webTestTags - Updated tag information to set into the web test instance.
// options - WebTestsClientUpdateTagsOptions contains the optional parameters for the WebTestsClient.UpdateTags method.
func (client *WebTestsClient) UpdateTags(ctx context.Context, resourceGroupName string, webTestName string, webTestTags TagsResource, options *WebTestsClientUpdateTagsOptions) (WebTestsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, webTestName, webTestTags, options)
	if err != nil {
		return WebTestsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WebTestsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WebTestsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *WebTestsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, webTestName string, webTestTags TagsResource, options *WebTestsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/webtests/{webTestName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if webTestName == "" {
		return nil, errors.New("parameter webTestName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webTestName}", url.PathEscape(webTestName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, webTestTags)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *WebTestsClient) updateTagsHandleResponse(resp *http.Response) (WebTestsClientUpdateTagsResponse, error) {
	result := WebTestsClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebTest); err != nil {
		return WebTestsClientUpdateTagsResponse{}, err
	}
	return result, nil
}