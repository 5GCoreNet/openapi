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

// checks if the SliceInfoForUEConfigurationUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SliceInfoForUEConfigurationUpdate{}

// SliceInfoForUEConfigurationUpdate Contains the slice information requested during UE configuration update procedure
type SliceInfoForUEConfigurationUpdate struct {
	SubscribedNssai []SubscribedSnssai `json:"subscribedNssai,omitempty"`
	AllowedNssaiCurrentAccess *AllowedNssai `json:"allowedNssaiCurrentAccess,omitempty"`
	AllowedNssaiOtherAccess *AllowedNssai `json:"allowedNssaiOtherAccess,omitempty"`
	DefaultConfiguredSnssaiInd *bool `json:"defaultConfiguredSnssaiInd,omitempty"`
	RequestedNssai []Snssai `json:"requestedNssai,omitempty"`
	MappingOfNssai []MappingOfSnssai `json:"mappingOfNssai,omitempty"`
	UeSupNssrgInd *bool `json:"ueSupNssrgInd,omitempty"`
	SuppressNssrgInd *bool `json:"suppressNssrgInd,omitempty"`
	RejectedNssaiRa []Snssai `json:"rejectedNssaiRa,omitempty"`
	NsagSupported *bool `json:"nsagSupported,omitempty"`
}

// NewSliceInfoForUEConfigurationUpdate instantiates a new SliceInfoForUEConfigurationUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliceInfoForUEConfigurationUpdate() *SliceInfoForUEConfigurationUpdate {
	this := SliceInfoForUEConfigurationUpdate{}
	var nsagSupported bool = false
	this.NsagSupported = &nsagSupported
	return &this
}

// NewSliceInfoForUEConfigurationUpdateWithDefaults instantiates a new SliceInfoForUEConfigurationUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliceInfoForUEConfigurationUpdateWithDefaults() *SliceInfoForUEConfigurationUpdate {
	this := SliceInfoForUEConfigurationUpdate{}
	var nsagSupported bool = false
	this.NsagSupported = &nsagSupported
	return &this
}

// GetSubscribedNssai returns the SubscribedNssai field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetSubscribedNssai() []SubscribedSnssai {
	if o == nil || isNil(o.SubscribedNssai) {
		var ret []SubscribedSnssai
		return ret
	}
	return o.SubscribedNssai
}

// GetSubscribedNssaiOk returns a tuple with the SubscribedNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetSubscribedNssaiOk() ([]SubscribedSnssai, bool) {
	if o == nil || isNil(o.SubscribedNssai) {
		return nil, false
	}
	return o.SubscribedNssai, true
}

// HasSubscribedNssai returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasSubscribedNssai() bool {
	if o != nil && !isNil(o.SubscribedNssai) {
		return true
	}

	return false
}

// SetSubscribedNssai gets a reference to the given []SubscribedSnssai and assigns it to the SubscribedNssai field.
func (o *SliceInfoForUEConfigurationUpdate) SetSubscribedNssai(v []SubscribedSnssai) {
	o.SubscribedNssai = v
}

// GetAllowedNssaiCurrentAccess returns the AllowedNssaiCurrentAccess field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetAllowedNssaiCurrentAccess() AllowedNssai {
	if o == nil || isNil(o.AllowedNssaiCurrentAccess) {
		var ret AllowedNssai
		return ret
	}
	return *o.AllowedNssaiCurrentAccess
}

// GetAllowedNssaiCurrentAccessOk returns a tuple with the AllowedNssaiCurrentAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetAllowedNssaiCurrentAccessOk() (*AllowedNssai, bool) {
	if o == nil || isNil(o.AllowedNssaiCurrentAccess) {
		return nil, false
	}
	return o.AllowedNssaiCurrentAccess, true
}

// HasAllowedNssaiCurrentAccess returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasAllowedNssaiCurrentAccess() bool {
	if o != nil && !isNil(o.AllowedNssaiCurrentAccess) {
		return true
	}

	return false
}

// SetAllowedNssaiCurrentAccess gets a reference to the given AllowedNssai and assigns it to the AllowedNssaiCurrentAccess field.
func (o *SliceInfoForUEConfigurationUpdate) SetAllowedNssaiCurrentAccess(v AllowedNssai) {
	o.AllowedNssaiCurrentAccess = &v
}

// GetAllowedNssaiOtherAccess returns the AllowedNssaiOtherAccess field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetAllowedNssaiOtherAccess() AllowedNssai {
	if o == nil || isNil(o.AllowedNssaiOtherAccess) {
		var ret AllowedNssai
		return ret
	}
	return *o.AllowedNssaiOtherAccess
}

// GetAllowedNssaiOtherAccessOk returns a tuple with the AllowedNssaiOtherAccess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetAllowedNssaiOtherAccessOk() (*AllowedNssai, bool) {
	if o == nil || isNil(o.AllowedNssaiOtherAccess) {
		return nil, false
	}
	return o.AllowedNssaiOtherAccess, true
}

