# Go API client for openapi_Namf_Location

AMF Location Service.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.3.0-alpha.1
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
import openapi_Namf_Location "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Namf_Location"
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
ctx := context.WithValue(context.Background(), openapi_Namf_Location.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Namf_Location.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Namf_Location.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Namf_Location.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/namf-loc/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualUEContextDocumentApi* | [**CancelLocation**](docs/IndividualUEContextDocumentApi.md#cancellocation) | **Post** /{ueContextId}/cancel-pos-info | Namf_Location CancelLocation service operation
*IndividualUEContextDocumentApi* | [**ProvideLocationInfo**](docs/IndividualUEContextDocumentApi.md#providelocationinfo) | **Post** /{ueContextId}/provide-loc-info | Namf_Location ProvideLocationInfo service Operation
*IndividualUEContextDocumentApi* | [**ProvidePositioningInfo**](docs/IndividualUEContextDocumentApi.md#providepositioninginfo) | **Post** /{ueContextId}/provide-pos-info | Namf_Location ProvidePositioningInfo service Operation


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AccuracyFulfilmentIndicator](docs/AccuracyFulfilmentIndicator.md)
 - [AreaEventInfo](docs/AreaEventInfo.md)
 - [CancelPosInfo](docs/CancelPosInfo.md)
 - [CellGlobalId](docs/CellGlobalId.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [Ecgi](docs/Ecgi.md)
 - [EllipsoidArc](docs/EllipsoidArc.md)
 - [EllipsoidArcAllOf](docs/EllipsoidArcAllOf.md)
 - [EutraLocation](docs/EutraLocation.md)
 - [ExternalClientType](docs/ExternalClientType.md)
 - [GADShape](docs/GADShape.md)
 - [GNbId](docs/GNbId.md)
 - [GeographicArea](docs/GeographicArea.md)
 - [GeographicalCoordinates](docs/GeographicalCoordinates.md)
 - [GeraLocation](docs/GeraLocation.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [GnssId](docs/GnssId.md)
 - [GnssPositioningMethodAndUsage](docs/GnssPositioningMethodAndUsage.md)
 - [Guami](docs/Guami.md)
 - [HfcNodeId](docs/HfcNodeId.md)
 - [HorizontalVelocity](docs/HorizontalVelocity.md)
 - [HorizontalVelocityWithUncertainty](docs/HorizontalVelocityWithUncertainty.md)
 - [HorizontalWithVerticalVelocity](docs/HorizontalWithVerticalVelocity.md)
 - [HorizontalWithVerticalVelocityAndUncertainty](docs/HorizontalWithVerticalVelocityAndUncertainty.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [LcsPriority](docs/LcsPriority.md)
 - [LcsQosClass](docs/LcsQosClass.md)
 - [LcsServiceAuth](docs/LcsServiceAuth.md)
 - [LdrType](docs/LdrType.md)
 - [LineType](docs/LineType.md)
 - [Local2dPointUncertaintyEllipse](docs/Local2dPointUncertaintyEllipse.md)
 - [Local2dPointUncertaintyEllipseAllOf](docs/Local2dPointUncertaintyEllipseAllOf.md)
 - [Local3dPointUncertaintyEllipsoid](docs/Local3dPointUncertaintyEllipsoid.md)
 - [Local3dPointUncertaintyEllipsoidAllOf](docs/Local3dPointUncertaintyEllipsoidAllOf.md)
 - [LocalArea](docs/LocalArea.md)
 - [LocalOrigin](docs/LocalOrigin.md)
 - [LocationAreaId](docs/LocationAreaId.md)
 - [LocationEvent](docs/LocationEvent.md)
 - [LocationPrivacyVerResult](docs/LocationPrivacyVerResult.md)
 - [LocationQoS](docs/LocationQoS.md)
 - [LocationType](docs/LocationType.md)
 - [MinorLocationQoS](docs/MinorLocationQoS.md)
 - [MotionEventInfo](docs/MotionEventInfo.md)
 - [N3gaLocation](docs/N3gaLocation.md)
 - [NFType](docs/NFType.md)
 - [Ncgi](docs/Ncgi.md)
 - [NotifiedPosInfo](docs/NotifiedPosInfo.md)
 - [NrLocation](docs/NrLocation.md)
 - [OccurrenceInfo](docs/OccurrenceInfo.md)
 - [PeriodicEventInfo](docs/PeriodicEventInfo.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
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
 - [PositioningMethod](docs/PositioningMethod.md)
 - [PositioningMethodAndUsage](docs/PositioningMethodAndUsage.md)
 - [PositioningMode](docs/PositioningMode.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [ProvideLocInfo](docs/ProvideLocInfo.md)
 - [ProvidePosInfo](docs/ProvidePosInfo.md)
 - [RatType](docs/RatType.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [RelativeCartesianLocation](docs/RelativeCartesianLocation.md)
 - [ReportingArea](docs/ReportingArea.md)
 - [ReportingAreaType](docs/ReportingAreaType.md)
 - [RequestLocInfo](docs/RequestLocInfo.md)
 - [RequestPosInfo](docs/RequestPosInfo.md)
 - [ResponseTime](docs/ResponseTime.md)
 - [RoutingAreaId](docs/RoutingAreaId.md)
 - [ServiceAreaId](docs/ServiceAreaId.md)
 - [Snssai](docs/Snssai.md)
 - [SupportedGADShapes](docs/SupportedGADShapes.md)
 - [Tai](docs/Tai.md)
 - [TerminationCause](docs/TerminationCause.md)
 - [TnapId](docs/TnapId.md)
 - [TransportProtocol](docs/TransportProtocol.md)
 - [TwapId](docs/TwapId.md)
 - [UePrivacyRequirements](docs/UePrivacyRequirements.md)
 - [UncertaintyEllipse](docs/UncertaintyEllipse.md)
 - [UncertaintyEllipsoid](docs/UncertaintyEllipsoid.md)
 - [Usage](docs/Usage.md)
 - [UserLocation](docs/UserLocation.md)
 - [UtraLocation](docs/UtraLocation.md)
 - [VelocityEstimate](docs/VelocityEstimate.md)
 - [VelocityRequested](docs/VelocityRequested.md)
 - [VerticalDirection](docs/VerticalDirection.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **namf-loc**: Access to the Namf_Location API

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



