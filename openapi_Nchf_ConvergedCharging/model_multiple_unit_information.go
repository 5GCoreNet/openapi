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

// checks if the MultipleUnitInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipleUnitInformation{}

// MultipleUnitInformation struct for MultipleUnitInformation
type MultipleUnitInformation struct {
	ResultCode *ResultCode `json:"resultCode,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer.
	RatingGroup int32        `json:"ratingGroup"`
	GrantedUnit *GrantedUnit `json:"grantedUnit,omitempty"`
	Triggers    []Trigger    `json:"triggers,omitempty"`
	// indicating a time in seconds.
	ValidityTime *int32 `json:"validityTime,omitempty"`
	// indicating a time in seconds.
	QuotaHoldingTime    *int32               `json:"quotaHoldingTime,omitempty"`
	FinalUnitIndication *FinalUnitIndication `json:"finalUnitIndication,omitempty"`
	TimeQuotaThreshold  *int32               `json:"timeQuotaThreshold,omitempty"`
	// Integer where the allowed values correspond to the value range of an unsigned 64-bit integer.
	VolumeQuotaThreshold *int32 `json:"volumeQuotaThreshold,omitempty"`
	UnitQuotaThreshold   *int32 `json:"unitQuotaThreshold,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	UPFID                   *string                  `json:"uPFID,omitempty"`
	AnnouncementInformation *AnnouncementInformation `json:"announcementInformation,omitempty"`
}

// NewMultipleUnitInformation instantiates a new MultipleUnitInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleUnitInformation(ratingGroup int32) *MultipleUnitInformation {
	this := MultipleUnitInformation{}
	this.RatingGroup = ratingGroup
	return &this
}

// NewMultipleUnitInformationWithDefaults instantiates a new MultipleUnitInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleUnitInformationWithDefaults() *MultipleUnitInformation {
	this := MultipleUnitInformation{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetResultCode() ResultCode {
	if o == nil || IsNil(o.ResultCode) {
		var ret ResultCode
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetResultCodeOk() (*ResultCode, bool) {
	if o == nil || IsNil(o.ResultCode) {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasResultCode() bool {
	if o != nil && !IsNil(o.ResultCode) {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given ResultCode and assigns it to the ResultCode field.
func (o *MultipleUnitInformation) SetResultCode(v ResultCode) {
	o.ResultCode = &v
}

// GetRatingGroup returns the RatingGroup field value
func (o *MultipleUnitInformation) GetRatingGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RatingGroup
}

// GetRatingGroupOk returns a tuple with the RatingGroup field value
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetRatingGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatingGroup, true
}

// SetRatingGroup sets field value
func (o *MultipleUnitInformation) SetRatingGroup(v int32) {
	o.RatingGroup = v
}

// GetGrantedUnit returns the GrantedUnit field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetGrantedUnit() GrantedUnit {
	if o == nil || IsNil(o.GrantedUnit) {
		var ret GrantedUnit
		return ret
	}
	return *o.GrantedUnit
}

// GetGrantedUnitOk returns a tuple with the GrantedUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetGrantedUnitOk() (*GrantedUnit, bool) {
	if o == nil || IsNil(o.GrantedUnit) {
		return nil, false
	}
	return o.GrantedUnit, true
}

// HasGrantedUnit returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasGrantedUnit() bool {
	if o != nil && !IsNil(o.GrantedUnit) {
		return true
	}

	return false
}

// SetGrantedUnit gets a reference to the given GrantedUnit and assigns it to the GrantedUnit field.
func (o *MultipleUnitInformation) SetGrantedUnit(v GrantedUnit) {
	o.GrantedUnit = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetTriggers() []Trigger {
	if o == nil || IsNil(o.Triggers) {
		var ret []Trigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetTriggersOk() ([]Trigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []Trigger and assigns it to the Triggers field.
func (o *MultipleUnitInformation) SetTriggers(v []Trigger) {
	o.Triggers = v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetValidityTime() int32 {
	if o == nil || IsNil(o.ValidityTime) {
		var ret int32
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetValidityTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given int32 and assigns it to the ValidityTime field.
func (o *MultipleUnitInformation) SetValidityTime(v int32) {
	o.ValidityTime = &v
}

// GetQuotaHoldingTime returns the QuotaHoldingTime field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetQuotaHoldingTime() int32 {
	if o == nil || IsNil(o.QuotaHoldingTime) {
		var ret int32
		return ret
	}
	return *o.QuotaHoldingTime
}

// GetQuotaHoldingTimeOk returns a tuple with the QuotaHoldingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetQuotaHoldingTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.QuotaHoldingTime) {
		return nil, false
	}
	return o.QuotaHoldingTime, true
}

// HasQuotaHoldingTime returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasQuotaHoldingTime() bool {
	if o != nil && !IsNil(o.QuotaHoldingTime) {
		return true
	}

	return false
}

// SetQuotaHoldingTime gets a reference to the given int32 and assigns it to the QuotaHoldingTime field.
func (o *MultipleUnitInformation) SetQuotaHoldingTime(v int32) {
	o.QuotaHoldingTime = &v
}

// GetFinalUnitIndication returns the FinalUnitIndication field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetFinalUnitIndication() FinalUnitIndication {
	if o == nil || IsNil(o.FinalUnitIndication) {
		var ret FinalUnitIndication
		return ret
	}
	return *o.FinalUnitIndication
}

// GetFinalUnitIndicationOk returns a tuple with the FinalUnitIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetFinalUnitIndicationOk() (*FinalUnitIndication, bool) {
	if o == nil || IsNil(o.FinalUnitIndication) {
		return nil, false
	}
	return o.FinalUnitIndication, true
}

// HasFinalUnitIndication returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasFinalUnitIndication() bool {
	if o != nil && !IsNil(o.FinalUnitIndication) {
		return true
	}

	return false
}

// SetFinalUnitIndication gets a reference to the given FinalUnitIndication and assigns it to the FinalUnitIndication field.
func (o *MultipleUnitInformation) SetFinalUnitIndication(v FinalUnitIndication) {
	o.FinalUnitIndication = &v
}

// GetTimeQuotaThreshold returns the TimeQuotaThreshold field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetTimeQuotaThreshold() int32 {
	if o == nil || IsNil(o.TimeQuotaThreshold) {
		var ret int32
		return ret
	}
	return *o.TimeQuotaThreshold
}

// GetTimeQuotaThresholdOk returns a tuple with the TimeQuotaThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetTimeQuotaThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeQuotaThreshold) {
		return nil, false
	}
	return o.TimeQuotaThreshold, true
}

