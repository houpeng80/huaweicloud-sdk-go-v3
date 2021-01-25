package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListCustomPoliciesResponse struct {
	Links *Links `json:"links,omitempty"`
	// 自定义策略信息列表。
	Roles *[]PolicyRoleResult `json:"roles,omitempty"`
	// 返回自定义策略的总条数
	TotalNumber    *int32 `json:"total_number,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListCustomPoliciesResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListCustomPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListCustomPoliciesResponse", string(data)}, " ")
}
