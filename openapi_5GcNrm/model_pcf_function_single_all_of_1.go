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

// checks if the PcfFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfFunctionSingleAllOf1{}

// PcfFunctionSingleAllOf1 struct for PcfFunctionSingleAllOf1
type PcfFunctionSingleAllOf1 struct {
	EPN5                 []EPN5Single                `json:"EP_N5,omitempty"`
	EPN7                 []EPN7Single                `json:"EP_N7,omitempty"`
	EPN15                []EPN15Single               `json:"EP_N15,omitempty"`
	EPN16                []EPN16Single               `json:"EP_N16,omitempty"`
	EPRx                 []EPRxSingle                `json:"EP_Rx,omitempty"`
	PredefinedPccRuleSet *PredefinedPccRuleSetSingle `json:"PredefinedPccRuleSet,omitempty"`
}

// NewPcfFunctionSingleAllOf1 instantiates a new PcfFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfFunctionSingleAllOf1() *PcfFunctionSingleAllOf1 {
	this := PcfFunctionSingleAllOf1{}
	return &this
}

// NewPcfFunctionSingleAllOf1WithDefaults instantiates a new PcfFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfFunctionSingleAllOf1WithDefaults() *PcfFunctionSingleAllOf1 {
	this := PcfFunctionSingleAllOf1{}
	return &this
}

// GetEPN5 returns the EPN5 field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOf1) GetEPN5() []EPN5Single {
	if o == nil || IsNil(o.EPN5) {
		var ret []EPN5Single
		return ret
	}
	return o.EPN5
}

// GetEPN5Ok returns a tuple with the EPN5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOf1) GetEPN5Ok() ([]EPN5Single, bool) {
	if o == nil || IsNil(o.EPN5) {
		return nil, false
	}
	return o.EPN5, true
}

// HasEPN5 returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOf1) HasEPN5() bool {
	if o != nil && !IsNil(o.EPN5) {
		return true
	}

	return false
}

// SetEPN5 gets a reference to the given []EPN5Single and assigns it to the EPN5 field.
func (o *PcfFunctionSingleAllOf1) SetEPN5(v []EPN5Single) {
	o.EPN5 = v
}

// GetEPN7 returns the EPN7 field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOf1) GetEPN7() []EPN7Single {
	if o == nil || IsNil(o.EPN7) {
		var ret []EPN7Single
		return ret
	}
	return o.EPN7
}

// GetEPN7Ok returns a tuple with the EPN7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOf1) GetEPN7Ok() ([]EPN7Single, bool) {
	if o == nil || IsNil(o.EPN7) {
		return nil, false
	}
	return o.EPN7, true
}

// HasEPN7 returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOf1) HasEPN7() bool {
	if o != nil && !IsNil(o.EPN7) {
		return true
	}

	return false
}

// SetEPN7 gets a reference to the given []EPN7Single and assigns it to the EPN7 field.
func (o *PcfFunctionSingleAllOf1) SetEPN7(v []EPN7Single) {
	o.EPN7 = v
}

// GetEPN15 returns the EPN15 field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOf1) GetEPN15() []EPN15Single {
	if o == nil || IsNil(o.EPN15) {
		var ret []EPN15Single
		return ret
	}
	return o.EPN15
}

// GetEPN15Ok returns a tuple with the EPN15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOf1) GetEPN15Ok() ([]EPN15Single, bool) {
	if o == nil || IsNil(o.EPN15) {
		return nil, false
	}
	return o.EPN15, true
}

// HasEPN15 returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOf1) HasEPN15() bool {
	if o != nil && !IsNil(o.EPN15) {
		return true
	}

	return false
}

// SetEPN15 gets a reference to the given []EPN15Single and assigns it to the EPN15 field.
func (o *PcfFunctionSingleAllOf1) SetEPN15(v []EPN15Single) {
	o.EPN15 = v
}

// GetEPN16 returns the EPN16 field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOf1) GetEPN16() []EPN16Single {
	if o == nil || IsNil(o.EPN16) {
		var ret []EPN16Single
		return ret
	}
	return o.EPN16
}

