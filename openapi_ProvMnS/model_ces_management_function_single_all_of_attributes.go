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

// checks if the CESManagementFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CESManagementFunctionSingleAllOfAttributes{}

// CESManagementFunctionSingleAllOfAttributes struct for CESManagementFunctionSingleAllOfAttributes
type CESManagementFunctionSingleAllOfAttributes struct {
	CesSwitch *bool `json:"cesSwitch,omitempty"`
	IntraRatEsActivationOriginalCellLoadParameters *IntraRatEsActivationOriginalCellLoadParameters `json:"intraRatEsActivationOriginalCellLoadParameters,omitempty"`
	IntraRatEsActivationCandidateCellsLoadParameters *IntraRatEsActivationCandidateCellsLoadParameters `json:"intraRatEsActivationCandidateCellsLoadParameters,omitempty"`
	IntraRatEsDeactivationCandidateCellsLoadParameters *IntraRatEsDeactivationCandidateCellsLoadParameters `json:"intraRatEsDeactivationCandidateCellsLoadParameters,omitempty"`
	EsNotAllowedTimePeriod *EsNotAllowedTimePeriod `json:"esNotAllowedTimePeriod,omitempty"`
	InterRatEsActivationOriginalCellParameters *IntraRatEsActivationOriginalCellLoadParameters `json:"interRatEsActivationOriginalCellParameters,omitempty"`
	InterRatEsActivationCandidateCellParameters *IntraRatEsActivationOriginalCellLoadParameters `json:"interRatEsActivationCandidateCellParameters,omitempty"`
	InterRatEsDeactivationCandidateCellParameters *IntraRatEsActivationOriginalCellLoadParameters `json:"interRatEsDeactivationCandidateCellParameters,omitempty"`
	EnergySavingControl *string `json:"energySavingControl,omitempty"`
	EnergySavingState *string `json:"energySavingState,omitempty"`
}

// NewCESManagementFunctionSingleAllOfAttributes instantiates a new CESManagementFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCESManagementFunctionSingleAllOfAttributes() *CESManagementFunctionSingleAllOfAttributes {
	this := CESManagementFunctionSingleAllOfAttributes{}
	return &this
}

// NewCESManagementFunctionSingleAllOfAttributesWithDefaults instantiates a new CESManagementFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCESManagementFunctionSingleAllOfAttributesWithDefaults() *CESManagementFunctionSingleAllOfAttributes {
	this := CESManagementFunctionSingleAllOfAttributes{}
	return &this
}

// GetCesSwitch returns the CesSwitch field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetCesSwitch() bool {
	if o == nil || IsNil(o.CesSwitch) {
		var ret bool
		return ret
	}
	return *o.CesSwitch
}

// GetCesSwitchOk returns a tuple with the CesSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetCesSwitchOk() (*bool, bool) {
	if o == nil || IsNil(o.CesSwitch) {
		return nil, false
	}
	return o.CesSwitch, true
}

// HasCesSwitch returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasCesSwitch() bool {
	if o != nil && !IsNil(o.CesSwitch) {
		return true
	}

	return false
}

// SetCesSwitch gets a reference to the given bool and assigns it to the CesSwitch field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetCesSwitch(v bool) {
	o.CesSwitch = &v
}

// GetIntraRatEsActivationOriginalCellLoadParameters returns the IntraRatEsActivationOriginalCellLoadParameters field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetIntraRatEsActivationOriginalCellLoadParameters() IntraRatEsActivationOriginalCellLoadParameters {
	if o == nil || IsNil(o.IntraRatEsActivationOriginalCellLoadParameters) {
		var ret IntraRatEsActivationOriginalCellLoadParameters
		return ret
	}
	return *o.IntraRatEsActivationOriginalCellLoadParameters
}

// GetIntraRatEsActivationOriginalCellLoadParametersOk returns a tuple with the IntraRatEsActivationOriginalCellLoadParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetIntraRatEsActivationOriginalCellLoadParametersOk() (*IntraRatEsActivationOriginalCellLoadParameters, bool) {
	if o == nil || IsNil(o.IntraRatEsActivationOriginalCellLoadParameters) {
		return nil, false
	}
	return o.IntraRatEsActivationOriginalCellLoadParameters, true
}

// HasIntraRatEsActivationOriginalCellLoadParameters returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasIntraRatEsActivationOriginalCellLoadParameters() bool {
	if o != nil && !IsNil(o.IntraRatEsActivationOriginalCellLoadParameters) {
		return true
	}

	return false
}

