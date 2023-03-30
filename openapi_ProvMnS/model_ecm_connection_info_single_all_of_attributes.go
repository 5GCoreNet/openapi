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

// checks if the EcmConnectionInfoSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EcmConnectionInfoSingleAllOfAttributes{}

// EcmConnectionInfoSingleAllOfAttributes struct for EcmConnectionInfoSingleAllOfAttributes
type EcmConnectionInfoSingleAllOfAttributes struct {
	EASServiceArea *ServingLocation `json:"eASServiceArea,omitempty"`
	EESServiceArea *ServingLocation `json:"eESServiceArea,omitempty"`
	EDNServiceArea *ServingLocation `json:"eDNServiceArea,omitempty"`
	EASIpAddress *string `json:"eASIpAddress,omitempty"`
	EESIpAddress *string `json:"eESIpAddress,omitempty"`
	ECSIpAddress *string `json:"eCSIpAddress,omitempty"`
	EdnIdentifier *string `json:"ednIdentifier,omitempty"`
	EcmConnectionType *string `json:"ecmConnectionType,omitempty"`
	Var5GCNfConnEcmInfoList []Model5GCNfConnEcmInfo `json:"5GCNfConnEcmInfoList,omitempty"`
	UPFConnectionInfo *UPFConnectionInfo `json:"uPFConnectionInfo,omitempty"`
}

// NewEcmConnectionInfoSingleAllOfAttributes instantiates a new EcmConnectionInfoSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcmConnectionInfoSingleAllOfAttributes() *EcmConnectionInfoSingleAllOfAttributes {
	this := EcmConnectionInfoSingleAllOfAttributes{}
	return &this
}

// NewEcmConnectionInfoSingleAllOfAttributesWithDefaults instantiates a new EcmConnectionInfoSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcmConnectionInfoSingleAllOfAttributesWithDefaults() *EcmConnectionInfoSingleAllOfAttributes {
	this := EcmConnectionInfoSingleAllOfAttributes{}
	return &this
}

// GetEASServiceArea returns the EASServiceArea field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEASServiceArea() ServingLocation {
	if o == nil || IsNil(o.EASServiceArea) {
		var ret ServingLocation
		return ret
	}
	return *o.EASServiceArea
}

// GetEASServiceAreaOk returns a tuple with the EASServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEASServiceAreaOk() (*ServingLocation, bool) {
	if o == nil || IsNil(o.EASServiceArea) {
		return nil, false
	}
	return o.EASServiceArea, true
}

// HasEASServiceArea returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasEASServiceArea() bool {
	if o != nil && !IsNil(o.EASServiceArea) {
		return true
	}

	return false
}

// SetEASServiceArea gets a reference to the given ServingLocation and assigns it to the EASServiceArea field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetEASServiceArea(v ServingLocation) {
	o.EASServiceArea = &v
}

// GetEESServiceArea returns the EESServiceArea field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEESServiceArea() ServingLocation {
	if o == nil || IsNil(o.EESServiceArea) {
		var ret ServingLocation
		return ret
	}
	return *o.EESServiceArea
}

// GetEESServiceAreaOk returns a tuple with the EESServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEESServiceAreaOk() (*ServingLocation, bool) {
	if o == nil || IsNil(o.EESServiceArea) {
		return nil, false
	}
	return o.EESServiceArea, true
}

// HasEESServiceArea returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasEESServiceArea() bool {
	if o != nil && !IsNil(o.EESServiceArea) {
		return true
	}

	return false
}

// SetEESServiceArea gets a reference to the given ServingLocation and assigns it to the EESServiceArea field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetEESServiceArea(v ServingLocation) {
	o.EESServiceArea = &v
}

// GetEDNServiceArea returns the EDNServiceArea field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEDNServiceArea() ServingLocation {
	if o == nil || IsNil(o.EDNServiceArea) {
		var ret ServingLocation
		return ret
	}
	return *o.EDNServiceArea
}

// GetEDNServiceAreaOk returns a tuple with the EDNServiceArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEDNServiceAreaOk() (*ServingLocation, bool) {
	if o == nil || IsNil(o.EDNServiceArea) {
		return nil, false
	}
	return o.EDNServiceArea, true
}

// HasEDNServiceArea returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasEDNServiceArea() bool {
	if o != nil && !IsNil(o.EDNServiceArea) {
		return true
	}

	return false
}

