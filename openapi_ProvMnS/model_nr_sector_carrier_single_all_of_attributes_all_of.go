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

// checks if the NrSectorCarrierSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrSectorCarrierSingleAllOfAttributesAllOf{}

// NrSectorCarrierSingleAllOfAttributesAllOf struct for NrSectorCarrierSingleAllOfAttributesAllOf
type NrSectorCarrierSingleAllOfAttributesAllOf struct {
	TxDirection                *TxDirection `json:"txDirection,omitempty"`
	ConfiguredMaxTxPower       *int32       `json:"configuredMaxTxPower,omitempty"`
	ArfcnDL                    *int32       `json:"arfcnDL,omitempty"`
	ArfcnUL                    *int32       `json:"arfcnUL,omitempty"`
	BSChannelBwDL              *int32       `json:"bSChannelBwDL,omitempty"`
	BSChannelBwUL              *int32       `json:"bSChannelBwUL,omitempty"`
	SectorEquipmentFunctionRef *string      `json:"sectorEquipmentFunctionRef,omitempty"`
}

// NewNrSectorCarrierSingleAllOfAttributesAllOf instantiates a new NrSectorCarrierSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrSectorCarrierSingleAllOfAttributesAllOf() *NrSectorCarrierSingleAllOfAttributesAllOf {
	this := NrSectorCarrierSingleAllOfAttributesAllOf{}
	return &this
}

// NewNrSectorCarrierSingleAllOfAttributesAllOfWithDefaults instantiates a new NrSectorCarrierSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrSectorCarrierSingleAllOfAttributesAllOfWithDefaults() *NrSectorCarrierSingleAllOfAttributesAllOf {
	this := NrSectorCarrierSingleAllOfAttributesAllOf{}
	return &this
}

// GetTxDirection returns the TxDirection field value if set, zero value otherwise.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetTxDirection() TxDirection {
	if o == nil || IsNil(o.TxDirection) {
		var ret TxDirection
		return ret
	}
	return *o.TxDirection
}

// GetTxDirectionOk returns a tuple with the TxDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetTxDirectionOk() (*TxDirection, bool) {
	if o == nil || IsNil(o.TxDirection) {
		return nil, false
	}
	return o.TxDirection, true
}

// HasTxDirection returns a boolean if a field has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) HasTxDirection() bool {
	if o != nil && !IsNil(o.TxDirection) {
		return true
	}

	return false
}

// SetTxDirection gets a reference to the given TxDirection and assigns it to the TxDirection field.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) SetTxDirection(v TxDirection) {
	o.TxDirection = &v
}

// GetConfiguredMaxTxPower returns the ConfiguredMaxTxPower field value if set, zero value otherwise.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetConfiguredMaxTxPower() int32 {
	if o == nil || IsNil(o.ConfiguredMaxTxPower) {
		var ret int32
		return ret
	}
	return *o.ConfiguredMaxTxPower
}

// GetConfiguredMaxTxPowerOk returns a tuple with the ConfiguredMaxTxPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetConfiguredMaxTxPowerOk() (*int32, bool) {
	if o == nil || IsNil(o.ConfiguredMaxTxPower) {
		return nil, false
	}
	return o.ConfiguredMaxTxPower, true
}

// HasConfiguredMaxTxPower returns a boolean if a field has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) HasConfiguredMaxTxPower() bool {
	if o != nil && !IsNil(o.ConfiguredMaxTxPower) {
		return true
	}

	return false
}

// SetConfiguredMaxTxPower gets a reference to the given int32 and assigns it to the ConfiguredMaxTxPower field.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) SetConfiguredMaxTxPower(v int32) {
	o.ConfiguredMaxTxPower = &v
}

// GetArfcnDL returns the ArfcnDL field value if set, zero value otherwise.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetArfcnDL() int32 {
	if o == nil || IsNil(o.ArfcnDL) {
		var ret int32
		return ret
	}
	return *o.ArfcnDL
}

// GetArfcnDLOk returns a tuple with the ArfcnDL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetArfcnDLOk() (*int32, bool) {
	if o == nil || IsNil(o.ArfcnDL) {
		return nil, false
	}
	return o.ArfcnDL, true
}

// HasArfcnDL returns a boolean if a field has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) HasArfcnDL() bool {
	if o != nil && !IsNil(o.ArfcnDL) {
		return true
	}

	return false
}

// SetArfcnDL gets a reference to the given int32 and assigns it to the ArfcnDL field.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) SetArfcnDL(v int32) {
	o.ArfcnDL = &v
}

// GetArfcnUL returns the ArfcnUL field value if set, zero value otherwise.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetArfcnUL() int32 {
	if o == nil || IsNil(o.ArfcnUL) {
		var ret int32
		return ret
	}
	return *o.ArfcnUL
}

// GetArfcnULOk returns a tuple with the ArfcnUL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetArfcnULOk() (*int32, bool) {
	if o == nil || IsNil(o.ArfcnUL) {
		return nil, false
	}
	return o.ArfcnUL, true
}

// HasArfcnUL returns a boolean if a field has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) HasArfcnUL() bool {
	if o != nil && !IsNil(o.ArfcnUL) {
		return true
	}

	return false
}

