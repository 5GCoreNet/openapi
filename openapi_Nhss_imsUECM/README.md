# Go API client for openapi_Nhss_imsUECM

Nhss UE Context Management Service for IMS.  
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
import openapi_Nhss_imsUECM "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Nhss_imsUECM"
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
ctx := context.WithValue(context.Background(), openapi_Nhss_imsUECM.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Nhss_imsUECM.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Nhss_imsUECM.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Nhss_imsUECM.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nhss-ims-uecm/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthorizeCustomOperationApi* | [**Authorize**](docs/AuthorizeCustomOperationApi.md#authorize) | **Post** /{impu}/authorize | Authorize IMS Identities to register in the network or establish multimedia sessions and return CSCF location if it is stored 
*SCSCFRegistrationDocumentApi* | [**SCSCFRegistration**](docs/SCSCFRegistrationDocumentApi.md#scscfregistration) | **Put** /{imsUeId}/scscf-registration | SCSCF registration information
*SCSCFRestorationInformationDocumentApi* | [**DeleteScscfRestorationInfo**](docs/SCSCFRestorationInformationDocumentApi.md#deletescscfrestorationinfo) | **Delete** /{impu}/scscf-registration/scscf-restoration-info | Delete the S-CSCF restoration information of the UE
*SCSCFRestorationInformationDocumentApi* | [**GetScscfRestorationInfo**](docs/SCSCFRestorationInformationDocumentApi.md#getscscfrestorationinfo) | **Get** /{impu}/scscf-registration/scscf-restoration-info | Retrieve the S-CSCF restoration information of the UE
*SCSCFRestorationInformationDocumentApi* | [**UpdateScscfRestorationInfo**](docs/SCSCFRestorationInformationDocumentApi.md#updatescscfrestorationinfo) | **Put** /{impu}/scscf-registration/scscf-restoration-info | Update the S-CSCF restoration information of the UE


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AdditionalInfo](docs/AdditionalInfo.md)
 - [AuthorizationRequest](docs/AuthorizationRequest.md)
 - [AuthorizationResponse](docs/AuthorizationResponse.md)
 - [AuthorizationResult](docs/AuthorizationResult.md)
 - [AuthorizationType](docs/AuthorizationType.md)
 - [DeregistrationData](docs/DeregistrationData.md)
 - [DeregistrationReason](docs/DeregistrationReason.md)
 - [DeregistrationReasonCode](docs/DeregistrationReasonCode.md)
 - [EmergencyRegisteredIdentity](docs/EmergencyRegisteredIdentity.md)
 - [ExtendedProblemDetails](docs/ExtendedProblemDetails.md)
 - [ImsRegistrationType](docs/ImsRegistrationType.md)
 - [ImsSdmSubscription](docs/ImsSdmSubscription.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [LooseRouteIndication](docs/LooseRouteIndication.md)
 - [NFType](docs/NFType.md)
 - [PcscfSubscriptionInfo](docs/PcscfSubscriptionInfo.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [RestorationInfo](docs/RestorationInfo.md)
 - [ScscfCapabilityList](docs/ScscfCapabilityList.md)
 - [ScscfRegistration](docs/ScscfRegistration.md)
 - [ScscfRestorationInfo](docs/ScscfRestorationInfo.md)
 - [ScscfRestorationInfoRequest](docs/ScscfRestorationInfoRequest.md)
 - [ScscfRestorationInfoResponse](docs/ScscfRestorationInfoResponse.md)
 - [ScscfSelectionAssistanceInformation](docs/ScscfSelectionAssistanceInformation.md)
 - [SipAuthenticationScheme](docs/SipAuthenticationScheme.md)
 - [Snssai](docs/Snssai.md)
 - [UeSubscriptionInfo](docs/UeSubscriptionInfo.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nhss-ims-uecm**: Access to the Nhss IMS UE Context Management API
 - **nhss-ims-uecm:authorize:invoke**: Access to invoke the Authorize custom operation
 - **nhss-ims-uecm:registration:create**: Access to create the S-CSCF Registration resource
 - **nhss-ims-uecm:restoration:read**: Access to read the S-CSCF Restoration resource
 - **nhss-ims-uecm:restoration:modify**: Access to create/update/delete the S-CSCF Restoration resource

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



