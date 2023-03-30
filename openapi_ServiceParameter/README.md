# Go API client for openapi_ServiceParameter

API for AF service paramter  
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
import openapi_ServiceParameter "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_ServiceParameter"
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
ctx := context.WithValue(context.Background(), openapi_ServiceParameter.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_ServiceParameter.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_ServiceParameter.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_ServiceParameter.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/3gpp-service-parameter/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualServiceParameterSubscriptionApi* | [**DeleteAnSubscription**](docs/IndividualServiceParameterSubscriptionApi.md#deleteansubscription) | **Delete** /{afId}/subscriptions/{subscriptionId} | Deletes an already existing subscription
*IndividualServiceParameterSubscriptionApi* | [**FullyUpdateAnSubscription**](docs/IndividualServiceParameterSubscriptionApi.md#fullyupdateansubscription) | **Put** /{afId}/subscriptions/{subscriptionId} | Fully updates/replaces an existing subscription resource
*IndividualServiceParameterSubscriptionApi* | [**PartialUpdateAnSubscription**](docs/IndividualServiceParameterSubscriptionApi.md#partialupdateansubscription) | **Patch** /{afId}/subscriptions/{subscriptionId} | Partial updates/replaces an existing subscription resource
*IndividualServiceParameterSubscriptionApi* | [**ReadAnSubscription**](docs/IndividualServiceParameterSubscriptionApi.md#readansubscription) | **Get** /{afId}/subscriptions/{subscriptionId} | read an active subscriptions for the SCS/AS and the subscription Id
*ServiceParameterSubscriptionsApi* | [**CreateAnSubscription**](docs/ServiceParameterSubscriptionsApi.md#createansubscription) | **Post** /{afId}/subscriptions | Creates a new subscription resource
*ServiceParameterSubscriptionsApi* | [**ReadAllSubscriptions**](docs/ServiceParameterSubscriptionsApi.md#readallsubscriptions) | **Get** /{afId}/subscriptions | read all of the active subscriptions for the AF


## Documentation For Models

 - [AfNotification](docs/AfNotification.md)
 - [AppDescriptor](docs/AppDescriptor.md)
 - [AuthorizationResult](docs/AuthorizationResult.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [ConnectionCapabilities](docs/ConnectionCapabilities.md)
 - [EllipsoidArc](docs/EllipsoidArc.md)
 - [EllipsoidArcAllOf](docs/EllipsoidArcAllOf.md)
 - [EthFlowDescription](docs/EthFlowDescription.md)
 - [Event](docs/Event.md)
 - [EventInfo](docs/EventInfo.md)
 - [Failure](docs/Failure.md)
 - [FlowDirection](docs/FlowDirection.md)
 - [GADShape](docs/GADShape.md)
 - [GeographicArea](docs/GeographicArea.md)
 - [GeographicalArea](docs/GeographicalArea.md)
 - [GeographicalCoordinates](docs/GeographicalCoordinates.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [Local2dPointUncertaintyEllipse](docs/Local2dPointUncertaintyEllipse.md)
 - [Local2dPointUncertaintyEllipseAllOf](docs/Local2dPointUncertaintyEllipseAllOf.md)
 - [Local3dPointUncertaintyEllipsoid](docs/Local3dPointUncertaintyEllipsoid.md)
 - [Local3dPointUncertaintyEllipsoidAllOf](docs/Local3dPointUncertaintyEllipsoidAllOf.md)
 - [LocalOrigin](docs/LocalOrigin.md)
 - [PlmnId](docs/PlmnId.md)
 - [Point](docs/Point.md)
 - [PointAllOf](docs/PointAllOf.md)
 - [PointAltitude](docs/PointAltitude.md)
 - [PointAltitudeAllOf](docs/PointAltitudeAllOf.md)
 - [PointAltitudeUncertainty](docs/PointAltitudeUncertainty.md)
 - [PointAltitudeUncertaintyAllOf](docs/PointAltitudeUncertaintyAllOf.md)
 - [PointUncertaintyCircle](docs/PointUncertaintyCircle.md)
 - [PointUncertaintyCircleAllOf](docs/PointUncertaintyCircleAllOf.md)
 - [PointUncertaintyEllipse](docs/PointUncertaintyEllipse.md)
 - [PointUncertaintyEllipseAllOf](docs/PointUncertaintyEllipseAllOf.md)
 - [Polygon](docs/Polygon.md)
 - [PolygonAllOf](docs/PolygonAllOf.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RelativeCartesianLocation](docs/RelativeCartesianLocation.md)
 - [RouteSelectionParameterSet](docs/RouteSelectionParameterSet.md)
 - [ServiceParameterData](docs/ServiceParameterData.md)
 - [ServiceParameterDataPatch](docs/ServiceParameterDataPatch.md)
 - [Snssai](docs/Snssai.md)
 - [SupportedGADShapes](docs/SupportedGADShapes.md)
 - [Tai](docs/Tai.md)
 - [TrafficDescriptorComponents](docs/TrafficDescriptorComponents.md)
 - [UncertaintyEllipse](docs/UncertaintyEllipse.md)
 - [UncertaintyEllipsoid](docs/UncertaintyEllipsoid.md)
 - [UrspRuleRequest](docs/UrspRuleRequest.md)
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


