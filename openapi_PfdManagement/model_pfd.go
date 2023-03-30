/*
3gpp-pfd-management

API for PFD management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PfdManagement

import (
	"encoding/json"
)

// checks if the Pfd type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Pfd{}

// Pfd Represents a PFD for an external Application Identifier.
type Pfd struct {
	// Identifies a PDF of an application identifier.
	PfdId string `json:"pfdId"`
	// Represents a 3-tuple with protocol, server ip and server port for UL/DL application traffic. The content of the string has the same encoding as the IPFilterRule AVP value as defined in IETF RFC 6733.
	FlowDescriptions []string `json:"flowDescriptions,omitempty"`
	// Indicates a URL or a regular expression which is used to match the significant parts of the URL.
	Urls []string `json:"urls,omitempty"`
	// Indicates an FQDN or a regular expression as a domain name matching criteria.
	DomainNames []string `json:"domainNames,omitempty"`
	DnProtocol *DomainNameProtocol `json:"dnProtocol,omitempty"`
}

// NewPfd instantiates a new Pfd object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPfd(pfdId string) *Pfd {
	this := Pfd{}
	this.PfdId = pfdId
	return &this
}

// NewPfdWithDefaults instantiates a new Pfd object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPfdWithDefaults() *Pfd {
	this := Pfd{}
	return &this
}

// GetPfdId returns the PfdId field value
func (o *Pfd) GetPfdId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PfdId
}

// GetPfdIdOk returns a tuple with the PfdId field value
// and a boolean to check if the value has been set.
func (o *Pfd) GetPfdIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PfdId, true
}

// SetPfdId sets field value
func (o *Pfd) SetPfdId(v string) {
	o.PfdId = v
}

// GetFlowDescriptions returns the FlowDescriptions field value if set, zero value otherwise.
func (o *Pfd) GetFlowDescriptions() []string {
	if o == nil || IsNil(o.FlowDescriptions) {
		var ret []string
		return ret
	}
	return o.FlowDescriptions
}

// GetFlowDescriptionsOk returns a tuple with the FlowDescriptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pfd) GetFlowDescriptionsOk() ([]string, bool) {
	if o == nil || IsNil(o.FlowDescriptions) {
		return nil, false
	}
	return o.FlowDescriptions, true
}

// HasFlowDescriptions returns a boolean if a field has been set.
func (o *Pfd) HasFlowDescriptions() bool {
	if o != nil && !IsNil(o.FlowDescriptions) {
		return true
	}

	return false
}

// SetFlowDescriptions gets a reference to the given []string and assigns it to the FlowDescriptions field.
func (o *Pfd) SetFlowDescriptions(v []string) {
	o.FlowDescriptions = v
}

// GetUrls returns the Urls field value if set, zero value otherwise.
func (o *Pfd) GetUrls() []string {
	if o == nil || IsNil(o.Urls) {
		var ret []string
		return ret
	}
	return o.Urls
}

// GetUrlsOk returns a tuple with the Urls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pfd) GetUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.Urls) {
		return nil, false
	}
	return o.Urls, true
}

// HasUrls returns a boolean if a field has been set.
func (o *Pfd) HasUrls() bool {
	if o != nil && !IsNil(o.Urls) {
		return true
	}

	return false
}

// SetUrls gets a reference to the given []string and assigns it to the Urls field.
func (o *Pfd) SetUrls(v []string) {
	o.Urls = v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *Pfd) GetDomainNames() []string {
	if o == nil || IsNil(o.DomainNames) {
		var ret []string
		return ret
	}
	return o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pfd) GetDomainNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.DomainNames) {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *Pfd) HasDomainNames() bool {
	if o != nil && !IsNil(o.DomainNames) {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *Pfd) SetDomainNames(v []string) {
	o.DomainNames = v
}

// GetDnProtocol returns the DnProtocol field value if set, zero value otherwise.
func (o *Pfd) GetDnProtocol() DomainNameProtocol {
	if o == nil || IsNil(o.DnProtocol) {
		var ret DomainNameProtocol
		return ret
	}
	return *o.DnProtocol
}

// GetDnProtocolOk returns a tuple with the DnProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pfd) GetDnProtocolOk() (*DomainNameProtocol, bool) {
	if o == nil || IsNil(o.DnProtocol) {
		return nil, false
	}
	return o.DnProtocol, true
}

// HasDnProtocol returns a boolean if a field has been set.
func (o *Pfd) HasDnProtocol() bool {
	if o != nil && !IsNil(o.DnProtocol) {
		return true
	}

	return false
}

// SetDnProtocol gets a reference to the given DomainNameProtocol and assigns it to the DnProtocol field.
func (o *Pfd) SetDnProtocol(v DomainNameProtocol) {
	o.DnProtocol = &v
}

func (o Pfd) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Pfd) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pfdId"] = o.PfdId
	if !IsNil(o.FlowDescriptions) {
		toSerialize["flowDescriptions"] = o.FlowDescriptions
	}
	if !IsNil(o.Urls) {
		toSerialize["urls"] = o.Urls
	}
	if !IsNil(o.DomainNames) {
		toSerialize["domainNames"] = o.DomainNames
	}
	if !IsNil(o.DnProtocol) {
		toSerialize["dnProtocol"] = o.DnProtocol
	}
	return toSerialize, nil
}

type NullablePfd struct {
	value *Pfd
	isSet bool
}

func (v NullablePfd) Get() *Pfd {
	return v.value
}

func (v *NullablePfd) Set(val *Pfd) {
	v.value = val
	v.isSet = true
}

func (v NullablePfd) IsSet() bool {
	return v.isSet
}

func (v *NullablePfd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePfd(val *Pfd) *NullablePfd {
	return &NullablePfd{value: val, isSet: true}
}

func (v NullablePfd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePfd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


