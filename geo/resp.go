package geo

import "github.com/Equationzhao/qweather-go/statusCode"

type SearchResponse struct {
	Code     statusCode.Code `json:"code"` // 状态码
	Location []struct {
		Name      string `json:"name"`      // 地区/城市名称
		Id        string `json:"id"`        // 地区/城市ID
		Lat       string `json:"lat"`       // 地区/城市纬度
		Lon       string `json:"lon"`       // 地区/城市经度
		Adm2      string `json:"adm2"`      // 地区/城市上级行政区域
		Adm1      string `json:"adm1"`      // 地区/城市所属一级行政区域
		Country   string `json:"country"`   // 地区/城市所属国家名称
		Tz        string `json:"tz"`        // 地区/城市所在时区
		UtcOffset string `json:"utcOffset"` // 地区/城市所在时区偏移小时数 参考详细说明 https://dev.qweather.com/docs/resource/glossary/#utc-offset
		IsDst     string `json:"isDst"`     // 地区/城市是否当前处于夏令时 https://dev.qweather.com/docs/resource/glossary/#daylight-saving-time 1 表示当前处于夏令时，0 表示当前不是夏令时。
		Type      string `json:"type"`      // 地区/城市的属性
		Rank      string `json:"rank"`      // 地区评分 https://dev.qweather.com/docs/resource/glossary/#rank
		FxLink    string `json:"fxLink"`    // 地区/城市对应的天气预报网页链接，便于嵌入网站或应用
	} `json:"location"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}

type HitResponse struct {
	Code        statusCode.Code `json:"code"` // 状态码
	TopCityList []struct {
		Name      string `json:"name"`      // 地区/城市名称
		Id        string `json:"id"`        // 地区/城市ID
		Lat       string `json:"lat"`       // 地区/城市纬度
		Lon       string `json:"lon"`       // 地区/城市经度
		Adm2      string `json:"adm2"`      // 地区/城市上级行政区域
		Adm1      string `json:"adm1"`      // 地区/城市所属一级行政区域
		Country   string `json:"country"`   // 地区/城市所属国家名称
		Tz        string `json:"tz"`        // 地区/城市所在时区
		UtcOffset string `json:"utcOffset"` // 地区/城市所在时区偏移小时数 参考详细说明 https://dev.qweather.com/docs/resource/glossary/#utc-offset
		IsDst     string `json:"isDst"`     // 地区/城市是否当前处于夏令时 https://dev.qweather.com/docs/resource/glossary/#daylight-saving-time 1 表示当前处于夏令时，0 表示当前不是夏令时。
		Type      string `json:"type"`      // 地区/城市的属性
		Rank      string `json:"rank"`      // 地区评分 https://dev.qweather.com/docs/resource/glossary/#rank
		FxLink    string `json:"fxLink"`    // 地区/城市对应的天气预报网页链接，便于嵌入网站或应用
	} `json:"topCityList"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}

type POIResponse struct {
	Code statusCode.Code `json:"code"` // 状态码
	Poi  []struct {
		Name      string `json:"name"`      // POI（兴趣点）名称
		Id        string `json:"id"`        // POI（兴趣点）ID
		Lat       string `json:"lat"`       // POI（兴趣点）纬度
		Lon       string `json:"lon"`       // POI（兴趣点）经度
		Adm2      string `json:"adm2"`      // POI（兴趣点）上级行政区域
		Adm1      string `json:"adm1"`      // POI（兴趣点）所属一级行政区域
		Country   string `json:"country"`   // POI（兴趣点）所属国家名称
		Tz        string `json:"tz"`        // POI（兴趣点）所在时区
		UtcOffset string `json:"utcOffset"` // POI（兴趣点）所在时区偏移小时数 参考详细说明 https://dev.qweather.com/docs/resource/glossary/#utc-offset
		IsDst     string `json:"isDst"`     // POI（兴趣点）是否当前处于夏令时 https://dev.qweather.com/docs/resource/glossary/#daylight-saving-time 1 表示当前处于夏令时，0 表示当前不是夏令时。
		Type      string `json:"type"`      // POI（兴趣点）的属性
		Rank      string `json:"rank"`      // 地区评分 https://dev.qweather.com/docs/resource/glossary/#rank
		FxLink    string `json:"fxLink"`    // 该地区的天气预报网页链接，便于嵌入你的网站或应用
	} `json:"poi"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}
