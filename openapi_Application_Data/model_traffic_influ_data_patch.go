/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
	"time"
)

// checks if the TrafficInfluDataPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrafficInfluDataPatch{}

// TrafficInfluDataPatch Represents the Traffic Influence Data to be updated in the UDR.
type TrafficInfluDataPatch struct {
	// Contains the Notification Correlation Id allocated by the NEF for the UP path change notification. 
	UpPathChgNotifCorreId *string `json:"upPathChgNotifCorreId,omitempty"`
	// Identifies whether an application can be relocated once a location of the application has been selected. 
	AppReloInd *bool `json:"appReloInd,omitempty"`
	// Identifies Ethernet packet filters. Either \"trafficFilters\" or \"ethTrafficFilters\" shall be included if applicable. 
	EthTrafficFilters []EthFlowDescription `json:"ethTrafficFilters,omitempty"`
	// Identifies IP packet filters. Either \"trafficFilters\" or \"ethTrafficFilters\" shall be included if applicable. 
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`
	// Identifies the N6 traffic routing requirement.
	TrafficRoutes []RouteToLocation `json:"trafficRoutes,omitempty"`
	TraffCorreInd *bool `json:"traffCorreInd,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidStartTime *time.Time `json:"validStartTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidEndTime *time.Time `json:"validEndTime,omitempty"`
	// Identifies the temporal validities for the N6 traffic routing requirement.
	TempValidities []TemporalValidity `json:"tempValidities,omitempty"`
	NwAreaInfo *NetworkAreaInfo `json:"nwAreaInfo,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	UpPathChgNotifUri *string `json:"upPathChgNotifUri,omitempty"`
	Headers []string `json:"headers,omitempty"`
	AfAckInd *bool `json:"afAckInd,omitempty"`
	AddrPreserInd *bool `json:"addrPreserInd,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible with the OpenAPI 'nullable: true' property. 
	MaxAllowedUpLat NullableInt32 `json:"maxAllowedUpLat,omitempty"`
	// Indicates whether simultaneous connectivity should be temporarily maintained for the source and target PSA. 
	SimConnInd *bool `json:"simConnInd,omitempty"`
	// indicating a time in seconds with OpenAPI defined 'nullable: true' property.
	SimConnTerm NullableInt32 `json:"simConnTerm,omitempty"`
}

// NewTrafficInfluDataPatch instantiates a new TrafficInfluDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrafficInfluDataPatch() *TrafficInfluDataPatch {
	this := TrafficInfluDataPatch{}
	return &this
}

// NewTrafficInfluDataPatchWithDefaults instantiates a new TrafficInfluDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrafficInfluDataPatchWithDefaults() *TrafficInfluDataPatch {
	this := TrafficInfluDataPatch{}
	return &this
}

// GetUpPathChgNotifCorreId returns the UpPathChgNotifCorreId field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetUpPathChgNotifCorreId() string {
	if o == nil || IsNil(o.UpPathChgNotifCorreId) {
		var ret string
		return ret
	}
	return *o.UpPathChgNotifCorreId
}

// GetUpPathChgNotifCorreIdOk returns a tuple with the UpPathChgNotifCorreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetUpPathChgNotifCorreIdOk() (*string, bool) {
	if o == nil || IsNil(o.UpPathChgNotifCorreId) {
		return nil, false
	}
	return o.UpPathChgNotifCorreId, true
}

// HasUpPathChgNotifCorreId returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasUpPathChgNotifCorreId() bool {
	if o != nil && !IsNil(o.UpPathChgNotifCorreId) {
		return true
	}

	return false
}

// SetUpPathChgNotifCorreId gets a reference to the given string and assigns it to the UpPathChgNotifCorreId field.
func (o *TrafficInfluDataPatch) SetUpPathChgNotifCorreId(v string) {
	o.UpPathChgNotifCorreId = &v
}

// GetAppReloInd returns the AppReloInd field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetAppReloInd() bool {
	if o == nil || IsNil(o.AppReloInd) {
		var ret bool
		return ret
	}
	return *o.AppReloInd
}

// GetAppReloIndOk returns a tuple with the AppReloInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetAppReloIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AppReloInd) {
		return nil, false
	}
	return o.AppReloInd, true
}