// HasAllowedNssaiOtherAccess returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasAllowedNssaiOtherAccess() bool {
	if o != nil && !isNil(o.AllowedNssaiOtherAccess) {
		return true
	}

	return false
}

// SetAllowedNssaiOtherAccess gets a reference to the given AllowedNssai and assigns it to the AllowedNssaiOtherAccess field.
func (o *SliceInfoForUEConfigurationUpdate) SetAllowedNssaiOtherAccess(v AllowedNssai) {
	o.AllowedNssaiOtherAccess = &v
}

// GetDefaultConfiguredSnssaiInd returns the DefaultConfiguredSnssaiInd field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetDefaultConfiguredSnssaiInd() bool {
	if o == nil || isNil(o.DefaultConfiguredSnssaiInd) {
		var ret bool
		return ret
	}
	return *o.DefaultConfiguredSnssaiInd
}

// GetDefaultConfiguredSnssaiIndOk returns a tuple with the DefaultConfiguredSnssaiInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetDefaultConfiguredSnssaiIndOk() (*bool, bool) {
	if o == nil || isNil(o.DefaultConfiguredSnssaiInd) {
		return nil, false
	}
	return o.DefaultConfiguredSnssaiInd, true
}

// HasDefaultConfiguredSnssaiInd returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasDefaultConfiguredSnssaiInd() bool {
	if o != nil && !isNil(o.DefaultConfiguredSnssaiInd) {
		return true
	}

	return false
}

// SetDefaultConfiguredSnssaiInd gets a reference to the given bool and assigns it to the DefaultConfiguredSnssaiInd field.
func (o *SliceInfoForUEConfigurationUpdate) SetDefaultConfiguredSnssaiInd(v bool) {
	o.DefaultConfiguredSnssaiInd = &v
}

// GetRequestedNssai returns the RequestedNssai field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetRequestedNssai() []Snssai {
	if o == nil || isNil(o.RequestedNssai) {
		var ret []Snssai
		return ret
	}
	return o.RequestedNssai
}

// GetRequestedNssaiOk returns a tuple with the RequestedNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetRequestedNssaiOk() ([]Snssai, bool) {
	if o == nil || isNil(o.RequestedNssai) {
		return nil, false
	}
	return o.RequestedNssai, true
}

// HasRequestedNssai returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasRequestedNssai() bool {
	if o != nil && !isNil(o.RequestedNssai) {
		return true
	}

	return false
}

// SetRequestedNssai gets a reference to the given []Snssai and assigns it to the RequestedNssai field.
func (o *SliceInfoForUEConfigurationUpdate) SetRequestedNssai(v []Snssai) {
	o.RequestedNssai = v
}

// GetMappingOfNssai returns the MappingOfNssai field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetMappingOfNssai() []MappingOfSnssai {
	if o == nil || isNil(o.MappingOfNssai) {
		var ret []MappingOfSnssai
		return ret
	}
	return o.MappingOfNssai
}

// GetMappingOfNssaiOk returns a tuple with the MappingOfNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetMappingOfNssaiOk() ([]MappingOfSnssai, bool) {
	if o == nil || isNil(o.MappingOfNssai) {
		return nil, false
	}
	return o.MappingOfNssai, true
}

// HasMappingOfNssai returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasMappingOfNssai() bool {
	if o != nil && !isNil(o.MappingOfNssai) {
		return true
	}

	return false
}

// SetMappingOfNssai gets a reference to the given []MappingOfSnssai and assigns it to the MappingOfNssai field.
func (o *SliceInfoForUEConfigurationUpdate) SetMappingOfNssai(v []MappingOfSnssai) {
	o.MappingOfNssai = v
}

// GetUeSupNssrgInd returns the UeSupNssrgInd field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetUeSupNssrgInd() bool {
	if o == nil || isNil(o.UeSupNssrgInd) {
		var ret bool
		return ret
	}
	return *o.UeSupNssrgInd
}

// GetUeSupNssrgIndOk returns a tuple with the UeSupNssrgInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetUeSupNssrgIndOk() (*bool, bool) {
	if o == nil || isNil(o.UeSupNssrgInd) {
		return nil, false
	}
	return o.UeSupNssrgInd, true
}

// HasUeSupNssrgInd returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasUeSupNssrgInd() bool {
	if o != nil && !isNil(o.UeSupNssrgInd) {
		return true
	}

	return false
}

// SetUeSupNssrgInd gets a reference to the given bool and assigns it to the UeSupNssrgInd field.
func (o *SliceInfoForUEConfigurationUpdate) SetUeSupNssrgInd(v bool) {
	o.UeSupNssrgInd = &v
}

// GetSuppressNssrgInd returns the SuppressNssrgInd field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetSuppressNssrgInd() bool {
	if o == nil || isNil(o.SuppressNssrgInd) {
		var ret bool
		return ret
	}
	return *o.SuppressNssrgInd
}

// GetSuppressNssrgIndOk returns a tuple with the SuppressNssrgInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetSuppressNssrgIndOk() (*bool, bool) {
	if o == nil || isNil(o.SuppressNssrgInd) {
		return nil, false
	}
	return o.SuppressNssrgInd, true
}

