/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the AmfFunctionSingle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfFunctionSingle{}

// AmfFunctionSingle struct for AmfFunctionSingle
type AmfFunctionSingle struct {
	Top
	Attributes       *AmfFunctionSingleAllOfAttributes `json:"attributes,omitempty"`
	PerfMetricJob    []PerfMetricJobSingle             `json:"PerfMetricJob,omitempty"`
	ThresholdMonitor []ThresholdMonitorSingle          `json:"ThresholdMonitor,omitempty"`
	ManagedNFService []ManagedNFServiceSingle          `json:"ManagedNFService,omitempty"`
	TraceJob         []TraceJobSingle                  `json:"TraceJob,omitempty"`
	EPN2             []EPN2Single                      `json:"EP_N2,omitempty"`
	EPN8             []EPN8Single                      `json:"EP_N8,omitempty"`
	EPN11            []EPN11Single                     `json:"EP_N11,omitempty"`
	EPN12            []EPN12Single                     `json:"EP_N12,omitempty"`
	EPN14            []EPN14Single                     `json:"EP_N14,omitempty"`
	EPN15            []EPN15Single                     `json:"EP_N15,omitempty"`
	EPN17            []EPN17Single                     `json:"EP_N17,omitempty"`
	EPN20            []EPN20Single                     `json:"EP_N20,omitempty"`
	EPN22            []EPN22Single                     `json:"EP_N22,omitempty"`
	EPN26            []EPN26Single                     `json:"EP_N26,omitempty"`
	EP_NLS           []EPNLSSingle                     `json:"EP_NLS,omitempty"`
	EP_NLG           []EPNLGSingle                     `json:"EP_NLG,omitempty"`
}

// NewAmfFunctionSingle instantiates a new AmfFunctionSingle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfFunctionSingle(id NullableString) *AmfFunctionSingle {
	this := AmfFunctionSingle{}
	this.Id = id
	return &this
}

// NewAmfFunctionSingleWithDefaults instantiates a new AmfFunctionSingle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfFunctionSingleWithDefaults() *AmfFunctionSingle {
	this := AmfFunctionSingle{}
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetAttributes() AmfFunctionSingleAllOfAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AmfFunctionSingleAllOfAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetAttributesOk() (*AmfFunctionSingleAllOfAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AmfFunctionSingleAllOfAttributes and assigns it to the Attributes field.
func (o *AmfFunctionSingle) SetAttributes(v AmfFunctionSingleAllOfAttributes) {
	o.Attributes = &v
}

// GetPerfMetricJob returns the PerfMetricJob field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetPerfMetricJob() []PerfMetricJobSingle {
	if o == nil || IsNil(o.PerfMetricJob) {
		var ret []PerfMetricJobSingle
		return ret
	}
	return o.PerfMetricJob
}

// GetPerfMetricJobOk returns a tuple with the PerfMetricJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetPerfMetricJobOk() ([]PerfMetricJobSingle, bool) {
	if o == nil || IsNil(o.PerfMetricJob) {
		return nil, false
	}
	return o.PerfMetricJob, true
}

// HasPerfMetricJob returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasPerfMetricJob() bool {
	if o != nil && !IsNil(o.PerfMetricJob) {
		return true
	}

	return false
}

// SetPerfMetricJob gets a reference to the given []PerfMetricJobSingle and assigns it to the PerfMetricJob field.
func (o *AmfFunctionSingle) SetPerfMetricJob(v []PerfMetricJobSingle) {
	o.PerfMetricJob = v
}

// GetThresholdMonitor returns the ThresholdMonitor field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetThresholdMonitor() []ThresholdMonitorSingle {
	if o == nil || IsNil(o.ThresholdMonitor) {
		var ret []ThresholdMonitorSingle
		return ret
	}
	return o.ThresholdMonitor
}

// GetThresholdMonitorOk returns a tuple with the ThresholdMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetThresholdMonitorOk() ([]ThresholdMonitorSingle, bool) {
	if o == nil || IsNil(o.ThresholdMonitor) {
		return nil, false
	}
	return o.ThresholdMonitor, true
}

// HasThresholdMonitor returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasThresholdMonitor() bool {
	if o != nil && !IsNil(o.ThresholdMonitor) {
		return true
	}

	return false
}

// SetThresholdMonitor gets a reference to the given []ThresholdMonitorSingle and assigns it to the ThresholdMonitor field.
func (o *AmfFunctionSingle) SetThresholdMonitor(v []ThresholdMonitorSingle) {
	o.ThresholdMonitor = v
}

