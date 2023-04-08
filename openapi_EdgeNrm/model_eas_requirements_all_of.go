/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
)

// checks if the EASRequirementsAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EASRequirementsAllOf{}

// EASRequirementsAllOf struct for EASRequirementsAllOf
type EASRequirementsAllOf struct {
	RequiredEASservingLocation *ServingLocation `json:"requiredEASservingLocation,omitempty"`
	AffinityAntiAffinity *AffinityAntiAffinity `json:"affinityAntiAffinity,omitempty"`
	ServiceContinuity *bool `json:"serviceContinuity,omitempty"`
	VirtualResource *VirtualResource `json:"virtualResource,omitempty"`
	SoftwareImageInfo *SoftwareImageInfo `json:"softwareImageInfo,omitempty"`
}

// NewEASRequirementsAllOf instantiates a new EASRequirementsAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEASRequirementsAllOf() *EASRequirementsAllOf {
	this := EASRequirementsAllOf{}
	return &this
}

// NewEASRequirementsAllOfWithDefaults instantiates a new EASRequirementsAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEASRequirementsAllOfWithDefaults() *EASRequirementsAllOf {
	this := EASRequirementsAllOf{}
	return &this
}

// GetRequiredEASservingLocation returns the RequiredEASservingLocation field value if set, zero value otherwise.
func (o *EASRequirementsAllOf) GetRequiredEASservingLocation() ServingLocation {
	if o == nil || isNil(o.RequiredEASservingLocation) {
		var ret ServingLocation
		return ret
	}
	return *o.RequiredEASservingLocation
}

// GetRequiredEASservingLocationOk returns a tuple with the RequiredEASservingLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirementsAllOf) GetRequiredEASservingLocationOk() (*ServingLocation, bool) {
	if o == nil || isNil(o.RequiredEASservingLocation) {
		return nil, false
	}
	return o.RequiredEASservingLocation, true
}

// HasRequiredEASservingLocation returns a boolean if a field has been set.
func (o *EASRequirementsAllOf) HasRequiredEASservingLocation() bool {
	if o != nil && !isNil(o.RequiredEASservingLocation) {
		return true
	}

	return false
}

// SetRequiredEASservingLocation gets a reference to the given ServingLocation and assigns it to the RequiredEASservingLocation field.
func (o *EASRequirementsAllOf) SetRequiredEASservingLocation(v ServingLocation) {
	o.RequiredEASservingLocation = &v
}

// GetAffinityAntiAffinity returns the AffinityAntiAffinity field value if set, zero value otherwise.
func (o *EASRequirementsAllOf) GetAffinityAntiAffinity() AffinityAntiAffinity {
	if o == nil || isNil(o.AffinityAntiAffinity) {
		var ret AffinityAntiAffinity
		return ret
	}
	return *o.AffinityAntiAffinity
}

// GetAffinityAntiAffinityOk returns a tuple with the AffinityAntiAffinity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirementsAllOf) GetAffinityAntiAffinityOk() (*AffinityAntiAffinity, bool) {
	if o == nil || isNil(o.AffinityAntiAffinity) {
		return nil, false
	}
	return o.AffinityAntiAffinity, true
}

// HasAffinityAntiAffinity returns a boolean if a field has been set.
func (o *EASRequirementsAllOf) HasAffinityAntiAffinity() bool {
	if o != nil && !isNil(o.AffinityAntiAffinity) {
		return true
	}

	return false
}

// SetAffinityAntiAffinity gets a reference to the given AffinityAntiAffinity and assigns it to the AffinityAntiAffinity field.
func (o *EASRequirementsAllOf) SetAffinityAntiAffinity(v AffinityAntiAffinity) {
	o.AffinityAntiAffinity = &v
}

// GetServiceContinuity returns the ServiceContinuity field value if set, zero value otherwise.
func (o *EASRequirementsAllOf) GetServiceContinuity() bool {
	if o == nil || isNil(o.ServiceContinuity) {
		var ret bool
		return ret
	}
	return *o.ServiceContinuity
}

// GetServiceContinuityOk returns a tuple with the ServiceContinuity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirementsAllOf) GetServiceContinuityOk() (*bool, bool) {
	if o == nil || isNil(o.ServiceContinuity) {
		return nil, false
	}
	return o.ServiceContinuity, true
}

