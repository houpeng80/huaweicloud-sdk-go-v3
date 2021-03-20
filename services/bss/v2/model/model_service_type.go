package model

import (
	"encoding/json"

	"strings"
)

type ServiceType struct {
	// 云服务类型的名称。

	ServiceTypeName *string `json:"service_type_name,omitempty"`
	// 云服务类型的编码。

	ServiceTypeCode *string `json:"service_type_code,omitempty"`
	// 云服务类型的缩写。

	Abbreviation *string `json:"abbreviation,omitempty"`
}

func (o ServiceType) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServiceType struct{}"
	}

	return strings.Join([]string{"ServiceType", string(data)}, " ")
}
