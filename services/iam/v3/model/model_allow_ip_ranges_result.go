package model

import (
	"encoding/json"

	"strings"
)

//
type AllowIpRangesResult struct {
	// 描述信息。
	Description string `json:"description"`
	// IP地址区间，例如：0.0.0.0-255.255.255.255。
	IpRange string `json:"ip_range"`
}

func (o AllowIpRangesResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AllowIpRangesResult struct{}"
	}

	return strings.Join([]string{"AllowIpRangesResult", string(data)}, " ")
}
