/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the MbsServiceAreaInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsServiceAreaInfo{}

// MbsServiceAreaInfo MBS Service Area Information for location dependent MBS session
type MbsServiceAreaInfo struct {
	// Integer where the allowed values correspond to the value range of an unsigned 16-bit integer.
	AreaSessionId  int32          `json:"areaSessionId"`
	MbsServiceArea MbsServiceArea `json:"mbsServiceArea"`
}

// NewMbsServiceAreaInfo instantiates a new MbsServiceAreaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsServiceAreaInfo(areaSessionId int32, mbsServiceArea MbsServiceArea) *MbsServiceAreaInfo {
	this := MbsServiceAreaInfo{}
	this.AreaSessionId = areaSessionId
	this.MbsServiceArea = mbsServiceArea
	return &this
}

// NewMbsServiceAreaInfoWithDefaults instantiates a new MbsServiceAreaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsServiceAreaInfoWithDefaults() *MbsServiceAreaInfo {
	this := MbsServiceAreaInfo{}
	return &this
}

// GetAreaSessionId returns the AreaSessionId field value
func (o *MbsServiceAreaInfo) GetAreaSessionId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AreaSessionId
}

// GetAreaSessionIdOk returns a tuple with the AreaSessionId field value
// and a boolean to check if the value has been set.
func (o *MbsServiceAreaInfo) GetAreaSessionIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AreaSessionId, true
}

// SetAreaSessionId sets field value
func (o *MbsServiceAreaInfo) SetAreaSessionId(v int32) {
	o.AreaSessionId = v
}

// GetMbsServiceArea returns the MbsServiceArea field value
func (o *MbsServiceAreaInfo) GetMbsServiceArea() MbsServiceArea {
	if o == nil {
		var ret MbsServiceArea
		return ret
	}

	return o.MbsServiceArea
}

// GetMbsServiceAreaOk returns a tuple with the MbsServiceArea field value
// and a boolean to check if the value has been set.
func (o *MbsServiceAreaInfo) GetMbsServiceAreaOk() (*MbsServiceArea, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsServiceArea, true
}

// SetMbsServiceArea sets field value
func (o *MbsServiceAreaInfo) SetMbsServiceArea(v MbsServiceArea) {
	o.MbsServiceArea = v
}

func (o MbsServiceAreaInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsServiceAreaInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["areaSessionId"] = o.AreaSessionId
	toSerialize["mbsServiceArea"] = o.MbsServiceArea
	return toSerialize, nil
}

type NullableMbsServiceAreaInfo struct {
	value *MbsServiceAreaInfo
	isSet bool
}

func (v NullableMbsServiceAreaInfo) Get() *MbsServiceAreaInfo {
	return v.value
}

func (v *NullableMbsServiceAreaInfo) Set(val *MbsServiceAreaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsServiceAreaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsServiceAreaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsServiceAreaInfo(val *MbsServiceAreaInfo) *NullableMbsServiceAreaInfo {
	return &NullableMbsServiceAreaInfo{value: val, isSet: true}
}

func (v NullableMbsServiceAreaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsServiceAreaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
