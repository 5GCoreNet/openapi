/*
Npcf_AMPolicyControl

Access and Mobility Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_AMPolicyControl

import (
	"encoding/json"
)

// checks if the PolicyAssociationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAssociationUpdateRequest{}

// PolicyAssociationUpdateRequest Represents information that the NF service consumer provides when requesting the update of a policy association.
type PolicyAssociationUpdateRequest struct {
	// String providing an URI formatted according to RFC 3986.
	NotificationUri *string `json:"notificationUri,omitempty"`
	// Alternate or backup IPv4 Address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 Address(es) where to send Notifications.
	AltNotifIpv6Addrs []Ipv6Addr `json:"altNotifIpv6Addrs,omitempty"`
	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty"`
	// Request Triggers that the NF service consumer observes.
	Triggers      []RequestTrigger                `json:"triggers,omitempty"`
	ServAreaRes   *ServiceAreaRestriction         `json:"servAreaRes,omitempty"`
	WlServAreaRes *WirelineServiceAreaRestriction `json:"wlServAreaRes,omitempty"`
	// Unsigned integer representing the \"Subscriber Profile ID for RAT/Frequency Priority\"  as specified in 3GPP TS 36.413.
	Rfsp       *int32                   `json:"rfsp,omitempty"`
	SmfSelInfo NullableSmfSelectionData `json:"smfSelInfo,omitempty"`
	UeAmbr     *Ambr                    `json:"ueAmbr,omitempty"`
	// The subscribed UE-Slice-MBR for each subscribed S-NSSAI of the home PLMN mapping to a S-NSSAI of the serving PLMN Shall be provided for the \"UE_SLICE_MBR_CH\" policy control request trigger.
	UeSliceMbrs []UeSliceMbr `json:"ueSliceMbrs,omitempty"`
	// Contains the UE presence status for tracking area for which changes of the UE presence occurred. The praId attribute within the PresenceInfo data type is the key of the map.
	PraStatuses *map[string]PresenceInfo `json:"praStatuses,omitempty"`
	UserLoc     *UserLocation            `json:"userLoc,omitempty"`
	// array of allowed S-NSSAIs for the 3GPP access.
	AllowedSnssais []Snssai `json:"allowedSnssais,omitempty"`
	// array of target S-NSSAIs.
	TargetSnssais []Snssai `json:"targetSnssais,omitempty"`
	// mapping of each S-NSSAI of the Allowed NSSAI to the corresponding S-NSSAI of the HPLMN.
	MappingSnssais []MappingOfSnssai `json:"mappingSnssais,omitempty"`
	AccessTypes    []AccessType      `json:"accessTypes,omitempty"`
	RatTypes       []RatType         `json:"ratTypes,omitempty"`
	// array of allowed S-NSSAIs for the Non-3GPP access.
	N3gAllowedSnssais []Snssai          `json:"n3gAllowedSnssais,omitempty"`
	TraceReq          NullableTraceData `json:"traceReq,omitempty"`
	Guami             *Guami            `json:"guami,omitempty"`
	NwdafDatas        []NwdafData       `json:"nwdafDatas,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewPolicyAssociationUpdateRequest instantiates a new PolicyAssociationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAssociationUpdateRequest() *PolicyAssociationUpdateRequest {
	this := PolicyAssociationUpdateRequest{}
	return &this
}

// NewPolicyAssociationUpdateRequestWithDefaults instantiates a new PolicyAssociationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAssociationUpdateRequestWithDefaults() *PolicyAssociationUpdateRequest {
	this := PolicyAssociationUpdateRequest{}
	return &this
}

// GetNotificationUri returns the NotificationUri field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetNotificationUri() string {
	if o == nil || IsNil(o.NotificationUri) {
		var ret string
		return ret
	}
	return *o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetNotificationUriOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationUri) {
		return nil, false
	}
	return o.NotificationUri, true
}

// HasNotificationUri returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasNotificationUri() bool {
	if o != nil && !IsNil(o.NotificationUri) {
		return true
	}

	return false
}

// SetNotificationUri gets a reference to the given string and assigns it to the NotificationUri field.
func (o *PolicyAssociationUpdateRequest) SetNotificationUri(v string) {
	o.NotificationUri = &v
}

// GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv4Addrs() []string {
	if o == nil || IsNil(o.AltNotifIpv4Addrs) {
		var ret []string
		return ret
	}
	return o.AltNotifIpv4Addrs
}

// GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv4AddrsOk() ([]string, bool) {
	if o == nil || IsNil(o.AltNotifIpv4Addrs) {
		return nil, false
	}
	return o.AltNotifIpv4Addrs, true
}

// HasAltNotifIpv4Addrs returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasAltNotifIpv4Addrs() bool {
	if o != nil && !IsNil(o.AltNotifIpv4Addrs) {
		return true
	}

	return false
}

// SetAltNotifIpv4Addrs gets a reference to the given []string and assigns it to the AltNotifIpv4Addrs field.
func (o *PolicyAssociationUpdateRequest) SetAltNotifIpv4Addrs(v []string) {
	o.AltNotifIpv4Addrs = v
}

// GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv6Addrs() []Ipv6Addr {
	if o == nil || IsNil(o.AltNotifIpv6Addrs) {
		var ret []Ipv6Addr
		return ret
	}
	return o.AltNotifIpv6Addrs
}

// GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetAltNotifIpv6AddrsOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.AltNotifIpv6Addrs) {
		return nil, false
	}
	return o.AltNotifIpv6Addrs, true
}

// HasAltNotifIpv6Addrs returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasAltNotifIpv6Addrs() bool {
	if o != nil && !IsNil(o.AltNotifIpv6Addrs) {
		return true
	}

	return false
}

// SetAltNotifIpv6Addrs gets a reference to the given []Ipv6Addr and assigns it to the AltNotifIpv6Addrs field.
func (o *PolicyAssociationUpdateRequest) SetAltNotifIpv6Addrs(v []Ipv6Addr) {
	o.AltNotifIpv6Addrs = v
}

// GetAltNotifFqdns returns the AltNotifFqdns field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetAltNotifFqdns() []string {
	if o == nil || IsNil(o.AltNotifFqdns) {
		var ret []string
		return ret
	}
	return o.AltNotifFqdns
}

// GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetAltNotifFqdnsOk() ([]string, bool) {
	if o == nil || IsNil(o.AltNotifFqdns) {
		return nil, false
	}
	return o.AltNotifFqdns, true
}

// HasAltNotifFqdns returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasAltNotifFqdns() bool {
	if o != nil && !IsNil(o.AltNotifFqdns) {
		return true
	}

	return false
}

// SetAltNotifFqdns gets a reference to the given []string and assigns it to the AltNotifFqdns field.
func (o *PolicyAssociationUpdateRequest) SetAltNotifFqdns(v []string) {
	o.AltNotifFqdns = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetTriggers() []RequestTrigger {
	if o == nil || IsNil(o.Triggers) {
		var ret []RequestTrigger
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetTriggersOk() ([]RequestTrigger, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []RequestTrigger and assigns it to the Triggers field.
func (o *PolicyAssociationUpdateRequest) SetTriggers(v []RequestTrigger) {
	o.Triggers = v
}

// GetServAreaRes returns the ServAreaRes field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetServAreaRes() ServiceAreaRestriction {
	if o == nil || IsNil(o.ServAreaRes) {
		var ret ServiceAreaRestriction
		return ret
	}
	return *o.ServAreaRes
}

// GetServAreaResOk returns a tuple with the ServAreaRes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetServAreaResOk() (*ServiceAreaRestriction, bool) {
	if o == nil || IsNil(o.ServAreaRes) {
		return nil, false
	}
	return o.ServAreaRes, true
}

// HasServAreaRes returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasServAreaRes() bool {
	if o != nil && !IsNil(o.ServAreaRes) {
		return true
	}

	return false
}

// SetServAreaRes gets a reference to the given ServiceAreaRestriction and assigns it to the ServAreaRes field.
func (o *PolicyAssociationUpdateRequest) SetServAreaRes(v ServiceAreaRestriction) {
	o.ServAreaRes = &v
}

// GetWlServAreaRes returns the WlServAreaRes field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetWlServAreaRes() WirelineServiceAreaRestriction {
	if o == nil || IsNil(o.WlServAreaRes) {
		var ret WirelineServiceAreaRestriction
		return ret
	}
	return *o.WlServAreaRes
}

// GetWlServAreaResOk returns a tuple with the WlServAreaRes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetWlServAreaResOk() (*WirelineServiceAreaRestriction, bool) {
	if o == nil || IsNil(o.WlServAreaRes) {
		return nil, false
	}
	return o.WlServAreaRes, true
}

// HasWlServAreaRes returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasWlServAreaRes() bool {
	if o != nil && !IsNil(o.WlServAreaRes) {
		return true
	}

	return false
}

// SetWlServAreaRes gets a reference to the given WirelineServiceAreaRestriction and assigns it to the WlServAreaRes field.
func (o *PolicyAssociationUpdateRequest) SetWlServAreaRes(v WirelineServiceAreaRestriction) {
	o.WlServAreaRes = &v
}

// GetRfsp returns the Rfsp field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetRfsp() int32 {
	if o == nil || IsNil(o.Rfsp) {
		var ret int32
		return ret
	}
	return *o.Rfsp
}

// GetRfspOk returns a tuple with the Rfsp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetRfspOk() (*int32, bool) {
	if o == nil || IsNil(o.Rfsp) {
		return nil, false
	}
	return o.Rfsp, true
}

// HasRfsp returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasRfsp() bool {
	if o != nil && !IsNil(o.Rfsp) {
		return true
	}

	return false
}

// SetRfsp gets a reference to the given int32 and assigns it to the Rfsp field.
func (o *PolicyAssociationUpdateRequest) SetRfsp(v int32) {
	o.Rfsp = &v
}

// GetSmfSelInfo returns the SmfSelInfo field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAssociationUpdateRequest) GetSmfSelInfo() SmfSelectionData {
	if o == nil || IsNil(o.SmfSelInfo.Get()) {
		var ret SmfSelectionData
		return ret
	}
	return *o.SmfSelInfo.Get()
}

// GetSmfSelInfoOk returns a tuple with the SmfSelInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAssociationUpdateRequest) GetSmfSelInfoOk() (*SmfSelectionData, bool) {
	if o == nil {
		return nil, false
	}
	return o.SmfSelInfo.Get(), o.SmfSelInfo.IsSet()
}

// HasSmfSelInfo returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasSmfSelInfo() bool {
	if o != nil && o.SmfSelInfo.IsSet() {
		return true
	}

	return false
}

// SetSmfSelInfo gets a reference to the given NullableSmfSelectionData and assigns it to the SmfSelInfo field.
func (o *PolicyAssociationUpdateRequest) SetSmfSelInfo(v SmfSelectionData) {
	o.SmfSelInfo.Set(&v)
}

// SetSmfSelInfoNil sets the value for SmfSelInfo to be an explicit nil
func (o *PolicyAssociationUpdateRequest) SetSmfSelInfoNil() {
	o.SmfSelInfo.Set(nil)
}

// UnsetSmfSelInfo ensures that no value is present for SmfSelInfo, not even an explicit nil
func (o *PolicyAssociationUpdateRequest) UnsetSmfSelInfo() {
	o.SmfSelInfo.Unset()
}

// GetUeAmbr returns the UeAmbr field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetUeAmbr() Ambr {
	if o == nil || IsNil(o.UeAmbr) {
		var ret Ambr
		return ret
	}
	return *o.UeAmbr
}

// GetUeAmbrOk returns a tuple with the UeAmbr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetUeAmbrOk() (*Ambr, bool) {
	if o == nil || IsNil(o.UeAmbr) {
		return nil, false
	}
	return o.UeAmbr, true
}

// HasUeAmbr returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasUeAmbr() bool {
	if o != nil && !IsNil(o.UeAmbr) {
		return true
	}

	return false
}

// SetUeAmbr gets a reference to the given Ambr and assigns it to the UeAmbr field.
func (o *PolicyAssociationUpdateRequest) SetUeAmbr(v Ambr) {
	o.UeAmbr = &v
}

// GetUeSliceMbrs returns the UeSliceMbrs field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetUeSliceMbrs() []UeSliceMbr {
	if o == nil || IsNil(o.UeSliceMbrs) {
		var ret []UeSliceMbr
		return ret
	}
	return o.UeSliceMbrs
}

// GetUeSliceMbrsOk returns a tuple with the UeSliceMbrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetUeSliceMbrsOk() ([]UeSliceMbr, bool) {
	if o == nil || IsNil(o.UeSliceMbrs) {
		return nil, false
	}
	return o.UeSliceMbrs, true
}

// HasUeSliceMbrs returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasUeSliceMbrs() bool {
	if o != nil && !IsNil(o.UeSliceMbrs) {
		return true
	}

	return false
}

// SetUeSliceMbrs gets a reference to the given []UeSliceMbr and assigns it to the UeSliceMbrs field.
func (o *PolicyAssociationUpdateRequest) SetUeSliceMbrs(v []UeSliceMbr) {
	o.UeSliceMbrs = v
}

// GetPraStatuses returns the PraStatuses field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetPraStatuses() map[string]PresenceInfo {
	if o == nil || IsNil(o.PraStatuses) {
		var ret map[string]PresenceInfo
		return ret
	}
	return *o.PraStatuses
}

// GetPraStatusesOk returns a tuple with the PraStatuses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetPraStatusesOk() (*map[string]PresenceInfo, bool) {
	if o == nil || IsNil(o.PraStatuses) {
		return nil, false
	}
	return o.PraStatuses, true
}

// HasPraStatuses returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasPraStatuses() bool {
	if o != nil && !IsNil(o.PraStatuses) {
		return true
	}

	return false
}

// SetPraStatuses gets a reference to the given map[string]PresenceInfo and assigns it to the PraStatuses field.
func (o *PolicyAssociationUpdateRequest) SetPraStatuses(v map[string]PresenceInfo) {
	o.PraStatuses = &v
}

// GetUserLoc returns the UserLoc field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetUserLoc() UserLocation {
	if o == nil || IsNil(o.UserLoc) {
		var ret UserLocation
		return ret
	}
	return *o.UserLoc
}

// GetUserLocOk returns a tuple with the UserLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetUserLocOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLoc) {
		return nil, false
	}
	return o.UserLoc, true
}

// HasUserLoc returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasUserLoc() bool {
	if o != nil && !IsNil(o.UserLoc) {
		return true
	}

	return false
}

// SetUserLoc gets a reference to the given UserLocation and assigns it to the UserLoc field.
func (o *PolicyAssociationUpdateRequest) SetUserLoc(v UserLocation) {
	o.UserLoc = &v
}

// GetAllowedSnssais returns the AllowedSnssais field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetAllowedSnssais() []Snssai {
	if o == nil || IsNil(o.AllowedSnssais) {
		var ret []Snssai
		return ret
	}
	return o.AllowedSnssais
}

// GetAllowedSnssaisOk returns a tuple with the AllowedSnssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetAllowedSnssaisOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.AllowedSnssais) {
		return nil, false
	}
	return o.AllowedSnssais, true
}

// HasAllowedSnssais returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasAllowedSnssais() bool {
	if o != nil && !IsNil(o.AllowedSnssais) {
		return true
	}

	return false
}

// SetAllowedSnssais gets a reference to the given []Snssai and assigns it to the AllowedSnssais field.
func (o *PolicyAssociationUpdateRequest) SetAllowedSnssais(v []Snssai) {
	o.AllowedSnssais = v
}

// GetTargetSnssais returns the TargetSnssais field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetTargetSnssais() []Snssai {
	if o == nil || IsNil(o.TargetSnssais) {
		var ret []Snssai
		return ret
	}
	return o.TargetSnssais
}

// GetTargetSnssaisOk returns a tuple with the TargetSnssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetTargetSnssaisOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.TargetSnssais) {
		return nil, false
	}
	return o.TargetSnssais, true
}

// HasTargetSnssais returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasTargetSnssais() bool {
	if o != nil && !IsNil(o.TargetSnssais) {
		return true
	}

	return false
}

// SetTargetSnssais gets a reference to the given []Snssai and assigns it to the TargetSnssais field.
func (o *PolicyAssociationUpdateRequest) SetTargetSnssais(v []Snssai) {
	o.TargetSnssais = v
}

// GetMappingSnssais returns the MappingSnssais field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetMappingSnssais() []MappingOfSnssai {
	if o == nil || IsNil(o.MappingSnssais) {
		var ret []MappingOfSnssai
		return ret
	}
	return o.MappingSnssais
}

// GetMappingSnssaisOk returns a tuple with the MappingSnssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetMappingSnssaisOk() ([]MappingOfSnssai, bool) {
	if o == nil || IsNil(o.MappingSnssais) {
		return nil, false
	}
	return o.MappingSnssais, true
}

// HasMappingSnssais returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasMappingSnssais() bool {
	if o != nil && !IsNil(o.MappingSnssais) {
		return true
	}

	return false
}

// SetMappingSnssais gets a reference to the given []MappingOfSnssai and assigns it to the MappingSnssais field.
func (o *PolicyAssociationUpdateRequest) SetMappingSnssais(v []MappingOfSnssai) {
	o.MappingSnssais = v
}

// GetAccessTypes returns the AccessTypes field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetAccessTypes() []AccessType {
	if o == nil || IsNil(o.AccessTypes) {
		var ret []AccessType
		return ret
	}
	return o.AccessTypes
}

// GetAccessTypesOk returns a tuple with the AccessTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetAccessTypesOk() ([]AccessType, bool) {
	if o == nil || IsNil(o.AccessTypes) {
		return nil, false
	}
	return o.AccessTypes, true
}

// HasAccessTypes returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasAccessTypes() bool {
	if o != nil && !IsNil(o.AccessTypes) {
		return true
	}

	return false
}

// SetAccessTypes gets a reference to the given []AccessType and assigns it to the AccessTypes field.
func (o *PolicyAssociationUpdateRequest) SetAccessTypes(v []AccessType) {
	o.AccessTypes = v
}

// GetRatTypes returns the RatTypes field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetRatTypes() []RatType {
	if o == nil || IsNil(o.RatTypes) {
		var ret []RatType
		return ret
	}
	return o.RatTypes
}

// GetRatTypesOk returns a tuple with the RatTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetRatTypesOk() ([]RatType, bool) {
	if o == nil || IsNil(o.RatTypes) {
		return nil, false
	}
	return o.RatTypes, true
}

// HasRatTypes returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasRatTypes() bool {
	if o != nil && !IsNil(o.RatTypes) {
		return true
	}

	return false
}

// SetRatTypes gets a reference to the given []RatType and assigns it to the RatTypes field.
func (o *PolicyAssociationUpdateRequest) SetRatTypes(v []RatType) {
	o.RatTypes = v
}

// GetN3gAllowedSnssais returns the N3gAllowedSnssais field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetN3gAllowedSnssais() []Snssai {
	if o == nil || IsNil(o.N3gAllowedSnssais) {
		var ret []Snssai
		return ret
	}
	return o.N3gAllowedSnssais
}

// GetN3gAllowedSnssaisOk returns a tuple with the N3gAllowedSnssais field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetN3gAllowedSnssaisOk() ([]Snssai, bool) {
	if o == nil || IsNil(o.N3gAllowedSnssais) {
		return nil, false
	}
	return o.N3gAllowedSnssais, true
}

// HasN3gAllowedSnssais returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasN3gAllowedSnssais() bool {
	if o != nil && !IsNil(o.N3gAllowedSnssais) {
		return true
	}

	return false
}

// SetN3gAllowedSnssais gets a reference to the given []Snssai and assigns it to the N3gAllowedSnssais field.
func (o *PolicyAssociationUpdateRequest) SetN3gAllowedSnssais(v []Snssai) {
	o.N3gAllowedSnssais = v
}

// GetTraceReq returns the TraceReq field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAssociationUpdateRequest) GetTraceReq() TraceData {
	if o == nil || IsNil(o.TraceReq.Get()) {
		var ret TraceData
		return ret
	}
	return *o.TraceReq.Get()
}

// GetTraceReqOk returns a tuple with the TraceReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAssociationUpdateRequest) GetTraceReqOk() (*TraceData, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceReq.Get(), o.TraceReq.IsSet()
}

// HasTraceReq returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasTraceReq() bool {
	if o != nil && o.TraceReq.IsSet() {
		return true
	}

	return false
}

// SetTraceReq gets a reference to the given NullableTraceData and assigns it to the TraceReq field.
func (o *PolicyAssociationUpdateRequest) SetTraceReq(v TraceData) {
	o.TraceReq.Set(&v)
}

// SetTraceReqNil sets the value for TraceReq to be an explicit nil
func (o *PolicyAssociationUpdateRequest) SetTraceReqNil() {
	o.TraceReq.Set(nil)
}

// UnsetTraceReq ensures that no value is present for TraceReq, not even an explicit nil
func (o *PolicyAssociationUpdateRequest) UnsetTraceReq() {
	o.TraceReq.Unset()
}

// GetGuami returns the Guami field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetGuami() Guami {
	if o == nil || IsNil(o.Guami) {
		var ret Guami
		return ret
	}
	return *o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetGuamiOk() (*Guami, bool) {
	if o == nil || IsNil(o.Guami) {
		return nil, false
	}
	return o.Guami, true
}

// HasGuami returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasGuami() bool {
	if o != nil && !IsNil(o.Guami) {
		return true
	}

	return false
}

// SetGuami gets a reference to the given Guami and assigns it to the Guami field.
func (o *PolicyAssociationUpdateRequest) SetGuami(v Guami) {
	o.Guami = &v
}

// GetNwdafDatas returns the NwdafDatas field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PolicyAssociationUpdateRequest) GetNwdafDatas() []NwdafData {
	if o == nil {
		var ret []NwdafData
		return ret
	}
	return o.NwdafDatas
}

// GetNwdafDatasOk returns a tuple with the NwdafDatas field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PolicyAssociationUpdateRequest) GetNwdafDatasOk() ([]NwdafData, bool) {
	if o == nil || IsNil(o.NwdafDatas) {
		return nil, false
	}
	return o.NwdafDatas, true
}

// HasNwdafDatas returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasNwdafDatas() bool {
	if o != nil && IsNil(o.NwdafDatas) {
		return true
	}

	return false
}

// SetNwdafDatas gets a reference to the given []NwdafData and assigns it to the NwdafDatas field.
func (o *PolicyAssociationUpdateRequest) SetNwdafDatas(v []NwdafData) {
	o.NwdafDatas = v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *PolicyAssociationUpdateRequest) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationUpdateRequest) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *PolicyAssociationUpdateRequest) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *PolicyAssociationUpdateRequest) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o PolicyAssociationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAssociationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationUri) {
		toSerialize["notificationUri"] = o.NotificationUri
	}
	if !IsNil(o.AltNotifIpv4Addrs) {
		toSerialize["altNotifIpv4Addrs"] = o.AltNotifIpv4Addrs
	}
	if !IsNil(o.AltNotifIpv6Addrs) {
		toSerialize["altNotifIpv6Addrs"] = o.AltNotifIpv6Addrs
	}
	if !IsNil(o.AltNotifFqdns) {
		toSerialize["altNotifFqdns"] = o.AltNotifFqdns
	}
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	if !IsNil(o.ServAreaRes) {
		toSerialize["servAreaRes"] = o.ServAreaRes
	}
	if !IsNil(o.WlServAreaRes) {
		toSerialize["wlServAreaRes"] = o.WlServAreaRes
	}
	if !IsNil(o.Rfsp) {
		toSerialize["rfsp"] = o.Rfsp
	}
	if o.SmfSelInfo.IsSet() {
		toSerialize["smfSelInfo"] = o.SmfSelInfo.Get()
	}
	if !IsNil(o.UeAmbr) {
		toSerialize["ueAmbr"] = o.UeAmbr
	}
	if !IsNil(o.UeSliceMbrs) {
		toSerialize["ueSliceMbrs"] = o.UeSliceMbrs
	}
	if !IsNil(o.PraStatuses) {
		toSerialize["praStatuses"] = o.PraStatuses
	}
	if !IsNil(o.UserLoc) {
		toSerialize["userLoc"] = o.UserLoc
	}
	if !IsNil(o.AllowedSnssais) {
		toSerialize["allowedSnssais"] = o.AllowedSnssais
	}
	if !IsNil(o.TargetSnssais) {
		toSerialize["targetSnssais"] = o.TargetSnssais
	}
	if !IsNil(o.MappingSnssais) {
		toSerialize["mappingSnssais"] = o.MappingSnssais
	}
	if !IsNil(o.AccessTypes) {
		toSerialize["accessTypes"] = o.AccessTypes
	}
	if !IsNil(o.RatTypes) {
		toSerialize["ratTypes"] = o.RatTypes
	}
	if !IsNil(o.N3gAllowedSnssais) {
		toSerialize["n3gAllowedSnssais"] = o.N3gAllowedSnssais
	}
	if o.TraceReq.IsSet() {
		toSerialize["traceReq"] = o.TraceReq.Get()
	}
	if !IsNil(o.Guami) {
		toSerialize["guami"] = o.Guami
	}
	if o.NwdafDatas != nil {
		toSerialize["nwdafDatas"] = o.NwdafDatas
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullablePolicyAssociationUpdateRequest struct {
	value *PolicyAssociationUpdateRequest
	isSet bool
}

func (v NullablePolicyAssociationUpdateRequest) Get() *PolicyAssociationUpdateRequest {
	return v.value
}

func (v *NullablePolicyAssociationUpdateRequest) Set(val *PolicyAssociationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAssociationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAssociationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAssociationUpdateRequest(val *PolicyAssociationUpdateRequest) *NullablePolicyAssociationUpdateRequest {
	return &NullablePolicyAssociationUpdateRequest{value: val, isSet: true}
}

func (v NullablePolicyAssociationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAssociationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
