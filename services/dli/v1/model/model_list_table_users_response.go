package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTableUsersResponse Response Object
type ListTableUsersResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 权限信息。
	Privileges     *[]ShowTableUsersRespPrivilege `json:"privileges,omitempty"`
	HttpStatusCode int                            `json:"-"`
}

func (o ListTableUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTableUsersResponse struct{}"
	}

	return strings.Join([]string{"ListTableUsersResponse", string(data)}, " ")
}
