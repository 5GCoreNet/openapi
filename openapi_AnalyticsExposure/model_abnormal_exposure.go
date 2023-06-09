/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
)

// checks if the AbnormalExposure type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AbnormalExposure{}

// AbnormalExposure Represents a user's abnormal behavior information.
type AbnormalExposure struct {
	Gpsis []string `json:"gpsis,omitempty"`
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn    *string   `json:"dnn,omitempty"`
	Snssai *Snssai   `json:"snssai,omitempty"`
	Excep  Exception `json:"excep"`
	// Unsigned integer indicating Sampling Ratio (see clauses 4.15.1 of 3GPP TS 23.502), expressed in percent.
	Ratio *int32 `json:"ratio,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Confidence   *int32                 `json:"confidence,omitempty"`
	AddtMeasInfo *AdditionalMeasurement `json:"addtMeasInfo,omitempty"`
}

// NewAbnormalExposure instantiates a new AbnormalExposure object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAbnormalExposure(excep Exception) *AbnormalExposure {
	this := AbnormalExposure{}
	this.Excep = excep
	return &this
}

// NewAbnormalExposureWithDefaults instantiates a new AbnormalExposure object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAbnormalExposureWithDefaults() *AbnormalExposure {
	this := AbnormalExposure{}
	return &this
}

// GetGpsis returns the Gpsis field value if set, zero value otherwise.
func (o *AbnormalExposure) GetGpsis() []string {
	if o == nil || IsNil(o.Gpsis) {
		var ret []string
		return ret
	}
	return o.Gpsis
}

// GetGpsisOk returns a tuple with the Gpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalExposure) GetGpsisOk() ([]string, bool) {
	if o == nil || IsNil(o.Gpsis) {
		return nil, false
	}
	return o.Gpsis, true
}

// HasGpsis returns a boolean if a field has been set.
func (o *AbnormalExposure) HasGpsis() bool {
	if o != nil && !IsNil(o.Gpsis) {
		return true
	}

	return false
}

// SetGpsis gets a reference to the given []string and assigns it to the Gpsis field.
func (o *AbnormalExposure) SetGpsis(v []string) {
	o.Gpsis = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AbnormalExposure) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalExposure) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AbnormalExposure) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *AbnormalExposure) SetAppId(v string) {
	o.AppId = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *AbnormalExposure) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalExposure) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *AbnormalExposure) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *AbnormalExposure) SetDnn(v string) {
	o.Dnn = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *AbnormalExposure) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalExposure) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *AbnormalExposure) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *AbnormalExposure) SetSnssai(v Snssai) {
	o.Snssai = &v
}

// GetExcep returns the Excep field value
func (o *AbnormalExposure) GetExcep() Exception {
	if o == nil {
		var ret Exception
		return ret
	}

	return o.Excep
}

// GetExcepOk returns a tuple with the Excep field value
// and a boolean to check if the value has been set.
func (o *AbnormalExposure) GetExcepOk() (*Exception, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Excep, true
}

// SetExcep sets field value
func (o *AbnormalExposure) SetExcep(v Exception) {
	o.Excep = v
}

// GetRatio returns the Ratio field value if set, zero value otherwise.
func (o *AbnormalExposure) GetRatio() int32 {
	if o == nil || IsNil(o.Ratio) {
		var ret int32
		return ret
	}
	return *o.Ratio
}

// GetRatioOk returns a tuple with the Ratio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalExposure) GetRatioOk() (*int32, bool) {
	if o == nil || IsNil(o.Ratio) {
		return nil, false
	}
	return o.Ratio, true
}

// HasRatio returns a boolean if a field has been set.
func (o *AbnormalExposure) HasRatio() bool {
	if o != nil && !IsNil(o.Ratio) {
		return true
	}

	return false
}

// SetRatio gets a reference to the given int32 and assigns it to the Ratio field.
func (o *AbnormalExposure) SetRatio(v int32) {
	o.Ratio = &v
}

// GetConfidence returns the Confidence field value if set, zero value otherwise.
func (o *AbnormalExposure) GetConfidence() int32 {
	if o == nil || IsNil(o.Confidence) {
		var ret int32
		return ret
	}
	return *o.Confidence
}

// GetConfidenceOk returns a tuple with the Confidence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalExposure) GetConfidenceOk() (*int32, bool) {
	if o == nil || IsNil(o.Confidence) {
		return nil, false
	}
	return o.Confidence, true
}

// HasConfidence returns a boolean if a field has been set.
func (o *AbnormalExposure) HasConfidence() bool {
	if o != nil && !IsNil(o.Confidence) {
		return true
	}

	return false
}

// SetConfidence gets a reference to the given int32 and assigns it to the Confidence field.
func (o *AbnormalExposure) SetConfidence(v int32) {
	o.Confidence = &v
}

// GetAddtMeasInfo returns the AddtMeasInfo field value if set, zero value otherwise.
func (o *AbnormalExposure) GetAddtMeasInfo() AdditionalMeasurement {
	if o == nil || IsNil(o.AddtMeasInfo) {
		var ret AdditionalMeasurement
		return ret
	}
	return *o.AddtMeasInfo
}

// GetAddtMeasInfoOk returns a tuple with the AddtMeasInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AbnormalExposure) GetAddtMeasInfoOk() (*AdditionalMeasurement, bool) {
	if o == nil || IsNil(o.AddtMeasInfo) {
		return nil, false
	}
	return o.AddtMeasInfo, true
}

// HasAddtMeasInfo returns a boolean if a field has been set.
func (o *AbnormalExposure) HasAddtMeasInfo() bool {
	if o != nil && !IsNil(o.AddtMeasInfo) {
		return true
	}

	return false
}

// SetAddtMeasInfo gets a reference to the given AdditionalMeasurement and assigns it to the AddtMeasInfo field.
func (o *AbnormalExposure) SetAddtMeasInfo(v AdditionalMeasurement) {
	o.AddtMeasInfo = &v
}

func (o AbnormalExposure) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AbnormalExposure) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Gpsis) {
		toSerialize["gpsis"] = o.Gpsis
	}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	toSerialize["excep"] = o.Excep
	if !IsNil(o.Ratio) {
		toSerialize["ratio"] = o.Ratio
	}
	if !IsNil(o.Confidence) {
		toSerialize["confidence"] = o.Confidence
	}
	if !IsNil(o.AddtMeasInfo) {
		toSerialize["addtMeasInfo"] = o.AddtMeasInfo
	}
	return toSerialize, nil
}

type NullableAbnormalExposure struct {
	value *AbnormalExposure
	isSet bool
}

func (v NullableAbnormalExposure) Get() *AbnormalExposure {
	return v.value
}

func (v *NullableAbnormalExposure) Set(val *AbnormalExposure) {
	v.value = val
	v.isSet = true
}

func (v NullableAbnormalExposure) IsSet() bool {
	return v.isSet
}

func (v *NullableAbnormalExposure) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAbnormalExposure(val *AbnormalExposure) *NullableAbnormalExposure {
	return &NullableAbnormalExposure{value: val, isSet: true}
}

func (v NullableAbnormalExposure) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAbnormalExposure) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
