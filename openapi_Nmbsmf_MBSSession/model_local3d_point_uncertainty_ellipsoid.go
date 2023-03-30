/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
)

// checks if the Local3dPointUncertaintyEllipsoid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Local3dPointUncertaintyEllipsoid{}

// Local3dPointUncertaintyEllipsoid Local 3D point with uncertainty ellipsoid
type Local3dPointUncertaintyEllipsoid struct {
	GADShape
	LocalOrigin LocalOrigin `json:"localOrigin"`
	Point RelativeCartesianLocation `json:"point"`
	UncertaintyEllipsoid UncertaintyEllipsoid `json:"uncertaintyEllipsoid"`
	// Indicates value of confidence.
	Confidence int32 `json:"confidence"`
}

// NewLocal3dPointUncertaintyEllipsoid instantiates a new Local3dPointUncertaintyEllipsoid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocal3dPointUncertaintyEllipsoid(localOrigin LocalOrigin, point RelativeCartesianLocation, uncertaintyEllipsoid UncertaintyEllipsoid, confidence int32, shape SupportedGADShapes) *Local3dPointUncertaintyEllipsoid {
	this := Local3dPointUncertaintyEllipsoid{}
	this.Shape = shape
	this.LocalOrigin = localOrigin
	this.Point = point
	this.UncertaintyEllipsoid = uncertaintyEllipsoid
	this.Confidence = confidence
	return &this
}

// NewLocal3dPointUncertaintyEllipsoidWithDefaults instantiates a new Local3dPointUncertaintyEllipsoid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocal3dPointUncertaintyEllipsoidWithDefaults() *Local3dPointUncertaintyEllipsoid {
	this := Local3dPointUncertaintyEllipsoid{}
	return &this
}

// GetLocalOrigin returns the LocalOrigin field value
func (o *Local3dPointUncertaintyEllipsoid) GetLocalOrigin() LocalOrigin {
	if o == nil {
		var ret LocalOrigin
		return ret
	}

	return o.LocalOrigin
}

// GetLocalOriginOk returns a tuple with the LocalOrigin field value
// and a boolean to check if the value has been set.
func (o *Local3dPointUncertaintyEllipsoid) GetLocalOriginOk() (*LocalOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LocalOrigin, true
}

// SetLocalOrigin sets field value
func (o *Local3dPointUncertaintyEllipsoid) SetLocalOrigin(v LocalOrigin) {
	o.LocalOrigin = v
}

// GetPoint returns the Point field value
func (o *Local3dPointUncertaintyEllipsoid) GetPoint() RelativeCartesianLocation {
	if o == nil {
		var ret RelativeCartesianLocation
		return ret
	}

	return o.Point
}

// GetPointOk returns a tuple with the Point field value
// and a boolean to check if the value has been set.
func (o *Local3dPointUncertaintyEllipsoid) GetPointOk() (*RelativeCartesianLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Point, true
}

// SetPoint sets field value
func (o *Local3dPointUncertaintyEllipsoid) SetPoint(v RelativeCartesianLocation) {
	o.Point = v
}

// GetUncertaintyEllipsoid returns the UncertaintyEllipsoid field value
func (o *Local3dPointUncertaintyEllipsoid) GetUncertaintyEllipsoid() UncertaintyEllipsoid {
	if o == nil {
		var ret UncertaintyEllipsoid
		return ret
	}

	return o.UncertaintyEllipsoid
}

// GetUncertaintyEllipsoidOk returns a tuple with the UncertaintyEllipsoid field value
// and a boolean to check if the value has been set.
func (o *Local3dPointUncertaintyEllipsoid) GetUncertaintyEllipsoidOk() (*UncertaintyEllipsoid, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UncertaintyEllipsoid, true
}

// SetUncertaintyEllipsoid sets field value
func (o *Local3dPointUncertaintyEllipsoid) SetUncertaintyEllipsoid(v UncertaintyEllipsoid) {
	o.UncertaintyEllipsoid = v
}

// GetConfidence returns the Confidence field value
func (o *Local3dPointUncertaintyEllipsoid) GetConfidence() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value
// and a boolean to check if the value has been set.
func (o *Local3dPointUncertaintyEllipsoid) GetConfidenceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Confidence, true
}

// SetConfidence sets field value
func (o *Local3dPointUncertaintyEllipsoid) SetConfidence(v int32) {
	o.Confidence = v
}

func (o Local3dPointUncertaintyEllipsoid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Local3dPointUncertaintyEllipsoid) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedGADShape, errGADShape := json.Marshal(o.GADShape)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	errGADShape = json.Unmarshal([]byte(serializedGADShape), &toSerialize)
	if errGADShape != nil {
		return map[string]interface{}{}, errGADShape
	}
	toSerialize["localOrigin"] = o.LocalOrigin
	toSerialize["point"] = o.Point
	toSerialize["uncertaintyEllipsoid"] = o.UncertaintyEllipsoid
	toSerialize["confidence"] = o.Confidence
	return toSerialize, nil
}

type NullableLocal3dPointUncertaintyEllipsoid struct {
	value *Local3dPointUncertaintyEllipsoid
	isSet bool
}

func (v NullableLocal3dPointUncertaintyEllipsoid) Get() *Local3dPointUncertaintyEllipsoid {
	return v.value
}

func (v *NullableLocal3dPointUncertaintyEllipsoid) Set(val *Local3dPointUncertaintyEllipsoid) {
	v.value = val
	v.isSet = true
}

func (v NullableLocal3dPointUncertaintyEllipsoid) IsSet() bool {
	return v.isSet
}

func (v *NullableLocal3dPointUncertaintyEllipsoid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocal3dPointUncertaintyEllipsoid(val *Local3dPointUncertaintyEllipsoid) *NullableLocal3dPointUncertaintyEllipsoid {
	return &NullableLocal3dPointUncertaintyEllipsoid{value: val, isSet: true}
}

func (v NullableLocal3dPointUncertaintyEllipsoid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocal3dPointUncertaintyEllipsoid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


