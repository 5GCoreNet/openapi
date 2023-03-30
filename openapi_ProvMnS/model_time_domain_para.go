/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the TimeDomainPara type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeDomainPara{}

// TimeDomainPara struct for TimeDomainPara
type TimeDomainPara struct {
	DlULSwitchingPeriod1 *string `json:"dlULSwitchingPeriod1,omitempty"`
	SymbolOffsetOfReferencePoint1 *int32 `json:"symbolOffsetOfReferencePoint1,omitempty"`
	DlULSwitchingPeriod2 *string `json:"dlULSwitchingPeriod2,omitempty"`
	SymbolOffsetOfReferencePoint2 *int32 `json:"symbolOffsetOfReferencePoint2,omitempty"`
	TotalnrofSetIdofRS1 *int32 `json:"totalnrofSetIdofRS1,omitempty"`
	TotalnrofSetIdofRS2 *int32 `json:"totalnrofSetIdofRS2,omitempty"`
	NrofConsecutiveRIMRS1 *int32 `json:"nrofConsecutiveRIMRS1,omitempty"`
	NrofConsecutiveRIMRS2 *int32 `json:"nrofConsecutiveRIMRS2,omitempty"`
	ConsecutiveRIMRS1List []int32 `json:"consecutiveRIMRS1List,omitempty"`
	ConsecutiveRIMRS2List []int32 `json:"consecutiveRIMRS2List,omitempty"`
	EnablenearfarIndicationRS1 *string `json:"enablenearfarIndicationRS1,omitempty"`
	EnablenearfarIndicationRS2 *string `json:"enablenearfarIndicationRS2,omitempty"`
}

// NewTimeDomainPara instantiates a new TimeDomainPara object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeDomainPara() *TimeDomainPara {
	this := TimeDomainPara{}
	return &this
}

// NewTimeDomainParaWithDefaults instantiates a new TimeDomainPara object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeDomainParaWithDefaults() *TimeDomainPara {
	this := TimeDomainPara{}
	return &this
}

// GetDlULSwitchingPeriod1 returns the DlULSwitchingPeriod1 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetDlULSwitchingPeriod1() string {
	if o == nil || IsNil(o.DlULSwitchingPeriod1) {
		var ret string
		return ret
	}
	return *o.DlULSwitchingPeriod1
}

// GetDlULSwitchingPeriod1Ok returns a tuple with the DlULSwitchingPeriod1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetDlULSwitchingPeriod1Ok() (*string, bool) {
	if o == nil || IsNil(o.DlULSwitchingPeriod1) {
		return nil, false
	}
	return o.DlULSwitchingPeriod1, true
}

// HasDlULSwitchingPeriod1 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasDlULSwitchingPeriod1() bool {
	if o != nil && !IsNil(o.DlULSwitchingPeriod1) {
		return true
	}

	return false
}

// SetDlULSwitchingPeriod1 gets a reference to the given string and assigns it to the DlULSwitchingPeriod1 field.
func (o *TimeDomainPara) SetDlULSwitchingPeriod1(v string) {
	o.DlULSwitchingPeriod1 = &v
}

// GetSymbolOffsetOfReferencePoint1 returns the SymbolOffsetOfReferencePoint1 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetSymbolOffsetOfReferencePoint1() int32 {
	if o == nil || IsNil(o.SymbolOffsetOfReferencePoint1) {
		var ret int32
		return ret
	}
	return *o.SymbolOffsetOfReferencePoint1
}

// GetSymbolOffsetOfReferencePoint1Ok returns a tuple with the SymbolOffsetOfReferencePoint1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetSymbolOffsetOfReferencePoint1Ok() (*int32, bool) {
	if o == nil || IsNil(o.SymbolOffsetOfReferencePoint1) {
		return nil, false
	}
	return o.SymbolOffsetOfReferencePoint1, true
}

// HasSymbolOffsetOfReferencePoint1 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasSymbolOffsetOfReferencePoint1() bool {
	if o != nil && !IsNil(o.SymbolOffsetOfReferencePoint1) {
		return true
	}

	return false
}

// SetSymbolOffsetOfReferencePoint1 gets a reference to the given int32 and assigns it to the SymbolOffsetOfReferencePoint1 field.
func (o *TimeDomainPara) SetSymbolOffsetOfReferencePoint1(v int32) {
	o.SymbolOffsetOfReferencePoint1 = &v
}

// GetDlULSwitchingPeriod2 returns the DlULSwitchingPeriod2 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetDlULSwitchingPeriod2() string {
	if o == nil || IsNil(o.DlULSwitchingPeriod2) {
		var ret string
		return ret
	}
	return *o.DlULSwitchingPeriod2
}

