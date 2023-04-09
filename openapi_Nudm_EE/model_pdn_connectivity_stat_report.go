/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
)

// checks if the PdnConnectivityStatReport type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PdnConnectivityStatReport{}

// PdnConnectivityStatReport struct for PdnConnectivityStatReport
type PdnConnectivityStatReport struct {
	PdnConnStat PdnConnectivityStatus `json:"pdnConnStat"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn *string `json:"dnn,omitempty"`
	// Unsigned integer identifying a PDU session, within the range 0 to 255, as specified in  clause 11.2.3.1b, bits 1 to 8, of 3GPP TS 24.007. If the PDU Session ID is allocated by the  Core Network for UEs not supporting N1 mode, reserved range 64 to 95 is used. PDU Session ID  within the reserved range is only visible in the Core Network.
	PduSeId *int32 `json:"pduSeId,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr     *string         `json:"ipv4Addr,omitempty"`
	Ipv6Prefixes []Ipv6Prefix    `json:"ipv6Prefixes,omitempty"`
	Ipv6Addrs    []Ipv6Addr      `json:"ipv6Addrs,omitempty"`
	PduSessType  *PduSessionType `json:"pduSessType,omitempty"`
}

// NewPdnConnectivityStatReport instantiates a new PdnConnectivityStatReport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPdnConnectivityStatReport(pdnConnStat PdnConnectivityStatus) *PdnConnectivityStatReport {
	this := PdnConnectivityStatReport{}
	this.PdnConnStat = pdnConnStat
	return &this
}

// NewPdnConnectivityStatReportWithDefaults instantiates a new PdnConnectivityStatReport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPdnConnectivityStatReportWithDefaults() *PdnConnectivityStatReport {
	this := PdnConnectivityStatReport{}
	return &this
}

// GetPdnConnStat returns the PdnConnStat field value
func (o *PdnConnectivityStatReport) GetPdnConnStat() PdnConnectivityStatus {
	if o == nil {
		var ret PdnConnectivityStatus
		return ret
	}

	return o.PdnConnStat
}

// GetPdnConnStatOk returns a tuple with the PdnConnStat field value
// and a boolean to check if the value has been set.
func (o *PdnConnectivityStatReport) GetPdnConnStatOk() (*PdnConnectivityStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PdnConnStat, true
}

// SetPdnConnStat sets field value
func (o *PdnConnectivityStatReport) SetPdnConnStat(v PdnConnectivityStatus) {
	o.PdnConnStat = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *PdnConnectivityStatReport) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectivityStatReport) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *PdnConnectivityStatReport) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *PdnConnectivityStatReport) SetDnn(v string) {
	o.Dnn = &v
}

// GetPduSeId returns the PduSeId field value if set, zero value otherwise.
func (o *PdnConnectivityStatReport) GetPduSeId() int32 {
	if o == nil || IsNil(o.PduSeId) {
		var ret int32
		return ret
	}
	return *o.PduSeId
}

// GetPduSeIdOk returns a tuple with the PduSeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectivityStatReport) GetPduSeIdOk() (*int32, bool) {
	if o == nil || IsNil(o.PduSeId) {
		return nil, false
	}
	return o.PduSeId, true
}

// HasPduSeId returns a boolean if a field has been set.
func (o *PdnConnectivityStatReport) HasPduSeId() bool {
	if o != nil && !IsNil(o.PduSeId) {
		return true
	}

	return false
}

// SetPduSeId gets a reference to the given int32 and assigns it to the PduSeId field.
func (o *PdnConnectivityStatReport) SetPduSeId(v int32) {
	o.PduSeId = &v
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *PdnConnectivityStatReport) GetIpv4Addr() string {
	if o == nil || IsNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectivityStatReport) GetIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Addr) {
		return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *PdnConnectivityStatReport) HasIpv4Addr() bool {
	if o != nil && !IsNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *PdnConnectivityStatReport) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Prefixes returns the Ipv6Prefixes field value if set, zero value otherwise.
func (o *PdnConnectivityStatReport) GetIpv6Prefixes() []Ipv6Prefix {
	if o == nil || IsNil(o.Ipv6Prefixes) {
		var ret []Ipv6Prefix
		return ret
	}
	return o.Ipv6Prefixes
}

