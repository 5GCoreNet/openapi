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

// checks if the NrCellCuSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NrCellCuSingleAllOf1{}

// NrCellCuSingleAllOf1 struct for NrCellCuSingleAllOf1
type NrCellCuSingleAllOf1 struct {
	RRMPolicyRatio []RRMPolicyRatioSingle `json:"RRMPolicyRatio,omitempty"`
	NRCellRelation []NRCellRelationSingle `json:"NRCellRelation,omitempty"`
	EUtranCellRelation []EUtranCellRelationSingle `json:"EUtranCellRelation,omitempty"`
	NRFreqRelation []NRFreqRelationSingle `json:"NRFreqRelation,omitempty"`
	EUtranFreqRelation []EUtranFreqRelationSingle `json:"EUtranFreqRelation,omitempty"`
	DESManagementFunction *DESManagementFunctionSingle `json:"DESManagementFunction,omitempty"`
	DMROFunction *DMROFunctionSingle `json:"DMROFunction,omitempty"`
	DLBOFunction *DLBOFunctionSingle `json:"DLBOFunction,omitempty"`
	CESManagementFunction *CESManagementFunctionSingle `json:"CESManagementFunction,omitempty"`
	DPCIConfigurationFunction *DPCIConfigurationFunctionSingle `json:"DPCIConfigurationFunction,omitempty"`
}

// NewNrCellCuSingleAllOf1 instantiates a new NrCellCuSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNrCellCuSingleAllOf1() *NrCellCuSingleAllOf1 {
	this := NrCellCuSingleAllOf1{}
	return &this
}

// NewNrCellCuSingleAllOf1WithDefaults instantiates a new NrCellCuSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNrCellCuSingleAllOf1WithDefaults() *NrCellCuSingleAllOf1 {
	this := NrCellCuSingleAllOf1{}
	return &this
}

// GetRRMPolicyRatio returns the RRMPolicyRatio field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetRRMPolicyRatio() []RRMPolicyRatioSingle {
	if o == nil || IsNil(o.RRMPolicyRatio) {
		var ret []RRMPolicyRatioSingle
		return ret
	}
	return o.RRMPolicyRatio
}

// GetRRMPolicyRatioOk returns a tuple with the RRMPolicyRatio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetRRMPolicyRatioOk() ([]RRMPolicyRatioSingle, bool) {
	if o == nil || IsNil(o.RRMPolicyRatio) {
		return nil, false
	}
	return o.RRMPolicyRatio, true
}

// HasRRMPolicyRatio returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasRRMPolicyRatio() bool {
	if o != nil && !IsNil(o.RRMPolicyRatio) {
		return true
	}

	return false
}

// SetRRMPolicyRatio gets a reference to the given []RRMPolicyRatioSingle and assigns it to the RRMPolicyRatio field.
func (o *NrCellCuSingleAllOf1) SetRRMPolicyRatio(v []RRMPolicyRatioSingle) {
	o.RRMPolicyRatio = v
}

// GetNRCellRelation returns the NRCellRelation field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetNRCellRelation() []NRCellRelationSingle {
	if o == nil || IsNil(o.NRCellRelation) {
		var ret []NRCellRelationSingle
		return ret
	}
	return o.NRCellRelation
}

// GetNRCellRelationOk returns a tuple with the NRCellRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetNRCellRelationOk() ([]NRCellRelationSingle, bool) {
	if o == nil || IsNil(o.NRCellRelation) {
		return nil, false
	}
	return o.NRCellRelation, true
}

// HasNRCellRelation returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasNRCellRelation() bool {
	if o != nil && !IsNil(o.NRCellRelation) {
		return true
	}

	return false
}

// SetNRCellRelation gets a reference to the given []NRCellRelationSingle and assigns it to the NRCellRelation field.
func (o *NrCellCuSingleAllOf1) SetNRCellRelation(v []NRCellRelationSingle) {
	o.NRCellRelation = v
}

// GetEUtranCellRelation returns the EUtranCellRelation field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetEUtranCellRelation() []EUtranCellRelationSingle {
	if o == nil || IsNil(o.EUtranCellRelation) {
		var ret []EUtranCellRelationSingle
		return ret
	}
	return o.EUtranCellRelation
}

// GetEUtranCellRelationOk returns a tuple with the EUtranCellRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetEUtranCellRelationOk() ([]EUtranCellRelationSingle, bool) {
	if o == nil || IsNil(o.EUtranCellRelation) {
		return nil, false
	}
	return o.EUtranCellRelation, true
}