// SetArfcnUL gets a reference to the given int32 and assigns it to the ArfcnUL field.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) SetArfcnUL(v int32) {
	o.ArfcnUL = &v
}

// GetBSChannelBwDL returns the BSChannelBwDL field value if set, zero value otherwise.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetBSChannelBwDL() int32 {
	if o == nil || IsNil(o.BSChannelBwDL) {
		var ret int32
		return ret
	}
	return *o.BSChannelBwDL
}

// GetBSChannelBwDLOk returns a tuple with the BSChannelBwDL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetBSChannelBwDLOk() (*int32, bool) {
	if o == nil || IsNil(o.BSChannelBwDL) {
		return nil, false
	}
	return o.BSChannelBwDL, true
}

// HasBSChannelBwDL returns a boolean if a field has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) HasBSChannelBwDL() bool {
	if o != nil && !IsNil(o.BSChannelBwDL) {
		return true
	}

	return false
}

// SetBSChannelBwDL gets a reference to the given int32 and assigns it to the BSChannelBwDL field.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) SetBSChannelBwDL(v int32) {
	o.BSChannelBwDL = &v
}

// GetBSChannelBwUL returns the BSChannelBwUL field value if set, zero value otherwise.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetBSChannelBwUL() int32 {
	if o == nil || IsNil(o.BSChannelBwUL) {
		var ret int32
		return ret
	}
	return *o.BSChannelBwUL
}

// GetBSChannelBwULOk returns a tuple with the BSChannelBwUL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetBSChannelBwULOk() (*int32, bool) {
	if o == nil || IsNil(o.BSChannelBwUL) {
		return nil, false
	}
	return o.BSChannelBwUL, true
}

// HasBSChannelBwUL returns a boolean if a field has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) HasBSChannelBwUL() bool {
	if o != nil && !IsNil(o.BSChannelBwUL) {
		return true
	}

	return false
}

// SetBSChannelBwUL gets a reference to the given int32 and assigns it to the BSChannelBwUL field.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) SetBSChannelBwUL(v int32) {
	o.BSChannelBwUL = &v
}

// GetSectorEquipmentFunctionRef returns the SectorEquipmentFunctionRef field value if set, zero value otherwise.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetSectorEquipmentFunctionRef() string {
	if o == nil || IsNil(o.SectorEquipmentFunctionRef) {
		var ret string
		return ret
	}
	return *o.SectorEquipmentFunctionRef
}

// GetSectorEquipmentFunctionRefOk returns a tuple with the SectorEquipmentFunctionRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) GetSectorEquipmentFunctionRefOk() (*string, bool) {
	if o == nil || IsNil(o.SectorEquipmentFunctionRef) {
		return nil, false
	}
	return o.SectorEquipmentFunctionRef, true
}

// HasSectorEquipmentFunctionRef returns a boolean if a field has been set.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) HasSectorEquipmentFunctionRef() bool {
	if o != nil && !IsNil(o.SectorEquipmentFunctionRef) {
		return true
	}

	return false
}

// SetSectorEquipmentFunctionRef gets a reference to the given string and assigns it to the SectorEquipmentFunctionRef field.
func (o *NrSectorCarrierSingleAllOfAttributesAllOf) SetSectorEquipmentFunctionRef(v string) {
	o.SectorEquipmentFunctionRef = &v
}

func (o NrSectorCarrierSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrSectorCarrierSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TxDirection) {
		toSerialize["txDirection"] = o.TxDirection
	}
	if !IsNil(o.ConfiguredMaxTxPower) {
		toSerialize["configuredMaxTxPower"] = o.ConfiguredMaxTxPower
	}
	if !IsNil(o.ArfcnDL) {
		toSerialize["arfcnDL"] = o.ArfcnDL
	}
	if !IsNil(o.ArfcnUL) {
		toSerialize["arfcnUL"] = o.ArfcnUL
	}
	if !IsNil(o.BSChannelBwDL) {
		toSerialize["bSChannelBwDL"] = o.BSChannelBwDL
	}
	if !IsNil(o.BSChannelBwUL) {
		toSerialize["bSChannelBwUL"] = o.BSChannelBwUL
	}
	if !IsNil(o.SectorEquipmentFunctionRef) {
		toSerialize["sectorEquipmentFunctionRef"] = o.SectorEquipmentFunctionRef
	}
	return toSerialize, nil
}

type NullableNrSectorCarrierSingleAllOfAttributesAllOf struct {
	value *NrSectorCarrierSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableNrSectorCarrierSingleAllOfAttributesAllOf) Get() *NrSectorCarrierSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableNrSectorCarrierSingleAllOfAttributesAllOf) Set(val *NrSectorCarrierSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNrSectorCarrierSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNrSectorCarrierSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrSectorCarrierSingleAllOfAttributesAllOf(val *NrSectorCarrierSingleAllOfAttributesAllOf) *NullableNrSectorCarrierSingleAllOfAttributesAllOf {
	return &NullableNrSectorCarrierSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableNrSectorCarrierSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrSectorCarrierSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
