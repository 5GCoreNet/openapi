/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the RegistrationContextContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationContextContainer{}

// RegistrationContextContainer Registration Context Container used to send the UE context information, N1 message from UE, AN address etc during Registration with AMF re-allocation procedure
type RegistrationContextContainer struct {
	UeContext UeContext `json:"ueContext"`
	// String with format \"time-numoffset\" optionally appended by \"daylightSavingTime\", where  - \"time-numoffset\" shall represent the time zone adjusted for daylight saving time and be    encoded as time-numoffset as defined in clause 5.6 of IETF RFC 3339;  - \"daylightSavingTime\" shall represent the adjustment that has been made and shall be    encoded as \"+1\" or \"+2\" for a +1 or +2 hours adjustment.   The example is for 8 hours behind UTC, +1 hour adjustment for Daylight Saving Time. 
	LocalTimeZone *string `json:"localTimeZone,omitempty"`
	AnType AccessType `json:"anType"`
	AnN2ApId int32 `json:"anN2ApId"`
	RanNodeId GlobalRanNodeId `json:"ranNodeId"`
	// Fully Qualified Domain Name
	InitialAmfName string `json:"initialAmfName"`
	UserLocation UserLocation `json:"userLocation"`
	RrcEstCause *string `json:"rrcEstCause,omitempty"`
	UeContextRequest *bool `json:"ueContextRequest,omitempty"`
	InitialAmfN2ApId *int32 `json:"initialAmfN2ApId,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	AnN2IPv4Addr *string `json:"anN2IPv4Addr,omitempty"`
	AnN2IPv6Addr *Ipv6Addr `json:"anN2IPv6Addr,omitempty"`
	AllowedNssai *AllowedNssai `json:"allowedNssai,omitempty"`
	ConfiguredNssai []ConfiguredSnssai `json:"configuredNssai,omitempty"`
	RejectedNssaiInPlmn []Snssai `json:"rejectedNssaiInPlmn,omitempty"`
	RejectedNssaiInTa []Snssai `json:"rejectedNssaiInTa,omitempty"`
	SelectedPlmnId *PlmnId `json:"selectedPlmnId,omitempty"`
	IabNodeInd *bool `json:"iabNodeInd,omitempty"`
	CeModeBInd *CeModeBInd `json:"ceModeBInd,omitempty"`
	LteMInd *LteMInd `json:"lteMInd,omitempty"`
	AuthenticatedInd *bool `json:"authenticatedInd,omitempty"`
	NpnAccessInfo *NpnAccessInfo `json:"npnAccessInfo,omitempty"`
}

// NewRegistrationContextContainer instantiates a new RegistrationContextContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationContextContainer(ueContext UeContext, anType AccessType, anN2ApId int32, ranNodeId GlobalRanNodeId, initialAmfName string, userLocation UserLocation) *RegistrationContextContainer {
	this := RegistrationContextContainer{}
	this.UeContext = ueContext
	this.AnType = anType
	this.AnN2ApId = anN2ApId
	this.RanNodeId = ranNodeId
	this.InitialAmfName = initialAmfName
	this.UserLocation = userLocation
	var ueContextRequest bool = false
	this.UeContextRequest = &ueContextRequest
	var iabNodeInd bool = false
	this.IabNodeInd = &iabNodeInd
	var authenticatedInd bool = false
	this.AuthenticatedInd = &authenticatedInd
	return &this
}

// NewRegistrationContextContainerWithDefaults instantiates a new RegistrationContextContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationContextContainerWithDefaults() *RegistrationContextContainer {
	this := RegistrationContextContainer{}
	var ueContextRequest bool = false
	this.UeContextRequest = &ueContextRequest
	var iabNodeInd bool = false
	this.IabNodeInd = &iabNodeInd
	var authenticatedInd bool = false
	this.AuthenticatedInd = &authenticatedInd
	return &this
}

// GetUeContext returns the UeContext field value
func (o *RegistrationContextContainer) GetUeContext() UeContext {
	if o == nil {
		var ret UeContext
		return ret
	}

	return o.UeContext
}

// GetUeContextOk returns a tuple with the UeContext field value
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetUeContextOk() (*UeContext, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UeContext, true
}

// SetUeContext sets field value
func (o *RegistrationContextContainer) SetUeContext(v UeContext) {
	o.UeContext = v
}

// GetLocalTimeZone returns the LocalTimeZone field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetLocalTimeZone() string {
	if o == nil || isNil(o.LocalTimeZone) {
		var ret string
		return ret
	}
	return *o.LocalTimeZone
}

// GetLocalTimeZoneOk returns a tuple with the LocalTimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetLocalTimeZoneOk() (*string, bool) {
	if o == nil || isNil(o.LocalTimeZone) {
		return nil, false
	}
	return o.LocalTimeZone, true
}

// HasLocalTimeZone returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasLocalTimeZone() bool {
	if o != nil && !isNil(o.LocalTimeZone) {
		return true
	}

	return false
}

// SetLocalTimeZone gets a reference to the given string and assigns it to the LocalTimeZone field.
func (o *RegistrationContextContainer) SetLocalTimeZone(v string) {
	o.LocalTimeZone = &v
}

// GetAnType returns the AnType field value
func (o *RegistrationContextContainer) GetAnType() AccessType {
	if o == nil {
		var ret AccessType
		return ret
	}

	return o.AnType
}

// GetAnTypeOk returns a tuple with the AnType field value
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetAnTypeOk() (*AccessType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnType, true
}

// SetAnType sets field value
func (o *RegistrationContextContainer) SetAnType(v AccessType) {
	o.AnType = v
}

// GetAnN2ApId returns the AnN2ApId field value
func (o *RegistrationContextContainer) GetAnN2ApId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AnN2ApId
}

// GetAnN2ApIdOk returns a tuple with the AnN2ApId field value
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetAnN2ApIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AnN2ApId, true
}

// SetAnN2ApId sets field value
func (o *RegistrationContextContainer) SetAnN2ApId(v int32) {
	o.AnN2ApId = v
}

// GetRanNodeId returns the RanNodeId field value
func (o *RegistrationContextContainer) GetRanNodeId() GlobalRanNodeId {
	if o == nil {
		var ret GlobalRanNodeId
		return ret
	}

	return o.RanNodeId
}

// GetRanNodeIdOk returns a tuple with the RanNodeId field value
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetRanNodeIdOk() (*GlobalRanNodeId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RanNodeId, true
}

// SetRanNodeId sets field value
func (o *RegistrationContextContainer) SetRanNodeId(v GlobalRanNodeId) {
	o.RanNodeId = v
}

// GetInitialAmfName returns the InitialAmfName field value
func (o *RegistrationContextContainer) GetInitialAmfName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialAmfName
}

// GetInitialAmfNameOk returns a tuple with the InitialAmfName field value
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetInitialAmfNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialAmfName, true
}

// SetInitialAmfName sets field value
func (o *RegistrationContextContainer) SetInitialAmfName(v string) {
	o.InitialAmfName = v
}

// GetUserLocation returns the UserLocation field value
func (o *RegistrationContextContainer) GetUserLocation() UserLocation {
	if o == nil {
		var ret UserLocation
		return ret
	}

	return o.UserLocation
}

// GetUserLocationOk returns a tuple with the UserLocation field value
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetUserLocationOk() (*UserLocation, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserLocation, true
}

// SetUserLocation sets field value
func (o *RegistrationContextContainer) SetUserLocation(v UserLocation) {
	o.UserLocation = v
}

// GetRrcEstCause returns the RrcEstCause field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetRrcEstCause() string {
	if o == nil || isNil(o.RrcEstCause) {
		var ret string
		return ret
	}
	return *o.RrcEstCause
}

// GetRrcEstCauseOk returns a tuple with the RrcEstCause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetRrcEstCauseOk() (*string, bool) {
	if o == nil || isNil(o.RrcEstCause) {
		return nil, false
	}
	return o.RrcEstCause, true
}

// HasRrcEstCause returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasRrcEstCause() bool {
	if o != nil && !isNil(o.RrcEstCause) {
		return true
	}

	return false
}

// SetRrcEstCause gets a reference to the given string and assigns it to the RrcEstCause field.
func (o *RegistrationContextContainer) SetRrcEstCause(v string) {
	o.RrcEstCause = &v
}

// GetUeContextRequest returns the UeContextRequest field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetUeContextRequest() bool {
	if o == nil || isNil(o.UeContextRequest) {
		var ret bool
		return ret
	}
	return *o.UeContextRequest
}

// GetUeContextRequestOk returns a tuple with the UeContextRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetUeContextRequestOk() (*bool, bool) {
	if o == nil || isNil(o.UeContextRequest) {
		return nil, false
	}
	return o.UeContextRequest, true
}

// HasUeContextRequest returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasUeContextRequest() bool {
	if o != nil && !isNil(o.UeContextRequest) {
		return true
	}

	return false
}

// SetUeContextRequest gets a reference to the given bool and assigns it to the UeContextRequest field.
func (o *RegistrationContextContainer) SetUeContextRequest(v bool) {
	o.UeContextRequest = &v
}

// GetInitialAmfN2ApId returns the InitialAmfN2ApId field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetInitialAmfN2ApId() int32 {
	if o == nil || isNil(o.InitialAmfN2ApId) {
		var ret int32
		return ret
	}
	return *o.InitialAmfN2ApId
}

// GetInitialAmfN2ApIdOk returns a tuple with the InitialAmfN2ApId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetInitialAmfN2ApIdOk() (*int32, bool) {
	if o == nil || isNil(o.InitialAmfN2ApId) {
		return nil, false
	}
	return o.InitialAmfN2ApId, true
}

// HasInitialAmfN2ApId returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasInitialAmfN2ApId() bool {
	if o != nil && !isNil(o.InitialAmfN2ApId) {
		return true
	}

	return false
}

// SetInitialAmfN2ApId gets a reference to the given int32 and assigns it to the InitialAmfN2ApId field.
func (o *RegistrationContextContainer) SetInitialAmfN2ApId(v int32) {
	o.InitialAmfN2ApId = &v
}

// GetAnN2IPv4Addr returns the AnN2IPv4Addr field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetAnN2IPv4Addr() string {
	if o == nil || isNil(o.AnN2IPv4Addr) {
		var ret string
		return ret
	}
	return *o.AnN2IPv4Addr
}

// GetAnN2IPv4AddrOk returns a tuple with the AnN2IPv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetAnN2IPv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.AnN2IPv4Addr) {
		return nil, false
	}
	return o.AnN2IPv4Addr, true
}

// HasAnN2IPv4Addr returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasAnN2IPv4Addr() bool {
	if o != nil && !isNil(o.AnN2IPv4Addr) {
		return true
	}

	return false
}

// SetAnN2IPv4Addr gets a reference to the given string and assigns it to the AnN2IPv4Addr field.
func (o *RegistrationContextContainer) SetAnN2IPv4Addr(v string) {
	o.AnN2IPv4Addr = &v
}

// GetAnN2IPv6Addr returns the AnN2IPv6Addr field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetAnN2IPv6Addr() Ipv6Addr {
	if o == nil || isNil(o.AnN2IPv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.AnN2IPv6Addr
}

// GetAnN2IPv6AddrOk returns a tuple with the AnN2IPv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetAnN2IPv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.AnN2IPv6Addr) {
		return nil, false
	}
	return o.AnN2IPv6Addr, true
}

// HasAnN2IPv6Addr returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasAnN2IPv6Addr() bool {
	if o != nil && !isNil(o.AnN2IPv6Addr) {
		return true
	}

	return false
}

// SetAnN2IPv6Addr gets a reference to the given Ipv6Addr and assigns it to the AnN2IPv6Addr field.
func (o *RegistrationContextContainer) SetAnN2IPv6Addr(v Ipv6Addr) {
	o.AnN2IPv6Addr = &v
}

// GetAllowedNssai returns the AllowedNssai field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetAllowedNssai() AllowedNssai {
	if o == nil || isNil(o.AllowedNssai) {
		var ret AllowedNssai
		return ret
	}
	return *o.AllowedNssai
}

// GetAllowedNssaiOk returns a tuple with the AllowedNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetAllowedNssaiOk() (*AllowedNssai, bool) {
	if o == nil || isNil(o.AllowedNssai) {
		return nil, false
	}
	return o.AllowedNssai, true
}

// HasAllowedNssai returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasAllowedNssai() bool {
	if o != nil && !isNil(o.AllowedNssai) {
		return true
	}

	return false
}

// SetAllowedNssai gets a reference to the given AllowedNssai and assigns it to the AllowedNssai field.
func (o *RegistrationContextContainer) SetAllowedNssai(v AllowedNssai) {
	o.AllowedNssai = &v
}

// GetConfiguredNssai returns the ConfiguredNssai field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetConfiguredNssai() []ConfiguredSnssai {
	if o == nil || isNil(o.ConfiguredNssai) {
		var ret []ConfiguredSnssai
		return ret
	}
	return o.ConfiguredNssai
}

// GetConfiguredNssaiOk returns a tuple with the ConfiguredNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetConfiguredNssaiOk() ([]ConfiguredSnssai, bool) {
	if o == nil || isNil(o.ConfiguredNssai) {
		return nil, false
	}
	return o.ConfiguredNssai, true
}

// HasConfiguredNssai returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasConfiguredNssai() bool {
	if o != nil && !isNil(o.ConfiguredNssai) {
		return true
	}

	return false
}

// SetConfiguredNssai gets a reference to the given []ConfiguredSnssai and assigns it to the ConfiguredNssai field.
func (o *RegistrationContextContainer) SetConfiguredNssai(v []ConfiguredSnssai) {
	o.ConfiguredNssai = v
}

// GetRejectedNssaiInPlmn returns the RejectedNssaiInPlmn field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetRejectedNssaiInPlmn() []Snssai {
	if o == nil || isNil(o.RejectedNssaiInPlmn) {
		var ret []Snssai
		return ret
	}
	return o.RejectedNssaiInPlmn
}

// GetRejectedNssaiInPlmnOk returns a tuple with the RejectedNssaiInPlmn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetRejectedNssaiInPlmnOk() ([]Snssai, bool) {
	if o == nil || isNil(o.RejectedNssaiInPlmn) {
		return nil, false
	}
	return o.RejectedNssaiInPlmn, true
}

// HasRejectedNssaiInPlmn returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasRejectedNssaiInPlmn() bool {
	if o != nil && !isNil(o.RejectedNssaiInPlmn) {
		return true
	}

	return false
}

// SetRejectedNssaiInPlmn gets a reference to the given []Snssai and assigns it to the RejectedNssaiInPlmn field.
func (o *RegistrationContextContainer) SetRejectedNssaiInPlmn(v []Snssai) {
	o.RejectedNssaiInPlmn = v
}

// GetRejectedNssaiInTa returns the RejectedNssaiInTa field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetRejectedNssaiInTa() []Snssai {
	if o == nil || isNil(o.RejectedNssaiInTa) {
		var ret []Snssai
		return ret
	}
	return o.RejectedNssaiInTa
}

// GetRejectedNssaiInTaOk returns a tuple with the RejectedNssaiInTa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetRejectedNssaiInTaOk() ([]Snssai, bool) {
	if o == nil || isNil(o.RejectedNssaiInTa) {
		return nil, false
	}
	return o.RejectedNssaiInTa, true
}

// HasRejectedNssaiInTa returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasRejectedNssaiInTa() bool {
	if o != nil && !isNil(o.RejectedNssaiInTa) {
		return true
	}

	return false
}

// SetRejectedNssaiInTa gets a reference to the given []Snssai and assigns it to the RejectedNssaiInTa field.
func (o *RegistrationContextContainer) SetRejectedNssaiInTa(v []Snssai) {
	o.RejectedNssaiInTa = v
}

// GetSelectedPlmnId returns the SelectedPlmnId field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetSelectedPlmnId() PlmnId {
	if o == nil || isNil(o.SelectedPlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.SelectedPlmnId
}

// GetSelectedPlmnIdOk returns a tuple with the SelectedPlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetSelectedPlmnIdOk() (*PlmnId, bool) {
	if o == nil || isNil(o.SelectedPlmnId) {
		return nil, false
	}
	return o.SelectedPlmnId, true
}

// HasSelectedPlmnId returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasSelectedPlmnId() bool {
	if o != nil && !isNil(o.SelectedPlmnId) {
		return true
	}

	return false
}

// SetSelectedPlmnId gets a reference to the given PlmnId and assigns it to the SelectedPlmnId field.
func (o *RegistrationContextContainer) SetSelectedPlmnId(v PlmnId) {
	o.SelectedPlmnId = &v
}

// GetIabNodeInd returns the IabNodeInd field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetIabNodeInd() bool {
	if o == nil || isNil(o.IabNodeInd) {
		var ret bool
		return ret
	}
	return *o.IabNodeInd
}

// GetIabNodeIndOk returns a tuple with the IabNodeInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetIabNodeIndOk() (*bool, bool) {
	if o == nil || isNil(o.IabNodeInd) {
		return nil, false
	}
	return o.IabNodeInd, true
}

// HasIabNodeInd returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasIabNodeInd() bool {
	if o != nil && !isNil(o.IabNodeInd) {
		return true
	}

	return false
}

// SetIabNodeInd gets a reference to the given bool and assigns it to the IabNodeInd field.
func (o *RegistrationContextContainer) SetIabNodeInd(v bool) {
	o.IabNodeInd = &v
}

// GetCeModeBInd returns the CeModeBInd field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetCeModeBInd() CeModeBInd {
	if o == nil || isNil(o.CeModeBInd) {
		var ret CeModeBInd
		return ret
	}
	return *o.CeModeBInd
}

// GetCeModeBIndOk returns a tuple with the CeModeBInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetCeModeBIndOk() (*CeModeBInd, bool) {
	if o == nil || isNil(o.CeModeBInd) {
		return nil, false
	}
	return o.CeModeBInd, true
}

// HasCeModeBInd returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasCeModeBInd() bool {
	if o != nil && !isNil(o.CeModeBInd) {
		return true
	}

	return false
}

// SetCeModeBInd gets a reference to the given CeModeBInd and assigns it to the CeModeBInd field.
func (o *RegistrationContextContainer) SetCeModeBInd(v CeModeBInd) {
	o.CeModeBInd = &v
}

// GetLteMInd returns the LteMInd field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetLteMInd() LteMInd {
	if o == nil || isNil(o.LteMInd) {
		var ret LteMInd
		return ret
	}
	return *o.LteMInd
}

// GetLteMIndOk returns a tuple with the LteMInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetLteMIndOk() (*LteMInd, bool) {
	if o == nil || isNil(o.LteMInd) {
		return nil, false
	}
	return o.LteMInd, true
}

// HasLteMInd returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasLteMInd() bool {
	if o != nil && !isNil(o.LteMInd) {
		return true
	}

	return false
}

// SetLteMInd gets a reference to the given LteMInd and assigns it to the LteMInd field.
func (o *RegistrationContextContainer) SetLteMInd(v LteMInd) {
	o.LteMInd = &v
}

// GetAuthenticatedInd returns the AuthenticatedInd field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetAuthenticatedInd() bool {
	if o == nil || isNil(o.AuthenticatedInd) {
		var ret bool
		return ret
	}
	return *o.AuthenticatedInd
}

// GetAuthenticatedIndOk returns a tuple with the AuthenticatedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetAuthenticatedIndOk() (*bool, bool) {
	if o == nil || isNil(o.AuthenticatedInd) {
		return nil, false
	}
	return o.AuthenticatedInd, true
}

// HasAuthenticatedInd returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasAuthenticatedInd() bool {
	if o != nil && !isNil(o.AuthenticatedInd) {
		return true
	}

	return false
}

// SetAuthenticatedInd gets a reference to the given bool and assigns it to the AuthenticatedInd field.
func (o *RegistrationContextContainer) SetAuthenticatedInd(v bool) {
	o.AuthenticatedInd = &v
}

// GetNpnAccessInfo returns the NpnAccessInfo field value if set, zero value otherwise.
func (o *RegistrationContextContainer) GetNpnAccessInfo() NpnAccessInfo {
	if o == nil || isNil(o.NpnAccessInfo) {
		var ret NpnAccessInfo
		return ret
	}
	return *o.NpnAccessInfo
}

// GetNpnAccessInfoOk returns a tuple with the NpnAccessInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationContextContainer) GetNpnAccessInfoOk() (*NpnAccessInfo, bool) {
	if o == nil || isNil(o.NpnAccessInfo) {
		return nil, false
	}
	return o.NpnAccessInfo, true
}

// HasNpnAccessInfo returns a boolean if a field has been set.
func (o *RegistrationContextContainer) HasNpnAccessInfo() bool {
	if o != nil && !isNil(o.NpnAccessInfo) {
		return true
	}

	return false
}

// SetNpnAccessInfo gets a reference to the given NpnAccessInfo and assigns it to the NpnAccessInfo field.
func (o *RegistrationContextContainer) SetNpnAccessInfo(v NpnAccessInfo) {
	o.NpnAccessInfo = &v
}

func (o RegistrationContextContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationContextContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ueContext"] = o.UeContext
	if !isNil(o.LocalTimeZone) {
		toSerialize["localTimeZone"] = o.LocalTimeZone
	}
	toSerialize["anType"] = o.AnType
	toSerialize["anN2ApId"] = o.AnN2ApId
	toSerialize["ranNodeId"] = o.RanNodeId
	toSerialize["initialAmfName"] = o.InitialAmfName
	toSerialize["userLocation"] = o.UserLocation
	if !isNil(o.RrcEstCause) {
		toSerialize["rrcEstCause"] = o.RrcEstCause
	}
	if !isNil(o.UeContextRequest) {
		toSerialize["ueContextRequest"] = o.UeContextRequest
	}
	if !isNil(o.InitialAmfN2ApId) {
		toSerialize["initialAmfN2ApId"] = o.InitialAmfN2ApId
	}
	if !isNil(o.AnN2IPv4Addr) {
		toSerialize["anN2IPv4Addr"] = o.AnN2IPv4Addr
	}
	if !isNil(o.AnN2IPv6Addr) {
		toSerialize["anN2IPv6Addr"] = o.AnN2IPv6Addr
	}
	if !isNil(o.AllowedNssai) {
		toSerialize["allowedNssai"] = o.AllowedNssai
	}
	if !isNil(o.ConfiguredNssai) {
		toSerialize["configuredNssai"] = o.ConfiguredNssai
	}
	if !isNil(o.RejectedNssaiInPlmn) {
		toSerialize["rejectedNssaiInPlmn"] = o.RejectedNssaiInPlmn
	}
	if !isNil(o.RejectedNssaiInTa) {
		toSerialize["rejectedNssaiInTa"] = o.RejectedNssaiInTa
	}
	if !isNil(o.SelectedPlmnId) {
		toSerialize["selectedPlmnId"] = o.SelectedPlmnId
	}
	if !isNil(o.IabNodeInd) {
		toSerialize["iabNodeInd"] = o.IabNodeInd
	}
	if !isNil(o.CeModeBInd) {
		toSerialize["ceModeBInd"] = o.CeModeBInd
	}
	if !isNil(o.LteMInd) {
		toSerialize["lteMInd"] = o.LteMInd
	}
	if !isNil(o.AuthenticatedInd) {
		toSerialize["authenticatedInd"] = o.AuthenticatedInd
	}
	if !isNil(o.NpnAccessInfo) {
		toSerialize["npnAccessInfo"] = o.NpnAccessInfo
	}
	return toSerialize, nil
}

type NullableRegistrationContextContainer struct {
	value *RegistrationContextContainer
	isSet bool
}

func (v NullableRegistrationContextContainer) Get() *RegistrationContextContainer {
	return v.value
}

func (v *NullableRegistrationContextContainer) Set(val *RegistrationContextContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationContextContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationContextContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationContextContainer(val *RegistrationContextContainer) *NullableRegistrationContextContainer {
	return &NullableRegistrationContextContainer{value: val, isSet: true}
}

func (v NullableRegistrationContextContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationContextContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


