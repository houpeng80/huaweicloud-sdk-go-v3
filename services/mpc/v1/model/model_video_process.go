package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type VideoProcess struct {
	// 视频顺时针旋转角度。  - 0：表示不旋转 - 1：表示顺时针旋转90度 - 2：表示顺时针旋转180度 - 3：表示顺时针旋转270度
	Rotate       *int32        `json:"rotate,omitempty"`
	BlackEnhance *BlackEnhance `json:"black_enhance,omitempty"`
	// 长短边自适应控制字段： - SHORT：表示短边自适应 - LONG：表示长边自适应 - NONE：表示不自适应
	Adaptation *VideoProcessAdaptation `json:"adaptation,omitempty"`
	// 是否开启上采样，如支持从480P的片源转为720P，可取值为:  - 0：表示上采样关闭， - 1：表示上采样开启.
	Upsample *int32 `json:"upsample,omitempty"`
}

func (o VideoProcess) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "VideoProcess struct{}"
	}

	return strings.Join([]string{"VideoProcess", string(data)}, " ")
}

type VideoProcessAdaptation struct {
	value string
}

type VideoProcessAdaptationEnum struct {
	SHORT VideoProcessAdaptation
	LONG  VideoProcessAdaptation
	NONE  VideoProcessAdaptation
}

func GetVideoProcessAdaptationEnum() VideoProcessAdaptationEnum {
	return VideoProcessAdaptationEnum{
		SHORT: VideoProcessAdaptation{
			value: "SHORT",
		},
		LONG: VideoProcessAdaptation{
			value: "LONG",
		},
		NONE: VideoProcessAdaptation{
			value: "NONE",
		},
	}
}

func (c VideoProcessAdaptation) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *VideoProcessAdaptation) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
