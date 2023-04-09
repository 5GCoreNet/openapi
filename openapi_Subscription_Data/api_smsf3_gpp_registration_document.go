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

// SMSF3GPPRegistrationDocumentApiService SMSF3GPPRegistrationDocumentApi service
type SMSF3GPPRegistrationDocumentApiService service

type ApiCreateSmsfContext3gppRequest struct {
	ctx              context.Context
	ApiService       *SMSF3GPPRegistrationDocumentApiService
	ueId             string
	smsfRegistration *SmsfRegistration
}

func (r ApiCreateSmsfContext3gppRequest) SmsfRegistration(smsfRegistration SmsfRegistration) ApiCreateSmsfContext3gppRequest {
	r.smsfRegistration = &smsfRegistration
	return r
}

func (r ApiCreateSmsfContext3gppRequest) Execute() (*SmsfRegistration, *http.Response, error) {
	return r.ApiService.CreateSmsfContext3gppExecute(r)
}

/*
CreateSmsfContext3gpp Create the SMSF context data of a UE via 3GPP access

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiCreateSmsfContext3gppRequest
*/
func (a *SMSF3GPPRegistrationDocumentApiService) CreateSmsfContext3gpp(ctx context.Context, ueId string) ApiCreateSmsfContext3gppRequest {
	return ApiCreateSmsfContext3gppRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
//
//	@return SmsfRegistration
func (a *SMSF3GPPRegistrationDocumentApiService) CreateSmsfContext3gppExecute(r ApiCreateSmsfContext3gppRequest) (*SmsfRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SmsfRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSF3GPPRegistrationDocumentApiService.CreateSmsfContext3gpp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smsf-3gpp-access"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.smsfRegistration == nil {
		return localVarReturnValue, nil, reportError("smsfRegistration is required and must be specified")
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
	localVarPostBody = r.smsfRegistration
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

type ApiDeleteSmsfContext3gppRequest struct {
	ctx        context.Context
	ApiService *SMSF3GPPRegistrationDocumentApiService
	ueId       string
}

func (r ApiDeleteSmsfContext3gppRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSmsfContext3gppExecute(r)
}

/*
DeleteSmsfContext3gpp To remove the SMSF context data of a UE via 3GPP access

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiDeleteSmsfContext3gppRequest
*/
func (a *SMSF3GPPRegistrationDocumentApiService) DeleteSmsfContext3gpp(ctx context.Context, ueId string) ApiDeleteSmsfContext3gppRequest {
	return ApiDeleteSmsfContext3gppRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
func (a *SMSF3GPPRegistrationDocumentApiService) DeleteSmsfContext3gppExecute(r ApiDeleteSmsfContext3gppRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSF3GPPRegistrationDocumentApiService.DeleteSmsfContext3gpp")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smsf-3gpp-access"
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

type ApiQuerySmsfContext3gppRequest struct {
	ctx               context.Context
	ApiService        *SMSF3GPPRegistrationDocumentApiService
	ueId              string
	fields            *[]string
	supportedFeatures *string
}

// attributes to be retrieved
func (r ApiQuerySmsfContext3gppRequest) Fields(fields []string) ApiQuerySmsfContext3gppRequest {
	r.fields = &fields
	return r
}

// Supported Features
func (r ApiQuerySmsfContext3gppRequest) SupportedFeatures(supportedFeatures string) ApiQuerySmsfContext3gppRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiQuerySmsfContext3gppRequest) Execute() (*SmsfRegistration, *http.Response, error) {
	return r.ApiService.QuerySmsfContext3gppExecute(r)
}

/*
QuerySmsfContext3gpp Retrieves the SMSF context data of a UE using 3gpp access

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiQuerySmsfContext3gppRequest
*/
func (a *SMSF3GPPRegistrationDocumentApiService) QuerySmsfContext3gpp(ctx context.Context, ueId string) ApiQuerySmsfContext3gppRequest {
	return ApiQuerySmsfContext3gppRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
//
//	@return SmsfRegistration
func (a *SMSF3GPPRegistrationDocumentApiService) QuerySmsfContext3gppExecute(r ApiQuerySmsfContext3gppRequest) (*SmsfRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *SmsfRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSF3GPPRegistrationDocumentApiService.QuerySmsfContext3gpp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smsf-3gpp-access"
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
