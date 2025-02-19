package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// DeleteGatewayResponseTypeV2Request Request Object
type DeleteGatewayResponseTypeV2Request struct {

	// 实例ID，在API网关控制台的“实例信息”中获取。
	InstanceId string `json:"instance_id"`

	// 分组的编号
	GroupId string `json:"group_id"`

	// 响应编号
	ResponseId string `json:"response_id"`

	// 错误类型 - AUTH_FAILURE: 认证失败，IAM或APP认证校验失败 - AUTH_HEADER_MISSING: 认证身份来源信息缺失 - AUTHORIZER_FAILURE: 自定义认证方返回认证失败 - AUTHORIZER_CONF_FAILURE:自定义认证方异常，通信失败、返回异常响应等错误 - AUTHORIZER_IDENTITIES_FAILURE: 前端自定义认证的身份来源信息缺失或不合法错误 - BACKEND_UNAVAILABLE: 后端不可用，网络不可达错误 - BACKEND_TIMEOUT: 后端超时，与后端的网络交互超过预配置的时间错误 - THROTTLED: API调用次数超出所配置的流量策略阈值 - UNAUTHORIZED: 使用的凭据未被授权访问该API - ACCESS_DENIED: 拒绝访问，如触发配置的访问控制策略、或异常攻击检测拦截 - NOT_FOUND: 未匹配到API错误 - REQUEST_PARAMETERS_FAILURE: 请求参数校验失败、不支持的HTTP方法 - DEFAULT_4XX: 其它4XX类错误 - DEFAULT_5XX: 其它5XX类错误 - THIRD_AUTH_FAILURE: 第三方认证方返回认证失败 - THIRD_AUTH_IDENTITIES_FAILURE: 第三方认证的身份来源信息缺失或不合法错误 - THIRD_AUTH_CONF_FAILURE: 第三方认证方异常，通信失败、返回异常响应等错误
	ResponseType DeleteGatewayResponseTypeV2RequestResponseType `json:"response_type"`
}

func (o DeleteGatewayResponseTypeV2Request) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteGatewayResponseTypeV2Request struct{}"
	}

	return strings.Join([]string{"DeleteGatewayResponseTypeV2Request", string(data)}, " ")
}

type DeleteGatewayResponseTypeV2RequestResponseType struct {
	value string
}

type DeleteGatewayResponseTypeV2RequestResponseTypeEnum struct {
	AUTH_FAILURE                  DeleteGatewayResponseTypeV2RequestResponseType
	AUTH_HEADER_MISSING           DeleteGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_FAILURE            DeleteGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_CONF_FAILURE       DeleteGatewayResponseTypeV2RequestResponseType
	AUTHORIZER_IDENTITIES_FAILURE DeleteGatewayResponseTypeV2RequestResponseType
	BACKEND_UNAVAILABLE           DeleteGatewayResponseTypeV2RequestResponseType
	BACKEND_TIMEOUT               DeleteGatewayResponseTypeV2RequestResponseType
	THROTTLED                     DeleteGatewayResponseTypeV2RequestResponseType
	UNAUTHORIZED                  DeleteGatewayResponseTypeV2RequestResponseType
	ACCESS_DENIED                 DeleteGatewayResponseTypeV2RequestResponseType
	NOT_FOUND                     DeleteGatewayResponseTypeV2RequestResponseType
	REQUEST_PARAMETERS_FAILURE    DeleteGatewayResponseTypeV2RequestResponseType
	DEFAULT_4_XX                  DeleteGatewayResponseTypeV2RequestResponseType
	DEFAULT_5_XX                  DeleteGatewayResponseTypeV2RequestResponseType
	THIRD_AUTH_FAILURE            DeleteGatewayResponseTypeV2RequestResponseType
	THIRD_AUTH_IDENTITIES_FAILURE DeleteGatewayResponseTypeV2RequestResponseType
	THIRD_AUTH_CONF_FAILURE       DeleteGatewayResponseTypeV2RequestResponseType
}

func GetDeleteGatewayResponseTypeV2RequestResponseTypeEnum() DeleteGatewayResponseTypeV2RequestResponseTypeEnum {
	return DeleteGatewayResponseTypeV2RequestResponseTypeEnum{
		AUTH_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTH_FAILURE",
		},
		AUTH_HEADER_MISSING: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTH_HEADER_MISSING",
		},
		AUTHORIZER_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_FAILURE",
		},
		AUTHORIZER_CONF_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_CONF_FAILURE",
		},
		AUTHORIZER_IDENTITIES_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "AUTHORIZER_IDENTITIES_FAILURE",
		},
		BACKEND_UNAVAILABLE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "BACKEND_UNAVAILABLE",
		},
		BACKEND_TIMEOUT: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "BACKEND_TIMEOUT",
		},
		THROTTLED: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "THROTTLED",
		},
		UNAUTHORIZED: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "UNAUTHORIZED",
		},
		ACCESS_DENIED: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "ACCESS_DENIED",
		},
		NOT_FOUND: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "NOT_FOUND",
		},
		REQUEST_PARAMETERS_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "REQUEST_PARAMETERS_FAILURE",
		},
		DEFAULT_4_XX: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "DEFAULT_4XX",
		},
		DEFAULT_5_XX: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "DEFAULT_5XX",
		},
		THIRD_AUTH_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "THIRD_AUTH_FAILURE",
		},
		THIRD_AUTH_IDENTITIES_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "THIRD_AUTH_IDENTITIES_FAILURE",
		},
		THIRD_AUTH_CONF_FAILURE: DeleteGatewayResponseTypeV2RequestResponseType{
			value: "THIRD_AUTH_CONF_FAILURE",
		},
	}
}

func (c DeleteGatewayResponseTypeV2RequestResponseType) Value() string {
	return c.value
}

func (c DeleteGatewayResponseTypeV2RequestResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *DeleteGatewayResponseTypeV2RequestResponseType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
