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

// checks if the IntentSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntentSingle{}

// IntentSingle struct for IntentSingle
type IntentSingle struct {
	Top
	UserLabel            *string                                    `json:"userLabel,omitempty"`
	IntentExpectations   []IntentSingleAllOfIntentExpectationsInner `json:"intentExpectations,omitempty"`
	IntentContexts       []IntentContext                            `json:"intentContexts,omitempty"`
	IntentFulfilmentInfo *FulfilmentInfo                            `json:"intentFulfilmentInfo,omitempty"`
}

// NewIntentSingle instantiates a new IntentSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntentSingle(id NullableString) *IntentSingle {
	this := IntentSingle{}
	this.Id = id
	return &this
}

// NewIntentSingleWithDefaults instantiates a new IntentSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntentSingleWithDefaults() *IntentSingle {
	this := IntentSingle{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *IntentSingle) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *IntentSingle) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *IntentSingle) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetIntentExpectations returns the IntentExpectations field value if set, zero value otherwise.
func (o *IntentSingle) GetIntentExpectations() []IntentSingleAllOfIntentExpectationsInner {
	if o == nil || IsNil(o.IntentExpectations) {
		var ret []IntentSingleAllOfIntentExpectationsInner
		return ret
	}
	return o.IntentExpectations
}

// GetIntentExpectationsOk returns a tuple with the IntentExpectations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetIntentExpectationsOk() ([]IntentSingleAllOfIntentExpectationsInner, bool) {
	if o == nil || IsNil(o.IntentExpectations) {
		return nil, false
	}
	return o.IntentExpectations, true
}

// HasIntentExpectations returns a boolean if a field has been set.
func (o *IntentSingle) HasIntentExpectations() bool {
	if o != nil && !IsNil(o.IntentExpectations) {
		return true
	}

	return false
}

// SetIntentExpectations gets a reference to the given []IntentSingleAllOfIntentExpectationsInner and assigns it to the IntentExpectations field.
func (o *IntentSingle) SetIntentExpectations(v []IntentSingleAllOfIntentExpectationsInner) {
	o.IntentExpectations = v
}

// GetIntentContexts returns the IntentContexts field value if set, zero value otherwise.
func (o *IntentSingle) GetIntentContexts() []IntentContext {
	if o == nil || IsNil(o.IntentContexts) {
		var ret []IntentContext
		return ret
	}
	return o.IntentContexts
}

// GetIntentContextsOk returns a tuple with the IntentContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetIntentContextsOk() ([]IntentContext, bool) {
	if o == nil || IsNil(o.IntentContexts) {
		return nil, false
	}
	return o.IntentContexts, true
}

// HasIntentContexts returns a boolean if a field has been set.
func (o *IntentSingle) HasIntentContexts() bool {
	if o != nil && !IsNil(o.IntentContexts) {
		return true
	}

	return false
}

// SetIntentContexts gets a reference to the given []IntentContext and assigns it to the IntentContexts field.
func (o *IntentSingle) SetIntentContexts(v []IntentContext) {
	o.IntentContexts = v
}

// GetIntentFulfilmentInfo returns the IntentFulfilmentInfo field value if set, zero value otherwise.
func (o *IntentSingle) GetIntentFulfilmentInfo() FulfilmentInfo {
	if o == nil || IsNil(o.IntentFulfilmentInfo) {
		var ret FulfilmentInfo
		return ret
	}
	return *o.IntentFulfilmentInfo
}

// GetIntentFulfilmentInfoOk returns a tuple with the IntentFulfilmentInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntentSingle) GetIntentFulfilmentInfoOk() (*FulfilmentInfo, bool) {
	if o == nil || IsNil(o.IntentFulfilmentInfo) {
		return nil, false
	}
	return o.IntentFulfilmentInfo, true
}

// HasIntentFulfilmentInfo returns a boolean if a field has been set.
func (o *IntentSingle) HasIntentFulfilmentInfo() bool {
	if o != nil && !IsNil(o.IntentFulfilmentInfo) {
		return true
	}

	return false
}

// SetIntentFulfilmentInfo gets a reference to the given FulfilmentInfo and assigns it to the IntentFulfilmentInfo field.
func (o *IntentSingle) SetIntentFulfilmentInfo(v FulfilmentInfo) {
	o.IntentFulfilmentInfo = &v
}

func (o IntentSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntentSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
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

type NullableIntentSingle struct {
	value *IntentSingle
	isSet bool
}

func (v NullableIntentSingle) Get() *IntentSingle {
	return v.value
}

func (v *NullableIntentSingle) Set(val *IntentSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableIntentSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableIntentSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntentSingle(val *IntentSingle) *NullableIntentSingle {
	return &NullableIntentSingle{value: val, isSet: true}
}

func (v NullableIntentSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntentSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
