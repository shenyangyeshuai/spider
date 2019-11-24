package model

import (
	"encoding/json"
)

type Profile struct {
	// 姓名
	Name string

	// 性别
	Gender string

	// 年龄
	Age int

	// 身高
	Height int

	// 体重
	// Weight int

	// 收入(通常是一个范围)
	Income string

	// 婚况
	Marriage string

	// 学历
	// Education string

	// 职业
	Occupation string

	// 户口(籍贯)
	// Hukou string

	// 星座
	Xingzuo string

	// 是否有房
	// House string

	// 是否有车
	// Car string
}

func FromJsonObj(o interface{}) (*Profile, error) {
	var profile Profile

	s, err := json.Marshal(o)
	if err != nil {
		return &profile, err
	}

	err = json.Unmarshal(s, &profile)
	return &profile, err
}
