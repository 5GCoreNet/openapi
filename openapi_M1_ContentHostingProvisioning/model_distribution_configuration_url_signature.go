/*
M1_ContentHostingProvisioning

5GMS AF M1 Content Hosting Provisioning API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ContentHostingProvisioning

import (
	"encoding/json"
)

// checks if the DistributionConfigurationUrlSignature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributionConfigurationUrlSignature{}

// DistributionConfigurationUrlSignature struct for DistributionConfigurationUrlSignature
type DistributionConfigurationUrlSignature struct {
	UrlPattern string `json:"urlPattern"`
	TokenName string `json:"tokenName"`
	PassphraseName string `json:"passphraseName"`
	Passphrase string `json:"passphrase"`
	TokenExpiryName string `json:"tokenExpiryName"`
	UseIPAddress bool `json:"useIPAddress"`
	IpAddressName *string `json:"ipAddressName,omitempty"`
}

// NewDistributionConfigurationUrlSignature instantiates a new DistributionConfigurationUrlSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionConfigurationUrlSignature(urlPattern string, tokenName string, passphraseName string, passphrase string, tokenExpiryName string, useIPAddress bool) *DistributionConfigurationUrlSignature {
	this := DistributionConfigurationUrlSignature{}
	this.UrlPattern = urlPattern
	this.TokenName = tokenName
	this.PassphraseName = passphraseName
	this.Passphrase = passphrase
	this.TokenExpiryName = tokenExpiryName
	this.UseIPAddress = useIPAddress
	return &this
}

// NewDistributionConfigurationUrlSignatureWithDefaults instantiates a new DistributionConfigurationUrlSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionConfigurationUrlSignatureWithDefaults() *DistributionConfigurationUrlSignature {
	this := DistributionConfigurationUrlSignature{}
	return &this
}

// GetUrlPattern returns the UrlPattern field value
func (o *DistributionConfigurationUrlSignature) GetUrlPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlPattern
}

// GetUrlPatternOk returns a tuple with the UrlPattern field value
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationUrlSignature) GetUrlPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlPattern, true
}

// SetUrlPattern sets field value
func (o *DistributionConfigurationUrlSignature) SetUrlPattern(v string) {
	o.UrlPattern = v
}

// GetTokenName returns the TokenName field value
func (o *DistributionConfigurationUrlSignature) GetTokenName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenName
}

// GetTokenNameOk returns a tuple with the TokenName field value
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationUrlSignature) GetTokenNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenName, true
}

// SetTokenName sets field value
func (o *DistributionConfigurationUrlSignature) SetTokenName(v string) {
	o.TokenName = v
}

// GetPassphraseName returns the PassphraseName field value
func (o *DistributionConfigurationUrlSignature) GetPassphraseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PassphraseName
}

// GetPassphraseNameOk returns a tuple with the PassphraseName field value
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationUrlSignature) GetPassphraseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PassphraseName, true
}

// SetPassphraseName sets field value
func (o *DistributionConfigurationUrlSignature) SetPassphraseName(v string) {
	o.PassphraseName = v
}

// GetPassphrase returns the Passphrase field value
func (o *DistributionConfigurationUrlSignature) GetPassphrase() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Passphrase
}

// GetPassphraseOk returns a tuple with the Passphrase field value
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationUrlSignature) GetPassphraseOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Passphrase, true
}

// SetPassphrase sets field value
func (o *DistributionConfigurationUrlSignature) SetPassphrase(v string) {
	o.Passphrase = v
}

// GetTokenExpiryName returns the TokenExpiryName field value
func (o *DistributionConfigurationUrlSignature) GetTokenExpiryName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenExpiryName
}

// GetTokenExpiryNameOk returns a tuple with the TokenExpiryName field value
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationUrlSignature) GetTokenExpiryNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenExpiryName, true
}

// SetTokenExpiryName sets field value
func (o *DistributionConfigurationUrlSignature) SetTokenExpiryName(v string) {
	o.TokenExpiryName = v
}

// GetUseIPAddress returns the UseIPAddress field value
func (o *DistributionConfigurationUrlSignature) GetUseIPAddress() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseIPAddress
}

// GetUseIPAddressOk returns a tuple with the UseIPAddress field value
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationUrlSignature) GetUseIPAddressOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseIPAddress, true
}

// SetUseIPAddress sets field value
func (o *DistributionConfigurationUrlSignature) SetUseIPAddress(v bool) {
	o.UseIPAddress = v
}

// GetIpAddressName returns the IpAddressName field value if set, zero value otherwise.
func (o *DistributionConfigurationUrlSignature) GetIpAddressName() string {
	if o == nil || IsNil(o.IpAddressName) {
		var ret string
		return ret
	}
	return *o.IpAddressName
}

// GetIpAddressNameOk returns a tuple with the IpAddressName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionConfigurationUrlSignature) GetIpAddressNameOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddressName) {
		return nil, false
	}
	return o.IpAddressName, true
}

// HasIpAddressName returns a boolean if a field has been set.
func (o *DistributionConfigurationUrlSignature) HasIpAddressName() bool {
	if o != nil && !IsNil(o.IpAddressName) {
		return true
	}

	return false
}

// SetIpAddressName gets a reference to the given string and assigns it to the IpAddressName field.
func (o *DistributionConfigurationUrlSignature) SetIpAddressName(v string) {
	o.IpAddressName = &v
}

func (o DistributionConfigurationUrlSignature) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistributionConfigurationUrlSignature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["urlPattern"] = o.UrlPattern
	toSerialize["tokenName"] = o.TokenName
	toSerialize["passphraseName"] = o.PassphraseName
	toSerialize["passphrase"] = o.Passphrase
	toSerialize["tokenExpiryName"] = o.TokenExpiryName
	toSerialize["useIPAddress"] = o.UseIPAddress
	if !IsNil(o.IpAddressName) {
		toSerialize["ipAddressName"] = o.IpAddressName
	}
	return toSerialize, nil
}

type NullableDistributionConfigurationUrlSignature struct {
	value *DistributionConfigurationUrlSignature
	isSet bool
}

func (v NullableDistributionConfigurationUrlSignature) Get() *DistributionConfigurationUrlSignature {
	return v.value
}

func (v *NullableDistributionConfigurationUrlSignature) Set(val *DistributionConfigurationUrlSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionConfigurationUrlSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionConfigurationUrlSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionConfigurationUrlSignature(val *DistributionConfigurationUrlSignature) *NullableDistributionConfigurationUrlSignature {
	return &NullableDistributionConfigurationUrlSignature{value: val, isSet: true}
}

func (v NullableDistributionConfigurationUrlSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionConfigurationUrlSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


