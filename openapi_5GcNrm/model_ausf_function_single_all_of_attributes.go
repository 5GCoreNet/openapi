/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the AusfFunctionSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AusfFunctionSingleAllOfAttributes{}

// AusfFunctionSingleAllOfAttributes struct for AusfFunctionSingleAllOfAttributes
type AusfFunctionSingleAllOfAttributes struct {
	ManagedFunctionAttr
	PlmnInfoList     []PlmnInfo        `json:"plmnInfoList,omitempty"`
	SBIFqdn          *string           `json:"sBIFqdn,omitempty"`
	ManagedNFProfile *ManagedNFProfile `json:"managedNFProfile,omitempty"`
	CommModelList    []CommModel       `json:"commModelList,omitempty"`
	AusfInfo         *AusfInfo         `json:"ausfInfo,omitempty"`
}

// NewAusfFunctionSingleAllOfAttributes instantiates a new AusfFunctionSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAusfFunctionSingleAllOfAttributes() *AusfFunctionSingleAllOfAttributes {
	this := AusfFunctionSingleAllOfAttributes{}
	return &this
}

// NewAusfFunctionSingleAllOfAttributesWithDefaults instantiates a new AusfFunctionSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAusfFunctionSingleAllOfAttributesWithDefaults() *AusfFunctionSingleAllOfAttributes {
	this := AusfFunctionSingleAllOfAttributes{}
	return &this
}

// GetPlmnInfoList returns the PlmnInfoList field value if set, zero value otherwise.
func (o *AusfFunctionSingleAllOfAttributes) GetPlmnInfoList() []PlmnInfo {
	if o == nil || IsNil(o.PlmnInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PlmnInfoList
}

// GetPlmnInfoListOk returns a tuple with the PlmnInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingleAllOfAttributes) GetPlmnInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || IsNil(o.PlmnInfoList) {
		return nil, false
	}
	return o.PlmnInfoList, true
}

// HasPlmnInfoList returns a boolean if a field has been set.
func (o *AusfFunctionSingleAllOfAttributes) HasPlmnInfoList() bool {
	if o != nil && !IsNil(o.PlmnInfoList) {
		return true
	}

	return false
}

// SetPlmnInfoList gets a reference to the given []PlmnInfo and assigns it to the PlmnInfoList field.
func (o *AusfFunctionSingleAllOfAttributes) SetPlmnInfoList(v []PlmnInfo) {
	o.PlmnInfoList = v
}

// GetSBIFqdn returns the SBIFqdn field value if set, zero value otherwise.
func (o *AusfFunctionSingleAllOfAttributes) GetSBIFqdn() string {
	if o == nil || IsNil(o.SBIFqdn) {
		var ret string
		return ret
	}
	return *o.SBIFqdn
}

// GetSBIFqdnOk returns a tuple with the SBIFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingleAllOfAttributes) GetSBIFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.SBIFqdn) {
		return nil, false
	}
	return o.SBIFqdn, true
}

// HasSBIFqdn returns a boolean if a field has been set.
func (o *AusfFunctionSingleAllOfAttributes) HasSBIFqdn() bool {
	if o != nil && !IsNil(o.SBIFqdn) {
		return true
	}

	return false
}

// SetSBIFqdn gets a reference to the given string and assigns it to the SBIFqdn field.
func (o *AusfFunctionSingleAllOfAttributes) SetSBIFqdn(v string) {
	o.SBIFqdn = &v
}

// GetManagedNFProfile returns the ManagedNFProfile field value if set, zero value otherwise.
func (o *AusfFunctionSingleAllOfAttributes) GetManagedNFProfile() ManagedNFProfile {
	if o == nil || IsNil(o.ManagedNFProfile) {
		var ret ManagedNFProfile
		return ret
	}
	return *o.ManagedNFProfile
}

// GetManagedNFProfileOk returns a tuple with the ManagedNFProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingleAllOfAttributes) GetManagedNFProfileOk() (*ManagedNFProfile, bool) {
	if o == nil || IsNil(o.ManagedNFProfile) {
		return nil, false
	}
	return o.ManagedNFProfile, true
}

// HasManagedNFProfile returns a boolean if a field has been set.
func (o *AusfFunctionSingleAllOfAttributes) HasManagedNFProfile() bool {
	if o != nil && !IsNil(o.ManagedNFProfile) {
		return true
	}

	return false
}

