package model

import (
	"encoding/json"

	"strings"
)

type ListRetirableGrantsRequestBody struct {
	// 指定查询可退役授权返回记录条数，如果查询记录条数小于存在的条数，响应参数“truncated”将返回“true”，表示存在分页。 取值在授权最大个数范围以内。例如：100
	Limit *string `json:"limit,omitempty"`
	// 分页查询起始位置标识。 分页查询收到的响应参数“truncated”为“true”时，可以发送连续的请求获取更多的记录条数，“marker”设置为响应的“next_marker”的值。例如：10。
	Marker *string `json:"marker,omitempty"`
	// 请求消息序列号，36字节序列号。 例如：919c82d4-8046-4722-9094-35c3c6524cff
	Sequence *string `json:"sequence,omitempty"`
}

func (o ListRetirableGrantsRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListRetirableGrantsRequestBody struct{}"
	}

	return strings.Join([]string{"ListRetirableGrantsRequestBody", string(data)}, " ")
}
