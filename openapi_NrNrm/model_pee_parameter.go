/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
)

// checks if the PeeParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PeeParameter{}

// PeeParameter struct for PeeParameter
type PeeParameter struct {
	SiteIdentification *string  `json:"siteIdentification,omitempty"`
	SiteDescription    *string  `json:"siteDescription,omitempty"`
	SiteLatitude       *float32 `json:"siteLatitude,omitempty"`
	SiteLongitude      *float32 `json:"siteLongitude,omitempty"`
	SiteAltitude       *float32 `json:"siteAltitude,omitempty"`
	EquipmentType      *string  `json:"equipmentType,omitempty"`
	EnvironmentType    *string  `json:"environmentType,omitempty"`
	PowerInterface     *string  `json:"powerInterface,omitempty"`
}

// NewPeeParameter instantiates a new PeeParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPeeParameter() *PeeParameter {
	this := PeeParameter{}
	return &this
}

// NewPeeParameterWithDefaults instantiates a new PeeParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPeeParameterWithDefaults() *PeeParameter {
	this := PeeParameter{}
	return &this
}

// GetSiteIdentification returns the SiteIdentification field value if set, zero value otherwise.
func (o *PeeParameter) GetSiteIdentification() string {
	if o == nil || IsNil(o.SiteIdentification) {
		var ret string
		return ret
	}
	return *o.SiteIdentification
}

// GetSiteIdentificationOk returns a tuple with the SiteIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeeParameter) GetSiteIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.SiteIdentification) {
		return nil, false
	}
	return o.SiteIdentification, true
}

// HasSiteIdentification returns a boolean if a field has been set.
func (o *PeeParameter) HasSiteIdentification() bool {
	if o != nil && !IsNil(o.SiteIdentification) {
		return true
	}

	return false
}

// SetSiteIdentification gets a reference to the given string and assigns it to the SiteIdentification field.
func (o *PeeParameter) SetSiteIdentification(v string) {
	o.SiteIdentification = &v
}

// GetSiteDescription returns the SiteDescription field value if set, zero value otherwise.
func (o *PeeParameter) GetSiteDescription() string {
	if o == nil || IsNil(o.SiteDescription) {
		var ret string
		return ret
	}
	return *o.SiteDescription
}

// GetSiteDescriptionOk returns a tuple with the SiteDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeeParameter) GetSiteDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.SiteDescription) {
		return nil, false
	}
	return o.SiteDescription, true
}

// HasSiteDescription returns a boolean if a field has been set.
func (o *PeeParameter) HasSiteDescription() bool {
	if o != nil && !IsNil(o.SiteDescription) {
		return true
	}

	return false
}

// SetSiteDescription gets a reference to the given string and assigns it to the SiteDescription field.
func (o *PeeParameter) SetSiteDescription(v string) {
	o.SiteDescription = &v
}

// GetSiteLatitude returns the SiteLatitude field value if set, zero value otherwise.
func (o *PeeParameter) GetSiteLatitude() float32 {
	if o == nil || IsNil(o.SiteLatitude) {
		var ret float32
		return ret
	}
	return *o.SiteLatitude
}

// GetSiteLatitudeOk returns a tuple with the SiteLatitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeeParameter) GetSiteLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.SiteLatitude) {
		return nil, false
	}
	return o.SiteLatitude, true
}

// HasSiteLatitude returns a boolean if a field has been set.
func (o *PeeParameter) HasSiteLatitude() bool {
	if o != nil && !IsNil(o.SiteLatitude) {
		return true
	}

	return false
}

// SetSiteLatitude gets a reference to the given float32 and assigns it to the SiteLatitude field.
func (o *PeeParameter) SetSiteLatitude(v float32) {
	o.SiteLatitude = &v
}

// GetSiteLongitude returns the SiteLongitude field value if set, zero value otherwise.
func (o *PeeParameter) GetSiteLongitude() float32 {
	if o == nil || IsNil(o.SiteLongitude) {
		var ret float32
		return ret
	}
	return *o.SiteLongitude
}

// GetSiteLongitudeOk returns a tuple with the SiteLongitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeeParameter) GetSiteLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.SiteLongitude) {
		return nil, false
	}
	return o.SiteLongitude, true
}

// HasSiteLongitude returns a boolean if a field has been set.
func (o *PeeParameter) HasSiteLongitude() bool {
	if o != nil && !IsNil(o.SiteLongitude) {
		return true
	}

	return false
}

// SetSiteLongitude gets a reference to the given float32 and assigns it to the SiteLongitude field.
func (o *PeeParameter) SetSiteLongitude(v float32) {
	o.SiteLongitude = &v
}

// GetSiteAltitude returns the SiteAltitude field value if set, zero value otherwise.
func (o *PeeParameter) GetSiteAltitude() float32 {
	if o == nil || IsNil(o.SiteAltitude) {
		var ret float32
		return ret
	}
	return *o.SiteAltitude
}

