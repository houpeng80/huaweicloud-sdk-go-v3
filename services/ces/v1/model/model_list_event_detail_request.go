package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// Request Object
type ListEventDetailRequest struct {
	ContentType string                          `json:"Content-Type"`
	EventName   string                          `json:"event_name"`
	EventType   ListEventDetailRequestEventType `json:"event_type"`
	EventSource *string                         `json:"event_source,omitempty"`
	EventLevel  *string                         `json:"event_level,omitempty"`
	EventUser   *string                         `json:"event_user,omitempty"`
	EventState  *string                         `json:"event_state,omitempty"`
	From        *int64                          `json:"from,omitempty"`
	To          *int64                          `json:"to,omitempty"`
	Start       *int32                          `json:"start,omitempty"`
	Limit       *int32                          `json:"limit,omitempty"`
}

func (o ListEventDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEventDetailRequest struct{}"
	}

	return strings.Join([]string{"ListEventDetailRequest", string(data)}, " ")
}

type ListEventDetailRequestEventType struct {
	value string
}

type ListEventDetailRequestEventTypeEnum struct {
	EVENT_SYS    ListEventDetailRequestEventType
	EVENT_CUSTOM ListEventDetailRequestEventType
}

func GetListEventDetailRequestEventTypeEnum() ListEventDetailRequestEventTypeEnum {
	return ListEventDetailRequestEventTypeEnum{
		EVENT_SYS: ListEventDetailRequestEventType{
			value: "EVENT.SYS",
		},
		EVENT_CUSTOM: ListEventDetailRequestEventType{
			value: "EVENT.CUSTOM",
		},
	}
}

func (c ListEventDetailRequestEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListEventDetailRequestEventType) UnmarshalJSON(b []byte) error {
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
