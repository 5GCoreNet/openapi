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

// checks if the DESManagementFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DESManagementFunctionSingleAllOfAttributes{}

// DESManagementFunctionSingleAllOfAttributes struct for DESManagementFunctionSingleAllOfAttributes
type DESManagementFunctionSingleAllOfAttributes struct {
	DesSwitch                                          *bool                                               `json:"desSwitch,omitempty"`
	IntraRatEsActivationOriginalCellLoadParameters     *IntraRatEsActivationOriginalCellLoadParameters     `json:"intraRatEsActivationOriginalCellLoadParameters,omitempty"`
	IntraRatEsActivationCandidateCellsLoadParameters   *IntraRatEsActivationCandidateCellsLoadParameters   `json:"intraRatEsActivationCandidateCellsLoadParameters,omitempty"`
	IntraRatEsDeactivationCandidateCellsLoadParameters *IntraRatEsDeactivationCandidateCellsLoadParameters `json:"intraRatEsDeactivationCandidateCellsLoadParameters,omitempty"`
	EsNotAllowedTimePeriod                             *EsNotAllowedTimePeriod                             `json:"esNotAllowedTimePeriod,omitempty"`
	InterRatEsActivationOriginalCellParameters         *InterRatEsActivationOriginalCellParameters         `json:"interRatEsActivationOriginalCellParameters,omitempty"`
	InterRatEsActivationCandidateCellParameters        *InterRatEsActivationCandidateCellParameters        `json:"interRatEsActivationCandidateCellParameters,omitempty"`
	InterRatEsDeactivationCandidateCellParameters      *InterRatEsDeactivationCandidateCellParameters      `json:"interRatEsDeactivationCandidateCellParameters,omitempty"`
	IsProbingCapable                                   *string                                             `json:"isProbingCapable,omitempty"`
	EnergySavingState                                  *string                                             `json:"energySavingState,omitempty"`
}

// NewDESManagementFunctionSingleAllOfAttributes instantiates a new DESManagementFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDESManagementFunctionSingleAllOfAttributes() *DESManagementFunctionSingleAllOfAttributes {
	this := DESManagementFunctionSingleAllOfAttributes{}
	return &this
}

// NewDESManagementFunctionSingleAllOfAttributesWithDefaults instantiates a new DESManagementFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDESManagementFunctionSingleAllOfAttributesWithDefaults() *DESManagementFunctionSingleAllOfAttributes {
	this := DESManagementFunctionSingleAllOfAttributes{}
	return &this
}

// GetDesSwitch returns the DesSwitch field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetDesSwitch() bool {
	if o == nil || IsNil(o.DesSwitch) {
		var ret bool
		return ret
	}
	return *o.DesSwitch
}

// GetDesSwitchOk returns a tuple with the DesSwitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetDesSwitchOk() (*bool, bool) {
	if o == nil || IsNil(o.DesSwitch) {
		return nil, false
	}
	return o.DesSwitch, true
}

// HasDesSwitch returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasDesSwitch() bool {
	if o != nil && !IsNil(o.DesSwitch) {
		return true
	}

	return false
}

// SetDesSwitch gets a reference to the given bool and assigns it to the DesSwitch field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetDesSwitch(v bool) {
	o.DesSwitch = &v
}

// GetIntraRatEsActivationOriginalCellLoadParameters returns the IntraRatEsActivationOriginalCellLoadParameters field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetIntraRatEsActivationOriginalCellLoadParameters() IntraRatEsActivationOriginalCellLoadParameters {
	if o == nil || IsNil(o.IntraRatEsActivationOriginalCellLoadParameters) {
		var ret IntraRatEsActivationOriginalCellLoadParameters
		return ret
	}
	return *o.IntraRatEsActivationOriginalCellLoadParameters
}

// GetIntraRatEsActivationOriginalCellLoadParametersOk returns a tuple with the IntraRatEsActivationOriginalCellLoadParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetIntraRatEsActivationOriginalCellLoadParametersOk() (*IntraRatEsActivationOriginalCellLoadParameters, bool) {
	if o == nil || IsNil(o.IntraRatEsActivationOriginalCellLoadParameters) {
		return nil, false
	}
	return o.IntraRatEsActivationOriginalCellLoadParameters, true
}

