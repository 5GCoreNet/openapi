/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
)

// checks if the DnnUpfInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnnUpfInfoItem{}

// DnnUpfInfoItem Set of parameters supported by UPF for a given DNN
type DnnUpfInfoItem struct {
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn               string             `json:"dnn"`
	DnaiList          []string           `json:"dnaiList,omitempty"`
	PduSessionTypes   []PduSessionType   `json:"pduSessionTypes,omitempty"`
	Ipv4AddressRanges []Ipv4AddressRange `json:"ipv4AddressRanges,omitempty"`
	Ipv6PrefixRanges  []Ipv6PrefixRange  `json:"ipv6PrefixRanges,omitempty"`
	Ipv4IndexList     []IpIndex          `json:"ipv4IndexList,omitempty"`
	Ipv6IndexList     []IpIndex          `json:"ipv6IndexList,omitempty"`
	// The N6 Network Instance associated with the S-NSSAI and DNN.
	NetworkInstance *string `json:"networkInstance,omitempty"`
	// Map of network instance per DNAI for the DNN, where the key of the map is the DNAI. When present, the value of each entry of the map shall contain a N6 network instance that is configured for the DNAI indicated by the key.
	DnaiNwInstanceList *map[string]string `json:"dnaiNwInstanceList,omitempty"`
}

// NewDnnUpfInfoItem instantiates a new DnnUpfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnnUpfInfoItem(dnn string) *DnnUpfInfoItem {
	this := DnnUpfInfoItem{}
	this.Dnn = dnn
	return &this
}

// NewDnnUpfInfoItemWithDefaults instantiates a new DnnUpfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnnUpfInfoItemWithDefaults() *DnnUpfInfoItem {
	this := DnnUpfInfoItem{}
	return &this
}

// GetDnn returns the Dnn field value
func (o *DnnUpfInfoItem) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *DnnUpfInfoItem) SetDnn(v string) {
	o.Dnn = v
}

// GetDnaiList returns the DnaiList field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetDnaiList() []string {
	if o == nil || IsNil(o.DnaiList) {
		var ret []string
		return ret
	}
	return o.DnaiList
}

// GetDnaiListOk returns a tuple with the DnaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetDnaiListOk() ([]string, bool) {
	if o == nil || IsNil(o.DnaiList) {
		return nil, false
	}
	return o.DnaiList, true
}

// HasDnaiList returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasDnaiList() bool {
	if o != nil && !IsNil(o.DnaiList) {
		return true
	}

	return false
}

// SetDnaiList gets a reference to the given []string and assigns it to the DnaiList field.
func (o *DnnUpfInfoItem) SetDnaiList(v []string) {
	o.DnaiList = v
}

// GetPduSessionTypes returns the PduSessionTypes field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetPduSessionTypes() []PduSessionType {
	if o == nil || IsNil(o.PduSessionTypes) {
		var ret []PduSessionType
		return ret
	}
	return o.PduSessionTypes
}

// GetPduSessionTypesOk returns a tuple with the PduSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetPduSessionTypesOk() ([]PduSessionType, bool) {
	if o == nil || IsNil(o.PduSessionTypes) {
		return nil, false
	}
	return o.PduSessionTypes, true
}

// HasPduSessionTypes returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasPduSessionTypes() bool {
	if o != nil && !IsNil(o.PduSessionTypes) {
		return true
	}

	return false
}

// SetPduSessionTypes gets a reference to the given []PduSessionType and assigns it to the PduSessionTypes field.
func (o *DnnUpfInfoItem) SetPduSessionTypes(v []PduSessionType) {
	o.PduSessionTypes = v
}

// GetIpv4AddressRanges returns the Ipv4AddressRanges field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetIpv4AddressRanges() []Ipv4AddressRange {
	if o == nil || IsNil(o.Ipv4AddressRanges) {
		var ret []Ipv4AddressRange
		return ret
	}
	return o.Ipv4AddressRanges
}

// GetIpv4AddressRangesOk returns a tuple with the Ipv4AddressRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetIpv4AddressRangesOk() ([]Ipv4AddressRange, bool) {
	if o == nil || IsNil(o.Ipv4AddressRanges) {
		return nil, false
	}
	return o.Ipv4AddressRanges, true
}

// HasIpv4AddressRanges returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasIpv4AddressRanges() bool {
	if o != nil && !IsNil(o.Ipv4AddressRanges) {
		return true
	}

	return false
}

// SetIpv4AddressRanges gets a reference to the given []Ipv4AddressRange and assigns it to the Ipv4AddressRanges field.
func (o *DnnUpfInfoItem) SetIpv4AddressRanges(v []Ipv4AddressRange) {
	o.Ipv4AddressRanges = v
}

// GetIpv6PrefixRanges returns the Ipv6PrefixRanges field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetIpv6PrefixRanges() []Ipv6PrefixRange {
	if o == nil || IsNil(o.Ipv6PrefixRanges) {
		var ret []Ipv6PrefixRange
		return ret
	}
	return o.Ipv6PrefixRanges
}

// GetIpv6PrefixRangesOk returns a tuple with the Ipv6PrefixRanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetIpv6PrefixRangesOk() ([]Ipv6PrefixRange, bool) {
	if o == nil || IsNil(o.Ipv6PrefixRanges) {
		return nil, false
	}
	return o.Ipv6PrefixRanges, true
}

// HasIpv6PrefixRanges returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasIpv6PrefixRanges() bool {
	if o != nil && !IsNil(o.Ipv6PrefixRanges) {
		return true
	}

	return false
}

// SetIpv6PrefixRanges gets a reference to the given []Ipv6PrefixRange and assigns it to the Ipv6PrefixRanges field.
func (o *DnnUpfInfoItem) SetIpv6PrefixRanges(v []Ipv6PrefixRange) {
	o.Ipv6PrefixRanges = v
}

