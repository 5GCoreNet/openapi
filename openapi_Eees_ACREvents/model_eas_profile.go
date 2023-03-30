/*
Eees_ACREvents

API for ACR events subscription and notification. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACREvents

import (
	"encoding/json"
)

// checks if the EASProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EASProfile{}

// EASProfile Represents the EAS profile information.
type EASProfile struct {
	// Identifier of the EAS.
	EasId string `json:"easId"`
	EndPt EndPoint `json:"endPt"`
	// Identities of application clients that are served by the EAS.
	AcIds []string `json:"acIds,omitempty"`
	// Identifier of the ASP that provides the EAS.
	ProvId *string `json:"provId,omitempty"`
	Type *EASCategory `json:"type,omitempty"`
	// The availability schedule of the EAS.
	Scheds []ScheduledCommunicationTime `json:"scheds,omitempty"`
	SvcArea *ServiceArea `json:"svcArea,omitempty"`
	SvcKpi *EASServiceKPI `json:"svcKpi,omitempty"`
	// level of service permissions supported by the EAS.
	PermLvl []PermissionLevel `json:"permLvl,omitempty"`
	// Service specific features supported by EAS.
	EasFeats []string `json:"easFeats,omitempty"`
	// List of DNAI(s) and the N6 traffic information associated with the EAS.
	AppLocs []RouteToLocation `json:"appLocs,omitempty"`
	// The ACR scenarios supported by the EAS for service continuity.
	SvcContSupp []ACRScenario `json:"svcContSupp,omitempty"`
	// Unsigned integer identifying a period of time in units of seconds.
	AvlRep *int32 `json:"avlRep,omitempty"`
	// EAS status information.
	Status *string `json:"status,omitempty"`
}

// NewEASProfile instantiates a new EASProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEASProfile(easId string, endPt EndPoint) *EASProfile {
	this := EASProfile{}
	this.EasId = easId
	this.EndPt = endPt
	return &this
}

// NewEASProfileWithDefaults instantiates a new EASProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEASProfileWithDefaults() *EASProfile {
	this := EASProfile{}
	return &this
}

// GetEasId returns the EasId field value
func (o *EASProfile) GetEasId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EasId
}

// GetEasIdOk returns a tuple with the EasId field value
// and a boolean to check if the value has been set.
func (o *EASProfile) GetEasIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EasId, true
}

// SetEasId sets field value
func (o *EASProfile) SetEasId(v string) {
	o.EasId = v
}

// GetEndPt returns the EndPt field value
func (o *EASProfile) GetEndPt() EndPoint {
	if o == nil {
		var ret EndPoint
		return ret
	}

	return o.EndPt
}

// GetEndPtOk returns a tuple with the EndPt field value
// and a boolean to check if the value has been set.
func (o *EASProfile) GetEndPtOk() (*EndPoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndPt, true
}

// SetEndPt sets field value
func (o *EASProfile) SetEndPt(v EndPoint) {
	o.EndPt = v
}

// GetAcIds returns the AcIds field value if set, zero value otherwise.
func (o *EASProfile) GetAcIds() []string {
	if o == nil || IsNil(o.AcIds) {
		var ret []string
		return ret
	}
	return o.AcIds
}

// GetAcIdsOk returns a tuple with the AcIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetAcIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AcIds) {
		return nil, false
	}
	return o.AcIds, true
}

// HasAcIds returns a boolean if a field has been set.
func (o *EASProfile) HasAcIds() bool {
	if o != nil && !IsNil(o.AcIds) {
		return true
	}

	return false
}

// SetAcIds gets a reference to the given []string and assigns it to the AcIds field.
func (o *EASProfile) SetAcIds(v []string) {
	o.AcIds = v
}

// GetProvId returns the ProvId field value if set, zero value otherwise.
func (o *EASProfile) GetProvId() string {
	if o == nil || IsNil(o.ProvId) {
		var ret string
		return ret
	}
	return *o.ProvId
}

// GetProvIdOk returns a tuple with the ProvId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetProvIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProvId) {
		return nil, false
	}
	return o.ProvId, true
}

// HasProvId returns a boolean if a field has been set.
func (o *EASProfile) HasProvId() bool {
	if o != nil && !IsNil(o.ProvId) {
		return true
	}

	return false
}

// SetProvId gets a reference to the given string and assigns it to the ProvId field.
func (o *EASProfile) SetProvId(v string) {
	o.ProvId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *EASProfile) GetType() EASCategory {
	if o == nil || IsNil(o.Type) {
		var ret EASCategory
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetTypeOk() (*EASCategory, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *EASProfile) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EASCategory and assigns it to the Type field.
func (o *EASProfile) SetType(v EASCategory) {
	o.Type = &v
}

// GetScheds returns the Scheds field value if set, zero value otherwise.
func (o *EASProfile) GetScheds() []ScheduledCommunicationTime {
	if o == nil || IsNil(o.Scheds) {
		var ret []ScheduledCommunicationTime
		return ret
	}
	return o.Scheds
}

// GetSchedsOk returns a tuple with the Scheds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetSchedsOk() ([]ScheduledCommunicationTime, bool) {
	if o == nil || IsNil(o.Scheds) {
		return nil, false
	}
	return o.Scheds, true
}

// HasScheds returns a boolean if a field has been set.
func (o *EASProfile) HasScheds() bool {
	if o != nil && !IsNil(o.Scheds) {
		return true
	}

	return false
}

// SetScheds gets a reference to the given []ScheduledCommunicationTime and assigns it to the Scheds field.
func (o *EASProfile) SetScheds(v []ScheduledCommunicationTime) {
	o.Scheds = v
}

// GetSvcArea returns the SvcArea field value if set, zero value otherwise.
func (o *EASProfile) GetSvcArea() ServiceArea {
	if o == nil || IsNil(o.SvcArea) {
		var ret ServiceArea
		return ret
	}
	return *o.SvcArea
}

// GetSvcAreaOk returns a tuple with the SvcArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetSvcAreaOk() (*ServiceArea, bool) {
	if o == nil || IsNil(o.SvcArea) {
		return nil, false
	}
	return o.SvcArea, true
}

// HasSvcArea returns a boolean if a field has been set.
func (o *EASProfile) HasSvcArea() bool {
	if o != nil && !IsNil(o.SvcArea) {
		return true
	}

	return false
}

// SetSvcArea gets a reference to the given ServiceArea and assigns it to the SvcArea field.
func (o *EASProfile) SetSvcArea(v ServiceArea) {
	o.SvcArea = &v
}

// GetSvcKpi returns the SvcKpi field value if set, zero value otherwise.
func (o *EASProfile) GetSvcKpi() EASServiceKPI {
	if o == nil || IsNil(o.SvcKpi) {
		var ret EASServiceKPI
		return ret
	}
	return *o.SvcKpi
}

// GetSvcKpiOk returns a tuple with the SvcKpi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetSvcKpiOk() (*EASServiceKPI, bool) {
	if o == nil || IsNil(o.SvcKpi) {
		return nil, false
	}
	return o.SvcKpi, true
}

// HasSvcKpi returns a boolean if a field has been set.
func (o *EASProfile) HasSvcKpi() bool {
	if o != nil && !IsNil(o.SvcKpi) {
		return true
	}

	return false
}

// SetSvcKpi gets a reference to the given EASServiceKPI and assigns it to the SvcKpi field.
func (o *EASProfile) SetSvcKpi(v EASServiceKPI) {
	o.SvcKpi = &v
}

// GetPermLvl returns the PermLvl field value if set, zero value otherwise.
func (o *EASProfile) GetPermLvl() []PermissionLevel {
	if o == nil || IsNil(o.PermLvl) {
		var ret []PermissionLevel
		return ret
	}
	return o.PermLvl
}

// GetPermLvlOk returns a tuple with the PermLvl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetPermLvlOk() ([]PermissionLevel, bool) {
	if o == nil || IsNil(o.PermLvl) {
		return nil, false
	}
	return o.PermLvl, true
}

// HasPermLvl returns a boolean if a field has been set.
func (o *EASProfile) HasPermLvl() bool {
	if o != nil && !IsNil(o.PermLvl) {
		return true
	}

	return false
}

// SetPermLvl gets a reference to the given []PermissionLevel and assigns it to the PermLvl field.
func (o *EASProfile) SetPermLvl(v []PermissionLevel) {
	o.PermLvl = v
}

// GetEasFeats returns the EasFeats field value if set, zero value otherwise.
func (o *EASProfile) GetEasFeats() []string {
	if o == nil || IsNil(o.EasFeats) {
		var ret []string
		return ret
	}
	return o.EasFeats
}

// GetEasFeatsOk returns a tuple with the EasFeats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetEasFeatsOk() ([]string, bool) {
	if o == nil || IsNil(o.EasFeats) {
		return nil, false
	}
	return o.EasFeats, true
}

// HasEasFeats returns a boolean if a field has been set.
func (o *EASProfile) HasEasFeats() bool {
	if o != nil && !IsNil(o.EasFeats) {
		return true
	}

	return false
}

// SetEasFeats gets a reference to the given []string and assigns it to the EasFeats field.
func (o *EASProfile) SetEasFeats(v []string) {
	o.EasFeats = v
}

// GetAppLocs returns the AppLocs field value if set, zero value otherwise.
func (o *EASProfile) GetAppLocs() []RouteToLocation {
	if o == nil || IsNil(o.AppLocs) {
		var ret []RouteToLocation
		return ret
	}
	return o.AppLocs
}

// GetAppLocsOk returns a tuple with the AppLocs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetAppLocsOk() ([]RouteToLocation, bool) {
	if o == nil || IsNil(o.AppLocs) {
		return nil, false
	}
	return o.AppLocs, true
}

// HasAppLocs returns a boolean if a field has been set.
func (o *EASProfile) HasAppLocs() bool {
	if o != nil && !IsNil(o.AppLocs) {
		return true
	}

	return false
}

// SetAppLocs gets a reference to the given []RouteToLocation and assigns it to the AppLocs field.
func (o *EASProfile) SetAppLocs(v []RouteToLocation) {
	o.AppLocs = v
}

// GetSvcContSupp returns the SvcContSupp field value if set, zero value otherwise.
func (o *EASProfile) GetSvcContSupp() []ACRScenario {
	if o == nil || IsNil(o.SvcContSupp) {
		var ret []ACRScenario
		return ret
	}
	return o.SvcContSupp
}

// GetSvcContSuppOk returns a tuple with the SvcContSupp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetSvcContSuppOk() ([]ACRScenario, bool) {
	if o == nil || IsNil(o.SvcContSupp) {
		return nil, false
	}
	return o.SvcContSupp, true
}

// HasSvcContSupp returns a boolean if a field has been set.
func (o *EASProfile) HasSvcContSupp() bool {
	if o != nil && !IsNil(o.SvcContSupp) {
		return true
	}

	return false
}

// SetSvcContSupp gets a reference to the given []ACRScenario and assigns it to the SvcContSupp field.
func (o *EASProfile) SetSvcContSupp(v []ACRScenario) {
	o.SvcContSupp = v
}

// GetAvlRep returns the AvlRep field value if set, zero value otherwise.
func (o *EASProfile) GetAvlRep() int32 {
	if o == nil || IsNil(o.AvlRep) {
		var ret int32
		return ret
	}
	return *o.AvlRep
}

// GetAvlRepOk returns a tuple with the AvlRep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetAvlRepOk() (*int32, bool) {
	if o == nil || IsNil(o.AvlRep) {
		return nil, false
	}
	return o.AvlRep, true
}

// HasAvlRep returns a boolean if a field has been set.
func (o *EASProfile) HasAvlRep() bool {
	if o != nil && !IsNil(o.AvlRep) {
		return true
	}

	return false
}

// SetAvlRep gets a reference to the given int32 and assigns it to the AvlRep field.
func (o *EASProfile) SetAvlRep(v int32) {
	o.AvlRep = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EASProfile) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASProfile) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EASProfile) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *EASProfile) SetStatus(v string) {
	o.Status = &v
}

func (o EASProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EASProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["easId"] = o.EasId
	toSerialize["endPt"] = o.EndPt
	if !IsNil(o.AcIds) {
		toSerialize["acIds"] = o.AcIds
	}
	if !IsNil(o.ProvId) {
		toSerialize["provId"] = o.ProvId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Scheds) {
		toSerialize["scheds"] = o.Scheds
	}
	if !IsNil(o.SvcArea) {
		toSerialize["svcArea"] = o.SvcArea
	}
	if !IsNil(o.SvcKpi) {
		toSerialize["svcKpi"] = o.SvcKpi
	}
	if !IsNil(o.PermLvl) {
		toSerialize["permLvl"] = o.PermLvl
	}
	if !IsNil(o.EasFeats) {
		toSerialize["easFeats"] = o.EasFeats
	}
	if !IsNil(o.AppLocs) {
		toSerialize["appLocs"] = o.AppLocs
	}
	if !IsNil(o.SvcContSupp) {
		toSerialize["svcContSupp"] = o.SvcContSupp
	}
	if !IsNil(o.AvlRep) {
		toSerialize["avlRep"] = o.AvlRep
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableEASProfile struct {
	value *EASProfile
	isSet bool
}

func (v NullableEASProfile) Get() *EASProfile {
	return v.value
}

func (v *NullableEASProfile) Set(val *EASProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableEASProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableEASProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASProfile(val *EASProfile) *NullableEASProfile {
	return &NullableEASProfile{value: val, isSet: true}
}

func (v NullableEASProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


