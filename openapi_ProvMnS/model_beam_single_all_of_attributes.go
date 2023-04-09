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

// checks if the BeamSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BeamSingleAllOfAttributes{}

// BeamSingleAllOfAttributes struct for BeamSingleAllOfAttributes
type BeamSingleAllOfAttributes struct {
	BeamIndex      *int32  `json:"beamIndex,omitempty"`
	BeamType       *string `json:"beamType,omitempty"`
	BeamAzimuth    *int32  `json:"beamAzimuth,omitempty"`
	BeamTilt       *int32  `json:"beamTilt,omitempty"`
	BeamHorizWidth *int32  `json:"beamHorizWidth,omitempty"`
	BeamVertWidth  *int32  `json:"beamVertWidth,omitempty"`
}

// NewBeamSingleAllOfAttributes instantiates a new BeamSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBeamSingleAllOfAttributes() *BeamSingleAllOfAttributes {
	this := BeamSingleAllOfAttributes{}
	return &this
}

// NewBeamSingleAllOfAttributesWithDefaults instantiates a new BeamSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBeamSingleAllOfAttributesWithDefaults() *BeamSingleAllOfAttributes {
	this := BeamSingleAllOfAttributes{}
	return &this
}

// GetBeamIndex returns the BeamIndex field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributes) GetBeamIndex() int32 {
	if o == nil || IsNil(o.BeamIndex) {
		var ret int32
		return ret
	}
	return *o.BeamIndex
}

// GetBeamIndexOk returns a tuple with the BeamIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributes) GetBeamIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamIndex) {
		return nil, false
	}
	return o.BeamIndex, true
}

// HasBeamIndex returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributes) HasBeamIndex() bool {
	if o != nil && !IsNil(o.BeamIndex) {
		return true
	}

	return false
}

// SetBeamIndex gets a reference to the given int32 and assigns it to the BeamIndex field.
func (o *BeamSingleAllOfAttributes) SetBeamIndex(v int32) {
	o.BeamIndex = &v
}

// GetBeamType returns the BeamType field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributes) GetBeamType() string {
	if o == nil || IsNil(o.BeamType) {
		var ret string
		return ret
	}
	return *o.BeamType
}

// GetBeamTypeOk returns a tuple with the BeamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributes) GetBeamTypeOk() (*string, bool) {
	if o == nil || IsNil(o.BeamType) {
		return nil, false
	}
	return o.BeamType, true
}

// HasBeamType returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributes) HasBeamType() bool {
	if o != nil && !IsNil(o.BeamType) {
		return true
	}

	return false
}

// SetBeamType gets a reference to the given string and assigns it to the BeamType field.
func (o *BeamSingleAllOfAttributes) SetBeamType(v string) {
	o.BeamType = &v
}

// GetBeamAzimuth returns the BeamAzimuth field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributes) GetBeamAzimuth() int32 {
	if o == nil || IsNil(o.BeamAzimuth) {
		var ret int32
		return ret
	}
	return *o.BeamAzimuth
}

// GetBeamAzimuthOk returns a tuple with the BeamAzimuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributes) GetBeamAzimuthOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamAzimuth) {
		return nil, false
	}
	return o.BeamAzimuth, true
}

// HasBeamAzimuth returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributes) HasBeamAzimuth() bool {
	if o != nil && !IsNil(o.BeamAzimuth) {
		return true
	}

	return false
}

// SetBeamAzimuth gets a reference to the given int32 and assigns it to the BeamAzimuth field.
func (o *BeamSingleAllOfAttributes) SetBeamAzimuth(v int32) {
	o.BeamAzimuth = &v
}

// GetBeamTilt returns the BeamTilt field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributes) GetBeamTilt() int32 {
	if o == nil || IsNil(o.BeamTilt) {
		var ret int32
		return ret
	}
	return *o.BeamTilt
}

// GetBeamTiltOk returns a tuple with the BeamTilt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributes) GetBeamTiltOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamTilt) {
		return nil, false
	}
	return o.BeamTilt, true
}

// HasBeamTilt returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributes) HasBeamTilt() bool {
	if o != nil && !IsNil(o.BeamTilt) {
		return true
	}

	return false
}

// SetBeamTilt gets a reference to the given int32 and assigns it to the BeamTilt field.
func (o *BeamSingleAllOfAttributes) SetBeamTilt(v int32) {
	o.BeamTilt = &v
}

// GetBeamHorizWidth returns the BeamHorizWidth field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributes) GetBeamHorizWidth() int32 {
	if o == nil || IsNil(o.BeamHorizWidth) {
		var ret int32
		return ret
	}
	return *o.BeamHorizWidth
}

// GetBeamHorizWidthOk returns a tuple with the BeamHorizWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributes) GetBeamHorizWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamHorizWidth) {
		return nil, false
	}
	return o.BeamHorizWidth, true
}

// HasBeamHorizWidth returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributes) HasBeamHorizWidth() bool {
	if o != nil && !IsNil(o.BeamHorizWidth) {
		return true
	}

	return false
}

// SetBeamHorizWidth gets a reference to the given int32 and assigns it to the BeamHorizWidth field.
func (o *BeamSingleAllOfAttributes) SetBeamHorizWidth(v int32) {
	o.BeamHorizWidth = &v
}

// GetBeamVertWidth returns the BeamVertWidth field value if set, zero value otherwise.
func (o *BeamSingleAllOfAttributes) GetBeamVertWidth() int32 {
	if o == nil || IsNil(o.BeamVertWidth) {
		var ret int32
		return ret
	}
	return *o.BeamVertWidth
}

// GetBeamVertWidthOk returns a tuple with the BeamVertWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BeamSingleAllOfAttributes) GetBeamVertWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.BeamVertWidth) {
		return nil, false
	}
	return o.BeamVertWidth, true
}

// HasBeamVertWidth returns a boolean if a field has been set.
func (o *BeamSingleAllOfAttributes) HasBeamVertWidth() bool {
	if o != nil && !IsNil(o.BeamVertWidth) {
		return true
	}

	return false
}

// SetBeamVertWidth gets a reference to the given int32 and assigns it to the BeamVertWidth field.
func (o *BeamSingleAllOfAttributes) SetBeamVertWidth(v int32) {
	o.BeamVertWidth = &v
}

func (o BeamSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BeamSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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

type NullableBeamSingleAllOfAttributes struct {
	value *BeamSingleAllOfAttributes
	isSet bool
}

func (v NullableBeamSingleAllOfAttributes) Get() *BeamSingleAllOfAttributes {
	return v.value
}

func (v *NullableBeamSingleAllOfAttributes) Set(val *BeamSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBeamSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBeamSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBeamSingleAllOfAttributes(val *BeamSingleAllOfAttributes) *NullableBeamSingleAllOfAttributes {
	return &NullableBeamSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableBeamSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBeamSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
