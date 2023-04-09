/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"time"
)

// checks if the SorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SorInfo{}

// SorInfo struct for SorInfo
type SorInfo struct {
	SteeringContainer *SteeringContainer `json:"steeringContainer,omitempty"`
	// Contains indication whether the acknowledgement from UE is needed.
	AckInd bool `json:"ackInd"`
	// MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE).
	SorMacIausf *string `json:"sorMacIausf,omitempty"`
	// CounterSoR.
	Countersor *string `json:"countersor,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`
	// string with format 'bytes' as defined in OpenAPI
	SorTransparentContainer *string `json:"sorTransparentContainer,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	SorCmci              *string `json:"sorCmci,omitempty"`
	StoreSorCmciInMe     *bool   `json:"storeSorCmciInMe,omitempty"`
	UsimSupportOfSorCmci *bool   `json:"usimSupportOfSorCmci,omitempty"`
}

// NewSorInfo instantiates a new SorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSorInfo(ackInd bool, provisioningTime time.Time) *SorInfo {
	this := SorInfo{}
	this.AckInd = ackInd
	this.ProvisioningTime = provisioningTime
	return &this
}

// NewSorInfoWithDefaults instantiates a new SorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSorInfoWithDefaults() *SorInfo {
	this := SorInfo{}
	return &this
}

// GetSteeringContainer returns the SteeringContainer field value if set, zero value otherwise.
func (o *SorInfo) GetSteeringContainer() SteeringContainer {
	if o == nil || IsNil(o.SteeringContainer) {
		var ret SteeringContainer
		return ret
	}
	return *o.SteeringContainer
}

// GetSteeringContainerOk returns a tuple with the SteeringContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInfo) GetSteeringContainerOk() (*SteeringContainer, bool) {
	if o == nil || IsNil(o.SteeringContainer) {
		return nil, false
	}
	return o.SteeringContainer, true
}

// HasSteeringContainer returns a boolean if a field has been set.
func (o *SorInfo) HasSteeringContainer() bool {
	if o != nil && !IsNil(o.SteeringContainer) {
		return true
	}

	return false
}

// SetSteeringContainer gets a reference to the given SteeringContainer and assigns it to the SteeringContainer field.
func (o *SorInfo) SetSteeringContainer(v SteeringContainer) {
	o.SteeringContainer = &v
}

// GetAckInd returns the AckInd field value
func (o *SorInfo) GetAckInd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AckInd
}

// GetAckIndOk returns a tuple with the AckInd field value
// and a boolean to check if the value has been set.
func (o *SorInfo) GetAckIndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AckInd, true
}

// SetAckInd sets field value
func (o *SorInfo) SetAckInd(v bool) {
	o.AckInd = v
}

// GetSorMacIausf returns the SorMacIausf field value if set, zero value otherwise.
func (o *SorInfo) GetSorMacIausf() string {
	if o == nil || IsNil(o.SorMacIausf) {
		var ret string
		return ret
	}
	return *o.SorMacIausf
}

// GetSorMacIausfOk returns a tuple with the SorMacIausf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInfo) GetSorMacIausfOk() (*string, bool) {
	if o == nil || IsNil(o.SorMacIausf) {
		return nil, false
	}
	return o.SorMacIausf, true
}

// HasSorMacIausf returns a boolean if a field has been set.
func (o *SorInfo) HasSorMacIausf() bool {
	if o != nil && !IsNil(o.SorMacIausf) {
		return true
	}

	return false
}

// SetSorMacIausf gets a reference to the given string and assigns it to the SorMacIausf field.
func (o *SorInfo) SetSorMacIausf(v string) {
	o.SorMacIausf = &v
}

// GetCountersor returns the Countersor field value if set, zero value otherwise.
func (o *SorInfo) GetCountersor() string {
	if o == nil || IsNil(o.Countersor) {
		var ret string
		return ret
	}
	return *o.Countersor
}

// GetCountersorOk returns a tuple with the Countersor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInfo) GetCountersorOk() (*string, bool) {
	if o == nil || IsNil(o.Countersor) {
		return nil, false
	}
	return o.Countersor, true
}

// HasCountersor returns a boolean if a field has been set.
func (o *SorInfo) HasCountersor() bool {
	if o != nil && !IsNil(o.Countersor) {
		return true
	}

	return false
}

// SetCountersor gets a reference to the given string and assigns it to the Countersor field.
func (o *SorInfo) SetCountersor(v string) {
	o.Countersor = &v
}

// GetProvisioningTime returns the ProvisioningTime field value
func (o *SorInfo) GetProvisioningTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value
// and a boolean to check if the value has been set.
func (o *SorInfo) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningTime, true
}

// SetProvisioningTime sets field value
func (o *SorInfo) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = v
}

// GetSorTransparentContainer returns the SorTransparentContainer field value if set, zero value otherwise.
func (o *SorInfo) GetSorTransparentContainer() string {
	if o == nil || IsNil(o.SorTransparentContainer) {
		var ret string
		return ret
	}
	return *o.SorTransparentContainer
}

// GetSorTransparentContainerOk returns a tuple with the SorTransparentContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInfo) GetSorTransparentContainerOk() (*string, bool) {
	if o == nil || IsNil(o.SorTransparentContainer) {
		return nil, false
	}
	return o.SorTransparentContainer, true
}

// HasSorTransparentContainer returns a boolean if a field has been set.
func (o *SorInfo) HasSorTransparentContainer() bool {
	if o != nil && !IsNil(o.SorTransparentContainer) {
		return true
	}

	return false
}

// SetSorTransparentContainer gets a reference to the given string and assigns it to the SorTransparentContainer field.
func (o *SorInfo) SetSorTransparentContainer(v string) {
	o.SorTransparentContainer = &v
}

// GetSorCmci returns the SorCmci field value if set, zero value otherwise.
func (o *SorInfo) GetSorCmci() string {
	if o == nil || IsNil(o.SorCmci) {
		var ret string
		return ret
	}
	return *o.SorCmci
}

// GetSorCmciOk returns a tuple with the SorCmci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInfo) GetSorCmciOk() (*string, bool) {
	if o == nil || IsNil(o.SorCmci) {
		return nil, false
	}
	return o.SorCmci, true
}

// HasSorCmci returns a boolean if a field has been set.
func (o *SorInfo) HasSorCmci() bool {
	if o != nil && !IsNil(o.SorCmci) {
		return true
	}

	return false
}

// SetSorCmci gets a reference to the given string and assigns it to the SorCmci field.
func (o *SorInfo) SetSorCmci(v string) {
	o.SorCmci = &v
}

// GetStoreSorCmciInMe returns the StoreSorCmciInMe field value if set, zero value otherwise.
func (o *SorInfo) GetStoreSorCmciInMe() bool {
	if o == nil || IsNil(o.StoreSorCmciInMe) {
		var ret bool
		return ret
	}
	return *o.StoreSorCmciInMe
}

// GetStoreSorCmciInMeOk returns a tuple with the StoreSorCmciInMe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInfo) GetStoreSorCmciInMeOk() (*bool, bool) {
	if o == nil || IsNil(o.StoreSorCmciInMe) {
		return nil, false
	}
	return o.StoreSorCmciInMe, true
}

// HasStoreSorCmciInMe returns a boolean if a field has been set.
func (o *SorInfo) HasStoreSorCmciInMe() bool {
	if o != nil && !IsNil(o.StoreSorCmciInMe) {
		return true
	}

	return false
}

// SetStoreSorCmciInMe gets a reference to the given bool and assigns it to the StoreSorCmciInMe field.
func (o *SorInfo) SetStoreSorCmciInMe(v bool) {
	o.StoreSorCmciInMe = &v
}

// GetUsimSupportOfSorCmci returns the UsimSupportOfSorCmci field value if set, zero value otherwise.
func (o *SorInfo) GetUsimSupportOfSorCmci() bool {
	if o == nil || IsNil(o.UsimSupportOfSorCmci) {
		var ret bool
		return ret
	}
	return *o.UsimSupportOfSorCmci
}

// GetUsimSupportOfSorCmciOk returns a tuple with the UsimSupportOfSorCmci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SorInfo) GetUsimSupportOfSorCmciOk() (*bool, bool) {
	if o == nil || IsNil(o.UsimSupportOfSorCmci) {
		return nil, false
	}
	return o.UsimSupportOfSorCmci, true
}

// HasUsimSupportOfSorCmci returns a boolean if a field has been set.
func (o *SorInfo) HasUsimSupportOfSorCmci() bool {
	if o != nil && !IsNil(o.UsimSupportOfSorCmci) {
		return true
	}

	return false
}

// SetUsimSupportOfSorCmci gets a reference to the given bool and assigns it to the UsimSupportOfSorCmci field.
func (o *SorInfo) SetUsimSupportOfSorCmci(v bool) {
	o.UsimSupportOfSorCmci = &v
}

func (o SorInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SteeringContainer) {
		toSerialize["steeringContainer"] = o.SteeringContainer
	}
	toSerialize["ackInd"] = o.AckInd
	if !IsNil(o.SorMacIausf) {
		toSerialize["sorMacIausf"] = o.SorMacIausf
	}
	if !IsNil(o.Countersor) {
		toSerialize["countersor"] = o.Countersor
	}
	toSerialize["provisioningTime"] = o.ProvisioningTime
	if !IsNil(o.SorTransparentContainer) {
		toSerialize["sorTransparentContainer"] = o.SorTransparentContainer
	}
	if !IsNil(o.SorCmci) {
		toSerialize["sorCmci"] = o.SorCmci
	}
	if !IsNil(o.StoreSorCmciInMe) {
		toSerialize["storeSorCmciInMe"] = o.StoreSorCmciInMe
	}
	if !IsNil(o.UsimSupportOfSorCmci) {
		toSerialize["usimSupportOfSorCmci"] = o.UsimSupportOfSorCmci
	}
	return toSerialize, nil
}

type NullableSorInfo struct {
	value *SorInfo
	isSet bool
}

func (v NullableSorInfo) Get() *SorInfo {
	return v.value
}

func (v *NullableSorInfo) Set(val *SorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSorInfo(val *SorInfo) *NullableSorInfo {
	return &NullableSorInfo{value: val, isSet: true}
}

func (v NullableSorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
