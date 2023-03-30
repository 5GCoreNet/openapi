/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the TrafficControlData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficControlData{}

// TrafficControlData struct for TrafficControlData
type TrafficControlData struct {
	TcId *string `json:"tcId,omitempty"`
	FlowStatus *FlowStatus `json:"flowStatus,omitempty"`
	RedirectInfo *RedirectInformation `json:"redirectInfo,omitempty"`
	AddRedirectInfo []RedirectInformation `json:"addRedirectInfo,omitempty"`
	MuteNotif *bool `json:"muteNotif,omitempty"`
	TrafficSteeringPolIdDl NullableString `json:"trafficSteeringPolIdDl,omitempty"`
	TrafficSteeringPolIdUl NullableString `json:"trafficSteeringPolIdUl,omitempty"`
	RouteToLocs []RouteToLocation `json:"routeToLocs,omitempty"`
	TraffCorreInd *bool `json:"traffCorreInd,omitempty"`
	UpPathChgEvent NullableUpPathChgEvent `json:"upPathChgEvent,omitempty"`
	SteerFun *SteeringFunctionality `json:"steerFun,omitempty"`
	SteerModeDl *SteeringMode `json:"steerModeDl,omitempty"`
	SteerModeUl *SteeringMode `json:"steerModeUl,omitempty"`
	MulAccCtrl *MulticastAccessControl `json:"mulAccCtrl,omitempty"`
	SnssaiList []Snssai `json:"snssaiList,omitempty"`
}

// NewTrafficControlData instantiates a new TrafficControlData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficControlData() *TrafficControlData {
	this := TrafficControlData{}
	return &this
}

// NewTrafficControlDataWithDefaults instantiates a new TrafficControlData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficControlDataWithDefaults() *TrafficControlData {
	this := TrafficControlData{}
	return &this
}

// GetTcId returns the TcId field value if set, zero value otherwise.
func (o *TrafficControlData) GetTcId() string {
	if o == nil || IsNil(o.TcId) {
		var ret string
		return ret
	}
	return *o.TcId
}

// GetTcIdOk returns a tuple with the TcId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetTcIdOk() (*string, bool) {
	if o == nil || IsNil(o.TcId) {
		return nil, false
	}
	return o.TcId, true
}

// HasTcId returns a boolean if a field has been set.
func (o *TrafficControlData) HasTcId() bool {
	if o != nil && !IsNil(o.TcId) {
		return true
	}

	return false
}

// SetTcId gets a reference to the given string and assigns it to the TcId field.
func (o *TrafficControlData) SetTcId(v string) {
	o.TcId = &v
}

// GetFlowStatus returns the FlowStatus field value if set, zero value otherwise.
func (o *TrafficControlData) GetFlowStatus() FlowStatus {
	if o == nil || IsNil(o.FlowStatus) {
		var ret FlowStatus
		return ret
	}
	return *o.FlowStatus
}

// GetFlowStatusOk returns a tuple with the FlowStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetFlowStatusOk() (*FlowStatus, bool) {
	if o == nil || IsNil(o.FlowStatus) {
		return nil, false
	}
	return o.FlowStatus, true
}

// HasFlowStatus returns a boolean if a field has been set.
func (o *TrafficControlData) HasFlowStatus() bool {
	if o != nil && !IsNil(o.FlowStatus) {
		return true
	}

	return false
}

// SetFlowStatus gets a reference to the given FlowStatus and assigns it to the FlowStatus field.
func (o *TrafficControlData) SetFlowStatus(v FlowStatus) {
	o.FlowStatus = &v
}

// GetRedirectInfo returns the RedirectInfo field value if set, zero value otherwise.
func (o *TrafficControlData) GetRedirectInfo() RedirectInformation {
	if o == nil || IsNil(o.RedirectInfo) {
		var ret RedirectInformation
		return ret
	}
	return *o.RedirectInfo
}

// GetRedirectInfoOk returns a tuple with the RedirectInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetRedirectInfoOk() (*RedirectInformation, bool) {
	if o == nil || IsNil(o.RedirectInfo) {
		return nil, false
	}
	return o.RedirectInfo, true
}