// SetEDNServiceArea gets a reference to the given ServingLocation and assigns it to the EDNServiceArea field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetEDNServiceArea(v ServingLocation) {
	o.EDNServiceArea = &v
}

// GetEASIpAddress returns the EASIpAddress field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEASIpAddress() string {
	if o == nil || IsNil(o.EASIpAddress) {
		var ret string
		return ret
	}
	return *o.EASIpAddress
}

// GetEASIpAddressOk returns a tuple with the EASIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEASIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EASIpAddress) {
		return nil, false
	}
	return o.EASIpAddress, true
}

// HasEASIpAddress returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasEASIpAddress() bool {
	if o != nil && !IsNil(o.EASIpAddress) {
		return true
	}

	return false
}

// SetEASIpAddress gets a reference to the given string and assigns it to the EASIpAddress field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetEASIpAddress(v string) {
	o.EASIpAddress = &v
}

// GetEESIpAddress returns the EESIpAddress field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEESIpAddress() string {
	if o == nil || IsNil(o.EESIpAddress) {
		var ret string
		return ret
	}
	return *o.EESIpAddress
}

// GetEESIpAddressOk returns a tuple with the EESIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEESIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EESIpAddress) {
		return nil, false
	}
	return o.EESIpAddress, true
}

// HasEESIpAddress returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasEESIpAddress() bool {
	if o != nil && !IsNil(o.EESIpAddress) {
		return true
	}

	return false
}

// SetEESIpAddress gets a reference to the given string and assigns it to the EESIpAddress field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetEESIpAddress(v string) {
	o.EESIpAddress = &v
}

// GetECSIpAddress returns the ECSIpAddress field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetECSIpAddress() string {
	if o == nil || IsNil(o.ECSIpAddress) {
		var ret string
		return ret
	}
	return *o.ECSIpAddress
}

// GetECSIpAddressOk returns a tuple with the ECSIpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetECSIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.ECSIpAddress) {
		return nil, false
	}
	return o.ECSIpAddress, true
}

// HasECSIpAddress returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasECSIpAddress() bool {
	if o != nil && !IsNil(o.ECSIpAddress) {
		return true
	}

	return false
}

// SetECSIpAddress gets a reference to the given string and assigns it to the ECSIpAddress field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetECSIpAddress(v string) {
	o.ECSIpAddress = &v
}

// GetEdnIdentifier returns the EdnIdentifier field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEdnIdentifier() string {
	if o == nil || IsNil(o.EdnIdentifier) {
		var ret string
		return ret
	}
	return *o.EdnIdentifier
}

// GetEdnIdentifierOk returns a tuple with the EdnIdentifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEdnIdentifierOk() (*string, bool) {
	if o == nil || IsNil(o.EdnIdentifier) {
		return nil, false
	}
	return o.EdnIdentifier, true
}

// HasEdnIdentifier returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasEdnIdentifier() bool {
	if o != nil && !IsNil(o.EdnIdentifier) {
		return true
	}

	return false
}

// SetEdnIdentifier gets a reference to the given string and assigns it to the EdnIdentifier field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetEdnIdentifier(v string) {
	o.EdnIdentifier = &v
}

// GetEcmConnectionType returns the EcmConnectionType field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEcmConnectionType() string {
	if o == nil || IsNil(o.EcmConnectionType) {
		var ret string
		return ret
	}
	return *o.EcmConnectionType
}

// GetEcmConnectionTypeOk returns a tuple with the EcmConnectionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetEcmConnectionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EcmConnectionType) {
		return nil, false
	}
	return o.EcmConnectionType, true
}

// HasEcmConnectionType returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasEcmConnectionType() bool {
	if o != nil && !IsNil(o.EcmConnectionType) {
		return true
	}

	return false
}

// SetEcmConnectionType gets a reference to the given string and assigns it to the EcmConnectionType field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetEcmConnectionType(v string) {
	o.EcmConnectionType = &v
}

// GetVar5GCNfConnEcmInfoList returns the Var5GCNfConnEcmInfoList field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetVar5GCNfConnEcmInfoList() []Model5GCNfConnEcmInfo {
	if o == nil || IsNil(o.Var5GCNfConnEcmInfoList) {
		var ret []Model5GCNfConnEcmInfo
		return ret
	}
	return o.Var5GCNfConnEcmInfoList
}

