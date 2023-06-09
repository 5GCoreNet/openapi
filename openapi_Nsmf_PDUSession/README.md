# Go API client for openapi_Nsmf_PDUSession

SMF PDU Session Service.  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.3.0-alpha.2
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
import openapi_Nsmf_PDUSession "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Nsmf_PDUSession"
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
ctx := context.WithValue(context.Background(), openapi_Nsmf_PDUSession.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Nsmf_PDUSession.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Nsmf_PDUSession.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Nsmf_PDUSession.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/nsmf-pdusession/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualPDUSessionHSMFOrSMFApi* | [**ReleasePduSession**](docs/IndividualPDUSessionHSMFOrSMFApi.md#releasepdusession) | **Post** /pdu-sessions/{pduSessionRef}/release | Release
*IndividualPDUSessionHSMFOrSMFApi* | [**RetrievePduSession**](docs/IndividualPDUSessionHSMFOrSMFApi.md#retrievepdusession) | **Post** /pdu-sessions/{pduSessionRef}/retrieve | Retrieve
*IndividualPDUSessionHSMFOrSMFApi* | [**TransferMoData**](docs/IndividualPDUSessionHSMFOrSMFApi.md#transfermodata) | **Post** /pdu-sessions/{pduSessionRef}/transfer-mo-data | Transfer MO Data
*IndividualPDUSessionHSMFOrSMFApi* | [**UpdatePduSession**](docs/IndividualPDUSessionHSMFOrSMFApi.md#updatepdusession) | **Post** /pdu-sessions/{pduSessionRef}/modify | Update (initiated by V-SMF or I-SMF)
*IndividualSMContextApi* | [**ReleaseSmContext**](docs/IndividualSMContextApi.md#releasesmcontext) | **Post** /sm-contexts/{smContextRef}/release | Release SM Context
*IndividualSMContextApi* | [**RetrieveSmContext**](docs/IndividualSMContextApi.md#retrievesmcontext) | **Post** /sm-contexts/{smContextRef}/retrieve | Retrieve SM Context
*IndividualSMContextApi* | [**SendMoData**](docs/IndividualSMContextApi.md#sendmodata) | **Post** /sm-contexts/{smContextRef}/send-mo-data | Send MO Data
*IndividualSMContextApi* | [**UpdateSmContext**](docs/IndividualSMContextApi.md#updatesmcontext) | **Post** /sm-contexts/{smContextRef}/modify | Update SM Context
*PDUSessionsCollectionApi* | [**PostPduSessions**](docs/PDUSessionsCollectionApi.md#postpdusessions) | **Post** /pdu-sessions | Create
*SMContextsCollectionApi* | [**PostSmContexts**](docs/SMContextsCollectionApi.md#postsmcontexts) | **Post** /sm-contexts | Create SM Context


## Documentation For Models

 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AccessType](docs/AccessType.md)
 - [AdditionalQosFlowInfo](docs/AdditionalQosFlowInfo.md)
 - [AdditionalQosFlowInfoAnyOf](docs/AdditionalQosFlowInfoAnyOf.md)
 - [AfCoordinationInfo](docs/AfCoordinationInfo.md)
 - [AlternativeQosProfile](docs/AlternativeQosProfile.md)
 - [Ambr](docs/Ambr.md)
 - [AnchorSmfFeatures](docs/AnchorSmfFeatures.md)
 - [ApnRateStatus](docs/ApnRateStatus.md)
 - [Arp](docs/Arp.md)
 - [BackupAmfInfo](docs/BackupAmfInfo.md)
 - [BatteryIndication](docs/BatteryIndication.md)
 - [Cause](docs/Cause.md)
 - [CellGlobalId](docs/CellGlobalId.md)
 - [ChargingInformation](docs/ChargingInformation.md)
 - [CnAssistedRanPara](docs/CnAssistedRanPara.md)
 - [DddTrafficDescriptor](docs/DddTrafficDescriptor.md)
 - [DdnFailureSubInfo](docs/DdnFailureSubInfo.md)
 - [DdnFailureSubs](docs/DdnFailureSubs.md)
 - [DnaiInformation](docs/DnaiInformation.md)
 - [DnnSelectionMode](docs/DnnSelectionMode.md)
 - [Dynamic5Qi](docs/Dynamic5Qi.md)
 - [EbiArpMapping](docs/EbiArpMapping.md)
 - [Ecgi](docs/Ecgi.md)
 - [EpsBearerInfo](docs/EpsBearerInfo.md)
 - [EpsInterworkingIndication](docs/EpsInterworkingIndication.md)
 - [EpsPdnCnxInfo](docs/EpsPdnCnxInfo.md)
 - [EutraLocation](docs/EutraLocation.md)
 - [ExemptionInd](docs/ExemptionInd.md)
 - [ExtProblemDetails](docs/ExtProblemDetails.md)
 - [GNbId](docs/GNbId.md)
 - [GbrQosFlowInformation](docs/GbrQosFlowInformation.md)
 - [GeraLocation](docs/GeraLocation.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [Guami](docs/Guami.md)
 - [HfcNodeId](docs/HfcNodeId.md)
 - [HoState](docs/HoState.md)
 - [HsmfUpdateData](docs/HsmfUpdateData.md)
 - [HsmfUpdateError](docs/HsmfUpdateError.md)
 - [HsmfUpdatedData](docs/HsmfUpdatedData.md)
 - [IndirectDataForwardingTunnelInfo](docs/IndirectDataForwardingTunnelInfo.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddress](docs/IpAddress.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [LineType](docs/LineType.md)
 - [LocationAreaId](docs/LocationAreaId.md)
 - [MaReleaseIndication](docs/MaReleaseIndication.md)
 - [MaxIntegrityProtectedDataRate](docs/MaxIntegrityProtectedDataRate.md)
 - [MmeCapabilities](docs/MmeCapabilities.md)
 - [MoExpDataCounter](docs/MoExpDataCounter.md)
 - [N2SmInfoType](docs/N2SmInfoType.md)
 - [N3gaLocation](docs/N3gaLocation.md)
 - [N4Information](docs/N4Information.md)
 - [N4MessageType](docs/N4MessageType.md)
 - [NFType](docs/NFType.md)
 - [Ncgi](docs/Ncgi.md)
 - [NgApCause](docs/NgApCause.md)
 - [NgRanTargetId](docs/NgRanTargetId.md)
 - [NonDynamic5Qi](docs/NonDynamic5Qi.md)
 - [NotificationCause](docs/NotificationCause.md)
 - [NotificationControl](docs/NotificationControl.md)
 - [NotificationInfo](docs/NotificationInfo.md)
 - [NrLocation](docs/NrLocation.md)
 - [NullValue](docs/NullValue.md)
 - [PartialRecordMethod](docs/PartialRecordMethod.md)
 - [PcfUeCallbackInfo](docs/PcfUeCallbackInfo.md)
 - [PduSessionContextType](docs/PduSessionContextType.md)
 - [PduSessionCreateData](docs/PduSessionCreateData.md)
 - [PduSessionCreateError](docs/PduSessionCreateError.md)
 - [PduSessionCreatedData](docs/PduSessionCreatedData.md)
 - [PduSessionNotifyItem](docs/PduSessionNotifyItem.md)
 - [PduSessionType](docs/PduSessionType.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnIdNid](docs/PlmnIdNid.md)
 - [PreemptionCapability](docs/PreemptionCapability.md)
 - [PreemptionVulnerability](docs/PreemptionVulnerability.md)
 - [PresenceState](docs/PresenceState.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [ProblemDetailsAddInfo](docs/ProblemDetailsAddInfo.md)
 - [ProtectionResult](docs/ProtectionResult.md)
 - [PsaIndication](docs/PsaIndication.md)
 - [PsaInformation](docs/PsaInformation.md)
 - [QosFlowAccessType](docs/QosFlowAccessType.md)
 - [QosFlowAddModifyRequestItem](docs/QosFlowAddModifyRequestItem.md)
 - [QosFlowItem](docs/QosFlowItem.md)
 - [QosFlowNotifyItem](docs/QosFlowNotifyItem.md)
 - [QosFlowProfile](docs/QosFlowProfile.md)
 - [QosFlowReleaseRequestItem](docs/QosFlowReleaseRequestItem.md)
 - [QosFlowSetupItem](docs/QosFlowSetupItem.md)
 - [QosFlowTunnel](docs/QosFlowTunnel.md)
 - [QosFlowUsageReport](docs/QosFlowUsageReport.md)
 - [QosMonitoringInfo](docs/QosMonitoringInfo.md)
 - [QosMonitoringReq](docs/QosMonitoringReq.md)
 - [QosResourceType](docs/QosResourceType.md)
 - [RatType](docs/RatType.md)
 - [RedirectResponse](docs/RedirectResponse.md)
 - [RedundantPduSessionInformation](docs/RedundantPduSessionInformation.md)
 - [RefToBinaryData](docs/RefToBinaryData.md)
 - [ReflectiveQoSAttribute](docs/ReflectiveQoSAttribute.md)
 - [ReleaseData](docs/ReleaseData.md)
 - [ReleasePduSession200Response](docs/ReleasePduSession200Response.md)
 - [ReleasePduSessionRequest](docs/ReleasePduSessionRequest.md)
 - [ReleaseSmContextRequest](docs/ReleaseSmContextRequest.md)
 - [ReleasedData](docs/ReleasedData.md)
 - [RequestIndication](docs/RequestIndication.md)
 - [RequestType](docs/RequestType.md)
 - [ResourceStatus](docs/ResourceStatus.md)
 - [RetrieveData](docs/RetrieveData.md)
 - [RetrievedData](docs/RetrievedData.md)
 - [RoamingChargingProfile](docs/RoamingChargingProfile.md)
 - [RoutingAreaId](docs/RoutingAreaId.md)
 - [Rsn](docs/Rsn.md)
 - [SatelliteBackhaulCategory](docs/SatelliteBackhaulCategory.md)
 - [SbiBindingLevel](docs/SbiBindingLevel.md)
 - [ScheduledCommunicationTime](docs/ScheduledCommunicationTime.md)
 - [ScheduledCommunicationType](docs/ScheduledCommunicationType.md)
 - [SecondaryRatUsageInfo](docs/SecondaryRatUsageInfo.md)
 - [SecondaryRatUsageReport](docs/SecondaryRatUsageReport.md)
 - [SecurityResult](docs/SecurityResult.md)
 - [SendMoDataReqData](docs/SendMoDataReqData.md)
 - [SendMoDataRequest](docs/SendMoDataRequest.md)
 - [ServerAddressingInfo](docs/ServerAddressingInfo.md)
 - [ServiceAreaId](docs/ServiceAreaId.md)
 - [ServiceName](docs/ServiceName.md)
 - [SmContext](docs/SmContext.md)
 - [SmContextCreateData](docs/SmContextCreateData.md)
 - [SmContextCreateError](docs/SmContextCreateError.md)
 - [SmContextCreatedData](docs/SmContextCreatedData.md)
 - [SmContextReleaseData](docs/SmContextReleaseData.md)
 - [SmContextReleasedData](docs/SmContextReleasedData.md)
 - [SmContextRetrieveData](docs/SmContextRetrieveData.md)
 - [SmContextRetrievedData](docs/SmContextRetrievedData.md)
 - [SmContextStatusNotification](docs/SmContextStatusNotification.md)
 - [SmContextType](docs/SmContextType.md)
 - [SmContextUpdateData](docs/SmContextUpdateData.md)
 - [SmContextUpdateError](docs/SmContextUpdateError.md)
 - [SmContextUpdatedData](docs/SmContextUpdatedData.md)
 - [SmallDataRateStatus](docs/SmallDataRateStatus.md)
 - [SmfSelectionType](docs/SmfSelectionType.md)
 - [Snssai](docs/Snssai.md)
 - [StationaryIndication](docs/StationaryIndication.md)
 - [StatusInfo](docs/StatusInfo.md)
 - [StatusNotification](docs/StatusNotification.md)
 - [Tai](docs/Tai.md)
 - [TargetDnaiInfo](docs/TargetDnaiInfo.md)
 - [TnapId](docs/TnapId.md)
 - [TngfInfo](docs/TngfInfo.md)
 - [TraceData](docs/TraceData.md)
 - [TraceDepth](docs/TraceDepth.md)
 - [TrafficProfile](docs/TrafficProfile.md)
 - [TransferMoDataReqData](docs/TransferMoDataReqData.md)
 - [TransferMoDataRequest](docs/TransferMoDataRequest.md)
 - [TransferMtDataAddInfo](docs/TransferMtDataAddInfo.md)
 - [TransferMtDataError](docs/TransferMtDataError.md)
 - [TransferMtDataIsmfRequest](docs/TransferMtDataIsmfRequest.md)
 - [TransferMtDataReqData](docs/TransferMtDataReqData.md)
 - [TransportProtocol](docs/TransportProtocol.md)
 - [Trigger](docs/Trigger.md)
 - [TriggerCategory](docs/TriggerCategory.md)
 - [TriggerType](docs/TriggerType.md)
 - [TunnelInfo](docs/TunnelInfo.md)
 - [TwapId](docs/TwapId.md)
 - [TwifInfo](docs/TwifInfo.md)
 - [UlclBpInformation](docs/UlclBpInformation.md)
 - [UnavailableAccessIndication](docs/UnavailableAccessIndication.md)
 - [UpCnxState](docs/UpCnxState.md)
 - [UpConfidentiality](docs/UpConfidentiality.md)
 - [UpIntegrity](docs/UpIntegrity.md)
 - [UpSecurity](docs/UpSecurity.md)
 - [UpSecurityInfo](docs/UpSecurityInfo.md)
 - [UpdatePduSession200Response](docs/UpdatePduSession200Response.md)
 - [UpdatePduSession400Response](docs/UpdatePduSession400Response.md)
 - [UpdatePduSessionRequest](docs/UpdatePduSessionRequest.md)
 - [UpdateSmContext200Response](docs/UpdateSmContext200Response.md)
 - [UpdateSmContext400Response](docs/UpdateSmContext400Response.md)
 - [UpdateSmContextRequest](docs/UpdateSmContextRequest.md)
 - [UserLocation](docs/UserLocation.md)
 - [UtraLocation](docs/UtraLocation.md)
 - [VolumeTimedReport](docs/VolumeTimedReport.md)
 - [VplmnQos](docs/VplmnQos.md)
 - [VsmfUpdateData](docs/VsmfUpdateData.md)
 - [VsmfUpdateError](docs/VsmfUpdateError.md)
 - [VsmfUpdatedData](docs/VsmfUpdatedData.md)
 - [WAgfInfo](docs/WAgfInfo.md)


## Documentation For Authorization



### oAuth2ClientCredentials


- **Type**: OAuth
- **Flow**: application
- **Authorization URL**: 
- **Scopes**: 
 - **nsmf-pdusession**: Access to the nsmf-pdusession API

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



