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


// OperatorSpecificDataContainerDocumentApiService OperatorSpecificDataContainerDocumentApi service
type OperatorSpecificDataContainerDocumentApiService service

type ApiCreateOperSpecDataRequest struct {
	ctx context.Context
	ApiService *OperatorSpecificDataContainerDocumentApiService
	ueId string
	requestBody *map[string]OperatorSpecificDataContainer
	supportedFeatures *string
}

func (r ApiCreateOperSpecDataRequest) RequestBody(requestBody map[string]OperatorSpecificDataContainer) ApiCreateOperSpecDataRequest {
	r.requestBody = &requestBody
	return r
}

// Supported Features
func (r ApiCreateOperSpecDataRequest) SupportedFeatures(supportedFeatures string) ApiCreateOperSpecDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiCreateOperSpecDataRequest) Execute() (*map[string]OperatorSpecificDataContainer, *http.Response, error) {
	return r.ApiService.CreateOperSpecDataExecute(r)
}

/*
CreateOperSpecData To create an operator-specific data resource of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiCreateOperSpecDataRequest
*/
func (a *OperatorSpecificDataContainerDocumentApiService) CreateOperSpecData(ctx context.Context, ueId string) ApiCreateOperSpecDataRequest {
	return ApiCreateOperSpecDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return map[string]OperatorSpecificDataContainer
func (a *OperatorSpecificDataContainerDocumentApiService) CreateOperSpecDataExecute(r ApiCreateOperSpecDataRequest) (*map[string]OperatorSpecificDataContainer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string]OperatorSpecificDataContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorSpecificDataContainerDocumentApiService.CreateOperSpecData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/operator-specific-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.requestBody == nil {
		return localVarReturnValue, nil, reportError("requestBody is required and must be specified")
	}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.requestBody
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

type ApiDeleteOperSpecDataRequest struct {
	ctx context.Context
	ApiService *OperatorSpecificDataContainerDocumentApiService
	ueId string
}

func (r ApiDeleteOperSpecDataRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteOperSpecDataExecute(r)
}

/*
DeleteOperSpecData To remove an operator-specific data resource of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiDeleteOperSpecDataRequest
*/
func (a *OperatorSpecificDataContainerDocumentApiService) DeleteOperSpecData(ctx context.Context, ueId string) ApiDeleteOperSpecDataRequest {
	return ApiDeleteOperSpecDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
func (a *OperatorSpecificDataContainerDocumentApiService) DeleteOperSpecDataExecute(r ApiDeleteOperSpecDataRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorSpecificDataContainerDocumentApiService.DeleteOperSpecData")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/operator-specific-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/problem+json"}

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
		if localVarHTTPResponse.StatusCode == 403 {
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

type ApiModifyOperSpecDataRequest struct {
	ctx context.Context
	ApiService *OperatorSpecificDataContainerDocumentApiService
	ueId string
	patchItem *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifyOperSpecDataRequest) PatchItem(patchItem []PatchItem) ApiModifyOperSpecDataRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifyOperSpecDataRequest) SupportedFeatures(supportedFeatures string) ApiModifyOperSpecDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifyOperSpecDataRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.ModifyOperSpecDataExecute(r)
}

/*
ModifyOperSpecData To modify operator specific data of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiModifyOperSpecDataRequest
*/
func (a *OperatorSpecificDataContainerDocumentApiService) ModifyOperSpecData(ctx context.Context, ueId string) ApiModifyOperSpecDataRequest {
	return ApiModifyOperSpecDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return PatchResult
func (a *OperatorSpecificDataContainerDocumentApiService) ModifyOperSpecDataExecute(r ApiModifyOperSpecDataRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorSpecificDataContainerDocumentApiService.ModifyOperSpecData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/operator-specific-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

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

type ApiQueryOperSpecDataRequest struct {
	ctx context.Context
	ApiService *OperatorSpecificDataContainerDocumentApiService
	ueId string
	fields *[]string
	supportedFeatures *string
	ifNoneMatch *string
	ifModifiedSince *string
}

// attributes to be retrieved
func (r ApiQueryOperSpecDataRequest) Fields(fields []string) ApiQueryOperSpecDataRequest {
	r.fields = &fields
	return r
}

// Supported Features
func (r ApiQueryOperSpecDataRequest) SupportedFeatures(supportedFeatures string) ApiQueryOperSpecDataRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiQueryOperSpecDataRequest) IfNoneMatch(ifNoneMatch string) ApiQueryOperSpecDataRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiQueryOperSpecDataRequest) IfModifiedSince(ifModifiedSince string) ApiQueryOperSpecDataRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiQueryOperSpecDataRequest) Execute() (*map[string]OperatorSpecificDataContainer, *http.Response, error) {
	return r.ApiService.QueryOperSpecDataExecute(r)
}

/*
QueryOperSpecData Retrieves the operator specific data of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiQueryOperSpecDataRequest
*/
func (a *OperatorSpecificDataContainerDocumentApiService) QueryOperSpecData(ctx context.Context, ueId string) ApiQueryOperSpecDataRequest {
	return ApiQueryOperSpecDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return map[string]OperatorSpecificDataContainer
func (a *OperatorSpecificDataContainerDocumentApiService) QueryOperSpecDataExecute(r ApiQueryOperSpecDataRequest) (*map[string]OperatorSpecificDataContainer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string]OperatorSpecificDataContainer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OperatorSpecificDataContainerDocumentApiService.QueryOperSpecData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/operator-specific-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fields != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fields", r.fields, "csv")
	}
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
