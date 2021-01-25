package model

import (
	"encoding/json"

	"strings"
)

// SSD类型云硬盘预留的云硬盘个数，键值对，包含：reserved（预留）、allocated（预留）、limit（最大）和in_use（已使用）。
type QuotaDetailVolumesSsd struct {
	// 已使用的数量。
	InUse int32 `json:"in_use"`
	// 最大的数量。
	Limit int32 `json:"limit"`
	// 预留属性。
	Reserved int32 `json:"reserved"`
	// 预留属性。
	Allocated string `json:"allocated"`
}

func (o QuotaDetailVolumesSsd) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "QuotaDetailVolumesSsd struct{}"
	}

	return strings.Join([]string{"QuotaDetailVolumesSsd", string(data)}, " ")
}
