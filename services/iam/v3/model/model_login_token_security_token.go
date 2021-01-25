package model

import (
	"encoding/json"

	"strings"
)

//
type LoginTokenSecurityToken struct {
	// AK。
	Access string `json:"access"`
	// SK。
	Secret string `json:"secret"`
	// securitytoken，即临时身份的安全token。    支持使用自定义代理用户或普通用户获取的securitytoken换取logintoken，详情请参见：[通过token获取临时访问密钥和securitytoken](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=CreateTemporaryAccessKeyByToken)。    支持委托的方式，但获取securitytoken时，请求体中必须填写session_user.name参数，详情请参见：[通过委托获取临时访问密钥和securitytoken](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=IAM&api=CreateTemporaryAccessKeyByAgency)。
	Id string `json:"id"`
}

func (o LoginTokenSecurityToken) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "LoginTokenSecurityToken struct{}"
	}

	return strings.Join([]string{"LoginTokenSecurityToken", string(data)}, " ")
}
