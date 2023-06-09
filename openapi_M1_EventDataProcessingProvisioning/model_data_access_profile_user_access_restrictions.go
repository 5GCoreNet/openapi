/*
M1_EventDataProcessingProvisioning

5GMS AF M1 Event Data Processing Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_EventDataProcessingProvisioning

import (
	"encoding/json"
)

// checks if the DataAccessProfileUserAccessRestrictions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataAccessProfileUserAccessRestrictions{}

// DataAccessProfileUserAccessRestrictions struct for DataAccessProfileUserAccessRestrictions
type DataAccessProfileUserAccessRestrictions struct {
	GroupIds             []string                                              `json:"groupIds"`
	UserIds              []DataAccessProfileUserAccessRestrictionsUserIdsInner `json:"userIds"`
	AggregationFunctions []DataAggregationFunctionType                         `json:"aggregationFunctions"`
}

// NewDataAccessProfileUserAccessRestrictions instantiates a new DataAccessProfileUserAccessRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataAccessProfileUserAccessRestrictions(groupIds []string, userIds []DataAccessProfileUserAccessRestrictionsUserIdsInner, aggregationFunctions []DataAggregationFunctionType) *DataAccessProfileUserAccessRestrictions {
	this := DataAccessProfileUserAccessRestrictions{}
	this.GroupIds = groupIds
	this.UserIds = userIds
	this.AggregationFunctions = aggregationFunctions
	return &this
}

// NewDataAccessProfileUserAccessRestrictionsWithDefaults instantiates a new DataAccessProfileUserAccessRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataAccessProfileUserAccessRestrictionsWithDefaults() *DataAccessProfileUserAccessRestrictions {
	this := DataAccessProfileUserAccessRestrictions{}
	return &this
}

// GetGroupIds returns the GroupIds field value
func (o *DataAccessProfileUserAccessRestrictions) GetGroupIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value
// and a boolean to check if the value has been set.
func (o *DataAccessProfileUserAccessRestrictions) GetGroupIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupIds, true
}

// SetGroupIds sets field value
func (o *DataAccessProfileUserAccessRestrictions) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetUserIds returns the UserIds field value
func (o *DataAccessProfileUserAccessRestrictions) GetUserIds() []DataAccessProfileUserAccessRestrictionsUserIdsInner {
	if o == nil {
		var ret []DataAccessProfileUserAccessRestrictionsUserIdsInner
		return ret
	}

	return o.UserIds
}

// GetUserIdsOk returns a tuple with the UserIds field value
// and a boolean to check if the value has been set.
func (o *DataAccessProfileUserAccessRestrictions) GetUserIdsOk() ([]DataAccessProfileUserAccessRestrictionsUserIdsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserIds, true
}

// SetUserIds sets field value
func (o *DataAccessProfileUserAccessRestrictions) SetUserIds(v []DataAccessProfileUserAccessRestrictionsUserIdsInner) {
	o.UserIds = v
}

// GetAggregationFunctions returns the AggregationFunctions field value
func (o *DataAccessProfileUserAccessRestrictions) GetAggregationFunctions() []DataAggregationFunctionType {
	if o == nil {
		var ret []DataAggregationFunctionType
		return ret
	}

	return o.AggregationFunctions
}

// GetAggregationFunctionsOk returns a tuple with the AggregationFunctions field value
// and a boolean to check if the value has been set.
func (o *DataAccessProfileUserAccessRestrictions) GetAggregationFunctionsOk() ([]DataAggregationFunctionType, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregationFunctions, true
}

// SetAggregationFunctions sets field value
func (o *DataAccessProfileUserAccessRestrictions) SetAggregationFunctions(v []DataAggregationFunctionType) {
	o.AggregationFunctions = v
}

func (o DataAccessProfileUserAccessRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataAccessProfileUserAccessRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["groupIds"] = o.GroupIds
	toSerialize["userIds"] = o.UserIds
	toSerialize["aggregationFunctions"] = o.AggregationFunctions
	return toSerialize, nil
}

type NullableDataAccessProfileUserAccessRestrictions struct {
	value *DataAccessProfileUserAccessRestrictions
	isSet bool
}

func (v NullableDataAccessProfileUserAccessRestrictions) Get() *DataAccessProfileUserAccessRestrictions {
	return v.value
}

func (v *NullableDataAccessProfileUserAccessRestrictions) Set(val *DataAccessProfileUserAccessRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAccessProfileUserAccessRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAccessProfileUserAccessRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAccessProfileUserAccessRestrictions(val *DataAccessProfileUserAccessRestrictions) *NullableDataAccessProfileUserAccessRestrictions {
	return &NullableDataAccessProfileUserAccessRestrictions{value: val, isSet: true}
}

func (v NullableDataAccessProfileUserAccessRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAccessProfileUserAccessRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
