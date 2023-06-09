/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the SDPMediaComponent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SDPMediaComponent{}

// SDPMediaComponent struct for SDPMediaComponent
type SDPMediaComponent struct {
	SDPMediaName                         *string             `json:"sDPMediaName,omitempty"`
	SDPMediaDescription                  []string            `json:"SDPMediaDescription,omitempty"`
	LocalGWInsertedIndication            *bool               `json:"localGWInsertedIndication,omitempty"`
	IpRealmDefaultIndication             *bool               `json:"ipRealmDefaultIndication,omitempty"`
	TranscoderInsertedIndication         *bool               `json:"transcoderInsertedIndication,omitempty"`
	MediaInitiatorFlag                   *MediaInitiatorFlag `json:"mediaInitiatorFlag,omitempty"`
	MediaInitiatorParty                  *string             `json:"mediaInitiatorParty,omitempty"`
	ThreeGPPChargingId                   *string             `json:"threeGPPChargingId,omitempty"`
	AccessNetworkChargingIdentifierValue *string             `json:"accessNetworkChargingIdentifierValue,omitempty"`
	SDPType                              *SDPType            `json:"sDPType,omitempty"`
}

// NewSDPMediaComponent instantiates a new SDPMediaComponent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSDPMediaComponent() *SDPMediaComponent {
	this := SDPMediaComponent{}
	return &this
}

// NewSDPMediaComponentWithDefaults instantiates a new SDPMediaComponent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSDPMediaComponentWithDefaults() *SDPMediaComponent {
	this := SDPMediaComponent{}
	return &this
}

// GetSDPMediaName returns the SDPMediaName field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetSDPMediaName() string {
	if o == nil || IsNil(o.SDPMediaName) {
		var ret string
		return ret
	}
	return *o.SDPMediaName
}

// GetSDPMediaNameOk returns a tuple with the SDPMediaName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetSDPMediaNameOk() (*string, bool) {
	if o == nil || IsNil(o.SDPMediaName) {
		return nil, false
	}
	return o.SDPMediaName, true
}

// HasSDPMediaName returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasSDPMediaName() bool {
	if o != nil && !IsNil(o.SDPMediaName) {
		return true
	}

	return false
}

// SetSDPMediaName gets a reference to the given string and assigns it to the SDPMediaName field.
func (o *SDPMediaComponent) SetSDPMediaName(v string) {
	o.SDPMediaName = &v
}

// GetSDPMediaDescription returns the SDPMediaDescription field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetSDPMediaDescription() []string {
	if o == nil || IsNil(o.SDPMediaDescription) {
		var ret []string
		return ret
	}
	return o.SDPMediaDescription
}

// GetSDPMediaDescriptionOk returns a tuple with the SDPMediaDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetSDPMediaDescriptionOk() ([]string, bool) {
	if o == nil || IsNil(o.SDPMediaDescription) {
		return nil, false
	}
	return o.SDPMediaDescription, true
}

// HasSDPMediaDescription returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasSDPMediaDescription() bool {
	if o != nil && !IsNil(o.SDPMediaDescription) {
		return true
	}

	return false
}

// SetSDPMediaDescription gets a reference to the given []string and assigns it to the SDPMediaDescription field.
func (o *SDPMediaComponent) SetSDPMediaDescription(v []string) {
	o.SDPMediaDescription = v
}

// GetLocalGWInsertedIndication returns the LocalGWInsertedIndication field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetLocalGWInsertedIndication() bool {
	if o == nil || IsNil(o.LocalGWInsertedIndication) {
		var ret bool
		return ret
	}
	return *o.LocalGWInsertedIndication
}

// GetLocalGWInsertedIndicationOk returns a tuple with the LocalGWInsertedIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetLocalGWInsertedIndicationOk() (*bool, bool) {
	if o == nil || IsNil(o.LocalGWInsertedIndication) {
		return nil, false
	}
	return o.LocalGWInsertedIndication, true
}

// HasLocalGWInsertedIndication returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasLocalGWInsertedIndication() bool {
	if o != nil && !IsNil(o.LocalGWInsertedIndication) {
		return true
	}

	return false
}

// SetLocalGWInsertedIndication gets a reference to the given bool and assigns it to the LocalGWInsertedIndication field.
func (o *SDPMediaComponent) SetLocalGWInsertedIndication(v bool) {
	o.LocalGWInsertedIndication = &v
}

// GetIpRealmDefaultIndication returns the IpRealmDefaultIndication field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetIpRealmDefaultIndication() bool {
	if o == nil || IsNil(o.IpRealmDefaultIndication) {
		var ret bool
		return ret
	}
	return *o.IpRealmDefaultIndication
}

