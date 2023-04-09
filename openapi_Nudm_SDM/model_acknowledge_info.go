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

// checks if the AcknowledgeInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcknowledgeInfo{}

// AcknowledgeInfo struct for AcknowledgeInfo
type AcknowledgeInfo struct {
	// MAC value for protecting SOR procedure (SoR-MAC-IAUSF and SoR-XMAC-IUE).
	SorMacIue *string `json:"sorMacIue,omitempty"`
	// MAC value for protecting UPU procedure (UPU-MAC-IAUSF and UPU-MAC-IUE).
	UpuMacIue *string `json:"upuMacIue,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ProvisioningTime time.Time `json:"provisioningTime"`
	// string with format 'bytes' as defined in OpenAPI
	SorTransparentContainer *string `json:"sorTransparentContainer,omitempty"`
	UeNotReachable          *bool   `json:"ueNotReachable,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	UpuTransparentContainer *string `json:"upuTransparentContainer,omitempty"`
}

// NewAcknowledgeInfo instantiates a new AcknowledgeInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcknowledgeInfo(provisioningTime time.Time) *AcknowledgeInfo {
	this := AcknowledgeInfo{}
	this.ProvisioningTime = provisioningTime
	var ueNotReachable bool = false
	this.UeNotReachable = &ueNotReachable
	return &this
}

// NewAcknowledgeInfoWithDefaults instantiates a new AcknowledgeInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcknowledgeInfoWithDefaults() *AcknowledgeInfo {
	this := AcknowledgeInfo{}
	var ueNotReachable bool = false
	this.UeNotReachable = &ueNotReachable
	return &this
}

// GetSorMacIue returns the SorMacIue field value if set, zero value otherwise.
func (o *AcknowledgeInfo) GetSorMacIue() string {
	if o == nil || IsNil(o.SorMacIue) {
		var ret string
		return ret
	}
	return *o.SorMacIue
}

// GetSorMacIueOk returns a tuple with the SorMacIue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcknowledgeInfo) GetSorMacIueOk() (*string, bool) {
	if o == nil || IsNil(o.SorMacIue) {
		return nil, false
	}
	return o.SorMacIue, true
}

// HasSorMacIue returns a boolean if a field has been set.
func (o *AcknowledgeInfo) HasSorMacIue() bool {
	if o != nil && !IsNil(o.SorMacIue) {
		return true
	}

	return false
}

// SetSorMacIue gets a reference to the given string and assigns it to the SorMacIue field.
func (o *AcknowledgeInfo) SetSorMacIue(v string) {
	o.SorMacIue = &v
}

// GetUpuMacIue returns the UpuMacIue field value if set, zero value otherwise.
func (o *AcknowledgeInfo) GetUpuMacIue() string {
	if o == nil || IsNil(o.UpuMacIue) {
		var ret string
		return ret
	}
	return *o.UpuMacIue
}

// GetUpuMacIueOk returns a tuple with the UpuMacIue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcknowledgeInfo) GetUpuMacIueOk() (*string, bool) {
	if o == nil || IsNil(o.UpuMacIue) {
		return nil, false
	}
	return o.UpuMacIue, true
}

// HasUpuMacIue returns a boolean if a field has been set.
func (o *AcknowledgeInfo) HasUpuMacIue() bool {
	if o != nil && !IsNil(o.UpuMacIue) {
		return true
	}

	return false
}

// SetUpuMacIue gets a reference to the given string and assigns it to the UpuMacIue field.
func (o *AcknowledgeInfo) SetUpuMacIue(v string) {
	o.UpuMacIue = &v
}

// GetProvisioningTime returns the ProvisioningTime field value
func (o *AcknowledgeInfo) GetProvisioningTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value
// and a boolean to check if the value has been set.
func (o *AcknowledgeInfo) GetProvisioningTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningTime, true
}

// SetProvisioningTime sets field value
func (o *AcknowledgeInfo) SetProvisioningTime(v time.Time) {
	o.ProvisioningTime = v
}

// GetSorTransparentContainer returns the SorTransparentContainer field value if set, zero value otherwise.
func (o *AcknowledgeInfo) GetSorTransparentContainer() string {
	if o == nil || IsNil(o.SorTransparentContainer) {
		var ret string
		return ret
	}
	return *o.SorTransparentContainer
}

