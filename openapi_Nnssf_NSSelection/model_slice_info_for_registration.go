/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSelection

import (
	"encoding/json"
)

// checks if the SliceInfoForRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SliceInfoForRegistration{}

// SliceInfoForRegistration Contains the slice information requested during a Registration procedure
type SliceInfoForRegistration struct {
	SubscribedNssai            []SubscribedSnssai `json:"subscribedNssai,omitempty"`
	AllowedNssaiCurrentAccess  *AllowedNssai      `json:"allowedNssaiCurrentAccess,omitempty"`
	AllowedNssaiOtherAccess    *AllowedNssai      `json:"allowedNssaiOtherAccess,omitempty"`
	SNssaiForMapping           []Snssai           `json:"sNssaiForMapping,omitempty"`
	RequestedNssai             []Snssai           `json:"requestedNssai,omitempty"`
	DefaultConfiguredSnssaiInd *bool              `json:"defaultConfiguredSnssaiInd,omitempty"`
	MappingOfNssai             []MappingOfSnssai  `json:"mappingOfNssai,omitempty"`
	RequestMapping             *bool              `json:"requestMapping,omitempty"`
	UeSupNssrgInd              *bool              `json:"ueSupNssrgInd,omitempty"`
	SuppressNssrgInd           *bool              `json:"suppressNssrgInd,omitempty"`
	NsagSupported              *bool              `json:"nsagSupported,omitempty"`
}

// NewSliceInfoForRegistration instantiates a new SliceInfoForRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliceInfoForRegistration() *SliceInfoForRegistration {
	this := SliceInfoForRegistration{}
	var defaultConfiguredSnssaiInd bool = false
	this.DefaultConfiguredSnssaiInd = &defaultConfiguredSnssaiInd
	var nsagSupported bool = false
	this.NsagSupported = &nsagSupported
	return &this
}

// NewSliceInfoForRegistrationWithDefaults instantiates a new SliceInfoForRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliceInfoForRegistrationWithDefaults() *SliceInfoForRegistration {
	this := SliceInfoForRegistration{}
	var defaultConfiguredSnssaiInd bool = false
	this.DefaultConfiguredSnssaiInd = &defaultConfiguredSnssaiInd
	var nsagSupported bool = false
	this.NsagSupported = &nsagSupported
	return &this
}

// GetSubscribedNssai returns the SubscribedNssai field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetSubscribedNssai() []SubscribedSnssai {
	if o == nil || IsNil(o.SubscribedNssai) {
		var ret []SubscribedSnssai
		return ret
	}
	return o.SubscribedNssai
}

// GetSubscribedNssaiOk returns a tuple with the SubscribedNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetSubscribedNssaiOk() ([]SubscribedSnssai, bool) {
	if o == nil || IsNil(o.SubscribedNssai) {
		return nil, false
	}
	return o.SubscribedNssai, true
}

// HasSubscribedNssai returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasSubscribedNssai() bool {
	if o != nil && !IsNil(o.SubscribedNssai) {
		return true
	}

	return false
}

// SetSubscribedNssai gets a reference to the given []SubscribedSnssai and assigns it to the SubscribedNssai field.
func (o *SliceInfoForRegistration) SetSubscribedNssai(v []SubscribedSnssai) {
	o.SubscribedNssai = v
}

// GetAllowedNssaiCurrentAccess returns the AllowedNssaiCurrentAccess field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetAllowedNssaiCurrentAccess() AllowedNssai {
	if o == nil || IsNil(o.AllowedNssaiCurrentAccess) {
		var ret AllowedNssai
		return ret
	}
	return *o.AllowedNssaiCurrentAccess
}

// GetAllowedNssaiCurrentAccessOk returns a tuple with the AllowedNssaiCurrentAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetAllowedNssaiCurrentAccessOk() (*AllowedNssai, bool) {
	if o == nil || IsNil(o.AllowedNssaiCurrentAccess) {
		return nil, false
	}
	return o.AllowedNssaiCurrentAccess, true
}

