/*
M1_ContentProtocolsDiscovery

5GMS AF M1 Content Protocols Discovery API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M1_ContentProtocolsDiscovery

import (
	"encoding/json"
)

// checks if the ContentProtocols type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContentProtocols{}

// ContentProtocols A representation of the Content Protocols resource.
type ContentProtocols struct {
	DownlinkIngestProtocols []ContentProtocolDescriptor `json:"downlinkIngestProtocols,omitempty"`
	UplinkEgestProtocols    []ContentProtocolDescriptor `json:"uplinkEgestProtocols,omitempty"`
	GeoFencingLocatorTypes  []string                    `json:"geoFencingLocatorTypes,omitempty"`
}

// NewContentProtocols instantiates a new ContentProtocols object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContentProtocols() *ContentProtocols {
	this := ContentProtocols{}
	return &this
}

// NewContentProtocolsWithDefaults instantiates a new ContentProtocols object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentProtocolsWithDefaults() *ContentProtocols {
	this := ContentProtocols{}
	return &this
}

// GetDownlinkIngestProtocols returns the DownlinkIngestProtocols field value if set, zero value otherwise.
func (o *ContentProtocols) GetDownlinkIngestProtocols() []ContentProtocolDescriptor {
	if o == nil || IsNil(o.DownlinkIngestProtocols) {
		var ret []ContentProtocolDescriptor
		return ret
	}
	return o.DownlinkIngestProtocols
}

// GetDownlinkIngestProtocolsOk returns a tuple with the DownlinkIngestProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentProtocols) GetDownlinkIngestProtocolsOk() ([]ContentProtocolDescriptor, bool) {
	if o == nil || IsNil(o.DownlinkIngestProtocols) {
		return nil, false
	}
	return o.DownlinkIngestProtocols, true
}

// HasDownlinkIngestProtocols returns a boolean if a field has been set.
func (o *ContentProtocols) HasDownlinkIngestProtocols() bool {
	if o != nil && !IsNil(o.DownlinkIngestProtocols) {
		return true
	}

	return false
}

// SetDownlinkIngestProtocols gets a reference to the given []ContentProtocolDescriptor and assigns it to the DownlinkIngestProtocols field.
func (o *ContentProtocols) SetDownlinkIngestProtocols(v []ContentProtocolDescriptor) {
	o.DownlinkIngestProtocols = v
}

// GetUplinkEgestProtocols returns the UplinkEgestProtocols field value if set, zero value otherwise.
func (o *ContentProtocols) GetUplinkEgestProtocols() []ContentProtocolDescriptor {
	if o == nil || IsNil(o.UplinkEgestProtocols) {
		var ret []ContentProtocolDescriptor
		return ret
	}
	return o.UplinkEgestProtocols
}

// GetUplinkEgestProtocolsOk returns a tuple with the UplinkEgestProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentProtocols) GetUplinkEgestProtocolsOk() ([]ContentProtocolDescriptor, bool) {
	if o == nil || IsNil(o.UplinkEgestProtocols) {
		return nil, false
	}
	return o.UplinkEgestProtocols, true
}

// HasUplinkEgestProtocols returns a boolean if a field has been set.
func (o *ContentProtocols) HasUplinkEgestProtocols() bool {
	if o != nil && !IsNil(o.UplinkEgestProtocols) {
		return true
	}

	return false
}

// SetUplinkEgestProtocols gets a reference to the given []ContentProtocolDescriptor and assigns it to the UplinkEgestProtocols field.
func (o *ContentProtocols) SetUplinkEgestProtocols(v []ContentProtocolDescriptor) {
	o.UplinkEgestProtocols = v
}

// GetGeoFencingLocatorTypes returns the GeoFencingLocatorTypes field value if set, zero value otherwise.
func (o *ContentProtocols) GetGeoFencingLocatorTypes() []string {
	if o == nil || IsNil(o.GeoFencingLocatorTypes) {
		var ret []string
		return ret
	}
	return o.GeoFencingLocatorTypes
}

// GetGeoFencingLocatorTypesOk returns a tuple with the GeoFencingLocatorTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContentProtocols) GetGeoFencingLocatorTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.GeoFencingLocatorTypes) {
		return nil, false
	}
	return o.GeoFencingLocatorTypes, true
}

// HasGeoFencingLocatorTypes returns a boolean if a field has been set.
func (o *ContentProtocols) HasGeoFencingLocatorTypes() bool {
	if o != nil && !IsNil(o.GeoFencingLocatorTypes) {
		return true
	}

	return false
}

// SetGeoFencingLocatorTypes gets a reference to the given []string and assigns it to the GeoFencingLocatorTypes field.
func (o *ContentProtocols) SetGeoFencingLocatorTypes(v []string) {
	o.GeoFencingLocatorTypes = v
}

func (o ContentProtocols) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContentProtocols) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DownlinkIngestProtocols) {
		toSerialize["downlinkIngestProtocols"] = o.DownlinkIngestProtocols
	}
	if !IsNil(o.UplinkEgestProtocols) {
		toSerialize["uplinkEgestProtocols"] = o.UplinkEgestProtocols
	}
	if !IsNil(o.GeoFencingLocatorTypes) {
		toSerialize["geoFencingLocatorTypes"] = o.GeoFencingLocatorTypes
	}
	return toSerialize, nil
}

type NullableContentProtocols struct {
	value *ContentProtocols
	isSet bool
}

func (v NullableContentProtocols) Get() *ContentProtocols {
	return v.value
}

func (v *NullableContentProtocols) Set(val *ContentProtocols) {
	v.value = val
	v.isSet = true
}

func (v NullableContentProtocols) IsSet() bool {
	return v.isSet
}

func (v *NullableContentProtocols) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContentProtocols(val *ContentProtocols) *NullableContentProtocols {
	return &NullableContentProtocols{value: val, isSet: true}
}

func (v NullableContentProtocols) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContentProtocols) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
