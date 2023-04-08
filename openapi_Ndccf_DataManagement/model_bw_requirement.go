/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the BwRequirement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BwRequirement{}

// BwRequirement Represents bandwidth requirements.
type BwRequirement struct {
	// String providing an application identifier.
	AppId string `json:"appId"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MarBwDl *string `json:"marBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MarBwUl *string `json:"marBwUl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MirBwDl *string `json:"mirBwDl,omitempty"`
	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	MirBwUl *string `json:"mirBwUl,omitempty"`
}

// NewBwRequirement instantiates a new BwRequirement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBwRequirement(appId string) *BwRequirement {
	this := BwRequirement{}
	this.AppId = appId
	return &this
}

// NewBwRequirementWithDefaults instantiates a new BwRequirement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBwRequirementWithDefaults() *BwRequirement {
	this := BwRequirement{}
	return &this
}

// GetAppId returns the AppId field value
func (o *BwRequirement) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *BwRequirement) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *BwRequirement) SetAppId(v string) {
	o.AppId = v
}

// GetMarBwDl returns the MarBwDl field value if set, zero value otherwise.
func (o *BwRequirement) GetMarBwDl() string {
	if o == nil || isNil(o.MarBwDl) {
		var ret string
		return ret
	}
	return *o.MarBwDl
}

// GetMarBwDlOk returns a tuple with the MarBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BwRequirement) GetMarBwDlOk() (*string, bool) {
	if o == nil || isNil(o.MarBwDl) {
		return nil, false
	}
	return o.MarBwDl, true
}

// HasMarBwDl returns a boolean if a field has been set.
func (o *BwRequirement) HasMarBwDl() bool {
	if o != nil && !isNil(o.MarBwDl) {
		return true
	}

	return false
}

// SetMarBwDl gets a reference to the given string and assigns it to the MarBwDl field.
func (o *BwRequirement) SetMarBwDl(v string) {
	o.MarBwDl = &v
}

// GetMarBwUl returns the MarBwUl field value if set, zero value otherwise.
func (o *BwRequirement) GetMarBwUl() string {
	if o == nil || isNil(o.MarBwUl) {
		var ret string
		return ret
	}
	return *o.MarBwUl
}

// GetMarBwUlOk returns a tuple with the MarBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BwRequirement) GetMarBwUlOk() (*string, bool) {
	if o == nil || isNil(o.MarBwUl) {
		return nil, false
	}
	return o.MarBwUl, true
}

// HasMarBwUl returns a boolean if a field has been set.
func (o *BwRequirement) HasMarBwUl() bool {
	if o != nil && !isNil(o.MarBwUl) {
		return true
	}

	return false
}

// SetMarBwUl gets a reference to the given string and assigns it to the MarBwUl field.
func (o *BwRequirement) SetMarBwUl(v string) {
	o.MarBwUl = &v
}

// GetMirBwDl returns the MirBwDl field value if set, zero value otherwise.
func (o *BwRequirement) GetMirBwDl() string {
	if o == nil || isNil(o.MirBwDl) {
		var ret string
		return ret
	}
	return *o.MirBwDl
}

// GetMirBwDlOk returns a tuple with the MirBwDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BwRequirement) GetMirBwDlOk() (*string, bool) {
	if o == nil || isNil(o.MirBwDl) {
		return nil, false
	}
	return o.MirBwDl, true
}

// HasMirBwDl returns a boolean if a field has been set.
func (o *BwRequirement) HasMirBwDl() bool {
	if o != nil && !isNil(o.MirBwDl) {
		return true
	}

	return false
}

// SetMirBwDl gets a reference to the given string and assigns it to the MirBwDl field.
func (o *BwRequirement) SetMirBwDl(v string) {
	o.MirBwDl = &v
}

// GetMirBwUl returns the MirBwUl field value if set, zero value otherwise.
func (o *BwRequirement) GetMirBwUl() string {
	if o == nil || isNil(o.MirBwUl) {
		var ret string
		return ret
	}
	return *o.MirBwUl
}

// GetMirBwUlOk returns a tuple with the MirBwUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BwRequirement) GetMirBwUlOk() (*string, bool) {
	if o == nil || isNil(o.MirBwUl) {
		return nil, false
	}
	return o.MirBwUl, true
}

// HasMirBwUl returns a boolean if a field has been set.
func (o *BwRequirement) HasMirBwUl() bool {
	if o != nil && !isNil(o.MirBwUl) {
		return true
	}

	return false
}

// SetMirBwUl gets a reference to the given string and assigns it to the MirBwUl field.
func (o *BwRequirement) SetMirBwUl(v string) {
	o.MirBwUl = &v
}

func (o BwRequirement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BwRequirement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["appId"] = o.AppId
	if !isNil(o.MarBwDl) {
		toSerialize["marBwDl"] = o.MarBwDl
	}
	if !isNil(o.MarBwUl) {
		toSerialize["marBwUl"] = o.MarBwUl
	}
	if !isNil(o.MirBwDl) {
		toSerialize["mirBwDl"] = o.MirBwDl
	}
	if !isNil(o.MirBwUl) {
		toSerialize["mirBwUl"] = o.MirBwUl
	}
	return toSerialize, nil
}

type NullableBwRequirement struct {
	value *BwRequirement
	isSet bool
}

func (v NullableBwRequirement) Get() *BwRequirement {
	return v.value
}

func (v *NullableBwRequirement) Set(val *BwRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableBwRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableBwRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBwRequirement(val *BwRequirement) *NullableBwRequirement {
	return &NullableBwRequirement{value: val, isSet: true}
}

func (v NullableBwRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBwRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


