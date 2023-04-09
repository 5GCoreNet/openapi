/*
TS 28.532 Performance Threshold Monitoring MnS

OAS 3.0.1 definition of the Performance Threshold Monitoring MnS © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMnS

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiNotificationSinkPostRequest struct {
	ctx                     context.Context
	ApiService              *DefaultApiService
	notifyThresholdCrossing *NotifyThresholdCrossing
}

func (r ApiNotificationSinkPostRequest) NotifyThresholdCrossing(notifyThresholdCrossing NotifyThresholdCrossing) ApiNotificationSinkPostRequest {
	r.notifyThresholdCrossing = &notifyThresholdCrossing
	return r
}

func (r ApiNotificationSinkPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.NotificationSinkPostExecute(r)
}

/*
NotificationSinkPost Send notifications about performance threshold crossing

To send a notifyThresholdCrossing notification

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiNotificationSinkPostRequest
*/
func (a *DefaultApiService) NotificationSinkPost(ctx context.Context) ApiNotificationSinkPostRequest {
	return ApiNotificationSinkPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) NotificationSinkPostExecute(r ApiNotificationSinkPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.NotificationSinkPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/notificationSink"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.notifyThresholdCrossing == nil {
		return nil, reportError("notifyThresholdCrossing is required and must be specified")
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
	localVarPostBody = r.notifyThresholdCrossing
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
		var v ErrorResponse
		err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
		if err != nil {
			newErr.error = err.Error()
			return localVarHTTPResponse, newErr
		}
		newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
		newErr.model = v
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
