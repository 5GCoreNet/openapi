/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the NdccfAnalyticsSubscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NdccfAnalyticsSubscription{}

// NdccfAnalyticsSubscription Represents an Individual DCCF Analytics Subscription.
type NdccfAnalyticsSubscription struct {
	AnaSub NnwdafEventsSubscription `json:"anaSub"`
	// String providing an URI formatted according to RFC 3986.
	AnaNotifUri string `json:"anaNotifUri"`
	// Notification correlation identifier.
	AnaNotifCorrId string `json:"anaNotifCorrId"`
	FormatInstruct *FormattingInstruction `json:"formatInstruct,omitempty"`
	// Processing instructions to be used for sending event notifications.
	ProcInstructs []ProcessingInstruction `json:"procInstructs,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	TargetNfId *string `json:"targetNfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	TargetNfSetId *string `json:"targetNfSetId,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AdrfId *string `json:"adrfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.  
	ArdfSetId *string `json:"ardfSetId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	TimePeriod *TimeWindow `json:"timePeriod,omitempty"`
	// The purposes of data collection. This attribute may only be provided if user consent is  required depending on local policy and regulations and the consumer has not checked user consent. 
	DataCollectPurposes []DataCollectionPurpose `json:"dataCollectPurposes,omitempty"`
	// Indication that the NF service consumer has already checked the user consent.
	CheckedConsentInd *bool `json:"checkedConsentInd,omitempty"`
}

// NewNdccfAnalyticsSubscription instantiates a new NdccfAnalyticsSubscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNdccfAnalyticsSubscription(anaSub NnwdafEventsSubscription, anaNotifUri string, anaNotifCorrId string) *NdccfAnalyticsSubscription {
	this := NdccfAnalyticsSubscription{}
	this.AnaSub = anaSub
	this.AnaNotifUri = anaNotifUri
	this.AnaNotifCorrId = anaNotifCorrId
	return &this
}

// NewNdccfAnalyticsSubscriptionWithDefaults instantiates a new NdccfAnalyticsSubscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNdccfAnalyticsSubscriptionWithDefaults() *NdccfAnalyticsSubscription {
	this := NdccfAnalyticsSubscription{}
	return &this
}

// GetAnaSub returns the AnaSub field value
func (o *NdccfAnalyticsSubscription) GetAnaSub() NnwdafEventsSubscription {
	if o == nil {
		var ret NnwdafEventsSubscription
		return ret
	}

	return o.AnaSub
}

// GetAnaSubOk returns a tuple with the AnaSub field value
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetAnaSubOk() (*NnwdafEventsSubscription, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnaSub, true
}

// SetAnaSub sets field value
func (o *NdccfAnalyticsSubscription) SetAnaSub(v NnwdafEventsSubscription) {
	o.AnaSub = v
}

// GetAnaNotifUri returns the AnaNotifUri field value
func (o *NdccfAnalyticsSubscription) GetAnaNotifUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnaNotifUri
}

// GetAnaNotifUriOk returns a tuple with the AnaNotifUri field value
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetAnaNotifUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnaNotifUri, true
}

// SetAnaNotifUri sets field value
func (o *NdccfAnalyticsSubscription) SetAnaNotifUri(v string) {
	o.AnaNotifUri = v
}

// GetAnaNotifCorrId returns the AnaNotifCorrId field value
func (o *NdccfAnalyticsSubscription) GetAnaNotifCorrId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AnaNotifCorrId
}

// GetAnaNotifCorrIdOk returns a tuple with the AnaNotifCorrId field value
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetAnaNotifCorrIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnaNotifCorrId, true
}

// SetAnaNotifCorrId sets field value
func (o *NdccfAnalyticsSubscription) SetAnaNotifCorrId(v string) {
	o.AnaNotifCorrId = v
}

// GetFormatInstruct returns the FormatInstruct field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetFormatInstruct() FormattingInstruction {
	if o == nil || IsNil(o.FormatInstruct) {
		var ret FormattingInstruction
		return ret
	}
	return *o.FormatInstruct
}

// GetFormatInstructOk returns a tuple with the FormatInstruct field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetFormatInstructOk() (*FormattingInstruction, bool) {
	if o == nil || IsNil(o.FormatInstruct) {
		return nil, false
	}
	return o.FormatInstruct, true
}

// HasFormatInstruct returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasFormatInstruct() bool {
	if o != nil && !IsNil(o.FormatInstruct) {
		return true
	}

	return false
}

// SetFormatInstruct gets a reference to the given FormattingInstruction and assigns it to the FormatInstruct field.
func (o *NdccfAnalyticsSubscription) SetFormatInstruct(v FormattingInstruction) {
	o.FormatInstruct = &v
}

