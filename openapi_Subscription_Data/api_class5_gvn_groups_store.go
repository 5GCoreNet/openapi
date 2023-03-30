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
)


// Class5GVNGroupsStoreApiService Class5GVNGroupsStoreApi service
type Class5GVNGroupsStoreApiService service

type ApiQuery5GVnGroupRequest struct {
	ctx context.Context
	ApiService *Class5GVNGroupsStoreApiService
	gpsis *[]string
}

// List of GPSIs
func (r ApiQuery5GVnGroupRequest) Gpsis(gpsis []string) ApiQuery5GVnGroupRequest {
	r.gpsis = &gpsis
	return r
}

func (r ApiQuery5GVnGroupRequest) Execute() (*map[string]Model5GVnGroupConfiguration, *http.Response, error) {
	return r.ApiService.Query5GVnGroupExecute(r)
}

/*
Query5GVnGroup Retrieves the data of a 5G VN Group

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQuery5GVnGroupRequest
*/
func (a *Class5GVNGroupsStoreApiService) Query5GVnGroup(ctx context.Context) ApiQuery5GVnGroupRequest {
	return ApiQuery5GVnGroupRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]Model5GVnGroupConfiguration
func (a *Class5GVNGroupsStoreApiService) Query5GVnGroupExecute(r ApiQuery5GVnGroupRequest) (*map[string]Model5GVnGroupConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string]Model5GVnGroupConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "Class5GVNGroupsStoreApiService.Query5GVnGroup")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/group-data/5g-vn-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.gpsis != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "gpsis", r.gpsis, "csv")
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
