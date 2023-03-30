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

// checks if the SmsfFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmsfFunctionSingleAllOf1{}

// SmsfFunctionSingleAllOf1 struct for SmsfFunctionSingleAllOf1
type SmsfFunctionSingleAllOf1 struct {
	EPN20 []EPN20Single `json:"EP_N20,omitempty"`
	EPN21 []EPN21Single `json:"EP_N21,omitempty"`
	EP_MAP_SMSC []EPMAPSMSCSingle `json:"EP_MAP_SMSC,omitempty"`
}

// NewSmsfFunctionSingleAllOf1 instantiates a new SmsfFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmsfFunctionSingleAllOf1() *SmsfFunctionSingleAllOf1 {
	this := SmsfFunctionSingleAllOf1{}
	return &this
}

// NewSmsfFunctionSingleAllOf1WithDefaults instantiates a new SmsfFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmsfFunctionSingleAllOf1WithDefaults() *SmsfFunctionSingleAllOf1 {
	this := SmsfFunctionSingleAllOf1{}
	return &this
}

// GetEPN20 returns the EPN20 field value if set, zero value otherwise.
func (o *SmsfFunctionSingleAllOf1) GetEPN20() []EPN20Single {
	if o == nil || IsNil(o.EPN20) {
		var ret []EPN20Single
		return ret
	}
	return o.EPN20
}

// GetEPN20Ok returns a tuple with the EPN20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfFunctionSingleAllOf1) GetEPN20Ok() ([]EPN20Single, bool) {
	if o == nil || IsNil(o.EPN20) {
		return nil, false
	}
	return o.EPN20, true
}

// HasEPN20 returns a boolean if a field has been set.
func (o *SmsfFunctionSingleAllOf1) HasEPN20() bool {
	if o != nil && !IsNil(o.EPN20) {
		return true
	}

	return false
}

// SetEPN20 gets a reference to the given []EPN20Single and assigns it to the EPN20 field.
func (o *SmsfFunctionSingleAllOf1) SetEPN20(v []EPN20Single) {
	o.EPN20 = v
}

// GetEPN21 returns the EPN21 field value if set, zero value otherwise.
func (o *SmsfFunctionSingleAllOf1) GetEPN21() []EPN21Single {
	if o == nil || IsNil(o.EPN21) {
		var ret []EPN21Single
		return ret
	}
	return o.EPN21
}

// GetEPN21Ok returns a tuple with the EPN21 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfFunctionSingleAllOf1) GetEPN21Ok() ([]EPN21Single, bool) {
	if o == nil || IsNil(o.EPN21) {
		return nil, false
	}
	return o.EPN21, true
}

// HasEPN21 returns a boolean if a field has been set.
func (o *SmsfFunctionSingleAllOf1) HasEPN21() bool {
	if o != nil && !IsNil(o.EPN21) {
		return true
	}

	return false
}

// SetEPN21 gets a reference to the given []EPN21Single and assigns it to the EPN21 field.
func (o *SmsfFunctionSingleAllOf1) SetEPN21(v []EPN21Single) {
	o.EPN21 = v
}

// GetEP_MAP_SMSC returns the EP_MAP_SMSC field value if set, zero value otherwise.
func (o *SmsfFunctionSingleAllOf1) GetEP_MAP_SMSC() []EPMAPSMSCSingle {
	if o == nil || IsNil(o.EP_MAP_SMSC) {
		var ret []EPMAPSMSCSingle
		return ret
	}
	return o.EP_MAP_SMSC
}

// GetEP_MAP_SMSCOk returns a tuple with the EP_MAP_SMSC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmsfFunctionSingleAllOf1) GetEP_MAP_SMSCOk() ([]EPMAPSMSCSingle, bool) {
	if o == nil || IsNil(o.EP_MAP_SMSC) {
		return nil, false
	}
	return o.EP_MAP_SMSC, true
}

// HasEP_MAP_SMSC returns a boolean if a field has been set.
func (o *SmsfFunctionSingleAllOf1) HasEP_MAP_SMSC() bool {
	if o != nil && !IsNil(o.EP_MAP_SMSC) {
		return true
	}

	return false
}

// SetEP_MAP_SMSC gets a reference to the given []EPMAPSMSCSingle and assigns it to the EP_MAP_SMSC field.
func (o *SmsfFunctionSingleAllOf1) SetEP_MAP_SMSC(v []EPMAPSMSCSingle) {
	o.EP_MAP_SMSC = v
}

func (o SmsfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmsfFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EPN20) {
		toSerialize["EP_N20"] = o.EPN20
	}
	if !IsNil(o.EPN21) {
		toSerialize["EP_N21"] = o.EPN21
	}
	if !IsNil(o.EP_MAP_SMSC) {
		toSerialize["EP_MAP_SMSC"] = o.EP_MAP_SMSC
	}
	return toSerialize, nil
}

type NullableSmsfFunctionSingleAllOf1 struct {
	value *SmsfFunctionSingleAllOf1
	isSet bool
}

func (v NullableSmsfFunctionSingleAllOf1) Get() *SmsfFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableSmsfFunctionSingleAllOf1) Set(val *SmsfFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableSmsfFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSmsfFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmsfFunctionSingleAllOf1(val *SmsfFunctionSingleAllOf1) *NullableSmsfFunctionSingleAllOf1 {
	return &NullableSmsfFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableSmsfFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmsfFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

