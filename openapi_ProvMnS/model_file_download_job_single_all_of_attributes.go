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

// checks if the FileDownloadJobSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileDownloadJobSingleAllOfAttributes{}

// FileDownloadJobSingleAllOfAttributes struct for FileDownloadJobSingleAllOfAttributes
type FileDownloadJobSingleAllOfAttributes struct {
	FileLocation *string `json:"fileLocation,omitempty"`
	NotificationRecipientAddress *string `json:"notificationRecipientAddress,omitempty"`
	CancelJob *string `json:"cancelJob,omitempty"`
	JobMonitor *FileDownloadJobProcessMonitor `json:"jobMonitor,omitempty"`
}

// NewFileDownloadJobSingleAllOfAttributes instantiates a new FileDownloadJobSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileDownloadJobSingleAllOfAttributes() *FileDownloadJobSingleAllOfAttributes {
	this := FileDownloadJobSingleAllOfAttributes{}
	return &this
}

// NewFileDownloadJobSingleAllOfAttributesWithDefaults instantiates a new FileDownloadJobSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileDownloadJobSingleAllOfAttributesWithDefaults() *FileDownloadJobSingleAllOfAttributes {
	this := FileDownloadJobSingleAllOfAttributes{}
	return &this
}

// GetFileLocation returns the FileLocation field value if set, zero value otherwise.
func (o *FileDownloadJobSingleAllOfAttributes) GetFileLocation() string {
	if o == nil || IsNil(o.FileLocation) {
		var ret string
		return ret
	}
	return *o.FileLocation
}

// GetFileLocationOk returns a tuple with the FileLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingleAllOfAttributes) GetFileLocationOk() (*string, bool) {
	if o == nil || IsNil(o.FileLocation) {
		return nil, false
	}
	return o.FileLocation, true
}

// HasFileLocation returns a boolean if a field has been set.
func (o *FileDownloadJobSingleAllOfAttributes) HasFileLocation() bool {
	if o != nil && !IsNil(o.FileLocation) {
		return true
	}

	return false
}

// SetFileLocation gets a reference to the given string and assigns it to the FileLocation field.
func (o *FileDownloadJobSingleAllOfAttributes) SetFileLocation(v string) {
	o.FileLocation = &v
}

// GetNotificationRecipientAddress returns the NotificationRecipientAddress field value if set, zero value otherwise.
func (o *FileDownloadJobSingleAllOfAttributes) GetNotificationRecipientAddress() string {
	if o == nil || IsNil(o.NotificationRecipientAddress) {
		var ret string
		return ret
	}
	return *o.NotificationRecipientAddress
}

// GetNotificationRecipientAddressOk returns a tuple with the NotificationRecipientAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingleAllOfAttributes) GetNotificationRecipientAddressOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationRecipientAddress) {
		return nil, false
	}
	return o.NotificationRecipientAddress, true
}

// HasNotificationRecipientAddress returns a boolean if a field has been set.
func (o *FileDownloadJobSingleAllOfAttributes) HasNotificationRecipientAddress() bool {
	if o != nil && !IsNil(o.NotificationRecipientAddress) {
		return true
	}

	return false
}

// SetNotificationRecipientAddress gets a reference to the given string and assigns it to the NotificationRecipientAddress field.
func (o *FileDownloadJobSingleAllOfAttributes) SetNotificationRecipientAddress(v string) {
	o.NotificationRecipientAddress = &v
}

// GetCancelJob returns the CancelJob field value if set, zero value otherwise.
func (o *FileDownloadJobSingleAllOfAttributes) GetCancelJob() string {
	if o == nil || IsNil(o.CancelJob) {
		var ret string
		return ret
	}
	return *o.CancelJob
}

// GetCancelJobOk returns a tuple with the CancelJob field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingleAllOfAttributes) GetCancelJobOk() (*string, bool) {
	if o == nil || IsNil(o.CancelJob) {
		return nil, false
	}
	return o.CancelJob, true
}

// HasCancelJob returns a boolean if a field has been set.
func (o *FileDownloadJobSingleAllOfAttributes) HasCancelJob() bool {
	if o != nil && !IsNil(o.CancelJob) {
		return true
	}

	return false
}

// SetCancelJob gets a reference to the given string and assigns it to the CancelJob field.
func (o *FileDownloadJobSingleAllOfAttributes) SetCancelJob(v string) {
	o.CancelJob = &v
}

// GetJobMonitor returns the JobMonitor field value if set, zero value otherwise.
func (o *FileDownloadJobSingleAllOfAttributes) GetJobMonitor() FileDownloadJobProcessMonitor {
	if o == nil || IsNil(o.JobMonitor) {
		var ret FileDownloadJobProcessMonitor
		return ret
	}
	return *o.JobMonitor
}

// GetJobMonitorOk returns a tuple with the JobMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileDownloadJobSingleAllOfAttributes) GetJobMonitorOk() (*FileDownloadJobProcessMonitor, bool) {
	if o == nil || IsNil(o.JobMonitor) {
		return nil, false
	}
	return o.JobMonitor, true
}

// HasJobMonitor returns a boolean if a field has been set.
func (o *FileDownloadJobSingleAllOfAttributes) HasJobMonitor() bool {
	if o != nil && !IsNil(o.JobMonitor) {
		return true
	}

	return false
}

// SetJobMonitor gets a reference to the given FileDownloadJobProcessMonitor and assigns it to the JobMonitor field.
func (o *FileDownloadJobSingleAllOfAttributes) SetJobMonitor(v FileDownloadJobProcessMonitor) {
	o.JobMonitor = &v
}

func (o FileDownloadJobSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileDownloadJobSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileLocation) {
		toSerialize["fileLocation"] = o.FileLocation
	}
	if !IsNil(o.NotificationRecipientAddress) {
		toSerialize["notificationRecipientAddress"] = o.NotificationRecipientAddress
	}
	if !IsNil(o.CancelJob) {
		toSerialize["cancelJob"] = o.CancelJob
	}
	if !IsNil(o.JobMonitor) {
		toSerialize["jobMonitor"] = o.JobMonitor
	}
	return toSerialize, nil
}

type NullableFileDownloadJobSingleAllOfAttributes struct {
	value *FileDownloadJobSingleAllOfAttributes
	isSet bool
}

func (v NullableFileDownloadJobSingleAllOfAttributes) Get() *FileDownloadJobSingleAllOfAttributes {
	return v.value
}

func (v *NullableFileDownloadJobSingleAllOfAttributes) Set(val *FileDownloadJobSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableFileDownloadJobSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableFileDownloadJobSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileDownloadJobSingleAllOfAttributes(val *FileDownloadJobSingleAllOfAttributes) *NullableFileDownloadJobSingleAllOfAttributes {
	return &NullableFileDownloadJobSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableFileDownloadJobSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileDownloadJobSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

