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

// ProvisionedDataDocumentApiService ProvisionedDataDocumentApi service
type ProvisionedDataDocumentApiService service

type ApiQueryProvisionedDataRequest struct {
	ctx           context.Context
	ApiService    *ProvisionedDataDocumentApiService
	ueId          string
	servingPlmnId string
	datasetNames  *[]DataSetName
}

// List of dataset names
func (r ApiQueryProvisionedDataRequest) DatasetNames(datasetNames []DataSetName) ApiQueryProvisionedDataRequest {
	r.datasetNames = &datasetNames
	return r
}

func (r ApiQueryProvisionedDataRequest) Execute() (*ProvisionedDataSets, *http.Response, error) {
	return r.ApiService.QueryProvisionedDataExecute(r)
}

/*
QueryProvisionedData Retrieve multiple provisioned data sets of a UE

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@param servingPlmnId PLMN ID
	@return ApiQueryProvisionedDataRequest
*/
func (a *ProvisionedDataDocumentApiService) QueryProvisionedData(ctx context.Context, ueId string, servingPlmnId string) ApiQueryProvisionedDataRequest {
	return ApiQueryProvisionedDataRequest{
		ApiService:    a,
		ctx:           ctx,
		ueId:          ueId,
		servingPlmnId: servingPlmnId,
	}
}

// Execute executes the request
//
//	@return ProvisionedDataSets
func (a *ProvisionedDataDocumentApiService) QueryProvisionedDataExecute(r ApiQueryProvisionedDataRequest) (*ProvisionedDataSets, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *ProvisionedDataSets
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisionedDataDocumentApiService.QueryProvisionedData")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/{servingPlmnId}/provisioned-data"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"servingPlmnId"+"}", url.PathEscape(parameterValueToString(r.servingPlmnId, "servingPlmnId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.datasetNames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dataset-names", r.datasetNames, "csv")
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
