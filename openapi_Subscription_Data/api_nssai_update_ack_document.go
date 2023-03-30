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


// NSSAIUpdateAckDocumentApiService NSSAIUpdateAckDocumentApi service
type NSSAIUpdateAckDocumentApiService service

type ApiCreateOrUpdateNssaiAckRequest struct {
	ctx context.Context
	ApiService *NSSAIUpdateAckDocumentApiService
	ueId string
	nssaiAckData *NssaiAckData
	supportedFeatures *string
}

func (r ApiCreateOrUpdateNssaiAckRequest) NssaiAckData(nssaiAckData NssaiAckData) ApiCreateOrUpdateNssaiAckRequest {
	r.nssaiAckData = &nssaiAckData
	return r
}

// Supported Features
func (r ApiCreateOrUpdateNssaiAckRequest) SupportedFeatures(supportedFeatures string) ApiCreateOrUpdateNssaiAckRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

func (r ApiCreateOrUpdateNssaiAckRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateOrUpdateNssaiAckExecute(r)
}

/*
CreateOrUpdateNssaiAck To store the NSSAI update acknowledgement information of a UE

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param ueId UE id
 @return ApiCreateOrUpdateNssaiAckRequest
*/
func (a *NSSAIUpdateAckDocumentApiService) CreateOrUpdateNssaiAck(ctx context.Context, ueId string) ApiCreateOrUpdateNssaiAckRequest {
	return ApiCreateOrUpdateNssaiAckRequest{
		ApiService: a,
		ctx: ctx,
		ueId: ueId,
	}
}

// Execute executes the request
func (a *NSSAIUpdateAckDocumentApiService) CreateOrUpdateNssaiAckExecute(r ApiCreateOrUpdateNssaiAckRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NSSAIUpdateAckDocumentApiService.CreateOrUpdateNssaiAck")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/subscription-data/{ueId}/ue-update-confirmation-data/subscribed-snssais"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nssaiAckData == nil {
		return nil, reportError("nssaiAckData is required and must be specified")
	}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
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
	localVarPostBody = r.nssaiAckData
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
