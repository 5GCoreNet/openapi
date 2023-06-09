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

// checks if the SeppFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeppFunctionSingleAllOfAttributes{}

// SeppFunctionSingleAllOfAttributes struct for SeppFunctionSingleAllOfAttributes
type SeppFunctionSingleAllOfAttributes struct {
	ManagedFunctionAttr
	PlmnId   *PlmnId   `json:"plmnId,omitempty"`
	SEPPType *SEPPType `json:"sEPPType,omitempty"`
	SEPPId   *int32    `json:"sEPPId,omitempty"`
	Fqdn     *string   `json:"fqdn,omitempty"`
	SeppInfo *SeppInfo `json:"seppInfo,omitempty"`
}

// NewSeppFunctionSingleAllOfAttributes instantiates a new SeppFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeppFunctionSingleAllOfAttributes() *SeppFunctionSingleAllOfAttributes {
	this := SeppFunctionSingleAllOfAttributes{}
	return &this
}

// NewSeppFunctionSingleAllOfAttributesWithDefaults instantiates a new SeppFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeppFunctionSingleAllOfAttributesWithDefaults() *SeppFunctionSingleAllOfAttributes {
	this := SeppFunctionSingleAllOfAttributes{}
	return &this
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *SeppFunctionSingleAllOfAttributes) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppFunctionSingleAllOfAttributes) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *SeppFunctionSingleAllOfAttributes) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *SeppFunctionSingleAllOfAttributes) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetSEPPType returns the SEPPType field value if set, zero value otherwise.
func (o *SeppFunctionSingleAllOfAttributes) GetSEPPType() SEPPType {
	if o == nil || IsNil(o.SEPPType) {
		var ret SEPPType
		return ret
	}
	return *o.SEPPType
}

// GetSEPPTypeOk returns a tuple with the SEPPType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppFunctionSingleAllOfAttributes) GetSEPPTypeOk() (*SEPPType, bool) {
	if o == nil || IsNil(o.SEPPType) {
		return nil, false
	}
	return o.SEPPType, true
}

// HasSEPPType returns a boolean if a field has been set.
func (o *SeppFunctionSingleAllOfAttributes) HasSEPPType() bool {
	if o != nil && !IsNil(o.SEPPType) {
		return true
	}

	return false
}

// SetSEPPType gets a reference to the given SEPPType and assigns it to the SEPPType field.
func (o *SeppFunctionSingleAllOfAttributes) SetSEPPType(v SEPPType) {
	o.SEPPType = &v
}

// GetSEPPId returns the SEPPId field value if set, zero value otherwise.
func (o *SeppFunctionSingleAllOfAttributes) GetSEPPId() int32 {
	if o == nil || IsNil(o.SEPPId) {
		var ret int32
		return ret
	}
	return *o.SEPPId
}

// GetSEPPIdOk returns a tuple with the SEPPId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppFunctionSingleAllOfAttributes) GetSEPPIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SEPPId) {
		return nil, false
	}
	return o.SEPPId, true
}

// HasSEPPId returns a boolean if a field has been set.
func (o *SeppFunctionSingleAllOfAttributes) HasSEPPId() bool {
	if o != nil && !IsNil(o.SEPPId) {
		return true
	}

	return false
}

// SetSEPPId gets a reference to the given int32 and assigns it to the SEPPId field.
func (o *SeppFunctionSingleAllOfAttributes) SetSEPPId(v int32) {
	o.SEPPId = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *SeppFunctionSingleAllOfAttributes) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppFunctionSingleAllOfAttributes) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// HasFqdn returns a boolean if a field has been set.
func (o *SeppFunctionSingleAllOfAttributes) HasFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *SeppFunctionSingleAllOfAttributes) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetSeppInfo returns the SeppInfo field value if set, zero value otherwise.
func (o *SeppFunctionSingleAllOfAttributes) GetSeppInfo() SeppInfo {
	if o == nil || IsNil(o.SeppInfo) {
		var ret SeppInfo
		return ret
	}
	return *o.SeppInfo
}

// GetSeppInfoOk returns a tuple with the SeppInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppFunctionSingleAllOfAttributes) GetSeppInfoOk() (*SeppInfo, bool) {
	if o == nil || IsNil(o.SeppInfo) {
		return nil, false
	}
	return o.SeppInfo, true
}

// HasSeppInfo returns a boolean if a field has been set.
func (o *SeppFunctionSingleAllOfAttributes) HasSeppInfo() bool {
	if o != nil && !IsNil(o.SeppInfo) {
		return true
	}

	return false
}

// SetSeppInfo gets a reference to the given SeppInfo and assigns it to the SeppInfo field.
func (o *SeppFunctionSingleAllOfAttributes) SetSeppInfo(v SeppInfo) {
	o.SeppInfo = &v
}

func (o SeppFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeppFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedFunctionAttr, errManagedFunctionAttr := json.Marshal(o.ManagedFunctionAttr)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	errManagedFunctionAttr = json.Unmarshal([]byte(serializedManagedFunctionAttr), &toSerialize)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.SEPPType) {
		toSerialize["sEPPType"] = o.SEPPType
	}
	if !IsNil(o.SEPPId) {
		toSerialize["sEPPId"] = o.SEPPId
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.SeppInfo) {
		toSerialize["seppInfo"] = o.SeppInfo
	}
	return toSerialize, nil
}

type NullableSeppFunctionSingleAllOfAttributes struct {
	value *SeppFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableSeppFunctionSingleAllOfAttributes) Get() *SeppFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableSeppFunctionSingleAllOfAttributes) Set(val *SeppFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSeppFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSeppFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeppFunctionSingleAllOfAttributes(val *SeppFunctionSingleAllOfAttributes) *NullableSeppFunctionSingleAllOfAttributes {
	return &NullableSeppFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableSeppFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeppFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
