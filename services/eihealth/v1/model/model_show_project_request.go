package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowProjectRequest Request Object
type ShowProjectRequest struct {

	// X-Bucket-Name
	XBucketName *string `json:"X-Bucket-Name,omitempty"`

	// X-Namespace
	XNamespaceName *string `json:"X-Namespace-Name,omitempty"`

	// 医疗智能体平台项目ID，您可以在EIHealth平台单击所需的项目名称，进入项目设置页面查看。
	EihealthProjectId string `json:"eihealth_project_id"`
}

func (o ShowProjectRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowProjectRequest struct{}"
	}

	return strings.Join([]string{"ShowProjectRequest", string(data)}, " ")
}
