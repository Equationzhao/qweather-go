package airqualityBeta

import "github.com/Equationzhao/qweather-go/statusCode"

type NowAirQualityResponse struct {
	Code       statusCode.Code `json:"code"`       // 状态码
	UpdateTime string          `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/resource/glossary/#update-time
	Aqi        []struct {
		Code            string `json:"code"`            // 空气质量指数 Code https://dev.qweather.com/docs/resource/air-info/#supported-aqis
		Name            string `json:"name"`            // 空气质量指数的名字
		DefaultLocalAqi bool   `json:"defaultLocalAqi"` // 是否是默认/推荐的当地AQI https://dev.qweather.com/docs/resource/air-info/#default-local-aqi
		Value           int    `json:"value"`           // 空气质量指数的值 https://dev.qweather.com/docs/resource/air-info/#aqi-value
		ValueDisplay    string `json:"valueDisplay"`    // 空气质量指数的值的文本显示 https://dev.qweather.com/docs/resource/air-info/#aqi-value
		Level           string `json:"level"`           // 空气质量指数等级，可能为空
		Category        string `json:"category"`        // 空气质量指数类别，可能为空
		Color           string `json:"color"`           // 空气质量指数的颜色，RGB格式
		Health          struct {
			Effect string `json:"effect"` // 空气质量对健康的影响，可能为空 https://dev.qweather.com/docs/resource/air-info/#health-effects-and-advice
			Advice struct {
				GeneralPopulation   string `json:"generalPopulation"`   // 对一般人群的健康指导意见，可能为空
				SensitivePopulation string `json:"sensitivePopulation"` // 对敏感人群的健康指导意见，可能为空
			} `json:"advice"`
		} `json:"health"`
		PrimaryPollutant struct {
			Code     string `json:"code"`     // 首要污染物的Code，可能为空 https://dev.qweather.com/docs/resource/air-info/#primary-pollutant
			Name     string `json:"name"`     // 首要污染物的名字，可能为空
			FullName string `json:"fullName"` // 首要污染物的全称，可能为空
		} `json:"primaryPollutant,omitempty"`
	} `json:"aqi"`
	Pollutant []struct {
		Code          string `json:"code"`     // 污染物的 Code https://dev.qweather.com/docs/resource/air-info/#pollutants
		Name          string `json:"name"`     // 污染物的名字
		FullName      string `json:"fullName"` // 污染物的全称
		Concentration struct {
			Value float64 `json:"value"` // 污染物的浓度值
			Unit  string  `json:"unit"`  // 污染物的浓度值的单位
		} `json:"concentration"`
		SubIndex struct {
			Value        int    `json:"value"`        // 污染物的分指数的数值，可能为空 https://dev.qweather.com/docs/resource/air-info/#pollutant-sub-index
			ValueDisplay string `json:"valueDisplay"` // 污染物的分指数数值的显示名称
		} `json:"subIndex"`
	} `json:"pollutant"`
	Station []struct {
		Id   string `json:"id"`   // AQI相关联的监测站Location ID，可能为空
		Name string `json:"name"` // AQI相关联的监测站名称
	} `json:"station"`
	Source []string `json:"source"` // 数据来源或提供商名字以及他们的声明，必须与空气质量数据一起展示。可能为空
}

type StationResponse struct {
	Code       statusCode.Code `json:"code"`       // 状态码
	UpdateTime string          `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/resource/glossary/#update-time
	Pollutant  []struct {
		Code          string `json:"code"`     // 污染物的 Code https://dev.qweather.com/docs/resource/air-info/#pollutants
		Name          string `json:"name"`     // 污染物的名字
		FullName      string `json:"fullName"` // 污染物的全称
		Concentration struct {
			Value string `json:"value"` // 污染物的浓度值
			Unit  string `json:"unit"`  // 污染物的浓度值的单位
		} `json:"concentration"`
	} `json:"pollutant"`
	Source []string `json:"source"` // 数据来源或提供商名字以及他们的声明，必须与空气质量数据一起展示。可能为空
}
