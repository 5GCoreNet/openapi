/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// checks if the EdgeDataNetworkSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeDataNetworkSingleAllOf1{}

// EdgeDataNetworkSingleAllOf1 struct for EdgeDataNetworkSingleAllOf1
type EdgeDataNetworkSingleAllOf1 struct {
	EASFunction []EASFunctionSingle `json:"EASFunction,omitempty"`
	EESFunction []EESFunctionSingle `json:"EESFunction,omitempty"`
}

// NewEdgeDataNetworkSingleAllOf1 instantiates a new EdgeDataNetworkSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeDataNetworkSingleAllOf1() *EdgeDataNetworkSingleAllOf1 {
	this := EdgeDataNetworkSingleAllOf1{}
	return &this
}

// NewEdgeDataNetworkSingleAllOf1WithDefaults instantiates a new EdgeDataNetworkSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeDataNetworkSingleAllOf1WithDefaults() *EdgeDataNetworkSingleAllOf1 {
	this := EdgeDataNetworkSingleAllOf1{}
	return &this
}

// GetEASFunction returns the EASFunction field value if set, zero value otherwise.
func (o *EdgeDataNetworkSingleAllOf1) GetEASFunction() []EASFunctionSingle {
	if o == nil || isNil(o.EASFunction) {
		var ret []EASFunctionSingle
		return ret
	}
	return o.EASFunction
}

// GetEASFunctionOk returns a tuple with the EASFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeDataNetworkSingleAllOf1) GetEASFunctionOk() ([]EASFunctionSingle, bool) {
	if o == nil || isNil(o.EASFunction) {
		return nil, false
	}
	return o.EASFunction, true
}

// HasEASFunction returns a boolean if a field has been set.
func (o *EdgeDataNetworkSingleAllOf1) HasEASFunction() bool {
	if o != nil && !isNil(o.EASFunction) {
		return true
	}

	return false
}

// SetEASFunction gets a reference to the given []EASFunctionSingle and assigns it to the EASFunction field.
func (o *EdgeDataNetworkSingleAllOf1) SetEASFunction(v []EASFunctionSingle) {
	o.EASFunction = v
}

// GetEESFunction returns the EESFunction field value if set, zero value otherwise.
func (o *EdgeDataNetworkSingleAllOf1) GetEESFunction() []EESFunctionSingle {
	if o == nil || isNil(o.EESFunction) {
		var ret []EESFunctionSingle
		return ret
	}
	return o.EESFunction
}

// GetEESFunctionOk returns a tuple with the EESFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EdgeDataNetworkSingleAllOf1) GetEESFunctionOk() ([]EESFunctionSingle, bool) {
	if o == nil || isNil(o.EESFunction) {
		return nil, false
	}
	return o.EESFunction, true
}

// HasEESFunction returns a boolean if a field has been set.
func (o *EdgeDataNetworkSingleAllOf1) HasEESFunction() bool {
	if o != nil && !isNil(o.EESFunction) {
		return true
	}

	return false
}

// SetEESFunction gets a reference to the given []EESFunctionSingle and assigns it to the EESFunction field.
func (o *EdgeDataNetworkSingleAllOf1) SetEESFunction(v []EESFunctionSingle) {
	o.EESFunction = v
}

func (o EdgeDataNetworkSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeDataNetworkSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.EASFunction) {
		toSerialize["EASFunction"] = o.EASFunction
	}
	if !isNil(o.EESFunction) {
		toSerialize["EESFunction"] = o.EESFunction
	}
	return toSerialize, nil
}

type NullableEdgeDataNetworkSingleAllOf1 struct {
	value *EdgeDataNetworkSingleAllOf1
	isSet bool
}

func (v NullableEdgeDataNetworkSingleAllOf1) Get() *EdgeDataNetworkSingleAllOf1 {
	return v.value
}

func (v *NullableEdgeDataNetworkSingleAllOf1) Set(val *EdgeDataNetworkSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeDataNetworkSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeDataNetworkSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeDataNetworkSingleAllOf1(val *EdgeDataNetworkSingleAllOf1) *NullableEdgeDataNetworkSingleAllOf1 {
	return &NullableEdgeDataNetworkSingleAllOf1{value: val, isSet: true}
}

func (v NullableEdgeDataNetworkSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeDataNetworkSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


