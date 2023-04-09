/*
3gpp-racs-parameter-provisioning

API for provisioning UE radio capability parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_RacsParameterProvisioning

import (
	"encoding/json"
)

// checks if the RacsProvisioningData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RacsProvisioningData{}

// RacsProvisioningData Represents a UE's radio capability data.
type RacsProvisioningData struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// Identifies the configuration related to manufacturer specific UE radio capability. Each element uniquely identifies an RACS configuration for an RACS ID and is identified in the map via the RACS ID as key. The response shall include successfully provisioned RACS data.
	RacsConfigs map[string]RacsConfiguration `json:"racsConfigs"`
	// Supplied by the SCEF. Contains the RACS IDs for which the RACS data are not provisioned successfully. Any string value can be used as a key of the map.
	RacsReports *map[string]RacsFailureReport `json:"racsReports,omitempty"`
}

// NewRacsProvisioningData instantiates a new RacsProvisioningData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRacsProvisioningData(racsConfigs map[string]RacsConfiguration) *RacsProvisioningData {
	this := RacsProvisioningData{}
	this.RacsConfigs = racsConfigs
	return &this
}

// NewRacsProvisioningDataWithDefaults instantiates a new RacsProvisioningData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRacsProvisioningDataWithDefaults() *RacsProvisioningData {
	this := RacsProvisioningData{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *RacsProvisioningData) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RacsProvisioningData) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *RacsProvisioningData) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *RacsProvisioningData) SetSelf(v string) {
	o.Self = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *RacsProvisioningData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RacsProvisioningData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *RacsProvisioningData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *RacsProvisioningData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetRacsConfigs returns the RacsConfigs field value
func (o *RacsProvisioningData) GetRacsConfigs() map[string]RacsConfiguration {
	if o == nil {
		var ret map[string]RacsConfiguration
		return ret
	}

	return o.RacsConfigs
}

// GetRacsConfigsOk returns a tuple with the RacsConfigs field value
// and a boolean to check if the value has been set.
func (o *RacsProvisioningData) GetRacsConfigsOk() (*map[string]RacsConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RacsConfigs, true
}

// SetRacsConfigs sets field value
func (o *RacsProvisioningData) SetRacsConfigs(v map[string]RacsConfiguration) {
	o.RacsConfigs = v
}

// GetRacsReports returns the RacsReports field value if set, zero value otherwise.
func (o *RacsProvisioningData) GetRacsReports() map[string]RacsFailureReport {
	if o == nil || IsNil(o.RacsReports) {
		var ret map[string]RacsFailureReport
		return ret
	}
	return *o.RacsReports
}

// GetRacsReportsOk returns a tuple with the RacsReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RacsProvisioningData) GetRacsReportsOk() (*map[string]RacsFailureReport, bool) {
	if o == nil || IsNil(o.RacsReports) {
		return nil, false
	}
	return o.RacsReports, true
}

// HasRacsReports returns a boolean if a field has been set.
func (o *RacsProvisioningData) HasRacsReports() bool {
	if o != nil && !IsNil(o.RacsReports) {
		return true
	}

	return false
}

// SetRacsReports gets a reference to the given map[string]RacsFailureReport and assigns it to the RacsReports field.
func (o *RacsProvisioningData) SetRacsReports(v map[string]RacsFailureReport) {
	o.RacsReports = &v
}

func (o RacsProvisioningData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RacsProvisioningData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	toSerialize["racsConfigs"] = o.RacsConfigs
	// skip: racsReports is readOnly
	return toSerialize, nil
}

type NullableRacsProvisioningData struct {
	value *RacsProvisioningData
	isSet bool
}

func (v NullableRacsProvisioningData) Get() *RacsProvisioningData {
	return v.value
}

func (v *NullableRacsProvisioningData) Set(val *RacsProvisioningData) {
	v.value = val
	v.isSet = true
}

func (v NullableRacsProvisioningData) IsSet() bool {
	return v.isSet
}

func (v *NullableRacsProvisioningData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRacsProvisioningData(val *RacsProvisioningData) *NullableRacsProvisioningData {
	return &NullableRacsProvisioningData{value: val, isSet: true}
}

func (v NullableRacsProvisioningData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRacsProvisioningData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
