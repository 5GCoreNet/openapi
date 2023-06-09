/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
	"time"
)

// checks if the PcfBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PcfBinding{}

// PcfBinding Identifies an Individual PCF for a PDU Session binding.
type PcfBinding struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string `json:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi *string `json:"gpsi,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	Ipv4Addr   *string     `json:"ipv4Addr,omitempty"`
	Ipv6Prefix *Ipv6Prefix `json:"ipv6Prefix,omitempty"`
	// The additional IPv6 Address Prefixes of the served UE.
	AddIpv6Prefixes []Ipv6Prefix `json:"addIpv6Prefixes,omitempty"`
	IpDomain        *string      `json:"ipDomain,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042.
	MacAddr48 *string `json:"macAddr48,omitempty"`
	// The additional MAC Addresses of the served UE.
	AddMacAddrs []string `json:"addMacAddrs,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	Dnn string `json:"dnn"`
	// Fully Qualified Domain Name
	PcfFqdn *string `json:"pcfFqdn,omitempty"`
	// IP end points of the PCF hosting the Npcf_PolicyAuthorization service
	PcfIpEndPoints []IpEndPoint `json:"pcfIpEndPoints,omitempty"`
	// Fully Qualified Domain Name
	PcfDiamHost *string `json:"pcfDiamHost,omitempty"`
	// Fully Qualified Domain Name
	PcfDiamRealm *string `json:"pcfDiamRealm,omitempty"`
	// Fully Qualified Domain Name
	PcfSmFqdn *string `json:"pcfSmFqdn,omitempty"`
	// IP end points of the PCF hosting the Npcf_SMPolicyControl service.
	PcfSmIpEndPoints []IpEndPoint `json:"pcfSmIpEndPoints,omitempty"`
	Snssai           Snssai       `json:"snssai"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.
	PcfId *string `json:"pcfId,omitempty"`
	// NF Set Identifier (see clause 28.12 of 3GPP TS 23.003), formatted as the following string \"set<Set ID>.<nftype>set.5gc.mnc<MNC>.mcc<MCC>\", or  \"set<SetID>.<NFType>set.5gc.nid<NID>.mnc<MNC>.mcc<MCC>\" with  <MCC> encoded as defined in clause 5.4.2 (\"Mcc\" data type definition)  <MNC> encoding the Mobile Network Code part of the PLMN, comprising 3 digits.    If there are only 2 significant digits in the MNC, one \"0\" digit shall be inserted    at the left side to fill the 3 digits coding of MNC.  Pattern: '^[0-9]{3}$' <NFType> encoded as a value defined in Table 6.1.6.3.3-1 of 3GPP TS 29.510 but    with lower case characters <Set ID> encoded as a string of characters consisting of    alphabetic characters (A-Z and a-z), digits (0-9) and/or the hyphen (-) and that    shall end with either an alphabetic character or a digit.
	PcfSetId *string `json:"pcfSetId,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	RecoveryTime       *time.Time            `json:"recoveryTime,omitempty"`
	ParaCom            *ParameterCombination `json:"paraCom,omitempty"`
	BindLevel          *BindingLevel         `json:"bindLevel,omitempty"`
	Ipv4FrameRouteList []string              `json:"ipv4FrameRouteList,omitempty"`
	Ipv6FrameRouteList []Ipv6Prefix          `json:"ipv6FrameRouteList,omitempty"`
}

// NewPcfBinding instantiates a new PcfBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPcfBinding(dnn string, snssai Snssai) *PcfBinding {
	this := PcfBinding{}
	this.Dnn = dnn
	this.Snssai = snssai
	return &this
}

// NewPcfBindingWithDefaults instantiates a new PcfBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPcfBindingWithDefaults() *PcfBinding {
	this := PcfBinding{}
	return &this
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *PcfBinding) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *PcfBinding) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *PcfBinding) SetSupi(v string) {
	o.Supi = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *PcfBinding) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *PcfBinding) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *PcfBinding) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *PcfBinding) GetIpv4Addr() string {
	if o == nil || IsNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Addr) {
		return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *PcfBinding) HasIpv4Addr() bool {
	if o != nil && !IsNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *PcfBinding) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Prefix returns the Ipv6Prefix field value if set, zero value otherwise.