// HasEUtranCellRelation returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasEUtranCellRelation() bool {
	if o != nil && !IsNil(o.EUtranCellRelation) {
		return true
	}

	return false
}

// SetEUtranCellRelation gets a reference to the given []EUtranCellRelationSingle and assigns it to the EUtranCellRelation field.
func (o *NrCellCuSingleAllOf1) SetEUtranCellRelation(v []EUtranCellRelationSingle) {
	o.EUtranCellRelation = v
}

// GetNRFreqRelation returns the NRFreqRelation field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetNRFreqRelation() []NRFreqRelationSingle {
	if o == nil || IsNil(o.NRFreqRelation) {
		var ret []NRFreqRelationSingle
		return ret
	}
	return o.NRFreqRelation
}

// GetNRFreqRelationOk returns a tuple with the NRFreqRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetNRFreqRelationOk() ([]NRFreqRelationSingle, bool) {
	if o == nil || IsNil(o.NRFreqRelation) {
		return nil, false
	}
	return o.NRFreqRelation, true
}

// HasNRFreqRelation returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasNRFreqRelation() bool {
	if o != nil && !IsNil(o.NRFreqRelation) {
		return true
	}

	return false
}

// SetNRFreqRelation gets a reference to the given []NRFreqRelationSingle and assigns it to the NRFreqRelation field.
func (o *NrCellCuSingleAllOf1) SetNRFreqRelation(v []NRFreqRelationSingle) {
	o.NRFreqRelation = v
}

// GetEUtranFreqRelation returns the EUtranFreqRelation field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetEUtranFreqRelation() []EUtranFreqRelationSingle {
	if o == nil || IsNil(o.EUtranFreqRelation) {
		var ret []EUtranFreqRelationSingle
		return ret
	}
	return o.EUtranFreqRelation
}

// GetEUtranFreqRelationOk returns a tuple with the EUtranFreqRelation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetEUtranFreqRelationOk() ([]EUtranFreqRelationSingle, bool) {
	if o == nil || IsNil(o.EUtranFreqRelation) {
		return nil, false
	}
	return o.EUtranFreqRelation, true
}

// HasEUtranFreqRelation returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasEUtranFreqRelation() bool {
	if o != nil && !IsNil(o.EUtranFreqRelation) {
		return true
	}

	return false
}

// SetEUtranFreqRelation gets a reference to the given []EUtranFreqRelationSingle and assigns it to the EUtranFreqRelation field.
func (o *NrCellCuSingleAllOf1) SetEUtranFreqRelation(v []EUtranFreqRelationSingle) {
	o.EUtranFreqRelation = v
}

// GetDESManagementFunction returns the DESManagementFunction field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetDESManagementFunction() DESManagementFunctionSingle {
	if o == nil || IsNil(o.DESManagementFunction) {
		var ret DESManagementFunctionSingle
		return ret
	}
	return *o.DESManagementFunction
}

// GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool) {
	if o == nil || IsNil(o.DESManagementFunction) {
		return nil, false
	}
	return o.DESManagementFunction, true
}

// HasDESManagementFunction returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasDESManagementFunction() bool {
	if o != nil && !IsNil(o.DESManagementFunction) {
		return true
	}

	return false
}

// SetDESManagementFunction gets a reference to the given DESManagementFunctionSingle and assigns it to the DESManagementFunction field.
func (o *NrCellCuSingleAllOf1) SetDESManagementFunction(v DESManagementFunctionSingle) {
	o.DESManagementFunction = &v
}

// GetDMROFunction returns the DMROFunction field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetDMROFunction() DMROFunctionSingle {
	if o == nil || IsNil(o.DMROFunction) {
		var ret DMROFunctionSingle
		return ret
	}
	return *o.DMROFunction
}

// GetDMROFunctionOk returns a tuple with the DMROFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetDMROFunctionOk() (*DMROFunctionSingle, bool) {
	if o == nil || IsNil(o.DMROFunction) {
		return nil, false
	}
	return o.DMROFunction, true
}

// HasDMROFunction returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasDMROFunction() bool {
	if o != nil && !IsNil(o.DMROFunction) {
		return true
	}

	return false
}

// SetDMROFunction gets a reference to the given DMROFunctionSingle and assigns it to the DMROFunction field.
func (o *NrCellCuSingleAllOf1) SetDMROFunction(v DMROFunctionSingle) {
	o.DMROFunction = &v
}

// GetDLBOFunction returns the DLBOFunction field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetDLBOFunction() DLBOFunctionSingle {
	if o == nil || IsNil(o.DLBOFunction) {
		var ret DLBOFunctionSingle
		return ret
	}
	return *o.DLBOFunction
}

