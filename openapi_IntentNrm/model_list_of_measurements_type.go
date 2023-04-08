/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the ListOfMeasurementsType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListOfMeasurementsType{}

// ListOfMeasurementsType See details in 3GPP TS 32.422 clause 5.10.3 for details.
type ListOfMeasurementsType struct {
	UMTS []string `json:"UMTS,omitempty"`
	LTE []string `json:"LTE,omitempty"`
	NR []string `json:"NR,omitempty"`
}

// NewListOfMeasurementsType instantiates a new ListOfMeasurementsType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListOfMeasurementsType() *ListOfMeasurementsType {
	this := ListOfMeasurementsType{}
	return &this
}

// NewListOfMeasurementsTypeWithDefaults instantiates a new ListOfMeasurementsType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListOfMeasurementsTypeWithDefaults() *ListOfMeasurementsType {
	this := ListOfMeasurementsType{}
	return &this
}

// GetUMTS returns the UMTS field value if set, zero value otherwise.
func (o *ListOfMeasurementsType) GetUMTS() []string {
	if o == nil || isNil(o.UMTS) {
		var ret []string
		return ret
	}
	return o.UMTS
}

// GetUMTSOk returns a tuple with the UMTS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOfMeasurementsType) GetUMTSOk() ([]string, bool) {
	if o == nil || isNil(o.UMTS) {
		return nil, false
	}
	return o.UMTS, true
}

// HasUMTS returns a boolean if a field has been set.
func (o *ListOfMeasurementsType) HasUMTS() bool {
	if o != nil && !isNil(o.UMTS) {
		return true
	}

	return false
}

// SetUMTS gets a reference to the given []string and assigns it to the UMTS field.
func (o *ListOfMeasurementsType) SetUMTS(v []string) {
	o.UMTS = v
}

// GetLTE returns the LTE field value if set, zero value otherwise.
func (o *ListOfMeasurementsType) GetLTE() []string {
	if o == nil || isNil(o.LTE) {
		var ret []string
		return ret
	}
	return o.LTE
}

// GetLTEOk returns a tuple with the LTE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOfMeasurementsType) GetLTEOk() ([]string, bool) {
	if o == nil || isNil(o.LTE) {
		return nil, false
	}
	return o.LTE, true
}

// HasLTE returns a boolean if a field has been set.
func (o *ListOfMeasurementsType) HasLTE() bool {
	if o != nil && !isNil(o.LTE) {
		return true
	}

	return false
}

// SetLTE gets a reference to the given []string and assigns it to the LTE field.
func (o *ListOfMeasurementsType) SetLTE(v []string) {
	o.LTE = v
}

// GetNR returns the NR field value if set, zero value otherwise.
func (o *ListOfMeasurementsType) GetNR() []string {
	if o == nil || isNil(o.NR) {
		var ret []string
		return ret
	}
	return o.NR
}

// GetNROk returns a tuple with the NR field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListOfMeasurementsType) GetNROk() ([]string, bool) {
	if o == nil || isNil(o.NR) {
		return nil, false
	}
	return o.NR, true
}

// HasNR returns a boolean if a field has been set.
func (o *ListOfMeasurementsType) HasNR() bool {
	if o != nil && !isNil(o.NR) {
		return true
	}

	return false
}

// SetNR gets a reference to the given []string and assigns it to the NR field.
func (o *ListOfMeasurementsType) SetNR(v []string) {
	o.NR = v
}

func (o ListOfMeasurementsType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListOfMeasurementsType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UMTS) {
		toSerialize["UMTS"] = o.UMTS
	}
	if !isNil(o.LTE) {
		toSerialize["LTE"] = o.LTE
	}
	if !isNil(o.NR) {
		toSerialize["NR"] = o.NR
	}
	return toSerialize, nil
}

type NullableListOfMeasurementsType struct {
	value *ListOfMeasurementsType
	isSet bool
}

func (v NullableListOfMeasurementsType) Get() *ListOfMeasurementsType {
	return v.value
}

func (v *NullableListOfMeasurementsType) Set(val *ListOfMeasurementsType) {
	v.value = val
	v.isSet = true
}

func (v NullableListOfMeasurementsType) IsSet() bool {
	return v.isSet
}

func (v *NullableListOfMeasurementsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListOfMeasurementsType(val *ListOfMeasurementsType) *NullableListOfMeasurementsType {
	return &NullableListOfMeasurementsType{value: val, isSet: true}
}

func (v NullableListOfMeasurementsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListOfMeasurementsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


