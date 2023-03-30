# Go API client for openapi_Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.  
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
import openapi_Nnwdaf_EventsSubscription "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Nnwdaf_EventsSubscription"
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
ctx := context.WithValue(context.Background(), openapi_Nnwdaf_EventsSubscription.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Nnwdaf_EventsSubscription.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Nnwdaf_EventsSubscription.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Nnwdaf_EventsSubscription.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nnwdaf-eventssubscription/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualNWDAFEventSubscriptionTransferDocumentApi* | [**DeleteNWDAFEventSubscriptionTransfer**](docs/IndividualNWDAFEventSubscriptionTransferDocumentApi.md#deletenwdafeventsubscriptiontransfer) | **Delete** /transfers/{transferId} | Delete an existing Individual NWDAF Event Subscription Transfer
*IndividualNWDAFEventSubscriptionTransferDocumentApi* | [**UpdateNWDAFEventSubscriptionTransfer**](docs/IndividualNWDAFEventSubscriptionTransferDocumentApi.md#updatenwdafeventsubscriptiontransfer) | **Put** /transfers/{transferId} | Update an existing Individual NWDAF Event Subscription Transfer
*IndividualNWDAFEventsSubscriptionDocumentApi* | [**DeleteNWDAFEventsSubscription**](docs/IndividualNWDAFEventsSubscriptionDocumentApi.md#deletenwdafeventssubscription) | **Delete** /subscriptions/{subscriptionId} | Delete an existing Individual NWDAF Events Subscription
*IndividualNWDAFEventsSubscriptionDocumentApi* | [**UpdateNWDAFEventsSubscription**](docs/IndividualNWDAFEventsSubscriptionDocumentApi.md#updatenwdafeventssubscription) | **Put** /subscriptions/{subscriptionId} | Update an existing Individual NWDAF Events Subscription
*NWDAFEventSubscriptionTransfersCollectionApi* | [**CreateNWDAFEventSubscriptionTransfer**](docs/NWDAFEventSubscriptionTransfersCollectionApi.md#createnwdafeventsubscriptiontransfer) | **Post** /transfers | Provide information about requested analytics subscriptions transfer and potentially create a new Individual NWDAF Event Subscription Transfer resource.
*NWDAFEventsSubscriptionsCollectionApi* | [**CreateNWDAFEventsSubscription**](docs/NWDAFEventsSubscriptionsCollectionApi.md#createnwdafeventssubscription) | **Post** /subscriptions | Create a new Individual NWDAF Events Subscription


## Documentation For Models

 - [AbnormalBehaviour](docs/AbnormalBehaviour.md)
 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [Accuracy](docs/Accuracy.md)
 - [AdditionalMeasurement](docs/AdditionalMeasurement.md)
 - [AddrFqdn](docs/AddrFqdn.md)
 - [AddressList](docs/AddressList.md)
 - [AnalyticsContextIdentifier](docs/AnalyticsContextIdentifier.md)
 - [AnalyticsMetadata](docs/AnalyticsMetadata.md)
 - [AnalyticsMetadataIndication](docs/AnalyticsMetadataIndication.md)
 - [AnalyticsMetadataInfo](docs/AnalyticsMetadataInfo.md)
 - [AnalyticsSubscriptionsTransfer](docs/AnalyticsSubscriptionsTransfer.md)
 - [AnalyticsSubset](docs/AnalyticsSubset.md)
 - [AppListForUeComm](docs/AppListForUeComm.md)
 - [ApplicationVolume](docs/ApplicationVolume.md)
 - [BatteryIndication](docs/BatteryIndication.md)
 - [BwRequirement](docs/BwRequirement.md)
 - [CellGlobalId](docs/CellGlobalId.md)
 - [CircumstanceDescription](docs/CircumstanceDescription.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [ClassCriterion](docs/ClassCriterion.md)
 - [CongestionInfo](docs/CongestionInfo.md)
 - [CongestionType](docs/CongestionType.md)
 - [ConsumerNfInformation](docs/ConsumerNfInformation.md)
 - [DatasetStatisticalProperty](docs/DatasetStatisticalProperty.md)
 - [DispersionClass](docs/DispersionClass.md)
 - [DispersionCollection](docs/DispersionCollection.md)
 - [DispersionInfo](docs/DispersionInfo.md)
 - [DispersionOrderingCriterion](docs/DispersionOrderingCriterion.md)
 - [DispersionRequirement](docs/DispersionRequirement.md)
 - [DispersionType](docs/DispersionType.md)
 - [DnPerf](docs/DnPerf.md)
 - [DnPerfInfo](docs/DnPerfInfo.md)
 - [DnPerfOrderingCriterion](docs/DnPerfOrderingCriterion.md)
 - [DnPerformanceReq](docs/DnPerformanceReq.md)
 - [Ecgi](docs/Ecgi.md)
 - [EllipsoidArc](docs/EllipsoidArc.md)
 - [EllipsoidArcAllOf](docs/EllipsoidArcAllOf.md)
 - [EthFlowDescription](docs/EthFlowDescription.md)
 - [EutraLocation](docs/EutraLocation.md)
 - [EventNotification](docs/EventNotification.md)
 - [EventReportingRequirement](docs/EventReportingRequirement.md)
 - [EventSubscription](docs/EventSubscription.md)
 - [Exception](docs/Exception.md)
 - [ExceptionId](docs/ExceptionId.md)
 - [ExceptionTrend](docs/ExceptionTrend.md)
 - [ExpectedAnalyticsType](docs/ExpectedAnalyticsType.md)
 - [ExpectedUeBehaviourData](docs/ExpectedUeBehaviourData.md)
 - [FailureEventInfo](docs/FailureEventInfo.md)
 - [FlowDirection](docs/FlowDirection.md)
 - [FlowInfo](docs/FlowInfo.md)
 - [GADShape](docs/GADShape.md)
 - [GNbId](docs/GNbId.md)
 - [GeographicArea](docs/GeographicArea.md)
 - [GeographicalCoordinates](docs/GeographicalCoordinates.md)
 - [GeraLocation](docs/GeraLocation.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [HfcNodeId](docs/HfcNodeId.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [IpEthFlowDescription](docs/IpEthFlowDescription.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [LineType](docs/LineType.md)
 - [Local2dPointUncertaintyEllipse](docs/Local2dPointUncertaintyEllipse.md)
 - [Local2dPointUncertaintyEllipseAllOf](docs/Local2dPointUncertaintyEllipseAllOf.md)
 - [Local3dPointUncertaintyEllipsoid](docs/Local3dPointUncertaintyEllipsoid.md)
 - [Local3dPointUncertaintyEllipsoidAllOf](docs/Local3dPointUncertaintyEllipsoidAllOf.md)
 - [LocalOrigin](docs/LocalOrigin.md)
 - [LocationArea](docs/LocationArea.md)
 - [LocationAreaId](docs/LocationAreaId.md)
 - [LocationInfo](docs/LocationInfo.md)
 - [MLModelAddr](docs/MLModelAddr.md)
 - [MLModelInfo](docs/MLModelInfo.md)
 - [MatchingDirection](docs/MatchingDirection.md)
 - [ModelInfo](docs/ModelInfo.md)
 - [N3gaLocation](docs/N3gaLocation.md)
 - [NFType](docs/NFType.md)
 - [Ncgi](docs/Ncgi.md)
 - [NetworkAreaInfo](docs/NetworkAreaInfo.md)
 - [NetworkAreaInfo1](docs/NetworkAreaInfo1.md)
 - [NetworkPerfInfo](docs/NetworkPerfInfo.md)
 - [NetworkPerfRequirement](docs/NetworkPerfRequirement.md)
 - [NetworkPerfType](docs/NetworkPerfType.md)
 - [NfLoadLevelInformation](docs/NfLoadLevelInformation.md)
 - [NfStatus](docs/NfStatus.md)
 - [NnwdafEventsSubscription](docs/NnwdafEventsSubscription.md)
 - [NnwdafEventsSubscriptionNotification](docs/NnwdafEventsSubscriptionNotification.md)
 - [NotificationFlag](docs/NotificationFlag.md)
 - [NotificationMethod](docs/NotificationMethod.md)
 - [NotificationMethod1](docs/NotificationMethod1.md)
 - [NrLocation](docs/NrLocation.md)
 - [NsiIdInfo](docs/NsiIdInfo.md)
 - [NsiLoadLevelInfo](docs/NsiLoadLevelInfo.md)
 - [NumberAverage](docs/NumberAverage.md)
 - [NwdafEvent](docs/NwdafEvent.md)
 - [NwdafFailureCode](docs/NwdafFailureCode.md)
 - [ObservedRedundantTransExp](docs/ObservedRedundantTransExp.md)
 - [OutputStrategy](docs/OutputStrategy.md)
 - [PartitioningCriteria](docs/PartitioningCriteria.md)
 - [PerfData](docs/PerfData.md)
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
 - [PrevSubInfo](docs/PrevSubInfo.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [QosRequirement](docs/QosRequirement.md)
 - [QosResourceType](docs/QosResourceType.md)
 - [QosSustainabilityInfo](docs/QosSustainabilityInfo.md)
 - [RankingCriterion](docs/RankingCriterion.md)
 - [RatFreqInformation](docs/RatFreqInformation.md)
 - [RatType](docs/RatType.md)
 - [RedTransExpOrderingCriterion](docs/RedTransExpOrderingCriterion.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [RedundantTransmissionExpInfo](docs/RedundantTransmissionExpInfo.md)
 - [RedundantTransmissionExpPerTS](docs/RedundantTransmissionExpPerTS.md)
 - [RedundantTransmissionExpReq](docs/RedundantTransmissionExpReq.md)
 - [RelativeCartesianLocation](docs/RelativeCartesianLocation.md)
 - [ReportingInformation](docs/ReportingInformation.md)
 - [ResourceUsage](docs/ResourceUsage.md)
 - [RetainabilityThreshold](docs/RetainabilityThreshold.md)
 - [RoutingAreaId](docs/RoutingAreaId.md)
 - [ScheduledCommunicationTime](docs/ScheduledCommunicationTime.md)
 - [ScheduledCommunicationTime1](docs/ScheduledCommunicationTime1.md)
 - [ScheduledCommunicationType](docs/ScheduledCommunicationType.md)
 - [ServiceAreaId](docs/ServiceAreaId.md)
 - [ServiceExperienceInfo](docs/ServiceExperienceInfo.md)
 - [ServiceExperienceType](docs/ServiceExperienceType.md)
 - [SessInactTimerForUeComm](docs/SessInactTimerForUeComm.md)
 - [SliceLoadLevelInformation](docs/SliceLoadLevelInformation.md)
 - [SmcceInfo](docs/SmcceInfo.md)
 - [SmcceUeList](docs/SmcceUeList.md)
 - [Snssai](docs/Snssai.md)
 - [StationaryIndication](docs/StationaryIndication.md)
 - [SubscriptionTransferInfo](docs/SubscriptionTransferInfo.md)
 - [SupportedGADShapes](docs/SupportedGADShapes.md)
 - [SvcExperience](docs/SvcExperience.md)
 - [Tai](docs/Tai.md)
 - [TargetUeInformation](docs/TargetUeInformation.md)
 - [TermCause](docs/TermCause.md)
 - [ThresholdLevel](docs/ThresholdLevel.md)
 - [TimeUnit](docs/TimeUnit.md)
 - [TimeWindow](docs/TimeWindow.md)
 - [TnapId](docs/TnapId.md)
 - [TopApplication](docs/TopApplication.md)
 - [TrafficCharacterization](docs/TrafficCharacterization.md)
 - [TrafficInformation](docs/TrafficInformation.md)
 - [TrafficProfile](docs/TrafficProfile.md)
 - [TransferRequestType](docs/TransferRequestType.md)
 - [TransportProtocol](docs/TransportProtocol.md)
 - [TwapId](docs/TwapId.md)
 - [UeAnalyticsContextDescriptor](docs/UeAnalyticsContextDescriptor.md)
 - [UeCommunication](docs/UeCommunication.md)
 - [UeMobility](docs/UeMobility.md)
 - [UmtTime](docs/UmtTime.md)
 - [UncertaintyEllipse](docs/UncertaintyEllipse.md)
 - [UncertaintyEllipsoid](docs/UncertaintyEllipsoid.md)
 - [UpfInformation](docs/UpfInformation.md)
 - [UserDataCongestionInfo](docs/UserDataCongestionInfo.md)
 - [UserLocation](docs/UserLocation.md)
 - [UtraLocation](docs/UtraLocation.md)
 - [WlanOrderingCriterion](docs/WlanOrderingCriterion.md)
 - [WlanPerSsIdPerformanceInfo](docs/WlanPerSsIdPerformanceInfo.md)
 - [WlanPerTsPerformanceInfo](docs/WlanPerTsPerformanceInfo.md)
 - [WlanPerformanceInfo](docs/WlanPerformanceInfo.md)
 - [WlanPerformanceReq](docs/WlanPerformanceReq.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nnwdaf-eventssubscription**: Access to the Nnwdaf_EventsSubscription API

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


