/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the N2ConnectionChargingInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N2ConnectionChargingInformation{}

// N2ConnectionChargingInformation struct for N2ConnectionChargingInformation
type N2ConnectionChargingInformation struct {
	N2ConnectionMessageType int32              `json:"n2ConnectionMessageType"`
	UserInformation         *UserInformation   `json:"userInformation,omitempty"`
	UserLocationinfo        *UserLocation      `json:"userLocationinfo,omitempty"`
	PSCellInformation       *PSCellInformation `json:"pSCellInformation,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	UetimeZone             *string                  `json:"uetimeZone,omitempty"`
	RATType                *RatType                 `json:"rATType,omitempty"`
	AmfUeNgapId            *int32                   `json:"amfUeNgapId,omitempty"`
	RanUeNgapId            *int32                   `json:"ranUeNgapId,omitempty"`
	RanNodeId              *GlobalRanNodeId         `json:"ranNodeId,omitempty"`
	RestrictedRatList      []RatType                `json:"restrictedRatList,omitempty"`
	ForbiddenAreaList      []Area                   `json:"forbiddenAreaList,omitempty"`
	ServiceAreaRestriction []ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`
	RestrictedCnList       []CoreNetworkType        `json:"restrictedCnList,omitempty"`
	AllowedNSSAI           []Snssai                 `json:"allowedNSSAI,omitempty"`
	RrcEstCause            *string                  `json:"rrcEstCause,omitempty"`
}

// NewN2ConnectionChargingInformation instantiates a new N2ConnectionChargingInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN2ConnectionChargingInformation(n2ConnectionMessageType int32) *N2ConnectionChargingInformation {
	this := N2ConnectionChargingInformation{}
	this.N2ConnectionMessageType = n2ConnectionMessageType
	return &this
}

// NewN2ConnectionChargingInformationWithDefaults instantiates a new N2ConnectionChargingInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN2ConnectionChargingInformationWithDefaults() *N2ConnectionChargingInformation {
	this := N2ConnectionChargingInformation{}
	return &this
}

// GetN2ConnectionMessageType returns the N2ConnectionMessageType field value
func (o *N2ConnectionChargingInformation) GetN2ConnectionMessageType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.N2ConnectionMessageType
}

// GetN2ConnectionMessageTypeOk returns a tuple with the N2ConnectionMessageType field value
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetN2ConnectionMessageTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N2ConnectionMessageType, true
}

// SetN2ConnectionMessageType sets field value
func (o *N2ConnectionChargingInformation) SetN2ConnectionMessageType(v int32) {
	o.N2ConnectionMessageType = v
}

// GetUserInformation returns the UserInformation field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetUserInformation() UserInformation {
	if o == nil || IsNil(o.UserInformation) {
		var ret UserInformation
		return ret
	}
	return *o.UserInformation
}

// GetUserInformationOk returns a tuple with the UserInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetUserInformationOk() (*UserInformation, bool) {
	if o == nil || IsNil(o.UserInformation) {
		return nil, false
	}
	return o.UserInformation, true
}

// HasUserInformation returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasUserInformation() bool {
	if o != nil && !IsNil(o.UserInformation) {
		return true
	}

	return false
}

// SetUserInformation gets a reference to the given UserInformation and assigns it to the UserInformation field.
func (o *N2ConnectionChargingInformation) SetUserInformation(v UserInformation) {
	o.UserInformation = &v
}

// GetUserLocationinfo returns the UserLocationinfo field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetUserLocationinfo() UserLocation {
	if o == nil || IsNil(o.UserLocationinfo) {
		var ret UserLocation
		return ret
	}
	return *o.UserLocationinfo
}

// GetUserLocationinfoOk returns a tuple with the UserLocationinfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetUserLocationinfoOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLocationinfo) {
		return nil, false
	}
	return o.UserLocationinfo, true
}

// HasUserLocationinfo returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasUserLocationinfo() bool {
	if o != nil && !IsNil(o.UserLocationinfo) {
		return true
	}

	return false
}

// SetUserLocationinfo gets a reference to the given UserLocation and assigns it to the UserLocationinfo field.
func (o *N2ConnectionChargingInformation) SetUserLocationinfo(v UserLocation) {
	o.UserLocationinfo = &v
}

// GetPSCellInformation returns the PSCellInformation field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetPSCellInformation() PSCellInformation {
	if o == nil || IsNil(o.PSCellInformation) {
		var ret PSCellInformation
		return ret
	}
	return *o.PSCellInformation
}

