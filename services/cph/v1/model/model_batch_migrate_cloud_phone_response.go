package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchMigrateCloudPhoneResponse Response Object
type BatchMigrateCloudPhoneResponse struct {

	// 请求的唯一标识ID。
	RequestId *string `json:"request_id,omitempty"`

	// 任务列表
	Jobs           *[]PhoneMigrateJob `json:"jobs,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o BatchMigrateCloudPhoneResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchMigrateCloudPhoneResponse struct{}"
	}

	return strings.Join([]string{"BatchMigrateCloudPhoneResponse", string(data)}, " ")
}
