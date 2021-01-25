package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type EncryptDataResponse struct {
	// 密钥ID。
	KeyId *string `json:"key_id,omitempty"`
	// DEK密文16进制，两位表示1byte。
	CipherText     *string `json:"cipher_text,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o EncryptDataResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EncryptDataResponse struct{}"
	}

	return strings.Join([]string{"EncryptDataResponse", string(data)}, " ")
}