// HasAppReloInd returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasAppReloInd() bool {
	if o != nil && !IsNil(o.AppReloInd) {
		return true
	}

	return false
}

// SetAppReloInd gets a reference to the given bool and assigns it to the AppReloInd field.
func (o *TrafficInfluDataPatch) SetAppReloInd(v bool) {
	o.AppReloInd = &v
}

// GetEthTrafficFilters returns the EthTrafficFilters field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetEthTrafficFilters() []EthFlowDescription {
	if o == nil || IsNil(o.EthTrafficFilters) {
		var ret []EthFlowDescription
		return ret
	}
	return o.EthTrafficFilters
}

// GetEthTrafficFiltersOk returns a tuple with the EthTrafficFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetEthTrafficFiltersOk() ([]EthFlowDescription, bool) {
	if o == nil || IsNil(o.EthTrafficFilters) {
		return nil, false
	}
	return o.EthTrafficFilters, true
}

// HasEthTrafficFilters returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasEthTrafficFilters() bool {
	if o != nil && !IsNil(o.EthTrafficFilters) {
		return true
	}

	return false
}

// SetEthTrafficFilters gets a reference to the given []EthFlowDescription and assigns it to the EthTrafficFilters field.
func (o *TrafficInfluDataPatch) SetEthTrafficFilters(v []EthFlowDescription) {
	o.EthTrafficFilters = v
}

// GetTrafficFilters returns the TrafficFilters field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetTrafficFilters() []FlowInfo {
	if o == nil || IsNil(o.TrafficFilters) {
		var ret []FlowInfo
		return ret
	}
	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetTrafficFiltersOk() ([]FlowInfo, bool) {
	if o == nil || IsNil(o.TrafficFilters) {
		return nil, false
	}
	return o.TrafficFilters, true
}

// HasTrafficFilters returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasTrafficFilters() bool {
	if o != nil && !IsNil(o.TrafficFilters) {
		return true
	}

	return false
}

// SetTrafficFilters gets a reference to the given []FlowInfo and assigns it to the TrafficFilters field.
func (o *TrafficInfluDataPatch) SetTrafficFilters(v []FlowInfo) {
	o.TrafficFilters = v
}

// GetTrafficRoutes returns the TrafficRoutes field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetTrafficRoutes() []RouteToLocation {
	if o == nil || IsNil(o.TrafficRoutes) {
		var ret []RouteToLocation
		return ret
	}
	return o.TrafficRoutes
}

// GetTrafficRoutesOk returns a tuple with the TrafficRoutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetTrafficRoutesOk() ([]RouteToLocation, bool) {
	if o == nil || IsNil(o.TrafficRoutes) {
		return nil, false
	}
	return o.TrafficRoutes, true
}

// HasTrafficRoutes returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasTrafficRoutes() bool {
	if o != nil && !IsNil(o.TrafficRoutes) {
		return true
	}

	return false
}

// SetTrafficRoutes gets a reference to the given []RouteToLocation and assigns it to the TrafficRoutes field.
func (o *TrafficInfluDataPatch) SetTrafficRoutes(v []RouteToLocation) {
	o.TrafficRoutes = v
}

// GetTraffCorreInd returns the TraffCorreInd field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetTraffCorreInd() bool {
	if o == nil || IsNil(o.TraffCorreInd) {
		var ret bool
		return ret
	}
	return *o.TraffCorreInd
}

// GetTraffCorreIndOk returns a tuple with the TraffCorreInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetTraffCorreIndOk() (*bool, bool) {
	if o == nil || IsNil(o.TraffCorreInd) {
		return nil, false
	}
	return o.TraffCorreInd, true
}

// HasTraffCorreInd returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasTraffCorreInd() bool {
	if o != nil && !IsNil(o.TraffCorreInd) {
		return true
	}

	return false
}

// SetTraffCorreInd gets a reference to the given bool and assigns it to the TraffCorreInd field.
func (o *TrafficInfluDataPatch) SetTraffCorreInd(v bool) {
	o.TraffCorreInd = &v
}

// GetValidStartTime returns the ValidStartTime field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetValidStartTime() time.Time {
	if o == nil || IsNil(o.ValidStartTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidStartTime
}

