package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListListenersResponse struct {
	// 监听器对象列表
	Listeners      *[]ListenerResp `json:"listeners,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListListenersResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListListenersResponse struct{}"
	}

	return strings.Join([]string{"ListListenersResponse", string(data)}, " ")
}
