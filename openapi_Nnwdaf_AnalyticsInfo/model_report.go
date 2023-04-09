/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// Report - struct for Report
type Report struct {
	ChangeOfSupiPeiAssociationReport *ChangeOfSupiPeiAssociationReport
	CmInfoReport                     *CmInfoReport
	CnTypeChangeReport               *CnTypeChangeReport
	LocationReport                   *LocationReport
	LossConnectivityReport           *LossConnectivityReport
	PdnConnectivityStatReport        *PdnConnectivityStatReport
	RoamingStatusReport              *RoamingStatusReport
}

// ChangeOfSupiPeiAssociationReportAsReport is a convenience function that returns ChangeOfSupiPeiAssociationReport wrapped in Report
func ChangeOfSupiPeiAssociationReportAsReport(v *ChangeOfSupiPeiAssociationReport) Report {
	return Report{
		ChangeOfSupiPeiAssociationReport: v,
	}
}

// CmInfoReportAsReport is a convenience function that returns CmInfoReport wrapped in Report
func CmInfoReportAsReport(v *CmInfoReport) Report {
	return Report{
		CmInfoReport: v,
	}
}

// CnTypeChangeReportAsReport is a convenience function that returns CnTypeChangeReport wrapped in Report
func CnTypeChangeReportAsReport(v *CnTypeChangeReport) Report {
	return Report{
		CnTypeChangeReport: v,
	}
}

// LocationReportAsReport is a convenience function that returns LocationReport wrapped in Report
func LocationReportAsReport(v *LocationReport) Report {
	return Report{
		LocationReport: v,
	}
}

// LossConnectivityReportAsReport is a convenience function that returns LossConnectivityReport wrapped in Report
func LossConnectivityReportAsReport(v *LossConnectivityReport) Report {
	return Report{
		LossConnectivityReport: v,
	}
}

// PdnConnectivityStatReportAsReport is a convenience function that returns PdnConnectivityStatReport wrapped in Report
func PdnConnectivityStatReportAsReport(v *PdnConnectivityStatReport) Report {
	return Report{
		PdnConnectivityStatReport: v,
	}
}

