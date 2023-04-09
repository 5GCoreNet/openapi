/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_JOSEProtectedMessageForwarding

import (
	"encoding/json"
)

// checks if the FlatJweJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlatJweJson{}

// FlatJweJson Contains the integrity protected reformatted block
type FlatJweJson struct {
	Protected    *string                `json:"protected,omitempty"`
	Unprotected  map[string]interface{} `json:"unprotected,omitempty"`
	Header       map[string]interface{} `json:"header,omitempty"`
	EncryptedKey *string                `json:"encrypted_key,omitempty"`
	Aad          *string                `json:"aad,omitempty"`
	Iv           *string                `json:"iv,omitempty"`
	Ciphertext   string                 `json:"ciphertext"`
	Tag          *string                `json:"tag,omitempty"`
}

// NewFlatJweJson instantiates a new FlatJweJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlatJweJson(ciphertext string) *FlatJweJson {
	this := FlatJweJson{}
	this.Ciphertext = ciphertext
	return &this
}

// NewFlatJweJsonWithDefaults instantiates a new FlatJweJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlatJweJsonWithDefaults() *FlatJweJson {
	this := FlatJweJson{}
	return &this
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *FlatJweJson) GetProtected() string {
	if o == nil || IsNil(o.Protected) {
		var ret string
		return ret
	}
	return *o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJweJson) GetProtectedOk() (*string, bool) {
	if o == nil || IsNil(o.Protected) {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *FlatJweJson) HasProtected() bool {
	if o != nil && !IsNil(o.Protected) {
		return true
	}

	return false
}

// SetProtected gets a reference to the given string and assigns it to the Protected field.
func (o *FlatJweJson) SetProtected(v string) {
	o.Protected = &v
}

// GetUnprotected returns the Unprotected field value if set, zero value otherwise.
func (o *FlatJweJson) GetUnprotected() map[string]interface{} {
	if o == nil || IsNil(o.Unprotected) {
		var ret map[string]interface{}
		return ret
	}
	return o.Unprotected
}

// GetUnprotectedOk returns a tuple with the Unprotected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJweJson) GetUnprotectedOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Unprotected) {
		return map[string]interface{}{}, false
	}
	return o.Unprotected, true
}

// HasUnprotected returns a boolean if a field has been set.
func (o *FlatJweJson) HasUnprotected() bool {
	if o != nil && !IsNil(o.Unprotected) {
		return true
	}

	return false
}

// SetUnprotected gets a reference to the given map[string]interface{} and assigns it to the Unprotected field.
func (o *FlatJweJson) SetUnprotected(v map[string]interface{}) {
	o.Unprotected = v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *FlatJweJson) GetHeader() map[string]interface{} {
	if o == nil || IsNil(o.Header) {
		var ret map[string]interface{}
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJweJson) GetHeaderOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Header) {
		return map[string]interface{}{}, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *FlatJweJson) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given map[string]interface{} and assigns it to the Header field.
func (o *FlatJweJson) SetHeader(v map[string]interface{}) {
	o.Header = v
}

// GetEncryptedKey returns the EncryptedKey field value if set, zero value otherwise.
func (o *FlatJweJson) GetEncryptedKey() string {
	if o == nil || IsNil(o.EncryptedKey) {
		var ret string
		return ret
	}
	return *o.EncryptedKey
}

// GetEncryptedKeyOk returns a tuple with the EncryptedKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJweJson) GetEncryptedKeyOk() (*string, bool) {
	if o == nil || IsNil(o.EncryptedKey) {
		return nil, false
	}
	return o.EncryptedKey, true
}

// HasEncryptedKey returns a boolean if a field has been set.
func (o *FlatJweJson) HasEncryptedKey() bool {
	if o != nil && !IsNil(o.EncryptedKey) {
		return true
	}

	return false
}

// SetEncryptedKey gets a reference to the given string and assigns it to the EncryptedKey field.
func (o *FlatJweJson) SetEncryptedKey(v string) {
	o.EncryptedKey = &v
}

// GetAad returns the Aad field value if set, zero value otherwise.
func (o *FlatJweJson) GetAad() string {
	if o == nil || IsNil(o.Aad) {
		var ret string
		return ret
	}
	return *o.Aad
}

// GetAadOk returns a tuple with the Aad field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJweJson) GetAadOk() (*string, bool) {
	if o == nil || IsNil(o.Aad) {
		return nil, false
	}
	return o.Aad, true
}

// HasAad returns a boolean if a field has been set.
func (o *FlatJweJson) HasAad() bool {
	if o != nil && !IsNil(o.Aad) {
		return true
	}

	return false
}

// SetAad gets a reference to the given string and assigns it to the Aad field.
func (o *FlatJweJson) SetAad(v string) {
	o.Aad = &v
}

// GetIv returns the Iv field value if set, zero value otherwise.
func (o *FlatJweJson) GetIv() string {
	if o == nil || IsNil(o.Iv) {
		var ret string
		return ret
	}
	return *o.Iv
}

// GetIvOk returns a tuple with the Iv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJweJson) GetIvOk() (*string, bool) {
	if o == nil || IsNil(o.Iv) {
		return nil, false
	}
	return o.Iv, true
}

// HasIv returns a boolean if a field has been set.
func (o *FlatJweJson) HasIv() bool {
	if o != nil && !IsNil(o.Iv) {
		return true
	}

	return false
}

// SetIv gets a reference to the given string and assigns it to the Iv field.
func (o *FlatJweJson) SetIv(v string) {
	o.Iv = &v
}

// GetCiphertext returns the Ciphertext field value
func (o *FlatJweJson) GetCiphertext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ciphertext
}

// GetCiphertextOk returns a tuple with the Ciphertext field value
// and a boolean to check if the value has been set.
func (o *FlatJweJson) GetCiphertextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ciphertext, true
}

// SetCiphertext sets field value
func (o *FlatJweJson) SetCiphertext(v string) {
	o.Ciphertext = v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *FlatJweJson) GetTag() string {
	if o == nil || IsNil(o.Tag) {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJweJson) GetTagOk() (*string, bool) {
	if o == nil || IsNil(o.Tag) {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *FlatJweJson) HasTag() bool {
	if o != nil && !IsNil(o.Tag) {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *FlatJweJson) SetTag(v string) {
	o.Tag = &v
}

func (o FlatJweJson) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlatJweJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Protected) {
		toSerialize["protected"] = o.Protected
	}
	if !IsNil(o.Unprotected) {
		toSerialize["unprotected"] = o.Unprotected
	}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	if !IsNil(o.EncryptedKey) {
		toSerialize["encrypted_key"] = o.EncryptedKey
	}
	if !IsNil(o.Aad) {
		toSerialize["aad"] = o.Aad
	}
	if !IsNil(o.Iv) {
		toSerialize["iv"] = o.Iv
	}
	toSerialize["ciphertext"] = o.Ciphertext
	if !IsNil(o.Tag) {
		toSerialize["tag"] = o.Tag
	}
	return toSerialize, nil
}

type NullableFlatJweJson struct {
	value *FlatJweJson
	isSet bool
}

func (v NullableFlatJweJson) Get() *FlatJweJson {
	return v.value
}

func (v *NullableFlatJweJson) Set(val *FlatJweJson) {
	v.value = val
	v.isSet = true
}

func (v NullableFlatJweJson) IsSet() bool {
	return v.isSet
}

func (v *NullableFlatJweJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlatJweJson(val *FlatJweJson) *NullableFlatJweJson {
	return &NullableFlatJweJson{value: val, isSet: true}
}

func (v NullableFlatJweJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlatJweJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
