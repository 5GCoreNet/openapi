/*
SS_LocationAreaInfoRetrieval

API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_LocationAreaInfoRetrieval

import (
	"encoding/json"
)

// checks if the LocationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationInfo{}

// LocationInfo Represents the user location information.
type LocationInfo struct {
	// Unsigned integer identifying a period of time in units of minutes.
	AgeOfLocationInfo *int32 `json:"ageOfLocationInfo,omitempty"`
	// Indicates the Cell Global Identification of the user which identifies the cell the UE is registered.
	CellId *string `json:"cellId,omitempty"`
	// Indicates the eNodeB in which the UE is currently located.
	EnodeBId *string `json:"enodeBId,omitempty"`
	// Identifies the Routing Area Identity of the user where the UE is located.
	RoutingAreaId *string `json:"routingAreaId,omitempty"`
	// Identifies the Tracking Area Identity of the user where the UE is located.
	TrackingAreaId *string `json:"trackingAreaId,omitempty"`
	// Identifies the PLMN Identity of the user where the UE is located.
	PlmnId *string `json:"plmnId,omitempty"`
	// Identifies the TWAN Identity of the user where the UE is located.
	TwanId         *string                      `json:"twanId,omitempty"`
	GeographicArea *GeographicArea              `json:"geographicArea,omitempty"`
	CivicAddress   *CivicAddress                `json:"civicAddress,omitempty"`
	PositionMethod *PositioningMethod           `json:"positionMethod,omitempty"`
	QosFulfilInd   *AccuracyFulfilmentIndicator `json:"qosFulfilInd,omitempty"`
	UeVelocity     *VelocityEstimate            `json:"ueVelocity,omitempty"`
	LdrType        *LdrType                     `json:"ldrType,omitempty"`
	AchievedQos    *MinorLocationQoS            `json:"achievedQos,omitempty"`
}

// NewLocationInfo instantiates a new LocationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationInfo() *LocationInfo {
	this := LocationInfo{}
	return &this
}

// NewLocationInfoWithDefaults instantiates a new LocationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationInfoWithDefaults() *LocationInfo {
	this := LocationInfo{}
	return &this
}

// GetAgeOfLocationInfo returns the AgeOfLocationInfo field value if set, zero value otherwise.
func (o *LocationInfo) GetAgeOfLocationInfo() int32 {
	if o == nil || IsNil(o.AgeOfLocationInfo) {
		var ret int32
		return ret
	}
	return *o.AgeOfLocationInfo
}

// GetAgeOfLocationInfoOk returns a tuple with the AgeOfLocationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetAgeOfLocationInfoOk() (*int32, bool) {
	if o == nil || IsNil(o.AgeOfLocationInfo) {
		return nil, false
	}
	return o.AgeOfLocationInfo, true
}

// HasAgeOfLocationInfo returns a boolean if a field has been set.
func (o *LocationInfo) HasAgeOfLocationInfo() bool {
	if o != nil && !IsNil(o.AgeOfLocationInfo) {
		return true
	}

	return false
}

// SetAgeOfLocationInfo gets a reference to the given int32 and assigns it to the AgeOfLocationInfo field.
func (o *LocationInfo) SetAgeOfLocationInfo(v int32) {
	o.AgeOfLocationInfo = &v
}

// GetCellId returns the CellId field value if set, zero value otherwise.
func (o *LocationInfo) GetCellId() string {
	if o == nil || IsNil(o.CellId) {
		var ret string
		return ret
	}
	return *o.CellId
}

// GetCellIdOk returns a tuple with the CellId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetCellIdOk() (*string, bool) {
	if o == nil || IsNil(o.CellId) {
		return nil, false
	}
	return o.CellId, true
}

// HasCellId returns a boolean if a field has been set.
func (o *LocationInfo) HasCellId() bool {
	if o != nil && !IsNil(o.CellId) {
		return true
	}

	return false
}

// SetCellId gets a reference to the given string and assigns it to the CellId field.
func (o *LocationInfo) SetCellId(v string) {
	o.CellId = &v
}

// GetEnodeBId returns the EnodeBId field value if set, zero value otherwise.
func (o *LocationInfo) GetEnodeBId() string {
	if o == nil || IsNil(o.EnodeBId) {
		var ret string
		return ret
	}
	return *o.EnodeBId
}

// GetEnodeBIdOk returns a tuple with the EnodeBId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetEnodeBIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnodeBId) {
		return nil, false
	}
	return o.EnodeBId, true
}

// HasEnodeBId returns a boolean if a field has been set.
func (o *LocationInfo) HasEnodeBId() bool {
	if o != nil && !IsNil(o.EnodeBId) {
		return true
	}

	return false
}

// SetEnodeBId gets a reference to the given string and assigns it to the EnodeBId field.
func (o *LocationInfo) SetEnodeBId(v string) {
	o.EnodeBId = &v
}

// GetRoutingAreaId returns the RoutingAreaId field value if set, zero value otherwise.
func (o *LocationInfo) GetRoutingAreaId() string {
	if o == nil || IsNil(o.RoutingAreaId) {
		var ret string
		return ret
	}
	return *o.RoutingAreaId
}

// GetRoutingAreaIdOk returns a tuple with the RoutingAreaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetRoutingAreaIdOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingAreaId) {
		return nil, false
	}
	return o.RoutingAreaId, true
}

// HasRoutingAreaId returns a boolean if a field has been set.
func (o *LocationInfo) HasRoutingAreaId() bool {
	if o != nil && !IsNil(o.RoutingAreaId) {
		return true
	}

	return false
}

// SetRoutingAreaId gets a reference to the given string and assigns it to the RoutingAreaId field.
func (o *LocationInfo) SetRoutingAreaId(v string) {
	o.RoutingAreaId = &v
}

// GetTrackingAreaId returns the TrackingAreaId field value if set, zero value otherwise.
func (o *LocationInfo) GetTrackingAreaId() string {
	if o == nil || IsNil(o.TrackingAreaId) {
		var ret string
		return ret
	}
	return *o.TrackingAreaId
}

// GetTrackingAreaIdOk returns a tuple with the TrackingAreaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetTrackingAreaIdOk() (*string, bool) {
	if o == nil || IsNil(o.TrackingAreaId) {
		return nil, false
	}
	return o.TrackingAreaId, true
}

// HasTrackingAreaId returns a boolean if a field has been set.
func (o *LocationInfo) HasTrackingAreaId() bool {
	if o != nil && !IsNil(o.TrackingAreaId) {
		return true
	}

	return false
}

// SetTrackingAreaId gets a reference to the given string and assigns it to the TrackingAreaId field.
func (o *LocationInfo) SetTrackingAreaId(v string) {
	o.TrackingAreaId = &v
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *LocationInfo) GetPlmnId() string {
	if o == nil || IsNil(o.PlmnId) {
		var ret string
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetPlmnIdOk() (*string, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *LocationInfo) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given string and assigns it to the PlmnId field.
func (o *LocationInfo) SetPlmnId(v string) {
	o.PlmnId = &v
}

// GetTwanId returns the TwanId field value if set, zero value otherwise.
func (o *LocationInfo) GetTwanId() string {
	if o == nil || IsNil(o.TwanId) {
		var ret string
		return ret
	}
	return *o.TwanId
}

// GetTwanIdOk returns a tuple with the TwanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetTwanIdOk() (*string, bool) {
	if o == nil || IsNil(o.TwanId) {
		return nil, false
	}
	return o.TwanId, true
}

// HasTwanId returns a boolean if a field has been set.
func (o *LocationInfo) HasTwanId() bool {
	if o != nil && !IsNil(o.TwanId) {
		return true
	}

	return false
}

// SetTwanId gets a reference to the given string and assigns it to the TwanId field.
func (o *LocationInfo) SetTwanId(v string) {
	o.TwanId = &v
}

// GetGeographicArea returns the GeographicArea field value if set, zero value otherwise.
func (o *LocationInfo) GetGeographicArea() GeographicArea {
	if o == nil || IsNil(o.GeographicArea) {
		var ret GeographicArea
		return ret
	}
	return *o.GeographicArea
}

// GetGeographicAreaOk returns a tuple with the GeographicArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetGeographicAreaOk() (*GeographicArea, bool) {
	if o == nil || IsNil(o.GeographicArea) {
		return nil, false
	}
	return o.GeographicArea, true
}

// HasGeographicArea returns a boolean if a field has been set.
func (o *LocationInfo) HasGeographicArea() bool {
	if o != nil && !IsNil(o.GeographicArea) {
		return true
	}

	return false
}

// SetGeographicArea gets a reference to the given GeographicArea and assigns it to the GeographicArea field.
func (o *LocationInfo) SetGeographicArea(v GeographicArea) {
	o.GeographicArea = &v
}

// GetCivicAddress returns the CivicAddress field value if set, zero value otherwise.
func (o *LocationInfo) GetCivicAddress() CivicAddress {
	if o == nil || IsNil(o.CivicAddress) {
		var ret CivicAddress
		return ret
	}
	return *o.CivicAddress
}

// GetCivicAddressOk returns a tuple with the CivicAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetCivicAddressOk() (*CivicAddress, bool) {
	if o == nil || IsNil(o.CivicAddress) {
		return nil, false
	}
	return o.CivicAddress, true
}

// HasCivicAddress returns a boolean if a field has been set.
func (o *LocationInfo) HasCivicAddress() bool {
	if o != nil && !IsNil(o.CivicAddress) {
		return true
	}

	return false
}

// SetCivicAddress gets a reference to the given CivicAddress and assigns it to the CivicAddress field.
func (o *LocationInfo) SetCivicAddress(v CivicAddress) {
	o.CivicAddress = &v
}

// GetPositionMethod returns the PositionMethod field value if set, zero value otherwise.
func (o *LocationInfo) GetPositionMethod() PositioningMethod {
	if o == nil || IsNil(o.PositionMethod) {
		var ret PositioningMethod
		return ret
	}
	return *o.PositionMethod
}

// GetPositionMethodOk returns a tuple with the PositionMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetPositionMethodOk() (*PositioningMethod, bool) {
	if o == nil || IsNil(o.PositionMethod) {
		return nil, false
	}
	return o.PositionMethod, true
}

// HasPositionMethod returns a boolean if a field has been set.
func (o *LocationInfo) HasPositionMethod() bool {
	if o != nil && !IsNil(o.PositionMethod) {
		return true
	}

	return false
}

// SetPositionMethod gets a reference to the given PositioningMethod and assigns it to the PositionMethod field.
func (o *LocationInfo) SetPositionMethod(v PositioningMethod) {
	o.PositionMethod = &v
}

// GetQosFulfilInd returns the QosFulfilInd field value if set, zero value otherwise.
func (o *LocationInfo) GetQosFulfilInd() AccuracyFulfilmentIndicator {
	if o == nil || IsNil(o.QosFulfilInd) {
		var ret AccuracyFulfilmentIndicator
		return ret
	}
	return *o.QosFulfilInd
}

// GetQosFulfilIndOk returns a tuple with the QosFulfilInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetQosFulfilIndOk() (*AccuracyFulfilmentIndicator, bool) {
	if o == nil || IsNil(o.QosFulfilInd) {
		return nil, false
	}
	return o.QosFulfilInd, true
}

// HasQosFulfilInd returns a boolean if a field has been set.
func (o *LocationInfo) HasQosFulfilInd() bool {
	if o != nil && !IsNil(o.QosFulfilInd) {
		return true
	}

	return false
}

// SetQosFulfilInd gets a reference to the given AccuracyFulfilmentIndicator and assigns it to the QosFulfilInd field.
func (o *LocationInfo) SetQosFulfilInd(v AccuracyFulfilmentIndicator) {
	o.QosFulfilInd = &v
}

// GetUeVelocity returns the UeVelocity field value if set, zero value otherwise.
func (o *LocationInfo) GetUeVelocity() VelocityEstimate {
	if o == nil || IsNil(o.UeVelocity) {
		var ret VelocityEstimate
		return ret
	}
	return *o.UeVelocity
}

// GetUeVelocityOk returns a tuple with the UeVelocity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetUeVelocityOk() (*VelocityEstimate, bool) {
	if o == nil || IsNil(o.UeVelocity) {
		return nil, false
	}
	return o.UeVelocity, true
}

// HasUeVelocity returns a boolean if a field has been set.
func (o *LocationInfo) HasUeVelocity() bool {
	if o != nil && !IsNil(o.UeVelocity) {
		return true
	}

	return false
}

// SetUeVelocity gets a reference to the given VelocityEstimate and assigns it to the UeVelocity field.
func (o *LocationInfo) SetUeVelocity(v VelocityEstimate) {
	o.UeVelocity = &v
}

// GetLdrType returns the LdrType field value if set, zero value otherwise.
func (o *LocationInfo) GetLdrType() LdrType {
	if o == nil || IsNil(o.LdrType) {
		var ret LdrType
		return ret
	}
	return *o.LdrType
}

// GetLdrTypeOk returns a tuple with the LdrType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetLdrTypeOk() (*LdrType, bool) {
	if o == nil || IsNil(o.LdrType) {
		return nil, false
	}
	return o.LdrType, true
}

// HasLdrType returns a boolean if a field has been set.
func (o *LocationInfo) HasLdrType() bool {
	if o != nil && !IsNil(o.LdrType) {
		return true
	}

	return false
}

// SetLdrType gets a reference to the given LdrType and assigns it to the LdrType field.
func (o *LocationInfo) SetLdrType(v LdrType) {
	o.LdrType = &v
}

// GetAchievedQos returns the AchievedQos field value if set, zero value otherwise.
func (o *LocationInfo) GetAchievedQos() MinorLocationQoS {
	if o == nil || IsNil(o.AchievedQos) {
		var ret MinorLocationQoS
		return ret
	}
	return *o.AchievedQos
}

// GetAchievedQosOk returns a tuple with the AchievedQos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationInfo) GetAchievedQosOk() (*MinorLocationQoS, bool) {
	if o == nil || IsNil(o.AchievedQos) {
		return nil, false
	}
	return o.AchievedQos, true
}

// HasAchievedQos returns a boolean if a field has been set.
func (o *LocationInfo) HasAchievedQos() bool {
	if o != nil && !IsNil(o.AchievedQos) {
		return true
	}

	return false
}

// SetAchievedQos gets a reference to the given MinorLocationQoS and assigns it to the AchievedQos field.
func (o *LocationInfo) SetAchievedQos(v MinorLocationQoS) {
	o.AchievedQos = &v
}

func (o LocationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgeOfLocationInfo) {
		toSerialize["ageOfLocationInfo"] = o.AgeOfLocationInfo
	}
	if !IsNil(o.CellId) {
		toSerialize["cellId"] = o.CellId
	}
	if !IsNil(o.EnodeBId) {
		toSerialize["enodeBId"] = o.EnodeBId
	}
	if !IsNil(o.RoutingAreaId) {
		toSerialize["routingAreaId"] = o.RoutingAreaId
	}
	if !IsNil(o.TrackingAreaId) {
		toSerialize["trackingAreaId"] = o.TrackingAreaId
	}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.TwanId) {
		toSerialize["twanId"] = o.TwanId
	}
	if !IsNil(o.GeographicArea) {
		toSerialize["geographicArea"] = o.GeographicArea
	}
	if !IsNil(o.CivicAddress) {
		toSerialize["civicAddress"] = o.CivicAddress
	}
	if !IsNil(o.PositionMethod) {
		toSerialize["positionMethod"] = o.PositionMethod
	}
	if !IsNil(o.QosFulfilInd) {
		toSerialize["qosFulfilInd"] = o.QosFulfilInd
	}
	if !IsNil(o.UeVelocity) {
		toSerialize["ueVelocity"] = o.UeVelocity
	}
	if !IsNil(o.LdrType) {
		toSerialize["ldrType"] = o.LdrType
	}
	if !IsNil(o.AchievedQos) {
		toSerialize["achievedQos"] = o.AchievedQos
	}
	return toSerialize, nil
}

type NullableLocationInfo struct {
	value *LocationInfo
	isSet bool
}

func (v NullableLocationInfo) Get() *LocationInfo {
	return v.value
}

func (v *NullableLocationInfo) Set(val *LocationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationInfo(val *LocationInfo) *NullableLocationInfo {
	return &NullableLocationInfo{value: val, isSet: true}
}

func (v NullableLocationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
