package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMembersDetailResponse struct {
	// 添加备份共享成员响应信息

	Members *[]Member `json:"members,omitempty"`
	//

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMembersDetailResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMembersDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowMembersDetailResponse", string(data)}, " ")
}
