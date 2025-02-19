package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAllTablesResponse Response Object
type ListAllTablesResponse struct {

	// 执行请求是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message *string `json:"message,omitempty"`

	// 表的总个数。
	TableCount *int32 `json:"table_count,omitempty"`

	// 表的信息。
	Tables         *[]Table `json:"tables,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListAllTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAllTablesResponse struct{}"
	}

	return strings.Join([]string{"ListAllTablesResponse", string(data)}, " ")
}