// SetIntraRatEsActivationOriginalCellLoadParameters gets a reference to the given IntraRatEsActivationOriginalCellLoadParameters and assigns it to the IntraRatEsActivationOriginalCellLoadParameters field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetIntraRatEsActivationOriginalCellLoadParameters(v IntraRatEsActivationOriginalCellLoadParameters) {
	o.IntraRatEsActivationOriginalCellLoadParameters = &v
}

// GetIntraRatEsActivationCandidateCellsLoadParameters returns the IntraRatEsActivationCandidateCellsLoadParameters field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetIntraRatEsActivationCandidateCellsLoadParameters() IntraRatEsActivationCandidateCellsLoadParameters {
	if o == nil || IsNil(o.IntraRatEsActivationCandidateCellsLoadParameters) {
		var ret IntraRatEsActivationCandidateCellsLoadParameters
		return ret
	}
	return *o.IntraRatEsActivationCandidateCellsLoadParameters
}

// GetIntraRatEsActivationCandidateCellsLoadParametersOk returns a tuple with the IntraRatEsActivationCandidateCellsLoadParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetIntraRatEsActivationCandidateCellsLoadParametersOk() (*IntraRatEsActivationCandidateCellsLoadParameters, bool) {
	if o == nil || IsNil(o.IntraRatEsActivationCandidateCellsLoadParameters) {
		return nil, false
	}
	return o.IntraRatEsActivationCandidateCellsLoadParameters, true
}

// HasIntraRatEsActivationCandidateCellsLoadParameters returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasIntraRatEsActivationCandidateCellsLoadParameters() bool {
	if o != nil && !IsNil(o.IntraRatEsActivationCandidateCellsLoadParameters) {
		return true
	}

	return false
}

// SetIntraRatEsActivationCandidateCellsLoadParameters gets a reference to the given IntraRatEsActivationCandidateCellsLoadParameters and assigns it to the IntraRatEsActivationCandidateCellsLoadParameters field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetIntraRatEsActivationCandidateCellsLoadParameters(v IntraRatEsActivationCandidateCellsLoadParameters) {
	o.IntraRatEsActivationCandidateCellsLoadParameters = &v
}

// GetIntraRatEsDeactivationCandidateCellsLoadParameters returns the IntraRatEsDeactivationCandidateCellsLoadParameters field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetIntraRatEsDeactivationCandidateCellsLoadParameters() IntraRatEsDeactivationCandidateCellsLoadParameters {
	if o == nil || IsNil(o.IntraRatEsDeactivationCandidateCellsLoadParameters) {
		var ret IntraRatEsDeactivationCandidateCellsLoadParameters
		return ret
	}
	return *o.IntraRatEsDeactivationCandidateCellsLoadParameters
}

// GetIntraRatEsDeactivationCandidateCellsLoadParametersOk returns a tuple with the IntraRatEsDeactivationCandidateCellsLoadParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetIntraRatEsDeactivationCandidateCellsLoadParametersOk() (*IntraRatEsDeactivationCandidateCellsLoadParameters, bool) {
	if o == nil || IsNil(o.IntraRatEsDeactivationCandidateCellsLoadParameters) {
		return nil, false
	}
	return o.IntraRatEsDeactivationCandidateCellsLoadParameters, true
}

// HasIntraRatEsDeactivationCandidateCellsLoadParameters returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasIntraRatEsDeactivationCandidateCellsLoadParameters() bool {
	if o != nil && !IsNil(o.IntraRatEsDeactivationCandidateCellsLoadParameters) {
		return true
	}

	return false
}

// SetIntraRatEsDeactivationCandidateCellsLoadParameters gets a reference to the given IntraRatEsDeactivationCandidateCellsLoadParameters and assigns it to the IntraRatEsDeactivationCandidateCellsLoadParameters field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetIntraRatEsDeactivationCandidateCellsLoadParameters(v IntraRatEsDeactivationCandidateCellsLoadParameters) {
	o.IntraRatEsDeactivationCandidateCellsLoadParameters = &v
}

// GetEsNotAllowedTimePeriod returns the EsNotAllowedTimePeriod field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetEsNotAllowedTimePeriod() EsNotAllowedTimePeriod {
	if o == nil || IsNil(o.EsNotAllowedTimePeriod) {
		var ret EsNotAllowedTimePeriod
		return ret
	}
	return *o.EsNotAllowedTimePeriod
}

// GetEsNotAllowedTimePeriodOk returns a tuple with the EsNotAllowedTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetEsNotAllowedTimePeriodOk() (*EsNotAllowedTimePeriod, bool) {
	if o == nil || IsNil(o.EsNotAllowedTimePeriod) {
		return nil, false
	}
	return o.EsNotAllowedTimePeriod, true
}

// HasEsNotAllowedTimePeriod returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasEsNotAllowedTimePeriod() bool {
	if o != nil && !IsNil(o.EsNotAllowedTimePeriod) {
		return true
	}

	return false
}

