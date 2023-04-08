/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the AssignEbiData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignEbiData{}

// AssignEbiData Data within an EBI assignment request
type AssignEbiData struct {
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.  
	PduSessionId int32 `json:"pduSessionId"`
	ArpList []Arp `json:"arpList,omitempty"`
	ReleasedEbiList []int32 `json:"releasedEbiList,omitempty"`
	OldGuami *Guami `json:"oldGuami,omitempty"`
	ModifiedEbiList []EbiArpMapping `json:"modifiedEbiList,omitempty"`
}

// NewAssignEbiData instantiates a new AssignEbiData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignEbiData(pduSessionId int32) *AssignEbiData {
	this := AssignEbiData{}
	this.PduSessionId = pduSessionId
	return &this
}

// NewAssignEbiDataWithDefaults instantiates a new AssignEbiData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignEbiDataWithDefaults() *AssignEbiData {
	this := AssignEbiData{}
	return &this
}

// GetPduSessionId returns the PduSessionId field value
func (o *AssignEbiData) GetPduSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PduSessionId
}

// GetPduSessionIdOk returns a tuple with the PduSessionId field value
// and a boolean to check if the value has been set.
func (o *AssignEbiData) GetPduSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PduSessionId, true
}

// SetPduSessionId sets field value
func (o *AssignEbiData) SetPduSessionId(v int32) {
	o.PduSessionId = v
}

// GetArpList returns the ArpList field value if set, zero value otherwise.
func (o *AssignEbiData) GetArpList() []Arp {
	if o == nil || isNil(o.ArpList) {
		var ret []Arp
		return ret
	}
	return o.ArpList
}

// GetArpListOk returns a tuple with the ArpList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignEbiData) GetArpListOk() ([]Arp, bool) {
	if o == nil || isNil(o.ArpList) {
		return nil, false
	}
	return o.ArpList, true
}

// HasArpList returns a boolean if a field has been set.
func (o *AssignEbiData) HasArpList() bool {
	if o != nil && !isNil(o.ArpList) {
		return true
	}

	return false
}

// SetArpList gets a reference to the given []Arp and assigns it to the ArpList field.
func (o *AssignEbiData) SetArpList(v []Arp) {
	o.ArpList = v
}

// GetReleasedEbiList returns the ReleasedEbiList field value if set, zero value otherwise.
func (o *AssignEbiData) GetReleasedEbiList() []int32 {
	if o == nil || isNil(o.ReleasedEbiList) {
		var ret []int32
		return ret
	}
	return o.ReleasedEbiList
}

// GetReleasedEbiListOk returns a tuple with the ReleasedEbiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignEbiData) GetReleasedEbiListOk() ([]int32, bool) {
	if o == nil || isNil(o.ReleasedEbiList) {
		return nil, false
	}
	return o.ReleasedEbiList, true
}

// HasReleasedEbiList returns a boolean if a field has been set.
func (o *AssignEbiData) HasReleasedEbiList() bool {
	if o != nil && !isNil(o.ReleasedEbiList) {
		return true
	}

	return false
}

// SetReleasedEbiList gets a reference to the given []int32 and assigns it to the ReleasedEbiList field.
func (o *AssignEbiData) SetReleasedEbiList(v []int32) {
	o.ReleasedEbiList = v
}

// GetOldGuami returns the OldGuami field value if set, zero value otherwise.
func (o *AssignEbiData) GetOldGuami() Guami {
	if o == nil || isNil(o.OldGuami) {
		var ret Guami
		return ret
	}
	return *o.OldGuami
}

// GetOldGuamiOk returns a tuple with the OldGuami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignEbiData) GetOldGuamiOk() (*Guami, bool) {
	if o == nil || isNil(o.OldGuami) {
		return nil, false
	}
	return o.OldGuami, true
}

// HasOldGuami returns a boolean if a field has been set.
func (o *AssignEbiData) HasOldGuami() bool {
	if o != nil && !isNil(o.OldGuami) {
		return true
	}

	return false
}

// SetOldGuami gets a reference to the given Guami and assigns it to the OldGuami field.
func (o *AssignEbiData) SetOldGuami(v Guami) {
	o.OldGuami = &v
}

// GetModifiedEbiList returns the ModifiedEbiList field value if set, zero value otherwise.
func (o *AssignEbiData) GetModifiedEbiList() []EbiArpMapping {
	if o == nil || isNil(o.ModifiedEbiList) {
		var ret []EbiArpMapping
		return ret
	}
	return o.ModifiedEbiList
}

// GetModifiedEbiListOk returns a tuple with the ModifiedEbiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignEbiData) GetModifiedEbiListOk() ([]EbiArpMapping, bool) {
	if o == nil || isNil(o.ModifiedEbiList) {
		return nil, false
	}
	return o.ModifiedEbiList, true
}

// HasModifiedEbiList returns a boolean if a field has been set.
func (o *AssignEbiData) HasModifiedEbiList() bool {
	if o != nil && !isNil(o.ModifiedEbiList) {
		return true
	}

	return false
}

// SetModifiedEbiList gets a reference to the given []EbiArpMapping and assigns it to the ModifiedEbiList field.
func (o *AssignEbiData) SetModifiedEbiList(v []EbiArpMapping) {
	o.ModifiedEbiList = v
}

func (o AssignEbiData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignEbiData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pduSessionId"] = o.PduSessionId
	if !isNil(o.ArpList) {
		toSerialize["arpList"] = o.ArpList
	}
	if !isNil(o.ReleasedEbiList) {
		toSerialize["releasedEbiList"] = o.ReleasedEbiList
	}
	if !isNil(o.OldGuami) {
		toSerialize["oldGuami"] = o.OldGuami
	}
	if !isNil(o.ModifiedEbiList) {
		toSerialize["modifiedEbiList"] = o.ModifiedEbiList
	}
	return toSerialize, nil
}

type NullableAssignEbiData struct {
	value *AssignEbiData
	isSet bool
}

func (v NullableAssignEbiData) Get() *AssignEbiData {
	return v.value
}

func (v *NullableAssignEbiData) Set(val *AssignEbiData) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignEbiData) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignEbiData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignEbiData(val *AssignEbiData) *NullableAssignEbiData {
	return &NullableAssignEbiData{value: val, isSet: true}
}

func (v NullableAssignEbiData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignEbiData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


