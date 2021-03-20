package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListOpLogsResponse struct {
	// 任务列表

	OperationLogs *[]OperationLog `json:"operation_logs,omitempty"`
	// 任务个数

	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListOpLogsResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListOpLogsResponse struct{}"
	}

	return strings.Join([]string{"ListOpLogsResponse", string(data)}, " ")
}
