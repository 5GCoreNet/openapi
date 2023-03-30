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


// IndividualAuthEventDocumentApiService IndividualAuthEventDocumentApi service
type IndividualAuthEventDocumentApiService service

type ApiDeleteIndividualAuthenticationStatusRequest struct {
	ctx context.Context
	ApiService *IndividualAuthEventDocumentApiService
	ueId string
	servingNetworkName string
}

func (r ApiDeleteIndividualAuthenticationStatusRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteIndividualAuthenticationStatusExecute(r)
}

/*
DeleteIndividualAuthenticationStatus To remove the Individual Authentication Status of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param servingNetworkName Serving Network Name
 @return ApiDeleteIndividualAuthenticationStatusRequest
*/
func (a *IndividualAuthEventDocumentApiService) DeleteIndividualAuthenticationStatus(ctx context.Context, ueId string, servingNetworkName string) ApiDeleteIndividualAuthenticationStatusRequest {
	return ApiDeleteIndividualAuthenticationStatusRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		servingNetworkName: servingNetworkName,
	}
}

// Execute executes the request
func (a *IndividualAuthEventDocumentApiService) DeleteIndividualAuthenticationStatusExecute(r ApiDeleteIndividualAuthenticationStatusRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndividualAuthEventDocumentApiService.DeleteIndividualAuthenticationStatus")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/authentication-data/authentication-status/{servingNetworkName}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"servingNetworkName"+"}", url.PathEscape(parameterValueToString(r.servingNetworkName, "servingNetworkName")), -1)

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

type ApiQueryIndividualAuthenticationStatusRequest struct {
	ctx context.Context
	ApiService *IndividualAuthEventDocumentApiService
	ueId string
	servingNetworkName string
	fields *[]string
	supportedFeatures *string
}

// attributes to be retrieved
func (r ApiQueryIndividualAuthenticationStatusRequest) Fields(fields []string) ApiQueryIndividualAuthenticationStatusRequest {
	r.fields = &fields
	return r
}

// Supported Features
func (r ApiQueryIndividualAuthenticationStatusRequest) SupportedFeatures(supportedFeatures string) ApiQueryIndividualAuthenticationStatusRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiQueryIndividualAuthenticationStatusRequest) Execute() (*AuthEvent, *http.Response, error) {
	return r.ApiService.QueryIndividualAuthenticationStatusExecute(r)
}

/*
QueryIndividualAuthenticationStatus Retrieves the Individual Authentication Status of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @param servingNetworkName Serving Network Name
 @return ApiQueryIndividualAuthenticationStatusRequest
*/
func (a *IndividualAuthEventDocumentApiService) QueryIndividualAuthenticationStatus(ctx context.Context, ueId string, servingNetworkName string) ApiQueryIndividualAuthenticationStatusRequest {
	return ApiQueryIndividualAuthenticationStatusRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
		servingNetworkName: servingNetworkName,
	}
}

// Execute executes the request
//  @return AuthEvent
func (a *IndividualAuthEventDocumentApiService) QueryIndividualAuthenticationStatusExecute(r ApiQueryIndividualAuthenticationStatusRequest) (*AuthEvent, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AuthEvent
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IndividualAuthEventDocumentApiService.QueryIndividualAuthenticationStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/authentication-data/authentication-status/{servingNetworkName}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"servingNetworkName"+"}", url.PathEscape(parameterValueToString(r.servingNetworkName, "servingNetworkName")), -1)

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