// HasIntraRatEsActivationOriginalCellLoadParameters returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasIntraRatEsActivationOriginalCellLoadParameters() bool {
	if o != nil && !IsNil(o.IntraRatEsActivationOriginalCellLoadParameters) {
		return true
	}

	return false
}

// SetIntraRatEsActivationOriginalCellLoadParameters gets a reference to the given IntraRatEsActivationOriginalCellLoadParameters and assigns it to the IntraRatEsActivationOriginalCellLoadParameters field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetIntraRatEsActivationOriginalCellLoadParameters(v IntraRatEsActivationOriginalCellLoadParameters) {
	o.IntraRatEsActivationOriginalCellLoadParameters = &v
}

// GetIntraRatEsActivationCandidateCellsLoadParameters returns the IntraRatEsActivationCandidateCellsLoadParameters field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetIntraRatEsActivationCandidateCellsLoadParameters() IntraRatEsActivationCandidateCellsLoadParameters {
	if o == nil || IsNil(o.IntraRatEsActivationCandidateCellsLoadParameters) {
		var ret IntraRatEsActivationCandidateCellsLoadParameters
		return ret
	}
	return *o.IntraRatEsActivationCandidateCellsLoadParameters
}

// GetIntraRatEsActivationCandidateCellsLoadParametersOk returns a tuple with the IntraRatEsActivationCandidateCellsLoadParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetIntraRatEsActivationCandidateCellsLoadParametersOk() (*IntraRatEsActivationCandidateCellsLoadParameters, bool) {
	if o == nil || IsNil(o.IntraRatEsActivationCandidateCellsLoadParameters) {
		return nil, false
	}
	return o.IntraRatEsActivationCandidateCellsLoadParameters, true
}

// HasIntraRatEsActivationCandidateCellsLoadParameters returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasIntraRatEsActivationCandidateCellsLoadParameters() bool {
	if o != nil && !IsNil(o.IntraRatEsActivationCandidateCellsLoadParameters) {
		return true
	}

	return false
}

// SetIntraRatEsActivationCandidateCellsLoadParameters gets a reference to the given IntraRatEsActivationCandidateCellsLoadParameters and assigns it to the IntraRatEsActivationCandidateCellsLoadParameters field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetIntraRatEsActivationCandidateCellsLoadParameters(v IntraRatEsActivationCandidateCellsLoadParameters) {
	o.IntraRatEsActivationCandidateCellsLoadParameters = &v
}

// GetIntraRatEsDeactivationCandidateCellsLoadParameters returns the IntraRatEsDeactivationCandidateCellsLoadParameters field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetIntraRatEsDeactivationCandidateCellsLoadParameters() IntraRatEsDeactivationCandidateCellsLoadParameters {
	if o == nil || IsNil(o.IntraRatEsDeactivationCandidateCellsLoadParameters) {
		var ret IntraRatEsDeactivationCandidateCellsLoadParameters
		return ret
	}
	return *o.IntraRatEsDeactivationCandidateCellsLoadParameters
}

// GetIntraRatEsDeactivationCandidateCellsLoadParametersOk returns a tuple with the IntraRatEsDeactivationCandidateCellsLoadParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetIntraRatEsDeactivationCandidateCellsLoadParametersOk() (*IntraRatEsDeactivationCandidateCellsLoadParameters, bool) {
	if o == nil || IsNil(o.IntraRatEsDeactivationCandidateCellsLoadParameters) {
		return nil, false
	}
	return o.IntraRatEsDeactivationCandidateCellsLoadParameters, true
}

// HasIntraRatEsDeactivationCandidateCellsLoadParameters returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasIntraRatEsDeactivationCandidateCellsLoadParameters() bool {
	if o != nil && !IsNil(o.IntraRatEsDeactivationCandidateCellsLoadParameters) {
		return true
	}

	return false
}

// SetIntraRatEsDeactivationCandidateCellsLoadParameters gets a reference to the given IntraRatEsDeactivationCandidateCellsLoadParameters and assigns it to the IntraRatEsDeactivationCandidateCellsLoadParameters field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetIntraRatEsDeactivationCandidateCellsLoadParameters(v IntraRatEsDeactivationCandidateCellsLoadParameters) {
	o.IntraRatEsDeactivationCandidateCellsLoadParameters = &v
}