// GetValidStartTimeOk returns a tuple with the ValidStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetValidStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidStartTime) {
		return nil, false
	}
	return o.ValidStartTime, true
}

// HasValidStartTime returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasValidStartTime() bool {
	if o != nil && !IsNil(o.ValidStartTime) {
		return true
	}

	return false
}

// SetValidStartTime gets a reference to the given time.Time and assigns it to the ValidStartTime field.
func (o *TrafficInfluDataPatch) SetValidStartTime(v time.Time) {
	o.ValidStartTime = &v
}

// GetValidEndTime returns the ValidEndTime field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetValidEndTime() time.Time {
	if o == nil || IsNil(o.ValidEndTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidEndTime
}

// GetValidEndTimeOk returns a tuple with the ValidEndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetValidEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidEndTime) {
		return nil, false
	}
	return o.ValidEndTime, true
}

// HasValidEndTime returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasValidEndTime() bool {
	if o != nil && !IsNil(o.ValidEndTime) {
		return true
	}

	return false
}

// SetValidEndTime gets a reference to the given time.Time and assigns it to the ValidEndTime field.
func (o *TrafficInfluDataPatch) SetValidEndTime(v time.Time) {
	o.ValidEndTime = &v
}

// GetTempValidities returns the TempValidities field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficInfluDataPatch) GetTempValidities() []TemporalValidity {
	if o == nil {
		var ret []TemporalValidity
		return ret
	}
	return o.TempValidities
}

// GetTempValiditiesOk returns a tuple with the TempValidities field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficInfluDataPatch) GetTempValiditiesOk() ([]TemporalValidity, bool) {
	if o == nil || IsNil(o.TempValidities) {
		return nil, false
	}
	return o.TempValidities, true
}

// HasTempValidities returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasTempValidities() bool {
	if o != nil && IsNil(o.TempValidities) {
		return true
	}

	return false
}

// SetTempValidities gets a reference to the given []TemporalValidity and assigns it to the TempValidities field.
func (o *TrafficInfluDataPatch) SetTempValidities(v []TemporalValidity) {
	o.TempValidities = v
}

// GetNwAreaInfo returns the NwAreaInfo field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetNwAreaInfo() NetworkAreaInfo {
	if o == nil || IsNil(o.NwAreaInfo) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NwAreaInfo
}

// GetNwAreaInfoOk returns a tuple with the NwAreaInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetNwAreaInfoOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.NwAreaInfo) {
		return nil, false
	}
	return o.NwAreaInfo, true
}

// HasNwAreaInfo returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasNwAreaInfo() bool {
	if o != nil && !IsNil(o.NwAreaInfo) {
		return true
	}

	return false
}

// SetNwAreaInfo gets a reference to the given NetworkAreaInfo and assigns it to the NwAreaInfo field.
func (o *TrafficInfluDataPatch) SetNwAreaInfo(v NetworkAreaInfo) {
	o.NwAreaInfo = &v
}

// GetUpPathChgNotifUri returns the UpPathChgNotifUri field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetUpPathChgNotifUri() string {
	if o == nil || IsNil(o.UpPathChgNotifUri) {
		var ret string
		return ret
	}
	return *o.UpPathChgNotifUri
}

// GetUpPathChgNotifUriOk returns a tuple with the UpPathChgNotifUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetUpPathChgNotifUriOk() (*string, bool) {
	if o == nil || IsNil(o.UpPathChgNotifUri) {
		return nil, false
	}
	return o.UpPathChgNotifUri, true
}

// HasUpPathChgNotifUri returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasUpPathChgNotifUri() bool {
	if o != nil && !IsNil(o.UpPathChgNotifUri) {
		return true
	}

	return false
}

