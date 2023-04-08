/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the PositioningRANSubnet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PositioningRANSubnet{}

// PositioningRANSubnet struct for PositioningRANSubnet
type PositioningRANSubnet struct {
	Availability []string `json:"availability,omitempty"`
	Predictionfrequency *Predictionfrequency `json:"predictionfrequency,omitempty"`
	Accuracy *float32 `json:"accuracy,omitempty"`
}

// NewPositioningRANSubnet instantiates a new PositioningRANSubnet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPositioningRANSubnet() *PositioningRANSubnet {
	this := PositioningRANSubnet{}
	return &this
}

// NewPositioningRANSubnetWithDefaults instantiates a new PositioningRANSubnet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPositioningRANSubnetWithDefaults() *PositioningRANSubnet {
	this := PositioningRANSubnet{}
	return &this
}

// GetAvailability returns the Availability field value if set, zero value otherwise.
func (o *PositioningRANSubnet) GetAvailability() []string {
	if o == nil || isNil(o.Availability) {
		var ret []string
		return ret
	}
	return o.Availability
}

// GetAvailabilityOk returns a tuple with the Availability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositioningRANSubnet) GetAvailabilityOk() ([]string, bool) {
	if o == nil || isNil(o.Availability) {
		return nil, false
	}
	return o.Availability, true
}

// HasAvailability returns a boolean if a field has been set.
func (o *PositioningRANSubnet) HasAvailability() bool {
	if o != nil && !isNil(o.Availability) {
		return true
	}

	return false
}

// SetAvailability gets a reference to the given []string and assigns it to the Availability field.
func (o *PositioningRANSubnet) SetAvailability(v []string) {
	o.Availability = v
}

// GetPredictionfrequency returns the Predictionfrequency field value if set, zero value otherwise.
func (o *PositioningRANSubnet) GetPredictionfrequency() Predictionfrequency {
	if o == nil || isNil(o.Predictionfrequency) {
		var ret Predictionfrequency
		return ret
	}
	return *o.Predictionfrequency
}

// GetPredictionfrequencyOk returns a tuple with the Predictionfrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositioningRANSubnet) GetPredictionfrequencyOk() (*Predictionfrequency, bool) {
	if o == nil || isNil(o.Predictionfrequency) {
		return nil, false
	}
	return o.Predictionfrequency, true
}

// HasPredictionfrequency returns a boolean if a field has been set.
func (o *PositioningRANSubnet) HasPredictionfrequency() bool {
	if o != nil && !isNil(o.Predictionfrequency) {
		return true
	}

	return false
}

// SetPredictionfrequency gets a reference to the given Predictionfrequency and assigns it to the Predictionfrequency field.
func (o *PositioningRANSubnet) SetPredictionfrequency(v Predictionfrequency) {
	o.Predictionfrequency = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *PositioningRANSubnet) GetAccuracy() float32 {
	if o == nil || isNil(o.Accuracy) {
		var ret float32
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PositioningRANSubnet) GetAccuracyOk() (*float32, bool) {
	if o == nil || isNil(o.Accuracy) {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *PositioningRANSubnet) HasAccuracy() bool {
	if o != nil && !isNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given float32 and assigns it to the Accuracy field.
func (o *PositioningRANSubnet) SetAccuracy(v float32) {
	o.Accuracy = &v
}

func (o PositioningRANSubnet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PositioningRANSubnet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Availability) {
		toSerialize["availability"] = o.Availability
	}
	if !isNil(o.Predictionfrequency) {
		toSerialize["predictionfrequency"] = o.Predictionfrequency
	}
	if !isNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	return toSerialize, nil
}

type NullablePositioningRANSubnet struct {
	value *PositioningRANSubnet
	isSet bool
}

func (v NullablePositioningRANSubnet) Get() *PositioningRANSubnet {
	return v.value
}

func (v *NullablePositioningRANSubnet) Set(val *PositioningRANSubnet) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningRANSubnet) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningRANSubnet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningRANSubnet(val *PositioningRANSubnet) *NullablePositioningRANSubnet {
	return &NullablePositioningRANSubnet{value: val, isSet: true}
}

func (v NullablePositioningRANSubnet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningRANSubnet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