// GetEPN16Ok returns a tuple with the EPN16 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOf1) GetEPN16Ok() ([]EPN16Single, bool) {
	if o == nil || IsNil(o.EPN16) {
		return nil, false
	}
	return o.EPN16, true
}

// HasEPN16 returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOf1) HasEPN16() bool {
	if o != nil && !IsNil(o.EPN16) {
		return true
	}

	return false
}

// SetEPN16 gets a reference to the given []EPN16Single and assigns it to the EPN16 field.
func (o *PcfFunctionSingleAllOf1) SetEPN16(v []EPN16Single) {
	o.EPN16 = v
}

// GetEPRx returns the EPRx field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOf1) GetEPRx() []EPRxSingle {
	if o == nil || IsNil(o.EPRx) {
		var ret []EPRxSingle
		return ret
	}
	return o.EPRx
}

// GetEPRxOk returns a tuple with the EPRx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOf1) GetEPRxOk() ([]EPRxSingle, bool) {
	if o == nil || IsNil(o.EPRx) {
		return nil, false
	}
	return o.EPRx, true
}

// HasEPRx returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOf1) HasEPRx() bool {
	if o != nil && !IsNil(o.EPRx) {
		return true
	}

	return false
}

// SetEPRx gets a reference to the given []EPRxSingle and assigns it to the EPRx field.
func (o *PcfFunctionSingleAllOf1) SetEPRx(v []EPRxSingle) {
	o.EPRx = v
}

// GetPredefinedPccRuleSet returns the PredefinedPccRuleSet field value if set, zero value otherwise.
func (o *PcfFunctionSingleAllOf1) GetPredefinedPccRuleSet() PredefinedPccRuleSetSingle {
	if o == nil || IsNil(o.PredefinedPccRuleSet) {
		var ret PredefinedPccRuleSetSingle
		return ret
	}
	return *o.PredefinedPccRuleSet
}

// GetPredefinedPccRuleSetOk returns a tuple with the PredefinedPccRuleSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfFunctionSingleAllOf1) GetPredefinedPccRuleSetOk() (*PredefinedPccRuleSetSingle, bool) {
	if o == nil || IsNil(o.PredefinedPccRuleSet) {
		return nil, false
	}
	return o.PredefinedPccRuleSet, true
}

// HasPredefinedPccRuleSet returns a boolean if a field has been set.
func (o *PcfFunctionSingleAllOf1) HasPredefinedPccRuleSet() bool {
	if o != nil && !IsNil(o.PredefinedPccRuleSet) {
		return true
	}

	return false
}

// SetPredefinedPccRuleSet gets a reference to the given PredefinedPccRuleSetSingle and assigns it to the PredefinedPccRuleSet field.
func (o *PcfFunctionSingleAllOf1) SetPredefinedPccRuleSet(v PredefinedPccRuleSetSingle) {
	o.PredefinedPccRuleSet = &v
}

func (o PcfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EPN5) {
		toSerialize["EP_N5"] = o.EPN5
	}
	if !IsNil(o.EPN7) {
		toSerialize["EP_N7"] = o.EPN7
	}
	if !IsNil(o.EPN15) {
		toSerialize["EP_N15"] = o.EPN15
	}
	if !IsNil(o.EPN16) {
		toSerialize["EP_N16"] = o.EPN16
	}
	if !IsNil(o.EPRx) {
		toSerialize["EP_Rx"] = o.EPRx
	}
	if !IsNil(o.PredefinedPccRuleSet) {
		toSerialize["PredefinedPccRuleSet"] = o.PredefinedPccRuleSet
	}
	return toSerialize, nil
}

type NullablePcfFunctionSingleAllOf1 struct {
	value *PcfFunctionSingleAllOf1
	isSet bool
}

func (v NullablePcfFunctionSingleAllOf1) Get() *PcfFunctionSingleAllOf1 {
	return v.value
}

func (v *NullablePcfFunctionSingleAllOf1) Set(val *PcfFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfFunctionSingleAllOf1(val *PcfFunctionSingleAllOf1) *NullablePcfFunctionSingleAllOf1 {
	return &NullablePcfFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullablePcfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
