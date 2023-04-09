/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
)

// checks if the ACProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ACProfile{}

// ACProfile ECS service provisioning response information.
type ACProfile struct {
	// Identity of the AC.
	AcId string `json:"acId"`
	// The category or type of AC.
	AcType *string `json:"acType,omitempty"`
	// Indicates to the ECS which ECSPs are preferred for the AC.
	PrefEcsps        []string                    `json:"prefEcsps,omitempty"`
	AcSchedule       *ScheduledCommunicationTime `json:"acSchedule,omitempty"`
	ExpAcGeoServArea *LocationArea5G             `json:"expAcGeoServArea,omitempty"`
	// Profiles of ACs for which the EEC provides edge enabling services.
	AcSvcContSupp []ACRScenario `json:"acSvcContSupp,omitempty"`
	// List of EAS information.
	Eass []EasDetail `json:"eass,omitempty"`
}

// NewACProfile instantiates a new ACProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewACProfile(acId string) *ACProfile {
	this := ACProfile{}
	this.AcId = acId
	return &this
}

// NewACProfileWithDefaults instantiates a new ACProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewACProfileWithDefaults() *ACProfile {
	this := ACProfile{}
	return &this
}

// GetAcId returns the AcId field value
func (o *ACProfile) GetAcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AcId
}

// GetAcIdOk returns a tuple with the AcId field value
// and a boolean to check if the value has been set.
func (o *ACProfile) GetAcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AcId, true
}

// SetAcId sets field value
func (o *ACProfile) SetAcId(v string) {
	o.AcId = v
}

// GetAcType returns the AcType field value if set, zero value otherwise.
func (o *ACProfile) GetAcType() string {
	if o == nil || IsNil(o.AcType) {
		var ret string
		return ret
	}
	return *o.AcType
}

// GetAcTypeOk returns a tuple with the AcType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACProfile) GetAcTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AcType) {
		return nil, false
	}
	return o.AcType, true
}

// HasAcType returns a boolean if a field has been set.
func (o *ACProfile) HasAcType() bool {
	if o != nil && !IsNil(o.AcType) {
		return true
	}

	return false
}

// SetAcType gets a reference to the given string and assigns it to the AcType field.
func (o *ACProfile) SetAcType(v string) {
	o.AcType = &v
}

// GetPrefEcsps returns the PrefEcsps field value if set, zero value otherwise.
func (o *ACProfile) GetPrefEcsps() []string {
	if o == nil || IsNil(o.PrefEcsps) {
		var ret []string
		return ret
	}
	return o.PrefEcsps
}

// GetPrefEcspsOk returns a tuple with the PrefEcsps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACProfile) GetPrefEcspsOk() ([]string, bool) {
	if o == nil || IsNil(o.PrefEcsps) {
		return nil, false
	}
	return o.PrefEcsps, true
}

// HasPrefEcsps returns a boolean if a field has been set.
func (o *ACProfile) HasPrefEcsps() bool {
	if o != nil && !IsNil(o.PrefEcsps) {
		return true
	}

	return false
}

// SetPrefEcsps gets a reference to the given []string and assigns it to the PrefEcsps field.
func (o *ACProfile) SetPrefEcsps(v []string) {
	o.PrefEcsps = v
}

// GetAcSchedule returns the AcSchedule field value if set, zero value otherwise.
func (o *ACProfile) GetAcSchedule() ScheduledCommunicationTime {
	if o == nil || IsNil(o.AcSchedule) {
		var ret ScheduledCommunicationTime
		return ret
	}
	return *o.AcSchedule
}

// GetAcScheduleOk returns a tuple with the AcSchedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACProfile) GetAcScheduleOk() (*ScheduledCommunicationTime, bool) {
	if o == nil || IsNil(o.AcSchedule) {
		return nil, false
	}
	return o.AcSchedule, true
}

// HasAcSchedule returns a boolean if a field has been set.
func (o *ACProfile) HasAcSchedule() bool {
	if o != nil && !IsNil(o.AcSchedule) {
		return true
	}

	return false
}

// SetAcSchedule gets a reference to the given ScheduledCommunicationTime and assigns it to the AcSchedule field.
func (o *ACProfile) SetAcSchedule(v ScheduledCommunicationTime) {
	o.AcSchedule = &v
}

