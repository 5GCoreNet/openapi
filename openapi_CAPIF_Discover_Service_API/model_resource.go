/*
CAPIF_Discover_Service_API

API for discovering service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Discover_Service_API

import (
	"encoding/json"
)

// checks if the Resource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Resource{}

// Resource Represents the API resource data.
type Resource struct {
	// Resource name
	ResourceName string `json:"resourceName"`
	CommType CommunicationType `json:"commType"`
	// Relative URI of the API resource, it is set as {apiSpecificSuffixes} part of the URI structure as defined in clause 5.2.4 of 3GPP TS 29.122. 
	Uri string `json:"uri"`
	// it is set as {custOpName} part of the URI structure for a custom operation associated with a resource as defined in clause 5.2.4 of 3GPP TS 29.122. 
	CustOpName *string `json:"custOpName,omitempty"`
	// Custom operations associated with this resource. 
	CustOperations []CustomOperation `json:"custOperations,omitempty"`
	// Supported HTTP methods for the API resource. Only applicable when the protocol in AefProfile indicates HTTP. 
	Operations []Operation `json:"operations,omitempty"`
	// Text description of the API resource
	Description *string `json:"description,omitempty"`
}

// NewResource instantiates a new Resource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResource(resourceName string, commType CommunicationType, uri string) *Resource {
	this := Resource{}
	this.ResourceName = resourceName
	this.CommType = commType
	this.Uri = uri
	return &this
}

// NewResourceWithDefaults instantiates a new Resource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceWithDefaults() *Resource {
	this := Resource{}
	return &this
}

// GetResourceName returns the ResourceName field value
func (o *Resource) GetResourceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value
// and a boolean to check if the value has been set.
func (o *Resource) GetResourceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceName, true
}

// SetResourceName sets field value
func (o *Resource) SetResourceName(v string) {
	o.ResourceName = v
}

// GetCommType returns the CommType field value
func (o *Resource) GetCommType() CommunicationType {
	if o == nil {
		var ret CommunicationType
		return ret
	}

	return o.CommType
}

// GetCommTypeOk returns a tuple with the CommType field value
// and a boolean to check if the value has been set.
func (o *Resource) GetCommTypeOk() (*CommunicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CommType, true
}

// SetCommType sets field value
func (o *Resource) SetCommType(v CommunicationType) {
	o.CommType = v
}

// GetUri returns the Uri field value
func (o *Resource) GetUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uri
}

// GetUriOk returns a tuple with the Uri field value
// and a boolean to check if the value has been set.
func (o *Resource) GetUriOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uri, true
}

// SetUri sets field value
func (o *Resource) SetUri(v string) {
	o.Uri = v
}

// GetCustOpName returns the CustOpName field value if set, zero value otherwise.
func (o *Resource) GetCustOpName() string {
	if o == nil || IsNil(o.CustOpName) {
		var ret string
		return ret
	}
	return *o.CustOpName
}

// GetCustOpNameOk returns a tuple with the CustOpName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetCustOpNameOk() (*string, bool) {
	if o == nil || IsNil(o.CustOpName) {
		return nil, false
	}
	return o.CustOpName, true
}

// HasCustOpName returns a boolean if a field has been set.
func (o *Resource) HasCustOpName() bool {
	if o != nil && !IsNil(o.CustOpName) {
		return true
	}

	return false
}

// SetCustOpName gets a reference to the given string and assigns it to the CustOpName field.
func (o *Resource) SetCustOpName(v string) {
	o.CustOpName = &v
}

// GetCustOperations returns the CustOperations field value if set, zero value otherwise.
func (o *Resource) GetCustOperations() []CustomOperation {
	if o == nil || IsNil(o.CustOperations) {
		var ret []CustomOperation
		return ret
	}
	return o.CustOperations
}

// GetCustOperationsOk returns a tuple with the CustOperations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetCustOperationsOk() ([]CustomOperation, bool) {
	if o == nil || IsNil(o.CustOperations) {
		return nil, false
	}
	return o.CustOperations, true
}

// HasCustOperations returns a boolean if a field has been set.
func (o *Resource) HasCustOperations() bool {
	if o != nil && !IsNil(o.CustOperations) {
		return true
	}

	return false
}

// SetCustOperations gets a reference to the given []CustomOperation and assigns it to the CustOperations field.
func (o *Resource) SetCustOperations(v []CustomOperation) {
	o.CustOperations = v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *Resource) GetOperations() []Operation {
	if o == nil || IsNil(o.Operations) {
		var ret []Operation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetOperationsOk() ([]Operation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *Resource) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []Operation and assigns it to the Operations field.
func (o *Resource) SetOperations(v []Operation) {
	o.Operations = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Resource) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resource) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Resource) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Resource) SetDescription(v string) {
	o.Description = &v
}

func (o Resource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["resourceName"] = o.ResourceName
	toSerialize["commType"] = o.CommType
	toSerialize["uri"] = o.Uri
	if !IsNil(o.CustOpName) {
		toSerialize["custOpName"] = o.CustOpName
	}
	if !IsNil(o.CustOperations) {
		toSerialize["custOperations"] = o.CustOperations
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableResource struct {
	value *Resource
	isSet bool
}

func (v NullableResource) Get() *Resource {
	return v.value
}

func (v *NullableResource) Set(val *Resource) {
	v.value = val
	v.isSet = true
}

func (v NullableResource) IsSet() bool {
	return v.isSet
}

func (v *NullableResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResource(val *Resource) *NullableResource {
	return &NullableResource{value: val, isSet: true}
}

func (v NullableResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


