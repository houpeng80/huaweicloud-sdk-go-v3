package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 转发DIS服务消息结构
type ActionDisForwarding struct {
	// **参数说明**：DIS服务对应的region区域。

	RegionName string `json:"region_name"`
	// **参数说明**：DIS服务对应的projectId信息。

	ProjectId string `json:"project_id"`
	// **参数说明**：DIS服务对应的通道名称，和通道ID参数中至少一个不为空，和通道ID参数都存在时，以通道ID参数值为准。通过调用DIS服务 [查询通道列表](https://support.huaweicloud.com/api-dis/dis_02_0024.html)接口获取。

	StreamName *string `json:"stream_name,omitempty"`
	// **参数说明**：DIS服务对应的通道ID，和通道名称参数中至少一个不为空，和通道名称参数都存在时，以本参数值为准。通过调用DIS服务 [查询通道详情](https://support.huaweicloud.com/api-dis/dis_02_0025.html)接口获取。

	StreamId *string `json:"stream_id,omitempty"`
}

func (o ActionDisForwarding) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionDisForwarding struct{}"
	}

	return strings.Join([]string{"ActionDisForwarding", string(data)}, " ")
}
