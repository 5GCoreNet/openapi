/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
)

// checks if the ManagedElementAttr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedElementAttr{}

// ManagedElementAttr struct for ManagedElementAttr
type ManagedElementAttr struct {
	DnPrefix *string `json:"dnPrefix,omitempty"`
	ManagedElementTypeList []string `json:"managedElementTypeList,omitempty"`
	UserLabel *string `json:"userLabel,omitempty"`
	LocationName *string `json:"locationName,omitempty"`
	ManagedBy []string `json:"managedBy,omitempty"`
	VendorName *string `json:"vendorName,omitempty"`
	UserDefinedState *string `json:"userDefinedState,omitempty"`
	SwVersion *string `json:"swVersion,omitempty"`
	PriorityLabel *int32 `json:"priorityLabel,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	SupportedTraceMetrics []string `json:"supportedTraceMetrics,omitempty"`
}

// NewManagedElementAttr instantiates a new ManagedElementAttr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedElementAttr() *ManagedElementAttr {
	this := ManagedElementAttr{}
	return &this
}

// NewManagedElementAttrWithDefaults instantiates a new ManagedElementAttr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedElementAttrWithDefaults() *ManagedElementAttr {
	this := ManagedElementAttr{}
	return &this
}

// GetDnPrefix returns the DnPrefix field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetDnPrefix() string {
	if o == nil || IsNil(o.DnPrefix) {
		var ret string
		return ret
	}
	return *o.DnPrefix
}

// GetDnPrefixOk returns a tuple with the DnPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetDnPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.DnPrefix) {
		return nil, false
	}
	return o.DnPrefix, true
}

// HasDnPrefix returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasDnPrefix() bool {
	if o != nil && !IsNil(o.DnPrefix) {
		return true
	}

	return false
}

// SetDnPrefix gets a reference to the given string and assigns it to the DnPrefix field.
func (o *ManagedElementAttr) SetDnPrefix(v string) {
	o.DnPrefix = &v
}

// GetManagedElementTypeList returns the ManagedElementTypeList field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetManagedElementTypeList() []string {
	if o == nil || IsNil(o.ManagedElementTypeList) {
		var ret []string
		return ret
	}
	return o.ManagedElementTypeList
}

// GetManagedElementTypeListOk returns a tuple with the ManagedElementTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetManagedElementTypeListOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagedElementTypeList) {
		return nil, false
	}
	return o.ManagedElementTypeList, true
}

// HasManagedElementTypeList returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasManagedElementTypeList() bool {
	if o != nil && !IsNil(o.ManagedElementTypeList) {
		return true
	}

	return false
}

// SetManagedElementTypeList gets a reference to the given []string and assigns it to the ManagedElementTypeList field.
func (o *ManagedElementAttr) SetManagedElementTypeList(v []string) {
	o.ManagedElementTypeList = v
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ManagedElementAttr) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetLocationName returns the LocationName field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetLocationName() string {
	if o == nil || IsNil(o.LocationName) {
		var ret string
		return ret
	}
	return *o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetLocationNameOk() (*string, bool) {
	if o == nil || IsNil(o.LocationName) {
		return nil, false
	}
	return o.LocationName, true
}

// HasLocationName returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasLocationName() bool {
	if o != nil && !IsNil(o.LocationName) {
		return true
	}

	return false
}

// SetLocationName gets a reference to the given string and assigns it to the LocationName field.
func (o *ManagedElementAttr) SetLocationName(v string) {
	o.LocationName = &v
}

// GetManagedBy returns the ManagedBy field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetManagedBy() []string {
	if o == nil || IsNil(o.ManagedBy) {
		var ret []string
		return ret
	}
	return o.ManagedBy
}

// GetManagedByOk returns a tuple with the ManagedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetManagedByOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagedBy) {
		return nil, false
	}
	return o.ManagedBy, true
}

// HasManagedBy returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasManagedBy() bool {
	if o != nil && !IsNil(o.ManagedBy) {
		return true
	}

	return false
}

// SetManagedBy gets a reference to the given []string and assigns it to the ManagedBy field.
func (o *ManagedElementAttr) SetManagedBy(v []string) {
	o.ManagedBy = v
}

// GetVendorName returns the VendorName field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetVendorName() string {
	if o == nil || IsNil(o.VendorName) {
		var ret string
		return ret
	}
	return *o.VendorName
}

// GetVendorNameOk returns a tuple with the VendorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetVendorNameOk() (*string, bool) {
	if o == nil || IsNil(o.VendorName) {
		return nil, false
	}
	return o.VendorName, true
}

// HasVendorName returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasVendorName() bool {
	if o != nil && !IsNil(o.VendorName) {
		return true
	}

	return false
}

// SetVendorName gets a reference to the given string and assigns it to the VendorName field.
func (o *ManagedElementAttr) SetVendorName(v string) {
	o.VendorName = &v
}

// GetUserDefinedState returns the UserDefinedState field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetUserDefinedState() string {
	if o == nil || IsNil(o.UserDefinedState) {
		var ret string
		return ret
	}
	return *o.UserDefinedState
}

// GetUserDefinedStateOk returns a tuple with the UserDefinedState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetUserDefinedStateOk() (*string, bool) {
	if o == nil || IsNil(o.UserDefinedState) {
		return nil, false
	}
	return o.UserDefinedState, true
}

// HasUserDefinedState returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasUserDefinedState() bool {
	if o != nil && !IsNil(o.UserDefinedState) {
		return true
	}

	return false
}

// SetUserDefinedState gets a reference to the given string and assigns it to the UserDefinedState field.
func (o *ManagedElementAttr) SetUserDefinedState(v string) {
	o.UserDefinedState = &v
}

// GetSwVersion returns the SwVersion field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetSwVersion() string {
	if o == nil || IsNil(o.SwVersion) {
		var ret string
		return ret
	}
	return *o.SwVersion
}

// GetSwVersionOk returns a tuple with the SwVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetSwVersionOk() (*string, bool) {
	if o == nil || IsNil(o.SwVersion) {
		return nil, false
	}
	return o.SwVersion, true
}

// HasSwVersion returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasSwVersion() bool {
	if o != nil && !IsNil(o.SwVersion) {
		return true
	}

	return false
}

// SetSwVersion gets a reference to the given string and assigns it to the SwVersion field.
func (o *ManagedElementAttr) SetSwVersion(v string) {
	o.SwVersion = &v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *ManagedElementAttr) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *ManagedElementAttr) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetSupportedTraceMetrics returns the SupportedTraceMetrics field value if set, zero value otherwise.
func (o *ManagedElementAttr) GetSupportedTraceMetrics() []string {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		var ret []string
		return ret
	}
	return o.SupportedTraceMetrics
}

// GetSupportedTraceMetricsOk returns a tuple with the SupportedTraceMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedElementAttr) GetSupportedTraceMetricsOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedTraceMetrics) {
		return nil, false
	}
	return o.SupportedTraceMetrics, true
}

// HasSupportedTraceMetrics returns a boolean if a field has been set.
func (o *ManagedElementAttr) HasSupportedTraceMetrics() bool {
	if o != nil && !IsNil(o.SupportedTraceMetrics) {
		return true
	}

	return false
}

// SetSupportedTraceMetrics gets a reference to the given []string and assigns it to the SupportedTraceMetrics field.
func (o *ManagedElementAttr) SetSupportedTraceMetrics(v []string) {
	o.SupportedTraceMetrics = v
}

func (o ManagedElementAttr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedElementAttr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnPrefix) {
		toSerialize["dnPrefix"] = o.DnPrefix
	}
	if !IsNil(o.ManagedElementTypeList) {
		toSerialize["managedElementTypeList"] = o.ManagedElementTypeList
	}
	if !IsNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !IsNil(o.LocationName) {
		toSerialize["locationName"] = o.LocationName
	}
	if !IsNil(o.ManagedBy) {
		toSerialize["managedBy"] = o.ManagedBy
	}
	if !IsNil(o.VendorName) {
		toSerialize["vendorName"] = o.VendorName
	}
	if !IsNil(o.UserDefinedState) {
		toSerialize["userDefinedState"] = o.UserDefinedState
	}
	if !IsNil(o.SwVersion) {
		toSerialize["swVersion"] = o.SwVersion
	}
	if !IsNil(o.PriorityLabel) {
		toSerialize["priorityLabel"] = o.PriorityLabel
	}
	if !IsNil(o.SupportedPerfMetricGroups) {
		toSerialize["supportedPerfMetricGroups"] = o.SupportedPerfMetricGroups
	}
	if !IsNil(o.SupportedTraceMetrics) {
		toSerialize["supportedTraceMetrics"] = o.SupportedTraceMetrics
	}
	return toSerialize, nil
}

type NullableManagedElementAttr struct {
	value *ManagedElementAttr
	isSet bool
}

func (v NullableManagedElementAttr) Get() *ManagedElementAttr {
	return v.value
}

func (v *NullableManagedElementAttr) Set(val *ManagedElementAttr) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedElementAttr) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedElementAttr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedElementAttr(val *ManagedElementAttr) *NullableManagedElementAttr {
	return &NullableManagedElementAttr{value: val, isSet: true}
}

func (v NullableManagedElementAttr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedElementAttr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


