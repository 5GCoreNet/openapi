/*
Ndcaf_DataReportingProvisioning

Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReportingProvisioning

import (
	"encoding/json"
)

// checks if the DataAccessProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataAccessProfile{}

// DataAccessProfile A data access profile.
type DataAccessProfile struct {
	DataAccessProfileId string `json:"dataAccessProfileId"`
	TargetEventConsumerTypes []EventConsumerType `json:"targetEventConsumerTypes"`
	Parameters []string `json:"parameters"`
	TimeAccessRestrictions *DataAccessProfileTimeAccessRestrictions `json:"timeAccessRestrictions,omitempty"`
	UserAccessRestrictions *DataAccessProfileUserAccessRestrictions `json:"userAccessRestrictions,omitempty"`
	LocationAccessRestrictions *DataAccessProfileLocationAccessRestrictions `json:"locationAccessRestrictions,omitempty"`
}

// NewDataAccessProfile instantiates a new DataAccessProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAccessProfile(dataAccessProfileId string, targetEventConsumerTypes []EventConsumerType, parameters []string) *DataAccessProfile {
	this := DataAccessProfile{}
	this.DataAccessProfileId = dataAccessProfileId
	this.TargetEventConsumerTypes = targetEventConsumerTypes
	this.Parameters = parameters
	return &this
}

// NewDataAccessProfileWithDefaults instantiates a new DataAccessProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAccessProfileWithDefaults() *DataAccessProfile {
	this := DataAccessProfile{}
	return &this
}

// GetDataAccessProfileId returns the DataAccessProfileId field value
func (o *DataAccessProfile) GetDataAccessProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DataAccessProfileId
}

// GetDataAccessProfileIdOk returns a tuple with the DataAccessProfileId field value
// and a boolean to check if the value has been set.
func (o *DataAccessProfile) GetDataAccessProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataAccessProfileId, true
}

// SetDataAccessProfileId sets field value
func (o *DataAccessProfile) SetDataAccessProfileId(v string) {
	o.DataAccessProfileId = v
}

// GetTargetEventConsumerTypes returns the TargetEventConsumerTypes field value
func (o *DataAccessProfile) GetTargetEventConsumerTypes() []EventConsumerType {
	if o == nil {
		var ret []EventConsumerType
		return ret
	}

	return o.TargetEventConsumerTypes
}

// GetTargetEventConsumerTypesOk returns a tuple with the TargetEventConsumerTypes field value
// and a boolean to check if the value has been set.
func (o *DataAccessProfile) GetTargetEventConsumerTypesOk() ([]EventConsumerType, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetEventConsumerTypes, true
}

// SetTargetEventConsumerTypes sets field value
func (o *DataAccessProfile) SetTargetEventConsumerTypes(v []EventConsumerType) {
	o.TargetEventConsumerTypes = v
}

// GetParameters returns the Parameters field value
func (o *DataAccessProfile) GetParameters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *DataAccessProfile) GetParametersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Parameters, true
}

// SetParameters sets field value
func (o *DataAccessProfile) SetParameters(v []string) {
	o.Parameters = v
}

// GetTimeAccessRestrictions returns the TimeAccessRestrictions field value if set, zero value otherwise.
func (o *DataAccessProfile) GetTimeAccessRestrictions() DataAccessProfileTimeAccessRestrictions {
	if o == nil || IsNil(o.TimeAccessRestrictions) {
		var ret DataAccessProfileTimeAccessRestrictions
		return ret
	}
	return *o.TimeAccessRestrictions
}

// GetTimeAccessRestrictionsOk returns a tuple with the TimeAccessRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAccessProfile) GetTimeAccessRestrictionsOk() (*DataAccessProfileTimeAccessRestrictions, bool) {
	if o == nil || IsNil(o.TimeAccessRestrictions) {
		return nil, false
	}
	return o.TimeAccessRestrictions, true
}

// HasTimeAccessRestrictions returns a boolean if a field has been set.
func (o *DataAccessProfile) HasTimeAccessRestrictions() bool {
	if o != nil && !IsNil(o.TimeAccessRestrictions) {
		return true
	}

	return false
}

// SetTimeAccessRestrictions gets a reference to the given DataAccessProfileTimeAccessRestrictions and assigns it to the TimeAccessRestrictions field.
func (o *DataAccessProfile) SetTimeAccessRestrictions(v DataAccessProfileTimeAccessRestrictions) {
	o.TimeAccessRestrictions = &v
}

// GetUserAccessRestrictions returns the UserAccessRestrictions field value if set, zero value otherwise.
func (o *DataAccessProfile) GetUserAccessRestrictions() DataAccessProfileUserAccessRestrictions {
	if o == nil || IsNil(o.UserAccessRestrictions) {
		var ret DataAccessProfileUserAccessRestrictions
		return ret
	}
	return *o.UserAccessRestrictions
}

// GetUserAccessRestrictionsOk returns a tuple with the UserAccessRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAccessProfile) GetUserAccessRestrictionsOk() (*DataAccessProfileUserAccessRestrictions, bool) {
	if o == nil || IsNil(o.UserAccessRestrictions) {
		return nil, false
	}
	return o.UserAccessRestrictions, true
}

// HasUserAccessRestrictions returns a boolean if a field has been set.
func (o *DataAccessProfile) HasUserAccessRestrictions() bool {
	if o != nil && !IsNil(o.UserAccessRestrictions) {
		return true
	}

	return false
}

// SetUserAccessRestrictions gets a reference to the given DataAccessProfileUserAccessRestrictions and assigns it to the UserAccessRestrictions field.
func (o *DataAccessProfile) SetUserAccessRestrictions(v DataAccessProfileUserAccessRestrictions) {
	o.UserAccessRestrictions = &v
}

// GetLocationAccessRestrictions returns the LocationAccessRestrictions field value if set, zero value otherwise.
func (o *DataAccessProfile) GetLocationAccessRestrictions() DataAccessProfileLocationAccessRestrictions {
	if o == nil || IsNil(o.LocationAccessRestrictions) {
		var ret DataAccessProfileLocationAccessRestrictions
		return ret
	}
	return *o.LocationAccessRestrictions
}

// GetLocationAccessRestrictionsOk returns a tuple with the LocationAccessRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataAccessProfile) GetLocationAccessRestrictionsOk() (*DataAccessProfileLocationAccessRestrictions, bool) {
	if o == nil || IsNil(o.LocationAccessRestrictions) {
		return nil, false
	}
	return o.LocationAccessRestrictions, true
}

// HasLocationAccessRestrictions returns a boolean if a field has been set.
func (o *DataAccessProfile) HasLocationAccessRestrictions() bool {
	if o != nil && !IsNil(o.LocationAccessRestrictions) {
		return true
	}

	return false
}

// SetLocationAccessRestrictions gets a reference to the given DataAccessProfileLocationAccessRestrictions and assigns it to the LocationAccessRestrictions field.
func (o *DataAccessProfile) SetLocationAccessRestrictions(v DataAccessProfileLocationAccessRestrictions) {
	o.LocationAccessRestrictions = &v
}

func (o DataAccessProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataAccessProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dataAccessProfileId"] = o.DataAccessProfileId
	toSerialize["targetEventConsumerTypes"] = o.TargetEventConsumerTypes
	toSerialize["parameters"] = o.Parameters
	if !IsNil(o.TimeAccessRestrictions) {
		toSerialize["timeAccessRestrictions"] = o.TimeAccessRestrictions
	}
	if !IsNil(o.UserAccessRestrictions) {
		toSerialize["userAccessRestrictions"] = o.UserAccessRestrictions
	}
	if !IsNil(o.LocationAccessRestrictions) {
		toSerialize["locationAccessRestrictions"] = o.LocationAccessRestrictions
	}
	return toSerialize, nil
}

type NullableDataAccessProfile struct {
	value *DataAccessProfile
	isSet bool
}

func (v NullableDataAccessProfile) Get() *DataAccessProfile {
	return v.value
}

func (v *NullableDataAccessProfile) Set(val *DataAccessProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAccessProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAccessProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAccessProfile(val *DataAccessProfile) *NullableDataAccessProfile {
	return &NullableDataAccessProfile{value: val, isSet: true}
}

func (v NullableDataAccessProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAccessProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


