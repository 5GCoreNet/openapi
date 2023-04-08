/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the FlowInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowInformation{}

// FlowInformation Contains the flow information.
type FlowInformation struct {
	// Defines a packet filter for an IP flow.
	FlowDescription *string `json:"flowDescription,omitempty"`
	EthFlowDescription *EthFlowDescription `json:"ethFlowDescription,omitempty"`
	// An identifier of packet filter.
	PackFiltId *string `json:"packFiltId,omitempty"`
	// The packet shall be sent to the UE.
	PacketFilterUsage *bool `json:"packetFilterUsage,omitempty"`
	// Contains the Ipv4 Type-of-Service and mask field or the Ipv6 Traffic-Class field and  mask field. 
	TosTrafficClass NullableString `json:"tosTrafficClass,omitempty"`
	// the security parameter index of the IPSec packet.
	Spi NullableString `json:"spi,omitempty"`
	// the Ipv6 flow label header field.
	FlowLabel NullableString `json:"flowLabel,omitempty"`
	FlowDirection *FlowDirectionRm `json:"flowDirection,omitempty"`
}

// NewFlowInformation instantiates a new FlowInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowInformation() *FlowInformation {
	this := FlowInformation{}
	return &this
}

// NewFlowInformationWithDefaults instantiates a new FlowInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowInformationWithDefaults() *FlowInformation {
	this := FlowInformation{}
	return &this
}

// GetFlowDescription returns the FlowDescription field value if set, zero value otherwise.
func (o *FlowInformation) GetFlowDescription() string {
	if o == nil || isNil(o.FlowDescription) {
		var ret string
		return ret
	}
	return *o.FlowDescription
}

// GetFlowDescriptionOk returns a tuple with the FlowDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInformation) GetFlowDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.FlowDescription) {
		return nil, false
	}
	return o.FlowDescription, true
}

// HasFlowDescription returns a boolean if a field has been set.
func (o *FlowInformation) HasFlowDescription() bool {
	if o != nil && !isNil(o.FlowDescription) {
		return true
	}

	return false
}

// SetFlowDescription gets a reference to the given string and assigns it to the FlowDescription field.
func (o *FlowInformation) SetFlowDescription(v string) {
	o.FlowDescription = &v
}

// GetEthFlowDescription returns the EthFlowDescription field value if set, zero value otherwise.
func (o *FlowInformation) GetEthFlowDescription() EthFlowDescription {
	if o == nil || isNil(o.EthFlowDescription) {
		var ret EthFlowDescription
		return ret
	}
	return *o.EthFlowDescription
}

// GetEthFlowDescriptionOk returns a tuple with the EthFlowDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInformation) GetEthFlowDescriptionOk() (*EthFlowDescription, bool) {
	if o == nil || isNil(o.EthFlowDescription) {
		return nil, false
	}
	return o.EthFlowDescription, true
}

// HasEthFlowDescription returns a boolean if a field has been set.
func (o *FlowInformation) HasEthFlowDescription() bool {
	if o != nil && !isNil(o.EthFlowDescription) {
		return true
	}

	return false
}

// SetEthFlowDescription gets a reference to the given EthFlowDescription and assigns it to the EthFlowDescription field.
func (o *FlowInformation) SetEthFlowDescription(v EthFlowDescription) {
	o.EthFlowDescription = &v
}

// GetPackFiltId returns the PackFiltId field value if set, zero value otherwise.
func (o *FlowInformation) GetPackFiltId() string {
	if o == nil || isNil(o.PackFiltId) {
		var ret string
		return ret
	}
	return *o.PackFiltId
}

// GetPackFiltIdOk returns a tuple with the PackFiltId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInformation) GetPackFiltIdOk() (*string, bool) {
	if o == nil || isNil(o.PackFiltId) {
		return nil, false
	}
	return o.PackFiltId, true
}

// HasPackFiltId returns a boolean if a field has been set.
func (o *FlowInformation) HasPackFiltId() bool {
	if o != nil && !isNil(o.PackFiltId) {
		return true
	}

	return false
}

// SetPackFiltId gets a reference to the given string and assigns it to the PackFiltId field.
func (o *FlowInformation) SetPackFiltId(v string) {
	o.PackFiltId = &v
}

// GetPacketFilterUsage returns the PacketFilterUsage field value if set, zero value otherwise.
func (o *FlowInformation) GetPacketFilterUsage() bool {
	if o == nil || isNil(o.PacketFilterUsage) {
		var ret bool
		return ret
	}
	return *o.PacketFilterUsage
}

// GetPacketFilterUsageOk returns a tuple with the PacketFilterUsage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInformation) GetPacketFilterUsageOk() (*bool, bool) {
	if o == nil || isNil(o.PacketFilterUsage) {
		return nil, false
	}
	return o.PacketFilterUsage, true
}

// HasPacketFilterUsage returns a boolean if a field has been set.
func (o *FlowInformation) HasPacketFilterUsage() bool {
	if o != nil && !isNil(o.PacketFilterUsage) {
		return true
	}

	return false
}

// SetPacketFilterUsage gets a reference to the given bool and assigns it to the PacketFilterUsage field.
func (o *FlowInformation) SetPacketFilterUsage(v bool) {
	o.PacketFilterUsage = &v
}

// GetTosTrafficClass returns the TosTrafficClass field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowInformation) GetTosTrafficClass() string {
	if o == nil || isNil(o.TosTrafficClass.Get()) {
		var ret string
		return ret
	}
	return *o.TosTrafficClass.Get()
}

