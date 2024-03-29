package cityWeather

import "github.com/Equationzhao/qweather-go/statusCode"

type RealTimeResponse struct {
	Code       statusCode.Code `json:"code"`       // 状态码
	UpdateTime string          `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/resource/glossary/#update-time
	FxLink     string          `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Now        struct {
		ObsTime   string `json:"obsTime"`   // 数据观测时间
		Temp      string `json:"temp"`      // 温度 默认单位：摄氏度
		FeelsLike string `json:"feelsLike"` // 体感温度，默认单位：摄氏度
		Icon      string `json:"icon"`      // 天气状况的图标代码 https://dev.qweather.com/docs/resource/icons/ 另请参考天气图标项目 https://icons.qweather.com/
		Text      string `json:"text"`      // 天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`   // 风向 https://dev.qweather.com/docs/resource/wind-info/#wind-direction 360 角度
		WindDir   string `json:"windDir"`   // 风向 https://dev.qweather.com/docs/resource/wind-info/#wind-direction
		WindScale string `json:"windScale"` // 风力等级 https://dev.qweather.com/docs/resource/wind-info/#wind-direction
		WindSpeed string `json:"windSpeed"` // 风速 https://dev.qweather.com/docs/resource/wind-info/#wind-speed 公里/小时
		Humidity  string `json:"humidity"`  // 相对湿度 百分比
		Precip    string `json:"precip"`    // 当前小时累计降水量，默认单位：毫米
		Pressure  string `json:"pressure"`  // 大气压强，默认单位：百帕
		Vis       string `json:"vis"`       // 能见度，默认单位：公里
		Cloud     string `json:"cloud"`     // 云量，百分比数值。**可能为空**
		Dew       string `json:"dew"`       // 露点温度。**可能为空**
	} `json:"now"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，或数据源说明，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}

type DailyResponse struct {
	Code       statusCode.Code `json:"code"`       // 状态码
	UpdateTime string          `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/resource/glossary/#update-time
	FxLink     string          `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Daily      []struct {
		FxDate         string `json:"fxDate"`         // 预报日期
		Sunrise        string `json:"sunrise"`        // 日出时间 https://dev.qweather.com/docs/resource/sun-moon-info/#sunrise-and-sunset **在高纬度地区可能为空**
		Sunset         string `json:"sunset"`         // 日落时间 https://dev.qweather.com/docs/resource/sun-moon-info/#sunrise-and-sunset **在高纬度地区可能为空**
		Moonrise       string `json:"moonrise"`       // 当天月升时间 https://dev.qweather.com/docs/resource/sun-moon-info/#moonrise-and-moonset **可能为空**
		Moonset        string `json:"moonset"`        // 当天月落时间 https://dev.qweather.com/docs/resource/sun-moon-info/#moonrise-and-moonset **可能为空**
		MoonPhase      string `json:"moonPhase"`      // 月相名称 https://dev.qweather.com/docs/resource/sun-moon-info/#moon-phase
		MoonPhaseIcon  string `json:"moonPhaseIcon"`  // 月相图标代码 https://dev.qweather.com/docs/resource/icons/ 另请参考天气图标项目 https://icons.qweather.com/
		TempMax        string `json:"tempMax"`        // 预报当天最高温度
		TempMin        string `json:"tempMin"`        // 预报当天最低温度
		IconDay        string `json:"iconDay"`        // 预报白天天气状况的图标代码 https://dev.qweather.com/docs/resource/icons/  另请参考天气图标项目 https://icons.qweather.com/
		TextDay        string `json:"textDay"`        // 预报白天天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		IconNight      string `json:"iconNight"`      // 预报夜间天气状况的图标代码 https://dev.qweather.com/docs/resource/icons/  另请参考天气图标项目 https://icons.qweather.com/
		TextNight      string `json:"textNight"`      // 预报夜间天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360Day     string `json:"wind360Day"`     // 预报白天风向 https://dev.qweather.com/docs/resource/wind-info/#wind-direction 360 角度
		WindDirDay     string `json:"windDirDay"`     // 预报白天风向 https://dev.qweather.com/docs/resource/wind-info/#wind-direction
		WindScaleDay   string `json:"windScaleDay"`   // 预报白天风力等级 https://dev.qweather.com/docs/resource/wind-info/#wind-direction
		WindSpeedDay   string `json:"windSpeedDay"`   // 预报白天风速 https://dev.qweather.com/docs/resource/wind-info/#wind-speed 公里/小时
		Wind360Night   string `json:"wind360Night"`   // 预报夜间风向 https://dev.qweather.com/docs/resource/wind-info/#wind-direction 360 角度
		WindDirNight   string `json:"windDirNight"`   // 预报夜间当天风向 https://dev.qweather.com/docs/resource/wind-info/#wind-direction
		WindScaleNight string `json:"windScaleNight"` // 预报夜间风力等级 https://dev.qweather.com/docs/resource/wind-info/#wind-direction
		WindSpeedNight string `json:"windSpeedNight"` // 预报夜间风速 https://dev.qweather.com/docs/resource/wind-info/#wind-speed 公里/小时
		Humidity       string `json:"humidity"`       // 预报当天相对湿度，百分比数值
		Precip         string `json:"precip"`         // 预报当天累计降水量，默认单位：毫米
		Pressure       string `json:"pressure"`       // 预报当天大气压强，默认单位：百帕
		Vis            string `json:"vis"`            // 预报当天能见度，默认单位：公里
		Cloud          string `json:"cloud"`          // 预报当天平均云量，百分比数值 **可能为空**
		UvIndex        string `json:"uvIndex"`        // 预报当天紫外线强度指数
	} `json:"daily"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}