// GetDlULSwitchingPeriod2Ok returns a tuple with the DlULSwitchingPeriod2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetDlULSwitchingPeriod2Ok() (*string, bool) {
	if o == nil || IsNil(o.DlULSwitchingPeriod2) {
		return nil, false
	}
	return o.DlULSwitchingPeriod2, true
}

// HasDlULSwitchingPeriod2 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasDlULSwitchingPeriod2() bool {
	if o != nil && !IsNil(o.DlULSwitchingPeriod2) {
		return true
	}

	return false
}

// SetDlULSwitchingPeriod2 gets a reference to the given string and assigns it to the DlULSwitchingPeriod2 field.
func (o *TimeDomainPara) SetDlULSwitchingPeriod2(v string) {
	o.DlULSwitchingPeriod2 = &v
}

// GetSymbolOffsetOfReferencePoint2 returns the SymbolOffsetOfReferencePoint2 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetSymbolOffsetOfReferencePoint2() int32 {
	if o == nil || IsNil(o.SymbolOffsetOfReferencePoint2) {
		var ret int32
		return ret
	}
	return *o.SymbolOffsetOfReferencePoint2
}

// GetSymbolOffsetOfReferencePoint2Ok returns a tuple with the SymbolOffsetOfReferencePoint2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetSymbolOffsetOfReferencePoint2Ok() (*int32, bool) {
	if o == nil || IsNil(o.SymbolOffsetOfReferencePoint2) {
		return nil, false
	}
	return o.SymbolOffsetOfReferencePoint2, true
}

// HasSymbolOffsetOfReferencePoint2 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasSymbolOffsetOfReferencePoint2() bool {
	if o != nil && !IsNil(o.SymbolOffsetOfReferencePoint2) {
		return true
	}

	return false
}

// SetSymbolOffsetOfReferencePoint2 gets a reference to the given int32 and assigns it to the SymbolOffsetOfReferencePoint2 field.
func (o *TimeDomainPara) SetSymbolOffsetOfReferencePoint2(v int32) {
	o.SymbolOffsetOfReferencePoint2 = &v
}

// GetTotalnrofSetIdofRS1 returns the TotalnrofSetIdofRS1 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetTotalnrofSetIdofRS1() int32 {
	if o == nil || IsNil(o.TotalnrofSetIdofRS1) {
		var ret int32
		return ret
	}
	return *o.TotalnrofSetIdofRS1
}

// GetTotalnrofSetIdofRS1Ok returns a tuple with the TotalnrofSetIdofRS1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetTotalnrofSetIdofRS1Ok() (*int32, bool) {
	if o == nil || IsNil(o.TotalnrofSetIdofRS1) {
		return nil, false
	}
	return o.TotalnrofSetIdofRS1, true
}

// HasTotalnrofSetIdofRS1 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasTotalnrofSetIdofRS1() bool {
	if o != nil && !IsNil(o.TotalnrofSetIdofRS1) {
		return true
	}

	return false
}

// SetTotalnrofSetIdofRS1 gets a reference to the given int32 and assigns it to the TotalnrofSetIdofRS1 field.
func (o *TimeDomainPara) SetTotalnrofSetIdofRS1(v int32) {
	o.TotalnrofSetIdofRS1 = &v
}

// GetTotalnrofSetIdofRS2 returns the TotalnrofSetIdofRS2 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetTotalnrofSetIdofRS2() int32 {
	if o == nil || IsNil(o.TotalnrofSetIdofRS2) {
		var ret int32
		return ret
	}
	return *o.TotalnrofSetIdofRS2
}

// GetTotalnrofSetIdofRS2Ok returns a tuple with the TotalnrofSetIdofRS2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetTotalnrofSetIdofRS2Ok() (*int32, bool) {
	if o == nil || IsNil(o.TotalnrofSetIdofRS2) {
		return nil, false
	}
	return o.TotalnrofSetIdofRS2, true
}

// HasTotalnrofSetIdofRS2 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasTotalnrofSetIdofRS2() bool {
	if o != nil && !IsNil(o.TotalnrofSetIdofRS2) {
		return true
	}

	return false
}

// SetTotalnrofSetIdofRS2 gets a reference to the given int32 and assigns it to the TotalnrofSetIdofRS2 field.
func (o *TimeDomainPara) SetTotalnrofSetIdofRS2(v int32) {
	o.TotalnrofSetIdofRS2 = &v
}

// GetNrofConsecutiveRIMRS1 returns the NrofConsecutiveRIMRS1 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetNrofConsecutiveRIMRS1() int32 {
	if o == nil || IsNil(o.NrofConsecutiveRIMRS1) {
		var ret int32
		return ret
	}
	return *o.NrofConsecutiveRIMRS1
}

// GetNrofConsecutiveRIMRS1Ok returns a tuple with the NrofConsecutiveRIMRS1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetNrofConsecutiveRIMRS1Ok() (*int32, bool) {
	if o == nil || IsNil(o.NrofConsecutiveRIMRS1) {
		return nil, false
	}
	return o.NrofConsecutiveRIMRS1, true
}

