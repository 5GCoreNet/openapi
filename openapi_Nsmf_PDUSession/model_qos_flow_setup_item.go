/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the QosFlowSetupItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QosFlowSetupItem{}

// QosFlowSetupItem Individual QoS flow to setup
type QosFlowSetupItem struct {
	// Unsigned integer identifying a QoS flow, within the range 0 to 63.
	Qfi int32 `json:"qfi"`
	// string with format 'bytes' as defined in OpenAPI
	QosRules string `json:"qosRules"`
	// EPS Bearer Identifier
	Ebi *int32 `json:"ebi,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	QosFlowDescription *string `json:"qosFlowDescription,omitempty"`
	QosFlowProfile *QosFlowProfile `json:"qosFlowProfile,omitempty"`
	AssociatedAnType *QosFlowAccessType `json:"associatedAnType,omitempty"`
	DefaultQosRuleInd *bool `json:"defaultQosRuleInd,omitempty"`
}

// NewQosFlowSetupItem instantiates a new QosFlowSetupItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQosFlowSetupItem(qfi int32, qosRules string) *QosFlowSetupItem {
	this := QosFlowSetupItem{}
	this.Qfi = qfi
	this.QosRules = qosRules
	return &this
}

// NewQosFlowSetupItemWithDefaults instantiates a new QosFlowSetupItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQosFlowSetupItemWithDefaults() *QosFlowSetupItem {
	this := QosFlowSetupItem{}
	return &this
}

// GetQfi returns the Qfi field value
func (o *QosFlowSetupItem) GetQfi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Qfi
}

// GetQfiOk returns a tuple with the Qfi field value
// and a boolean to check if the value has been set.
func (o *QosFlowSetupItem) GetQfiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Qfi, true
}

// SetQfi sets field value
func (o *QosFlowSetupItem) SetQfi(v int32) {
	o.Qfi = v
}

// GetQosRules returns the QosRules field value
func (o *QosFlowSetupItem) GetQosRules() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QosRules
}

// GetQosRulesOk returns a tuple with the QosRules field value
// and a boolean to check if the value has been set.
func (o *QosFlowSetupItem) GetQosRulesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QosRules, true
}

// SetQosRules sets field value
func (o *QosFlowSetupItem) SetQosRules(v string) {
	o.QosRules = v
}

// GetEbi returns the Ebi field value if set, zero value otherwise.
func (o *QosFlowSetupItem) GetEbi() int32 {
	if o == nil || IsNil(o.Ebi) {
		var ret int32
		return ret
	}
	return *o.Ebi
}

// GetEbiOk returns a tuple with the Ebi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowSetupItem) GetEbiOk() (*int32, bool) {
	if o == nil || IsNil(o.Ebi) {
		return nil, false
	}
	return o.Ebi, true
}

// HasEbi returns a boolean if a field has been set.
func (o *QosFlowSetupItem) HasEbi() bool {
	if o != nil && !IsNil(o.Ebi) {
		return true
	}

	return false
}

// SetEbi gets a reference to the given int32 and assigns it to the Ebi field.
func (o *QosFlowSetupItem) SetEbi(v int32) {
	o.Ebi = &v
}

// GetQosFlowDescription returns the QosFlowDescription field value if set, zero value otherwise.
func (o *QosFlowSetupItem) GetQosFlowDescription() string {
	if o == nil || IsNil(o.QosFlowDescription) {
		var ret string
		return ret
	}
	return *o.QosFlowDescription
}

// GetQosFlowDescriptionOk returns a tuple with the QosFlowDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowSetupItem) GetQosFlowDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.QosFlowDescription) {
		return nil, false
	}
	return o.QosFlowDescription, true
}

// HasQosFlowDescription returns a boolean if a field has been set.
func (o *QosFlowSetupItem) HasQosFlowDescription() bool {
	if o != nil && !IsNil(o.QosFlowDescription) {
		return true
	}

	return false
}

// SetQosFlowDescription gets a reference to the given string and assigns it to the QosFlowDescription field.
func (o *QosFlowSetupItem) SetQosFlowDescription(v string) {
	o.QosFlowDescription = &v
}

// GetQosFlowProfile returns the QosFlowProfile field value if set, zero value otherwise.
func (o *QosFlowSetupItem) GetQosFlowProfile() QosFlowProfile {
	if o == nil || IsNil(o.QosFlowProfile) {
		var ret QosFlowProfile
		return ret
	}
	return *o.QosFlowProfile
}

// GetQosFlowProfileOk returns a tuple with the QosFlowProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowSetupItem) GetQosFlowProfileOk() (*QosFlowProfile, bool) {
	if o == nil || IsNil(o.QosFlowProfile) {
		return nil, false
	}
	return o.QosFlowProfile, true
}

// HasQosFlowProfile returns a boolean if a field has been set.
func (o *QosFlowSetupItem) HasQosFlowProfile() bool {
	if o != nil && !IsNil(o.QosFlowProfile) {
		return true
	}

	return false
}

// SetQosFlowProfile gets a reference to the given QosFlowProfile and assigns it to the QosFlowProfile field.
func (o *QosFlowSetupItem) SetQosFlowProfile(v QosFlowProfile) {
	o.QosFlowProfile = &v
}

// GetAssociatedAnType returns the AssociatedAnType field value if set, zero value otherwise.
func (o *QosFlowSetupItem) GetAssociatedAnType() QosFlowAccessType {
	if o == nil || IsNil(o.AssociatedAnType) {
		var ret QosFlowAccessType
		return ret
	}
	return *o.AssociatedAnType
}

// GetAssociatedAnTypeOk returns a tuple with the AssociatedAnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowSetupItem) GetAssociatedAnTypeOk() (*QosFlowAccessType, bool) {
	if o == nil || IsNil(o.AssociatedAnType) {
		return nil, false
	}
	return o.AssociatedAnType, true
}

// HasAssociatedAnType returns a boolean if a field has been set.
func (o *QosFlowSetupItem) HasAssociatedAnType() bool {
	if o != nil && !IsNil(o.AssociatedAnType) {
		return true
	}

	return false
}

// SetAssociatedAnType gets a reference to the given QosFlowAccessType and assigns it to the AssociatedAnType field.
func (o *QosFlowSetupItem) SetAssociatedAnType(v QosFlowAccessType) {
	o.AssociatedAnType = &v
}

// GetDefaultQosRuleInd returns the DefaultQosRuleInd field value if set, zero value otherwise.
func (o *QosFlowSetupItem) GetDefaultQosRuleInd() bool {
	if o == nil || IsNil(o.DefaultQosRuleInd) {
		var ret bool
		return ret
	}
	return *o.DefaultQosRuleInd
}

// GetDefaultQosRuleIndOk returns a tuple with the DefaultQosRuleInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QosFlowSetupItem) GetDefaultQosRuleIndOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultQosRuleInd) {
		return nil, false
	}
	return o.DefaultQosRuleInd, true
}

// HasDefaultQosRuleInd returns a boolean if a field has been set.
func (o *QosFlowSetupItem) HasDefaultQosRuleInd() bool {
	if o != nil && !IsNil(o.DefaultQosRuleInd) {
		return true
	}

	return false
}

// SetDefaultQosRuleInd gets a reference to the given bool and assigns it to the DefaultQosRuleInd field.
func (o *QosFlowSetupItem) SetDefaultQosRuleInd(v bool) {
	o.DefaultQosRuleInd = &v
}

func (o QosFlowSetupItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QosFlowSetupItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["qfi"] = o.Qfi
	toSerialize["qosRules"] = o.QosRules
	if !IsNil(o.Ebi) {
		toSerialize["ebi"] = o.Ebi
	}
	if !IsNil(o.QosFlowDescription) {
		toSerialize["qosFlowDescription"] = o.QosFlowDescription
	}
	if !IsNil(o.QosFlowProfile) {
		toSerialize["qosFlowProfile"] = o.QosFlowProfile
	}
	if !IsNil(o.AssociatedAnType) {
		toSerialize["associatedAnType"] = o.AssociatedAnType
	}
	if !IsNil(o.DefaultQosRuleInd) {
		toSerialize["defaultQosRuleInd"] = o.DefaultQosRuleInd
	}
	return toSerialize, nil
}

type NullableQosFlowSetupItem struct {
	value *QosFlowSetupItem
	isSet bool
}

func (v NullableQosFlowSetupItem) Get() *QosFlowSetupItem {
	return v.value
}

func (v *NullableQosFlowSetupItem) Set(val *QosFlowSetupItem) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowSetupItem) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowSetupItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowSetupItem(val *QosFlowSetupItem) *NullableQosFlowSetupItem {
	return &NullableQosFlowSetupItem{value: val, isSet: true}
}

func (v NullableQosFlowSetupItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowSetupItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


