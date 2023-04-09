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

// checks if the CommonBeamformingFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonBeamformingFunctionSingleAllOfAttributesAllOf{}

// CommonBeamformingFunctionSingleAllOfAttributesAllOf struct for CommonBeamformingFunctionSingleAllOfAttributesAllOf
type CommonBeamformingFunctionSingleAllOfAttributesAllOf struct {
	CoverageShape  *int32 `json:"coverageShape,omitempty"`
	DigitalAzimuth *int32 `json:"digitalAzimuth,omitempty"`
	DigitalTilt    *int32 `json:"digitalTilt,omitempty"`
}

// NewCommonBeamformingFunctionSingleAllOfAttributesAllOf instantiates a new CommonBeamformingFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBeamformingFunctionSingleAllOfAttributesAllOf() *CommonBeamformingFunctionSingleAllOfAttributesAllOf {
	this := CommonBeamformingFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewCommonBeamformingFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new CommonBeamformingFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBeamformingFunctionSingleAllOfAttributesAllOfWithDefaults() *CommonBeamformingFunctionSingleAllOfAttributesAllOf {
	this := CommonBeamformingFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetCoverageShape returns the CoverageShape field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) GetCoverageShape() int32 {
	if o == nil || IsNil(o.CoverageShape) {
		var ret int32
		return ret
	}
	return *o.CoverageShape
}

// GetCoverageShapeOk returns a tuple with the CoverageShape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) GetCoverageShapeOk() (*int32, bool) {
	if o == nil || IsNil(o.CoverageShape) {
		return nil, false
	}
	return o.CoverageShape, true
}

// HasCoverageShape returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) HasCoverageShape() bool {
	if o != nil && !IsNil(o.CoverageShape) {
		return true
	}

	return false
}

// SetCoverageShape gets a reference to the given int32 and assigns it to the CoverageShape field.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) SetCoverageShape(v int32) {
	o.CoverageShape = &v
}

// GetDigitalAzimuth returns the DigitalAzimuth field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) GetDigitalAzimuth() int32 {
	if o == nil || IsNil(o.DigitalAzimuth) {
		var ret int32
		return ret
	}
	return *o.DigitalAzimuth
}

// GetDigitalAzimuthOk returns a tuple with the DigitalAzimuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) GetDigitalAzimuthOk() (*int32, bool) {
	if o == nil || IsNil(o.DigitalAzimuth) {
		return nil, false
	}
	return o.DigitalAzimuth, true
}

// HasDigitalAzimuth returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) HasDigitalAzimuth() bool {
	if o != nil && !IsNil(o.DigitalAzimuth) {
		return true
	}

	return false
}

// SetDigitalAzimuth gets a reference to the given int32 and assigns it to the DigitalAzimuth field.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) SetDigitalAzimuth(v int32) {
	o.DigitalAzimuth = &v
}

// GetDigitalTilt returns the DigitalTilt field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) GetDigitalTilt() int32 {
	if o == nil || IsNil(o.DigitalTilt) {
		var ret int32
		return ret
	}
	return *o.DigitalTilt
}

// GetDigitalTiltOk returns a tuple with the DigitalTilt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) GetDigitalTiltOk() (*int32, bool) {
	if o == nil || IsNil(o.DigitalTilt) {
		return nil, false
	}
	return o.DigitalTilt, true
}

// HasDigitalTilt returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) HasDigitalTilt() bool {
	if o != nil && !IsNil(o.DigitalTilt) {
		return true
	}

	return false
}

// SetDigitalTilt gets a reference to the given int32 and assigns it to the DigitalTilt field.
func (o *CommonBeamformingFunctionSingleAllOfAttributesAllOf) SetDigitalTilt(v int32) {
	o.DigitalTilt = &v
}

func (o CommonBeamformingFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonBeamformingFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CoverageShape) {
		toSerialize["coverageShape"] = o.CoverageShape
	}
	if !IsNil(o.DigitalAzimuth) {
		toSerialize["digitalAzimuth"] = o.DigitalAzimuth
	}
	if !IsNil(o.DigitalTilt) {
		toSerialize["digitalTilt"] = o.DigitalTilt
	}
	return toSerialize, nil
}

type NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf struct {
	value *CommonBeamformingFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf) Get() *CommonBeamformingFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf) Set(val *CommonBeamformingFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBeamformingFunctionSingleAllOfAttributesAllOf(val *CommonBeamformingFunctionSingleAllOfAttributesAllOf) *NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf {
	return &NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBeamformingFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
