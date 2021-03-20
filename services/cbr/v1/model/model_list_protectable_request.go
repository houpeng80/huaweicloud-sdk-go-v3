package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListProtectableRequest struct {
	Limit *int32 `json:"limit,omitempty"`

	Marker *string `json:"marker,omitempty"`

	Name *string `json:"name,omitempty"`

	Offset *int32 `json:"offset,omitempty"`

	ProtectableType ListProtectableRequestProtectableType `json:"protectable_type"`

	Status *string `json:"status,omitempty"`

	Id *string `json:"id,omitempty"`

	ServerId *string `json:"server_id,omitempty"`
}

func (o ListProtectableRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListProtectableRequest struct{}"
	}

	return strings.Join([]string{"ListProtectableRequest", string(data)}, " ")
}

type ListProtectableRequestProtectableType struct {
	value string
}

type ListProtectableRequestProtectableTypeEnum struct {
	SERVER ListProtectableRequestProtectableType
	DISK   ListProtectableRequestProtectableType
}

func GetListProtectableRequestProtectableTypeEnum() ListProtectableRequestProtectableTypeEnum {
	return ListProtectableRequestProtectableTypeEnum{
		SERVER: ListProtectableRequestProtectableType{
			value: "server",
		},
		DISK: ListProtectableRequestProtectableType{
			value: "disk",
		},
	}
}

func (c ListProtectableRequestProtectableType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListProtectableRequestProtectableType) UnmarshalJSON(b []byte) error {
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
