package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePortalFastappModel 快应用模型。
type UpdatePortalFastappModel struct {

	// 快应用名。  > 长度范围为1-30个字符，中文占2个字符，英文占1个字符。
	Name *string `json:"name,omitempty"`

	// 快应用LOGO图片资源ID。  > 图片格式为jpg、bmp、jpeg，分辨率大于等于192*192，比例+-0.15，大小不超过5M。参数值为上传智能信息服务号图片资源API返回的resource_id。
	LogoImg *string `json:"logo_img,omitempty"`

	// 快应用描述。  > 长度范围为1-38个字符，中文占2个字符，英文占1个字符。
	Description *string `json:"description,omitempty"`

	// 快应用跳转链接。
	Deeplink *string `json:"deeplink,omitempty"`

	// 快应用依赖引擎版本。  > 长度范围为1-50个字符，中文占2个字符，英文占1个字符。
	DependEngineVersion *string `json:"depend_engine_version,omitempty"`
}

func (o UpdatePortalFastappModel) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePortalFastappModel struct{}"
	}

	return strings.Join([]string{"UpdatePortalFastappModel", string(data)}, " ")
}
