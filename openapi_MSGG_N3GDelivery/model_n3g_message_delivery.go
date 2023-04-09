/*
MSGG_N3GDelivery

API for MSGG N3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSGG_N3GDelivery

import (
	"encoding/json"
)

// checks if the N3gMessageDelivery type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N3gMessageDelivery{}

// N3gMessageDelivery N3G message delivery data
type N3gMessageDelivery struct {
	OriAddr       Address                   `json:"oriAddr"`
	DestAddr      Address                   `json:"destAddr"`
	AppId         *string                   `json:"appId,omitempty"`
	MsgId         string                    `json:"msgId"`
	DelivStReqInd *bool                     `json:"delivStReqInd,omitempty"`
	Payload       *string                   `json:"payload,omitempty"`
	SegInd        *bool                     `json:"segInd,omitempty"`
	SegParams     *MessageSegmentParameters `json:"segParams,omitempty"`
}

// NewN3gMessageDelivery instantiates a new N3gMessageDelivery object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN3gMessageDelivery(oriAddr Address, destAddr Address, msgId string) *N3gMessageDelivery {
	this := N3gMessageDelivery{}
	this.OriAddr = oriAddr
	this.DestAddr = destAddr
	this.MsgId = msgId
	return &this
}

// NewN3gMessageDeliveryWithDefaults instantiates a new N3gMessageDelivery object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN3gMessageDeliveryWithDefaults() *N3gMessageDelivery {
	this := N3gMessageDelivery{}
	return &this
}

// GetOriAddr returns the OriAddr field value
func (o *N3gMessageDelivery) GetOriAddr() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.OriAddr
}

// GetOriAddrOk returns a tuple with the OriAddr field value
// and a boolean to check if the value has been set.
func (o *N3gMessageDelivery) GetOriAddrOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriAddr, true
}

// SetOriAddr sets field value
func (o *N3gMessageDelivery) SetOriAddr(v Address) {
	o.OriAddr = v
}

// GetDestAddr returns the DestAddr field value
func (o *N3gMessageDelivery) GetDestAddr() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.DestAddr
}

// GetDestAddrOk returns a tuple with the DestAddr field value
// and a boolean to check if the value has been set.
func (o *N3gMessageDelivery) GetDestAddrOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestAddr, true
}

// SetDestAddr sets field value
func (o *N3gMessageDelivery) SetDestAddr(v Address) {
	o.DestAddr = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *N3gMessageDelivery) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gMessageDelivery) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *N3gMessageDelivery) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *N3gMessageDelivery) SetAppId(v string) {
	o.AppId = &v
}

// GetMsgId returns the MsgId field value
func (o *N3gMessageDelivery) GetMsgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MsgId
}

// GetMsgIdOk returns a tuple with the MsgId field value
// and a boolean to check if the value has been set.
func (o *N3gMessageDelivery) GetMsgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MsgId, true
}

// SetMsgId sets field value
func (o *N3gMessageDelivery) SetMsgId(v string) {
	o.MsgId = v
}

// GetDelivStReqInd returns the DelivStReqInd field value if set, zero value otherwise.
func (o *N3gMessageDelivery) GetDelivStReqInd() bool {
	if o == nil || IsNil(o.DelivStReqInd) {
		var ret bool
		return ret
	}
	return *o.DelivStReqInd
}

// GetDelivStReqIndOk returns a tuple with the DelivStReqInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gMessageDelivery) GetDelivStReqIndOk() (*bool, bool) {
	if o == nil || IsNil(o.DelivStReqInd) {
		return nil, false
	}
	return o.DelivStReqInd, true
}

// HasDelivStReqInd returns a boolean if a field has been set.
func (o *N3gMessageDelivery) HasDelivStReqInd() bool {
	if o != nil && !IsNil(o.DelivStReqInd) {
		return true
	}

	return false
}

// SetDelivStReqInd gets a reference to the given bool and assigns it to the DelivStReqInd field.
func (o *N3gMessageDelivery) SetDelivStReqInd(v bool) {
	o.DelivStReqInd = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *N3gMessageDelivery) GetPayload() string {
	if o == nil || IsNil(o.Payload) {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gMessageDelivery) GetPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *N3gMessageDelivery) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *N3gMessageDelivery) SetPayload(v string) {
	o.Payload = &v
}

// GetSegInd returns the SegInd field value if set, zero value otherwise.
func (o *N3gMessageDelivery) GetSegInd() bool {
	if o == nil || IsNil(o.SegInd) {
		var ret bool
		return ret
	}
	return *o.SegInd
}

// GetSegIndOk returns a tuple with the SegInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gMessageDelivery) GetSegIndOk() (*bool, bool) {
	if o == nil || IsNil(o.SegInd) {
		return nil, false
	}
	return o.SegInd, true
}

// HasSegInd returns a boolean if a field has been set.
func (o *N3gMessageDelivery) HasSegInd() bool {
	if o != nil && !IsNil(o.SegInd) {
		return true
	}

	return false
}

// SetSegInd gets a reference to the given bool and assigns it to the SegInd field.
func (o *N3gMessageDelivery) SetSegInd(v bool) {
	o.SegInd = &v
}

// GetSegParams returns the SegParams field value if set, zero value otherwise.
func (o *N3gMessageDelivery) GetSegParams() MessageSegmentParameters {
	if o == nil || IsNil(o.SegParams) {
		var ret MessageSegmentParameters
		return ret
	}
	return *o.SegParams
}

// GetSegParamsOk returns a tuple with the SegParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N3gMessageDelivery) GetSegParamsOk() (*MessageSegmentParameters, bool) {
	if o == nil || IsNil(o.SegParams) {
		return nil, false
	}
	return o.SegParams, true
}

// HasSegParams returns a boolean if a field has been set.
func (o *N3gMessageDelivery) HasSegParams() bool {
	if o != nil && !IsNil(o.SegParams) {
		return true
	}

	return false
}

// SetSegParams gets a reference to the given MessageSegmentParameters and assigns it to the SegParams field.
func (o *N3gMessageDelivery) SetSegParams(v MessageSegmentParameters) {
	o.SegParams = &v
}

func (o N3gMessageDelivery) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N3gMessageDelivery) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["oriAddr"] = o.OriAddr
	toSerialize["destAddr"] = o.DestAddr
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	toSerialize["msgId"] = o.MsgId
	if !IsNil(o.DelivStReqInd) {
		toSerialize["delivStReqInd"] = o.DelivStReqInd
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.SegInd) {
		toSerialize["segInd"] = o.SegInd
	}
	if !IsNil(o.SegParams) {
		toSerialize["segParams"] = o.SegParams
	}
	return toSerialize, nil
}

type NullableN3gMessageDelivery struct {
	value *N3gMessageDelivery
	isSet bool
}

func (v NullableN3gMessageDelivery) Get() *N3gMessageDelivery {
	return v.value
}

func (v *NullableN3gMessageDelivery) Set(val *N3gMessageDelivery) {
	v.value = val
	v.isSet = true
}

func (v NullableN3gMessageDelivery) IsSet() bool {
	return v.isSet
}

func (v *NullableN3gMessageDelivery) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN3gMessageDelivery(val *N3gMessageDelivery) *NullableN3gMessageDelivery {
	return &NullableN3gMessageDelivery{value: val, isSet: true}
}

func (v NullableN3gMessageDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN3gMessageDelivery) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
