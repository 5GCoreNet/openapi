/*
UPF Event Exposure Service

UPF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nupf_EventExposure

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SubscriptionsCollectionApiService SubscriptionsCollectionApi service
type SubscriptionsCollectionApiService service

type ApiCreateIndividualSubcriptionRequest struct {
	ctx context.Context
	ApiService *SubscriptionsCollectionApiService
	body *interface{}
}

func (r ApiCreateIndividualSubcriptionRequest) Body(body interface{}) ApiCreateIndividualSubcriptionRequest {
	r.body = &body
	return r
}

func (r ApiCreateIndividualSubcriptionRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateIndividualSubcriptionExecute(r)
}

/*
CreateIndividualSubcription subscribe to notifications

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateIndividualSubcriptionRequest
*/
func (a *SubscriptionsCollectionApiService) CreateIndividualSubcription(ctx context.Context) ApiCreateIndividualSubcriptionRequest {
	return ApiCreateIndividualSubcriptionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *SubscriptionsCollectionApiService) CreateIndividualSubcriptionExecute(r ApiCreateIndividualSubcriptionRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SubscriptionsCollectionApiService.CreateIndividualSubcription")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ee-subscriptions"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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
