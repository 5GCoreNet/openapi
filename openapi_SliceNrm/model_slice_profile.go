/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
)

// checks if the SliceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SliceProfile{}

// SliceProfile struct for SliceProfile
type SliceProfile struct {
	ServiceProfileId *string `json:"serviceProfileId,omitempty"`
	PlmnInfoList []PlmnInfo `json:"plmnInfoList,omitempty"`
	CNSliceSubnetProfile *CNSliceSubnetProfile `json:"cNSliceSubnetProfile,omitempty"`
	RANSliceSubnetProfile *RANSliceSubnetProfile `json:"rANSliceSubnetProfile,omitempty"`
	TopSliceSubnetProfile *TopSliceSubnetProfile `json:"topSliceSubnetProfile,omitempty"`
}

// NewSliceProfile instantiates a new SliceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSliceProfile() *SliceProfile {
	this := SliceProfile{}
	return &this
}

// NewSliceProfileWithDefaults instantiates a new SliceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSliceProfileWithDefaults() *SliceProfile {
	this := SliceProfile{}
	return &this
}

// GetServiceProfileId returns the ServiceProfileId field value if set, zero value otherwise.
func (o *SliceProfile) GetServiceProfileId() string {
	if o == nil || isNil(o.ServiceProfileId) {
		var ret string
		return ret
	}
	return *o.ServiceProfileId
}

// GetServiceProfileIdOk returns a tuple with the ServiceProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceProfile) GetServiceProfileIdOk() (*string, bool) {
	if o == nil || isNil(o.ServiceProfileId) {
		return nil, false
	}
	return o.ServiceProfileId, true
}

// HasServiceProfileId returns a boolean if a field has been set.
func (o *SliceProfile) HasServiceProfileId() bool {
	if o != nil && !isNil(o.ServiceProfileId) {
		return true
	}

	return false
}

// SetServiceProfileId gets a reference to the given string and assigns it to the ServiceProfileId field.
func (o *SliceProfile) SetServiceProfileId(v string) {
	o.ServiceProfileId = &v
}

// GetPlmnInfoList returns the PlmnInfoList field value if set, zero value otherwise.
func (o *SliceProfile) GetPlmnInfoList() []PlmnInfo {
	if o == nil || isNil(o.PlmnInfoList) {
		var ret []PlmnInfo
		return ret
	}
	return o.PlmnInfoList
}

// GetPlmnInfoListOk returns a tuple with the PlmnInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceProfile) GetPlmnInfoListOk() ([]PlmnInfo, bool) {
	if o == nil || isNil(o.PlmnInfoList) {
		return nil, false
	}
	return o.PlmnInfoList, true
}

// HasPlmnInfoList returns a boolean if a field has been set.
func (o *SliceProfile) HasPlmnInfoList() bool {
	if o != nil && !isNil(o.PlmnInfoList) {
		return true
	}

	return false
}

// SetPlmnInfoList gets a reference to the given []PlmnInfo and assigns it to the PlmnInfoList field.
func (o *SliceProfile) SetPlmnInfoList(v []PlmnInfo) {
	o.PlmnInfoList = v
}

// GetCNSliceSubnetProfile returns the CNSliceSubnetProfile field value if set, zero value otherwise.
func (o *SliceProfile) GetCNSliceSubnetProfile() CNSliceSubnetProfile {
	if o == nil || isNil(o.CNSliceSubnetProfile) {
		var ret CNSliceSubnetProfile
		return ret
	}
	return *o.CNSliceSubnetProfile
}

// GetCNSliceSubnetProfileOk returns a tuple with the CNSliceSubnetProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceProfile) GetCNSliceSubnetProfileOk() (*CNSliceSubnetProfile, bool) {
	if o == nil || isNil(o.CNSliceSubnetProfile) {
		return nil, false
	}
	return o.CNSliceSubnetProfile, true
}

// HasCNSliceSubnetProfile returns a boolean if a field has been set.
func (o *SliceProfile) HasCNSliceSubnetProfile() bool {
	if o != nil && !isNil(o.CNSliceSubnetProfile) {
		return true
	}

	return false
}

