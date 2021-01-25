package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListMembersRequest struct {
	PoolId              string    `json:"pool_id"`
	Address             *[]string `json:"address,omitempty"`
	AdminStateUp        *bool     `json:"admin_state_up,omitempty"`
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	Id                  *[]string `json:"id,omitempty"`
	Limit               *int32    `json:"limit,omitempty"`
	Marker              *string   `json:"marker,omitempty"`
	Name                *[]string `json:"name,omitempty"`
	OperatingStatus     *[]string `json:"operating_status,omitempty"`
	PageReverse         *bool     `json:"page_reverse,omitempty"`
	ProtocolPort        *[]int32  `json:"protocol_port,omitempty"`
	SubnetCidrId        *[]string `json:"subnet_cidr_id,omitempty"`
	Weight              *[]int32  `json:"weight,omitempty"`
}

func (o ListMembersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMembersRequest struct{}"
	}

	return strings.Join([]string{"ListMembersRequest", string(data)}, " ")
}
