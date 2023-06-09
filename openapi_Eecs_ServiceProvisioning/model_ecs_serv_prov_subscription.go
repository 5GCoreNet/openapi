/*
Eecs_ServiceProvisioning

API for ECS Service Provisioning. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eecs_ServiceProvisioning

import (
	"encoding/json"
	"time"
)

// checks if the ECSServProvSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ECSServProvSubscription{}

// ECSServProvSubscription Represents an individual service provisioning subscription resource.
type ECSServProvSubscription struct {
	// Represents a unique identifier of the EEC.
	EecId string `json:"eecId"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	UeId *string `json:"ueId,omitempty"`
	// Information about services the EEC wants to connect to.
	AcProfs []ACProfile `json:"acProfs,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	ExpTime *time.Time `json:"expTime,omitempty"`
	// Indicates if the EEC supports service continuity or not, also indicates which ACR scenarios are supported by the EEC.
	EecSvcContSupp []ACRScenario `json:"eecSvcContSupp,omitempty"`
	// List of connectivity information for the UE.
	ConnInfo []ConnectivityInfo `json:"connInfo,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	// Set to true by Subscriber to request the ECS to send a test notification. Set to  false or omitted otherwise.
	RequestTestNotification *bool               `json:"requestTestNotification,omitempty"`
	WebsockNotifConfig      *WebsockNotifConfig `json:"websockNotifConfig,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewECSServProvSubscription instantiates a new ECSServProvSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewECSServProvSubscription(eecId string) *ECSServProvSubscription {
	this := ECSServProvSubscription{}
	this.EecId = eecId
	return &this
}

// NewECSServProvSubscriptionWithDefaults instantiates a new ECSServProvSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewECSServProvSubscriptionWithDefaults() *ECSServProvSubscription {
	this := ECSServProvSubscription{}
	return &this
}

// GetEecId returns the EecId field value
func (o *ECSServProvSubscription) GetEecId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EecId
}

// GetEecIdOk returns a tuple with the EecId field value
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetEecIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EecId, true
}

// SetEecId sets field value
func (o *ECSServProvSubscription) SetEecId(v string) {
	o.EecId = v
}

// GetUeId returns the UeId field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetUeId() string {
	if o == nil || IsNil(o.UeId) {
		var ret string
		return ret
	}
	return *o.UeId
}

// GetUeIdOk returns a tuple with the UeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetUeIdOk() (*string, bool) {
	if o == nil || IsNil(o.UeId) {
		return nil, false
	}
	return o.UeId, true
}

// HasUeId returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasUeId() bool {
	if o != nil && !IsNil(o.UeId) {
		return true
	}

	return false
}

// SetUeId gets a reference to the given string and assigns it to the UeId field.
func (o *ECSServProvSubscription) SetUeId(v string) {
	o.UeId = &v
}

// GetAcProfs returns the AcProfs field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetAcProfs() []ACProfile {
	if o == nil || IsNil(o.AcProfs) {
		var ret []ACProfile
		return ret
	}
	return o.AcProfs
}

// GetAcProfsOk returns a tuple with the AcProfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetAcProfsOk() ([]ACProfile, bool) {
	if o == nil || IsNil(o.AcProfs) {
		return nil, false
	}
	return o.AcProfs, true
}

// HasAcProfs returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasAcProfs() bool {
	if o != nil && !IsNil(o.AcProfs) {
		return true
	}

	return false
}

// SetAcProfs gets a reference to the given []ACProfile and assigns it to the AcProfs field.
func (o *ECSServProvSubscription) SetAcProfs(v []ACProfile) {
	o.AcProfs = v
}

// GetExpTime returns the ExpTime field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetExpTime() time.Time {
	if o == nil || IsNil(o.ExpTime) {
		var ret time.Time
		return ret
	}
	return *o.ExpTime
}

// GetExpTimeOk returns a tuple with the ExpTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetExpTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpTime) {
		return nil, false
	}
	return o.ExpTime, true
}

// HasExpTime returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasExpTime() bool {
	if o != nil && !IsNil(o.ExpTime) {
		return true
	}

	return false
}

// SetExpTime gets a reference to the given time.Time and assigns it to the ExpTime field.
func (o *ECSServProvSubscription) SetExpTime(v time.Time) {
	o.ExpTime = &v
}

// GetEecSvcContSupp returns the EecSvcContSupp field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetEecSvcContSupp() []ACRScenario {
	if o == nil || IsNil(o.EecSvcContSupp) {
		var ret []ACRScenario
		return ret
	}
	return o.EecSvcContSupp
}

// GetEecSvcContSuppOk returns a tuple with the EecSvcContSupp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetEecSvcContSuppOk() ([]ACRScenario, bool) {
	if o == nil || IsNil(o.EecSvcContSupp) {
		return nil, false
	}
	return o.EecSvcContSupp, true
}

// HasEecSvcContSupp returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasEecSvcContSupp() bool {
	if o != nil && !IsNil(o.EecSvcContSupp) {
		return true
	}

	return false
}

// SetEecSvcContSupp gets a reference to the given []ACRScenario and assigns it to the EecSvcContSupp field.
func (o *ECSServProvSubscription) SetEecSvcContSupp(v []ACRScenario) {
	o.EecSvcContSupp = v
}

// GetConnInfo returns the ConnInfo field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetConnInfo() []ConnectivityInfo {
	if o == nil || IsNil(o.ConnInfo) {
		var ret []ConnectivityInfo
		return ret
	}
	return o.ConnInfo
}

// GetConnInfoOk returns a tuple with the ConnInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetConnInfoOk() ([]ConnectivityInfo, bool) {
	if o == nil || IsNil(o.ConnInfo) {
		return nil, false
	}
	return o.ConnInfo, true
}

// HasConnInfo returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasConnInfo() bool {
	if o != nil && !IsNil(o.ConnInfo) {
		return true
	}

	return false
}

// SetConnInfo gets a reference to the given []ConnectivityInfo and assigns it to the ConnInfo field.
func (o *ECSServProvSubscription) SetConnInfo(v []ConnectivityInfo) {
	o.ConnInfo = v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetNotificationDestination() string {
	if o == nil || IsNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationDestination) {
		return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasNotificationDestination() bool {
	if o != nil && !IsNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *ECSServProvSubscription) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetRequestTestNotification returns the RequestTestNotification field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetRequestTestNotification() bool {
	if o == nil || IsNil(o.RequestTestNotification) {
		var ret bool
		return ret
	}
	return *o.RequestTestNotification
}

// GetRequestTestNotificationOk returns a tuple with the RequestTestNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetRequestTestNotificationOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestTestNotification) {
		return nil, false
	}
	return o.RequestTestNotification, true
}

// HasRequestTestNotification returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasRequestTestNotification() bool {
	if o != nil && !IsNil(o.RequestTestNotification) {
		return true
	}

	return false
}

// SetRequestTestNotification gets a reference to the given bool and assigns it to the RequestTestNotification field.
func (o *ECSServProvSubscription) SetRequestTestNotification(v bool) {
	o.RequestTestNotification = &v
}

// GetWebsockNotifConfig returns the WebsockNotifConfig field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetWebsockNotifConfig() WebsockNotifConfig {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		var ret WebsockNotifConfig
		return ret
	}
	return *o.WebsockNotifConfig
}

// GetWebsockNotifConfigOk returns a tuple with the WebsockNotifConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetWebsockNotifConfigOk() (*WebsockNotifConfig, bool) {
	if o == nil || IsNil(o.WebsockNotifConfig) {
		return nil, false
	}
	return o.WebsockNotifConfig, true
}

// HasWebsockNotifConfig returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasWebsockNotifConfig() bool {
	if o != nil && !IsNil(o.WebsockNotifConfig) {
		return true
	}

	return false
}

// SetWebsockNotifConfig gets a reference to the given WebsockNotifConfig and assigns it to the WebsockNotifConfig field.
func (o *ECSServProvSubscription) SetWebsockNotifConfig(v WebsockNotifConfig) {
	o.WebsockNotifConfig = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *ECSServProvSubscription) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ECSServProvSubscription) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *ECSServProvSubscription) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *ECSServProvSubscription) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o ECSServProvSubscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ECSServProvSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eecId"] = o.EecId
	if !IsNil(o.UeId) {
		toSerialize["ueId"] = o.UeId
	}
	if !IsNil(o.AcProfs) {
		toSerialize["acProfs"] = o.AcProfs
	}
	if !IsNil(o.ExpTime) {
		toSerialize["expTime"] = o.ExpTime
	}
	if !IsNil(o.EecSvcContSupp) {
		toSerialize["eecSvcContSupp"] = o.EecSvcContSupp
	}
	if !IsNil(o.ConnInfo) {
		toSerialize["connInfo"] = o.ConnInfo
	}
	if !IsNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if !IsNil(o.RequestTestNotification) {
		toSerialize["requestTestNotification"] = o.RequestTestNotification
	}
	if !IsNil(o.WebsockNotifConfig) {
		toSerialize["websockNotifConfig"] = o.WebsockNotifConfig
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableECSServProvSubscription struct {
	value *ECSServProvSubscription
	isSet bool
}

func (v NullableECSServProvSubscription) Get() *ECSServProvSubscription {
	return v.value
}

func (v *NullableECSServProvSubscription) Set(val *ECSServProvSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableECSServProvSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableECSServProvSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableECSServProvSubscription(val *ECSServProvSubscription) *NullableECSServProvSubscription {
	return &NullableECSServProvSubscription{value: val, isSet: true}
}

func (v NullableECSServProvSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableECSServProvSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
