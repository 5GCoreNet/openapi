/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
)

// checks if the Local2dPointUncertaintyEllipseAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Local2dPointUncertaintyEllipseAllOf{}

// Local2dPointUncertaintyEllipseAllOf struct for Local2dPointUncertaintyEllipseAllOf
type Local2dPointUncertaintyEllipseAllOf struct {
	LocalOrigin LocalOrigin `json:"localOrigin"`
	Point RelativeCartesianLocation `json:"point"`
	UncertaintyEllipse UncertaintyEllipse `json:"uncertaintyEllipse"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewLocal2dPointUncertaintyEllipseAllOf instantiates a new Local2dPointUncertaintyEllipseAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocal2dPointUncertaintyEllipseAllOf(localOrigin LocalOrigin, point RelativeCartesianLocation, uncertaintyEllipse UncertaintyEllipse, confidence int32) *Local2dPointUncertaintyEllipseAllOf {
	this := Local2dPointUncertaintyEllipseAllOf{}
	this.LocalOrigin = localOrigin
	this.Point = point
	this.UncertaintyEllipse = uncertaintyEllipse
	this.Confidence = confidence
	return &this
}

// NewLocal2dPointUncertaintyEllipseAllOfWithDefaults instantiates a new Local2dPointUncertaintyEllipseAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocal2dPointUncertaintyEllipseAllOfWithDefaults() *Local2dPointUncertaintyEllipseAllOf {
	this := Local2dPointUncertaintyEllipseAllOf{}
	return &this
}

// GetLocalOrigin returns the LocalOrigin field value
func (o *Local2dPointUncertaintyEllipseAllOf) GetLocalOrigin() LocalOrigin {
	if o == nil {
		var ret LocalOrigin
		return ret
	}

	return o.LocalOrigin
}

// GetLocalOriginOk returns a tuple with the LocalOrigin field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipseAllOf) GetLocalOriginOk() (*LocalOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalOrigin, true
}

// SetLocalOrigin sets field value
func (o *Local2dPointUncertaintyEllipseAllOf) SetLocalOrigin(v LocalOrigin) {
	o.LocalOrigin = v
}

// GetPoint returns the Point field value
func (o *Local2dPointUncertaintyEllipseAllOf) GetPoint() RelativeCartesianLocation {
	if o == nil {
		var ret RelativeCartesianLocation
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipseAllOf) GetPointOk() (*RelativeCartesianLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *Local2dPointUncertaintyEllipseAllOf) SetPoint(v RelativeCartesianLocation) {
	o.Point = v
}

// GetUncertaintyEllipse returns the UncertaintyEllipse field value
func (o *Local2dPointUncertaintyEllipseAllOf) GetUncertaintyEllipse() UncertaintyEllipse {
	if o == nil {
		var ret UncertaintyEllipse
		return ret
	}

	return o.UncertaintyEllipse
}

// GetUncertaintyEllipseOk returns a tuple with the UncertaintyEllipse field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipseAllOf) GetUncertaintyEllipseOk() (*UncertaintyEllipse, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncertaintyEllipse, true
}

// SetUncertaintyEllipse sets field value
func (o *Local2dPointUncertaintyEllipseAllOf) SetUncertaintyEllipse(v UncertaintyEllipse) {
	o.UncertaintyEllipse = v
}

// GetConfidence returns the Confidence field value
func (o *Local2dPointUncertaintyEllipseAllOf) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *Local2dPointUncertaintyEllipseAllOf) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *Local2dPointUncertaintyEllipseAllOf) SetConfidence(v int32) {
	o.Confidence = v
}

func (o Local2dPointUncertaintyEllipseAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Local2dPointUncertaintyEllipseAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["localOrigin"] = o.LocalOrigin
	toSerialize["point"] = o.Point
	toSerialize["uncertaintyEllipse"] = o.UncertaintyEllipse
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

type NullableLocal2dPointUncertaintyEllipseAllOf struct {
	value *Local2dPointUncertaintyEllipseAllOf
	isSet bool
}

func (v NullableLocal2dPointUncertaintyEllipseAllOf) Get() *Local2dPointUncertaintyEllipseAllOf {
	return v.value
}

func (v *NullableLocal2dPointUncertaintyEllipseAllOf) Set(val *Local2dPointUncertaintyEllipseAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocal2dPointUncertaintyEllipseAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocal2dPointUncertaintyEllipseAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocal2dPointUncertaintyEllipseAllOf(val *Local2dPointUncertaintyEllipseAllOf) *NullableLocal2dPointUncertaintyEllipseAllOf {
	return &NullableLocal2dPointUncertaintyEllipseAllOf{value: val, isSet: true}
}

func (v NullableLocal2dPointUncertaintyEllipseAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocal2dPointUncertaintyEllipseAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


