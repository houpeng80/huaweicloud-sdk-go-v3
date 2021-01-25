package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type DeleteTrackerRequest struct {
	TrackerName *string                          `json:"tracker_name,omitempty"`
	TrackerType *DeleteTrackerRequestTrackerType `json:"tracker_type,omitempty"`
}

func (o DeleteTrackerRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteTrackerRequest struct{}"
	}

	return strings.Join([]string{"DeleteTrackerRequest", string(data)}, " ")
}

type DeleteTrackerRequestTrackerType struct {
	value string
}

type DeleteTrackerRequestTrackerTypeEnum struct {
	DATA DeleteTrackerRequestTrackerType
}

func GetDeleteTrackerRequestTrackerTypeEnum() DeleteTrackerRequestTrackerTypeEnum {
	return DeleteTrackerRequestTrackerTypeEnum{
		DATA: DeleteTrackerRequestTrackerType{
			value: "data",
		},
	}
}

func (c DeleteTrackerRequestTrackerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DeleteTrackerRequestTrackerType) UnmarshalJSON(b []byte) error {
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
