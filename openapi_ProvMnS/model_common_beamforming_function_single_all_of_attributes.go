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

// checks if the CommonBeamformingFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommonBeamformingFunctionSingleAllOfAttributes{}

// CommonBeamformingFunctionSingleAllOfAttributes struct for CommonBeamformingFunctionSingleAllOfAttributes
type CommonBeamformingFunctionSingleAllOfAttributes struct {
	CoverageShape  *int32 `json:"coverageShape,omitempty"`
	DigitalAzimuth *int32 `json:"digitalAzimuth,omitempty"`
	DigitalTilt    *int32 `json:"digitalTilt,omitempty"`
}

// NewCommonBeamformingFunctionSingleAllOfAttributes instantiates a new CommonBeamformingFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommonBeamformingFunctionSingleAllOfAttributes() *CommonBeamformingFunctionSingleAllOfAttributes {
	this := CommonBeamformingFunctionSingleAllOfAttributes{}
	return &this
}

// NewCommonBeamformingFunctionSingleAllOfAttributesWithDefaults instantiates a new CommonBeamformingFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommonBeamformingFunctionSingleAllOfAttributesWithDefaults() *CommonBeamformingFunctionSingleAllOfAttributes {
	this := CommonBeamformingFunctionSingleAllOfAttributes{}
	return &this
}

// GetCoverageShape returns the CoverageShape field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) GetCoverageShape() int32 {
	if o == nil || IsNil(o.CoverageShape) {
		var ret int32
		return ret
	}
	return *o.CoverageShape
}

// GetCoverageShapeOk returns a tuple with the CoverageShape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) GetCoverageShapeOk() (*int32, bool) {
	if o == nil || IsNil(o.CoverageShape) {
		return nil, false
	}
	return o.CoverageShape, true
}

// HasCoverageShape returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) HasCoverageShape() bool {
	if o != nil && !IsNil(o.CoverageShape) {
		return true
	}

	return false
}

// SetCoverageShape gets a reference to the given int32 and assigns it to the CoverageShape field.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) SetCoverageShape(v int32) {
	o.CoverageShape = &v
}

// GetDigitalAzimuth returns the DigitalAzimuth field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) GetDigitalAzimuth() int32 {
	if o == nil || IsNil(o.DigitalAzimuth) {
		var ret int32
		return ret
	}
	return *o.DigitalAzimuth
}

// GetDigitalAzimuthOk returns a tuple with the DigitalAzimuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) GetDigitalAzimuthOk() (*int32, bool) {
	if o == nil || IsNil(o.DigitalAzimuth) {
		return nil, false
	}
	return o.DigitalAzimuth, true
}

// HasDigitalAzimuth returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) HasDigitalAzimuth() bool {
	if o != nil && !IsNil(o.DigitalAzimuth) {
		return true
	}

	return false
}

// SetDigitalAzimuth gets a reference to the given int32 and assigns it to the DigitalAzimuth field.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) SetDigitalAzimuth(v int32) {
	o.DigitalAzimuth = &v
}

// GetDigitalTilt returns the DigitalTilt field value if set, zero value otherwise.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) GetDigitalTilt() int32 {
	if o == nil || IsNil(o.DigitalTilt) {
		var ret int32
		return ret
	}
	return *o.DigitalTilt
}

// GetDigitalTiltOk returns a tuple with the DigitalTilt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) GetDigitalTiltOk() (*int32, bool) {
	if o == nil || IsNil(o.DigitalTilt) {
		return nil, false
	}
	return o.DigitalTilt, true
}

// HasDigitalTilt returns a boolean if a field has been set.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) HasDigitalTilt() bool {
	if o != nil && !IsNil(o.DigitalTilt) {
		return true
	}

	return false
}

// SetDigitalTilt gets a reference to the given int32 and assigns it to the DigitalTilt field.
func (o *CommonBeamformingFunctionSingleAllOfAttributes) SetDigitalTilt(v int32) {
	o.DigitalTilt = &v
}

func (o CommonBeamformingFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommonBeamformingFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
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

type NullableCommonBeamformingFunctionSingleAllOfAttributes struct {
	value *CommonBeamformingFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableCommonBeamformingFunctionSingleAllOfAttributes) Get() *CommonBeamformingFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableCommonBeamformingFunctionSingleAllOfAttributes) Set(val *CommonBeamformingFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCommonBeamformingFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCommonBeamformingFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommonBeamformingFunctionSingleAllOfAttributes(val *CommonBeamformingFunctionSingleAllOfAttributes) *NullableCommonBeamformingFunctionSingleAllOfAttributes {
	return &NullableCommonBeamformingFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableCommonBeamformingFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommonBeamformingFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
