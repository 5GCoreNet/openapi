/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the CCOPilotPollutionParametersSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CCOPilotPollutionParametersSingle{}

// CCOPilotPollutionParametersSingle struct for CCOPilotPollutionParametersSingle
type CCOPilotPollutionParametersSingle struct {
	CCOParametersAttr
}

// NewCCOPilotPollutionParametersSingle instantiates a new CCOPilotPollutionParametersSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCCOPilotPollutionParametersSingle(id NullableString) *CCOPilotPollutionParametersSingle {
	this := CCOPilotPollutionParametersSingle{}
	this.Id = id
	return &this
}

// NewCCOPilotPollutionParametersSingleWithDefaults instantiates a new CCOPilotPollutionParametersSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCCOPilotPollutionParametersSingleWithDefaults() *CCOPilotPollutionParametersSingle {
	this := CCOPilotPollutionParametersSingle{}
	return &this
}

func (o CCOPilotPollutionParametersSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CCOPilotPollutionParametersSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedCCOParametersAttr, errCCOParametersAttr := json.Marshal(o.CCOParametersAttr)
	if errCCOParametersAttr != nil {
		return map[string]interface{}{}, errCCOParametersAttr
	}
	errCCOParametersAttr = json.Unmarshal([]byte(serializedCCOParametersAttr), &toSerialize)
	if errCCOParametersAttr != nil {
		return map[string]interface{}{}, errCCOParametersAttr
	}
	return toSerialize, nil
}

type NullableCCOPilotPollutionParametersSingle struct {
	value *CCOPilotPollutionParametersSingle
	isSet bool
}

func (v NullableCCOPilotPollutionParametersSingle) Get() *CCOPilotPollutionParametersSingle {
	return v.value
}

func (v *NullableCCOPilotPollutionParametersSingle) Set(val *CCOPilotPollutionParametersSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableCCOPilotPollutionParametersSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableCCOPilotPollutionParametersSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCCOPilotPollutionParametersSingle(val *CCOPilotPollutionParametersSingle) *NullableCCOPilotPollutionParametersSingle {
	return &NullableCCOPilotPollutionParametersSingle{value: val, isSet: true}
}

func (v NullableCCOPilotPollutionParametersSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCCOPilotPollutionParametersSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
