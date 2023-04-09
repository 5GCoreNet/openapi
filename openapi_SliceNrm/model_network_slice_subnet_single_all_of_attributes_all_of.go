/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// checks if the NetworkSliceSubnetSingleAllOfAttributesAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSliceSubnetSingleAllOfAttributesAllOf{}

// NetworkSliceSubnetSingleAllOfAttributesAllOf struct for NetworkSliceSubnetSingleAllOfAttributesAllOf
type NetworkSliceSubnetSingleAllOfAttributesAllOf struct {
	ManagedFunctionRefList    []string             `json:"managedFunctionRefList,omitempty"`
	NetworkSliceSubnetRefList []string             `json:"networkSliceSubnetRefList,omitempty"`
	OperationalState          *OperationalState    `json:"operationalState,omitempty"`
	AdministrativeState       *AdministrativeState `json:"administrativeState,omitempty"`
	NsInfo                    *NsInfo              `json:"nsInfo,omitempty"`
	SliceProfileList          []SliceProfile       `json:"sliceProfileList,omitempty"`
	EpTransportRefList        []string             `json:"epTransportRefList,omitempty"`
	PriorityLabel             *int32               `json:"priorityLabel,omitempty"`
	NetworkSliceSubnetType    *string              `json:"networkSliceSubnetType,omitempty"`
}

// NewNetworkSliceSubnetSingleAllOfAttributesAllOf instantiates a new NetworkSliceSubnetSingleAllOfAttributesAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSliceSubnetSingleAllOfAttributesAllOf() *NetworkSliceSubnetSingleAllOfAttributesAllOf {
	this := NetworkSliceSubnetSingleAllOfAttributesAllOf{}
	return &this
}

// NewNetworkSliceSubnetSingleAllOfAttributesAllOfWithDefaults instantiates a new NetworkSliceSubnetSingleAllOfAttributesAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSliceSubnetSingleAllOfAttributesAllOfWithDefaults() *NetworkSliceSubnetSingleAllOfAttributesAllOf {
	this := NetworkSliceSubnetSingleAllOfAttributesAllOf{}
	return &this
}

// GetManagedFunctionRefList returns the ManagedFunctionRefList field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetManagedFunctionRefList() []string {
	if o == nil || IsNil(o.ManagedFunctionRefList) {
		var ret []string
		return ret
	}
	return o.ManagedFunctionRefList
}

// GetManagedFunctionRefListOk returns a tuple with the ManagedFunctionRefList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetManagedFunctionRefListOk() ([]string, bool) {
	if o == nil || IsNil(o.ManagedFunctionRefList) {
		return nil, false
	}
	return o.ManagedFunctionRefList, true
}

// HasManagedFunctionRefList returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasManagedFunctionRefList() bool {
	if o != nil && !IsNil(o.ManagedFunctionRefList) {
		return true
	}

	return false
}

// SetManagedFunctionRefList gets a reference to the given []string and assigns it to the ManagedFunctionRefList field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetManagedFunctionRefList(v []string) {
	o.ManagedFunctionRefList = v
}

// GetNetworkSliceSubnetRefList returns the NetworkSliceSubnetRefList field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetNetworkSliceSubnetRefList() []string {
	if o == nil || IsNil(o.NetworkSliceSubnetRefList) {
		var ret []string
		return ret
	}
	return o.NetworkSliceSubnetRefList
}

// GetNetworkSliceSubnetRefListOk returns a tuple with the NetworkSliceSubnetRefList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetNetworkSliceSubnetRefListOk() ([]string, bool) {
	if o == nil || IsNil(o.NetworkSliceSubnetRefList) {
		return nil, false
	}
	return o.NetworkSliceSubnetRefList, true
}

// HasNetworkSliceSubnetRefList returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasNetworkSliceSubnetRefList() bool {
	if o != nil && !IsNil(o.NetworkSliceSubnetRefList) {
		return true
	}

	return false
}

// SetNetworkSliceSubnetRefList gets a reference to the given []string and assigns it to the NetworkSliceSubnetRefList field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetNetworkSliceSubnetRefList(v []string) {
	o.NetworkSliceSubnetRefList = v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetOperationalState() OperationalState {
	if o == nil || IsNil(o.OperationalState) {
		var ret OperationalState
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetOperationalStateOk() (*OperationalState, bool) {
	if o == nil || IsNil(o.OperationalState) {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasOperationalState() bool {
	if o != nil && !IsNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given OperationalState and assigns it to the OperationalState field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetOperationalState(v OperationalState) {
	o.OperationalState = &v
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetAdministrativeState() AdministrativeState {
	if o == nil || IsNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || IsNil(o.AdministrativeState) {
		return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasAdministrativeState() bool {
	if o != nil && !IsNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetNsInfo returns the NsInfo field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetNsInfo() NsInfo {
	if o == nil || IsNil(o.NsInfo) {
		var ret NsInfo
		return ret
	}
	return *o.NsInfo
}

// GetNsInfoOk returns a tuple with the NsInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetNsInfoOk() (*NsInfo, bool) {
	if o == nil || IsNil(o.NsInfo) {
		return nil, false
	}
	return o.NsInfo, true
}

// HasNsInfo returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasNsInfo() bool {
	if o != nil && !IsNil(o.NsInfo) {
		return true
	}

	return false
}

// SetNsInfo gets a reference to the given NsInfo and assigns it to the NsInfo field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetNsInfo(v NsInfo) {
	o.NsInfo = &v
}

// GetSliceProfileList returns the SliceProfileList field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetSliceProfileList() []SliceProfile {
	if o == nil || IsNil(o.SliceProfileList) {
		var ret []SliceProfile
		return ret
	}
	return o.SliceProfileList
}

// GetSliceProfileListOk returns a tuple with the SliceProfileList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetSliceProfileListOk() ([]SliceProfile, bool) {
	if o == nil || IsNil(o.SliceProfileList) {
		return nil, false
	}
	return o.SliceProfileList, true
}

// HasSliceProfileList returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasSliceProfileList() bool {
	if o != nil && !IsNil(o.SliceProfileList) {
		return true
	}

	return false
}

// SetSliceProfileList gets a reference to the given []SliceProfile and assigns it to the SliceProfileList field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetSliceProfileList(v []SliceProfile) {
	o.SliceProfileList = v
}

// GetEpTransportRefList returns the EpTransportRefList field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetEpTransportRefList() []string {
	if o == nil || IsNil(o.EpTransportRefList) {
		var ret []string
		return ret
	}
	return o.EpTransportRefList
}

// GetEpTransportRefListOk returns a tuple with the EpTransportRefList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetEpTransportRefListOk() ([]string, bool) {
	if o == nil || IsNil(o.EpTransportRefList) {
		return nil, false
	}
	return o.EpTransportRefList, true
}