// GetIpRealmDefaultIndicationOk returns a tuple with the IpRealmDefaultIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetIpRealmDefaultIndicationOk() (*bool, bool) {
	if o == nil || IsNil(o.IpRealmDefaultIndication) {
		return nil, false
	}
	return o.IpRealmDefaultIndication, true
}

// HasIpRealmDefaultIndication returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasIpRealmDefaultIndication() bool {
	if o != nil && !IsNil(o.IpRealmDefaultIndication) {
		return true
	}

	return false
}

// SetIpRealmDefaultIndication gets a reference to the given bool and assigns it to the IpRealmDefaultIndication field.
func (o *SDPMediaComponent) SetIpRealmDefaultIndication(v bool) {
	o.IpRealmDefaultIndication = &v
}

// GetTranscoderInsertedIndication returns the TranscoderInsertedIndication field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetTranscoderInsertedIndication() bool {
	if o == nil || IsNil(o.TranscoderInsertedIndication) {
		var ret bool
		return ret
	}
	return *o.TranscoderInsertedIndication
}

// GetTranscoderInsertedIndicationOk returns a tuple with the TranscoderInsertedIndication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetTranscoderInsertedIndicationOk() (*bool, bool) {
	if o == nil || IsNil(o.TranscoderInsertedIndication) {
		return nil, false
	}
	return o.TranscoderInsertedIndication, true
}

// HasTranscoderInsertedIndication returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasTranscoderInsertedIndication() bool {
	if o != nil && !IsNil(o.TranscoderInsertedIndication) {
		return true
	}

	return false
}

// SetTranscoderInsertedIndication gets a reference to the given bool and assigns it to the TranscoderInsertedIndication field.
func (o *SDPMediaComponent) SetTranscoderInsertedIndication(v bool) {
	o.TranscoderInsertedIndication = &v
}

// GetMediaInitiatorFlag returns the MediaInitiatorFlag field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetMediaInitiatorFlag() MediaInitiatorFlag {
	if o == nil || IsNil(o.MediaInitiatorFlag) {
		var ret MediaInitiatorFlag
		return ret
	}
	return *o.MediaInitiatorFlag
}

// GetMediaInitiatorFlagOk returns a tuple with the MediaInitiatorFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetMediaInitiatorFlagOk() (*MediaInitiatorFlag, bool) {
	if o == nil || IsNil(o.MediaInitiatorFlag) {
		return nil, false
	}
	return o.MediaInitiatorFlag, true
}

// HasMediaInitiatorFlag returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasMediaInitiatorFlag() bool {
	if o != nil && !IsNil(o.MediaInitiatorFlag) {
		return true
	}

	return false
}

// SetMediaInitiatorFlag gets a reference to the given MediaInitiatorFlag and assigns it to the MediaInitiatorFlag field.
func (o *SDPMediaComponent) SetMediaInitiatorFlag(v MediaInitiatorFlag) {
	o.MediaInitiatorFlag = &v
}

// GetMediaInitiatorParty returns the MediaInitiatorParty field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetMediaInitiatorParty() string {
	if o == nil || IsNil(o.MediaInitiatorParty) {
		var ret string
		return ret
	}
	return *o.MediaInitiatorParty
}

// GetMediaInitiatorPartyOk returns a tuple with the MediaInitiatorParty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetMediaInitiatorPartyOk() (*string, bool) {
	if o == nil || IsNil(o.MediaInitiatorParty) {
		return nil, false
	}
	return o.MediaInitiatorParty, true
}

// HasMediaInitiatorParty returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasMediaInitiatorParty() bool {
	if o != nil && !IsNil(o.MediaInitiatorParty) {
		return true
	}

	return false
}

// SetMediaInitiatorParty gets a reference to the given string and assigns it to the MediaInitiatorParty field.
func (o *SDPMediaComponent) SetMediaInitiatorParty(v string) {
	o.MediaInitiatorParty = &v
}

// GetThreeGPPChargingId returns the ThreeGPPChargingId field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetThreeGPPChargingId() string {
	if o == nil || IsNil(o.ThreeGPPChargingId) {
		var ret string
		return ret
	}
	return *o.ThreeGPPChargingId
}

// GetThreeGPPChargingIdOk returns a tuple with the ThreeGPPChargingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetThreeGPPChargingIdOk() (*string, bool) {
	if o == nil || IsNil(o.ThreeGPPChargingId) {
		return nil, false
	}
	return o.ThreeGPPChargingId, true
}

// HasThreeGPPChargingId returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasThreeGPPChargingId() bool {
	if o != nil && !IsNil(o.ThreeGPPChargingId) {
		return true
	}

	return false
}