// SetCNSliceSubnetProfile gets a reference to the given CNSliceSubnetProfile and assigns it to the CNSliceSubnetProfile field.
func (o *SliceProfile) SetCNSliceSubnetProfile(v CNSliceSubnetProfile) {
	o.CNSliceSubnetProfile = &v
}

// GetRANSliceSubnetProfile returns the RANSliceSubnetProfile field value if set, zero value otherwise.
func (o *SliceProfile) GetRANSliceSubnetProfile() RANSliceSubnetProfile {
	if o == nil || isNil(o.RANSliceSubnetProfile) {
		var ret RANSliceSubnetProfile
		return ret
	}
	return *o.RANSliceSubnetProfile
}

// GetRANSliceSubnetProfileOk returns a tuple with the RANSliceSubnetProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceProfile) GetRANSliceSubnetProfileOk() (*RANSliceSubnetProfile, bool) {
	if o == nil || isNil(o.RANSliceSubnetProfile) {
		return nil, false
	}
	return o.RANSliceSubnetProfile, true
}

// HasRANSliceSubnetProfile returns a boolean if a field has been set.
func (o *SliceProfile) HasRANSliceSubnetProfile() bool {
	if o != nil && !isNil(o.RANSliceSubnetProfile) {
		return true
	}

	return false
}

// SetRANSliceSubnetProfile gets a reference to the given RANSliceSubnetProfile and assigns it to the RANSliceSubnetProfile field.
func (o *SliceProfile) SetRANSliceSubnetProfile(v RANSliceSubnetProfile) {
	o.RANSliceSubnetProfile = &v
}

// GetTopSliceSubnetProfile returns the TopSliceSubnetProfile field value if set, zero value otherwise.
func (o *SliceProfile) GetTopSliceSubnetProfile() TopSliceSubnetProfile {
	if o == nil || isNil(o.TopSliceSubnetProfile) {
		var ret TopSliceSubnetProfile
		return ret
	}
	return *o.TopSliceSubnetProfile
}

// GetTopSliceSubnetProfileOk returns a tuple with the TopSliceSubnetProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SliceProfile) GetTopSliceSubnetProfileOk() (*TopSliceSubnetProfile, bool) {
	if o == nil || isNil(o.TopSliceSubnetProfile) {
		return nil, false
	}
	return o.TopSliceSubnetProfile, true
}

// HasTopSliceSubnetProfile returns a boolean if a field has been set.
func (o *SliceProfile) HasTopSliceSubnetProfile() bool {
	if o != nil && !isNil(o.TopSliceSubnetProfile) {
		return true
	}

	return false
}

// SetTopSliceSubnetProfile gets a reference to the given TopSliceSubnetProfile and assigns it to the TopSliceSubnetProfile field.
func (o *SliceProfile) SetTopSliceSubnetProfile(v TopSliceSubnetProfile) {
	o.TopSliceSubnetProfile = &v
}

func (o SliceProfile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SliceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ServiceProfileId) {
		toSerialize["serviceProfileId"] = o.ServiceProfileId
	}
	if !isNil(o.PlmnInfoList) {
		toSerialize["plmnInfoList"] = o.PlmnInfoList
	}
	if !isNil(o.CNSliceSubnetProfile) {
		toSerialize["cNSliceSubnetProfile"] = o.CNSliceSubnetProfile
	}
	if !isNil(o.RANSliceSubnetProfile) {
		toSerialize["rANSliceSubnetProfile"] = o.RANSliceSubnetProfile
	}
	if !isNil(o.TopSliceSubnetProfile) {
		toSerialize["topSliceSubnetProfile"] = o.TopSliceSubnetProfile
	}
	return toSerialize, nil
}

type NullableSliceProfile struct {
	value *SliceProfile
	isSet bool
}

func (v NullableSliceProfile) Get() *SliceProfile {
	return v.value
}

func (v *NullableSliceProfile) Set(val *SliceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceProfile(val *SliceProfile) *NullableSliceProfile {
	return &NullableSliceProfile{value: val, isSet: true}
}

func (v NullableSliceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