// HasAllowedNssaiCurrentAccess returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasAllowedNssaiCurrentAccess() bool {
	if o != nil && !IsNil(o.AllowedNssaiCurrentAccess) {
		return true
	}

	return false
}

// SetAllowedNssaiCurrentAccess gets a reference to the given AllowedNssai and assigns it to the AllowedNssaiCurrentAccess field.
func (o *SliceInfoForRegistration) SetAllowedNssaiCurrentAccess(v AllowedNssai) {
	o.AllowedNssaiCurrentAccess = &v
}

// GetAllowedNssaiOtherAccess returns the AllowedNssaiOtherAccess field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetAllowedNssaiOtherAccess() AllowedNssai {
	if o == nil || IsNil(o.AllowedNssaiOtherAccess) {
		var ret AllowedNssai
		return ret
	}
	return *o.AllowedNssaiOtherAccess
}

// GetAllowedNssaiOtherAccessOk returns a tuple with the AllowedNssaiOtherAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetAllowedNssaiOtherAccessOk() (*AllowedNssai, bool) {
	if o == nil || IsNil(o.AllowedNssaiOtherAccess) {
		return nil, false
	}
	return o.AllowedNssaiOtherAccess, true
}

// HasAllowedNssaiOtherAccess returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasAllowedNssaiOtherAccess() bool {
	if o != nil && !IsNil(o.AllowedNssaiOtherAccess) {
		return true
	}

	return false
}

// SetAllowedNssaiOtherAccess gets a reference to the given AllowedNssai and assigns it to the AllowedNssaiOtherAccess field.
func (o *SliceInfoForRegistration) SetAllowedNssaiOtherAccess(v AllowedNssai) {
	o.AllowedNssaiOtherAccess = &v
}

// GetSNssaiForMapping returns the SNssaiForMapping field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetSNssaiForMapping() []Snssai {
	if o == nil || IsNil(o.SNssaiForMapping) {
		var ret []Snssai
		return ret
	}
	return o.SNssaiForMapping
}

// GetSNssaiForMappingOk returns a tuple with the SNssaiForMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetSNssaiForMappingOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SNssaiForMapping) {
		return nil, false
	}
	return o.SNssaiForMapping, true
}

// HasSNssaiForMapping returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasSNssaiForMapping() bool {
	if o != nil && !IsNil(o.SNssaiForMapping) {
		return true
	}

	return false
}

// SetSNssaiForMapping gets a reference to the given []Snssai and assigns it to the SNssaiForMapping field.
func (o *SliceInfoForRegistration) SetSNssaiForMapping(v []Snssai) {
	o.SNssaiForMapping = v
}

// GetRequestedNssai returns the RequestedNssai field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetRequestedNssai() []Snssai {
	if o == nil || IsNil(o.RequestedNssai) {
		var ret []Snssai
		return ret
	}
	return o.RequestedNssai
}

// GetRequestedNssaiOk returns a tuple with the RequestedNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetRequestedNssaiOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.RequestedNssai) {
		return nil, false
	}
	return o.RequestedNssai, true
}

// HasRequestedNssai returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasRequestedNssai() bool {
	if o != nil && !IsNil(o.RequestedNssai) {
		return true
	}

	return false
}

// SetRequestedNssai gets a reference to the given []Snssai and assigns it to the RequestedNssai field.
func (o *SliceInfoForRegistration) SetRequestedNssai(v []Snssai) {
	o.RequestedNssai = v
}

// GetDefaultConfiguredSnssaiInd returns the DefaultConfiguredSnssaiInd field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetDefaultConfiguredSnssaiInd() bool {
	if o == nil || IsNil(o.DefaultConfiguredSnssaiInd) {
		var ret bool
		return ret
	}
	return *o.DefaultConfiguredSnssaiInd
}

