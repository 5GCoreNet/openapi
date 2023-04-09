/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
)

// checks if the LocationRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationRecord{}

// LocationRecord struct for LocationRecord
type LocationRecord struct {
	BaseRecord
	Location LocationData `json:"location"`
}

// NewLocationRecord instantiates a new LocationRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationRecord(location LocationData, timestamp time.Time) *LocationRecord {
	this := LocationRecord{}
	this.Timestamp = timestamp
	this.Location = location
	return &this
}

// NewLocationRecordWithDefaults instantiates a new LocationRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationRecordWithDefaults() *LocationRecord {
	this := LocationRecord{}
	return &this
}

// GetLocation returns the Location field value
func (o *LocationRecord) GetLocation() LocationData {
	if o == nil {
		var ret LocationData
		return ret
	}

	return o.Location
}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
func (o *LocationRecord) GetLocationOk() (*LocationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Location, true
}

// SetLocation sets field value
func (o *LocationRecord) SetLocation(v LocationData) {
	o.Location = v
}

func (o LocationRecord) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseRecord, errBaseRecord := json.Marshal(o.BaseRecord)
	if errBaseRecord != nil {
		return map[string]interface{}{}, errBaseRecord
	}
	errBaseRecord = json.Unmarshal([]byte(serializedBaseRecord), &toSerialize)
	if errBaseRecord != nil {
		return map[string]interface{}{}, errBaseRecord
	}
	toSerialize["location"] = o.Location
	return toSerialize, nil
}

type NullableLocationRecord struct {
	value *LocationRecord
	isSet bool
}

func (v NullableLocationRecord) Get() *LocationRecord {
	return v.value
}

func (v *NullableLocationRecord) Set(val *LocationRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRecord(val *LocationRecord) *NullableLocationRecord {
	return &NullableLocationRecord{value: val, isSet: true}
}

func (v NullableLocationRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
