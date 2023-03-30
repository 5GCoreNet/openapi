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

// checks if the EPF1USingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EPF1USingleAllOfAttributes{}

// EPF1USingleAllOfAttributes struct for EPF1USingleAllOfAttributes
type EPF1USingleAllOfAttributes struct {
	UserLabel *string `json:"userLabel,omitempty"`
	FarEndEntity *string `json:"farEndEntity,omitempty"`
	SupportedPerfMetricGroups []SupportedPerfMetricGroup `json:"supportedPerfMetricGroups,omitempty"`
	LocalAddress *LocalAddress `json:"localAddress,omitempty"`
	RemoteAddress *RemoteAddress `json:"remoteAddress,omitempty"`
	EpTransportRefs []string `json:"epTransportRefs,omitempty"`
}

// NewEPF1USingleAllOfAttributes instantiates a new EPF1USingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEPF1USingleAllOfAttributes() *EPF1USingleAllOfAttributes {
	this := EPF1USingleAllOfAttributes{}
	return &this
}

// NewEPF1USingleAllOfAttributesWithDefaults instantiates a new EPF1USingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEPF1USingleAllOfAttributesWithDefaults() *EPF1USingleAllOfAttributes {
	this := EPF1USingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *EPF1USingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1USingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *EPF1USingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *EPF1USingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetFarEndEntity returns the FarEndEntity field value if set, zero value otherwise.
func (o *EPF1USingleAllOfAttributes) GetFarEndEntity() string {
	if o == nil || IsNil(o.FarEndEntity) {
		var ret string
		return ret
	}
	return *o.FarEndEntity
}

// GetFarEndEntityOk returns a tuple with the FarEndEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1USingleAllOfAttributes) GetFarEndEntityOk() (*string, bool) {
	if o == nil || IsNil(o.FarEndEntity) {
		return nil, false
	}
	return o.FarEndEntity, true
}

// HasFarEndEntity returns a boolean if a field has been set.
func (o *EPF1USingleAllOfAttributes) HasFarEndEntity() bool {
	if o != nil && !IsNil(o.FarEndEntity) {
		return true
	}

	return false
}

// SetFarEndEntity gets a reference to the given string and assigns it to the FarEndEntity field.
func (o *EPF1USingleAllOfAttributes) SetFarEndEntity(v string) {
	o.FarEndEntity = &v
}

// GetSupportedPerfMetricGroups returns the SupportedPerfMetricGroups field value if set, zero value otherwise.
func (o *EPF1USingleAllOfAttributes) GetSupportedPerfMetricGroups() []SupportedPerfMetricGroup {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		var ret []SupportedPerfMetricGroup
		return ret
	}
	return o.SupportedPerfMetricGroups
}

// GetSupportedPerfMetricGroupsOk returns a tuple with the SupportedPerfMetricGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1USingleAllOfAttributes) GetSupportedPerfMetricGroupsOk() ([]SupportedPerfMetricGroup, bool) {
	if o == nil || IsNil(o.SupportedPerfMetricGroups) {
		return nil, false
	}
	return o.SupportedPerfMetricGroups, true
}

// HasSupportedPerfMetricGroups returns a boolean if a field has been set.
func (o *EPF1USingleAllOfAttributes) HasSupportedPerfMetricGroups() bool {
	if o != nil && !IsNil(o.SupportedPerfMetricGroups) {
		return true
	}

	return false
}

// SetSupportedPerfMetricGroups gets a reference to the given []SupportedPerfMetricGroup and assigns it to the SupportedPerfMetricGroups field.
func (o *EPF1USingleAllOfAttributes) SetSupportedPerfMetricGroups(v []SupportedPerfMetricGroup) {
	o.SupportedPerfMetricGroups = v
}

// GetLocalAddress returns the LocalAddress field value if set, zero value otherwise.
func (o *EPF1USingleAllOfAttributes) GetLocalAddress() LocalAddress {
	if o == nil || IsNil(o.LocalAddress) {
		var ret LocalAddress
		return ret
	}
	return *o.LocalAddress
}