func (o *PcfBinding) GetIpv6Prefix() Ipv6Prefix {
	if o == nil || IsNil(o.Ipv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.Ipv6Prefix
}

// GetIpv6PrefixOk returns a tuple with the Ipv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.Ipv6Prefix) {
		return nil, false
	}
	return o.Ipv6Prefix, true
}

// HasIpv6Prefix returns a boolean if a field has been set.
func (o *PcfBinding) HasIpv6Prefix() bool {
	if o != nil && !IsNil(o.Ipv6Prefix) {
		return true
	}

	return false
}

// SetIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the Ipv6Prefix field.
func (o *PcfBinding) SetIpv6Prefix(v Ipv6Prefix) {
	o.Ipv6Prefix = &v
}

// GetAddIpv6Prefixes returns the AddIpv6Prefixes field value if set, zero value otherwise.
func (o *PcfBinding) GetAddIpv6Prefixes() []Ipv6Prefix {
	if o == nil || IsNil(o.AddIpv6Prefixes) {
		var ret []Ipv6Prefix
		return ret
	}
	return o.AddIpv6Prefixes
}

// GetAddIpv6PrefixesOk returns a tuple with the AddIpv6Prefixes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetAddIpv6PrefixesOk() ([]Ipv6Prefix, bool) {
	if o == nil || IsNil(o.AddIpv6Prefixes) {
		return nil, false
	}
	return o.AddIpv6Prefixes, true
}

// HasAddIpv6Prefixes returns a boolean if a field has been set.
func (o *PcfBinding) HasAddIpv6Prefixes() bool {
	if o != nil && !IsNil(o.AddIpv6Prefixes) {
		return true
	}

	return false
}

// SetAddIpv6Prefixes gets a reference to the given []Ipv6Prefix and assigns it to the AddIpv6Prefixes field.
func (o *PcfBinding) SetAddIpv6Prefixes(v []Ipv6Prefix) {
	o.AddIpv6Prefixes = v
}

// GetIpDomain returns the IpDomain field value if set, zero value otherwise.
func (o *PcfBinding) GetIpDomain() string {
	if o == nil || IsNil(o.IpDomain) {
		var ret string
		return ret
	}
	return *o.IpDomain
}

// GetIpDomainOk returns a tuple with the IpDomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetIpDomainOk() (*string, bool) {
	if o == nil || IsNil(o.IpDomain) {
		return nil, false
	}
	return o.IpDomain, true
}

// HasIpDomain returns a boolean if a field has been set.
func (o *PcfBinding) HasIpDomain() bool {
	if o != nil && !IsNil(o.IpDomain) {
		return true
	}

	return false
}

// SetIpDomain gets a reference to the given string and assigns it to the IpDomain field.
func (o *PcfBinding) SetIpDomain(v string) {
	o.IpDomain = &v
}

// GetMacAddr48 returns the MacAddr48 field value if set, zero value otherwise.
func (o *PcfBinding) GetMacAddr48() string {
	if o == nil || IsNil(o.MacAddr48) {
		var ret string
		return ret
	}
	return *o.MacAddr48
}

// GetMacAddr48Ok returns a tuple with the MacAddr48 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetMacAddr48Ok() (*string, bool) {
	if o == nil || IsNil(o.MacAddr48) {
		return nil, false
	}
	return o.MacAddr48, true
}

// HasMacAddr48 returns a boolean if a field has been set.
func (o *PcfBinding) HasMacAddr48() bool {
	if o != nil && !IsNil(o.MacAddr48) {
		return true
	}

	return false
}

// SetMacAddr48 gets a reference to the given string and assigns it to the MacAddr48 field.
func (o *PcfBinding) SetMacAddr48(v string) {
	o.MacAddr48 = &v
}

// GetAddMacAddrs returns the AddMacAddrs field value if set, zero value otherwise.
func (o *PcfBinding) GetAddMacAddrs() []string {
	if o == nil || IsNil(o.AddMacAddrs) {
		var ret []string
		return ret
	}
	return o.AddMacAddrs
}

// GetAddMacAddrsOk returns a tuple with the AddMacAddrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetAddMacAddrsOk() ([]string, bool) {
	if o == nil || IsNil(o.AddMacAddrs) {
		return nil, false
	}
	return o.AddMacAddrs, true
}