// GetProcInstructs returns the ProcInstructs field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetProcInstructs() []ProcessingInstruction {
	if o == nil || IsNil(o.ProcInstructs) {
		var ret []ProcessingInstruction
		return ret
	}
	return o.ProcInstructs
}

// GetProcInstructsOk returns a tuple with the ProcInstructs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetProcInstructsOk() ([]ProcessingInstruction, bool) {
	if o == nil || IsNil(o.ProcInstructs) {
		return nil, false
	}
	return o.ProcInstructs, true
}

// HasProcInstructs returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasProcInstructs() bool {
	if o != nil && !IsNil(o.ProcInstructs) {
		return true
	}

	return false
}

// SetProcInstructs gets a reference to the given []ProcessingInstruction and assigns it to the ProcInstructs field.
func (o *NdccfAnalyticsSubscription) SetProcInstructs(v []ProcessingInstruction) {
	o.ProcInstructs = v
}

// GetTargetNfId returns the TargetNfId field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetTargetNfId() string {
	if o == nil || IsNil(o.TargetNfId) {
		var ret string
		return ret
	}
	return *o.TargetNfId
}

// GetTargetNfIdOk returns a tuple with the TargetNfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetTargetNfIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNfId) {
		return nil, false
	}
	return o.TargetNfId, true
}

// HasTargetNfId returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasTargetNfId() bool {
	if o != nil && !IsNil(o.TargetNfId) {
		return true
	}

	return false
}

// SetTargetNfId gets a reference to the given string and assigns it to the TargetNfId field.
func (o *NdccfAnalyticsSubscription) SetTargetNfId(v string) {
	o.TargetNfId = &v
}

// GetTargetNfSetId returns the TargetNfSetId field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetTargetNfSetId() string {
	if o == nil || IsNil(o.TargetNfSetId) {
		var ret string
		return ret
	}
	return *o.TargetNfSetId
}

// GetTargetNfSetIdOk returns a tuple with the TargetNfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetTargetNfSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.TargetNfSetId) {
		return nil, false
	}
	return o.TargetNfSetId, true
}

// HasTargetNfSetId returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasTargetNfSetId() bool {
	if o != nil && !IsNil(o.TargetNfSetId) {
		return true
	}

	return false
}

// SetTargetNfSetId gets a reference to the given string and assigns it to the TargetNfSetId field.
func (o *NdccfAnalyticsSubscription) SetTargetNfSetId(v string) {
	o.TargetNfSetId = &v
}

// GetAdrfId returns the AdrfId field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetAdrfId() string {
	if o == nil || IsNil(o.AdrfId) {
		var ret string
		return ret
	}
	return *o.AdrfId
}

// GetAdrfIdOk returns a tuple with the AdrfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetAdrfIdOk() (*string, bool) {
	if o == nil || IsNil(o.AdrfId) {
		return nil, false
	}
	return o.AdrfId, true
}

// HasAdrfId returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasAdrfId() bool {
	if o != nil && !IsNil(o.AdrfId) {
		return true
	}

	return false
}

// SetAdrfId gets a reference to the given string and assigns it to the AdrfId field.
func (o *NdccfAnalyticsSubscription) SetAdrfId(v string) {
	o.AdrfId = &v
}

// GetArdfSetId returns the ArdfSetId field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetArdfSetId() string {
	if o == nil || IsNil(o.ArdfSetId) {
		var ret string
		return ret
	}
	return *o.ArdfSetId
}

// GetArdfSetIdOk returns a tuple with the ArdfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetArdfSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.ArdfSetId) {
		return nil, false
	}
	return o.ArdfSetId, true
}

// HasArdfSetId returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasArdfSetId() bool {
	if o != nil && !IsNil(o.ArdfSetId) {
		return true
	}

	return false
}

// SetArdfSetId gets a reference to the given string and assigns it to the ArdfSetId field.
func (o *NdccfAnalyticsSubscription) SetArdfSetId(v string) {
	o.ArdfSetId = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *NdccfAnalyticsSubscription) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetTimePeriod returns the TimePeriod field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetTimePeriod() TimeWindow {
	if o == nil || IsNil(o.TimePeriod) {
		var ret TimeWindow
		return ret
	}
	return *o.TimePeriod
}

// GetTimePeriodOk returns a tuple with the TimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetTimePeriodOk() (*TimeWindow, bool) {
	if o == nil || IsNil(o.TimePeriod) {
		return nil, false
	}
	return o.TimePeriod, true
}

