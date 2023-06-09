# Go API client for openapi_Application_Data

The API version is defined in 3GPP TS 29.504  
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).  
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: -
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
import openapi_Application_Data "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_Application_Data"
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
ctx := context.WithValue(context.Background(), openapi_Application_Data.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_Application_Data.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_Application_Data.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_Application_Data.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AMInfluenceDataStoreApi* | [**ReadAmInfluenceData**](docs/AMInfluenceDataStoreApi.md#readaminfluencedata) | **Get** /application-data/am-influence-data | Retrieve AM Influence Data
*ApplicationDataSubscriptionsCollectionApi* | [**CreateIndividualApplicationDataSubscription**](docs/ApplicationDataSubscriptionsCollectionApi.md#createindividualapplicationdatasubscription) | **Post** /application-data/subs-to-notify | Create a subscription to receive notification of application data changes
*ApplicationDataSubscriptionsCollectionApi* | [**ReadApplicationDataChangeSubscriptions**](docs/ApplicationDataSubscriptionsCollectionApi.md#readapplicationdatachangesubscriptions) | **Get** /application-data/subs-to-notify | Read Application Data change Subscriptions
*BdtPolicyDataStoreApi* | [**ReadBdtPolicyData**](docs/BdtPolicyDataStoreApi.md#readbdtpolicydata) | **Get** /application-data/bdtPolicyData | Retrieve applied BDT Policy Data
*EASDeploymentDataStoreApi* | [**ReadEasDeployData**](docs/EASDeploymentDataStoreApi.md#readeasdeploydata) | **Get** /application-data/eas-deploy-data | Retrieve EAS Deployment Information Data
*IPTVConfigurationDataStoreApi* | [**ReadIPTVCongifurationData**](docs/IPTVConfigurationDataStoreApi.md#readiptvcongifurationdata) | **Get** /application-data/iptvConfigData | Retrieve IPTV configuration Data
*IndividualAMInfluenceDataDocumentApi* | [**CreateOrReplaceIndividualAmInfluenceData**](docs/IndividualAMInfluenceDataDocumentApi.md#createorreplaceindividualaminfluencedata) | **Put** /application-data/am-influence-data/{amInfluenceId} | Create or update an individual AM Influence Data resource
*IndividualAMInfluenceDataDocumentApi* | [**DeleteIndividualAmInfluenceData**](docs/IndividualAMInfluenceDataDocumentApi.md#deleteindividualaminfluencedata) | **Delete** /application-data/am-influence-data/{amInfluenceId} | Delete an individual AM Influence Data resource
*IndividualAMInfluenceDataDocumentApi* | [**UpdateIndividualAmInfluenceData**](docs/IndividualAMInfluenceDataDocumentApi.md#updateindividualaminfluencedata) | **Patch** /application-data/am-influence-data/{amInfluenceId} | Modify part of the properties of an individual AM Influence Data resource
*IndividualApplicationDataSubscriptionDocumentApi* | [**DeleteIndividualApplicationDataSubscription**](docs/IndividualApplicationDataSubscriptionDocumentApi.md#deleteindividualapplicationdatasubscription) | **Delete** /application-data/subs-to-notify/{subsId} | Delete the individual Application Data subscription
*IndividualApplicationDataSubscriptionDocumentApi* | [**ReadIndividualApplicationDataSubscription**](docs/IndividualApplicationDataSubscriptionDocumentApi.md#readindividualapplicationdatasubscription) | **Get** /application-data/subs-to-notify/{subsId} | Get an existing individual Application Data Subscription resource
*IndividualApplicationDataSubscriptionDocumentApi* | [**ReplaceIndividualApplicationDataSubscription**](docs/IndividualApplicationDataSubscriptionDocumentApi.md#replaceindividualapplicationdatasubscription) | **Put** /application-data/subs-to-notify/{subsId} | Modify a subscription to receive notification of application data changes
*IndividualAppliedBDTPolicyDataDocumentApi* | [**CreateIndividualAppliedBdtPolicyData**](docs/IndividualAppliedBDTPolicyDataDocumentApi.md#createindividualappliedbdtpolicydata) | **Put** /application-data/bdtPolicyData/{bdtPolicyId} | Create an individual applied BDT Policy Data resource
*IndividualAppliedBDTPolicyDataDocumentApi* | [**DeleteIndividualAppliedBdtPolicyData**](docs/IndividualAppliedBDTPolicyDataDocumentApi.md#deleteindividualappliedbdtpolicydata) | **Delete** /application-data/bdtPolicyData/{bdtPolicyId} | Delete an individual Applied BDT Policy Data resource
*IndividualAppliedBDTPolicyDataDocumentApi* | [**UpdateIndividualAppliedBdtPolicyData**](docs/IndividualAppliedBDTPolicyDataDocumentApi.md#updateindividualappliedbdtpolicydata) | **Patch** /application-data/bdtPolicyData/{bdtPolicyId} | Modify part of the properties of an individual Applied BDT Policy Data resource
*IndividualEASDeploymentDataDocumentApi* | [**CreateOrReplaceIndividualEasDeployData**](docs/IndividualEASDeploymentDataDocumentApi.md#createorreplaceindividualeasdeploydata) | **Put** /application-data/eas-deploy-data/{easDeployInfoId} | Create or update an individual EAS Deployment Data resource
*IndividualEASDeploymentDataDocumentApi* | [**ReadIndividualEasDeployData**](docs/IndividualEASDeploymentDataDocumentApi.md#readindividualeasdeploydata) | **Get** /application-data/eas-deploy-data/{easDeployInfoId} | Retrieve an individual EAS Deployment Data resource
*IndividualEasDeploymentDataDocumentApi* | [**DeleteIndividualEasDeployData**](docs/IndividualEasDeploymentDataDocumentApi.md#deleteindividualeasdeploydata) | **Delete** /application-data/eas-deploy-data/{easDeployInfoId} | Delete an individual EAS Deployment Data resource
*IndividualIPTVConfigurationDataDocumentApi* | [**CreateOrReplaceIndividualIPTVConfigurationData**](docs/IndividualIPTVConfigurationDataDocumentApi.md#createorreplaceindividualiptvconfigurationdata) | **Put** /application-data/iptvConfigData/{configurationId} | Create or update an individual IPTV configuration resource
*IndividualIPTVConfigurationDataDocumentApi* | [**DeleteIndividualIPTVConfigurationData**](docs/IndividualIPTVConfigurationDataDocumentApi.md#deleteindividualiptvconfigurationdata) | **Delete** /application-data/iptvConfigData/{configurationId} | Delete an individual IPTV configuration resource
*IndividualIPTVConfigurationDataDocumentApi* | [**PartialReplaceIndividualIPTVConfigurationData**](docs/IndividualIPTVConfigurationDataDocumentApi.md#partialreplaceindividualiptvconfigurationdata) | **Patch** /application-data/iptvConfigData/{configurationId} | Partial update an individual IPTV configuration resource
*IndividualInfluenceDataDocumentApi* | [**CreateOrReplaceIndividualInfluenceData**](docs/IndividualInfluenceDataDocumentApi.md#createorreplaceindividualinfluencedata) | **Put** /application-data/influenceData/{influenceId} | Create or update an individual Influence Data resource
*IndividualInfluenceDataDocumentApi* | [**DeleteIndividualInfluenceData**](docs/IndividualInfluenceDataDocumentApi.md#deleteindividualinfluencedata) | **Delete** /application-data/influenceData/{influenceId} | Delete an individual Influence Data resource
*IndividualInfluenceDataDocumentApi* | [**UpdateIndividualInfluenceData**](docs/IndividualInfluenceDataDocumentApi.md#updateindividualinfluencedata) | **Patch** /application-data/influenceData/{influenceId} | Modify part of the properties of an individual Influence Data resource
*IndividualInfluenceDataSubscriptionDocumentApi* | [**DeleteIndividualInfluenceDataSubscription**](docs/IndividualInfluenceDataSubscriptionDocumentApi.md#deleteindividualinfluencedatasubscription) | **Delete** /application-data/influenceData/subs-to-notify/{subscriptionId} | Delete an individual Influence Data Subscription resource
*IndividualInfluenceDataSubscriptionDocumentApi* | [**ReadIndividualInfluenceDataSubscription**](docs/IndividualInfluenceDataSubscriptionDocumentApi.md#readindividualinfluencedatasubscription) | **Get** /application-data/influenceData/subs-to-notify/{subscriptionId} | Get an existing individual Influence Data Subscription resource
*IndividualInfluenceDataSubscriptionDocumentApi* | [**ReplaceIndividualInfluenceDataSubscription**](docs/IndividualInfluenceDataSubscriptionDocumentApi.md#replaceindividualinfluencedatasubscription) | **Put** /application-data/influenceData/subs-to-notify/{subscriptionId} | Modify an existing individual Influence Data Subscription resource
*IndividualPFDDataDocumentApi* | [**CreateOrReplaceIndividualPFDData**](docs/IndividualPFDDataDocumentApi.md#createorreplaceindividualpfddata) | **Put** /application-data/pfds/{appId} | Create or update the corresponding PFDs for the specified application identifier
*IndividualPFDDataDocumentApi* | [**DeleteIndividualPFDData**](docs/IndividualPFDDataDocumentApi.md#deleteindividualpfddata) | **Delete** /application-data/pfds/{appId} | Delete the corresponding PFDs of the specified application identifier
*IndividualPFDDataDocumentApi* | [**ReadIndividualPFDData**](docs/IndividualPFDDataDocumentApi.md#readindividualpfddata) | **Get** /application-data/pfds/{appId} | Retrieve the corresponding PFDs of the specified application identifier
*IndividualServiceParameterDataDocumentApi* | [**CreateOrReplaceServiceParameterData**](docs/IndividualServiceParameterDataDocumentApi.md#createorreplaceserviceparameterdata) | **Put** /application-data/serviceParamData/{serviceParamId} | Create or update an individual Service Parameter Data resource
*IndividualServiceParameterDataDocumentApi* | [**DeleteIndividualServiceParameterData**](docs/IndividualServiceParameterDataDocumentApi.md#deleteindividualserviceparameterdata) | **Delete** /application-data/serviceParamData/{serviceParamId} | Delete an individual Service Parameter Data resource
*IndividualServiceParameterDataDocumentApi* | [**UpdateIndividualServiceParameterData**](docs/IndividualServiceParameterDataDocumentApi.md#updateindividualserviceparameterdata) | **Patch** /application-data/serviceParamData/{serviceParamId} | Modify part of the properties of an individual Service Parameter Data resource
*InfluenceDataStoreApi* | [**ReadInfluenceData**](docs/InfluenceDataStoreApi.md#readinfluencedata) | **Get** /application-data/influenceData | Retrieve Traffic Influence Data
*InfluenceDataSubscriptionsCollectionApi* | [**CreateIndividualInfluenceDataSubscription**](docs/InfluenceDataSubscriptionsCollectionApi.md#createindividualinfluencedatasubscription) | **Post** /application-data/influenceData/subs-to-notify | Create a new Individual Influence Data Subscription resource
*InfluenceDataSubscriptionsCollectionApi* | [**ReadInfluenceDataSubscriptions**](docs/InfluenceDataSubscriptionsCollectionApi.md#readinfluencedatasubscriptions) | **Get** /application-data/influenceData/subs-to-notify | Read Influence Data Subscriptions
*PFDDataStoreApi* | [**ReadPFDData**](docs/PFDDataStoreApi.md#readpfddata) | **Get** /application-data/pfds | Retrieve PFDs for application identifier(s)
*ServiceParameterDataStoreApi* | [**ReadServiceParameterData**](docs/ServiceParameterDataStoreApi.md#readserviceparameterdata) | **Get** /application-data/serviceParamData | Retrieve Service Parameter Data


## Documentation For Models

 - [AccessRightStatus](docs/AccessRightStatus.md)
 - [AccessTokenErr](docs/AccessTokenErr.md)
 - [AccessTokenReq](docs/AccessTokenReq.md)
 - [AmInfluData](docs/AmInfluData.md)
 - [AmInfluDataPatch](docs/AmInfluDataPatch.md)
 - [AmInfluEvent](docs/AmInfluEvent.md)
 - [AppDescriptor](docs/AppDescriptor.md)
 - [ApplicationDataChangeNotif](docs/ApplicationDataChangeNotif.md)
 - [ApplicationDataInfluenceDataSubsToNotifyPostRequestInner](docs/ApplicationDataInfluenceDataSubsToNotifyPostRequestInner.md)
 - [ApplicationDataSubs](docs/ApplicationDataSubs.md)
 - [BdtPolicyData](docs/BdtPolicyData.md)
 - [BdtPolicyDataPatch](docs/BdtPolicyDataPatch.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [ConnectionCapabilities](docs/ConnectionCapabilities.md)
 - [DataFilter](docs/DataFilter.md)
 - [DataInd](docs/DataInd.md)
 - [DnaiChangeType](docs/DnaiChangeType.md)
 - [DnaiInformation](docs/DnaiInformation.md)
 - [DnnSnssaiInformation](docs/DnnSnssaiInformation.md)
 - [DnsServerIdentifier](docs/DnsServerIdentifier.md)
 - [DomainNameProtocol](docs/DomainNameProtocol.md)
 - [EasDeployInfoData](docs/EasDeployInfoData.md)
 - [Ecgi](docs/Ecgi.md)
 - [EllipsoidArc](docs/EllipsoidArc.md)
 - [EllipsoidArcAllOf](docs/EllipsoidArcAllOf.md)
 - [EthFlowDescription](docs/EthFlowDescription.md)
 - [Event](docs/Event.md)
 - [FlowDirection](docs/FlowDirection.md)
 - [FlowInfo](docs/FlowInfo.md)
 - [FqdnPatternMatchingRule](docs/FqdnPatternMatchingRule.md)
 - [GADShape](docs/GADShape.md)
 - [GNbId](docs/GNbId.md)
 - [GeographicArea](docs/GeographicArea.md)
 - [GeographicalArea](docs/GeographicalArea.md)
 - [GeographicalCoordinates](docs/GeographicalCoordinates.md)
 - [GlobalRanNodeId](docs/GlobalRanNodeId.md)
 - [InvalidParam](docs/InvalidParam.md)
 - [IpAddr](docs/IpAddr.md)
 - [IptvConfigData](docs/IptvConfigData.md)
 - [IptvConfigDataPatch](docs/IptvConfigDataPatch.md)
 - [Ipv6Addr](docs/Ipv6Addr.md)
 - [Ipv6Prefix](docs/Ipv6Prefix.md)
 - [Local2dPointUncertaintyEllipse](docs/Local2dPointUncertaintyEllipse.md)
 - [Local2dPointUncertaintyEllipseAllOf](docs/Local2dPointUncertaintyEllipseAllOf.md)
 - [Local3dPointUncertaintyEllipsoid](docs/Local3dPointUncertaintyEllipsoid.md)
 - [Local3dPointUncertaintyEllipsoidAllOf](docs/Local3dPointUncertaintyEllipsoidAllOf.md)
 - [LocalOrigin](docs/LocalOrigin.md)
 - [MatchingOperator](docs/MatchingOperator.md)
 - [MulticastAccessControl](docs/MulticastAccessControl.md)
 - [NFType](docs/NFType.md)
 - [Ncgi](docs/Ncgi.md)
 - [NetworkAreaInfo](docs/NetworkAreaInfo.md)
 - [PfdChangeNotification](docs/PfdChangeNotification.md)
 - [PfdContent](docs/PfdContent.md)
 - [PfdDataForAppExt](docs/PfdDataForAppExt.md)
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
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RelativeCartesianLocation](docs/RelativeCartesianLocation.md)
 - [RouteInformation](docs/RouteInformation.md)
 - [RouteSelectionParameterSet](docs/RouteSelectionParameterSet.md)
 - [RouteToLocation](docs/RouteToLocation.md)
 - [ServiceAreaCoverageInfo](docs/ServiceAreaCoverageInfo.md)
 - [ServiceParameterData](docs/ServiceParameterData.md)
 - [ServiceParameterDataPatch](docs/ServiceParameterDataPatch.md)
 - [Snssai](docs/Snssai.md)
 - [StringMatchingCondition](docs/StringMatchingCondition.md)
 - [StringMatchingRule](docs/StringMatchingRule.md)
 - [SubscribedEvent](docs/SubscribedEvent.md)
 - [SupportedGADShapes](docs/SupportedGADShapes.md)
 - [Tai](docs/Tai.md)
 - [TemporalValidity](docs/TemporalValidity.md)
 - [TrafficDescriptorComponents](docs/TrafficDescriptorComponents.md)
 - [TrafficInfluData](docs/TrafficInfluData.md)
 - [TrafficInfluDataNotif](docs/TrafficInfluDataNotif.md)
 - [TrafficInfluDataPatch](docs/TrafficInfluDataPatch.md)
 - [TrafficInfluSub](docs/TrafficInfluSub.md)
 - [UncertaintyEllipse](docs/UncertaintyEllipse.md)
 - [UncertaintyEllipsoid](docs/UncertaintyEllipsoid.md)
 - [UrspRuleRequest](docs/UrspRuleRequest.md)


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



