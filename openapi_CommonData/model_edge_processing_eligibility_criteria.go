/*
5GMS Common Data Types

5GMS Common Data Types © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// checks if the EdgeProcessingEligibilityCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EdgeProcessingEligibilityCriteria{}

// EdgeProcessingEligibilityCriteria struct for EdgeProcessingEligibilityCriteria
type EdgeProcessingEligibilityCriteria struct {
	ServiceDataFlowDescriptions []ServiceDataFlowDescription `json:"serviceDataFlowDescriptions"`
	UeLocations []LocationArea5G `json:"ueLocations"`
	TimeWindows []TimeWindow `json:"timeWindows"`
	AppRequest bool `json:"appRequest"`
}

// NewEdgeProcessingEligibilityCriteria instantiates a new EdgeProcessingEligibilityCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEdgeProcessingEligibilityCriteria(serviceDataFlowDescriptions []ServiceDataFlowDescription, ueLocations []LocationArea5G, timeWindows []TimeWindow, appRequest bool) *EdgeProcessingEligibilityCriteria {
	this := EdgeProcessingEligibilityCriteria{}
	this.ServiceDataFlowDescriptions = serviceDataFlowDescriptions
	this.UeLocations = ueLocations
	this.TimeWindows = timeWindows
	this.AppRequest = appRequest
	return &this
}

// NewEdgeProcessingEligibilityCriteriaWithDefaults instantiates a new EdgeProcessingEligibilityCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEdgeProcessingEligibilityCriteriaWithDefaults() *EdgeProcessingEligibilityCriteria {
	this := EdgeProcessingEligibilityCriteria{}
	return &this
}

// GetServiceDataFlowDescriptions returns the ServiceDataFlowDescriptions field value
func (o *EdgeProcessingEligibilityCriteria) GetServiceDataFlowDescriptions() []ServiceDataFlowDescription {
	if o == nil {
		var ret []ServiceDataFlowDescription
		return ret
	}

	return o.ServiceDataFlowDescriptions
}

// GetServiceDataFlowDescriptionsOk returns a tuple with the ServiceDataFlowDescriptions field value
// and a boolean to check if the value has been set.
func (o *EdgeProcessingEligibilityCriteria) GetServiceDataFlowDescriptionsOk() ([]ServiceDataFlowDescription, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServiceDataFlowDescriptions, true
}

// SetServiceDataFlowDescriptions sets field value
func (o *EdgeProcessingEligibilityCriteria) SetServiceDataFlowDescriptions(v []ServiceDataFlowDescription) {
	o.ServiceDataFlowDescriptions = v
}

// GetUeLocations returns the UeLocations field value
func (o *EdgeProcessingEligibilityCriteria) GetUeLocations() []LocationArea5G {
	if o == nil {
		var ret []LocationArea5G
		return ret
	}

	return o.UeLocations
}

// GetUeLocationsOk returns a tuple with the UeLocations field value
// and a boolean to check if the value has been set.
func (o *EdgeProcessingEligibilityCriteria) GetUeLocationsOk() ([]LocationArea5G, bool) {
	if o == nil {
		return nil, false
	}
	return o.UeLocations, true
}

// SetUeLocations sets field value
func (o *EdgeProcessingEligibilityCriteria) SetUeLocations(v []LocationArea5G) {
	o.UeLocations = v
}

// GetTimeWindows returns the TimeWindows field value
func (o *EdgeProcessingEligibilityCriteria) GetTimeWindows() []TimeWindow {
	if o == nil {
		var ret []TimeWindow
		return ret
	}

	return o.TimeWindows
}

// GetTimeWindowsOk returns a tuple with the TimeWindows field value
// and a boolean to check if the value has been set.
func (o *EdgeProcessingEligibilityCriteria) GetTimeWindowsOk() ([]TimeWindow, bool) {
	if o == nil {
		return nil, false
	}
	return o.TimeWindows, true
}

// SetTimeWindows sets field value
func (o *EdgeProcessingEligibilityCriteria) SetTimeWindows(v []TimeWindow) {
	o.TimeWindows = v
}

// GetAppRequest returns the AppRequest field value
func (o *EdgeProcessingEligibilityCriteria) GetAppRequest() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AppRequest
}

// GetAppRequestOk returns a tuple with the AppRequest field value
// and a boolean to check if the value has been set.
func (o *EdgeProcessingEligibilityCriteria) GetAppRequestOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppRequest, true
}

// SetAppRequest sets field value
func (o *EdgeProcessingEligibilityCriteria) SetAppRequest(v bool) {
	o.AppRequest = v
}

func (o EdgeProcessingEligibilityCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EdgeProcessingEligibilityCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serviceDataFlowDescriptions"] = o.ServiceDataFlowDescriptions
	toSerialize["ueLocations"] = o.UeLocations
	toSerialize["timeWindows"] = o.TimeWindows
	toSerialize["appRequest"] = o.AppRequest
	return toSerialize, nil
}

type NullableEdgeProcessingEligibilityCriteria struct {
	value *EdgeProcessingEligibilityCriteria
	isSet bool
}

func (v NullableEdgeProcessingEligibilityCriteria) Get() *EdgeProcessingEligibilityCriteria {
	return v.value
}

func (v *NullableEdgeProcessingEligibilityCriteria) Set(val *EdgeProcessingEligibilityCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableEdgeProcessingEligibilityCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableEdgeProcessingEligibilityCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEdgeProcessingEligibilityCriteria(val *EdgeProcessingEligibilityCriteria) *NullableEdgeProcessingEligibilityCriteria {
	return &NullableEdgeProcessingEligibilityCriteria{value: val, isSet: true}
}

func (v NullableEdgeProcessingEligibilityCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEdgeProcessingEligibilityCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


