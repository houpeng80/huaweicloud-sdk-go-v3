package model

import (
	"encoding/json"

	"strings"
)

type IefNode struct {
	// 节点ID

	Id string `json:"id"`
	// 节点状态:\"ACTIVE\"

	Status string `json:"status"`
	// 节点公有IP

	PublicIpAddress string `json:"public_ip_address"`
}

func (o IefNode) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "IefNode struct{}"
	}

	return strings.Join([]string{"IefNode", string(data)}, " ")
}
