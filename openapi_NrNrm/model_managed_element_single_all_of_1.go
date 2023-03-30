/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the ManagedElementSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedElementSingleAllOf1{}

// ManagedElementSingleAllOf1 struct for ManagedElementSingleAllOf1
type ManagedElementSingleAllOf1 struct {
	GnbDuFunction []GnbDuFunctionSingle `json:"GnbDuFunction,omitempty"`
	GnbCuUpFunction []GnbCuUpFunctionSingle `json:"GnbCuUpFunction,omitempty"`
	GnbCuCpFunction []GnbCuCpFunctionSingle `json:"GnbCuCpFunction,omitempty"`
	DESManagementFunction *DESManagementFunctionSingle `json:"DESManagementFunction,omitempty"`
	DRACHOptimizationFunction *DRACHOptimizationFunctionSingle `json:"DRACHOptimizationFunction,omitempty"`
	DMROFunction *DMROFunctionSingle `json:"DMROFunction,omitempty"`
	DLBOFunction *DLBOFunctionSingle `json:"DLBOFunction,omitempty"`
	DPCIConfigurationFunction *DPCIConfigurationFunctionSingle `json:"DPCIConfigurationFunction,omitempty"`
	CPCIConfigurationFunction *CPCIConfigurationFunctionSingle `json:"CPCIConfigurationFunction,omitempty"`
	CESManagementFunction *CESManagementFunctionSingle `json:"CESManagementFunction,omitempty"`
	Configurable5QISet []Configurable5QISetSingle `json:"Configurable5QISet,omitempty"`
	Dynamic5QISet []Dynamic5QISetSingle `json:"Dynamic5QISet,omitempty"`
}

// NewManagedElementSingleAllOf1 instantiates a new ManagedElementSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedElementSingleAllOf1() *ManagedElementSingleAllOf1 {
	this := ManagedElementSingleAllOf1{}
	return &this
}

// NewManagedElementSingleAllOf1WithDefaults instantiates a new ManagedElementSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedElementSingleAllOf1WithDefaults() *ManagedElementSingleAllOf1 {
	this := ManagedElementSingleAllOf1{}
	return &this
}

// GetGnbDuFunction returns the GnbDuFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetGnbDuFunction() []GnbDuFunctionSingle {
	if o == nil || IsNil(o.GnbDuFunction) {
		var ret []GnbDuFunctionSingle
		return ret
	}
	return o.GnbDuFunction
}

// GetGnbDuFunctionOk returns a tuple with the GnbDuFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetGnbDuFunctionOk() ([]GnbDuFunctionSingle, bool) {
	if o == nil || IsNil(o.GnbDuFunction) {
		return nil, false
	}
	return o.GnbDuFunction, true
}

// HasGnbDuFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasGnbDuFunction() bool {
	if o != nil && !IsNil(o.GnbDuFunction) {
		return true
	}

	return false
}

// SetGnbDuFunction gets a reference to the given []GnbDuFunctionSingle and assigns it to the GnbDuFunction field.
func (o *ManagedElementSingleAllOf1) SetGnbDuFunction(v []GnbDuFunctionSingle) {
	o.GnbDuFunction = v
}

// GetGnbCuUpFunction returns the GnbCuUpFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetGnbCuUpFunction() []GnbCuUpFunctionSingle {
	if o == nil || IsNil(o.GnbCuUpFunction) {
		var ret []GnbCuUpFunctionSingle
		return ret
	}
	return o.GnbCuUpFunction
}

// GetGnbCuUpFunctionOk returns a tuple with the GnbCuUpFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetGnbCuUpFunctionOk() ([]GnbCuUpFunctionSingle, bool) {
	if o == nil || IsNil(o.GnbCuUpFunction) {
		return nil, false
	}
	return o.GnbCuUpFunction, true
}

// HasGnbCuUpFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasGnbCuUpFunction() bool {
	if o != nil && !IsNil(o.GnbCuUpFunction) {
		return true
	}

	return false
}

// SetGnbCuUpFunction gets a reference to the given []GnbCuUpFunctionSingle and assigns it to the GnbCuUpFunction field.
func (o *ManagedElementSingleAllOf1) SetGnbCuUpFunction(v []GnbCuUpFunctionSingle) {
	o.GnbCuUpFunction = v
}