// GetPSCellInformationOk returns a tuple with the PSCellInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetPSCellInformationOk() (*PSCellInformation, bool) {
	if o == nil || IsNil(o.PSCellInformation) {
		return nil, false
	}
	return o.PSCellInformation, true
}

// HasPSCellInformation returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasPSCellInformation() bool {
	if o != nil && !IsNil(o.PSCellInformation) {
		return true
	}

	return false
}

// SetPSCellInformation gets a reference to the given PSCellInformation and assigns it to the PSCellInformation field.
func (o *N2ConnectionChargingInformation) SetPSCellInformation(v PSCellInformation) {
	o.PSCellInformation = &v
}

// GetUetimeZone returns the UetimeZone field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetUetimeZone() string {
	if o == nil || IsNil(o.UetimeZone) {
		var ret string
		return ret
	}
	return *o.UetimeZone
}

// GetUetimeZoneOk returns a tuple with the UetimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetUetimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.UetimeZone) {
		return nil, false
	}
	return o.UetimeZone, true
}

// HasUetimeZone returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasUetimeZone() bool {
	if o != nil && !IsNil(o.UetimeZone) {
		return true
	}

	return false
}

// SetUetimeZone gets a reference to the given string and assigns it to the UetimeZone field.
func (o *N2ConnectionChargingInformation) SetUetimeZone(v string) {
	o.UetimeZone = &v
}

// GetRATType returns the RATType field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetRATType() RatType {
	if o == nil || IsNil(o.RATType) {
		var ret RatType
		return ret
	}
	return *o.RATType
}

// GetRATTypeOk returns a tuple with the RATType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetRATTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RATType) {
		return nil, false
	}
	return o.RATType, true
}

// HasRATType returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasRATType() bool {
	if o != nil && !IsNil(o.RATType) {
		return true
	}

	return false
}

// SetRATType gets a reference to the given RatType and assigns it to the RATType field.
func (o *N2ConnectionChargingInformation) SetRATType(v RatType) {
	o.RATType = &v
}

// GetAmfUeNgapId returns the AmfUeNgapId field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetAmfUeNgapId() int32 {
	if o == nil || IsNil(o.AmfUeNgapId) {
		var ret int32
		return ret
	}
	return *o.AmfUeNgapId
}

// GetAmfUeNgapIdOk returns a tuple with the AmfUeNgapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetAmfUeNgapIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AmfUeNgapId) {
		return nil, false
	}
	return o.AmfUeNgapId, true
}

// HasAmfUeNgapId returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasAmfUeNgapId() bool {
	if o != nil && !IsNil(o.AmfUeNgapId) {
		return true
	}

	return false
}

// SetAmfUeNgapId gets a reference to the given int32 and assigns it to the AmfUeNgapId field.
func (o *N2ConnectionChargingInformation) SetAmfUeNgapId(v int32) {
	o.AmfUeNgapId = &v
}

// GetRanUeNgapId returns the RanUeNgapId field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetRanUeNgapId() int32 {
	if o == nil || IsNil(o.RanUeNgapId) {
		var ret int32
		return ret
	}
	return *o.RanUeNgapId
}

// GetRanUeNgapIdOk returns a tuple with the RanUeNgapId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetRanUeNgapIdOk() (*int32, bool) {
	if o == nil || IsNil(o.RanUeNgapId) {
		return nil, false
	}
	return o.RanUeNgapId, true
}

// HasRanUeNgapId returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasRanUeNgapId() bool {
	if o != nil && !IsNil(o.RanUeNgapId) {
		return true
	}

	return false
}

// SetRanUeNgapId gets a reference to the given int32 and assigns it to the RanUeNgapId field.
func (o *N2ConnectionChargingInformation) SetRanUeNgapId(v int32) {
	o.RanUeNgapId = &v
}

// GetRanNodeId returns the RanNodeId field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetRanNodeId() GlobalRanNodeId {
	if o == nil || IsNil(o.RanNodeId) {
		var ret GlobalRanNodeId
		return ret
	}
	return *o.RanNodeId
}

// GetRanNodeIdOk returns a tuple with the RanNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetRanNodeIdOk() (*GlobalRanNodeId, bool) {
	if o == nil || IsNil(o.RanNodeId) {
		return nil, false
	}
	return o.RanNodeId, true
}

// HasRanNodeId returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasRanNodeId() bool {
	if o != nil && !IsNil(o.RanNodeId) {
		return true
	}

	return false
}

