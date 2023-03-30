# Go API client for openapi_Namf_MBSBroadcast

AMF MBSBroadcast Service.  
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
import openapi_Namf_MBSBroadcast "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Namf_MBSBroadcast"
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
ctx := context.WithValue(context.Background(), openapi_Namf_MBSBroadcast.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Namf_MBSBroadcast.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Namf_MBSBroadcast.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Namf_MBSBroadcast.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/namf-mbs-bc/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BroadcastMBSSessionContextsCollectionCollectionApi* | [**ContextCreate**](docs/BroadcastMBSSessionContextsCollectionCollectionApi.md#contextcreate) | **Post** /mbs-contexts | Namf_MBSBroadcast ContextCreate service Operation
*IndividualBroadcastMBSSessionContextDocumentApi* | [**ContextDelete**](docs/IndividualBroadcastMBSSessionContextDocumentApi.md#contextdelete) | **Delete** /mbs-contexts/{mbsContextRef} | Namf_MBSBroadcast ContextDelete service Operation
*IndividualBroadcastMBSSessionContextDocumentApi* | [**ContextUpdate**](docs/IndividualBroadcastMBSSessionContextDocumentApi.md#contextupdate) | **Post** /mbs-contexts/{mbsContextRef}/update | Namf_MBSBroadcast ContextUpdate service Operation


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [ContextCreateReqData](docs/ContextCreateReqData.md)
 - [ContextCreateRspData](docs/ContextCreateRspData.md)
 - [ContextStatusNotification](docs/ContextStatusNotification.md)
 - [ContextUpdate200Response](docs/ContextUpdate200Response.md)
 - [ContextUpdateReqData](docs/ContextUpdateReqData.md)
 - [ContextUpdateRequest](docs/ContextUpdateRequest.md)
 - [ContextUpdateRspData](docs/ContextUpdateRspData.md)
 - [GNbId](docs/GNbId.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [MbsContextsPostRequest](docs/MbsContextsPostRequest.md)
 - [MbsServiceArea](docs/MbsServiceArea.md)
 - [MbsServiceAreaInfo](docs/MbsServiceAreaInfo.md)
 - [MbsSessionId](docs/MbsSessionId.md)
 - [N2MbsSmInfo](docs/N2MbsSmInfo.md)
 - [NFType](docs/NFType.md)
 - [Ncgi](docs/Ncgi.md)
 - [NcgiTai](docs/NcgiTai.md)
 - [NgapIeType](docs/NgapIeType.md)
 - [NgranFailureEvent](docs/NgranFailureEvent.md)
 - [NgranFailureIndication](docs/NgranFailureIndication.md)
 - [OpEventType](docs/OpEventType.md)
 - [OperationEvent](docs/OperationEvent.md)
 - [OperationStatus](docs/OperationStatus.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [RefToBinaryData](docs/RefToBinaryData.md)
 - [Snssai](docs/Snssai.md)
 - [Ssm](docs/Ssm.md)
 - [Tai](docs/Tai.md)
 - [Tmgi](docs/Tmgi.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **namf-mbs-bc**: Access to the Namf_MBSBroadcast API

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