// HasAddMacAddrs returns a boolean if a field has been set.
func (o *PcfBinding) HasAddMacAddrs() bool {
	if o != nil && !IsNil(o.AddMacAddrs) {
		return true
	}

	return false
}

// SetAddMacAddrs gets a reference to the given []string and assigns it to the AddMacAddrs field.
func (o *PcfBinding) SetAddMacAddrs(v []string) {
	o.AddMacAddrs = v
}

// GetDnn returns the Dnn field value
func (o *PcfBinding) GetDnn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetDnnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dnn, true
}

// SetDnn sets field value
func (o *PcfBinding) SetDnn(v string) {
	o.Dnn = v
}

// GetPcfFqdn returns the PcfFqdn field value if set, zero value otherwise.
func (o *PcfBinding) GetPcfFqdn() string {
	if o == nil || IsNil(o.PcfFqdn) {
		var ret string
		return ret
	}
	return *o.PcfFqdn
}

// GetPcfFqdnOk returns a tuple with the PcfFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetPcfFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.PcfFqdn) {
		return nil, false
	}
	return o.PcfFqdn, true
}

// HasPcfFqdn returns a boolean if a field has been set.
func (o *PcfBinding) HasPcfFqdn() bool {
	if o != nil && !IsNil(o.PcfFqdn) {
		return true
	}

	return false
}

// SetPcfFqdn gets a reference to the given string and assigns it to the PcfFqdn field.
func (o *PcfBinding) SetPcfFqdn(v string) {
	o.PcfFqdn = &v
}

// GetPcfIpEndPoints returns the PcfIpEndPoints field value if set, zero value otherwise.
func (o *PcfBinding) GetPcfIpEndPoints() []IpEndPoint {
	if o == nil || IsNil(o.PcfIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.PcfIpEndPoints
}

// GetPcfIpEndPointsOk returns a tuple with the PcfIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetPcfIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || IsNil(o.PcfIpEndPoints) {
		return nil, false
	}
	return o.PcfIpEndPoints, true
}

// HasPcfIpEndPoints returns a boolean if a field has been set.
func (o *PcfBinding) HasPcfIpEndPoints() bool {
	if o != nil && !IsNil(o.PcfIpEndPoints) {
		return true
	}

	return false
}

// SetPcfIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the PcfIpEndPoints field.
func (o *PcfBinding) SetPcfIpEndPoints(v []IpEndPoint) {
	o.PcfIpEndPoints = v
}

// GetPcfDiamHost returns the PcfDiamHost field value if set, zero value otherwise.
func (o *PcfBinding) GetPcfDiamHost() string {
	if o == nil || IsNil(o.PcfDiamHost) {
		var ret string
		return ret
	}
	return *o.PcfDiamHost
}

// GetPcfDiamHostOk returns a tuple with the PcfDiamHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetPcfDiamHostOk() (*string, bool) {
	if o == nil || IsNil(o.PcfDiamHost) {
		return nil, false
	}
	return o.PcfDiamHost, true
}

// HasPcfDiamHost returns a boolean if a field has been set.
func (o *PcfBinding) HasPcfDiamHost() bool {
	if o != nil && !IsNil(o.PcfDiamHost) {
		return true
	}

	return false
}

// SetPcfDiamHost gets a reference to the given string and assigns it to the PcfDiamHost field.
func (o *PcfBinding) SetPcfDiamHost(v string) {
	o.PcfDiamHost = &v
}

// GetPcfDiamRealm returns the PcfDiamRealm field value if set, zero value otherwise.
func (o *PcfBinding) GetPcfDiamRealm() string {
	if o == nil || IsNil(o.PcfDiamRealm) {
		var ret string
		return ret
	}
	return *o.PcfDiamRealm
}

// GetPcfDiamRealmOk returns a tuple with the PcfDiamRealm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetPcfDiamRealmOk() (*string, bool) {
	if o == nil || IsNil(o.PcfDiamRealm) {
		return nil, false
	}
	return o.PcfDiamRealm, true
}

// HasPcfDiamRealm returns a boolean if a field has been set.
func (o *PcfBinding) HasPcfDiamRealm() bool {
	if o != nil && !IsNil(o.PcfDiamRealm) {
		return true
	}

	return false
}

