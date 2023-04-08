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

// checks if the N1MessageNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &N1MessageNotification{}

// N1MessageNotification Data within a N1 message notification request
type N1MessageNotification struct {
	N1NotifySubscriptionId *string `json:"n1NotifySubscriptionId,omitempty"`
	N1MessageContainer N1MessageContainer `json:"n1MessageContainer"`
	// LCS Correlation ID.
	LcsCorrelationId *string `json:"lcsCorrelationId,omitempty"`
	RegistrationCtxtContainer *RegistrationContextContainer `json:"registrationCtxtContainer,omitempty"`
	// LMF identification.
	NewLmfIdentification *string `json:"newLmfIdentification,omitempty"`
	Guami *Guami `json:"guami,omitempty"`
	CIoT5GSOptimisation *bool `json:"cIoT5GSOptimisation,omitempty"`
	Ecgi *Ecgi `json:"ecgi,omitempty"`
	Ncgi *Ncgi `json:"ncgi,omitempty"`
}

// NewN1MessageNotification instantiates a new N1MessageNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewN1MessageNotification(n1MessageContainer N1MessageContainer) *N1MessageNotification {
	this := N1MessageNotification{}
	this.N1MessageContainer = n1MessageContainer
	var cIoT5GSOptimisation bool = false
	this.CIoT5GSOptimisation = &cIoT5GSOptimisation
	return &this
}

// NewN1MessageNotificationWithDefaults instantiates a new N1MessageNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewN1MessageNotificationWithDefaults() *N1MessageNotification {
	this := N1MessageNotification{}
	var cIoT5GSOptimisation bool = false
	this.CIoT5GSOptimisation = &cIoT5GSOptimisation
	return &this
}

// GetN1NotifySubscriptionId returns the N1NotifySubscriptionId field value if set, zero value otherwise.
func (o *N1MessageNotification) GetN1NotifySubscriptionId() string {
	if o == nil || isNil(o.N1NotifySubscriptionId) {
		var ret string
		return ret
	}
	return *o.N1NotifySubscriptionId
}

// GetN1NotifySubscriptionIdOk returns a tuple with the N1NotifySubscriptionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetN1NotifySubscriptionIdOk() (*string, bool) {
	if o == nil || isNil(o.N1NotifySubscriptionId) {
		return nil, false
	}
	return o.N1NotifySubscriptionId, true
}

// HasN1NotifySubscriptionId returns a boolean if a field has been set.
func (o *N1MessageNotification) HasN1NotifySubscriptionId() bool {
	if o != nil && !isNil(o.N1NotifySubscriptionId) {
		return true
	}

	return false
}

// SetN1NotifySubscriptionId gets a reference to the given string and assigns it to the N1NotifySubscriptionId field.
func (o *N1MessageNotification) SetN1NotifySubscriptionId(v string) {
	o.N1NotifySubscriptionId = &v
}

// GetN1MessageContainer returns the N1MessageContainer field value
func (o *N1MessageNotification) GetN1MessageContainer() N1MessageContainer {
	if o == nil {
		var ret N1MessageContainer
		return ret
	}

	return o.N1MessageContainer
}

// GetN1MessageContainerOk returns a tuple with the N1MessageContainer field value
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetN1MessageContainerOk() (*N1MessageContainer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.N1MessageContainer, true
}

// SetN1MessageContainer sets field value
func (o *N1MessageNotification) SetN1MessageContainer(v N1MessageContainer) {
	o.N1MessageContainer = v
}

// GetLcsCorrelationId returns the LcsCorrelationId field value if set, zero value otherwise.
func (o *N1MessageNotification) GetLcsCorrelationId() string {
	if o == nil || isNil(o.LcsCorrelationId) {
		var ret string
		return ret
	}
	return *o.LcsCorrelationId
}

// GetLcsCorrelationIdOk returns a tuple with the LcsCorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetLcsCorrelationIdOk() (*string, bool) {
	if o == nil || isNil(o.LcsCorrelationId) {
		return nil, false
	}
	return o.LcsCorrelationId, true
}

