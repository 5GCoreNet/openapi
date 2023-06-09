/*
GMDviaMBMSbyMB2

API for Group Message Delivery via MBMS by MB2   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GMDviaMBMSbyMB2

import (
	"encoding/json"
	"time"
)

// checks if the TMGIAllocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TMGIAllocation{}

// TMGIAllocation Represents an individual TMGI Allocation resource.
type TMGIAllocation struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExternalGroupId *string      `json:"externalGroupId,omitempty"`
	MbmsLocArea     *MbmsLocArea `json:"mbmsLocArea,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI with \"readOnly=true\" property.
	TmgiExpiration *time.Time `json:"tmgiExpiration,omitempty"`
}

// NewTMGIAllocation instantiates a new TMGIAllocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTMGIAllocation() *TMGIAllocation {
	this := TMGIAllocation{}
	return &this
}

// NewTMGIAllocationWithDefaults instantiates a new TMGIAllocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTMGIAllocationWithDefaults() *TMGIAllocation {
	this := TMGIAllocation{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *TMGIAllocation) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TMGIAllocation) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *TMGIAllocation) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *TMGIAllocation) SetSelf(v string) {
	o.Self = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *TMGIAllocation) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TMGIAllocation) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *TMGIAllocation) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *TMGIAllocation) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetExternalGroupId returns the ExternalGroupId field value if set, zero value otherwise.
func (o *TMGIAllocation) GetExternalGroupId() string {
	if o == nil || IsNil(o.ExternalGroupId) {
		var ret string
		return ret
	}
	return *o.ExternalGroupId
}

// GetExternalGroupIdOk returns a tuple with the ExternalGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TMGIAllocation) GetExternalGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalGroupId) {
		return nil, false
	}
	return o.ExternalGroupId, true
}

// HasExternalGroupId returns a boolean if a field has been set.
func (o *TMGIAllocation) HasExternalGroupId() bool {
	if o != nil && !IsNil(o.ExternalGroupId) {
		return true
	}

	return false
}

// SetExternalGroupId gets a reference to the given string and assigns it to the ExternalGroupId field.
func (o *TMGIAllocation) SetExternalGroupId(v string) {
	o.ExternalGroupId = &v
}

// GetMbmsLocArea returns the MbmsLocArea field value if set, zero value otherwise.
func (o *TMGIAllocation) GetMbmsLocArea() MbmsLocArea {
	if o == nil || IsNil(o.MbmsLocArea) {
		var ret MbmsLocArea
		return ret
	}
	return *o.MbmsLocArea
}

// GetMbmsLocAreaOk returns a tuple with the MbmsLocArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TMGIAllocation) GetMbmsLocAreaOk() (*MbmsLocArea, bool) {
	if o == nil || IsNil(o.MbmsLocArea) {
		return nil, false
	}
	return o.MbmsLocArea, true
}

// HasMbmsLocArea returns a boolean if a field has been set.
func (o *TMGIAllocation) HasMbmsLocArea() bool {
	if o != nil && !IsNil(o.MbmsLocArea) {
		return true
	}

	return false
}

// SetMbmsLocArea gets a reference to the given MbmsLocArea and assigns it to the MbmsLocArea field.
func (o *TMGIAllocation) SetMbmsLocArea(v MbmsLocArea) {
	o.MbmsLocArea = &v
}

// GetTmgiExpiration returns the TmgiExpiration field value if set, zero value otherwise.
func (o *TMGIAllocation) GetTmgiExpiration() time.Time {
	if o == nil || IsNil(o.TmgiExpiration) {
		var ret time.Time
		return ret
	}
	return *o.TmgiExpiration
}

// GetTmgiExpirationOk returns a tuple with the TmgiExpiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TMGIAllocation) GetTmgiExpirationOk() (*time.Time, bool) {
	if o == nil || IsNil(o.TmgiExpiration) {
		return nil, false
	}
	return o.TmgiExpiration, true
}

// HasTmgiExpiration returns a boolean if a field has been set.
func (o *TMGIAllocation) HasTmgiExpiration() bool {
	if o != nil && !IsNil(o.TmgiExpiration) {
		return true
	}

	return false
}

// SetTmgiExpiration gets a reference to the given time.Time and assigns it to the TmgiExpiration field.
func (o *TMGIAllocation) SetTmgiExpiration(v time.Time) {
	o.TmgiExpiration = &v
}

func (o TMGIAllocation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TMGIAllocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.ExternalGroupId) {
		toSerialize["externalGroupId"] = o.ExternalGroupId
	}
	if !IsNil(o.MbmsLocArea) {
		toSerialize["mbmsLocArea"] = o.MbmsLocArea
	}
	// skip: tmgiExpiration is readOnly
	return toSerialize, nil
}

type NullableTMGIAllocation struct {
	value *TMGIAllocation
	isSet bool
}

func (v NullableTMGIAllocation) Get() *TMGIAllocation {
	return v.value
}

func (v *NullableTMGIAllocation) Set(val *TMGIAllocation) {
	v.value = val
	v.isSet = true
}

func (v NullableTMGIAllocation) IsSet() bool {
	return v.isSet
}

func (v *NullableTMGIAllocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTMGIAllocation(val *TMGIAllocation) *NullableTMGIAllocation {
	return &NullableTMGIAllocation{value: val, isSet: true}
}

func (v NullableTMGIAllocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTMGIAllocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
