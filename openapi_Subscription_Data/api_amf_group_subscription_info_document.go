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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// AMFGroupSubscriptionInfoDocumentApiService AMFGroupSubscriptionInfoDocumentApi service
type AMFGroupSubscriptionInfoDocumentApiService service

type ApiCreateAmfGroupSubscriptionsRequest struct {
	ctx context.Context
	ApiService *AMFGroupSubscriptionInfoDocumentApiService
	ueGroupId string
	subsId string
	amfSubscriptionInfo *[]AmfSubscriptionInfo
}

func (r ApiCreateAmfGroupSubscriptionsRequest) AmfSubscriptionInfo(amfSubscriptionInfo []AmfSubscriptionInfo) ApiCreateAmfGroupSubscriptionsRequest {
	r.amfSubscriptionInfo = &amfSubscriptionInfo
	return r
}

func (r ApiCreateAmfGroupSubscriptionsRequest) Execute() ([]AmfSubscriptionInfo, *http.Response, error) {
	return r.ApiService.CreateAmfGroupSubscriptionsExecute(r)
}

/*
CreateAmfGroupSubscriptions Create AmfSubscriptions for a group of UEs or any UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueGroupId
 @param subsId
 @return ApiCreateAmfGroupSubscriptionsRequest
*/
func (a *AMFGroupSubscriptionInfoDocumentApiService) CreateAmfGroupSubscriptions(ctx context.Context, ueGroupId string, subsId string) ApiCreateAmfGroupSubscriptionsRequest {
	return ApiCreateAmfGroupSubscriptionsRequest{
		ApiService: a,
		ctx: ctx,
		ueGroupId: ueGroupId,
		subsId: subsId,
	}
}

// Execute executes the request
//  @return []AmfSubscriptionInfo
func (a *AMFGroupSubscriptionInfoDocumentApiService) CreateAmfGroupSubscriptionsExecute(r ApiCreateAmfGroupSubscriptionsRequest) ([]AmfSubscriptionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []AmfSubscriptionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AMFGroupSubscriptionInfoDocumentApiService.CreateAmfGroupSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/amf-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"ueGroupId"+"}", url.PathEscape(parameterValueToString(r.ueGroupId, "ueGroupId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.amfSubscriptionInfo == nil {
		return localVarReturnValue, nil, reportError("amfSubscriptionInfo is required and must be specified")
	}
	if len(*r.amfSubscriptionInfo) < 1 {
		return localVarReturnValue, nil, reportError("amfSubscriptionInfo must have at least 1 elements")
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
	localVarPostBody = r.amfSubscriptionInfo
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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
