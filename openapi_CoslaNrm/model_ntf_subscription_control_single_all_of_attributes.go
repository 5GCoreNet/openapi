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

// checks if the NtfSubscriptionControlSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NtfSubscriptionControlSingleAllOfAttributes{}

// NtfSubscriptionControlSingleAllOfAttributes struct for NtfSubscriptionControlSingleAllOfAttributes
type NtfSubscriptionControlSingleAllOfAttributes struct {
	NotificationRecipientAddress *string `json:"notificationRecipientAddress,omitempty"`
	NotificationTypes []NotificationType `json:"notificationTypes,omitempty"`
	Scope *Scope `json:"scope,omitempty"`
	// The filter format shall be compliant to XPath 1.0.
	NotificationFilter *string `json:"notificationFilter,omitempty"`
}

// NewNtfSubscriptionControlSingleAllOfAttributes instantiates a new NtfSubscriptionControlSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNtfSubscriptionControlSingleAllOfAttributes() *NtfSubscriptionControlSingleAllOfAttributes {
	this := NtfSubscriptionControlSingleAllOfAttributes{}
	return &this
}

// NewNtfSubscriptionControlSingleAllOfAttributesWithDefaults instantiates a new NtfSubscriptionControlSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNtfSubscriptionControlSingleAllOfAttributesWithDefaults() *NtfSubscriptionControlSingleAllOfAttributes {
	this := NtfSubscriptionControlSingleAllOfAttributes{}
	return &this
}

// GetNotificationRecipientAddress returns the NotificationRecipientAddress field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationRecipientAddress() string {
	if o == nil || IsNil(o.NotificationRecipientAddress) {
		var ret string
		return ret
	}
	return *o.NotificationRecipientAddress
}

// GetNotificationRecipientAddressOk returns a tuple with the NotificationRecipientAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationRecipientAddressOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationRecipientAddress) {
		return nil, false
	}
	return o.NotificationRecipientAddress, true
}

// HasNotificationRecipientAddress returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingleAllOfAttributes) HasNotificationRecipientAddress() bool {
	if o != nil && !IsNil(o.NotificationRecipientAddress) {
		return true
	}

	return false
}

// SetNotificationRecipientAddress gets a reference to the given string and assigns it to the NotificationRecipientAddress field.
func (o *NtfSubscriptionControlSingleAllOfAttributes) SetNotificationRecipientAddress(v string) {
	o.NotificationRecipientAddress = &v
}

// GetNotificationTypes returns the NotificationTypes field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationTypes() []NotificationType {
	if o == nil || IsNil(o.NotificationTypes) {
		var ret []NotificationType
		return ret
	}
	return o.NotificationTypes
}

// GetNotificationTypesOk returns a tuple with the NotificationTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationTypesOk() ([]NotificationType, bool) {
	if o == nil || IsNil(o.NotificationTypes) {
		return nil, false
	}
	return o.NotificationTypes, true
}

// HasNotificationTypes returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingleAllOfAttributes) HasNotificationTypes() bool {
	if o != nil && !IsNil(o.NotificationTypes) {
		return true
	}

	return false
}

// SetNotificationTypes gets a reference to the given []NotificationType and assigns it to the NotificationTypes field.
func (o *NtfSubscriptionControlSingleAllOfAttributes) SetNotificationTypes(v []NotificationType) {
	o.NotificationTypes = v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingleAllOfAttributes) GetScope() Scope {
	if o == nil || IsNil(o.Scope) {
		var ret Scope
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingleAllOfAttributes) GetScopeOk() (*Scope, bool) {
	if o == nil || IsNil(o.Scope) {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingleAllOfAttributes) HasScope() bool {
	if o != nil && !IsNil(o.Scope) {
		return true
	}

	return false
}

// SetScope gets a reference to the given Scope and assigns it to the Scope field.
func (o *NtfSubscriptionControlSingleAllOfAttributes) SetScope(v Scope) {
	o.Scope = &v
}

// GetNotificationFilter returns the NotificationFilter field value if set, zero value otherwise.
func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationFilter() string {
	if o == nil || IsNil(o.NotificationFilter) {
		var ret string
		return ret
	}
	return *o.NotificationFilter
}

// GetNotificationFilterOk returns a tuple with the NotificationFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NtfSubscriptionControlSingleAllOfAttributes) GetNotificationFilterOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationFilter) {
		return nil, false
	}
	return o.NotificationFilter, true
}

// HasNotificationFilter returns a boolean if a field has been set.
func (o *NtfSubscriptionControlSingleAllOfAttributes) HasNotificationFilter() bool {
	if o != nil && !IsNil(o.NotificationFilter) {
		return true
	}

	return false
}

// SetNotificationFilter gets a reference to the given string and assigns it to the NotificationFilter field.
func (o *NtfSubscriptionControlSingleAllOfAttributes) SetNotificationFilter(v string) {
	o.NotificationFilter = &v
}

func (o NtfSubscriptionControlSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NtfSubscriptionControlSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationRecipientAddress) {
		toSerialize["notificationRecipientAddress"] = o.NotificationRecipientAddress
	}
	if !IsNil(o.NotificationTypes) {
		toSerialize["notificationTypes"] = o.NotificationTypes
	}
	if !IsNil(o.Scope) {
		toSerialize["scope"] = o.Scope
	}
	if !IsNil(o.NotificationFilter) {
		toSerialize["notificationFilter"] = o.NotificationFilter
	}
	return toSerialize, nil
}

type NullableNtfSubscriptionControlSingleAllOfAttributes struct {
	value *NtfSubscriptionControlSingleAllOfAttributes
	isSet bool
}

func (v NullableNtfSubscriptionControlSingleAllOfAttributes) Get() *NtfSubscriptionControlSingleAllOfAttributes {
	return v.value
}

func (v *NullableNtfSubscriptionControlSingleAllOfAttributes) Set(val *NtfSubscriptionControlSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNtfSubscriptionControlSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNtfSubscriptionControlSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNtfSubscriptionControlSingleAllOfAttributes(val *NtfSubscriptionControlSingleAllOfAttributes) *NullableNtfSubscriptionControlSingleAllOfAttributes {
	return &NullableNtfSubscriptionControlSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableNtfSubscriptionControlSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNtfSubscriptionControlSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


