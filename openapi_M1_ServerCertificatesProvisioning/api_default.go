/*
M1_ServerCertificatesProvisioning

5GMS AF M1 Server Certificates Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ServerCertificatesProvisioning

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// DefaultApiService DefaultApi service
type DefaultApiService service

type ApiCreateOrReserveServerCertificateRequest struct {
	ctx                   context.Context
	ApiService            *DefaultApiService
	provisioningSessionId string
	csr                   *string
}

// When present, return a Certificate Signing Request instead of generating an X.509 certificate
func (r ApiCreateOrReserveServerCertificateRequest) Csr(csr string) ApiCreateOrReserveServerCertificateRequest {
	r.csr = &csr
	return r
}

func (r ApiCreateOrReserveServerCertificateRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.CreateOrReserveServerCertificateExecute(r)
}

/*
CreateOrReserveServerCertificate Create or reserve a Service Certificate resource

Without the optional csr query parameter, an X.509 certificate is generated and this is returned. If the csr query parameter is present, a Certificate Signing Request is instead generated and returned, allowing the X.509 certificate to be generated by the invoker and later uploaded.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param provisioningSessionId The resource identifier of an existing Provisioning Session.'
	@return ApiCreateOrReserveServerCertificateRequest
*/
func (a *DefaultApiService) CreateOrReserveServerCertificate(ctx context.Context, provisioningSessionId string) ApiCreateOrReserveServerCertificateRequest {
	return ApiCreateOrReserveServerCertificateRequest{
		ApiService:            a,
		ctx:                   ctx,
		provisioningSessionId: provisioningSessionId,
	}
}

// Execute executes the request
//
//	@return string
func (a *DefaultApiService) CreateOrReserveServerCertificateExecute(r ApiCreateOrReserveServerCertificateRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.CreateOrReserveServerCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/certificates"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterValueToString(r.provisioningSessionId, "provisioningSessionId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.csr != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "csr", r.csr, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-pem-file"}

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

type ApiDestroyServerCertificateRequest struct {
	ctx                   context.Context
	ApiService            *DefaultApiService
	provisioningSessionId string
	certificateId         string
}

func (r ApiDestroyServerCertificateRequest) Execute() (*http.Response, error) {
	return r.ApiService.DestroyServerCertificateExecute(r)
}

/*
DestroyServerCertificate Destroy an existing Server Certificate resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param provisioningSessionId The resource identifier of an existing Provisioning Session.'
	@param certificateId The resource identifier of an existing Server Certificate
	@return ApiDestroyServerCertificateRequest
*/
func (a *DefaultApiService) DestroyServerCertificate(ctx context.Context, provisioningSessionId string, certificateId string) ApiDestroyServerCertificateRequest {
	return ApiDestroyServerCertificateRequest{
		ApiService:            a,
		ctx:                   ctx,
		provisioningSessionId: provisioningSessionId,
		certificateId:         certificateId,
	}
}

// Execute executes the request
func (a *DefaultApiService) DestroyServerCertificateExecute(r ApiDestroyServerCertificateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.DestroyServerCertificate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/certificates/{certificateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterValueToString(r.provisioningSessionId, "provisioningSessionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"certificateId"+"}", url.PathEscape(parameterValueToString(r.certificateId, "certificateId")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiRetrieveServerCertificateRequest struct {
	ctx                   context.Context
	ApiService            *DefaultApiService
	provisioningSessionId string
	certificateId         string
}

func (r ApiRetrieveServerCertificateRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.RetrieveServerCertificateExecute(r)
}

/*
RetrieveServerCertificate Retrieve the X.509 certificate representation of the specified Server Certificate resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param provisioningSessionId The resource identifier of an existing Provisioning Session.'
	@param certificateId The resource identifier of an existing Server Certificate
	@return ApiRetrieveServerCertificateRequest
*/
func (a *DefaultApiService) RetrieveServerCertificate(ctx context.Context, provisioningSessionId string, certificateId string) ApiRetrieveServerCertificateRequest {
	return ApiRetrieveServerCertificateRequest{
		ApiService:            a,
		ctx:                   ctx,
		provisioningSessionId: provisioningSessionId,
		certificateId:         certificateId,
	}
}

// Execute executes the request
//
//	@return string
func (a *DefaultApiService) RetrieveServerCertificateExecute(r ApiRetrieveServerCertificateRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.RetrieveServerCertificate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/certificates/{certificateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterValueToString(r.provisioningSessionId, "provisioningSessionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"certificateId"+"}", url.PathEscape(parameterValueToString(r.certificateId, "certificateId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/x-pem-file"}

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

type ApiUploadServerCertificateRequest struct {
	ctx                   context.Context
	ApiService            *DefaultApiService
	provisioningSessionId string
	certificateId         string
	body                  *string
}

func (r ApiUploadServerCertificateRequest) Body(body string) ApiUploadServerCertificateRequest {
	r.body = &body
	return r
}

func (r ApiUploadServerCertificateRequest) Execute() (*http.Response, error) {
	return r.ApiService.UploadServerCertificateExecute(r)
}

/*
UploadServerCertificate Upload the X.509 certificate for a previously reserved Server Certificate resource

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param provisioningSessionId The resource identifier of an existing Provisioning Session.'
	@param certificateId The resource identifier of an existing Server Certificate
	@return ApiUploadServerCertificateRequest
*/
func (a *DefaultApiService) UploadServerCertificate(ctx context.Context, provisioningSessionId string, certificateId string) ApiUploadServerCertificateRequest {
	return ApiUploadServerCertificateRequest{
		ApiService:            a,
		ctx:                   ctx,
		provisioningSessionId: provisioningSessionId,
		certificateId:         certificateId,
	}
}

// Execute executes the request
func (a *DefaultApiService) UploadServerCertificateExecute(r ApiUploadServerCertificateRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultApiService.UploadServerCertificate")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/provisioning-sessions/{provisioningSessionId}/certificates/{certificateId}"
	localVarPath = strings.Replace(localVarPath, "{"+"provisioningSessionId"+"}", url.PathEscape(parameterValueToString(r.provisioningSessionId, "provisioningSessionId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"certificateId"+"}", url.PathEscape(parameterValueToString(r.certificateId, "certificateId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-pem-file"}

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
