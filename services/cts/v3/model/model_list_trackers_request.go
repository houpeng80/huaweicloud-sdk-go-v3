package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListTrackersRequest struct {
	TrackerName *string `json:"tracker_name,omitempty"`

	TrackerType *ListTrackersRequestTrackerType `json:"tracker_type,omitempty"`
}

func (o ListTrackersRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListTrackersRequest struct{}"
	}

	return strings.Join([]string{"ListTrackersRequest", string(data)}, " ")
}

type ListTrackersRequestTrackerType struct {
	value string
}

type ListTrackersRequestTrackerTypeEnum struct {
	SYSTEM ListTrackersRequestTrackerType
	DATA   ListTrackersRequestTrackerType
}

func GetListTrackersRequestTrackerTypeEnum() ListTrackersRequestTrackerTypeEnum {
	return ListTrackersRequestTrackerTypeEnum{
		SYSTEM: ListTrackersRequestTrackerType{
			value: "system",
		},
		DATA: ListTrackersRequestTrackerType{
			value: "data",
		},
	}
}

func (c ListTrackersRequestTrackerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListTrackersRequestTrackerType) UnmarshalJSON(b []byte) error {
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
