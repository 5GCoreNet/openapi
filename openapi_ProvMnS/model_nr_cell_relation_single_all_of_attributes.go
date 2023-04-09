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

// checks if the NRCellRelationSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NRCellRelationSingleAllOfAttributes{}

// NRCellRelationSingleAllOfAttributes struct for NRCellRelationSingleAllOfAttributes
type NRCellRelationSingleAllOfAttributes struct {
	NRTCI                *int32                `json:"nRTCI,omitempty"`
	CellIndividualOffset *CellIndividualOffset `json:"cellIndividualOffset,omitempty"`
	AdjacentNRCellRef    *string               `json:"adjacentNRCellRef,omitempty"`
	NRFreqRelationRef    *string               `json:"nRFreqRelationRef,omitempty"`
	IsRemoveAllowed      *bool                 `json:"isRemoveAllowed,omitempty"`
	IsHOAllowed          *bool                 `json:"isHOAllowed,omitempty"`
	IsESCoveredBy        *IsESCoveredBy        `json:"isESCoveredBy,omitempty"`
	IsENDCAllowed        *bool                 `json:"isENDCAllowed,omitempty"`
	IsMLBAllowed         *bool                 `json:"isMLBAllowed,omitempty"`
}

// NewNRCellRelationSingleAllOfAttributes instantiates a new NRCellRelationSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNRCellRelationSingleAllOfAttributes() *NRCellRelationSingleAllOfAttributes {
	this := NRCellRelationSingleAllOfAttributes{}
	return &this
}

// NewNRCellRelationSingleAllOfAttributesWithDefaults instantiates a new NRCellRelationSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNRCellRelationSingleAllOfAttributesWithDefaults() *NRCellRelationSingleAllOfAttributes {
	this := NRCellRelationSingleAllOfAttributes{}
	return &this
}

// GetNRTCI returns the NRTCI field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetNRTCI() int32 {
	if o == nil || IsNil(o.NRTCI) {
		var ret int32
		return ret
	}
	return *o.NRTCI
}

// GetNRTCIOk returns a tuple with the NRTCI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetNRTCIOk() (*int32, bool) {
	if o == nil || IsNil(o.NRTCI) {
		return nil, false
	}
	return o.NRTCI, true
}

// HasNRTCI returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasNRTCI() bool {
	if o != nil && !IsNil(o.NRTCI) {
		return true
	}

	return false
}

// SetNRTCI gets a reference to the given int32 and assigns it to the NRTCI field.
func (o *NRCellRelationSingleAllOfAttributes) SetNRTCI(v int32) {
	o.NRTCI = &v
}

// GetCellIndividualOffset returns the CellIndividualOffset field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetCellIndividualOffset() CellIndividualOffset {
	if o == nil || IsNil(o.CellIndividualOffset) {
		var ret CellIndividualOffset
		return ret
	}
	return *o.CellIndividualOffset
}

// GetCellIndividualOffsetOk returns a tuple with the CellIndividualOffset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetCellIndividualOffsetOk() (*CellIndividualOffset, bool) {
	if o == nil || IsNil(o.CellIndividualOffset) {
		return nil, false
	}
	return o.CellIndividualOffset, true
}

// HasCellIndividualOffset returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasCellIndividualOffset() bool {
	if o != nil && !IsNil(o.CellIndividualOffset) {
		return true
	}

	return false
}

// SetCellIndividualOffset gets a reference to the given CellIndividualOffset and assigns it to the CellIndividualOffset field.
func (o *NRCellRelationSingleAllOfAttributes) SetCellIndividualOffset(v CellIndividualOffset) {
	o.CellIndividualOffset = &v
}

// GetAdjacentNRCellRef returns the AdjacentNRCellRef field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetAdjacentNRCellRef() string {
	if o == nil || IsNil(o.AdjacentNRCellRef) {
		var ret string
		return ret
	}
	return *o.AdjacentNRCellRef
}

// GetAdjacentNRCellRefOk returns a tuple with the AdjacentNRCellRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetAdjacentNRCellRefOk() (*string, bool) {
	if o == nil || IsNil(o.AdjacentNRCellRef) {
		return nil, false
	}
	return o.AdjacentNRCellRef, true
}