// SetPcfDiamRealm gets a reference to the given string and assigns it to the PcfDiamRealm field.
func (o *PcfBinding) SetPcfDiamRealm(v string) {
	o.PcfDiamRealm = &v
}

// GetPcfSmFqdn returns the PcfSmFqdn field value if set, zero value otherwise.
func (o *PcfBinding) GetPcfSmFqdn() string {
	if o == nil || IsNil(o.PcfSmFqdn) {
		var ret string
		return ret
	}
	return *o.PcfSmFqdn
}

// GetPcfSmFqdnOk returns a tuple with the PcfSmFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetPcfSmFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.PcfSmFqdn) {
		return nil, false
	}
	return o.PcfSmFqdn, true
}

// HasPcfSmFqdn returns a boolean if a field has been set.
func (o *PcfBinding) HasPcfSmFqdn() bool {
	if o != nil && !IsNil(o.PcfSmFqdn) {
		return true
	}

	return false
}

// SetPcfSmFqdn gets a reference to the given string and assigns it to the PcfSmFqdn field.
func (o *PcfBinding) SetPcfSmFqdn(v string) {
	o.PcfSmFqdn = &v
}

// GetPcfSmIpEndPoints returns the PcfSmIpEndPoints field value if set, zero value otherwise.
func (o *PcfBinding) GetPcfSmIpEndPoints() []IpEndPoint {
	if o == nil || IsNil(o.PcfSmIpEndPoints) {
		var ret []IpEndPoint
		return ret
	}
	return o.PcfSmIpEndPoints
}

// GetPcfSmIpEndPointsOk returns a tuple with the PcfSmIpEndPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetPcfSmIpEndPointsOk() ([]IpEndPoint, bool) {
	if o == nil || IsNil(o.PcfSmIpEndPoints) {
		return nil, false
	}
	return o.PcfSmIpEndPoints, true
}

// HasPcfSmIpEndPoints returns a boolean if a field has been set.
func (o *PcfBinding) HasPcfSmIpEndPoints() bool {
	if o != nil && !IsNil(o.PcfSmIpEndPoints) {
		return true
	}

	return false
}

// SetPcfSmIpEndPoints gets a reference to the given []IpEndPoint and assigns it to the PcfSmIpEndPoints field.
func (o *PcfBinding) SetPcfSmIpEndPoints(v []IpEndPoint) {
	o.PcfSmIpEndPoints = v
}

// GetSnssai returns the Snssai field value
func (o *PcfBinding) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *PcfBinding) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *PcfBinding) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *PcfBinding) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *PcfBinding) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetPcfId returns the PcfId field value if set, zero value otherwise.
func (o *PcfBinding) GetPcfId() string {
	if o == nil || IsNil(o.PcfId) {
		var ret string
		return ret
	}
	return *o.PcfId
}

// GetPcfIdOk returns a tuple with the PcfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetPcfIdOk() (*string, bool) {
	if o == nil || IsNil(o.PcfId) {
		return nil, false
	}
	return o.PcfId, true
}

// HasPcfId returns a boolean if a field has been set.
func (o *PcfBinding) HasPcfId() bool {
	if o != nil && !IsNil(o.PcfId) {
		return true
	}

	return false
}

// SetPcfId gets a reference to the given string and assigns it to the PcfId field.
func (o *PcfBinding) SetPcfId(v string) {
	o.PcfId = &v
}

// GetPcfSetId returns the PcfSetId field value if set, zero value otherwise.
func (o *PcfBinding) GetPcfSetId() string {
	if o == nil || IsNil(o.PcfSetId) {
		var ret string
		return ret
	}
	return *o.PcfSetId
}

// GetPcfSetIdOk returns a tuple with the PcfSetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetPcfSetIdOk() (*string, bool) {
	if o == nil || IsNil(o.PcfSetId) {
		return nil, false
	}
	return o.PcfSetId, true
}

// HasPcfSetId returns a boolean if a field has been set.
func (o *PcfBinding) HasPcfSetId() bool {
	if o != nil && !IsNil(o.PcfSetId) {
		return true
	}

	return false
}

