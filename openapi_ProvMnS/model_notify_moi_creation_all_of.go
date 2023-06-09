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

// checks if the NotifyMoiCreationAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotifyMoiCreationAllOf{}

// NotifyMoiCreationAllOf struct for NotifyMoiCreationAllOf
type NotifyMoiCreationAllOf struct {
	CorrelatedNotifications []CorrelatedNotification `json:"correlatedNotifications,omitempty"`
	AdditionalText          *string                  `json:"additionalText,omitempty"`
	SourceIndicator         *SourceIndicator         `json:"sourceIndicator,omitempty"`
	// The key of this map is the attribute name, and the value the attribute value.
	AttributeList map[string]interface{} `json:"attributeList,omitempty"`
}

// NewNotifyMoiCreationAllOf instantiates a new NotifyMoiCreationAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotifyMoiCreationAllOf() *NotifyMoiCreationAllOf {
	this := NotifyMoiCreationAllOf{}
	return &this
}

// NewNotifyMoiCreationAllOfWithDefaults instantiates a new NotifyMoiCreationAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotifyMoiCreationAllOfWithDefaults() *NotifyMoiCreationAllOf {
	this := NotifyMoiCreationAllOf{}
	return &this
}

// GetCorrelatedNotifications returns the CorrelatedNotifications field value if set, zero value otherwise.
func (o *NotifyMoiCreationAllOf) GetCorrelatedNotifications() []CorrelatedNotification {
	if o == nil || IsNil(o.CorrelatedNotifications) {
		var ret []CorrelatedNotification
		return ret
	}
	return o.CorrelatedNotifications
}

// GetCorrelatedNotificationsOk returns a tuple with the CorrelatedNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreationAllOf) GetCorrelatedNotificationsOk() ([]CorrelatedNotification, bool) {
	if o == nil || IsNil(o.CorrelatedNotifications) {
		return nil, false
	}
	return o.CorrelatedNotifications, true
}

// HasCorrelatedNotifications returns a boolean if a field has been set.
func (o *NotifyMoiCreationAllOf) HasCorrelatedNotifications() bool {
	if o != nil && !IsNil(o.CorrelatedNotifications) {
		return true
	}

	return false
}

// SetCorrelatedNotifications gets a reference to the given []CorrelatedNotification and assigns it to the CorrelatedNotifications field.
func (o *NotifyMoiCreationAllOf) SetCorrelatedNotifications(v []CorrelatedNotification) {
	o.CorrelatedNotifications = v
}

// GetAdditionalText returns the AdditionalText field value if set, zero value otherwise.
func (o *NotifyMoiCreationAllOf) GetAdditionalText() string {
	if o == nil || IsNil(o.AdditionalText) {
		var ret string
		return ret
	}
	return *o.AdditionalText
}

// GetAdditionalTextOk returns a tuple with the AdditionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreationAllOf) GetAdditionalTextOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalText) {
		return nil, false
	}
	return o.AdditionalText, true
}

// HasAdditionalText returns a boolean if a field has been set.
func (o *NotifyMoiCreationAllOf) HasAdditionalText() bool {
	if o != nil && !IsNil(o.AdditionalText) {
		return true
	}

	return false
}

// SetAdditionalText gets a reference to the given string and assigns it to the AdditionalText field.
func (o *NotifyMoiCreationAllOf) SetAdditionalText(v string) {
	o.AdditionalText = &v
}

// GetSourceIndicator returns the SourceIndicator field value if set, zero value otherwise.
func (o *NotifyMoiCreationAllOf) GetSourceIndicator() SourceIndicator {
	if o == nil || IsNil(o.SourceIndicator) {
		var ret SourceIndicator
		return ret
	}
	return *o.SourceIndicator
}

// GetSourceIndicatorOk returns a tuple with the SourceIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreationAllOf) GetSourceIndicatorOk() (*SourceIndicator, bool) {
	if o == nil || IsNil(o.SourceIndicator) {
		return nil, false
	}
	return o.SourceIndicator, true
}

// HasSourceIndicator returns a boolean if a field has been set.
func (o *NotifyMoiCreationAllOf) HasSourceIndicator() bool {
	if o != nil && !IsNil(o.SourceIndicator) {
		return true
	}

	return false
}

// SetSourceIndicator gets a reference to the given SourceIndicator and assigns it to the SourceIndicator field.
func (o *NotifyMoiCreationAllOf) SetSourceIndicator(v SourceIndicator) {
	o.SourceIndicator = &v
}

// GetAttributeList returns the AttributeList field value if set, zero value otherwise.
func (o *NotifyMoiCreationAllOf) GetAttributeList() map[string]interface{} {
	if o == nil || IsNil(o.AttributeList) {
		var ret map[string]interface{}
		return ret
	}
	return o.AttributeList
}

// GetAttributeListOk returns a tuple with the AttributeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotifyMoiCreationAllOf) GetAttributeListOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.AttributeList) {
		return map[string]interface{}{}, false
	}
	return o.AttributeList, true
}

// HasAttributeList returns a boolean if a field has been set.
func (o *NotifyMoiCreationAllOf) HasAttributeList() bool {
	if o != nil && !IsNil(o.AttributeList) {
		return true
	}

	return false
}

// SetAttributeList gets a reference to the given map[string]interface{} and assigns it to the AttributeList field.
func (o *NotifyMoiCreationAllOf) SetAttributeList(v map[string]interface{}) {
	o.AttributeList = v
}

func (o NotifyMoiCreationAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotifyMoiCreationAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CorrelatedNotifications) {
		toSerialize["correlatedNotifications"] = o.CorrelatedNotifications
	}
	if !IsNil(o.AdditionalText) {
		toSerialize["additionalText"] = o.AdditionalText
	}
	if !IsNil(o.SourceIndicator) {
		toSerialize["sourceIndicator"] = o.SourceIndicator
	}
	if !IsNil(o.AttributeList) {
		toSerialize["attributeList"] = o.AttributeList
	}
	return toSerialize, nil
}

type NullableNotifyMoiCreationAllOf struct {
	value *NotifyMoiCreationAllOf
	isSet bool
}

func (v NullableNotifyMoiCreationAllOf) Get() *NotifyMoiCreationAllOf {
	return v.value
}

func (v *NullableNotifyMoiCreationAllOf) Set(val *NotifyMoiCreationAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyMoiCreationAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyMoiCreationAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyMoiCreationAllOf(val *NotifyMoiCreationAllOf) *NullableNotifyMoiCreationAllOf {
	return &NullableNotifyMoiCreationAllOf{value: val, isSet: true}
}

func (v NullableNotifyMoiCreationAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyMoiCreationAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
