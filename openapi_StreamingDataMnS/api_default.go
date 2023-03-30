/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_StreamingDataMnS

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
	"reflect"
)


// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiConnectionsConnectionIdGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	connectionId string
	connection *WebsocketHeaderConnectionType
	secWebSocketExtensions *string
	secWebSocketKey *string
	secWebSocketProtocol *string
	secWebSocketVersion *string
}

func (r ApiConnectionsConnectionIdGetRequest) Connection(connection WebsocketHeaderConnectionType) ApiConnectionsConnectionIdGetRequest {
	r.connection = &connection
	return r
}

func (r ApiConnectionsConnectionIdGetRequest) SecWebSocketExtensions(secWebSocketExtensions string) ApiConnectionsConnectionIdGetRequest {
	r.secWebSocketExtensions = &secWebSocketExtensions
	return r
}

func (r ApiConnectionsConnectionIdGetRequest) SecWebSocketKey(secWebSocketKey string) ApiConnectionsConnectionIdGetRequest {
	r.secWebSocketKey = &secWebSocketKey
	return r
}

func (r ApiConnectionsConnectionIdGetRequest) SecWebSocketProtocol(secWebSocketProtocol string) ApiConnectionsConnectionIdGetRequest {
	r.secWebSocketProtocol = &secWebSocketProtocol
	return r
}

func (r ApiConnectionsConnectionIdGetRequest) SecWebSocketVersion(secWebSocketVersion string) ApiConnectionsConnectionIdGetRequest {
	r.secWebSocketVersion = &secWebSocketVersion
	return r
}

func (r ApiConnectionsConnectionIdGetRequest) Execute() (*ConnectionInfoType, *http.Response, error) {
	return r.ApiService.ConnectionsConnectionIdGetExecute(r)
}

/*
ConnectionsConnectionIdGet Obtain information about a connection.

Enables the streaming data reporting service producer to obtain information about one streaming connection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId Indicate the ID (URI) of the connection for which the information is being retrieved
 @return ApiConnectionsConnectionIdGetRequest
*/
func (a *DefaultApiService) ConnectionsConnectionIdGet(ctx context.Context, connectionId string) ApiConnectionsConnectionIdGetRequest {
	return ApiConnectionsConnectionIdGetRequest{
		ApiService: a,
		ctx: ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return ConnectionInfoType
func (a *DefaultApiService) ConnectionsConnectionIdGetExecute(r ApiConnectionsConnectionIdGetRequest) (*ConnectionInfoType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConnectionInfoType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectionsConnectionIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{connectionId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterValueToString(r.connectionId, "connectionId")), -1)

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
	if r.connection != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Connection", r.connection, "")
	}
	if r.secWebSocketExtensions != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Sec-WebSocket-Extensions", r.secWebSocketExtensions, "")
	}
	if r.secWebSocketKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Sec-WebSocket-Key", r.secWebSocketKey, "")
	}
	if r.secWebSocketProtocol != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Sec-WebSocket-Protocol", r.secWebSocketProtocol, "")
	}
	if r.secWebSocketVersion != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Sec-WebSocket-Version", r.secWebSocketVersion, "")
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
			var v ErrorResponseType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiConnectionsConnectionIdStreamsDeleteRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	connectionId string
	streamIds *[]string
}

// The list of streamId for the stream(s) to be deleted.
func (r ApiConnectionsConnectionIdStreamsDeleteRequest) StreamIds(streamIds []string) ApiConnectionsConnectionIdStreamsDeleteRequest {
	r.streamIds = &streamIds
	return r
}

func (r ApiConnectionsConnectionIdStreamsDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ConnectionsConnectionIdStreamsDeleteExecute(r)
}

