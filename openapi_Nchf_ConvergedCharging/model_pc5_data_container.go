/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"time"
)

// checks if the PC5DataContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PC5DataContainer{}

// PC5DataContainer struct for PC5DataContainer
type PC5DataContainer struct {
	LocalSequenceNumber *string `json:"localSequenceNumber,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ChangeTime              *time.Time    `json:"changeTime,omitempty"`
	CoverageStatus          *bool         `json:"coverageStatus,omitempty"`
	UserLocationInformation *UserLocation `json:"userLocationInformation,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	DataVolume         *int32            `json:"dataVolume,omitempty"`
	ChangeCondition    *string           `json:"changeCondition,omitempty"`
	RadioResourcesId   *RadioResourcesId `json:"radioResourcesId,omitempty"`
	RadioFrequency     *string           `json:"radioFrequency,omitempty"`
	PC5RadioTechnology *string           `json:"pC5RadioTechnology,omitempty"`
}

// NewPC5DataContainer instantiates a new PC5DataContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPC5DataContainer() *PC5DataContainer {
	this := PC5DataContainer{}
	return &this
}

// NewPC5DataContainerWithDefaults instantiates a new PC5DataContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPC5DataContainerWithDefaults() *PC5DataContainer {
	this := PC5DataContainer{}
	return &this
}

// GetLocalSequenceNumber returns the LocalSequenceNumber field value if set, zero value otherwise.
func (o *PC5DataContainer) GetLocalSequenceNumber() string {
	if o == nil || IsNil(o.LocalSequenceNumber) {
		var ret string
		return ret
	}
	return *o.LocalSequenceNumber
}

// GetLocalSequenceNumberOk returns a tuple with the LocalSequenceNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetLocalSequenceNumberOk() (*string, bool) {
	if o == nil || IsNil(o.LocalSequenceNumber) {
		return nil, false
	}
	return o.LocalSequenceNumber, true
}

// HasLocalSequenceNumber returns a boolean if a field has been set.
func (o *PC5DataContainer) HasLocalSequenceNumber() bool {
	if o != nil && !IsNil(o.LocalSequenceNumber) {
		return true
	}

	return false
}

// SetLocalSequenceNumber gets a reference to the given string and assigns it to the LocalSequenceNumber field.
func (o *PC5DataContainer) SetLocalSequenceNumber(v string) {
	o.LocalSequenceNumber = &v
}

// GetChangeTime returns the ChangeTime field value if set, zero value otherwise.
func (o *PC5DataContainer) GetChangeTime() time.Time {
	if o == nil || IsNil(o.ChangeTime) {
		var ret time.Time
		return ret
	}
	return *o.ChangeTime
}

// GetChangeTimeOk returns a tuple with the ChangeTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetChangeTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ChangeTime) {
		return nil, false
	}
	return o.ChangeTime, true
}

// HasChangeTime returns a boolean if a field has been set.
func (o *PC5DataContainer) HasChangeTime() bool {
	if o != nil && !IsNil(o.ChangeTime) {
		return true
	}

	return false
}

// SetChangeTime gets a reference to the given time.Time and assigns it to the ChangeTime field.
func (o *PC5DataContainer) SetChangeTime(v time.Time) {
	o.ChangeTime = &v
}

// GetCoverageStatus returns the CoverageStatus field value if set, zero value otherwise.
func (o *PC5DataContainer) GetCoverageStatus() bool {
	if o == nil || IsNil(o.CoverageStatus) {
		var ret bool
		return ret
	}
	return *o.CoverageStatus
}

// GetCoverageStatusOk returns a tuple with the CoverageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetCoverageStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.CoverageStatus) {
		return nil, false
	}
	return o.CoverageStatus, true
}

// HasCoverageStatus returns a boolean if a field has been set.
func (o *PC5DataContainer) HasCoverageStatus() bool {
	if o != nil && !IsNil(o.CoverageStatus) {
		return true
	}

	return false
}

