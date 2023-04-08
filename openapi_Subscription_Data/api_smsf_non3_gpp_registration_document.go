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


// SMSFNon3GPPRegistrationDocumentApiService SMSFNon3GPPRegistrationDocumentApi service
type SMSFNon3GPPRegistrationDocumentApiService service

type ApiCreateSmsfContextNon3gppRequest struct {
	ctx context.Context
	ApiService *SMSFNon3GPPRegistrationDocumentApiService
	ueId string
	smsfRegistration *SmsfRegistration
}

func (r ApiCreateSmsfContextNon3gppRequest) SmsfRegistration(smsfRegistration SmsfRegistration) ApiCreateSmsfContextNon3gppRequest {
	r.smsfRegistration = &smsfRegistration
	return r
}

func (r ApiCreateSmsfContextNon3gppRequest) Execute() (*SmsfRegistration, *http.Response, error) {
	return r.ApiService.CreateSmsfContextNon3gppExecute(r)
}

/*
CreateSmsfContextNon3gpp Create the SMSF context data of a UE via non-3GPP access

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiCreateSmsfContextNon3gppRequest
*/
func (a *SMSFNon3GPPRegistrationDocumentApiService) CreateSmsfContextNon3gpp(ctx context.Context, ueId string) ApiCreateSmsfContextNon3gppRequest {
	return ApiCreateSmsfContextNon3gppRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return SmsfRegistration
func (a *SMSFNon3GPPRegistrationDocumentApiService) CreateSmsfContextNon3gppExecute(r ApiCreateSmsfContextNon3gppRequest) (*SmsfRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmsfRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSFNon3GPPRegistrationDocumentApiService.CreateSmsfContextNon3gpp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smsf-non-3gpp-access"
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

type ApiDeleteSmsfContextNon3gppRequest struct {
	ctx context.Context
	ApiService *SMSFNon3GPPRegistrationDocumentApiService
	ueId string
}

func (r ApiDeleteSmsfContextNon3gppRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteSmsfContextNon3gppExecute(r)
}

/*
DeleteSmsfContextNon3gpp To remove the SMSF context data of a UE via non-3GPP access

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiDeleteSmsfContextNon3gppRequest
*/
func (a *SMSFNon3GPPRegistrationDocumentApiService) DeleteSmsfContextNon3gpp(ctx context.Context, ueId string) ApiDeleteSmsfContextNon3gppRequest {
	return ApiDeleteSmsfContextNon3gppRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
func (a *SMSFNon3GPPRegistrationDocumentApiService) DeleteSmsfContextNon3gppExecute(r ApiDeleteSmsfContextNon3gppRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSFNon3GPPRegistrationDocumentApiService.DeleteSmsfContextNon3gpp")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smsf-non-3gpp-access"
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

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
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

type ApiQuerySmsfContextNon3gppRequest struct {
	ctx context.Context
	ApiService *SMSFNon3GPPRegistrationDocumentApiService
	ueId string
	fields *[]string
	supportedFeatures *string
}

// attributes to be retrieved
func (r ApiQuerySmsfContextNon3gppRequest) Fields(fields []string) ApiQuerySmsfContextNon3gppRequest {
	r.fields = &fields
	return r
}

// Supported Features
func (r ApiQuerySmsfContextNon3gppRequest) SupportedFeatures(supportedFeatures string) ApiQuerySmsfContextNon3gppRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiQuerySmsfContextNon3gppRequest) Execute() (*SmsfRegistration, *http.Response, error) {
	return r.ApiService.QuerySmsfContextNon3gppExecute(r)
}

/*
QuerySmsfContextNon3gpp Retrieves the SMSF context data of a UE using non-3gpp access

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiQuerySmsfContextNon3gppRequest
*/
func (a *SMSFNon3GPPRegistrationDocumentApiService) QuerySmsfContextNon3gpp(ctx context.Context, ueId string) ApiQuerySmsfContextNon3gppRequest {
	return ApiQuerySmsfContextNon3gppRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return SmsfRegistration
func (a *SMSFNon3GPPRegistrationDocumentApiService) QuerySmsfContextNon3gppExecute(r ApiQuerySmsfContextNon3gppRequest) (*SmsfRegistration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SmsfRegistration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SMSFNon3GPPRegistrationDocumentApiService.QuerySmsfContextNon3gpp")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/smsf-non-3gpp-access"
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