// GetGnbCuCpFunction returns the GnbCuCpFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetGnbCuCpFunction() []GnbCuCpFunctionSingle {
	if o == nil || IsNil(o.GnbCuCpFunction) {
		var ret []GnbCuCpFunctionSingle
		return ret
	}
	return o.GnbCuCpFunction
}

// GetGnbCuCpFunctionOk returns a tuple with the GnbCuCpFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetGnbCuCpFunctionOk() ([]GnbCuCpFunctionSingle, bool) {
	if o == nil || IsNil(o.GnbCuCpFunction) {
		return nil, false
	}
	return o.GnbCuCpFunction, true
}

// HasGnbCuCpFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasGnbCuCpFunction() bool {
	if o != nil && !IsNil(o.GnbCuCpFunction) {
		return true
	}

	return false
}

// SetGnbCuCpFunction gets a reference to the given []GnbCuCpFunctionSingle and assigns it to the GnbCuCpFunction field.
func (o *ManagedElementSingleAllOf1) SetGnbCuCpFunction(v []GnbCuCpFunctionSingle) {
	o.GnbCuCpFunction = v
}

// GetDESManagementFunction returns the DESManagementFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetDESManagementFunction() DESManagementFunctionSingle {
	if o == nil || IsNil(o.DESManagementFunction) {
		var ret DESManagementFunctionSingle
		return ret
	}
	return *o.DESManagementFunction
}

// GetDESManagementFunctionOk returns a tuple with the DESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetDESManagementFunctionOk() (*DESManagementFunctionSingle, bool) {
	if o == nil || IsNil(o.DESManagementFunction) {
		return nil, false
	}
	return o.DESManagementFunction, true
}

// HasDESManagementFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasDESManagementFunction() bool {
	if o != nil && !IsNil(o.DESManagementFunction) {
		return true
	}

	return false
}

// SetDESManagementFunction gets a reference to the given DESManagementFunctionSingle and assigns it to the DESManagementFunction field.
func (o *ManagedElementSingleAllOf1) SetDESManagementFunction(v DESManagementFunctionSingle) {
	o.DESManagementFunction = &v
}

// GetDRACHOptimizationFunction returns the DRACHOptimizationFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetDRACHOptimizationFunction() DRACHOptimizationFunctionSingle {
	if o == nil || IsNil(o.DRACHOptimizationFunction) {
		var ret DRACHOptimizationFunctionSingle
		return ret
	}
	return *o.DRACHOptimizationFunction
}

// GetDRACHOptimizationFunctionOk returns a tuple with the DRACHOptimizationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetDRACHOptimizationFunctionOk() (*DRACHOptimizationFunctionSingle, bool) {
	if o == nil || IsNil(o.DRACHOptimizationFunction) {
		return nil, false
	}
	return o.DRACHOptimizationFunction, true
}

// HasDRACHOptimizationFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasDRACHOptimizationFunction() bool {
	if o != nil && !IsNil(o.DRACHOptimizationFunction) {
		return true
	}

	return false
}

// SetDRACHOptimizationFunction gets a reference to the given DRACHOptimizationFunctionSingle and assigns it to the DRACHOptimizationFunction field.
func (o *ManagedElementSingleAllOf1) SetDRACHOptimizationFunction(v DRACHOptimizationFunctionSingle) {
	o.DRACHOptimizationFunction = &v
}

// GetDMROFunction returns the DMROFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetDMROFunction() DMROFunctionSingle {
	if o == nil || IsNil(o.DMROFunction) {
		var ret DMROFunctionSingle
		return ret
	}
	return *o.DMROFunction
}

// GetDMROFunctionOk returns a tuple with the DMROFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetDMROFunctionOk() (*DMROFunctionSingle, bool) {
	if o == nil || IsNil(o.DMROFunction) {
		return nil, false
	}
	return o.DMROFunction, true
}

// HasDMROFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasDMROFunction() bool {
	if o != nil && !IsNil(o.DMROFunction) {
		return true
	}

	return false
}

// SetDMROFunction gets a reference to the given DMROFunctionSingle and assigns it to the DMROFunction field.
func (o *ManagedElementSingleAllOf1) SetDMROFunction(v DMROFunctionSingle) {
	o.DMROFunction = &v
}

// GetDLBOFunction returns the DLBOFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetDLBOFunction() DLBOFunctionSingle {
	if o == nil || IsNil(o.DLBOFunction) {
		var ret DLBOFunctionSingle
		return ret
	}
	return *o.DLBOFunction
}

