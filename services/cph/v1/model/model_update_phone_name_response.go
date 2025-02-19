package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePhoneNameResponse Response Object
type UpdatePhoneNameResponse struct {

	// 请求的唯一标识ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePhoneNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePhoneNameResponse struct{}"
	}

	return strings.Join([]string{"UpdatePhoneNameResponse", string(data)}, " ")
}
