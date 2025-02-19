package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateHealthmonitorReq 创建健康检查请求
type CreateHealthmonitorReq struct {

	// 健康检查所在的项目ID。
	TenantId *string `json:"tenant_id,omitempty"`

	// 健康检查名称。
	Name *string `json:"name,omitempty"`

	// 健康检查的管理状态；该字段虽然支持创建、更新，但实际取值决定于后端云服务器对应的弹性云服务器是否存在。若存在，该值为true，否则，该值为false。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 健康检查端口号。默认为空，表示使用后端云服务器组的端口。
	MonitorPort *int32 `json:"monitor_port,omitempty"`

	// 健康检查的超时时间。建议该值小于delay的值。
	Timeout int32 `json:"timeout"`

	// 健康检查类型
	Type CreateHealthmonitorReqType `json:"type"`

	// 期望HTTP响应状态码，指定下列值：单值，例如200；列表，例如200，202；区间，例如200-204。仅当type为HTTP时生效。该字段为预留字段，暂未启用。
	ExpectedCodes *string `json:"expected_codes,omitempty"`

	// 功能说明：健康检查测试member健康状态时，发送的http请求的域名。仅当type为HTTP时生效。使用说明：默认为空，表示使用负载均衡器的vip作为http请求的目的地址。以数字或字母开头，只能包含数字、字母、’-’、’.’。
	DomainName *string `json:"domain_name,omitempty"`

	// HTTP方法，可以为GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH。仅当type为HTTP时生效。该字段为预留字段，暂未启用。
	UrlPath *string `json:"url_path,omitempty"`

	// HTTP方法，可以为GET、HEAD、POST、PUT、DELETE、TRACE、OPTIONS、CONNECT、PATCH。仅当type为HTTP时生效。该字段为预留字段，暂未启用。
	HttpMethod *string `json:"http_method,omitempty"`

	// 健康检查间隔
	Delay int32 `json:"delay"`

	// 最大重试次数
	MaxRetries int32 `json:"max_retries"`

	// 健康检查关联的后端云服务器组ID
	PoolId string `json:"pool_id"`
}

func (o CreateHealthmonitorReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateHealthmonitorReq struct{}"
	}

	return strings.Join([]string{"CreateHealthmonitorReq", string(data)}, " ")
}

type CreateHealthmonitorReqType struct {
	value string
}

type CreateHealthmonitorReqTypeEnum struct {
	TCP         CreateHealthmonitorReqType
	UDP_CONNECT CreateHealthmonitorReqType
	HTTP        CreateHealthmonitorReqType
}

func GetCreateHealthmonitorReqTypeEnum() CreateHealthmonitorReqTypeEnum {
	return CreateHealthmonitorReqTypeEnum{
		TCP: CreateHealthmonitorReqType{
			value: "TCP",
		},
		UDP_CONNECT: CreateHealthmonitorReqType{
			value: "UDP_CONNECT",
		},
		HTTP: CreateHealthmonitorReqType{
			value: "HTTP",
		},
	}
}

func (c CreateHealthmonitorReqType) Value() string {
	return c.value
}

func (c CreateHealthmonitorReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateHealthmonitorReqType) UnmarshalJSON(b []byte) error {
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
