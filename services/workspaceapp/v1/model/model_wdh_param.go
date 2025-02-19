package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WdhParam wdh参数
type WdhParam struct {

	// 云办公主机id
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	// 专属主机类型，目前只支持dedicated - dedicated：专属型 - shared： 共享型
	Tenancy *string `json:"tenancy,omitempty"`
}

func (o WdhParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WdhParam struct{}"
	}

	return strings.Join([]string{"WdhParam", string(data)}, " ")
}
