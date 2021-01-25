package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListPredefineTagsRequest struct {
	Key         *string                              `json:"key,omitempty"`
	Value       *string                              `json:"value,omitempty"`
	Limit       *int32                               `json:"limit,omitempty"`
	Marker      *string                              `json:"marker,omitempty"`
	OrderField  *string                              `json:"order_field,omitempty"`
	OrderMethod *ListPredefineTagsRequestOrderMethod `json:"order_method,omitempty"`
}

func (o ListPredefineTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPredefineTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPredefineTagsRequest", string(data)}, " ")
}

type ListPredefineTagsRequestOrderMethod struct {
	value string
}

type ListPredefineTagsRequestOrderMethodEnum struct {
	ASC  ListPredefineTagsRequestOrderMethod
	DESC ListPredefineTagsRequestOrderMethod
}

func GetListPredefineTagsRequestOrderMethodEnum() ListPredefineTagsRequestOrderMethodEnum {
	return ListPredefineTagsRequestOrderMethodEnum{
		ASC: ListPredefineTagsRequestOrderMethod{
			value: "asc",
		},
		DESC: ListPredefineTagsRequestOrderMethod{
			value: "desc",
		},
	}
}

func (c ListPredefineTagsRequestOrderMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListPredefineTagsRequestOrderMethod) UnmarshalJSON(b []byte) error {
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
