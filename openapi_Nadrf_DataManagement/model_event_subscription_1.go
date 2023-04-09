/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the EventSubscription1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventSubscription1{}

// EventSubscription1 Represents a subscription to a single event.
type EventSubscription1 struct {
	Event             SmfEvent               `json:"event"`
	DnaiChgType       *DnaiChangeType        `json:"dnaiChgType,omitempty"`
	DddTraDescriptors []DddTrafficDescriptor `json:"dddTraDescriptors,omitempty"`
	DddStati          []DlDataDeliveryStatus `json:"dddStati,omitempty"`
	AppIds            []string               `json:"appIds,omitempty"`
	TargetPeriod      *TimeWindow            `json:"targetPeriod,omitempty"`
	// Indicates the subscription for UE transaction dispersion collectionon, if it is included and set to \"true\". Default value is \"false\".
	TransacDispInd *bool `json:"transacDispInd,omitempty"`
	// Indicates Session Management Transaction metrics.
	TransacMetrics []TransactionMetric `json:"transacMetrics,omitempty"`
	UeIpAddr       *IpAddr             `json:"ueIpAddr,omitempty"`
}

// NewEventSubscription1 instantiates a new EventSubscription1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventSubscription1(event SmfEvent) *EventSubscription1 {
	this := EventSubscription1{}
	this.Event = event
	return &this
}

// NewEventSubscription1WithDefaults instantiates a new EventSubscription1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventSubscription1WithDefaults() *EventSubscription1 {
	this := EventSubscription1{}
	return &this
}

// GetEvent returns the Event field value
func (o *EventSubscription1) GetEvent() SmfEvent {
	if o == nil {
		var ret SmfEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetEventOk() (*SmfEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventSubscription1) SetEvent(v SmfEvent) {
	o.Event = v
}

// GetDnaiChgType returns the DnaiChgType field value if set, zero value otherwise.
func (o *EventSubscription1) GetDnaiChgType() DnaiChangeType {
	if o == nil || IsNil(o.DnaiChgType) {
		var ret DnaiChangeType
		return ret
	}
	return *o.DnaiChgType
}

// GetDnaiChgTypeOk returns a tuple with the DnaiChgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetDnaiChgTypeOk() (*DnaiChangeType, bool) {
	if o == nil || IsNil(o.DnaiChgType) {
		return nil, false
	}
	return o.DnaiChgType, true
}

// HasDnaiChgType returns a boolean if a field has been set.
func (o *EventSubscription1) HasDnaiChgType() bool {
	if o != nil && !IsNil(o.DnaiChgType) {
		return true
	}

	return false
}

// SetDnaiChgType gets a reference to the given DnaiChangeType and assigns it to the DnaiChgType field.
func (o *EventSubscription1) SetDnaiChgType(v DnaiChangeType) {
	o.DnaiChgType = &v
}

// GetDddTraDescriptors returns the DddTraDescriptors field value if set, zero value otherwise.
func (o *EventSubscription1) GetDddTraDescriptors() []DddTrafficDescriptor {
	if o == nil || IsNil(o.DddTraDescriptors) {
		var ret []DddTrafficDescriptor
		return ret
	}
	return o.DddTraDescriptors
}

// GetDddTraDescriptorsOk returns a tuple with the DddTraDescriptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetDddTraDescriptorsOk() ([]DddTrafficDescriptor, bool) {
	if o == nil || IsNil(o.DddTraDescriptors) {
		return nil, false
	}
	return o.DddTraDescriptors, true
}

// HasDddTraDescriptors returns a boolean if a field has been set.
func (o *EventSubscription1) HasDddTraDescriptors() bool {
	if o != nil && !IsNil(o.DddTraDescriptors) {
		return true
	}

	return false
}

// SetDddTraDescriptors gets a reference to the given []DddTrafficDescriptor and assigns it to the DddTraDescriptors field.
func (o *EventSubscription1) SetDddTraDescriptors(v []DddTrafficDescriptor) {
	o.DddTraDescriptors = v
}

// GetDddStati returns the DddStati field value if set, zero value otherwise.
func (o *EventSubscription1) GetDddStati() []DlDataDeliveryStatus {
	if o == nil || IsNil(o.DddStati) {
		var ret []DlDataDeliveryStatus
		return ret
	}
	return o.DddStati
}

// GetDddStatiOk returns a tuple with the DddStati field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetDddStatiOk() ([]DlDataDeliveryStatus, bool) {
	if o == nil || IsNil(o.DddStati) {
		return nil, false
	}
	return o.DddStati, true
}

// HasDddStati returns a boolean if a field has been set.
func (o *EventSubscription1) HasDddStati() bool {
	if o != nil && !IsNil(o.DddStati) {
		return true
	}

	return false
}

// SetDddStati gets a reference to the given []DlDataDeliveryStatus and assigns it to the DddStati field.
func (o *EventSubscription1) SetDddStati(v []DlDataDeliveryStatus) {
	o.DddStati = v
}

// GetAppIds returns the AppIds field value if set, zero value otherwise.
func (o *EventSubscription1) GetAppIds() []string {
	if o == nil || IsNil(o.AppIds) {
		var ret []string
		return ret
	}
	return o.AppIds
}

// GetAppIdsOk returns a tuple with the AppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetAppIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AppIds) {
		return nil, false
	}
	return o.AppIds, true
}

// HasAppIds returns a boolean if a field has been set.
func (o *EventSubscription1) HasAppIds() bool {
	if o != nil && !IsNil(o.AppIds) {
		return true
	}

	return false
}

// SetAppIds gets a reference to the given []string and assigns it to the AppIds field.
func (o *EventSubscription1) SetAppIds(v []string) {
	o.AppIds = v
}