// SetCoverageStatus gets a reference to the given bool and assigns it to the CoverageStatus field.
func (o *PC5DataContainer) SetCoverageStatus(v bool) {
	o.CoverageStatus = &v
}

// GetUserLocationInformation returns the UserLocationInformation field value if set, zero value otherwise.
func (o *PC5DataContainer) GetUserLocationInformation() UserLocation {
	if o == nil || IsNil(o.UserLocationInformation) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationInformation
}

// GetUserLocationInformationOk returns a tuple with the UserLocationInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetUserLocationInformationOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLocationInformation) {
		return nil, false
	}
	return o.UserLocationInformation, true
}

// HasUserLocationInformation returns a boolean if a field has been set.
func (o *PC5DataContainer) HasUserLocationInformation() bool {
	if o != nil && !IsNil(o.UserLocationInformation) {
		return true
	}

	return false
}

// SetUserLocationInformation gets a reference to the given UserLocation and assigns it to the UserLocationInformation field.
func (o *PC5DataContainer) SetUserLocationInformation(v UserLocation) {
	o.UserLocationInformation = &v
}

// GetDataVolume returns the DataVolume field value if set, zero value otherwise.
func (o *PC5DataContainer) GetDataVolume() int32 {
	if o == nil || IsNil(o.DataVolume) {
		var ret int32
		return ret
	}
	return *o.DataVolume
}

// GetDataVolumeOk returns a tuple with the DataVolume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetDataVolumeOk() (*int32, bool) {
	if o == nil || IsNil(o.DataVolume) {
		return nil, false
	}
	return o.DataVolume, true
}

// HasDataVolume returns a boolean if a field has been set.
func (o *PC5DataContainer) HasDataVolume() bool {
	if o != nil && !IsNil(o.DataVolume) {
		return true
	}

	return false
}

// SetDataVolume gets a reference to the given int32 and assigns it to the DataVolume field.
func (o *PC5DataContainer) SetDataVolume(v int32) {
	o.DataVolume = &v
}

// GetChangeCondition returns the ChangeCondition field value if set, zero value otherwise.
func (o *PC5DataContainer) GetChangeCondition() string {
	if o == nil || IsNil(o.ChangeCondition) {
		var ret string
		return ret
	}
	return *o.ChangeCondition
}

// GetChangeConditionOk returns a tuple with the ChangeCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetChangeConditionOk() (*string, bool) {
	if o == nil || IsNil(o.ChangeCondition) {
		return nil, false
	}
	return o.ChangeCondition, true
}

// HasChangeCondition returns a boolean if a field has been set.
func (o *PC5DataContainer) HasChangeCondition() bool {
	if o != nil && !IsNil(o.ChangeCondition) {
		return true
	}

	return false
}

// SetChangeCondition gets a reference to the given string and assigns it to the ChangeCondition field.
func (o *PC5DataContainer) SetChangeCondition(v string) {
	o.ChangeCondition = &v
}

// GetRadioResourcesId returns the RadioResourcesId field value if set, zero value otherwise.
func (o *PC5DataContainer) GetRadioResourcesId() RadioResourcesId {
	if o == nil || IsNil(o.RadioResourcesId) {
		var ret RadioResourcesId
		return ret
	}
	return *o.RadioResourcesId
}

// GetRadioResourcesIdOk returns a tuple with the RadioResourcesId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetRadioResourcesIdOk() (*RadioResourcesId, bool) {
	if o == nil || IsNil(o.RadioResourcesId) {
		return nil, false
	}
	return o.RadioResourcesId, true
}

// HasRadioResourcesId returns a boolean if a field has been set.
func (o *PC5DataContainer) HasRadioResourcesId() bool {
	if o != nil && !IsNil(o.RadioResourcesId) {
		return true
	}

	return false
}

// SetRadioResourcesId gets a reference to the given RadioResourcesId and assigns it to the RadioResourcesId field.
func (o *PC5DataContainer) SetRadioResourcesId(v RadioResourcesId) {
	o.RadioResourcesId = &v
}

