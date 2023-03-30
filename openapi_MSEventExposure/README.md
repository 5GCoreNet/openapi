# Go API client for openapi_MSEventExposure

API for Media Streaming Event Exposure.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.1
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
import openapi_MSEventExposure "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_MSEventExposure"
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
ctx := context.WithValue(context.Background(), openapi_MSEventExposure.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_MSEventExposure.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_MSEventExposure.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_MSEventExposure.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/3gpp-ms-event-exposure/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualMediaStreamingEventExposureSubscriptionDocumentApi* | [**DeleteIndivMSEventExposureSubsc**](docs/IndividualMediaStreamingEventExposureSubscriptionDocumentApi.md#deleteindivmseventexposuresubsc) | **Delete** /subscriptions/{subscriptionId} | Request the deletion of an existing Individual Media Streaming Event Exposure Subscription resource.
*IndividualMediaStreamingEventExposureSubscriptionDocumentApi* | [**RetrieveIndivMSEventExposureSubsc**](docs/IndividualMediaStreamingEventExposureSubscriptionDocumentApi.md#retrieveindivmseventexposuresubsc) | **Get** /subscriptions/{subscriptionId} | Retrieve an existing Individual Media Streaming Event Exposure Subscription resource.
*IndividualMediaStreamingEventExposureSubscriptionDocumentApi* | [**UpdateIndivMSEventExposureSubsc**](docs/IndividualMediaStreamingEventExposureSubscriptionDocumentApi.md#updateindivmseventexposuresubsc) | **Put** /subscriptions/{subscriptionId} | Request the update of an existing Individual Media Streaming Event Exposure Subscription resource.
*MediaStreamingEventExposureSubscriptionsCollectionApi* | [**CreateMSEventExposureSubsc**](docs/MediaStreamingEventExposureSubscriptionsCollectionApi.md#createmseventexposuresubsc) | **Post** /subscriptions | Request the creation of a new Individual Media Streaming Event Exposure Subscription resource.
*MediaStreamingEventExposureSubscriptionsCollectionApi* | [**RetrieveMSEventExposureSubscs**](docs/MediaStreamingEventExposureSubscriptionsCollectionApi.md#retrievemseventexposuresubscs) | **Get** /subscriptions | Retrieve all the active Media Streaming Event Exposure Subscription resources managed by the NEF.


## Documentation For Models

 - [AddrFqdn](docs/AddrFqdn.md)
 - [AfEvent](docs/AfEvent.md)
 - [AfEventExposureNotif](docs/AfEventExposureNotif.md)
 - [AfEventExposureSubsc](docs/AfEventExposureSubsc.md)
 - [AfEventNotification](docs/AfEventNotification.md)
 - [BaseRecord](docs/BaseRecord.md)
 - [CacheStatus](docs/CacheStatus.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [CollectiveBehaviourFilter](docs/CollectiveBehaviourFilter.md)
 - [CollectiveBehaviourFilterType](docs/CollectiveBehaviourFilterType.md)
 - [CollectiveBehaviourInfo](docs/CollectiveBehaviourInfo.md)
 - [CommunicationCollection](docs/CommunicationCollection.md)
 - [DispersionCollection](docs/DispersionCollection.md)
 - [DynamicPolicy](docs/DynamicPolicy.md)
 - [Ecgi](docs/Ecgi.md)
 - [EllipsoidArc](docs/EllipsoidArc.md)
 - [EllipsoidArcAllOf](docs/EllipsoidArcAllOf.md)
 - [EndpointAddress](docs/EndpointAddress.md)
 - [EthFlowDescription](docs/EthFlowDescription.md)
 - [EventFilter](docs/EventFilter.md)
 - [EventsSubs](docs/EventsSubs.md)
 - [Exception](docs/Exception.md)
 - [ExceptionId](docs/ExceptionId.md)
 - [ExceptionInfo](docs/ExceptionInfo.md)
 - [ExceptionTrend](docs/ExceptionTrend.md)
 - [FlowDirection](docs/FlowDirection.md)
 - [FlowInfo](docs/FlowInfo.md)
 - [GADShape](docs/GADShape.md)
 - [GNbId](docs/GNbId.md)
 - [GeographicArea](docs/GeographicArea.md)
 - [GeographicalCoordinates](docs/GeographicalCoordinates.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [IpPacketFilterSet](docs/IpPacketFilterSet.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [Local2dPointUncertaintyEllipse](docs/Local2dPointUncertaintyEllipse.md)
 - [Local2dPointUncertaintyEllipseAllOf](docs/Local2dPointUncertaintyEllipseAllOf.md)
 - [Local3dPointUncertaintyEllipsoid](docs/Local3dPointUncertaintyEllipsoid.md)
 - [Local3dPointUncertaintyEllipsoidAllOf](docs/Local3dPointUncertaintyEllipsoidAllOf.md)
 - [LocalOrigin](docs/LocalOrigin.md)
 - [LocationArea5G](docs/LocationArea5G.md)
 - [M5QoSSpecification](docs/M5QoSSpecification.md)
 - [MSAccessActivityCollection](docs/MSAccessActivityCollection.md)
 - [MediaStreamingAccessRecord](docs/MediaStreamingAccessRecord.md)
 - [MediaStreamingAccessRecordAllOf](docs/MediaStreamingAccessRecordAllOf.md)
 - [MediaStreamingAccessRecordAllOfConnectionMetrics](docs/MediaStreamingAccessRecordAllOfConnectionMetrics.md)
 - [MediaStreamingAccessRecordAllOfRequestMessage](docs/MediaStreamingAccessRecordAllOfRequestMessage.md)
 - [MediaStreamingAccessRecordAllOfResponseMessage](docs/MediaStreamingAccessRecordAllOfResponseMessage.md)
 - [MsConsumptionCollection](docs/MsConsumptionCollection.md)
 - [MsDynPolicyInvocationCollection](docs/MsDynPolicyInvocationCollection.md)
 - [MsNetAssInvocationCollection](docs/MsNetAssInvocationCollection.md)
 - [MsQoeMetricsCollection](docs/MsQoeMetricsCollection.md)
 - [Ncgi](docs/Ncgi.md)
 - [NetworkAreaInfo](docs/NetworkAreaInfo.md)
 - [NetworkAssistanceSession](docs/NetworkAssistanceSession.md)
 - [NotificationFlag](docs/NotificationFlag.md)
 - [NotificationMethod](docs/NotificationMethod.md)
 - [PartitioningCriteria](docs/PartitioningCriteria.md)
 - [PerUeAttribute](docs/PerUeAttribute.md)
 - [PerformanceData](docs/PerformanceData.md)
 - [PerformanceDataCollection](docs/PerformanceDataCollection.md)
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
 - [ReportingInformation](docs/ReportingInformation.md)
 - [ServiceDataFlowDescription](docs/ServiceDataFlowDescription.md)
 - [ServiceExperienceInfoPerApp](docs/ServiceExperienceInfoPerApp.md)
 - [ServiceExperienceInfoPerFlow](docs/ServiceExperienceInfoPerFlow.md)
 - [SupportedGADShapes](docs/SupportedGADShapes.md)
 - [SvcExperience](docs/SvcExperience.md)
 - [Tai](docs/Tai.md)
 - [TimeWindow](docs/TimeWindow.md)
 - [UeCommunicationCollection](docs/UeCommunicationCollection.md)
 - [UeMobilityCollection](docs/UeMobilityCollection.md)
 - [UeTrajectoryCollection](docs/UeTrajectoryCollection.md)
 - [UncertaintyEllipse](docs/UncertaintyEllipse.md)
 - [UncertaintyEllipsoid](docs/UncertaintyEllipsoid.md)
 - [UsageThreshold](docs/UsageThreshold.md)
 - [UserDataCongestionCollection](docs/UserDataCongestionCollection.md)


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


