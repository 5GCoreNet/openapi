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

// checks if the NSMChargingInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NSMChargingInformation{}

// NSMChargingInformation struct for NSMChargingInformation
type NSMChargingInformation struct {
	ManagementOperation                     ManagementOperation                 `json:"managementOperation"`
	IdNetworkSliceInstance                  *string                             `json:"idNetworkSliceInstance,omitempty"`
	ListOfserviceProfileChargingInformation []ServiceProfileChargingInformation `json:"listOfserviceProfileChargingInformation,omitempty"`
	ManagementOperationStatus               *ManagementOperationStatus          `json:"managementOperationStatus,omitempty"`
}

// NewNSMChargingInformation instantiates a new NSMChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNSMChargingInformation(managementOperation ManagementOperation) *NSMChargingInformation {
	this := NSMChargingInformation{}
	this.ManagementOperation = managementOperation
	return &this
}

// NewNSMChargingInformationWithDefaults instantiates a new NSMChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNSMChargingInformationWithDefaults() *NSMChargingInformation {
	this := NSMChargingInformation{}
	return &this
}

// GetManagementOperation returns the ManagementOperation field value
func (o *NSMChargingInformation) GetManagementOperation() ManagementOperation {
	if o == nil {
		var ret ManagementOperation
		return ret
	}

	return o.ManagementOperation
}

// GetManagementOperationOk returns a tuple with the ManagementOperation field value
// and a boolean to check if the value has been set.
func (o *NSMChargingInformation) GetManagementOperationOk() (*ManagementOperation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManagementOperation, true
}

// SetManagementOperation sets field value
func (o *NSMChargingInformation) SetManagementOperation(v ManagementOperation) {
	o.ManagementOperation = v
}

// GetIdNetworkSliceInstance returns the IdNetworkSliceInstance field value if set, zero value otherwise.
func (o *NSMChargingInformation) GetIdNetworkSliceInstance() string {
	if o == nil || IsNil(o.IdNetworkSliceInstance) {
		var ret string
		return ret
	}
	return *o.IdNetworkSliceInstance
}

// GetIdNetworkSliceInstanceOk returns a tuple with the IdNetworkSliceInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSMChargingInformation) GetIdNetworkSliceInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.IdNetworkSliceInstance) {
		return nil, false
	}
	return o.IdNetworkSliceInstance, true
}

// HasIdNetworkSliceInstance returns a boolean if a field has been set.
func (o *NSMChargingInformation) HasIdNetworkSliceInstance() bool {
	if o != nil && !IsNil(o.IdNetworkSliceInstance) {
		return true
	}

	return false
}

// SetIdNetworkSliceInstance gets a reference to the given string and assigns it to the IdNetworkSliceInstance field.
func (o *NSMChargingInformation) SetIdNetworkSliceInstance(v string) {
	o.IdNetworkSliceInstance = &v
}

// GetListOfserviceProfileChargingInformation returns the ListOfserviceProfileChargingInformation field value if set, zero value otherwise.
func (o *NSMChargingInformation) GetListOfserviceProfileChargingInformation() []ServiceProfileChargingInformation {
	if o == nil || IsNil(o.ListOfserviceProfileChargingInformation) {
		var ret []ServiceProfileChargingInformation
		return ret
	}
	return o.ListOfserviceProfileChargingInformation
}

// GetListOfserviceProfileChargingInformationOk returns a tuple with the ListOfserviceProfileChargingInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSMChargingInformation) GetListOfserviceProfileChargingInformationOk() ([]ServiceProfileChargingInformation, bool) {
	if o == nil || IsNil(o.ListOfserviceProfileChargingInformation) {
		return nil, false
	}
	return o.ListOfserviceProfileChargingInformation, true
}

// HasListOfserviceProfileChargingInformation returns a boolean if a field has been set.
func (o *NSMChargingInformation) HasListOfserviceProfileChargingInformation() bool {
	if o != nil && !IsNil(o.ListOfserviceProfileChargingInformation) {
		return true
	}

	return false
}

// SetListOfserviceProfileChargingInformation gets a reference to the given []ServiceProfileChargingInformation and assigns it to the ListOfserviceProfileChargingInformation field.
func (o *NSMChargingInformation) SetListOfserviceProfileChargingInformation(v []ServiceProfileChargingInformation) {
	o.ListOfserviceProfileChargingInformation = v
}

// GetManagementOperationStatus returns the ManagementOperationStatus field value if set, zero value otherwise.
func (o *NSMChargingInformation) GetManagementOperationStatus() ManagementOperationStatus {
	if o == nil || IsNil(o.ManagementOperationStatus) {
		var ret ManagementOperationStatus
		return ret
	}
	return *o.ManagementOperationStatus
}

// GetManagementOperationStatusOk returns a tuple with the ManagementOperationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NSMChargingInformation) GetManagementOperationStatusOk() (*ManagementOperationStatus, bool) {
	if o == nil || IsNil(o.ManagementOperationStatus) {
		return nil, false
	}
	return o.ManagementOperationStatus, true
}

// HasManagementOperationStatus returns a boolean if a field has been set.
func (o *NSMChargingInformation) HasManagementOperationStatus() bool {
	if o != nil && !IsNil(o.ManagementOperationStatus) {
		return true
	}

	return false
}

// SetManagementOperationStatus gets a reference to the given ManagementOperationStatus and assigns it to the ManagementOperationStatus field.
func (o *NSMChargingInformation) SetManagementOperationStatus(v ManagementOperationStatus) {
	o.ManagementOperationStatus = &v
}

func (o NSMChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NSMChargingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["managementOperation"] = o.ManagementOperation
	if !IsNil(o.IdNetworkSliceInstance) {
		toSerialize["idNetworkSliceInstance"] = o.IdNetworkSliceInstance
	}
	if !IsNil(o.ListOfserviceProfileChargingInformation) {
		toSerialize["listOfserviceProfileChargingInformation"] = o.ListOfserviceProfileChargingInformation
	}
	if !IsNil(o.ManagementOperationStatus) {
		toSerialize["managementOperationStatus"] = o.ManagementOperationStatus
	}
	return toSerialize, nil
}

type NullableNSMChargingInformation struct {
	value *NSMChargingInformation
	isSet bool
}

func (v NullableNSMChargingInformation) Get() *NSMChargingInformation {
	return v.value
}

func (v *NullableNSMChargingInformation) Set(val *NSMChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableNSMChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableNSMChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNSMChargingInformation(val *NSMChargingInformation) *NullableNSMChargingInformation {
	return &NullableNSMChargingInformation{value: val, isSet: true}
}

func (v NullableNSMChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNSMChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
