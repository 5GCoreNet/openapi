/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
	"time"
)

// checks if the FileSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileSingleAllOfAttributes{}

// FileSingleAllOfAttributes struct for FileSingleAllOfAttributes
type FileSingleAllOfAttributes struct {
	FileLocation       *string    `json:"fileLocation,omitempty"`
	FileCompression    *string    `json:"fileCompression,omitempty"`
	FileSize           *int32     `json:"fileSize,omitempty"`
	FileDataType       *string    `json:"fileDataType,omitempty"`
	FileFormat         *string    `json:"fileFormat,omitempty"`
	FileReadyTime      *time.Time `json:"fileReadyTime,omitempty"`
	FileExpirationTime *time.Time `json:"fileExpirationTime,omitempty"`
	FileContent        *string    `json:"fileContent,omitempty"`
	JobRef             *string    `json:"jobRef,omitempty"`
	JobId              *string    `json:"jobId,omitempty"`
}

// NewFileSingleAllOfAttributes instantiates a new FileSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileSingleAllOfAttributes() *FileSingleAllOfAttributes {
	this := FileSingleAllOfAttributes{}
	return &this
}

// NewFileSingleAllOfAttributesWithDefaults instantiates a new FileSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileSingleAllOfAttributesWithDefaults() *FileSingleAllOfAttributes {
	this := FileSingleAllOfAttributes{}
	return &this
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetFileLocation() string {
	if o == nil || IsNil(o.FileLocation) {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetFileLocationOk() (*string, bool) {
	if o == nil || IsNil(o.FileLocation) {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasFileLocation() bool {
	if o != nil && !IsNil(o.FileLocation) {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *FileSingleAllOfAttributes) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetFileCompression returns the FileCompression field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetFileCompression() string {
	if o == nil || IsNil(o.FileCompression) {
		var ret string
		return ret
	}
	return *o.FileCompression
}

// GetFileCompressionOk returns a tuple with the FileCompression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetFileCompressionOk() (*string, bool) {
	if o == nil || IsNil(o.FileCompression) {
		return nil, false
	}
	return o.FileCompression, true
}

// HasFileCompression returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasFileCompression() bool {
	if o != nil && !IsNil(o.FileCompression) {
		return true
	}

	return false
}

// SetFileCompression gets a reference to the given string and assigns it to the FileCompression field.
func (o *FileSingleAllOfAttributes) SetFileCompression(v string) {
	o.FileCompression = &v
}

// GetFileSize returns the FileSize field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetFileSize() int32 {
	if o == nil || IsNil(o.FileSize) {
		var ret int32
		return ret
	}
	return *o.FileSize
}

// GetFileSizeOk returns a tuple with the FileSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetFileSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FileSize) {
		return nil, false
	}
	return o.FileSize, true
}

// HasFileSize returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasFileSize() bool {
	if o != nil && !IsNil(o.FileSize) {
		return true
	}

	return false
}

// SetFileSize gets a reference to the given int32 and assigns it to the FileSize field.
func (o *FileSingleAllOfAttributes) SetFileSize(v int32) {
	o.FileSize = &v
}

// GetFileDataType returns the FileDataType field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetFileDataType() string {
	if o == nil || IsNil(o.FileDataType) {
		var ret string
		return ret
	}
	return *o.FileDataType
}

// GetFileDataTypeOk returns a tuple with the FileDataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetFileDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FileDataType) {
		return nil, false
	}
	return o.FileDataType, true
}

// HasFileDataType returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasFileDataType() bool {
	if o != nil && !IsNil(o.FileDataType) {
		return true
	}

	return false
}

// SetFileDataType gets a reference to the given string and assigns it to the FileDataType field.
func (o *FileSingleAllOfAttributes) SetFileDataType(v string) {
	o.FileDataType = &v
}

// GetFileFormat returns the FileFormat field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetFileFormat() string {
	if o == nil || IsNil(o.FileFormat) {
		var ret string
		return ret
	}
	return *o.FileFormat
}

// GetFileFormatOk returns a tuple with the FileFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetFileFormatOk() (*string, bool) {
	if o == nil || IsNil(o.FileFormat) {
		return nil, false
	}
	return o.FileFormat, true
}

// HasFileFormat returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasFileFormat() bool {
	if o != nil && !IsNil(o.FileFormat) {
		return true
	}

	return false
}

// SetFileFormat gets a reference to the given string and assigns it to the FileFormat field.
func (o *FileSingleAllOfAttributes) SetFileFormat(v string) {
	o.FileFormat = &v
}

// GetFileReadyTime returns the FileReadyTime field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetFileReadyTime() time.Time {
	if o == nil || IsNil(o.FileReadyTime) {
		var ret time.Time
		return ret
	}
	return *o.FileReadyTime
}

