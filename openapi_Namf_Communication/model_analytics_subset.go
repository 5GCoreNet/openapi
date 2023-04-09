/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// AnalyticsSubset Possible values are: - NUM_OF_UE_REG: The number of UE registered. This value is only applicable to NSI_LOAD_LEVEL event. - NUM_OF_PDU_SESS_ESTBL: The number of PDU sessions established. This value is only applicable to NSI_LOAD_LEVEL event. - RES_USAGE: The current usage of the virtual resources assigned to the NF instances belonging to a particular network slice instance. This value is only applicable to NSI_LOAD_LEVEL event. - NUM_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR: The number of times the resource usage threshold of the network slice instance is reached or exceeded if a threshold value is provided by the consumer. This value is only applicable to NSI_LOAD_LEVEL event. - PERIOD_OF_EXCEED_RES_USAGE_LOAD_LEVEL_THR: The time interval between each time the threshold being met or exceeded on the network slice (instance). This value is only applicable to NSI_LOAD_LEVEL event. - EXCEED_LOAD_LEVEL_THR_IND: Whether the Load Level Threshold is met or exceeded by the statistics value. This value is only applicable to NSI_LOAD_LEVEL event. - LIST_OF_TOP_APP_UL: The list of applications that contribute the most to the traffic in the UL direction. This value is only applicable to USER_DATA_CONGESTION event. - LIST_OF_TOP_APP_DL: The list of applications that contribute the most to the traffic in the DL direction. This value is only applicable to USER_DATA_CONGESTION event. - NF_STATUS: The availability status of the NF on the Analytics target period, expressed as a percentage of time per status value (registered, suspended, undiscoverable). This value is only applicable to NF_LOAD event. - NF_RESOURCE_USAGE: The average usage of assigned resources (CPU, memory, storage). This value is only applicable to NF_LOAD event. - NF_LOAD: The average load of the NF instance over the Analytics target period. This value is only applicable to NF_LOAD event. - NF_PEAK_LOAD: The maximum load of the NF instance over the Analytics target period. This value is only applicable to NF_LOAD event. - NF_LOAD_AVG_IN_AOI: The average load of the NF instances over the area of interest. This value is only applicable to NF_LOAD event. - DISPER_AMOUNT: Indicates the dispersion amount of the reported data volume or transaction dispersion type. This value is only applicable to DISPERSION event. - DISPER_CLASS: Indicates the dispersion mobility class: fixed, camper, traveller upon set its usage threshold, and/or the top-heavy class upon set its percentile rating threshold. This value is only applicable to DISPERSION event. - RANKING: Data/transaction usage ranking high (i.e.value 1), medium (2) or low (3). This value is only applicable to DISPERSION event. - PERCENTILE_RANKING: Percentile ranking of the target UE in the Cumulative Distribution Function of data usage for the population of all UEs. This value is only applicable to DISPERSION event. - RSSI: Indicated the RSSI in the unit of dBm. This value is only applicable to WLAN_PERFORMANCE event. - RTT: Indicates the RTT in the unit of millisecond. This value is only applicable to WLAN_PERFORMANCE event. - TRAFFIC_INFO: Traffic information including UL/DL data rate and/or Traffic volume. This value is only applicable to WLAN_PERFORMANCE event. - NUMBER_OF_UES: Number of UEs observed for the SSID. This value is only applicable to WLAN_PERFORMANCE event. - APP_LIST_FOR_UE_COMM: The analytics of the application list used by UE. This value is only applicable to UE_COMM event. - N4_SESS_INACT_TIMER_FOR_UE_COMM: The N4 Session inactivity timer. This value is only applicable to UE_COMM event. - AVG_TRAFFIC_RATE: Indicates average traffic rate. This value is only applicable to DN_PERFORMANCE event. - MAX_TRAFFIC_RATE: Indicates maximum traffic rate. This value is only applicable to DN_PERFORMANCE event. - AVG_PACKET_DELAY: Indicates average Packet Delay. This value is only applicable to DN_PERFORMANCE event. - MAX_PACKET_DELAY: Indicates maximum Packet Delay. This value is only applicable to DN_PERFORMANCE event. - AVG_PACKET_LOSS_RATE: Indicates average Loss Rate. This value is only applicable to DN_PERFORMANCE event. - UE_LOCATION: Indicates UE location information. This value is only applicable to SERVICE_EXPERIENCE event. - LIST_OF_HIGH_EXP_UE: Indicates list of high experienced UE. This value is only applicable to SM_CONGESTION event. - LIST_OF_MEDIUM_EXP_UE: Indicates list of medium experienced UE. This value is only applicable to SM_CONGESTION event. - LIST_OF_LOW_EXP_UE: Indicates list of low experienced UE. This value is only applicable to SM_CONGESTION event. - AVG_UL_PKT_DROP_RATE: Indicates average uplink packet drop rate on GTP-U path on N3. This value is only applicable to RED_TRANS_EXP event. - VAR_UL_PKT_DROP_RATE: Indicates variance of uplink packet drop rate on GTP-U path on N3. This value is only applicable to RED_TRANS_EXP event. - AVG_DL_PKT_DROP_RATE: Indicates average downlink packet drop rate on GTP-U path on N3. This value is only applicable to RED_TRANS_EXP event. - VAR_DL_PKT_DROP_RATE: Indicates variance of downlink packet drop rate on GTP-U path on N3. This value is only applicable to RED_TRANS_EXP event. - AVG_UL_PKT_DELAY: Indicates average uplink packet delay round trip on GTP-U path on N3. This value is only applicable to RED_TRANS_EXP event. - VAR_UL_PKT_DELAY: Indicates variance uplink packet delay round trip on GTP-U path on N3. This value is only applicable to RED_TRANS_EXP event. - AVG_DL_PKT_DELAY: Indicates average downlink packet delay round trip on GTP-U path on N3. This value is only applicable to RED_TRANS_EXP event. - VAR_DL_PKT_DELAY: Indicates variance downlink packet delay round trip on GTP-U path on N3. This value is only applicable to RED_TRANS_EXP event.
type AnalyticsSubset struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnalyticsSubset) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AnalyticsSubset)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnalyticsSubset) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAnalyticsSubset struct {
	value *AnalyticsSubset
	isSet bool
}

func (v NullableAnalyticsSubset) Get() *AnalyticsSubset {
	return v.value
}

func (v *NullableAnalyticsSubset) Set(val *AnalyticsSubset) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsSubset) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsSubset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsSubset(val *AnalyticsSubset) *NullableAnalyticsSubset {
	return &NullableAnalyticsSubset{value: val, isSet: true}
}

func (v NullableAnalyticsSubset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsSubset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
