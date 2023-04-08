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

// checks if the MultipleUnitUsage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipleUnitUsage{}

// MultipleUnitUsage struct for MultipleUnitUsage
type MultipleUnitUsage struct {
	// Integer where the allowed values correspond to the value range of an unsigned 32-bit integer. 
	RatingGroup int32 `json:"ratingGroup"`
	RequestedUnit *RequestedUnit `json:"requestedUnit,omitempty"`
	UsedUnitContainer []UsedUnitContainer `json:"usedUnitContainer,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	UPFID *string `json:"uPFID,omitempty"`
	MultihomedPDUAddress *PDUAddress `json:"multihomedPDUAddress,omitempty"`
}

// NewMultipleUnitUsage instantiates a new MultipleUnitUsage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipleUnitUsage(ratingGroup int32) *MultipleUnitUsage {
	this := MultipleUnitUsage{}
	this.RatingGroup = ratingGroup
	return &this
}

// NewMultipleUnitUsageWithDefaults instantiates a new MultipleUnitUsage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipleUnitUsageWithDefaults() *MultipleUnitUsage {
	this := MultipleUnitUsage{}
	return &this
}

// GetRatingGroup returns the RatingGroup field value
func (o *MultipleUnitUsage) GetRatingGroup() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RatingGroup
}

// GetRatingGroupOk returns a tuple with the RatingGroup field value
// and a boolean to check if the value has been set.
func (o *MultipleUnitUsage) GetRatingGroupOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RatingGroup, true
}

// SetRatingGroup sets field value
func (o *MultipleUnitUsage) SetRatingGroup(v int32) {
	o.RatingGroup = v
}

// GetRequestedUnit returns the RequestedUnit field value if set, zero value otherwise.
func (o *MultipleUnitUsage) GetRequestedUnit() RequestedUnit {
	if o == nil || isNil(o.RequestedUnit) {
		var ret RequestedUnit
		return ret
	}
	return *o.RequestedUnit
}

// GetRequestedUnitOk returns a tuple with the RequestedUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitUsage) GetRequestedUnitOk() (*RequestedUnit, bool) {
	if o == nil || isNil(o.RequestedUnit) {
		return nil, false
	}
	return o.RequestedUnit, true
}

// HasRequestedUnit returns a boolean if a field has been set.
func (o *MultipleUnitUsage) HasRequestedUnit() bool {
	if o != nil && !isNil(o.RequestedUnit) {
		return true
	}

	return false
}

// SetRequestedUnit gets a reference to the given RequestedUnit and assigns it to the RequestedUnit field.
func (o *MultipleUnitUsage) SetRequestedUnit(v RequestedUnit) {
	o.RequestedUnit = &v
}

// GetUsedUnitContainer returns the UsedUnitContainer field value if set, zero value otherwise.
func (o *MultipleUnitUsage) GetUsedUnitContainer() []UsedUnitContainer {
	if o == nil || isNil(o.UsedUnitContainer) {
		var ret []UsedUnitContainer
		return ret
	}
	return o.UsedUnitContainer
}

// GetUsedUnitContainerOk returns a tuple with the UsedUnitContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitUsage) GetUsedUnitContainerOk() ([]UsedUnitContainer, bool) {
	if o == nil || isNil(o.UsedUnitContainer) {
		return nil, false
	}
	return o.UsedUnitContainer, true
}

// HasUsedUnitContainer returns a boolean if a field has been set.
func (o *MultipleUnitUsage) HasUsedUnitContainer() bool {
	if o != nil && !isNil(o.UsedUnitContainer) {
		return true
	}

	return false
}

// SetUsedUnitContainer gets a reference to the given []UsedUnitContainer and assigns it to the UsedUnitContainer field.
func (o *MultipleUnitUsage) SetUsedUnitContainer(v []UsedUnitContainer) {
	o.UsedUnitContainer = v
}

// GetUPFID returns the UPFID field value if set, zero value otherwise.
func (o *MultipleUnitUsage) GetUPFID() string {
	if o == nil || isNil(o.UPFID) {
		var ret string
		return ret
	}
	return *o.UPFID
}

// GetUPFIDOk returns a tuple with the UPFID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitUsage) GetUPFIDOk() (*string, bool) {
	if o == nil || isNil(o.UPFID) {
		return nil, false
	}
	return o.UPFID, true
}

// HasUPFID returns a boolean if a field has been set.
func (o *MultipleUnitUsage) HasUPFID() bool {
	if o != nil && !isNil(o.UPFID) {
		return true
	}

	return false
}

// SetUPFID gets a reference to the given string and assigns it to the UPFID field.
func (o *MultipleUnitUsage) SetUPFID(v string) {
	o.UPFID = &v
}

// GetMultihomedPDUAddress returns the MultihomedPDUAddress field value if set, zero value otherwise.
func (o *MultipleUnitUsage) GetMultihomedPDUAddress() PDUAddress {
	if o == nil || isNil(o.MultihomedPDUAddress) {
		var ret PDUAddress
		return ret
	}
	return *o.MultihomedPDUAddress
}

// GetMultihomedPDUAddressOk returns a tuple with the MultihomedPDUAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipleUnitUsage) GetMultihomedPDUAddressOk() (*PDUAddress, bool) {
	if o == nil || isNil(o.MultihomedPDUAddress) {
		return nil, false
	}
	return o.MultihomedPDUAddress, true
}

// HasMultihomedPDUAddress returns a boolean if a field has been set.
func (o *MultipleUnitUsage) HasMultihomedPDUAddress() bool {
	if o != nil && !isNil(o.MultihomedPDUAddress) {
		return true
	}

	return false
}

// SetMultihomedPDUAddress gets a reference to the given PDUAddress and assigns it to the MultihomedPDUAddress field.
func (o *MultipleUnitUsage) SetMultihomedPDUAddress(v PDUAddress) {
	o.MultihomedPDUAddress = &v
}

func (o MultipleUnitUsage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipleUnitUsage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ratingGroup"] = o.RatingGroup
	if !isNil(o.RequestedUnit) {
		toSerialize["requestedUnit"] = o.RequestedUnit
	}
	if !isNil(o.UsedUnitContainer) {
		toSerialize["usedUnitContainer"] = o.UsedUnitContainer
	}
	if !isNil(o.UPFID) {
		toSerialize["uPFID"] = o.UPFID
	}
	if !isNil(o.MultihomedPDUAddress) {
		toSerialize["multihomedPDUAddress"] = o.MultihomedPDUAddress
	}
	return toSerialize, nil
}

type NullableMultipleUnitUsage struct {
	value *MultipleUnitUsage
	isSet bool
}

func (v NullableMultipleUnitUsage) Get() *MultipleUnitUsage {
	return v.value
}

func (v *NullableMultipleUnitUsage) Set(val *MultipleUnitUsage) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipleUnitUsage) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipleUnitUsage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipleUnitUsage(val *MultipleUnitUsage) *NullableMultipleUnitUsage {
	return &NullableMultipleUnitUsage{value: val, isSet: true}
}

func (v NullableMultipleUnitUsage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipleUnitUsage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