// GetTosTrafficClassOk returns a tuple with the TosTrafficClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowInformation) GetTosTrafficClassOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TosTrafficClass.Get(), o.TosTrafficClass.IsSet()
}

// HasTosTrafficClass returns a boolean if a field has been set.
func (o *FlowInformation) HasTosTrafficClass() bool {
	if o != nil && o.TosTrafficClass.IsSet() {
		return true
	}

	return false
}

// SetTosTrafficClass gets a reference to the given NullableString and assigns it to the TosTrafficClass field.
func (o *FlowInformation) SetTosTrafficClass(v string) {
	o.TosTrafficClass.Set(&v)
}
// SetTosTrafficClassNil sets the value for TosTrafficClass to be an explicit nil
func (o *FlowInformation) SetTosTrafficClassNil() {
	o.TosTrafficClass.Set(nil)
}

// UnsetTosTrafficClass ensures that no value is present for TosTrafficClass, not even an explicit nil
func (o *FlowInformation) UnsetTosTrafficClass() {
	o.TosTrafficClass.Unset()
}

// GetSpi returns the Spi field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowInformation) GetSpi() string {
	if o == nil || isNil(o.Spi.Get()) {
		var ret string
		return ret
	}
	return *o.Spi.Get()
}

// GetSpiOk returns a tuple with the Spi field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowInformation) GetSpiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Spi.Get(), o.Spi.IsSet()
}

// HasSpi returns a boolean if a field has been set.
func (o *FlowInformation) HasSpi() bool {
	if o != nil && o.Spi.IsSet() {
		return true
	}

	return false
}

// SetSpi gets a reference to the given NullableString and assigns it to the Spi field.
func (o *FlowInformation) SetSpi(v string) {
	o.Spi.Set(&v)
}
// SetSpiNil sets the value for Spi to be an explicit nil
func (o *FlowInformation) SetSpiNil() {
	o.Spi.Set(nil)
}

// UnsetSpi ensures that no value is present for Spi, not even an explicit nil
func (o *FlowInformation) UnsetSpi() {
	o.Spi.Unset()
}

// GetFlowLabel returns the FlowLabel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FlowInformation) GetFlowLabel() string {
	if o == nil || isNil(o.FlowLabel.Get()) {
		var ret string
		return ret
	}
	return *o.FlowLabel.Get()
}

// GetFlowLabelOk returns a tuple with the FlowLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FlowInformation) GetFlowLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowLabel.Get(), o.FlowLabel.IsSet()
}

// HasFlowLabel returns a boolean if a field has been set.
func (o *FlowInformation) HasFlowLabel() bool {
	if o != nil && o.FlowLabel.IsSet() {
		return true
	}

	return false
}

// SetFlowLabel gets a reference to the given NullableString and assigns it to the FlowLabel field.
func (o *FlowInformation) SetFlowLabel(v string) {
	o.FlowLabel.Set(&v)
}
// SetFlowLabelNil sets the value for FlowLabel to be an explicit nil
func (o *FlowInformation) SetFlowLabelNil() {
	o.FlowLabel.Set(nil)
}

// UnsetFlowLabel ensures that no value is present for FlowLabel, not even an explicit nil
func (o *FlowInformation) UnsetFlowLabel() {
	o.FlowLabel.Unset()
}

// GetFlowDirection returns the FlowDirection field value if set, zero value otherwise.
func (o *FlowInformation) GetFlowDirection() FlowDirectionRm {
	if o == nil || isNil(o.FlowDirection) {
		var ret FlowDirectionRm
		return ret
	}
	return *o.FlowDirection
}

// GetFlowDirectionOk returns a tuple with the FlowDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlowInformation) GetFlowDirectionOk() (*FlowDirectionRm, bool) {
	if o == nil || isNil(o.FlowDirection) {
		return nil, false
	}
	return o.FlowDirection, true
}

// HasFlowDirection returns a boolean if a field has been set.
func (o *FlowInformation) HasFlowDirection() bool {
	if o != nil && !isNil(o.FlowDirection) {
		return true
	}

	return false
}

// SetFlowDirection gets a reference to the given FlowDirectionRm and assigns it to the FlowDirection field.
func (o *FlowInformation) SetFlowDirection(v FlowDirectionRm) {
	o.FlowDirection = &v
}

func (o FlowInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.FlowDescription) {
		toSerialize["flowDescription"] = o.FlowDescription
	}
	if !isNil(o.EthFlowDescription) {
		toSerialize["ethFlowDescription"] = o.EthFlowDescription
	}
	if !isNil(o.PackFiltId) {
		toSerialize["packFiltId"] = o.PackFiltId
	}
	if !isNil(o.PacketFilterUsage) {
		toSerialize["packetFilterUsage"] = o.PacketFilterUsage
	}
	if o.TosTrafficClass.IsSet() {
		toSerialize["tosTrafficClass"] = o.TosTrafficClass.Get()
	}
	if o.Spi.IsSet() {
		toSerialize["spi"] = o.Spi.Get()
	}
	if o.FlowLabel.IsSet() {
		toSerialize["flowLabel"] = o.FlowLabel.Get()
	}
	if !isNil(o.FlowDirection) {
		toSerialize["flowDirection"] = o.FlowDirection
	}
	return toSerialize, nil
}

type NullableFlowInformation struct {
	value *FlowInformation
	isSet bool
}

func (v NullableFlowInformation) Get() *FlowInformation {
	return v.value
}

func (v *NullableFlowInformation) Set(val *FlowInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowInformation(val *FlowInformation) *NullableFlowInformation {
	return &NullableFlowInformation{value: val, isSet: true}
}

func (v NullableFlowInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


