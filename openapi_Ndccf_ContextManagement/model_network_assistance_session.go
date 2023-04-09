/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the NetworkAssistanceSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkAssistanceSession{}

// NetworkAssistanceSession A representation of a Network Assistance Session resource.
type NetworkAssistanceSession struct {
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	NaSessionId                string                       `json:"naSessionId"`
	ServiceDataFlowDescription []ServiceDataFlowDescription `json:"serviceDataFlowDescription,omitempty"`
	// String chosen by the 5GMS AF to serve as an identifier in a resource URI.
	PolicyTemplateId *string             `json:"policyTemplateId,omitempty"`
	RequestedQoS     *M5QoSSpecification `json:"requestedQoS,omitempty"`
	RecommendedQoS   *M5QoSSpecification `json:"recommendedQoS,omitempty"`
	// Uniform Resource Locator, comforming with the URI Generic Syntax specified in IETF RFC 3986.
	NotficationURL *string `json:"notficationURL,omitempty"`
}

// NewNetworkAssistanceSession instantiates a new NetworkAssistanceSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkAssistanceSession(naSessionId string) *NetworkAssistanceSession {
	this := NetworkAssistanceSession{}
	this.NaSessionId = naSessionId
	return &this
}

// NewNetworkAssistanceSessionWithDefaults instantiates a new NetworkAssistanceSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkAssistanceSessionWithDefaults() *NetworkAssistanceSession {
	this := NetworkAssistanceSession{}
	return &this
}

// GetNaSessionId returns the NaSessionId field value
func (o *NetworkAssistanceSession) GetNaSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NaSessionId
}

// GetNaSessionIdOk returns a tuple with the NaSessionId field value
// and a boolean to check if the value has been set.
func (o *NetworkAssistanceSession) GetNaSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NaSessionId, true
}

// SetNaSessionId sets field value
func (o *NetworkAssistanceSession) SetNaSessionId(v string) {
	o.NaSessionId = v
}

// GetServiceDataFlowDescription returns the ServiceDataFlowDescription field value if set, zero value otherwise.
func (o *NetworkAssistanceSession) GetServiceDataFlowDescription() []ServiceDataFlowDescription {
	if o == nil || IsNil(o.ServiceDataFlowDescription) {
		var ret []ServiceDataFlowDescription
		return ret
	}
	return o.ServiceDataFlowDescription
}

// GetServiceDataFlowDescriptionOk returns a tuple with the ServiceDataFlowDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAssistanceSession) GetServiceDataFlowDescriptionOk() ([]ServiceDataFlowDescription, bool) {
	if o == nil || IsNil(o.ServiceDataFlowDescription) {
		return nil, false
	}
	return o.ServiceDataFlowDescription, true
}

// HasServiceDataFlowDescription returns a boolean if a field has been set.
func (o *NetworkAssistanceSession) HasServiceDataFlowDescription() bool {
	if o != nil && !IsNil(o.ServiceDataFlowDescription) {
		return true
	}

	return false
}

// SetServiceDataFlowDescription gets a reference to the given []ServiceDataFlowDescription and assigns it to the ServiceDataFlowDescription field.
func (o *NetworkAssistanceSession) SetServiceDataFlowDescription(v []ServiceDataFlowDescription) {
	o.ServiceDataFlowDescription = v
}

// GetPolicyTemplateId returns the PolicyTemplateId field value if set, zero value otherwise.
func (o *NetworkAssistanceSession) GetPolicyTemplateId() string {
	if o == nil || IsNil(o.PolicyTemplateId) {
		var ret string
		return ret
	}
	return *o.PolicyTemplateId
}

// GetPolicyTemplateIdOk returns a tuple with the PolicyTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAssistanceSession) GetPolicyTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.PolicyTemplateId) {
		return nil, false
	}
	return o.PolicyTemplateId, true
}

// HasPolicyTemplateId returns a boolean if a field has been set.
func (o *NetworkAssistanceSession) HasPolicyTemplateId() bool {
	if o != nil && !IsNil(o.PolicyTemplateId) {
		return true
	}

	return false
}

// SetPolicyTemplateId gets a reference to the given string and assigns it to the PolicyTemplateId field.
func (o *NetworkAssistanceSession) SetPolicyTemplateId(v string) {
	o.PolicyTemplateId = &v
}

