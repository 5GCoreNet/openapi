/*
EES Session with QoS API

API for EES Session with Qos service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_SessionWithQoS

import (
	"encoding/json"
)

// checks if the SessionWithQoSPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionWithQoSPatch{}

// SessionWithQoSPatch Represents a modification request of Individual Session with QoS Subscription.
type SessionWithQoSPatch struct {
	// Contains the flow description for the Uplink and/or Downlink IP flows.
	IpFlows []string `json:"ipFlows,omitempty"`
	// Identifies a pre-defined QoS information.
	QosReference *string `json:"qosReference,omitempty"`
	// Identifies an ordered list of pre-defined QoS information. The lower the index of the array for a given entry, the higher the priority. 
	AltQosReference []string `json:"altQosReference,omitempty"`
	// Indicates the events subscribed by the EAS.
	Events []UserPlaneEvent `json:"events,omitempty"`
	SponsorInformation *SponsorInformation `json:"sponsorInformation,omitempty"`
	QosMonInfo *QosMonitoringInformationRm `json:"qosMonInfo,omitempty"`
	// string providing an URI formatted according to IETF RFC 3986.
	NotificationDestination *string `json:"notificationDestination,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	MaxbrUl NullableString `json:"maxbrUl,omitempty"`
	// This data type is defined in the same way as the 'BitRate' data type, but with the OpenAPI 'nullable: true' property. 
	MaxbrDl NullableString `json:"maxbrDl,omitempty"`
	DisUeNotif *bool `json:"disUeNotif,omitempty"`
}

// NewSessionWithQoSPatch instantiates a new SessionWithQoSPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionWithQoSPatch() *SessionWithQoSPatch {
	this := SessionWithQoSPatch{}
	return &this
}

// NewSessionWithQoSPatchWithDefaults instantiates a new SessionWithQoSPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWithQoSPatchWithDefaults() *SessionWithQoSPatch {
	this := SessionWithQoSPatch{}
	return &this
}

// GetIpFlows returns the IpFlows field value if set, zero value otherwise.
func (o *SessionWithQoSPatch) GetIpFlows() []string {
	if o == nil || isNil(o.IpFlows) {
		var ret []string
		return ret
	}
	return o.IpFlows
}

// GetIpFlowsOk returns a tuple with the IpFlows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoSPatch) GetIpFlowsOk() ([]string, bool) {
	if o == nil || isNil(o.IpFlows) {
		return nil, false
	}
	return o.IpFlows, true
}

// HasIpFlows returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasIpFlows() bool {
	if o != nil && !isNil(o.IpFlows) {
		return true
	}

	return false
}

// SetIpFlows gets a reference to the given []string and assigns it to the IpFlows field.
func (o *SessionWithQoSPatch) SetIpFlows(v []string) {
	o.IpFlows = v
}

// GetQosReference returns the QosReference field value if set, zero value otherwise.
func (o *SessionWithQoSPatch) GetQosReference() string {
	if o == nil || isNil(o.QosReference) {
		var ret string
		return ret
	}
	return *o.QosReference
}

// GetQosReferenceOk returns a tuple with the QosReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoSPatch) GetQosReferenceOk() (*string, bool) {
	if o == nil || isNil(o.QosReference) {
		return nil, false
	}
	return o.QosReference, true
}

// HasQosReference returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasQosReference() bool {
	if o != nil && !isNil(o.QosReference) {
		return true
	}

	return false
}

// SetQosReference gets a reference to the given string and assigns it to the QosReference field.
func (o *SessionWithQoSPatch) SetQosReference(v string) {
	o.QosReference = &v
}

// GetAltQosReference returns the AltQosReference field value if set, zero value otherwise.
func (o *SessionWithQoSPatch) GetAltQosReference() []string {
	if o == nil || isNil(o.AltQosReference) {
		var ret []string
		return ret
	}
	return o.AltQosReference
}

// GetAltQosReferenceOk returns a tuple with the AltQosReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoSPatch) GetAltQosReferenceOk() ([]string, bool) {
	if o == nil || isNil(o.AltQosReference) {
		return nil, false
	}
	return o.AltQosReference, true
}

// HasAltQosReference returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasAltQosReference() bool {
	if o != nil && !isNil(o.AltQosReference) {
		return true
	}

	return false
}

// SetAltQosReference gets a reference to the given []string and assigns it to the AltQosReference field.
func (o *SessionWithQoSPatch) SetAltQosReference(v []string) {
	o.AltQosReference = v
}

// GetEvents returns the Events field value if set, zero value otherwise.
func (o *SessionWithQoSPatch) GetEvents() []UserPlaneEvent {
	if o == nil || isNil(o.Events) {
		var ret []UserPlaneEvent
		return ret
	}
	return o.Events
}

// GetEventsOk returns a tuple with the Events field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoSPatch) GetEventsOk() ([]UserPlaneEvent, bool) {
	if o == nil || isNil(o.Events) {
		return nil, false
	}
	return o.Events, true
}

// HasEvents returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasEvents() bool {
	if o != nil && !isNil(o.Events) {
		return true
	}

	return false
}

// SetEvents gets a reference to the given []UserPlaneEvent and assigns it to the Events field.
func (o *SessionWithQoSPatch) SetEvents(v []UserPlaneEvent) {
	o.Events = v
}

// GetSponsorInformation returns the SponsorInformation field value if set, zero value otherwise.
func (o *SessionWithQoSPatch) GetSponsorInformation() SponsorInformation {
	if o == nil || isNil(o.SponsorInformation) {
		var ret SponsorInformation
		return ret
	}
	return *o.SponsorInformation
}

// GetSponsorInformationOk returns a tuple with the SponsorInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoSPatch) GetSponsorInformationOk() (*SponsorInformation, bool) {
	if o == nil || isNil(o.SponsorInformation) {
		return nil, false
	}
	return o.SponsorInformation, true
}

// HasSponsorInformation returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasSponsorInformation() bool {
	if o != nil && !isNil(o.SponsorInformation) {
		return true
	}

	return false
}

// SetSponsorInformation gets a reference to the given SponsorInformation and assigns it to the SponsorInformation field.
func (o *SessionWithQoSPatch) SetSponsorInformation(v SponsorInformation) {
	o.SponsorInformation = &v
}

// GetQosMonInfo returns the QosMonInfo field value if set, zero value otherwise.
func (o *SessionWithQoSPatch) GetQosMonInfo() QosMonitoringInformationRm {
	if o == nil || isNil(o.QosMonInfo) {
		var ret QosMonitoringInformationRm
		return ret
	}
	return *o.QosMonInfo
}

// GetQosMonInfoOk returns a tuple with the QosMonInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoSPatch) GetQosMonInfoOk() (*QosMonitoringInformationRm, bool) {
	if o == nil || isNil(o.QosMonInfo) {
		return nil, false
	}
	return o.QosMonInfo, true
}

// HasQosMonInfo returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasQosMonInfo() bool {
	if o != nil && !isNil(o.QosMonInfo) {
		return true
	}

	return false
}

// SetQosMonInfo gets a reference to the given QosMonitoringInformationRm and assigns it to the QosMonInfo field.
func (o *SessionWithQoSPatch) SetQosMonInfo(v QosMonitoringInformationRm) {
	o.QosMonInfo = &v
}

// GetNotificationDestination returns the NotificationDestination field value if set, zero value otherwise.
func (o *SessionWithQoSPatch) GetNotificationDestination() string {
	if o == nil || isNil(o.NotificationDestination) {
		var ret string
		return ret
	}
	return *o.NotificationDestination
}

// GetNotificationDestinationOk returns a tuple with the NotificationDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoSPatch) GetNotificationDestinationOk() (*string, bool) {
	if o == nil || isNil(o.NotificationDestination) {
		return nil, false
	}
	return o.NotificationDestination, true
}

// HasNotificationDestination returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasNotificationDestination() bool {
	if o != nil && !isNil(o.NotificationDestination) {
		return true
	}

	return false
}

// SetNotificationDestination gets a reference to the given string and assigns it to the NotificationDestination field.
func (o *SessionWithQoSPatch) SetNotificationDestination(v string) {
	o.NotificationDestination = &v
}

// GetMaxbrUl returns the MaxbrUl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionWithQoSPatch) GetMaxbrUl() string {
	if o == nil || isNil(o.MaxbrUl.Get()) {
		var ret string
		return ret
	}
	return *o.MaxbrUl.Get()
}

// GetMaxbrUlOk returns a tuple with the MaxbrUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionWithQoSPatch) GetMaxbrUlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxbrUl.Get(), o.MaxbrUl.IsSet()
}

// HasMaxbrUl returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasMaxbrUl() bool {
	if o != nil && o.MaxbrUl.IsSet() {
		return true
	}

	return false
}

// SetMaxbrUl gets a reference to the given NullableString and assigns it to the MaxbrUl field.
func (o *SessionWithQoSPatch) SetMaxbrUl(v string) {
	o.MaxbrUl.Set(&v)
}
// SetMaxbrUlNil sets the value for MaxbrUl to be an explicit nil
func (o *SessionWithQoSPatch) SetMaxbrUlNil() {
	o.MaxbrUl.Set(nil)
}

// UnsetMaxbrUl ensures that no value is present for MaxbrUl, not even an explicit nil
func (o *SessionWithQoSPatch) UnsetMaxbrUl() {
	o.MaxbrUl.Unset()
}

// GetMaxbrDl returns the MaxbrDl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SessionWithQoSPatch) GetMaxbrDl() string {
	if o == nil || isNil(o.MaxbrDl.Get()) {
		var ret string
		return ret
	}
	return *o.MaxbrDl.Get()
}

// GetMaxbrDlOk returns a tuple with the MaxbrDl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SessionWithQoSPatch) GetMaxbrDlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxbrDl.Get(), o.MaxbrDl.IsSet()
}

// HasMaxbrDl returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasMaxbrDl() bool {
	if o != nil && o.MaxbrDl.IsSet() {
		return true
	}

	return false
}

// SetMaxbrDl gets a reference to the given NullableString and assigns it to the MaxbrDl field.
func (o *SessionWithQoSPatch) SetMaxbrDl(v string) {
	o.MaxbrDl.Set(&v)
}
// SetMaxbrDlNil sets the value for MaxbrDl to be an explicit nil
func (o *SessionWithQoSPatch) SetMaxbrDlNil() {
	o.MaxbrDl.Set(nil)
}

// UnsetMaxbrDl ensures that no value is present for MaxbrDl, not even an explicit nil
func (o *SessionWithQoSPatch) UnsetMaxbrDl() {
	o.MaxbrDl.Unset()
}

// GetDisUeNotif returns the DisUeNotif field value if set, zero value otherwise.
func (o *SessionWithQoSPatch) GetDisUeNotif() bool {
	if o == nil || isNil(o.DisUeNotif) {
		var ret bool
		return ret
	}
	return *o.DisUeNotif
}

// GetDisUeNotifOk returns a tuple with the DisUeNotif field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionWithQoSPatch) GetDisUeNotifOk() (*bool, bool) {
	if o == nil || isNil(o.DisUeNotif) {
		return nil, false
	}
	return o.DisUeNotif, true
}

// HasDisUeNotif returns a boolean if a field has been set.
func (o *SessionWithQoSPatch) HasDisUeNotif() bool {
	if o != nil && !isNil(o.DisUeNotif) {
		return true
	}

	return false
}

// SetDisUeNotif gets a reference to the given bool and assigns it to the DisUeNotif field.
func (o *SessionWithQoSPatch) SetDisUeNotif(v bool) {
	o.DisUeNotif = &v
}

func (o SessionWithQoSPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionWithQoSPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.IpFlows) {
		toSerialize["ipFlows"] = o.IpFlows
	}
	if !isNil(o.QosReference) {
		toSerialize["qosReference"] = o.QosReference
	}
	if !isNil(o.AltQosReference) {
		toSerialize["altQosReference"] = o.AltQosReference
	}
	if !isNil(o.Events) {
		toSerialize["events"] = o.Events
	}
	if !isNil(o.SponsorInformation) {
		toSerialize["sponsorInformation"] = o.SponsorInformation
	}
	if !isNil(o.QosMonInfo) {
		toSerialize["qosMonInfo"] = o.QosMonInfo
	}
	if !isNil(o.NotificationDestination) {
		toSerialize["notificationDestination"] = o.NotificationDestination
	}
	if o.MaxbrUl.IsSet() {
		toSerialize["maxbrUl"] = o.MaxbrUl.Get()
	}
	if o.MaxbrDl.IsSet() {
		toSerialize["maxbrDl"] = o.MaxbrDl.Get()
	}
	if !isNil(o.DisUeNotif) {
		toSerialize["disUeNotif"] = o.DisUeNotif
	}
	return toSerialize, nil
}

type NullableSessionWithQoSPatch struct {
	value *SessionWithQoSPatch
	isSet bool
}

func (v NullableSessionWithQoSPatch) Get() *SessionWithQoSPatch {
	return v.value
}

func (v *NullableSessionWithQoSPatch) Set(val *SessionWithQoSPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionWithQoSPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionWithQoSPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionWithQoSPatch(val *SessionWithQoSPatch) *NullableSessionWithQoSPatch {
	return &NullableSessionWithQoSPatch{value: val, isSet: true}
}

func (v NullableSessionWithQoSPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionWithQoSPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


