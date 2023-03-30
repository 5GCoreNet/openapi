/*
5GMS Common Data Types

5GMS Common Data Types © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
)

// checks if the M1QoSSpecification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &M1QoSSpecification{}

// M1QoSSpecification struct for M1QoSSpecification
type M1QoSSpecification struct {
	QosReference *string `json:"qosReference,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxBtrUl *string `json:"maxBtrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxBtrDl *string `json:"maxBtrDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxAuthBtrUl *string `json:"maxAuthBtrUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MaxAuthBtrDl *string `json:"maxAuthBtrDl,omitempty"`
	DefPacketLossRateDl *int32 `json:"defPacketLossRateDl,omitempty"`
	DefPacketLossRateUl *int32 `json:"defPacketLossRateUl,omitempty"`
}

// NewM1QoSSpecification instantiates a new M1QoSSpecification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewM1QoSSpecification() *M1QoSSpecification {
	this := M1QoSSpecification{}
	return &this
}

// NewM1QoSSpecificationWithDefaults instantiates a new M1QoSSpecification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewM1QoSSpecificationWithDefaults() *M1QoSSpecification {
	this := M1QoSSpecification{}
	return &this
}

// GetQosReference returns the QosReference field value if set, zero value otherwise.
func (o *M1QoSSpecification) GetQosReference() string {
	if o == nil || IsNil(o.QosReference) {
		var ret string
		return ret
	}
	return *o.QosReference
}

// GetQosReferenceOk returns a tuple with the QosReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M1QoSSpecification) GetQosReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.QosReference) {
		return nil, false
	}
	return o.QosReference, true
}

// HasQosReference returns a boolean if a field has been set.
func (o *M1QoSSpecification) HasQosReference() bool {
	if o != nil && !IsNil(o.QosReference) {
		return true
	}

	return false
}

// SetQosReference gets a reference to the given string and assigns it to the QosReference field.
func (o *M1QoSSpecification) SetQosReference(v string) {
	o.QosReference = &v
}

// GetMaxBtrUl returns the MaxBtrUl field value if set, zero value otherwise.
func (o *M1QoSSpecification) GetMaxBtrUl() string {
	if o == nil || IsNil(o.MaxBtrUl) {
		var ret string
		return ret
	}
	return *o.MaxBtrUl
}

// GetMaxBtrUlOk returns a tuple with the MaxBtrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M1QoSSpecification) GetMaxBtrUlOk() (*string, bool) {
	if o == nil || IsNil(o.MaxBtrUl) {
		return nil, false
	}
	return o.MaxBtrUl, true
}

// HasMaxBtrUl returns a boolean if a field has been set.
func (o *M1QoSSpecification) HasMaxBtrUl() bool {
	if o != nil && !IsNil(o.MaxBtrUl) {
		return true
	}

	return false
}

// SetMaxBtrUl gets a reference to the given string and assigns it to the MaxBtrUl field.
func (o *M1QoSSpecification) SetMaxBtrUl(v string) {
	o.MaxBtrUl = &v
}

// GetMaxBtrDl returns the MaxBtrDl field value if set, zero value otherwise.
func (o *M1QoSSpecification) GetMaxBtrDl() string {
	if o == nil || IsNil(o.MaxBtrDl) {
		var ret string
		return ret
	}
	return *o.MaxBtrDl
}

// GetMaxBtrDlOk returns a tuple with the MaxBtrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M1QoSSpecification) GetMaxBtrDlOk() (*string, bool) {
	if o == nil || IsNil(o.MaxBtrDl) {
		return nil, false
	}
	return o.MaxBtrDl, true
}

// HasMaxBtrDl returns a boolean if a field has been set.
func (o *M1QoSSpecification) HasMaxBtrDl() bool {
	if o != nil && !IsNil(o.MaxBtrDl) {
		return true
	}

	return false
}

// SetMaxBtrDl gets a reference to the given string and assigns it to the MaxBtrDl field.
func (o *M1QoSSpecification) SetMaxBtrDl(v string) {
	o.MaxBtrDl = &v
}

// GetMaxAuthBtrUl returns the MaxAuthBtrUl field value if set, zero value otherwise.
func (o *M1QoSSpecification) GetMaxAuthBtrUl() string {
	if o == nil || IsNil(o.MaxAuthBtrUl) {
		var ret string
		return ret
	}
	return *o.MaxAuthBtrUl
}

// GetMaxAuthBtrUlOk returns a tuple with the MaxAuthBtrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M1QoSSpecification) GetMaxAuthBtrUlOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAuthBtrUl) {
		return nil, false
	}
	return o.MaxAuthBtrUl, true
}

// HasMaxAuthBtrUl returns a boolean if a field has been set.
func (o *M1QoSSpecification) HasMaxAuthBtrUl() bool {
	if o != nil && !IsNil(o.MaxAuthBtrUl) {
		return true
	}

	return false
}

// SetMaxAuthBtrUl gets a reference to the given string and assigns it to the MaxAuthBtrUl field.
func (o *M1QoSSpecification) SetMaxAuthBtrUl(v string) {
	o.MaxAuthBtrUl = &v
}

// GetMaxAuthBtrDl returns the MaxAuthBtrDl field value if set, zero value otherwise.
func (o *M1QoSSpecification) GetMaxAuthBtrDl() string {
	if o == nil || IsNil(o.MaxAuthBtrDl) {
		var ret string
		return ret
	}
	return *o.MaxAuthBtrDl
}

// GetMaxAuthBtrDlOk returns a tuple with the MaxAuthBtrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M1QoSSpecification) GetMaxAuthBtrDlOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAuthBtrDl) {
		return nil, false
	}
	return o.MaxAuthBtrDl, true
}

// HasMaxAuthBtrDl returns a boolean if a field has been set.
func (o *M1QoSSpecification) HasMaxAuthBtrDl() bool {
	if o != nil && !IsNil(o.MaxAuthBtrDl) {
		return true
	}

	return false
}

// SetMaxAuthBtrDl gets a reference to the given string and assigns it to the MaxAuthBtrDl field.
func (o *M1QoSSpecification) SetMaxAuthBtrDl(v string) {
	o.MaxAuthBtrDl = &v
}

// GetDefPacketLossRateDl returns the DefPacketLossRateDl field value if set, zero value otherwise.
func (o *M1QoSSpecification) GetDefPacketLossRateDl() int32 {
	if o == nil || IsNil(o.DefPacketLossRateDl) {
		var ret int32
		return ret
	}
	return *o.DefPacketLossRateDl
}

// GetDefPacketLossRateDlOk returns a tuple with the DefPacketLossRateDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M1QoSSpecification) GetDefPacketLossRateDlOk() (*int32, bool) {
	if o == nil || IsNil(o.DefPacketLossRateDl) {
		return nil, false
	}
	return o.DefPacketLossRateDl, true
}

// HasDefPacketLossRateDl returns a boolean if a field has been set.
func (o *M1QoSSpecification) HasDefPacketLossRateDl() bool {
	if o != nil && !IsNil(o.DefPacketLossRateDl) {
		return true
	}

	return false
}

// SetDefPacketLossRateDl gets a reference to the given int32 and assigns it to the DefPacketLossRateDl field.
func (o *M1QoSSpecification) SetDefPacketLossRateDl(v int32) {
	o.DefPacketLossRateDl = &v
}

// GetDefPacketLossRateUl returns the DefPacketLossRateUl field value if set, zero value otherwise.
func (o *M1QoSSpecification) GetDefPacketLossRateUl() int32 {
	if o == nil || IsNil(o.DefPacketLossRateUl) {
		var ret int32
		return ret
	}
	return *o.DefPacketLossRateUl
}

// GetDefPacketLossRateUlOk returns a tuple with the DefPacketLossRateUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *M1QoSSpecification) GetDefPacketLossRateUlOk() (*int32, bool) {
	if o == nil || IsNil(o.DefPacketLossRateUl) {
		return nil, false
	}
	return o.DefPacketLossRateUl, true
}

// HasDefPacketLossRateUl returns a boolean if a field has been set.
func (o *M1QoSSpecification) HasDefPacketLossRateUl() bool {
	if o != nil && !IsNil(o.DefPacketLossRateUl) {
		return true
	}

	return false
}

// SetDefPacketLossRateUl gets a reference to the given int32 and assigns it to the DefPacketLossRateUl field.
func (o *M1QoSSpecification) SetDefPacketLossRateUl(v int32) {
	o.DefPacketLossRateUl = &v
}

func (o M1QoSSpecification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o M1QoSSpecification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.QosReference) {
		toSerialize["qosReference"] = o.QosReference
	}
	if !IsNil(o.MaxBtrUl) {
		toSerialize["maxBtrUl"] = o.MaxBtrUl
	}
	if !IsNil(o.MaxBtrDl) {
		toSerialize["maxBtrDl"] = o.MaxBtrDl
	}
	if !IsNil(o.MaxAuthBtrUl) {
		toSerialize["maxAuthBtrUl"] = o.MaxAuthBtrUl
	}
	if !IsNil(o.MaxAuthBtrDl) {
		toSerialize["maxAuthBtrDl"] = o.MaxAuthBtrDl
	}
	if !IsNil(o.DefPacketLossRateDl) {
		toSerialize["defPacketLossRateDl"] = o.DefPacketLossRateDl
	}
	if !IsNil(o.DefPacketLossRateUl) {
		toSerialize["defPacketLossRateUl"] = o.DefPacketLossRateUl
	}
	return toSerialize, nil
}

type NullableM1QoSSpecification struct {
	value *M1QoSSpecification
	isSet bool
}

func (v NullableM1QoSSpecification) Get() *M1QoSSpecification {
	return v.value
}

func (v *NullableM1QoSSpecification) Set(val *M1QoSSpecification) {
	v.value = val
	v.isSet = true
}

func (v NullableM1QoSSpecification) IsSet() bool {
	return v.isSet
}

func (v *NullableM1QoSSpecification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableM1QoSSpecification(val *M1QoSSpecification) *NullableM1QoSSpecification {
	return &NullableM1QoSSpecification{value: val, isSet: true}
}

func (v NullableM1QoSSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableM1QoSSpecification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

