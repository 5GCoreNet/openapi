/*
Neasdf_BaselineDNSPattern

EASDF Baseline DNS Pattern Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_BaselineDNSPattern

import (
	"encoding/json"
)

// checks if the BaselineDnsMdt type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaselineDnsMdt{}

// BaselineDnsMdt Baseline DNS message detection template
type BaselineDnsMdt struct {
	MdtId string `json:"mdtId"`
	Label *string `json:"label,omitempty"`
	// map of DNS query message detection templates where a valid JSON string serves as key
	DnsQueryMdtList *map[string]DnsQueryMdt `json:"dnsQueryMdtList,omitempty"`
	// map of DNS response message detection templates where a valid JSON string serves as key
	DnsRspMdtList *map[string]DnsRspMdt `json:"dnsRspMdtList,omitempty"`
}

// NewBaselineDnsMdt instantiates a new BaselineDnsMdt object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaselineDnsMdt(mdtId string) *BaselineDnsMdt {
	this := BaselineDnsMdt{}
	this.MdtId = mdtId
	return &this
}

// NewBaselineDnsMdtWithDefaults instantiates a new BaselineDnsMdt object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaselineDnsMdtWithDefaults() *BaselineDnsMdt {
	this := BaselineDnsMdt{}
	return &this
}

// GetMdtId returns the MdtId field value
func (o *BaselineDnsMdt) GetMdtId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MdtId
}

// GetMdtIdOk returns a tuple with the MdtId field value
// and a boolean to check if the value has been set.
func (o *BaselineDnsMdt) GetMdtIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MdtId, true
}

// SetMdtId sets field value
func (o *BaselineDnsMdt) SetMdtId(v string) {
	o.MdtId = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *BaselineDnsMdt) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaselineDnsMdt) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *BaselineDnsMdt) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *BaselineDnsMdt) SetLabel(v string) {
	o.Label = &v
}

// GetDnsQueryMdtList returns the DnsQueryMdtList field value if set, zero value otherwise.
func (o *BaselineDnsMdt) GetDnsQueryMdtList() map[string]DnsQueryMdt {
	if o == nil || IsNil(o.DnsQueryMdtList) {
		var ret map[string]DnsQueryMdt
		return ret
	}
	return *o.DnsQueryMdtList
}

// GetDnsQueryMdtListOk returns a tuple with the DnsQueryMdtList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaselineDnsMdt) GetDnsQueryMdtListOk() (*map[string]DnsQueryMdt, bool) {
	if o == nil || IsNil(o.DnsQueryMdtList) {
		return nil, false
	}
	return o.DnsQueryMdtList, true
}

// HasDnsQueryMdtList returns a boolean if a field has been set.
func (o *BaselineDnsMdt) HasDnsQueryMdtList() bool {
	if o != nil && !IsNil(o.DnsQueryMdtList) {
		return true
	}

	return false
}

// SetDnsQueryMdtList gets a reference to the given map[string]DnsQueryMdt and assigns it to the DnsQueryMdtList field.
func (o *BaselineDnsMdt) SetDnsQueryMdtList(v map[string]DnsQueryMdt) {
	o.DnsQueryMdtList = &v
}

// GetDnsRspMdtList returns the DnsRspMdtList field value if set, zero value otherwise.
func (o *BaselineDnsMdt) GetDnsRspMdtList() map[string]DnsRspMdt {
	if o == nil || IsNil(o.DnsRspMdtList) {
		var ret map[string]DnsRspMdt
		return ret
	}
	return *o.DnsRspMdtList
}

// GetDnsRspMdtListOk returns a tuple with the DnsRspMdtList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaselineDnsMdt) GetDnsRspMdtListOk() (*map[string]DnsRspMdt, bool) {
	if o == nil || IsNil(o.DnsRspMdtList) {
		return nil, false
	}
	return o.DnsRspMdtList, true
}

// HasDnsRspMdtList returns a boolean if a field has been set.
func (o *BaselineDnsMdt) HasDnsRspMdtList() bool {
	if o != nil && !IsNil(o.DnsRspMdtList) {
		return true
	}

	return false
}

// SetDnsRspMdtList gets a reference to the given map[string]DnsRspMdt and assigns it to the DnsRspMdtList field.
func (o *BaselineDnsMdt) SetDnsRspMdtList(v map[string]DnsRspMdt) {
	o.DnsRspMdtList = &v
}

func (o BaselineDnsMdt) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaselineDnsMdt) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mdtId"] = o.MdtId
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.DnsQueryMdtList) {
		toSerialize["dnsQueryMdtList"] = o.DnsQueryMdtList
	}
	if !IsNil(o.DnsRspMdtList) {
		toSerialize["dnsRspMdtList"] = o.DnsRspMdtList
	}
	return toSerialize, nil
}

type NullableBaselineDnsMdt struct {
	value *BaselineDnsMdt
	isSet bool
}

func (v NullableBaselineDnsMdt) Get() *BaselineDnsMdt {
	return v.value
}

func (v *NullableBaselineDnsMdt) Set(val *BaselineDnsMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableBaselineDnsMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableBaselineDnsMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaselineDnsMdt(val *BaselineDnsMdt) *NullableBaselineDnsMdt {
	return &NullableBaselineDnsMdt{value: val, isSet: true}
}

func (v NullableBaselineDnsMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaselineDnsMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

