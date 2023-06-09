/*
3gpp-5glan-pp

API for 5G LAN Parameter Provision.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GLANParameterProvision

import (
	"encoding/json"
)

// checks if the Model5GLanParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model5GLanParameters{}

// Model5GLanParameters Represents 5G LAN service related parameters that need to be provisioned.
type Model5GLanParameters struct {
	// string containing a local identifier followed by \"@\" and a domain identifier. Both the local identifier and the domain identifier shall be encoded as strings that do not contain any \"@\" characters. See Clauses 4.6.2 and 4.6.3 of 3GPP TS 23.682 for more information.
	ExterGroupId string `json:"exterGroupId"`
	// Contains the list of 5G VN Group members, each member is identified by GPSI. Any string value can be used as a key of the map.
	Gpsis map[string]string `json:"gpsis"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	AaaIpv4Addr *string    `json:"aaaIpv4Addr,omitempty"`
	AaaIpv6Addr *Ipv6Addr  `json:"aaaIpv6Addr,omitempty"`
	AaaUsgs     []AaaUsage `json:"aaaUsgs,omitempty"`
	// String uniquely identifying MTC provider information.
	MtcProviderId *string        `json:"mtcProviderId,omitempty"`
	Snssai        Snssai         `json:"snssai"`
	SessionType   PduSessionType `json:"sessionType"`
	// Further allowed PDU Session types.
	SessionTypes []PduSessionType `json:"sessionTypes,omitempty"`
	// Describes the operation systems and the corresponding applications for each operation systems. The key of map is osId.
	AppDesps map[string]AppDescriptor `json:"appDesps"`
}

// NewModel5GLanParameters instantiates a new Model5GLanParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel5GLanParameters(exterGroupId string, gpsis map[string]string, dnn string, snssai Snssai, sessionType PduSessionType, appDesps map[string]AppDescriptor) *Model5GLanParameters {
	this := Model5GLanParameters{}
	this.ExterGroupId = exterGroupId
	this.Gpsis = gpsis
	this.Dnn = dnn
	this.Snssai = snssai
	this.SessionType = sessionType
	this.AppDesps = appDesps
	return &this
}

// NewModel5GLanParametersWithDefaults instantiates a new Model5GLanParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel5GLanParametersWithDefaults() *Model5GLanParameters {
	this := Model5GLanParameters{}
	return &this
}

// GetExterGroupId returns the ExterGroupId field value
func (o *Model5GLanParameters) GetExterGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExterGroupId
}

// GetExterGroupIdOk returns a tuple with the ExterGroupId field value
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetExterGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExterGroupId, true
}

// SetExterGroupId sets field value
func (o *Model5GLanParameters) SetExterGroupId(v string) {
	o.ExterGroupId = v
}

// GetGpsis returns the Gpsis field value
func (o *Model5GLanParameters) GetGpsis() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetGpsisOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gpsis, true
}

// SetGpsis sets field value
func (o *Model5GLanParameters) SetGpsis(v map[string]string) {
	o.Gpsis = v
}

// GetDnn returns the Dnn field value
func (o *Model5GLanParameters) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *Model5GLanParameters) SetDnn(v string) {
	o.Dnn = v
}

// GetAaaIpv4Addr returns the AaaIpv4Addr field value if set, zero value otherwise.
func (o *Model5GLanParameters) GetAaaIpv4Addr() string {
	if o == nil || IsNil(o.AaaIpv4Addr) {
		var ret string
		return ret
	}
	return *o.AaaIpv4Addr
}

// GetAaaIpv4AddrOk returns a tuple with the AaaIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetAaaIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.AaaIpv4Addr) {
		return nil, false
	}
	return o.AaaIpv4Addr, true
}

// HasAaaIpv4Addr returns a boolean if a field has been set.
func (o *Model5GLanParameters) HasAaaIpv4Addr() bool {
	if o != nil && !IsNil(o.AaaIpv4Addr) {
		return true
	}

	return false
}

// SetAaaIpv4Addr gets a reference to the given string and assigns it to the AaaIpv4Addr field.
func (o *Model5GLanParameters) SetAaaIpv4Addr(v string) {
	o.AaaIpv4Addr = &v
}

// GetAaaIpv6Addr returns the AaaIpv6Addr field value if set, zero value otherwise.
func (o *Model5GLanParameters) GetAaaIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.AaaIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.AaaIpv6Addr
}

// GetAaaIpv6AddrOk returns a tuple with the AaaIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetAaaIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.AaaIpv6Addr) {
		return nil, false
	}
	return o.AaaIpv6Addr, true
}

// HasAaaIpv6Addr returns a boolean if a field has been set.
func (o *Model5GLanParameters) HasAaaIpv6Addr() bool {
	if o != nil && !IsNil(o.AaaIpv6Addr) {
		return true
	}

	return false
}

// SetAaaIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the AaaIpv6Addr field.
func (o *Model5GLanParameters) SetAaaIpv6Addr(v Ipv6Addr) {
	o.AaaIpv6Addr = &v
}

// GetAaaUsgs returns the AaaUsgs field value if set, zero value otherwise.
func (o *Model5GLanParameters) GetAaaUsgs() []AaaUsage {
	if o == nil || IsNil(o.AaaUsgs) {
		var ret []AaaUsage
		return ret
	}
	return o.AaaUsgs
}

// GetAaaUsgsOk returns a tuple with the AaaUsgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetAaaUsgsOk() ([]AaaUsage, bool) {
	if o == nil || IsNil(o.AaaUsgs) {
		return nil, false
	}
	return o.AaaUsgs, true
}

// HasAaaUsgs returns a boolean if a field has been set.
func (o *Model5GLanParameters) HasAaaUsgs() bool {
	if o != nil && !IsNil(o.AaaUsgs) {
		return true
	}

	return false
}

// SetAaaUsgs gets a reference to the given []AaaUsage and assigns it to the AaaUsgs field.
func (o *Model5GLanParameters) SetAaaUsgs(v []AaaUsage) {
	o.AaaUsgs = v
}

// GetMtcProviderId returns the MtcProviderId field value if set, zero value otherwise.
func (o *Model5GLanParameters) GetMtcProviderId() string {
	if o == nil || IsNil(o.MtcProviderId) {
		var ret string
		return ret
	}
	return *o.MtcProviderId
}

// GetMtcProviderIdOk returns a tuple with the MtcProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetMtcProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderId) {
		return nil, false
	}
	return o.MtcProviderId, true
}

// HasMtcProviderId returns a boolean if a field has been set.
func (o *Model5GLanParameters) HasMtcProviderId() bool {
	if o != nil && !IsNil(o.MtcProviderId) {
		return true
	}

	return false
}

// SetMtcProviderId gets a reference to the given string and assigns it to the MtcProviderId field.
func (o *Model5GLanParameters) SetMtcProviderId(v string) {
	o.MtcProviderId = &v
}

// GetSnssai returns the Snssai field value
func (o *Model5GLanParameters) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *Model5GLanParameters) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetSessionType returns the SessionType field value
func (o *Model5GLanParameters) GetSessionType() PduSessionType {
	if o == nil {
		var ret PduSessionType
		return ret
	}

	return o.SessionType
}

// GetSessionTypeOk returns a tuple with the SessionType field value
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetSessionTypeOk() (*PduSessionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionType, true
}

// SetSessionType sets field value
func (o *Model5GLanParameters) SetSessionType(v PduSessionType) {
	o.SessionType = v
}

// GetSessionTypes returns the SessionTypes field value if set, zero value otherwise.
func (o *Model5GLanParameters) GetSessionTypes() []PduSessionType {
	if o == nil || IsNil(o.SessionTypes) {
		var ret []PduSessionType
		return ret
	}
	return o.SessionTypes
}

// GetSessionTypesOk returns a tuple with the SessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetSessionTypesOk() ([]PduSessionType, bool) {
	if o == nil || IsNil(o.SessionTypes) {
		return nil, false
	}
	return o.SessionTypes, true
}

// HasSessionTypes returns a boolean if a field has been set.
func (o *Model5GLanParameters) HasSessionTypes() bool {
	if o != nil && !IsNil(o.SessionTypes) {
		return true
	}

	return false
}

// SetSessionTypes gets a reference to the given []PduSessionType and assigns it to the SessionTypes field.
func (o *Model5GLanParameters) SetSessionTypes(v []PduSessionType) {
	o.SessionTypes = v
}

// GetAppDesps returns the AppDesps field value
func (o *Model5GLanParameters) GetAppDesps() map[string]AppDescriptor {
	if o == nil {
		var ret map[string]AppDescriptor
		return ret
	}

	return o.AppDesps
}

// GetAppDespsOk returns a tuple with the AppDesps field value
// and a boolean to check if the value has been set.
func (o *Model5GLanParameters) GetAppDespsOk() (*map[string]AppDescriptor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppDesps, true
}

// SetAppDesps sets field value
func (o *Model5GLanParameters) SetAppDesps(v map[string]AppDescriptor) {
	o.AppDesps = v
}

func (o Model5GLanParameters) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model5GLanParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exterGroupId"] = o.ExterGroupId
	toSerialize["gpsis"] = o.Gpsis
	toSerialize["dnn"] = o.Dnn
	if !IsNil(o.AaaIpv4Addr) {
		toSerialize["aaaIpv4Addr"] = o.AaaIpv4Addr
	}
	if !IsNil(o.AaaIpv6Addr) {
		toSerialize["aaaIpv6Addr"] = o.AaaIpv6Addr
	}
	if !IsNil(o.AaaUsgs) {
		toSerialize["aaaUsgs"] = o.AaaUsgs
	}
	if !IsNil(o.MtcProviderId) {
		toSerialize["mtcProviderId"] = o.MtcProviderId
	}
	toSerialize["snssai"] = o.Snssai
	toSerialize["sessionType"] = o.SessionType
	if !IsNil(o.SessionTypes) {
		toSerialize["sessionTypes"] = o.SessionTypes
	}
	toSerialize["appDesps"] = o.AppDesps
	return toSerialize, nil
}

type NullableModel5GLanParameters struct {
	value *Model5GLanParameters
	isSet bool
}

func (v NullableModel5GLanParameters) Get() *Model5GLanParameters {
	return v.value
}

func (v *NullableModel5GLanParameters) Set(val *Model5GLanParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GLanParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GLanParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GLanParameters(val *Model5GLanParameters) *NullableModel5GLanParameters {
	return &NullableModel5GLanParameters{value: val, isSet: true}
}

func (v NullableModel5GLanParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GLanParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
