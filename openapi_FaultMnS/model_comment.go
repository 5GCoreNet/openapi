/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
	"time"
)

// checks if the Comment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Comment{}

// Comment struct for Comment
type Comment struct {
	CommentTime *time.Time `json:"commentTime,omitempty"`
	CommentUserId *string `json:"commentUserId,omitempty"`
	CommentSystemId *string `json:"commentSystemId,omitempty"`
	CommentText *string `json:"commentText,omitempty"`
}

// NewComment instantiates a new Comment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComment() *Comment {
	this := Comment{}
	return &this
}

// NewCommentWithDefaults instantiates a new Comment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommentWithDefaults() *Comment {
	this := Comment{}
	return &this
}

// GetCommentTime returns the CommentTime field value if set, zero value otherwise.
func (o *Comment) GetCommentTime() time.Time {
	if o == nil || IsNil(o.CommentTime) {
		var ret time.Time
		return ret
	}
	return *o.CommentTime
}

// GetCommentTimeOk returns a tuple with the CommentTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetCommentTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CommentTime) {
		return nil, false
	}
	return o.CommentTime, true
}

// HasCommentTime returns a boolean if a field has been set.
func (o *Comment) HasCommentTime() bool {
	if o != nil && !IsNil(o.CommentTime) {
		return true
	}

	return false
}

// SetCommentTime gets a reference to the given time.Time and assigns it to the CommentTime field.
func (o *Comment) SetCommentTime(v time.Time) {
	o.CommentTime = &v
}

// GetCommentUserId returns the CommentUserId field value if set, zero value otherwise.
func (o *Comment) GetCommentUserId() string {
	if o == nil || IsNil(o.CommentUserId) {
		var ret string
		return ret
	}
	return *o.CommentUserId
}

// GetCommentUserIdOk returns a tuple with the CommentUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetCommentUserIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommentUserId) {
		return nil, false
	}
	return o.CommentUserId, true
}

// HasCommentUserId returns a boolean if a field has been set.
func (o *Comment) HasCommentUserId() bool {
	if o != nil && !IsNil(o.CommentUserId) {
		return true
	}

	return false
}

// SetCommentUserId gets a reference to the given string and assigns it to the CommentUserId field.
func (o *Comment) SetCommentUserId(v string) {
	o.CommentUserId = &v
}

// GetCommentSystemId returns the CommentSystemId field value if set, zero value otherwise.
func (o *Comment) GetCommentSystemId() string {
	if o == nil || IsNil(o.CommentSystemId) {
		var ret string
		return ret
	}
	return *o.CommentSystemId
}

// GetCommentSystemIdOk returns a tuple with the CommentSystemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetCommentSystemIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommentSystemId) {
		return nil, false
	}
	return o.CommentSystemId, true
}

// HasCommentSystemId returns a boolean if a field has been set.
func (o *Comment) HasCommentSystemId() bool {
	if o != nil && !IsNil(o.CommentSystemId) {
		return true
	}

	return false
}

// SetCommentSystemId gets a reference to the given string and assigns it to the CommentSystemId field.
func (o *Comment) SetCommentSystemId(v string) {
	o.CommentSystemId = &v
}

// GetCommentText returns the CommentText field value if set, zero value otherwise.
func (o *Comment) GetCommentText() string {
	if o == nil || IsNil(o.CommentText) {
		var ret string
		return ret
	}
	return *o.CommentText
}

// GetCommentTextOk returns a tuple with the CommentText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Comment) GetCommentTextOk() (*string, bool) {
	if o == nil || IsNil(o.CommentText) {
		return nil, false
	}
	return o.CommentText, true
}

// HasCommentText returns a boolean if a field has been set.
func (o *Comment) HasCommentText() bool {
	if o != nil && !IsNil(o.CommentText) {
		return true
	}

	return false
}

// SetCommentText gets a reference to the given string and assigns it to the CommentText field.
func (o *Comment) SetCommentText(v string) {
	o.CommentText = &v
}

func (o Comment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Comment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommentTime) {
		toSerialize["commentTime"] = o.CommentTime
	}
	if !IsNil(o.CommentUserId) {
		toSerialize["commentUserId"] = o.CommentUserId
	}
	if !IsNil(o.CommentSystemId) {
		toSerialize["commentSystemId"] = o.CommentSystemId
	}
	if !IsNil(o.CommentText) {
		toSerialize["commentText"] = o.CommentText
	}
	return toSerialize, nil
}

type NullableComment struct {
	value *Comment
	isSet bool
}

func (v NullableComment) Get() *Comment {
	return v.value
}

func (v *NullableComment) Set(val *Comment) {
	v.value = val
	v.isSet = true
}

func (v NullableComment) IsSet() bool {
	return v.isSet
}

func (v *NullableComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComment(val *Comment) *NullableComment {
	return &NullableComment{value: val, isSet: true}
}

func (v NullableComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