// GetSiteAltitudeOk returns a tuple with the SiteAltitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeeParameter) GetSiteAltitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.SiteAltitude) {
		return nil, false
	}
	return o.SiteAltitude, true
}

// HasSiteAltitude returns a boolean if a field has been set.
func (o *PeeParameter) HasSiteAltitude() bool {
	if o != nil && !IsNil(o.SiteAltitude) {
		return true
	}

	return false
}

// SetSiteAltitude gets a reference to the given float32 and assigns it to the SiteAltitude field.
func (o *PeeParameter) SetSiteAltitude(v float32) {
	o.SiteAltitude = &v
}

// GetEquipmentType returns the EquipmentType field value if set, zero value otherwise.
func (o *PeeParameter) GetEquipmentType() string {
	if o == nil || IsNil(o.EquipmentType) {
		var ret string
		return ret
	}
	return *o.EquipmentType
}

// GetEquipmentTypeOk returns a tuple with the EquipmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeeParameter) GetEquipmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EquipmentType) {
		return nil, false
	}
	return o.EquipmentType, true
}

// HasEquipmentType returns a boolean if a field has been set.
func (o *PeeParameter) HasEquipmentType() bool {
	if o != nil && !IsNil(o.EquipmentType) {
		return true
	}

	return false
}

// SetEquipmentType gets a reference to the given string and assigns it to the EquipmentType field.
func (o *PeeParameter) SetEquipmentType(v string) {
	o.EquipmentType = &v
}

// GetEnvironmentType returns the EnvironmentType field value if set, zero value otherwise.
func (o *PeeParameter) GetEnvironmentType() string {
	if o == nil || IsNil(o.EnvironmentType) {
		var ret string
		return ret
	}
	return *o.EnvironmentType
}

// GetEnvironmentTypeOk returns a tuple with the EnvironmentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeeParameter) GetEnvironmentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentType) {
		return nil, false
	}
	return o.EnvironmentType, true
}

// HasEnvironmentType returns a boolean if a field has been set.
func (o *PeeParameter) HasEnvironmentType() bool {
	if o != nil && !IsNil(o.EnvironmentType) {
		return true
	}

	return false
}

// SetEnvironmentType gets a reference to the given string and assigns it to the EnvironmentType field.
func (o *PeeParameter) SetEnvironmentType(v string) {
	o.EnvironmentType = &v
}

// GetPowerInterface returns the PowerInterface field value if set, zero value otherwise.
func (o *PeeParameter) GetPowerInterface() string {
	if o == nil || IsNil(o.PowerInterface) {
		var ret string
		return ret
	}
	return *o.PowerInterface
}

// GetPowerInterfaceOk returns a tuple with the PowerInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PeeParameter) GetPowerInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.PowerInterface) {
		return nil, false
	}
	return o.PowerInterface, true
}

// HasPowerInterface returns a boolean if a field has been set.
func (o *PeeParameter) HasPowerInterface() bool {
	if o != nil && !IsNil(o.PowerInterface) {
		return true
	}

	return false
}

// SetPowerInterface gets a reference to the given string and assigns it to the PowerInterface field.
func (o *PeeParameter) SetPowerInterface(v string) {
	o.PowerInterface = &v
}

func (o PeeParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PeeParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SiteIdentification) {
		toSerialize["siteIdentification"] = o.SiteIdentification
	}
	if !IsNil(o.SiteDescription) {
		toSerialize["siteDescription"] = o.SiteDescription
	}
	if !IsNil(o.SiteLatitude) {
		toSerialize["siteLatitude"] = o.SiteLatitude
	}
	if !IsNil(o.SiteLongitude) {
		toSerialize["siteLongitude"] = o.SiteLongitude
	}
	if !IsNil(o.SiteAltitude) {
		toSerialize["siteAltitude"] = o.SiteAltitude
	}
	if !IsNil(o.EquipmentType) {
		toSerialize["equipmentType"] = o.EquipmentType
	}
	if !IsNil(o.EnvironmentType) {
		toSerialize["environmentType"] = o.EnvironmentType
	}
	if !IsNil(o.PowerInterface) {
		toSerialize["powerInterface"] = o.PowerInterface
	}
	return toSerialize, nil
}

type NullablePeeParameter struct {
	value *PeeParameter
	isSet bool
}

func (v NullablePeeParameter) Get() *PeeParameter {
	return v.value
}

func (v *NullablePeeParameter) Set(val *PeeParameter) {
	v.value = val
	v.isSet = true
}

func (v NullablePeeParameter) IsSet() bool {
	return v.isSet
}

func (v *NullablePeeParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeeParameter(val *PeeParameter) *NullablePeeParameter {
	return &NullablePeeParameter{value: val, isSet: true}
}

func (v NullablePeeParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeeParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
