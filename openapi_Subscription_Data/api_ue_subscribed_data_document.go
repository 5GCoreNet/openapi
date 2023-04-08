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


// UESubscribedDataDocumentApiService UESubscribedDataDocumentApi service
type UESubscribedDataDocumentApiService service

type ApiQueryUeSubscribedDataRequest struct {
	ctx context.Context
	ApiService *UESubscribedDataDocumentApiService
	ueId string
	datasetNames *[]DataSetName
	servingPlmn *string
}

// List of dataset names
func (r ApiQueryUeSubscribedDataRequest) DatasetNames(datasetNames []DataSetName) ApiQueryUeSubscribedDataRequest {
	r.datasetNames = &datasetNames
	return r
}

// Serving PLMN Id
func (r ApiQueryUeSubscribedDataRequest) ServingPlmn(servingPlmn string) ApiQueryUeSubscribedDataRequest {
	r.servingPlmn = &servingPlmn
	return r
}

func (r ApiQueryUeSubscribedDataRequest) Execute() (*UeSubscribedDataSets, *http.Response, error) {
	return r.ApiService.QueryUeSubscribedDataExecute(r)
}

/*
QueryUeSubscribedData Retrieve multiple subscribed data sets of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE Id
 @return ApiQueryUeSubscribedDataRequest
*/
func (a *UESubscribedDataDocumentApiService) QueryUeSubscribedData(ctx context.Context, ueId string) ApiQueryUeSubscribedDataRequest {
	return ApiQueryUeSubscribedDataRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
//  @return UeSubscribedDataSets
func (a *UESubscribedDataDocumentApiService) QueryUeSubscribedDataExecute(r ApiQueryUeSubscribedDataRequest) (*UeSubscribedDataSets, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UeSubscribedDataSets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UESubscribedDataDocumentApiService.QueryUeSubscribedData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.datasetNames != nil {
		parameterAddToQuery(localVarQueryParams, "dataset-names", r.datasetNames, "csv")
	}
	if r.servingPlmn != nil {
		parameterAddToQuery(localVarQueryParams, "serving-plmn", r.servingPlmn, "")
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