// HasTimeQuotaThreshold returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasTimeQuotaThreshold() bool {
	if o != nil && !IsNil(o.TimeQuotaThreshold) {
		return true
	}

	return false
}

// SetTimeQuotaThreshold gets a reference to the given int32 and assigns it to the TimeQuotaThreshold field.
func (o *MultipleUnitInformation) SetTimeQuotaThreshold(v int32) {
	o.TimeQuotaThreshold = &v
}

// GetVolumeQuotaThreshold returns the VolumeQuotaThreshold field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetVolumeQuotaThreshold() int32 {
	if o == nil || IsNil(o.VolumeQuotaThreshold) {
		var ret int32
		return ret
	}
	return *o.VolumeQuotaThreshold
}

// GetVolumeQuotaThresholdOk returns a tuple with the VolumeQuotaThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetVolumeQuotaThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.VolumeQuotaThreshold) {
		return nil, false
	}
	return o.VolumeQuotaThreshold, true
}

// HasVolumeQuotaThreshold returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasVolumeQuotaThreshold() bool {
	if o != nil && !IsNil(o.VolumeQuotaThreshold) {
		return true
	}

	return false
}

// SetVolumeQuotaThreshold gets a reference to the given int32 and assigns it to the VolumeQuotaThreshold field.
func (o *MultipleUnitInformation) SetVolumeQuotaThreshold(v int32) {
	o.VolumeQuotaThreshold = &v
}

// GetUnitQuotaThreshold returns the UnitQuotaThreshold field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetUnitQuotaThreshold() int32 {
	if o == nil || IsNil(o.UnitQuotaThreshold) {
		var ret int32
		return ret
	}
	return *o.UnitQuotaThreshold
}

// GetUnitQuotaThresholdOk returns a tuple with the UnitQuotaThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetUnitQuotaThresholdOk() (*int32, bool) {
	if o == nil || IsNil(o.UnitQuotaThreshold) {
		return nil, false
	}
	return o.UnitQuotaThreshold, true
}

