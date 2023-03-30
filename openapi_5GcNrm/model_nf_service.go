/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the NFService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NFService{}

// NFService NF Service is defined in TS 29.510
type NFService struct {
	ServiceInstanceId *string `json:"serviceInstanceId,omitempty"`
	ServiceName *string `json:"serviceName,omitempty"`
	Version *string `json:"version,omitempty"`
	Schema *string `json:"schema,omitempty"`
	Fqdn *string `json:"fqdn,omitempty"`
	InterPlmnFqdn *string `json:"interPlmnFqdn,omitempty"`
	IpEndPoints []IpEndPoint `json:"ipEndPoints,omitempty"`
	ApiPrfix *string `json:"apiPrfix,omitempty"`
	AllowedPlmns *PlmnId `json:"allowedPlmns,omitempty"`
	AllowedNfTypes []NFType `json:"allowedNfTypes,omitempty"`
	AllowedNssais []Snssai `json:"allowedNssais,omitempty"`
}

// NewNFService instantiates a new NFService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNFService() *NFService {
	this := NFService{}
	return &this
}

// NewNFServiceWithDefaults instantiates a new NFService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNFServiceWithDefaults() *NFService {
	this := NFService{}
	return &this
}

// GetServiceInstanceId returns the ServiceInstanceId field value if set, zero value otherwise.
func (o *NFService) GetServiceInstanceId() string {
	if o == nil || IsNil(o.ServiceInstanceId) {
		var ret string
		return ret
	}
	return *o.ServiceInstanceId
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceInstanceId) {
		return nil, false
	}
	return o.ServiceInstanceId, true
}

// HasServiceInstanceId returns a boolean if a field has been set.
func (o *NFService) HasServiceInstanceId() bool {
	if o != nil && !IsNil(o.ServiceInstanceId) {
		return true
	}

	return false
}

// SetServiceInstanceId gets a reference to the given string and assigns it to the ServiceInstanceId field.
func (o *NFService) SetServiceInstanceId(v string) {
	o.ServiceInstanceId = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *NFService) GetServiceName() string {
	if o == nil || IsNil(o.ServiceName) {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *NFService) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *NFService) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *NFService) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *NFService) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *NFService) SetVersion(v string) {
	o.Version = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *NFService) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *NFService) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *NFService) SetSchema(v string) {
	o.Schema = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *NFService) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *NFService) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *NFService) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetInterPlmnFqdn returns the InterPlmnFqdn field value if set, zero value otherwise.
func (o *NFService) GetInterPlmnFqdn() string {
	if o == nil || IsNil(o.InterPlmnFqdn) {
		var ret string
		return ret
	}
	return *o.InterPlmnFqdn
}

// GetInterPlmnFqdnOk returns a tuple with the InterPlmnFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetInterPlmnFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.InterPlmnFqdn) {
		return nil, false
	}
	return o.InterPlmnFqdn, true
}

// HasInterPlmnFqdn returns a boolean if a field has been set.
func (o *NFService) HasInterPlmnFqdn() bool {
	if o != nil && !IsNil(o.InterPlmnFqdn) {
		return true
	}

	return false
}

// SetInterPlmnFqdn gets a reference to the given string and assigns it to the InterPlmnFqdn field.
func (o *NFService) SetInterPlmnFqdn(v string) {
	o.InterPlmnFqdn = &v
}

// GetIpEndPoints returns the IpEndPoints field value if set, zero value otherwise.
func (o *NFService) GetIpEndPoints() []IpEndPoint {
	if o == nil || IsNil(o.IpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.IpEndPoints
}

// GetIpEndPointsOk returns a tuple with the IpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || IsNil(o.IpEndPoints) {
		return nil, false
	}
	return o.IpEndPoints, true
}

// HasIpEndPoints returns a boolean if a field has been set.
func (o *NFService) HasIpEndPoints() bool {
	if o != nil && !IsNil(o.IpEndPoints) {
		return true
	}

	return false
}

// SetIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the IpEndPoints field.
func (o *NFService) SetIpEndPoints(v []IpEndPoint) {
	o.IpEndPoints = v
}

// GetApiPrfix returns the ApiPrfix field value if set, zero value otherwise.
func (o *NFService) GetApiPrfix() string {
	if o == nil || IsNil(o.ApiPrfix) {
		var ret string
		return ret
	}
	return *o.ApiPrfix
}

// GetApiPrfixOk returns a tuple with the ApiPrfix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetApiPrfixOk() (*string, bool) {
	if o == nil || IsNil(o.ApiPrfix) {
		return nil, false
	}
	return o.ApiPrfix, true
}

