/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 创建裸金属服务器接口请求结构体
type CreateBaremetalServersBody struct {
	Server *CreateServers `json:"server"`
}

func (o CreateBaremetalServersBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateBaremetalServersBody", string(data)}, " ")
}