// GetDLBOFunctionOk returns a tuple with the DLBOFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetDLBOFunctionOk() (*DLBOFunctionSingle, bool) {
	if o == nil || IsNil(o.DLBOFunction) {
		return nil, false
	}
	return o.DLBOFunction, true
}

// HasDLBOFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasDLBOFunction() bool {
	if o != nil && !IsNil(o.DLBOFunction) {
		return true
	}

	return false
}

// SetDLBOFunction gets a reference to the given DLBOFunctionSingle and assigns it to the DLBOFunction field.
func (o *ManagedElementSingleAllOf1) SetDLBOFunction(v DLBOFunctionSingle) {
	o.DLBOFunction = &v
}

// GetDPCIConfigurationFunction returns the DPCIConfigurationFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetDPCIConfigurationFunction() DPCIConfigurationFunctionSingle {
	if o == nil || IsNil(o.DPCIConfigurationFunction) {
		var ret DPCIConfigurationFunctionSingle
		return ret
	}
	return *o.DPCIConfigurationFunction
}

// GetDPCIConfigurationFunctionOk returns a tuple with the DPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetDPCIConfigurationFunctionOk() (*DPCIConfigurationFunctionSingle, bool) {
	if o == nil || IsNil(o.DPCIConfigurationFunction) {
		return nil, false
	}
	return o.DPCIConfigurationFunction, true
}

// HasDPCIConfigurationFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasDPCIConfigurationFunction() bool {
	if o != nil && !IsNil(o.DPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetDPCIConfigurationFunction gets a reference to the given DPCIConfigurationFunctionSingle and assigns it to the DPCIConfigurationFunction field.
func (o *ManagedElementSingleAllOf1) SetDPCIConfigurationFunction(v DPCIConfigurationFunctionSingle) {
	o.DPCIConfigurationFunction = &v
}

// GetCPCIConfigurationFunction returns the CPCIConfigurationFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetCPCIConfigurationFunction() CPCIConfigurationFunctionSingle {
	if o == nil || IsNil(o.CPCIConfigurationFunction) {
		var ret CPCIConfigurationFunctionSingle
		return ret
	}
	return *o.CPCIConfigurationFunction
}

// GetCPCIConfigurationFunctionOk returns a tuple with the CPCIConfigurationFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetCPCIConfigurationFunctionOk() (*CPCIConfigurationFunctionSingle, bool) {
	if o == nil || IsNil(o.CPCIConfigurationFunction) {
		return nil, false
	}
	return o.CPCIConfigurationFunction, true
}

// HasCPCIConfigurationFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasCPCIConfigurationFunction() bool {
	if o != nil && !IsNil(o.CPCIConfigurationFunction) {
		return true
	}

	return false
}

// SetCPCIConfigurationFunction gets a reference to the given CPCIConfigurationFunctionSingle and assigns it to the CPCIConfigurationFunction field.
func (o *ManagedElementSingleAllOf1) SetCPCIConfigurationFunction(v CPCIConfigurationFunctionSingle) {
	o.CPCIConfigurationFunction = &v
}

// GetCESManagementFunction returns the CESManagementFunction field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetCESManagementFunction() CESManagementFunctionSingle {
	if o == nil || IsNil(o.CESManagementFunction) {
		var ret CESManagementFunctionSingle
		return ret
	}
	return *o.CESManagementFunction
}

// GetCESManagementFunctionOk returns a tuple with the CESManagementFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetCESManagementFunctionOk() (*CESManagementFunctionSingle, bool) {
	if o == nil || IsNil(o.CESManagementFunction) {
		return nil, false
	}
	return o.CESManagementFunction, true
}

// HasCESManagementFunction returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasCESManagementFunction() bool {
	if o != nil && !IsNil(o.CESManagementFunction) {
		return true
	}

	return false
}

// SetCESManagementFunction gets a reference to the given CESManagementFunctionSingle and assigns it to the CESManagementFunction field.
func (o *ManagedElementSingleAllOf1) SetCESManagementFunction(v CESManagementFunctionSingle) {
	o.CESManagementFunction = &v
}

// GetConfigurable5QISet returns the Configurable5QISet field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetConfigurable5QISet() []Configurable5QISetSingle {
	if o == nil || IsNil(o.Configurable5QISet) {
		var ret []Configurable5QISetSingle
		return ret
	}
	return o.Configurable5QISet
}

