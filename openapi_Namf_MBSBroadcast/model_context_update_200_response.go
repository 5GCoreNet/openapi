/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
	"os"
)

// checks if the ContextUpdate200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextUpdate200Response{}

// ContextUpdate200Response struct for ContextUpdate200Response
type ContextUpdate200Response struct {
	JsonData                  *ContextUpdateRspData `json:"jsonData,omitempty"`
	BinaryDataN2Information1  **os.File             `json:"binaryDataN2Information1,omitempty"`
	BinaryDataN2Information2  **os.File             `json:"binaryDataN2Information2,omitempty"`
	BinaryDataN2Information3  **os.File             `json:"binaryDataN2Information3,omitempty"`
	BinaryDataN2Information4  **os.File             `json:"binaryDataN2Information4,omitempty"`
	BinaryDataN2Information5  **os.File             `json:"binaryDataN2Information5,omitempty"`
	BinaryDataN2Information6  **os.File             `json:"binaryDataN2Information6,omitempty"`
	BinaryDataN2Information7  **os.File             `json:"binaryDataN2Information7,omitempty"`
	BinaryDataN2Information8  **os.File             `json:"binaryDataN2Information8,omitempty"`
	BinaryDataN2Information9  **os.File             `json:"binaryDataN2Information9,omitempty"`
	BinaryDataN2Information10 **os.File             `json:"binaryDataN2Information10,omitempty"`
}

// NewContextUpdate200Response instantiates a new ContextUpdate200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextUpdate200Response() *ContextUpdate200Response {
	this := ContextUpdate200Response{}
	return &this
}

// NewContextUpdate200ResponseWithDefaults instantiates a new ContextUpdate200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextUpdate200ResponseWithDefaults() *ContextUpdate200Response {
	this := ContextUpdate200Response{}
	return &this
}

// GetJsonData returns the JsonData field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetJsonData() ContextUpdateRspData {
	if o == nil || IsNil(o.JsonData) {
		var ret ContextUpdateRspData
		return ret
	}
	return *o.JsonData
}

// GetJsonDataOk returns a tuple with the JsonData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetJsonDataOk() (*ContextUpdateRspData, bool) {
	if o == nil || IsNil(o.JsonData) {
		return nil, false
	}
	return o.JsonData, true
}

// HasJsonData returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasJsonData() bool {
	if o != nil && !IsNil(o.JsonData) {
		return true
	}

	return false
}

// SetJsonData gets a reference to the given ContextUpdateRspData and assigns it to the JsonData field.
func (o *ContextUpdate200Response) SetJsonData(v ContextUpdateRspData) {
	o.JsonData = &v
}

// GetBinaryDataN2Information1 returns the BinaryDataN2Information1 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information1() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information1) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information1
}

// GetBinaryDataN2Information1Ok returns a tuple with the BinaryDataN2Information1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information1Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information1) {
		return nil, false
	}
	return o.BinaryDataN2Information1, true
}

// HasBinaryDataN2Information1 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information1() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information1) {
		return true
	}

	return false
}

// SetBinaryDataN2Information1 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information1 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information1(v *os.File) {
	o.BinaryDataN2Information1 = &v
}

// GetBinaryDataN2Information2 returns the BinaryDataN2Information2 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information2() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information2) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information2
}

// GetBinaryDataN2Information2Ok returns a tuple with the BinaryDataN2Information2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information2Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information2) {
		return nil, false
	}
	return o.BinaryDataN2Information2, true
}

// HasBinaryDataN2Information2 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information2() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information2) {
		return true
	}

	return false
}

// SetBinaryDataN2Information2 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information2 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information2(v *os.File) {
	o.BinaryDataN2Information2 = &v
}

// GetBinaryDataN2Information3 returns the BinaryDataN2Information3 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information3() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information3) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information3
}

// GetBinaryDataN2Information3Ok returns a tuple with the BinaryDataN2Information3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information3Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information3) {
		return nil, false
	}
	return o.BinaryDataN2Information3, true
}

// HasBinaryDataN2Information3 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information3() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information3) {
		return true
	}

	return false
}

// SetBinaryDataN2Information3 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information3 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information3(v *os.File) {
	o.BinaryDataN2Information3 = &v
}

// GetBinaryDataN2Information4 returns the BinaryDataN2Information4 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information4() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information4) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information4
}

// GetBinaryDataN2Information4Ok returns a tuple with the BinaryDataN2Information4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information4Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information4) {
		return nil, false
	}
	return o.BinaryDataN2Information4, true
}

// HasBinaryDataN2Information4 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information4() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information4) {
		return true
	}

	return false
}

// SetBinaryDataN2Information4 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information4 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information4(v *os.File) {
	o.BinaryDataN2Information4 = &v
}

// GetBinaryDataN2Information5 returns the BinaryDataN2Information5 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information5() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information5) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information5
}

// GetBinaryDataN2Information5Ok returns a tuple with the BinaryDataN2Information5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information5Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information5) {
		return nil, false
	}
	return o.BinaryDataN2Information5, true
}

// HasBinaryDataN2Information5 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information5() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information5) {
		return true
	}

	return false
}

// SetBinaryDataN2Information5 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information5 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information5(v *os.File) {
	o.BinaryDataN2Information5 = &v
}

// GetBinaryDataN2Information6 returns the BinaryDataN2Information6 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information6() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information6) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information6
}