// SetEsNotAllowedTimePeriod gets a reference to the given EsNotAllowedTimePeriod and assigns it to the EsNotAllowedTimePeriod field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetEsNotAllowedTimePeriod(v EsNotAllowedTimePeriod) {
	o.EsNotAllowedTimePeriod = &v
}

// GetInterRatEsActivationOriginalCellParameters returns the InterRatEsActivationOriginalCellParameters field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetInterRatEsActivationOriginalCellParameters() IntraRatEsActivationOriginalCellLoadParameters {
	if o == nil || IsNil(o.InterRatEsActivationOriginalCellParameters) {
		var ret IntraRatEsActivationOriginalCellLoadParameters
		return ret
	}
	return *o.InterRatEsActivationOriginalCellParameters
}

// GetInterRatEsActivationOriginalCellParametersOk returns a tuple with the InterRatEsActivationOriginalCellParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetInterRatEsActivationOriginalCellParametersOk() (*IntraRatEsActivationOriginalCellLoadParameters, bool) {
	if o == nil || IsNil(o.InterRatEsActivationOriginalCellParameters) {
		return nil, false
	}
	return o.InterRatEsActivationOriginalCellParameters, true
}

// HasInterRatEsActivationOriginalCellParameters returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasInterRatEsActivationOriginalCellParameters() bool {
	if o != nil && !IsNil(o.InterRatEsActivationOriginalCellParameters) {
		return true
	}

	return false
}

// SetInterRatEsActivationOriginalCellParameters gets a reference to the given IntraRatEsActivationOriginalCellLoadParameters and assigns it to the InterRatEsActivationOriginalCellParameters field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetInterRatEsActivationOriginalCellParameters(v IntraRatEsActivationOriginalCellLoadParameters) {
	o.InterRatEsActivationOriginalCellParameters = &v
}

// GetInterRatEsActivationCandidateCellParameters returns the InterRatEsActivationCandidateCellParameters field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetInterRatEsActivationCandidateCellParameters() IntraRatEsActivationOriginalCellLoadParameters {
	if o == nil || IsNil(o.InterRatEsActivationCandidateCellParameters) {
		var ret IntraRatEsActivationOriginalCellLoadParameters
		return ret
	}
	return *o.InterRatEsActivationCandidateCellParameters
}

// GetInterRatEsActivationCandidateCellParametersOk returns a tuple with the InterRatEsActivationCandidateCellParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetInterRatEsActivationCandidateCellParametersOk() (*IntraRatEsActivationOriginalCellLoadParameters, bool) {
	if o == nil || IsNil(o.InterRatEsActivationCandidateCellParameters) {
		return nil, false
	}
	return o.InterRatEsActivationCandidateCellParameters, true
}

// HasInterRatEsActivationCandidateCellParameters returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasInterRatEsActivationCandidateCellParameters() bool {
	if o != nil && !IsNil(o.InterRatEsActivationCandidateCellParameters) {
		return true
	}

	return false
}

// SetInterRatEsActivationCandidateCellParameters gets a reference to the given IntraRatEsActivationOriginalCellLoadParameters and assigns it to the InterRatEsActivationCandidateCellParameters field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetInterRatEsActivationCandidateCellParameters(v IntraRatEsActivationOriginalCellLoadParameters) {
	o.InterRatEsActivationCandidateCellParameters = &v
}

// GetInterRatEsDeactivationCandidateCellParameters returns the InterRatEsDeactivationCandidateCellParameters field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetInterRatEsDeactivationCandidateCellParameters() IntraRatEsActivationOriginalCellLoadParameters {
	if o == nil || IsNil(o.InterRatEsDeactivationCandidateCellParameters) {
		var ret IntraRatEsActivationOriginalCellLoadParameters
		return ret
	}
	return *o.InterRatEsDeactivationCandidateCellParameters
}

// GetInterRatEsDeactivationCandidateCellParametersOk returns a tuple with the InterRatEsDeactivationCandidateCellParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetInterRatEsDeactivationCandidateCellParametersOk() (*IntraRatEsActivationOriginalCellLoadParameters, bool) {
	if o == nil || IsNil(o.InterRatEsDeactivationCandidateCellParameters) {
		return nil, false
	}
	return o.InterRatEsDeactivationCandidateCellParameters, true
}

// HasInterRatEsDeactivationCandidateCellParameters returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasInterRatEsDeactivationCandidateCellParameters() bool {
	if o != nil && !IsNil(o.InterRatEsDeactivationCandidateCellParameters) {
		return true
	}

	return false
}

