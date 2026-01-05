// Copyright 2025 EURECOM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Contributors:
//   Giulio CAROTA
//   Thomas DU
//   Adlen KSENTINI

/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.5
*/

package models

import (
	"encoding/json"
	"fmt"
)

// TransactionMetric Possible values are - PDU_SES_EST: PDU Session Establishment - PDU_SES_AUTH: PDU Session Authentication - PDU_SES_MODIF: PDU Session Modification - PDU_SES_REL: PDU Session Release
type TransactionMetric struct {
	TransactionMetricAnyOf *TransactionMetricAnyOf
	string                 *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransactionMetric) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TransactionMetricAnyOf
	err = json.Unmarshal(data, &dst.TransactionMetricAnyOf)
	if err == nil {
		jsonTransactionMetricAnyOf, _ := json.Marshal(dst.TransactionMetricAnyOf)
		if string(jsonTransactionMetricAnyOf) == "{}" { // empty struct
			dst.TransactionMetricAnyOf = nil
		} else {
			return nil // data stored in dst.TransactionMetricAnyOf, return on the first match
		}
	} else {
		dst.TransactionMetricAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TransactionMetric)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransactionMetric) MarshalJSON() ([]byte, error) {
	if src.TransactionMetricAnyOf != nil {
		return json.Marshal(&src.TransactionMetricAnyOf)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTransactionMetric struct {
	value *TransactionMetric
	isSet bool
}

func (v NullableTransactionMetric) Get() *TransactionMetric {
	return v.value
}

func (v *NullableTransactionMetric) Set(val *TransactionMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMetric(val *TransactionMetric) *NullableTransactionMetric {
	return &NullableTransactionMetric{value: val, isSet: true}
}

func (v NullableTransactionMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
