/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// ServiceNameAnyOf the model 'ServiceNameAnyOf'
type ServiceNameAnyOf string

// List of ServiceName_anyOf
const (
	NNRF_NFM ServiceNameAnyOf = "nnrf-nfm"
	NNRF_DISC ServiceNameAnyOf = "nnrf-disc"
	NNRF_OAUTH2 ServiceNameAnyOf = "nnrf-oauth2"
	NUDM_SDM ServiceNameAnyOf = "nudm-sdm"
	NUDM_UECM ServiceNameAnyOf = "nudm-uecm"
	NUDM_UEAU ServiceNameAnyOf = "nudm-ueau"
	NUDM_EE ServiceNameAnyOf = "nudm-ee"
	NUDM_PP ServiceNameAnyOf = "nudm-pp"
	NUDM_NIDDAU ServiceNameAnyOf = "nudm-niddau"
	NUDM_MT ServiceNameAnyOf = "nudm-mt"
	NUDM_SSAU ServiceNameAnyOf = "nudm-ssau"
	NUDM_RSDS ServiceNameAnyOf = "nudm-rsds"
	NAMF_COMM ServiceNameAnyOf = "namf-comm"
	NAMF_EVTS ServiceNameAnyOf = "namf-evts"
	NAMF_MT ServiceNameAnyOf = "namf-mt"
	NAMF_LOC ServiceNameAnyOf = "namf-loc"
	NAMF_MBS_COMM ServiceNameAnyOf = "namf-mbs-comm"
	NAMF_MBS_BC ServiceNameAnyOf = "namf-mbs-bc"
	NSMF_PDUSESSION ServiceNameAnyOf = "nsmf-pdusession"
	NSMF_EVENT_EXPOSURE ServiceNameAnyOf = "nsmf-event-exposure"
	NSMF_NIDD ServiceNameAnyOf = "nsmf-nidd"
	NAUSF_AUTH ServiceNameAnyOf = "nausf-auth"
	NAUSF_SORPROTECTION ServiceNameAnyOf = "nausf-sorprotection"
	NAUSF_UPUPROTECTION ServiceNameAnyOf = "nausf-upuprotection"
	NNEF_PFDMANAGEMENT ServiceNameAnyOf = "nnef-pfdmanagement"
	NNEF_SMCONTEXT ServiceNameAnyOf = "nnef-smcontext"
	NNEF_EVENTEXPOSURE ServiceNameAnyOf = "nnef-eventexposure"
	NNEF_EAS_DEPLOYMENT_INFO ServiceNameAnyOf = "nnef-eas-deployment-info"
	_3GPP_CP_PARAMETER_PROVISIONING ServiceNameAnyOf = "3gpp-cp-parameter-provisioning"
	_3GPP_DEVICE_TRIGGERING ServiceNameAnyOf = "3gpp-device-triggering"
	_3GPP_BDT ServiceNameAnyOf = "3gpp-bdt"
	_3GPP_TRAFFIC_INFLUENCE ServiceNameAnyOf = "3gpp-traffic-influence"
	_3GPP_CHARGEABLE_PARTY ServiceNameAnyOf = "3gpp-chargeable-party"
	_3GPP_AS_SESSION_WITH_QOS ServiceNameAnyOf = "3gpp-as-session-with-qos"
	_3GPP_MSISDN_LESS_MO_SMS ServiceNameAnyOf = "3gpp-msisdn-less-mo-sms"
	_3GPP_SERVICE_PARAMETER ServiceNameAnyOf = "3gpp-service-parameter"
	_3GPP_MONITORING_EVENT ServiceNameAnyOf = "3gpp-monitoring-event"
	_3GPP_NIDD_CONFIGURATION_TRIGGER ServiceNameAnyOf = "3gpp-nidd-configuration-trigger"
	_3GPP_NIDD ServiceNameAnyOf = "3gpp-nidd"
	_3GPP_ANALYTICSEXPOSURE ServiceNameAnyOf = "3gpp-analyticsexposure"
	_3GPP_RACS_PARAMETER_PROVISIONING ServiceNameAnyOf = "3gpp-racs-parameter-provisioning"
	_3GPP_ECR_CONTROL ServiceNameAnyOf = "3gpp-ecr-control"
	_3GPP_APPLYING_BDT_POLICY ServiceNameAnyOf = "3gpp-applying-bdt-policy"
	_3GPP_MO_LCS_NOTIFY ServiceNameAnyOf = "3gpp-mo-lcs-notify"
	_3GPP_TIME_SYNC ServiceNameAnyOf = "3gpp-time-sync"
	_3GPP_AM_INFLUENCE ServiceNameAnyOf = "3gpp-am-influence"
	_3GPP_AM_POLICYAUTHORIZATION ServiceNameAnyOf = "3gpp-am-policyauthorization"
	_3GPP_AKMA ServiceNameAnyOf = "3gpp-akma"
	_3GPP_EAS_DEPLOYMENT ServiceNameAnyOf = "3gpp-eas-deployment"
	_3GPP_IPTVCONFIGURATION ServiceNameAnyOf = "3gpp-iptvconfiguration"
	_3GPP_MBS_TMGI ServiceNameAnyOf = "3gpp-mbs-tmgi"
	_3GPP_MBS_SESSION ServiceNameAnyOf = "3gpp-mbs-session"
	_3GPP_AUTHENTICATION ServiceNameAnyOf = "3gpp-authentication"
	_3GPP_ASTI ServiceNameAnyOf = "3gpp-asti"
	NPCF_AM_POLICY_CONTROL ServiceNameAnyOf = "npcf-am-policy-control"
	NPCF_SMPOLICYCONTROL ServiceNameAnyOf = "npcf-smpolicycontrol"
	NPCF_POLICYAUTHORIZATION ServiceNameAnyOf = "npcf-policyauthorization"
	NPCF_BDTPOLICYCONTROL ServiceNameAnyOf = "npcf-bdtpolicycontrol"
	NPCF_EVENTEXPOSURE ServiceNameAnyOf = "npcf-eventexposure"
	NPCF_UE_POLICY_CONTROL ServiceNameAnyOf = "npcf-ue-policy-control"
	NPCF_AM_POLICYAUTHORIZATION ServiceNameAnyOf = "npcf-am-policyauthorization"
	NSMSF_SMS ServiceNameAnyOf = "nsmsf-sms"
	NNSSF_NSSELECTION ServiceNameAnyOf = "nnssf-nsselection"
	NNSSF_NSSAIAVAILABILITY ServiceNameAnyOf = "nnssf-nssaiavailability"
	NUDR_DR ServiceNameAnyOf = "nudr-dr"
	NUDR_GROUP_ID_MAP ServiceNameAnyOf = "nudr-group-id-map"
	NLMF_LOC ServiceNameAnyOf = "nlmf-loc"
	N5G_EIR_EIC ServiceNameAnyOf = "n5g-eir-eic"
	NBSF_MANAGEMENT ServiceNameAnyOf = "nbsf-management"
	NCHF_SPENDINGLIMITCONTROL ServiceNameAnyOf = "nchf-spendinglimitcontrol"
	NCHF_CONVERGEDCHARGING ServiceNameAnyOf = "nchf-convergedcharging"
	NCHF_OFFLINEONLYCHARGING ServiceNameAnyOf = "nchf-offlineonlycharging"
	NNWDAF_EVENTSSUBSCRIPTION ServiceNameAnyOf = "nnwdaf-eventssubscription"
	NNWDAF_ANALYTICSINFO ServiceNameAnyOf = "nnwdaf-analyticsinfo"
	NNWDAF_DATAMANAGEMENT ServiceNameAnyOf = "nnwdaf-datamanagement"
	NNWDAF_MLMODELPROVISION ServiceNameAnyOf = "nnwdaf-mlmodelprovision"
	NGMLC_LOC ServiceNameAnyOf = "ngmlc-loc"
	NUCMF_PROVISIONING ServiceNameAnyOf = "nucmf-provisioning"
	NUCMF_UECAPABILITYMANAGEMENT ServiceNameAnyOf = "nucmf-uecapabilitymanagement"
	NHSS_SDM ServiceNameAnyOf = "nhss-sdm"
	NHSS_UECM ServiceNameAnyOf = "nhss-uecm"
	NHSS_UEAU ServiceNameAnyOf = "nhss-ueau"
	NHSS_EE ServiceNameAnyOf = "nhss-ee"
	NHSS_IMS_SDM ServiceNameAnyOf = "nhss-ims-sdm"
	NHSS_IMS_UECM ServiceNameAnyOf = "nhss-ims-uecm"
	NHSS_IMS_UEAU ServiceNameAnyOf = "nhss-ims-ueau"
	NHSS_GBA_SDM ServiceNameAnyOf = "nhss-gba-sdm"
	NHSS_GBA_UEAU ServiceNameAnyOf = "nhss-gba-ueau"
	NSEPP_TELESCOPIC ServiceNameAnyOf = "nsepp-telescopic"
	NSORAF_SOR ServiceNameAnyOf = "nsoraf-sor"
	NSPAF_SECURED_PACKET ServiceNameAnyOf = "nspaf-secured-packet"
	NUDSF_DR ServiceNameAnyOf = "nudsf-dr"
	NUDSF_TIMER ServiceNameAnyOf = "nudsf-timer"
	NNSSAAF_NSSAA ServiceNameAnyOf = "nnssaaf-nssaa"
	NNSSAAF_AIW ServiceNameAnyOf = "nnssaaf-aiw"
	NAANF_AKMA ServiceNameAnyOf = "naanf-akma"
	N5GDDNMF_DISCOVERY ServiceNameAnyOf = "n5gddnmf-discovery"
	NMFAF_3DADM ServiceNameAnyOf = "nmfaf-3dadm"
	NMFAF_3CADM ServiceNameAnyOf = "nmfaf-3cadm"
	NEASDF_DNSCONTEXT ServiceNameAnyOf = "neasdf-dnscontext"
	NEASDF_BASELINEDNSPATTERN ServiceNameAnyOf = "neasdf-baselinednspattern"
	NDCCF_DM ServiceNameAnyOf = "ndccf-dm"
	NDCCF_CM ServiceNameAnyOf = "ndccf-cm"
	NNSACF_NSAC ServiceNameAnyOf = "nnsacf-nsac"
	NNSACF_SLICE_EE ServiceNameAnyOf = "nnsacf-slice-ee"
	NMBSMF_TMGI ServiceNameAnyOf = "nmbsmf-tmgi"
	NMBSMF_MBSSESSION ServiceNameAnyOf = "nmbsmf-mbssession"
	NADRF_DM ServiceNameAnyOf = "nadrf-dm"
	NBSP_GBA ServiceNameAnyOf = "nbsp-gba"
	NTSCTSF_TIME_SYNC ServiceNameAnyOf = "ntsctsf-time-sync"
	NTSCTSF_QOS_TSCAI ServiceNameAnyOf = "ntsctsf-qos-tscai"
	NTSCTSF_ASTI ServiceNameAnyOf = "ntsctsf-asti"
	NPKMF_KEYREQ ServiceNameAnyOf = "npkmf-keyreq"
	NMNPF_NPSTATUS ServiceNameAnyOf = "nmnpf-npstatus"
	NIWMSC_SMSERVICE ServiceNameAnyOf = "niwmsc-smservice"
	NMBSF_MBSUSERSERV ServiceNameAnyOf = "nmbsf-mbsuserserv"
	NMBSF_MBSUSERDATAING ServiceNameAnyOf = "nmbsf-mbsuserdataing"
	NMBSTF_DISTSESSION ServiceNameAnyOf = "nmbstf-distsession"
	NPANF_PROSEKEY ServiceNameAnyOf = "npanf-prosekey"
)