// GetFileReadyTimeOk returns a tuple with the FileReadyTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetFileReadyTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FileReadyTime) {
		return nil, false
	}
	return o.FileReadyTime, true
}

// HasFileReadyTime returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasFileReadyTime() bool {
	if o != nil && !IsNil(o.FileReadyTime) {
		return true
	}

	return false
}

// SetFileReadyTime gets a reference to the given time.Time and assigns it to the FileReadyTime field.
func (o *FileSingleAllOfAttributes) SetFileReadyTime(v time.Time) {
	o.FileReadyTime = &v
}

// GetFileExpirationTime returns the FileExpirationTime field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetFileExpirationTime() time.Time {
	if o == nil || IsNil(o.FileExpirationTime) {
		var ret time.Time
		return ret
	}
	return *o.FileExpirationTime
}

// GetFileExpirationTimeOk returns a tuple with the FileExpirationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetFileExpirationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.FileExpirationTime) {
		return nil, false
	}
	return o.FileExpirationTime, true
}

// HasFileExpirationTime returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasFileExpirationTime() bool {
	if o != nil && !IsNil(o.FileExpirationTime) {
		return true
	}

	return false
}

// SetFileExpirationTime gets a reference to the given time.Time and assigns it to the FileExpirationTime field.
func (o *FileSingleAllOfAttributes) SetFileExpirationTime(v time.Time) {
	o.FileExpirationTime = &v
}

// GetFileContent returns the FileContent field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetFileContent() string {
	if o == nil || IsNil(o.FileContent) {
		var ret string
		return ret
	}
	return *o.FileContent
}

// GetFileContentOk returns a tuple with the FileContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetFileContentOk() (*string, bool) {
	if o == nil || IsNil(o.FileContent) {
		return nil, false
	}
	return o.FileContent, true
}

// HasFileContent returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasFileContent() bool {
	if o != nil && !IsNil(o.FileContent) {
		return true
	}

	return false
}

// SetFileContent gets a reference to the given string and assigns it to the FileContent field.
func (o *FileSingleAllOfAttributes) SetFileContent(v string) {
	o.FileContent = &v
}

// GetJobRef returns the JobRef field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetJobRef() string {
	if o == nil || IsNil(o.JobRef) {
		var ret string
		return ret
	}
	return *o.JobRef
}

// GetJobRefOk returns a tuple with the JobRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetJobRefOk() (*string, bool) {
	if o == nil || IsNil(o.JobRef) {
		return nil, false
	}
	return o.JobRef, true
}

// HasJobRef returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasJobRef() bool {
	if o != nil && !IsNil(o.JobRef) {
		return true
	}

	return false
}

// SetJobRef gets a reference to the given string and assigns it to the JobRef field.
func (o *FileSingleAllOfAttributes) SetJobRef(v string) {
	o.JobRef = &v
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *FileSingleAllOfAttributes) GetJobId() string {
	if o == nil || IsNil(o.JobId) {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileSingleAllOfAttributes) GetJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.JobId) {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *FileSingleAllOfAttributes) HasJobId() bool {
	if o != nil && !IsNil(o.JobId) {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *FileSingleAllOfAttributes) SetJobId(v string) {
	o.JobId = &v
}

func (o FileSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileLocation) {
		toSerialize["fileLocation"] = o.FileLocation
	}
	if !IsNil(o.FileCompression) {
		toSerialize["fileCompression"] = o.FileCompression
	}
	if !IsNil(o.FileSize) {
		toSerialize["fileSize"] = o.FileSize
	}
	if !IsNil(o.FileDataType) {
		toSerialize["fileDataType"] = o.FileDataType
	}
	if !IsNil(o.FileFormat) {
		toSerialize["fileFormat"] = o.FileFormat
	}
	if !IsNil(o.FileReadyTime) {
		toSerialize["fileReadyTime"] = o.FileReadyTime
	}
	if !IsNil(o.FileExpirationTime) {
		toSerialize["fileExpirationTime"] = o.FileExpirationTime
	}
	if !IsNil(o.FileContent) {
		toSerialize["fileContent"] = o.FileContent
	}
	if !IsNil(o.JobRef) {
		toSerialize["jobRef"] = o.JobRef
	}
	if !IsNil(o.JobId) {
		toSerialize["jobId"] = o.JobId
	}
	return toSerialize, nil
}

type NullableFileSingleAllOfAttributes struct {
	value *FileSingleAllOfAttributes
	isSet bool
}

func (v NullableFileSingleAllOfAttributes) Get() *FileSingleAllOfAttributes {
	return v.value
}

func (v *NullableFileSingleAllOfAttributes) Set(val *FileSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFileSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFileSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileSingleAllOfAttributes(val *FileSingleAllOfAttributes) *NullableFileSingleAllOfAttributes {
	return &NullableFileSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableFileSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
