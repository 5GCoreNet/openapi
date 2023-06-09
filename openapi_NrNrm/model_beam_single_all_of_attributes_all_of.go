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

// checks if the BeamSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeamSingleAllOfAttributesAllOf{}

// BeamSingleAllOfAttributesAllOf struct for BeamSingleAllOfAttributesAllOf
type BeamSingleAllOfAttributesAllOf struct {
	BeamIndex      *int32  `json:"beamIndex,omitempty"`
	BeamType       *string `json:"beamType,omitempty"`
	BeamAzimuth    *int32  `json:"beamAzimuth,omitempty"`
	BeamTilt       *int32  `json:"beamTilt,omitempty"`
	BeamHorizWidth *int32  `json:"beamHorizWidth,omitempty"`
	BeamVertWidth  *int32  `json:"beamVertWidth,omitempty"`
}

// NewBeamSingleAllOfAttributesAllOf instantiates a new BeamSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeamSingleAllOfAttributesAllOf() *BeamSingleAllOfAttributesAllOf {
	this := BeamSingleAllOfAttributesAllOf{}
	return &this
}

// NewBeamSingleAllOfAttributesAllOfWithDefaults instantiates a new BeamSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeamSingleAllOfAttributesAllOfWithDefaults() *BeamSingleAllOfAttributesAllOf {
	this := BeamSingleAllOfAttributesAllOf{}
	return &this
}

// GetBeamIndex returns the BeamIndex field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamIndex() int32 {
	if o == nil || IsNil(o.BeamIndex) {
		var ret int32
		return ret
	}
	return *o.BeamIndex
}

// GetBeamIndexOk returns a tuple with the BeamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamIndex) {
		return nil, false
	}
	return o.BeamIndex, true
}

// HasBeamIndex returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributesAllOf) HasBeamIndex() bool {
	if o != nil && !IsNil(o.BeamIndex) {
		return true
	}

	return false
}

// SetBeamIndex gets a reference to the given int32 and assigns it to the BeamIndex field.
func (o *BeamSingleAllOfAttributesAllOf) SetBeamIndex(v int32) {
	o.BeamIndex = &v
}

// GetBeamType returns the BeamType field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamType() string {
	if o == nil || IsNil(o.BeamType) {
		var ret string
		return ret
	}
	return *o.BeamType
}

// GetBeamTypeOk returns a tuple with the BeamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BeamType) {
		return nil, false
	}
	return o.BeamType, true
}

// HasBeamType returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributesAllOf) HasBeamType() bool {
	if o != nil && !IsNil(o.BeamType) {
		return true
	}

	return false
}

// SetBeamType gets a reference to the given string and assigns it to the BeamType field.
func (o *BeamSingleAllOfAttributesAllOf) SetBeamType(v string) {
	o.BeamType = &v
}

// GetBeamAzimuth returns the BeamAzimuth field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamAzimuth() int32 {
	if o == nil || IsNil(o.BeamAzimuth) {
		var ret int32
		return ret
	}
	return *o.BeamAzimuth
}

// GetBeamAzimuthOk returns a tuple with the BeamAzimuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamAzimuthOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamAzimuth) {
		return nil, false
	}
	return o.BeamAzimuth, true
}

// HasBeamAzimuth returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributesAllOf) HasBeamAzimuth() bool {
	if o != nil && !IsNil(o.BeamAzimuth) {
		return true
	}

	return false
}

// SetBeamAzimuth gets a reference to the given int32 and assigns it to the BeamAzimuth field.
func (o *BeamSingleAllOfAttributesAllOf) SetBeamAzimuth(v int32) {
	o.BeamAzimuth = &v
}

// GetBeamTilt returns the BeamTilt field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamTilt() int32 {
	if o == nil || IsNil(o.BeamTilt) {
		var ret int32
		return ret
	}
	return *o.BeamTilt
}

// GetBeamTiltOk returns a tuple with the BeamTilt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamTiltOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamTilt) {
		return nil, false
	}
	return o.BeamTilt, true
}

// HasBeamTilt returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributesAllOf) HasBeamTilt() bool {
	if o != nil && !IsNil(o.BeamTilt) {
		return true
	}

	return false
}

// SetBeamTilt gets a reference to the given int32 and assigns it to the BeamTilt field.
func (o *BeamSingleAllOfAttributesAllOf) SetBeamTilt(v int32) {
	o.BeamTilt = &v
}

// GetBeamHorizWidth returns the BeamHorizWidth field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamHorizWidth() int32 {
	if o == nil || IsNil(o.BeamHorizWidth) {
		var ret int32
		return ret
	}
	return *o.BeamHorizWidth
}

// GetBeamHorizWidthOk returns a tuple with the BeamHorizWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamHorizWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamHorizWidth) {
		return nil, false
	}
	return o.BeamHorizWidth, true
}

// HasBeamHorizWidth returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributesAllOf) HasBeamHorizWidth() bool {
	if o != nil && !IsNil(o.BeamHorizWidth) {
		return true
	}

	return false
}

// SetBeamHorizWidth gets a reference to the given int32 and assigns it to the BeamHorizWidth field.
func (o *BeamSingleAllOfAttributesAllOf) SetBeamHorizWidth(v int32) {
	o.BeamHorizWidth = &v
}

// GetBeamVertWidth returns the BeamVertWidth field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamVertWidth() int32 {
	if o == nil || IsNil(o.BeamVertWidth) {
		var ret int32
		return ret
	}
	return *o.BeamVertWidth
}

// GetBeamVertWidthOk returns a tuple with the BeamVertWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributesAllOf) GetBeamVertWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamVertWidth) {
		return nil, false
	}
	return o.BeamVertWidth, true
}

// HasBeamVertWidth returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributesAllOf) HasBeamVertWidth() bool {
	if o != nil && !IsNil(o.BeamVertWidth) {
		return true
	}

	return false
}

// SetBeamVertWidth gets a reference to the given int32 and assigns it to the BeamVertWidth field.
func (o *BeamSingleAllOfAttributesAllOf) SetBeamVertWidth(v int32) {
	o.BeamVertWidth = &v
}

func (o BeamSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeamSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BeamIndex) {
		toSerialize["beamIndex"] = o.BeamIndex
	}
	if !IsNil(o.BeamType) {
		toSerialize["beamType"] = o.BeamType
	}
	if !IsNil(o.BeamAzimuth) {
		toSerialize["beamAzimuth"] = o.BeamAzimuth
	}
	if !IsNil(o.BeamTilt) {
		toSerialize["beamTilt"] = o.BeamTilt
	}
	if !IsNil(o.BeamHorizWidth) {
		toSerialize["beamHorizWidth"] = o.BeamHorizWidth
	}
	if !IsNil(o.BeamVertWidth) {
		toSerialize["beamVertWidth"] = o.BeamVertWidth
	}
	return toSerialize, nil
}

type NullableBeamSingleAllOfAttributesAllOf struct {
	value *BeamSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableBeamSingleAllOfAttributesAllOf) Get() *BeamSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableBeamSingleAllOfAttributesAllOf) Set(val *BeamSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableBeamSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableBeamSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeamSingleAllOfAttributesAllOf(val *BeamSingleAllOfAttributesAllOf) *NullableBeamSingleAllOfAttributesAllOf {
	return &NullableBeamSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableBeamSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeamSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