// All allowed values of ServiceNameAnyOf enum
var AllowedServiceNameAnyOfEnumValues = []ServiceNameAnyOf{
	"nnrf-nfm",
	"nnrf-disc",
	"nnrf-oauth2",
	"nudm-sdm",
	"nudm-uecm",
	"nudm-ueau",
	"nudm-ee",
	"nudm-pp",
	"nudm-niddau",
	"nudm-mt",
	"nudm-ssau",
	"nudm-rsds",
	"namf-comm",
	"namf-evts",
	"namf-mt",
	"namf-loc",
	"namf-mbs-comm",
	"namf-mbs-bc",
	"nsmf-pdusession",
	"nsmf-event-exposure",
	"nsmf-nidd",
	"nausf-auth",
	"nausf-sorprotection",
	"nausf-upuprotection",
	"nnef-pfdmanagement",
	"nnef-smcontext",
	"nnef-eventexposure",
	"nnef-eas-deployment-info",
	"3gpp-cp-parameter-provisioning",
	"3gpp-device-triggering",
	"3gpp-bdt",
	"3gpp-traffic-influence",
	"3gpp-chargeable-party",
	"3gpp-as-session-with-qos",
	"3gpp-msisdn-less-mo-sms",
	"3gpp-service-parameter",
	"3gpp-monitoring-event",
	"3gpp-nidd-configuration-trigger",
	"3gpp-nidd",
	"3gpp-analyticsexposure",
	"3gpp-racs-parameter-provisioning",
	"3gpp-ecr-control",
	"3gpp-applying-bdt-policy",
	"3gpp-mo-lcs-notify",
	"3gpp-time-sync",
	"3gpp-am-influence",
	"3gpp-am-policyauthorization",
	"3gpp-akma",
	"3gpp-eas-deployment",
	"3gpp-iptvconfiguration",
	"3gpp-mbs-tmgi",
	"3gpp-mbs-session",
	"3gpp-authentication",
	"3gpp-asti",
	"npcf-am-policy-control",
	"npcf-smpolicycontrol",
	"npcf-policyauthorization",
	"npcf-bdtpolicycontrol",
	"npcf-eventexposure",
	"npcf-ue-policy-control",
	"npcf-am-policyauthorization",
	"nsmsf-sms",
	"nnssf-nsselection",
	"nnssf-nssaiavailability",
	"nudr-dr",
	"nudr-group-id-map",
	"nlmf-loc",
	"n5g-eir-eic",
	"nbsf-management",
	"nchf-spendinglimitcontrol",
	"nchf-convergedcharging",
	"nchf-offlineonlycharging",
	"nnwdaf-eventssubscription",
	"nnwdaf-analyticsinfo",
	"nnwdaf-datamanagement",
	"nnwdaf-mlmodelprovision",
	"ngmlc-loc",
	"nucmf-provisioning",
	"nucmf-uecapabilitymanagement",
	"nhss-sdm",
	"nhss-uecm",
	"nhss-ueau",
	"nhss-ee",
	"nhss-ims-sdm",
	"nhss-ims-uecm",
	"nhss-ims-ueau",
	"nhss-gba-sdm",
	"nhss-gba-ueau",
	"nsepp-telescopic",
	"nsoraf-sor",
	"nspaf-secured-packet",
	"nudsf-dr",
	"nudsf-timer",
	"nnssaaf-nssaa",
	"nnssaaf-aiw",
	"naanf-akma",
	"n5gddnmf-discovery",
	"nmfaf-3dadm",
	"nmfaf-3cadm",
	"neasdf-dnscontext",
	"neasdf-baselinednspattern",
	"ndccf-dm",
	"ndccf-cm",
	"nnsacf-nsac",
	"nnsacf-slice-ee",
	"nmbsmf-tmgi",
	"nmbsmf-mbssession",
	"nadrf-dm",
	"nbsp-gba",
	"ntsctsf-time-sync",
	"ntsctsf-qos-tscai",
	"ntsctsf-asti",
	"npkmf-keyreq",
	"nmnpf-npstatus",
	"niwmsc-smservice",
	"nmbsf-mbsuserserv",
	"nmbsf-mbsuserdataing",
	"nmbstf-distsession",
	"npanf-prosekey",
}

func (v *ServiceNameAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceNameAnyOf(value)
	for _, existing := range AllowedServiceNameAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceNameAnyOf", value)
}

// NewServiceNameAnyOfFromValue returns a pointer to a valid ServiceNameAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceNameAnyOfFromValue(v string) (*ServiceNameAnyOf, error) {
	ev := ServiceNameAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceNameAnyOf: valid values are %v", v, AllowedServiceNameAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceNameAnyOf) IsValid() bool {
	for _, existing := range AllowedServiceNameAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceName_anyOf value
func (v ServiceNameAnyOf) Ptr() *ServiceNameAnyOf {
	return &v
}

type NullableServiceNameAnyOf struct {
	value *ServiceNameAnyOf
	isSet bool
}

func (v NullableServiceNameAnyOf) Get() *ServiceNameAnyOf {
	return v.value
}

func (v *NullableServiceNameAnyOf) Set(val *ServiceNameAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNameAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNameAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNameAnyOf(val *ServiceNameAnyOf) *NullableServiceNameAnyOf {
	return &NullableServiceNameAnyOf{value: val, isSet: true}
}

func (v NullableServiceNameAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNameAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