// SetInterRatEsDeactivationCandidateCellParameters gets a reference to the given IntraRatEsActivationOriginalCellLoadParameters and assigns it to the InterRatEsDeactivationCandidateCellParameters field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetInterRatEsDeactivationCandidateCellParameters(v IntraRatEsActivationOriginalCellLoadParameters) {
	o.InterRatEsDeactivationCandidateCellParameters = &v
}

// GetEnergySavingControl returns the EnergySavingControl field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetEnergySavingControl() string {
	if o == nil || IsNil(o.EnergySavingControl) {
		var ret string
		return ret
	}
	return *o.EnergySavingControl
}

// GetEnergySavingControlOk returns a tuple with the EnergySavingControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetEnergySavingControlOk() (*string, bool) {
	if o == nil || IsNil(o.EnergySavingControl) {
		return nil, false
	}
	return o.EnergySavingControl, true
}

// HasEnergySavingControl returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasEnergySavingControl() bool {
	if o != nil && !IsNil(o.EnergySavingControl) {
		return true
	}

	return false
}

// SetEnergySavingControl gets a reference to the given string and assigns it to the EnergySavingControl field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetEnergySavingControl(v string) {
	o.EnergySavingControl = &v
}

// GetEnergySavingState returns the EnergySavingState field value if set, zero value otherwise.
func (o *CESManagementFunctionSingleAllOfAttributes) GetEnergySavingState() string {
	if o == nil || IsNil(o.EnergySavingState) {
		var ret string
		return ret
	}
	return *o.EnergySavingState
}

// GetEnergySavingStateOk returns a tuple with the EnergySavingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) GetEnergySavingStateOk() (*string, bool) {
	if o == nil || IsNil(o.EnergySavingState) {
		return nil, false
	}
	return o.EnergySavingState, true
}

// HasEnergySavingState returns a boolean if a field has been set.
func (o *CESManagementFunctionSingleAllOfAttributes) HasEnergySavingState() bool {
	if o != nil && !IsNil(o.EnergySavingState) {
		return true
	}

	return false
}

// SetEnergySavingState gets a reference to the given string and assigns it to the EnergySavingState field.
func (o *CESManagementFunctionSingleAllOfAttributes) SetEnergySavingState(v string) {
	o.EnergySavingState = &v
}

func (o CESManagementFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CESManagementFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CesSwitch) {
		toSerialize["cesSwitch"] = o.CesSwitch
	}
	if !IsNil(o.IntraRatEsActivationOriginalCellLoadParameters) {
		toSerialize["intraRatEsActivationOriginalCellLoadParameters"] = o.IntraRatEsActivationOriginalCellLoadParameters
	}
	if !IsNil(o.IntraRatEsActivationCandidateCellsLoadParameters) {
		toSerialize["intraRatEsActivationCandidateCellsLoadParameters"] = o.IntraRatEsActivationCandidateCellsLoadParameters
	}
	if !IsNil(o.IntraRatEsDeactivationCandidateCellsLoadParameters) {
		toSerialize["intraRatEsDeactivationCandidateCellsLoadParameters"] = o.IntraRatEsDeactivationCandidateCellsLoadParameters
	}
	if !IsNil(o.EsNotAllowedTimePeriod) {
		toSerialize["esNotAllowedTimePeriod"] = o.EsNotAllowedTimePeriod
	}
	if !IsNil(o.InterRatEsActivationOriginalCellParameters) {
		toSerialize["interRatEsActivationOriginalCellParameters"] = o.InterRatEsActivationOriginalCellParameters
	}
	if !IsNil(o.InterRatEsActivationCandidateCellParameters) {
		toSerialize["interRatEsActivationCandidateCellParameters"] = o.InterRatEsActivationCandidateCellParameters
	}
	if !IsNil(o.InterRatEsDeactivationCandidateCellParameters) {
		toSerialize["interRatEsDeactivationCandidateCellParameters"] = o.InterRatEsDeactivationCandidateCellParameters
	}
	if !IsNil(o.EnergySavingControl) {
		toSerialize["energySavingControl"] = o.EnergySavingControl
	}
	if !IsNil(o.EnergySavingState) {
		toSerialize["energySavingState"] = o.EnergySavingState
	}
	return toSerialize, nil
}

type NullableCESManagementFunctionSingleAllOfAttributes struct {
	value *CESManagementFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableCESManagementFunctionSingleAllOfAttributes) Get() *CESManagementFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableCESManagementFunctionSingleAllOfAttributes) Set(val *CESManagementFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableCESManagementFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableCESManagementFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCESManagementFunctionSingleAllOfAttributes(val *CESManagementFunctionSingleAllOfAttributes) *NullableCESManagementFunctionSingleAllOfAttributes {
	return &NullableCESManagementFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableCESManagementFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCESManagementFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


