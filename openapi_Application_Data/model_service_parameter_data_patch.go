/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
)

// checks if the ServiceParameterDataPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceParameterDataPatch{}

// ServiceParameterDataPatch Represents the service parameter data that can be updated.
type ServiceParameterDataPatch struct {
	// Represents configuration parameters for V2X communications over PC5 reference point.
	ParamOverPc5 *string `json:"paramOverPc5,omitempty"`
	// Represents configuration parameters for V2X communications over Uu reference point.
	ParamOverUu *string `json:"paramOverUu,omitempty"`
	// Represents the service parameters for 5G ProSe direct discovery.
	ParamForProSeDd *string `json:"paramForProSeDd,omitempty"`
	// Represents the service parameters for 5G ProSe direct communications.
	ParamForProSeDc *string `json:"paramForProSeDc,omitempty"`
	// Represents the service parameters for 5G ProSe UE-to-network relay UE.
	ParamForProSeU2NRelUe *string `json:"paramForProSeU2NRelUe,omitempty"`
	// Represents the service parameters for 5G ProSe Remate UE.
	ParamForProSeRemUe *string `json:"paramForProSeRemUe,omitempty"`
	// Contains the service parameter used to influence the URSP.
	UrspInfluence []UrspRuleRequest `json:"urspInfluence,omitempty"`
	// Contains the outcome of the UE Policy Delivery.
	DeliveryEvents []Event `json:"deliveryEvents,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	PolicDelivNotifUri *string `json:"policDelivNotifUri,omitempty"`
}

// NewServiceParameterDataPatch instantiates a new ServiceParameterDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceParameterDataPatch() *ServiceParameterDataPatch {
	this := ServiceParameterDataPatch{}
	return &this
}

// NewServiceParameterDataPatchWithDefaults instantiates a new ServiceParameterDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceParameterDataPatchWithDefaults() *ServiceParameterDataPatch {
	this := ServiceParameterDataPatch{}
	return &this
}

// GetParamOverPc5 returns the ParamOverPc5 field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetParamOverPc5() string {
	if o == nil || IsNil(o.ParamOverPc5) {
		var ret string
		return ret
	}
	return *o.ParamOverPc5
}

// GetParamOverPc5Ok returns a tuple with the ParamOverPc5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetParamOverPc5Ok() (*string, bool) {
	if o == nil || IsNil(o.ParamOverPc5) {
		return nil, false
	}
	return o.ParamOverPc5, true
}

// HasParamOverPc5 returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasParamOverPc5() bool {
	if o != nil && !IsNil(o.ParamOverPc5) {
		return true
	}

	return false
}

// SetParamOverPc5 gets a reference to the given string and assigns it to the ParamOverPc5 field.
func (o *ServiceParameterDataPatch) SetParamOverPc5(v string) {
	o.ParamOverPc5 = &v
}

// GetParamOverUu returns the ParamOverUu field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetParamOverUu() string {
	if o == nil || IsNil(o.ParamOverUu) {
		var ret string
		return ret
	}
	return *o.ParamOverUu
}

// GetParamOverUuOk returns a tuple with the ParamOverUu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetParamOverUuOk() (*string, bool) {
	if o == nil || IsNil(o.ParamOverUu) {
		return nil, false
	}
	return o.ParamOverUu, true
}

// HasParamOverUu returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasParamOverUu() bool {
	if o != nil && !IsNil(o.ParamOverUu) {
		return true
	}

	return false
}

// SetParamOverUu gets a reference to the given string and assigns it to the ParamOverUu field.
func (o *ServiceParameterDataPatch) SetParamOverUu(v string) {
	o.ParamOverUu = &v
}

// GetParamForProSeDd returns the ParamForProSeDd field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetParamForProSeDd() string {
	if o == nil || IsNil(o.ParamForProSeDd) {
		var ret string
		return ret
	}
	return *o.ParamForProSeDd
}

// GetParamForProSeDdOk returns a tuple with the ParamForProSeDd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetParamForProSeDdOk() (*string, bool) {
	if o == nil || IsNil(o.ParamForProSeDd) {
		return nil, false
	}
	return o.ParamForProSeDd, true
}

// HasParamForProSeDd returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasParamForProSeDd() bool {
	if o != nil && !IsNil(o.ParamForProSeDd) {
		return true
	}

	return false
}

// SetParamForProSeDd gets a reference to the given string and assigns it to the ParamForProSeDd field.
func (o *ServiceParameterDataPatch) SetParamForProSeDd(v string) {
	o.ParamForProSeDd = &v
}

// GetParamForProSeDc returns the ParamForProSeDc field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetParamForProSeDc() string {
	if o == nil || IsNil(o.ParamForProSeDc) {
		var ret string
		return ret
	}
	return *o.ParamForProSeDc
}

// GetParamForProSeDcOk returns a tuple with the ParamForProSeDc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetParamForProSeDcOk() (*string, bool) {
	if o == nil || IsNil(o.ParamForProSeDc) {
		return nil, false
	}
	return o.ParamForProSeDc, true
}

// HasParamForProSeDc returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasParamForProSeDc() bool {
	if o != nil && !IsNil(o.ParamForProSeDc) {
		return true
	}

	return false
}

// SetParamForProSeDc gets a reference to the given string and assigns it to the ParamForProSeDc field.
func (o *ServiceParameterDataPatch) SetParamForProSeDc(v string) {
	o.ParamForProSeDc = &v
}

// GetParamForProSeU2NRelUe returns the ParamForProSeU2NRelUe field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetParamForProSeU2NRelUe() string {
	if o == nil || IsNil(o.ParamForProSeU2NRelUe) {
		var ret string
		return ret
	}
	return *o.ParamForProSeU2NRelUe
}

// GetParamForProSeU2NRelUeOk returns a tuple with the ParamForProSeU2NRelUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetParamForProSeU2NRelUeOk() (*string, bool) {
	if o == nil || IsNil(o.ParamForProSeU2NRelUe) {
		return nil, false
	}
	return o.ParamForProSeU2NRelUe, true
}

// HasParamForProSeU2NRelUe returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasParamForProSeU2NRelUe() bool {
	if o != nil && !IsNil(o.ParamForProSeU2NRelUe) {
		return true
	}

	return false
}

// SetParamForProSeU2NRelUe gets a reference to the given string and assigns it to the ParamForProSeU2NRelUe field.
func (o *ServiceParameterDataPatch) SetParamForProSeU2NRelUe(v string) {
	o.ParamForProSeU2NRelUe = &v
}

// GetParamForProSeRemUe returns the ParamForProSeRemUe field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetParamForProSeRemUe() string {
	if o == nil || IsNil(o.ParamForProSeRemUe) {
		var ret string
		return ret
	}
	return *o.ParamForProSeRemUe
}

// GetParamForProSeRemUeOk returns a tuple with the ParamForProSeRemUe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetParamForProSeRemUeOk() (*string, bool) {
	if o == nil || IsNil(o.ParamForProSeRemUe) {
		return nil, false
	}
	return o.ParamForProSeRemUe, true
}

// HasParamForProSeRemUe returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasParamForProSeRemUe() bool {
	if o != nil && !IsNil(o.ParamForProSeRemUe) {
		return true
	}

	return false
}

// SetParamForProSeRemUe gets a reference to the given string and assigns it to the ParamForProSeRemUe field.
func (o *ServiceParameterDataPatch) SetParamForProSeRemUe(v string) {
	o.ParamForProSeRemUe = &v
}

// GetUrspInfluence returns the UrspInfluence field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetUrspInfluence() []UrspRuleRequest {
	if o == nil || IsNil(o.UrspInfluence) {
		var ret []UrspRuleRequest
		return ret
	}
	return o.UrspInfluence
}

// GetUrspInfluenceOk returns a tuple with the UrspInfluence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetUrspInfluenceOk() ([]UrspRuleRequest, bool) {
	if o == nil || IsNil(o.UrspInfluence) {
		return nil, false
	}
	return o.UrspInfluence, true
}

// HasUrspInfluence returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasUrspInfluence() bool {
	if o != nil && !IsNil(o.UrspInfluence) {
		return true
	}

	return false
}

// SetUrspInfluence gets a reference to the given []UrspRuleRequest and assigns it to the UrspInfluence field.
func (o *ServiceParameterDataPatch) SetUrspInfluence(v []UrspRuleRequest) {
	o.UrspInfluence = v
}

// GetDeliveryEvents returns the DeliveryEvents field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetDeliveryEvents() []Event {
	if o == nil || IsNil(o.DeliveryEvents) {
		var ret []Event
		return ret
	}
	return o.DeliveryEvents
}

// GetDeliveryEventsOk returns a tuple with the DeliveryEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetDeliveryEventsOk() ([]Event, bool) {
	if o == nil || IsNil(o.DeliveryEvents) {
		return nil, false
	}
	return o.DeliveryEvents, true
}

// HasDeliveryEvents returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasDeliveryEvents() bool {
	if o != nil && !IsNil(o.DeliveryEvents) {
		return true
	}

	return false
}

// SetDeliveryEvents gets a reference to the given []Event and assigns it to the DeliveryEvents field.
func (o *ServiceParameterDataPatch) SetDeliveryEvents(v []Event) {
	o.DeliveryEvents = v
}

// GetPolicDelivNotifUri returns the PolicDelivNotifUri field value if set, zero value otherwise.
func (o *ServiceParameterDataPatch) GetPolicDelivNotifUri() string {
	if o == nil || IsNil(o.PolicDelivNotifUri) {
		var ret string
		return ret
	}
	return *o.PolicDelivNotifUri
}

// GetPolicDelivNotifUriOk returns a tuple with the PolicDelivNotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceParameterDataPatch) GetPolicDelivNotifUriOk() (*string, bool) {
	if o == nil || IsNil(o.PolicDelivNotifUri) {
		return nil, false
	}
	return o.PolicDelivNotifUri, true
}

// HasPolicDelivNotifUri returns a boolean if a field has been set.
func (o *ServiceParameterDataPatch) HasPolicDelivNotifUri() bool {
	if o != nil && !IsNil(o.PolicDelivNotifUri) {
		return true
	}

	return false
}

// SetPolicDelivNotifUri gets a reference to the given string and assigns it to the PolicDelivNotifUri field.
func (o *ServiceParameterDataPatch) SetPolicDelivNotifUri(v string) {
	o.PolicDelivNotifUri = &v
}

func (o ServiceParameterDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceParameterDataPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParamOverPc5) {
		toSerialize["paramOverPc5"] = o.ParamOverPc5
	}
	if !IsNil(o.ParamOverUu) {
		toSerialize["paramOverUu"] = o.ParamOverUu
	}
	if !IsNil(o.ParamForProSeDd) {
		toSerialize["paramForProSeDd"] = o.ParamForProSeDd
	}
	if !IsNil(o.ParamForProSeDc) {
		toSerialize["paramForProSeDc"] = o.ParamForProSeDc
	}
	if !IsNil(o.ParamForProSeU2NRelUe) {
		toSerialize["paramForProSeU2NRelUe"] = o.ParamForProSeU2NRelUe
	}
	if !IsNil(o.ParamForProSeRemUe) {
		toSerialize["paramForProSeRemUe"] = o.ParamForProSeRemUe
	}
	if !IsNil(o.UrspInfluence) {
		toSerialize["urspInfluence"] = o.UrspInfluence
	}
	if !IsNil(o.DeliveryEvents) {
		toSerialize["deliveryEvents"] = o.DeliveryEvents
	}
	if !IsNil(o.PolicDelivNotifUri) {
		toSerialize["policDelivNotifUri"] = o.PolicDelivNotifUri
	}
	return toSerialize, nil
}

type NullableServiceParameterDataPatch struct {
	value *ServiceParameterDataPatch
	isSet bool
}

func (v NullableServiceParameterDataPatch) Get() *ServiceParameterDataPatch {
	return v.value
}

func (v *NullableServiceParameterDataPatch) Set(val *ServiceParameterDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceParameterDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceParameterDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceParameterDataPatch(val *ServiceParameterDataPatch) *NullableServiceParameterDataPatch {
	return &NullableServiceParameterDataPatch{value: val, isSet: true}
}

func (v NullableServiceParameterDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceParameterDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
