package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TextToImageInference struct {

	// 文生图引导词
	Prompt string `json:"prompt"`

	// 随机数种子
	Seed *int32 `json:"seed,omitempty"`

	// 生成图片分辨率，限定字符串\"512*768\",\"768*512\",\"512*512\",\"1024*768\",\"768*1024\"，默认\"512*512\"
	Resolution *string `json:"resolution,omitempty"`
}

func (o TextToImageInference) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TextToImageInference struct{}"
	}

	return strings.Join([]string{"TextToImageInference", string(data)}, " ")
}
