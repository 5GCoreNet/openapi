/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the MLTrainingFunctionSingleAllOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MLTrainingFunctionSingleAllOf1{}

// MLTrainingFunctionSingleAllOf1 struct for MLTrainingFunctionSingleAllOf1
type MLTrainingFunctionSingleAllOf1 struct {
	MLTrainingRequest []MLTrainingRequestSingle `json:"MLTrainingRequest,omitempty"`
	MLTrainingProcess []MLTrainingProcessSingle `json:"MLTrainingProcess,omitempty"`
	MLTrainingReport  []MLTrainingReportSingle  `json:"MLTrainingReport,omitempty"`
}

// NewMLTrainingFunctionSingleAllOf1 instantiates a new MLTrainingFunctionSingleAllOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMLTrainingFunctionSingleAllOf1() *MLTrainingFunctionSingleAllOf1 {
	this := MLTrainingFunctionSingleAllOf1{}
	return &this
}

// NewMLTrainingFunctionSingleAllOf1WithDefaults instantiates a new MLTrainingFunctionSingleAllOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMLTrainingFunctionSingleAllOf1WithDefaults() *MLTrainingFunctionSingleAllOf1 {
	this := MLTrainingFunctionSingleAllOf1{}
	return &this
}

// GetMLTrainingRequest returns the MLTrainingRequest field value if set, zero value otherwise.
func (o *MLTrainingFunctionSingleAllOf1) GetMLTrainingRequest() []MLTrainingRequestSingle {
	if o == nil || IsNil(o.MLTrainingRequest) {
		var ret []MLTrainingRequestSingle
		return ret
	}
	return o.MLTrainingRequest
}

// GetMLTrainingRequestOk returns a tuple with the MLTrainingRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingFunctionSingleAllOf1) GetMLTrainingRequestOk() ([]MLTrainingRequestSingle, bool) {
	if o == nil || IsNil(o.MLTrainingRequest) {
		return nil, false
	}
	return o.MLTrainingRequest, true
}

// HasMLTrainingRequest returns a boolean if a field has been set.
func (o *MLTrainingFunctionSingleAllOf1) HasMLTrainingRequest() bool {
	if o != nil && !IsNil(o.MLTrainingRequest) {
		return true
	}

	return false
}

// SetMLTrainingRequest gets a reference to the given []MLTrainingRequestSingle and assigns it to the MLTrainingRequest field.
func (o *MLTrainingFunctionSingleAllOf1) SetMLTrainingRequest(v []MLTrainingRequestSingle) {
	o.MLTrainingRequest = v
}

// GetMLTrainingProcess returns the MLTrainingProcess field value if set, zero value otherwise.
func (o *MLTrainingFunctionSingleAllOf1) GetMLTrainingProcess() []MLTrainingProcessSingle {
	if o == nil || IsNil(o.MLTrainingProcess) {
		var ret []MLTrainingProcessSingle
		return ret
	}
	return o.MLTrainingProcess
}

// GetMLTrainingProcessOk returns a tuple with the MLTrainingProcess field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingFunctionSingleAllOf1) GetMLTrainingProcessOk() ([]MLTrainingProcessSingle, bool) {
	if o == nil || IsNil(o.MLTrainingProcess) {
		return nil, false
	}
	return o.MLTrainingProcess, true
}

// HasMLTrainingProcess returns a boolean if a field has been set.
func (o *MLTrainingFunctionSingleAllOf1) HasMLTrainingProcess() bool {
	if o != nil && !IsNil(o.MLTrainingProcess) {
		return true
	}

	return false
}

// SetMLTrainingProcess gets a reference to the given []MLTrainingProcessSingle and assigns it to the MLTrainingProcess field.
func (o *MLTrainingFunctionSingleAllOf1) SetMLTrainingProcess(v []MLTrainingProcessSingle) {
	o.MLTrainingProcess = v
}

// GetMLTrainingReport returns the MLTrainingReport field value if set, zero value otherwise.
func (o *MLTrainingFunctionSingleAllOf1) GetMLTrainingReport() []MLTrainingReportSingle {
	if o == nil || IsNil(o.MLTrainingReport) {
		var ret []MLTrainingReportSingle
		return ret
	}
	return o.MLTrainingReport
}

// GetMLTrainingReportOk returns a tuple with the MLTrainingReport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MLTrainingFunctionSingleAllOf1) GetMLTrainingReportOk() ([]MLTrainingReportSingle, bool) {
	if o == nil || IsNil(o.MLTrainingReport) {
		return nil, false
	}
	return o.MLTrainingReport, true
}

// HasMLTrainingReport returns a boolean if a field has been set.
func (o *MLTrainingFunctionSingleAllOf1) HasMLTrainingReport() bool {
	if o != nil && !IsNil(o.MLTrainingReport) {
		return true
	}

	return false
}

// SetMLTrainingReport gets a reference to the given []MLTrainingReportSingle and assigns it to the MLTrainingReport field.
func (o *MLTrainingFunctionSingleAllOf1) SetMLTrainingReport(v []MLTrainingReportSingle) {
	o.MLTrainingReport = v
}

func (o MLTrainingFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MLTrainingFunctionSingleAllOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MLTrainingRequest) {
		toSerialize["MLTrainingRequest"] = o.MLTrainingRequest
	}
	if !IsNil(o.MLTrainingProcess) {
		toSerialize["MLTrainingProcess"] = o.MLTrainingProcess
	}
	if !IsNil(o.MLTrainingReport) {
		toSerialize["MLTrainingReport"] = o.MLTrainingReport
	}
	return toSerialize, nil
}

type NullableMLTrainingFunctionSingleAllOf1 struct {
	value *MLTrainingFunctionSingleAllOf1
	isSet bool
}

func (v NullableMLTrainingFunctionSingleAllOf1) Get() *MLTrainingFunctionSingleAllOf1 {
	return v.value
}

func (v *NullableMLTrainingFunctionSingleAllOf1) Set(val *MLTrainingFunctionSingleAllOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableMLTrainingFunctionSingleAllOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableMLTrainingFunctionSingleAllOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMLTrainingFunctionSingleAllOf1(val *MLTrainingFunctionSingleAllOf1) *NullableMLTrainingFunctionSingleAllOf1 {
	return &NullableMLTrainingFunctionSingleAllOf1{value: val, isSet: true}
}

func (v NullableMLTrainingFunctionSingleAllOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMLTrainingFunctionSingleAllOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