// HasRedirectInfo returns a boolean if a field has been set.
func (o *TrafficControlData) HasRedirectInfo() bool {
	if o != nil && !IsNil(o.RedirectInfo) {
		return true
	}

	return false
}

// SetRedirectInfo gets a reference to the given RedirectInformation and assigns it to the RedirectInfo field.
func (o *TrafficControlData) SetRedirectInfo(v RedirectInformation) {
	o.RedirectInfo = &v
}

// GetAddRedirectInfo returns the AddRedirectInfo field value if set, zero value otherwise.
func (o *TrafficControlData) GetAddRedirectInfo() []RedirectInformation {
	if o == nil || IsNil(o.AddRedirectInfo) {
		var ret []RedirectInformation
		return ret
	}
	return o.AddRedirectInfo
}

// GetAddRedirectInfoOk returns a tuple with the AddRedirectInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetAddRedirectInfoOk() ([]RedirectInformation, bool) {
	if o == nil || IsNil(o.AddRedirectInfo) {
		return nil, false
	}
	return o.AddRedirectInfo, true
}

// HasAddRedirectInfo returns a boolean if a field has been set.
func (o *TrafficControlData) HasAddRedirectInfo() bool {
	if o != nil && !IsNil(o.AddRedirectInfo) {
		return true
	}

	return false
}

// SetAddRedirectInfo gets a reference to the given []RedirectInformation and assigns it to the AddRedirectInfo field.
func (o *TrafficControlData) SetAddRedirectInfo(v []RedirectInformation) {
	o.AddRedirectInfo = v
}

// GetMuteNotif returns the MuteNotif field value if set, zero value otherwise.
func (o *TrafficControlData) GetMuteNotif() bool {
	if o == nil || IsNil(o.MuteNotif) {
		var ret bool
		return ret
	}
	return *o.MuteNotif
}

// GetMuteNotifOk returns a tuple with the MuteNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetMuteNotifOk() (*bool, bool) {
	if o == nil || IsNil(o.MuteNotif) {
		return nil, false
	}
	return o.MuteNotif, true
}

// HasMuteNotif returns a boolean if a field has been set.
func (o *TrafficControlData) HasMuteNotif() bool {
	if o != nil && !IsNil(o.MuteNotif) {
		return true
	}

	return false
}

// SetMuteNotif gets a reference to the given bool and assigns it to the MuteNotif field.
func (o *TrafficControlData) SetMuteNotif(v bool) {
	o.MuteNotif = &v
}

// GetTrafficSteeringPolIdDl returns the TrafficSteeringPolIdDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficControlData) GetTrafficSteeringPolIdDl() string {
	if o == nil || IsNil(o.TrafficSteeringPolIdDl.Get()) {
		var ret string
		return ret
	}
	return *o.TrafficSteeringPolIdDl.Get()
}

// GetTrafficSteeringPolIdDlOk returns a tuple with the TrafficSteeringPolIdDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficControlData) GetTrafficSteeringPolIdDlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrafficSteeringPolIdDl.Get(), o.TrafficSteeringPolIdDl.IsSet()
}

// HasTrafficSteeringPolIdDl returns a boolean if a field has been set.
func (o *TrafficControlData) HasTrafficSteeringPolIdDl() bool {
	if o != nil && o.TrafficSteeringPolIdDl.IsSet() {
		return true
	}

	return false
}

// SetTrafficSteeringPolIdDl gets a reference to the given NullableString and assigns it to the TrafficSteeringPolIdDl field.
func (o *TrafficControlData) SetTrafficSteeringPolIdDl(v string) {
	o.TrafficSteeringPolIdDl.Set(&v)
}
// SetTrafficSteeringPolIdDlNil sets the value for TrafficSteeringPolIdDl to be an explicit nil
func (o *TrafficControlData) SetTrafficSteeringPolIdDlNil() {
	o.TrafficSteeringPolIdDl.Set(nil)
}

// UnsetTrafficSteeringPolIdDl ensures that no value is present for TrafficSteeringPolIdDl, not even an explicit nil
func (o *TrafficControlData) UnsetTrafficSteeringPolIdDl() {
	o.TrafficSteeringPolIdDl.Unset()
}