// GetIpv6PrefixesOk returns a tuple with the Ipv6Prefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectivityStatReport) GetIpv6PrefixesOk() ([]Ipv6Prefix, bool) {
	if o == nil || IsNil(o.Ipv6Prefixes) {
		return nil, false
	}
	return o.Ipv6Prefixes, true
}

// HasIpv6Prefixes returns a boolean if a field has been set.
func (o *PdnConnectivityStatReport) HasIpv6Prefixes() bool {
	if o != nil && !IsNil(o.Ipv6Prefixes) {
		return true
	}

	return false
}

// SetIpv6Prefixes gets a reference to the given []Ipv6Prefix and assigns it to the Ipv6Prefixes field.
func (o *PdnConnectivityStatReport) SetIpv6Prefixes(v []Ipv6Prefix) {
	o.Ipv6Prefixes = v
}

// GetIpv6Addrs returns the Ipv6Addrs field value if set, zero value otherwise.
func (o *PdnConnectivityStatReport) GetIpv6Addrs() []Ipv6Addr {
	if o == nil || IsNil(o.Ipv6Addrs) {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6Addrs
}

// GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectivityStatReport) GetIpv6AddrsOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.Ipv6Addrs) {
		return nil, false
	}
	return o.Ipv6Addrs, true
}

// HasIpv6Addrs returns a boolean if a field has been set.
func (o *PdnConnectivityStatReport) HasIpv6Addrs() bool {
	if o != nil && !IsNil(o.Ipv6Addrs) {
		return true
	}

	return false
}

// SetIpv6Addrs gets a reference to the given []Ipv6Addr and assigns it to the Ipv6Addrs field.
func (o *PdnConnectivityStatReport) SetIpv6Addrs(v []Ipv6Addr) {
	o.Ipv6Addrs = v
}

// GetPduSessType returns the PduSessType field value if set, zero value otherwise.
func (o *PdnConnectivityStatReport) GetPduSessType() PduSessionType {
	if o == nil || IsNil(o.PduSessType) {
		var ret PduSessionType
		return ret
	}
	return *o.PduSessType
}

// GetPduSessTypeOk returns a tuple with the PduSessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PdnConnectivityStatReport) GetPduSessTypeOk() (*PduSessionType, bool) {
	if o == nil || IsNil(o.PduSessType) {
		return nil, false
	}
	return o.PduSessType, true
}

// HasPduSessType returns a boolean if a field has been set.
func (o *PdnConnectivityStatReport) HasPduSessType() bool {
	if o != nil && !IsNil(o.PduSessType) {
		return true
	}

	return false
}

// SetPduSessType gets a reference to the given PduSessionType and assigns it to the PduSessType field.
func (o *PdnConnectivityStatReport) SetPduSessType(v PduSessionType) {
	o.PduSessType = &v
}

func (o PdnConnectivityStatReport) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PdnConnectivityStatReport) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pdnConnStat"] = o.PdnConnStat
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.PduSeId) {
		toSerialize["pduSeId"] = o.PduSeId
	}
	if !IsNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !IsNil(o.Ipv6Prefixes) {
		toSerialize["ipv6Prefixes"] = o.Ipv6Prefixes
	}
	if !IsNil(o.Ipv6Addrs) {
		toSerialize["ipv6Addrs"] = o.Ipv6Addrs
	}
	if !IsNil(o.PduSessType) {
		toSerialize["pduSessType"] = o.PduSessType
	}
	return toSerialize, nil
}

type NullablePdnConnectivityStatReport struct {
	value *PdnConnectivityStatReport
	isSet bool
}

func (v NullablePdnConnectivityStatReport) Get() *PdnConnectivityStatReport {
	return v.value
}

func (v *NullablePdnConnectivityStatReport) Set(val *PdnConnectivityStatReport) {
	v.value = val
	v.isSet = true
}

func (v NullablePdnConnectivityStatReport) IsSet() bool {
	return v.isSet
}

func (v *NullablePdnConnectivityStatReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePdnConnectivityStatReport(val *PdnConnectivityStatReport) *NullablePdnConnectivityStatReport {
	return &NullablePdnConnectivityStatReport{value: val, isSet: true}
}

func (v NullablePdnConnectivityStatReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePdnConnectivityStatReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
