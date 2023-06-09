/*
3gpp-pfd-management

API for PFD management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PfdManagement

import (
	"encoding/json"
)

// checks if the PfdManagement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PfdManagement{}

// PfdManagement Represents a PFD management resource for a PFD management request.
type PfdManagement struct {
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	Self *string `json:"self,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// Each element uniquely identifies the PFDs for an external application identifier. Each element is identified in the map via an external application identifier as key. The response shall include successfully provisioned PFD data of application(s).
	PfdDatas map[string]PfdData `json:"pfdDatas"`
	// Supplied by the SCEF and contains the external application identifiers for which PFD(s) are not added or modified successfully. The failure reason is also included. Each element provides the related information for one or more external application identifier(s) and is identified in the map via the failure identifier as key.
	PfdReports *map[string]PfdReport `json:"pfdReports,omitempty"`
	// string formatted according to IETF RFC 3986 identifying a referenced resource.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	// Set to true by the SCS/AS to request the SCEF to send a test notification as defined in clause 5.2.5.3. Set to false or omitted otherwise.
	RequestTestNotification *bool               `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig      *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
}

// NewPfdManagement instantiates a new PfdManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPfdManagement(pfdDatas map[string]PfdData) *PfdManagement {
	this := PfdManagement{}
	this.PfdDatas = pfdDatas
	return &this
}

// NewPfdManagementWithDefaults instantiates a new PfdManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPfdManagementWithDefaults() *PfdManagement {
	this := PfdManagement{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PfdManagement) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdManagement) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PfdManagement) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *PfdManagement) SetSelf(v string) {
	o.Self = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *PfdManagement) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdManagement) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *PfdManagement) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *PfdManagement) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetPfdDatas returns the PfdDatas field value
func (o *PfdManagement) GetPfdDatas() map[string]PfdData {
	if o == nil {
		var ret map[string]PfdData
		return ret
	}

	return o.PfdDatas
}

// GetPfdDatasOk returns a tuple with the PfdDatas field value
// and a boolean to check if the value has been set.
func (o *PfdManagement) GetPfdDatasOk() (*map[string]PfdData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PfdDatas, true
}

// SetPfdDatas sets field value
func (o *PfdManagement) SetPfdDatas(v map[string]PfdData) {
	o.PfdDatas = v
}

// GetPfdReports returns the PfdReports field value if set, zero value otherwise.
func (o *PfdManagement) GetPfdReports() map[string]PfdReport {
	if o == nil || IsNil(o.PfdReports) {
		var ret map[string]PfdReport
		return ret
	}
	return *o.PfdReports
}

// GetPfdReportsOk returns a tuple with the PfdReports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdManagement) GetPfdReportsOk() (*map[string]PfdReport, bool) {
	if o == nil || IsNil(o.PfdReports) {
		return nil, false
	}
	return o.PfdReports, true
}

// HasPfdReports returns a boolean if a field has been set.
func (o *PfdManagement) HasPfdReports() bool {
	if o != nil && !IsNil(o.PfdReports) {
		return true
	}

	return false
}

// SetPfdReports gets a reference to the given map[string]PfdReport and assigns it to the PfdReports field.
func (o *PfdManagement) SetPfdReports(v map[string]PfdReport) {
	o.PfdReports = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *PfdManagement) GetNotificationDestination() string {
	if o == nil || IsNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdManagement) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationDestination) {
		return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *PfdManagement) HasNotificationDestination() bool {
	if o != nil && !IsNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *PfdManagement) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *PfdManagement) GetRequestTestNotification() bool {
	if o == nil || IsNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdManagement) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestTestNotification) {
		return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *PfdManagement) HasRequestTestNotification() bool {
	if o != nil && !IsNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *PfdManagement) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *PfdManagement) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PfdManagement) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *PfdManagement) HasWebsockNotifConfig() bool {
	if o != nil && !IsNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *PfdManagement) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

func (o PfdManagement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PfdManagement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	toSerialize["pfdDatas"] = o.PfdDatas
	// skip: pfdReports is readOnly
	if !IsNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !IsNil(o.RequestTestNotification) {
		toSerialize["requestTestNotification"] = o.RequestTestNotification
	}
	if !IsNil(o.WebsockNotifConfig) {
		toSerialize["websockNotifConfig"] = o.WebsockNotifConfig
	}
	return toSerialize, nil
}

type NullablePfdManagement struct {
	value *PfdManagement
	isSet bool
}

func (v NullablePfdManagement) Get() *PfdManagement {
	return v.value
}

func (v *NullablePfdManagement) Set(val *PfdManagement) {
	v.value = val
	v.isSet = true
}

func (v NullablePfdManagement) IsSet() bool {
	return v.isSet
}

func (v *NullablePfdManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePfdManagement(val *PfdManagement) *NullablePfdManagement {
	return &NullablePfdManagement{value: val, isSet: true}
}

func (v NullablePfdManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePfdManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