// HasEpTransportRefList returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasEpTransportRefList() bool {
	if o != nil && !IsNil(o.EpTransportRefList) {
		return true
	}

	return false
}

// SetEpTransportRefList gets a reference to the given []string and assigns it to the EpTransportRefList field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetEpTransportRefList(v []string) {
	o.EpTransportRefList = v
}

// GetPriorityLabel returns the PriorityLabel field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetPriorityLabel() int32 {
	if o == nil || IsNil(o.PriorityLabel) {
		var ret int32
		return ret
	}
	return *o.PriorityLabel
}

// GetPriorityLabelOk returns a tuple with the PriorityLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetPriorityLabelOk() (*int32, bool) {
	if o == nil || IsNil(o.PriorityLabel) {
		return nil, false
	}
	return o.PriorityLabel, true
}

// HasPriorityLabel returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasPriorityLabel() bool {
	if o != nil && !IsNil(o.PriorityLabel) {
		return true
	}

	return false
}

// SetPriorityLabel gets a reference to the given int32 and assigns it to the PriorityLabel field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetPriorityLabel(v int32) {
	o.PriorityLabel = &v
}

// GetNetworkSliceSubnetType returns the NetworkSliceSubnetType field value if set, zero value otherwise.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetNetworkSliceSubnetType() string {
	if o == nil || IsNil(o.NetworkSliceSubnetType) {
		var ret string
		return ret
	}
	return *o.NetworkSliceSubnetType
}

// GetNetworkSliceSubnetTypeOk returns a tuple with the NetworkSliceSubnetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) GetNetworkSliceSubnetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkSliceSubnetType) {
		return nil, false
	}
	return o.NetworkSliceSubnetType, true
}

// HasNetworkSliceSubnetType returns a boolean if a field has been set.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) HasNetworkSliceSubnetType() bool {
	if o != nil && !IsNil(o.NetworkSliceSubnetType) {
		return true
	}

	return false
}

// SetNetworkSliceSubnetType gets a reference to the given string and assigns it to the NetworkSliceSubnetType field.
func (o *NetworkSliceSubnetSingleAllOfAttributesAllOf) SetNetworkSliceSubnetType(v string) {
	o.NetworkSliceSubnetType = &v
}

func (o NetworkSliceSubnetSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSliceSubnetSingleAllOfAttributesAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ManagedFunctionRefList) {
		toSerialize["managedFunctionRefList"] = o.ManagedFunctionRefList
	}
	if !IsNil(o.NetworkSliceSubnetRefList) {
		toSerialize["networkSliceSubnetRefList"] = o.NetworkSliceSubnetRefList
	}
	if !IsNil(o.OperationalState) {
		toSerialize["operationalState"] = o.OperationalState
	}
	if !IsNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !IsNil(o.NsInfo) {
		toSerialize["nsInfo"] = o.NsInfo
	}
	if !IsNil(o.SliceProfileList) {
		toSerialize["sliceProfileList"] = o.SliceProfileList
	}
	if !IsNil(o.EpTransportRefList) {
		toSerialize["epTransportRefList"] = o.EpTransportRefList
	}
	if !IsNil(o.PriorityLabel) {
		toSerialize["priorityLabel"] = o.PriorityLabel
	}
	if !IsNil(o.NetworkSliceSubnetType) {
		toSerialize["networkSliceSubnetType"] = o.NetworkSliceSubnetType
	}
	return toSerialize, nil
}

type NullableNetworkSliceSubnetSingleAllOfAttributesAllOf struct {
	value *NetworkSliceSubnetSingleAllOfAttributesAllOf
	isSet bool
}

func (v NullableNetworkSliceSubnetSingleAllOfAttributesAllOf) Get() *NetworkSliceSubnetSingleAllOfAttributesAllOf {
	return v.value
}

func (v *NullableNetworkSliceSubnetSingleAllOfAttributesAllOf) Set(val *NetworkSliceSubnetSingleAllOfAttributesAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSliceSubnetSingleAllOfAttributesAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSliceSubnetSingleAllOfAttributesAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSliceSubnetSingleAllOfAttributesAllOf(val *NetworkSliceSubnetSingleAllOfAttributesAllOf) *NullableNetworkSliceSubnetSingleAllOfAttributesAllOf {
	return &NullableNetworkSliceSubnetSingleAllOfAttributesAllOf{value: val, isSet: true}
}

func (v NullableNetworkSliceSubnetSingleAllOfAttributesAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSliceSubnetSingleAllOfAttributesAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