// SetPcfSetId gets a reference to the given string and assigns it to the PcfSetId field.
func (o *PcfBinding) SetPcfSetId(v string) {
	o.PcfSetId = &v
}

// GetRecoveryTime returns the RecoveryTime field value if set, zero value otherwise.
func (o *PcfBinding) GetRecoveryTime() time.Time {
	if o == nil || IsNil(o.RecoveryTime) {
		var ret time.Time
		return ret
	}
	return *o.RecoveryTime
}

// GetRecoveryTimeOk returns a tuple with the RecoveryTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetRecoveryTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.RecoveryTime) {
		return nil, false
	}
	return o.RecoveryTime, true
}

// HasRecoveryTime returns a boolean if a field has been set.
func (o *PcfBinding) HasRecoveryTime() bool {
	if o != nil && !IsNil(o.RecoveryTime) {
		return true
	}

	return false
}

// SetRecoveryTime gets a reference to the given time.Time and assigns it to the RecoveryTime field.
func (o *PcfBinding) SetRecoveryTime(v time.Time) {
	o.RecoveryTime = &v
}

// GetParaCom returns the ParaCom field value if set, zero value otherwise.
func (o *PcfBinding) GetParaCom() ParameterCombination {
	if o == nil || IsNil(o.ParaCom) {
		var ret ParameterCombination
		return ret
	}
	return *o.ParaCom
}

// GetParaComOk returns a tuple with the ParaCom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetParaComOk() (*ParameterCombination, bool) {
	if o == nil || IsNil(o.ParaCom) {
		return nil, false
	}
	return o.ParaCom, true
}

// HasParaCom returns a boolean if a field has been set.
func (o *PcfBinding) HasParaCom() bool {
	if o != nil && !IsNil(o.ParaCom) {
		return true
	}

	return false
}

// SetParaCom gets a reference to the given ParameterCombination and assigns it to the ParaCom field.
func (o *PcfBinding) SetParaCom(v ParameterCombination) {
	o.ParaCom = &v
}

// GetBindLevel returns the BindLevel field value if set, zero value otherwise.
func (o *PcfBinding) GetBindLevel() BindingLevel {
	if o == nil || IsNil(o.BindLevel) {
		var ret BindingLevel
		return ret
	}
	return *o.BindLevel
}

// GetBindLevelOk returns a tuple with the BindLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetBindLevelOk() (*BindingLevel, bool) {
	if o == nil || IsNil(o.BindLevel) {
		return nil, false
	}
	return o.BindLevel, true
}

// HasBindLevel returns a boolean if a field has been set.
func (o *PcfBinding) HasBindLevel() bool {
	if o != nil && !IsNil(o.BindLevel) {
		return true
	}

	return false
}

// SetBindLevel gets a reference to the given BindingLevel and assigns it to the BindLevel field.
func (o *PcfBinding) SetBindLevel(v BindingLevel) {
	o.BindLevel = &v
}

// GetIpv4FrameRouteList returns the Ipv4FrameRouteList field value if set, zero value otherwise.
func (o *PcfBinding) GetIpv4FrameRouteList() []string {
	if o == nil || IsNil(o.Ipv4FrameRouteList) {
		var ret []string
		return ret
	}
	return o.Ipv4FrameRouteList
}

// GetIpv4FrameRouteListOk returns a tuple with the Ipv4FrameRouteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetIpv4FrameRouteListOk() ([]string, bool) {
	if o == nil || IsNil(o.Ipv4FrameRouteList) {
		return nil, false
	}
	return o.Ipv4FrameRouteList, true
}

// HasIpv4FrameRouteList returns a boolean if a field has been set.
func (o *PcfBinding) HasIpv4FrameRouteList() bool {
	if o != nil && !IsNil(o.Ipv4FrameRouteList) {
		return true
	}

	return false
}

// SetIpv4FrameRouteList gets a reference to the given []string and assigns it to the Ipv4FrameRouteList field.
func (o *PcfBinding) SetIpv4FrameRouteList(v []string) {
	o.Ipv4FrameRouteList = v
}

