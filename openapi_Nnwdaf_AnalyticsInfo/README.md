# Go API client for openapi_Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.  
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
import openapi_Nnwdaf_AnalyticsInfo "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Nnwdaf_AnalyticsInfo"
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
ctx := context.WithValue(context.Background(), openapi_Nnwdaf_AnalyticsInfo.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Nnwdaf_AnalyticsInfo.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Nnwdaf_AnalyticsInfo.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Nnwdaf_AnalyticsInfo.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nnwdaf-analyticsinfo/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*NWDAFAnalyticsDocumentApi* | [**GetNWDAFAnalytics**](docs/NWDAFAnalyticsDocumentApi.md#getnwdafanalytics) | **Get** /analytics | Read a NWDAF Analytics
*NWDAFContextDocumentApi* | [**GetNwdafContext**](docs/NWDAFContextDocumentApi.md#getnwdafcontext) | **Get** /context | Get context information related to analytics subscriptions.


## Documentation For Models

 - [AanfInfo](docs/AanfInfo.md)
 - [AbnormalBehaviour](docs/AbnormalBehaviour.md)
 - [AccessStateTransitionType](docs/AccessStateTransitionType.md)
 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AccessType](docs/AccessType.md)
 - [Accuracy](docs/Accuracy.md)
 - [AdditionInfoAnalyticsInfoRequest](docs/AdditionInfoAnalyticsInfoRequest.md)
 - [AdditionalMeasurement](docs/AdditionalMeasurement.md)
 - [AddrFqdn](docs/AddrFqdn.md)
 - [AddressList](docs/AddressList.md)
 - [AdrfDataType](docs/AdrfDataType.md)
 - [AfEvent](docs/AfEvent.md)
 - [AfEventExposureData](docs/AfEventExposureData.md)
 - [AfEventExposureNotif](docs/AfEventExposureNotif.md)
 - [AfEventExposureSubsc](docs/AfEventExposureSubsc.md)
 - [AfEventNotification](docs/AfEventNotification.md)
 - [AmfCond](docs/AmfCond.md)
 - [AmfEvent](docs/AmfEvent.md)
 - [AmfEventArea](docs/AmfEventArea.md)
 - [AmfEventMode](docs/AmfEventMode.md)
 - [AmfEventNotification](docs/AmfEventNotification.md)
 - [AmfEventReport](docs/AmfEventReport.md)
 - [AmfEventState](docs/AmfEventState.md)
 - [AmfEventSubsSyncInfo](docs/AmfEventSubsSyncInfo.md)
 - [AmfEventSubscription](docs/AmfEventSubscription.md)
 - [AmfEventSubscriptionInfo](docs/AmfEventSubscriptionInfo.md)
 - [AmfEventTrigger](docs/AmfEventTrigger.md)
 - [AmfEventType](docs/AmfEventType.md)
 - [AmfInfo](docs/AmfInfo.md)
 - [AnNodeType](docs/AnNodeType.md)
 - [AnalyticsContextIdentifier](docs/AnalyticsContextIdentifier.md)
 - [AnalyticsData](docs/AnalyticsData.md)
 - [AnalyticsMetadata](docs/AnalyticsMetadata.md)
 - [AnalyticsMetadataIndication](docs/AnalyticsMetadataIndication.md)
 - [AnalyticsMetadataInfo](docs/AnalyticsMetadataInfo.md)
 - [AnalyticsSubset](docs/AnalyticsSubset.md)
 - [AppListForUeComm](docs/AppListForUeComm.md)
 - [ApplicationVolume](docs/ApplicationVolume.md)
 - [AppliedSmccType](docs/AppliedSmccType.md)
 - [AssociationType](docs/AssociationType.md)
 - [AtsssCapability](docs/AtsssCapability.md)
 - [AusfInfo](docs/AusfInfo.md)
 - [BaseRecord](docs/BaseRecord.md)
 - [BatteryIndication](docs/BatteryIndication.md)
 - [BsfInfo](docs/BsfInfo.md)
 - [BwRequirement](docs/BwRequirement.md)
 - [CacheStatus](docs/CacheStatus.md)
 - [CellGlobalId](docs/CellGlobalId.md)
 - [ChangeItem](docs/ChangeItem.md)
 - [ChangeOfSupiPeiAssociationReport](docs/ChangeOfSupiPeiAssociationReport.md)
 - [ChangeType](docs/ChangeType.md)
 - [ChfInfo](docs/ChfInfo.md)
 - [CircumstanceDescription](docs/CircumstanceDescription.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [ClassCriterion](docs/ClassCriterion.md)
 - [CmInfo](docs/CmInfo.md)
 - [CmInfoReport](docs/CmInfoReport.md)
 - [CmState](docs/CmState.md)
 - [CnType](docs/CnType.md)
 - [CnTypeChangeReport](docs/CnTypeChangeReport.md)
 - [CollectiveBehaviourFilter](docs/CollectiveBehaviourFilter.md)
 - [CollectiveBehaviourFilterType](docs/CollectiveBehaviourFilterType.md)
 - [CollectiveBehaviourInfo](docs/CollectiveBehaviourInfo.md)
 - [CollocatedNfInstance](docs/CollocatedNfInstance.md)
 - [CollocatedNfType](docs/CollocatedNfType.md)
 - [CommunicationCollection](docs/CommunicationCollection.md)
 - [CommunicationFailure](docs/CommunicationFailure.md)
 - [ConditionEventType](docs/ConditionEventType.md)
 - [CongestionInfo](docs/CongestionInfo.md)
 - [CongestionType](docs/CongestionType.md)
 - [ConsumerNfInformation](docs/ConsumerNfInformation.md)
 - [ContextData](docs/ContextData.md)
 - [ContextElement](docs/ContextElement.md)
 - [ContextIdList](docs/ContextIdList.md)
 - [ContextInfo](docs/ContextInfo.md)
 - [ContextType](docs/ContextType.md)
 - [DataNotification](docs/DataNotification.md)
 - [DataSetId](docs/DataSetId.md)
 - [DataSubscription](docs/DataSubscription.md)
 - [DatalinkReportingConfiguration](docs/DatalinkReportingConfiguration.md)
 - [DatasetStatisticalProperty](docs/DatasetStatisticalProperty.md)
 - [DccfCond](docs/DccfCond.md)
 - [DccfInfo](docs/DccfInfo.md)
 - [DddTrafficDescriptor](docs/DddTrafficDescriptor.md)
 - [DefSubServiceInfo](docs/DefSubServiceInfo.md)
 - [DefaultNotificationSubscription](docs/DefaultNotificationSubscription.md)
 - [DispersionArea](docs/DispersionArea.md)
 - [DispersionClass](docs/DispersionClass.md)
 - [DispersionCollection](docs/DispersionCollection.md)
 - [DispersionCollection1](docs/DispersionCollection1.md)
 - [DispersionInfo](docs/DispersionInfo.md)
 - [DispersionOrderingCriterion](docs/DispersionOrderingCriterion.md)
 - [DispersionRequirement](docs/DispersionRequirement.md)
 - [DispersionType](docs/DispersionType.md)
 - [DlDataDeliveryStatus](docs/DlDataDeliveryStatus.md)
 - [DnPerf](docs/DnPerf.md)
 - [DnPerfInfo](docs/DnPerfInfo.md)
 - [DnPerfOrderingCriterion](docs/DnPerfOrderingCriterion.md)
 - [DnPerformanceReq](docs/DnPerformanceReq.md)
 - [DnaiChangeType](docs/DnaiChangeType.md)
 - [DnnEasdfInfoItem](docs/DnnEasdfInfoItem.md)
 - [DnnInfoItem](docs/DnnInfoItem.md)
 - [DnnMbSmfInfoItem](docs/DnnMbSmfInfoItem.md)
 - [DnnSmfInfoItem](docs/DnnSmfInfoItem.md)
 - [DnnSmfInfoItemDnaiListInner](docs/DnnSmfInfoItemDnaiListInner.md)
 - [DnnSmfInfoItemDnn](docs/DnnSmfInfoItemDnn.md)
 - [DnnTsctsfInfoItem](docs/DnnTsctsfInfoItem.md)
 - [DnnUpfInfoItem](docs/DnnUpfInfoItem.md)
 - [DynamicPolicy](docs/DynamicPolicy.md)
 - [EasdfInfo](docs/EasdfInfo.md)
 - [Ecgi](docs/Ecgi.md)
 - [EeSubscription](docs/EeSubscription.md)
 - [EllipsoidArc](docs/EllipsoidArc.md)
 - [EllipsoidArcAllOf](docs/EllipsoidArcAllOf.md)
 - [EndpointAddress](docs/EndpointAddress.md)
 - [EthFlowDescription](docs/EthFlowDescription.md)
 - [EutraLocation](docs/EutraLocation.md)
 - [EventFilter](docs/EventFilter.md)
 - [EventFilter1](docs/EventFilter1.md)
 - [EventId](docs/EventId.md)
 - [EventNotification](docs/EventNotification.md)
 - [EventNotification1](docs/EventNotification1.md)
 - [EventReportMode](docs/EventReportMode.md)
 - [EventReportingRequirement](docs/EventReportingRequirement.md)
 - [EventSubscription](docs/EventSubscription.md)
 - [EventSubscription1](docs/EventSubscription1.md)
 - [EventType](docs/EventType.md)
 - [EventsSubs](docs/EventsSubs.md)
 - [Exception](docs/Exception.md)
 - [ExceptionId](docs/ExceptionId.md)
 - [ExceptionInfo](docs/ExceptionInfo.md)
 - [ExceptionTrend](docs/ExceptionTrend.md)
 - [ExpectedAnalyticsType](docs/ExpectedAnalyticsType.md)
 - [ExpectedUeBehaviourData](docs/ExpectedUeBehaviourData.md)
 - [ExtSnssai](docs/ExtSnssai.md)
 - [ExternalClientType](docs/ExternalClientType.md)
 - [FailureEventInfo](docs/FailureEventInfo.md)
 - [FlowDirection](docs/FlowDirection.md)
 - [FlowInfo](docs/FlowInfo.md)
 - [GADShape](docs/GADShape.md)
 - [GNbId](docs/GNbId.md)
 - [GeographicArea](docs/GeographicArea.md)
 - [GeographicalCoordinates](docs/GeographicalCoordinates.md)
 - [GeraLocation](docs/GeraLocation.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [GmlcInfo](docs/GmlcInfo.md)
 - [Guami](docs/Guami.md)
 - [GuamiListCond](docs/GuamiListCond.md)
 - [HfcNodeId](docs/HfcNodeId.md)
 - [HistoricalData](docs/HistoricalData.md)
 - [HssInfo](docs/HssInfo.md)
 - [IdentityRange](docs/IdentityRange.md)
 - [IdleStatusIndication](docs/IdleStatusIndication.md)
 - [ImsiRange](docs/ImsiRange.md)
 - [InterfaceUpfInfoItem](docs/InterfaceUpfInfoItem.md)
 - [InternalGroupIdRange](docs/InternalGroupIdRange.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [IpEndPoint](docs/IpEndPoint.md)
 - [IpEthFlowDescription](docs/IpEthFlowDescription.md)
 - [IpIndex](docs/IpIndex.md)
 - [IpPacketFilterSet](docs/IpPacketFilterSet.md)
 - [IpReachability](docs/IpReachability.md)
 - [Ipv4AddressRange](docs/Ipv4AddressRange.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [Ipv6PrefixRange](docs/Ipv6PrefixRange.md)
 - [IwmscInfo](docs/IwmscInfo.md)
 - [LadnInfo](docs/LadnInfo.md)
 - [LineType](docs/LineType.md)
 - [LmfInfo](docs/LmfInfo.md)
 - [Local2dPointUncertaintyEllipse](docs/Local2dPointUncertaintyEllipse.md)
 - [Local2dPointUncertaintyEllipseAllOf](docs/Local2dPointUncertaintyEllipseAllOf.md)
 - [Local3dPointUncertaintyEllipsoid](docs/Local3dPointUncertaintyEllipsoid.md)
 - [Local3dPointUncertaintyEllipsoidAllOf](docs/Local3dPointUncertaintyEllipsoidAllOf.md)
 - [LocalOrigin](docs/LocalOrigin.md)
 - [LocalityDescription](docs/LocalityDescription.md)
 - [LocalityDescriptionItem](docs/LocalityDescriptionItem.md)
 - [LocalityType](docs/LocalityType.md)
 - [LocationAccuracy](docs/LocationAccuracy.md)
 - [LocationArea](docs/LocationArea.md)
 - [LocationArea5G](docs/LocationArea5G.md)
 - [LocationAreaId](docs/LocationAreaId.md)
 - [LocationFilter](docs/LocationFilter.md)
 - [LocationInfo](docs/LocationInfo.md)
 - [LocationReport](docs/LocationReport.md)
 - [LocationReportingConfiguration](docs/LocationReportingConfiguration.md)
 - [LossConnectivityCfg](docs/LossConnectivityCfg.md)
 - [LossConnectivityReport](docs/LossConnectivityReport.md)
 - [LossOfConnectivityReason](docs/LossOfConnectivityReason.md)
 - [M5QoSSpecification](docs/M5QoSSpecification.md)
 - [MLModelAddr](docs/MLModelAddr.md)
 - [MLModelInfo](docs/MLModelInfo.md)
 - [MSAccessActivityCollection](docs/MSAccessActivityCollection.md)
 - [MatchingDirection](docs/MatchingDirection.md)
 - [MbSmfInfo](docs/MbSmfInfo.md)
 - [MbUpfInfo](docs/MbUpfInfo.md)
 - [MbsServiceArea](docs/MbsServiceArea.md)
 - [MbsServiceAreaInfo](docs/MbsServiceAreaInfo.md)
 - [MbsSession](docs/MbsSession.md)
 - [MbsSessionId](docs/MbsSessionId.md)
 - [MediaStreamingAccessRecord](docs/MediaStreamingAccessRecord.md)
 - [MediaStreamingAccessRecordAllOf](docs/MediaStreamingAccessRecordAllOf.md)
 - [MediaStreamingAccessRecordAllOfConnectionMetrics](docs/MediaStreamingAccessRecordAllOfConnectionMetrics.md)
 - [MediaStreamingAccessRecordAllOfRequestMessage](docs/MediaStreamingAccessRecordAllOfRequestMessage.md)
 - [MediaStreamingAccessRecordAllOfResponseMessage](docs/MediaStreamingAccessRecordAllOfResponseMessage.md)
 - [MfafInfo](docs/MfafInfo.md)
 - [MlAnalyticsInfo](docs/MlAnalyticsInfo.md)
 - [MmTransactionLocationReportItem](docs/MmTransactionLocationReportItem.md)
 - [MmTransactionSliceReportItem](docs/MmTransactionSliceReportItem.md)
 - [MnpfInfo](docs/MnpfInfo.md)
 - [Model5GDdnmfInfo](docs/Model5GDdnmfInfo.md)
 - [Model5GsUserState](docs/Model5GsUserState.md)
 - [Model5GsUserStateInfo](docs/Model5GsUserStateInfo.md)
 - [ModelInfo](docs/ModelInfo.md)
 - [MonitoringConfiguration](docs/MonitoringConfiguration.md)
 - [MonitoringReport](docs/MonitoringReport.md)
 - [MsConsumptionCollection](docs/MsConsumptionCollection.md)
 - [MsDynPolicyInvocationCollection](docs/MsDynPolicyInvocationCollection.md)
 - [MsNetAssInvocationCollection](docs/MsNetAssInvocationCollection.md)
 - [MsQoeMetricsCollection](docs/MsQoeMetricsCollection.md)
 - [N1MessageClass](docs/N1MessageClass.md)
 - [N2InformationClass](docs/N2InformationClass.md)
 - [N2InterfaceAmfInfo](docs/N2InterfaceAmfInfo.md)
 - [N32Purpose](docs/N32Purpose.md)
 - [N3gaLocation](docs/N3gaLocation.md)
 - [NFProfile](docs/NFProfile.md)
 - [NFService](docs/NFService.md)
 - [NFServiceStatus](docs/NFServiceStatus.md)
 - [NFServiceVersion](docs/NFServiceVersion.md)
 - [NFType](docs/NFType.md)
 - [Ncgi](docs/Ncgi.md)
 - [NcgiTai](docs/NcgiTai.md)
 - [NefCond](docs/NefCond.md)
 - [NefEvent](docs/NefEvent.md)
 - [NefEventExposureNotif](docs/NefEventExposureNotif.md)
 - [NefEventExposureSubsc](docs/NefEventExposureSubsc.md)
 - [NefEventFilter](docs/NefEventFilter.md)
 - [NefEventNotification](docs/NefEventNotification.md)
 - [NefEventSubs](docs/NefEventSubs.md)
 - [NefInfo](docs/NefInfo.md)
 - [NetworkAreaInfo](docs/NetworkAreaInfo.md)
 - [NetworkAreaInfo1](docs/NetworkAreaInfo1.md)
 - [NetworkAssistanceSession](docs/NetworkAssistanceSession.md)
 - [NetworkNodeDiameterAddress](docs/NetworkNodeDiameterAddress.md)
 - [NetworkPerfInfo](docs/NetworkPerfInfo.md)
 - [NetworkPerfRequirement](docs/NetworkPerfRequirement.md)
 - [NetworkPerfType](docs/NetworkPerfType.md)
 - [NetworkSliceCond](docs/NetworkSliceCond.md)
 - [NfGroupCond](docs/NfGroupCond.md)
 - [NfGroupListCond](docs/NfGroupListCond.md)
 - [NfInfo](docs/NfInfo.md)
 - [NfInstanceIdCond](docs/NfInstanceIdCond.md)
 - [NfInstanceIdListCond](docs/NfInstanceIdListCond.md)
 - [NfLoadLevelInformation](docs/NfLoadLevelInformation.md)
 - [NfServiceSetCond](docs/NfServiceSetCond.md)
 - [NfSetCond](docs/NfSetCond.md)
 - [NfStatus](docs/NfStatus.md)
 - [NfTypeCond](docs/NfTypeCond.md)
 - [NgApCause](docs/NgApCause.md)
 - [NnwdafEventsSubscription](docs/NnwdafEventsSubscription.md)
 - [NotifCondition](docs/NotifCondition.md)
 - [NotificationData](docs/NotificationData.md)
 - [NotificationEventType](docs/NotificationEventType.md)
 - [NotificationFlag](docs/NotificationFlag.md)
 - [NotificationMethod](docs/NotificationMethod.md)
 - [NotificationMethod1](docs/NotificationMethod1.md)
 - [NotificationType](docs/NotificationType.md)
 - [NrLocation](docs/NrLocation.md)
 - [NrfInfo](docs/NrfInfo.md)
 - [NrfInfoServedAanfInfoListValueValue](docs/NrfInfoServedAanfInfoListValueValue.md)
 - [NrfInfoServedAmfInfoValue](docs/NrfInfoServedAmfInfoValue.md)
 - [NrfInfoServedAusfInfoValue](docs/NrfInfoServedAusfInfoValue.md)
 - [NrfInfoServedBsfInfoValue](docs/NrfInfoServedBsfInfoValue.md)
 - [NrfInfoServedChfInfoValue](docs/NrfInfoServedChfInfoValue.md)
 - [NrfInfoServedGmlcInfoValue](docs/NrfInfoServedGmlcInfoValue.md)
 - [NrfInfoServedHssInfoListValueValue](docs/NrfInfoServedHssInfoListValueValue.md)
 - [NrfInfoServedLmfInfoValue](docs/NrfInfoServedLmfInfoValue.md)
 - [NrfInfoServedMbSmfInfoListValueValue](docs/NrfInfoServedMbSmfInfoListValueValue.md)
 - [NrfInfoServedNefInfoValue](docs/NrfInfoServedNefInfoValue.md)
 - [NrfInfoServedNwdafInfoValue](docs/NrfInfoServedNwdafInfoValue.md)
 - [NrfInfoServedPcfInfoValue](docs/NrfInfoServedPcfInfoValue.md)
 - [NrfInfoServedPcscfInfoListValueValue](docs/NrfInfoServedPcscfInfoListValueValue.md)
 - [NrfInfoServedScpInfoListValue](docs/NrfInfoServedScpInfoListValue.md)
 - [NrfInfoServedSeppInfoListValue](docs/NrfInfoServedSeppInfoListValue.md)
 - [NrfInfoServedSmfInfoValue](docs/NrfInfoServedSmfInfoValue.md)
 - [NrfInfoServedUdmInfoValue](docs/NrfInfoServedUdmInfoValue.md)
 - [NrfInfoServedUdrInfoValue](docs/NrfInfoServedUdrInfoValue.md)
 - [NrfInfoServedUdsfInfoValue](docs/NrfInfoServedUdsfInfoValue.md)
 - [NrfInfoServedUpfInfoValue](docs/NrfInfoServedUpfInfoValue.md)
 - [NsacfCapability](docs/NsacfCapability.md)
 - [NsacfInfo](docs/NsacfInfo.md)
 - [NsiIdInfo](docs/NsiIdInfo.md)
 - [NsiLoadLevelInfo](docs/NsiLoadLevelInfo.md)
 - [NsmfEventExposure](docs/NsmfEventExposure.md)
 - [NsmfEventExposureNotification](docs/NsmfEventExposureNotification.md)
 - [NssaafInfo](docs/NssaafInfo.md)
 - [NumberAverage](docs/NumberAverage.md)
 - [NwdafCapability](docs/NwdafCapability.md)
 - [NwdafCond](docs/NwdafCond.md)
 - [NwdafEvent](docs/NwdafEvent.md)
 - [NwdafFailureCode](docs/NwdafFailureCode.md)
 - [NwdafInfo](docs/NwdafInfo.md)
 - [ObservedRedundantTransExp](docs/ObservedRedundantTransExp.md)
 - [OutputStrategy](docs/OutputStrategy.md)
 - [PartitioningCriteria](docs/PartitioningCriteria.md)
 - [PcfInfo](docs/PcfInfo.md)
 - [PcscfInfo](docs/PcscfInfo.md)
 - [PdnConnectivityStatReport](docs/PdnConnectivityStatReport.md)
 - [PdnConnectivityStatus](docs/PdnConnectivityStatus.md)
 - [PduSessionInfo](docs/PduSessionInfo.md)
 - [PduSessionInformation](docs/PduSessionInformation.md)
 - [PduSessionStatus](docs/PduSessionStatus.md)
 - [PduSessionStatusCfg](docs/PduSessionStatusCfg.md)
 - [PduSessionType](docs/PduSessionType.md)
 - [PerUeAttribute](docs/PerUeAttribute.md)
 - [PerfData](docs/PerfData.md)
 - [PerformanceData](docs/PerformanceData.md)
 - [PerformanceDataCollection](docs/PerformanceDataCollection.md)
 - [PerformanceDataInfo](docs/PerformanceDataInfo.md)
 - [PfdData](docs/PfdData.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [PlmnOauth2](docs/PlmnOauth2.md)
 - [PlmnRange](docs/PlmnRange.md)
 - [PlmnSnssai](docs/PlmnSnssai.md)
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
 - [PresenceInfo](docs/PresenceInfo.md)
 - [PresenceState](docs/PresenceState.md)
 - [PrevSubInfo](docs/PrevSubInfo.md)
 - [ProSeCapability](docs/ProSeCapability.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [ProblemDetailsAnalyticsInfoRequest](docs/ProblemDetailsAnalyticsInfoRequest.md)
 - [QosRequirement](docs/QosRequirement.md)
 - [QosResourceType](docs/QosResourceType.md)
 - [QosSustainabilityInfo](docs/QosSustainabilityInfo.md)
 - [RankingCriterion](docs/RankingCriterion.md)
 - [RatFreqInformation](docs/RatFreqInformation.md)
 - [RatType](docs/RatType.md)
 - [ReachabilityFilter](docs/ReachabilityFilter.md)
 - [ReachabilityForDataConfiguration](docs/ReachabilityForDataConfiguration.md)
 - [ReachabilityForDataReportConfig](docs/ReachabilityForDataReportConfig.md)
 - [ReachabilityForSmsConfiguration](docs/ReachabilityForSmsConfiguration.md)
 - [ReachabilityForSmsReport](docs/ReachabilityForSmsReport.md)
 - [ReachabilityReport](docs/ReachabilityReport.md)
 - [RedTransExpOrderingCriterion](docs/RedTransExpOrderingCriterion.md)
 - [RedundantTransmissionExpInfo](docs/RedundantTransmissionExpInfo.md)
 - [RedundantTransmissionExpPerTS](docs/RedundantTransmissionExpPerTS.md)
 - [RedundantTransmissionExpReq](docs/RedundantTransmissionExpReq.md)
 - [RelativeCartesianLocation](docs/RelativeCartesianLocation.md)
 - [Report](docs/Report.md)
 - [ReportingInformation](docs/ReportingInformation.md)
 - [ReportingOptions](docs/ReportingOptions.md)
 - [RequestedContext](docs/RequestedContext.md)
 - [ResourceUsage](docs/ResourceUsage.md)
 - [RetainabilityThreshold](docs/RetainabilityThreshold.md)
 - [RmInfo](docs/RmInfo.md)
 - [RmState](docs/RmState.md)
 - [RoamingStatusReport](docs/RoamingStatusReport.md)
 - [RouteInformation](docs/RouteInformation.md)
 - [RouteToLocation](docs/RouteToLocation.md)
 - [RoutingAreaId](docs/RoutingAreaId.md)
 - [SACEvent](docs/SACEvent.md)
 - [SACEventReport](docs/SACEventReport.md)
 - [SACEventReportItem](docs/SACEventReportItem.md)
 - [SACEventState](docs/SACEventState.md)
 - [SACEventStatus](docs/SACEventStatus.md)
 - [SACEventSubscription](docs/SACEventSubscription.md)
 - [SACEventTrigger](docs/SACEventTrigger.md)
 - [SACEventType](docs/SACEventType.md)
 - [SACInfo](docs/SACInfo.md)
 - [ScheduledCommunicationTime](docs/ScheduledCommunicationTime.md)
 - [ScheduledCommunicationTime1](docs/ScheduledCommunicationTime1.md)
 - [ScheduledCommunicationType](docs/ScheduledCommunicationType.md)
 - [ScpCapability](docs/ScpCapability.md)
 - [ScpDomainCond](docs/ScpDomainCond.md)
 - [ScpDomainInfo](docs/ScpDomainInfo.md)
 - [ScpInfo](docs/ScpInfo.md)
 - [SdRange](docs/SdRange.md)
 - [SeppInfo](docs/SeppInfo.md)
 - [ServiceAreaId](docs/ServiceAreaId.md)
 - [ServiceDataFlowDescription](docs/ServiceDataFlowDescription.md)
 - [ServiceExperienceInfo](docs/ServiceExperienceInfo.md)
 - [ServiceExperienceInfo1](docs/ServiceExperienceInfo1.md)
 - [ServiceExperienceInfoPerApp](docs/ServiceExperienceInfoPerApp.md)
 - [ServiceExperienceInfoPerFlow](docs/ServiceExperienceInfoPerFlow.md)
 - [ServiceExperienceType](docs/ServiceExperienceType.md)
 - [ServiceName](docs/ServiceName.md)
 - [ServiceNameCond](docs/ServiceNameCond.md)
 - [ServiceNameListCond](docs/ServiceNameListCond.md)
 - [SessInactTimerForUeComm](docs/SessInactTimerForUeComm.md)
 - [SharedDataIdRange](docs/SharedDataIdRange.md)
 - [SliceLoadLevelInformation](docs/SliceLoadLevelInformation.md)
 - [SmNasFromSmf](docs/SmNasFromSmf.md)
 - [SmNasFromUe](docs/SmNasFromUe.md)
 - [SmcceInfo](docs/SmcceInfo.md)
 - [SmcceInfo1](docs/SmcceInfo1.md)
 - [SmcceUeList](docs/SmcceUeList.md)
 - [SmcceUeList1](docs/SmcceUeList1.md)
 - [SmfEvent](docs/SmfEvent.md)
 - [SmfInfo](docs/SmfInfo.md)
 - [SmsfInfo](docs/SmsfInfo.md)
 - [Snssai](docs/Snssai.md)
 - [SnssaiEasdfInfoItem](docs/SnssaiEasdfInfoItem.md)
 - [SnssaiExtension](docs/SnssaiExtension.md)
 - [SnssaiInfoItem](docs/SnssaiInfoItem.md)
 - [SnssaiMbSmfInfoItem](docs/SnssaiMbSmfInfoItem.md)
 - [SnssaiSmfInfoItem](docs/SnssaiSmfInfoItem.md)
 - [SnssaiTaiMapping](docs/SnssaiTaiMapping.md)
 - [SnssaiTsctsfInfoItem](docs/SnssaiTsctsfInfoItem.md)
 - [SnssaiUpfInfoItem](docs/SnssaiUpfInfoItem.md)
 - [SpecificAnalyticsSubscription](docs/SpecificAnalyticsSubscription.md)
 - [SpecificDataSubscription](docs/SpecificDataSubscription.md)
 - [Ssm](docs/Ssm.md)
 - [StationaryIndication](docs/StationaryIndication.md)
 - [SubTerminationReason](docs/SubTerminationReason.md)
 - [SubscrCond](docs/SubscrCond.md)
 - [SubscriptionContext](docs/SubscriptionContext.md)
 - [SubscriptionData](docs/SubscriptionData.md)
 - [SuciInfo](docs/SuciInfo.md)
 - [SupiRange](docs/SupiRange.md)
 - [SupportedGADShapes](docs/SupportedGADShapes.md)
 - [SupportedSnssai](docs/SupportedSnssai.md)
 - [SvcExperience](docs/SvcExperience.md)
 - [TacRange](docs/TacRange.md)
 - [Tai](docs/Tai.md)
 - [TaiRange](docs/TaiRange.md)
 - [TargetArea](docs/TargetArea.md)
 - [TargetUeIdentification](docs/TargetUeIdentification.md)
 - [TargetUeInformation](docs/TargetUeInformation.md)
 - [ThresholdLevel](docs/ThresholdLevel.md)
 - [TimeUnit](docs/TimeUnit.md)
 - [TimeWindow](docs/TimeWindow.md)
 - [Tmgi](docs/Tmgi.md)
 - [TmgiRange](docs/TmgiRange.md)
 - [TnapId](docs/TnapId.md)
 - [TngfInfo](docs/TngfInfo.md)
 - [TopApplication](docs/TopApplication.md)
 - [TrafficCharacterization](docs/TrafficCharacterization.md)
 - [TrafficDescriptor](docs/TrafficDescriptor.md)
 - [TrafficInformation](docs/TrafficInformation.md)
 - [TrafficProfile](docs/TrafficProfile.md)
 - [TransactionInfo](docs/TransactionInfo.md)
 - [TransactionMetric](docs/TransactionMetric.md)
 - [TransportProtocol](docs/TransportProtocol.md)
 - [TransportProtocol1](docs/TransportProtocol1.md)
 - [TrustAfInfo](docs/TrustAfInfo.md)
 - [TsctsfInfo](docs/TsctsfInfo.md)
 - [TwapId](docs/TwapId.md)
 - [TwifInfo](docs/TwifInfo.md)
 - [UEIdExt](docs/UEIdExt.md)
 - [UPInterfaceType](docs/UPInterfaceType.md)
 - [UdmInfo](docs/UdmInfo.md)
 - [UdrInfo](docs/UdrInfo.md)
 - [UdsfInfo](docs/UdsfInfo.md)
 - [UeAccessBehaviorReportItem](docs/UeAccessBehaviorReportItem.md)
 - [UeAnalyticsContextDescriptor](docs/UeAnalyticsContextDescriptor.md)
 - [UeCommunication](docs/UeCommunication.md)
 - [UeCommunicationCollection](docs/UeCommunicationCollection.md)
 - [UeCommunicationInfo](docs/UeCommunicationInfo.md)
 - [UeInAreaFilter](docs/UeInAreaFilter.md)
 - [UeLocationTrendsReportItem](docs/UeLocationTrendsReportItem.md)
 - [UeMobility](docs/UeMobility.md)
 - [UeMobilityCollection](docs/UeMobilityCollection.md)
 - [UeMobilityInfo](docs/UeMobilityInfo.md)
 - [UeReachability](docs/UeReachability.md)
 - [UeTrajectoryCollection](docs/UeTrajectoryCollection.md)
 - [UeTrajectoryInfo](docs/UeTrajectoryInfo.md)
 - [UeType](docs/UeType.md)
 - [UmtTime](docs/UmtTime.md)
 - [UnTrustAfInfo](docs/UnTrustAfInfo.md)
 - [UncertaintyEllipse](docs/UncertaintyEllipse.md)
 - [UncertaintyEllipsoid](docs/UncertaintyEllipsoid.md)
 - [UpfCond](docs/UpfCond.md)
 - [UpfInfo](docs/UpfInfo.md)
 - [UpfInformation](docs/UpfInformation.md)
 - [UriScheme](docs/UriScheme.md)
 - [UsageThreshold](docs/UsageThreshold.md)
 - [UserDataCongestionCollection](docs/UserDataCongestionCollection.md)
 - [UserDataCongestionInfo](docs/UserDataCongestionInfo.md)
 - [UserLocation](docs/UserLocation.md)
 - [UtraLocation](docs/UtraLocation.md)
 - [V2xCapability](docs/V2xCapability.md)
 - [VendorSpecificFeature](docs/VendorSpecificFeature.md)
 - [WAgfInfo](docs/WAgfInfo.md)
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
 - **nnwdaf-analyticsinfo**: Access to the Nnwdaf_AnalyticsInfo API

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


