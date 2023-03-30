/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// HSSEventSubscriptionInfoDocumentApiService HSSEventSubscriptionInfoDocumentApi service
type HSSEventSubscriptionInfoDocumentApiService service

type ApiCreateHSSSubscriptionsRequest struct {
	ctx context.Context
	ApiService *HSSEventSubscriptionInfoDocumentApiService
	ueId string
	subsId string
	hssSubscriptionInfo *HssSubscriptionInfo
}

func (r ApiCreateHSSSubscriptionsRequest) HssSubscriptionInfo(hssSubscriptionInfo HssSubscriptionInfo) ApiCreateHSSSubscriptionsRequest {
	r.hssSubscriptionInfo = &hssSubscriptionInfo
	return r
}

func (r ApiCreateHSSSubscriptionsRequest) Execute() (*HssSubscriptionInfo, *http.Response, error) {
	return r.ApiService.CreateHSSSubscriptionsExecute(r)
}

/*
CreateHSSSubscriptions Create HSS Subscription Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param subsId
 @return ApiCreateHSSSubscriptionsRequest
*/
func (a *HSSEventSubscriptionInfoDocumentApiService) CreateHSSSubscriptions(ctx context.Context, ueId string, subsId string) ApiCreateHSSSubscriptionsRequest {
	return ApiCreateHSSSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return HssSubscriptionInfo
func (a *HSSEventSubscriptionInfoDocumentApiService) CreateHSSSubscriptionsExecute(r ApiCreateHSSSubscriptionsRequest) (*HssSubscriptionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HssSubscriptionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HSSEventSubscriptionInfoDocumentApiService.CreateHSSSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.hssSubscriptionInfo == nil {
		return localVarReturnValue, nil, reportError("hssSubscriptionInfo is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.hssSubscriptionInfo
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetHssGroupSubscriptionsRequest struct {
	ctx context.Context
	ApiService *HSSEventSubscriptionInfoDocumentApiService
	externalGroupId string
	subsId string
}

func (r ApiGetHssGroupSubscriptionsRequest) Execute() (*HssSubscriptionInfo, *http.Response, error) {
	return r.ApiService.GetHssGroupSubscriptionsExecute(r)
}

/*
GetHssGroupSubscriptions Retrieve HSS Subscription Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGroupId
 @param subsId
 @return ApiGetHssGroupSubscriptionsRequest
*/
func (a *HSSEventSubscriptionInfoDocumentApiService) GetHssGroupSubscriptions(ctx context.Context, externalGroupId string, subsId string) ApiGetHssGroupSubscriptionsRequest {
	return ApiGetHssGroupSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		externalGroupId: externalGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return HssSubscriptionInfo
func (a *HSSEventSubscriptionInfoDocumentApiService) GetHssGroupSubscriptionsExecute(r ApiGetHssGroupSubscriptionsRequest) (*HssSubscriptionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *HssSubscriptionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HSSEventSubscriptionInfoDocumentApiService.GetHssGroupSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", url.PathEscape(parameterValueToString(r.externalGroupId, "externalGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetHssSubscriptionInfoRequest struct {
	ctx context.Context
	ApiService *HSSEventSubscriptionInfoDocumentApiService
	ueId string
	subsId string
}

func (r ApiGetHssSubscriptionInfoRequest) Execute() (*SmfSubscriptionInfo, *http.Response, error) {
	return r.ApiService.GetHssSubscriptionInfoExecute(r)
}

/*
GetHssSubscriptionInfo Retrieve HSS Subscription Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param subsId
 @return ApiGetHssSubscriptionInfoRequest
*/
func (a *HSSEventSubscriptionInfoDocumentApiService) GetHssSubscriptionInfo(ctx context.Context, ueId string, subsId string) ApiGetHssSubscriptionInfoRequest {
	return ApiGetHssSubscriptionInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return SmfSubscriptionInfo
func (a *HSSEventSubscriptionInfoDocumentApiService) GetHssSubscriptionInfoExecute(r ApiGetHssSubscriptionInfoRequest) (*SmfSubscriptionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmfSubscriptionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HSSEventSubscriptionInfoDocumentApiService.GetHssSubscriptionInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiModifyHssGroupSubscriptionsRequest struct {
	ctx context.Context
	ApiService *HSSEventSubscriptionInfoDocumentApiService
	externalGroupId string
	subsId string
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifyHssGroupSubscriptionsRequest) PatchItem(patchItem []PatchItem) ApiModifyHssGroupSubscriptionsRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifyHssGroupSubscriptionsRequest) SupportedFeatures(supportedFeatures string) ApiModifyHssGroupSubscriptionsRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifyHssGroupSubscriptionsRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.ModifyHssGroupSubscriptionsExecute(r)
}

/*
ModifyHssGroupSubscriptions Modify HSS Subscription Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGroupId
 @param subsId
 @return ApiModifyHssGroupSubscriptionsRequest
*/
func (a *HSSEventSubscriptionInfoDocumentApiService) ModifyHssGroupSubscriptions(ctx context.Context, externalGroupId string, subsId string) ApiModifyHssGroupSubscriptionsRequest {
	return ApiModifyHssGroupSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		externalGroupId: externalGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return PatchResult
func (a *HSSEventSubscriptionInfoDocumentApiService) ModifyHssGroupSubscriptionsExecute(r ApiModifyHssGroupSubscriptionsRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HSSEventSubscriptionInfoDocumentApiService.ModifyHssGroupSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", url.PathEscape(parameterValueToString(r.externalGroupId, "externalGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchItem == nil {
		return localVarReturnValue, nil, reportError("patchItem is required and must be specified")
	}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.patchItem
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiModifyHssSubscriptionInfoRequest struct {
	ctx context.Context
	ApiService *HSSEventSubscriptionInfoDocumentApiService
	ueId string
	subsId string
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifyHssSubscriptionInfoRequest) PatchItem(patchItem []PatchItem) ApiModifyHssSubscriptionInfoRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifyHssSubscriptionInfoRequest) SupportedFeatures(supportedFeatures string) ApiModifyHssSubscriptionInfoRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifyHssSubscriptionInfoRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.ModifyHssSubscriptionInfoExecute(r)
}

/*
ModifyHssSubscriptionInfo Modify HSS Subscription Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param subsId
 @return ApiModifyHssSubscriptionInfoRequest
*/
func (a *HSSEventSubscriptionInfoDocumentApiService) ModifyHssSubscriptionInfo(ctx context.Context, ueId string, subsId string) ApiModifyHssSubscriptionInfoRequest {
	return ApiModifyHssSubscriptionInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return PatchResult
func (a *HSSEventSubscriptionInfoDocumentApiService) ModifyHssSubscriptionInfoExecute(r ApiModifyHssSubscriptionInfoRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HSSEventSubscriptionInfoDocumentApiService.ModifyHssSubscriptionInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchItem == nil {
		return localVarReturnValue, nil, reportError("patchItem is required and must be specified")
	}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json-patch+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.patchItem
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveHssGroupSubscriptionsRequest struct {
	ctx context.Context
	ApiService *HSSEventSubscriptionInfoDocumentApiService
	externalGroupId string
	subsId string
}

func (r ApiRemoveHssGroupSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveHssGroupSubscriptionsExecute(r)
}

/*
RemoveHssGroupSubscriptions Delete HSS Subscription Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param externalGroupId
 @param subsId
 @return ApiRemoveHssGroupSubscriptionsRequest
*/
func (a *HSSEventSubscriptionInfoDocumentApiService) RemoveHssGroupSubscriptions(ctx context.Context, externalGroupId string, subsId string) ApiRemoveHssGroupSubscriptionsRequest {
	return ApiRemoveHssGroupSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		externalGroupId: externalGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
func (a *HSSEventSubscriptionInfoDocumentApiService) RemoveHssGroupSubscriptionsExecute(r ApiRemoveHssGroupSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HSSEventSubscriptionInfoDocumentApiService.RemoveHssGroupSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", url.PathEscape(parameterValueToString(r.externalGroupId, "externalGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRemoveHssSubscriptionsInfoRequest struct {
	ctx context.Context
	ApiService *HSSEventSubscriptionInfoDocumentApiService
	ueId string
	subsId string
}

func (r ApiRemoveHssSubscriptionsInfoRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveHssSubscriptionsInfoExecute(r)
}

/*
RemoveHssSubscriptionsInfo Delete HSS Subscription Info

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param subsId
 @return ApiRemoveHssSubscriptionsInfoRequest
*/
func (a *HSSEventSubscriptionInfoDocumentApiService) RemoveHssSubscriptionsInfo(ctx context.Context, ueId string, subsId string) ApiRemoveHssSubscriptionsInfoRequest {
	return ApiRemoveHssSubscriptionsInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		subsId: subsId,
	}
}

// Execute executes the request
func (a *HSSEventSubscriptionInfoDocumentApiService) RemoveHssSubscriptionsInfoExecute(r ApiRemoveHssSubscriptionsInfoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HSSEventSubscriptionInfoDocumentApiService.RemoveHssSubscriptionsInfo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
