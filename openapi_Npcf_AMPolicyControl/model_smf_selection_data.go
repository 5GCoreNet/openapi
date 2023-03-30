/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_AMPolicyControl

import (
	"encoding/json"
)

// checks if the SmfSelectionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmfSelectionData{}

// SmfSelectionData Represents the SMF Selection information that may be replaced by the PCF.
type SmfSelectionData struct {
	UnsuppDnn *bool `json:"unsuppDnn,omitempty"`
	// Contains the list of DNNs per S-NSSAI that are candidates for replacement. The snssai attribute within the CandidateForReplacement data type is the key of the map. 
	Candidates map[string]CandidateForReplacement `json:"candidates,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
	MappingSnssai *Snssai `json:"mappingSnssai,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
}

// NewSmfSelectionData instantiates a new SmfSelectionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmfSelectionData() *SmfSelectionData {
	this := SmfSelectionData{}
	return &this
}

// NewSmfSelectionDataWithDefaults instantiates a new SmfSelectionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmfSelectionDataWithDefaults() *SmfSelectionData {
	this := SmfSelectionData{}
	return &this
}

// GetUnsuppDnn returns the UnsuppDnn field value if set, zero value otherwise.
func (o *SmfSelectionData) GetUnsuppDnn() bool {
	if o == nil || IsNil(o.UnsuppDnn) {
		var ret bool
		return ret
	}
	return *o.UnsuppDnn
}

// GetUnsuppDnnOk returns a tuple with the UnsuppDnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfSelectionData) GetUnsuppDnnOk() (*bool, bool) {
	if o == nil || IsNil(o.UnsuppDnn) {
		return nil, false
	}
	return o.UnsuppDnn, true
}

// HasUnsuppDnn returns a boolean if a field has been set.
func (o *SmfSelectionData) HasUnsuppDnn() bool {
	if o != nil && !IsNil(o.UnsuppDnn) {
		return true
	}

	return false
}

// SetUnsuppDnn gets a reference to the given bool and assigns it to the UnsuppDnn field.
func (o *SmfSelectionData) SetUnsuppDnn(v bool) {
	o.UnsuppDnn = &v
}

// GetCandidates returns the Candidates field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SmfSelectionData) GetCandidates() map[string]CandidateForReplacement {
	if o == nil {
		var ret map[string]CandidateForReplacement
		return ret
	}
	return o.Candidates
}

// GetCandidatesOk returns a tuple with the Candidates field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SmfSelectionData) GetCandidatesOk() (*map[string]CandidateForReplacement, bool) {
	if o == nil || IsNil(o.Candidates) {
		return nil, false
	}
	return &o.Candidates, true
}

// HasCandidates returns a boolean if a field has been set.
func (o *SmfSelectionData) HasCandidates() bool {
	if o != nil && IsNil(o.Candidates) {
		return true
	}

	return false
}

// SetCandidates gets a reference to the given map[string]CandidateForReplacement and assigns it to the Candidates field.
func (o *SmfSelectionData) SetCandidates(v map[string]CandidateForReplacement) {
	o.Candidates = v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *SmfSelectionData) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfSelectionData) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *SmfSelectionData) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *SmfSelectionData) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetMappingSnssai returns the MappingSnssai field value if set, zero value otherwise.
func (o *SmfSelectionData) GetMappingSnssai() Snssai {
	if o == nil || IsNil(o.MappingSnssai) {
		var ret Snssai
		return ret
	}
	return *o.MappingSnssai
}

// GetMappingSnssaiOk returns a tuple with the MappingSnssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfSelectionData) GetMappingSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.MappingSnssai) {
		return nil, false
	}
	return o.MappingSnssai, true
}

// HasMappingSnssai returns a boolean if a field has been set.
func (o *SmfSelectionData) HasMappingSnssai() bool {
	if o != nil && !IsNil(o.MappingSnssai) {
		return true
	}

	return false
}

// SetMappingSnssai gets a reference to the given Snssai and assigns it to the MappingSnssai field.
func (o *SmfSelectionData) SetMappingSnssai(v Snssai) {
	o.MappingSnssai = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *SmfSelectionData) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmfSelectionData) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *SmfSelectionData) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *SmfSelectionData) SetDnn(v string) {
	o.Dnn = &v
}

func (o SmfSelectionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmfSelectionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UnsuppDnn) {
		toSerialize["unsuppDnn"] = o.UnsuppDnn
	}
	if o.Candidates != nil {
		toSerialize["candidates"] = o.Candidates
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	if !IsNil(o.MappingSnssai) {
		toSerialize["mappingSnssai"] = o.MappingSnssai
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	return toSerialize, nil
}

type NullableSmfSelectionData struct {
	value *SmfSelectionData
	isSet bool
}

func (v NullableSmfSelectionData) Get() *SmfSelectionData {
	return v.value
}

func (v *NullableSmfSelectionData) Set(val *SmfSelectionData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmfSelectionData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmfSelectionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmfSelectionData(val *SmfSelectionData) *NullableSmfSelectionData {
	return &NullableSmfSelectionData{value: val, isSet: true}
}

func (v NullableSmfSelectionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmfSelectionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


