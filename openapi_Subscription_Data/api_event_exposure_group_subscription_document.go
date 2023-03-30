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


// EventExposureGroupSubscriptionDocumentApiService EventExposureGroupSubscriptionDocumentApi service
type EventExposureGroupSubscriptionDocumentApiService service

type ApiModifyEeGroupSubscriptionRequest struct {
	ctx context.Context
	ApiService *EventExposureGroupSubscriptionDocumentApiService
	ueGroupId string
	subsId string
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifyEeGroupSubscriptionRequest) PatchItem(patchItem []PatchItem) ApiModifyEeGroupSubscriptionRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifyEeGroupSubscriptionRequest) SupportedFeatures(supportedFeatures string) ApiModifyEeGroupSubscriptionRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifyEeGroupSubscriptionRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.ModifyEeGroupSubscriptionExecute(r)
}

/*
ModifyEeGroupSubscription Modify an individual ee subscription for a group of a UEs

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueGroupId
 @param subsId
 @return ApiModifyEeGroupSubscriptionRequest
*/
func (a *EventExposureGroupSubscriptionDocumentApiService) ModifyEeGroupSubscription(ctx context.Context, ueGroupId string, subsId string) ApiModifyEeGroupSubscriptionRequest {
	return ApiModifyEeGroupSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		ueGroupId: ueGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return PatchResult
func (a *EventExposureGroupSubscriptionDocumentApiService) ModifyEeGroupSubscriptionExecute(r ApiModifyEeGroupSubscriptionRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventExposureGroupSubscriptionDocumentApiService.ModifyEeGroupSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", url.PathEscape(parameterValueToString(r.ueGroupId, "ueGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchItem == nil {
		return localVarReturnValue, nil, reportError("patchItem is required and must be specified")
	}
	if len(*r.patchItem) < 1 {
		return localVarReturnValue, nil, reportError("patchItem must have at least 1 elements")
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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiQueryEeGroupSubscriptionRequest struct {
	ctx context.Context
	ApiService *EventExposureGroupSubscriptionDocumentApiService
	ueGroupId string
	subsId string
}

func (r ApiQueryEeGroupSubscriptionRequest) Execute() ([]EeSubscription, *http.Response, error) {
	return r.ApiService.QueryEeGroupSubscriptionExecute(r)
}

/*
QueryEeGroupSubscription Retrieve a individual eeSubscription for a group of UEs or any UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueGroupId
 @param subsId Unique ID of the subscription to remove
 @return ApiQueryEeGroupSubscriptionRequest
*/
func (a *EventExposureGroupSubscriptionDocumentApiService) QueryEeGroupSubscription(ctx context.Context, ueGroupId string, subsId string) ApiQueryEeGroupSubscriptionRequest {
	return ApiQueryEeGroupSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		ueGroupId: ueGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return []EeSubscription
func (a *EventExposureGroupSubscriptionDocumentApiService) QueryEeGroupSubscriptionExecute(r ApiQueryEeGroupSubscriptionRequest) ([]EeSubscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []EeSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventExposureGroupSubscriptionDocumentApiService.QueryEeGroupSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", url.PathEscape(parameterValueToString(r.ueGroupId, "ueGroupId")), -1)
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

type ApiRemoveEeGroupSubscriptionsRequest struct {
	ctx context.Context
	ApiService *EventExposureGroupSubscriptionDocumentApiService
	ueGroupId string
	subsId string
}

func (r ApiRemoveEeGroupSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveEeGroupSubscriptionsExecute(r)
}

/*
RemoveEeGroupSubscriptions Deletes a eeSubscription for a group of UEs or any UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueGroupId
 @param subsId Unique ID of the subscription to remove
 @return ApiRemoveEeGroupSubscriptionsRequest
*/
func (a *EventExposureGroupSubscriptionDocumentApiService) RemoveEeGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) ApiRemoveEeGroupSubscriptionsRequest {
	return ApiRemoveEeGroupSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		ueGroupId: ueGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
func (a *EventExposureGroupSubscriptionDocumentApiService) RemoveEeGroupSubscriptionsExecute(r ApiRemoveEeGroupSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventExposureGroupSubscriptionDocumentApiService.RemoveEeGroupSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", url.PathEscape(parameterValueToString(r.ueGroupId, "ueGroupId")), -1)
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

type ApiUpdateEeGroupSubscriptionsRequest struct {
	ctx context.Context
	ApiService *EventExposureGroupSubscriptionDocumentApiService
	ueGroupId string
	subsId string
	eeSubscription *EeSubscription
}

func (r ApiUpdateEeGroupSubscriptionsRequest) EeSubscription(eeSubscription EeSubscription) ApiUpdateEeGroupSubscriptionsRequest {
	r.eeSubscription = &eeSubscription
	return r
}

func (r ApiUpdateEeGroupSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateEeGroupSubscriptionsExecute(r)
}

/*
UpdateEeGroupSubscriptions Update an individual ee subscription of a group of UEs or any UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueGroupId
 @param subsId
 @return ApiUpdateEeGroupSubscriptionsRequest
*/
func (a *EventExposureGroupSubscriptionDocumentApiService) UpdateEeGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) ApiUpdateEeGroupSubscriptionsRequest {
	return ApiUpdateEeGroupSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		ueGroupId: ueGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
func (a *EventExposureGroupSubscriptionDocumentApiService) UpdateEeGroupSubscriptionsExecute(r ApiUpdateEeGroupSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventExposureGroupSubscriptionDocumentApiService.UpdateEeGroupSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", url.PathEscape(parameterValueToString(r.ueGroupId, "ueGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.eeSubscription == nil {
		return nil, reportError("eeSubscription is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.eeSubscription
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v ProblemDetails
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
