/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the RedirectInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RedirectInformation{}

// RedirectInformation Contains the redirect information.
type RedirectInformation struct {
	// Indicates the redirect is enable.
	RedirectEnabled     *bool                `json:"redirectEnabled,omitempty"`
	RedirectAddressType *RedirectAddressType `json:"redirectAddressType,omitempty"`
	// Indicates the address of the redirect server. If \"redirectAddressType\" attribute indicates the IPV4_ADDR, the encoding is the same as the Ipv4Addr data type defined in 3GPP TS 29.571.If \"redirectAddressType\" attribute indicates the IPV6_ADDR, the encoding is the same as the Ipv6Addr data type defined in 3GPP TS 29.571.If \"redirectAddressType\" attribute indicates the URL or SIP_URI, the encoding is the same as the Uri data type defined in 3GPP TS 29.571.
	RedirectServerAddress *string `json:"redirectServerAddress,omitempty"`
}

// NewRedirectInformation instantiates a new RedirectInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRedirectInformation() *RedirectInformation {
	this := RedirectInformation{}
	return &this
}

// NewRedirectInformationWithDefaults instantiates a new RedirectInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRedirectInformationWithDefaults() *RedirectInformation {
	this := RedirectInformation{}
	return &this
}

// GetRedirectEnabled returns the RedirectEnabled field value if set, zero value otherwise.
func (o *RedirectInformation) GetRedirectEnabled() bool {
	if o == nil || IsNil(o.RedirectEnabled) {
		var ret bool
		return ret
	}
	return *o.RedirectEnabled
}

// GetRedirectEnabledOk returns a tuple with the RedirectEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectInformation) GetRedirectEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.RedirectEnabled) {
		return nil, false
	}
	return o.RedirectEnabled, true
}

// HasRedirectEnabled returns a boolean if a field has been set.
func (o *RedirectInformation) HasRedirectEnabled() bool {
	if o != nil && !IsNil(o.RedirectEnabled) {
		return true
	}

	return false
}

// SetRedirectEnabled gets a reference to the given bool and assigns it to the RedirectEnabled field.
func (o *RedirectInformation) SetRedirectEnabled(v bool) {
	o.RedirectEnabled = &v
}

// GetRedirectAddressType returns the RedirectAddressType field value if set, zero value otherwise.
func (o *RedirectInformation) GetRedirectAddressType() RedirectAddressType {
	if o == nil || IsNil(o.RedirectAddressType) {
		var ret RedirectAddressType
		return ret
	}
	return *o.RedirectAddressType
}

// GetRedirectAddressTypeOk returns a tuple with the RedirectAddressType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectInformation) GetRedirectAddressTypeOk() (*RedirectAddressType, bool) {
	if o == nil || IsNil(o.RedirectAddressType) {
		return nil, false
	}
	return o.RedirectAddressType, true
}

// HasRedirectAddressType returns a boolean if a field has been set.
func (o *RedirectInformation) HasRedirectAddressType() bool {
	if o != nil && !IsNil(o.RedirectAddressType) {
		return true
	}

	return false
}

// SetRedirectAddressType gets a reference to the given RedirectAddressType and assigns it to the RedirectAddressType field.
func (o *RedirectInformation) SetRedirectAddressType(v RedirectAddressType) {
	o.RedirectAddressType = &v
}

// GetRedirectServerAddress returns the RedirectServerAddress field value if set, zero value otherwise.
func (o *RedirectInformation) GetRedirectServerAddress() string {
	if o == nil || IsNil(o.RedirectServerAddress) {
		var ret string
		return ret
	}
	return *o.RedirectServerAddress
}

// GetRedirectServerAddressOk returns a tuple with the RedirectServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RedirectInformation) GetRedirectServerAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RedirectServerAddress) {
		return nil, false
	}
	return o.RedirectServerAddress, true
}

// HasRedirectServerAddress returns a boolean if a field has been set.
func (o *RedirectInformation) HasRedirectServerAddress() bool {
	if o != nil && !IsNil(o.RedirectServerAddress) {
		return true
	}

	return false
}

// SetRedirectServerAddress gets a reference to the given string and assigns it to the RedirectServerAddress field.
func (o *RedirectInformation) SetRedirectServerAddress(v string) {
	o.RedirectServerAddress = &v
}

func (o RedirectInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RedirectInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RedirectEnabled) {
		toSerialize["redirectEnabled"] = o.RedirectEnabled
	}
	if !IsNil(o.RedirectAddressType) {
		toSerialize["redirectAddressType"] = o.RedirectAddressType
	}
	if !IsNil(o.RedirectServerAddress) {
		toSerialize["redirectServerAddress"] = o.RedirectServerAddress
	}
	return toSerialize, nil
}

type NullableRedirectInformation struct {
	value *RedirectInformation
	isSet bool
}

func (v NullableRedirectInformation) Get() *RedirectInformation {
	return v.value
}

func (v *NullableRedirectInformation) Set(val *RedirectInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectInformation(val *RedirectInformation) *NullableRedirectInformation {
	return &NullableRedirectInformation{value: val, isSet: true}
}

func (v NullableRedirectInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