type HourlyResponse struct {
	Code       statusCode.Code `json:"code"`       // 状态码
	UpdateTime string          `json:"updateTime"` // 当前API的最近更新时间 https://dev.qweather.com/docs/resource/glossary/#update-time
	FxLink     string          `json:"fxLink"`     // 当前数据的响应式页面，便于嵌入网站或应用
	Hourly     []struct {
		FxTime    string `json:"fxTime"`    // 预报时间
		Temp      string `json:"temp"`      // 温度 默认单位：摄氏度
		Icon      string `json:"icon"`      // 天气状况的图标代码 https://dev.qweather.com/docs/resource/icons/ 另请参考天气图标项目 https://icons.qweather.com/
		Text      string `json:"text"`      // 天气状况的文字描述，包括阴晴雨雪等天气状态的描述
		Wind360   string `json:"wind360"`   // 风向 https://dev.qweather.com/docs/resource/wind-info/#wind-direction 360 角度
		WindDir   string `json:"windDir"`   // 风向 https://dev.qweather.com/docs/resource/wind-info/#wind-direction
		WindScale string `json:"windScale"` // 风力等级 https://dev.qweather.com/docs/resource/wind-info/#wind-direction
		WindSpeed string `json:"windSpeed"` // 风速 https://dev.qweather.com/docs/resource/wind-info/#wind-speed 公里/小时
		Humidity  string `json:"humidity"`  // 相对湿度 百分比
		Pop       string `json:"pop"`       // 降水概率 百分比数值，**可能为空**
		Precip    string `json:"precip"`    // 当前小时累计降水量 默认单位：毫米
		Pressure  string `json:"pressure"`  // 大气压强 默认单位：百帕
		Cloud     string `json:"cloud"`     // 云量，百分比数值，**可能为空**
		Dew       string `json:"dew"`       // 露点温度 **可能为空**
	} `json:"hourly"`
	Refer struct {
		Sources []string `json:"sources"` // 原始数据来源，**可能为空**
		License []string `json:"license"` // 数据许可或版权声明，**可能为空**
	} `json:"refer"`
}
