/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the ManagedNFServiceSingleAllOfAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManagedNFServiceSingleAllOfAttributes{}

// ManagedNFServiceSingleAllOfAttributes struct for ManagedNFServiceSingleAllOfAttributes
type ManagedNFServiceSingleAllOfAttributes struct {
	UserLabel           *string              `json:"userLabel,omitempty"`
	NFServiceType       *NFServiceType       `json:"nFServiceType,omitempty"`
	SAP                 *SAP                 `json:"sAP,omitempty"`
	Operations          []Operation          `json:"operations,omitempty"`
	AdministrativeState *AdministrativeState `json:"administrativeState,omitempty"`
	OperationalState    *OperationalState    `json:"operationalState,omitempty"`
	UsageState          *UsageState          `json:"usageState,omitempty"`
	RegistrationState   *RegistrationState   `json:"registrationState,omitempty"`
}

// NewManagedNFServiceSingleAllOfAttributes instantiates a new ManagedNFServiceSingleAllOfAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedNFServiceSingleAllOfAttributes() *ManagedNFServiceSingleAllOfAttributes {
	this := ManagedNFServiceSingleAllOfAttributes{}
	return &this
}

// NewManagedNFServiceSingleAllOfAttributesWithDefaults instantiates a new ManagedNFServiceSingleAllOfAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedNFServiceSingleAllOfAttributesWithDefaults() *ManagedNFServiceSingleAllOfAttributes {
	this := ManagedNFServiceSingleAllOfAttributes{}
	return &this
}

// GetUserLabel returns the UserLabel field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOfAttributes) GetUserLabel() string {
	if o == nil || IsNil(o.UserLabel) {
		var ret string
		return ret
	}
	return *o.UserLabel
}

// GetUserLabelOk returns a tuple with the UserLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) GetUserLabelOk() (*string, bool) {
	if o == nil || IsNil(o.UserLabel) {
		return nil, false
	}
	return o.UserLabel, true
}

// HasUserLabel returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) HasUserLabel() bool {
	if o != nil && !IsNil(o.UserLabel) {
		return true
	}

	return false
}

// SetUserLabel gets a reference to the given string and assigns it to the UserLabel field.
func (o *ManagedNFServiceSingleAllOfAttributes) SetUserLabel(v string) {
	o.UserLabel = &v
}

// GetNFServiceType returns the NFServiceType field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOfAttributes) GetNFServiceType() NFServiceType {
	if o == nil || IsNil(o.NFServiceType) {
		var ret NFServiceType
		return ret
	}
	return *o.NFServiceType
}

// GetNFServiceTypeOk returns a tuple with the NFServiceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) GetNFServiceTypeOk() (*NFServiceType, bool) {
	if o == nil || IsNil(o.NFServiceType) {
		return nil, false
	}
	return o.NFServiceType, true
}

// HasNFServiceType returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) HasNFServiceType() bool {
	if o != nil && !IsNil(o.NFServiceType) {
		return true
	}

	return false
}

// SetNFServiceType gets a reference to the given NFServiceType and assigns it to the NFServiceType field.
func (o *ManagedNFServiceSingleAllOfAttributes) SetNFServiceType(v NFServiceType) {
	o.NFServiceType = &v
}

// GetSAP returns the SAP field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOfAttributes) GetSAP() SAP {
	if o == nil || IsNil(o.SAP) {
		var ret SAP
		return ret
	}
	return *o.SAP
}

// GetSAPOk returns a tuple with the SAP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) GetSAPOk() (*SAP, bool) {
	if o == nil || IsNil(o.SAP) {
		return nil, false
	}
	return o.SAP, true
}

// HasSAP returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) HasSAP() bool {
	if o != nil && !IsNil(o.SAP) {
		return true
	}

	return false
}

// SetSAP gets a reference to the given SAP and assigns it to the SAP field.
func (o *ManagedNFServiceSingleAllOfAttributes) SetSAP(v SAP) {
	o.SAP = &v
}

// GetOperations returns the Operations field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOfAttributes) GetOperations() []Operation {
	if o == nil || IsNil(o.Operations) {
		var ret []Operation
		return ret
	}
	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) GetOperationsOk() ([]Operation, bool) {
	if o == nil || IsNil(o.Operations) {
		return nil, false
	}
	return o.Operations, true
}

// HasOperations returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) HasOperations() bool {
	if o != nil && !IsNil(o.Operations) {
		return true
	}

	return false
}

// SetOperations gets a reference to the given []Operation and assigns it to the Operations field.
func (o *ManagedNFServiceSingleAllOfAttributes) SetOperations(v []Operation) {
	o.Operations = v
}

// GetAdministrativeState returns the AdministrativeState field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOfAttributes) GetAdministrativeState() AdministrativeState {
	if o == nil || IsNil(o.AdministrativeState) {
		var ret AdministrativeState
		return ret
	}
	return *o.AdministrativeState
}

// GetAdministrativeStateOk returns a tuple with the AdministrativeState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) GetAdministrativeStateOk() (*AdministrativeState, bool) {
	if o == nil || IsNil(o.AdministrativeState) {
		return nil, false
	}
	return o.AdministrativeState, true
}

