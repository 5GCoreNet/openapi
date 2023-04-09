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

// checks if the CellIndividualOffset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CellIndividualOffset{}

// CellIndividualOffset struct for CellIndividualOffset
type CellIndividualOffset struct {
	RsrpOffsetSSB   *int32 `json:"rsrpOffsetSSB,omitempty"`
	RsrqOffsetSSB   *int32 `json:"rsrqOffsetSSB,omitempty"`
	SinrOffsetSSB   *int32 `json:"sinrOffsetSSB,omitempty"`
	RsrpOffsetCSIRS *int32 `json:"rsrpOffsetCSI-RS,omitempty"`
	RsrqOffsetCSIRS *int32 `json:"rsrqOffsetCSI-RS,omitempty"`
	SinrOffsetCSIRS *int32 `json:"sinrOffsetCSI-RS,omitempty"`
}

// NewCellIndividualOffset instantiates a new CellIndividualOffset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCellIndividualOffset() *CellIndividualOffset {
	this := CellIndividualOffset{}
	return &this
}

// NewCellIndividualOffsetWithDefaults instantiates a new CellIndividualOffset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCellIndividualOffsetWithDefaults() *CellIndividualOffset {
	this := CellIndividualOffset{}
	return &this
}

// GetRsrpOffsetSSB returns the RsrpOffsetSSB field value if set, zero value otherwise.
func (o *CellIndividualOffset) GetRsrpOffsetSSB() int32 {
	if o == nil || IsNil(o.RsrpOffsetSSB) {
		var ret int32
		return ret
	}
	return *o.RsrpOffsetSSB
}

// GetRsrpOffsetSSBOk returns a tuple with the RsrpOffsetSSB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellIndividualOffset) GetRsrpOffsetSSBOk() (*int32, bool) {
	if o == nil || IsNil(o.RsrpOffsetSSB) {
		return nil, false
	}
	return o.RsrpOffsetSSB, true
}

// HasRsrpOffsetSSB returns a boolean if a field has been set.
func (o *CellIndividualOffset) HasRsrpOffsetSSB() bool {
	if o != nil && !IsNil(o.RsrpOffsetSSB) {
		return true
	}

	return false
}

// SetRsrpOffsetSSB gets a reference to the given int32 and assigns it to the RsrpOffsetSSB field.
func (o *CellIndividualOffset) SetRsrpOffsetSSB(v int32) {
	o.RsrpOffsetSSB = &v
}

// GetRsrqOffsetSSB returns the RsrqOffsetSSB field value if set, zero value otherwise.
func (o *CellIndividualOffset) GetRsrqOffsetSSB() int32 {
	if o == nil || IsNil(o.RsrqOffsetSSB) {
		var ret int32
		return ret
	}
	return *o.RsrqOffsetSSB
}

// GetRsrqOffsetSSBOk returns a tuple with the RsrqOffsetSSB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellIndividualOffset) GetRsrqOffsetSSBOk() (*int32, bool) {
	if o == nil || IsNil(o.RsrqOffsetSSB) {
		return nil, false
	}
	return o.RsrqOffsetSSB, true
}

// HasRsrqOffsetSSB returns a boolean if a field has been set.
func (o *CellIndividualOffset) HasRsrqOffsetSSB() bool {
	if o != nil && !IsNil(o.RsrqOffsetSSB) {
		return true
	}

	return false
}

// SetRsrqOffsetSSB gets a reference to the given int32 and assigns it to the RsrqOffsetSSB field.
func (o *CellIndividualOffset) SetRsrqOffsetSSB(v int32) {
	o.RsrqOffsetSSB = &v
}

// GetSinrOffsetSSB returns the SinrOffsetSSB field value if set, zero value otherwise.
func (o *CellIndividualOffset) GetSinrOffsetSSB() int32 {
	if o == nil || IsNil(o.SinrOffsetSSB) {
		var ret int32
		return ret
	}
	return *o.SinrOffsetSSB
}

// GetSinrOffsetSSBOk returns a tuple with the SinrOffsetSSB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellIndividualOffset) GetSinrOffsetSSBOk() (*int32, bool) {
	if o == nil || IsNil(o.SinrOffsetSSB) {
		return nil, false
	}
	return o.SinrOffsetSSB, true
}

// HasSinrOffsetSSB returns a boolean if a field has been set.
func (o *CellIndividualOffset) HasSinrOffsetSSB() bool {
	if o != nil && !IsNil(o.SinrOffsetSSB) {
		return true
	}

	return false
}

// SetSinrOffsetSSB gets a reference to the given int32 and assigns it to the SinrOffsetSSB field.
func (o *CellIndividualOffset) SetSinrOffsetSSB(v int32) {
	o.SinrOffsetSSB = &v
}

// GetRsrpOffsetCSIRS returns the RsrpOffsetCSIRS field value if set, zero value otherwise.
func (o *CellIndividualOffset) GetRsrpOffsetCSIRS() int32 {
	if o == nil || IsNil(o.RsrpOffsetCSIRS) {
		var ret int32
		return ret
	}
	return *o.RsrpOffsetCSIRS
}

