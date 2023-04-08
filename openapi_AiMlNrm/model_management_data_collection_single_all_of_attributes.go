/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
)

// checks if the ManagementDataCollectionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagementDataCollectionSingleAllOfAttributes{}

// ManagementDataCollectionSingleAllOfAttributes struct for ManagementDataCollectionSingleAllOfAttributes
type ManagementDataCollectionSingleAllOfAttributes struct {
	ManagementData *ManagementData `json:"managementData,omitempty"`
	TargetNodeFilter *NodeFilter `json:"targetNodeFilter,omitempty"`
	CollectionTimeWindow *TimeWindow `json:"collectionTimeWindow,omitempty"`
	ReportingCtrl *ReportingCtrl `json:"reportingCtrl,omitempty"`
	DataScope *string `json:"dataScope,omitempty"`
}

// NewManagementDataCollectionSingleAllOfAttributes instantiates a new ManagementDataCollectionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagementDataCollectionSingleAllOfAttributes() *ManagementDataCollectionSingleAllOfAttributes {
	this := ManagementDataCollectionSingleAllOfAttributes{}
	return &this
}

// NewManagementDataCollectionSingleAllOfAttributesWithDefaults instantiates a new ManagementDataCollectionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagementDataCollectionSingleAllOfAttributesWithDefaults() *ManagementDataCollectionSingleAllOfAttributes {
	this := ManagementDataCollectionSingleAllOfAttributes{}
	return &this
}

// GetManagementData returns the ManagementData field value if set, zero value otherwise.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetManagementData() ManagementData {
	if o == nil || isNil(o.ManagementData) {
		var ret ManagementData
		return ret
	}
	return *o.ManagementData
}

// GetManagementDataOk returns a tuple with the ManagementData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetManagementDataOk() (*ManagementData, bool) {
	if o == nil || isNil(o.ManagementData) {
		return nil, false
	}
	return o.ManagementData, true
}

// HasManagementData returns a boolean if a field has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) HasManagementData() bool {
	if o != nil && !isNil(o.ManagementData) {
		return true
	}

	return false
}

// SetManagementData gets a reference to the given ManagementData and assigns it to the ManagementData field.
func (o *ManagementDataCollectionSingleAllOfAttributes) SetManagementData(v ManagementData) {
	o.ManagementData = &v
}

// GetTargetNodeFilter returns the TargetNodeFilter field value if set, zero value otherwise.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetTargetNodeFilter() NodeFilter {
	if o == nil || isNil(o.TargetNodeFilter) {
		var ret NodeFilter
		return ret
	}
	return *o.TargetNodeFilter
}

// GetTargetNodeFilterOk returns a tuple with the TargetNodeFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetTargetNodeFilterOk() (*NodeFilter, bool) {
	if o == nil || isNil(o.TargetNodeFilter) {
		return nil, false
	}
	return o.TargetNodeFilter, true
}

// HasTargetNodeFilter returns a boolean if a field has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) HasTargetNodeFilter() bool {
	if o != nil && !isNil(o.TargetNodeFilter) {
		return true
	}

	return false
}

// SetTargetNodeFilter gets a reference to the given NodeFilter and assigns it to the TargetNodeFilter field.
func (o *ManagementDataCollectionSingleAllOfAttributes) SetTargetNodeFilter(v NodeFilter) {
	o.TargetNodeFilter = &v
}

// GetCollectionTimeWindow returns the CollectionTimeWindow field value if set, zero value otherwise.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetCollectionTimeWindow() TimeWindow {
	if o == nil || isNil(o.CollectionTimeWindow) {
		var ret TimeWindow
		return ret
	}
	return *o.CollectionTimeWindow
}

// GetCollectionTimeWindowOk returns a tuple with the CollectionTimeWindow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetCollectionTimeWindowOk() (*TimeWindow, bool) {
	if o == nil || isNil(o.CollectionTimeWindow) {
		return nil, false
	}
	return o.CollectionTimeWindow, true
}

