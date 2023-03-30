/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the ACCLDisallowedAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACCLDisallowedAttributes{}

// ACCLDisallowedAttributes struct for ACCLDisallowedAttributes
type ACCLDisallowedAttributes struct {
	ManagedEntityIdentifier *string `json:"managedEntityIdentifier,omitempty"`
	AttributeNameList []string `json:"attributeNameList,omitempty"`
}

// NewACCLDisallowedAttributes instantiates a new ACCLDisallowedAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACCLDisallowedAttributes() *ACCLDisallowedAttributes {
	this := ACCLDisallowedAttributes{}
	return &this
}

// NewACCLDisallowedAttributesWithDefaults instantiates a new ACCLDisallowedAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACCLDisallowedAttributesWithDefaults() *ACCLDisallowedAttributes {
	this := ACCLDisallowedAttributes{}
	return &this
}

// GetManagedEntityIdentifier returns the ManagedEntityIdentifier field value if set, zero value otherwise.
func (o *ACCLDisallowedAttributes) GetManagedEntityIdentifier() string {
	if o == nil || IsNil(o.ManagedEntityIdentifier) {
		var ret string
		return ret
	}
	return *o.ManagedEntityIdentifier
}

// GetManagedEntityIdentifierOk returns a tuple with the ManagedEntityIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACCLDisallowedAttributes) GetManagedEntityIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.ManagedEntityIdentifier) {
		return nil, false
	}
	return o.ManagedEntityIdentifier, true
}

// HasManagedEntityIdentifier returns a boolean if a field has been set.
func (o *ACCLDisallowedAttributes) HasManagedEntityIdentifier() bool {
	if o != nil && !IsNil(o.ManagedEntityIdentifier) {
		return true
	}

	return false
}

// SetManagedEntityIdentifier gets a reference to the given string and assigns it to the ManagedEntityIdentifier field.
func (o *ACCLDisallowedAttributes) SetManagedEntityIdentifier(v string) {
	o.ManagedEntityIdentifier = &v
}

// GetAttributeNameList returns the AttributeNameList field value if set, zero value otherwise.
func (o *ACCLDisallowedAttributes) GetAttributeNameList() []string {
	if o == nil || IsNil(o.AttributeNameList) {
		var ret []string
		return ret
	}
	return o.AttributeNameList
}

// GetAttributeNameListOk returns a tuple with the AttributeNameList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACCLDisallowedAttributes) GetAttributeNameListOk() ([]string, bool) {
	if o == nil || IsNil(o.AttributeNameList) {
		return nil, false
	}
	return o.AttributeNameList, true
}

// HasAttributeNameList returns a boolean if a field has been set.
func (o *ACCLDisallowedAttributes) HasAttributeNameList() bool {
	if o != nil && !IsNil(o.AttributeNameList) {
		return true
	}

	return false
}

// SetAttributeNameList gets a reference to the given []string and assigns it to the AttributeNameList field.
func (o *ACCLDisallowedAttributes) SetAttributeNameList(v []string) {
	o.AttributeNameList = v
}

func (o ACCLDisallowedAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACCLDisallowedAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagedEntityIdentifier) {
		toSerialize["managedEntityIdentifier"] = o.ManagedEntityIdentifier
	}
	if !IsNil(o.AttributeNameList) {
		toSerialize["attributeNameList"] = o.AttributeNameList
	}
	return toSerialize, nil
}

type NullableACCLDisallowedAttributes struct {
	value *ACCLDisallowedAttributes
	isSet bool
}

func (v NullableACCLDisallowedAttributes) Get() *ACCLDisallowedAttributes {
	return v.value
}

func (v *NullableACCLDisallowedAttributes) Set(val *ACCLDisallowedAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableACCLDisallowedAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableACCLDisallowedAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACCLDisallowedAttributes(val *ACCLDisallowedAttributes) *NullableACCLDisallowedAttributes {
	return &NullableACCLDisallowedAttributes{value: val, isSet: true}
}

func (v NullableACCLDisallowedAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACCLDisallowedAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