// HasApiPrfix returns a boolean if a field has been set.
func (o *NFService) HasApiPrfix() bool {
	if o != nil && !IsNil(o.ApiPrfix) {
		return true
	}

	return false
}

// SetApiPrfix gets a reference to the given string and assigns it to the ApiPrfix field.
func (o *NFService) SetApiPrfix(v string) {
	o.ApiPrfix = &v
}

// GetAllowedPlmns returns the AllowedPlmns field value if set, zero value otherwise.
func (o *NFService) GetAllowedPlmns() PlmnId {
	if o == nil || IsNil(o.AllowedPlmns) {
		var ret PlmnId
		return ret
	}
	return *o.AllowedPlmns
}

// GetAllowedPlmnsOk returns a tuple with the AllowedPlmns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedPlmnsOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.AllowedPlmns) {
		return nil, false
	}
	return o.AllowedPlmns, true
}

// HasAllowedPlmns returns a boolean if a field has been set.
func (o *NFService) HasAllowedPlmns() bool {
	if o != nil && !IsNil(o.AllowedPlmns) {
		return true
	}

	return false
}

// SetAllowedPlmns gets a reference to the given PlmnId and assigns it to the AllowedPlmns field.
func (o *NFService) SetAllowedPlmns(v PlmnId) {
	o.AllowedPlmns = &v
}

// GetAllowedNfTypes returns the AllowedNfTypes field value if set, zero value otherwise.
func (o *NFService) GetAllowedNfTypes() []NFType {
	if o == nil || IsNil(o.AllowedNfTypes) {
		var ret []NFType
		return ret
	}
	return o.AllowedNfTypes
}

// GetAllowedNfTypesOk returns a tuple with the AllowedNfTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedNfTypesOk() ([]NFType, bool) {
	if o == nil || IsNil(o.AllowedNfTypes) {
		return nil, false
	}
	return o.AllowedNfTypes, true
}

// HasAllowedNfTypes returns a boolean if a field has been set.
func (o *NFService) HasAllowedNfTypes() bool {
	if o != nil && !IsNil(o.AllowedNfTypes) {
		return true
	}

	return false
}

// SetAllowedNfTypes gets a reference to the given []NFType and assigns it to the AllowedNfTypes field.
func (o *NFService) SetAllowedNfTypes(v []NFType) {
	o.AllowedNfTypes = v
}

// GetAllowedNssais returns the AllowedNssais field value if set, zero value otherwise.
func (o *NFService) GetAllowedNssais() []Snssai {
	if o == nil || IsNil(o.AllowedNssais) {
		var ret []Snssai
		return ret
	}
	return o.AllowedNssais
}

// GetAllowedNssaisOk returns a tuple with the AllowedNssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NFService) GetAllowedNssaisOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.AllowedNssais) {
		return nil, false
	}
	return o.AllowedNssais, true
}

// HasAllowedNssais returns a boolean if a field has been set.
func (o *NFService) HasAllowedNssais() bool {
	if o != nil && !IsNil(o.AllowedNssais) {
		return true
	}

	return false
}

// SetAllowedNssais gets a reference to the given []Snssai and assigns it to the AllowedNssais field.
func (o *NFService) SetAllowedNssais(v []Snssai) {
	o.AllowedNssais = v
}

func (o NFService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NFService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceInstanceId) {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.InterPlmnFqdn) {
		toSerialize["interPlmnFqdn"] = o.InterPlmnFqdn
	}
	if !IsNil(o.IpEndPoints) {
		toSerialize["ipEndPoints"] = o.IpEndPoints
	}
	if !IsNil(o.ApiPrfix) {
		toSerialize["apiPrfix"] = o.ApiPrfix
	}
	if !IsNil(o.AllowedPlmns) {
		toSerialize["allowedPlmns"] = o.AllowedPlmns
	}
	if !IsNil(o.AllowedNfTypes) {
		toSerialize["allowedNfTypes"] = o.AllowedNfTypes
	}
	if !IsNil(o.AllowedNssais) {
		toSerialize["allowedNssais"] = o.AllowedNssais
	}
	return toSerialize, nil
}

type NullableNFService struct {
	value *NFService
	isSet bool
}

func (v NullableNFService) Get() *NFService {
	return v.value
}

func (v *NullableNFService) Set(val *NFService) {
	v.value = val
	v.isSet = true
}

func (v NullableNFService) IsSet() bool {
	return v.isSet
}

func (v *NullableNFService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFService(val *NFService) *NullableNFService {
	return &NullableNFService{value: val, isSet: true}
}

func (v NullableNFService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