// GetExpAcGeoServArea returns the ExpAcGeoServArea field value if set, zero value otherwise.
func (o *ACProfile) GetExpAcGeoServArea() LocationArea5G {
	if o == nil || IsNil(o.ExpAcGeoServArea) {
		var ret LocationArea5G
		return ret
	}
	return *o.ExpAcGeoServArea
}

// GetExpAcGeoServAreaOk returns a tuple with the ExpAcGeoServArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACProfile) GetExpAcGeoServAreaOk() (*LocationArea5G, bool) {
	if o == nil || IsNil(o.ExpAcGeoServArea) {
		return nil, false
	}
	return o.ExpAcGeoServArea, true
}

// HasExpAcGeoServArea returns a boolean if a field has been set.
func (o *ACProfile) HasExpAcGeoServArea() bool {
	if o != nil && !IsNil(o.ExpAcGeoServArea) {
		return true
	}

	return false
}

// SetExpAcGeoServArea gets a reference to the given LocationArea5G and assigns it to the ExpAcGeoServArea field.
func (o *ACProfile) SetExpAcGeoServArea(v LocationArea5G) {
	o.ExpAcGeoServArea = &v
}

// GetAcSvcContSupp returns the AcSvcContSupp field value if set, zero value otherwise.
func (o *ACProfile) GetAcSvcContSupp() []ACRScenario {
	if o == nil || IsNil(o.AcSvcContSupp) {
		var ret []ACRScenario
		return ret
	}
	return o.AcSvcContSupp
}

// GetAcSvcContSuppOk returns a tuple with the AcSvcContSupp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACProfile) GetAcSvcContSuppOk() ([]ACRScenario, bool) {
	if o == nil || IsNil(o.AcSvcContSupp) {
		return nil, false
	}
	return o.AcSvcContSupp, true
}

// HasAcSvcContSupp returns a boolean if a field has been set.
func (o *ACProfile) HasAcSvcContSupp() bool {
	if o != nil && !IsNil(o.AcSvcContSupp) {
		return true
	}

	return false
}

// SetAcSvcContSupp gets a reference to the given []ACRScenario and assigns it to the AcSvcContSupp field.
func (o *ACProfile) SetAcSvcContSupp(v []ACRScenario) {
	o.AcSvcContSupp = v
}

// GetEass returns the Eass field value if set, zero value otherwise.
func (o *ACProfile) GetEass() []EasDetail {
	if o == nil || IsNil(o.Eass) {
		var ret []EasDetail
		return ret
	}
	return o.Eass
}

// GetEassOk returns a tuple with the Eass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ACProfile) GetEassOk() ([]EasDetail, bool) {
	if o == nil || IsNil(o.Eass) {
		return nil, false
	}
	return o.Eass, true
}

// HasEass returns a boolean if a field has been set.
func (o *ACProfile) HasEass() bool {
	if o != nil && !IsNil(o.Eass) {
		return true
	}

	return false
}

// SetEass gets a reference to the given []EasDetail and assigns it to the Eass field.
func (o *ACProfile) SetEass(v []EasDetail) {
	o.Eass = v
}

func (o ACProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ACProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["acId"] = o.AcId
	if !IsNil(o.AcType) {
		toSerialize["acType"] = o.AcType
	}
	if !IsNil(o.PrefEcsps) {
		toSerialize["prefEcsps"] = o.PrefEcsps
	}
	if !IsNil(o.AcSchedule) {
		toSerialize["acSchedule"] = o.AcSchedule
	}
	if !IsNil(o.ExpAcGeoServArea) {
		toSerialize["expAcGeoServArea"] = o.ExpAcGeoServArea
	}
	if !IsNil(o.AcSvcContSupp) {
		toSerialize["acSvcContSupp"] = o.AcSvcContSupp
	}
	if !IsNil(o.Eass) {
		toSerialize["eass"] = o.Eass
	}
	return toSerialize, nil
}

type NullableACProfile struct {
	value *ACProfile
	isSet bool
}

func (v NullableACProfile) Get() *ACProfile {
	return v.value
}

func (v *NullableACProfile) Set(val *ACProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableACProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableACProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACProfile(val *ACProfile) *NullableACProfile {
	return &NullableACProfile{value: val, isSet: true}
}

func (v NullableACProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
