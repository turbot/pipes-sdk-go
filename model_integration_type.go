/*
Turbot Pipes

Turbot Pipes is an intelligence, automation & security platform built specifically for DevOps.

API version: {{OPEN_API_VERSION}}
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipes

import (
	"encoding/json"
	"fmt"
)

// IntegrationType the model 'IntegrationType'
type IntegrationType string

// List of IntegrationType
const (
	IntegrationTypeAWS    IntegrationType = "aws"
	IntegrationTypeAzure  IntegrationType = "azure"
	IntegrationTypeEmail  IntegrationType = "email"
	IntegrationTypeGCP    IntegrationType = "gcp"
	IntegrationTypeGitHub IntegrationType = "github"
	IntegrationTypeSlack  IntegrationType = "slack"
	IntegrationTypeGitlab IntegrationType = "gitlab"
)

// All allowed values of IntegrationType enum
var AllowedIntegrationTypeEnumValues = []IntegrationType{
	"aws",
	"azure",
	"email",
	"gcp",
	"github",
	"slack",
	"gitlab",
}

func (v *IntegrationType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IntegrationType(value)
	for _, existing := range AllowedIntegrationTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IntegrationType", value)
}

// NewIntegrationTypeFromValue returns a pointer to a valid IntegrationType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIntegrationTypeFromValue(v string) (*IntegrationType, error) {
	ev := IntegrationType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IntegrationType: valid values are %v", v, AllowedIntegrationTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IntegrationType) IsValid() bool {
	for _, existing := range AllowedIntegrationTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IntegrationType value
func (v IntegrationType) Ptr() *IntegrationType {
	return &v
}

type NullableIntegrationType struct {
	value *IntegrationType
	isSet bool
}

func (v NullableIntegrationType) Get() *IntegrationType {
	return v.value
}

func (v *NullableIntegrationType) Set(val *IntegrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationType(val *IntegrationType) *NullableIntegrationType {
	return &NullableIntegrationType{value: val, isSet: true}
}

func (v NullableIntegrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
