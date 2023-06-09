/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the IntentSingleAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntentSingleAllOf{}

// IntentSingleAllOf struct for IntentSingleAllOf
type IntentSingleAllOf struct {
	UserLabel            *string                                    `json:"userLabel,omitempty"`
	IntentExpectations   []IntentSingleAllOfIntentExpectationsInner `json:"intentExpectations,omitempty"`
	IntentContexts       []IntentContext                            `json:"intentContexts,omitempty"`
	IntentFulfilmentInfo *FulfilmentInfo                            `json:"intentFulfilmentInfo,omitempty"`
}

// NewIntentSingleAllOf instantiates a new IntentSingleAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntentSingleAllOf() *IntentSingleAllOf {
	this := IntentSingleAllOf{}
	return &this
}

// NewIntentSingleAllOfWithDefaults instantiates a new IntentSingleAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntentSingleAllOfWithDefaults() *IntentSingleAllOf {
	this := IntentSingleAllOf{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *IntentSingleAllOf) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingleAllOf) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *IntentSingleAllOf) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *IntentSingleAllOf) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetIntentExpectations returns the IntentExpectations field value if set, zero value otherwise.
func (o *IntentSingleAllOf) GetIntentExpectations() []IntentSingleAllOfIntentExpectationsInner {
	if o == nil || IsNil(o.IntentExpectations) {
		var ret []IntentSingleAllOfIntentExpectationsInner
		return ret
	}
	return o.IntentExpectations
}

// GetIntentExpectationsOk returns a tuple with the IntentExpectations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingleAllOf) GetIntentExpectationsOk() ([]IntentSingleAllOfIntentExpectationsInner, bool) {
	if o == nil || IsNil(o.IntentExpectations) {
		return nil, false
	}
	return o.IntentExpectations, true
}

// HasIntentExpectations returns a boolean if a field has been set.
func (o *IntentSingleAllOf) HasIntentExpectations() bool {
	if o != nil && !IsNil(o.IntentExpectations) {
		return true
	}

	return false
}

// SetIntentExpectations gets a reference to the given []IntentSingleAllOfIntentExpectationsInner and assigns it to the IntentExpectations field.
func (o *IntentSingleAllOf) SetIntentExpectations(v []IntentSingleAllOfIntentExpectationsInner) {
	o.IntentExpectations = v
}

// GetIntentContexts returns the IntentContexts field value if set, zero value otherwise.
func (o *IntentSingleAllOf) GetIntentContexts() []IntentContext {
	if o == nil || IsNil(o.IntentContexts) {
		var ret []IntentContext
		return ret
	}
	return o.IntentContexts
}

// GetIntentContextsOk returns a tuple with the IntentContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingleAllOf) GetIntentContextsOk() ([]IntentContext, bool) {
	if o == nil || IsNil(o.IntentContexts) {
		return nil, false
	}
	return o.IntentContexts, true
}

// HasIntentContexts returns a boolean if a field has been set.
func (o *IntentSingleAllOf) HasIntentContexts() bool {
	if o != nil && !IsNil(o.IntentContexts) {
		return true
	}

	return false
}

// SetIntentContexts gets a reference to the given []IntentContext and assigns it to the IntentContexts field.
func (o *IntentSingleAllOf) SetIntentContexts(v []IntentContext) {
	o.IntentContexts = v
}

// GetIntentFulfilmentInfo returns the IntentFulfilmentInfo field value if set, zero value otherwise.
func (o *IntentSingleAllOf) GetIntentFulfilmentInfo() FulfilmentInfo {
	if o == nil || IsNil(o.IntentFulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.IntentFulfilmentInfo
}

// GetIntentFulfilmentInfoOk returns a tuple with the IntentFulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingleAllOf) GetIntentFulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || IsNil(o.IntentFulfilmentInfo) {
		return nil, false
	}
	return o.IntentFulfilmentInfo, true
}

// HasIntentFulfilmentInfo returns a boolean if a field has been set.
func (o *IntentSingleAllOf) HasIntentFulfilmentInfo() bool {
	if o != nil && !IsNil(o.IntentFulfilmentInfo) {
		return true
	}

	return false
}

// SetIntentFulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the IntentFulfilmentInfo field.
func (o *IntentSingleAllOf) SetIntentFulfilmentInfo(v FulfilmentInfo) {
	o.IntentFulfilmentInfo = &v
}

func (o IntentSingleAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntentSingleAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !IsNil(o.IntentExpectations) {
		toSerialize["intentExpectations"] = o.IntentExpectations
	}
	if !IsNil(o.IntentContexts) {
		toSerialize["intentContexts"] = o.IntentContexts
	}
	if !IsNil(o.IntentFulfilmentInfo) {
		toSerialize["intentFulfilmentInfo"] = o.IntentFulfilmentInfo
	}
	return toSerialize, nil
}

type NullableIntentSingleAllOf struct {
	value *IntentSingleAllOf
	isSet bool
}

func (v NullableIntentSingleAllOf) Get() *IntentSingleAllOf {
	return v.value
}

func (v *NullableIntentSingleAllOf) Set(val *IntentSingleAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIntentSingleAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIntentSingleAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntentSingleAllOf(val *IntentSingleAllOf) *NullableIntentSingleAllOf {
	return &NullableIntentSingleAllOf{value: val, isSet: true}
}

func (v NullableIntentSingleAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntentSingleAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