// GetEsNotAllowedTimePeriod returns the EsNotAllowedTimePeriod field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetEsNotAllowedTimePeriod() EsNotAllowedTimePeriod {
	if o == nil || IsNil(o.EsNotAllowedTimePeriod) {
		var ret EsNotAllowedTimePeriod
		return ret
	}
	return *o.EsNotAllowedTimePeriod
}

// GetEsNotAllowedTimePeriodOk returns a tuple with the EsNotAllowedTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetEsNotAllowedTimePeriodOk() (*EsNotAllowedTimePeriod, bool) {
	if o == nil || IsNil(o.EsNotAllowedTimePeriod) {
		return nil, false
	}
	return o.EsNotAllowedTimePeriod, true
}

// HasEsNotAllowedTimePeriod returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasEsNotAllowedTimePeriod() bool {
	if o != nil && !IsNil(o.EsNotAllowedTimePeriod) {
		return true
	}

	return false
}

// SetEsNotAllowedTimePeriod gets a reference to the given EsNotAllowedTimePeriod and assigns it to the EsNotAllowedTimePeriod field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetEsNotAllowedTimePeriod(v EsNotAllowedTimePeriod) {
	o.EsNotAllowedTimePeriod = &v
}

// GetInterRatEsActivationOriginalCellParameters returns the InterRatEsActivationOriginalCellParameters field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetInterRatEsActivationOriginalCellParameters() InterRatEsActivationOriginalCellParameters {
	if o == nil || IsNil(o.InterRatEsActivationOriginalCellParameters) {
		var ret InterRatEsActivationOriginalCellParameters
		return ret
	}
	return *o.InterRatEsActivationOriginalCellParameters
}

// GetInterRatEsActivationOriginalCellParametersOk returns a tuple with the InterRatEsActivationOriginalCellParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetInterRatEsActivationOriginalCellParametersOk() (*InterRatEsActivationOriginalCellParameters, bool) {
	if o == nil || IsNil(o.InterRatEsActivationOriginalCellParameters) {
		return nil, false
	}
	return o.InterRatEsActivationOriginalCellParameters, true
}

// HasInterRatEsActivationOriginalCellParameters returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasInterRatEsActivationOriginalCellParameters() bool {
	if o != nil && !IsNil(o.InterRatEsActivationOriginalCellParameters) {
		return true
	}

	return false
}

// SetInterRatEsActivationOriginalCellParameters gets a reference to the given InterRatEsActivationOriginalCellParameters and assigns it to the InterRatEsActivationOriginalCellParameters field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetInterRatEsActivationOriginalCellParameters(v InterRatEsActivationOriginalCellParameters) {
	o.InterRatEsActivationOriginalCellParameters = &v
}

// GetInterRatEsActivationCandidateCellParameters returns the InterRatEsActivationCandidateCellParameters field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetInterRatEsActivationCandidateCellParameters() InterRatEsActivationCandidateCellParameters {
	if o == nil || IsNil(o.InterRatEsActivationCandidateCellParameters) {
		var ret InterRatEsActivationCandidateCellParameters
		return ret
	}
	return *o.InterRatEsActivationCandidateCellParameters
}

// GetInterRatEsActivationCandidateCellParametersOk returns a tuple with the InterRatEsActivationCandidateCellParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetInterRatEsActivationCandidateCellParametersOk() (*InterRatEsActivationCandidateCellParameters, bool) {
	if o == nil || IsNil(o.InterRatEsActivationCandidateCellParameters) {
		return nil, false
	}
	return o.InterRatEsActivationCandidateCellParameters, true
}

// HasInterRatEsActivationCandidateCellParameters returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasInterRatEsActivationCandidateCellParameters() bool {
	if o != nil && !IsNil(o.InterRatEsActivationCandidateCellParameters) {
		return true
	}

	return false
}

// SetInterRatEsActivationCandidateCellParameters gets a reference to the given InterRatEsActivationCandidateCellParameters and assigns it to the InterRatEsActivationCandidateCellParameters field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetInterRatEsActivationCandidateCellParameters(v InterRatEsActivationCandidateCellParameters) {
	o.InterRatEsActivationCandidateCellParameters = &v
}