// GetIpv4IndexList returns the Ipv4IndexList field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetIpv4IndexList() []IpIndex {
	if o == nil || IsNil(o.Ipv4IndexList) {
		var ret []IpIndex
		return ret
	}
	return o.Ipv4IndexList
}

// GetIpv4IndexListOk returns a tuple with the Ipv4IndexList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetIpv4IndexListOk() ([]IpIndex, bool) {
	if o == nil || IsNil(o.Ipv4IndexList) {
		return nil, false
	}
	return o.Ipv4IndexList, true
}

// HasIpv4IndexList returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasIpv4IndexList() bool {
	if o != nil && !IsNil(o.Ipv4IndexList) {
		return true
	}

	return false
}

// SetIpv4IndexList gets a reference to the given []IpIndex and assigns it to the Ipv4IndexList field.
func (o *DnnUpfInfoItem) SetIpv4IndexList(v []IpIndex) {
	o.Ipv4IndexList = v
}

// GetIpv6IndexList returns the Ipv6IndexList field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetIpv6IndexList() []IpIndex {
	if o == nil || IsNil(o.Ipv6IndexList) {
		var ret []IpIndex
		return ret
	}
	return o.Ipv6IndexList
}

// GetIpv6IndexListOk returns a tuple with the Ipv6IndexList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetIpv6IndexListOk() ([]IpIndex, bool) {
	if o == nil || IsNil(o.Ipv6IndexList) {
		return nil, false
	}
	return o.Ipv6IndexList, true
}

// HasIpv6IndexList returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasIpv6IndexList() bool {
	if o != nil && !IsNil(o.Ipv6IndexList) {
		return true
	}

	return false
}

// SetIpv6IndexList gets a reference to the given []IpIndex and assigns it to the Ipv6IndexList field.
func (o *DnnUpfInfoItem) SetIpv6IndexList(v []IpIndex) {
	o.Ipv6IndexList = v
}

// GetNetworkInstance returns the NetworkInstance field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetNetworkInstance() string {
	if o == nil || IsNil(o.NetworkInstance) {
		var ret string
		return ret
	}
	return *o.NetworkInstance
}

// GetNetworkInstanceOk returns a tuple with the NetworkInstance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetNetworkInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkInstance) {
		return nil, false
	}
	return o.NetworkInstance, true
}

// HasNetworkInstance returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasNetworkInstance() bool {
	if o != nil && !IsNil(o.NetworkInstance) {
		return true
	}

	return false
}

// SetNetworkInstance gets a reference to the given string and assigns it to the NetworkInstance field.
func (o *DnnUpfInfoItem) SetNetworkInstance(v string) {
	o.NetworkInstance = &v
}

// GetDnaiNwInstanceList returns the DnaiNwInstanceList field value if set, zero value otherwise.
func (o *DnnUpfInfoItem) GetDnaiNwInstanceList() map[string]string {
	if o == nil || IsNil(o.DnaiNwInstanceList) {
		var ret map[string]string
		return ret
	}
	return *o.DnaiNwInstanceList
}

// GetDnaiNwInstanceListOk returns a tuple with the DnaiNwInstanceList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnnUpfInfoItem) GetDnaiNwInstanceListOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.DnaiNwInstanceList) {
		return nil, false
	}
	return o.DnaiNwInstanceList, true
}

// HasDnaiNwInstanceList returns a boolean if a field has been set.
func (o *DnnUpfInfoItem) HasDnaiNwInstanceList() bool {
	if o != nil && !IsNil(o.DnaiNwInstanceList) {
		return true
	}

	return false
}

// SetDnaiNwInstanceList gets a reference to the given map[string]string and assigns it to the DnaiNwInstanceList field.
func (o *DnnUpfInfoItem) SetDnaiNwInstanceList(v map[string]string) {
	o.DnaiNwInstanceList = &v
}

func (o DnnUpfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnnUpfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnn"] = o.Dnn
	if !IsNil(o.DnaiList) {
		toSerialize["dnaiList"] = o.DnaiList
	}
	if !IsNil(o.PduSessionTypes) {
		toSerialize["pduSessionTypes"] = o.PduSessionTypes
	}
	if !IsNil(o.Ipv4AddressRanges) {
		toSerialize["ipv4AddressRanges"] = o.Ipv4AddressRanges
	}
	if !IsNil(o.Ipv6PrefixRanges) {
		toSerialize["ipv6PrefixRanges"] = o.Ipv6PrefixRanges
	}
	if !IsNil(o.Ipv4IndexList) {
		toSerialize["ipv4IndexList"] = o.Ipv4IndexList
	}
	if !IsNil(o.Ipv6IndexList) {
		toSerialize["ipv6IndexList"] = o.Ipv6IndexList
	}
	if !IsNil(o.NetworkInstance) {
		toSerialize["networkInstance"] = o.NetworkInstance
	}
	if !IsNil(o.DnaiNwInstanceList) {
		toSerialize["dnaiNwInstanceList"] = o.DnaiNwInstanceList
	}
	return toSerialize, nil
}

type NullableDnnUpfInfoItem struct {
	value *DnnUpfInfoItem
	isSet bool
}

func (v NullableDnnUpfInfoItem) Get() *DnnUpfInfoItem {
	return v.value
}

func (v *NullableDnnUpfInfoItem) Set(val *DnnUpfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnUpfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnUpfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnUpfInfoItem(val *DnnUpfInfoItem) *NullableDnnUpfInfoItem {
	return &NullableDnnUpfInfoItem{value: val, isSet: true}
}

func (v NullableDnnUpfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnUpfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
