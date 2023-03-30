/*
Nrouter_SMService Service API

SMS Router SMService.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nrouter_SMService

import (
	"encoding/json"
)

// checks if the CreatedRoutingData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedRoutingData{}

// CreatedRoutingData Information used for receiving the MT SMS.
type CreatedRoutingData struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	RouterIpv4 *string `json:"routerIpv4,omitempty"`
	RouterIpv6 *Ipv6Addr `json:"routerIpv6,omitempty"`
	// Fully Qualified Domain Name
	RouterFqdn *string `json:"routerFqdn,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewCreatedRoutingData instantiates a new CreatedRoutingData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedRoutingData() *CreatedRoutingData {
	this := CreatedRoutingData{}
	return &this
}

// NewCreatedRoutingDataWithDefaults instantiates a new CreatedRoutingData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedRoutingDataWithDefaults() *CreatedRoutingData {
	this := CreatedRoutingData{}
	return &this
}

// GetRouterIpv4 returns the RouterIpv4 field value if set, zero value otherwise.
func (o *CreatedRoutingData) GetRouterIpv4() string {
	if o == nil || IsNil(o.RouterIpv4) {
		var ret string
		return ret
	}
	return *o.RouterIpv4
}

// GetRouterIpv4Ok returns a tuple with the RouterIpv4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedRoutingData) GetRouterIpv4Ok() (*string, bool) {
	if o == nil || IsNil(o.RouterIpv4) {
		return nil, false
	}
	return o.RouterIpv4, true
}

// HasRouterIpv4 returns a boolean if a field has been set.
func (o *CreatedRoutingData) HasRouterIpv4() bool {
	if o != nil && !IsNil(o.RouterIpv4) {
		return true
	}

	return false
}

// SetRouterIpv4 gets a reference to the given string and assigns it to the RouterIpv4 field.
func (o *CreatedRoutingData) SetRouterIpv4(v string) {
	o.RouterIpv4 = &v
}

// GetRouterIpv6 returns the RouterIpv6 field value if set, zero value otherwise.
func (o *CreatedRoutingData) GetRouterIpv6() Ipv6Addr {
	if o == nil || IsNil(o.RouterIpv6) {
		var ret Ipv6Addr
		return ret
	}
	return *o.RouterIpv6
}

// GetRouterIpv6Ok returns a tuple with the RouterIpv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedRoutingData) GetRouterIpv6Ok() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.RouterIpv6) {
		return nil, false
	}
	return o.RouterIpv6, true
}

// HasRouterIpv6 returns a boolean if a field has been set.
func (o *CreatedRoutingData) HasRouterIpv6() bool {
	if o != nil && !IsNil(o.RouterIpv6) {
		return true
	}

	return false
}

// SetRouterIpv6 gets a reference to the given Ipv6Addr and assigns it to the RouterIpv6 field.
func (o *CreatedRoutingData) SetRouterIpv6(v Ipv6Addr) {
	o.RouterIpv6 = &v
}

// GetRouterFqdn returns the RouterFqdn field value if set, zero value otherwise.
func (o *CreatedRoutingData) GetRouterFqdn() string {
	if o == nil || IsNil(o.RouterFqdn) {
		var ret string
		return ret
	}
	return *o.RouterFqdn
}

// GetRouterFqdnOk returns a tuple with the RouterFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedRoutingData) GetRouterFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.RouterFqdn) {
		return nil, false
	}
	return o.RouterFqdn, true
}

// HasRouterFqdn returns a boolean if a field has been set.
func (o *CreatedRoutingData) HasRouterFqdn() bool {
	if o != nil && !IsNil(o.RouterFqdn) {
		return true
	}

	return false
}

// SetRouterFqdn gets a reference to the given string and assigns it to the RouterFqdn field.
func (o *CreatedRoutingData) SetRouterFqdn(v string) {
	o.RouterFqdn = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *CreatedRoutingData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreatedRoutingData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *CreatedRoutingData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *CreatedRoutingData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o CreatedRoutingData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedRoutingData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RouterIpv4) {
		toSerialize["routerIpv4"] = o.RouterIpv4
	}
	if !IsNil(o.RouterIpv6) {
		toSerialize["routerIpv6"] = o.RouterIpv6
	}
	if !IsNil(o.RouterFqdn) {
		toSerialize["routerFqdn"] = o.RouterFqdn
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableCreatedRoutingData struct {
	value *CreatedRoutingData
	isSet bool
}

func (v NullableCreatedRoutingData) Get() *CreatedRoutingData {
	return v.value
}

func (v *NullableCreatedRoutingData) Set(val *CreatedRoutingData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedRoutingData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedRoutingData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedRoutingData(val *CreatedRoutingData) *NullableCreatedRoutingData {
	return &NullableCreatedRoutingData{value: val, isSet: true}
}

func (v NullableCreatedRoutingData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedRoutingData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

