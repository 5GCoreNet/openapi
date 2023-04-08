/*
Ndcaf_DataReportingProvisioning

Data Collection AF: Provisioning Sessions API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReportingProvisioning

import (
	"encoding/json"
	"fmt"
)

// DataAccessProfileUserAccessRestrictionsUserIdsInner struct for DataAccessProfileUserAccessRestrictionsUserIdsInner
type DataAccessProfileUserAccessRestrictionsUserIdsInner struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DataAccessProfileUserAccessRestrictionsUserIdsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
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

	return fmt.Errorf("data failed to match schemas in anyOf(DataAccessProfileUserAccessRestrictionsUserIdsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DataAccessProfileUserAccessRestrictionsUserIdsInner) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDataAccessProfileUserAccessRestrictionsUserIdsInner struct {
	value *DataAccessProfileUserAccessRestrictionsUserIdsInner
	isSet bool
}

func (v NullableDataAccessProfileUserAccessRestrictionsUserIdsInner) Get() *DataAccessProfileUserAccessRestrictionsUserIdsInner {
	return v.value
}

func (v *NullableDataAccessProfileUserAccessRestrictionsUserIdsInner) Set(val *DataAccessProfileUserAccessRestrictionsUserIdsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAccessProfileUserAccessRestrictionsUserIdsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAccessProfileUserAccessRestrictionsUserIdsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAccessProfileUserAccessRestrictionsUserIdsInner(val *DataAccessProfileUserAccessRestrictionsUserIdsInner) *NullableDataAccessProfileUserAccessRestrictionsUserIdsInner {
	return &NullableDataAccessProfileUserAccessRestrictionsUserIdsInner{value: val, isSet: true}
}

func (v NullableDataAccessProfileUserAccessRestrictionsUserIdsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAccessProfileUserAccessRestrictionsUserIdsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


