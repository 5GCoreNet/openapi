/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// PolicyControlRequestTrigger Possible values are: - PLMN_CH: PLMN Change - RES_MO_RE: A request for resource modification has been received by the SMF. The SMF  always reports to the PCF. - AC_TY_CH: Access Type Change - UE_IP_CH: UE IP address change. The SMF always reports to the PCF. - UE_MAC_CH: A new UE MAC address is detected or a used UE MAC address is inactive for a  specific period - AN_CH_COR: Access Network Charging Correlation Information - US_RE: The PDU Session or the Monitoring key specific resources consumed by a UE either  reached the threshold or needs to be reported for other reasons. - APP_STA: The start of application traffic has been detected. - APP_STO: The stop of application traffic has been detected. - AN_INFO: Access Network Information report - CM_SES_FAIL: Credit management session failure - PS_DA_OFF: The SMF reports when the 3GPP PS Data Off status changes. The SMF always  reports to the PCF. - DEF_QOS_CH: Default QoS Change. The SMF always reports to the PCF. - SE_AMBR_CH: Session-AMBR Change. The SMF always reports to the PCF. - QOS_NOTIF: The SMF notify the PCF when receiving notification from RAN that QoS targets of  the QoS Flow cannot be guranteed or gurateed again. - NO_CREDIT: Out of credit - REALLO_OF_CREDIT: Reallocation of credit - PRA_CH: Change of UE presence in Presence Reporting Area - SAREA_CH: Location Change with respect to the Serving Area - SCNN_CH: Location Change with respect to the Serving CN node - RE_TIMEOUT: Indicates the SMF generated the request because there has been a PCC  revalidation timeout - RES_RELEASE: Indicate that the SMF can inform the PCF of the outcome of the release of  resources for those rules that require so. - SUCC_RES_ALLO: Indicates that the requested rule data is the successful resource  allocation. - RAI_CH: Location Change with respect to the RAI of GERAN and UTRAN. - RAT_TY_CH: RAT Type Change. - REF_QOS_IND_CH: Reflective QoS indication Change - NUM_OF_PACKET_FILTER: Indicates that the SMF shall report the number of supported packet  filter for signalled QoS rules - UE_STATUS_RESUME: Indicates that the UE's status is resumed. - UE_TZ_CH: UE Time Zone Change - AUTH_PROF_CH: The DN-AAA authorization profile index has changed - QOS_MONITORING: Indicate that the SMF notifies the PCF of the QoS Monitoring information. - SCELL_CH: Location Change with respect to the Serving Cell. - USER_LOCATION_CH: Indicate that user location has been changed, applicable to serving area  change and serving cell change. - EPS_FALLBACK: EPS Fallback report is enabled in the SMF. - MA_PDU: UE Indicates that the SMF notifies the PCF of the MA PDU session request - TSN_BRIDGE_INFO: TSC user plane node information available - 5G_RG_JOIN: The 5G-RG has joined to an IP Multicast Group. - 5G_RG_LEAVE: The 5G-RG has left an IP Multicast Group. - DDN_FAILURE: Event subscription for DDN Failure event received. - DDN_DELIVERY_STATUS: Event subscription for DDN Delivery Status received. - GROUP_ID_LIST_CHG: UE Internal Group Identifier(s) has changed: the SMF reports that UDM  provided list of group Ids has changed. - DDN_FAILURE_CANCELLATION: The event subscription for DDN Failure event is cancelled. - DDN_DELIVERY_STATUS_CANCELLATION: The event subscription for DDD STATUS is cancelled. - VPLMN_QOS_CH: Change of the QoS supported in the VPLMN. - SUCC_QOS_UPDATE: Indicates that the requested MPS Action is successful. - SAT_CATEGORY_CHG: Indicates that the SMF has detected a change between different satellite  backhaul categories, or between a satellite backhaul and a non-satellite backhaul. - PCF_UE_NOTIF_IND: Indicates the SMF has detected the AMF forwarded the PCF for the UE  indication to receive/stop receiving notifications of SM Policy association  established/terminated events. - NWDAF_DATA_CHG: Indicates that the NWDAF instance IDs used for the PDU session and/or  associated Analytics IDs used for the PDU session and available in the SMF have changed.
type PolicyControlRequestTrigger struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PolicyControlRequestTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PolicyControlRequestTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PolicyControlRequestTrigger) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePolicyControlRequestTrigger struct {
	value *PolicyControlRequestTrigger
	isSet bool
}

func (v NullablePolicyControlRequestTrigger) Get() *PolicyControlRequestTrigger {
	return v.value
}

func (v *NullablePolicyControlRequestTrigger) Set(val *PolicyControlRequestTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyControlRequestTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyControlRequestTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyControlRequestTrigger(val *PolicyControlRequestTrigger) *NullablePolicyControlRequestTrigger {
	return &NullablePolicyControlRequestTrigger{value: val, isSet: true}
}

func (v NullablePolicyControlRequestTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyControlRequestTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