// GetVar5GCNfConnEcmInfoListOk returns a tuple with the Var5GCNfConnEcmInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetVar5GCNfConnEcmInfoListOk() ([]Model5GCNfConnEcmInfo, bool) {
	if o == nil || IsNil(o.Var5GCNfConnEcmInfoList) {
		return nil, false
	}
	return o.Var5GCNfConnEcmInfoList, true
}

// HasVar5GCNfConnEcmInfoList returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasVar5GCNfConnEcmInfoList() bool {
	if o != nil && !IsNil(o.Var5GCNfConnEcmInfoList) {
		return true
	}

	return false
}

// SetVar5GCNfConnEcmInfoList gets a reference to the given []Model5GCNfConnEcmInfo and assigns it to the Var5GCNfConnEcmInfoList field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetVar5GCNfConnEcmInfoList(v []Model5GCNfConnEcmInfo) {
	o.Var5GCNfConnEcmInfoList = v
}

// GetUPFConnectionInfo returns the UPFConnectionInfo field value if set, zero value otherwise.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetUPFConnectionInfo() UPFConnectionInfo {
	if o == nil || IsNil(o.UPFConnectionInfo) {
		var ret UPFConnectionInfo
		return ret
	}
	return *o.UPFConnectionInfo
}

// GetUPFConnectionInfoOk returns a tuple with the UPFConnectionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) GetUPFConnectionInfoOk() (*UPFConnectionInfo, bool) {
	if o == nil || IsNil(o.UPFConnectionInfo) {
		return nil, false
	}
	return o.UPFConnectionInfo, true
}

// HasUPFConnectionInfo returns a boolean if a field has been set.
func (o *EcmConnectionInfoSingleAllOfAttributes) HasUPFConnectionInfo() bool {
	if o != nil && !IsNil(o.UPFConnectionInfo) {
		return true
	}

	return false
}

// SetUPFConnectionInfo gets a reference to the given UPFConnectionInfo and assigns it to the UPFConnectionInfo field.
func (o *EcmConnectionInfoSingleAllOfAttributes) SetUPFConnectionInfo(v UPFConnectionInfo) {
	o.UPFConnectionInfo = &v
}

func (o EcmConnectionInfoSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcmConnectionInfoSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EASServiceArea) {
		toSerialize["eASServiceArea"] = o.EASServiceArea
	}
	if !IsNil(o.EESServiceArea) {
		toSerialize["eESServiceArea"] = o.EESServiceArea
	}
	if !IsNil(o.EDNServiceArea) {
		toSerialize["eDNServiceArea"] = o.EDNServiceArea
	}
	if !IsNil(o.EASIpAddress) {
		toSerialize["eASIpAddress"] = o.EASIpAddress
	}
	if !IsNil(o.EESIpAddress) {
		toSerialize["eESIpAddress"] = o.EESIpAddress
	}
	if !IsNil(o.ECSIpAddress) {
		toSerialize["eCSIpAddress"] = o.ECSIpAddress
	}
	if !IsNil(o.EdnIdentifier) {
		toSerialize["ednIdentifier"] = o.EdnIdentifier
	}
	if !IsNil(o.EcmConnectionType) {
		toSerialize["ecmConnectionType"] = o.EcmConnectionType
	}
	if !IsNil(o.Var5GCNfConnEcmInfoList) {
		toSerialize["5GCNfConnEcmInfoList"] = o.Var5GCNfConnEcmInfoList
	}
	if !IsNil(o.UPFConnectionInfo) {
		toSerialize["uPFConnectionInfo"] = o.UPFConnectionInfo
	}
	return toSerialize, nil
}

type NullableEcmConnectionInfoSingleAllOfAttributes struct {
	value *EcmConnectionInfoSingleAllOfAttributes
	isSet bool
}

func (v NullableEcmConnectionInfoSingleAllOfAttributes) Get() *EcmConnectionInfoSingleAllOfAttributes {
	return v.value
}

func (v *NullableEcmConnectionInfoSingleAllOfAttributes) Set(val *EcmConnectionInfoSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEcmConnectionInfoSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEcmConnectionInfoSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcmConnectionInfoSingleAllOfAttributes(val *EcmConnectionInfoSingleAllOfAttributes) *NullableEcmConnectionInfoSingleAllOfAttributes {
	return &NullableEcmConnectionInfoSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableEcmConnectionInfoSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcmConnectionInfoSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


