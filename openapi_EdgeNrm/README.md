# Go API client for openapi_EdgeNrm

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 17.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi_EdgeNrm "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_EdgeNrm"
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
ctx := context.WithValue(context.Background(), openapi_EdgeNrm.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_EdgeNrm.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_EdgeNrm.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_EdgeNrm.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------


## Documentation For Models

 - [AckState](docs/AckState.md)
 - [AdministrativeState](docs/AdministrativeState.md)
 - [AffinityAntiAffinity](docs/AffinityAntiAffinity.md)
 - [AlarmListSingle](docs/AlarmListSingle.md)
 - [AlarmListSingleAllOf](docs/AlarmListSingleAllOf.md)
 - [AlarmListSingleAllOfAttributes](docs/AlarmListSingleAllOfAttributes.md)
 - [AlarmNotificationTypes](docs/AlarmNotificationTypes.md)
 - [AlarmRecord](docs/AlarmRecord.md)
 - [AlarmType](docs/AlarmType.md)
 - [AnonymizationOfMdtDataType](docs/AnonymizationOfMdtDataType.md)
 - [AreaConfig](docs/AreaConfig.md)
 - [AreaOfInterest](docs/AreaOfInterest.md)
 - [AreaScope](docs/AreaScope.md)
 - [CmNotificationTypes](docs/CmNotificationTypes.md)
 - [CollectionPeriodM6LteType](docs/CollectionPeriodM6LteType.md)
 - [CollectionPeriodM6NrType](docs/CollectionPeriodM6NrType.md)
 - [CollectionPeriodRrmLteType](docs/CollectionPeriodRrmLteType.md)
 - [CollectionPeriodRrmNrType](docs/CollectionPeriodRrmNrType.md)
 - [CollectionPeriodRrmUmtsType](docs/CollectionPeriodRrmUmtsType.md)
 - [CorrelatedNotification](docs/CorrelatedNotification.md)
 - [EASFunctionSingle](docs/EASFunctionSingle.md)
 - [EASFunctionSingleAllOf](docs/EASFunctionSingleAllOf.md)
 - [EASFunctionSingleAllOfAttributes](docs/EASFunctionSingleAllOfAttributes.md)
 - [EASFunctionSingleAllOfAttributesAllOf](docs/EASFunctionSingleAllOfAttributesAllOf.md)
 - [EASRequirements](docs/EASRequirements.md)
 - [EASRequirementsAllOf](docs/EASRequirementsAllOf.md)
 - [ECSFunctionSingle](docs/ECSFunctionSingle.md)
 - [ECSFunctionSingleAllOf](docs/ECSFunctionSingleAllOf.md)
 - [ECSFunctionSingleAllOfAttributes](docs/ECSFunctionSingleAllOfAttributes.md)
 - [ECSFunctionSingleAllOfAttributesAllOf](docs/ECSFunctionSingleAllOfAttributesAllOf.md)
 - [EDNConnectionInfo](docs/EDNConnectionInfo.md)
 - [EESFunctionSingle](docs/EESFunctionSingle.md)
 - [EESFunctionSingleAllOf](docs/EESFunctionSingleAllOf.md)
 - [EESFunctionSingleAllOfAttributes](docs/EESFunctionSingleAllOfAttributes.md)
 - [EESFunctionSingleAllOfAttributesAllOf](docs/EESFunctionSingleAllOfAttributesAllOf.md)
 - [EdgeDataNetworkSingle](docs/EdgeDataNetworkSingle.md)
 - [EdgeDataNetworkSingleAllOf](docs/EdgeDataNetworkSingleAllOf.md)
 - [EdgeDataNetworkSingleAllOf1](docs/EdgeDataNetworkSingleAllOf1.md)
 - [EventListForEventTriggeredMeasurementType](docs/EventListForEventTriggeredMeasurementType.md)
 - [EventThresholdL1Type](docs/EventThresholdL1Type.md)
 - [EventThresholdType](docs/EventThresholdType.md)
 - [EventThresholdTypeEventThreshold1F](docs/EventThresholdTypeEventThreshold1F.md)
 - [EventThresholdTypeEventThresholdRSRP](docs/EventThresholdTypeEventThresholdRSRP.md)
 - [EventThresholdTypeEventThresholdRSRQ](docs/EventThresholdTypeEventThresholdRSRQ.md)
 - [FileDownloadJobProcessMonitor](docs/FileDownloadJobProcessMonitor.md)
 - [FileDownloadJobProcessMonitorResultStateInfo](docs/FileDownloadJobProcessMonitorResultStateInfo.md)
 - [FileDownloadJobSingle](docs/FileDownloadJobSingle.md)
 - [FileDownloadJobSingleAllOf](docs/FileDownloadJobSingleAllOf.md)
 - [FileDownloadJobSingleAllOfAttributes](docs/FileDownloadJobSingleAllOfAttributes.md)
 - [FileNotificationTypes](docs/FileNotificationTypes.md)
 - [FileSingle](docs/FileSingle.md)
 - [FileSingleAllOf](docs/FileSingleAllOf.md)
 - [FileSingleAllOfAttributes](docs/FileSingleAllOfAttributes.md)
 - [FilesSingle](docs/FilesSingle.md)
 - [FilesSingleAllOf](docs/FilesSingleAllOf.md)
 - [FilesSingleAllOfAttributes](docs/FilesSingleAllOfAttributes.md)
 - [FreqInfo](docs/FreqInfo.md)
 - [GeoArea](docs/GeoArea.md)
 - [GeoAreaToCellMapping](docs/GeoAreaToCellMapping.md)
 - [GeoCoordinate](docs/GeoCoordinate.md)
 - [GeoLoc](docs/GeoLoc.md)
 - [GeographicalCoordinates](docs/GeographicalCoordinates.md)
 - [HeartbeatControlSingle](docs/HeartbeatControlSingle.md)
 - [HeartbeatControlSingleAllOf](docs/HeartbeatControlSingleAllOf.md)
 - [HeartbeatControlSingleAllOfAttributes](docs/HeartbeatControlSingleAllOfAttributes.md)
 - [HeartbeatNotificationTypes](docs/HeartbeatNotificationTypes.md)
 - [HostAddr](docs/HostAddr.md)
 - [IpAddr](docs/IpAddr.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [JobTypeType](docs/JobTypeType.md)
 - [ListOfInterfacesType](docs/ListOfInterfacesType.md)
 - [ListOfMeasurementsType](docs/ListOfMeasurementsType.md)
 - [LoggingDurationType](docs/LoggingDurationType.md)
 - [LoggingIntervalType](docs/LoggingIntervalType.md)
 - [ManagedFunctionAttr](docs/ManagedFunctionAttr.md)
 - [ManagedFunctionNcO](docs/ManagedFunctionNcO.md)
 - [ManagedNFServiceSingle](docs/ManagedNFServiceSingle.md)
 - [ManagedNFServiceSingleAllOf](docs/ManagedNFServiceSingleAllOf.md)
 - [ManagedNFServiceSingleAllOfAttributes](docs/ManagedNFServiceSingleAllOfAttributes.md)
 - [ManagementData](docs/ManagementData.md)
 - [ManagementDataCollectionSingle](docs/ManagementDataCollectionSingle.md)
 - [ManagementDataCollectionSingleAllOf](docs/ManagementDataCollectionSingleAllOf.md)
 - [ManagementDataCollectionSingleAllOfAttributes](docs/ManagementDataCollectionSingleAllOfAttributes.md)
 - [ManagementNodeSingle](docs/ManagementNodeSingle.md)
 - [ManagementNodeSingleAllOf](docs/ManagementNodeSingleAllOf.md)
 - [ManagementNodeSingleAllOfAttributes](docs/ManagementNodeSingleAllOfAttributes.md)
 - [MbsfnArea](docs/MbsfnArea.md)
 - [MeContextSingle](docs/MeContextSingle.md)
 - [MeContextSingleAllOf](docs/MeContextSingleAllOf.md)
 - [MeContextSingleAllOfAttributes](docs/MeContextSingleAllOfAttributes.md)
 - [MeasurementPeriodLteType](docs/MeasurementPeriodLteType.md)
 - [MeasurementPeriodUmtsType](docs/MeasurementPeriodUmtsType.md)
 - [MeasurementQuantityType](docs/MeasurementQuantityType.md)
 - [MnS](docs/MnS.md)
 - [MnSOneOf](docs/MnSOneOf.md)
 - [MnsAgentSingle](docs/MnsAgentSingle.md)
 - [MnsAgentSingleAllOf](docs/MnsAgentSingleAllOf.md)
 - [MnsAgentSingleAllOfAttributes](docs/MnsAgentSingleAllOfAttributes.md)
 - [MnsInfoSingle](docs/MnsInfoSingle.md)
 - [MnsRegistrySingle](docs/MnsRegistrySingle.md)
 - [NFServiceType](docs/NFServiceType.md)
 - [NFType](docs/NFType.md)
 - [NodeFilter](docs/NodeFilter.md)
 - [NotificationType](docs/NotificationType.md)
 - [NtfSubscriptionControlSingle](docs/NtfSubscriptionControlSingle.md)
 - [NtfSubscriptionControlSingleAllOf](docs/NtfSubscriptionControlSingleAllOf.md)
 - [NtfSubscriptionControlSingleAllOfAttributes](docs/NtfSubscriptionControlSingleAllOfAttributes.md)
 - [Operation](docs/Operation.md)
 - [OperationSemantics](docs/OperationSemantics.md)
 - [OperationalState](docs/OperationalState.md)
 - [PeeParameter](docs/PeeParameter.md)
 - [PerceivedSeverity](docs/PerceivedSeverity.md)
 - [PerfMetricJobSingle](docs/PerfMetricJobSingle.md)
 - [PerfMetricJobSingleAllOf](docs/PerfMetricJobSingleAllOf.md)
 - [PerfMetricJobSingleAllOfAttributes](docs/PerfMetricJobSingleAllOfAttributes.md)
 - [PerfNotificationTypes](docs/PerfNotificationTypes.md)
 - [PlmnId](docs/PlmnId.md)
 - [PlmnListTypeInner](docs/PlmnListTypeInner.md)
 - [PlmnTargetType](docs/PlmnTargetType.md)
 - [PositioningMethodType](docs/PositioningMethodType.md)
 - [ProbableCause](docs/ProbableCause.md)
 - [ProbableCauseOneOf](docs/ProbableCauseOneOf.md)
 - [RegistrationInfo](docs/RegistrationInfo.md)
 - [RegistrationState](docs/RegistrationState.md)
 - [ReportAmountType](docs/ReportAmountType.md)
 - [ReportIntervalType](docs/ReportIntervalType.md)
 - [ReportTypeType](docs/ReportTypeType.md)
 - [ReportingCtrl](docs/ReportingCtrl.md)
 - [ReportingCtrlOneOf](docs/ReportingCtrlOneOf.md)
 - [ReportingCtrlOneOf1](docs/ReportingCtrlOneOf1.md)
 - [ReportingCtrlOneOf2](docs/ReportingCtrlOneOf2.md)
 - [ResourcesEdgeNrm](docs/ResourcesEdgeNrm.md)
 - [SAP](docs/SAP.md)
 - [Scope](docs/Scope.md)
 - [ServingLocation](docs/ServingLocation.md)
 - [SoftwareImageInfo](docs/SoftwareImageInfo.md)
 - [SpecificProblem](docs/SpecificProblem.md)
 - [SubNetworkAttr](docs/SubNetworkAttr.md)
 - [SubNetworkNcO](docs/SubNetworkNcO.md)
 - [SubNetworkSingle](docs/SubNetworkSingle.md)
 - [SubNetworkSingleAllOf](docs/SubNetworkSingleAllOf.md)
 - [SubNetworkSingleAllOf1](docs/SubNetworkSingleAllOf1.md)
 - [SubNetworkSingleAllOfAttributes](docs/SubNetworkSingleAllOfAttributes.md)
 - [SupportedPerfMetricGroup](docs/SupportedPerfMetricGroup.md)
 - [Tai](docs/Tai.md)
 - [Tai1](docs/Tai1.md)
 - [ThresholdInfo](docs/ThresholdInfo.md)
 - [ThresholdInfo1](docs/ThresholdInfo1.md)
 - [ThresholdInfoHysteresis](docs/ThresholdInfoHysteresis.md)
 - [ThresholdInfoThresholdValue](docs/ThresholdInfoThresholdValue.md)
 - [ThresholdLevelInd](docs/ThresholdLevelInd.md)
 - [ThresholdLevelIndOneOf](docs/ThresholdLevelIndOneOf.md)
 - [ThresholdLevelIndOneOf1](docs/ThresholdLevelIndOneOf1.md)
 - [ThresholdMonitorSingle](docs/ThresholdMonitorSingle.md)
 - [ThresholdMonitorSingleAllOf](docs/ThresholdMonitorSingleAllOf.md)
 - [ThresholdMonitorSingleAllOfAttributes](docs/ThresholdMonitorSingleAllOfAttributes.md)
 - [TimeToTriggerL1Type](docs/TimeToTriggerL1Type.md)
 - [TimeWindow](docs/TimeWindow.md)
 - [Top](docs/Top.md)
 - [TopologicalServiceArea](docs/TopologicalServiceArea.md)
 - [TraceDepthType](docs/TraceDepthType.md)
 - [TraceJobAttr](docs/TraceJobAttr.md)
 - [TraceJobSingle](docs/TraceJobSingle.md)
 - [TraceJobSingleAllOf](docs/TraceJobSingleAllOf.md)
 - [TraceReferenceType](docs/TraceReferenceType.md)
 - [TraceReportingFormatType](docs/TraceReportingFormatType.md)
 - [TraceTargetType](docs/TraceTargetType.md)
 - [TrendIndication](docs/TrendIndication.md)
 - [TriggeringEventsType](docs/TriggeringEventsType.md)
 - [UsageState](docs/UsageState.md)
 - [VirtualResource](docs/VirtualResource.md)
 - [VnfParameter](docs/VnfParameter.md)
 - [VsDataContainerSingle](docs/VsDataContainerSingle.md)
 - [VsDataContainerSingleAttributes](docs/VsDataContainerSingleAttributes.md)


## Documentation For Authorization

 Endpoints do not require authorization.


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

