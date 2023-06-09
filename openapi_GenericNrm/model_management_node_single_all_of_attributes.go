/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
)

// checks if the ManagementNodeSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementNodeSingleAllOfAttributes{}

// ManagementNodeSingleAllOfAttributes struct for ManagementNodeSingleAllOfAttributes
type ManagementNodeSingleAllOfAttributes struct {
	UserLabel        *string  `json:"userLabel,omitempty"`
	ManagedElements  []string `json:"managedElements,omitempty"`
	VendorName       *string  `json:"vendorName,omitempty"`
	UserDefinedState *string  `json:"userDefinedState,omitempty"`
	LocationName     *string  `json:"locationName,omitempty"`
	SwVersion        *string  `json:"swVersion,omitempty"`
}

// NewManagementNodeSingleAllOfAttributes instantiates a new ManagementNodeSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementNodeSingleAllOfAttributes() *ManagementNodeSingleAllOfAttributes {
	this := ManagementNodeSingleAllOfAttributes{}
	return &this
}

// NewManagementNodeSingleAllOfAttributesWithDefaults instantiates a new ManagementNodeSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementNodeSingleAllOfAttributesWithDefaults() *ManagementNodeSingleAllOfAttributes {
	this := ManagementNodeSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ManagementNodeSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ManagementNodeSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ManagementNodeSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetManagedElements returns the ManagedElements field value if set, zero value otherwise.
func (o *ManagementNodeSingleAllOfAttributes) GetManagedElements() []string {
	if o == nil || IsNil(o.ManagedElements) {
		var ret []string
		return ret
	}
	return o.ManagedElements
}

// GetManagedElementsOk returns a tuple with the ManagedElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingleAllOfAttributes) GetManagedElementsOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagedElements) {
		return nil, false
	}
	return o.ManagedElements, true
}

// HasManagedElements returns a boolean if a field has been set.
func (o *ManagementNodeSingleAllOfAttributes) HasManagedElements() bool {
	if o != nil && !IsNil(o.ManagedElements) {
		return true
	}

	return false
}

// SetManagedElements gets a reference to the given []string and assigns it to the ManagedElements field.
func (o *ManagementNodeSingleAllOfAttributes) SetManagedElements(v []string) {
	o.ManagedElements = v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *ManagementNodeSingleAllOfAttributes) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingleAllOfAttributes) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *ManagementNodeSingleAllOfAttributes) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *ManagementNodeSingleAllOfAttributes) SetVendorName(v string) {
	o.VendorName = &v
}

// GetUserDefinedState returns the UserDefinedState field value if set, zero value otherwise.
func (o *ManagementNodeSingleAllOfAttributes) GetUserDefinedState() string {
	if o == nil || IsNil(o.UserDefinedState) {
		var ret string
		return ret
	}
	return *o.UserDefinedState
}

// GetUserDefinedStateOk returns a tuple with the UserDefinedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingleAllOfAttributes) GetUserDefinedStateOk() (*string, bool) {
	if o == nil || IsNil(o.UserDefinedState) {
		return nil, false
	}
	return o.UserDefinedState, true
}

// HasUserDefinedState returns a boolean if a field has been set.
func (o *ManagementNodeSingleAllOfAttributes) HasUserDefinedState() bool {
	if o != nil && !IsNil(o.UserDefinedState) {
		return true
	}

	return false
}

// SetUserDefinedState gets a reference to the given string and assigns it to the UserDefinedState field.
func (o *ManagementNodeSingleAllOfAttributes) SetUserDefinedState(v string) {
	o.UserDefinedState = &v
}

// GetLocationName returns the LocationName field value if set, zero value otherwise.
func (o *ManagementNodeSingleAllOfAttributes) GetLocationName() string {
	if o == nil || IsNil(o.LocationName) {
		var ret string
		return ret
	}
	return *o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingleAllOfAttributes) GetLocationNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocationName) {
		return nil, false
	}
	return o.LocationName, true
}

// HasLocationName returns a boolean if a field has been set.
func (o *ManagementNodeSingleAllOfAttributes) HasLocationName() bool {
	if o != nil && !IsNil(o.LocationName) {
		return true
	}

	return false
}

// SetLocationName gets a reference to the given string and assigns it to the LocationName field.
func (o *ManagementNodeSingleAllOfAttributes) SetLocationName(v string) {
	o.LocationName = &v
}

// GetSwVersion returns the SwVersion field value if set, zero value otherwise.
func (o *ManagementNodeSingleAllOfAttributes) GetSwVersion() string {
	if o == nil || IsNil(o.SwVersion) {
		var ret string
		return ret
	}
	return *o.SwVersion
}

// GetSwVersionOk returns a tuple with the SwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementNodeSingleAllOfAttributes) GetSwVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SwVersion) {
		return nil, false
	}
	return o.SwVersion, true
}

// HasSwVersion returns a boolean if a field has been set.
func (o *ManagementNodeSingleAllOfAttributes) HasSwVersion() bool {
	if o != nil && !IsNil(o.SwVersion) {
		return true
	}

	return false
}

// SetSwVersion gets a reference to the given string and assigns it to the SwVersion field.
func (o *ManagementNodeSingleAllOfAttributes) SetSwVersion(v string) {
	o.SwVersion = &v
}

func (o ManagementNodeSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementNodeSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !IsNil(o.ManagedElements) {
		toSerialize["managedElements"] = o.ManagedElements
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendorName"] = o.VendorName
	}
	if !IsNil(o.UserDefinedState) {
		toSerialize["userDefinedState"] = o.UserDefinedState
	}
	if !IsNil(o.LocationName) {
		toSerialize["locationName"] = o.LocationName
	}
	if !IsNil(o.SwVersion) {
		toSerialize["swVersion"] = o.SwVersion
	}
	return toSerialize, nil
}

type NullableManagementNodeSingleAllOfAttributes struct {
	value *ManagementNodeSingleAllOfAttributes
	isSet bool
}

func (v NullableManagementNodeSingleAllOfAttributes) Get() *ManagementNodeSingleAllOfAttributes {
	return v.value
}

func (v *NullableManagementNodeSingleAllOfAttributes) Set(val *ManagementNodeSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementNodeSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementNodeSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementNodeSingleAllOfAttributes(val *ManagementNodeSingleAllOfAttributes) *NullableManagementNodeSingleAllOfAttributes {
	return &NullableManagementNodeSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableManagementNodeSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementNodeSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