// HasNrofConsecutiveRIMRS1 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasNrofConsecutiveRIMRS1() bool {
	if o != nil && !IsNil(o.NrofConsecutiveRIMRS1) {
		return true
	}

	return false
}

// SetNrofConsecutiveRIMRS1 gets a reference to the given int32 and assigns it to the NrofConsecutiveRIMRS1 field.
func (o *TimeDomainPara) SetNrofConsecutiveRIMRS1(v int32) {
	o.NrofConsecutiveRIMRS1 = &v
}

// GetNrofConsecutiveRIMRS2 returns the NrofConsecutiveRIMRS2 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetNrofConsecutiveRIMRS2() int32 {
	if o == nil || IsNil(o.NrofConsecutiveRIMRS2) {
		var ret int32
		return ret
	}
	return *o.NrofConsecutiveRIMRS2
}

// GetNrofConsecutiveRIMRS2Ok returns a tuple with the NrofConsecutiveRIMRS2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetNrofConsecutiveRIMRS2Ok() (*int32, bool) {
	if o == nil || IsNil(o.NrofConsecutiveRIMRS2) {
		return nil, false
	}
	return o.NrofConsecutiveRIMRS2, true
}

// HasNrofConsecutiveRIMRS2 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasNrofConsecutiveRIMRS2() bool {
	if o != nil && !IsNil(o.NrofConsecutiveRIMRS2) {
		return true
	}

	return false
}

// SetNrofConsecutiveRIMRS2 gets a reference to the given int32 and assigns it to the NrofConsecutiveRIMRS2 field.
func (o *TimeDomainPara) SetNrofConsecutiveRIMRS2(v int32) {
	o.NrofConsecutiveRIMRS2 = &v
}

// GetConsecutiveRIMRS1List returns the ConsecutiveRIMRS1List field value if set, zero value otherwise.
func (o *TimeDomainPara) GetConsecutiveRIMRS1List() []int32 {
	if o == nil || IsNil(o.ConsecutiveRIMRS1List) {
		var ret []int32
		return ret
	}
	return o.ConsecutiveRIMRS1List
}

// GetConsecutiveRIMRS1ListOk returns a tuple with the ConsecutiveRIMRS1List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetConsecutiveRIMRS1ListOk() ([]int32, bool) {
	if o == nil || IsNil(o.ConsecutiveRIMRS1List) {
		return nil, false
	}
	return o.ConsecutiveRIMRS1List, true
}

// HasConsecutiveRIMRS1List returns a boolean if a field has been set.
func (o *TimeDomainPara) HasConsecutiveRIMRS1List() bool {
	if o != nil && !IsNil(o.ConsecutiveRIMRS1List) {
		return true
	}

	return false
}

// SetConsecutiveRIMRS1List gets a reference to the given []int32 and assigns it to the ConsecutiveRIMRS1List field.
func (o *TimeDomainPara) SetConsecutiveRIMRS1List(v []int32) {
	o.ConsecutiveRIMRS1List = v
}

// GetConsecutiveRIMRS2List returns the ConsecutiveRIMRS2List field value if set, zero value otherwise.
func (o *TimeDomainPara) GetConsecutiveRIMRS2List() []int32 {
	if o == nil || IsNil(o.ConsecutiveRIMRS2List) {
		var ret []int32
		return ret
	}
	return o.ConsecutiveRIMRS2List
}

// GetConsecutiveRIMRS2ListOk returns a tuple with the ConsecutiveRIMRS2List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetConsecutiveRIMRS2ListOk() ([]int32, bool) {
	if o == nil || IsNil(o.ConsecutiveRIMRS2List) {
		return nil, false
	}
	return o.ConsecutiveRIMRS2List, true
}

// HasConsecutiveRIMRS2List returns a boolean if a field has been set.
func (o *TimeDomainPara) HasConsecutiveRIMRS2List() bool {
	if o != nil && !IsNil(o.ConsecutiveRIMRS2List) {
		return true
	}

	return false
}

// SetConsecutiveRIMRS2List gets a reference to the given []int32 and assigns it to the ConsecutiveRIMRS2List field.
func (o *TimeDomainPara) SetConsecutiveRIMRS2List(v []int32) {
	o.ConsecutiveRIMRS2List = v
}

// GetEnablenearfarIndicationRS1 returns the EnablenearfarIndicationRS1 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetEnablenearfarIndicationRS1() string {
	if o == nil || IsNil(o.EnablenearfarIndicationRS1) {
		var ret string
		return ret
	}
	return *o.EnablenearfarIndicationRS1
}