// GetIpv6FrameRouteList returns the Ipv6FrameRouteList field value if set, zero value otherwise.
func (o *PcfBinding) GetIpv6FrameRouteList() []Ipv6Prefix {
	if o == nil || IsNil(o.Ipv6FrameRouteList) {
		var ret []Ipv6Prefix
		return ret
	}
	return o.Ipv6FrameRouteList
}

// GetIpv6FrameRouteListOk returns a tuple with the Ipv6FrameRouteList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PcfBinding) GetIpv6FrameRouteListOk() ([]Ipv6Prefix, bool) {
	if o == nil || IsNil(o.Ipv6FrameRouteList) {
		return nil, false
	}
	return o.Ipv6FrameRouteList, true
}

// HasIpv6FrameRouteList returns a boolean if a field has been set.
func (o *PcfBinding) HasIpv6FrameRouteList() bool {
	if o != nil && !IsNil(o.Ipv6FrameRouteList) {
		return true
	}

	return false
}

// SetIpv6FrameRouteList gets a reference to the given []Ipv6Prefix and assigns it to the Ipv6FrameRouteList field.
func (o *PcfBinding) SetIpv6FrameRouteList(v []Ipv6Prefix) {
	o.Ipv6FrameRouteList = v
}

func (o PcfBinding) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PcfBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !IsNil(o.Ipv6Prefix) {
		toSerialize["ipv6Prefix"] = o.Ipv6Prefix
	}
	if !IsNil(o.AddIpv6Prefixes) {
		toSerialize["addIpv6Prefixes"] = o.AddIpv6Prefixes
	}
	if !IsNil(o.IpDomain) {
		toSerialize["ipDomain"] = o.IpDomain
	}
	if !IsNil(o.MacAddr48) {
		toSerialize["macAddr48"] = o.MacAddr48
	}
	if !IsNil(o.AddMacAddrs) {
		toSerialize["addMacAddrs"] = o.AddMacAddrs
	}
	toSerialize["dnn"] = o.Dnn
	if !IsNil(o.PcfFqdn) {
		toSerialize["pcfFqdn"] = o.PcfFqdn
	}
	if !IsNil(o.PcfIpEndPoints) {
		toSerialize["pcfIpEndPoints"] = o.PcfIpEndPoints
	}
	if !IsNil(o.PcfDiamHost) {
		toSerialize["pcfDiamHost"] = o.PcfDiamHost
	}
	if !IsNil(o.PcfDiamRealm) {
		toSerialize["pcfDiamRealm"] = o.PcfDiamRealm
	}
	if !IsNil(o.PcfSmFqdn) {
		toSerialize["pcfSmFqdn"] = o.PcfSmFqdn
	}
	if !IsNil(o.PcfSmIpEndPoints) {
		toSerialize["pcfSmIpEndPoints"] = o.PcfSmIpEndPoints
	}
	toSerialize["snssai"] = o.Snssai
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !IsNil(o.PcfId) {
		toSerialize["pcfId"] = o.PcfId
	}
	if !IsNil(o.PcfSetId) {
		toSerialize["pcfSetId"] = o.PcfSetId
	}
	if !IsNil(o.RecoveryTime) {
		toSerialize["recoveryTime"] = o.RecoveryTime
	}
	if !IsNil(o.ParaCom) {
		toSerialize["paraCom"] = o.ParaCom
	}
	if !IsNil(o.BindLevel) {
		toSerialize["bindLevel"] = o.BindLevel
	}
	if !IsNil(o.Ipv4FrameRouteList) {
		toSerialize["ipv4FrameRouteList"] = o.Ipv4FrameRouteList
	}
	if !IsNil(o.Ipv6FrameRouteList) {
		toSerialize["ipv6FrameRouteList"] = o.Ipv6FrameRouteList
	}
	return toSerialize, nil
}

type NullablePcfBinding struct {
	value *PcfBinding
	isSet bool
}

func (v NullablePcfBinding) Get() *PcfBinding {
	return v.value
}

func (v *NullablePcfBinding) Set(val *PcfBinding) {
	v.value = val
	v.isSet = true
}

func (v NullablePcfBinding) IsSet() bool {
	return v.isSet
}

func (v *NullablePcfBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePcfBinding(val *PcfBinding) *NullablePcfBinding {
	return &NullablePcfBinding{value: val, isSet: true}
}

func (v NullablePcfBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePcfBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
