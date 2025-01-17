/*
SLOs

OpenAPI schema for SLOs endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package slo

import (
	"encoding/json"
	"fmt"
)

// BudgetingMethod The budgeting method to use
type BudgetingMethod string

// List of budgeting_method
const (
	OCCURRENCES BudgetingMethod = "occurrences"
	TIMESLICES  BudgetingMethod = "timeslices"
)

// All allowed values of BudgetingMethod enum
var AllowedBudgetingMethodEnumValues = []BudgetingMethod{
	"occurrences",
	"timeslices",
}

func (v *BudgetingMethod) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BudgetingMethod(value)
	for _, existing := range AllowedBudgetingMethodEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BudgetingMethod", value)
}

// NewBudgetingMethodFromValue returns a pointer to a valid BudgetingMethod
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBudgetingMethodFromValue(v string) (*BudgetingMethod, error) {
	ev := BudgetingMethod(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BudgetingMethod: valid values are %v", v, AllowedBudgetingMethodEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BudgetingMethod) IsValid() bool {
	for _, existing := range AllowedBudgetingMethodEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to budgeting_method value
func (v BudgetingMethod) Ptr() *BudgetingMethod {
	return &v
}

type NullableBudgetingMethod struct {
	value *BudgetingMethod
	isSet bool
}

func (v NullableBudgetingMethod) Get() *BudgetingMethod {
	return v.value
}

func (v *NullableBudgetingMethod) Set(val *BudgetingMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableBudgetingMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableBudgetingMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBudgetingMethod(val *BudgetingMethod) *NullableBudgetingMethod {
	return &NullableBudgetingMethod{value: val, isSet: true}
}

func (v NullableBudgetingMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBudgetingMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