// HasAdjacentNRCellRef returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasAdjacentNRCellRef() bool {
	if o != nil && !IsNil(o.AdjacentNRCellRef) {
		return true
	}

	return false
}

// SetAdjacentNRCellRef gets a reference to the given string and assigns it to the AdjacentNRCellRef field.
func (o *NRCellRelationSingleAllOfAttributes) SetAdjacentNRCellRef(v string) {
	o.AdjacentNRCellRef = &v
}

// GetNRFreqRelationRef returns the NRFreqRelationRef field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetNRFreqRelationRef() string {
	if o == nil || IsNil(o.NRFreqRelationRef) {
		var ret string
		return ret
	}
	return *o.NRFreqRelationRef
}

// GetNRFreqRelationRefOk returns a tuple with the NRFreqRelationRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetNRFreqRelationRefOk() (*string, bool) {
	if o == nil || IsNil(o.NRFreqRelationRef) {
		return nil, false
	}
	return o.NRFreqRelationRef, true
}

// HasNRFreqRelationRef returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasNRFreqRelationRef() bool {
	if o != nil && !IsNil(o.NRFreqRelationRef) {
		return true
	}

	return false
}

// SetNRFreqRelationRef gets a reference to the given string and assigns it to the NRFreqRelationRef field.
func (o *NRCellRelationSingleAllOfAttributes) SetNRFreqRelationRef(v string) {
	o.NRFreqRelationRef = &v
}

// GetIsRemoveAllowed returns the IsRemoveAllowed field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetIsRemoveAllowed() bool {
	if o == nil || IsNil(o.IsRemoveAllowed) {
		var ret bool
		return ret
	}
	return *o.IsRemoveAllowed
}

// GetIsRemoveAllowedOk returns a tuple with the IsRemoveAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetIsRemoveAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRemoveAllowed) {
		return nil, false
	}
	return o.IsRemoveAllowed, true
}

// HasIsRemoveAllowed returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasIsRemoveAllowed() bool {
	if o != nil && !IsNil(o.IsRemoveAllowed) {
		return true
	}

	return false
}

// SetIsRemoveAllowed gets a reference to the given bool and assigns it to the IsRemoveAllowed field.
func (o *NRCellRelationSingleAllOfAttributes) SetIsRemoveAllowed(v bool) {
	o.IsRemoveAllowed = &v
}

// GetIsHOAllowed returns the IsHOAllowed field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetIsHOAllowed() bool {
	if o == nil || IsNil(o.IsHOAllowed) {
		var ret bool
		return ret
	}
	return *o.IsHOAllowed
}

// GetIsHOAllowedOk returns a tuple with the IsHOAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetIsHOAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHOAllowed) {
		return nil, false
	}
	return o.IsHOAllowed, true
}

// HasIsHOAllowed returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasIsHOAllowed() bool {
	if o != nil && !IsNil(o.IsHOAllowed) {
		return true
	}

	return false
}

// SetIsHOAllowed gets a reference to the given bool and assigns it to the IsHOAllowed field.
func (o *NRCellRelationSingleAllOfAttributes) SetIsHOAllowed(v bool) {
	o.IsHOAllowed = &v
}

// GetIsESCoveredBy returns the IsESCoveredBy field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetIsESCoveredBy() IsESCoveredBy {
	if o == nil || IsNil(o.IsESCoveredBy) {
		var ret IsESCoveredBy
		return ret
	}
	return *o.IsESCoveredBy
}

// GetIsESCoveredByOk returns a tuple with the IsESCoveredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetIsESCoveredByOk() (*IsESCoveredBy, bool) {
	if o == nil || IsNil(o.IsESCoveredBy) {
		return nil, false
	}
	return o.IsESCoveredBy, true
}

// HasIsESCoveredBy returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasIsESCoveredBy() bool {
	if o != nil && !IsNil(o.IsESCoveredBy) {
		return true
	}

	return false
}

