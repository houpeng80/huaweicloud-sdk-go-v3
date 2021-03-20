package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 转储桶信息。
type TrackerObsInfo struct {
	// 标识OBS桶名称。由数字或字母开头，支持小写字母、数字、“-”、“.”，长度为3～63个字符。

	BucketName *string `json:"bucket_name,omitempty"`
	// 标识需要存储于OBS的日志文件前缀，0-9，a-z，A-Z，'-'，'.'，'_'长度为0～64字符。

	FilePrefixName *string `json:"file_prefix_name,omitempty"`
	// 是否支持新建OBS桶。   值为“true”时，表示新创建OBS桶存储事件文件；   值为“false”时，选择已存在的OBS桶存储事件文件。

	IsObsCreated *bool `json:"is_obs_created,omitempty"`
	// 标识配置桶内对象存储周期。 当\"tracker_type\"参数值为\"data\"时该参数值有效。

	BucketLifecycle *TrackerObsInfoBucketLifecycle `json:"bucket_lifecycle,omitempty"`
}

func (o TrackerObsInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "TrackerObsInfo struct{}"
	}

	return strings.Join([]string{"TrackerObsInfo", string(data)}, " ")
}

type TrackerObsInfoBucketLifecycle struct {
	value int32
}

type TrackerObsInfoBucketLifecycleEnum struct {
	E_30   TrackerObsInfoBucketLifecycle
	E_60   TrackerObsInfoBucketLifecycle
	E_90   TrackerObsInfoBucketLifecycle
	E_180  TrackerObsInfoBucketLifecycle
	E_1095 TrackerObsInfoBucketLifecycle
}

func GetTrackerObsInfoBucketLifecycleEnum() TrackerObsInfoBucketLifecycleEnum {
	return TrackerObsInfoBucketLifecycleEnum{
		E_30: TrackerObsInfoBucketLifecycle{
			value: 30,
		}, E_60: TrackerObsInfoBucketLifecycle{
			value: 60,
		}, E_90: TrackerObsInfoBucketLifecycle{
			value: 90,
		}, E_180: TrackerObsInfoBucketLifecycle{
			value: 180,
		}, E_1095: TrackerObsInfoBucketLifecycle{
			value: 1095,
		},
	}
}

func (c TrackerObsInfoBucketLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *TrackerObsInfoBucketLifecycle) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
