/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the N4Information type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N4Information{}

// N4Information N4 Information
type N4Information struct {
	N4MessageType N4MessageType `json:"n4MessageType"`
	N4MessagePayload RefToBinaryData `json:"n4MessagePayload"`
	N4DnaiInfo *DnaiInformation `json:"n4DnaiInfo,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	PsaUpfId *string `json:"psaUpfId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	UlClBpId *string `json:"ulClBpId,omitempty"`
	N9UlPdrIdList []int32 `json:"n9UlPdrIdList,omitempty"`
}

// NewN4Information instantiates a new N4Information object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN4Information(n4MessageType N4MessageType, n4MessagePayload RefToBinaryData) *N4Information {
	this := N4Information{}
	this.N4MessageType = n4MessageType
	this.N4MessagePayload = n4MessagePayload
	return &this
}

// NewN4InformationWithDefaults instantiates a new N4Information object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN4InformationWithDefaults() *N4Information {
	this := N4Information{}
	return &this
}

// GetN4MessageType returns the N4MessageType field value
func (o *N4Information) GetN4MessageType() N4MessageType {
	if o == nil {
		var ret N4MessageType
		return ret
	}

	return o.N4MessageType
}

// GetN4MessageTypeOk returns a tuple with the N4MessageType field value
// and a boolean to check if the value has been set.
func (o *N4Information) GetN4MessageTypeOk() (*N4MessageType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N4MessageType, true
}

// SetN4MessageType sets field value
func (o *N4Information) SetN4MessageType(v N4MessageType) {
	o.N4MessageType = v
}

// GetN4MessagePayload returns the N4MessagePayload field value
func (o *N4Information) GetN4MessagePayload() RefToBinaryData {
	if o == nil {
		var ret RefToBinaryData
		return ret
	}

	return o.N4MessagePayload
}

// GetN4MessagePayloadOk returns a tuple with the N4MessagePayload field value
// and a boolean to check if the value has been set.
func (o *N4Information) GetN4MessagePayloadOk() (*RefToBinaryData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N4MessagePayload, true
}

// SetN4MessagePayload sets field value
func (o *N4Information) SetN4MessagePayload(v RefToBinaryData) {
	o.N4MessagePayload = v
}

// GetN4DnaiInfo returns the N4DnaiInfo field value if set, zero value otherwise.
func (o *N4Information) GetN4DnaiInfo() DnaiInformation {
	if o == nil || isNil(o.N4DnaiInfo) {
		var ret DnaiInformation
		return ret
	}
	return *o.N4DnaiInfo
}

// GetN4DnaiInfoOk returns a tuple with the N4DnaiInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N4Information) GetN4DnaiInfoOk() (*DnaiInformation, bool) {
	if o == nil || isNil(o.N4DnaiInfo) {
		return nil, false
	}
	return o.N4DnaiInfo, true
}

// HasN4DnaiInfo returns a boolean if a field has been set.
func (o *N4Information) HasN4DnaiInfo() bool {
	if o != nil && !isNil(o.N4DnaiInfo) {
		return true
	}

	return false
}

// SetN4DnaiInfo gets a reference to the given DnaiInformation and assigns it to the N4DnaiInfo field.
func (o *N4Information) SetN4DnaiInfo(v DnaiInformation) {
	o.N4DnaiInfo = &v
}

// GetPsaUpfId returns the PsaUpfId field value if set, zero value otherwise.
func (o *N4Information) GetPsaUpfId() string {
	if o == nil || isNil(o.PsaUpfId) {
		var ret string
		return ret
	}
	return *o.PsaUpfId
}

// GetPsaUpfIdOk returns a tuple with the PsaUpfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N4Information) GetPsaUpfIdOk() (*string, bool) {
	if o == nil || isNil(o.PsaUpfId) {
		return nil, false
	}
	return o.PsaUpfId, true
}

// HasPsaUpfId returns a boolean if a field has been set.
func (o *N4Information) HasPsaUpfId() bool {
	if o != nil && !isNil(o.PsaUpfId) {
		return true
	}

	return false
}

// SetPsaUpfId gets a reference to the given string and assigns it to the PsaUpfId field.
func (o *N4Information) SetPsaUpfId(v string) {
	o.PsaUpfId = &v
}

// GetUlClBpId returns the UlClBpId field value if set, zero value otherwise.
func (o *N4Information) GetUlClBpId() string {
	if o == nil || isNil(o.UlClBpId) {
		var ret string
		return ret
	}
	return *o.UlClBpId
}

// GetUlClBpIdOk returns a tuple with the UlClBpId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N4Information) GetUlClBpIdOk() (*string, bool) {
	if o == nil || isNil(o.UlClBpId) {
		return nil, false
	}
	return o.UlClBpId, true
}

// HasUlClBpId returns a boolean if a field has been set.
func (o *N4Information) HasUlClBpId() bool {
	if o != nil && !isNil(o.UlClBpId) {
		return true
	}

	return false
}

// SetUlClBpId gets a reference to the given string and assigns it to the UlClBpId field.
func (o *N4Information) SetUlClBpId(v string) {
	o.UlClBpId = &v
}

// GetN9UlPdrIdList returns the N9UlPdrIdList field value if set, zero value otherwise.
func (o *N4Information) GetN9UlPdrIdList() []int32 {
	if o == nil || isNil(o.N9UlPdrIdList) {
		var ret []int32
		return ret
	}
	return o.N9UlPdrIdList
}

// GetN9UlPdrIdListOk returns a tuple with the N9UlPdrIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N4Information) GetN9UlPdrIdListOk() ([]int32, bool) {
	if o == nil || isNil(o.N9UlPdrIdList) {
		return nil, false
	}
	return o.N9UlPdrIdList, true
}

// HasN9UlPdrIdList returns a boolean if a field has been set.
func (o *N4Information) HasN9UlPdrIdList() bool {
	if o != nil && !isNil(o.N9UlPdrIdList) {
		return true
	}

	return false
}

// SetN9UlPdrIdList gets a reference to the given []int32 and assigns it to the N9UlPdrIdList field.
func (o *N4Information) SetN9UlPdrIdList(v []int32) {
	o.N9UlPdrIdList = v
}

func (o N4Information) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N4Information) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["n4MessageType"] = o.N4MessageType
	toSerialize["n4MessagePayload"] = o.N4MessagePayload
	if !isNil(o.N4DnaiInfo) {
		toSerialize["n4DnaiInfo"] = o.N4DnaiInfo
	}
	if !isNil(o.PsaUpfId) {
		toSerialize["psaUpfId"] = o.PsaUpfId
	}
	if !isNil(o.UlClBpId) {
		toSerialize["ulClBpId"] = o.UlClBpId
	}
	if !isNil(o.N9UlPdrIdList) {
		toSerialize["n9UlPdrIdList"] = o.N9UlPdrIdList
	}
	return toSerialize, nil
}

type NullableN4Information struct {
	value *N4Information
	isSet bool
}

func (v NullableN4Information) Get() *N4Information {
	return v.value
}

func (v *NullableN4Information) Set(val *N4Information) {
	v.value = val
	v.isSet = true
}

func (v NullableN4Information) IsSet() bool {
	return v.isSet
}

func (v *NullableN4Information) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN4Information(val *N4Information) *NullableN4Information {
	return &NullableN4Information{value: val, isSet: true}
}

func (v NullableN4Information) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN4Information) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