// GetLocalAddressOk returns a tuple with the LocalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1USingleAllOfAttributes) GetLocalAddressOk() (*LocalAddress, bool) {
	if o == nil || IsNil(o.LocalAddress) {
		return nil, false
	}
	return o.LocalAddress, true
}

// HasLocalAddress returns a boolean if a field has been set.
func (o *EPF1USingleAllOfAttributes) HasLocalAddress() bool {
	if o != nil && !IsNil(o.LocalAddress) {
		return true
	}

	return false
}

// SetLocalAddress gets a reference to the given LocalAddress and assigns it to the LocalAddress field.
func (o *EPF1USingleAllOfAttributes) SetLocalAddress(v LocalAddress) {
	o.LocalAddress = &v
}

// GetRemoteAddress returns the RemoteAddress field value if set, zero value otherwise.
func (o *EPF1USingleAllOfAttributes) GetRemoteAddress() RemoteAddress {
	if o == nil || IsNil(o.RemoteAddress) {
		var ret RemoteAddress
		return ret
	}
	return *o.RemoteAddress
}

// GetRemoteAddressOk returns a tuple with the RemoteAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1USingleAllOfAttributes) GetRemoteAddressOk() (*RemoteAddress, bool) {
	if o == nil || IsNil(o.RemoteAddress) {
		return nil, false
	}
	return o.RemoteAddress, true
}

// HasRemoteAddress returns a boolean if a field has been set.
func (o *EPF1USingleAllOfAttributes) HasRemoteAddress() bool {
	if o != nil && !IsNil(o.RemoteAddress) {
		return true
	}

	return false
}

// SetRemoteAddress gets a reference to the given RemoteAddress and assigns it to the RemoteAddress field.
func (o *EPF1USingleAllOfAttributes) SetRemoteAddress(v RemoteAddress) {
	o.RemoteAddress = &v
}

// GetEpTransportRefs returns the EpTransportRefs field value if set, zero value otherwise.
func (o *EPF1USingleAllOfAttributes) GetEpTransportRefs() []string {
	if o == nil || IsNil(o.EpTransportRefs) {
		var ret []string
		return ret
	}
	return o.EpTransportRefs
}

// GetEpTransportRefsOk returns a tuple with the EpTransportRefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EPF1USingleAllOfAttributes) GetEpTransportRefsOk() ([]string, bool) {
	if o == nil || IsNil(o.EpTransportRefs) {
		return nil, false
	}
	return o.EpTransportRefs, true
}

// HasEpTransportRefs returns a boolean if a field has been set.
func (o *EPF1USingleAllOfAttributes) HasEpTransportRefs() bool {
	if o != nil && !IsNil(o.EpTransportRefs) {
		return true
	}

	return false
}

// SetEpTransportRefs gets a reference to the given []string and assigns it to the EpTransportRefs field.
func (o *EPF1USingleAllOfAttributes) SetEpTransportRefs(v []string) {
	o.EpTransportRefs = v
}

func (o EPF1USingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EPF1USingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !IsNil(o.FarEndEntity) {
		toSerialize["farEndEntity"] = o.FarEndEntity
	}
	if !IsNil(o.SupportedPerfMetricGroups) {
		toSerialize["supportedPerfMetricGroups"] = o.SupportedPerfMetricGroups
	}
	if !IsNil(o.LocalAddress) {
		toSerialize["localAddress"] = o.LocalAddress
	}
	if !IsNil(o.RemoteAddress) {
		toSerialize["remoteAddress"] = o.RemoteAddress
	}
	if !IsNil(o.EpTransportRefs) {
		toSerialize["epTransportRefs"] = o.EpTransportRefs
	}
	return toSerialize, nil
}

type NullableEPF1USingleAllOfAttributes struct {
	value *EPF1USingleAllOfAttributes
	isSet bool
}

func (v NullableEPF1USingleAllOfAttributes) Get() *EPF1USingleAllOfAttributes {
	return v.value
}

func (v *NullableEPF1USingleAllOfAttributes) Set(val *EPF1USingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEPF1USingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEPF1USingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEPF1USingleAllOfAttributes(val *EPF1USingleAllOfAttributes) *NullableEPF1USingleAllOfAttributes {
	return &NullableEPF1USingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableEPF1USingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEPF1USingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


