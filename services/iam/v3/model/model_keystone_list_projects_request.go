package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type KeystoneListProjectsRequest struct {
	DomainId *string `json:"domain_id,omitempty"`
	Name     *string `json:"name,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
	Enabled  *bool   `json:"enabled,omitempty"`
	IsDomain *bool   `json:"is_domain,omitempty"`
	Page     *int32  `json:"page,omitempty"`
	PerPage  *int32  `json:"per_page,omitempty"`
}

func (o KeystoneListProjectsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "KeystoneListProjectsRequest struct{}"
	}

	return strings.Join([]string{"KeystoneListProjectsRequest", string(data)}, " ")
}