// GetTargetPeriod returns the TargetPeriod field value if set, zero value otherwise.
func (o *EventSubscription1) GetTargetPeriod() TimeWindow {
	if o == nil || IsNil(o.TargetPeriod) {
		var ret TimeWindow
		return ret
	}
	return *o.TargetPeriod
}

// GetTargetPeriodOk returns a tuple with the TargetPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetTargetPeriodOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.TargetPeriod) {
		return nil, false
	}
	return o.TargetPeriod, true
}

// HasTargetPeriod returns a boolean if a field has been set.
func (o *EventSubscription1) HasTargetPeriod() bool {
	if o != nil && !IsNil(o.TargetPeriod) {
		return true
	}

	return false
}

// SetTargetPeriod gets a reference to the given TimeWindow and assigns it to the TargetPeriod field.
func (o *EventSubscription1) SetTargetPeriod(v TimeWindow) {
	o.TargetPeriod = &v
}

// GetTransacDispInd returns the TransacDispInd field value if set, zero value otherwise.
func (o *EventSubscription1) GetTransacDispInd() bool {
	if o == nil || IsNil(o.TransacDispInd) {
		var ret bool
		return ret
	}
	return *o.TransacDispInd
}

// GetTransacDispIndOk returns a tuple with the TransacDispInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetTransacDispIndOk() (*bool, bool) {
	if o == nil || IsNil(o.TransacDispInd) {
		return nil, false
	}
	return o.TransacDispInd, true
}

// HasTransacDispInd returns a boolean if a field has been set.
func (o *EventSubscription1) HasTransacDispInd() bool {
	if o != nil && !IsNil(o.TransacDispInd) {
		return true
	}

	return false
}

// SetTransacDispInd gets a reference to the given bool and assigns it to the TransacDispInd field.
func (o *EventSubscription1) SetTransacDispInd(v bool) {
	o.TransacDispInd = &v
}

// GetTransacMetrics returns the TransacMetrics field value if set, zero value otherwise.
func (o *EventSubscription1) GetTransacMetrics() []TransactionMetric {
	if o == nil || IsNil(o.TransacMetrics) {
		var ret []TransactionMetric
		return ret
	}
	return o.TransacMetrics
}

// GetTransacMetricsOk returns a tuple with the TransacMetrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetTransacMetricsOk() ([]TransactionMetric, bool) {
	if o == nil || IsNil(o.TransacMetrics) {
		return nil, false
	}
	return o.TransacMetrics, true
}

// HasTransacMetrics returns a boolean if a field has been set.
func (o *EventSubscription1) HasTransacMetrics() bool {
	if o != nil && !IsNil(o.TransacMetrics) {
		return true
	}

	return false
}

// SetTransacMetrics gets a reference to the given []TransactionMetric and assigns it to the TransacMetrics field.
func (o *EventSubscription1) SetTransacMetrics(v []TransactionMetric) {
	o.TransacMetrics = v
}

// GetUeIpAddr returns the UeIpAddr field value if set, zero value otherwise.
func (o *EventSubscription1) GetUeIpAddr() IpAddr {
	if o == nil || IsNil(o.UeIpAddr) {
		var ret IpAddr
		return ret
	}
	return *o.UeIpAddr
}

// GetUeIpAddrOk returns a tuple with the UeIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventSubscription1) GetUeIpAddrOk() (*IpAddr, bool) {
	if o == nil || IsNil(o.UeIpAddr) {
		return nil, false
	}
	return o.UeIpAddr, true
}

// HasUeIpAddr returns a boolean if a field has been set.
func (o *EventSubscription1) HasUeIpAddr() bool {
	if o != nil && !IsNil(o.UeIpAddr) {
		return true
	}

	return false
}

// SetUeIpAddr gets a reference to the given IpAddr and assigns it to the UeIpAddr field.
func (o *EventSubscription1) SetUeIpAddr(v IpAddr) {
	o.UeIpAddr = &v
}

func (o EventSubscription1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventSubscription1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.DnaiChgType) {
		toSerialize["dnaiChgType"] = o.DnaiChgType
	}
	if !IsNil(o.DddTraDescriptors) {
		toSerialize["dddTraDescriptors"] = o.DddTraDescriptors
	}
	if !IsNil(o.DddStati) {
		toSerialize["dddStati"] = o.DddStati
	}
	if !IsNil(o.AppIds) {
		toSerialize["appIds"] = o.AppIds
	}
	if !IsNil(o.TargetPeriod) {
		toSerialize["targetPeriod"] = o.TargetPeriod
	}
	if !IsNil(o.TransacDispInd) {
		toSerialize["transacDispInd"] = o.TransacDispInd
	}
	if !IsNil(o.TransacMetrics) {
		toSerialize["transacMetrics"] = o.TransacMetrics
	}
	if !IsNil(o.UeIpAddr) {
		toSerialize["ueIpAddr"] = o.UeIpAddr
	}
	return toSerialize, nil
}

type NullableEventSubscription1 struct {
	value *EventSubscription1
	isSet bool
}

func (v NullableEventSubscription1) Get() *EventSubscription1 {
	return v.value
}

func (v *NullableEventSubscription1) Set(val *EventSubscription1) {
	v.value = val
	v.isSet = true
}

func (v NullableEventSubscription1) IsSet() bool {
	return v.isSet
}

func (v *NullableEventSubscription1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventSubscription1(val *EventSubscription1) *NullableEventSubscription1 {
	return &NullableEventSubscription1{value: val, isSet: true}
}

func (v NullableEventSubscription1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventSubscription1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