// SetRanNodeId gets a reference to the given GlobalRanNodeId and assigns it to the RanNodeId field.
func (o *N2ConnectionChargingInformation) SetRanNodeId(v GlobalRanNodeId) {
	o.RanNodeId = &v
}

// GetRestrictedRatList returns the RestrictedRatList field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetRestrictedRatList() []RatType {
	if o == nil || IsNil(o.RestrictedRatList) {
		var ret []RatType
		return ret
	}
	return o.RestrictedRatList
}

// GetRestrictedRatListOk returns a tuple with the RestrictedRatList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetRestrictedRatListOk() ([]RatType, bool) {
	if o == nil || IsNil(o.RestrictedRatList) {
		return nil, false
	}
	return o.RestrictedRatList, true
}

// HasRestrictedRatList returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasRestrictedRatList() bool {
	if o != nil && !IsNil(o.RestrictedRatList) {
		return true
	}

	return false
}

// SetRestrictedRatList gets a reference to the given []RatType and assigns it to the RestrictedRatList field.
func (o *N2ConnectionChargingInformation) SetRestrictedRatList(v []RatType) {
	o.RestrictedRatList = v
}

// GetForbiddenAreaList returns the ForbiddenAreaList field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetForbiddenAreaList() []Area {
	if o == nil || IsNil(o.ForbiddenAreaList) {
		var ret []Area
		return ret
	}
	return o.ForbiddenAreaList
}

// GetForbiddenAreaListOk returns a tuple with the ForbiddenAreaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetForbiddenAreaListOk() ([]Area, bool) {
	if o == nil || IsNil(o.ForbiddenAreaList) {
		return nil, false
	}
	return o.ForbiddenAreaList, true
}

// HasForbiddenAreaList returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasForbiddenAreaList() bool {
	if o != nil && !IsNil(o.ForbiddenAreaList) {
		return true
	}

	return false
}

// SetForbiddenAreaList gets a reference to the given []Area and assigns it to the ForbiddenAreaList field.
func (o *N2ConnectionChargingInformation) SetForbiddenAreaList(v []Area) {
	o.ForbiddenAreaList = v
}

// GetServiceAreaRestriction returns the ServiceAreaRestriction field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetServiceAreaRestriction() []ServiceAreaRestriction {
	if o == nil || IsNil(o.ServiceAreaRestriction) {
		var ret []ServiceAreaRestriction
		return ret
	}
	return o.ServiceAreaRestriction
}

// GetServiceAreaRestrictionOk returns a tuple with the ServiceAreaRestriction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetServiceAreaRestrictionOk() ([]ServiceAreaRestriction, bool) {
	if o == nil || IsNil(o.ServiceAreaRestriction) {
		return nil, false
	}
	return o.ServiceAreaRestriction, true
}

// HasServiceAreaRestriction returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasServiceAreaRestriction() bool {
	if o != nil && !IsNil(o.ServiceAreaRestriction) {
		return true
	}

	return false
}

// SetServiceAreaRestriction gets a reference to the given []ServiceAreaRestriction and assigns it to the ServiceAreaRestriction field.
func (o *N2ConnectionChargingInformation) SetServiceAreaRestriction(v []ServiceAreaRestriction) {
	o.ServiceAreaRestriction = v
}

// GetRestrictedCnList returns the RestrictedCnList field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetRestrictedCnList() []CoreNetworkType {
	if o == nil || IsNil(o.RestrictedCnList) {
		var ret []CoreNetworkType
		return ret
	}
	return o.RestrictedCnList
}

// GetRestrictedCnListOk returns a tuple with the RestrictedCnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetRestrictedCnListOk() ([]CoreNetworkType, bool) {
	if o == nil || IsNil(o.RestrictedCnList) {
		return nil, false
	}
	return o.RestrictedCnList, true
}

// HasRestrictedCnList returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasRestrictedCnList() bool {
	if o != nil && !IsNil(o.RestrictedCnList) {
		return true
	}

	return false
}

// SetRestrictedCnList gets a reference to the given []CoreNetworkType and assigns it to the RestrictedCnList field.
func (o *N2ConnectionChargingInformation) SetRestrictedCnList(v []CoreNetworkType) {
	o.RestrictedCnList = v
}

// GetAllowedNSSAI returns the AllowedNSSAI field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetAllowedNSSAI() []Snssai {
	if o == nil || IsNil(o.AllowedNSSAI) {
		var ret []Snssai
		return ret
	}
	return o.AllowedNSSAI
}

