package minForecast

type MinPrecipitationResponse struct {
	Code       string `json:"code"`       // 状态码
	UpdateTime string `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/resource/glossary/#update-time
	FxLink     string `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Summary    string `json:"summary"`    // 分钟降水描述
	Minutely   []struct {
		FxTime string `json:"fxTime"` // 预报时间
		Precip string `json:"precip"` // 5分钟累计降水量，单位毫米
		Type   string `json:"type"`   // 降水类型：rain = 雨，snow = 雪
	} `json:"minutely"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}
