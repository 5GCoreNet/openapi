/*
M5_NetworkAssistance

5GMS AF M5 Network Assistance API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_NetworkAssistance

import (
	"encoding/json"
)

// checks if the ServiceDataFlowDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDataFlowDescription{}

// ServiceDataFlowDescription struct for ServiceDataFlowDescription
type ServiceDataFlowDescription struct {
	FlowDescription *IpPacketFilterSet `json:"flowDescription,omitempty"`
	DomainName *string `json:"domainName,omitempty"`
}

// NewServiceDataFlowDescription instantiates a new ServiceDataFlowDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDataFlowDescription() *ServiceDataFlowDescription {
	this := ServiceDataFlowDescription{}
	return &this
}

// NewServiceDataFlowDescriptionWithDefaults instantiates a new ServiceDataFlowDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDataFlowDescriptionWithDefaults() *ServiceDataFlowDescription {
	this := ServiceDataFlowDescription{}
	return &this
}

// GetFlowDescription returns the FlowDescription field value if set, zero value otherwise.
func (o *ServiceDataFlowDescription) GetFlowDescription() IpPacketFilterSet {
	if o == nil || IsNil(o.FlowDescription) {
		var ret IpPacketFilterSet
		return ret
	}
	return *o.FlowDescription
}

// GetFlowDescriptionOk returns a tuple with the FlowDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDataFlowDescription) GetFlowDescriptionOk() (*IpPacketFilterSet, bool) {
	if o == nil || IsNil(o.FlowDescription) {
		return nil, false
	}
	return o.FlowDescription, true
}

// HasFlowDescription returns a boolean if a field has been set.
func (o *ServiceDataFlowDescription) HasFlowDescription() bool {
	if o != nil && !IsNil(o.FlowDescription) {
		return true
	}

	return false
}

// SetFlowDescription gets a reference to the given IpPacketFilterSet and assigns it to the FlowDescription field.
func (o *ServiceDataFlowDescription) SetFlowDescription(v IpPacketFilterSet) {
	o.FlowDescription = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *ServiceDataFlowDescription) GetDomainName() string {
	if o == nil || IsNil(o.DomainName) {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDataFlowDescription) GetDomainNameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainName) {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *ServiceDataFlowDescription) HasDomainName() bool {
	if o != nil && !IsNil(o.DomainName) {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *ServiceDataFlowDescription) SetDomainName(v string) {
	o.DomainName = &v
}

func (o ServiceDataFlowDescription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDataFlowDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FlowDescription) {
		toSerialize["flowDescription"] = o.FlowDescription
	}
	if !IsNil(o.DomainName) {
		toSerialize["domainName"] = o.DomainName
	}
	return toSerialize, nil
}

type NullableServiceDataFlowDescription struct {
	value *ServiceDataFlowDescription
	isSet bool
}

func (v NullableServiceDataFlowDescription) Get() *ServiceDataFlowDescription {
	return v.value
}

func (v *NullableServiceDataFlowDescription) Set(val *ServiceDataFlowDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDataFlowDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDataFlowDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDataFlowDescription(val *ServiceDataFlowDescription) *NullableServiceDataFlowDescription {
	return &NullableServiceDataFlowDescription{value: val, isSet: true}
}

func (v NullableServiceDataFlowDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDataFlowDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


