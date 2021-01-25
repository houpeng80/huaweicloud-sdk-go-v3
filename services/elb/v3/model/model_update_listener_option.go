package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

// 更新监听器请求体
type UpdateListenerOption struct {
	// 监听器的管理状态。只支持设定为true，该字段的值无实际意义。
	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 监听器使用的CA证书ID。
	ClientCaTlsContainerRef *string `json:"client_ca_tls_container_ref,omitempty"`
	// 监听器的默认后端云服务器组ID。当请求没有匹配的转发策略时，转发到默认后端云服务器上处理。
	DefaultPoolId *string `json:"default_pool_id,omitempty"`
	// 监听器使用的服务器证书ID。
	DefaultTlsContainerRef *string `json:"default_tls_container_ref,omitempty"`
	// 监听器的描述信息
	Description *string `json:"description,omitempty"`
	// HTTP2功能的开启状态。该字段只有当监听器的协议是TERMINATED_HTTPS时生效。
	Http2Enable   *bool                  `json:"http2_enable,omitempty"`
	InsertHeaders *ListenerInsertHeaders `json:"insert_headers,omitempty"`
	// 监听器名称
	Name *string `json:"name,omitempty"`
	// 监听器使用的SNI证书（带域名的服务器证书）ID的列表。 各SNI证书的域名不允许重复。 各SNI证书域名总数不超过30。
	SniContainerRefs *[]string `json:"sni_container_refs,omitempty"`
	// 监听器使用的安全策略，仅对TERMINATED_HTTPS协议类型的监听器有效，且默认值为tls-1-0。 取值包括：tls-1-0-inherit,tls-1-0, tls-1-1, tls-1-2, tls-1-2-strict, tls-1-2-fst 六种安全策略
	TlsCiphersPolicy *UpdateListenerOptionTlsCiphersPolicy `json:"tls_ciphers_policy,omitempty"`
	// 是否关闭后端服务器的重试。 当前仅七层的性能共享型实例支持指定该字段。
	EnableMemberRetry *bool `json:"enable_member_retry,omitempty"`
	// 等待后端服务器请求超时时间，仅限协议为HTTP， TERMINATED_HTTPS的监听器配置。取值范围为1-300s，默认为60s TCP，UDP协议的监听器不支持此字段
	MemberTimeout *int32 `json:"member_timeout,omitempty"`
	// 等待客户端请求超时时间，仅限协议为HTTP， TERMINATED_HTTPS的监听器配置。取值范围为1-60s, 默认值为60s TCP，UDP协议的监听器不支持此字段
	ClientTimeout *int32 `json:"client_timeout,omitempty"`
	// TCP监听器配置空闲超时时间，取值范围为（10-900s）默认值为300s，HTTP/TERMINATED_HTTPS监听器为客户端连接空闲超时时间，取值范围为（1-300s）默认值为15s。 UDP监听器不支持此字段
	KeepaliveTimeout *int32                       `json:"keepalive_timeout,omitempty"`
	Ipgroup          *UpdateListenerIpGroupOption `json:"ipgroup,omitempty"`
	// 获取客户端真实IP 共享型实例的TCP/UDP监听器支持修改，共享型实例的HTTP/TERMINATED_HTTPS监听器和独享型实例的所有类型监听器都不支持修改
	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`
}

func (o UpdateListenerOption) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateListenerOption struct{}"
	}

	return strings.Join([]string{"UpdateListenerOption", string(data)}, " ")
}

type UpdateListenerOptionTlsCiphersPolicy struct {
	value string
}

type UpdateListenerOptionTlsCiphersPolicyEnum struct {
	TLS_1_0_INHERIT UpdateListenerOptionTlsCiphersPolicy
	TLS_1_0         UpdateListenerOptionTlsCiphersPolicy
	TLS_1_1         UpdateListenerOptionTlsCiphersPolicy
	TLS_1_2         UpdateListenerOptionTlsCiphersPolicy
	TLS_1_2_STRICT  UpdateListenerOptionTlsCiphersPolicy
	TLS_1_2_FS      UpdateListenerOptionTlsCiphersPolicy
}

func GetUpdateListenerOptionTlsCiphersPolicyEnum() UpdateListenerOptionTlsCiphersPolicyEnum {
	return UpdateListenerOptionTlsCiphersPolicyEnum{
		TLS_1_0_INHERIT: UpdateListenerOptionTlsCiphersPolicy{
			value: "tls-1-0-inherit",
		},
		TLS_1_0: UpdateListenerOptionTlsCiphersPolicy{
			value: "tls-1-0",
		},
		TLS_1_1: UpdateListenerOptionTlsCiphersPolicy{
			value: " tls-1-1",
		},
		TLS_1_2: UpdateListenerOptionTlsCiphersPolicy{
			value: " tls-1-2",
		},
		TLS_1_2_STRICT: UpdateListenerOptionTlsCiphersPolicy{
			value: " tls-1-2-strict",
		},
		TLS_1_2_FS: UpdateListenerOptionTlsCiphersPolicy{
			value: "tls-1-2-fs",
		},
	}
}

func (c UpdateListenerOptionTlsCiphersPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *UpdateListenerOptionTlsCiphersPolicy) UnmarshalJSON(b []byte) error {
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
