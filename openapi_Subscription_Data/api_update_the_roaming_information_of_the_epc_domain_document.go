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

// UpdateTheRoamingInformationOfTheEPCDomainDocumentApiService UpdateTheRoamingInformationOfTheEPCDomainDocumentApi service
type UpdateTheRoamingInformationOfTheEPCDomainDocumentApiService service

type ApiUpdateRoamingInformationRequest struct {
	ctx               context.Context
	ApiService        *UpdateTheRoamingInformationOfTheEPCDomainDocumentApiService
	ueId              string
	roamingInfoUpdate *RoamingInfoUpdate
}

func (r ApiUpdateRoamingInformationRequest) RoamingInfoUpdate(roamingInfoUpdate RoamingInfoUpdate) ApiUpdateRoamingInformationRequest {
	r.roamingInfoUpdate = &roamingInfoUpdate
	return r
}

func (r ApiUpdateRoamingInformationRequest) Execute() (*RoamingInfoUpdate, *http.Response, error) {
	return r.ApiService.UpdateRoamingInformationExecute(r)
}

/*
UpdateRoamingInformation Update the Roaming Information of the EPC domain

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId UE id
	@return ApiUpdateRoamingInformationRequest
*/
func (a *UpdateTheRoamingInformationOfTheEPCDomainDocumentApiService) UpdateRoamingInformation(ctx context.Context, ueId string) ApiUpdateRoamingInformationRequest {
	return ApiUpdateRoamingInformationRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
//
//	@return RoamingInfoUpdate
func (a *UpdateTheRoamingInformationOfTheEPCDomainDocumentApiService) UpdateRoamingInformationExecute(r ApiUpdateRoamingInformationRequest) (*RoamingInfoUpdate, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *RoamingInfoUpdate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UpdateTheRoamingInformationOfTheEPCDomainDocumentApiService.UpdateRoamingInformation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/context-data/roaming-information"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.roamingInfoUpdate == nil {
		return localVarReturnValue, nil, reportError("roamingInfoUpdate is required and must be specified")
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
	localVarPostBody = r.roamingInfoUpdate
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
