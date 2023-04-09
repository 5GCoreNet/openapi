/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the N2InfoContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N2InfoContainer{}

// N2InfoContainer N2 information container
type N2InfoContainer struct {
	N2InformationClass N2InformationClass `json:"n2InformationClass"`
	SmInfo             *N2SmInformation   `json:"smInfo,omitempty"`
	RanInfo            *N2RanInformation  `json:"ranInfo,omitempty"`
	NrppaInfo          *NrppaInformation  `json:"nrppaInfo,omitempty"`
	PwsInfo            *PwsInformation    `json:"pwsInfo,omitempty"`
	V2xInfo            *V2xInformation    `json:"v2xInfo,omitempty"`
	ProseInfo          *ProSeInformation  `json:"proseInfo,omitempty"`
}

// NewN2InfoContainer instantiates a new N2InfoContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN2InfoContainer(n2InformationClass N2InformationClass) *N2InfoContainer {
	this := N2InfoContainer{}
	this.N2InformationClass = n2InformationClass
	return &this
}

// NewN2InfoContainerWithDefaults instantiates a new N2InfoContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN2InfoContainerWithDefaults() *N2InfoContainer {
	this := N2InfoContainer{}
	return &this
}

// GetN2InformationClass returns the N2InformationClass field value
func (o *N2InfoContainer) GetN2InformationClass() N2InformationClass {
	if o == nil {
		var ret N2InformationClass
		return ret
	}

	return o.N2InformationClass
}

// GetN2InformationClassOk returns a tuple with the N2InformationClass field value
// and a boolean to check if the value has been set.
func (o *N2InfoContainer) GetN2InformationClassOk() (*N2InformationClass, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N2InformationClass, true
}

// SetN2InformationClass sets field value
func (o *N2InfoContainer) SetN2InformationClass(v N2InformationClass) {
	o.N2InformationClass = v
}

// GetSmInfo returns the SmInfo field value if set, zero value otherwise.
func (o *N2InfoContainer) GetSmInfo() N2SmInformation {
	if o == nil || IsNil(o.SmInfo) {
		var ret N2SmInformation
		return ret
	}
	return *o.SmInfo
}

// GetSmInfoOk returns a tuple with the SmInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InfoContainer) GetSmInfoOk() (*N2SmInformation, bool) {
	if o == nil || IsNil(o.SmInfo) {
		return nil, false
	}
	return o.SmInfo, true
}

// HasSmInfo returns a boolean if a field has been set.
func (o *N2InfoContainer) HasSmInfo() bool {
	if o != nil && !IsNil(o.SmInfo) {
		return true
	}

	return false
}

// SetSmInfo gets a reference to the given N2SmInformation and assigns it to the SmInfo field.
func (o *N2InfoContainer) SetSmInfo(v N2SmInformation) {
	o.SmInfo = &v
}

// GetRanInfo returns the RanInfo field value if set, zero value otherwise.
func (o *N2InfoContainer) GetRanInfo() N2RanInformation {
	if o == nil || IsNil(o.RanInfo) {
		var ret N2RanInformation
		return ret
	}
	return *o.RanInfo
}

// GetRanInfoOk returns a tuple with the RanInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InfoContainer) GetRanInfoOk() (*N2RanInformation, bool) {
	if o == nil || IsNil(o.RanInfo) {
		return nil, false
	}
	return o.RanInfo, true
}

// HasRanInfo returns a boolean if a field has been set.
func (o *N2InfoContainer) HasRanInfo() bool {
	if o != nil && !IsNil(o.RanInfo) {
		return true
	}

	return false
}

// SetRanInfo gets a reference to the given N2RanInformation and assigns it to the RanInfo field.
func (o *N2InfoContainer) SetRanInfo(v N2RanInformation) {
	o.RanInfo = &v
}

// GetNrppaInfo returns the NrppaInfo field value if set, zero value otherwise.
func (o *N2InfoContainer) GetNrppaInfo() NrppaInformation {
	if o == nil || IsNil(o.NrppaInfo) {
		var ret NrppaInformation
		return ret
	}
	return *o.NrppaInfo
}

// GetNrppaInfoOk returns a tuple with the NrppaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InfoContainer) GetNrppaInfoOk() (*NrppaInformation, bool) {
	if o == nil || IsNil(o.NrppaInfo) {
		return nil, false
	}
	return o.NrppaInfo, true
}

// HasNrppaInfo returns a boolean if a field has been set.
func (o *N2InfoContainer) HasNrppaInfo() bool {
	if o != nil && !IsNil(o.NrppaInfo) {
		return true
	}

	return false
}