// GetRadioFrequency returns the RadioFrequency field value if set, zero value otherwise.
func (o *PC5DataContainer) GetRadioFrequency() string {
	if o == nil || IsNil(o.RadioFrequency) {
		var ret string
		return ret
	}
	return *o.RadioFrequency
}

// GetRadioFrequencyOk returns a tuple with the RadioFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetRadioFrequencyOk() (*string, bool) {
	if o == nil || IsNil(o.RadioFrequency) {
		return nil, false
	}
	return o.RadioFrequency, true
}

// HasRadioFrequency returns a boolean if a field has been set.
func (o *PC5DataContainer) HasRadioFrequency() bool {
	if o != nil && !IsNil(o.RadioFrequency) {
		return true
	}

	return false
}

// SetRadioFrequency gets a reference to the given string and assigns it to the RadioFrequency field.
func (o *PC5DataContainer) SetRadioFrequency(v string) {
	o.RadioFrequency = &v
}

// GetPC5RadioTechnology returns the PC5RadioTechnology field value if set, zero value otherwise.
func (o *PC5DataContainer) GetPC5RadioTechnology() string {
	if o == nil || IsNil(o.PC5RadioTechnology) {
		var ret string
		return ret
	}
	return *o.PC5RadioTechnology
}

// GetPC5RadioTechnologyOk returns a tuple with the PC5RadioTechnology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PC5DataContainer) GetPC5RadioTechnologyOk() (*string, bool) {
	if o == nil || IsNil(o.PC5RadioTechnology) {
		return nil, false
	}
	return o.PC5RadioTechnology, true
}

// HasPC5RadioTechnology returns a boolean if a field has been set.
func (o *PC5DataContainer) HasPC5RadioTechnology() bool {
	if o != nil && !IsNil(o.PC5RadioTechnology) {
		return true
	}

	return false
}

// SetPC5RadioTechnology gets a reference to the given string and assigns it to the PC5RadioTechnology field.
func (o *PC5DataContainer) SetPC5RadioTechnology(v string) {
	o.PC5RadioTechnology = &v
}

func (o PC5DataContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PC5DataContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocalSequenceNumber) {
		toSerialize["localSequenceNumber"] = o.LocalSequenceNumber
	}
	if !IsNil(o.ChangeTime) {
		toSerialize["changeTime"] = o.ChangeTime
	}
	if !IsNil(o.CoverageStatus) {
		toSerialize["coverageStatus"] = o.CoverageStatus
	}
	if !IsNil(o.UserLocationInformation) {
		toSerialize["userLocationInformation"] = o.UserLocationInformation
	}
	if !IsNil(o.DataVolume) {
		toSerialize["dataVolume"] = o.DataVolume
	}
	if !IsNil(o.ChangeCondition) {
		toSerialize["changeCondition"] = o.ChangeCondition
	}
	if !IsNil(o.RadioResourcesId) {
		toSerialize["radioResourcesId"] = o.RadioResourcesId
	}
	if !IsNil(o.RadioFrequency) {
		toSerialize["radioFrequency"] = o.RadioFrequency
	}
	if !IsNil(o.PC5RadioTechnology) {
		toSerialize["pC5RadioTechnology"] = o.PC5RadioTechnology
	}
	return toSerialize, nil
}

type NullablePC5DataContainer struct {
	value *PC5DataContainer
	isSet bool
}

func (v NullablePC5DataContainer) Get() *PC5DataContainer {
	return v.value
}

func (v *NullablePC5DataContainer) Set(val *PC5DataContainer) {
	v.value = val
	v.isSet = true
}

func (v NullablePC5DataContainer) IsSet() bool {
	return v.isSet
}

func (v *NullablePC5DataContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePC5DataContainer(val *PC5DataContainer) *NullablePC5DataContainer {
	return &NullablePC5DataContainer{value: val, isSet: true}
}

func (v NullablePC5DataContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePC5DataContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
