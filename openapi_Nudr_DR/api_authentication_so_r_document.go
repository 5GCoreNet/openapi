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

// AuthenticationSoRDocumentApiService AuthenticationSoRDocumentApi service
type AuthenticationSoRDocumentApiService service

type ApiCreateAuthenticationSoRRequest struct {
	ctx               context.Context
	ApiService        *AuthenticationSoRDocumentApiService
	ueId              string
	sorData           *SorData
	supportedFeatures *string
}

func (r ApiCreateAuthenticationSoRRequest) SorData(sorData SorData) ApiCreateAuthenticationSoRRequest {
	r.sorData = &sorData
	return r
}

// Supported Features
func (r ApiCreateAuthenticationSoRRequest) SupportedFeatures(supportedFeatures string) ApiCreateAuthenticationSoRRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiCreateAuthenticationSoRRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateAuthenticationSoRExecute(r)
}

/*
CreateAuthenticationSoR To store the SoR acknowledgement information of a UE and ME support of SOR CMCI

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiCreateAuthenticationSoRRequest
*/
func (a *AuthenticationSoRDocumentApiService) CreateAuthenticationSoR(ctx context.Context, ueId string) ApiCreateAuthenticationSoRRequest {
	return ApiCreateAuthenticationSoRRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
func (a *AuthenticationSoRDocumentApiService) CreateAuthenticationSoRExecute(r ApiCreateAuthenticationSoRRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationSoRDocumentApiService.CreateAuthenticationSoR")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/ue-update-confirmation-data/sor-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sorData == nil {
		return nil, reportError("sorData is required and must be specified")
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
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sorData
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

type ApiQueryAuthSoRRequest struct {
	ctx               context.Context
	ApiService        *AuthenticationSoRDocumentApiService
	ueId              string
	supportedFeatures *string
}

// Supported Features
func (r ApiQueryAuthSoRRequest) SupportedFeatures(supportedFeatures string) ApiQueryAuthSoRRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiQueryAuthSoRRequest) Execute() (*SorData, *http.Response, error) {
	return r.ApiService.QueryAuthSoRExecute(r)
}

/*
QueryAuthSoR Retrieves the SoR acknowledgement information of a UE and ME support of SOR CMCI

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiQueryAuthSoRRequest
*/
func (a *AuthenticationSoRDocumentApiService) QueryAuthSoR(ctx context.Context, ueId string) ApiQueryAuthSoRRequest {
	return ApiQueryAuthSoRRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
//
//	@return SorData
func (a *AuthenticationSoRDocumentApiService) QueryAuthSoRExecute(r ApiQueryAuthSoRRequest) (*SorData, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SorData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationSoRDocumentApiService.QueryAuthSoR")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/ue-update-confirmation-data/sor-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

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

type ApiUpdateAuthenticationSoRRequest struct {
	ctx               context.Context
	ApiService        *AuthenticationSoRDocumentApiService
	ueId              string
	patchItem         *[]PatchItem
	supportedFeatures *string
}

func (r ApiUpdateAuthenticationSoRRequest) PatchItem(patchItem []PatchItem) ApiUpdateAuthenticationSoRRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiUpdateAuthenticationSoRRequest) SupportedFeatures(supportedFeatures string) ApiUpdateAuthenticationSoRRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiUpdateAuthenticationSoRRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.UpdateAuthenticationSoRExecute(r)
}

/*
UpdateAuthenticationSoR Updates the ME support of SOR CMCI information of a UE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId
	@return ApiUpdateAuthenticationSoRRequest
*/
func (a *AuthenticationSoRDocumentApiService) UpdateAuthenticationSoR(ctx context.Context, ueId string) ApiUpdateAuthenticationSoRRequest {
	return ApiUpdateAuthenticationSoRRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
//
//	@return PatchResult
func (a *AuthenticationSoRDocumentApiService) UpdateAuthenticationSoRExecute(r ApiUpdateAuthenticationSoRRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationSoRDocumentApiService.UpdateAuthenticationSoR")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/ue-update-confirmation-data/sor-data"
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
