# Go API client for openapi_AsSessionWithQoS

API for setting us an AS session with required QoS.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.2.2
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
import openapi_AsSessionWithQoS "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_AsSessionWithQoS"
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
ctx := context.WithValue(context.Background(), openapi_AsSessionWithQoS.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_AsSessionWithQoS.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_AsSessionWithQoS.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_AsSessionWithQoS.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/3gpp-as-session-with-qos/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ASSessionWithRequiredQoSSubscriptionsApi* | [**CreateASSessionWithQoSSubscription**](docs/ASSessionWithRequiredQoSSubscriptionsApi.md#createassessionwithqossubscription) | **Post** /{scsAsId}/subscriptions | Creates a new subscription resource.
*ASSessionWithRequiredQoSSubscriptionsApi* | [**FetchAllASSessionWithQoSSubscriptions**](docs/ASSessionWithRequiredQoSSubscriptionsApi.md#fetchallassessionwithqossubscriptions) | **Get** /{scsAsId}/subscriptions | Read all or queried active subscriptions for the SCS/AS.
*IndividualASSessionWithRequiredQoSSubscriptionApi* | [**DeleteIndASSessionWithQoSSubscription**](docs/IndividualASSessionWithRequiredQoSSubscriptionApi.md#deleteindassessionwithqossubscription) | **Delete** /{scsAsId}/subscriptions/{subscriptionId} | Deletes an already existing subscription.
*IndividualASSessionWithRequiredQoSSubscriptionApi* | [**FetchIndASSessionWithQoSSubscription**](docs/IndividualASSessionWithRequiredQoSSubscriptionApi.md#fetchindassessionwithqossubscription) | **Get** /{scsAsId}/subscriptions/{subscriptionId} | Read an active subscriptions for the SCS/AS and the subscription Id.
*IndividualASSessionWithRequiredQoSSubscriptionApi* | [**ModifyIndASSessionWithQoSSubscription**](docs/IndividualASSessionWithRequiredQoSSubscriptionApi.md#modifyindassessionwithqossubscription) | **Patch** /{scsAsId}/subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource.
*IndividualASSessionWithRequiredQoSSubscriptionApi* | [**UpdateIndASSessionWithQoSSubscription**](docs/IndividualASSessionWithRequiredQoSSubscriptionApi.md#updateindassessionwithqossubscription) | **Put** /{scsAsId}/subscriptions/{subscriptionId} | Updates/replaces an existing subscription resource.


## Documentation For Models

 - [AccumulatedUsage](docs/AccumulatedUsage.md)
 - [AlternativeServiceRequirementsData](docs/AlternativeServiceRequirementsData.md)
 - [AsSessionWithQoSSubscription](docs/AsSessionWithQoSSubscription.md)
 - [AsSessionWithQoSSubscriptionPatch](docs/AsSessionWithQoSSubscriptionPatch.md)
 - [EthFlowDescription](docs/EthFlowDescription.md)
 - [EthFlowInfo](docs/EthFlowInfo.md)
 - [FlowDirection](docs/FlowDirection.md)
 - [FlowInfo](docs/FlowInfo.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [QosMonitoringInformation](docs/QosMonitoringInformation.md)
 - [QosMonitoringInformationRm](docs/QosMonitoringInformationRm.md)
 - [QosMonitoringReport](docs/QosMonitoringReport.md)
 - [RatType](docs/RatType.md)
 - [ReportingFrequency](docs/ReportingFrequency.md)
 - [RequestedQosMonitoringParameter](docs/RequestedQosMonitoringParameter.md)
 - [Snssai](docs/Snssai.md)
 - [SponsorInformation](docs/SponsorInformation.md)
 - [TscQosRequirement](docs/TscQosRequirement.md)
 - [TscQosRequirementRm](docs/TscQosRequirementRm.md)
 - [TscaiInputContainer](docs/TscaiInputContainer.md)
 - [UsageThreshold](docs/UsageThreshold.md)
 - [UsageThresholdRm](docs/UsageThresholdRm.md)
 - [UserPlaneEvent](docs/UserPlaneEvent.md)
 - [UserPlaneEventReport](docs/UserPlaneEventReport.md)
 - [UserPlaneNotificationData](docs/UserPlaneNotificationData.md)
 - [WebsockNotifConfig](docs/WebsockNotifConfig.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: N/A

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