// GetAllowedNSSAIOk returns a tuple with the AllowedNSSAI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetAllowedNSSAIOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.AllowedNSSAI) {
		return nil, false
	}
	return o.AllowedNSSAI, true
}

// HasAllowedNSSAI returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasAllowedNSSAI() bool {
	if o != nil && !IsNil(o.AllowedNSSAI) {
		return true
	}

	return false
}

// SetAllowedNSSAI gets a reference to the given []Snssai and assigns it to the AllowedNSSAI field.
func (o *N2ConnectionChargingInformation) SetAllowedNSSAI(v []Snssai) {
	o.AllowedNSSAI = v
}

// GetRrcEstCause returns the RrcEstCause field value if set, zero value otherwise.
func (o *N2ConnectionChargingInformation) GetRrcEstCause() string {
	if o == nil || IsNil(o.RrcEstCause) {
		var ret string
		return ret
	}
	return *o.RrcEstCause
}

// GetRrcEstCauseOk returns a tuple with the RrcEstCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N2ConnectionChargingInformation) GetRrcEstCauseOk() (*string, bool) {
	if o == nil || IsNil(o.RrcEstCause) {
		return nil, false
	}
	return o.RrcEstCause, true
}

// HasRrcEstCause returns a boolean if a field has been set.
func (o *N2ConnectionChargingInformation) HasRrcEstCause() bool {
	if o != nil && !IsNil(o.RrcEstCause) {
		return true
	}

	return false
}

// SetRrcEstCause gets a reference to the given string and assigns it to the RrcEstCause field.
func (o *N2ConnectionChargingInformation) SetRrcEstCause(v string) {
	o.RrcEstCause = &v
}

func (o N2ConnectionChargingInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N2ConnectionChargingInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["n2ConnectionMessageType"] = o.N2ConnectionMessageType
	if !IsNil(o.UserInformation) {
		toSerialize["userInformation"] = o.UserInformation
	}
	if !IsNil(o.UserLocationinfo) {
		toSerialize["userLocationinfo"] = o.UserLocationinfo
	}
	if !IsNil(o.PSCellInformation) {
		toSerialize["pSCellInformation"] = o.PSCellInformation
	}
	if !IsNil(o.UetimeZone) {
		toSerialize["uetimeZone"] = o.UetimeZone
	}
	if !IsNil(o.RATType) {
		toSerialize["rATType"] = o.RATType
	}
	if !IsNil(o.AmfUeNgapId) {
		toSerialize["amfUeNgapId"] = o.AmfUeNgapId
	}
	if !IsNil(o.RanUeNgapId) {
		toSerialize["ranUeNgapId"] = o.RanUeNgapId
	}
	if !IsNil(o.RanNodeId) {
		toSerialize["ranNodeId"] = o.RanNodeId
	}
	if !IsNil(o.RestrictedRatList) {
		toSerialize["restrictedRatList"] = o.RestrictedRatList
	}
	if !IsNil(o.ForbiddenAreaList) {
		toSerialize["forbiddenAreaList"] = o.ForbiddenAreaList
	}
	if !IsNil(o.ServiceAreaRestriction) {
		toSerialize["serviceAreaRestriction"] = o.ServiceAreaRestriction
	}
	if !IsNil(o.RestrictedCnList) {
		toSerialize["restrictedCnList"] = o.RestrictedCnList
	}
	if !IsNil(o.AllowedNSSAI) {
		toSerialize["allowedNSSAI"] = o.AllowedNSSAI
	}
	if !IsNil(o.RrcEstCause) {
		toSerialize["rrcEstCause"] = o.RrcEstCause
	}
	return toSerialize, nil
}

type NullableN2ConnectionChargingInformation struct {
	value *N2ConnectionChargingInformation
	isSet bool
}

func (v NullableN2ConnectionChargingInformation) Get() *N2ConnectionChargingInformation {
	return v.value
}

func (v *NullableN2ConnectionChargingInformation) Set(val *N2ConnectionChargingInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableN2ConnectionChargingInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableN2ConnectionChargingInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2ConnectionChargingInformation(val *N2ConnectionChargingInformation) *NullableN2ConnectionChargingInformation {
	return &NullableN2ConnectionChargingInformation{value: val, isSet: true}
}

func (v NullableN2ConnectionChargingInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2ConnectionChargingInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