// SetIsESCoveredBy gets a reference to the given IsESCoveredBy and assigns it to the IsESCoveredBy field.
func (o *NRCellRelationSingleAllOfAttributes) SetIsESCoveredBy(v IsESCoveredBy) {
	o.IsESCoveredBy = &v
}

// GetIsENDCAllowed returns the IsENDCAllowed field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetIsENDCAllowed() bool {
	if o == nil || IsNil(o.IsENDCAllowed) {
		var ret bool
		return ret
	}
	return *o.IsENDCAllowed
}

// GetIsENDCAllowedOk returns a tuple with the IsENDCAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetIsENDCAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsENDCAllowed) {
		return nil, false
	}
	return o.IsENDCAllowed, true
}

// HasIsENDCAllowed returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasIsENDCAllowed() bool {
	if o != nil && !IsNil(o.IsENDCAllowed) {
		return true
	}

	return false
}

// SetIsENDCAllowed gets a reference to the given bool and assigns it to the IsENDCAllowed field.
func (o *NRCellRelationSingleAllOfAttributes) SetIsENDCAllowed(v bool) {
	o.IsENDCAllowed = &v
}

// GetIsMLBAllowed returns the IsMLBAllowed field value if set, zero value otherwise.
func (o *NRCellRelationSingleAllOfAttributes) GetIsMLBAllowed() bool {
	if o == nil || IsNil(o.IsMLBAllowed) {
		var ret bool
		return ret
	}
	return *o.IsMLBAllowed
}

// GetIsMLBAllowedOk returns a tuple with the IsMLBAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NRCellRelationSingleAllOfAttributes) GetIsMLBAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMLBAllowed) {
		return nil, false
	}
	return o.IsMLBAllowed, true
}

// HasIsMLBAllowed returns a boolean if a field has been set.
func (o *NRCellRelationSingleAllOfAttributes) HasIsMLBAllowed() bool {
	if o != nil && !IsNil(o.IsMLBAllowed) {
		return true
	}

	return false
}

// SetIsMLBAllowed gets a reference to the given bool and assigns it to the IsMLBAllowed field.
func (o *NRCellRelationSingleAllOfAttributes) SetIsMLBAllowed(v bool) {
	o.IsMLBAllowed = &v
}

func (o NRCellRelationSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NRCellRelationSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NRTCI) {
		toSerialize["nRTCI"] = o.NRTCI
	}
	if !IsNil(o.CellIndividualOffset) {
		toSerialize["cellIndividualOffset"] = o.CellIndividualOffset
	}
	if !IsNil(o.AdjacentNRCellRef) {
		toSerialize["adjacentNRCellRef"] = o.AdjacentNRCellRef
	}
	if !IsNil(o.NRFreqRelationRef) {
		toSerialize["nRFreqRelationRef"] = o.NRFreqRelationRef
	}
	if !IsNil(o.IsRemoveAllowed) {
		toSerialize["isRemoveAllowed"] = o.IsRemoveAllowed
	}
	if !IsNil(o.IsHOAllowed) {
		toSerialize["isHOAllowed"] = o.IsHOAllowed
	}
	if !IsNil(o.IsESCoveredBy) {
		toSerialize["isESCoveredBy"] = o.IsESCoveredBy
	}
	if !IsNil(o.IsENDCAllowed) {
		toSerialize["isENDCAllowed"] = o.IsENDCAllowed
	}
	if !IsNil(o.IsMLBAllowed) {
		toSerialize["isMLBAllowed"] = o.IsMLBAllowed
	}
	return toSerialize, nil
}

type NullableNRCellRelationSingleAllOfAttributes struct {
	value *NRCellRelationSingleAllOfAttributes
	isSet bool
}

func (v NullableNRCellRelationSingleAllOfAttributes) Get() *NRCellRelationSingleAllOfAttributes {
	return v.value
}

func (v *NullableNRCellRelationSingleAllOfAttributes) Set(val *NRCellRelationSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNRCellRelationSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNRCellRelationSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNRCellRelationSingleAllOfAttributes(val *NRCellRelationSingleAllOfAttributes) *NullableNRCellRelationSingleAllOfAttributes {
	return &NullableNRCellRelationSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableNRCellRelationSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNRCellRelationSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