// HasAdministrativeState returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) HasAdministrativeState() bool {
	if o != nil && !IsNil(o.AdministrativeState) {
		return true
	}

	return false
}

// SetAdministrativeState gets a reference to the given AdministrativeState and assigns it to the AdministrativeState field.
func (o *ManagedNFServiceSingleAllOfAttributes) SetAdministrativeState(v AdministrativeState) {
	o.AdministrativeState = &v
}

// GetOperationalState returns the OperationalState field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOfAttributes) GetOperationalState() OperationalState {
	if o == nil || IsNil(o.OperationalState) {
		var ret OperationalState
		return ret
	}
	return *o.OperationalState
}

// GetOperationalStateOk returns a tuple with the OperationalState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) GetOperationalStateOk() (*OperationalState, bool) {
	if o == nil || IsNil(o.OperationalState) {
		return nil, false
	}
	return o.OperationalState, true
}

// HasOperationalState returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) HasOperationalState() bool {
	if o != nil && !IsNil(o.OperationalState) {
		return true
	}

	return false
}

// SetOperationalState gets a reference to the given OperationalState and assigns it to the OperationalState field.
func (o *ManagedNFServiceSingleAllOfAttributes) SetOperationalState(v OperationalState) {
	o.OperationalState = &v
}

// GetUsageState returns the UsageState field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOfAttributes) GetUsageState() UsageState {
	if o == nil || IsNil(o.UsageState) {
		var ret UsageState
		return ret
	}
	return *o.UsageState
}

// GetUsageStateOk returns a tuple with the UsageState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) GetUsageStateOk() (*UsageState, bool) {
	if o == nil || IsNil(o.UsageState) {
		return nil, false
	}
	return o.UsageState, true
}

// HasUsageState returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) HasUsageState() bool {
	if o != nil && !IsNil(o.UsageState) {
		return true
	}

	return false
}

// SetUsageState gets a reference to the given UsageState and assigns it to the UsageState field.
func (o *ManagedNFServiceSingleAllOfAttributes) SetUsageState(v UsageState) {
	o.UsageState = &v
}

// GetRegistrationState returns the RegistrationState field value if set, zero value otherwise.
func (o *ManagedNFServiceSingleAllOfAttributes) GetRegistrationState() RegistrationState {
	if o == nil || IsNil(o.RegistrationState) {
		var ret RegistrationState
		return ret
	}
	return *o.RegistrationState
}

// GetRegistrationStateOk returns a tuple with the RegistrationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) GetRegistrationStateOk() (*RegistrationState, bool) {
	if o == nil || IsNil(o.RegistrationState) {
		return nil, false
	}
	return o.RegistrationState, true
}

// HasRegistrationState returns a boolean if a field has been set.
func (o *ManagedNFServiceSingleAllOfAttributes) HasRegistrationState() bool {
	if o != nil && !IsNil(o.RegistrationState) {
		return true
	}

	return false
}

// SetRegistrationState gets a reference to the given RegistrationState and assigns it to the RegistrationState field.
func (o *ManagedNFServiceSingleAllOfAttributes) SetRegistrationState(v RegistrationState) {
	o.RegistrationState = &v
}

func (o ManagedNFServiceSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManagedNFServiceSingleAllOfAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserLabel) {
		toSerialize["userLabel"] = o.UserLabel
	}
	if !IsNil(o.NFServiceType) {
		toSerialize["nFServiceType"] = o.NFServiceType
	}
	if !IsNil(o.SAP) {
		toSerialize["sAP"] = o.SAP
	}
	if !IsNil(o.Operations) {
		toSerialize["operations"] = o.Operations
	}
	if !IsNil(o.AdministrativeState) {
		toSerialize["administrativeState"] = o.AdministrativeState
	}
	if !IsNil(o.OperationalState) {
		toSerialize["operationalState"] = o.OperationalState
	}
	if !IsNil(o.UsageState) {
		toSerialize["usageState"] = o.UsageState
	}
	if !IsNil(o.RegistrationState) {
		toSerialize["registrationState"] = o.RegistrationState
	}
	return toSerialize, nil
}

type NullableManagedNFServiceSingleAllOfAttributes struct {
	value *ManagedNFServiceSingleAllOfAttributes
	isSet bool
}

func (v NullableManagedNFServiceSingleAllOfAttributes) Get() *ManagedNFServiceSingleAllOfAttributes {
	return v.value
}

func (v *NullableManagedNFServiceSingleAllOfAttributes) Set(val *ManagedNFServiceSingleAllOfAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedNFServiceSingleAllOfAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedNFServiceSingleAllOfAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedNFServiceSingleAllOfAttributes(val *ManagedNFServiceSingleAllOfAttributes) *NullableManagedNFServiceSingleAllOfAttributes {
	return &NullableManagedNFServiceSingleAllOfAttributes{value: val, isSet: true}
}

func (v NullableManagedNFServiceSingleAllOfAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedNFServiceSingleAllOfAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
