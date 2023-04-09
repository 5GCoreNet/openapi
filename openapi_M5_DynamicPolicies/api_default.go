/*
M5_DynamicPolicies

5GMS AF M5 Dynamic Policy API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_DynamicPolicies

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateDynamicPolicyRequest struct {
	ctx           context.Context
	ApiService    *DefaultApiService
	dynamicPolicy *DynamicPolicy
}

// An optional JSON representation of a Dynamic Policy resource
func (r ApiCreateDynamicPolicyRequest) DynamicPolicy(dynamicPolicy DynamicPolicy) ApiCreateDynamicPolicyRequest {
	r.dynamicPolicy = &dynamicPolicy
	return r
}

func (r ApiCreateDynamicPolicyRequest) Execute() (*DynamicPolicy, *http.Response, error) {
	return r.ApiService.CreateDynamicPolicyExecute(r)
}

/*
CreateDynamicPolicy Create (and optionally upload) a new Dynamic Policy resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateDynamicPolicyRequest
*/
func (a *DefaultApiService) CreateDynamicPolicy(ctx context.Context) ApiCreateDynamicPolicyRequest {
	return ApiCreateDynamicPolicyRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return DynamicPolicy
func (a *DefaultApiService) CreateDynamicPolicyExecute(r ApiCreateDynamicPolicyRequest) (*DynamicPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DynamicPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateDynamicPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dynamic-policies"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	localVarPostBody = r.dynamicPolicy
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

type ApiDestroyDynamicPolicyRequest struct {
	ctx             context.Context
	ApiService      *DefaultApiService
	dynamicPolicyId string
}

func (r ApiDestroyDynamicPolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.DestroyDynamicPolicyExecute(r)
}

/*
DestroyDynamicPolicy Destroy an existing Dynamic Policy resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param dynamicPolicyId The resource identifier of a Dynamic Policy resource
	@return ApiDestroyDynamicPolicyRequest
*/
func (a *DefaultApiService) DestroyDynamicPolicy(ctx context.Context, dynamicPolicyId string) ApiDestroyDynamicPolicyRequest {
	return ApiDestroyDynamicPolicyRequest{
		ApiService:      a,
		ctx:             ctx,
		dynamicPolicyId: dynamicPolicyId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DestroyDynamicPolicyExecute(r ApiDestroyDynamicPolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DestroyDynamicPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dynamic-policies/{dynamicPolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dynamicPolicyId"+"}", url.PathEscape(parameterValueToString(r.dynamicPolicyId, "dynamicPolicyId")), -1)

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

type ApiPatchDynamicPolicyRequest struct {
	ctx             context.Context
	ApiService      *DefaultApiService
	dynamicPolicyId string
	dynamicPolicy   *DynamicPolicy
}

// A JSON patch to a Dynamic Policy resource
func (r ApiPatchDynamicPolicyRequest) DynamicPolicy(dynamicPolicy DynamicPolicy) ApiPatchDynamicPolicyRequest {
	r.dynamicPolicy = &dynamicPolicy
	return r
}

func (r ApiPatchDynamicPolicyRequest) Execute() (*DynamicPolicy, *http.Response, error) {
	return r.ApiService.PatchDynamicPolicyExecute(r)
}

/*
PatchDynamicPolicy Patch an existing Dynamic Policy resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param dynamicPolicyId The resource identifier of a Dynamic Policy resource
	@return ApiPatchDynamicPolicyRequest
*/
func (a *DefaultApiService) PatchDynamicPolicy(ctx context.Context, dynamicPolicyId string) ApiPatchDynamicPolicyRequest {
	return ApiPatchDynamicPolicyRequest{
		ApiService:      a,
		ctx:             ctx,
		dynamicPolicyId: dynamicPolicyId,
	}
}

// Execute executes the request
//
//	@return DynamicPolicy
func (a *DefaultApiService) PatchDynamicPolicyExecute(r ApiPatchDynamicPolicyRequest) (*DynamicPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DynamicPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.PatchDynamicPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dynamic-policies/{dynamicPolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dynamicPolicyId"+"}", url.PathEscape(parameterValueToString(r.dynamicPolicyId, "dynamicPolicyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dynamicPolicy == nil {
		return localVarReturnValue, nil, reportError("dynamicPolicy is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/merge-patch+json", "application/json-patch+json"}

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
	localVarPostBody = r.dynamicPolicy
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

type ApiRetrieveDynamicPolicyRequest struct {
	ctx             context.Context
	ApiService      *DefaultApiService
	dynamicPolicyId string
}

func (r ApiRetrieveDynamicPolicyRequest) Execute() (*DynamicPolicy, *http.Response, error) {
	return r.ApiService.RetrieveDynamicPolicyExecute(r)
}

/*
RetrieveDynamicPolicy Retrieve an existing Dynamic Policy resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param dynamicPolicyId The resource identifier of a Dynamic Policy resource
	@return ApiRetrieveDynamicPolicyRequest
*/
func (a *DefaultApiService) RetrieveDynamicPolicy(ctx context.Context, dynamicPolicyId string) ApiRetrieveDynamicPolicyRequest {
	return ApiRetrieveDynamicPolicyRequest{
		ApiService:      a,
		ctx:             ctx,
		dynamicPolicyId: dynamicPolicyId,
	}
}

// Execute executes the request
//
//	@return DynamicPolicy
func (a *DefaultApiService) RetrieveDynamicPolicyExecute(r ApiRetrieveDynamicPolicyRequest) (*DynamicPolicy, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *DynamicPolicy
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.RetrieveDynamicPolicy")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dynamic-policies/{dynamicPolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dynamicPolicyId"+"}", url.PathEscape(parameterValueToString(r.dynamicPolicyId, "dynamicPolicyId")), -1)

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

type ApiUpdateDynamicPolicyRequest struct {
	ctx             context.Context
	ApiService      *DefaultApiService
	dynamicPolicyId string
	dynamicPolicy   *DynamicPolicy
}

// A replacement JSON representation of a Dynamic Policy resource
func (r ApiUpdateDynamicPolicyRequest) DynamicPolicy(dynamicPolicy DynamicPolicy) ApiUpdateDynamicPolicyRequest {
	r.dynamicPolicy = &dynamicPolicy
	return r
}

func (r ApiUpdateDynamicPolicyRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdateDynamicPolicyExecute(r)
}

/*
UpdateDynamicPolicy Update an existing Dynamic Policy resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param dynamicPolicyId The resource identifier of a Dynamic Policy resource
	@return ApiUpdateDynamicPolicyRequest
*/
func (a *DefaultApiService) UpdateDynamicPolicy(ctx context.Context, dynamicPolicyId string) ApiUpdateDynamicPolicyRequest {
	return ApiUpdateDynamicPolicyRequest{
		ApiService:      a,
		ctx:             ctx,
		dynamicPolicyId: dynamicPolicyId,
	}
}

// Execute executes the request
func (a *DefaultApiService) UpdateDynamicPolicyExecute(r ApiUpdateDynamicPolicyRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UpdateDynamicPolicy")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/dynamic-policies/{dynamicPolicyId}"
	localVarPath = strings.Replace(localVarPath, "{"+"dynamicPolicyId"+"}", url.PathEscape(parameterValueToString(r.dynamicPolicyId, "dynamicPolicyId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.dynamicPolicy == nil {
		return nil, reportError("dynamicPolicy is required and must be specified")
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
	localVarPostBody = r.dynamicPolicy
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
