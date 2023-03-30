# Go API client for openapi_MBSUserDataIngestSession

API for MBS User Data Ingest Session.  
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
import openapi_MBSUserDataIngestSession "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_MBSUserDataIngestSession"
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
ctx := context.WithValue(context.Background(), openapi_MBSUserDataIngestSession.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_MBSUserDataIngestSession.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_MBSUserDataIngestSession.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_MBSUserDataIngestSession.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/3gpp-mbs-ud-ingest/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*IndividualMBSUserDataIngestSessionDocumentApi* | [**DeleteIndivMBSUserDataIngestSession**](docs/IndividualMBSUserDataIngestSessionDocumentApi.md#deleteindivmbsuserdataingestsession) | **Delete** /sessions/{sessionId} | Deletes an existing Individual MBS User Data Ingest Session resource.
*IndividualMBSUserDataIngestSessionDocumentApi* | [**ModifyIndivMBSUserDataIngestSession**](docs/IndividualMBSUserDataIngestSessionDocumentApi.md#modifyindivmbsuserdataingestsession) | **Patch** /sessions/{sessionId} | Request the modification of an existing Individual MBS User Data Ingest Session resource.
*IndividualMBSUserDataIngestSessionDocumentApi* | [**RetrieveIndivMBSUserDataIngestSession**](docs/IndividualMBSUserDataIngestSessionDocumentApi.md#retrieveindivmbsuserdataingestsession) | **Get** /sessions/{sessionId} | Retrieve an existing Individual MBS User Data Ingest Session resource.
*IndividualMBSUserDataIngestSessionDocumentApi* | [**UpdateIndivMBSUserDataIngestSession**](docs/IndividualMBSUserDataIngestSessionDocumentApi.md#updateindivmbsuserdataingestsession) | **Put** /sessions/{sessionId} | Request the update of an existing Individual MBS User Data Ingest Session resource.
*IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi* | [**DeleteIndMBSUserDataIngStatSubsc**](docs/IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.md#deleteindmbsuserdataingstatsubsc) | **Delete** /status-subscriptions/{subscriptionId} | Deletes an existing Individual MBS User Data Ingest Session Status Subscription resource.
*IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi* | [**ModifyIndMBSUserDataIngStatSubsc**](docs/IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.md#modifyindmbsuserdataingstatsubsc) | **Patch** /status-subscriptions/{subscriptionId} | Request the modification of an existing Individual MBS User Data Ingest Session Status Subscription resource.
*IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi* | [**RetrieveIndMBSUserDataIngStatSubsc**](docs/IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.md#retrieveindmbsuserdataingstatsubsc) | **Get** /status-subscriptions/{subscriptionId} | Retrieve an existing Individual MBS User Data Ingest Session Status Subscription resource.
*IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi* | [**UpdateIndMBSUserDataIngStatSubsc**](docs/IndividualMBSUserDataIngestSessionStatusSubscriptionDocumentApi.md#updateindmbsuserdataingstatsubsc) | **Put** /status-subscriptions/{subscriptionId} | Request the update of an existing Individual MBS User Data Ingest Session Status Subscription resource.
*MBSUserDataIngestSessionStatusSubscriptionsCollectionApi* | [**CreateMBSUserDataIngStatSubsc**](docs/MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.md#creatembsuserdataingstatsubsc) | **Post** /status-subscriptions | Creates a new Individual MBS User Data Ingest Session Status Subscription resource.
*MBSUserDataIngestSessionStatusSubscriptionsCollectionApi* | [**RetrieveMBSUserDataIngStatSubscs**](docs/MBSUserDataIngestSessionStatusSubscriptionsCollectionApi.md#retrievembsuserdataingstatsubscs) | **Get** /status-subscriptions | Retrieve all the active MBS User Data Ingest Session Status Subscriptions resources managed by the NEF.
*MBSUserDataIngestSessionsCollectionApi* | [**CreateMBSUserDataIngestSession**](docs/MBSUserDataIngestSessionsCollectionApi.md#creatembsuserdataingestsession) | **Post** /sessions | Request the creation of a new Individual MBS User Data Ingest Session resource.
*MBSUserDataIngestSessionsCollectionApi* | [**RetrieveMBSUserDataIngestSessions**](docs/MBSUserDataIngestSessionsCollectionApi.md#retrievembsuserdataingestsessions) | **Get** /sessions | Retrieve all the active MBS User Data Ingest Sessions managed by the NEF.


## Documentation For Models

 - [AddFecParams](docs/AddFecParams.md)
 - [Arp](docs/Arp.md)
 - [CivicAddress](docs/CivicAddress.md)
 - [DistSessionState](docs/DistSessionState.md)
 - [DistributionMethod](docs/DistributionMethod.md)
 - [EllipsoidArc](docs/EllipsoidArc.md)
 - [EllipsoidArcAllOf](docs/EllipsoidArcAllOf.md)
 - [Event](docs/Event.md)
 - [EventNotification](docs/EventNotification.md)
 - [ExtSsm](docs/ExtSsm.md)
 - [ExternalMbsServiceArea](docs/ExternalMbsServiceArea.md)
 - [FECConfig](docs/FECConfig.md)
 - [GADShape](docs/GADShape.md)
 - [GeographicArea](docs/GeographicArea.md)
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
 - [MBSDistSessionAnmt](docs/MBSDistSessionAnmt.md)
 - [MBSDistributionSessionInfo](docs/MBSDistributionSessionInfo.md)
 - [MBSUserDataIngSession](docs/MBSUserDataIngSession.md)
 - [MBSUserDataIngSessionPatch](docs/MBSUserDataIngSessionPatch.md)
 - [MBSUserDataIngStatNotif](docs/MBSUserDataIngStatNotif.md)
 - [MBSUserDataIngStatSubsc](docs/MBSUserDataIngStatSubsc.md)
 - [MBSUserDataIngStatSubscPatch](docs/MBSUserDataIngStatSubscPatch.md)
 - [MBSUserServAnmt](docs/MBSUserServAnmt.md)
 - [MbStfIngestAddr](docs/MbStfIngestAddr.md)
 - [MbStfIngestAddrAfEgressTunAddr](docs/MbStfIngestAddrAfEgressTunAddr.md)
 - [MbStfIngestAddrAfSsm](docs/MbStfIngestAddrAfSsm.md)
 - [MbsMediaComp](docs/MbsMediaComp.md)
 - [MbsMediaCompRm](docs/MbsMediaCompRm.md)
 - [MbsMediaInfo](docs/MbsMediaInfo.md)
 - [MbsQoSReq](docs/MbsQoSReq.md)
 - [MbsServiceArea](docs/MbsServiceArea.md)
 - [MbsServiceInfo](docs/MbsServiceInfo.md)
 - [MbsSessionId](docs/MbsSessionId.md)
 - [MediaType](docs/MediaType.md)
 - [Ncgi](docs/Ncgi.md)
 - [NcgiTai](docs/NcgiTai.md)
 - [NullValue](docs/NullValue.md)
 - [ObjAcquisitionMethod](docs/ObjAcquisitionMethod.md)
 - [ObjDistributionOperatingMode](docs/ObjDistributionOperatingMode.md)
 - [ObjectDistMethAnmtInfo](docs/ObjectDistMethAnmtInfo.md)
 - [ObjectDistrMethInfo](docs/ObjectDistrMethInfo.md)
 - [PacketDistrMethInfo](docs/PacketDistrMethInfo.md)
 - [PktDistributionOperatingMode](docs/PktDistributionOperatingMode.md)
 - [PktIngestMethod](docs/PktIngestMethod.md)
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
 - [PreemptionCapability](docs/PreemptionCapability.md)
 - [PreemptionVulnerability](docs/PreemptionVulnerability.md)
 - [ProblemDetails](docs/ProblemDetails.md)
 - [RelativeCartesianLocation](docs/RelativeCartesianLocation.md)
 - [ReservPriority](docs/ReservPriority.md)
 - [ServiceNameDescription](docs/ServiceNameDescription.md)
 - [Ssm](docs/Ssm.md)
 - [SubscribedEvent](docs/SubscribedEvent.md)
 - [SupportedGADShapes](docs/SupportedGADShapes.md)
 - [Tai](docs/Tai.md)
 - [TimeWindow](docs/TimeWindow.md)
 - [Tmgi](docs/Tmgi.md)
 - [TunnelAddress](docs/TunnelAddress.md)
 - [UncertaintyEllipse](docs/UncertaintyEllipse.md)
 - [UncertaintyEllipsoid](docs/UncertaintyEllipsoid.md)


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


