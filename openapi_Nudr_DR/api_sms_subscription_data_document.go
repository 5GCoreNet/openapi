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


// SMSSubscriptionDataDocumentApiService SMSSubscriptionDataDocumentApi service
type SMSSubscriptionDataDocumentApiService service

type ApiQuerySmsDataRequest struct {
	ctx context.Context
	ApiService *SMSSubscriptionDataDocumentApiService
	ueId string
	servingPlmnId string
	supportedFeatures *string
	ifNoneMatch *string
	ifModifiedSince *string
}

// Supported Features
func (r ApiQuerySmsDataRequest) SupportedFeatures(supportedFeatures string) ApiQuerySmsDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiQuerySmsDataRequest) IfNoneMatch(ifNoneMatch string) ApiQuerySmsDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiQuerySmsDataRequest) IfModifiedSince(ifModifiedSince string) ApiQuerySmsDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiQuerySmsDataRequest) Execute() (*SmsSubscriptionData, *http.Response, error) {
	return r.ApiService.QuerySmsDataExecute(r)
}

/*
QuerySmsData Retrieves the SMS subscription data of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param servingPlmnId PLMN ID
 @return ApiQuerySmsDataRequest
*/
func (a *SMSSubscriptionDataDocumentApiService) QuerySmsData(ctx context.Context, ueId string, servingPlmnId string) ApiQuerySmsDataRequest {
	return ApiQuerySmsDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		servingPlmnId: servingPlmnId,
	}
}

// Execute executes the request
//  @return SmsSubscriptionData
func (a *SMSSubscriptionDataDocumentApiService) QuerySmsDataExecute(r ApiQuerySmsDataRequest) (*SmsSubscriptionData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmsSubscriptionData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSSubscriptionDataDocumentApiService.QuerySmsData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/{servingPlmnId}/provisioned-data/sms-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"servingPlmnId"+"}", url.PathEscape(parameterValueToString(r.servingPlmnId, "servingPlmnId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
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