// SetUpPathChgNotifUri gets a reference to the given string and assigns it to the UpPathChgNotifUri field.
func (o *TrafficInfluDataPatch) SetUpPathChgNotifUri(v string) {
	o.UpPathChgNotifUri = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetHeaders() []string {
	if o == nil || IsNil(o.Headers) {
		var ret []string
		return ret
	}
	return o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetHeadersOk() ([]string, bool) {
	if o == nil || IsNil(o.Headers) {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasHeaders() bool {
	if o != nil && !IsNil(o.Headers) {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given []string and assigns it to the Headers field.
func (o *TrafficInfluDataPatch) SetHeaders(v []string) {
	o.Headers = v
}

// GetAfAckInd returns the AfAckInd field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetAfAckInd() bool {
	if o == nil || IsNil(o.AfAckInd) {
		var ret bool
		return ret
	}
	return *o.AfAckInd
}

// GetAfAckIndOk returns a tuple with the AfAckInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetAfAckIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AfAckInd) {
		return nil, false
	}
	return o.AfAckInd, true
}

// HasAfAckInd returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasAfAckInd() bool {
	if o != nil && !IsNil(o.AfAckInd) {
		return true
	}

	return false
}

// SetAfAckInd gets a reference to the given bool and assigns it to the AfAckInd field.
func (o *TrafficInfluDataPatch) SetAfAckInd(v bool) {
	o.AfAckInd = &v
}

// GetAddrPreserInd returns the AddrPreserInd field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetAddrPreserInd() bool {
	if o == nil || IsNil(o.AddrPreserInd) {
		var ret bool
		return ret
	}
	return *o.AddrPreserInd
}

// GetAddrPreserIndOk returns a tuple with the AddrPreserInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetAddrPreserIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AddrPreserInd) {
		return nil, false
	}
	return o.AddrPreserInd, true
}

// HasAddrPreserInd returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasAddrPreserInd() bool {
	if o != nil && !IsNil(o.AddrPreserInd) {
		return true
	}

	return false
}

// SetAddrPreserInd gets a reference to the given bool and assigns it to the AddrPreserInd field.
func (o *TrafficInfluDataPatch) SetAddrPreserInd(v bool) {
	o.AddrPreserInd = &v
}

// GetMaxAllowedUpLat returns the MaxAllowedUpLat field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficInfluDataPatch) GetMaxAllowedUpLat() int32 {
	if o == nil || IsNil(o.MaxAllowedUpLat.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxAllowedUpLat.Get()
}

// GetMaxAllowedUpLatOk returns a tuple with the MaxAllowedUpLat field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficInfluDataPatch) GetMaxAllowedUpLatOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxAllowedUpLat.Get(), o.MaxAllowedUpLat.IsSet()
}

// HasMaxAllowedUpLat returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasMaxAllowedUpLat() bool {
	if o != nil && o.MaxAllowedUpLat.IsSet() {
		return true
	}

	return false
}

// SetMaxAllowedUpLat gets a reference to the given NullableInt32 and assigns it to the MaxAllowedUpLat field.
func (o *TrafficInfluDataPatch) SetMaxAllowedUpLat(v int32) {
	o.MaxAllowedUpLat.Set(&v)
}
// SetMaxAllowedUpLatNil sets the value for MaxAllowedUpLat to be an explicit nil
func (o *TrafficInfluDataPatch) SetMaxAllowedUpLatNil() {
	o.MaxAllowedUpLat.Set(nil)
}

// UnsetMaxAllowedUpLat ensures that no value is present for MaxAllowedUpLat, not even an explicit nil
func (o *TrafficInfluDataPatch) UnsetMaxAllowedUpLat() {
	o.MaxAllowedUpLat.Unset()
}

// GetSimConnInd returns the SimConnInd field value if set, zero value otherwise.
func (o *TrafficInfluDataPatch) GetSimConnInd() bool {
	if o == nil || IsNil(o.SimConnInd) {
		var ret bool
		return ret
	}
	return *o.SimConnInd
}

// GetSimConnIndOk returns a tuple with the SimConnInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrafficInfluDataPatch) GetSimConnIndOk() (*bool, bool) {
	if o == nil || IsNil(o.SimConnInd) {
		return nil, false
	}
	return o.SimConnInd, true
}

// HasSimConnInd returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasSimConnInd() bool {
	if o != nil && !IsNil(o.SimConnInd) {
		return true
	}

	return false
}

// SetSimConnInd gets a reference to the given bool and assigns it to the SimConnInd field.
func (o *TrafficInfluDataPatch) SetSimConnInd(v bool) {
	o.SimConnInd = &v
}

