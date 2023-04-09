/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the DefaultNotificationSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefaultNotificationSubscription{}

// DefaultNotificationSubscription Data structure for specifying the notifications the NF service subscribes by default, along with callback URI
type DefaultNotificationSubscription struct {
	NotificationType NotificationType `json:"notificationType"`
	// String providing an URI formatted according to RFC 3986.
	CallbackUri        string              `json:"callbackUri"`
	N1MessageClass     *N1MessageClass     `json:"n1MessageClass,omitempty"`
	N2InformationClass *N2InformationClass `json:"n2InformationClass,omitempty"`
	Versions           []string            `json:"versions,omitempty"`
	Binding            *string             `json:"binding,omitempty"`
	AcceptedEncoding   *string             `json:"acceptedEncoding,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// A map of service specific information. The name of the corresponding service (as specified in ServiceName data type) is the key.
	ServiceInfoList *map[string]DefSubServiceInfo `json:"serviceInfoList,omitempty"`
}

// NewDefaultNotificationSubscription instantiates a new DefaultNotificationSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefaultNotificationSubscription(notificationType NotificationType, callbackUri string) *DefaultNotificationSubscription {
	this := DefaultNotificationSubscription{}
	this.NotificationType = notificationType
	this.CallbackUri = callbackUri
	return &this
}

// NewDefaultNotificationSubscriptionWithDefaults instantiates a new DefaultNotificationSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefaultNotificationSubscriptionWithDefaults() *DefaultNotificationSubscription {
	this := DefaultNotificationSubscription{}
	return &this
}

// GetNotificationType returns the NotificationType field value
func (o *DefaultNotificationSubscription) GetNotificationType() NotificationType {
	if o == nil {
		var ret NotificationType
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetNotificationTypeOk() (*NotificationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *DefaultNotificationSubscription) SetNotificationType(v NotificationType) {
	o.NotificationType = v
}

// GetCallbackUri returns the CallbackUri field value
func (o *DefaultNotificationSubscription) GetCallbackUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CallbackUri
}

// GetCallbackUriOk returns a tuple with the CallbackUri field value
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetCallbackUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CallbackUri, true
}

// SetCallbackUri sets field value
func (o *DefaultNotificationSubscription) SetCallbackUri(v string) {
	o.CallbackUri = v
}

// GetN1MessageClass returns the N1MessageClass field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetN1MessageClass() N1MessageClass {
	if o == nil || IsNil(o.N1MessageClass) {
		var ret N1MessageClass
		return ret
	}
	return *o.N1MessageClass
}

// GetN1MessageClassOk returns a tuple with the N1MessageClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetN1MessageClassOk() (*N1MessageClass, bool) {
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

// SetN1MessageClass gets a reference to the given N1MessageClass and assigns it to the N1MessageClass field.
func (o *DefaultNotificationSubscription) SetN1MessageClass(v N1MessageClass) {
	o.N1MessageClass = &v
}

// GetN2InformationClass returns the N2InformationClass field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetN2InformationClass() N2InformationClass {
	if o == nil || IsNil(o.N2InformationClass) {
		var ret N2InformationClass
		return ret
	}
	return *o.N2InformationClass
}

// GetN2InformationClassOk returns a tuple with the N2InformationClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetN2InformationClassOk() (*N2InformationClass, bool) {
	if o == nil || IsNil(o.N2InformationClass) {
		return nil, false
	}
	return o.N2InformationClass, true
}

// HasN2InformationClass returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasN2InformationClass() bool {
	if o != nil && !IsNil(o.N2InformationClass) {
		return true
	}

	return false
}

// SetN2InformationClass gets a reference to the given N2InformationClass and assigns it to the N2InformationClass field.
func (o *DefaultNotificationSubscription) SetN2InformationClass(v N2InformationClass) {
	o.N2InformationClass = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetVersions() []string {
	if o == nil || IsNil(o.Versions) {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetVersionsOk() ([]string, bool) {
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

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *DefaultNotificationSubscription) SetVersions(v []string) {
	o.Versions = v
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

// GetAcceptedEncoding returns the AcceptedEncoding field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetAcceptedEncoding() string {
	if o == nil || IsNil(o.AcceptedEncoding) {
		var ret string
		return ret
	}
	return *o.AcceptedEncoding
}

// GetAcceptedEncodingOk returns a tuple with the AcceptedEncoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetAcceptedEncodingOk() (*string, bool) {
	if o == nil || IsNil(o.AcceptedEncoding) {
		return nil, false
	}
	return o.AcceptedEncoding, true
}

// HasAcceptedEncoding returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasAcceptedEncoding() bool {
	if o != nil && !IsNil(o.AcceptedEncoding) {
		return true
	}

	return false
}

// SetAcceptedEncoding gets a reference to the given string and assigns it to the AcceptedEncoding field.
func (o *DefaultNotificationSubscription) SetAcceptedEncoding(v string) {
	o.AcceptedEncoding = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *DefaultNotificationSubscription) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetServiceInfoList returns the ServiceInfoList field value if set, zero value otherwise.
func (o *DefaultNotificationSubscription) GetServiceInfoList() map[string]DefSubServiceInfo {
	if o == nil || IsNil(o.ServiceInfoList) {
		var ret map[string]DefSubServiceInfo
		return ret
	}
	return *o.ServiceInfoList
}

// GetServiceInfoListOk returns a tuple with the ServiceInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefaultNotificationSubscription) GetServiceInfoListOk() (*map[string]DefSubServiceInfo, bool) {
	if o == nil || IsNil(o.ServiceInfoList) {
		return nil, false
	}
	return o.ServiceInfoList, true
}

// HasServiceInfoList returns a boolean if a field has been set.
func (o *DefaultNotificationSubscription) HasServiceInfoList() bool {
	if o != nil && !IsNil(o.ServiceInfoList) {
		return true
	}

	return false
}

// SetServiceInfoList gets a reference to the given map[string]DefSubServiceInfo and assigns it to the ServiceInfoList field.
func (o *DefaultNotificationSubscription) SetServiceInfoList(v map[string]DefSubServiceInfo) {
	o.ServiceInfoList = &v
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
	toSerialize["notificationType"] = o.NotificationType
	toSerialize["callbackUri"] = o.CallbackUri
	if !IsNil(o.N1MessageClass) {
		toSerialize["n1MessageClass"] = o.N1MessageClass
	}
	if !IsNil(o.N2InformationClass) {
		toSerialize["n2InformationClass"] = o.N2InformationClass
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.Binding) {
		toSerialize["binding"] = o.Binding
	}
	if !IsNil(o.AcceptedEncoding) {
		toSerialize["acceptedEncoding"] = o.AcceptedEncoding
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.ServiceInfoList) {
		toSerialize["serviceInfoList"] = o.ServiceInfoList
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
