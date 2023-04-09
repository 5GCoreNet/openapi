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

// checks if the ApplicationSpecificRecordAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationSpecificRecordAllOf{}

// ApplicationSpecificRecordAllOf struct for ApplicationSpecificRecordAllOf
type ApplicationSpecificRecordAllOf struct {
	// String providing an URI formatted according to RFC 3986.
	RecordType      string      `json:"recordType"`
	RecordContainer interface{} `json:"recordContainer"`
}

// NewApplicationSpecificRecordAllOf instantiates a new ApplicationSpecificRecordAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSpecificRecordAllOf(recordType string, recordContainer interface{}) *ApplicationSpecificRecordAllOf {
	this := ApplicationSpecificRecordAllOf{}
	this.RecordType = recordType
	this.RecordContainer = recordContainer
	return &this
}

// NewApplicationSpecificRecordAllOfWithDefaults instantiates a new ApplicationSpecificRecordAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSpecificRecordAllOfWithDefaults() *ApplicationSpecificRecordAllOf {
	this := ApplicationSpecificRecordAllOf{}
	return &this
}

// GetRecordType returns the RecordType field value
func (o *ApplicationSpecificRecordAllOf) GetRecordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value
// and a boolean to check if the value has been set.
func (o *ApplicationSpecificRecordAllOf) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordType, true
}

// SetRecordType sets field value
func (o *ApplicationSpecificRecordAllOf) SetRecordType(v string) {
	o.RecordType = v
}

// GetRecordContainer returns the RecordContainer field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *ApplicationSpecificRecordAllOf) GetRecordContainer() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.RecordContainer
}

// GetRecordContainerOk returns a tuple with the RecordContainer field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApplicationSpecificRecordAllOf) GetRecordContainerOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RecordContainer) {
		return nil, false
	}
	return &o.RecordContainer, true
}

// SetRecordContainer sets field value
func (o *ApplicationSpecificRecordAllOf) SetRecordContainer(v interface{}) {
	o.RecordContainer = v
}

func (o ApplicationSpecificRecordAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationSpecificRecordAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["recordType"] = o.RecordType
	if o.RecordContainer != nil {
		toSerialize["recordContainer"] = o.RecordContainer
	}
	return toSerialize, nil
}

type NullableApplicationSpecificRecordAllOf struct {
	value *ApplicationSpecificRecordAllOf
	isSet bool
}

func (v NullableApplicationSpecificRecordAllOf) Get() *ApplicationSpecificRecordAllOf {
	return v.value
}

func (v *NullableApplicationSpecificRecordAllOf) Set(val *ApplicationSpecificRecordAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSpecificRecordAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSpecificRecordAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSpecificRecordAllOf(val *ApplicationSpecificRecordAllOf) *NullableApplicationSpecificRecordAllOf {
	return &NullableApplicationSpecificRecordAllOf{value: val, isSet: true}
}

func (v NullableApplicationSpecificRecordAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSpecificRecordAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
