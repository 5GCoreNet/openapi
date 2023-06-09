/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the MMTelChargingInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MMTelChargingInformation{}

// MMTelChargingInformation struct for MMTelChargingInformation
type MMTelChargingInformation struct {
	SupplementaryServices []SupplementaryService `json:"supplementaryServices,omitempty"`
}

// NewMMTelChargingInformation instantiates a new MMTelChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMMTelChargingInformation() *MMTelChargingInformation {
	this := MMTelChargingInformation{}
	return &this
}

// NewMMTelChargingInformationWithDefaults instantiates a new MMTelChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMMTelChargingInformationWithDefaults() *MMTelChargingInformation {
	this := MMTelChargingInformation{}
	return &this
}

// GetSupplementaryServices returns the SupplementaryServices field value if set, zero value otherwise.
func (o *MMTelChargingInformation) GetSupplementaryServices() []SupplementaryService {
	if o == nil || IsNil(o.SupplementaryServices) {
		var ret []SupplementaryService
		return ret
	}
	return o.SupplementaryServices
}

// GetSupplementaryServicesOk returns a tuple with the SupplementaryServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MMTelChargingInformation) GetSupplementaryServicesOk() ([]SupplementaryService, bool) {
	if o == nil || IsNil(o.SupplementaryServices) {
		return nil, false
	}
	return o.SupplementaryServices, true
}

// HasSupplementaryServices returns a boolean if a field has been set.
func (o *MMTelChargingInformation) HasSupplementaryServices() bool {
	if o != nil && !IsNil(o.SupplementaryServices) {
		return true
	}

	return false
}

// SetSupplementaryServices gets a reference to the given []SupplementaryService and assigns it to the SupplementaryServices field.
func (o *MMTelChargingInformation) SetSupplementaryServices(v []SupplementaryService) {
	o.SupplementaryServices = v
}

func (o MMTelChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MMTelChargingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupplementaryServices) {
		toSerialize["supplementaryServices"] = o.SupplementaryServices
	}
	return toSerialize, nil
}

type NullableMMTelChargingInformation struct {
	value *MMTelChargingInformation
	isSet bool
}

func (v NullableMMTelChargingInformation) Get() *MMTelChargingInformation {
	return v.value
}

func (v *NullableMMTelChargingInformation) Set(val *MMTelChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableMMTelChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableMMTelChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMMTelChargingInformation(val *MMTelChargingInformation) *NullableMMTelChargingInformation {
	return &NullableMMTelChargingInformation{value: val, isSet: true}
}

func (v NullableMMTelChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMMTelChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
