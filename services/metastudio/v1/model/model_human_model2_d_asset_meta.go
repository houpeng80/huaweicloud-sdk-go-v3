package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type HumanModel2DAssetMeta struct {

	// 分身数字人的动作是否可编辑。默认不可编辑。
	IsActionEditable *bool `json:"is_action_editable,omitempty"`

	// 是否是实景分身数字人。实景分身数字人不做背景替换。
	IsRealBackground *bool `json:"is_real_background,omitempty"`

	// 是否支持直播
	SupportLive *bool `json:"support_live,omitempty"`

	// V3模型分身数字人是否需要实时抠图。
	IsRealtimeMatting *bool `json:"is_realtime_matting,omitempty"`

	// 分身数字人模型版本。默认是V2版本模型。 * V2: V2版本模型 * V3：V3版本模型
	ModelVersion *HumanModel2DAssetMetaModelVersion `json:"model_version,omitempty"`
}

func (o HumanModel2DAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HumanModel2DAssetMeta struct{}"
	}

	return strings.Join([]string{"HumanModel2DAssetMeta", string(data)}, " ")
}

type HumanModel2DAssetMetaModelVersion struct {
	value string
}

type HumanModel2DAssetMetaModelVersionEnum struct {
	V2 HumanModel2DAssetMetaModelVersion
	V3 HumanModel2DAssetMetaModelVersion
}

func GetHumanModel2DAssetMetaModelVersionEnum() HumanModel2DAssetMetaModelVersionEnum {
	return HumanModel2DAssetMetaModelVersionEnum{
		V2: HumanModel2DAssetMetaModelVersion{
			value: "V2",
		},
		V3: HumanModel2DAssetMetaModelVersion{
			value: "V3",
		},
	}
}

func (c HumanModel2DAssetMetaModelVersion) Value() string {
	return c.value
}

func (c HumanModel2DAssetMetaModelVersion) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *HumanModel2DAssetMetaModelVersion) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
