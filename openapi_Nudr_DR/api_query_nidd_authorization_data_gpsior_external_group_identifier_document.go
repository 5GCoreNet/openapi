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


// QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApiService QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApi service
type QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApiService service

type ApiGetNiddAuDataRequest struct {
	ctx context.Context
	ApiService *QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApiService
	ueId string
	singleNssai *Snssai
	dnn *string
	mtcProviderInformation *string
	afId *string
	ifNoneMatch *string
	ifModifiedSince *string
}

// single NSSAI
func (r ApiGetNiddAuDataRequest) SingleNssai(singleNssai Snssai) ApiGetNiddAuDataRequest {
	r.singleNssai = &singleNssai
	return r
}

// DNN
func (r ApiGetNiddAuDataRequest) Dnn(dnn string) ApiGetNiddAuDataRequest {
	r.dnn = &dnn
	return r
}

// MTC Provider Information
func (r ApiGetNiddAuDataRequest) MtcProviderInformation(mtcProviderInformation string) ApiGetNiddAuDataRequest {
	r.mtcProviderInformation = &mtcProviderInformation
	return r
}

// Application Function Identifier
func (r ApiGetNiddAuDataRequest) AfId(afId string) ApiGetNiddAuDataRequest {
	r.afId = &afId
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiGetNiddAuDataRequest) IfNoneMatch(ifNoneMatch string) ApiGetNiddAuDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiGetNiddAuDataRequest) IfModifiedSince(ifModifiedSince string) ApiGetNiddAuDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiGetNiddAuDataRequest) Execute() (*AuthorizationData, *http.Response, error) {
	return r.ApiService.GetNiddAuDataExecute(r)
}

/*
GetNiddAuData Retrieve NIDD Authorization Data GPSI or External Group identifier

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE ID
 @return ApiGetNiddAuDataRequest
*/
func (a *QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApiService) GetNiddAuData(ctx context.Context, ueId string) ApiGetNiddAuDataRequest {
	return ApiGetNiddAuDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return AuthorizationData
func (a *QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApiService) GetNiddAuDataExecute(r ApiGetNiddAuDataRequest) (*AuthorizationData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthorizationData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryNIDDAuthorizationDataGPSIOrExternalGroupIdentifierDocumentApiService.GetNiddAuData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/nidd-authorization-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.singleNssai == nil {
		return localVarReturnValue, nil, reportError("singleNssai is required and must be specified")
	}
	if r.dnn == nil {
		return localVarReturnValue, nil, reportError("dnn is required and must be specified")
	}
	if r.mtcProviderInformation == nil {
		return localVarReturnValue, nil, reportError("mtcProviderInformation is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "single-nssai", r.singleNssai, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "dnn", r.dnn, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "mtc-provider-information", r.mtcProviderInformation, "")
	if r.afId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "af-id", r.afId, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
