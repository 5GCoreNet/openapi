/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the M5QoSSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &M5QoSSpecification{}

// M5QoSSpecification struct for M5QoSSpecification
type M5QoSSpecification struct {
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwDlBitRate string `json:"marBwDlBitRate"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MarBwUlBitRate string `json:"marBwUlBitRate"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwDlBitRate *string `json:"minDesBwDlBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MinDesBwUlBitRate *string `json:"minDesBwUlBitRate,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwDlBitRate string `json:"mirBwDlBitRate"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\".
	MirBwUlBitRate string `json:"mirBwUlBitRate"`
	DesLatency     *int32 `json:"desLatency,omitempty"`
	DesLoss        *int32 `json:"desLoss,omitempty"`
}

// NewM5QoSSpecification instantiates a new M5QoSSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewM5QoSSpecification(marBwDlBitRate string, marBwUlBitRate string, mirBwDlBitRate string, mirBwUlBitRate string) *M5QoSSpecification {
	this := M5QoSSpecification{}
	this.MarBwDlBitRate = marBwDlBitRate
	this.MarBwUlBitRate = marBwUlBitRate
	this.MirBwDlBitRate = mirBwDlBitRate
	this.MirBwUlBitRate = mirBwUlBitRate
	return &this
}

// NewM5QoSSpecificationWithDefaults instantiates a new M5QoSSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewM5QoSSpecificationWithDefaults() *M5QoSSpecification {
	this := M5QoSSpecification{}
	return &this
}

// GetMarBwDlBitRate returns the MarBwDlBitRate field value
func (o *M5QoSSpecification) GetMarBwDlBitRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarBwDlBitRate
}

// GetMarBwDlBitRateOk returns a tuple with the MarBwDlBitRate field value
// and a boolean to check if the value has been set.
func (o *M5QoSSpecification) GetMarBwDlBitRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarBwDlBitRate, true
}

// SetMarBwDlBitRate sets field value
func (o *M5QoSSpecification) SetMarBwDlBitRate(v string) {
	o.MarBwDlBitRate = v
}

// GetMarBwUlBitRate returns the MarBwUlBitRate field value
func (o *M5QoSSpecification) GetMarBwUlBitRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MarBwUlBitRate
}

// GetMarBwUlBitRateOk returns a tuple with the MarBwUlBitRate field value
// and a boolean to check if the value has been set.
func (o *M5QoSSpecification) GetMarBwUlBitRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MarBwUlBitRate, true
}

// SetMarBwUlBitRate sets field value
func (o *M5QoSSpecification) SetMarBwUlBitRate(v string) {
	o.MarBwUlBitRate = v
}

// GetMinDesBwDlBitRate returns the MinDesBwDlBitRate field value if set, zero value otherwise.
func (o *M5QoSSpecification) GetMinDesBwDlBitRate() string {
	if o == nil || IsNil(o.MinDesBwDlBitRate) {
		var ret string
		return ret
	}
	return *o.MinDesBwDlBitRate
}

// GetMinDesBwDlBitRateOk returns a tuple with the MinDesBwDlBitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M5QoSSpecification) GetMinDesBwDlBitRateOk() (*string, bool) {
	if o == nil || IsNil(o.MinDesBwDlBitRate) {
		return nil, false
	}
	return o.MinDesBwDlBitRate, true
}

// HasMinDesBwDlBitRate returns a boolean if a field has been set.
func (o *M5QoSSpecification) HasMinDesBwDlBitRate() bool {
	if o != nil && !IsNil(o.MinDesBwDlBitRate) {
		return true
	}

	return false
}

// SetMinDesBwDlBitRate gets a reference to the given string and assigns it to the MinDesBwDlBitRate field.
func (o *M5QoSSpecification) SetMinDesBwDlBitRate(v string) {
	o.MinDesBwDlBitRate = &v
}

// GetMinDesBwUlBitRate returns the MinDesBwUlBitRate field value if set, zero value otherwise.
func (o *M5QoSSpecification) GetMinDesBwUlBitRate() string {
	if o == nil || IsNil(o.MinDesBwUlBitRate) {
		var ret string
		return ret
	}
	return *o.MinDesBwUlBitRate
}

// GetMinDesBwUlBitRateOk returns a tuple with the MinDesBwUlBitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M5QoSSpecification) GetMinDesBwUlBitRateOk() (*string, bool) {
	if o == nil || IsNil(o.MinDesBwUlBitRate) {
		return nil, false
	}
	return o.MinDesBwUlBitRate, true
}

// HasMinDesBwUlBitRate returns a boolean if a field has been set.
func (o *M5QoSSpecification) HasMinDesBwUlBitRate() bool {
	if o != nil && !IsNil(o.MinDesBwUlBitRate) {
		return true
	}

	return false
}