// GetManagedNFService returns the ManagedNFService field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetManagedNFService() []ManagedNFServiceSingle {
	if o == nil || IsNil(o.ManagedNFService) {
		var ret []ManagedNFServiceSingle
		return ret
	}
	return o.ManagedNFService
}

// GetManagedNFServiceOk returns a tuple with the ManagedNFService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetManagedNFServiceOk() ([]ManagedNFServiceSingle, bool) {
	if o == nil || IsNil(o.ManagedNFService) {
		return nil, false
	}
	return o.ManagedNFService, true
}

// HasManagedNFService returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasManagedNFService() bool {
	if o != nil && !IsNil(o.ManagedNFService) {
		return true
	}

	return false
}

// SetManagedNFService gets a reference to the given []ManagedNFServiceSingle and assigns it to the ManagedNFService field.
func (o *AmfFunctionSingle) SetManagedNFService(v []ManagedNFServiceSingle) {
	o.ManagedNFService = v
}

// GetTraceJob returns the TraceJob field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetTraceJob() []TraceJobSingle {
	if o == nil || IsNil(o.TraceJob) {
		var ret []TraceJobSingle
		return ret
	}
	return o.TraceJob
}

// GetTraceJobOk returns a tuple with the TraceJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetTraceJobOk() ([]TraceJobSingle, bool) {
	if o == nil || IsNil(o.TraceJob) {
		return nil, false
	}
	return o.TraceJob, true
}

// HasTraceJob returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasTraceJob() bool {
	if o != nil && !IsNil(o.TraceJob) {
		return true
	}

	return false
}

// SetTraceJob gets a reference to the given []TraceJobSingle and assigns it to the TraceJob field.
func (o *AmfFunctionSingle) SetTraceJob(v []TraceJobSingle) {
	o.TraceJob = v
}

// GetEPN2 returns the EPN2 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN2() []EPN2Single {
	if o == nil || IsNil(o.EPN2) {
		var ret []EPN2Single
		return ret
	}
	return o.EPN2
}

// GetEPN2Ok returns a tuple with the EPN2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN2Ok() ([]EPN2Single, bool) {
	if o == nil || IsNil(o.EPN2) {
		return nil, false
	}
	return o.EPN2, true
}

// HasEPN2 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN2() bool {
	if o != nil && !IsNil(o.EPN2) {
		return true
	}

	return false
}

// SetEPN2 gets a reference to the given []EPN2Single and assigns it to the EPN2 field.
func (o *AmfFunctionSingle) SetEPN2(v []EPN2Single) {
	o.EPN2 = v
}

// GetEPN8 returns the EPN8 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN8() []EPN8Single {
	if o == nil || IsNil(o.EPN8) {
		var ret []EPN8Single
		return ret
	}
	return o.EPN8
}

// GetEPN8Ok returns a tuple with the EPN8 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN8Ok() ([]EPN8Single, bool) {
	if o == nil || IsNil(o.EPN8) {
		return nil, false
	}
	return o.EPN8, true
}

// HasEPN8 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN8() bool {
	if o != nil && !IsNil(o.EPN8) {
		return true
	}

	return false
}

// SetEPN8 gets a reference to the given []EPN8Single and assigns it to the EPN8 field.
func (o *AmfFunctionSingle) SetEPN8(v []EPN8Single) {
	o.EPN8 = v
}

// GetEPN11 returns the EPN11 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN11() []EPN11Single {
	if o == nil || IsNil(o.EPN11) {
		var ret []EPN11Single
		return ret
	}
	return o.EPN11
}

// GetEPN11Ok returns a tuple with the EPN11 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN11Ok() ([]EPN11Single, bool) {
	if o == nil || IsNil(o.EPN11) {
		return nil, false
	}
	return o.EPN11, true
}

// HasEPN11 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN11() bool {
	if o != nil && !IsNil(o.EPN11) {
		return true
	}

	return false
}

// SetEPN11 gets a reference to the given []EPN11Single and assigns it to the EPN11 field.
func (o *AmfFunctionSingle) SetEPN11(v []EPN11Single) {
	o.EPN11 = v
}

// GetEPN12 returns the EPN12 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN12() []EPN12Single {
	if o == nil || IsNil(o.EPN12) {
		var ret []EPN12Single
		return ret
	}
	return o.EPN12
}

// GetEPN12Ok returns a tuple with the EPN12 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN12Ok() ([]EPN12Single, bool) {
	if o == nil || IsNil(o.EPN12) {
		return nil, false
	}
	return o.EPN12, true
}

// HasEPN12 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN12() bool {
	if o != nil && !IsNil(o.EPN12) {
		return true
	}

	return false
}

// SetEPN12 gets a reference to the given []EPN12Single and assigns it to the EPN12 field.
func (o *AmfFunctionSingle) SetEPN12(v []EPN12Single) {
	o.EPN12 = v
}

