/*
M5_ServiceAccessInformation

5GMS AF M5 Service Access Information API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_ServiceAccessInformation

import (
	"encoding/json"
)

// checks if the ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration{}

// ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration struct for ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration
type ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration struct {
	// A set of application endpoint addresses.
	ServerAddresses        []string    `json:"serverAddresses"`
	ValidPolicyTemplateIds []string    `json:"validPolicyTemplateIds"`
	SdfMethods             []SdfMethod `json:"sdfMethods"`
	ExternalReferences     []string    `json:"externalReferences,omitempty"`
}

// NewServiceAccessInformationResourceDynamicPolicyInvocationConfiguration instantiates a new ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccessInformationResourceDynamicPolicyInvocationConfiguration(serverAddresses []string, validPolicyTemplateIds []string, sdfMethods []SdfMethod) *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration {
	this := ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration{}
	this.ServerAddresses = serverAddresses
	this.ValidPolicyTemplateIds = validPolicyTemplateIds
	this.SdfMethods = sdfMethods
	return &this
}

// NewServiceAccessInformationResourceDynamicPolicyInvocationConfigurationWithDefaults instantiates a new ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccessInformationResourceDynamicPolicyInvocationConfigurationWithDefaults() *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration {
	this := ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration{}
	return &this
}

// GetServerAddresses returns the ServerAddresses field value
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetServerAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ServerAddresses
}

// GetServerAddressesOk returns a tuple with the ServerAddresses field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetServerAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerAddresses, true
}

// SetServerAddresses sets field value
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) SetServerAddresses(v []string) {
	o.ServerAddresses = v
}

// GetValidPolicyTemplateIds returns the ValidPolicyTemplateIds field value
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetValidPolicyTemplateIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ValidPolicyTemplateIds
}

// GetValidPolicyTemplateIdsOk returns a tuple with the ValidPolicyTemplateIds field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetValidPolicyTemplateIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValidPolicyTemplateIds, true
}

// SetValidPolicyTemplateIds sets field value
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) SetValidPolicyTemplateIds(v []string) {
	o.ValidPolicyTemplateIds = v
}

// GetSdfMethods returns the SdfMethods field value
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetSdfMethods() []SdfMethod {
	if o == nil {
		var ret []SdfMethod
		return ret
	}

	return o.SdfMethods
}

// GetSdfMethodsOk returns a tuple with the SdfMethods field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetSdfMethodsOk() ([]SdfMethod, bool) {
	if o == nil {
		return nil, false
	}
	return o.SdfMethods, true
}

// SetSdfMethods sets field value
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) SetSdfMethods(v []SdfMethod) {
	o.SdfMethods = v
}

// GetExternalReferences returns the ExternalReferences field value if set, zero value otherwise.
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetExternalReferences() []string {
	if o == nil || IsNil(o.ExternalReferences) {
		var ret []string
		return ret
	}
	return o.ExternalReferences
}

// GetExternalReferencesOk returns a tuple with the ExternalReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) GetExternalReferencesOk() ([]string, bool) {
	if o == nil || IsNil(o.ExternalReferences) {
		return nil, false
	}
	return o.ExternalReferences, true
}

// HasExternalReferences returns a boolean if a field has been set.
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) HasExternalReferences() bool {
	if o != nil && !IsNil(o.ExternalReferences) {
		return true
	}

	return false
}

// SetExternalReferences gets a reference to the given []string and assigns it to the ExternalReferences field.
func (o *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) SetExternalReferences(v []string) {
	o.ExternalReferences = v
}

func (o ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverAddresses"] = o.ServerAddresses
	toSerialize["validPolicyTemplateIds"] = o.ValidPolicyTemplateIds
	toSerialize["sdfMethods"] = o.SdfMethods
	if !IsNil(o.ExternalReferences) {
		toSerialize["externalReferences"] = o.ExternalReferences
	}
	return toSerialize, nil
}

type NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration struct {
	value *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration
	isSet bool
}

func (v NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) Get() *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration {
	return v.value
}

func (v *NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) Set(val *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration(val *ServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) *NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration {
	return &NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration{value: val, isSet: true}
}

func (v NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccessInformationResourceDynamicPolicyInvocationConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