// GetRsrpOffsetCSIRSOk returns a tuple with the RsrpOffsetCSIRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellIndividualOffset) GetRsrpOffsetCSIRSOk() (*int32, bool) {
	if o == nil || IsNil(o.RsrpOffsetCSIRS) {
		return nil, false
	}
	return o.RsrpOffsetCSIRS, true
}

// HasRsrpOffsetCSIRS returns a boolean if a field has been set.
func (o *CellIndividualOffset) HasRsrpOffsetCSIRS() bool {
	if o != nil && !IsNil(o.RsrpOffsetCSIRS) {
		return true
	}

	return false
}

// SetRsrpOffsetCSIRS gets a reference to the given int32 and assigns it to the RsrpOffsetCSIRS field.
func (o *CellIndividualOffset) SetRsrpOffsetCSIRS(v int32) {
	o.RsrpOffsetCSIRS = &v
}

// GetRsrqOffsetCSIRS returns the RsrqOffsetCSIRS field value if set, zero value otherwise.
func (o *CellIndividualOffset) GetRsrqOffsetCSIRS() int32 {
	if o == nil || IsNil(o.RsrqOffsetCSIRS) {
		var ret int32
		return ret
	}
	return *o.RsrqOffsetCSIRS
}

// GetRsrqOffsetCSIRSOk returns a tuple with the RsrqOffsetCSIRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellIndividualOffset) GetRsrqOffsetCSIRSOk() (*int32, bool) {
	if o == nil || IsNil(o.RsrqOffsetCSIRS) {
		return nil, false
	}
	return o.RsrqOffsetCSIRS, true
}

// HasRsrqOffsetCSIRS returns a boolean if a field has been set.
func (o *CellIndividualOffset) HasRsrqOffsetCSIRS() bool {
	if o != nil && !IsNil(o.RsrqOffsetCSIRS) {
		return true
	}

	return false
}

// SetRsrqOffsetCSIRS gets a reference to the given int32 and assigns it to the RsrqOffsetCSIRS field.
func (o *CellIndividualOffset) SetRsrqOffsetCSIRS(v int32) {
	o.RsrqOffsetCSIRS = &v
}

// GetSinrOffsetCSIRS returns the SinrOffsetCSIRS field value if set, zero value otherwise.
func (o *CellIndividualOffset) GetSinrOffsetCSIRS() int32 {
	if o == nil || IsNil(o.SinrOffsetCSIRS) {
		var ret int32
		return ret
	}
	return *o.SinrOffsetCSIRS
}

// GetSinrOffsetCSIRSOk returns a tuple with the SinrOffsetCSIRS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CellIndividualOffset) GetSinrOffsetCSIRSOk() (*int32, bool) {
	if o == nil || IsNil(o.SinrOffsetCSIRS) {
		return nil, false
	}
	return o.SinrOffsetCSIRS, true
}

// HasSinrOffsetCSIRS returns a boolean if a field has been set.
func (o *CellIndividualOffset) HasSinrOffsetCSIRS() bool {
	if o != nil && !IsNil(o.SinrOffsetCSIRS) {
		return true
	}

	return false
}

// SetSinrOffsetCSIRS gets a reference to the given int32 and assigns it to the SinrOffsetCSIRS field.
func (o *CellIndividualOffset) SetSinrOffsetCSIRS(v int32) {
	o.SinrOffsetCSIRS = &v
}

func (o CellIndividualOffset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CellIndividualOffset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RsrpOffsetSSB) {
		toSerialize["rsrpOffsetSSB"] = o.RsrpOffsetSSB
	}
	if !IsNil(o.RsrqOffsetSSB) {
		toSerialize["rsrqOffsetSSB"] = o.RsrqOffsetSSB
	}
	if !IsNil(o.SinrOffsetSSB) {
		toSerialize["sinrOffsetSSB"] = o.SinrOffsetSSB
	}
	if !IsNil(o.RsrpOffsetCSIRS) {
		toSerialize["rsrpOffsetCSI-RS"] = o.RsrpOffsetCSIRS
	}
	if !IsNil(o.RsrqOffsetCSIRS) {
		toSerialize["rsrqOffsetCSI-RS"] = o.RsrqOffsetCSIRS
	}
	if !IsNil(o.SinrOffsetCSIRS) {
		toSerialize["sinrOffsetCSI-RS"] = o.SinrOffsetCSIRS
	}
	return toSerialize, nil
}

type NullableCellIndividualOffset struct {
	value *CellIndividualOffset
	isSet bool
}

func (v NullableCellIndividualOffset) Get() *CellIndividualOffset {
	return v.value
}

func (v *NullableCellIndividualOffset) Set(val *CellIndividualOffset) {
	v.value = val
	v.isSet = true
}

func (v NullableCellIndividualOffset) IsSet() bool {
	return v.isSet
}

func (v *NullableCellIndividualOffset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellIndividualOffset(val *CellIndividualOffset) *NullableCellIndividualOffset {
	return &NullableCellIndividualOffset{value: val, isSet: true}
}

func (v NullableCellIndividualOffset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellIndividualOffset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
