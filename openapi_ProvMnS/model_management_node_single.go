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

// checks if the ManagementNodeSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementNodeSingle{}

// ManagementNodeSingle struct for ManagementNodeSingle
type ManagementNodeSingle struct {
	Top
	Attributes *ManagementNodeSingleAllOfAttributes `json:"attributes,omitempty"`
	MnsAgent   []MnsAgentSingle                     `json:"MnsAgent,omitempty"`
}

// NewManagementNodeSingle instantiates a new ManagementNodeSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementNodeSingle(id NullableString) *ManagementNodeSingle {
	this := ManagementNodeSingle{}
	this.Id = id
	return &this
}

// NewManagementNodeSingleWithDefaults instantiates a new ManagementNodeSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementNodeSingleWithDefaults() *ManagementNodeSingle {
	this := ManagementNodeSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *ManagementNodeSingle) GetAttributes() ManagementNodeSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret ManagementNodeSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingle) GetAttributesOk() (*ManagementNodeSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *ManagementNodeSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given ManagementNodeSingleAllOfAttributes and assigns it to the Attributes field.
func (o *ManagementNodeSingle) SetAttributes(v ManagementNodeSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetMnsAgent returns the MnsAgent field value if set, zero value otherwise.
func (o *ManagementNodeSingle) GetMnsAgent() []MnsAgentSingle {
	if o == nil || IsNil(o.MnsAgent) {
		var ret []MnsAgentSingle
		return ret
	}
	return o.MnsAgent
}

// GetMnsAgentOk returns a tuple with the MnsAgent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingle) GetMnsAgentOk() ([]MnsAgentSingle, bool) {
	if o == nil || IsNil(o.MnsAgent) {
		return nil, false
	}
	return o.MnsAgent, true
}

// HasMnsAgent returns a boolean if a field has been set.
func (o *ManagementNodeSingle) HasMnsAgent() bool {
	if o != nil && !IsNil(o.MnsAgent) {
		return true
	}

	return false
}

// SetMnsAgent gets a reference to the given []MnsAgentSingle and assigns it to the MnsAgent field.
func (o *ManagementNodeSingle) SetMnsAgent(v []MnsAgentSingle) {
	o.MnsAgent = v
}

func (o ManagementNodeSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementNodeSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.MnsAgent) {
		toSerialize["MnsAgent"] = o.MnsAgent
	}
	return toSerialize, nil
}

type NullableManagementNodeSingle struct {
	value *ManagementNodeSingle
	isSet bool
}

func (v NullableManagementNodeSingle) Get() *ManagementNodeSingle {
	return v.value
}

func (v *NullableManagementNodeSingle) Set(val *ManagementNodeSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementNodeSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementNodeSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementNodeSingle(val *ManagementNodeSingle) *NullableManagementNodeSingle {
	return &NullableManagementNodeSingle{value: val, isSet: true}
}

func (v NullableManagementNodeSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementNodeSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