// HasCollectionTimeWindow returns a boolean if a field has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) HasCollectionTimeWindow() bool {
	if o != nil && !isNil(o.CollectionTimeWindow) {
		return true
	}

	return false
}

// SetCollectionTimeWindow gets a reference to the given TimeWindow and assigns it to the CollectionTimeWindow field.
func (o *ManagementDataCollectionSingleAllOfAttributes) SetCollectionTimeWindow(v TimeWindow) {
	o.CollectionTimeWindow = &v
}

// GetReportingCtrl returns the ReportingCtrl field value if set, zero value otherwise.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetReportingCtrl() ReportingCtrl {
	if o == nil || isNil(o.ReportingCtrl) {
		var ret ReportingCtrl
		return ret
	}
	return *o.ReportingCtrl
}

// GetReportingCtrlOk returns a tuple with the ReportingCtrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetReportingCtrlOk() (*ReportingCtrl, bool) {
	if o == nil || isNil(o.ReportingCtrl) {
		return nil, false
	}
	return o.ReportingCtrl, true
}

// HasReportingCtrl returns a boolean if a field has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) HasReportingCtrl() bool {
	if o != nil && !isNil(o.ReportingCtrl) {
		return true
	}

	return false
}

// SetReportingCtrl gets a reference to the given ReportingCtrl and assigns it to the ReportingCtrl field.
func (o *ManagementDataCollectionSingleAllOfAttributes) SetReportingCtrl(v ReportingCtrl) {
	o.ReportingCtrl = &v
}

// GetDataScope returns the DataScope field value if set, zero value otherwise.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetDataScope() string {
	if o == nil || isNil(o.DataScope) {
		var ret string
		return ret
	}
	return *o.DataScope
}

// GetDataScopeOk returns a tuple with the DataScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) GetDataScopeOk() (*string, bool) {
	if o == nil || isNil(o.DataScope) {
		return nil, false
	}
	return o.DataScope, true
}

// HasDataScope returns a boolean if a field has been set.
func (o *ManagementDataCollectionSingleAllOfAttributes) HasDataScope() bool {
	if o != nil && !isNil(o.DataScope) {
		return true
	}

	return false
}

// SetDataScope gets a reference to the given string and assigns it to the DataScope field.
func (o *ManagementDataCollectionSingleAllOfAttributes) SetDataScope(v string) {
	o.DataScope = &v
}

func (o ManagementDataCollectionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagementDataCollectionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ManagementData) {
		toSerialize["managementData"] = o.ManagementData
	}
	if !isNil(o.TargetNodeFilter) {
		toSerialize["targetNodeFilter"] = o.TargetNodeFilter
	}
	if !isNil(o.CollectionTimeWindow) {
		toSerialize["collectionTimeWindow"] = o.CollectionTimeWindow
	}
	if !isNil(o.ReportingCtrl) {
		toSerialize["reportingCtrl"] = o.ReportingCtrl
	}
	if !isNil(o.DataScope) {
		toSerialize["dataScope"] = o.DataScope
	}
	return toSerialize, nil
}

type NullableManagementDataCollectionSingleAllOfAttributes struct {
	value *ManagementDataCollectionSingleAllOfAttributes
	isSet bool
}

func (v NullableManagementDataCollectionSingleAllOfAttributes) Get() *ManagementDataCollectionSingleAllOfAttributes {
	return v.value
}

func (v *NullableManagementDataCollectionSingleAllOfAttributes) Set(val *ManagementDataCollectionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementDataCollectionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementDataCollectionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementDataCollectionSingleAllOfAttributes(val *ManagementDataCollectionSingleAllOfAttributes) *NullableManagementDataCollectionSingleAllOfAttributes {
	return &NullableManagementDataCollectionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableManagementDataCollectionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementDataCollectionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