// GetInterRatEsDeactivationCandidateCellParameters returns the InterRatEsDeactivationCandidateCellParameters field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetInterRatEsDeactivationCandidateCellParameters() InterRatEsDeactivationCandidateCellParameters {
	if o == nil || IsNil(o.InterRatEsDeactivationCandidateCellParameters) {
		var ret InterRatEsDeactivationCandidateCellParameters
		return ret
	}
	return *o.InterRatEsDeactivationCandidateCellParameters
}

// GetInterRatEsDeactivationCandidateCellParametersOk returns a tuple with the InterRatEsDeactivationCandidateCellParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetInterRatEsDeactivationCandidateCellParametersOk() (*InterRatEsDeactivationCandidateCellParameters, bool) {
	if o == nil || IsNil(o.InterRatEsDeactivationCandidateCellParameters) {
		return nil, false
	}
	return o.InterRatEsDeactivationCandidateCellParameters, true
}

// HasInterRatEsDeactivationCandidateCellParameters returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasInterRatEsDeactivationCandidateCellParameters() bool {
	if o != nil && !IsNil(o.InterRatEsDeactivationCandidateCellParameters) {
		return true
	}

	return false
}

// SetInterRatEsDeactivationCandidateCellParameters gets a reference to the given InterRatEsDeactivationCandidateCellParameters and assigns it to the InterRatEsDeactivationCandidateCellParameters field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetInterRatEsDeactivationCandidateCellParameters(v InterRatEsDeactivationCandidateCellParameters) {
	o.InterRatEsDeactivationCandidateCellParameters = &v
}

// GetIsProbingCapable returns the IsProbingCapable field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetIsProbingCapable() string {
	if o == nil || IsNil(o.IsProbingCapable) {
		var ret string
		return ret
	}
	return *o.IsProbingCapable
}

// GetIsProbingCapableOk returns a tuple with the IsProbingCapable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetIsProbingCapableOk() (*string, bool) {
	if o == nil || IsNil(o.IsProbingCapable) {
		return nil, false
	}
	return o.IsProbingCapable, true
}

// HasIsProbingCapable returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasIsProbingCapable() bool {
	if o != nil && !IsNil(o.IsProbingCapable) {
		return true
	}

	return false
}

// SetIsProbingCapable gets a reference to the given string and assigns it to the IsProbingCapable field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetIsProbingCapable(v string) {
	o.IsProbingCapable = &v
}

// GetEnergySavingState returns the EnergySavingState field value if set, zero value otherwise.
func (o *DESManagementFunctionSingleAllOfAttributes) GetEnergySavingState() string {
	if o == nil || IsNil(o.EnergySavingState) {
		var ret string
		return ret
	}
	return *o.EnergySavingState
}

// GetEnergySavingStateOk returns a tuple with the EnergySavingState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) GetEnergySavingStateOk() (*string, bool) {
	if o == nil || IsNil(o.EnergySavingState) {
		return nil, false
	}
	return o.EnergySavingState, true
}

// HasEnergySavingState returns a boolean if a field has been set.
func (o *DESManagementFunctionSingleAllOfAttributes) HasEnergySavingState() bool {
	if o != nil && !IsNil(o.EnergySavingState) {
		return true
	}

	return false
}

// SetEnergySavingState gets a reference to the given string and assigns it to the EnergySavingState field.
func (o *DESManagementFunctionSingleAllOfAttributes) SetEnergySavingState(v string) {
	o.EnergySavingState = &v
}

func (o DESManagementFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DESManagementFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DesSwitch) {
		toSerialize["desSwitch"] = o.DesSwitch
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
	if !IsNil(o.IsProbingCapable) {
		toSerialize["isProbingCapable"] = o.IsProbingCapable
	}
	if !IsNil(o.EnergySavingState) {
		toSerialize["energySavingState"] = o.EnergySavingState
	}
	return toSerialize, nil
}

type NullableDESManagementFunctionSingleAllOfAttributes struct {
	value *DESManagementFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableDESManagementFunctionSingleAllOfAttributes) Get() *DESManagementFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableDESManagementFunctionSingleAllOfAttributes) Set(val *DESManagementFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDESManagementFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDESManagementFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDESManagementFunctionSingleAllOfAttributes(val *DESManagementFunctionSingleAllOfAttributes) *NullableDESManagementFunctionSingleAllOfAttributes {
	return &NullableDESManagementFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableDESManagementFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDESManagementFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
