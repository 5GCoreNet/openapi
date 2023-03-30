/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the RequestedUnit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestedUnit{}

// RequestedUnit struct for RequestedUnit
type RequestedUnit struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	Time *int32 `json:"time,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	TotalVolume *int32 `json:"totalVolume,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	UplinkVolume *int32 `json:"uplinkVolume,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	DownlinkVolume *int32 `json:"downlinkVolume,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer. 
	ServiceSpecificUnits *int32 `json:"serviceSpecificUnits,omitempty"`
}

// NewRequestedUnit instantiates a new RequestedUnit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestedUnit() *RequestedUnit {
	this := RequestedUnit{}
	return &this
}

// NewRequestedUnitWithDefaults instantiates a new RequestedUnit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestedUnitWithDefaults() *RequestedUnit {
	this := RequestedUnit{}
	return &this
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *RequestedUnit) GetTime() int32 {
	if o == nil || IsNil(o.Time) {
		var ret int32
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUnit) GetTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.Time) {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *RequestedUnit) HasTime() bool {
	if o != nil && !IsNil(o.Time) {
		return true
	}

	return false
}

// SetTime gets a reference to the given int32 and assigns it to the Time field.
func (o *RequestedUnit) SetTime(v int32) {
	o.Time = &v
}

// GetTotalVolume returns the TotalVolume field value if set, zero value otherwise.
func (o *RequestedUnit) GetTotalVolume() int32 {
	if o == nil || IsNil(o.TotalVolume) {
		var ret int32
		return ret
	}
	return *o.TotalVolume
}

// GetTotalVolumeOk returns a tuple with the TotalVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUnit) GetTotalVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalVolume) {
		return nil, false
	}
	return o.TotalVolume, true
}

// HasTotalVolume returns a boolean if a field has been set.
func (o *RequestedUnit) HasTotalVolume() bool {
	if o != nil && !IsNil(o.TotalVolume) {
		return true
	}

	return false
}

// SetTotalVolume gets a reference to the given int32 and assigns it to the TotalVolume field.
func (o *RequestedUnit) SetTotalVolume(v int32) {
	o.TotalVolume = &v
}

// GetUplinkVolume returns the UplinkVolume field value if set, zero value otherwise.
func (o *RequestedUnit) GetUplinkVolume() int32 {
	if o == nil || IsNil(o.UplinkVolume) {
		var ret int32
		return ret
	}
	return *o.UplinkVolume
}

// GetUplinkVolumeOk returns a tuple with the UplinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUnit) GetUplinkVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.UplinkVolume) {
		return nil, false
	}
	return o.UplinkVolume, true
}

// HasUplinkVolume returns a boolean if a field has been set.
func (o *RequestedUnit) HasUplinkVolume() bool {
	if o != nil && !IsNil(o.UplinkVolume) {
		return true
	}

	return false
}

// SetUplinkVolume gets a reference to the given int32 and assigns it to the UplinkVolume field.
func (o *RequestedUnit) SetUplinkVolume(v int32) {
	o.UplinkVolume = &v
}

// GetDownlinkVolume returns the DownlinkVolume field value if set, zero value otherwise.
func (o *RequestedUnit) GetDownlinkVolume() int32 {
	if o == nil || IsNil(o.DownlinkVolume) {
		var ret int32
		return ret
	}
	return *o.DownlinkVolume
}

// GetDownlinkVolumeOk returns a tuple with the DownlinkVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUnit) GetDownlinkVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.DownlinkVolume) {
		return nil, false
	}
	return o.DownlinkVolume, true
}

// HasDownlinkVolume returns a boolean if a field has been set.
func (o *RequestedUnit) HasDownlinkVolume() bool {
	if o != nil && !IsNil(o.DownlinkVolume) {
		return true
	}

	return false
}

// SetDownlinkVolume gets a reference to the given int32 and assigns it to the DownlinkVolume field.
func (o *RequestedUnit) SetDownlinkVolume(v int32) {
	o.DownlinkVolume = &v
}

// GetServiceSpecificUnits returns the ServiceSpecificUnits field value if set, zero value otherwise.
func (o *RequestedUnit) GetServiceSpecificUnits() int32 {
	if o == nil || IsNil(o.ServiceSpecificUnits) {
		var ret int32
		return ret
	}
	return *o.ServiceSpecificUnits
}

// GetServiceSpecificUnitsOk returns a tuple with the ServiceSpecificUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestedUnit) GetServiceSpecificUnitsOk() (*int32, bool) {
	if o == nil || IsNil(o.ServiceSpecificUnits) {
		return nil, false
	}
	return o.ServiceSpecificUnits, true
}

// HasServiceSpecificUnits returns a boolean if a field has been set.
func (o *RequestedUnit) HasServiceSpecificUnits() bool {
	if o != nil && !IsNil(o.ServiceSpecificUnits) {
		return true
	}

	return false
}

// SetServiceSpecificUnits gets a reference to the given int32 and assigns it to the ServiceSpecificUnits field.
func (o *RequestedUnit) SetServiceSpecificUnits(v int32) {
	o.ServiceSpecificUnits = &v
}

func (o RequestedUnit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestedUnit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Time) {
		toSerialize["time"] = o.Time
	}
	if !IsNil(o.TotalVolume) {
		toSerialize["totalVolume"] = o.TotalVolume
	}
	if !IsNil(o.UplinkVolume) {
		toSerialize["uplinkVolume"] = o.UplinkVolume
	}
	if !IsNil(o.DownlinkVolume) {
		toSerialize["downlinkVolume"] = o.DownlinkVolume
	}
	if !IsNil(o.ServiceSpecificUnits) {
		toSerialize["serviceSpecificUnits"] = o.ServiceSpecificUnits
	}
	return toSerialize, nil
}

type NullableRequestedUnit struct {
	value *RequestedUnit
	isSet bool
}

func (v NullableRequestedUnit) Get() *RequestedUnit {
	return v.value
}

func (v *NullableRequestedUnit) Set(val *RequestedUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedUnit(val *RequestedUnit) *NullableRequestedUnit {
	return &NullableRequestedUnit{value: val, isSet: true}
}

func (v NullableRequestedUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


