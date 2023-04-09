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

// Delete5GVnGroupApiService Delete5GVnGroupApi service
type Delete5GVnGroupApiService service

type ApiDelete5GVnGroupRequest struct {
	ctx             context.Context
	ApiService      *Delete5GVnGroupApiService
	externalGroupId string
}

func (r ApiDelete5GVnGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.Delete5GVnGroupExecute(r)
}

/*
Delete5GVnGroup Deletes the 5GVnGroup

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param externalGroupId
	@return ApiDelete5GVnGroupRequest
*/
func (a *Delete5GVnGroupApiService) Delete5GVnGroup(ctx context.Context, externalGroupId string) ApiDelete5GVnGroupRequest {
	return ApiDelete5GVnGroupRequest{
		ApiService:      a,
		ctx:             ctx,
		externalGroupId: externalGroupId,
	}
}

// Execute executes the request
func (a *Delete5GVnGroupApiService) Delete5GVnGroupExecute(r ApiDelete5GVnGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Delete5GVnGroupApiService.Delete5GVnGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/5g-vn-groups/{externalGroupId}"
	localVarPath = strings.Replace(localVarPath, "{"+"externalGroupId"+"}", url.PathEscape(parameterValueToString(r.externalGroupId, "externalGroupId")), -1)

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