// GetEPN14 returns the EPN14 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN14() []EPN14Single {
	if o == nil || IsNil(o.EPN14) {
		var ret []EPN14Single
		return ret
	}
	return o.EPN14
}

// GetEPN14Ok returns a tuple with the EPN14 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN14Ok() ([]EPN14Single, bool) {
	if o == nil || IsNil(o.EPN14) {
		return nil, false
	}
	return o.EPN14, true
}

// HasEPN14 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN14() bool {
	if o != nil && !IsNil(o.EPN14) {
		return true
	}

	return false
}

// SetEPN14 gets a reference to the given []EPN14Single and assigns it to the EPN14 field.
func (o *AmfFunctionSingle) SetEPN14(v []EPN14Single) {
	o.EPN14 = v
}

// GetEPN15 returns the EPN15 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN15() []EPN15Single {
	if o == nil || IsNil(o.EPN15) {
		var ret []EPN15Single
		return ret
	}
	return o.EPN15
}

// GetEPN15Ok returns a tuple with the EPN15 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN15Ok() ([]EPN15Single, bool) {
	if o == nil || IsNil(o.EPN15) {
		return nil, false
	}
	return o.EPN15, true
}

// HasEPN15 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN15() bool {
	if o != nil && !IsNil(o.EPN15) {
		return true
	}

	return false
}

// SetEPN15 gets a reference to the given []EPN15Single and assigns it to the EPN15 field.
func (o *AmfFunctionSingle) SetEPN15(v []EPN15Single) {
	o.EPN15 = v
}

// GetEPN17 returns the EPN17 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN17() []EPN17Single {
	if o == nil || IsNil(o.EPN17) {
		var ret []EPN17Single
		return ret
	}
	return o.EPN17
}

// GetEPN17Ok returns a tuple with the EPN17 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN17Ok() ([]EPN17Single, bool) {
	if o == nil || IsNil(o.EPN17) {
		return nil, false
	}
	return o.EPN17, true
}

// HasEPN17 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN17() bool {
	if o != nil && !IsNil(o.EPN17) {
		return true
	}

	return false
}

// SetEPN17 gets a reference to the given []EPN17Single and assigns it to the EPN17 field.
func (o *AmfFunctionSingle) SetEPN17(v []EPN17Single) {
	o.EPN17 = v
}

// GetEPN20 returns the EPN20 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN20() []EPN20Single {
	if o == nil || IsNil(o.EPN20) {
		var ret []EPN20Single
		return ret
	}
	return o.EPN20
}

// GetEPN20Ok returns a tuple with the EPN20 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN20Ok() ([]EPN20Single, bool) {
	if o == nil || IsNil(o.EPN20) {
		return nil, false
	}
	return o.EPN20, true
}

// HasEPN20 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN20() bool {
	if o != nil && !IsNil(o.EPN20) {
		return true
	}

	return false
}

// SetEPN20 gets a reference to the given []EPN20Single and assigns it to the EPN20 field.
func (o *AmfFunctionSingle) SetEPN20(v []EPN20Single) {
	o.EPN20 = v
}

// GetEPN22 returns the EPN22 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN22() []EPN22Single {
	if o == nil || IsNil(o.EPN22) {
		var ret []EPN22Single
		return ret
	}
	return o.EPN22
}

// GetEPN22Ok returns a tuple with the EPN22 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN22Ok() ([]EPN22Single, bool) {
	if o == nil || IsNil(o.EPN22) {
		return nil, false
	}
	return o.EPN22, true
}

// HasEPN22 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN22() bool {
	if o != nil && !IsNil(o.EPN22) {
		return true
	}

	return false
}

// SetEPN22 gets a reference to the given []EPN22Single and assigns it to the EPN22 field.
func (o *AmfFunctionSingle) SetEPN22(v []EPN22Single) {
	o.EPN22 = v
}

// GetEPN26 returns the EPN26 field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEPN26() []EPN26Single {
	if o == nil || IsNil(o.EPN26) {
		var ret []EPN26Single
		return ret
	}
	return o.EPN26
}

// GetEPN26Ok returns a tuple with the EPN26 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEPN26Ok() ([]EPN26Single, bool) {
	if o == nil || IsNil(o.EPN26) {
		return nil, false
	}
	return o.EPN26, true
}

// HasEPN26 returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEPN26() bool {
	if o != nil && !IsNil(o.EPN26) {
		return true
	}

	return false
}

// SetEPN26 gets a reference to the given []EPN26Single and assigns it to the EPN26 field.
func (o *AmfFunctionSingle) SetEPN26(v []EPN26Single) {
	o.EPN26 = v
}

