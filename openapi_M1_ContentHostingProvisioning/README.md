# Go API client for openapi_M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API
© 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).
All rights reserved.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.1.0
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
import openapi_M1_ContentHostingProvisioning "github.com/GIT_USER_ID/GIT_REPO_ID/openapi_M1_ContentHostingProvisioning"
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
ctx := context.WithValue(context.Background(), openapi_M1_ContentHostingProvisioning.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi_M1_ContentHostingProvisioning.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), openapi_M1_ContentHostingProvisioning.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi_M1_ContentHostingProvisioning.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://example.com/3gpp-m1/v2*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**CreateContentHostingConfiguration**](docs/DefaultApi.md#createcontenthostingconfiguration) | **Post** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Create (and optionally upload) the Content Hosting Configuration for the specified Provisioning Session
*DefaultApi* | [**DestroyContentHostingConfiguration**](docs/DefaultApi.md#destroycontenthostingconfiguration) | **Delete** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Destroy the current Content Hosting Configuration of the specified Provisioning Session
*DefaultApi* | [**PatchContentHostingConfiguration**](docs/DefaultApi.md#patchcontenthostingconfiguration) | **Patch** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Patch the Content Hosting Configuration for the specified Provisioning Session
*DefaultApi* | [**PurgeContentHostingCache**](docs/DefaultApi.md#purgecontenthostingcache) | **Post** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration/purge | Purge the content of the cache for the Content Hosting Configuration of the specified Provisioning Session
*DefaultApi* | [**RetrieveContentHostingConfiguration**](docs/DefaultApi.md#retrievecontenthostingconfiguration) | **Get** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Retrieve the Content Hosting Configuration of the specified Provisioning Session
*DefaultApi* | [**UpdateContentHostingConfiguration**](docs/DefaultApi.md#updatecontenthostingconfiguration) | **Put** /provisioning-sessions/{provisioningSessionId}/content-hosting-configuration | Update the Content Hosting Configuration for the specified Provisioning Session


## Documentation For Models

 - [CachingConfiguration](docs/CachingConfiguration.md)
 - [CachingConfigurationCachingDirectives](docs/CachingConfigurationCachingDirectives.md)
 - [ContentHostingConfiguration](docs/ContentHostingConfiguration.md)
 - [DistributionConfiguration](docs/DistributionConfiguration.md)
 - [DistributionConfigurationGeoFencing](docs/DistributionConfigurationGeoFencing.md)
 - [DistributionConfigurationSupplementaryDistributionNetworksInner](docs/DistributionConfigurationSupplementaryDistributionNetworksInner.md)
 - [DistributionConfigurationUrlSignature](docs/DistributionConfigurationUrlSignature.md)
 - [DistributionMode](docs/DistributionMode.md)
 - [DistributionNetworkType](docs/DistributionNetworkType.md)
 - [IngestConfiguration](docs/IngestConfiguration.md)
 - [PathRewriteRule](docs/PathRewriteRule.md)


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