// GetTrafficSteeringPolIdUl returns the TrafficSteeringPolIdUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficControlData) GetTrafficSteeringPolIdUl() string {
	if o == nil || IsNil(o.TrafficSteeringPolIdUl.Get()) {
		var ret string
		return ret
	}
	return *o.TrafficSteeringPolIdUl.Get()
}

// GetTrafficSteeringPolIdUlOk returns a tuple with the TrafficSteeringPolIdUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficControlData) GetTrafficSteeringPolIdUlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TrafficSteeringPolIdUl.Get(), o.TrafficSteeringPolIdUl.IsSet()
}

// HasTrafficSteeringPolIdUl returns a boolean if a field has been set.
func (o *TrafficControlData) HasTrafficSteeringPolIdUl() bool {
	if o != nil && o.TrafficSteeringPolIdUl.IsSet() {
		return true
	}

	return false
}

// SetTrafficSteeringPolIdUl gets a reference to the given NullableString and assigns it to the TrafficSteeringPolIdUl field.
func (o *TrafficControlData) SetTrafficSteeringPolIdUl(v string) {
	o.TrafficSteeringPolIdUl.Set(&v)
}
// SetTrafficSteeringPolIdUlNil sets the value for TrafficSteeringPolIdUl to be an explicit nil
func (o *TrafficControlData) SetTrafficSteeringPolIdUlNil() {
	o.TrafficSteeringPolIdUl.Set(nil)
}

// UnsetTrafficSteeringPolIdUl ensures that no value is present for TrafficSteeringPolIdUl, not even an explicit nil
func (o *TrafficControlData) UnsetTrafficSteeringPolIdUl() {
	o.TrafficSteeringPolIdUl.Unset()
}

// GetRouteToLocs returns the RouteToLocs field value if set, zero value otherwise.
func (o *TrafficControlData) GetRouteToLocs() []RouteToLocation {
	if o == nil || IsNil(o.RouteToLocs) {
		var ret []RouteToLocation
		return ret
	}
	return o.RouteToLocs
}

// GetRouteToLocsOk returns a tuple with the RouteToLocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetRouteToLocsOk() ([]RouteToLocation, bool) {
	if o == nil || IsNil(o.RouteToLocs) {
		return nil, false
	}
	return o.RouteToLocs, true
}

// HasRouteToLocs returns a boolean if a field has been set.
func (o *TrafficControlData) HasRouteToLocs() bool {
	if o != nil && !IsNil(o.RouteToLocs) {
		return true
	}

	return false
}

// SetRouteToLocs gets a reference to the given []RouteToLocation and assigns it to the RouteToLocs field.
func (o *TrafficControlData) SetRouteToLocs(v []RouteToLocation) {
	o.RouteToLocs = v
}

// GetTraffCorreInd returns the TraffCorreInd field value if set, zero value otherwise.
func (o *TrafficControlData) GetTraffCorreInd() bool {
	if o == nil || IsNil(o.TraffCorreInd) {
		var ret bool
		return ret
	}
	return *o.TraffCorreInd
}

// GetTraffCorreIndOk returns a tuple with the TraffCorreInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetTraffCorreIndOk() (*bool, bool) {
	if o == nil || IsNil(o.TraffCorreInd) {
		return nil, false
	}
	return o.TraffCorreInd, true
}

// HasTraffCorreInd returns a boolean if a field has been set.
func (o *TrafficControlData) HasTraffCorreInd() bool {
	if o != nil && !IsNil(o.TraffCorreInd) {
		return true
	}

	return false
}

// SetTraffCorreInd gets a reference to the given bool and assigns it to the TraffCorreInd field.
func (o *TrafficControlData) SetTraffCorreInd(v bool) {
	o.TraffCorreInd = &v
}

// GetUpPathChgEvent returns the UpPathChgEvent field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficControlData) GetUpPathChgEvent() UpPathChgEvent {
	if o == nil || IsNil(o.UpPathChgEvent.Get()) {
		var ret UpPathChgEvent
		return ret
	}
	return *o.UpPathChgEvent.Get()
}