// GetRequestedQoS returns the RequestedQoS field value if set, zero value otherwise.
func (o *NetworkAssistanceSession) GetRequestedQoS() M5QoSSpecification {
	if o == nil || IsNil(o.RequestedQoS) {
		var ret M5QoSSpecification
		return ret
	}
	return *o.RequestedQoS
}

// GetRequestedQoSOk returns a tuple with the RequestedQoS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAssistanceSession) GetRequestedQoSOk() (*M5QoSSpecification, bool) {
	if o == nil || IsNil(o.RequestedQoS) {
		return nil, false
	}
	return o.RequestedQoS, true
}

// HasRequestedQoS returns a boolean if a field has been set.
func (o *NetworkAssistanceSession) HasRequestedQoS() bool {
	if o != nil && !IsNil(o.RequestedQoS) {
		return true
	}

	return false
}

// SetRequestedQoS gets a reference to the given M5QoSSpecification and assigns it to the RequestedQoS field.
func (o *NetworkAssistanceSession) SetRequestedQoS(v M5QoSSpecification) {
	o.RequestedQoS = &v
}

// GetRecommendedQoS returns the RecommendedQoS field value if set, zero value otherwise.
func (o *NetworkAssistanceSession) GetRecommendedQoS() M5QoSSpecification {
	if o == nil || IsNil(o.RecommendedQoS) {
		var ret M5QoSSpecification
		return ret
	}
	return *o.RecommendedQoS
}

// GetRecommendedQoSOk returns a tuple with the RecommendedQoS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAssistanceSession) GetRecommendedQoSOk() (*M5QoSSpecification, bool) {
	if o == nil || IsNil(o.RecommendedQoS) {
		return nil, false
	}
	return o.RecommendedQoS, true
}

// HasRecommendedQoS returns a boolean if a field has been set.
func (o *NetworkAssistanceSession) HasRecommendedQoS() bool {
	if o != nil && !IsNil(o.RecommendedQoS) {
		return true
	}

	return false
}

// SetRecommendedQoS gets a reference to the given M5QoSSpecification and assigns it to the RecommendedQoS field.
func (o *NetworkAssistanceSession) SetRecommendedQoS(v M5QoSSpecification) {
	o.RecommendedQoS = &v
}

// GetNotficationURL returns the NotficationURL field value if set, zero value otherwise.
func (o *NetworkAssistanceSession) GetNotficationURL() string {
	if o == nil || IsNil(o.NotficationURL) {
		var ret string
		return ret
	}
	return *o.NotficationURL
}

// GetNotficationURLOk returns a tuple with the NotficationURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkAssistanceSession) GetNotficationURLOk() (*string, bool) {
	if o == nil || IsNil(o.NotficationURL) {
		return nil, false
	}
	return o.NotficationURL, true
}

// HasNotficationURL returns a boolean if a field has been set.
func (o *NetworkAssistanceSession) HasNotficationURL() bool {
	if o != nil && !IsNil(o.NotficationURL) {
		return true
	}

	return false
}

// SetNotficationURL gets a reference to the given string and assigns it to the NotficationURL field.
func (o *NetworkAssistanceSession) SetNotficationURL(v string) {
	o.NotficationURL = &v
}

func (o NetworkAssistanceSession) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkAssistanceSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["naSessionId"] = o.NaSessionId
	if !IsNil(o.ServiceDataFlowDescription) {
		toSerialize["serviceDataFlowDescription"] = o.ServiceDataFlowDescription
	}
	if !IsNil(o.PolicyTemplateId) {
		toSerialize["policyTemplateId"] = o.PolicyTemplateId
	}
	if !IsNil(o.RequestedQoS) {
		toSerialize["requestedQoS"] = o.RequestedQoS
	}
	if !IsNil(o.RecommendedQoS) {
		toSerialize["recommendedQoS"] = o.RecommendedQoS
	}
	if !IsNil(o.NotficationURL) {
		toSerialize["notficationURL"] = o.NotficationURL
	}
	return toSerialize, nil
}

type NullableNetworkAssistanceSession struct {
	value *NetworkAssistanceSession
	isSet bool
}

func (v NullableNetworkAssistanceSession) Get() *NetworkAssistanceSession {
	return v.value
}

func (v *NullableNetworkAssistanceSession) Set(val *NetworkAssistanceSession) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkAssistanceSession) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkAssistanceSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkAssistanceSession(val *NetworkAssistanceSession) *NullableNetworkAssistanceSession {
	return &NullableNetworkAssistanceSession{value: val, isSet: true}
}

func (v NullableNetworkAssistanceSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkAssistanceSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
