/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the AmfEventSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AmfEventSubscription{}

// AmfEventSubscription Represents an individual event subscription resource on AMF
type AmfEventSubscription struct {
	EventList []AmfEvent `json:"eventList"`
	// String providing an URI formatted according to RFC 3986.
	EventNotifyUri string `json:"eventNotifyUri"`
	NotifyCorrelationId string `json:"notifyCorrelationId"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	NfId string `json:"nfId"`
	// String providing an URI formatted according to RFC 3986.
	SubsChangeNotifyUri *string `json:"subsChangeNotifyUri,omitempty"`
	SubsChangeNotifyCorrelationId *string `json:"subsChangeNotifyCorrelationId,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// String identifying a group of devices network internal globally unique ID which identifies a set of IMSIs, as specified in clause 19.9 of 3GPP TS 23.003.  
	GroupId *string `json:"groupId,omitempty"`
	ExcludeSupiList []string `json:"excludeSupiList,omitempty"`
	ExcludeGpsiList []string `json:"excludeGpsiList,omitempty"`
	IncludeSupiList []string `json:"includeSupiList,omitempty"`
	IncludeGpsiList []string `json:"includeGpsiList,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	Pei *string `json:"pei,omitempty"`
	AnyUE *bool `json:"anyUE,omitempty"`
	Options *AmfEventMode `json:"options,omitempty"`
	SourceNfType *NFType `json:"sourceNfType,omitempty"`
	TermNotifyInd *bool `json:"termNotifyInd,omitempty"`
}

// NewAmfEventSubscription instantiates a new AmfEventSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmfEventSubscription(eventList []AmfEvent, eventNotifyUri string, notifyCorrelationId string, nfId string) *AmfEventSubscription {
	this := AmfEventSubscription{}
	this.EventList = eventList
	this.EventNotifyUri = eventNotifyUri
	this.NotifyCorrelationId = notifyCorrelationId
	this.NfId = nfId
	return &this
}

// NewAmfEventSubscriptionWithDefaults instantiates a new AmfEventSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmfEventSubscriptionWithDefaults() *AmfEventSubscription {
	this := AmfEventSubscription{}
	return &this
}

// GetEventList returns the EventList field value
func (o *AmfEventSubscription) GetEventList() []AmfEvent {
	if o == nil {
		var ret []AmfEvent
		return ret
	}

	return o.EventList
}

// GetEventListOk returns a tuple with the EventList field value
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetEventListOk() ([]AmfEvent, bool) {
	if o == nil {
		return nil, false
	}
	return o.EventList, true
}

// SetEventList sets field value
func (o *AmfEventSubscription) SetEventList(v []AmfEvent) {
	o.EventList = v
}

// GetEventNotifyUri returns the EventNotifyUri field value
func (o *AmfEventSubscription) GetEventNotifyUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventNotifyUri
}

// GetEventNotifyUriOk returns a tuple with the EventNotifyUri field value
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetEventNotifyUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventNotifyUri, true
}

// SetEventNotifyUri sets field value
func (o *AmfEventSubscription) SetEventNotifyUri(v string) {
	o.EventNotifyUri = v
}

// GetNotifyCorrelationId returns the NotifyCorrelationId field value
func (o *AmfEventSubscription) GetNotifyCorrelationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotifyCorrelationId
}

// GetNotifyCorrelationIdOk returns a tuple with the NotifyCorrelationId field value
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetNotifyCorrelationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotifyCorrelationId, true
}

// SetNotifyCorrelationId sets field value
func (o *AmfEventSubscription) SetNotifyCorrelationId(v string) {
	o.NotifyCorrelationId = v
}

// GetNfId returns the NfId field value
func (o *AmfEventSubscription) GetNfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NfId
}

// GetNfIdOk returns a tuple with the NfId field value
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetNfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfId, true
}

// SetNfId sets field value
func (o *AmfEventSubscription) SetNfId(v string) {
	o.NfId = v
}

// GetSubsChangeNotifyUri returns the SubsChangeNotifyUri field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetSubsChangeNotifyUri() string {
	if o == nil || IsNil(o.SubsChangeNotifyUri) {
		var ret string
		return ret
	}
	return *o.SubsChangeNotifyUri
}

// GetSubsChangeNotifyUriOk returns a tuple with the SubsChangeNotifyUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetSubsChangeNotifyUriOk() (*string, bool) {
	if o == nil || IsNil(o.SubsChangeNotifyUri) {
		return nil, false
	}
	return o.SubsChangeNotifyUri, true
}

// HasSubsChangeNotifyUri returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasSubsChangeNotifyUri() bool {
	if o != nil && !IsNil(o.SubsChangeNotifyUri) {
		return true
	}

	return false
}

// SetSubsChangeNotifyUri gets a reference to the given string and assigns it to the SubsChangeNotifyUri field.
func (o *AmfEventSubscription) SetSubsChangeNotifyUri(v string) {
	o.SubsChangeNotifyUri = &v
}

// GetSubsChangeNotifyCorrelationId returns the SubsChangeNotifyCorrelationId field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetSubsChangeNotifyCorrelationId() string {
	if o == nil || IsNil(o.SubsChangeNotifyCorrelationId) {
		var ret string
		return ret
	}
	return *o.SubsChangeNotifyCorrelationId
}

// GetSubsChangeNotifyCorrelationIdOk returns a tuple with the SubsChangeNotifyCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetSubsChangeNotifyCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubsChangeNotifyCorrelationId) {
		return nil, false
	}
	return o.SubsChangeNotifyCorrelationId, true
}

// HasSubsChangeNotifyCorrelationId returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasSubsChangeNotifyCorrelationId() bool {
	if o != nil && !IsNil(o.SubsChangeNotifyCorrelationId) {
		return true
	}

	return false
}

// SetSubsChangeNotifyCorrelationId gets a reference to the given string and assigns it to the SubsChangeNotifyCorrelationId field.
func (o *AmfEventSubscription) SetSubsChangeNotifyCorrelationId(v string) {
	o.SubsChangeNotifyCorrelationId = &v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *AmfEventSubscription) SetSupi(v string) {
	o.Supi = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AmfEventSubscription) SetGroupId(v string) {
	o.GroupId = &v
}

// GetExcludeSupiList returns the ExcludeSupiList field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetExcludeSupiList() []string {
	if o == nil || IsNil(o.ExcludeSupiList) {
		var ret []string
		return ret
	}
	return o.ExcludeSupiList
}

// GetExcludeSupiListOk returns a tuple with the ExcludeSupiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetExcludeSupiListOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeSupiList) {
		return nil, false
	}
	return o.ExcludeSupiList, true
}

// HasExcludeSupiList returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasExcludeSupiList() bool {
	if o != nil && !IsNil(o.ExcludeSupiList) {
		return true
	}

	return false
}

// SetExcludeSupiList gets a reference to the given []string and assigns it to the ExcludeSupiList field.
func (o *AmfEventSubscription) SetExcludeSupiList(v []string) {
	o.ExcludeSupiList = v
}

// GetExcludeGpsiList returns the ExcludeGpsiList field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetExcludeGpsiList() []string {
	if o == nil || IsNil(o.ExcludeGpsiList) {
		var ret []string
		return ret
	}
	return o.ExcludeGpsiList
}

// GetExcludeGpsiListOk returns a tuple with the ExcludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetExcludeGpsiListOk() ([]string, bool) {
	if o == nil || IsNil(o.ExcludeGpsiList) {
		return nil, false
	}
	return o.ExcludeGpsiList, true
}

// HasExcludeGpsiList returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasExcludeGpsiList() bool {
	if o != nil && !IsNil(o.ExcludeGpsiList) {
		return true
	}

	return false
}

// SetExcludeGpsiList gets a reference to the given []string and assigns it to the ExcludeGpsiList field.
func (o *AmfEventSubscription) SetExcludeGpsiList(v []string) {
	o.ExcludeGpsiList = v
}

// GetIncludeSupiList returns the IncludeSupiList field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetIncludeSupiList() []string {
	if o == nil || IsNil(o.IncludeSupiList) {
		var ret []string
		return ret
	}
	return o.IncludeSupiList
}

// GetIncludeSupiListOk returns a tuple with the IncludeSupiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetIncludeSupiListOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeSupiList) {
		return nil, false
	}
	return o.IncludeSupiList, true
}

// HasIncludeSupiList returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasIncludeSupiList() bool {
	if o != nil && !IsNil(o.IncludeSupiList) {
		return true
	}

	return false
}

// SetIncludeSupiList gets a reference to the given []string and assigns it to the IncludeSupiList field.
func (o *AmfEventSubscription) SetIncludeSupiList(v []string) {
	o.IncludeSupiList = v
}

// GetIncludeGpsiList returns the IncludeGpsiList field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetIncludeGpsiList() []string {
	if o == nil || IsNil(o.IncludeGpsiList) {
		var ret []string
		return ret
	}
	return o.IncludeGpsiList
}

// GetIncludeGpsiListOk returns a tuple with the IncludeGpsiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetIncludeGpsiListOk() ([]string, bool) {
	if o == nil || IsNil(o.IncludeGpsiList) {
		return nil, false
	}
	return o.IncludeGpsiList, true
}

// HasIncludeGpsiList returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasIncludeGpsiList() bool {
	if o != nil && !IsNil(o.IncludeGpsiList) {
		return true
	}

	return false
}

// SetIncludeGpsiList gets a reference to the given []string and assigns it to the IncludeGpsiList field.
func (o *AmfEventSubscription) SetIncludeGpsiList(v []string) {
	o.IncludeGpsiList = v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *AmfEventSubscription) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetPei() string {
	if o == nil || IsNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetPeiOk() (*string, bool) {
	if o == nil || IsNil(o.Pei) {
		return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasPei() bool {
	if o != nil && !IsNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *AmfEventSubscription) SetPei(v string) {
	o.Pei = &v
}

// GetAnyUE returns the AnyUE field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetAnyUE() bool {
	if o == nil || IsNil(o.AnyUE) {
		var ret bool
		return ret
	}
	return *o.AnyUE
}

// GetAnyUEOk returns a tuple with the AnyUE field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetAnyUEOk() (*bool, bool) {
	if o == nil || IsNil(o.AnyUE) {
		return nil, false
	}
	return o.AnyUE, true
}

// HasAnyUE returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasAnyUE() bool {
	if o != nil && !IsNil(o.AnyUE) {
		return true
	}

	return false
}

// SetAnyUE gets a reference to the given bool and assigns it to the AnyUE field.
func (o *AmfEventSubscription) SetAnyUE(v bool) {
	o.AnyUE = &v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetOptions() AmfEventMode {
	if o == nil || IsNil(o.Options) {
		var ret AmfEventMode
		return ret
	}
	return *o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetOptionsOk() (*AmfEventMode, bool) {
	if o == nil || IsNil(o.Options) {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given AmfEventMode and assigns it to the Options field.
func (o *AmfEventSubscription) SetOptions(v AmfEventMode) {
	o.Options = &v
}

// GetSourceNfType returns the SourceNfType field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetSourceNfType() NFType {
	if o == nil || IsNil(o.SourceNfType) {
		var ret NFType
		return ret
	}
	return *o.SourceNfType
}

// GetSourceNfTypeOk returns a tuple with the SourceNfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetSourceNfTypeOk() (*NFType, bool) {
	if o == nil || IsNil(o.SourceNfType) {
		return nil, false
	}
	return o.SourceNfType, true
}

// HasSourceNfType returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasSourceNfType() bool {
	if o != nil && !IsNil(o.SourceNfType) {
		return true
	}

	return false
}

// SetSourceNfType gets a reference to the given NFType and assigns it to the SourceNfType field.
func (o *AmfEventSubscription) SetSourceNfType(v NFType) {
	o.SourceNfType = &v
}

// GetTermNotifyInd returns the TermNotifyInd field value if set, zero value otherwise.
func (o *AmfEventSubscription) GetTermNotifyInd() bool {
	if o == nil || IsNil(o.TermNotifyInd) {
		var ret bool
		return ret
	}
	return *o.TermNotifyInd
}

// GetTermNotifyIndOk returns a tuple with the TermNotifyInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AmfEventSubscription) GetTermNotifyIndOk() (*bool, bool) {
	if o == nil || IsNil(o.TermNotifyInd) {
		return nil, false
	}
	return o.TermNotifyInd, true
}

// HasTermNotifyInd returns a boolean if a field has been set.
func (o *AmfEventSubscription) HasTermNotifyInd() bool {
	if o != nil && !IsNil(o.TermNotifyInd) {
		return true
	}

	return false
}

// SetTermNotifyInd gets a reference to the given bool and assigns it to the TermNotifyInd field.
func (o *AmfEventSubscription) SetTermNotifyInd(v bool) {
	o.TermNotifyInd = &v
}

func (o AmfEventSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AmfEventSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eventList"] = o.EventList
	toSerialize["eventNotifyUri"] = o.EventNotifyUri
	toSerialize["notifyCorrelationId"] = o.NotifyCorrelationId
	toSerialize["nfId"] = o.NfId
	if !IsNil(o.SubsChangeNotifyUri) {
		toSerialize["subsChangeNotifyUri"] = o.SubsChangeNotifyUri
	}
	if !IsNil(o.SubsChangeNotifyCorrelationId) {
		toSerialize["subsChangeNotifyCorrelationId"] = o.SubsChangeNotifyCorrelationId
	}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.GroupId) {
		toSerialize["groupId"] = o.GroupId
	}
	if !IsNil(o.ExcludeSupiList) {
		toSerialize["excludeSupiList"] = o.ExcludeSupiList
	}
	if !IsNil(o.ExcludeGpsiList) {
		toSerialize["excludeGpsiList"] = o.ExcludeGpsiList
	}
	if !IsNil(o.IncludeSupiList) {
		toSerialize["includeSupiList"] = o.IncludeSupiList
	}
	if !IsNil(o.IncludeGpsiList) {
		toSerialize["includeGpsiList"] = o.IncludeGpsiList
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if !IsNil(o.AnyUE) {
		toSerialize["anyUE"] = o.AnyUE
	}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.SourceNfType) {
		toSerialize["sourceNfType"] = o.SourceNfType
	}
	if !IsNil(o.TermNotifyInd) {
		toSerialize["termNotifyInd"] = o.TermNotifyInd
	}
	return toSerialize, nil
}

type NullableAmfEventSubscription struct {
	value *AmfEventSubscription
	isSet bool
}

func (v NullableAmfEventSubscription) Get() *AmfEventSubscription {
	return v.value
}

func (v *NullableAmfEventSubscription) Set(val *AmfEventSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventSubscription(val *AmfEventSubscription) *NullableAmfEventSubscription {
	return &NullableAmfEventSubscription{value: val, isSet: true}
}

func (v NullableAmfEventSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


