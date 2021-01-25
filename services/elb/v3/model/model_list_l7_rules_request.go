package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListL7RulesRequest struct {
	L7policyId          string    `json:"l7policy_id"`
	AdminStateUp        *bool     `json:"admin_state_up,omitempty"`
	CompareType         *[]string `json:"compare_type,omitempty"`
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	Id                  *[]string `json:"id,omitempty"`
	Invert              *bool     `json:"invert,omitempty"`
	Key                 *[]string `json:"key,omitempty"`
	Limit               *int32    `json:"limit,omitempty"`
	Marker              *string   `json:"marker,omitempty"`
	PageReverse         *bool     `json:"page_reverse,omitempty"`
	ProvisioningStatus  *[]string `json:"provisioning_status,omitempty"`
	Type                *[]string `json:"type,omitempty"`
	Value               *[]string `json:"value,omitempty"`
}

func (o ListL7RulesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListL7RulesRequest struct{}"
	}

	return strings.Join([]string{"ListL7RulesRequest", string(data)}, " ")
}