// GetBinaryDataN2Information6Ok returns a tuple with the BinaryDataN2Information6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information6Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information6) {
		return nil, false
	}
	return o.BinaryDataN2Information6, true
}

// HasBinaryDataN2Information6 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information6() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information6) {
		return true
	}

	return false
}

// SetBinaryDataN2Information6 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information6 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information6(v *os.File) {
	o.BinaryDataN2Information6 = &v
}

// GetBinaryDataN2Information7 returns the BinaryDataN2Information7 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information7() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information7) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information7
}

// GetBinaryDataN2Information7Ok returns a tuple with the BinaryDataN2Information7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information7Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information7) {
		return nil, false
	}
	return o.BinaryDataN2Information7, true
}

// HasBinaryDataN2Information7 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information7() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information7) {
		return true
	}

	return false
}

// SetBinaryDataN2Information7 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information7 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information7(v *os.File) {
	o.BinaryDataN2Information7 = &v
}

// GetBinaryDataN2Information8 returns the BinaryDataN2Information8 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information8() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information8) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information8
}

// GetBinaryDataN2Information8Ok returns a tuple with the BinaryDataN2Information8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information8Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information8) {
		return nil, false
	}
	return o.BinaryDataN2Information8, true
}

// HasBinaryDataN2Information8 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information8() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information8) {
		return true
	}

	return false
}

// SetBinaryDataN2Information8 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information8 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information8(v *os.File) {
	o.BinaryDataN2Information8 = &v
}

// GetBinaryDataN2Information9 returns the BinaryDataN2Information9 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information9() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information9) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information9
}

// GetBinaryDataN2Information9Ok returns a tuple with the BinaryDataN2Information9 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information9Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information9) {
		return nil, false
	}
	return o.BinaryDataN2Information9, true
}

// HasBinaryDataN2Information9 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information9() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information9) {
		return true
	}

	return false
}

// SetBinaryDataN2Information9 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information9 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information9(v *os.File) {
	o.BinaryDataN2Information9 = &v
}

// GetBinaryDataN2Information10 returns the BinaryDataN2Information10 field value if set, zero value otherwise.
func (o *ContextUpdate200Response) GetBinaryDataN2Information10() *os.File {
	if o == nil || IsNil(o.BinaryDataN2Information10) {
		var ret *os.File
		return ret
	}
	return *o.BinaryDataN2Information10
}

// GetBinaryDataN2Information10Ok returns a tuple with the BinaryDataN2Information10 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextUpdate200Response) GetBinaryDataN2Information10Ok() (**os.File, bool) {
	if o == nil || IsNil(o.BinaryDataN2Information10) {
		return nil, false
	}
	return o.BinaryDataN2Information10, true
}

// HasBinaryDataN2Information10 returns a boolean if a field has been set.
func (o *ContextUpdate200Response) HasBinaryDataN2Information10() bool {
	if o != nil && !IsNil(o.BinaryDataN2Information10) {
		return true
	}

	return false
}

// SetBinaryDataN2Information10 gets a reference to the given *os.File and assigns it to the BinaryDataN2Information10 field.
func (o *ContextUpdate200Response) SetBinaryDataN2Information10(v *os.File) {
	o.BinaryDataN2Information10 = &v
}

func (o ContextUpdate200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextUpdate200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.JsonData) {
		toSerialize["jsonData"] = o.JsonData
	}
	if !IsNil(o.BinaryDataN2Information1) {
		toSerialize["binaryDataN2Information1"] = o.BinaryDataN2Information1
	}
	if !IsNil(o.BinaryDataN2Information2) {
		toSerialize["binaryDataN2Information2"] = o.BinaryDataN2Information2
	}
	if !IsNil(o.BinaryDataN2Information3) {
		toSerialize["binaryDataN2Information3"] = o.BinaryDataN2Information3
	}
	if !IsNil(o.BinaryDataN2Information4) {
		toSerialize["binaryDataN2Information4"] = o.BinaryDataN2Information4
	}
	if !IsNil(o.BinaryDataN2Information5) {
		toSerialize["binaryDataN2Information5"] = o.BinaryDataN2Information5
	}
	if !IsNil(o.BinaryDataN2Information6) {
		toSerialize["binaryDataN2Information6"] = o.BinaryDataN2Information6
	}
	if !IsNil(o.BinaryDataN2Information7) {
		toSerialize["binaryDataN2Information7"] = o.BinaryDataN2Information7
	}
	if !IsNil(o.BinaryDataN2Information8) {
		toSerialize["binaryDataN2Information8"] = o.BinaryDataN2Information8
	}
	if !IsNil(o.BinaryDataN2Information9) {
		toSerialize["binaryDataN2Information9"] = o.BinaryDataN2Information9
	}
	if !IsNil(o.BinaryDataN2Information10) {
		toSerialize["binaryDataN2Information10"] = o.BinaryDataN2Information10
	}
	return toSerialize, nil
}

type NullableContextUpdate200Response struct {
	value *ContextUpdate200Response
	isSet bool
}

func (v NullableContextUpdate200Response) Get() *ContextUpdate200Response {
	return v.value
}

func (v *NullableContextUpdate200Response) Set(val *ContextUpdate200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableContextUpdate200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableContextUpdate200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextUpdate200Response(val *ContextUpdate200Response) *NullableContextUpdate200Response {
	return &NullableContextUpdate200Response{value: val, isSet: true}
}

func (v NullableContextUpdate200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextUpdate200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