// GetDefaultConfiguredSnssaiIndOk returns a tuple with the DefaultConfiguredSnssaiInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetDefaultConfiguredSnssaiIndOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultConfiguredSnssaiInd) {
		return nil, false
	}
	return o.DefaultConfiguredSnssaiInd, true
}

// HasDefaultConfiguredSnssaiInd returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasDefaultConfiguredSnssaiInd() bool {
	if o != nil && !IsNil(o.DefaultConfiguredSnssaiInd) {
		return true
	}

	return false
}

// SetDefaultConfiguredSnssaiInd gets a reference to the given bool and assigns it to the DefaultConfiguredSnssaiInd field.
func (o *SliceInfoForRegistration) SetDefaultConfiguredSnssaiInd(v bool) {
	o.DefaultConfiguredSnssaiInd = &v
}

// GetMappingOfNssai returns the MappingOfNssai field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetMappingOfNssai() []MappingOfSnssai {
	if o == nil || IsNil(o.MappingOfNssai) {
		var ret []MappingOfSnssai
		return ret
	}
	return o.MappingOfNssai
}

// GetMappingOfNssaiOk returns a tuple with the MappingOfNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetMappingOfNssaiOk() ([]MappingOfSnssai, bool) {
	if o == nil || IsNil(o.MappingOfNssai) {
		return nil, false
	}
	return o.MappingOfNssai, true
}

// HasMappingOfNssai returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasMappingOfNssai() bool {
	if o != nil && !IsNil(o.MappingOfNssai) {
		return true
	}

	return false
}

// SetMappingOfNssai gets a reference to the given []MappingOfSnssai and assigns it to the MappingOfNssai field.
func (o *SliceInfoForRegistration) SetMappingOfNssai(v []MappingOfSnssai) {
	o.MappingOfNssai = v
}

// GetRequestMapping returns the RequestMapping field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetRequestMapping() bool {
	if o == nil || IsNil(o.RequestMapping) {
		var ret bool
		return ret
	}
	return *o.RequestMapping
}

// GetRequestMappingOk returns a tuple with the RequestMapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetRequestMappingOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestMapping) {
		return nil, false
	}
	return o.RequestMapping, true
}

// HasRequestMapping returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasRequestMapping() bool {
	if o != nil && !IsNil(o.RequestMapping) {
		return true
	}

	return false
}

// SetRequestMapping gets a reference to the given bool and assigns it to the RequestMapping field.
func (o *SliceInfoForRegistration) SetRequestMapping(v bool) {
	o.RequestMapping = &v
}

// GetUeSupNssrgInd returns the UeSupNssrgInd field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetUeSupNssrgInd() bool {
	if o == nil || IsNil(o.UeSupNssrgInd) {
		var ret bool
		return ret
	}
	return *o.UeSupNssrgInd
}

// GetUeSupNssrgIndOk returns a tuple with the UeSupNssrgInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetUeSupNssrgIndOk() (*bool, bool) {
	if o == nil || IsNil(o.UeSupNssrgInd) {
		return nil, false
	}
	return o.UeSupNssrgInd, true
}

// HasUeSupNssrgInd returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasUeSupNssrgInd() bool {
	if o != nil && !IsNil(o.UeSupNssrgInd) {
		return true
	}

	return false
}

// SetUeSupNssrgInd gets a reference to the given bool and assigns it to the UeSupNssrgInd field.
func (o *SliceInfoForRegistration) SetUeSupNssrgInd(v bool) {
	o.UeSupNssrgInd = &v
}

// GetSuppressNssrgInd returns the SuppressNssrgInd field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetSuppressNssrgInd() bool {
	if o == nil || IsNil(o.SuppressNssrgInd) {
		var ret bool
		return ret
	}
	return *o.SuppressNssrgInd
}

