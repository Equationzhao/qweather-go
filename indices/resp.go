package indices

import "github.com/Equationzhao/qweather-go/statusCode"

type Response struct {
	Code       statusCode.Code `json:"code"`       // 状态码
	UpdateTime string          `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/api/indices/
	FxLink     string          `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Daily      []struct {
		Date     string `json:"date"`     // 预报日期
		Type     string `json:"type"`     // 生活指数类型ID
		Name     string `json:"name"`     // 生活指数类型名称
		Level    string `json:"level"`    // 生活指数预报等级
		Category string `json:"category"` // 生活指数预报级别名称
		Text     string `json:"text"`     // 生活指数预报的详细描述，**可能为空**
	} `json:"daily"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}
