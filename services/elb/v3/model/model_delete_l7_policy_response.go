package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteL7PolicyResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteL7PolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteL7PolicyResponse struct{}"
	}

	return strings.Join([]string{"DeleteL7PolicyResponse", string(data)}, " ")
}