// GetConfigurable5QISetOk returns a tuple with the Configurable5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetConfigurable5QISetOk() ([]Configurable5QISetSingle, bool) {
	if o == nil || IsNil(o.Configurable5QISet) {
		return nil, false
	}
	return o.Configurable5QISet, true
}

// HasConfigurable5QISet returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasConfigurable5QISet() bool {
	if o != nil && !IsNil(o.Configurable5QISet) {
		return true
	}

	return false
}

// SetConfigurable5QISet gets a reference to the given []Configurable5QISetSingle and assigns it to the Configurable5QISet field.
func (o *ManagedElementSingleAllOf1) SetConfigurable5QISet(v []Configurable5QISetSingle) {
	o.Configurable5QISet = v
}

// GetDynamic5QISet returns the Dynamic5QISet field value if set, zero value otherwise.
func (o *ManagedElementSingleAllOf1) GetDynamic5QISet() []Dynamic5QISetSingle {
	if o == nil || IsNil(o.Dynamic5QISet) {
		var ret []Dynamic5QISetSingle
		return ret
	}
	return o.Dynamic5QISet
}

// GetDynamic5QISetOk returns a tuple with the Dynamic5QISet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementSingleAllOf1) GetDynamic5QISetOk() ([]Dynamic5QISetSingle, bool) {
	if o == nil || IsNil(o.Dynamic5QISet) {
		return nil, false
	}
	return o.Dynamic5QISet, true
}

// HasDynamic5QISet returns a boolean if a field has been set.
func (o *ManagedElementSingleAllOf1) HasDynamic5QISet() bool {
	if o != nil && !IsNil(o.Dynamic5QISet) {
		return true
	}

	return false
}

// SetDynamic5QISet gets a reference to the given []Dynamic5QISetSingle and assigns it to the Dynamic5QISet field.
func (o *ManagedElementSingleAllOf1) SetDynamic5QISet(v []Dynamic5QISetSingle) {
	o.Dynamic5QISet = v
}

func (o ManagedElementSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedElementSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GnbDuFunction) {
		toSerialize["GnbDuFunction"] = o.GnbDuFunction
	}
	if !IsNil(o.GnbCuUpFunction) {
		toSerialize["GnbCuUpFunction"] = o.GnbCuUpFunction
	}
	if !IsNil(o.GnbCuCpFunction) {
		toSerialize["GnbCuCpFunction"] = o.GnbCuCpFunction
	}
	if !IsNil(o.DESManagementFunction) {
		toSerialize["DESManagementFunction"] = o.DESManagementFunction
	}
	if !IsNil(o.DRACHOptimizationFunction) {
		toSerialize["DRACHOptimizationFunction"] = o.DRACHOptimizationFunction
	}
	if !IsNil(o.DMROFunction) {
		toSerialize["DMROFunction"] = o.DMROFunction
	}
	if !IsNil(o.DLBOFunction) {
		toSerialize["DLBOFunction"] = o.DLBOFunction
	}
	if !IsNil(o.DPCIConfigurationFunction) {
		toSerialize["DPCIConfigurationFunction"] = o.DPCIConfigurationFunction
	}
	if !IsNil(o.CPCIConfigurationFunction) {
		toSerialize["CPCIConfigurationFunction"] = o.CPCIConfigurationFunction
	}
	if !IsNil(o.CESManagementFunction) {
		toSerialize["CESManagementFunction"] = o.CESManagementFunction
	}
	if !IsNil(o.Configurable5QISet) {
		toSerialize["Configurable5QISet"] = o.Configurable5QISet
	}
	if !IsNil(o.Dynamic5QISet) {
		toSerialize["Dynamic5QISet"] = o.Dynamic5QISet
	}
	return toSerialize, nil
}

type NullableManagedElementSingleAllOf1 struct {
	value *ManagedElementSingleAllOf1
	isSet bool
}

func (v NullableManagedElementSingleAllOf1) Get() *ManagedElementSingleAllOf1 {
	return v.value
}

func (v *NullableManagedElementSingleAllOf1) Set(val *ManagedElementSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedElementSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedElementSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedElementSingleAllOf1(val *ManagedElementSingleAllOf1) *NullableManagedElementSingleAllOf1 {
	return &NullableManagedElementSingleAllOf1{value: val, isSet: true}
}

func (v NullableManagedElementSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedElementSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