// GetDLBOFunctionOk returns a tuple with the DLBOFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool) {
	if o == nil || IsNil(o.DLBOFunction) {
		return nil, false
	}
	return o.DLBOFunction, true
}

// HasDLBOFunction returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasDLBOFunction() bool {
	if o != nil && !IsNil(o.DLBOFunction) {
		return true
	}

	return false
}

// SetDLBOFunction gets a reference to the given DLBOFunctionSingle and assigns it to the DLBOFunction field.
func (o *NrCellCuSingleAllOf1) SetDLBOFunction(v DLBOFunctionSingle) {
	o.DLBOFunction = &v
}

// GetCESManagementFunction returns the CESManagementFunction field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetCESManagementFunction() CESManagementFunctionSingle {
	if o == nil || IsNil(o.CESManagementFunction) {
		var ret CESManagementFunctionSingle
		return ret
	}
	return *o.CESManagementFunction
}

// GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool) {
	if o == nil || IsNil(o.CESManagementFunction) {
		return nil, false
	}
	return o.CESManagementFunction, true
}

// HasCESManagementFunction returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasCESManagementFunction() bool {
	if o != nil && !IsNil(o.CESManagementFunction) {
		return true
	}

	return false
}

// SetCESManagementFunction gets a reference to the given CESManagementFunctionSingle and assigns it to the CESManagementFunction field.
func (o *NrCellCuSingleAllOf1) SetCESManagementFunction(v CESManagementFunctionSingle) {
	o.CESManagementFunction = &v
}

// GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field value if set, zero value otherwise.
func (o *NrCellCuSingleAllOf1) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle {
	if o == nil || IsNil(o.DPCIConfigurationFunction) {
		var ret DPCIConfigurationFunctionSingle
		return ret
	}
	return *o.DPCIConfigurationFunction
}

// GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NrCellCuSingleAllOf1) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool) {
	if o == nil || IsNil(o.DPCIConfigurationFunction) {
		return nil, false
	}
	return o.DPCIConfigurationFunction, true
}

// HasDPCIConfigurationFunction returns a boolean if a field has been set.
func (o *NrCellCuSingleAllOf1) HasDPCIConfigurationFunction() bool {
	if o != nil && !IsNil(o.DPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetDPCIConfigurationFunction gets a reference to the given DPCIConfigurationFunctionSingle and assigns it to the DPCIConfigurationFunction field.
func (o *NrCellCuSingleAllOf1) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle) {
	o.DPCIConfigurationFunction = &v
}

func (o NrCellCuSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NrCellCuSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RRMPolicyRatio) {
		toSerialize["RRMPolicyRatio"] = o.RRMPolicyRatio
	}
	if !IsNil(o.NRCellRelation) {
		toSerialize["NRCellRelation"] = o.NRCellRelation
	}
	if !IsNil(o.EUtranCellRelation) {
		toSerialize["EUtranCellRelation"] = o.EUtranCellRelation
	}
	if !IsNil(o.NRFreqRelation) {
		toSerialize["NRFreqRelation"] = o.NRFreqRelation
	}
	if !IsNil(o.EUtranFreqRelation) {
		toSerialize["EUtranFreqRelation"] = o.EUtranFreqRelation
	}
	if !IsNil(o.DESManagementFunction) {
		toSerialize["DESManagementFunction"] = o.DESManagementFunction
	}
	if !IsNil(o.DMROFunction) {
		toSerialize["DMROFunction"] = o.DMROFunction
	}
	if !IsNil(o.DLBOFunction) {
		toSerialize["DLBOFunction"] = o.DLBOFunction
	}
	if !IsNil(o.CESManagementFunction) {
		toSerialize["CESManagementFunction"] = o.CESManagementFunction
	}
	if !IsNil(o.DPCIConfigurationFunction) {
		toSerialize["DPCIConfigurationFunction"] = o.DPCIConfigurationFunction
	}
	return toSerialize, nil
}

type NullableNrCellCuSingleAllOf1 struct {
	value *NrCellCuSingleAllOf1
	isSet bool
}

func (v NullableNrCellCuSingleAllOf1) Get() *NrCellCuSingleAllOf1 {
	return v.value
}

func (v *NullableNrCellCuSingleAllOf1) Set(val *NrCellCuSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableNrCellCuSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableNrCellCuSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrCellCuSingleAllOf1(val *NrCellCuSingleAllOf1) *NullableNrCellCuSingleAllOf1 {
	return &NullableNrCellCuSingleAllOf1{value: val, isSet: true}
}

func (v NullableNrCellCuSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrCellCuSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


