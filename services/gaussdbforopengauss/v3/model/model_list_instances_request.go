package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ListInstancesRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`
	// 实例ID。  “\\*”为系统保留字符，如果id是以“\\*”起始，表示按照\\*后面的值模糊匹配，否则，按照id精确匹配查询。不能只传入“\\*”。

	Id *string `json:"id,omitempty"`
	// 实例名称。  “\\*”为系统保留字符，如果name是以“\\*”起始，表示按照\\*后面的值模糊匹配，否则，按照name精确匹配查询。不能只传入“\\*”。

	Name *string `json:"name,omitempty"`
	// 按照实例类型查询。目前仅支持取值“Enterprise”（区分大小写），对应分布式实例（企业版）。

	Type *ListInstancesRequestType `json:"type,omitempty"`
	// 数据库类型，区分大小写。  - GaussDB(openGauss)

	DatastoreType *ListInstancesRequestDatastoreType `json:"datastore_type,omitempty"`
	// 虚拟私有云ID。  方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询VPC列表](https://support.huaweicloud.com/api-vpc/vpc_api01_0003.html)。

	VpcId *string `json:"vpc_id,omitempty"`
	// 子网的网络ID信息。  - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体操作可参考[查询子网列表](https://support.huaweicloud.com/api-vpc/vpc_subnet01_0003.html)。

	SubnetId *string `json:"subnet_id,omitempty"`
	// 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询，默认为0（偏移0条数据，表示从第一条数据开始查询），必须为数字，不能为负数。

	Offset *int32 `json:"offset,omitempty"`
	// 查询记录数。默认为100，不能为负数，最小值为1，最大值为100

	Limit *int32 `json:"limit,omitempty"`
	// 根据实例标签键值对进行查询。  {key}表示标签键，不可以为空或重复。 {value}表示标签值，可以为空。 如果同时使用多个标签键值对进行查询，中间使用逗号分隔开，最多包含10组。

	Tags *[]string `json:"tags,omitempty"`
}

func (o ListInstancesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstancesRequest struct{}"
	}

	return strings.Join([]string{"ListInstancesRequest", string(data)}, " ")
}

type ListInstancesRequestType struct {
	value string
}

type ListInstancesRequestTypeEnum struct {
	ENTERPRISE ListInstancesRequestType
}

func GetListInstancesRequestTypeEnum() ListInstancesRequestTypeEnum {
	return ListInstancesRequestTypeEnum{
		ENTERPRISE: ListInstancesRequestType{
			value: "Enterprise",
		},
	}
}

func (c ListInstancesRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListInstancesRequestType) UnmarshalJSON(b []byte) error {
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

type ListInstancesRequestDatastoreType struct {
	value string
}

type ListInstancesRequestDatastoreTypeEnum struct {
	GAUSS_DB_OPEN_GAUSS ListInstancesRequestDatastoreType
}

func GetListInstancesRequestDatastoreTypeEnum() ListInstancesRequestDatastoreTypeEnum {
	return ListInstancesRequestDatastoreTypeEnum{
		GAUSS_DB_OPEN_GAUSS: ListInstancesRequestDatastoreType{
			value: "GaussDB(openGauss)",
		},
	}
}

func (c ListInstancesRequestDatastoreType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ListInstancesRequestDatastoreType) UnmarshalJSON(b []byte) error {
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
