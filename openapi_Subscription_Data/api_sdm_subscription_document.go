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

// SDMSubscriptionDocumentApiService SDMSubscriptionDocumentApi service
type SDMSubscriptionDocumentApiService service

type ApiModifysdmSubscriptionRequest struct {
	ctx               context.Context
	ApiService        *SDMSubscriptionDocumentApiService
	ueId              string
	subsId            string
	patchItem         *[]PatchItem
	supportedFeatures *string
}

func (r ApiModifysdmSubscriptionRequest) PatchItem(patchItem []PatchItem) ApiModifysdmSubscriptionRequest {
	r.patchItem = &patchItem
	return r
}

// Features required to be supported by the target NF
func (r ApiModifysdmSubscriptionRequest) SupportedFeatures(supportedFeatures string) ApiModifysdmSubscriptionRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiModifysdmSubscriptionRequest) Execute() (*PatchResult, *http.Response, error) {
	return r.ApiService.ModifysdmSubscriptionExecute(r)
}

/*
ModifysdmSubscription Modify an individual sdm subscription

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@param subsId
	@return ApiModifysdmSubscriptionRequest
*/
func (a *SDMSubscriptionDocumentApiService) ModifysdmSubscription(ctx context.Context, ueId string, subsId string) ApiModifysdmSubscriptionRequest {
	return ApiModifysdmSubscriptionRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
		subsId:     subsId,
	}
}

// Execute executes the request
//
//	@return PatchResult
func (a *SDMSubscriptionDocumentApiService) ModifysdmSubscriptionExecute(r ApiModifysdmSubscriptionRequest) (*PatchResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *PatchResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SDMSubscriptionDocumentApiService.ModifysdmSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.patchItem == nil {
		return localVarReturnValue, nil, reportError("patchItem is required and must be specified")
	}
	if len(*r.patchItem) < 1 {
		return localVarReturnValue, nil, reportError("patchItem must have at least 1 elements")
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

type ApiQuerysdmSubscriptionRequest struct {
	ctx        context.Context
	ApiService *SDMSubscriptionDocumentApiService
	ueId       string
	subsId     string
}

func (r ApiQuerysdmSubscriptionRequest) Execute() ([]SdmSubscription, *http.Response, error) {
	return r.ApiService.QuerysdmSubscriptionExecute(r)
}

/*
QuerysdmSubscription Retrieves a individual sdmSubscription identified by subsId

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId
	@param subsId Unique ID of the subscription to retrieve
	@return ApiQuerysdmSubscriptionRequest
*/
func (a *SDMSubscriptionDocumentApiService) QuerysdmSubscription(ctx context.Context, ueId string, subsId string) ApiQuerysdmSubscriptionRequest {
	return ApiQuerysdmSubscriptionRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
		subsId:     subsId,
	}
}

// Execute executes the request
//
//	@return []SdmSubscription
func (a *SDMSubscriptionDocumentApiService) QuerysdmSubscriptionExecute(r ApiQuerysdmSubscriptionRequest) ([]SdmSubscription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []SdmSubscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SDMSubscriptionDocumentApiService.QuerysdmSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

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

type ApiRemovesdmSubscriptionsRequest struct {
	ctx        context.Context
	ApiService *SDMSubscriptionDocumentApiService
	ueId       string
	subsId     string
}

func (r ApiRemovesdmSubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemovesdmSubscriptionsExecute(r)
}

/*
RemovesdmSubscriptions Deletes a sdmsubscriptions

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId
	@param subsId Unique ID of the subscription to remove
	@return ApiRemovesdmSubscriptionsRequest
*/
func (a *SDMSubscriptionDocumentApiService) RemovesdmSubscriptions(ctx context.Context, ueId string, subsId string) ApiRemovesdmSubscriptionsRequest {
	return ApiRemovesdmSubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
		subsId:     subsId,
	}
}

// Execute executes the request
func (a *SDMSubscriptionDocumentApiService) RemovesdmSubscriptionsExecute(r ApiRemovesdmSubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SDMSubscriptionDocumentApiService.RemovesdmSubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

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
		if localVarHTTPResponse.StatusCode == 404 {
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

type ApiUpdatesdmsubscriptionsRequest struct {
	ctx             context.Context
	ApiService      *SDMSubscriptionDocumentApiService
	ueId            string
	subsId          string
	sdmSubscription *SdmSubscription
}

func (r ApiUpdatesdmsubscriptionsRequest) SdmSubscription(sdmSubscription SdmSubscription) ApiUpdatesdmsubscriptionsRequest {
	r.sdmSubscription = &sdmSubscription
	return r
}

func (r ApiUpdatesdmsubscriptionsRequest) Execute() (*http.Response, error) {
	return r.ApiService.UpdatesdmsubscriptionsExecute(r)
}

/*
Updatesdmsubscriptions Update an individual sdm subscriptions of a UE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId
	@param subsId
	@return ApiUpdatesdmsubscriptionsRequest
*/
func (a *SDMSubscriptionDocumentApiService) Updatesdmsubscriptions(ctx context.Context, ueId string, subsId string) ApiUpdatesdmsubscriptionsRequest {
	return ApiUpdatesdmsubscriptionsRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
		subsId:     subsId,
	}
}

// Execute executes the request
func (a *SDMSubscriptionDocumentApiService) UpdatesdmsubscriptionsExecute(r ApiUpdatesdmsubscriptionsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SDMSubscriptionDocumentApiService.Updatesdmsubscriptions")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/sdm-subscriptions/{subsId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"subsId"+"}", url.PathEscape(parameterValueToString(r.subsId, "subsId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.sdmSubscription == nil {
		return nil, reportError("sdmSubscription is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.sdmSubscription
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
		if localVarHTTPResponse.StatusCode == 404 {
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