// HasSuppressNssrgInd returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasSuppressNssrgInd() bool {
	if o != nil && !isNil(o.SuppressNssrgInd) {
		return true
	}

	return false
}

// SetSuppressNssrgInd gets a reference to the given bool and assigns it to the SuppressNssrgInd field.
func (o *SliceInfoForUEConfigurationUpdate) SetSuppressNssrgInd(v bool) {
	o.SuppressNssrgInd = &v
}

// GetRejectedNssaiRa returns the RejectedNssaiRa field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetRejectedNssaiRa() []Snssai {
	if o == nil || isNil(o.RejectedNssaiRa) {
		var ret []Snssai
		return ret
	}
	return o.RejectedNssaiRa
}

// GetRejectedNssaiRaOk returns a tuple with the RejectedNssaiRa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetRejectedNssaiRaOk() ([]Snssai, bool) {
	if o == nil || isNil(o.RejectedNssaiRa) {
		return nil, false
	}
	return o.RejectedNssaiRa, true
}

// HasRejectedNssaiRa returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasRejectedNssaiRa() bool {
	if o != nil && !isNil(o.RejectedNssaiRa) {
		return true
	}

	return false
}

// SetRejectedNssaiRa gets a reference to the given []Snssai and assigns it to the RejectedNssaiRa field.
func (o *SliceInfoForUEConfigurationUpdate) SetRejectedNssaiRa(v []Snssai) {
	o.RejectedNssaiRa = v
}

// GetNsagSupported returns the NsagSupported field value if set, zero value otherwise.
func (o *SliceInfoForUEConfigurationUpdate) GetNsagSupported() bool {
	if o == nil || isNil(o.NsagSupported) {
		var ret bool
		return ret
	}
	return *o.NsagSupported
}

// GetNsagSupportedOk returns a tuple with the NsagSupported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceInfoForUEConfigurationUpdate) GetNsagSupportedOk() (*bool, bool) {
	if o == nil || isNil(o.NsagSupported) {
		return nil, false
	}
	return o.NsagSupported, true
}

// HasNsagSupported returns a boolean if a field has been set.
func (o *SliceInfoForUEConfigurationUpdate) HasNsagSupported() bool {
	if o != nil && !isNil(o.NsagSupported) {
		return true
	}

	return false
}

// SetNsagSupported gets a reference to the given bool and assigns it to the NsagSupported field.
func (o *SliceInfoForUEConfigurationUpdate) SetNsagSupported(v bool) {
	o.NsagSupported = &v
}

func (o SliceInfoForUEConfigurationUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SliceInfoForUEConfigurationUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SubscribedNssai) {
		toSerialize["subscribedNssai"] = o.SubscribedNssai
	}
	if !isNil(o.AllowedNssaiCurrentAccess) {
		toSerialize["allowedNssaiCurrentAccess"] = o.AllowedNssaiCurrentAccess
	}
	if !isNil(o.AllowedNssaiOtherAccess) {
		toSerialize["allowedNssaiOtherAccess"] = o.AllowedNssaiOtherAccess
	}
	if !isNil(o.DefaultConfiguredSnssaiInd) {
		toSerialize["defaultConfiguredSnssaiInd"] = o.DefaultConfiguredSnssaiInd
	}
	if !isNil(o.RequestedNssai) {
		toSerialize["requestedNssai"] = o.RequestedNssai
	}
	if !isNil(o.MappingOfNssai) {
		toSerialize["mappingOfNssai"] = o.MappingOfNssai
	}
	if !isNil(o.UeSupNssrgInd) {
		toSerialize["ueSupNssrgInd"] = o.UeSupNssrgInd
	}
	if !isNil(o.SuppressNssrgInd) {
		toSerialize["suppressNssrgInd"] = o.SuppressNssrgInd
	}
	if !isNil(o.RejectedNssaiRa) {
		toSerialize["rejectedNssaiRa"] = o.RejectedNssaiRa
	}
	if !isNil(o.NsagSupported) {
		toSerialize["nsagSupported"] = o.NsagSupported
	}
	return toSerialize, nil
}

type NullableSliceInfoForUEConfigurationUpdate struct {
	value *SliceInfoForUEConfigurationUpdate
	isSet bool
}

func (v NullableSliceInfoForUEConfigurationUpdate) Get() *SliceInfoForUEConfigurationUpdate {
	return v.value
}

func (v *NullableSliceInfoForUEConfigurationUpdate) Set(val *SliceInfoForUEConfigurationUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceInfoForUEConfigurationUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceInfoForUEConfigurationUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceInfoForUEConfigurationUpdate(val *SliceInfoForUEConfigurationUpdate) *NullableSliceInfoForUEConfigurationUpdate {
	return &NullableSliceInfoForUEConfigurationUpdate{value: val, isSet: true}
}

func (v NullableSliceInfoForUEConfigurationUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceInfoForUEConfigurationUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


