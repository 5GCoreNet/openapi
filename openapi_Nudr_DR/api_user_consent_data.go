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

// UserConsentDataApiService UserConsentDataApi service
type UserConsentDataApiService service

type ApiQueryUserConsentDataRequest struct {
	ctx               context.Context
	ApiService        *UserConsentDataApiService
	ueId              string
	supportedFeatures *string
	ucPurpose         *UcPurpose
	ifNoneMatch       *string
	ifModifiedSince   *string
}

// Supported Features
func (r ApiQueryUserConsentDataRequest) SupportedFeatures(supportedFeatures string) ApiQueryUserConsentDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// User consent purpose
func (r ApiQueryUserConsentDataRequest) UcPurpose(ucPurpose UcPurpose) ApiQueryUserConsentDataRequest {
	r.ucPurpose = &ucPurpose
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiQueryUserConsentDataRequest) IfNoneMatch(ifNoneMatch string) ApiQueryUserConsentDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiQueryUserConsentDataRequest) IfModifiedSince(ifModifiedSince string) ApiQueryUserConsentDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiQueryUserConsentDataRequest) Execute() (*UcSubscriptionData, *http.Response, error) {
	return r.ApiService.QueryUserConsentDataExecute(r)
}

/*
QueryUserConsentData Retrieves the subscribed User Consent Data of a UE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiQueryUserConsentDataRequest
*/
func (a *UserConsentDataApiService) QueryUserConsentData(ctx context.Context, ueId string) ApiQueryUserConsentDataRequest {
	return ApiQueryUserConsentDataRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
//
//	@return UcSubscriptionData
func (a *UserConsentDataApiService) QueryUserConsentDataExecute(r ApiQueryUserConsentDataRequest) (*UcSubscriptionData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *UcSubscriptionData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserConsentDataApiService.QueryUserConsentData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/uc-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	if r.ucPurpose != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ucPurpose", r.ucPurpose, "")
	}
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
	if r.ifNoneMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-None-Match", r.ifNoneMatch, "")
	}
	if r.ifModifiedSince != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Modified-Since", r.ifModifiedSince, "")
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
