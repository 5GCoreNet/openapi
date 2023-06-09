# Go API client for openapi_N5g_ddnmf_Discovery

N5g-ddnmf_Discovery Service.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0-alpha.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi_N5g_ddnmf_Discovery "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_N5g_ddnmf_Discovery"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi_N5g_ddnmf_Discovery.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_N5g_ddnmf_Discovery.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_N5g_ddnmf_Discovery.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_N5g_ddnmf_Discovery.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/n5g-ddnmf-disc/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ObtainTheAuthorizationForADiscovererUEApi* | [**ObtainDiscAuth**](docs/ObtainTheAuthorizationForADiscovererUEApi.md#obtaindiscauth) | **Put** /{ueId}/discovery-authorize/{discEntryId} | Obtain the authorization from the 5G DDNMF for a discoverer UE in the PLMN to operate Model B restricted discovery
*ObtainTheAuthorizationToAnnounceForAUEApi* | [**ObtainAnnounceAuth**](docs/ObtainTheAuthorizationToAnnounceForAUEApi.md#obtainannounceauth) | **Put** /{ueId}/announce-authorize/{discEntryId} | Obtain the authorization to announce for a UE
*ObtainTheAuthorizationToMonitorForAUEApi* | [**ObtainMonitorAuth**](docs/ObtainTheAuthorizationToMonitorForAUEApi.md#obtainmonitorauth) | **Put** /{ueId}/monitor-authorize/{discEntryId} | Obtain the authorization to monitor for a UE
*ObtainTheInformationAboutTheIndicatedDiscoveryCodeApi* | [**MatchReport**](docs/ObtainTheInformationAboutTheIndicatedDiscoveryCodeApi.md#matchreport) | **Post** /{ueId}/match-report | Obtain the information about the indicated discovery code from the 5G DDNMF
*UpdateTheAuthorizationForAnnouncingForAUEApi* | [**UpdateAnnounceAuth**](docs/UpdateTheAuthorizationForAnnouncingForAUEApi.md#updateannounceauth) | **Patch** /{ueId}/announce-authorize/{discEntryId} | Update the authorization for announcing for a UE
*UpdateTheAuthorizationForMonitoringForAUEApi* | [**UpdateMonitorAuth**](docs/UpdateTheAuthorizationForMonitoringForAUEApi.md#updatemonitorauth) | **Patch** /{ueId}/monitor-authorize/{discEntryId} | Update the authorization for monitoring for a UE


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AnnounceAuthReqData](docs/AnnounceAuthReqData.md)
 - [AnnounceDiscDataForOpen](docs/AnnounceDiscDataForOpen.md)
 - [AnnounceDiscDataForRestricted](docs/AnnounceDiscDataForRestricted.md)
 - [AnnounceUpdateData](docs/AnnounceUpdateData.md)
 - [AuthDataForRestricted](docs/AuthDataForRestricted.md)
 - [DiscDataForRestricted](docs/DiscDataForRestricted.md)
 - [DiscoveryAuthReqData](docs/DiscoveryAuthReqData.md)
 - [DiscoveryAuthRespData](docs/DiscoveryAuthRespData.md)
 - [DiscoveryType](docs/DiscoveryType.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [MatchInfoForOpen](docs/MatchInfoForOpen.md)
 - [MatchInfoForRestricted](docs/MatchInfoForRestricted.md)
 - [MatchInformation](docs/MatchInformation.md)
 - [MatchReportReqData](docs/MatchReportReqData.md)
 - [MatchReportRespData](docs/MatchReportRespData.md)
 - [MonitorAuthDataForOpen](docs/MonitorAuthDataForOpen.md)
 - [MonitorAuthDataForRestricted](docs/MonitorAuthDataForRestricted.md)
 - [MonitorAuthReqData](docs/MonitorAuthReqData.md)
 - [MonitorAuthRespData](docs/MonitorAuthRespData.md)
 - [MonitorDiscDataForOpen](docs/MonitorDiscDataForOpen.md)
 - [MonitorDiscDataForRestricted](docs/MonitorDiscDataForRestricted.md)
 - [MonitorUpdateData](docs/MonitorUpdateData.md)
 - [MonitorUpdateDataForOpen](docs/MonitorUpdateDataForOpen.md)
 - [MonitorUpdateDataForRestricted](docs/MonitorUpdateDataForRestricted.md)
 - [MonitorUpdateResult](docs/MonitorUpdateResult.md)
 - [NFType](docs/NFType.md)
 - [PatchResult](docs/PatchResult.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [ProseAppCodeSuffixRange](docs/ProseAppCodeSuffixRange.md)
 - [ProseApplicationCodeSuffixPool](docs/ProseApplicationCodeSuffixPool.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [ReportItem](docs/ReportItem.md)
 - [RestrictedCodeSuffixPool](docs/RestrictedCodeSuffixPool.md)
 - [RestrictedCodeSuffixRange](docs/RestrictedCodeSuffixRange.md)
 - [RevocationResult](docs/RevocationResult.md)
 - [Snssai](docs/Snssai.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **n5g-ddnmf-disc**: Access to the N5g-ddnmf_Discovery API

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```golang
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, sw.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



