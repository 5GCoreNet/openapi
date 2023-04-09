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

// checks if the DLBOFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DLBOFunctionSingleAllOfAttributes{}

// DLBOFunctionSingleAllOfAttributes struct for DLBOFunctionSingleAllOfAttributes
type DLBOFunctionSingleAllOfAttributes struct {
	DlboControl                       *bool  `json:"dlboControl,omitempty"`
	MaximumDeviationHoTrigger         *int32 `json:"maximumDeviationHoTrigger,omitempty"`
	MinimumTimeBetweenHoTriggerChange *int32 `json:"minimumTimeBetweenHoTriggerChange,omitempty"`
}

// NewDLBOFunctionSingleAllOfAttributes instantiates a new DLBOFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDLBOFunctionSingleAllOfAttributes() *DLBOFunctionSingleAllOfAttributes {
	this := DLBOFunctionSingleAllOfAttributes{}
	return &this
}

// NewDLBOFunctionSingleAllOfAttributesWithDefaults instantiates a new DLBOFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDLBOFunctionSingleAllOfAttributesWithDefaults() *DLBOFunctionSingleAllOfAttributes {
	this := DLBOFunctionSingleAllOfAttributes{}
	return &this
}

// GetDlboControl returns the DlboControl field value if set, zero value otherwise.
func (o *DLBOFunctionSingleAllOfAttributes) GetDlboControl() bool {
	if o == nil || IsNil(o.DlboControl) {
		var ret bool
		return ret
	}
	return *o.DlboControl
}

// GetDlboControlOk returns a tuple with the DlboControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLBOFunctionSingleAllOfAttributes) GetDlboControlOk() (*bool, bool) {
	if o == nil || IsNil(o.DlboControl) {
		return nil, false
	}
	return o.DlboControl, true
}

// HasDlboControl returns a boolean if a field has been set.
func (o *DLBOFunctionSingleAllOfAttributes) HasDlboControl() bool {
	if o != nil && !IsNil(o.DlboControl) {
		return true
	}

	return false
}

// SetDlboControl gets a reference to the given bool and assigns it to the DlboControl field.
func (o *DLBOFunctionSingleAllOfAttributes) SetDlboControl(v bool) {
	o.DlboControl = &v
}

// GetMaximumDeviationHoTrigger returns the MaximumDeviationHoTrigger field value if set, zero value otherwise.
func (o *DLBOFunctionSingleAllOfAttributes) GetMaximumDeviationHoTrigger() int32 {
	if o == nil || IsNil(o.MaximumDeviationHoTrigger) {
		var ret int32
		return ret
	}
	return *o.MaximumDeviationHoTrigger
}

// GetMaximumDeviationHoTriggerOk returns a tuple with the MaximumDeviationHoTrigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLBOFunctionSingleAllOfAttributes) GetMaximumDeviationHoTriggerOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumDeviationHoTrigger) {
		return nil, false
	}
	return o.MaximumDeviationHoTrigger, true
}

// HasMaximumDeviationHoTrigger returns a boolean if a field has been set.
func (o *DLBOFunctionSingleAllOfAttributes) HasMaximumDeviationHoTrigger() bool {
	if o != nil && !IsNil(o.MaximumDeviationHoTrigger) {
		return true
	}

	return false
}

// SetMaximumDeviationHoTrigger gets a reference to the given int32 and assigns it to the MaximumDeviationHoTrigger field.
func (o *DLBOFunctionSingleAllOfAttributes) SetMaximumDeviationHoTrigger(v int32) {
	o.MaximumDeviationHoTrigger = &v
}

// GetMinimumTimeBetweenHoTriggerChange returns the MinimumTimeBetweenHoTriggerChange field value if set, zero value otherwise.
func (o *DLBOFunctionSingleAllOfAttributes) GetMinimumTimeBetweenHoTriggerChange() int32 {
	if o == nil || IsNil(o.MinimumTimeBetweenHoTriggerChange) {
		var ret int32
		return ret
	}
	return *o.MinimumTimeBetweenHoTriggerChange
}

// GetMinimumTimeBetweenHoTriggerChangeOk returns a tuple with the MinimumTimeBetweenHoTriggerChange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DLBOFunctionSingleAllOfAttributes) GetMinimumTimeBetweenHoTriggerChangeOk() (*int32, bool) {
	if o == nil || IsNil(o.MinimumTimeBetweenHoTriggerChange) {
		return nil, false
	}
	return o.MinimumTimeBetweenHoTriggerChange, true
}

// HasMinimumTimeBetweenHoTriggerChange returns a boolean if a field has been set.
func (o *DLBOFunctionSingleAllOfAttributes) HasMinimumTimeBetweenHoTriggerChange() bool {
	if o != nil && !IsNil(o.MinimumTimeBetweenHoTriggerChange) {
		return true
	}

	return false
}

// SetMinimumTimeBetweenHoTriggerChange gets a reference to the given int32 and assigns it to the MinimumTimeBetweenHoTriggerChange field.
func (o *DLBOFunctionSingleAllOfAttributes) SetMinimumTimeBetweenHoTriggerChange(v int32) {
	o.MinimumTimeBetweenHoTriggerChange = &v
}

func (o DLBOFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DLBOFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DlboControl) {
		toSerialize["dlboControl"] = o.DlboControl
	}
	if !IsNil(o.MaximumDeviationHoTrigger) {
		toSerialize["maximumDeviationHoTrigger"] = o.MaximumDeviationHoTrigger
	}
	if !IsNil(o.MinimumTimeBetweenHoTriggerChange) {
		toSerialize["minimumTimeBetweenHoTriggerChange"] = o.MinimumTimeBetweenHoTriggerChange
	}
	return toSerialize, nil
}

type NullableDLBOFunctionSingleAllOfAttributes struct {
	value *DLBOFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableDLBOFunctionSingleAllOfAttributes) Get() *DLBOFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableDLBOFunctionSingleAllOfAttributes) Set(val *DLBOFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableDLBOFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableDLBOFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDLBOFunctionSingleAllOfAttributes(val *DLBOFunctionSingleAllOfAttributes) *NullableDLBOFunctionSingleAllOfAttributes {
	return &NullableDLBOFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableDLBOFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDLBOFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