// GetUpPathChgEventOk returns a tuple with the UpPathChgEvent field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficControlData) GetUpPathChgEventOk() (*UpPathChgEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.UpPathChgEvent.Get(), o.UpPathChgEvent.IsSet()
}

// HasUpPathChgEvent returns a boolean if a field has been set.
func (o *TrafficControlData) HasUpPathChgEvent() bool {
	if o != nil && o.UpPathChgEvent.IsSet() {
		return true
	}

	return false
}

// SetUpPathChgEvent gets a reference to the given NullableUpPathChgEvent and assigns it to the UpPathChgEvent field.
func (o *TrafficControlData) SetUpPathChgEvent(v UpPathChgEvent) {
	o.UpPathChgEvent.Set(&v)
}
// SetUpPathChgEventNil sets the value for UpPathChgEvent to be an explicit nil
func (o *TrafficControlData) SetUpPathChgEventNil() {
	o.UpPathChgEvent.Set(nil)
}

// UnsetUpPathChgEvent ensures that no value is present for UpPathChgEvent, not even an explicit nil
func (o *TrafficControlData) UnsetUpPathChgEvent() {
	o.UpPathChgEvent.Unset()
}

// GetSteerFun returns the SteerFun field value if set, zero value otherwise.
func (o *TrafficControlData) GetSteerFun() SteeringFunctionality {
	if o == nil || IsNil(o.SteerFun) {
		var ret SteeringFunctionality
		return ret
	}
	return *o.SteerFun
}

// GetSteerFunOk returns a tuple with the SteerFun field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetSteerFunOk() (*SteeringFunctionality, bool) {
	if o == nil || IsNil(o.SteerFun) {
		return nil, false
	}
	return o.SteerFun, true
}

// HasSteerFun returns a boolean if a field has been set.
func (o *TrafficControlData) HasSteerFun() bool {
	if o != nil && !IsNil(o.SteerFun) {
		return true
	}

	return false
}

// SetSteerFun gets a reference to the given SteeringFunctionality and assigns it to the SteerFun field.
func (o *TrafficControlData) SetSteerFun(v SteeringFunctionality) {
	o.SteerFun = &v
}

// GetSteerModeDl returns the SteerModeDl field value if set, zero value otherwise.
func (o *TrafficControlData) GetSteerModeDl() SteeringMode {
	if o == nil || IsNil(o.SteerModeDl) {
		var ret SteeringMode
		return ret
	}
	return *o.SteerModeDl
}

// GetSteerModeDlOk returns a tuple with the SteerModeDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetSteerModeDlOk() (*SteeringMode, bool) {
	if o == nil || IsNil(o.SteerModeDl) {
		return nil, false
	}
	return o.SteerModeDl, true
}

// HasSteerModeDl returns a boolean if a field has been set.
func (o *TrafficControlData) HasSteerModeDl() bool {
	if o != nil && !IsNil(o.SteerModeDl) {
		return true
	}

	return false
}

// SetSteerModeDl gets a reference to the given SteeringMode and assigns it to the SteerModeDl field.
func (o *TrafficControlData) SetSteerModeDl(v SteeringMode) {
	o.SteerModeDl = &v
}

// GetSteerModeUl returns the SteerModeUl field value if set, zero value otherwise.
func (o *TrafficControlData) GetSteerModeUl() SteeringMode {
	if o == nil || IsNil(o.SteerModeUl) {
		var ret SteeringMode
		return ret
	}
	return *o.SteerModeUl
}

// GetSteerModeUlOk returns a tuple with the SteerModeUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetSteerModeUlOk() (*SteeringMode, bool) {
	if o == nil || IsNil(o.SteerModeUl) {
		return nil, false
	}
	return o.SteerModeUl, true
}

// HasSteerModeUl returns a boolean if a field has been set.
func (o *TrafficControlData) HasSteerModeUl() bool {
	if o != nil && !IsNil(o.SteerModeUl) {
		return true
	}

	return false
}

// SetSteerModeUl gets a reference to the given SteeringMode and assigns it to the SteerModeUl field.
func (o *TrafficControlData) SetSteerModeUl(v SteeringMode) {
	o.SteerModeUl = &v
}