// GetEP_NLS returns the EP_NLS field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEP_NLS() []EPNLSSingle {
	if o == nil || IsNil(o.EP_NLS) {
		var ret []EPNLSSingle
		return ret
	}
	return o.EP_NLS
}

// GetEP_NLSOk returns a tuple with the EP_NLS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEP_NLSOk() ([]EPNLSSingle, bool) {
	if o == nil || IsNil(o.EP_NLS) {
		return nil, false
	}
	return o.EP_NLS, true
}

// HasEP_NLS returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEP_NLS() bool {
	if o != nil && !IsNil(o.EP_NLS) {
		return true
	}

	return false
}

// SetEP_NLS gets a reference to the given []EPNLSSingle and assigns it to the EP_NLS field.
func (o *AmfFunctionSingle) SetEP_NLS(v []EPNLSSingle) {
	o.EP_NLS = v
}

// GetEP_NLG returns the EP_NLG field value if set, zero value otherwise.
func (o *AmfFunctionSingle) GetEP_NLG() []EPNLGSingle {
	if o == nil || IsNil(o.EP_NLG) {
		var ret []EPNLGSingle
		return ret
	}
	return o.EP_NLG
}

// GetEP_NLGOk returns a tuple with the EP_NLG field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfFunctionSingle) GetEP_NLGOk() ([]EPNLGSingle, bool) {
	if o == nil || IsNil(o.EP_NLG) {
		return nil, false
	}
	return o.EP_NLG, true
}

// HasEP_NLG returns a boolean if a field has been set.
func (o *AmfFunctionSingle) HasEP_NLG() bool {
	if o != nil && !IsNil(o.EP_NLG) {
		return true
	}

	return false
}

// SetEP_NLG gets a reference to the given []EPNLGSingle and assigns it to the EP_NLG field.
func (o *AmfFunctionSingle) SetEP_NLG(v []EPNLGSingle) {
	o.EP_NLG = v
}

func (o AmfFunctionSingle) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfFunctionSingle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTop, errTop := json.Marshal(o.Top)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	errTop = json.Unmarshal([]byte(serializedTop), &toSerialize)
	if errTop != nil {
		return map[string]interface{}{}, errTop
	}
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.PerfMetricJob) {
		toSerialize["PerfMetricJob"] = o.PerfMetricJob
	}
	if !IsNil(o.ThresholdMonitor) {
		toSerialize["ThresholdMonitor"] = o.ThresholdMonitor
	}
	if !IsNil(o.ManagedNFService) {
		toSerialize["ManagedNFService"] = o.ManagedNFService
	}
	if !IsNil(o.TraceJob) {
		toSerialize["TraceJob"] = o.TraceJob
	}
	if !IsNil(o.EPN2) {
		toSerialize["EP_N2"] = o.EPN2
	}
	if !IsNil(o.EPN8) {
		toSerialize["EP_N8"] = o.EPN8
	}
	if !IsNil(o.EPN11) {
		toSerialize["EP_N11"] = o.EPN11
	}
	if !IsNil(o.EPN12) {
		toSerialize["EP_N12"] = o.EPN12
	}
	if !IsNil(o.EPN14) {
		toSerialize["EP_N14"] = o.EPN14
	}
	if !IsNil(o.EPN15) {
		toSerialize["EP_N15"] = o.EPN15
	}
	if !IsNil(o.EPN17) {
		toSerialize["EP_N17"] = o.EPN17
	}
	if !IsNil(o.EPN20) {
		toSerialize["EP_N20"] = o.EPN20
	}
	if !IsNil(o.EPN22) {
		toSerialize["EP_N22"] = o.EPN22
	}
	if !IsNil(o.EPN26) {
		toSerialize["EP_N26"] = o.EPN26
	}
	if !IsNil(o.EP_NLS) {
		toSerialize["EP_NLS"] = o.EP_NLS
	}
	if !IsNil(o.EP_NLG) {
		toSerialize["EP_NLG"] = o.EP_NLG
	}
	return toSerialize, nil
}

type NullableAmfFunctionSingle struct {
	value *AmfFunctionSingle
	isSet bool
}

func (v NullableAmfFunctionSingle) Get() *AmfFunctionSingle {
	return v.value
}

func (v *NullableAmfFunctionSingle) Set(val *AmfFunctionSingle) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfFunctionSingle) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfFunctionSingle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfFunctionSingle(val *AmfFunctionSingle) *NullableAmfFunctionSingle {
	return &NullableAmfFunctionSingle{value: val, isSet: true}
}

func (v NullableAmfFunctionSingle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfFunctionSingle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
