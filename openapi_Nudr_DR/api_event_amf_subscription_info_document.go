/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// EventAMFSubscriptionInfoDocumentApiService EventAMFSubscriptionInfoDocumentApi service
type EventAMFSubscriptionInfoDocumentApiService service

type ApiRemoveAmfGroupSubscriptionsRequest struct {
	ctx context.Context
	ApiService *EventAMFSubscriptionInfoDocumentApiService
	ueGroupId string
	subsId string
}

func (r ApiRemoveAmfGroupSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveAmfGroupSubscriptionsExecute(r)
}

/*
RemoveAmfGroupSubscriptions Deletes AMF Subscription Info for an eeSubscription for a group of UEs or any UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueGroupId
 @param subsId
 @return ApiRemoveAmfGroupSubscriptionsRequest
*/
func (a *EventAMFSubscriptionInfoDocumentApiService) RemoveAmfGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) ApiRemoveAmfGroupSubscriptionsRequest {
	return ApiRemoveAmfGroupSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		ueGroupId: ueGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
func (a *EventAMFSubscriptionInfoDocumentApiService) RemoveAmfGroupSubscriptionsExecute(r ApiRemoveAmfGroupSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventAMFSubscriptionInfoDocumentApiService.RemoveAmfGroupSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/amf-subscriptions"
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

type ApiRemoveAmfSubscriptionsInfoRequest struct {
	ctx context.Context
	ApiService *EventAMFSubscriptionInfoDocumentApiService
	ueId string
	subsId string
}

func (r ApiRemoveAmfSubscriptionsInfoRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveAmfSubscriptionsInfoExecute(r)
}

/*
RemoveAmfSubscriptionsInfo Deletes AMF Subscription Info for an eeSubscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId
 @param subsId
 @return ApiRemoveAmfSubscriptionsInfoRequest
*/
func (a *EventAMFSubscriptionInfoDocumentApiService) RemoveAmfSubscriptionsInfo(ctx context.Context, ueId string, subsId string) ApiRemoveAmfSubscriptionsInfoRequest {
	return ApiRemoveAmfSubscriptionsInfoRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		subsId: subsId,
	}
}

// Execute executes the request
func (a *EventAMFSubscriptionInfoDocumentApiService) RemoveAmfSubscriptionsInfoExecute(r ApiRemoveAmfSubscriptionsInfoRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EventAMFSubscriptionInfoDocumentApiService.RemoveAmfSubscriptionsInfo")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/ee-subscriptions/{subsId}/amf-subscriptions"
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