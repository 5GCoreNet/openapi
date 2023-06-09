/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
)

// checks if the Spt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Spt{}

// Spt Contains the data of a Service Point Trigger
type Spt struct {
	ConditionNegated   bool               `json:"conditionNegated"`
	SptGroup           []int32            `json:"sptGroup"`
	RegType            []RegistrationType `json:"regType,omitempty"`
	RequestUri         *string            `json:"requestUri,omitempty"`
	SipMethod          *string            `json:"sipMethod,omitempty"`
	SipHeader          *HeaderSipRequest  `json:"sipHeader,omitempty"`
	SessionCase        *RequestDirection  `json:"sessionCase,omitempty"`
	SessionDescription *SdpDescription    `json:"sessionDescription,omitempty"`
}

// NewSpt instantiates a new Spt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpt(conditionNegated bool, sptGroup []int32) *Spt {
	this := Spt{}
	this.ConditionNegated = conditionNegated
	this.SptGroup = sptGroup
	return &this
}

// NewSptWithDefaults instantiates a new Spt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSptWithDefaults() *Spt {
	this := Spt{}
	return &this
}

// GetConditionNegated returns the ConditionNegated field value
func (o *Spt) GetConditionNegated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ConditionNegated
}

// GetConditionNegatedOk returns a tuple with the ConditionNegated field value
// and a boolean to check if the value has been set.
func (o *Spt) GetConditionNegatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionNegated, true
}

// SetConditionNegated sets field value
func (o *Spt) SetConditionNegated(v bool) {
	o.ConditionNegated = v
}

// GetSptGroup returns the SptGroup field value
func (o *Spt) GetSptGroup() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.SptGroup
}

// GetSptGroupOk returns a tuple with the SptGroup field value
// and a boolean to check if the value has been set.
func (o *Spt) GetSptGroupOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SptGroup, true
}

// SetSptGroup sets field value
func (o *Spt) SetSptGroup(v []int32) {
	o.SptGroup = v
}

// GetRegType returns the RegType field value if set, zero value otherwise.
func (o *Spt) GetRegType() []RegistrationType {
	if o == nil || IsNil(o.RegType) {
		var ret []RegistrationType
		return ret
	}
	return o.RegType
}

// GetRegTypeOk returns a tuple with the RegType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spt) GetRegTypeOk() ([]RegistrationType, bool) {
	if o == nil || IsNil(o.RegType) {
		return nil, false
	}
	return o.RegType, true
}

// HasRegType returns a boolean if a field has been set.
func (o *Spt) HasRegType() bool {
	if o != nil && !IsNil(o.RegType) {
		return true
	}

	return false
}

// SetRegType gets a reference to the given []RegistrationType and assigns it to the RegType field.
func (o *Spt) SetRegType(v []RegistrationType) {
	o.RegType = v
}

// GetRequestUri returns the RequestUri field value if set, zero value otherwise.
func (o *Spt) GetRequestUri() string {
	if o == nil || IsNil(o.RequestUri) {
		var ret string
		return ret
	}
	return *o.RequestUri
}

// GetRequestUriOk returns a tuple with the RequestUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spt) GetRequestUriOk() (*string, bool) {
	if o == nil || IsNil(o.RequestUri) {
		return nil, false
	}
	return o.RequestUri, true
}

// HasRequestUri returns a boolean if a field has been set.
func (o *Spt) HasRequestUri() bool {
	if o != nil && !IsNil(o.RequestUri) {
		return true
	}

	return false
}

// SetRequestUri gets a reference to the given string and assigns it to the RequestUri field.
func (o *Spt) SetRequestUri(v string) {
	o.RequestUri = &v
}

// GetSipMethod returns the SipMethod field value if set, zero value otherwise.
func (o *Spt) GetSipMethod() string {
	if o == nil || IsNil(o.SipMethod) {
		var ret string
		return ret
	}
	return *o.SipMethod
}

// GetSipMethodOk returns a tuple with the SipMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spt) GetSipMethodOk() (*string, bool) {
	if o == nil || IsNil(o.SipMethod) {
		return nil, false
	}
	return o.SipMethod, true
}

// HasSipMethod returns a boolean if a field has been set.
func (o *Spt) HasSipMethod() bool {
	if o != nil && !IsNil(o.SipMethod) {
		return true
	}

	return false
}

