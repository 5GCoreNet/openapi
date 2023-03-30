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

// checks if the EarlyMediaDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EarlyMediaDescription{}

// EarlyMediaDescription struct for EarlyMediaDescription
type EarlyMediaDescription struct {
	SDPTimeStamps *SDPTimeStamps `json:"sDPTimeStamps,omitempty"`
	SDPMediaComponent []SDPMediaComponent `json:"sDPMediaComponent,omitempty"`
	SDPSessionDescription []string `json:"sDPSessionDescription,omitempty"`
}

// NewEarlyMediaDescription instantiates a new EarlyMediaDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarlyMediaDescription() *EarlyMediaDescription {
	this := EarlyMediaDescription{}
	return &this
}

// NewEarlyMediaDescriptionWithDefaults instantiates a new EarlyMediaDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarlyMediaDescriptionWithDefaults() *EarlyMediaDescription {
	this := EarlyMediaDescription{}
	return &this
}

// GetSDPTimeStamps returns the SDPTimeStamps field value if set, zero value otherwise.
func (o *EarlyMediaDescription) GetSDPTimeStamps() SDPTimeStamps {
	if o == nil || IsNil(o.SDPTimeStamps) {
		var ret SDPTimeStamps
		return ret
	}
	return *o.SDPTimeStamps
}

// GetSDPTimeStampsOk returns a tuple with the SDPTimeStamps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarlyMediaDescription) GetSDPTimeStampsOk() (*SDPTimeStamps, bool) {
	if o == nil || IsNil(o.SDPTimeStamps) {
		return nil, false
	}
	return o.SDPTimeStamps, true
}

// HasSDPTimeStamps returns a boolean if a field has been set.
func (o *EarlyMediaDescription) HasSDPTimeStamps() bool {
	if o != nil && !IsNil(o.SDPTimeStamps) {
		return true
	}

	return false
}

// SetSDPTimeStamps gets a reference to the given SDPTimeStamps and assigns it to the SDPTimeStamps field.
func (o *EarlyMediaDescription) SetSDPTimeStamps(v SDPTimeStamps) {
	o.SDPTimeStamps = &v
}

// GetSDPMediaComponent returns the SDPMediaComponent field value if set, zero value otherwise.
func (o *EarlyMediaDescription) GetSDPMediaComponent() []SDPMediaComponent {
	if o == nil || IsNil(o.SDPMediaComponent) {
		var ret []SDPMediaComponent
		return ret
	}
	return o.SDPMediaComponent
}

// GetSDPMediaComponentOk returns a tuple with the SDPMediaComponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarlyMediaDescription) GetSDPMediaComponentOk() ([]SDPMediaComponent, bool) {
	if o == nil || IsNil(o.SDPMediaComponent) {
		return nil, false
	}
	return o.SDPMediaComponent, true
}

// HasSDPMediaComponent returns a boolean if a field has been set.
func (o *EarlyMediaDescription) HasSDPMediaComponent() bool {
	if o != nil && !IsNil(o.SDPMediaComponent) {
		return true
	}

	return false
}

// SetSDPMediaComponent gets a reference to the given []SDPMediaComponent and assigns it to the SDPMediaComponent field.
func (o *EarlyMediaDescription) SetSDPMediaComponent(v []SDPMediaComponent) {
	o.SDPMediaComponent = v
}

// GetSDPSessionDescription returns the SDPSessionDescription field value if set, zero value otherwise.
func (o *EarlyMediaDescription) GetSDPSessionDescription() []string {
	if o == nil || IsNil(o.SDPSessionDescription) {
		var ret []string
		return ret
	}
	return o.SDPSessionDescription
}

// GetSDPSessionDescriptionOk returns a tuple with the SDPSessionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EarlyMediaDescription) GetSDPSessionDescriptionOk() ([]string, bool) {
	if o == nil || IsNil(o.SDPSessionDescription) {
		return nil, false
	}
	return o.SDPSessionDescription, true
}

// HasSDPSessionDescription returns a boolean if a field has been set.
func (o *EarlyMediaDescription) HasSDPSessionDescription() bool {
	if o != nil && !IsNil(o.SDPSessionDescription) {
		return true
	}

	return false
}

// SetSDPSessionDescription gets a reference to the given []string and assigns it to the SDPSessionDescription field.
func (o *EarlyMediaDescription) SetSDPSessionDescription(v []string) {
	o.SDPSessionDescription = v
}

func (o EarlyMediaDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EarlyMediaDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SDPTimeStamps) {
		toSerialize["sDPTimeStamps"] = o.SDPTimeStamps
	}
	if !IsNil(o.SDPMediaComponent) {
		toSerialize["sDPMediaComponent"] = o.SDPMediaComponent
	}
	if !IsNil(o.SDPSessionDescription) {
		toSerialize["sDPSessionDescription"] = o.SDPSessionDescription
	}
	return toSerialize, nil
}

type NullableEarlyMediaDescription struct {
	value *EarlyMediaDescription
	isSet bool
}

func (v NullableEarlyMediaDescription) Get() *EarlyMediaDescription {
	return v.value
}

func (v *NullableEarlyMediaDescription) Set(val *EarlyMediaDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableEarlyMediaDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableEarlyMediaDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarlyMediaDescription(val *EarlyMediaDescription) *NullableEarlyMediaDescription {
	return &NullableEarlyMediaDescription{value: val, isSet: true}
}

func (v NullableEarlyMediaDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarlyMediaDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


