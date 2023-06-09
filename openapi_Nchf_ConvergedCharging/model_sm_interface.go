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

// checks if the SMInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMInterface{}

// SMInterface struct for SMInterface
type SMInterface struct {
	InterfaceId   *string        `json:"interfaceId,omitempty"`
	InterfaceText *string        `json:"interfaceText,omitempty"`
	InterfacePort *string        `json:"interfacePort,omitempty"`
	InterfaceType *InterfaceType `json:"interfaceType,omitempty"`
}

// NewSMInterface instantiates a new SMInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMInterface() *SMInterface {
	this := SMInterface{}
	return &this
}

// NewSMInterfaceWithDefaults instantiates a new SMInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMInterfaceWithDefaults() *SMInterface {
	this := SMInterface{}
	return &this
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *SMInterface) GetInterfaceId() string {
	if o == nil || IsNil(o.InterfaceId) {
		var ret string
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMInterface) GetInterfaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceId) {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *SMInterface) HasInterfaceId() bool {
	if o != nil && !IsNil(o.InterfaceId) {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given string and assigns it to the InterfaceId field.
func (o *SMInterface) SetInterfaceId(v string) {
	o.InterfaceId = &v
}

// GetInterfaceText returns the InterfaceText field value if set, zero value otherwise.
func (o *SMInterface) GetInterfaceText() string {
	if o == nil || IsNil(o.InterfaceText) {
		var ret string
		return ret
	}
	return *o.InterfaceText
}

// GetInterfaceTextOk returns a tuple with the InterfaceText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMInterface) GetInterfaceTextOk() (*string, bool) {
	if o == nil || IsNil(o.InterfaceText) {
		return nil, false
	}
	return o.InterfaceText, true
}

// HasInterfaceText returns a boolean if a field has been set.
func (o *SMInterface) HasInterfaceText() bool {
	if o != nil && !IsNil(o.InterfaceText) {
		return true
	}

	return false
}

// SetInterfaceText gets a reference to the given string and assigns it to the InterfaceText field.
func (o *SMInterface) SetInterfaceText(v string) {
	o.InterfaceText = &v
}

// GetInterfacePort returns the InterfacePort field value if set, zero value otherwise.
func (o *SMInterface) GetInterfacePort() string {
	if o == nil || IsNil(o.InterfacePort) {
		var ret string
		return ret
	}
	return *o.InterfacePort
}

// GetInterfacePortOk returns a tuple with the InterfacePort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMInterface) GetInterfacePortOk() (*string, bool) {
	if o == nil || IsNil(o.InterfacePort) {
		return nil, false
	}
	return o.InterfacePort, true
}

// HasInterfacePort returns a boolean if a field has been set.
func (o *SMInterface) HasInterfacePort() bool {
	if o != nil && !IsNil(o.InterfacePort) {
		return true
	}

	return false
}

// SetInterfacePort gets a reference to the given string and assigns it to the InterfacePort field.
func (o *SMInterface) SetInterfacePort(v string) {
	o.InterfacePort = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *SMInterface) GetInterfaceType() InterfaceType {
	if o == nil || IsNil(o.InterfaceType) {
		var ret InterfaceType
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMInterface) GetInterfaceTypeOk() (*InterfaceType, bool) {
	if o == nil || IsNil(o.InterfaceType) {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *SMInterface) HasInterfaceType() bool {
	if o != nil && !IsNil(o.InterfaceType) {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given InterfaceType and assigns it to the InterfaceType field.
func (o *SMInterface) SetInterfaceType(v InterfaceType) {
	o.InterfaceType = &v
}

func (o SMInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InterfaceId) {
		toSerialize["interfaceId"] = o.InterfaceId
	}
	if !IsNil(o.InterfaceText) {
		toSerialize["interfaceText"] = o.InterfaceText
	}
	if !IsNil(o.InterfacePort) {
		toSerialize["interfacePort"] = o.InterfacePort
	}
	if !IsNil(o.InterfaceType) {
		toSerialize["interfaceType"] = o.InterfaceType
	}
	return toSerialize, nil
}

type NullableSMInterface struct {
	value *SMInterface
	isSet bool
}

func (v NullableSMInterface) Get() *SMInterface {
	return v.value
}

func (v *NullableSMInterface) Set(val *SMInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableSMInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableSMInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMInterface(val *SMInterface) *NullableSMInterface {
	return &NullableSMInterface{value: val, isSet: true}
}

func (v NullableSMInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
