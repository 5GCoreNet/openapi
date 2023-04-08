/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the StatusInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusInfo{}

// StatusInfo Status of SM context or of PDU session
type StatusInfo struct {
	ResourceStatus ResourceStatus `json:"resourceStatus"`
	Cause *Cause `json:"cause,omitempty"`
	CnAssistedRanPara *CnAssistedRanPara `json:"cnAssistedRanPara,omitempty"`
	AnType *AccessType `json:"anType,omitempty"`
}

// NewStatusInfo instantiates a new StatusInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusInfo(resourceStatus ResourceStatus) *StatusInfo {
	this := StatusInfo{}
	this.ResourceStatus = resourceStatus
	return &this
}

// NewStatusInfoWithDefaults instantiates a new StatusInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusInfoWithDefaults() *StatusInfo {
	this := StatusInfo{}
	return &this
}

// GetResourceStatus returns the ResourceStatus field value
func (o *StatusInfo) GetResourceStatus() ResourceStatus {
	if o == nil {
		var ret ResourceStatus
		return ret
	}

	return o.ResourceStatus
}

// GetResourceStatusOk returns a tuple with the ResourceStatus field value
// and a boolean to check if the value has been set.
func (o *StatusInfo) GetResourceStatusOk() (*ResourceStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceStatus, true
}

// SetResourceStatus sets field value
func (o *StatusInfo) SetResourceStatus(v ResourceStatus) {
	o.ResourceStatus = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *StatusInfo) GetCause() Cause {
	if o == nil || isNil(o.Cause) {
		var ret Cause
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusInfo) GetCauseOk() (*Cause, bool) {
	if o == nil || isNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *StatusInfo) HasCause() bool {
	if o != nil && !isNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given Cause and assigns it to the Cause field.
func (o *StatusInfo) SetCause(v Cause) {
	o.Cause = &v
}

// GetCnAssistedRanPara returns the CnAssistedRanPara field value if set, zero value otherwise.
func (o *StatusInfo) GetCnAssistedRanPara() CnAssistedRanPara {
	if o == nil || isNil(o.CnAssistedRanPara) {
		var ret CnAssistedRanPara
		return ret
	}
	return *o.CnAssistedRanPara
}

// GetCnAssistedRanParaOk returns a tuple with the CnAssistedRanPara field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusInfo) GetCnAssistedRanParaOk() (*CnAssistedRanPara, bool) {
	if o == nil || isNil(o.CnAssistedRanPara) {
		return nil, false
	}
	return o.CnAssistedRanPara, true
}

// HasCnAssistedRanPara returns a boolean if a field has been set.
func (o *StatusInfo) HasCnAssistedRanPara() bool {
	if o != nil && !isNil(o.CnAssistedRanPara) {
		return true
	}

	return false
}

// SetCnAssistedRanPara gets a reference to the given CnAssistedRanPara and assigns it to the CnAssistedRanPara field.
func (o *StatusInfo) SetCnAssistedRanPara(v CnAssistedRanPara) {
	o.CnAssistedRanPara = &v
}

// GetAnType returns the AnType field value if set, zero value otherwise.
func (o *StatusInfo) GetAnType() AccessType {
	if o == nil || isNil(o.AnType) {
		var ret AccessType
		return ret
	}
	return *o.AnType
}

// GetAnTypeOk returns a tuple with the AnType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusInfo) GetAnTypeOk() (*AccessType, bool) {
	if o == nil || isNil(o.AnType) {
		return nil, false
	}
	return o.AnType, true
}

// HasAnType returns a boolean if a field has been set.
func (o *StatusInfo) HasAnType() bool {
	if o != nil && !isNil(o.AnType) {
		return true
	}

	return false
}

// SetAnType gets a reference to the given AccessType and assigns it to the AnType field.
func (o *StatusInfo) SetAnType(v AccessType) {
	o.AnType = &v
}

func (o StatusInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceStatus"] = o.ResourceStatus
	if !isNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !isNil(o.CnAssistedRanPara) {
		toSerialize["cnAssistedRanPara"] = o.CnAssistedRanPara
	}
	if !isNil(o.AnType) {
		toSerialize["anType"] = o.AnType
	}
	return toSerialize, nil
}

type NullableStatusInfo struct {
	value *StatusInfo
	isSet bool
}

func (v NullableStatusInfo) Get() *StatusInfo {
	return v.value
}

func (v *NullableStatusInfo) Set(val *StatusInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusInfo(val *StatusInfo) *NullableStatusInfo {
	return &NullableStatusInfo{value: val, isSet: true}
}

func (v NullableStatusInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


