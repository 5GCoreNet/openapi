/*
AMInfluence

AMInfluence API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AMInfluence

import (
	"encoding/json"
)

// checks if the AmInfluSubPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmInfluSubPatch{}

// AmInfluSubPatch Represents parameters to request the modification of an AM influence subscription resource.
type AmInfluSubPatch struct {
	HighThruInd NullableBool `json:"highThruInd,omitempty"`
	// Identifies geographic areas of the user where the request is applicable.
	GeoAreas []GeographicArea `json:"geoAreas,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds with \"nullable=true\" property.
	PolicyDuration NullableInt32 `json:"policyDuration,omitempty"`
	// Each of the element identifies a (DNN, S-NSSAI) combination.
	DnnSnssaiInfos []DnnSnssaiInformation `json:"dnnSnssaiInfos,omitempty"`
	// Each of the element identifies an application.
	AfAppIds []string `json:"afAppIds,omitempty"`
	// Indicates one or more AM influence related events.
	SubscribedEvents []AmInfluEvent `json:"subscribedEvents,omitempty"`
	// String formatted according to IETF RFC 3986 identifying a referenced resource, but with the nullable property set to true.
	NotificationDestination NullableString `json:"notificationDestination,omitempty"`
}

// NewAmInfluSubPatch instantiates a new AmInfluSubPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmInfluSubPatch() *AmInfluSubPatch {
	this := AmInfluSubPatch{}
	return &this
}

// NewAmInfluSubPatchWithDefaults instantiates a new AmInfluSubPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmInfluSubPatchWithDefaults() *AmInfluSubPatch {
	this := AmInfluSubPatch{}
	return &this
}

// GetHighThruInd returns the HighThruInd field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluSubPatch) GetHighThruInd() bool {
	if o == nil || IsNil(o.HighThruInd.Get()) {
		var ret bool
		return ret
	}
	return *o.HighThruInd.Get()
}

// GetHighThruIndOk returns a tuple with the HighThruInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluSubPatch) GetHighThruIndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.HighThruInd.Get(), o.HighThruInd.IsSet()
}

// HasHighThruInd returns a boolean if a field has been set.
func (o *AmInfluSubPatch) HasHighThruInd() bool {
	if o != nil && o.HighThruInd.IsSet() {
		return true
	}

	return false
}

// SetHighThruInd gets a reference to the given NullableBool and assigns it to the HighThruInd field.
func (o *AmInfluSubPatch) SetHighThruInd(v bool) {
	o.HighThruInd.Set(&v)
}

// SetHighThruIndNil sets the value for HighThruInd to be an explicit nil
func (o *AmInfluSubPatch) SetHighThruIndNil() {
	o.HighThruInd.Set(nil)
}

// UnsetHighThruInd ensures that no value is present for HighThruInd, not even an explicit nil
func (o *AmInfluSubPatch) UnsetHighThruInd() {
	o.HighThruInd.Unset()
}

// GetGeoAreas returns the GeoAreas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluSubPatch) GetGeoAreas() []GeographicArea {
	if o == nil {
		var ret []GeographicArea
		return ret
	}
	return o.GeoAreas
}

// GetGeoAreasOk returns a tuple with the GeoAreas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluSubPatch) GetGeoAreasOk() ([]GeographicArea, bool) {
	if o == nil || IsNil(o.GeoAreas) {
		return nil, false
	}
	return o.GeoAreas, true
}

// HasGeoAreas returns a boolean if a field has been set.
func (o *AmInfluSubPatch) HasGeoAreas() bool {
	if o != nil && IsNil(o.GeoAreas) {
		return true
	}

	return false
}

// SetGeoAreas gets a reference to the given []GeographicArea and assigns it to the GeoAreas field.
func (o *AmInfluSubPatch) SetGeoAreas(v []GeographicArea) {
	o.GeoAreas = v
}

// GetPolicyDuration returns the PolicyDuration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluSubPatch) GetPolicyDuration() int32 {
	if o == nil || IsNil(o.PolicyDuration.Get()) {
		var ret int32
		return ret
	}
	return *o.PolicyDuration.Get()
}

// GetPolicyDurationOk returns a tuple with the PolicyDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluSubPatch) GetPolicyDurationOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PolicyDuration.Get(), o.PolicyDuration.IsSet()
}

// HasPolicyDuration returns a boolean if a field has been set.
func (o *AmInfluSubPatch) HasPolicyDuration() bool {
	if o != nil && o.PolicyDuration.IsSet() {
		return true
	}

	return false
}

// SetPolicyDuration gets a reference to the given NullableInt32 and assigns it to the PolicyDuration field.
func (o *AmInfluSubPatch) SetPolicyDuration(v int32) {
	o.PolicyDuration.Set(&v)
}

// SetPolicyDurationNil sets the value for PolicyDuration to be an explicit nil
func (o *AmInfluSubPatch) SetPolicyDurationNil() {
	o.PolicyDuration.Set(nil)
}

// UnsetPolicyDuration ensures that no value is present for PolicyDuration, not even an explicit nil
func (o *AmInfluSubPatch) UnsetPolicyDuration() {
	o.PolicyDuration.Unset()
}

// GetDnnSnssaiInfos returns the DnnSnssaiInfos field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluSubPatch) GetDnnSnssaiInfos() []DnnSnssaiInformation {
	if o == nil {
		var ret []DnnSnssaiInformation
		return ret
	}
	return o.DnnSnssaiInfos
}

// GetDnnSnssaiInfosOk returns a tuple with the DnnSnssaiInfos field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluSubPatch) GetDnnSnssaiInfosOk() ([]DnnSnssaiInformation, bool) {
	if o == nil || IsNil(o.DnnSnssaiInfos) {
		return nil, false
	}
	return o.DnnSnssaiInfos, true
}

// HasDnnSnssaiInfos returns a boolean if a field has been set.
func (o *AmInfluSubPatch) HasDnnSnssaiInfos() bool {
	if o != nil && IsNil(o.DnnSnssaiInfos) {
		return true
	}

	return false
}

// SetDnnSnssaiInfos gets a reference to the given []DnnSnssaiInformation and assigns it to the DnnSnssaiInfos field.
func (o *AmInfluSubPatch) SetDnnSnssaiInfos(v []DnnSnssaiInformation) {
	o.DnnSnssaiInfos = v
}

// GetAfAppIds returns the AfAppIds field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluSubPatch) GetAfAppIds() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AfAppIds
}

// GetAfAppIdsOk returns a tuple with the AfAppIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluSubPatch) GetAfAppIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AfAppIds) {
		return nil, false
	}
	return o.AfAppIds, true
}

// HasAfAppIds returns a boolean if a field has been set.
func (o *AmInfluSubPatch) HasAfAppIds() bool {
	if o != nil && IsNil(o.AfAppIds) {
		return true
	}

	return false
}

// SetAfAppIds gets a reference to the given []string and assigns it to the AfAppIds field.
func (o *AmInfluSubPatch) SetAfAppIds(v []string) {
	o.AfAppIds = v
}

// GetSubscribedEvents returns the SubscribedEvents field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluSubPatch) GetSubscribedEvents() []AmInfluEvent {
	if o == nil {
		var ret []AmInfluEvent
		return ret
	}
	return o.SubscribedEvents
}

// GetSubscribedEventsOk returns a tuple with the SubscribedEvents field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluSubPatch) GetSubscribedEventsOk() ([]AmInfluEvent, bool) {
	if o == nil || IsNil(o.SubscribedEvents) {
		return nil, false
	}
	return o.SubscribedEvents, true
}

// HasSubscribedEvents returns a boolean if a field has been set.
func (o *AmInfluSubPatch) HasSubscribedEvents() bool {
	if o != nil && IsNil(o.SubscribedEvents) {
		return true
	}

	return false
}

// SetSubscribedEvents gets a reference to the given []AmInfluEvent and assigns it to the SubscribedEvents field.
func (o *AmInfluSubPatch) SetSubscribedEvents(v []AmInfluEvent) {
	o.SubscribedEvents = v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AmInfluSubPatch) GetNotificationDestination() string {
	if o == nil || IsNil(o.NotificationDestination.Get()) {
		var ret string
		return ret
	}
	return *o.NotificationDestination.Get()
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AmInfluSubPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NotificationDestination.Get(), o.NotificationDestination.IsSet()
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *AmInfluSubPatch) HasNotificationDestination() bool {
	if o != nil && o.NotificationDestination.IsSet() {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given NullableString and assigns it to the NotificationDestination field.
func (o *AmInfluSubPatch) SetNotificationDestination(v string) {
	o.NotificationDestination.Set(&v)
}

// SetNotificationDestinationNil sets the value for NotificationDestination to be an explicit nil
func (o *AmInfluSubPatch) SetNotificationDestinationNil() {
	o.NotificationDestination.Set(nil)
}

// UnsetNotificationDestination ensures that no value is present for NotificationDestination, not even an explicit nil
func (o *AmInfluSubPatch) UnsetNotificationDestination() {
	o.NotificationDestination.Unset()
}

func (o AmInfluSubPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmInfluSubPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.HighThruInd.IsSet() {
		toSerialize["highThruInd"] = o.HighThruInd.Get()
	}
	if o.GeoAreas != nil {
		toSerialize["geoAreas"] = o.GeoAreas
	}
	if o.PolicyDuration.IsSet() {
		toSerialize["policyDuration"] = o.PolicyDuration.Get()
	}
	if o.DnnSnssaiInfos != nil {
		toSerialize["dnnSnssaiInfos"] = o.DnnSnssaiInfos
	}
	if o.AfAppIds != nil {
		toSerialize["afAppIds"] = o.AfAppIds
	}
	if o.SubscribedEvents != nil {
		toSerialize["subscribedEvents"] = o.SubscribedEvents
	}
	if o.NotificationDestination.IsSet() {
		toSerialize["notificationDestination"] = o.NotificationDestination.Get()
	}
	return toSerialize, nil
}

type NullableAmInfluSubPatch struct {
	value *AmInfluSubPatch
	isSet bool
}

func (v NullableAmInfluSubPatch) Get() *AmInfluSubPatch {
	return v.value
}

func (v *NullableAmInfluSubPatch) Set(val *AmInfluSubPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableAmInfluSubPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableAmInfluSubPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmInfluSubPatch(val *AmInfluSubPatch) *NullableAmInfluSubPatch {
	return &NullableAmInfluSubPatch{value: val, isSet: true}
}

func (v NullableAmInfluSubPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmInfluSubPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
