package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceRequest Request Object
type ShowInstanceRequest struct {

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`

	// 数据库实例id
	DatabaseId string `json:"database_id"`
}

func (o ShowInstanceRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceRequest", string(data)}, " ")
}