// SetMinDesBwUlBitRate gets a reference to the given string and assigns it to the MinDesBwUlBitRate field.
func (o *M5QoSSpecification) SetMinDesBwUlBitRate(v string) {
	o.MinDesBwUlBitRate = &v
}

// GetMirBwDlBitRate returns the MirBwDlBitRate field value
func (o *M5QoSSpecification) GetMirBwDlBitRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MirBwDlBitRate
}

// GetMirBwDlBitRateOk returns a tuple with the MirBwDlBitRate field value
// and a boolean to check if the value has been set.
func (o *M5QoSSpecification) GetMirBwDlBitRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirBwDlBitRate, true
}

// SetMirBwDlBitRate sets field value
func (o *M5QoSSpecification) SetMirBwDlBitRate(v string) {
	o.MirBwDlBitRate = v
}

// GetMirBwUlBitRate returns the MirBwUlBitRate field value
func (o *M5QoSSpecification) GetMirBwUlBitRate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MirBwUlBitRate
}

// GetMirBwUlBitRateOk returns a tuple with the MirBwUlBitRate field value
// and a boolean to check if the value has been set.
func (o *M5QoSSpecification) GetMirBwUlBitRateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MirBwUlBitRate, true
}

// SetMirBwUlBitRate sets field value
func (o *M5QoSSpecification) SetMirBwUlBitRate(v string) {
	o.MirBwUlBitRate = v
}

// GetDesLatency returns the DesLatency field value if set, zero value otherwise.
func (o *M5QoSSpecification) GetDesLatency() int32 {
	if o == nil || IsNil(o.DesLatency) {
		var ret int32
		return ret
	}
	return *o.DesLatency
}

// GetDesLatencyOk returns a tuple with the DesLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M5QoSSpecification) GetDesLatencyOk() (*int32, bool) {
	if o == nil || IsNil(o.DesLatency) {
		return nil, false
	}
	return o.DesLatency, true
}

// HasDesLatency returns a boolean if a field has been set.
func (o *M5QoSSpecification) HasDesLatency() bool {
	if o != nil && !IsNil(o.DesLatency) {
		return true
	}

	return false
}

// SetDesLatency gets a reference to the given int32 and assigns it to the DesLatency field.
func (o *M5QoSSpecification) SetDesLatency(v int32) {
	o.DesLatency = &v
}

// GetDesLoss returns the DesLoss field value if set, zero value otherwise.
func (o *M5QoSSpecification) GetDesLoss() int32 {
	if o == nil || IsNil(o.DesLoss) {
		var ret int32
		return ret
	}
	return *o.DesLoss
}

// GetDesLossOk returns a tuple with the DesLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M5QoSSpecification) GetDesLossOk() (*int32, bool) {
	if o == nil || IsNil(o.DesLoss) {
		return nil, false
	}
	return o.DesLoss, true
}

// HasDesLoss returns a boolean if a field has been set.
func (o *M5QoSSpecification) HasDesLoss() bool {
	if o != nil && !IsNil(o.DesLoss) {
		return true
	}

	return false
}

// SetDesLoss gets a reference to the given int32 and assigns it to the DesLoss field.
func (o *M5QoSSpecification) SetDesLoss(v int32) {
	o.DesLoss = &v
}

func (o M5QoSSpecification) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o M5QoSSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["marBwDlBitRate"] = o.MarBwDlBitRate
	toSerialize["marBwUlBitRate"] = o.MarBwUlBitRate
	if !IsNil(o.MinDesBwDlBitRate) {
		toSerialize["minDesBwDlBitRate"] = o.MinDesBwDlBitRate
	}
	if !IsNil(o.MinDesBwUlBitRate) {
		toSerialize["minDesBwUlBitRate"] = o.MinDesBwUlBitRate
	}
	toSerialize["mirBwDlBitRate"] = o.MirBwDlBitRate
	toSerialize["mirBwUlBitRate"] = o.MirBwUlBitRate
	if !IsNil(o.DesLatency) {
		toSerialize["desLatency"] = o.DesLatency
	}
	if !IsNil(o.DesLoss) {
		toSerialize["desLoss"] = o.DesLoss
	}
	return toSerialize, nil
}

type NullableM5QoSSpecification struct {
	value *M5QoSSpecification
	isSet bool
}

func (v NullableM5QoSSpecification) Get() *M5QoSSpecification {
	return v.value
}

func (v *NullableM5QoSSpecification) Set(val *M5QoSSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableM5QoSSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableM5QoSSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableM5QoSSpecification(val *M5QoSSpecification) *NullableM5QoSSpecification {
	return &NullableM5QoSSpecification{value: val, isSet: true}
}

func (v NullableM5QoSSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableM5QoSSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
