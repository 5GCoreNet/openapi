/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
)

// checks if the MbsMediaInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsMediaInfo{}

// MbsMediaInfo Represent MBS Media Information.
type MbsMediaInfo struct {
	MbsMedType *MediaType `json:"mbsMedType,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxReqMbsBwDl *string `json:"maxReqMbsBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MinReqMbsBwDl *string `json:"minReqMbsBwDl,omitempty"`
	Codecs []string `json:"codecs,omitempty"`
}

// NewMbsMediaInfo instantiates a new MbsMediaInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsMediaInfo() *MbsMediaInfo {
	this := MbsMediaInfo{}
	return &this
}

// NewMbsMediaInfoWithDefaults instantiates a new MbsMediaInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsMediaInfoWithDefaults() *MbsMediaInfo {
	this := MbsMediaInfo{}
	return &this
}

// GetMbsMedType returns the MbsMedType field value if set, zero value otherwise.
func (o *MbsMediaInfo) GetMbsMedType() MediaType {
	if o == nil || isNil(o.MbsMedType) {
		var ret MediaType
		return ret
	}
	return *o.MbsMedType
}

// GetMbsMedTypeOk returns a tuple with the MbsMedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaInfo) GetMbsMedTypeOk() (*MediaType, bool) {
	if o == nil || isNil(o.MbsMedType) {
		return nil, false
	}
	return o.MbsMedType, true
}

// HasMbsMedType returns a boolean if a field has been set.
func (o *MbsMediaInfo) HasMbsMedType() bool {
	if o != nil && !isNil(o.MbsMedType) {
		return true
	}

	return false
}

// SetMbsMedType gets a reference to the given MediaType and assigns it to the MbsMedType field.
func (o *MbsMediaInfo) SetMbsMedType(v MediaType) {
	o.MbsMedType = &v
}

// GetMaxReqMbsBwDl returns the MaxReqMbsBwDl field value if set, zero value otherwise.
func (o *MbsMediaInfo) GetMaxReqMbsBwDl() string {
	if o == nil || isNil(o.MaxReqMbsBwDl) {
		var ret string
		return ret
	}
	return *o.MaxReqMbsBwDl
}

// GetMaxReqMbsBwDlOk returns a tuple with the MaxReqMbsBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaInfo) GetMaxReqMbsBwDlOk() (*string, bool) {
	if o == nil || isNil(o.MaxReqMbsBwDl) {
		return nil, false
	}
	return o.MaxReqMbsBwDl, true
}

// HasMaxReqMbsBwDl returns a boolean if a field has been set.
func (o *MbsMediaInfo) HasMaxReqMbsBwDl() bool {
	if o != nil && !isNil(o.MaxReqMbsBwDl) {
		return true
	}

	return false
}

// SetMaxReqMbsBwDl gets a reference to the given string and assigns it to the MaxReqMbsBwDl field.
func (o *MbsMediaInfo) SetMaxReqMbsBwDl(v string) {
	o.MaxReqMbsBwDl = &v
}

// GetMinReqMbsBwDl returns the MinReqMbsBwDl field value if set, zero value otherwise.
func (o *MbsMediaInfo) GetMinReqMbsBwDl() string {
	if o == nil || isNil(o.MinReqMbsBwDl) {
		var ret string
		return ret
	}
	return *o.MinReqMbsBwDl
}

// GetMinReqMbsBwDlOk returns a tuple with the MinReqMbsBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaInfo) GetMinReqMbsBwDlOk() (*string, bool) {
	if o == nil || isNil(o.MinReqMbsBwDl) {
		return nil, false
	}
	return o.MinReqMbsBwDl, true
}

// HasMinReqMbsBwDl returns a boolean if a field has been set.
func (o *MbsMediaInfo) HasMinReqMbsBwDl() bool {
	if o != nil && !isNil(o.MinReqMbsBwDl) {
		return true
	}

	return false
}

// SetMinReqMbsBwDl gets a reference to the given string and assigns it to the MinReqMbsBwDl field.
func (o *MbsMediaInfo) SetMinReqMbsBwDl(v string) {
	o.MinReqMbsBwDl = &v
}

// GetCodecs returns the Codecs field value if set, zero value otherwise.
func (o *MbsMediaInfo) GetCodecs() []string {
	if o == nil || isNil(o.Codecs) {
		var ret []string
		return ret
	}
	return o.Codecs
}

// GetCodecsOk returns a tuple with the Codecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsMediaInfo) GetCodecsOk() ([]string, bool) {
	if o == nil || isNil(o.Codecs) {
		return nil, false
	}
	return o.Codecs, true
}

// HasCodecs returns a boolean if a field has been set.
func (o *MbsMediaInfo) HasCodecs() bool {
	if o != nil && !isNil(o.Codecs) {
		return true
	}

	return false
}

// SetCodecs gets a reference to the given []string and assigns it to the Codecs field.
func (o *MbsMediaInfo) SetCodecs(v []string) {
	o.Codecs = v
}

func (o MbsMediaInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsMediaInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MbsMedType) {
		toSerialize["mbsMedType"] = o.MbsMedType
	}
	if !isNil(o.MaxReqMbsBwDl) {
		toSerialize["maxReqMbsBwDl"] = o.MaxReqMbsBwDl
	}
	if !isNil(o.MinReqMbsBwDl) {
		toSerialize["minReqMbsBwDl"] = o.MinReqMbsBwDl
	}
	if !isNil(o.Codecs) {
		toSerialize["codecs"] = o.Codecs
	}
	return toSerialize, nil
}

type NullableMbsMediaInfo struct {
	value *MbsMediaInfo
	isSet bool
}

func (v NullableMbsMediaInfo) Get() *MbsMediaInfo {
	return v.value
}

func (v *NullableMbsMediaInfo) Set(val *MbsMediaInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsMediaInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsMediaInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsMediaInfo(val *MbsMediaInfo) *NullableMbsMediaInfo {
	return &NullableMbsMediaInfo{value: val, isSet: true}
}

func (v NullableMbsMediaInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsMediaInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


