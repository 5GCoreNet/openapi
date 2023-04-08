/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)


// NWDAFAnalyticsDocumentApiService NWDAFAnalyticsDocumentApi service
type NWDAFAnalyticsDocumentApiService service

type ApiGetNWDAFAnalyticsRequest struct {
	ctx context.Context
	ApiService *NWDAFAnalyticsDocumentApiService
	eventId *EventId
	anaReq *EventReportingRequirement
	eventFilter *EventFilter
	supportedFeatures *string
	tgtUe *TargetUeInformation
}

// Identify the analytics.
func (r ApiGetNWDAFAnalyticsRequest) EventId(eventId EventId) ApiGetNWDAFAnalyticsRequest {
	r.eventId = &eventId
	return r
}

// Identifies the analytics reporting requirement information.
func (r ApiGetNWDAFAnalyticsRequest) AnaReq(anaReq EventReportingRequirement) ApiGetNWDAFAnalyticsRequest {
	r.anaReq = &anaReq
	return r
}

// Identify the analytics.
func (r ApiGetNWDAFAnalyticsRequest) EventFilter(eventFilter EventFilter) ApiGetNWDAFAnalyticsRequest {
	r.eventFilter = &eventFilter
	return r
}

// To filter irrelevant responses related to unsupported features.
func (r ApiGetNWDAFAnalyticsRequest) SupportedFeatures(supportedFeatures string) ApiGetNWDAFAnalyticsRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// Identify the target UE information.
func (r ApiGetNWDAFAnalyticsRequest) TgtUe(tgtUe TargetUeInformation) ApiGetNWDAFAnalyticsRequest {
	r.tgtUe = &tgtUe
	return r
}

func (r ApiGetNWDAFAnalyticsRequest) Execute() (*AnalyticsData, *http.Response, error) {
	return r.ApiService.GetNWDAFAnalyticsExecute(r)
}

/*
GetNWDAFAnalytics Read a NWDAF Analytics

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetNWDAFAnalyticsRequest
*/
func (a *NWDAFAnalyticsDocumentApiService) GetNWDAFAnalytics(ctx context.Context) ApiGetNWDAFAnalyticsRequest {
	return ApiGetNWDAFAnalyticsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return AnalyticsData
func (a *NWDAFAnalyticsDocumentApiService) GetNWDAFAnalyticsExecute(r ApiGetNWDAFAnalyticsRequest) (*AnalyticsData, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AnalyticsData
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NWDAFAnalyticsDocumentApiService.GetNWDAFAnalytics")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/analytics"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.eventId == nil {
		return localVarReturnValue, nil, reportError("eventId is required and must be specified")
	}

	parameterAddToQuery(localVarQueryParams, "event-id", r.eventId, "")
	if r.anaReq != nil {
		parameterAddToQuery(localVarQueryParams, "ana-req", r.anaReq, "")
	}
	if r.eventFilter != nil {
		parameterAddToQuery(localVarQueryParams, "event-filter", r.eventFilter, "")
	}
	if r.supportedFeatures != nil {
		parameterAddToQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	if r.tgtUe != nil {
		parameterAddToQuery(localVarQueryParams, "tgt-ue", r.tgtUe, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

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
		if localVarHTTPResponse.StatusCode == 400 {
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
		if localVarHTTPResponse.StatusCode == 401 {
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
		if localVarHTTPResponse.StatusCode == 414 {
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
		if localVarHTTPResponse.StatusCode == 429 {
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
		if localVarHTTPResponse.StatusCode == 500 {
			var v ProblemDetailsAnalyticsInfoRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 502 {
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
		if localVarHTTPResponse.StatusCode == 503 {
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
