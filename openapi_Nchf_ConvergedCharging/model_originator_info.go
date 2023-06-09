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

// checks if the OriginatorInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OriginatorInfo{}

// OriginatorInfo struct for OriginatorInfo
type OriginatorInfo struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	OriginatorSUPI *string `json:"originatorSUPI,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	OriginatorGPSI            *string        `json:"originatorGPSI,omitempty"`
	OriginatorOtherAddress    *SMAddressInfo `json:"originatorOtherAddress,omitempty"`
	OriginatorReceivedAddress *SMAddressInfo `json:"originatorReceivedAddress,omitempty"`
	OriginatorSCCPAddress     *string        `json:"originatorSCCPAddress,omitempty"`
	SMOriginatorInterface     *SMInterface   `json:"sMOriginatorInterface,omitempty"`
	SMOriginatorProtocolId    *string        `json:"sMOriginatorProtocolId,omitempty"`
}

// NewOriginatorInfo instantiates a new OriginatorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOriginatorInfo() *OriginatorInfo {
	this := OriginatorInfo{}
	return &this
}

// NewOriginatorInfoWithDefaults instantiates a new OriginatorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOriginatorInfoWithDefaults() *OriginatorInfo {
	this := OriginatorInfo{}
	return &this
}

// GetOriginatorSUPI returns the OriginatorSUPI field value if set, zero value otherwise.
func (o *OriginatorInfo) GetOriginatorSUPI() string {
	if o == nil || IsNil(o.OriginatorSUPI) {
		var ret string
		return ret
	}
	return *o.OriginatorSUPI
}

// GetOriginatorSUPIOk returns a tuple with the OriginatorSUPI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginatorInfo) GetOriginatorSUPIOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatorSUPI) {
		return nil, false
	}
	return o.OriginatorSUPI, true
}

// HasOriginatorSUPI returns a boolean if a field has been set.
func (o *OriginatorInfo) HasOriginatorSUPI() bool {
	if o != nil && !IsNil(o.OriginatorSUPI) {
		return true
	}

	return false
}

// SetOriginatorSUPI gets a reference to the given string and assigns it to the OriginatorSUPI field.
func (o *OriginatorInfo) SetOriginatorSUPI(v string) {
	o.OriginatorSUPI = &v
}

// GetOriginatorGPSI returns the OriginatorGPSI field value if set, zero value otherwise.
func (o *OriginatorInfo) GetOriginatorGPSI() string {
	if o == nil || IsNil(o.OriginatorGPSI) {
		var ret string
		return ret
	}
	return *o.OriginatorGPSI
}

// GetOriginatorGPSIOk returns a tuple with the OriginatorGPSI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginatorInfo) GetOriginatorGPSIOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatorGPSI) {
		return nil, false
	}
	return o.OriginatorGPSI, true
}

// HasOriginatorGPSI returns a boolean if a field has been set.
func (o *OriginatorInfo) HasOriginatorGPSI() bool {
	if o != nil && !IsNil(o.OriginatorGPSI) {
		return true
	}

	return false
}

// SetOriginatorGPSI gets a reference to the given string and assigns it to the OriginatorGPSI field.
func (o *OriginatorInfo) SetOriginatorGPSI(v string) {
	o.OriginatorGPSI = &v
}

// GetOriginatorOtherAddress returns the OriginatorOtherAddress field value if set, zero value otherwise.
func (o *OriginatorInfo) GetOriginatorOtherAddress() SMAddressInfo {
	if o == nil || IsNil(o.OriginatorOtherAddress) {
		var ret SMAddressInfo
		return ret
	}
	return *o.OriginatorOtherAddress
}

// GetOriginatorOtherAddressOk returns a tuple with the OriginatorOtherAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginatorInfo) GetOriginatorOtherAddressOk() (*SMAddressInfo, bool) {
	if o == nil || IsNil(o.OriginatorOtherAddress) {
		return nil, false
	}
	return o.OriginatorOtherAddress, true
}

// HasOriginatorOtherAddress returns a boolean if a field has been set.
func (o *OriginatorInfo) HasOriginatorOtherAddress() bool {
	if o != nil && !IsNil(o.OriginatorOtherAddress) {
		return true
	}

	return false
}

// SetOriginatorOtherAddress gets a reference to the given SMAddressInfo and assigns it to the OriginatorOtherAddress field.
func (o *OriginatorInfo) SetOriginatorOtherAddress(v SMAddressInfo) {
	o.OriginatorOtherAddress = &v
}

// GetOriginatorReceivedAddress returns the OriginatorReceivedAddress field value if set, zero value otherwise.
func (o *OriginatorInfo) GetOriginatorReceivedAddress() SMAddressInfo {
	if o == nil || IsNil(o.OriginatorReceivedAddress) {
		var ret SMAddressInfo
		return ret
	}
	return *o.OriginatorReceivedAddress
}

// GetOriginatorReceivedAddressOk returns a tuple with the OriginatorReceivedAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginatorInfo) GetOriginatorReceivedAddressOk() (*SMAddressInfo, bool) {
	if o == nil || IsNil(o.OriginatorReceivedAddress) {
		return nil, false
	}
	return o.OriginatorReceivedAddress, true
}

// HasOriginatorReceivedAddress returns a boolean if a field has been set.
func (o *OriginatorInfo) HasOriginatorReceivedAddress() bool {
	if o != nil && !IsNil(o.OriginatorReceivedAddress) {
		return true
	}

	return false
}

// SetOriginatorReceivedAddress gets a reference to the given SMAddressInfo and assigns it to the OriginatorReceivedAddress field.
func (o *OriginatorInfo) SetOriginatorReceivedAddress(v SMAddressInfo) {
	o.OriginatorReceivedAddress = &v
}

// GetOriginatorSCCPAddress returns the OriginatorSCCPAddress field value if set, zero value otherwise.
func (o *OriginatorInfo) GetOriginatorSCCPAddress() string {
	if o == nil || IsNil(o.OriginatorSCCPAddress) {
		var ret string
		return ret
	}
	return *o.OriginatorSCCPAddress
}

// GetOriginatorSCCPAddressOk returns a tuple with the OriginatorSCCPAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginatorInfo) GetOriginatorSCCPAddressOk() (*string, bool) {
	if o == nil || IsNil(o.OriginatorSCCPAddress) {
		return nil, false
	}
	return o.OriginatorSCCPAddress, true
}

// HasOriginatorSCCPAddress returns a boolean if a field has been set.
func (o *OriginatorInfo) HasOriginatorSCCPAddress() bool {
	if o != nil && !IsNil(o.OriginatorSCCPAddress) {
		return true
	}

	return false
}

// SetOriginatorSCCPAddress gets a reference to the given string and assigns it to the OriginatorSCCPAddress field.
func (o *OriginatorInfo) SetOriginatorSCCPAddress(v string) {
	o.OriginatorSCCPAddress = &v
}

// GetSMOriginatorInterface returns the SMOriginatorInterface field value if set, zero value otherwise.
func (o *OriginatorInfo) GetSMOriginatorInterface() SMInterface {
	if o == nil || IsNil(o.SMOriginatorInterface) {
		var ret SMInterface
		return ret
	}
	return *o.SMOriginatorInterface
}

// GetSMOriginatorInterfaceOk returns a tuple with the SMOriginatorInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginatorInfo) GetSMOriginatorInterfaceOk() (*SMInterface, bool) {
	if o == nil || IsNil(o.SMOriginatorInterface) {
		return nil, false
	}
	return o.SMOriginatorInterface, true
}

// HasSMOriginatorInterface returns a boolean if a field has been set.
func (o *OriginatorInfo) HasSMOriginatorInterface() bool {
	if o != nil && !IsNil(o.SMOriginatorInterface) {
		return true
	}

	return false
}

// SetSMOriginatorInterface gets a reference to the given SMInterface and assigns it to the SMOriginatorInterface field.
func (o *OriginatorInfo) SetSMOriginatorInterface(v SMInterface) {
	o.SMOriginatorInterface = &v
}

// GetSMOriginatorProtocolId returns the SMOriginatorProtocolId field value if set, zero value otherwise.
func (o *OriginatorInfo) GetSMOriginatorProtocolId() string {
	if o == nil || IsNil(o.SMOriginatorProtocolId) {
		var ret string
		return ret
	}
	return *o.SMOriginatorProtocolId
}

// GetSMOriginatorProtocolIdOk returns a tuple with the SMOriginatorProtocolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OriginatorInfo) GetSMOriginatorProtocolIdOk() (*string, bool) {
	if o == nil || IsNil(o.SMOriginatorProtocolId) {
		return nil, false
	}
	return o.SMOriginatorProtocolId, true
}

// HasSMOriginatorProtocolId returns a boolean if a field has been set.
func (o *OriginatorInfo) HasSMOriginatorProtocolId() bool {
	if o != nil && !IsNil(o.SMOriginatorProtocolId) {
		return true
	}

	return false
}

// SetSMOriginatorProtocolId gets a reference to the given string and assigns it to the SMOriginatorProtocolId field.
func (o *OriginatorInfo) SetSMOriginatorProtocolId(v string) {
	o.SMOriginatorProtocolId = &v
}

func (o OriginatorInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OriginatorInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OriginatorSUPI) {
		toSerialize["originatorSUPI"] = o.OriginatorSUPI
	}
	if !IsNil(o.OriginatorGPSI) {
		toSerialize["originatorGPSI"] = o.OriginatorGPSI
	}
	if !IsNil(o.OriginatorOtherAddress) {
		toSerialize["originatorOtherAddress"] = o.OriginatorOtherAddress
	}
	if !IsNil(o.OriginatorReceivedAddress) {
		toSerialize["originatorReceivedAddress"] = o.OriginatorReceivedAddress
	}
	if !IsNil(o.OriginatorSCCPAddress) {
		toSerialize["originatorSCCPAddress"] = o.OriginatorSCCPAddress
	}
	if !IsNil(o.SMOriginatorInterface) {
		toSerialize["sMOriginatorInterface"] = o.SMOriginatorInterface
	}
	if !IsNil(o.SMOriginatorProtocolId) {
		toSerialize["sMOriginatorProtocolId"] = o.SMOriginatorProtocolId
	}
	return toSerialize, nil
}

type NullableOriginatorInfo struct {
	value *OriginatorInfo
	isSet bool
}

func (v NullableOriginatorInfo) Get() *OriginatorInfo {
	return v.value
}

func (v *NullableOriginatorInfo) Set(val *OriginatorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginatorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginatorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginatorInfo(val *OriginatorInfo) *NullableOriginatorInfo {
	return &NullableOriginatorInfo{value: val, isSet: true}
}

func (v NullableOriginatorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginatorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
