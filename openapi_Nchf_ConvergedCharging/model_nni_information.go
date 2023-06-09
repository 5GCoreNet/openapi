/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the NNIInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NNIInformation{}

// NNIInformation struct for NNIInformation
type NNIInformation struct {
	SessionDirection     *NNISessionDirection `json:"sessionDirection,omitempty"`
	NNIType              *NNIType             `json:"nNIType,omitempty"`
	RelationshipMode     *NNIRelationshipMode `json:"relationshipMode,omitempty"`
	NeighbourNodeAddress *IMSAddress          `json:"neighbourNodeAddress,omitempty"`
}

// NewNNIInformation instantiates a new NNIInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNNIInformation() *NNIInformation {
	this := NNIInformation{}
	return &this
}

// NewNNIInformationWithDefaults instantiates a new NNIInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNNIInformationWithDefaults() *NNIInformation {
	this := NNIInformation{}
	return &this
}

// GetSessionDirection returns the SessionDirection field value if set, zero value otherwise.
func (o *NNIInformation) GetSessionDirection() NNISessionDirection {
	if o == nil || IsNil(o.SessionDirection) {
		var ret NNISessionDirection
		return ret
	}
	return *o.SessionDirection
}

// GetSessionDirectionOk returns a tuple with the SessionDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NNIInformation) GetSessionDirectionOk() (*NNISessionDirection, bool) {
	if o == nil || IsNil(o.SessionDirection) {
		return nil, false
	}
	return o.SessionDirection, true
}

// HasSessionDirection returns a boolean if a field has been set.
func (o *NNIInformation) HasSessionDirection() bool {
	if o != nil && !IsNil(o.SessionDirection) {
		return true
	}

	return false
}

// SetSessionDirection gets a reference to the given NNISessionDirection and assigns it to the SessionDirection field.
func (o *NNIInformation) SetSessionDirection(v NNISessionDirection) {
	o.SessionDirection = &v
}

// GetNNIType returns the NNIType field value if set, zero value otherwise.
func (o *NNIInformation) GetNNIType() NNIType {
	if o == nil || IsNil(o.NNIType) {
		var ret NNIType
		return ret
	}
	return *o.NNIType
}

// GetNNITypeOk returns a tuple with the NNIType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NNIInformation) GetNNITypeOk() (*NNIType, bool) {
	if o == nil || IsNil(o.NNIType) {
		return nil, false
	}
	return o.NNIType, true
}

// HasNNIType returns a boolean if a field has been set.
func (o *NNIInformation) HasNNIType() bool {
	if o != nil && !IsNil(o.NNIType) {
		return true
	}

	return false
}

// SetNNIType gets a reference to the given NNIType and assigns it to the NNIType field.
func (o *NNIInformation) SetNNIType(v NNIType) {
	o.NNIType = &v
}

// GetRelationshipMode returns the RelationshipMode field value if set, zero value otherwise.
func (o *NNIInformation) GetRelationshipMode() NNIRelationshipMode {
	if o == nil || IsNil(o.RelationshipMode) {
		var ret NNIRelationshipMode
		return ret
	}
	return *o.RelationshipMode
}

// GetRelationshipModeOk returns a tuple with the RelationshipMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NNIInformation) GetRelationshipModeOk() (*NNIRelationshipMode, bool) {
	if o == nil || IsNil(o.RelationshipMode) {
		return nil, false
	}
	return o.RelationshipMode, true
}

// HasRelationshipMode returns a boolean if a field has been set.
func (o *NNIInformation) HasRelationshipMode() bool {
	if o != nil && !IsNil(o.RelationshipMode) {
		return true
	}

	return false
}

// SetRelationshipMode gets a reference to the given NNIRelationshipMode and assigns it to the RelationshipMode field.
func (o *NNIInformation) SetRelationshipMode(v NNIRelationshipMode) {
	o.RelationshipMode = &v
}

// GetNeighbourNodeAddress returns the NeighbourNodeAddress field value if set, zero value otherwise.
func (o *NNIInformation) GetNeighbourNodeAddress() IMSAddress {
	if o == nil || IsNil(o.NeighbourNodeAddress) {
		var ret IMSAddress
		return ret
	}
	return *o.NeighbourNodeAddress
}

// GetNeighbourNodeAddressOk returns a tuple with the NeighbourNodeAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NNIInformation) GetNeighbourNodeAddressOk() (*IMSAddress, bool) {
	if o == nil || IsNil(o.NeighbourNodeAddress) {
		return nil, false
	}
	return o.NeighbourNodeAddress, true
}

// HasNeighbourNodeAddress returns a boolean if a field has been set.
func (o *NNIInformation) HasNeighbourNodeAddress() bool {
	if o != nil && !IsNil(o.NeighbourNodeAddress) {
		return true
	}

	return false
}

// SetNeighbourNodeAddress gets a reference to the given IMSAddress and assigns it to the NeighbourNodeAddress field.
func (o *NNIInformation) SetNeighbourNodeAddress(v IMSAddress) {
	o.NeighbourNodeAddress = &v
}

func (o NNIInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NNIInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SessionDirection) {
		toSerialize["sessionDirection"] = o.SessionDirection
	}
	if !IsNil(o.NNIType) {
		toSerialize["nNIType"] = o.NNIType
	}
	if !IsNil(o.RelationshipMode) {
		toSerialize["relationshipMode"] = o.RelationshipMode
	}
	if !IsNil(o.NeighbourNodeAddress) {
		toSerialize["neighbourNodeAddress"] = o.NeighbourNodeAddress
	}
	return toSerialize, nil
}

type NullableNNIInformation struct {
	value *NNIInformation
	isSet bool
}

func (v NullableNNIInformation) Get() *NNIInformation {
	return v.value
}

func (v *NullableNNIInformation) Set(val *NNIInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableNNIInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableNNIInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNNIInformation(val *NNIInformation) *NullableNNIInformation {
	return &NullableNNIInformation{value: val, isSet: true}
}

func (v NullableNNIInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNNIInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
