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

// checks if the GtpUPathQoSMonitoringControlSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GtpUPathQoSMonitoringControlSingleAllOf{}

// GtpUPathQoSMonitoringControlSingleAllOf struct for GtpUPathQoSMonitoringControlSingleAllOf
type GtpUPathQoSMonitoringControlSingleAllOf struct {
	Attributes *interface{} `json:"attributes,omitempty"`
}

// NewGtpUPathQoSMonitoringControlSingleAllOf instantiates a new GtpUPathQoSMonitoringControlSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGtpUPathQoSMonitoringControlSingleAllOf() *GtpUPathQoSMonitoringControlSingleAllOf {
	this := GtpUPathQoSMonitoringControlSingleAllOf{}
	return &this
}

// NewGtpUPathQoSMonitoringControlSingleAllOfWithDefaults instantiates a new GtpUPathQoSMonitoringControlSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGtpUPathQoSMonitoringControlSingleAllOfWithDefaults() *GtpUPathQoSMonitoringControlSingleAllOf {
	this := GtpUPathQoSMonitoringControlSingleAllOf{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GtpUPathQoSMonitoringControlSingleAllOf) GetAttributes() interface{} {
	if o == nil || isNil(o.Attributes) {
		var ret interface{}
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOf) GetAttributesOk() (*interface{}, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GtpUPathQoSMonitoringControlSingleAllOf) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given interface{} and assigns it to the Attributes field.
func (o *GtpUPathQoSMonitoringControlSingleAllOf) SetAttributes(v interface{}) {
	o.Attributes = &v
}

func (o GtpUPathQoSMonitoringControlSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GtpUPathQoSMonitoringControlSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGtpUPathQoSMonitoringControlSingleAllOf struct {
	value *GtpUPathQoSMonitoringControlSingleAllOf
	isSet bool
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOf) Get() *GtpUPathQoSMonitoringControlSingleAllOf {
	return v.value
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOf) Set(val *GtpUPathQoSMonitoringControlSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGtpUPathQoSMonitoringControlSingleAllOf(val *GtpUPathQoSMonitoringControlSingleAllOf) *NullableGtpUPathQoSMonitoringControlSingleAllOf {
	return &NullableGtpUPathQoSMonitoringControlSingleAllOf{value: val, isSet: true}
}

func (v NullableGtpUPathQoSMonitoringControlSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGtpUPathQoSMonitoringControlSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


