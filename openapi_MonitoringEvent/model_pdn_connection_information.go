/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the PdnConnectionInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PdnConnectionInformation{}

// PdnConnectionInformation Represents the PDN connection information of the UE.
type PdnConnectionInformation struct {
	Status PdnConnectionStatus `json:"status"`
	// Identify the APN, it is depending on the SCEF local configuration whether or not this attribute is sent to the SCS/AS.
	Apn          *string              `json:"apn,omitempty"`
	PdnType      PdnType              `json:"pdnType"`
	InterfaceInd *InterfaceIndication `json:"interfaceInd,omitempty"`
	// string identifying a Ipv4 address formatted in the \"dotted decimal\" notation as defined in IETF RFC 1166.
	Ipv4Addr  *string  `json:"ipv4Addr,omitempty"`
	Ipv6Addrs []string `json:"ipv6Addrs,omitempty"`
	MacAddrs  []string `json:"macAddrs,omitempty"`
}

// NewPdnConnectionInformation instantiates a new PdnConnectionInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPdnConnectionInformation(status PdnConnectionStatus, pdnType PdnType) *PdnConnectionInformation {
	this := PdnConnectionInformation{}
	this.Status = status
	this.PdnType = pdnType
	return &this
}

// NewPdnConnectionInformationWithDefaults instantiates a new PdnConnectionInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPdnConnectionInformationWithDefaults() *PdnConnectionInformation {
	this := PdnConnectionInformation{}
	return &this
}

// GetStatus returns the Status field value
func (o *PdnConnectionInformation) GetStatus() PdnConnectionStatus {
	if o == nil {
		var ret PdnConnectionStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *PdnConnectionInformation) GetStatusOk() (*PdnConnectionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *PdnConnectionInformation) SetStatus(v PdnConnectionStatus) {
	o.Status = v
}

// GetApn returns the Apn field value if set, zero value otherwise.
func (o *PdnConnectionInformation) GetApn() string {
	if o == nil || IsNil(o.Apn) {
		var ret string
		return ret
	}
	return *o.Apn
}

// GetApnOk returns a tuple with the Apn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectionInformation) GetApnOk() (*string, bool) {
	if o == nil || IsNil(o.Apn) {
		return nil, false
	}
	return o.Apn, true
}

// HasApn returns a boolean if a field has been set.
func (o *PdnConnectionInformation) HasApn() bool {
	if o != nil && !IsNil(o.Apn) {
		return true
	}

	return false
}

// SetApn gets a reference to the given string and assigns it to the Apn field.
func (o *PdnConnectionInformation) SetApn(v string) {
	o.Apn = &v
}

// GetPdnType returns the PdnType field value
func (o *PdnConnectionInformation) GetPdnType() PdnType {
	if o == nil {
		var ret PdnType
		return ret
	}

	return o.PdnType
}

// GetPdnTypeOk returns a tuple with the PdnType field value
// and a boolean to check if the value has been set.
func (o *PdnConnectionInformation) GetPdnTypeOk() (*PdnType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PdnType, true
}

// SetPdnType sets field value
func (o *PdnConnectionInformation) SetPdnType(v PdnType) {
	o.PdnType = v
}

// GetInterfaceInd returns the InterfaceInd field value if set, zero value otherwise.
func (o *PdnConnectionInformation) GetInterfaceInd() InterfaceIndication {
	if o == nil || IsNil(o.InterfaceInd) {
		var ret InterfaceIndication
		return ret
	}
	return *o.InterfaceInd
}

// GetInterfaceIndOk returns a tuple with the InterfaceInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectionInformation) GetInterfaceIndOk() (*InterfaceIndication, bool) {
	if o == nil || IsNil(o.InterfaceInd) {
		return nil, false
	}
	return o.InterfaceInd, true
}

// HasInterfaceInd returns a boolean if a field has been set.
func (o *PdnConnectionInformation) HasInterfaceInd() bool {
	if o != nil && !IsNil(o.InterfaceInd) {
		return true
	}

	return false
}

// SetInterfaceInd gets a reference to the given InterfaceIndication and assigns it to the InterfaceInd field.
func (o *PdnConnectionInformation) SetInterfaceInd(v InterfaceIndication) {
	o.InterfaceInd = &v
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *PdnConnectionInformation) GetIpv4Addr() string {
	if o == nil || IsNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectionInformation) GetIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Addr) {
		return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *PdnConnectionInformation) HasIpv4Addr() bool {
	if o != nil && !IsNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *PdnConnectionInformation) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Addrs returns the Ipv6Addrs field value if set, zero value otherwise.
func (o *PdnConnectionInformation) GetIpv6Addrs() []string {
	if o == nil || IsNil(o.Ipv6Addrs) {
		var ret []string
		return ret
	}
	return o.Ipv6Addrs
}

// GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectionInformation) GetIpv6AddrsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ipv6Addrs) {
		return nil, false
	}
	return o.Ipv6Addrs, true
}

// HasIpv6Addrs returns a boolean if a field has been set.
func (o *PdnConnectionInformation) HasIpv6Addrs() bool {
	if o != nil && !IsNil(o.Ipv6Addrs) {
		return true
	}

	return false
}

// SetIpv6Addrs gets a reference to the given []string and assigns it to the Ipv6Addrs field.
func (o *PdnConnectionInformation) SetIpv6Addrs(v []string) {
	o.Ipv6Addrs = v
}

// GetMacAddrs returns the MacAddrs field value if set, zero value otherwise.
func (o *PdnConnectionInformation) GetMacAddrs() []string {
	if o == nil || IsNil(o.MacAddrs) {
		var ret []string
		return ret
	}
	return o.MacAddrs
}

// GetMacAddrsOk returns a tuple with the MacAddrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectionInformation) GetMacAddrsOk() ([]string, bool) {
	if o == nil || IsNil(o.MacAddrs) {
		return nil, false
	}
	return o.MacAddrs, true
}

// HasMacAddrs returns a boolean if a field has been set.
func (o *PdnConnectionInformation) HasMacAddrs() bool {
	if o != nil && !IsNil(o.MacAddrs) {
		return true
	}

	return false
}

// SetMacAddrs gets a reference to the given []string and assigns it to the MacAddrs field.
func (o *PdnConnectionInformation) SetMacAddrs(v []string) {
	o.MacAddrs = v
}

func (o PdnConnectionInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PdnConnectionInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Apn) {
		toSerialize["apn"] = o.Apn
	}
	toSerialize["pdnType"] = o.PdnType
	if !IsNil(o.InterfaceInd) {
		toSerialize["interfaceInd"] = o.InterfaceInd
	}
	if !IsNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !IsNil(o.Ipv6Addrs) {
		toSerialize["ipv6Addrs"] = o.Ipv6Addrs
	}
	if !IsNil(o.MacAddrs) {
		toSerialize["macAddrs"] = o.MacAddrs
	}
	return toSerialize, nil
}

type NullablePdnConnectionInformation struct {
	value *PdnConnectionInformation
	isSet bool
}

func (v NullablePdnConnectionInformation) Get() *PdnConnectionInformation {
	return v.value
}

func (v *NullablePdnConnectionInformation) Set(val *PdnConnectionInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePdnConnectionInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePdnConnectionInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdnConnectionInformation(val *PdnConnectionInformation) *NullablePdnConnectionInformation {
	return &NullablePdnConnectionInformation{value: val, isSet: true}
}

func (v NullablePdnConnectionInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdnConnectionInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
