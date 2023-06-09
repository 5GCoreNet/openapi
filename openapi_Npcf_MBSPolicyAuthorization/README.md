# Go API client for openapi_Npcf_MBSPolicyAuthorization

MBS Policy Authorization Service.  
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
import openapi_Npcf_MBSPolicyAuthorization "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Npcf_MBSPolicyAuthorization"
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
ctx := context.WithValue(context.Background(), openapi_Npcf_MBSPolicyAuthorization.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Npcf_MBSPolicyAuthorization.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Npcf_MBSPolicyAuthorization.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Npcf_MBSPolicyAuthorization.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/npcf-mbspolicyauth/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualMBSApplicationSessionContextDocumentApi* | [**DeleteMBSAppSessionCtxt**](docs/IndividualMBSApplicationSessionContextDocumentApi.md#deletembsappsessionctxt) | **Delete** /contexts/{contextId} | Request the deletion of an existing Individual MBS Application Session Context resource.
*IndividualMBSApplicationSessionContextDocumentApi* | [**GetMBSAppSessionCtxt**](docs/IndividualMBSApplicationSessionContextDocumentApi.md#getmbsappsessionctxt) | **Get** /contexts/{contextId} | Read an existing Individual MBS Application Session Context resource.
*IndividualMBSApplicationSessionContextDocumentApi* | [**ModifyMBSAppSessionCtxt**](docs/IndividualMBSApplicationSessionContextDocumentApi.md#modifymbsappsessionctxt) | **Patch** /contexts/{contextId} | Request the modification of an existing Individual MBS Application Session Context resource.
*MBSApplicationSessionContextsCollectionApi* | [**CreateMBSAppSessionCtxt**](docs/MBSApplicationSessionContextsCollectionApi.md#creatembsappsessionctxt) | **Post** /contexts | Request the creation of a new Individual MBS Application Session Context resource.


## Documentation For Models

 - [AcceptableMbsServInfo](docs/AcceptableMbsServInfo.md)
 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [Arp](docs/Arp.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [MbsAppSessionCtxt](docs/MbsAppSessionCtxt.md)
 - [MbsAppSessionCtxtPatch](docs/MbsAppSessionCtxtPatch.md)
 - [MbsExtProblemDetails](docs/MbsExtProblemDetails.md)
 - [MbsMediaComp](docs/MbsMediaComp.md)
 - [MbsMediaCompRm](docs/MbsMediaCompRm.md)
 - [MbsMediaInfo](docs/MbsMediaInfo.md)
 - [MbsQoSReq](docs/MbsQoSReq.md)
 - [MbsServiceInfo](docs/MbsServiceInfo.md)
 - [MbsSessionId](docs/MbsSessionId.md)
 - [MediaType](docs/MediaType.md)
 - [NFType](docs/NFType.md)
 - [NullValue](docs/NullValue.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [PreemptionCapability](docs/PreemptionCapability.md)
 - [PreemptionVulnerability](docs/PreemptionVulnerability.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [ReservPriority](docs/ReservPriority.md)
 - [Snssai](docs/Snssai.md)
 - [Ssm](docs/Ssm.md)
 - [Tmgi](docs/Tmgi.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **npcf-mbspolicyauth**: Access to the Npcf_MBSPolicyAuthorization API

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



