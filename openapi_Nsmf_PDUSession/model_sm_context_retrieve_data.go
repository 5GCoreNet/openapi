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

// checks if the SmContextRetrieveData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SmContextRetrieveData{}

// SmContextRetrieveData Data within Retrieve SM Context Request
type SmContextRetrieveData struct {
	TargetMmeCap         *MmeCapabilities `json:"targetMmeCap,omitempty"`
	SmContextType        *SmContextType   `json:"smContextType,omitempty"`
	ServingNetwork       *PlmnId          `json:"servingNetwork,omitempty"`
	NotToTransferEbiList []int32          `json:"notToTransferEbiList,omitempty"`
	RanUnchangedInd      *bool            `json:"ranUnchangedInd,omitempty"`
}

// NewSmContextRetrieveData instantiates a new SmContextRetrieveData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSmContextRetrieveData() *SmContextRetrieveData {
	this := SmContextRetrieveData{}
	var ranUnchangedInd bool = false
	this.RanUnchangedInd = &ranUnchangedInd
	return &this
}

// NewSmContextRetrieveDataWithDefaults instantiates a new SmContextRetrieveData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSmContextRetrieveDataWithDefaults() *SmContextRetrieveData {
	this := SmContextRetrieveData{}
	var ranUnchangedInd bool = false
	this.RanUnchangedInd = &ranUnchangedInd
	return &this
}

// GetTargetMmeCap returns the TargetMmeCap field value if set, zero value otherwise.
func (o *SmContextRetrieveData) GetTargetMmeCap() MmeCapabilities {
	if o == nil || IsNil(o.TargetMmeCap) {
		var ret MmeCapabilities
		return ret
	}
	return *o.TargetMmeCap
}

// GetTargetMmeCapOk returns a tuple with the TargetMmeCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrieveData) GetTargetMmeCapOk() (*MmeCapabilities, bool) {
	if o == nil || IsNil(o.TargetMmeCap) {
		return nil, false
	}
	return o.TargetMmeCap, true
}

// HasTargetMmeCap returns a boolean if a field has been set.
func (o *SmContextRetrieveData) HasTargetMmeCap() bool {
	if o != nil && !IsNil(o.TargetMmeCap) {
		return true
	}

	return false
}

// SetTargetMmeCap gets a reference to the given MmeCapabilities and assigns it to the TargetMmeCap field.
func (o *SmContextRetrieveData) SetTargetMmeCap(v MmeCapabilities) {
	o.TargetMmeCap = &v
}

// GetSmContextType returns the SmContextType field value if set, zero value otherwise.
func (o *SmContextRetrieveData) GetSmContextType() SmContextType {
	if o == nil || IsNil(o.SmContextType) {
		var ret SmContextType
		return ret
	}
	return *o.SmContextType
}

// GetSmContextTypeOk returns a tuple with the SmContextType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrieveData) GetSmContextTypeOk() (*SmContextType, bool) {
	if o == nil || IsNil(o.SmContextType) {
		return nil, false
	}
	return o.SmContextType, true
}

// HasSmContextType returns a boolean if a field has been set.
func (o *SmContextRetrieveData) HasSmContextType() bool {
	if o != nil && !IsNil(o.SmContextType) {
		return true
	}

	return false
}

// SetSmContextType gets a reference to the given SmContextType and assigns it to the SmContextType field.
func (o *SmContextRetrieveData) SetSmContextType(v SmContextType) {
	o.SmContextType = &v
}

// GetServingNetwork returns the ServingNetwork field value if set, zero value otherwise.
func (o *SmContextRetrieveData) GetServingNetwork() PlmnId {
	if o == nil || IsNil(o.ServingNetwork) {
		var ret PlmnId
		return ret
	}
	return *o.ServingNetwork
}

// GetServingNetworkOk returns a tuple with the ServingNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrieveData) GetServingNetworkOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.ServingNetwork) {
		return nil, false
	}
	return o.ServingNetwork, true
}