// HasServiceContinuity returns a boolean if a field has been set.
func (o *EASRequirementsAllOf) HasServiceContinuity() bool {
	if o != nil && !isNil(o.ServiceContinuity) {
		return true
	}

	return false
}

// SetServiceContinuity gets a reference to the given bool and assigns it to the ServiceContinuity field.
func (o *EASRequirementsAllOf) SetServiceContinuity(v bool) {
	o.ServiceContinuity = &v
}

// GetVirtualResource returns the VirtualResource field value if set, zero value otherwise.
func (o *EASRequirementsAllOf) GetVirtualResource() VirtualResource {
	if o == nil || isNil(o.VirtualResource) {
		var ret VirtualResource
		return ret
	}
	return *o.VirtualResource
}

// GetVirtualResourceOk returns a tuple with the VirtualResource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirementsAllOf) GetVirtualResourceOk() (*VirtualResource, bool) {
	if o == nil || isNil(o.VirtualResource) {
		return nil, false
	}
	return o.VirtualResource, true
}

// HasVirtualResource returns a boolean if a field has been set.
func (o *EASRequirementsAllOf) HasVirtualResource() bool {
	if o != nil && !isNil(o.VirtualResource) {
		return true
	}

	return false
}

// SetVirtualResource gets a reference to the given VirtualResource and assigns it to the VirtualResource field.
func (o *EASRequirementsAllOf) SetVirtualResource(v VirtualResource) {
	o.VirtualResource = &v
}

// GetSoftwareImageInfo returns the SoftwareImageInfo field value if set, zero value otherwise.
func (o *EASRequirementsAllOf) GetSoftwareImageInfo() SoftwareImageInfo {
	if o == nil || isNil(o.SoftwareImageInfo) {
		var ret SoftwareImageInfo
		return ret
	}
	return *o.SoftwareImageInfo
}

// GetSoftwareImageInfoOk returns a tuple with the SoftwareImageInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EASRequirementsAllOf) GetSoftwareImageInfoOk() (*SoftwareImageInfo, bool) {
	if o == nil || isNil(o.SoftwareImageInfo) {
		return nil, false
	}
	return o.SoftwareImageInfo, true
}

// HasSoftwareImageInfo returns a boolean if a field has been set.
func (o *EASRequirementsAllOf) HasSoftwareImageInfo() bool {
	if o != nil && !isNil(o.SoftwareImageInfo) {
		return true
	}

	return false
}

// SetSoftwareImageInfo gets a reference to the given SoftwareImageInfo and assigns it to the SoftwareImageInfo field.
func (o *EASRequirementsAllOf) SetSoftwareImageInfo(v SoftwareImageInfo) {
	o.SoftwareImageInfo = &v
}

func (o EASRequirementsAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EASRequirementsAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RequiredEASservingLocation) {
		toSerialize["requiredEASservingLocation"] = o.RequiredEASservingLocation
	}
	if !isNil(o.AffinityAntiAffinity) {
		toSerialize["affinityAntiAffinity"] = o.AffinityAntiAffinity
	}
	if !isNil(o.ServiceContinuity) {
		toSerialize["serviceContinuity"] = o.ServiceContinuity
	}
	if !isNil(o.VirtualResource) {
		toSerialize["virtualResource"] = o.VirtualResource
	}
	if !isNil(o.SoftwareImageInfo) {
		toSerialize["softwareImageInfo"] = o.SoftwareImageInfo
	}
	return toSerialize, nil
}

type NullableEASRequirementsAllOf struct {
	value *EASRequirementsAllOf
	isSet bool
}

func (v NullableEASRequirementsAllOf) Get() *EASRequirementsAllOf {
	return v.value
}

func (v *NullableEASRequirementsAllOf) Set(val *EASRequirementsAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEASRequirementsAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEASRequirementsAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEASRequirementsAllOf(val *EASRequirementsAllOf) *NullableEASRequirementsAllOf {
	return &NullableEASRequirementsAllOf{value: val, isSet: true}
}

func (v NullableEASRequirementsAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEASRequirementsAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


