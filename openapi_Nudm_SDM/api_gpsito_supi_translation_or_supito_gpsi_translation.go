/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// GPSIToSUPITranslationOrSUPIToGPSITranslationApiService GPSIToSUPITranslationOrSUPIToGPSITranslationApi service
type GPSIToSUPITranslationOrSUPIToGPSITranslationApiService service

type ApiGetSupiOrGpsiRequest struct {
	ctx               context.Context
	ApiService        *GPSIToSUPITranslationOrSUPIToGPSITranslationApiService
	ueId              string
	supportedFeatures *string
	afId              *string
	appPortId         *AppPortId
	afServiceId       *string
	mtcProviderInfo   *string
	requestedGpsiType *GpsiType
	ifNoneMatch       *string
	ifModifiedSince   *string
}

// Supported Features
func (r ApiGetSupiOrGpsiRequest) SupportedFeatures(supportedFeatures string) ApiGetSupiOrGpsiRequest {
	r.supportedFeatures = &supportedFeatures
	return r
}

// AF identifier
func (r ApiGetSupiOrGpsiRequest) AfId(afId string) ApiGetSupiOrGpsiRequest {
	r.afId = &afId
	return r
}

// Application port identifier
func (r ApiGetSupiOrGpsiRequest) AppPortId(appPortId AppPortId) ApiGetSupiOrGpsiRequest {
	r.appPortId = &appPortId
	return r
}

// AF Service Identifier
func (r ApiGetSupiOrGpsiRequest) AfServiceId(afServiceId string) ApiGetSupiOrGpsiRequest {
	r.afServiceId = &afServiceId
	return r
}

// MTC Provider Information
func (r ApiGetSupiOrGpsiRequest) MtcProviderInfo(mtcProviderInfo string) ApiGetSupiOrGpsiRequest {
	r.mtcProviderInfo = &mtcProviderInfo
	return r
}

// Requested GPSI Type
func (r ApiGetSupiOrGpsiRequest) RequestedGpsiType(requestedGpsiType GpsiType) ApiGetSupiOrGpsiRequest {
	r.requestedGpsiType = &requestedGpsiType
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.2
func (r ApiGetSupiOrGpsiRequest) IfNoneMatch(ifNoneMatch string) ApiGetSupiOrGpsiRequest {
	r.ifNoneMatch = &ifNoneMatch
	return r
}

// Validator for conditional requests, as described in RFC 7232, 3.3
func (r ApiGetSupiOrGpsiRequest) IfModifiedSince(ifModifiedSince string) ApiGetSupiOrGpsiRequest {
	r.ifModifiedSince = &ifModifiedSince
	return r
}

func (r ApiGetSupiOrGpsiRequest) Execute() (*IdTranslationResult, *http.Response, error) {
	return r.ApiService.GetSupiOrGpsiExecute(r)
}

/*
GetSupiOrGpsi retrieve a UE's SUPI or GPSI

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param ueId Identifier of the UE
	@return ApiGetSupiOrGpsiRequest
*/
func (a *GPSIToSUPITranslationOrSUPIToGPSITranslationApiService) GetSupiOrGpsi(ctx context.Context, ueId string) ApiGetSupiOrGpsiRequest {
	return ApiGetSupiOrGpsiRequest{
		ApiService: a,
		ctx:        ctx,
		ueId:       ueId,
	}
}

// Execute executes the request
//
//	@return IdTranslationResult
func (a *GPSIToSUPITranslationOrSUPIToGPSITranslationApiService) GetSupiOrGpsiExecute(r ApiGetSupiOrGpsiRequest) (*IdTranslationResult, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *IdTranslationResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GPSIToSUPITranslationOrSUPIToGPSITranslationApiService.GetSupiOrGpsi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/{ueId}/id-translation-result"
	localVarPath = strings.Replace(localVarPath, "{"+"ueId"+"}", url.PathEscape(parameterValueToString(r.ueId, "ueId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.supportedFeatures != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "supported-features", r.supportedFeatures, "")
	}
	if r.afId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "af-id", r.afId, "")
	}
	if r.appPortId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "app-port-id", r.appPortId, "")
	}
	if r.afServiceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "af-service-id", r.afServiceId, "")
	}
	if r.mtcProviderInfo != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "mtc-provider-info", r.mtcProviderInfo, "")
	}
	if r.requestedGpsiType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "requested-gpsi-type", r.requestedGpsiType, "")
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
	if r.ifNoneMatch != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-None-Match", r.ifNoneMatch, "")
	}
	if r.ifModifiedSince != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "If-Modified-Since", r.ifModifiedSince, "")
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
