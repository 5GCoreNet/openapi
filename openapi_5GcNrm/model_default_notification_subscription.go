/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the DefaultNotificationSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultNotificationSubscription{}

// DefaultNotificationSubscription struct for DefaultNotificationSubscription
type DefaultNotificationSubscription struct {
	NotificationType   *NotificationType `json:"notificationType,omitempty"`
	CallbackURI        *string           `json:"callbackURI,omitempty"`
	N1MessageClass     *bool             `json:"n1MessageClass,omitempty"`
	N2InfroamtionClass *bool             `json:"n2InfroamtionClass,omitempty"`
	Versions           *string           `json:"versions,omitempty"`
	Binding            *string           `json:"binding,omitempty"`
}

// NewDefaultNotificationSubscription instantiates a new DefaultNotificationSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultNotificationSubscription() *DefaultNotificationSubscription {
	this := DefaultNotificationSubscription{}
	return &this
}

// NewDefaultNotificationSubscriptionWithDefaults instantiates a new DefaultNotificationSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultNotificationSubscriptionWithDefaults() *DefaultNotificationSubscription {
	this := DefaultNotificationSubscription{}
	return &this
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetNotificationType() NotificationType {
	if o == nil || IsNil(o.NotificationType) {
		var ret NotificationType
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil || IsNil(o.NotificationType) {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasNotificationType() bool {
	if o != nil && !IsNil(o.NotificationType) {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given NotificationType and assigns it to the NotificationType field.
func (o *DefaultNotificationSubscription) SetNotificationType(v NotificationType) {
	o.NotificationType = &v
}

// GetCallbackURI returns the CallbackURI field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetCallbackURI() string {
	if o == nil || IsNil(o.CallbackURI) {
		var ret string
		return ret
	}
	return *o.CallbackURI
}

// GetCallbackURIOk returns a tuple with the CallbackURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetCallbackURIOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackURI) {
		return nil, false
	}
	return o.CallbackURI, true
}

// HasCallbackURI returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasCallbackURI() bool {
	if o != nil && !IsNil(o.CallbackURI) {
		return true
	}

	return false
}

// SetCallbackURI gets a reference to the given string and assigns it to the CallbackURI field.
func (o *DefaultNotificationSubscription) SetCallbackURI(v string) {
	o.CallbackURI = &v
}

// GetN1MessageClass returns the N1MessageClass field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetN1MessageClass() bool {
	if o == nil || IsNil(o.N1MessageClass) {
		var ret bool
		return ret
	}
	return *o.N1MessageClass
}

// GetN1MessageClassOk returns a tuple with the N1MessageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetN1MessageClassOk() (*bool, bool) {
	if o == nil || IsNil(o.N1MessageClass) {
		return nil, false
	}
	return o.N1MessageClass, true
}

// HasN1MessageClass returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasN1MessageClass() bool {
	if o != nil && !IsNil(o.N1MessageClass) {
		return true
	}

	return false
}

// SetN1MessageClass gets a reference to the given bool and assigns it to the N1MessageClass field.
func (o *DefaultNotificationSubscription) SetN1MessageClass(v bool) {
	o.N1MessageClass = &v
}

// GetN2InfroamtionClass returns the N2InfroamtionClass field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetN2InfroamtionClass() bool {
	if o == nil || IsNil(o.N2InfroamtionClass) {
		var ret bool
		return ret
	}
	return *o.N2InfroamtionClass
}

// GetN2InfroamtionClassOk returns a tuple with the N2InfroamtionClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetN2InfroamtionClassOk() (*bool, bool) {
	if o == nil || IsNil(o.N2InfroamtionClass) {
		return nil, false
	}
	return o.N2InfroamtionClass, true
}

// HasN2InfroamtionClass returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasN2InfroamtionClass() bool {
	if o != nil && !IsNil(o.N2InfroamtionClass) {
		return true
	}

	return false
}

// SetN2InfroamtionClass gets a reference to the given bool and assigns it to the N2InfroamtionClass field.
func (o *DefaultNotificationSubscription) SetN2InfroamtionClass(v bool) {
	o.N2InfroamtionClass = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetVersions() string {
	if o == nil || IsNil(o.Versions) {
		var ret string
		return ret
	}
	return *o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetVersionsOk() (*string, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given string and assigns it to the Versions field.
func (o *DefaultNotificationSubscription) SetVersions(v string) {
	o.Versions = &v
}

// GetBinding returns the Binding field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetBinding() string {
	if o == nil || IsNil(o.Binding) {
		var ret string
		return ret
	}
	return *o.Binding
}

// GetBindingOk returns a tuple with the Binding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetBindingOk() (*string, bool) {
	if o == nil || IsNil(o.Binding) {
		return nil, false
	}
	return o.Binding, true
}

// HasBinding returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasBinding() bool {
	if o != nil && !IsNil(o.Binding) {
		return true
	}

	return false
}

// SetBinding gets a reference to the given string and assigns it to the Binding field.
func (o *DefaultNotificationSubscription) SetBinding(v string) {
	o.Binding = &v
}

func (o DefaultNotificationSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefaultNotificationSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationType) {
		toSerialize["notificationType"] = o.NotificationType
	}
	if !IsNil(o.CallbackURI) {
		toSerialize["callbackURI"] = o.CallbackURI
	}
	if !IsNil(o.N1MessageClass) {
		toSerialize["n1MessageClass"] = o.N1MessageClass
	}
	if !IsNil(o.N2InfroamtionClass) {
		toSerialize["n2InfroamtionClass"] = o.N2InfroamtionClass
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.Binding) {
		toSerialize["binding"] = o.Binding
	}
	return toSerialize, nil
}

type NullableDefaultNotificationSubscription struct {
	value *DefaultNotificationSubscription
	isSet bool
}

func (v NullableDefaultNotificationSubscription) Get() *DefaultNotificationSubscription {
	return v.value
}

func (v *NullableDefaultNotificationSubscription) Set(val *DefaultNotificationSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableDefaultNotificationSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableDefaultNotificationSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefaultNotificationSubscription(val *DefaultNotificationSubscription) *NullableDefaultNotificationSubscription {
	return &NullableDefaultNotificationSubscription{value: val, isSet: true}
}

func (v NullableDefaultNotificationSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefaultNotificationSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
