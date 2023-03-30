/*
AUSF API

AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UEAuthentication

import (
	"encoding/json"
)

// checks if the AuthenticationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationInfo{}

// AuthenticationInfo Contains the UE id (i.e. SUCI or SUPI) and the Serving Network Name.
type AuthenticationInfo struct {
	// String identifying a SUPI or a SUCI.
	SupiOrSuci string `json:"supiOrSuci"`
	ServingNetworkName string `json:"servingNetworkName"`
	ResynchronizationInfo *ResynchronizationInfo `json:"resynchronizationInfo,omitempty"`
	// String representing a Permanent Equipment Identifier that may contain - an IMEI or IMEISV, as  specified in clause 6.2 of 3GPP TS 23.003; a MAC address for a 5G-RG or FN-RG via  wireline  access, with an indication that this address cannot be trusted for regulatory purpose if this  address cannot be used as an Equipment Identifier of the FN-RG, as specified in clause 4.7.7  of 3GPP TS23.316. Examples are imei-012345678901234 or imeisv-0123456789012345.  
	Pei *string `json:"pei,omitempty"`
	TraceData NullableTraceData `json:"traceData,omitempty"`
	// Identifier of a group of NFs.
	UdmGroupId *string `json:"udmGroupId,omitempty"`
	RoutingIndicator *string `json:"routingIndicator,omitempty"`
	CellCagInfo []string `json:"cellCagInfo,omitempty"`
	N5gcInd *bool `json:"n5gcInd,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	NswoInd *bool `json:"nswoInd,omitempty"`
	DisasterRoamingInd *bool `json:"disasterRoamingInd,omitempty"`
	OnboardingInd *bool `json:"onboardingInd,omitempty"`
}

// NewAuthenticationInfo instantiates a new AuthenticationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationInfo(supiOrSuci string, servingNetworkName string) *AuthenticationInfo {
	this := AuthenticationInfo{}
	this.SupiOrSuci = supiOrSuci
	this.ServingNetworkName = servingNetworkName
	var n5gcInd bool = false
	this.N5gcInd = &n5gcInd
	var nswoInd bool = false
	this.NswoInd = &nswoInd
	var disasterRoamingInd bool = false
	this.DisasterRoamingInd = &disasterRoamingInd
	var onboardingInd bool = false
	this.OnboardingInd = &onboardingInd
	return &this
}

// NewAuthenticationInfoWithDefaults instantiates a new AuthenticationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationInfoWithDefaults() *AuthenticationInfo {
	this := AuthenticationInfo{}
	var n5gcInd bool = false
	this.N5gcInd = &n5gcInd
	var nswoInd bool = false
	this.NswoInd = &nswoInd
	var disasterRoamingInd bool = false
	this.DisasterRoamingInd = &disasterRoamingInd
	var onboardingInd bool = false
	this.OnboardingInd = &onboardingInd
	return &this
}

// GetSupiOrSuci returns the SupiOrSuci field value
func (o *AuthenticationInfo) GetSupiOrSuci() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SupiOrSuci
}

// GetSupiOrSuciOk returns a tuple with the SupiOrSuci field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetSupiOrSuciOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SupiOrSuci, true
}

// SetSupiOrSuci sets field value
func (o *AuthenticationInfo) SetSupiOrSuci(v string) {
	o.SupiOrSuci = v
}

// GetServingNetworkName returns the ServingNetworkName field value
func (o *AuthenticationInfo) GetServingNetworkName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServingNetworkName
}

// GetServingNetworkNameOk returns a tuple with the ServingNetworkName field value
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetServingNetworkNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServingNetworkName, true
}

// SetServingNetworkName sets field value
func (o *AuthenticationInfo) SetServingNetworkName(v string) {
	o.ServingNetworkName = v
}

// GetResynchronizationInfo returns the ResynchronizationInfo field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetResynchronizationInfo() ResynchronizationInfo {
	if o == nil || IsNil(o.ResynchronizationInfo) {
		var ret ResynchronizationInfo
		return ret
	}
	return *o.ResynchronizationInfo
}

// GetResynchronizationInfoOk returns a tuple with the ResynchronizationInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetResynchronizationInfoOk() (*ResynchronizationInfo, bool) {
	if o == nil || IsNil(o.ResynchronizationInfo) {
		return nil, false
	}
	return o.ResynchronizationInfo, true
}

// HasResynchronizationInfo returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasResynchronizationInfo() bool {
	if o != nil && !IsNil(o.ResynchronizationInfo) {
		return true
	}

	return false
}

// SetResynchronizationInfo gets a reference to the given ResynchronizationInfo and assigns it to the ResynchronizationInfo field.
func (o *AuthenticationInfo) SetResynchronizationInfo(v ResynchronizationInfo) {
	o.ResynchronizationInfo = &v
}

// GetPei returns the Pei field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetPei() string {
	if o == nil || IsNil(o.Pei) {
		var ret string
		return ret
	}
	return *o.Pei
}

// GetPeiOk returns a tuple with the Pei field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetPeiOk() (*string, bool) {
	if o == nil || IsNil(o.Pei) {
		return nil, false
	}
	return o.Pei, true
}

// HasPei returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasPei() bool {
	if o != nil && !IsNil(o.Pei) {
		return true
	}

	return false
}

// SetPei gets a reference to the given string and assigns it to the Pei field.
func (o *AuthenticationInfo) SetPei(v string) {
	o.Pei = &v
}

// GetTraceData returns the TraceData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AuthenticationInfo) GetTraceData() TraceData {
	if o == nil || IsNil(o.TraceData.Get()) {
		var ret TraceData
		return ret
	}
	return *o.TraceData.Get()
}

// GetTraceDataOk returns a tuple with the TraceData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthenticationInfo) GetTraceDataOk() (*TraceData, bool) {
	if o == nil {
		return nil, false
	}
	return o.TraceData.Get(), o.TraceData.IsSet()
}

// HasTraceData returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasTraceData() bool {
	if o != nil && o.TraceData.IsSet() {
		return true
	}

	return false
}

// SetTraceData gets a reference to the given NullableTraceData and assigns it to the TraceData field.
func (o *AuthenticationInfo) SetTraceData(v TraceData) {
	o.TraceData.Set(&v)
}
// SetTraceDataNil sets the value for TraceData to be an explicit nil
func (o *AuthenticationInfo) SetTraceDataNil() {
	o.TraceData.Set(nil)
}

// UnsetTraceData ensures that no value is present for TraceData, not even an explicit nil
func (o *AuthenticationInfo) UnsetTraceData() {
	o.TraceData.Unset()
}

// GetUdmGroupId returns the UdmGroupId field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetUdmGroupId() string {
	if o == nil || IsNil(o.UdmGroupId) {
		var ret string
		return ret
	}
	return *o.UdmGroupId
}

// GetUdmGroupIdOk returns a tuple with the UdmGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetUdmGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.UdmGroupId) {
		return nil, false
	}
	return o.UdmGroupId, true
}

// HasUdmGroupId returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasUdmGroupId() bool {
	if o != nil && !IsNil(o.UdmGroupId) {
		return true
	}

	return false
}

// SetUdmGroupId gets a reference to the given string and assigns it to the UdmGroupId field.
func (o *AuthenticationInfo) SetUdmGroupId(v string) {
	o.UdmGroupId = &v
}

// GetRoutingIndicator returns the RoutingIndicator field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetRoutingIndicator() string {
	if o == nil || IsNil(o.RoutingIndicator) {
		var ret string
		return ret
	}
	return *o.RoutingIndicator
}

// GetRoutingIndicatorOk returns a tuple with the RoutingIndicator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetRoutingIndicatorOk() (*string, bool) {
	if o == nil || IsNil(o.RoutingIndicator) {
		return nil, false
	}
	return o.RoutingIndicator, true
}

// HasRoutingIndicator returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasRoutingIndicator() bool {
	if o != nil && !IsNil(o.RoutingIndicator) {
		return true
	}

	return false
}

// SetRoutingIndicator gets a reference to the given string and assigns it to the RoutingIndicator field.
func (o *AuthenticationInfo) SetRoutingIndicator(v string) {
	o.RoutingIndicator = &v
}

// GetCellCagInfo returns the CellCagInfo field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetCellCagInfo() []string {
	if o == nil || IsNil(o.CellCagInfo) {
		var ret []string
		return ret
	}
	return o.CellCagInfo
}

// GetCellCagInfoOk returns a tuple with the CellCagInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetCellCagInfoOk() ([]string, bool) {
	if o == nil || IsNil(o.CellCagInfo) {
		return nil, false
	}
	return o.CellCagInfo, true
}

// HasCellCagInfo returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasCellCagInfo() bool {
	if o != nil && !IsNil(o.CellCagInfo) {
		return true
	}

	return false
}

// SetCellCagInfo gets a reference to the given []string and assigns it to the CellCagInfo field.
func (o *AuthenticationInfo) SetCellCagInfo(v []string) {
	o.CellCagInfo = v
}

// GetN5gcInd returns the N5gcInd field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetN5gcInd() bool {
	if o == nil || IsNil(o.N5gcInd) {
		var ret bool
		return ret
	}
	return *o.N5gcInd
}

// GetN5gcIndOk returns a tuple with the N5gcInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetN5gcIndOk() (*bool, bool) {
	if o == nil || IsNil(o.N5gcInd) {
		return nil, false
	}
	return o.N5gcInd, true
}

// HasN5gcInd returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasN5gcInd() bool {
	if o != nil && !IsNil(o.N5gcInd) {
		return true
	}

	return false
}

// SetN5gcInd gets a reference to the given bool and assigns it to the N5gcInd field.
func (o *AuthenticationInfo) SetN5gcInd(v bool) {
	o.N5gcInd = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *AuthenticationInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetNswoInd returns the NswoInd field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetNswoInd() bool {
	if o == nil || IsNil(o.NswoInd) {
		var ret bool
		return ret
	}
	return *o.NswoInd
}

// GetNswoIndOk returns a tuple with the NswoInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetNswoIndOk() (*bool, bool) {
	if o == nil || IsNil(o.NswoInd) {
		return nil, false
	}
	return o.NswoInd, true
}

// HasNswoInd returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasNswoInd() bool {
	if o != nil && !IsNil(o.NswoInd) {
		return true
	}

	return false
}

// SetNswoInd gets a reference to the given bool and assigns it to the NswoInd field.
func (o *AuthenticationInfo) SetNswoInd(v bool) {
	o.NswoInd = &v
}

// GetDisasterRoamingInd returns the DisasterRoamingInd field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetDisasterRoamingInd() bool {
	if o == nil || IsNil(o.DisasterRoamingInd) {
		var ret bool
		return ret
	}
	return *o.DisasterRoamingInd
}

// GetDisasterRoamingIndOk returns a tuple with the DisasterRoamingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetDisasterRoamingIndOk() (*bool, bool) {
	if o == nil || IsNil(o.DisasterRoamingInd) {
		return nil, false
	}
	return o.DisasterRoamingInd, true
}

// HasDisasterRoamingInd returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasDisasterRoamingInd() bool {
	if o != nil && !IsNil(o.DisasterRoamingInd) {
		return true
	}

	return false
}

// SetDisasterRoamingInd gets a reference to the given bool and assigns it to the DisasterRoamingInd field.
func (o *AuthenticationInfo) SetDisasterRoamingInd(v bool) {
	o.DisasterRoamingInd = &v
}

// GetOnboardingInd returns the OnboardingInd field value if set, zero value otherwise.
func (o *AuthenticationInfo) GetOnboardingInd() bool {
	if o == nil || IsNil(o.OnboardingInd) {
		var ret bool
		return ret
	}
	return *o.OnboardingInd
}

// GetOnboardingIndOk returns a tuple with the OnboardingInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfo) GetOnboardingIndOk() (*bool, bool) {
	if o == nil || IsNil(o.OnboardingInd) {
		return nil, false
	}
	return o.OnboardingInd, true
}

// HasOnboardingInd returns a boolean if a field has been set.
func (o *AuthenticationInfo) HasOnboardingInd() bool {
	if o != nil && !IsNil(o.OnboardingInd) {
		return true
	}

	return false
}

// SetOnboardingInd gets a reference to the given bool and assigns it to the OnboardingInd field.
func (o *AuthenticationInfo) SetOnboardingInd(v bool) {
	o.OnboardingInd = &v
}

func (o AuthenticationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supiOrSuci"] = o.SupiOrSuci
	toSerialize["servingNetworkName"] = o.ServingNetworkName
	if !IsNil(o.ResynchronizationInfo) {
		toSerialize["resynchronizationInfo"] = o.ResynchronizationInfo
	}
	if !IsNil(o.Pei) {
		toSerialize["pei"] = o.Pei
	}
	if o.TraceData.IsSet() {
		toSerialize["traceData"] = o.TraceData.Get()
	}
	if !IsNil(o.UdmGroupId) {
		toSerialize["udmGroupId"] = o.UdmGroupId
	}
	if !IsNil(o.RoutingIndicator) {
		toSerialize["routingIndicator"] = o.RoutingIndicator
	}
	if !IsNil(o.CellCagInfo) {
		toSerialize["cellCagInfo"] = o.CellCagInfo
	}
	if !IsNil(o.N5gcInd) {
		toSerialize["n5gcInd"] = o.N5gcInd
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if !IsNil(o.NswoInd) {
		toSerialize["nswoInd"] = o.NswoInd
	}
	if !IsNil(o.DisasterRoamingInd) {
		toSerialize["disasterRoamingInd"] = o.DisasterRoamingInd
	}
	if !IsNil(o.OnboardingInd) {
		toSerialize["onboardingInd"] = o.OnboardingInd
	}
	return toSerialize, nil
}

type NullableAuthenticationInfo struct {
	value *AuthenticationInfo
	isSet bool
}

func (v NullableAuthenticationInfo) Get() *AuthenticationInfo {
	return v.value
}

func (v *NullableAuthenticationInfo) Set(val *AuthenticationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationInfo(val *AuthenticationInfo) *NullableAuthenticationInfo {
	return &NullableAuthenticationInfo{value: val, isSet: true}
}

func (v NullableAuthenticationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