// SetManagedNFProfile gets a reference to the given ManagedNFProfile and assigns it to the ManagedNFProfile field.
func (o *AusfFunctionSingleAllOfAttributes) SetManagedNFProfile(v ManagedNFProfile) {
	o.ManagedNFProfile = &v
}

// GetCommModelList returns the CommModelList field value if set, zero value otherwise.
func (o *AusfFunctionSingleAllOfAttributes) GetCommModelList() []CommModel {
	if o == nil || IsNil(o.CommModelList) {
		var ret []CommModel
		return ret
	}
	return o.CommModelList
}

// GetCommModelListOk returns a tuple with the CommModelList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingleAllOfAttributes) GetCommModelListOk() ([]CommModel, bool) {
	if o == nil || IsNil(o.CommModelList) {
		return nil, false
	}
	return o.CommModelList, true
}

// HasCommModelList returns a boolean if a field has been set.
func (o *AusfFunctionSingleAllOfAttributes) HasCommModelList() bool {
	if o != nil && !IsNil(o.CommModelList) {
		return true
	}

	return false
}

// SetCommModelList gets a reference to the given []CommModel and assigns it to the CommModelList field.
func (o *AusfFunctionSingleAllOfAttributes) SetCommModelList(v []CommModel) {
	o.CommModelList = v
}

// GetAusfInfo returns the AusfInfo field value if set, zero value otherwise.
func (o *AusfFunctionSingleAllOfAttributes) GetAusfInfo() AusfInfo {
	if o == nil || IsNil(o.AusfInfo) {
		var ret AusfInfo
		return ret
	}
	return *o.AusfInfo
}

// GetAusfInfoOk returns a tuple with the AusfInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AusfFunctionSingleAllOfAttributes) GetAusfInfoOk() (*AusfInfo, bool) {
	if o == nil || IsNil(o.AusfInfo) {
		return nil, false
	}
	return o.AusfInfo, true
}

// HasAusfInfo returns a boolean if a field has been set.
func (o *AusfFunctionSingleAllOfAttributes) HasAusfInfo() bool {
	if o != nil && !IsNil(o.AusfInfo) {
		return true
	}

	return false
}

// SetAusfInfo gets a reference to the given AusfInfo and assigns it to the AusfInfo field.
func (o *AusfFunctionSingleAllOfAttributes) SetAusfInfo(v AusfInfo) {
	o.AusfInfo = &v
}

func (o AusfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AusfFunctionSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedManagedFunctionAttr, errManagedFunctionAttr := json.Marshal(o.ManagedFunctionAttr)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	errManagedFunctionAttr = json.Unmarshal([]byte(serializedManagedFunctionAttr), &toSerialize)
	if errManagedFunctionAttr != nil {
		return map[string]interface{}{}, errManagedFunctionAttr
	}
	if !IsNil(o.PlmnInfoList) {
		toSerialize["plmnInfoList"] = o.PlmnInfoList
	}
	if !IsNil(o.SBIFqdn) {
		toSerialize["sBIFqdn"] = o.SBIFqdn
	}
	if !IsNil(o.ManagedNFProfile) {
		toSerialize["managedNFProfile"] = o.ManagedNFProfile
	}
	if !IsNil(o.CommModelList) {
		toSerialize["commModelList"] = o.CommModelList
	}
	if !IsNil(o.AusfInfo) {
		toSerialize["ausfInfo"] = o.AusfInfo
	}
	return toSerialize, nil
}

type NullableAusfFunctionSingleAllOfAttributes struct {
	value *AusfFunctionSingleAllOfAttributes
	isSet bool
}

func (v NullableAusfFunctionSingleAllOfAttributes) Get() *AusfFunctionSingleAllOfAttributes {
	return v.value
}

func (v *NullableAusfFunctionSingleAllOfAttributes) Set(val *AusfFunctionSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAusfFunctionSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAusfFunctionSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAusfFunctionSingleAllOfAttributes(val *AusfFunctionSingleAllOfAttributes) *NullableAusfFunctionSingleAllOfAttributes {
	return &NullableAusfFunctionSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableAusfFunctionSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAusfFunctionSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