// GetMulAccCtrl returns the MulAccCtrl field value if set, zero value otherwise.
func (o *TrafficControlData) GetMulAccCtrl() MulticastAccessControl {
	if o == nil || IsNil(o.MulAccCtrl) {
		var ret MulticastAccessControl
		return ret
	}
	return *o.MulAccCtrl
}

// GetMulAccCtrlOk returns a tuple with the MulAccCtrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetMulAccCtrlOk() (*MulticastAccessControl, bool) {
	if o == nil || IsNil(o.MulAccCtrl) {
		return nil, false
	}
	return o.MulAccCtrl, true
}

// HasMulAccCtrl returns a boolean if a field has been set.
func (o *TrafficControlData) HasMulAccCtrl() bool {
	if o != nil && !IsNil(o.MulAccCtrl) {
		return true
	}

	return false
}

// SetMulAccCtrl gets a reference to the given MulticastAccessControl and assigns it to the MulAccCtrl field.
func (o *TrafficControlData) SetMulAccCtrl(v MulticastAccessControl) {
	o.MulAccCtrl = &v
}

// GetSnssaiList returns the SnssaiList field value if set, zero value otherwise.
func (o *TrafficControlData) GetSnssaiList() []Snssai {
	if o == nil || IsNil(o.SnssaiList) {
		var ret []Snssai
		return ret
	}
	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficControlData) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.SnssaiList) {
		return nil, false
	}
	return o.SnssaiList, true
}

// HasSnssaiList returns a boolean if a field has been set.
func (o *TrafficControlData) HasSnssaiList() bool {
	if o != nil && !IsNil(o.SnssaiList) {
		return true
	}

	return false
}

// SetSnssaiList gets a reference to the given []Snssai and assigns it to the SnssaiList field.
func (o *TrafficControlData) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

func (o TrafficControlData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficControlData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TcId) {
		toSerialize["tcId"] = o.TcId
	}
	if !IsNil(o.FlowStatus) {
		toSerialize["flowStatus"] = o.FlowStatus
	}
	if !IsNil(o.RedirectInfo) {
		toSerialize["redirectInfo"] = o.RedirectInfo
	}
	if !IsNil(o.AddRedirectInfo) {
		toSerialize["addRedirectInfo"] = o.AddRedirectInfo
	}
	if !IsNil(o.MuteNotif) {
		toSerialize["muteNotif"] = o.MuteNotif
	}
	if o.TrafficSteeringPolIdDl.IsSet() {
		toSerialize["trafficSteeringPolIdDl"] = o.TrafficSteeringPolIdDl.Get()
	}
	if o.TrafficSteeringPolIdUl.IsSet() {
		toSerialize["trafficSteeringPolIdUl"] = o.TrafficSteeringPolIdUl.Get()
	}
	if !IsNil(o.RouteToLocs) {
		toSerialize["routeToLocs"] = o.RouteToLocs
	}
	if !IsNil(o.TraffCorreInd) {
		toSerialize["traffCorreInd"] = o.TraffCorreInd
	}
	if o.UpPathChgEvent.IsSet() {
		toSerialize["upPathChgEvent"] = o.UpPathChgEvent.Get()
	}
	if !IsNil(o.SteerFun) {
		toSerialize["steerFun"] = o.SteerFun
	}
	if !IsNil(o.SteerModeDl) {
		toSerialize["steerModeDl"] = o.SteerModeDl
	}
	if !IsNil(o.SteerModeUl) {
		toSerialize["steerModeUl"] = o.SteerModeUl
	}
	if !IsNil(o.MulAccCtrl) {
		toSerialize["mulAccCtrl"] = o.MulAccCtrl
	}
	if !IsNil(o.SnssaiList) {
		toSerialize["snssaiList"] = o.SnssaiList
	}
	return toSerialize, nil
}

type NullableTrafficControlData struct {
	value *TrafficControlData
	isSet bool
}

func (v NullableTrafficControlData) Get() *TrafficControlData {
	return v.value
}

func (v *NullableTrafficControlData) Set(val *TrafficControlData) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficControlData) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficControlData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficControlData(val *TrafficControlData) *NullableTrafficControlData {
	return &NullableTrafficControlData{value: val, isSet: true}
}

func (v NullableTrafficControlData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficControlData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


