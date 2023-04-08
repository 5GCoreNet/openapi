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
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// LCSPrivacySubscriptionDataApiService LCSPrivacySubscriptionDataApi service
type LCSPrivacySubscriptionDataApiService service

type ApiQueryLcsPrivacyDataRequest struct {
	ctx context.Context
	ApiService *LCSPrivacySubscriptionDataApiService
	ueId string
	fields *[]string
	supportedFeatures *string
	ifNoneMatch *string
	ifModifiedSince *string
}

// attributes to be retrieved
func (r ApiQueryLcsPrivacyDataRequest) Fields(fields []string) ApiQueryLcsPrivacyDataRequest {
	r.fields = &fields
	return r
}

// Supported Features
func (r ApiQueryLcsPrivacyDataRequest) SupportedFeatures(supportedFeatures string) ApiQueryLcsPrivacyDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiQueryLcsPrivacyDataRequest) IfNoneMatch(ifNoneMatch string) ApiQueryLcsPrivacyDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiQueryLcsPrivacyDataRequest) IfModifiedSince(ifModifiedSince string) ApiQueryLcsPrivacyDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiQueryLcsPrivacyDataRequest) Execute() (*LcsPrivacyData, *http.Response, error) {
	return r.ApiService.QueryLcsPrivacyDataExecute(r)
}

/*
QueryLcsPrivacyData Retrieves the LCS Privacy subscription data of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiQueryLcsPrivacyDataRequest
*/
func (a *LCSPrivacySubscriptionDataApiService) QueryLcsPrivacyData(ctx context.Context, ueId string) ApiQueryLcsPrivacyDataRequest {
	return ApiQueryLcsPrivacyDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return LcsPrivacyData
func (a *LCSPrivacySubscriptionDataApiService) QueryLcsPrivacyDataExecute(r ApiQueryLcsPrivacyDataRequest) (*LcsPrivacyData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LcsPrivacyData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LCSPrivacySubscriptionDataApiService.QueryLcsPrivacyData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/lcs-privacy-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
	if r.supportedFeatures != nil {
		parameterAddToQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
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
		parameterAddToQuery(localVarQueryParams, "If-None-Match", r.ifNoneMatch, "")
	}
	if r.ifModifiedSince != nil {
		parameterAddToQuery(localVarQueryParams, "If-Modified-Since", r.ifModifiedSince, "")
	}
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
