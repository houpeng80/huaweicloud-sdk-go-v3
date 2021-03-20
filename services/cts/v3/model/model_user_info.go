package model

import (
	"encoding/json"

	"strings"
)

// 用户信息。
type UserInfo struct {
	// 账号ID，参见《云审计服务API参考》“获取账号ID和项目ID”章节。

	Id *string `json:"id,omitempty"`
	// 账号名称。

	Name *string `json:"name,omitempty"`

	Domain *BaseUser `json:"domain,omitempty"`
}

func (o UserInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UserInfo struct{}"
	}

	return strings.Join([]string{"UserInfo", string(data)}, " ")
}