// HasTimePeriod returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasTimePeriod() bool {
	if o != nil && !IsNil(o.TimePeriod) {
		return true
	}

	return false
}

// SetTimePeriod gets a reference to the given TimeWindow and assigns it to the TimePeriod field.
func (o *NdccfAnalyticsSubscription) SetTimePeriod(v TimeWindow) {
	o.TimePeriod = &v
}

// GetDataCollectPurposes returns the DataCollectPurposes field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetDataCollectPurposes() []DataCollectionPurpose {
	if o == nil || IsNil(o.DataCollectPurposes) {
		var ret []DataCollectionPurpose
		return ret
	}
	return o.DataCollectPurposes
}

// GetDataCollectPurposesOk returns a tuple with the DataCollectPurposes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetDataCollectPurposesOk() ([]DataCollectionPurpose, bool) {
	if o == nil || IsNil(o.DataCollectPurposes) {
		return nil, false
	}
	return o.DataCollectPurposes, true
}

// HasDataCollectPurposes returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasDataCollectPurposes() bool {
	if o != nil && !IsNil(o.DataCollectPurposes) {
		return true
	}

	return false
}

// SetDataCollectPurposes gets a reference to the given []DataCollectionPurpose and assigns it to the DataCollectPurposes field.
func (o *NdccfAnalyticsSubscription) SetDataCollectPurposes(v []DataCollectionPurpose) {
	o.DataCollectPurposes = v
}

// GetCheckedConsentInd returns the CheckedConsentInd field value if set, zero value otherwise.
func (o *NdccfAnalyticsSubscription) GetCheckedConsentInd() bool {
	if o == nil || IsNil(o.CheckedConsentInd) {
		var ret bool
		return ret
	}
	return *o.CheckedConsentInd
}

// GetCheckedConsentIndOk returns a tuple with the CheckedConsentInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NdccfAnalyticsSubscription) GetCheckedConsentIndOk() (*bool, bool) {
	if o == nil || IsNil(o.CheckedConsentInd) {
		return nil, false
	}
	return o.CheckedConsentInd, true
}

// HasCheckedConsentInd returns a boolean if a field has been set.
func (o *NdccfAnalyticsSubscription) HasCheckedConsentInd() bool {
	if o != nil && !IsNil(o.CheckedConsentInd) {
		return true
	}

	return false
}

// SetCheckedConsentInd gets a reference to the given bool and assigns it to the CheckedConsentInd field.
func (o *NdccfAnalyticsSubscription) SetCheckedConsentInd(v bool) {
	o.CheckedConsentInd = &v
}

func (o NdccfAnalyticsSubscription) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NdccfAnalyticsSubscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["anaSub"] = o.AnaSub
	toSerialize["anaNotifUri"] = o.AnaNotifUri
	toSerialize["anaNotifCorrId"] = o.AnaNotifCorrId
	if !IsNil(o.FormatInstruct) {
		toSerialize["formatInstruct"] = o.FormatInstruct
	}
	if !IsNil(o.ProcInstructs) {
		toSerialize["procInstructs"] = o.ProcInstructs
	}
	if !IsNil(o.TargetNfId) {
		toSerialize["targetNfId"] = o.TargetNfId
	}
	if !IsNil(o.TargetNfSetId) {
		toSerialize["targetNfSetId"] = o.TargetNfSetId
	}
	if !IsNil(o.AdrfId) {
		toSerialize["adrfId"] = o.AdrfId
	}
	if !IsNil(o.ArdfSetId) {
		toSerialize["ardfSetId"] = o.ArdfSetId
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !IsNil(o.TimePeriod) {
		toSerialize["timePeriod"] = o.TimePeriod
	}
	if !IsNil(o.DataCollectPurposes) {
		toSerialize["dataCollectPurposes"] = o.DataCollectPurposes
	}
	if !IsNil(o.CheckedConsentInd) {
		toSerialize["checkedConsentInd"] = o.CheckedConsentInd
	}
	return toSerialize, nil
}

type NullableNdccfAnalyticsSubscription struct {
	value *NdccfAnalyticsSubscription
	isSet bool
}

func (v NullableNdccfAnalyticsSubscription) Get() *NdccfAnalyticsSubscription {
	return v.value
}

func (v *NullableNdccfAnalyticsSubscription) Set(val *NdccfAnalyticsSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableNdccfAnalyticsSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableNdccfAnalyticsSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNdccfAnalyticsSubscription(val *NdccfAnalyticsSubscription) *NullableNdccfAnalyticsSubscription {
	return &NullableNdccfAnalyticsSubscription{value: val, isSet: true}
}

func (v NullableNdccfAnalyticsSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNdccfAnalyticsSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


