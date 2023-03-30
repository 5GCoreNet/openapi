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

// checks if the NrfFunctionSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrfFunctionSingleAllOfAttributesAllOf{}

// NrfFunctionSingleAllOfAttributesAllOf struct for NrfFunctionSingleAllOfAttributesAllOf
type NrfFunctionSingleAllOfAttributesAllOf struct {
	PlmnIdList []PlmnId `json:"plmnIdList,omitempty"`
	SBIFqdn *string `json:"sBIFqdn,omitempty"`
	CNSIIdList []string `json:"cNSIIdList,omitempty"`
	// List of NF profile
	NFProfileList []NFProfile `json:"nFProfileList,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
}

// NewNrfFunctionSingleAllOfAttributesAllOf instantiates a new NrfFunctionSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrfFunctionSingleAllOfAttributesAllOf() *NrfFunctionSingleAllOfAttributesAllOf {
	this := NrfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// NewNrfFunctionSingleAllOfAttributesAllOfWithDefaults instantiates a new NrfFunctionSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrfFunctionSingleAllOfAttributesAllOfWithDefaults() *NrfFunctionSingleAllOfAttributesAllOf {
	this := NrfFunctionSingleAllOfAttributesAllOf{}
	return &this
}

// GetPlmnIdList returns the PlmnIdList field value if set, zero value otherwise.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetPlmnIdList() []PlmnId {
	if o == nil || IsNil(o.PlmnIdList) {
		var ret []PlmnId
		return ret
	}
	return o.PlmnIdList
}

// GetPlmnIdListOk returns a tuple with the PlmnIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetPlmnIdListOk() ([]PlmnId, bool) {
	if o == nil || IsNil(o.PlmnIdList) {
		return nil, false
	}
	return o.PlmnIdList, true
}

// HasPlmnIdList returns a boolean if a field has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) HasPlmnIdList() bool {
	if o != nil && !IsNil(o.PlmnIdList) {
		return true
	}

	return false
}

// SetPlmnIdList gets a reference to the given []PlmnId and assigns it to the PlmnIdList field.
func (o *NrfFunctionSingleAllOfAttributesAllOf) SetPlmnIdList(v []PlmnId) {
	o.PlmnIdList = v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *NrfFunctionSingleAllOfAttributesAllOf) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetCNSIIdList returns the CNSIIdList field value if set, zero value otherwise.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetCNSIIdList() []string {
	if o == nil || IsNil(o.CNSIIdList) {
		var ret []string
		return ret
	}
	return o.CNSIIdList
}

// GetCNSIIdListOk returns a tuple with the CNSIIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetCNSIIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.CNSIIdList) {
		return nil, false
	}
	return o.CNSIIdList, true
}

// HasCNSIIdList returns a boolean if a field has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) HasCNSIIdList() bool {
	if o != nil && !IsNil(o.CNSIIdList) {
		return true
	}

	return false
}

// SetCNSIIdList gets a reference to the given []string and assigns it to the CNSIIdList field.
func (o *NrfFunctionSingleAllOfAttributesAllOf) SetCNSIIdList(v []string) {
	o.CNSIIdList = v
}

// GetNFProfileList returns the NFProfileList field value if set, zero value otherwise.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetNFProfileList() []NFProfile {
	if o == nil || IsNil(o.NFProfileList) {
		var ret []NFProfile
		return ret
	}
	return o.NFProfileList
}

// GetNFProfileListOk returns a tuple with the NFProfileList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetNFProfileListOk() ([]NFProfile, bool) {
	if o == nil || IsNil(o.NFProfileList) {
		return nil, false
	}
	return o.NFProfileList, true
}

// HasNFProfileList returns a boolean if a field has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) HasNFProfileList() bool {
	if o != nil && !IsNil(o.NFProfileList) {
		return true
	}

	return false
}

// SetNFProfileList gets a reference to the given []NFProfile and assigns it to the NFProfileList field.
func (o *NrfFunctionSingleAllOfAttributesAllOf) SetNFProfileList(v []NFProfile) {
	o.NFProfileList = v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetSnssaiList() []Snssai {
	if o == nil || IsNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SnssaiList) {
		return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *NrfFunctionSingleAllOfAttributesAllOf) HasSnssaiList() bool {
	if o != nil && !IsNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *NrfFunctionSingleAllOfAttributesAllOf) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

func (o NrfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrfFunctionSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlmnIdList) {
		toSerialize["plmnIdList"] = o.PlmnIdList
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.CNSIIdList) {
		toSerialize["cNSIIdList"] = o.CNSIIdList
	}
	if !IsNil(o.NFProfileList) {
		toSerialize["nFProfileList"] = o.NFProfileList
	}
	if !IsNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	return toSerialize, nil
}

type NullableNrfFunctionSingleAllOfAttributesAllOf struct {
	value *NrfFunctionSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableNrfFunctionSingleAllOfAttributesAllOf) Get() *NrfFunctionSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableNrfFunctionSingleAllOfAttributesAllOf) Set(val *NrfFunctionSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfFunctionSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfFunctionSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfFunctionSingleAllOfAttributesAllOf(val *NrfFunctionSingleAllOfAttributesAllOf) *NullableNrfFunctionSingleAllOfAttributesAllOf {
	return &NullableNrfFunctionSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableNrfFunctionSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfFunctionSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


