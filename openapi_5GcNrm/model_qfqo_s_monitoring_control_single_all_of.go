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

// checks if the QFQoSMonitoringControlSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QFQoSMonitoringControlSingleAllOf{}

// QFQoSMonitoringControlSingleAllOf struct for QFQoSMonitoringControlSingleAllOf
type QFQoSMonitoringControlSingleAllOf struct {
	Attributes *QFQoSMonitoringControlSingleAllOfAttributes `json:"attributes,omitempty"`
}

// NewQFQoSMonitoringControlSingleAllOf instantiates a new QFQoSMonitoringControlSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQFQoSMonitoringControlSingleAllOf() *QFQoSMonitoringControlSingleAllOf {
	this := QFQoSMonitoringControlSingleAllOf{}
	return &this
}

// NewQFQoSMonitoringControlSingleAllOfWithDefaults instantiates a new QFQoSMonitoringControlSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQFQoSMonitoringControlSingleAllOfWithDefaults() *QFQoSMonitoringControlSingleAllOf {
	this := QFQoSMonitoringControlSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *QFQoSMonitoringControlSingleAllOf) GetAttributes() QFQoSMonitoringControlSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret QFQoSMonitoringControlSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QFQoSMonitoringControlSingleAllOf) GetAttributesOk() (*QFQoSMonitoringControlSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *QFQoSMonitoringControlSingleAllOf) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given QFQoSMonitoringControlSingleAllOfAttributes and assigns it to the Attributes field.
func (o *QFQoSMonitoringControlSingleAllOf) SetAttributes(v QFQoSMonitoringControlSingleAllOfAttributes) {
	o.Attributes = &v
}

func (o QFQoSMonitoringControlSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QFQoSMonitoringControlSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableQFQoSMonitoringControlSingleAllOf struct {
	value *QFQoSMonitoringControlSingleAllOf
	isSet bool
}

func (v NullableQFQoSMonitoringControlSingleAllOf) Get() *QFQoSMonitoringControlSingleAllOf {
	return v.value
}

func (v *NullableQFQoSMonitoringControlSingleAllOf) Set(val *QFQoSMonitoringControlSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableQFQoSMonitoringControlSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableQFQoSMonitoringControlSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQFQoSMonitoringControlSingleAllOf(val *QFQoSMonitoringControlSingleAllOf) *NullableQFQoSMonitoringControlSingleAllOf {
	return &NullableQFQoSMonitoringControlSingleAllOf{value: val, isSet: true}
}

func (v NullableQFQoSMonitoringControlSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQFQoSMonitoringControlSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