// SetSipMethod gets a reference to the given string and assigns it to the SipMethod field.
func (o *Spt) SetSipMethod(v string) {
	o.SipMethod = &v
}

// GetSipHeader returns the SipHeader field value if set, zero value otherwise.
func (o *Spt) GetSipHeader() HeaderSipRequest {
	if o == nil || IsNil(o.SipHeader) {
		var ret HeaderSipRequest
		return ret
	}
	return *o.SipHeader
}

// GetSipHeaderOk returns a tuple with the SipHeader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spt) GetSipHeaderOk() (*HeaderSipRequest, bool) {
	if o == nil || IsNil(o.SipHeader) {
		return nil, false
	}
	return o.SipHeader, true
}

// HasSipHeader returns a boolean if a field has been set.
func (o *Spt) HasSipHeader() bool {
	if o != nil && !IsNil(o.SipHeader) {
		return true
	}

	return false
}

// SetSipHeader gets a reference to the given HeaderSipRequest and assigns it to the SipHeader field.
func (o *Spt) SetSipHeader(v HeaderSipRequest) {
	o.SipHeader = &v
}

// GetSessionCase returns the SessionCase field value if set, zero value otherwise.
func (o *Spt) GetSessionCase() RequestDirection {
	if o == nil || IsNil(o.SessionCase) {
		var ret RequestDirection
		return ret
	}
	return *o.SessionCase
}

// GetSessionCaseOk returns a tuple with the SessionCase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spt) GetSessionCaseOk() (*RequestDirection, bool) {
	if o == nil || IsNil(o.SessionCase) {
		return nil, false
	}
	return o.SessionCase, true
}

// HasSessionCase returns a boolean if a field has been set.
func (o *Spt) HasSessionCase() bool {
	if o != nil && !IsNil(o.SessionCase) {
		return true
	}

	return false
}

// SetSessionCase gets a reference to the given RequestDirection and assigns it to the SessionCase field.
func (o *Spt) SetSessionCase(v RequestDirection) {
	o.SessionCase = &v
}

// GetSessionDescription returns the SessionDescription field value if set, zero value otherwise.
func (o *Spt) GetSessionDescription() SdpDescription {
	if o == nil || IsNil(o.SessionDescription) {
		var ret SdpDescription
		return ret
	}
	return *o.SessionDescription
}

// GetSessionDescriptionOk returns a tuple with the SessionDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Spt) GetSessionDescriptionOk() (*SdpDescription, bool) {
	if o == nil || IsNil(o.SessionDescription) {
		return nil, false
	}
	return o.SessionDescription, true
}

// HasSessionDescription returns a boolean if a field has been set.
func (o *Spt) HasSessionDescription() bool {
	if o != nil && !IsNil(o.SessionDescription) {
		return true
	}

	return false
}

// SetSessionDescription gets a reference to the given SdpDescription and assigns it to the SessionDescription field.
func (o *Spt) SetSessionDescription(v SdpDescription) {
	o.SessionDescription = &v
}

func (o Spt) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Spt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditionNegated"] = o.ConditionNegated
	toSerialize["sptGroup"] = o.SptGroup
	if !IsNil(o.RegType) {
		toSerialize["regType"] = o.RegType
	}
	if !IsNil(o.RequestUri) {
		toSerialize["requestUri"] = o.RequestUri
	}
	if !IsNil(o.SipMethod) {
		toSerialize["sipMethod"] = o.SipMethod
	}
	if !IsNil(o.SipHeader) {
		toSerialize["sipHeader"] = o.SipHeader
	}
	if !IsNil(o.SessionCase) {
		toSerialize["sessionCase"] = o.SessionCase
	}
	if !IsNil(o.SessionDescription) {
		toSerialize["sessionDescription"] = o.SessionDescription
	}
	return toSerialize, nil
}

type NullableSpt struct {
	value *Spt
	isSet bool
}

func (v NullableSpt) Get() *Spt {
	return v.value
}

func (v *NullableSpt) Set(val *Spt) {
	v.value = val
	v.isSet = true
}

func (v NullableSpt) IsSet() bool {
	return v.isSet
}

func (v *NullableSpt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpt(val *Spt) *NullableSpt {
	return &NullableSpt{value: val, isSet: true}
}

func (v NullableSpt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