// GetEnablenearfarIndicationRS1Ok returns a tuple with the EnablenearfarIndicationRS1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetEnablenearfarIndicationRS1Ok() (*string, bool) {
	if o == nil || IsNil(o.EnablenearfarIndicationRS1) {
		return nil, false
	}
	return o.EnablenearfarIndicationRS1, true
}

// HasEnablenearfarIndicationRS1 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasEnablenearfarIndicationRS1() bool {
	if o != nil && !IsNil(o.EnablenearfarIndicationRS1) {
		return true
	}

	return false
}

// SetEnablenearfarIndicationRS1 gets a reference to the given string and assigns it to the EnablenearfarIndicationRS1 field.
func (o *TimeDomainPara) SetEnablenearfarIndicationRS1(v string) {
	o.EnablenearfarIndicationRS1 = &v
}

// GetEnablenearfarIndicationRS2 returns the EnablenearfarIndicationRS2 field value if set, zero value otherwise.
func (o *TimeDomainPara) GetEnablenearfarIndicationRS2() string {
	if o == nil || IsNil(o.EnablenearfarIndicationRS2) {
		var ret string
		return ret
	}
	return *o.EnablenearfarIndicationRS2
}

// GetEnablenearfarIndicationRS2Ok returns a tuple with the EnablenearfarIndicationRS2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TimeDomainPara) GetEnablenearfarIndicationRS2Ok() (*string, bool) {
	if o == nil || IsNil(o.EnablenearfarIndicationRS2) {
		return nil, false
	}
	return o.EnablenearfarIndicationRS2, true
}

// HasEnablenearfarIndicationRS2 returns a boolean if a field has been set.
func (o *TimeDomainPara) HasEnablenearfarIndicationRS2() bool {
	if o != nil && !IsNil(o.EnablenearfarIndicationRS2) {
		return true
	}

	return false
}

// SetEnablenearfarIndicationRS2 gets a reference to the given string and assigns it to the EnablenearfarIndicationRS2 field.
func (o *TimeDomainPara) SetEnablenearfarIndicationRS2(v string) {
	o.EnablenearfarIndicationRS2 = &v
}

func (o TimeDomainPara) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeDomainPara) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DlULSwitchingPeriod1) {
		toSerialize["dlULSwitchingPeriod1"] = o.DlULSwitchingPeriod1
	}
	if !IsNil(o.SymbolOffsetOfReferencePoint1) {
		toSerialize["symbolOffsetOfReferencePoint1"] = o.SymbolOffsetOfReferencePoint1
	}
	if !IsNil(o.DlULSwitchingPeriod2) {
		toSerialize["dlULSwitchingPeriod2"] = o.DlULSwitchingPeriod2
	}
	if !IsNil(o.SymbolOffsetOfReferencePoint2) {
		toSerialize["symbolOffsetOfReferencePoint2"] = o.SymbolOffsetOfReferencePoint2
	}
	if !IsNil(o.TotalnrofSetIdofRS1) {
		toSerialize["totalnrofSetIdofRS1"] = o.TotalnrofSetIdofRS1
	}
	if !IsNil(o.TotalnrofSetIdofRS2) {
		toSerialize["totalnrofSetIdofRS2"] = o.TotalnrofSetIdofRS2
	}
	if !IsNil(o.NrofConsecutiveRIMRS1) {
		toSerialize["nrofConsecutiveRIMRS1"] = o.NrofConsecutiveRIMRS1
	}
	if !IsNil(o.NrofConsecutiveRIMRS2) {
		toSerialize["nrofConsecutiveRIMRS2"] = o.NrofConsecutiveRIMRS2
	}
	if !IsNil(o.ConsecutiveRIMRS1List) {
		toSerialize["consecutiveRIMRS1List"] = o.ConsecutiveRIMRS1List
	}
	if !IsNil(o.ConsecutiveRIMRS2List) {
		toSerialize["consecutiveRIMRS2List"] = o.ConsecutiveRIMRS2List
	}
	if !IsNil(o.EnablenearfarIndicationRS1) {
		toSerialize["enablenearfarIndicationRS1"] = o.EnablenearfarIndicationRS1
	}
	if !IsNil(o.EnablenearfarIndicationRS2) {
		toSerialize["enablenearfarIndicationRS2"] = o.EnablenearfarIndicationRS2
	}
	return toSerialize, nil
}

type NullableTimeDomainPara struct {
	value *TimeDomainPara
	isSet bool
}

func (v NullableTimeDomainPara) Get() *TimeDomainPara {
	return v.value
}

func (v *NullableTimeDomainPara) Set(val *TimeDomainPara) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeDomainPara) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeDomainPara) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeDomainPara(val *TimeDomainPara) *NullableTimeDomainPara {
	return &NullableTimeDomainPara{value: val, isSet: true}
}

func (v NullableTimeDomainPara) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeDomainPara) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


