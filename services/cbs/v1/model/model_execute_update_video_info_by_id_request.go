package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ExecuteUpdateVideoInfoByIdRequest struct {

	// 视频id
	VideoId string `json:"video_id"`

	Body *PutVideoInfo `json:"body,omitempty"`
}

func (o ExecuteUpdateVideoInfoByIdRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteUpdateVideoInfoByIdRequest struct{}"
	}

	return strings.Join([]string{"ExecuteUpdateVideoInfoByIdRequest", string(data)}, " ")
}