// HasLcsCorrelationId returns a boolean if a field has been set.
func (o *N1MessageNotification) HasLcsCorrelationId() bool {
	if o != nil && !isNil(o.LcsCorrelationId) {
		return true
	}

	return false
}

// SetLcsCorrelationId gets a reference to the given string and assigns it to the LcsCorrelationId field.
func (o *N1MessageNotification) SetLcsCorrelationId(v string) {
	o.LcsCorrelationId = &v
}

// GetRegistrationCtxtContainer returns the RegistrationCtxtContainer field value if set, zero value otherwise.
func (o *N1MessageNotification) GetRegistrationCtxtContainer() RegistrationContextContainer {
	if o == nil || isNil(o.RegistrationCtxtContainer) {
		var ret RegistrationContextContainer
		return ret
	}
	return *o.RegistrationCtxtContainer
}

// GetRegistrationCtxtContainerOk returns a tuple with the RegistrationCtxtContainer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetRegistrationCtxtContainerOk() (*RegistrationContextContainer, bool) {
	if o == nil || isNil(o.RegistrationCtxtContainer) {
		return nil, false
	}
	return o.RegistrationCtxtContainer, true
}

// HasRegistrationCtxtContainer returns a boolean if a field has been set.
func (o *N1MessageNotification) HasRegistrationCtxtContainer() bool {
	if o != nil && !isNil(o.RegistrationCtxtContainer) {
		return true
	}

	return false
}

// SetRegistrationCtxtContainer gets a reference to the given RegistrationContextContainer and assigns it to the RegistrationCtxtContainer field.
func (o *N1MessageNotification) SetRegistrationCtxtContainer(v RegistrationContextContainer) {
	o.RegistrationCtxtContainer = &v
}

// GetNewLmfIdentification returns the NewLmfIdentification field value if set, zero value otherwise.
func (o *N1MessageNotification) GetNewLmfIdentification() string {
	if o == nil || isNil(o.NewLmfIdentification) {
		var ret string
		return ret
	}
	return *o.NewLmfIdentification
}

// GetNewLmfIdentificationOk returns a tuple with the NewLmfIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetNewLmfIdentificationOk() (*string, bool) {
	if o == nil || isNil(o.NewLmfIdentification) {
		return nil, false
	}
	return o.NewLmfIdentification, true
}

// HasNewLmfIdentification returns a boolean if a field has been set.
func (o *N1MessageNotification) HasNewLmfIdentification() bool {
	if o != nil && !isNil(o.NewLmfIdentification) {
		return true
	}

	return false
}

// SetNewLmfIdentification gets a reference to the given string and assigns it to the NewLmfIdentification field.
func (o *N1MessageNotification) SetNewLmfIdentification(v string) {
	o.NewLmfIdentification = &v
}

// GetGuami returns the Guami field value if set, zero value otherwise.
func (o *N1MessageNotification) GetGuami() Guami {
	if o == nil || isNil(o.Guami) {
		var ret Guami
		return ret
	}
	return *o.Guami
}

// GetGuamiOk returns a tuple with the Guami field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetGuamiOk() (*Guami, bool) {
	if o == nil || isNil(o.Guami) {
		return nil, false
	}
	return o.Guami, true
}

// HasGuami returns a boolean if a field has been set.
func (o *N1MessageNotification) HasGuami() bool {
	if o != nil && !isNil(o.Guami) {
		return true
	}

	return false
}

// SetGuami gets a reference to the given Guami and assigns it to the Guami field.
func (o *N1MessageNotification) SetGuami(v Guami) {
	o.Guami = &v
}

// GetCIoT5GSOptimisation returns the CIoT5GSOptimisation field value if set, zero value otherwise.
func (o *N1MessageNotification) GetCIoT5GSOptimisation() bool {
	if o == nil || isNil(o.CIoT5GSOptimisation) {
		var ret bool
		return ret
	}
	return *o.CIoT5GSOptimisation
}

// GetCIoT5GSOptimisationOk returns a tuple with the CIoT5GSOptimisation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetCIoT5GSOptimisationOk() (*bool, bool) {
	if o == nil || isNil(o.CIoT5GSOptimisation) {
		return nil, false
	}
	return o.CIoT5GSOptimisation, true
}

