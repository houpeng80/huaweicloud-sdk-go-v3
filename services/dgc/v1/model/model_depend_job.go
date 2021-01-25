package model

import (
	"encoding/json"
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"strings"
)

type DependJob struct {
	Jobs             *string                    `json:"jobs,omitempty"`
	DependPeriod     *string                    `json:"dependPeriod,omitempty"`
	DependFailPolicy *DependJobDependFailPolicy `json:"dependFailPolicy,omitempty"`
}

func (o DependJob) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DependJob struct{}"
	}

	return strings.Join([]string{"DependJob", string(data)}, " ")
}

type DependJobDependFailPolicy struct {
	value string
}

type DependJobDependFailPolicyEnum struct {
	FAIL    DependJobDependFailPolicy
	IGNORE  DependJobDependFailPolicy
	SUSPEND DependJobDependFailPolicy
}

func GetDependJobDependFailPolicyEnum() DependJobDependFailPolicyEnum {
	return DependJobDependFailPolicyEnum{
		FAIL: DependJobDependFailPolicy{
			value: "FAIL",
		},
		IGNORE: DependJobDependFailPolicy{
			value: "IGNORE",
		},
		SUSPEND: DependJobDependFailPolicy{
			value: "SUSPEND",
		},
	}
}

func (c DependJobDependFailPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *DependJobDependFailPolicy) UnmarshalJSON(b []byte) error {
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
