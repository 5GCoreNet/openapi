/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// Event Possible values are: - USER_DATA_ING_SESS_STARTING: >     Indicates that the MBS User Data Ingest Session is starting. This is an \"MBS User Data     Ingest Session\" level event. - USER_DATA_ING_SESS_STARTED: >     Indicates that the MBS User Data Ingest Session started. This is an \"MBS User Data     Ingest Session\" level event. - USER_DATA_ING_SESS_TERMINATED: >     Indicates that the MBS User Data Ingest Session is terminated. This is an \"MBS User Data     Ingest Session\" level event. - DIST_SESS_STARTING: >     Indicates that the MBS Distribution Session is starting. This is an \"MBS Distribution     Session\" level event. - DIST_SESS_STARTED: >     Indicates that the MBS Distribution Session started. This is an \"MBS Distribution     Session\" level event. - DIST_SESS_TERMINATED: >     Indicates that the MBS Distribution Session is terminated. This is an \"MBS Distribution     Session\" level event. - DIST_SESS_SERV_MNGT_FAILURE: >     Indicates that the MBS Distribution Session could not be started (e.g. the necessary     resources could not be allocated by the MBS system). This is an \"MBS Distribution     Session\" level event. - DIST_SESS_POL_CRTL_FAILURE: >     Indicates that the MBS Distribution Session could not be started because of a policy     authorization/control failure or rejection. This is an \"MBS Distribution Session\"     level event. - DATA_INGEST_FAILURE: >     The MBS User Data Ingest is failed because the MBSTF is expecting data (the MBS Session     is active), but not receiving it. This is an \"MBS Distribution Session\" level event. - DELIVERY_STARTED: >     The MBS User Data delivery is started. - SESSION_TERMINATED: >     The MBS User Data Ingest Session is terminated.
type Event struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Event) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(Event)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Event) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