// HasServingNetwork returns a boolean if a field has been set.
func (o *SmContextRetrieveData) HasServingNetwork() bool {
	if o != nil && !IsNil(o.ServingNetwork) {
		return true
	}

	return false
}

// SetServingNetwork gets a reference to the given PlmnId and assigns it to the ServingNetwork field.
func (o *SmContextRetrieveData) SetServingNetwork(v PlmnId) {
	o.ServingNetwork = &v
}

// GetNotToTransferEbiList returns the NotToTransferEbiList field value if set, zero value otherwise.
func (o *SmContextRetrieveData) GetNotToTransferEbiList() []int32 {
	if o == nil || IsNil(o.NotToTransferEbiList) {
		var ret []int32
		return ret
	}
	return o.NotToTransferEbiList
}

// GetNotToTransferEbiListOk returns a tuple with the NotToTransferEbiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrieveData) GetNotToTransferEbiListOk() ([]int32, bool) {
	if o == nil || IsNil(o.NotToTransferEbiList) {
		return nil, false
	}
	return o.NotToTransferEbiList, true
}

// HasNotToTransferEbiList returns a boolean if a field has been set.
func (o *SmContextRetrieveData) HasNotToTransferEbiList() bool {
	if o != nil && !IsNil(o.NotToTransferEbiList) {
		return true
	}

	return false
}

// SetNotToTransferEbiList gets a reference to the given []int32 and assigns it to the NotToTransferEbiList field.
func (o *SmContextRetrieveData) SetNotToTransferEbiList(v []int32) {
	o.NotToTransferEbiList = v
}

// GetRanUnchangedInd returns the RanUnchangedInd field value if set, zero value otherwise.
func (o *SmContextRetrieveData) GetRanUnchangedInd() bool {
	if o == nil || IsNil(o.RanUnchangedInd) {
		var ret bool
		return ret
	}
	return *o.RanUnchangedInd
}

// GetRanUnchangedIndOk returns a tuple with the RanUnchangedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SmContextRetrieveData) GetRanUnchangedIndOk() (*bool, bool) {
	if o == nil || IsNil(o.RanUnchangedInd) {
		return nil, false
	}
	return o.RanUnchangedInd, true
}

// HasRanUnchangedInd returns a boolean if a field has been set.
func (o *SmContextRetrieveData) HasRanUnchangedInd() bool {
	if o != nil && !IsNil(o.RanUnchangedInd) {
		return true
	}

	return false
}

// SetRanUnchangedInd gets a reference to the given bool and assigns it to the RanUnchangedInd field.
func (o *SmContextRetrieveData) SetRanUnchangedInd(v bool) {
	o.RanUnchangedInd = &v
}

func (o SmContextRetrieveData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SmContextRetrieveData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TargetMmeCap) {
		toSerialize["targetMmeCap"] = o.TargetMmeCap
	}
	if !IsNil(o.SmContextType) {
		toSerialize["smContextType"] = o.SmContextType
	}
	if !IsNil(o.ServingNetwork) {
		toSerialize["servingNetwork"] = o.ServingNetwork
	}
	if !IsNil(o.NotToTransferEbiList) {
		toSerialize["notToTransferEbiList"] = o.NotToTransferEbiList
	}
	if !IsNil(o.RanUnchangedInd) {
		toSerialize["ranUnchangedInd"] = o.RanUnchangedInd
	}
	return toSerialize, nil
}

type NullableSmContextRetrieveData struct {
	value *SmContextRetrieveData
	isSet bool
}

func (v NullableSmContextRetrieveData) Get() *SmContextRetrieveData {
	return v.value
}

func (v *NullableSmContextRetrieveData) Set(val *SmContextRetrieveData) {
	v.value = val
	v.isSet = true
}

func (v NullableSmContextRetrieveData) IsSet() bool {
	return v.isSet
}

func (v *NullableSmContextRetrieveData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmContextRetrieveData(val *SmContextRetrieveData) *NullableSmContextRetrieveData {
	return &NullableSmContextRetrieveData{value: val, isSet: true}
}

func (v NullableSmContextRetrieveData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmContextRetrieveData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