/*
ConnectionsConnectionIdStreamsDelete Remove reporting streams from an existing connection

Allows the producer to remove one or more reporting streams from an already established streaming connection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId Indicate the ID (URI) of the connection for which the reporting stream information is being removed.
 @return ApiConnectionsConnectionIdStreamsDeleteRequest
*/
func (a *DefaultApiService) ConnectionsConnectionIdStreamsDelete(ctx context.Context, connectionId string) ApiConnectionsConnectionIdStreamsDeleteRequest {
	return ApiConnectionsConnectionIdStreamsDeleteRequest{
		ApiService: a,
		ctx: ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
func (a *DefaultApiService) ConnectionsConnectionIdStreamsDeleteExecute(r ApiConnectionsConnectionIdStreamsDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectionsConnectionIdStreamsDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{connectionId}/streams"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterValueToString(r.connectionId, "connectionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamIds == nil {
		return nil, reportError("streamIds is required and must be specified")
	}

	{
		t := *r.streamIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "streamIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "streamIds", t, "multi")
		}
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
			var v ErrorResponseType
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

type ApiConnectionsConnectionIdStreamsGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	connectionId string
	streamIds *[]string
}

// The list of streamId for which the stream information is to be retrieved.
func (r ApiConnectionsConnectionIdStreamsGetRequest) StreamIds(streamIds []string) ApiConnectionsConnectionIdStreamsGetRequest {
	r.streamIds = &streamIds
	return r
}

func (r ApiConnectionsConnectionIdStreamsGetRequest) Execute() ([]StreamInfoWithReportersType, *http.Response, error) {
	return r.ApiService.ConnectionsConnectionIdStreamsGetExecute(r)
}

/*
ConnectionsConnectionIdStreamsGet Obtain information about streams.

Enables the streaming data reporting service producer to obtain information about one or more reporting streams.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId Indicate the ID (URI) of the connection for which the information is being retrieved
 @return ApiConnectionsConnectionIdStreamsGetRequest
*/
func (a *DefaultApiService) ConnectionsConnectionIdStreamsGet(ctx context.Context, connectionId string) ApiConnectionsConnectionIdStreamsGetRequest {
	return ApiConnectionsConnectionIdStreamsGetRequest{
		ApiService: a,
		ctx: ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return []StreamInfoWithReportersType
func (a *DefaultApiService) ConnectionsConnectionIdStreamsGetExecute(r ApiConnectionsConnectionIdStreamsGetRequest) ([]StreamInfoWithReportersType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []StreamInfoWithReportersType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectionsConnectionIdStreamsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{connectionId}/streams"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterValueToString(r.connectionId, "connectionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamIds == nil {
		return localVarReturnValue, nil, reportError("streamIds is required and must be specified")
	}

	{
		t := *r.streamIds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "streamIds", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "streamIds", t, "multi")
		}
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
			var v ErrorResponseType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiConnectionsConnectionIdStreamsPostRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	connectionId string
	streamInfoType *[]StreamInfoType
}

func (r ApiConnectionsConnectionIdStreamsPostRequest) StreamInfoType(streamInfoType []StreamInfoType) ApiConnectionsConnectionIdStreamsPostRequest {
	r.streamInfoType = &streamInfoType
	return r
}

func (r ApiConnectionsConnectionIdStreamsPostRequest) Execute() ([]StreamInfoType, *http.Response, error) {
	return r.ApiService.ConnectionsConnectionIdStreamsPostExecute(r)
}

/*
ConnectionsConnectionIdStreamsPost Inform consumer about new reporting streams on an existing connection.

Allows the producer to add one or more reporting streams to an already established streaming connection.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId Indicate the ID (URI) of the connection for which the reporting stream information is being added.
 @return ApiConnectionsConnectionIdStreamsPostRequest
*/
func (a *DefaultApiService) ConnectionsConnectionIdStreamsPost(ctx context.Context, connectionId string) ApiConnectionsConnectionIdStreamsPostRequest {
	return ApiConnectionsConnectionIdStreamsPostRequest{
		ApiService: a,
		ctx: ctx,
		connectionId: connectionId,
	}
}

// Execute executes the request
//  @return []StreamInfoType
func (a *DefaultApiService) ConnectionsConnectionIdStreamsPostExecute(r ApiConnectionsConnectionIdStreamsPostRequest) ([]StreamInfoType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []StreamInfoType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectionsConnectionIdStreamsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{connectionId}/streams"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterValueToString(r.connectionId, "connectionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.streamInfoType == nil {
		return localVarReturnValue, nil, reportError("streamInfoType is required and must be specified")
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
	localVarPostBody = r.streamInfoType
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
			var v ErrorResponseType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiConnectionsConnectionIdStreamsStreamIdGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	connectionId string
	streamId string
}

func (r ApiConnectionsConnectionIdStreamsStreamIdGetRequest) Execute() (*StreamInfoWithReportersType, *http.Response, error) {
	return r.ApiService.ConnectionsConnectionIdStreamsStreamIdGetExecute(r)
}

/*
ConnectionsConnectionIdStreamsStreamIdGet Obtain information about stream

Enables the streaming data reporting service producer to obtain information about a reporting stream.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param connectionId Indicate the ID (URI) of the connection for which the information is being retrieved
 @param streamId Indicate the ID of the reporting stream for which the information is being retrieved
 @return ApiConnectionsConnectionIdStreamsStreamIdGetRequest
*/
func (a *DefaultApiService) ConnectionsConnectionIdStreamsStreamIdGet(ctx context.Context, connectionId string, streamId string) ApiConnectionsConnectionIdStreamsStreamIdGetRequest {
	return ApiConnectionsConnectionIdStreamsStreamIdGetRequest{
		ApiService: a,
		ctx: ctx,
		connectionId: connectionId,
		streamId: streamId,
	}
}

// Execute executes the request
//  @return StreamInfoWithReportersType
func (a *DefaultApiService) ConnectionsConnectionIdStreamsStreamIdGetExecute(r ApiConnectionsConnectionIdStreamsStreamIdGetRequest) (*StreamInfoWithReportersType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *StreamInfoWithReportersType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectionsConnectionIdStreamsStreamIdGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections/{connectionId}/streams/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"connectionId"+"}", url.PathEscape(parameterValueToString(r.connectionId, "connectionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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
			var v ErrorResponseType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiConnectionsGetRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	connectionIdList *[]string
}

// The list of connectionId for which the connection information is to be returned.
func (r ApiConnectionsGetRequest) ConnectionIdList(connectionIdList []string) ApiConnectionsGetRequest {
	r.connectionIdList = &connectionIdList
	return r
}

func (r ApiConnectionsGetRequest) Execute() ([]ConnectionInfoType, *http.Response, error) {
	return r.ApiService.ConnectionsGetExecute(r)
}

/*
ConnectionsGet Obtain information about connections.

Enables the streaming data reporting service producer to obtain information about one or more streaming connections.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConnectionsGetRequest
*/
func (a *DefaultApiService) ConnectionsGet(ctx context.Context) ApiConnectionsGetRequest {
	return ApiConnectionsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ConnectionInfoType
func (a *DefaultApiService) ConnectionsGetExecute(r ApiConnectionsGetRequest) ([]ConnectionInfoType, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ConnectionInfoType
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectionsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.connectionIdList != nil {
		t := *r.connectionIdList
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "connectionIdList", s.Index(i), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "connectionIdList", t, "multi")
		}
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
			var v ErrorResponseType
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiConnectionsPostRequest struct {
	ctx context.Context
	ApiService *DefaultApiService
	connectionRequestType *ConnectionRequestType
}

func (r ApiConnectionsPostRequest) ConnectionRequestType(connectionRequestType ConnectionRequestType) ApiConnectionsPostRequest {
	r.connectionRequestType = &connectionRequestType
	return r
}

func (r ApiConnectionsPostRequest) Execute() (*http.Response, error) {
	return r.ApiService.ConnectionsPostExecute(r)
}

/*
ConnectionsPost Inform consumer about reporting streams to be carried by the new connection and receive a new connection id.

Exchange of meta-data (producer informs consumer about its own identity and the nature of the data to be reported via streaming) phase of the connection establishement by streaming data reporting producer to the streaming data reporting consumer (i.e. streaming target).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiConnectionsPostRequest
*/
func (a *DefaultApiService) ConnectionsPost(ctx context.Context) ApiConnectionsPostRequest {
	return ApiConnectionsPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *DefaultApiService) ConnectionsPostExecute(r ApiConnectionsPostRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.ConnectionsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/connections"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.connectionRequestType == nil {
		return nil, reportError("connectionRequestType is required and must be specified")
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
	localVarPostBody = r.connectionRequestType
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
			var v FailedConnectionResponseType
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
