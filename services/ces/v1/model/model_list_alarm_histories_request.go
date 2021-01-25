package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListAlarmHistoriesRequest struct {
	ContentType string  `json:"Content-Type"`
	GroupId     *string `json:"group_id,omitempty"`
	AlarmId     *string `json:"alarm_id,omitempty"`
	AlarmName   *string `json:"alarm_name,omitempty"`
	AlarmStatus *string `json:"alarm_status,omitempty"`
	AlarmLevel  *string `json:"alarm_level,omitempty"`
	Namespace   *string `json:"namespace,omitempty"`
	From        *string `json:"from,omitempty"`
	To          *string `json:"to,omitempty"`
	Start       *string `json:"start,omitempty"`
	Limit       *string `json:"limit,omitempty"`
}

func (o ListAlarmHistoriesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListAlarmHistoriesRequest struct{}"
	}

	return strings.Join([]string{"ListAlarmHistoriesRequest", string(data)}, " ")
}
