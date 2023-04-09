/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
)

// checks if the ChangeItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChangeItem{}

// ChangeItem It contains data which need to be changed.
type ChangeItem struct {
	Op ChangeType `json:"op"`
	// contains a JSON pointer value (as defined in IETF RFC 6901) that references a target  location within the resource on which the change has been applied.
	Path string `json:"path"`
	// indicates the path of the source JSON element (according to JSON Pointer syntax)  being moved or copied to the location indicated by the \"path\" attribute. It shall  be present if the \"op\" attribute is of value \"MOVE\".
	From      *string     `json:"from,omitempty"`
	OrigValue interface{} `json:"origValue,omitempty"`
	NewValue  interface{} `json:"newValue,omitempty"`
}

// NewChangeItem instantiates a new ChangeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChangeItem(op ChangeType, path string) *ChangeItem {
	this := ChangeItem{}
	this.Op = op
	this.Path = path
	return &this
}

// NewChangeItemWithDefaults instantiates a new ChangeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChangeItemWithDefaults() *ChangeItem {
	this := ChangeItem{}
	return &this
}

// GetOp returns the Op field value
func (o *ChangeItem) GetOp() ChangeType {
	if o == nil {
		var ret ChangeType
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *ChangeItem) GetOpOk() (*ChangeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *ChangeItem) SetOp(v ChangeType) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *ChangeItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *ChangeItem) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *ChangeItem) SetPath(v string) {
	o.Path = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ChangeItem) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChangeItem) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ChangeItem) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *ChangeItem) SetFrom(v string) {
	o.From = &v
}

// GetOrigValue returns the OrigValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeItem) GetOrigValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.OrigValue
}

// GetOrigValueOk returns a tuple with the OrigValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeItem) GetOrigValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.OrigValue) {
		return nil, false
	}
	return &o.OrigValue, true
}

// HasOrigValue returns a boolean if a field has been set.
func (o *ChangeItem) HasOrigValue() bool {
	if o != nil && IsNil(o.OrigValue) {
		return true
	}

	return false
}

// SetOrigValue gets a reference to the given interface{} and assigns it to the OrigValue field.
func (o *ChangeItem) SetOrigValue(v interface{}) {
	o.OrigValue = v
}

// GetNewValue returns the NewValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChangeItem) GetNewValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.NewValue
}

// GetNewValueOk returns a tuple with the NewValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChangeItem) GetNewValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.NewValue) {
		return nil, false
	}
	return &o.NewValue, true
}

// HasNewValue returns a boolean if a field has been set.
func (o *ChangeItem) HasNewValue() bool {
	if o != nil && IsNil(o.NewValue) {
		return true
	}

	return false
}

// SetNewValue gets a reference to the given interface{} and assigns it to the NewValue field.
func (o *ChangeItem) SetNewValue(v interface{}) {
	o.NewValue = v
}

func (o ChangeItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChangeItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if o.OrigValue != nil {
		toSerialize["origValue"] = o.OrigValue
	}
	if o.NewValue != nil {
		toSerialize["newValue"] = o.NewValue
	}
	return toSerialize, nil
}

type NullableChangeItem struct {
	value *ChangeItem
	isSet bool
}

func (v NullableChangeItem) Get() *ChangeItem {
	return v.value
}

func (v *NullableChangeItem) Set(val *ChangeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableChangeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableChangeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChangeItem(val *ChangeItem) *NullableChangeItem {
	return &NullableChangeItem{value: val, isSet: true}
}

func (v NullableChangeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChangeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
