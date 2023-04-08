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

// checks if the RrmPolicyAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RrmPolicyAttr{}

// RrmPolicyAttr struct for RrmPolicyAttr
type RrmPolicyAttr struct {
	ResourceType *ResourceType `json:"resourceType,omitempty"`
	RRMPolicyMemberList []RrmPolicyMember `json:"rRMPolicyMemberList,omitempty"`
}

// NewRrmPolicyAttr instantiates a new RrmPolicyAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRrmPolicyAttr() *RrmPolicyAttr {
	this := RrmPolicyAttr{}
	return &this
}

// NewRrmPolicyAttrWithDefaults instantiates a new RrmPolicyAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRrmPolicyAttrWithDefaults() *RrmPolicyAttr {
	this := RrmPolicyAttr{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *RrmPolicyAttr) GetResourceType() ResourceType {
	if o == nil || isNil(o.ResourceType) {
		var ret ResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmPolicyAttr) GetResourceTypeOk() (*ResourceType, bool) {
	if o == nil || isNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *RrmPolicyAttr) HasResourceType() bool {
	if o != nil && !isNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given ResourceType and assigns it to the ResourceType field.
func (o *RrmPolicyAttr) SetResourceType(v ResourceType) {
	o.ResourceType = &v
}

// GetRRMPolicyMemberList returns the RRMPolicyMemberList field value if set, zero value otherwise.
func (o *RrmPolicyAttr) GetRRMPolicyMemberList() []RrmPolicyMember {
	if o == nil || isNil(o.RRMPolicyMemberList) {
		var ret []RrmPolicyMember
		return ret
	}
	return o.RRMPolicyMemberList
}

// GetRRMPolicyMemberListOk returns a tuple with the RRMPolicyMemberList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmPolicyAttr) GetRRMPolicyMemberListOk() ([]RrmPolicyMember, bool) {
	if o == nil || isNil(o.RRMPolicyMemberList) {
		return nil, false
	}
	return o.RRMPolicyMemberList, true
}

// HasRRMPolicyMemberList returns a boolean if a field has been set.
func (o *RrmPolicyAttr) HasRRMPolicyMemberList() bool {
	if o != nil && !isNil(o.RRMPolicyMemberList) {
		return true
	}

	return false
}

// SetRRMPolicyMemberList gets a reference to the given []RrmPolicyMember and assigns it to the RRMPolicyMemberList field.
func (o *RrmPolicyAttr) SetRRMPolicyMemberList(v []RrmPolicyMember) {
	o.RRMPolicyMemberList = v
}

func (o RrmPolicyAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RrmPolicyAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !isNil(o.RRMPolicyMemberList) {
		toSerialize["rRMPolicyMemberList"] = o.RRMPolicyMemberList
	}
	return toSerialize, nil
}

type NullableRrmPolicyAttr struct {
	value *RrmPolicyAttr
	isSet bool
}

func (v NullableRrmPolicyAttr) Get() *RrmPolicyAttr {
	return v.value
}

func (v *NullableRrmPolicyAttr) Set(val *RrmPolicyAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableRrmPolicyAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableRrmPolicyAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRrmPolicyAttr(val *RrmPolicyAttr) *NullableRrmPolicyAttr {
	return &NullableRrmPolicyAttr{value: val, isSet: true}
}

func (v NullableRrmPolicyAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRrmPolicyAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