// GetSorTransparentContainerOk returns a tuple with the SorTransparentContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcknowledgeInfo) GetSorTransparentContainerOk() (*string, bool) {
	if o == nil || IsNil(o.SorTransparentContainer) {
		return nil, false
	}
	return o.SorTransparentContainer, true
}

// HasSorTransparentContainer returns a boolean if a field has been set.
func (o *AcknowledgeInfo) HasSorTransparentContainer() bool {
	if o != nil && !IsNil(o.SorTransparentContainer) {
		return true
	}

	return false
}

// SetSorTransparentContainer gets a reference to the given string and assigns it to the SorTransparentContainer field.
func (o *AcknowledgeInfo) SetSorTransparentContainer(v string) {
	o.SorTransparentContainer = &v
}

// GetUeNotReachable returns the UeNotReachable field value if set, zero value otherwise.
func (o *AcknowledgeInfo) GetUeNotReachable() bool {
	if o == nil || IsNil(o.UeNotReachable) {
		var ret bool
		return ret
	}
	return *o.UeNotReachable
}

// GetUeNotReachableOk returns a tuple with the UeNotReachable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcknowledgeInfo) GetUeNotReachableOk() (*bool, bool) {
	if o == nil || IsNil(o.UeNotReachable) {
		return nil, false
	}
	return o.UeNotReachable, true
}

// HasUeNotReachable returns a boolean if a field has been set.
func (o *AcknowledgeInfo) HasUeNotReachable() bool {
	if o != nil && !IsNil(o.UeNotReachable) {
		return true
	}

	return false
}

// SetUeNotReachable gets a reference to the given bool and assigns it to the UeNotReachable field.
func (o *AcknowledgeInfo) SetUeNotReachable(v bool) {
	o.UeNotReachable = &v
}

// GetUpuTransparentContainer returns the UpuTransparentContainer field value if set, zero value otherwise.
func (o *AcknowledgeInfo) GetUpuTransparentContainer() string {
	if o == nil || IsNil(o.UpuTransparentContainer) {
		var ret string
		return ret
	}
	return *o.UpuTransparentContainer
}

// GetUpuTransparentContainerOk returns a tuple with the UpuTransparentContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcknowledgeInfo) GetUpuTransparentContainerOk() (*string, bool) {
	if o == nil || IsNil(o.UpuTransparentContainer) {
		return nil, false
	}
	return o.UpuTransparentContainer, true
}

// HasUpuTransparentContainer returns a boolean if a field has been set.
func (o *AcknowledgeInfo) HasUpuTransparentContainer() bool {
	if o != nil && !IsNil(o.UpuTransparentContainer) {
		return true
	}

	return false
}

// SetUpuTransparentContainer gets a reference to the given string and assigns it to the UpuTransparentContainer field.
func (o *AcknowledgeInfo) SetUpuTransparentContainer(v string) {
	o.UpuTransparentContainer = &v
}

func (o AcknowledgeInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcknowledgeInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SorMacIue) {
		toSerialize["sorMacIue"] = o.SorMacIue
	}
	if !IsNil(o.UpuMacIue) {
		toSerialize["upuMacIue"] = o.UpuMacIue
	}
	toSerialize["provisioningTime"] = o.ProvisioningTime
	if !IsNil(o.SorTransparentContainer) {
		toSerialize["sorTransparentContainer"] = o.SorTransparentContainer
	}
	if !IsNil(o.UeNotReachable) {
		toSerialize["ueNotReachable"] = o.UeNotReachable
	}
	if !IsNil(o.UpuTransparentContainer) {
		toSerialize["upuTransparentContainer"] = o.UpuTransparentContainer
	}
	return toSerialize, nil
}

type NullableAcknowledgeInfo struct {
	value *AcknowledgeInfo
	isSet bool
}

func (v NullableAcknowledgeInfo) Get() *AcknowledgeInfo {
	return v.value
}

func (v *NullableAcknowledgeInfo) Set(val *AcknowledgeInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAcknowledgeInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAcknowledgeInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcknowledgeInfo(val *AcknowledgeInfo) *NullableAcknowledgeInfo {
	return &NullableAcknowledgeInfo{value: val, isSet: true}
}

func (v NullableAcknowledgeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcknowledgeInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