// RoamingStatusReportAsReport is a convenience function that returns RoamingStatusReport wrapped in Report
func RoamingStatusReportAsReport(v *RoamingStatusReport) Report {
	return Report{
		RoamingStatusReport: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Report) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChangeOfSupiPeiAssociationReport
	err = newStrictDecoder(data).Decode(&dst.ChangeOfSupiPeiAssociationReport)
	if err == nil {
		jsonChangeOfSupiPeiAssociationReport, _ := json.Marshal(dst.ChangeOfSupiPeiAssociationReport)
		if string(jsonChangeOfSupiPeiAssociationReport) == "{}" { // empty struct
			dst.ChangeOfSupiPeiAssociationReport = nil
		} else {
			match++
		}
	} else {
		dst.ChangeOfSupiPeiAssociationReport = nil
	}

	// try to unmarshal data into CmInfoReport
	err = newStrictDecoder(data).Decode(&dst.CmInfoReport)
	if err == nil {
		jsonCmInfoReport, _ := json.Marshal(dst.CmInfoReport)
		if string(jsonCmInfoReport) == "{}" { // empty struct
			dst.CmInfoReport = nil
		} else {
			match++
		}
	} else {
		dst.CmInfoReport = nil
	}

	// try to unmarshal data into CnTypeChangeReport
	err = newStrictDecoder(data).Decode(&dst.CnTypeChangeReport)
	if err == nil {
		jsonCnTypeChangeReport, _ := json.Marshal(dst.CnTypeChangeReport)
		if string(jsonCnTypeChangeReport) == "{}" { // empty struct
			dst.CnTypeChangeReport = nil
		} else {
			match++
		}
	} else {
		dst.CnTypeChangeReport = nil
	}

	// try to unmarshal data into LocationReport
	err = newStrictDecoder(data).Decode(&dst.LocationReport)
	if err == nil {
		jsonLocationReport, _ := json.Marshal(dst.LocationReport)
		if string(jsonLocationReport) == "{}" { // empty struct
			dst.LocationReport = nil
		} else {
			match++
		}
	} else {
		dst.LocationReport = nil
	}

	// try to unmarshal data into LossConnectivityReport
	err = newStrictDecoder(data).Decode(&dst.LossConnectivityReport)
	if err == nil {
		jsonLossConnectivityReport, _ := json.Marshal(dst.LossConnectivityReport)
		if string(jsonLossConnectivityReport) == "{}" { // empty struct
			dst.LossConnectivityReport = nil
		} else {
			match++
		}
	} else {
		dst.LossConnectivityReport = nil
	}

	// try to unmarshal data into PdnConnectivityStatReport
	err = newStrictDecoder(data).Decode(&dst.PdnConnectivityStatReport)
	if err == nil {
		jsonPdnConnectivityStatReport, _ := json.Marshal(dst.PdnConnectivityStatReport)
		if string(jsonPdnConnectivityStatReport) == "{}" { // empty struct
			dst.PdnConnectivityStatReport = nil
		} else {
			match++
		}
	} else {
		dst.PdnConnectivityStatReport = nil
	}

	// try to unmarshal data into RoamingStatusReport
	err = newStrictDecoder(data).Decode(&dst.RoamingStatusReport)
	if err == nil {
		jsonRoamingStatusReport, _ := json.Marshal(dst.RoamingStatusReport)
		if string(jsonRoamingStatusReport) == "{}" { // empty struct
			dst.RoamingStatusReport = nil
		} else {
			match++
		}
	} else {
		dst.RoamingStatusReport = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ChangeOfSupiPeiAssociationReport = nil
		dst.CmInfoReport = nil
		dst.CnTypeChangeReport = nil
		dst.LocationReport = nil
		dst.LossConnectivityReport = nil
		dst.PdnConnectivityStatReport = nil
		dst.RoamingStatusReport = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Report)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Report)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Report) MarshalJSON() ([]byte, error) {
	if src.ChangeOfSupiPeiAssociationReport != nil {
		return json.Marshal(&src.ChangeOfSupiPeiAssociationReport)
	}

	if src.CmInfoReport != nil {
		return json.Marshal(&src.CmInfoReport)
	}

	if src.CnTypeChangeReport != nil {
		return json.Marshal(&src.CnTypeChangeReport)
	}

	if src.LocationReport != nil {
		return json.Marshal(&src.LocationReport)
	}

	if src.LossConnectivityReport != nil {
		return json.Marshal(&src.LossConnectivityReport)
	}

	if src.PdnConnectivityStatReport != nil {
		return json.Marshal(&src.PdnConnectivityStatReport)
	}

	if src.RoamingStatusReport != nil {
		return json.Marshal(&src.RoamingStatusReport)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Report) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ChangeOfSupiPeiAssociationReport != nil {
		return obj.ChangeOfSupiPeiAssociationReport
	}

	if obj.CmInfoReport != nil {
		return obj.CmInfoReport
	}

	if obj.CnTypeChangeReport != nil {
		return obj.CnTypeChangeReport
	}

	if obj.LocationReport != nil {
		return obj.LocationReport
	}

	if obj.LossConnectivityReport != nil {
		return obj.LossConnectivityReport
	}

	if obj.PdnConnectivityStatReport != nil {
		return obj.PdnConnectivityStatReport
	}

	if obj.RoamingStatusReport != nil {
		return obj.RoamingStatusReport
	}

	// all schemas are nil
	return nil
}

type NullableReport struct {
	value *Report
	isSet bool
}

func (v NullableReport) Get() *Report {
	return v.value
}

func (v *NullableReport) Set(val *Report) {
	v.value = val
	v.isSet = true
}

func (v NullableReport) IsSet() bool {
	return v.isSet
}

func (v *NullableReport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReport(val *Report) *NullableReport {
	return &NullableReport{value: val, isSet: true}
}

func (v NullableReport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
