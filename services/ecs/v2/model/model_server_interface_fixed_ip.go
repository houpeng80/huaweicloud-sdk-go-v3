package model

import (
	"encoding/json"

	"strings"
)

type ServerInterfaceFixedIp struct {
	// 网卡私网IP信息。
	IpAddress *string `json:"ip_address,omitempty"`
	// 网卡私网IP对应子网信息。
	SubnetId *string `json:"subnet_id,omitempty"`
}

func (o ServerInterfaceFixedIp) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServerInterfaceFixedIp struct{}"
	}

	return strings.Join([]string{"ServerInterfaceFixedIp", string(data)}, " ")
}