// SetNrppaInfo gets a reference to the given NrppaInformation and assigns it to the NrppaInfo field.
func (o *N2InfoContainer) SetNrppaInfo(v NrppaInformation) {
	o.NrppaInfo = &v
}

// GetPwsInfo returns the PwsInfo field value if set, zero value otherwise.
func (o *N2InfoContainer) GetPwsInfo() PwsInformation {
	if o == nil || IsNil(o.PwsInfo) {
		var ret PwsInformation
		return ret
	}
	return *o.PwsInfo
}

// GetPwsInfoOk returns a tuple with the PwsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InfoContainer) GetPwsInfoOk() (*PwsInformation, bool) {
	if o == nil || IsNil(o.PwsInfo) {
		return nil, false
	}
	return o.PwsInfo, true
}

// HasPwsInfo returns a boolean if a field has been set.
func (o *N2InfoContainer) HasPwsInfo() bool {
	if o != nil && !IsNil(o.PwsInfo) {
		return true
	}

	return false
}

// SetPwsInfo gets a reference to the given PwsInformation and assigns it to the PwsInfo field.
func (o *N2InfoContainer) SetPwsInfo(v PwsInformation) {
	o.PwsInfo = &v
}

// GetV2xInfo returns the V2xInfo field value if set, zero value otherwise.
func (o *N2InfoContainer) GetV2xInfo() V2xInformation {
	if o == nil || IsNil(o.V2xInfo) {
		var ret V2xInformation
		return ret
	}
	return *o.V2xInfo
}

// GetV2xInfoOk returns a tuple with the V2xInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InfoContainer) GetV2xInfoOk() (*V2xInformation, bool) {
	if o == nil || IsNil(o.V2xInfo) {
		return nil, false
	}
	return o.V2xInfo, true
}

// HasV2xInfo returns a boolean if a field has been set.
func (o *N2InfoContainer) HasV2xInfo() bool {
	if o != nil && !IsNil(o.V2xInfo) {
		return true
	}

	return false
}

// SetV2xInfo gets a reference to the given V2xInformation and assigns it to the V2xInfo field.
func (o *N2InfoContainer) SetV2xInfo(v V2xInformation) {
	o.V2xInfo = &v
}

// GetProseInfo returns the ProseInfo field value if set, zero value otherwise.
func (o *N2InfoContainer) GetProseInfo() ProSeInformation {
	if o == nil || IsNil(o.ProseInfo) {
		var ret ProSeInformation
		return ret
	}
	return *o.ProseInfo
}

// GetProseInfoOk returns a tuple with the ProseInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2InfoContainer) GetProseInfoOk() (*ProSeInformation, bool) {
	if o == nil || IsNil(o.ProseInfo) {
		return nil, false
	}
	return o.ProseInfo, true
}

// HasProseInfo returns a boolean if a field has been set.
func (o *N2InfoContainer) HasProseInfo() bool {
	if o != nil && !IsNil(o.ProseInfo) {
		return true
	}

	return false
}

// SetProseInfo gets a reference to the given ProSeInformation and assigns it to the ProseInfo field.
func (o *N2InfoContainer) SetProseInfo(v ProSeInformation) {
	o.ProseInfo = &v
}

func (o N2InfoContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N2InfoContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["n2InformationClass"] = o.N2InformationClass
	if !IsNil(o.SmInfo) {
		toSerialize["smInfo"] = o.SmInfo
	}
	if !IsNil(o.RanInfo) {
		toSerialize["ranInfo"] = o.RanInfo
	}
	if !IsNil(o.NrppaInfo) {
		toSerialize["nrppaInfo"] = o.NrppaInfo
	}
	if !IsNil(o.PwsInfo) {
		toSerialize["pwsInfo"] = o.PwsInfo
	}
	if !IsNil(o.V2xInfo) {
		toSerialize["v2xInfo"] = o.V2xInfo
	}
	if !IsNil(o.ProseInfo) {
		toSerialize["proseInfo"] = o.ProseInfo
	}
	return toSerialize, nil
}

type NullableN2InfoContainer struct {
	value *N2InfoContainer
	isSet bool
}

func (v NullableN2InfoContainer) Get() *N2InfoContainer {
	return v.value
}

func (v *NullableN2InfoContainer) Set(val *N2InfoContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InfoContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InfoContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InfoContainer(val *N2InfoContainer) *NullableN2InfoContainer {
	return &NullableN2InfoContainer{value: val, isSet: true}
}

func (v NullableN2InfoContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InfoContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
