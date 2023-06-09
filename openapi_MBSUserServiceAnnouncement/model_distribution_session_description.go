/*
MBS User Service Announcement Element units’ definition

MBS User Service Announcement Element units. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserServiceAnnouncement

import (
	"encoding/json"
)

// checks if the DistributionSessionDescription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributionSessionDescription{}

// DistributionSessionDescription struct for DistributionSessionDescription
type DistributionSessionDescription struct {
	ConformanceProfile *string `json:"conformanceProfile,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	SessionDescriptionURI  string                                                  `json:"sessionDescriptionURI"`
	ObjectRepairParameters *AssociatedProcedureDescription                         `json:"objectRepairParameters,omitempty"`
	DataNetworkName        *string                                                 `json:"dataNetworkName,omitempty"`
	MbsAppService          []ApplicationService                                    `json:"mbsAppService,omitempty"`
	UnicastAppServices     []DistributionSessionDescriptionUnicastAppServicesInner `json:"unicastAppServices,omitempty"`
}

// NewDistributionSessionDescription instantiates a new DistributionSessionDescription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributionSessionDescription(sessionDescriptionURI string) *DistributionSessionDescription {
	this := DistributionSessionDescription{}
	this.SessionDescriptionURI = sessionDescriptionURI
	return &this
}

// NewDistributionSessionDescriptionWithDefaults instantiates a new DistributionSessionDescription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributionSessionDescriptionWithDefaults() *DistributionSessionDescription {
	this := DistributionSessionDescription{}
	return &this
}

// GetConformanceProfile returns the ConformanceProfile field value if set, zero value otherwise.
func (o *DistributionSessionDescription) GetConformanceProfile() string {
	if o == nil || IsNil(o.ConformanceProfile) {
		var ret string
		return ret
	}
	return *o.ConformanceProfile
}

// GetConformanceProfileOk returns a tuple with the ConformanceProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionSessionDescription) GetConformanceProfileOk() (*string, bool) {
	if o == nil || IsNil(o.ConformanceProfile) {
		return nil, false
	}
	return o.ConformanceProfile, true
}

// HasConformanceProfile returns a boolean if a field has been set.
func (o *DistributionSessionDescription) HasConformanceProfile() bool {
	if o != nil && !IsNil(o.ConformanceProfile) {
		return true
	}

	return false
}

// SetConformanceProfile gets a reference to the given string and assigns it to the ConformanceProfile field.
func (o *DistributionSessionDescription) SetConformanceProfile(v string) {
	o.ConformanceProfile = &v
}

// GetSessionDescriptionURI returns the SessionDescriptionURI field value
func (o *DistributionSessionDescription) GetSessionDescriptionURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionDescriptionURI
}

// GetSessionDescriptionURIOk returns a tuple with the SessionDescriptionURI field value
// and a boolean to check if the value has been set.
func (o *DistributionSessionDescription) GetSessionDescriptionURIOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionDescriptionURI, true
}

// SetSessionDescriptionURI sets field value
func (o *DistributionSessionDescription) SetSessionDescriptionURI(v string) {
	o.SessionDescriptionURI = v
}

// GetObjectRepairParameters returns the ObjectRepairParameters field value if set, zero value otherwise.
func (o *DistributionSessionDescription) GetObjectRepairParameters() AssociatedProcedureDescription {
	if o == nil || IsNil(o.ObjectRepairParameters) {
		var ret AssociatedProcedureDescription
		return ret
	}
	return *o.ObjectRepairParameters
}

// GetObjectRepairParametersOk returns a tuple with the ObjectRepairParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionSessionDescription) GetObjectRepairParametersOk() (*AssociatedProcedureDescription, bool) {
	if o == nil || IsNil(o.ObjectRepairParameters) {
		return nil, false
	}
	return o.ObjectRepairParameters, true
}

// HasObjectRepairParameters returns a boolean if a field has been set.
func (o *DistributionSessionDescription) HasObjectRepairParameters() bool {
	if o != nil && !IsNil(o.ObjectRepairParameters) {
		return true
	}

	return false
}

// SetObjectRepairParameters gets a reference to the given AssociatedProcedureDescription and assigns it to the ObjectRepairParameters field.
func (o *DistributionSessionDescription) SetObjectRepairParameters(v AssociatedProcedureDescription) {
	o.ObjectRepairParameters = &v
}

// GetDataNetworkName returns the DataNetworkName field value if set, zero value otherwise.
func (o *DistributionSessionDescription) GetDataNetworkName() string {
	if o == nil || IsNil(o.DataNetworkName) {
		var ret string
		return ret
	}
	return *o.DataNetworkName
}

// GetDataNetworkNameOk returns a tuple with the DataNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionSessionDescription) GetDataNetworkNameOk() (*string, bool) {
	if o == nil || IsNil(o.DataNetworkName) {
		return nil, false
	}
	return o.DataNetworkName, true
}

// HasDataNetworkName returns a boolean if a field has been set.
func (o *DistributionSessionDescription) HasDataNetworkName() bool {
	if o != nil && !IsNil(o.DataNetworkName) {
		return true
	}

	return false
}

// SetDataNetworkName gets a reference to the given string and assigns it to the DataNetworkName field.
func (o *DistributionSessionDescription) SetDataNetworkName(v string) {
	o.DataNetworkName = &v
}

// GetMbsAppService returns the MbsAppService field value if set, zero value otherwise.
func (o *DistributionSessionDescription) GetMbsAppService() []ApplicationService {
	if o == nil || IsNil(o.MbsAppService) {
		var ret []ApplicationService
		return ret
	}
	return o.MbsAppService
}

// GetMbsAppServiceOk returns a tuple with the MbsAppService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionSessionDescription) GetMbsAppServiceOk() ([]ApplicationService, bool) {
	if o == nil || IsNil(o.MbsAppService) {
		return nil, false
	}
	return o.MbsAppService, true
}

// HasMbsAppService returns a boolean if a field has been set.
func (o *DistributionSessionDescription) HasMbsAppService() bool {
	if o != nil && !IsNil(o.MbsAppService) {
		return true
	}

	return false
}

// SetMbsAppService gets a reference to the given []ApplicationService and assigns it to the MbsAppService field.
func (o *DistributionSessionDescription) SetMbsAppService(v []ApplicationService) {
	o.MbsAppService = v
}

// GetUnicastAppServices returns the UnicastAppServices field value if set, zero value otherwise.
func (o *DistributionSessionDescription) GetUnicastAppServices() []DistributionSessionDescriptionUnicastAppServicesInner {
	if o == nil || IsNil(o.UnicastAppServices) {
		var ret []DistributionSessionDescriptionUnicastAppServicesInner
		return ret
	}
	return o.UnicastAppServices
}

// GetUnicastAppServicesOk returns a tuple with the UnicastAppServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DistributionSessionDescription) GetUnicastAppServicesOk() ([]DistributionSessionDescriptionUnicastAppServicesInner, bool) {
	if o == nil || IsNil(o.UnicastAppServices) {
		return nil, false
	}
	return o.UnicastAppServices, true
}

// HasUnicastAppServices returns a boolean if a field has been set.
func (o *DistributionSessionDescription) HasUnicastAppServices() bool {
	if o != nil && !IsNil(o.UnicastAppServices) {
		return true
	}

	return false
}

// SetUnicastAppServices gets a reference to the given []DistributionSessionDescriptionUnicastAppServicesInner and assigns it to the UnicastAppServices field.
func (o *DistributionSessionDescription) SetUnicastAppServices(v []DistributionSessionDescriptionUnicastAppServicesInner) {
	o.UnicastAppServices = v
}

func (o DistributionSessionDescription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistributionSessionDescription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConformanceProfile) {
		toSerialize["conformanceProfile"] = o.ConformanceProfile
	}
	toSerialize["sessionDescriptionURI"] = o.SessionDescriptionURI
	if !IsNil(o.ObjectRepairParameters) {
		toSerialize["objectRepairParameters"] = o.ObjectRepairParameters
	}
	if !IsNil(o.DataNetworkName) {
		toSerialize["dataNetworkName"] = o.DataNetworkName
	}
	if !IsNil(o.MbsAppService) {
		toSerialize["mbsAppService"] = o.MbsAppService
	}
	if !IsNil(o.UnicastAppServices) {
		toSerialize["unicastAppServices"] = o.UnicastAppServices
	}
	return toSerialize, nil
}

type NullableDistributionSessionDescription struct {
	value *DistributionSessionDescription
	isSet bool
}

func (v NullableDistributionSessionDescription) Get() *DistributionSessionDescription {
	return v.value
}

func (v *NullableDistributionSessionDescription) Set(val *DistributionSessionDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributionSessionDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributionSessionDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributionSessionDescription(val *DistributionSessionDescription) *NullableDistributionSessionDescription {
	return &NullableDistributionSessionDescription{value: val, isSet: true}
}

func (v NullableDistributionSessionDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributionSessionDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