// GetSuppressNssrgIndOk returns a tuple with the SuppressNssrgInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetSuppressNssrgIndOk() (*bool, bool) {
	if o == nil || IsNil(o.SuppressNssrgInd) {
		return nil, false
	}
	return o.SuppressNssrgInd, true
}

// HasSuppressNssrgInd returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasSuppressNssrgInd() bool {
	if o != nil && !IsNil(o.SuppressNssrgInd) {
		return true
	}

	return false
}

// SetSuppressNssrgInd gets a reference to the given bool and assigns it to the SuppressNssrgInd field.
func (o *SliceInfoForRegistration) SetSuppressNssrgInd(v bool) {
	o.SuppressNssrgInd = &v
}

// GetNsagSupported returns the NsagSupported field value if set, zero value otherwise.
func (o *SliceInfoForRegistration) GetNsagSupported() bool {
	if o == nil || IsNil(o.NsagSupported) {
		var ret bool
		return ret
	}
	return *o.NsagSupported
}

// GetNsagSupportedOk returns a tuple with the NsagSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForRegistration) GetNsagSupportedOk() (*bool, bool) {
	if o == nil || IsNil(o.NsagSupported) {
		return nil, false
	}
	return o.NsagSupported, true
}

// HasNsagSupported returns a boolean if a field has been set.
func (o *SliceInfoForRegistration) HasNsagSupported() bool {
	if o != nil && !IsNil(o.NsagSupported) {
		return true
	}

	return false
}

// SetNsagSupported gets a reference to the given bool and assigns it to the NsagSupported field.
func (o *SliceInfoForRegistration) SetNsagSupported(v bool) {
	o.NsagSupported = &v
}

func (o SliceInfoForRegistration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SliceInfoForRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubscribedNssai) {
		toSerialize["subscribedNssai"] = o.SubscribedNssai
	}
	if !IsNil(o.AllowedNssaiCurrentAccess) {
		toSerialize["allowedNssaiCurrentAccess"] = o.AllowedNssaiCurrentAccess
	}
	if !IsNil(o.AllowedNssaiOtherAccess) {
		toSerialize["allowedNssaiOtherAccess"] = o.AllowedNssaiOtherAccess
	}
	if !IsNil(o.SNssaiForMapping) {
		toSerialize["sNssaiForMapping"] = o.SNssaiForMapping
	}
	if !IsNil(o.RequestedNssai) {
		toSerialize["requestedNssai"] = o.RequestedNssai
	}
	if !IsNil(o.DefaultConfiguredSnssaiInd) {
		toSerialize["defaultConfiguredSnssaiInd"] = o.DefaultConfiguredSnssaiInd
	}
	if !IsNil(o.MappingOfNssai) {
		toSerialize["mappingOfNssai"] = o.MappingOfNssai
	}
	if !IsNil(o.RequestMapping) {
		toSerialize["requestMapping"] = o.RequestMapping
	}
	if !IsNil(o.UeSupNssrgInd) {
		toSerialize["ueSupNssrgInd"] = o.UeSupNssrgInd
	}
	if !IsNil(o.SuppressNssrgInd) {
		toSerialize["suppressNssrgInd"] = o.SuppressNssrgInd
	}
	if !IsNil(o.NsagSupported) {
		toSerialize["nsagSupported"] = o.NsagSupported
	}
	return toSerialize, nil
}

type NullableSliceInfoForRegistration struct {
	value *SliceInfoForRegistration
	isSet bool
}

func (v NullableSliceInfoForRegistration) Get() *SliceInfoForRegistration {
	return v.value
}

func (v *NullableSliceInfoForRegistration) Set(val *SliceInfoForRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceInfoForRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceInfoForRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceInfoForRegistration(val *SliceInfoForRegistration) *NullableSliceInfoForRegistration {
	return &NullableSliceInfoForRegistration{value: val, isSet: true}
}

func (v NullableSliceInfoForRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceInfoForRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
