/*
M1_PolicyTemplatesProvisioning

5GMS AF M1 Policy Templates Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_PolicyTemplatesProvisioning

import (
	"encoding/json"
)

// checks if the PolicyTemplateApplicationSessionContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyTemplateApplicationSessionContext{}

// PolicyTemplateApplicationSessionContext struct for PolicyTemplateApplicationSessionContext
type PolicyTemplateApplicationSessionContext struct {
	// Contains an AF application identifier.
	AfAppId   *string `json:"afAppId,omitempty"`
	SliceInfo *Snssai `json:"sliceInfo,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn *string `json:"dnn,omitempty"`
	// Contains an identity of an application service provider.
	AspId *string `json:"aspId,omitempty"`
}

// NewPolicyTemplateApplicationSessionContext instantiates a new PolicyTemplateApplicationSessionContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyTemplateApplicationSessionContext() *PolicyTemplateApplicationSessionContext {
	this := PolicyTemplateApplicationSessionContext{}
	return &this
}

// NewPolicyTemplateApplicationSessionContextWithDefaults instantiates a new PolicyTemplateApplicationSessionContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyTemplateApplicationSessionContextWithDefaults() *PolicyTemplateApplicationSessionContext {
	this := PolicyTemplateApplicationSessionContext{}
	return &this
}

// GetAfAppId returns the AfAppId field value if set, zero value otherwise.
func (o *PolicyTemplateApplicationSessionContext) GetAfAppId() string {
	if o == nil || IsNil(o.AfAppId) {
		var ret string
		return ret
	}
	return *o.AfAppId
}

// GetAfAppIdOk returns a tuple with the AfAppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTemplateApplicationSessionContext) GetAfAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfAppId) {
		return nil, false
	}
	return o.AfAppId, true
}

// HasAfAppId returns a boolean if a field has been set.
func (o *PolicyTemplateApplicationSessionContext) HasAfAppId() bool {
	if o != nil && !IsNil(o.AfAppId) {
		return true
	}

	return false
}

// SetAfAppId gets a reference to the given string and assigns it to the AfAppId field.
func (o *PolicyTemplateApplicationSessionContext) SetAfAppId(v string) {
	o.AfAppId = &v
}

// GetSliceInfo returns the SliceInfo field value if set, zero value otherwise.
func (o *PolicyTemplateApplicationSessionContext) GetSliceInfo() Snssai {
	if o == nil || IsNil(o.SliceInfo) {
		var ret Snssai
		return ret
	}
	return *o.SliceInfo
}

// GetSliceInfoOk returns a tuple with the SliceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTemplateApplicationSessionContext) GetSliceInfoOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SliceInfo) {
		return nil, false
	}
	return o.SliceInfo, true
}

// HasSliceInfo returns a boolean if a field has been set.
func (o *PolicyTemplateApplicationSessionContext) HasSliceInfo() bool {
	if o != nil && !IsNil(o.SliceInfo) {
		return true
	}

	return false
}

// SetSliceInfo gets a reference to the given Snssai and assigns it to the SliceInfo field.
func (o *PolicyTemplateApplicationSessionContext) SetSliceInfo(v Snssai) {
	o.SliceInfo = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *PolicyTemplateApplicationSessionContext) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTemplateApplicationSessionContext) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *PolicyTemplateApplicationSessionContext) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *PolicyTemplateApplicationSessionContext) SetDnn(v string) {
	o.Dnn = &v
}

// GetAspId returns the AspId field value if set, zero value otherwise.
func (o *PolicyTemplateApplicationSessionContext) GetAspId() string {
	if o == nil || IsNil(o.AspId) {
		var ret string
		return ret
	}
	return *o.AspId
}

// GetAspIdOk returns a tuple with the AspId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyTemplateApplicationSessionContext) GetAspIdOk() (*string, bool) {
	if o == nil || IsNil(o.AspId) {
		return nil, false
	}
	return o.AspId, true
}

// HasAspId returns a boolean if a field has been set.
func (o *PolicyTemplateApplicationSessionContext) HasAspId() bool {
	if o != nil && !IsNil(o.AspId) {
		return true
	}

	return false
}

// SetAspId gets a reference to the given string and assigns it to the AspId field.
func (o *PolicyTemplateApplicationSessionContext) SetAspId(v string) {
	o.AspId = &v
}

func (o PolicyTemplateApplicationSessionContext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyTemplateApplicationSessionContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AfAppId) {
		toSerialize["afAppId"] = o.AfAppId
	}
	if !IsNil(o.SliceInfo) {
		toSerialize["sliceInfo"] = o.SliceInfo
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.AspId) {
		toSerialize["aspId"] = o.AspId
	}
	return toSerialize, nil
}

type NullablePolicyTemplateApplicationSessionContext struct {
	value *PolicyTemplateApplicationSessionContext
	isSet bool
}

func (v NullablePolicyTemplateApplicationSessionContext) Get() *PolicyTemplateApplicationSessionContext {
	return v.value
}

func (v *NullablePolicyTemplateApplicationSessionContext) Set(val *PolicyTemplateApplicationSessionContext) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyTemplateApplicationSessionContext) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyTemplateApplicationSessionContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyTemplateApplicationSessionContext(val *PolicyTemplateApplicationSessionContext) *NullablePolicyTemplateApplicationSessionContext {
	return &NullablePolicyTemplateApplicationSessionContext{value: val, isSet: true}
}

func (v NullablePolicyTemplateApplicationSessionContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyTemplateApplicationSessionContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
