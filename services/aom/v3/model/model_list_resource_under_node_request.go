package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListResourceUnderNodeRequest struct {

	// 资源种类
	RfResourceType string `json:"rf_resource_type"`

	// 资源种类下的类型区分
	Type string `json:"type"`

	Body *PageResourceListParam `json:"body,omitempty"`
}

func (o ListResourceUnderNodeRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListResourceUnderNodeRequest struct{}"
	}

	return strings.Join([]string{"ListResourceUnderNodeRequest", string(data)}, " ")
}