// SetThreeGPPChargingId gets a reference to the given string and assigns it to the ThreeGPPChargingId field.
func (o *SDPMediaComponent) SetThreeGPPChargingId(v string) {
	o.ThreeGPPChargingId = &v
}

// GetAccessNetworkChargingIdentifierValue returns the AccessNetworkChargingIdentifierValue field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetAccessNetworkChargingIdentifierValue() string {
	if o == nil || IsNil(o.AccessNetworkChargingIdentifierValue) {
		var ret string
		return ret
	}
	return *o.AccessNetworkChargingIdentifierValue
}

// GetAccessNetworkChargingIdentifierValueOk returns a tuple with the AccessNetworkChargingIdentifierValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetAccessNetworkChargingIdentifierValueOk() (*string, bool) {
	if o == nil || IsNil(o.AccessNetworkChargingIdentifierValue) {
		return nil, false
	}
	return o.AccessNetworkChargingIdentifierValue, true
}

// HasAccessNetworkChargingIdentifierValue returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasAccessNetworkChargingIdentifierValue() bool {
	if o != nil && !IsNil(o.AccessNetworkChargingIdentifierValue) {
		return true
	}

	return false
}

// SetAccessNetworkChargingIdentifierValue gets a reference to the given string and assigns it to the AccessNetworkChargingIdentifierValue field.
func (o *SDPMediaComponent) SetAccessNetworkChargingIdentifierValue(v string) {
	o.AccessNetworkChargingIdentifierValue = &v
}

// GetSDPType returns the SDPType field value if set, zero value otherwise.
func (o *SDPMediaComponent) GetSDPType() SDPType {
	if o == nil || IsNil(o.SDPType) {
		var ret SDPType
		return ret
	}
	return *o.SDPType
}

// GetSDPTypeOk returns a tuple with the SDPType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SDPMediaComponent) GetSDPTypeOk() (*SDPType, bool) {
	if o == nil || IsNil(o.SDPType) {
		return nil, false
	}
	return o.SDPType, true
}

// HasSDPType returns a boolean if a field has been set.
func (o *SDPMediaComponent) HasSDPType() bool {
	if o != nil && !IsNil(o.SDPType) {
		return true
	}

	return false
}

// SetSDPType gets a reference to the given SDPType and assigns it to the SDPType field.
func (o *SDPMediaComponent) SetSDPType(v SDPType) {
	o.SDPType = &v
}

func (o SDPMediaComponent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SDPMediaComponent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SDPMediaName) {
		toSerialize["sDPMediaName"] = o.SDPMediaName
	}
	if !IsNil(o.SDPMediaDescription) {
		toSerialize["SDPMediaDescription"] = o.SDPMediaDescription
	}
	if !IsNil(o.LocalGWInsertedIndication) {
		toSerialize["localGWInsertedIndication"] = o.LocalGWInsertedIndication
	}
	if !IsNil(o.IpRealmDefaultIndication) {
		toSerialize["ipRealmDefaultIndication"] = o.IpRealmDefaultIndication
	}
	if !IsNil(o.TranscoderInsertedIndication) {
		toSerialize["transcoderInsertedIndication"] = o.TranscoderInsertedIndication
	}
	if !IsNil(o.MediaInitiatorFlag) {
		toSerialize["mediaInitiatorFlag"] = o.MediaInitiatorFlag
	}
	if !IsNil(o.MediaInitiatorParty) {
		toSerialize["mediaInitiatorParty"] = o.MediaInitiatorParty
	}
	if !IsNil(o.ThreeGPPChargingId) {
		toSerialize["threeGPPChargingId"] = o.ThreeGPPChargingId
	}
	if !IsNil(o.AccessNetworkChargingIdentifierValue) {
		toSerialize["accessNetworkChargingIdentifierValue"] = o.AccessNetworkChargingIdentifierValue
	}
	if !IsNil(o.SDPType) {
		toSerialize["sDPType"] = o.SDPType
	}
	return toSerialize, nil
}

type NullableSDPMediaComponent struct {
	value *SDPMediaComponent
	isSet bool
}

func (v NullableSDPMediaComponent) Get() *SDPMediaComponent {
	return v.value
}

func (v *NullableSDPMediaComponent) Set(val *SDPMediaComponent) {
	v.value = val
	v.isSet = true
}

func (v NullableSDPMediaComponent) IsSet() bool {
	return v.isSet
}

func (v *NullableSDPMediaComponent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSDPMediaComponent(val *SDPMediaComponent) *NullableSDPMediaComponent {
	return &NullableSDPMediaComponent{value: val, isSet: true}
}

func (v NullableSDPMediaComponent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSDPMediaComponent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
