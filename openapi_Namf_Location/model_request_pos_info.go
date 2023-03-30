/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
	"time"
)

// checks if the RequestPosInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RequestPosInfo{}

// RequestPosInfo Data within Provide Positioning Information Request
type RequestPosInfo struct {
	LcsClientType ExternalClientType `json:"lcsClientType"`
	LcsLocation LocationType `json:"lcsLocation"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	Priority *LcsPriority `json:"priority,omitempty"`
	LcsQoS *LocationQoS `json:"lcsQoS,omitempty"`
	VelocityRequested *VelocityRequested `json:"velocityRequested,omitempty"`
	LcsSupportedGADShapes *SupportedGADShapes `json:"lcsSupportedGADShapes,omitempty"`
	AdditionalLcsSuppGADShapes []SupportedGADShapes `json:"additionalLcsSuppGADShapes,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	LocationNotificationUri *string `json:"locationNotificationUri,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	OldGuami *Guami `json:"oldGuami,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	Pei *string `json:"pei,omitempty"`
	// LCS service type.
	LcsServiceType *int32 `json:"lcsServiceType,omitempty"`
	LdrType *LdrType `json:"ldrType,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	HgmlcCallBackURI *string `json:"hgmlcCallBackURI,omitempty"`
	// LDR Reference.
	LdrReference *string `json:"ldrReference,omitempty"`
	PeriodicEventInfo *PeriodicEventInfo `json:"periodicEventInfo,omitempty"`
	AreaEventInfo *AreaEventInfo `json:"areaEventInfo,omitempty"`
	MotionEventInfo *MotionEventInfo `json:"motionEventInfo,omitempty"`
	// Contains the external client identification
	ExternalClientIdentification *string `json:"externalClientIdentification,omitempty"`
	// String uniquely identifying a NF instance. The format of the NF Instance ID shall be a  Universally Unique Identifier (UUID) version 4, as described in IETF RFC 4122.  
	AfID *string `json:"afID,omitempty"`
	// Contains the codeword
	CodeWord *string `json:"codeWord,omitempty"`
	UePrivacyRequirements *UePrivacyRequirements `json:"uePrivacyRequirements,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ScheduledLocTime *time.Time `json:"scheduledLocTime,omitempty"`
	ReliableLocReq *bool `json:"reliableLocReq,omitempty"`
}

// NewRequestPosInfo instantiates a new RequestPosInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestPosInfo(lcsClientType ExternalClientType, lcsLocation LocationType) *RequestPosInfo {
	this := RequestPosInfo{}
	this.LcsClientType = lcsClientType
	this.LcsLocation = lcsLocation
	var reliableLocReq bool = false
	this.ReliableLocReq = &reliableLocReq
	return &this
}

// NewRequestPosInfoWithDefaults instantiates a new RequestPosInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestPosInfoWithDefaults() *RequestPosInfo {
	this := RequestPosInfo{}
	var reliableLocReq bool = false
	this.ReliableLocReq = &reliableLocReq
	return &this
}

// GetLcsClientType returns the LcsClientType field value
func (o *RequestPosInfo) GetLcsClientType() ExternalClientType {
	if o == nil {
		var ret ExternalClientType
		return ret
	}

	return o.LcsClientType
}

// GetLcsClientTypeOk returns a tuple with the LcsClientType field value
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetLcsClientTypeOk() (*ExternalClientType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LcsClientType, true
}

// SetLcsClientType sets field value
func (o *RequestPosInfo) SetLcsClientType(v ExternalClientType) {
	o.LcsClientType = v
}

// GetLcsLocation returns the LcsLocation field value
func (o *RequestPosInfo) GetLcsLocation() LocationType {
	if o == nil {
		var ret LocationType
		return ret
	}

	return o.LcsLocation
}

// GetLcsLocationOk returns a tuple with the LcsLocation field value
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetLcsLocationOk() (*LocationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LcsLocation, true
}

// SetLcsLocation sets field value
func (o *RequestPosInfo) SetLcsLocation(v LocationType) {
	o.LcsLocation = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *RequestPosInfo) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *RequestPosInfo) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *RequestPosInfo) SetSupi(v string) {
	o.Supi = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *RequestPosInfo) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *RequestPosInfo) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *RequestPosInfo) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *RequestPosInfo) GetPriority() LcsPriority {
	if o == nil || IsNil(o.Priority) {
		var ret LcsPriority
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetPriorityOk() (*LcsPriority, bool) {
	if o == nil || IsNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *RequestPosInfo) HasPriority() bool {
	if o != nil && !IsNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given LcsPriority and assigns it to the Priority field.
func (o *RequestPosInfo) SetPriority(v LcsPriority) {
	o.Priority = &v
}

// GetLcsQoS returns the LcsQoS field value if set, zero value otherwise.
func (o *RequestPosInfo) GetLcsQoS() LocationQoS {
	if o == nil || IsNil(o.LcsQoS) {
		var ret LocationQoS
		return ret
	}
	return *o.LcsQoS
}

// GetLcsQoSOk returns a tuple with the LcsQoS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetLcsQoSOk() (*LocationQoS, bool) {
	if o == nil || IsNil(o.LcsQoS) {
		return nil, false
	}
	return o.LcsQoS, true
}

// HasLcsQoS returns a boolean if a field has been set.
func (o *RequestPosInfo) HasLcsQoS() bool {
	if o != nil && !IsNil(o.LcsQoS) {
		return true
	}

	return false
}

// SetLcsQoS gets a reference to the given LocationQoS and assigns it to the LcsQoS field.
func (o *RequestPosInfo) SetLcsQoS(v LocationQoS) {
	o.LcsQoS = &v
}

// GetVelocityRequested returns the VelocityRequested field value if set, zero value otherwise.
func (o *RequestPosInfo) GetVelocityRequested() VelocityRequested {
	if o == nil || IsNil(o.VelocityRequested) {
		var ret VelocityRequested
		return ret
	}
	return *o.VelocityRequested
}

// GetVelocityRequestedOk returns a tuple with the VelocityRequested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetVelocityRequestedOk() (*VelocityRequested, bool) {
	if o == nil || IsNil(o.VelocityRequested) {
		return nil, false
	}
	return o.VelocityRequested, true
}

// HasVelocityRequested returns a boolean if a field has been set.
func (o *RequestPosInfo) HasVelocityRequested() bool {
	if o != nil && !IsNil(o.VelocityRequested) {
		return true
	}

	return false
}

// SetVelocityRequested gets a reference to the given VelocityRequested and assigns it to the VelocityRequested field.
func (o *RequestPosInfo) SetVelocityRequested(v VelocityRequested) {
	o.VelocityRequested = &v
}

// GetLcsSupportedGADShapes returns the LcsSupportedGADShapes field value if set, zero value otherwise.
func (o *RequestPosInfo) GetLcsSupportedGADShapes() SupportedGADShapes {
	if o == nil || IsNil(o.LcsSupportedGADShapes) {
		var ret SupportedGADShapes
		return ret
	}
	return *o.LcsSupportedGADShapes
}

// GetLcsSupportedGADShapesOk returns a tuple with the LcsSupportedGADShapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetLcsSupportedGADShapesOk() (*SupportedGADShapes, bool) {
	if o == nil || IsNil(o.LcsSupportedGADShapes) {
		return nil, false
	}
	return o.LcsSupportedGADShapes, true
}

// HasLcsSupportedGADShapes returns a boolean if a field has been set.
func (o *RequestPosInfo) HasLcsSupportedGADShapes() bool {
	if o != nil && !IsNil(o.LcsSupportedGADShapes) {
		return true
	}

	return false
}

// SetLcsSupportedGADShapes gets a reference to the given SupportedGADShapes and assigns it to the LcsSupportedGADShapes field.
func (o *RequestPosInfo) SetLcsSupportedGADShapes(v SupportedGADShapes) {
	o.LcsSupportedGADShapes = &v
}

// GetAdditionalLcsSuppGADShapes returns the AdditionalLcsSuppGADShapes field value if set, zero value otherwise.
func (o *RequestPosInfo) GetAdditionalLcsSuppGADShapes() []SupportedGADShapes {
	if o == nil || IsNil(o.AdditionalLcsSuppGADShapes) {
		var ret []SupportedGADShapes
		return ret
	}
	return o.AdditionalLcsSuppGADShapes
}

// GetAdditionalLcsSuppGADShapesOk returns a tuple with the AdditionalLcsSuppGADShapes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetAdditionalLcsSuppGADShapesOk() ([]SupportedGADShapes, bool) {
	if o == nil || IsNil(o.AdditionalLcsSuppGADShapes) {
		return nil, false
	}
	return o.AdditionalLcsSuppGADShapes, true
}

// HasAdditionalLcsSuppGADShapes returns a boolean if a field has been set.
func (o *RequestPosInfo) HasAdditionalLcsSuppGADShapes() bool {
	if o != nil && !IsNil(o.AdditionalLcsSuppGADShapes) {
		return true
	}

	return false
}

// SetAdditionalLcsSuppGADShapes gets a reference to the given []SupportedGADShapes and assigns it to the AdditionalLcsSuppGADShapes field.
func (o *RequestPosInfo) SetAdditionalLcsSuppGADShapes(v []SupportedGADShapes) {
	o.AdditionalLcsSuppGADShapes = v
}

// GetLocationNotificationUri returns the LocationNotificationUri field value if set, zero value otherwise.
func (o *RequestPosInfo) GetLocationNotificationUri() string {
	if o == nil || IsNil(o.LocationNotificationUri) {
		var ret string
		return ret
	}
	return *o.LocationNotificationUri
}

// GetLocationNotificationUriOk returns a tuple with the LocationNotificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetLocationNotificationUriOk() (*string, bool) {
	if o == nil || IsNil(o.LocationNotificationUri) {
		return nil, false
	}
	return o.LocationNotificationUri, true
}

// HasLocationNotificationUri returns a boolean if a field has been set.
func (o *RequestPosInfo) HasLocationNotificationUri() bool {
	if o != nil && !IsNil(o.LocationNotificationUri) {
		return true
	}

	return false
}

// SetLocationNotificationUri gets a reference to the given string and assigns it to the LocationNotificationUri field.
func (o *RequestPosInfo) SetLocationNotificationUri(v string) {
	o.LocationNotificationUri = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *RequestPosInfo) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *RequestPosInfo) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *RequestPosInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetOldGuami returns the OldGuami field value if set, zero value otherwise.
func (o *RequestPosInfo) GetOldGuami() Guami {
	if o == nil || IsNil(o.OldGuami) {
		var ret Guami
		return ret
	}
	return *o.OldGuami
}

// GetOldGuamiOk returns a tuple with the OldGuami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetOldGuamiOk() (*Guami, bool) {
	if o == nil || IsNil(o.OldGuami) {
		return nil, false
	}
	return o.OldGuami, true
}

// HasOldGuami returns a boolean if a field has been set.
func (o *RequestPosInfo) HasOldGuami() bool {
	if o != nil && !IsNil(o.OldGuami) {
		return true
	}

	return false
}

// SetOldGuami gets a reference to the given Guami and assigns it to the OldGuami field.
func (o *RequestPosInfo) SetOldGuami(v Guami) {
	o.OldGuami = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *RequestPosInfo) GetPei() string {
	if o == nil || IsNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetPeiOk() (*string, bool) {
	if o == nil || IsNil(o.Pei) {
		return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *RequestPosInfo) HasPei() bool {
	if o != nil && !IsNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *RequestPosInfo) SetPei(v string) {
	o.Pei = &v
}

// GetLcsServiceType returns the LcsServiceType field value if set, zero value otherwise.
func (o *RequestPosInfo) GetLcsServiceType() int32 {
	if o == nil || IsNil(o.LcsServiceType) {
		var ret int32
		return ret
	}
	return *o.LcsServiceType
}

// GetLcsServiceTypeOk returns a tuple with the LcsServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetLcsServiceTypeOk() (*int32, bool) {
	if o == nil || IsNil(o.LcsServiceType) {
		return nil, false
	}
	return o.LcsServiceType, true
}

// HasLcsServiceType returns a boolean if a field has been set.
func (o *RequestPosInfo) HasLcsServiceType() bool {
	if o != nil && !IsNil(o.LcsServiceType) {
		return true
	}

	return false
}

// SetLcsServiceType gets a reference to the given int32 and assigns it to the LcsServiceType field.
func (o *RequestPosInfo) SetLcsServiceType(v int32) {
	o.LcsServiceType = &v
}

// GetLdrType returns the LdrType field value if set, zero value otherwise.
func (o *RequestPosInfo) GetLdrType() LdrType {
	if o == nil || IsNil(o.LdrType) {
		var ret LdrType
		return ret
	}
	return *o.LdrType
}

// GetLdrTypeOk returns a tuple with the LdrType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetLdrTypeOk() (*LdrType, bool) {
	if o == nil || IsNil(o.LdrType) {
		return nil, false
	}
	return o.LdrType, true
}

// HasLdrType returns a boolean if a field has been set.
func (o *RequestPosInfo) HasLdrType() bool {
	if o != nil && !IsNil(o.LdrType) {
		return true
	}

	return false
}

// SetLdrType gets a reference to the given LdrType and assigns it to the LdrType field.
func (o *RequestPosInfo) SetLdrType(v LdrType) {
	o.LdrType = &v
}

// GetHgmlcCallBackURI returns the HgmlcCallBackURI field value if set, zero value otherwise.
func (o *RequestPosInfo) GetHgmlcCallBackURI() string {
	if o == nil || IsNil(o.HgmlcCallBackURI) {
		var ret string
		return ret
	}
	return *o.HgmlcCallBackURI
}

// GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetHgmlcCallBackURIOk() (*string, bool) {
	if o == nil || IsNil(o.HgmlcCallBackURI) {
		return nil, false
	}
	return o.HgmlcCallBackURI, true
}

// HasHgmlcCallBackURI returns a boolean if a field has been set.
func (o *RequestPosInfo) HasHgmlcCallBackURI() bool {
	if o != nil && !IsNil(o.HgmlcCallBackURI) {
		return true
	}

	return false
}

// SetHgmlcCallBackURI gets a reference to the given string and assigns it to the HgmlcCallBackURI field.
func (o *RequestPosInfo) SetHgmlcCallBackURI(v string) {
	o.HgmlcCallBackURI = &v
}

// GetLdrReference returns the LdrReference field value if set, zero value otherwise.
func (o *RequestPosInfo) GetLdrReference() string {
	if o == nil || IsNil(o.LdrReference) {
		var ret string
		return ret
	}
	return *o.LdrReference
}

// GetLdrReferenceOk returns a tuple with the LdrReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetLdrReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.LdrReference) {
		return nil, false
	}
	return o.LdrReference, true
}

// HasLdrReference returns a boolean if a field has been set.
func (o *RequestPosInfo) HasLdrReference() bool {
	if o != nil && !IsNil(o.LdrReference) {
		return true
	}

	return false
}

// SetLdrReference gets a reference to the given string and assigns it to the LdrReference field.
func (o *RequestPosInfo) SetLdrReference(v string) {
	o.LdrReference = &v
}

// GetPeriodicEventInfo returns the PeriodicEventInfo field value if set, zero value otherwise.
func (o *RequestPosInfo) GetPeriodicEventInfo() PeriodicEventInfo {
	if o == nil || IsNil(o.PeriodicEventInfo) {
		var ret PeriodicEventInfo
		return ret
	}
	return *o.PeriodicEventInfo
}

// GetPeriodicEventInfoOk returns a tuple with the PeriodicEventInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetPeriodicEventInfoOk() (*PeriodicEventInfo, bool) {
	if o == nil || IsNil(o.PeriodicEventInfo) {
		return nil, false
	}
	return o.PeriodicEventInfo, true
}

// HasPeriodicEventInfo returns a boolean if a field has been set.
func (o *RequestPosInfo) HasPeriodicEventInfo() bool {
	if o != nil && !IsNil(o.PeriodicEventInfo) {
		return true
	}

	return false
}

// SetPeriodicEventInfo gets a reference to the given PeriodicEventInfo and assigns it to the PeriodicEventInfo field.
func (o *RequestPosInfo) SetPeriodicEventInfo(v PeriodicEventInfo) {
	o.PeriodicEventInfo = &v
}

// GetAreaEventInfo returns the AreaEventInfo field value if set, zero value otherwise.
func (o *RequestPosInfo) GetAreaEventInfo() AreaEventInfo {
	if o == nil || IsNil(o.AreaEventInfo) {
		var ret AreaEventInfo
		return ret
	}
	return *o.AreaEventInfo
}

// GetAreaEventInfoOk returns a tuple with the AreaEventInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetAreaEventInfoOk() (*AreaEventInfo, bool) {
	if o == nil || IsNil(o.AreaEventInfo) {
		return nil, false
	}
	return o.AreaEventInfo, true
}

// HasAreaEventInfo returns a boolean if a field has been set.
func (o *RequestPosInfo) HasAreaEventInfo() bool {
	if o != nil && !IsNil(o.AreaEventInfo) {
		return true
	}

	return false
}

// SetAreaEventInfo gets a reference to the given AreaEventInfo and assigns it to the AreaEventInfo field.
func (o *RequestPosInfo) SetAreaEventInfo(v AreaEventInfo) {
	o.AreaEventInfo = &v
}

// GetMotionEventInfo returns the MotionEventInfo field value if set, zero value otherwise.
func (o *RequestPosInfo) GetMotionEventInfo() MotionEventInfo {
	if o == nil || IsNil(o.MotionEventInfo) {
		var ret MotionEventInfo
		return ret
	}
	return *o.MotionEventInfo
}

// GetMotionEventInfoOk returns a tuple with the MotionEventInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetMotionEventInfoOk() (*MotionEventInfo, bool) {
	if o == nil || IsNil(o.MotionEventInfo) {
		return nil, false
	}
	return o.MotionEventInfo, true
}

// HasMotionEventInfo returns a boolean if a field has been set.
func (o *RequestPosInfo) HasMotionEventInfo() bool {
	if o != nil && !IsNil(o.MotionEventInfo) {
		return true
	}

	return false
}

// SetMotionEventInfo gets a reference to the given MotionEventInfo and assigns it to the MotionEventInfo field.
func (o *RequestPosInfo) SetMotionEventInfo(v MotionEventInfo) {
	o.MotionEventInfo = &v
}

// GetExternalClientIdentification returns the ExternalClientIdentification field value if set, zero value otherwise.
func (o *RequestPosInfo) GetExternalClientIdentification() string {
	if o == nil || IsNil(o.ExternalClientIdentification) {
		var ret string
		return ret
	}
	return *o.ExternalClientIdentification
}

// GetExternalClientIdentificationOk returns a tuple with the ExternalClientIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetExternalClientIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalClientIdentification) {
		return nil, false
	}
	return o.ExternalClientIdentification, true
}

// HasExternalClientIdentification returns a boolean if a field has been set.
func (o *RequestPosInfo) HasExternalClientIdentification() bool {
	if o != nil && !IsNil(o.ExternalClientIdentification) {
		return true
	}

	return false
}

// SetExternalClientIdentification gets a reference to the given string and assigns it to the ExternalClientIdentification field.
func (o *RequestPosInfo) SetExternalClientIdentification(v string) {
	o.ExternalClientIdentification = &v
}

// GetAfID returns the AfID field value if set, zero value otherwise.
func (o *RequestPosInfo) GetAfID() string {
	if o == nil || IsNil(o.AfID) {
		var ret string
		return ret
	}
	return *o.AfID
}

// GetAfIDOk returns a tuple with the AfID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetAfIDOk() (*string, bool) {
	if o == nil || IsNil(o.AfID) {
		return nil, false
	}
	return o.AfID, true
}

// HasAfID returns a boolean if a field has been set.
func (o *RequestPosInfo) HasAfID() bool {
	if o != nil && !IsNil(o.AfID) {
		return true
	}

	return false
}

// SetAfID gets a reference to the given string and assigns it to the AfID field.
func (o *RequestPosInfo) SetAfID(v string) {
	o.AfID = &v
}

// GetCodeWord returns the CodeWord field value if set, zero value otherwise.
func (o *RequestPosInfo) GetCodeWord() string {
	if o == nil || IsNil(o.CodeWord) {
		var ret string
		return ret
	}
	return *o.CodeWord
}

// GetCodeWordOk returns a tuple with the CodeWord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetCodeWordOk() (*string, bool) {
	if o == nil || IsNil(o.CodeWord) {
		return nil, false
	}
	return o.CodeWord, true
}

// HasCodeWord returns a boolean if a field has been set.
func (o *RequestPosInfo) HasCodeWord() bool {
	if o != nil && !IsNil(o.CodeWord) {
		return true
	}

	return false
}

// SetCodeWord gets a reference to the given string and assigns it to the CodeWord field.
func (o *RequestPosInfo) SetCodeWord(v string) {
	o.CodeWord = &v
}

// GetUePrivacyRequirements returns the UePrivacyRequirements field value if set, zero value otherwise.
func (o *RequestPosInfo) GetUePrivacyRequirements() UePrivacyRequirements {
	if o == nil || IsNil(o.UePrivacyRequirements) {
		var ret UePrivacyRequirements
		return ret
	}
	return *o.UePrivacyRequirements
}

// GetUePrivacyRequirementsOk returns a tuple with the UePrivacyRequirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetUePrivacyRequirementsOk() (*UePrivacyRequirements, bool) {
	if o == nil || IsNil(o.UePrivacyRequirements) {
		return nil, false
	}
	return o.UePrivacyRequirements, true
}

// HasUePrivacyRequirements returns a boolean if a field has been set.
func (o *RequestPosInfo) HasUePrivacyRequirements() bool {
	if o != nil && !IsNil(o.UePrivacyRequirements) {
		return true
	}

	return false
}

// SetUePrivacyRequirements gets a reference to the given UePrivacyRequirements and assigns it to the UePrivacyRequirements field.
func (o *RequestPosInfo) SetUePrivacyRequirements(v UePrivacyRequirements) {
	o.UePrivacyRequirements = &v
}

// GetScheduledLocTime returns the ScheduledLocTime field value if set, zero value otherwise.
func (o *RequestPosInfo) GetScheduledLocTime() time.Time {
	if o == nil || IsNil(o.ScheduledLocTime) {
		var ret time.Time
		return ret
	}
	return *o.ScheduledLocTime
}

// GetScheduledLocTimeOk returns a tuple with the ScheduledLocTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetScheduledLocTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ScheduledLocTime) {
		return nil, false
	}
	return o.ScheduledLocTime, true
}

// HasScheduledLocTime returns a boolean if a field has been set.
func (o *RequestPosInfo) HasScheduledLocTime() bool {
	if o != nil && !IsNil(o.ScheduledLocTime) {
		return true
	}

	return false
}

// SetScheduledLocTime gets a reference to the given time.Time and assigns it to the ScheduledLocTime field.
func (o *RequestPosInfo) SetScheduledLocTime(v time.Time) {
	o.ScheduledLocTime = &v
}

// GetReliableLocReq returns the ReliableLocReq field value if set, zero value otherwise.
func (o *RequestPosInfo) GetReliableLocReq() bool {
	if o == nil || IsNil(o.ReliableLocReq) {
		var ret bool
		return ret
	}
	return *o.ReliableLocReq
}

// GetReliableLocReqOk returns a tuple with the ReliableLocReq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestPosInfo) GetReliableLocReqOk() (*bool, bool) {
	if o == nil || IsNil(o.ReliableLocReq) {
		return nil, false
	}
	return o.ReliableLocReq, true
}

// HasReliableLocReq returns a boolean if a field has been set.
func (o *RequestPosInfo) HasReliableLocReq() bool {
	if o != nil && !IsNil(o.ReliableLocReq) {
		return true
	}

	return false
}

// SetReliableLocReq gets a reference to the given bool and assigns it to the ReliableLocReq field.
func (o *RequestPosInfo) SetReliableLocReq(v bool) {
	o.ReliableLocReq = &v
}

func (o RequestPosInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RequestPosInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["lcsClientType"] = o.LcsClientType
	toSerialize["lcsLocation"] = o.LcsLocation
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !IsNil(o.LcsQoS) {
		toSerialize["lcsQoS"] = o.LcsQoS
	}
	if !IsNil(o.VelocityRequested) {
		toSerialize["velocityRequested"] = o.VelocityRequested
	}
	if !IsNil(o.LcsSupportedGADShapes) {
		toSerialize["lcsSupportedGADShapes"] = o.LcsSupportedGADShapes
	}
	if !IsNil(o.AdditionalLcsSuppGADShapes) {
		toSerialize["additionalLcsSuppGADShapes"] = o.AdditionalLcsSuppGADShapes
	}
	if !IsNil(o.LocationNotificationUri) {
		toSerialize["locationNotificationUri"] = o.LocationNotificationUri
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.OldGuami) {
		toSerialize["oldGuami"] = o.OldGuami
	}
	if !IsNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if !IsNil(o.LcsServiceType) {
		toSerialize["lcsServiceType"] = o.LcsServiceType
	}
	if !IsNil(o.LdrType) {
		toSerialize["ldrType"] = o.LdrType
	}
	if !IsNil(o.HgmlcCallBackURI) {
		toSerialize["hgmlcCallBackURI"] = o.HgmlcCallBackURI
	}
	if !IsNil(o.LdrReference) {
		toSerialize["ldrReference"] = o.LdrReference
	}
	if !IsNil(o.PeriodicEventInfo) {
		toSerialize["periodicEventInfo"] = o.PeriodicEventInfo
	}
	if !IsNil(o.AreaEventInfo) {
		toSerialize["areaEventInfo"] = o.AreaEventInfo
	}
	if !IsNil(o.MotionEventInfo) {
		toSerialize["motionEventInfo"] = o.MotionEventInfo
	}
	if !IsNil(o.ExternalClientIdentification) {
		toSerialize["externalClientIdentification"] = o.ExternalClientIdentification
	}
	if !IsNil(o.AfID) {
		toSerialize["afID"] = o.AfID
	}
	if !IsNil(o.CodeWord) {
		toSerialize["codeWord"] = o.CodeWord
	}
	if !IsNil(o.UePrivacyRequirements) {
		toSerialize["uePrivacyRequirements"] = o.UePrivacyRequirements
	}
	if !IsNil(o.ScheduledLocTime) {
		toSerialize["scheduledLocTime"] = o.ScheduledLocTime
	}
	if !IsNil(o.ReliableLocReq) {
		toSerialize["reliableLocReq"] = o.ReliableLocReq
	}
	return toSerialize, nil
}

type NullableRequestPosInfo struct {
	value *RequestPosInfo
	isSet bool
}

func (v NullableRequestPosInfo) Get() *RequestPosInfo {
	return v.value
}

func (v *NullableRequestPosInfo) Set(val *RequestPosInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestPosInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestPosInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestPosInfo(val *RequestPosInfo) *NullableRequestPosInfo {
	return &NullableRequestPosInfo{value: val, isSet: true}
}

func (v NullableRequestPosInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestPosInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