// GetSimConnTerm returns the SimConnTerm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TrafficInfluDataPatch) GetSimConnTerm() int32 {
	if o == nil || IsNil(o.SimConnTerm.Get()) {
		var ret int32
		return ret
	}
	return *o.SimConnTerm.Get()
}

// GetSimConnTermOk returns a tuple with the SimConnTerm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TrafficInfluDataPatch) GetSimConnTermOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SimConnTerm.Get(), o.SimConnTerm.IsSet()
}

// HasSimConnTerm returns a boolean if a field has been set.
func (o *TrafficInfluDataPatch) HasSimConnTerm() bool {
	if o != nil && o.SimConnTerm.IsSet() {
		return true
	}

	return false
}

// SetSimConnTerm gets a reference to the given NullableInt32 and assigns it to the SimConnTerm field.
func (o *TrafficInfluDataPatch) SetSimConnTerm(v int32) {
	o.SimConnTerm.Set(&v)
}
// SetSimConnTermNil sets the value for SimConnTerm to be an explicit nil
func (o *TrafficInfluDataPatch) SetSimConnTermNil() {
	o.SimConnTerm.Set(nil)
}

// UnsetSimConnTerm ensures that no value is present for SimConnTerm, not even an explicit nil
func (o *TrafficInfluDataPatch) UnsetSimConnTerm() {
	o.SimConnTerm.Unset()
}

func (o TrafficInfluDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrafficInfluDataPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UpPathChgNotifCorreId) {
		toSerialize["upPathChgNotifCorreId"] = o.UpPathChgNotifCorreId
	}
	if !IsNil(o.AppReloInd) {
		toSerialize["appReloInd"] = o.AppReloInd
	}
	if !IsNil(o.EthTrafficFilters) {
		toSerialize["ethTrafficFilters"] = o.EthTrafficFilters
	}
	if !IsNil(o.TrafficFilters) {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	if !IsNil(o.TrafficRoutes) {
		toSerialize["trafficRoutes"] = o.TrafficRoutes
	}
	if !IsNil(o.TraffCorreInd) {
		toSerialize["traffCorreInd"] = o.TraffCorreInd
	}
	if !IsNil(o.ValidStartTime) {
		toSerialize["validStartTime"] = o.ValidStartTime
	}
	if !IsNil(o.ValidEndTime) {
		toSerialize["validEndTime"] = o.ValidEndTime
	}
	if o.TempValidities != nil {
		toSerialize["tempValidities"] = o.TempValidities
	}
	if !IsNil(o.NwAreaInfo) {
		toSerialize["nwAreaInfo"] = o.NwAreaInfo
	}
	if !IsNil(o.UpPathChgNotifUri) {
		toSerialize["upPathChgNotifUri"] = o.UpPathChgNotifUri
	}
	if !IsNil(o.Headers) {
		toSerialize["headers"] = o.Headers
	}
	if !IsNil(o.AfAckInd) {
		toSerialize["afAckInd"] = o.AfAckInd
	}
	if !IsNil(o.AddrPreserInd) {
		toSerialize["addrPreserInd"] = o.AddrPreserInd
	}
	if o.MaxAllowedUpLat.IsSet() {
		toSerialize["maxAllowedUpLat"] = o.MaxAllowedUpLat.Get()
	}
	if !IsNil(o.SimConnInd) {
		toSerialize["simConnInd"] = o.SimConnInd
	}
	if o.SimConnTerm.IsSet() {
		toSerialize["simConnTerm"] = o.SimConnTerm.Get()
	}
	return toSerialize, nil
}

type NullableTrafficInfluDataPatch struct {
	value *TrafficInfluDataPatch
	isSet bool
}

func (v NullableTrafficInfluDataPatch) Get() *TrafficInfluDataPatch {
	return v.value
}

func (v *NullableTrafficInfluDataPatch) Set(val *TrafficInfluDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficInfluDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficInfluDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficInfluDataPatch(val *TrafficInfluDataPatch) *NullableTrafficInfluDataPatch {
	return &NullableTrafficInfluDataPatch{value: val, isSet: true}
}

func (v NullableTrafficInfluDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficInfluDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


