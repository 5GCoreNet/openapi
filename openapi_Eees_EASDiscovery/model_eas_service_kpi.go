/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
)

// checks if the EASServiceKPI type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EASServiceKPI{}

// EASServiceKPI Represents the EAS service KPI information.
type EASServiceKPI struct {
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxReqRate *int32 `json:"maxReqRate,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	MaxRespTime *int32 `json:"maxRespTime,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Avail *int32 `json:"avail,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AvlComp *int32 `json:"avlComp,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AvlGraComp *int32 `json:"avlGraComp,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AvlMem *int32 `json:"avlMem,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	AvlStrg *int32 `json:"avlStrg,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	ConnBand *string `json:"connBand,omitempty"`
}

// NewEASServiceKPI instantiates a new EASServiceKPI object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEASServiceKPI() *EASServiceKPI {
	this := EASServiceKPI{}
	return &this
}

// NewEASServiceKPIWithDefaults instantiates a new EASServiceKPI object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEASServiceKPIWithDefaults() *EASServiceKPI {
	this := EASServiceKPI{}
	return &this
}

// GetMaxReqRate returns the MaxReqRate field value if set, zero value otherwise.
func (o *EASServiceKPI) GetMaxReqRate() int32 {
	if o == nil || isNil(o.MaxReqRate) {
		var ret int32
		return ret
	}
	return *o.MaxReqRate
}

// GetMaxReqRateOk returns a tuple with the MaxReqRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASServiceKPI) GetMaxReqRateOk() (*int32, bool) {
	if o == nil || isNil(o.MaxReqRate) {
		return nil, false
	}
	return o.MaxReqRate, true
}

// HasMaxReqRate returns a boolean if a field has been set.
func (o *EASServiceKPI) HasMaxReqRate() bool {
	if o != nil && !isNil(o.MaxReqRate) {
		return true
	}

	return false
}

// SetMaxReqRate gets a reference to the given int32 and assigns it to the MaxReqRate field.
func (o *EASServiceKPI) SetMaxReqRate(v int32) {
	o.MaxReqRate = &v
}

// GetMaxRespTime returns the MaxRespTime field value if set, zero value otherwise.
func (o *EASServiceKPI) GetMaxRespTime() int32 {
	if o == nil || isNil(o.MaxRespTime) {
		var ret int32
		return ret
	}
	return *o.MaxRespTime
}

// GetMaxRespTimeOk returns a tuple with the MaxRespTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASServiceKPI) GetMaxRespTimeOk() (*int32, bool) {
	if o == nil || isNil(o.MaxRespTime) {
		return nil, false
	}
	return o.MaxRespTime, true
}

// HasMaxRespTime returns a boolean if a field has been set.
func (o *EASServiceKPI) HasMaxRespTime() bool {
	if o != nil && !isNil(o.MaxRespTime) {
		return true
	}

	return false
}

// SetMaxRespTime gets a reference to the given int32 and assigns it to the MaxRespTime field.
func (o *EASServiceKPI) SetMaxRespTime(v int32) {
	o.MaxRespTime = &v
}

// GetAvail returns the Avail field value if set, zero value otherwise.
func (o *EASServiceKPI) GetAvail() int32 {
	if o == nil || isNil(o.Avail) {
		var ret int32
		return ret
	}
	return *o.Avail
}

// GetAvailOk returns a tuple with the Avail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASServiceKPI) GetAvailOk() (*int32, bool) {
	if o == nil || isNil(o.Avail) {
		return nil, false
	}
	return o.Avail, true
}

// HasAvail returns a boolean if a field has been set.
func (o *EASServiceKPI) HasAvail() bool {
	if o != nil && !isNil(o.Avail) {
		return true
	}

	return false
}

// SetAvail gets a reference to the given int32 and assigns it to the Avail field.
func (o *EASServiceKPI) SetAvail(v int32) {
	o.Avail = &v
}

// GetAvlComp returns the AvlComp field value if set, zero value otherwise.
func (o *EASServiceKPI) GetAvlComp() int32 {
	if o == nil || isNil(o.AvlComp) {
		var ret int32
		return ret
	}
	return *o.AvlComp
}

// GetAvlCompOk returns a tuple with the AvlComp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASServiceKPI) GetAvlCompOk() (*int32, bool) {
	if o == nil || isNil(o.AvlComp) {
		return nil, false
	}
	return o.AvlComp, true
}

// HasAvlComp returns a boolean if a field has been set.
func (o *EASServiceKPI) HasAvlComp() bool {
	if o != nil && !isNil(o.AvlComp) {
		return true
	}

	return false
}

// SetAvlComp gets a reference to the given int32 and assigns it to the AvlComp field.
func (o *EASServiceKPI) SetAvlComp(v int32) {
	o.AvlComp = &v
}

// GetAvlGraComp returns the AvlGraComp field value if set, zero value otherwise.
func (o *EASServiceKPI) GetAvlGraComp() int32 {
	if o == nil || isNil(o.AvlGraComp) {
		var ret int32
		return ret
	}
	return *o.AvlGraComp
}

// GetAvlGraCompOk returns a tuple with the AvlGraComp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASServiceKPI) GetAvlGraCompOk() (*int32, bool) {
	if o == nil || isNil(o.AvlGraComp) {
		return nil, false
	}
	return o.AvlGraComp, true
}

// HasAvlGraComp returns a boolean if a field has been set.
func (o *EASServiceKPI) HasAvlGraComp() bool {
	if o != nil && !isNil(o.AvlGraComp) {
		return true
	}

	return false
}

// SetAvlGraComp gets a reference to the given int32 and assigns it to the AvlGraComp field.
func (o *EASServiceKPI) SetAvlGraComp(v int32) {
	o.AvlGraComp = &v
}

// GetAvlMem returns the AvlMem field value if set, zero value otherwise.
func (o *EASServiceKPI) GetAvlMem() int32 {
	if o == nil || isNil(o.AvlMem) {
		var ret int32
		return ret
	}
	return *o.AvlMem
}

// GetAvlMemOk returns a tuple with the AvlMem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASServiceKPI) GetAvlMemOk() (*int32, bool) {
	if o == nil || isNil(o.AvlMem) {
		return nil, false
	}
	return o.AvlMem, true
}

// HasAvlMem returns a boolean if a field has been set.
func (o *EASServiceKPI) HasAvlMem() bool {
	if o != nil && !isNil(o.AvlMem) {
		return true
	}

	return false
}

// SetAvlMem gets a reference to the given int32 and assigns it to the AvlMem field.
func (o *EASServiceKPI) SetAvlMem(v int32) {
	o.AvlMem = &v
}

// GetAvlStrg returns the AvlStrg field value if set, zero value otherwise.
func (o *EASServiceKPI) GetAvlStrg() int32 {
	if o == nil || isNil(o.AvlStrg) {
		var ret int32
		return ret
	}
	return *o.AvlStrg
}

// GetAvlStrgOk returns a tuple with the AvlStrg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASServiceKPI) GetAvlStrgOk() (*int32, bool) {
	if o == nil || isNil(o.AvlStrg) {
		return nil, false
	}
	return o.AvlStrg, true
}

// HasAvlStrg returns a boolean if a field has been set.
func (o *EASServiceKPI) HasAvlStrg() bool {
	if o != nil && !isNil(o.AvlStrg) {
		return true
	}

	return false
}

// SetAvlStrg gets a reference to the given int32 and assigns it to the AvlStrg field.
func (o *EASServiceKPI) SetAvlStrg(v int32) {
	o.AvlStrg = &v
}

// GetConnBand returns the ConnBand field value if set, zero value otherwise.
func (o *EASServiceKPI) GetConnBand() string {
	if o == nil || isNil(o.ConnBand) {
		var ret string
		return ret
	}
	return *o.ConnBand
}

// GetConnBandOk returns a tuple with the ConnBand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASServiceKPI) GetConnBandOk() (*string, bool) {
	if o == nil || isNil(o.ConnBand) {
		return nil, false
	}
	return o.ConnBand, true
}

// HasConnBand returns a boolean if a field has been set.
func (o *EASServiceKPI) HasConnBand() bool {
	if o != nil && !isNil(o.ConnBand) {
		return true
	}

	return false
}

// SetConnBand gets a reference to the given string and assigns it to the ConnBand field.
func (o *EASServiceKPI) SetConnBand(v string) {
	o.ConnBand = &v
}

func (o EASServiceKPI) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EASServiceKPI) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MaxReqRate) {
		toSerialize["maxReqRate"] = o.MaxReqRate
	}
	if !isNil(o.MaxRespTime) {
		toSerialize["maxRespTime"] = o.MaxRespTime
	}
	if !isNil(o.Avail) {
		toSerialize["avail"] = o.Avail
	}
	if !isNil(o.AvlComp) {
		toSerialize["avlComp"] = o.AvlComp
	}
	if !isNil(o.AvlGraComp) {
		toSerialize["avlGraComp"] = o.AvlGraComp
	}
	if !isNil(o.AvlMem) {
		toSerialize["avlMem"] = o.AvlMem
	}
	if !isNil(o.AvlStrg) {
		toSerialize["avlStrg"] = o.AvlStrg
	}
	if !isNil(o.ConnBand) {
		toSerialize["connBand"] = o.ConnBand
	}
	return toSerialize, nil
}

type NullableEASServiceKPI struct {
	value *EASServiceKPI
	isSet bool
}

func (v NullableEASServiceKPI) Get() *EASServiceKPI {
	return v.value
}

func (v *NullableEASServiceKPI) Set(val *EASServiceKPI) {
	v.value = val
	v.isSet = true
}

func (v NullableEASServiceKPI) IsSet() bool {
	return v.isSet
}

func (v *NullableEASServiceKPI) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASServiceKPI(val *EASServiceKPI) *NullableEASServiceKPI {
	return &NullableEASServiceKPI{value: val, isSet: true}
}

func (v NullableEASServiceKPI) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASServiceKPI) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


