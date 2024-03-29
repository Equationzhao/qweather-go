package air

import "github.com/Equationzhao/qweather-go/statusCode"

// NowResponse
// code 请参考状态码
// updateTime 当前API的最近更新时间
// fxLink 当前数据的响应式页面，便于嵌入网站或应用
// now.pubTime 空气质量数据发布时间
// now.aqi 空气质量指数
// now.level 空气质量指数等级
// now.category 空气质量指数级别
// now.primary 空气质量的主要污染物，空气质量为优时，返回值为NA
// now.pm10 PM10
// now.pm2p5 PM2.5
// now.no2 二氧化氮
// now.so2 二氧化硫
// now.co 一氧化碳
// now.o3 臭氧
// station.name 监测站名称
// station.id 监测站ID
// station.pubTime 空气质量数据发布时间
// station.aqi 空气质量指数
// station.level 空气质量指数等级
// station.category 空气质量指数级别
// station.primary 空气质量的主要污染物，空气质量为优时，返回值为NA
// station.pm10 PM10
// station.pm2p5 PM2.5
// station.no2 二氧化氮
// station.so2 二氧化硫
// station.co 一氧化碳
// station.o3 臭氧
// refer.sources 原始数据来源，或数据源说明，可能为空
// refer.license 数据许可或版权声明，可能为空
type NowResponse struct {
	Code       statusCode.Code `json:"code"`
	UpdateTime string          `json:"updateTime"`
	FxLink     string          `json:"fxLink"`
	Now        struct {
		PubTime  string `json:"pubTime"`
		Aqi      string `json:"aqi"`
		Level    string `json:"level"`
		Category string `json:"category"`
		Primary  string `json:"primary"`
		Pm10     string `json:"pm10"`
		Pm2P5    string `json:"pm2p5"`
		No2      string `json:"no2"`
		So2      string `json:"so2"`
		Co       string `json:"co"`
		O3       string `json:"o3"`
	} `json:"now"`
	Station []struct {
		PubTime  string `json:"pubTime"`
		Name     string `json:"name"`
		Id       string `json:"id"`
		Aqi      string `json:"aqi"`
		Level    string `json:"level"`
		Category string `json:"category"`
		Primary  string `json:"primary"`
		Pm10     string `json:"pm10"`
		Pm2P5    string `json:"pm2p5"`
		No2      string `json:"no2"`
		So2      string `json:"so2"`
		Co       string `json:"co"`
		O3       string `json:"o3"`
	} `json:"station"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}

// Day5Response
// code 请参考状态码
// updateTime 当前API的最近更新时间
// fxLink 当前数据的响应式页面，便于嵌入网站或应用
// daily.fxDate 预报日期
// daily.aqi 空气质量指数
// daily.level 空气质量指数等级
// daily.category 空气质量指数级别
// daily.primary 空气质量的主要污染物，空气质量为优时，返回值为NA
// refer.sources 原始数据来源，或数据源说明，可能为空
// refer.license 数据许可或版权声明，可能为空
type Day5Response struct {
	Code       statusCode.Code `json:"code"`
	UpdateTime string          `json:"updateTime"`
	FxLink     string          `json:"fxLink"`
	Daily      []struct {
		FxDate   string `json:"fxDate"`
		Aqi      string `json:"aqi"`
		Level    string `json:"level"`
		Category string `json:"category"`
		Primary  string `json:"primary"`
	} `json:"daily"`
	Refer struct {
		Sources []string `json:"sources"`
		License []string `json:"license"`
	} `json:"refer"`
}
