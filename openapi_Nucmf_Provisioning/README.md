# Go API client for openapi_Nucmf_Provisioning

UCMF_Provisioning Service.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.2.0-alpha.1
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
import openapi_Nucmf_Provisioning "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Nucmf_Provisioning"
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
ctx := context.WithValue(context.Background(), openapi_Nucmf_Provisioning.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Nucmf_Provisioning.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Nucmf_Provisioning.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Nucmf_Provisioning.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nucmf-provisioning/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualUERadioCapabilityProvisioningDocumentApi* | [**GetProvisioning**](docs/IndividualUERadioCapabilityProvisioningDocumentApi.md#getprovisioning) | **Get** /provisionings/{provisioningId} | Get an Individual UE radio capability provisioning
*IndividualUERadioCapabilityProvisioningDocumentApi* | [**RemoveProvisioning**](docs/IndividualUERadioCapabilityProvisioningDocumentApi.md#removeprovisioning) | **Delete** /provisionings/{provisioningId} | Remove an Individual UE radio capability provisioning
*IndividualUERadioCapabilityProvisioningDocumentApi* | [**ReplaceProvisioning**](docs/IndividualUERadioCapabilityProvisioningDocumentApi.md#replaceprovisioning) | **Put** /provisionings/{provisioningId} | Replace (PUT) an Individual UE radio capability provisioning
*IndividualUERadioCapabilityProvisioningDocumentApi* | [**UpdateProvisioning**](docs/IndividualUERadioCapabilityProvisioningDocumentApi.md#updateprovisioning) | **Patch** /provisionings/{provisioningId} | Update (PATCH) an Individual UE radio capability provisioning
*UERadioCapabilityProvisioningsCollectionApi* | [**CreateProvisioning**](docs/UERadioCapabilityProvisioningsCollectionApi.md#createprovisioning) | **Post** /provisionings | Create an Individual UE radio capability provisioning


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [InvalidParam1](docs/InvalidParam1.md)
 - [NFType](docs/NFType.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [ProblemDetails1](docs/ProblemDetails1.md)
 - [RacsConfiguration](docs/RacsConfiguration.md)
 - [RacsConfigurationRm](docs/RacsConfigurationRm.md)
 - [RacsData](docs/RacsData.md)
 - [RacsDataPatch](docs/RacsDataPatch.md)
 - [RacsFailureCode](docs/RacsFailureCode.md)
 - [RacsFailureReport](docs/RacsFailureReport.md)
 - [Snssai](docs/Snssai.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nucmf-provisioning**: Access to the Nucmf_Provisioning API

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


