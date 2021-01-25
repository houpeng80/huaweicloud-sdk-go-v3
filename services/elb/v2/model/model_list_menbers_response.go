package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListMenbersResponse struct {
	// 后端云服务器对象的列表
	Members        *[]MemberResp `json:"members,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ListMenbersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListMenbersResponse struct{}"
	}

	return strings.Join([]string{"ListMenbersResponse", string(data)}, " ")
}
