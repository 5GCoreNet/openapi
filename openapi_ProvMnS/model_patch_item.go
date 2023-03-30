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

// checks if the PatchItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchItem{}

// PatchItem struct for PatchItem
type PatchItem struct {
	Op *PatchOperation `json:"op,omitempty"`
	From *string `json:"from,omitempty"`
	Path *string `json:"path,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

// NewPatchItem instantiates a new PatchItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchItem() *PatchItem {
	this := PatchItem{}
	return &this
}

// NewPatchItemWithDefaults instantiates a new PatchItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchItemWithDefaults() *PatchItem {
	this := PatchItem{}
	return &this
}

// GetOp returns the Op field value if set, zero value otherwise.
func (o *PatchItem) GetOp() PatchOperation {
	if o == nil || IsNil(o.Op) {
		var ret PatchOperation
		return ret
	}
	return *o.Op
}

// GetOpOk returns a tuple with the Op field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchItem) GetOpOk() (*PatchOperation, bool) {
	if o == nil || IsNil(o.Op) {
		return nil, false
	}
	return o.Op, true
}

// HasOp returns a boolean if a field has been set.
func (o *PatchItem) HasOp() bool {
	if o != nil && !IsNil(o.Op) {
		return true
	}

	return false
}

// SetOp gets a reference to the given PatchOperation and assigns it to the Op field.
func (o *PatchItem) SetOp(v PatchOperation) {
	o.Op = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PatchItem) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchItem) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PatchItem) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PatchItem) SetFrom(v string) {
	o.From = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *PatchItem) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchItem) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *PatchItem) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *PatchItem) SetPath(v string) {
	o.Path = &v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchItem) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchItem) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchItem) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *PatchItem) SetValue(v interface{}) {
	o.Value = v
}

func (o PatchItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Op) {
		toSerialize["op"] = o.Op
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePatchItem struct {
	value *PatchItem
	isSet bool
}

func (v NullablePatchItem) Get() *PatchItem {
	return v.value
}

func (v *NullablePatchItem) Set(val *PatchItem) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchItem) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchItem(val *PatchItem) *NullablePatchItem {
	return &NullablePatchItem{value: val, isSet: true}
}

func (v NullablePatchItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


