/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
)

// checks if the PolicyAssociationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyAssociationRequest{}

// PolicyAssociationRequest Represents information that the NF service consumer provides when requesting the creation of a policy association.
type PolicyAssociationRequest struct {
	// String providing an URI formatted according to RFC 3986.
	NotificationUri string `json:"notificationUri"`
	// Alternate or backup IPv4 Address(es) where to send Notifications.
	AltNotifIpv4Addrs []string `json:"altNotifIpv4Addrs,omitempty"`
	// Alternate or backup IPv6 Address(es) where to send Notifications.
	AltNotifIpv6Addrs []Ipv6Addr `json:"altNotifIpv6Addrs,omitempty"`
	// Alternate or backup FQDN(s) where to send Notifications.
	AltNotifFqdns []string `json:"altNotifFqdns,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi       *string     `json:"gpsi,omitempty"`
	AccessType *AccessType `json:"accessType,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.
	Pei     *string       `json:"pei,omitempty"`
	UserLoc *UserLocation `json:"userLoc,omitempty"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time.
	TimeZone    *string    `json:"timeZone,omitempty"`
	ServingPlmn *PlmnIdNid `json:"servingPlmn,omitempty"`
	RatType     *RatType   `json:"ratType,omitempty"`
	GroupIds    []string   `json:"groupIds,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	HPcfId *string `json:"hPcfId,omitempty"`
	// string with format 'bytes' as defined in OpenAPI
	UePolReq    *string      `json:"uePolReq,omitempty"`
	Guami       *Guami       `json:"guami,omitempty"`
	ServiceName *ServiceName `json:"serviceName,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	ServingNfId *string           `json:"servingNfId,omitempty"`
	Pc5Capab    *Pc5Capability    `json:"pc5Capab,omitempty"`
	ProSeCapab  []ProSeCapability `json:"proSeCapab,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat"`
}

// NewPolicyAssociationRequest instantiates a new PolicyAssociationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyAssociationRequest(notificationUri string, supi string, suppFeat string) *PolicyAssociationRequest {
	this := PolicyAssociationRequest{}
	this.NotificationUri = notificationUri
	this.Supi = supi
	this.SuppFeat = suppFeat
	return &this
}

// NewPolicyAssociationRequestWithDefaults instantiates a new PolicyAssociationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyAssociationRequestWithDefaults() *PolicyAssociationRequest {
	this := PolicyAssociationRequest{}
	return &this
}

// GetNotificationUri returns the NotificationUri field value
func (o *PolicyAssociationRequest) GetNotificationUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationUri
}

// GetNotificationUriOk returns a tuple with the NotificationUri field value
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetNotificationUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationUri, true
}

// SetNotificationUri sets field value
func (o *PolicyAssociationRequest) SetNotificationUri(v string) {
	o.NotificationUri = v
}

// GetAltNotifIpv4Addrs returns the AltNotifIpv4Addrs field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetAltNotifIpv4Addrs() []string {
	if o == nil || IsNil(o.AltNotifIpv4Addrs) {
		var ret []string
		return ret
	}
	return o.AltNotifIpv4Addrs
}

// GetAltNotifIpv4AddrsOk returns a tuple with the AltNotifIpv4Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetAltNotifIpv4AddrsOk() ([]string, bool) {
	if o == nil || IsNil(o.AltNotifIpv4Addrs) {
		return nil, false
	}
	return o.AltNotifIpv4Addrs, true
}

// HasAltNotifIpv4Addrs returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasAltNotifIpv4Addrs() bool {
	if o != nil && !IsNil(o.AltNotifIpv4Addrs) {
		return true
	}

	return false
}

// SetAltNotifIpv4Addrs gets a reference to the given []string and assigns it to the AltNotifIpv4Addrs field.
func (o *PolicyAssociationRequest) SetAltNotifIpv4Addrs(v []string) {
	o.AltNotifIpv4Addrs = v
}

// GetAltNotifIpv6Addrs returns the AltNotifIpv6Addrs field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetAltNotifIpv6Addrs() []Ipv6Addr {
	if o == nil || IsNil(o.AltNotifIpv6Addrs) {
		var ret []Ipv6Addr
		return ret
	}
	return o.AltNotifIpv6Addrs
}

// GetAltNotifIpv6AddrsOk returns a tuple with the AltNotifIpv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetAltNotifIpv6AddrsOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.AltNotifIpv6Addrs) {
		return nil, false
	}
	return o.AltNotifIpv6Addrs, true
}

// HasAltNotifIpv6Addrs returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasAltNotifIpv6Addrs() bool {
	if o != nil && !IsNil(o.AltNotifIpv6Addrs) {
		return true
	}

	return false
}

// SetAltNotifIpv6Addrs gets a reference to the given []Ipv6Addr and assigns it to the AltNotifIpv6Addrs field.
func (o *PolicyAssociationRequest) SetAltNotifIpv6Addrs(v []Ipv6Addr) {
	o.AltNotifIpv6Addrs = v
}

// GetAltNotifFqdns returns the AltNotifFqdns field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetAltNotifFqdns() []string {
	if o == nil || IsNil(o.AltNotifFqdns) {
		var ret []string
		return ret
	}
	return o.AltNotifFqdns
}

// GetAltNotifFqdnsOk returns a tuple with the AltNotifFqdns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetAltNotifFqdnsOk() ([]string, bool) {
	if o == nil || IsNil(o.AltNotifFqdns) {
		return nil, false
	}
	return o.AltNotifFqdns, true
}

// HasAltNotifFqdns returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasAltNotifFqdns() bool {
	if o != nil && !IsNil(o.AltNotifFqdns) {
		return true
	}

	return false
}

// SetAltNotifFqdns gets a reference to the given []string and assigns it to the AltNotifFqdns field.
func (o *PolicyAssociationRequest) SetAltNotifFqdns(v []string) {
	o.AltNotifFqdns = v
}

// GetSupi returns the Supi field value
func (o *PolicyAssociationRequest) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetSupiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *PolicyAssociationRequest) SetSupi(v string) {
	o.Supi = v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *PolicyAssociationRequest) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *PolicyAssociationRequest) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetPei() string {
	if o == nil || IsNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetPeiOk() (*string, bool) {
	if o == nil || IsNil(o.Pei) {
		return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasPei() bool {
	if o != nil && !IsNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *PolicyAssociationRequest) SetPei(v string) {
	o.Pei = &v
}

// GetUserLoc returns the UserLoc field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetUserLoc() UserLocation {
	if o == nil || IsNil(o.UserLoc) {
		var ret UserLocation
		return ret
	}
	return *o.UserLoc
}

// GetUserLocOk returns a tuple with the UserLoc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetUserLocOk() (*UserLocation, bool) {
	if o == nil || IsNil(o.UserLoc) {
		return nil, false
	}
	return o.UserLoc, true
}

// HasUserLoc returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasUserLoc() bool {
	if o != nil && !IsNil(o.UserLoc) {
		return true
	}

	return false
}

// SetUserLoc gets a reference to the given UserLocation and assigns it to the UserLoc field.
func (o *PolicyAssociationRequest) SetUserLoc(v UserLocation) {
	o.UserLoc = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *PolicyAssociationRequest) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetServingPlmn returns the ServingPlmn field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetServingPlmn() PlmnIdNid {
	if o == nil || IsNil(o.ServingPlmn) {
		var ret PlmnIdNid
		return ret
	}
	return *o.ServingPlmn
}

// GetServingPlmnOk returns a tuple with the ServingPlmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetServingPlmnOk() (*PlmnIdNid, bool) {
	if o == nil || IsNil(o.ServingPlmn) {
		return nil, false
	}
	return o.ServingPlmn, true
}

// HasServingPlmn returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasServingPlmn() bool {
	if o != nil && !IsNil(o.ServingPlmn) {
		return true
	}

	return false
}

// SetServingPlmn gets a reference to the given PlmnIdNid and assigns it to the ServingPlmn field.
func (o *PolicyAssociationRequest) SetServingPlmn(v PlmnIdNid) {
	o.ServingPlmn = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetRatType() RatType {
	if o == nil || IsNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetRatTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RatType) {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasRatType() bool {
	if o != nil && !IsNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *PolicyAssociationRequest) SetRatType(v RatType) {
	o.RatType = &v
}

// GetGroupIds returns the GroupIds field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetGroupIds() []string {
	if o == nil || IsNil(o.GroupIds) {
		var ret []string
		return ret
	}
	return o.GroupIds
}

// GetGroupIdsOk returns a tuple with the GroupIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetGroupIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.GroupIds) {
		return nil, false
	}
	return o.GroupIds, true
}

// HasGroupIds returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasGroupIds() bool {
	if o != nil && !IsNil(o.GroupIds) {
		return true
	}

	return false
}

// SetGroupIds gets a reference to the given []string and assigns it to the GroupIds field.
func (o *PolicyAssociationRequest) SetGroupIds(v []string) {
	o.GroupIds = v
}

// GetHPcfId returns the HPcfId field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetHPcfId() string {
	if o == nil || IsNil(o.HPcfId) {
		var ret string
		return ret
	}
	return *o.HPcfId
}

// GetHPcfIdOk returns a tuple with the HPcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetHPcfIdOk() (*string, bool) {
	if o == nil || IsNil(o.HPcfId) {
		return nil, false
	}
	return o.HPcfId, true
}

// HasHPcfId returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasHPcfId() bool {
	if o != nil && !IsNil(o.HPcfId) {
		return true
	}

	return false
}

// SetHPcfId gets a reference to the given string and assigns it to the HPcfId field.
func (o *PolicyAssociationRequest) SetHPcfId(v string) {
	o.HPcfId = &v
}

// GetUePolReq returns the UePolReq field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetUePolReq() string {
	if o == nil || IsNil(o.UePolReq) {
		var ret string
		return ret
	}
	return *o.UePolReq
}

// GetUePolReqOk returns a tuple with the UePolReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetUePolReqOk() (*string, bool) {
	if o == nil || IsNil(o.UePolReq) {
		return nil, false
	}
	return o.UePolReq, true
}

// HasUePolReq returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasUePolReq() bool {
	if o != nil && !IsNil(o.UePolReq) {
		return true
	}

	return false
}

// SetUePolReq gets a reference to the given string and assigns it to the UePolReq field.
func (o *PolicyAssociationRequest) SetUePolReq(v string) {
	o.UePolReq = &v
}

// GetGuami returns the Guami field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetGuami() Guami {
	if o == nil || IsNil(o.Guami) {
		var ret Guami
		return ret
	}
	return *o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetGuamiOk() (*Guami, bool) {
	if o == nil || IsNil(o.Guami) {
		return nil, false
	}
	return o.Guami, true
}

// HasGuami returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasGuami() bool {
	if o != nil && !IsNil(o.Guami) {
		return true
	}

	return false
}

// SetGuami gets a reference to the given Guami and assigns it to the Guami field.
func (o *PolicyAssociationRequest) SetGuami(v Guami) {
	o.Guami = &v
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetServiceName() ServiceName {
	if o == nil || IsNil(o.ServiceName) {
		var ret ServiceName
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetServiceNameOk() (*ServiceName, bool) {
	if o == nil || IsNil(o.ServiceName) {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasServiceName() bool {
	if o != nil && !IsNil(o.ServiceName) {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given ServiceName and assigns it to the ServiceName field.
func (o *PolicyAssociationRequest) SetServiceName(v ServiceName) {
	o.ServiceName = &v
}

// GetServingNfId returns the ServingNfId field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetServingNfId() string {
	if o == nil || IsNil(o.ServingNfId) {
		var ret string
		return ret
	}
	return *o.ServingNfId
}

// GetServingNfIdOk returns a tuple with the ServingNfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetServingNfIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServingNfId) {
		return nil, false
	}
	return o.ServingNfId, true
}

// HasServingNfId returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasServingNfId() bool {
	if o != nil && !IsNil(o.ServingNfId) {
		return true
	}

	return false
}

// SetServingNfId gets a reference to the given string and assigns it to the ServingNfId field.
func (o *PolicyAssociationRequest) SetServingNfId(v string) {
	o.ServingNfId = &v
}

// GetPc5Capab returns the Pc5Capab field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetPc5Capab() Pc5Capability {
	if o == nil || IsNil(o.Pc5Capab) {
		var ret Pc5Capability
		return ret
	}
	return *o.Pc5Capab
}

// GetPc5CapabOk returns a tuple with the Pc5Capab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetPc5CapabOk() (*Pc5Capability, bool) {
	if o == nil || IsNil(o.Pc5Capab) {
		return nil, false
	}
	return o.Pc5Capab, true
}

// HasPc5Capab returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasPc5Capab() bool {
	if o != nil && !IsNil(o.Pc5Capab) {
		return true
	}

	return false
}

// SetPc5Capab gets a reference to the given Pc5Capability and assigns it to the Pc5Capab field.
func (o *PolicyAssociationRequest) SetPc5Capab(v Pc5Capability) {
	o.Pc5Capab = &v
}

// GetProSeCapab returns the ProSeCapab field value if set, zero value otherwise.
func (o *PolicyAssociationRequest) GetProSeCapab() []ProSeCapability {
	if o == nil || IsNil(o.ProSeCapab) {
		var ret []ProSeCapability
		return ret
	}
	return o.ProSeCapab
}

// GetProSeCapabOk returns a tuple with the ProSeCapab field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetProSeCapabOk() ([]ProSeCapability, bool) {
	if o == nil || IsNil(o.ProSeCapab) {
		return nil, false
	}
	return o.ProSeCapab, true
}

// HasProSeCapab returns a boolean if a field has been set.
func (o *PolicyAssociationRequest) HasProSeCapab() bool {
	if o != nil && !IsNil(o.ProSeCapab) {
		return true
	}

	return false
}

// SetProSeCapab gets a reference to the given []ProSeCapability and assigns it to the ProSeCapab field.
func (o *PolicyAssociationRequest) SetProSeCapab(v []ProSeCapability) {
	o.ProSeCapab = v
}

// GetSuppFeat returns the SuppFeat field value
func (o *PolicyAssociationRequest) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *PolicyAssociationRequest) GetSuppFeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *PolicyAssociationRequest) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o PolicyAssociationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyAssociationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["notificationUri"] = o.NotificationUri
	if !IsNil(o.AltNotifIpv4Addrs) {
		toSerialize["altNotifIpv4Addrs"] = o.AltNotifIpv4Addrs
	}
	if !IsNil(o.AltNotifIpv6Addrs) {
		toSerialize["altNotifIpv6Addrs"] = o.AltNotifIpv6Addrs
	}
	if !IsNil(o.AltNotifFqdns) {
		toSerialize["altNotifFqdns"] = o.AltNotifFqdns
	}
	toSerialize["supi"] = o.Supi
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if !IsNil(o.UserLoc) {
		toSerialize["userLoc"] = o.UserLoc
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.ServingPlmn) {
		toSerialize["servingPlmn"] = o.ServingPlmn
	}
	if !IsNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	if !IsNil(o.GroupIds) {
		toSerialize["groupIds"] = o.GroupIds
	}
	if !IsNil(o.HPcfId) {
		toSerialize["hPcfId"] = o.HPcfId
	}
	if !IsNil(o.UePolReq) {
		toSerialize["uePolReq"] = o.UePolReq
	}
	if !IsNil(o.Guami) {
		toSerialize["guami"] = o.Guami
	}
	if !IsNil(o.ServiceName) {
		toSerialize["serviceName"] = o.ServiceName
	}
	if !IsNil(o.ServingNfId) {
		toSerialize["servingNfId"] = o.ServingNfId
	}
	if !IsNil(o.Pc5Capab) {
		toSerialize["pc5Capab"] = o.Pc5Capab
	}
	if !IsNil(o.ProSeCapab) {
		toSerialize["proSeCapab"] = o.ProSeCapab
	}
	toSerialize["suppFeat"] = o.SuppFeat
	return toSerialize, nil
}

type NullablePolicyAssociationRequest struct {
	value *PolicyAssociationRequest
	isSet bool
}

func (v NullablePolicyAssociationRequest) Get() *PolicyAssociationRequest {
	return v.value
}

func (v *NullablePolicyAssociationRequest) Set(val *PolicyAssociationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyAssociationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyAssociationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyAssociationRequest(val *PolicyAssociationRequest) *NullablePolicyAssociationRequest {
	return &NullablePolicyAssociationRequest{value: val, isSet: true}
}

func (v NullablePolicyAssociationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyAssociationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