// HasUnitQuotaThreshold returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasUnitQuotaThreshold() bool {
	if o != nil && !IsNil(o.UnitQuotaThreshold) {
		return true
	}

	return false
}

// SetUnitQuotaThreshold gets a reference to the given int32 and assigns it to the UnitQuotaThreshold field.
func (o *MultipleUnitInformation) SetUnitQuotaThreshold(v int32) {
	o.UnitQuotaThreshold = &v
}

// GetUPFID returns the UPFID field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetUPFID() string {
	if o == nil || IsNil(o.UPFID) {
		var ret string
		return ret
	}
	return *o.UPFID
}

// GetUPFIDOk returns a tuple with the UPFID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetUPFIDOk() (*string, bool) {
	if o == nil || IsNil(o.UPFID) {
		return nil, false
	}
	return o.UPFID, true
}

// HasUPFID returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasUPFID() bool {
	if o != nil && !IsNil(o.UPFID) {
		return true
	}

	return false
}

// SetUPFID gets a reference to the given string and assigns it to the UPFID field.
func (o *MultipleUnitInformation) SetUPFID(v string) {
	o.UPFID = &v
}

// GetAnnouncementInformation returns the AnnouncementInformation field value if set, zero value otherwise.
func (o *MultipleUnitInformation) GetAnnouncementInformation() AnnouncementInformation {
	if o == nil || IsNil(o.AnnouncementInformation) {
		var ret AnnouncementInformation
		return ret
	}
	return *o.AnnouncementInformation
}

// GetAnnouncementInformationOk returns a tuple with the AnnouncementInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitInformation) GetAnnouncementInformationOk() (*AnnouncementInformation, bool) {
	if o == nil || IsNil(o.AnnouncementInformation) {
		return nil, false
	}
	return o.AnnouncementInformation, true
}

// HasAnnouncementInformation returns a boolean if a field has been set.
func (o *MultipleUnitInformation) HasAnnouncementInformation() bool {
	if o != nil && !IsNil(o.AnnouncementInformation) {
		return true
	}

	return false
}

// SetAnnouncementInformation gets a reference to the given AnnouncementInformation and assigns it to the AnnouncementInformation field.
func (o *MultipleUnitInformation) SetAnnouncementInformation(v AnnouncementInformation) {
	o.AnnouncementInformation = &v
}

func (o MultipleUnitInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipleUnitInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResultCode) {
		toSerialize["resultCode"] = o.ResultCode
	}
	toSerialize["ratingGroup"] = o.RatingGroup
	if !IsNil(o.GrantedUnit) {
		toSerialize["grantedUnit"] = o.GrantedUnit
	}
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !IsNil(o.QuotaHoldingTime) {
		toSerialize["quotaHoldingTime"] = o.QuotaHoldingTime
	}
	if !IsNil(o.FinalUnitIndication) {
		toSerialize["finalUnitIndication"] = o.FinalUnitIndication
	}
	if !IsNil(o.TimeQuotaThreshold) {
		toSerialize["timeQuotaThreshold"] = o.TimeQuotaThreshold
	}
	if !IsNil(o.VolumeQuotaThreshold) {
		toSerialize["volumeQuotaThreshold"] = o.VolumeQuotaThreshold
	}
	if !IsNil(o.UnitQuotaThreshold) {
		toSerialize["unitQuotaThreshold"] = o.UnitQuotaThreshold
	}
	if !IsNil(o.UPFID) {
		toSerialize["uPFID"] = o.UPFID
	}
	if !IsNil(o.AnnouncementInformation) {
		toSerialize["announcementInformation"] = o.AnnouncementInformation
	}
	return toSerialize, nil
}

type NullableMultipleUnitInformation struct {
	value *MultipleUnitInformation
	isSet bool
}

func (v NullableMultipleUnitInformation) Get() *MultipleUnitInformation {
	return v.value
}

func (v *NullableMultipleUnitInformation) Set(val *MultipleUnitInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleUnitInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleUnitInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleUnitInformation(val *MultipleUnitInformation) *NullableMultipleUnitInformation {
	return &NullableMultipleUnitInformation{value: val, isSet: true}
}

func (v NullableMultipleUnitInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleUnitInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
