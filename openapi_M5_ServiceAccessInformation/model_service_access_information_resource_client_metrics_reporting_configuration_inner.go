/*
M5_ServiceAccessInformation

5GMS AF M5 Service Access Information API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 2.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_M5_ServiceAccessInformation

import (
	"encoding/json"
)

// checks if the ServiceAccessInformationResourceClientMetricsReportingConfigurationInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceAccessInformationResourceClientMetricsReportingConfigurationInner{}

// ServiceAccessInformationResourceClientMetricsReportingConfigurationInner struct for ServiceAccessInformationResourceClientMetricsReportingConfigurationInner
type ServiceAccessInformationResourceClientMetricsReportingConfigurationInner struct {
	// A set of application endpoint addresses.
	ServerAddresses []string `json:"serverAddresses"`
	// String providing an URI formatted according to RFC 3986.
	Scheme string `json:"scheme"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\").
	DataNetworkName *string `json:"dataNetworkName,omitempty"`
	// indicating a time in seconds.
	ReportingInterval *int32   `json:"reportingInterval,omitempty"`
	SamplePercentage  float32  `json:"samplePercentage"`
	UrlFilters        []string `json:"urlFilters"`
	Metrics           []string `json:"metrics"`
}

// NewServiceAccessInformationResourceClientMetricsReportingConfigurationInner instantiates a new ServiceAccessInformationResourceClientMetricsReportingConfigurationInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceAccessInformationResourceClientMetricsReportingConfigurationInner(serverAddresses []string, scheme string, samplePercentage float32, urlFilters []string, metrics []string) *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner {
	this := ServiceAccessInformationResourceClientMetricsReportingConfigurationInner{}
	this.ServerAddresses = serverAddresses
	this.Scheme = scheme
	this.SamplePercentage = samplePercentage
	this.UrlFilters = urlFilters
	this.Metrics = metrics
	return &this
}

// NewServiceAccessInformationResourceClientMetricsReportingConfigurationInnerWithDefaults instantiates a new ServiceAccessInformationResourceClientMetricsReportingConfigurationInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceAccessInformationResourceClientMetricsReportingConfigurationInnerWithDefaults() *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner {
	this := ServiceAccessInformationResourceClientMetricsReportingConfigurationInner{}
	return &this
}

// GetServerAddresses returns the ServerAddresses field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetServerAddresses() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ServerAddresses
}

// GetServerAddressesOk returns a tuple with the ServerAddresses field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetServerAddressesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ServerAddresses, true
}

// SetServerAddresses sets field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetServerAddresses(v []string) {
	o.ServerAddresses = v
}

// GetScheme returns the Scheme field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scheme, true
}

// SetScheme sets field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetScheme(v string) {
	o.Scheme = v
}

// GetDataNetworkName returns the DataNetworkName field value if set, zero value otherwise.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetDataNetworkName() string {
	if o == nil || IsNil(o.DataNetworkName) {
		var ret string
		return ret
	}
	return *o.DataNetworkName
}

// GetDataNetworkNameOk returns a tuple with the DataNetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetDataNetworkNameOk() (*string, bool) {
	if o == nil || IsNil(o.DataNetworkName) {
		return nil, false
	}
	return o.DataNetworkName, true
}

// HasDataNetworkName returns a boolean if a field has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) HasDataNetworkName() bool {
	if o != nil && !IsNil(o.DataNetworkName) {
		return true
	}

	return false
}

// SetDataNetworkName gets a reference to the given string and assigns it to the DataNetworkName field.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetDataNetworkName(v string) {
	o.DataNetworkName = &v
}

// GetReportingInterval returns the ReportingInterval field value if set, zero value otherwise.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetReportingInterval() int32 {
	if o == nil || IsNil(o.ReportingInterval) {
		var ret int32
		return ret
	}
	return *o.ReportingInterval
}

// GetReportingIntervalOk returns a tuple with the ReportingInterval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetReportingIntervalOk() (*int32, bool) {
	if o == nil || IsNil(o.ReportingInterval) {
		return nil, false
	}
	return o.ReportingInterval, true
}

// HasReportingInterval returns a boolean if a field has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) HasReportingInterval() bool {
	if o != nil && !IsNil(o.ReportingInterval) {
		return true
	}

	return false
}

// SetReportingInterval gets a reference to the given int32 and assigns it to the ReportingInterval field.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetReportingInterval(v int32) {
	o.ReportingInterval = &v
}

// GetSamplePercentage returns the SamplePercentage field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetSamplePercentage() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.SamplePercentage
}

// GetSamplePercentageOk returns a tuple with the SamplePercentage field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetSamplePercentageOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SamplePercentage, true
}

// SetSamplePercentage sets field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetSamplePercentage(v float32) {
	o.SamplePercentage = v
}

// GetUrlFilters returns the UrlFilters field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetUrlFilters() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.UrlFilters
}

// GetUrlFiltersOk returns a tuple with the UrlFilters field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetUrlFiltersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.UrlFilters, true
}

// SetUrlFilters sets field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetUrlFilters(v []string) {
	o.UrlFilters = v
}

// GetMetrics returns the Metrics field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetMetrics() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) GetMetricsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Metrics, true
}

// SetMetrics sets field value
func (o *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) SetMetrics(v []string) {
	o.Metrics = v
}

func (o ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["serverAddresses"] = o.ServerAddresses
	toSerialize["scheme"] = o.Scheme
	if !IsNil(o.DataNetworkName) {
		toSerialize["dataNetworkName"] = o.DataNetworkName
	}
	if !IsNil(o.ReportingInterval) {
		toSerialize["reportingInterval"] = o.ReportingInterval
	}
	toSerialize["samplePercentage"] = o.SamplePercentage
	toSerialize["urlFilters"] = o.UrlFilters
	toSerialize["metrics"] = o.Metrics
	return toSerialize, nil
}

type NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner struct {
	value *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner
	isSet bool
}

func (v NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner) Get() *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner {
	return v.value
}

func (v *NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner) Set(val *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner(val *ServiceAccessInformationResourceClientMetricsReportingConfigurationInner) *NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner {
	return &NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner{value: val, isSet: true}
}

func (v NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceAccessInformationResourceClientMetricsReportingConfigurationInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
