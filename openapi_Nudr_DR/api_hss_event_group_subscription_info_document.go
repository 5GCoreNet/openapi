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

// HSSEventGroupSubscriptionInfoDocumentApiService HSSEventGroupSubscriptionInfoDocumentApi service
type HSSEventGroupSubscriptionInfoDocumentApiService service

type ApiCreateHssGroupSubscriptionsRequest struct {
	ctx                 context.Context
	ApiService          *HSSEventGroupSubscriptionInfoDocumentApiService
	externalGroupId     string
	subsId              string
	hssSubscriptionInfo *HssSubscriptionInfo
}

func (r ApiCreateHssGroupSubscriptionsRequest) HssSubscriptionInfo(hssSubscriptionInfo HssSubscriptionInfo) ApiCreateHssGroupSubscriptionsRequest {
	r.hssSubscriptionInfo = &hssSubscriptionInfo
	return r
}

func (r ApiCreateHssGroupSubscriptionsRequest) Execute() (*HssSubscriptionInfo, *http.Response, error) {
	return r.ApiService.CreateHssGroupSubscriptionsExecute(r)
}

/*
CreateHssGroupSubscriptions Create HSS Subscription Info for a group of UEs

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalGroupId
	@param subsId
	@return ApiCreateHssGroupSubscriptionsRequest
*/
func (a *HSSEventGroupSubscriptionInfoDocumentApiService) CreateHssGroupSubscriptions(ctx context.Context, externalGroupId string, subsId string) ApiCreateHssGroupSubscriptionsRequest {
	return ApiCreateHssGroupSubscriptionsRequest{
		ApiService:      a,
		ctx:             ctx,
		externalGroupId: externalGroupId,
		subsId:          subsId,
	}
}

// Execute executes the request
//
//	@return HssSubscriptionInfo
func (a *HSSEventGroupSubscriptionInfoDocumentApiService) CreateHssGroupSubscriptionsExecute(r ApiCreateHssGroupSubscriptionsRequest) (*HssSubscriptionInfo, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *HssSubscriptionInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "HSSEventGroupSubscriptionInfoDocumentApiService.CreateHssGroupSubscriptions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/{ueGroupId}/ee-subscriptions/{subsId}/hss-subscriptions"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", url.PathEscape(parameterValueToString(r.externalGroupId, "externalGroupId")), -1)
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