// HasCIoT5GSOptimisation returns a boolean if a field has been set.
func (o *N1MessageNotification) HasCIoT5GSOptimisation() bool {
	if o != nil && !isNil(o.CIoT5GSOptimisation) {
		return true
	}

	return false
}

// SetCIoT5GSOptimisation gets a reference to the given bool and assigns it to the CIoT5GSOptimisation field.
func (o *N1MessageNotification) SetCIoT5GSOptimisation(v bool) {
	o.CIoT5GSOptimisation = &v
}

// GetEcgi returns the Ecgi field value if set, zero value otherwise.
func (o *N1MessageNotification) GetEcgi() Ecgi {
	if o == nil || isNil(o.Ecgi) {
		var ret Ecgi
		return ret
	}
	return *o.Ecgi
}

// GetEcgiOk returns a tuple with the Ecgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetEcgiOk() (*Ecgi, bool) {
	if o == nil || isNil(o.Ecgi) {
		return nil, false
	}
	return o.Ecgi, true
}

// HasEcgi returns a boolean if a field has been set.
func (o *N1MessageNotification) HasEcgi() bool {
	if o != nil && !isNil(o.Ecgi) {
		return true
	}

	return false
}

// SetEcgi gets a reference to the given Ecgi and assigns it to the Ecgi field.
func (o *N1MessageNotification) SetEcgi(v Ecgi) {
	o.Ecgi = &v
}

// GetNcgi returns the Ncgi field value if set, zero value otherwise.
func (o *N1MessageNotification) GetNcgi() Ncgi {
	if o == nil || isNil(o.Ncgi) {
		var ret Ncgi
		return ret
	}
	return *o.Ncgi
}

// GetNcgiOk returns a tuple with the Ncgi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *N1MessageNotification) GetNcgiOk() (*Ncgi, bool) {
	if o == nil || isNil(o.Ncgi) {
		return nil, false
	}
	return o.Ncgi, true
}

// HasNcgi returns a boolean if a field has been set.
func (o *N1MessageNotification) HasNcgi() bool {
	if o != nil && !isNil(o.Ncgi) {
		return true
	}

	return false
}

// SetNcgi gets a reference to the given Ncgi and assigns it to the Ncgi field.
func (o *N1MessageNotification) SetNcgi(v Ncgi) {
	o.Ncgi = &v
}

func (o N1MessageNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o N1MessageNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.N1NotifySubscriptionId) {
		toSerialize["n1NotifySubscriptionId"] = o.N1NotifySubscriptionId
	}
	toSerialize["n1MessageContainer"] = o.N1MessageContainer
	if !isNil(o.LcsCorrelationId) {
		toSerialize["lcsCorrelationId"] = o.LcsCorrelationId
	}
	if !isNil(o.RegistrationCtxtContainer) {
		toSerialize["registrationCtxtContainer"] = o.RegistrationCtxtContainer
	}
	if !isNil(o.NewLmfIdentification) {
		toSerialize["newLmfIdentification"] = o.NewLmfIdentification
	}
	if !isNil(o.Guami) {
		toSerialize["guami"] = o.Guami
	}
	if !isNil(o.CIoT5GSOptimisation) {
		toSerialize["cIoT5GSOptimisation"] = o.CIoT5GSOptimisation
	}
	if !isNil(o.Ecgi) {
		toSerialize["ecgi"] = o.Ecgi
	}
	if !isNil(o.Ncgi) {
		toSerialize["ncgi"] = o.Ncgi
	}
	return toSerialize, nil
}

type NullableN1MessageNotification struct {
	value *N1MessageNotification
	isSet bool
}

func (v NullableN1MessageNotification) Get() *N1MessageNotification {
	return v.value
}

func (v *NullableN1MessageNotification) Set(val *N1MessageNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableN1MessageNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableN1MessageNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1MessageNotification(val *N1MessageNotification) *NullableN1MessageNotification {
	return &NullableN1MessageNotification{value: val, isSet: true}
}

func (v NullableN1MessageNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1MessageNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


