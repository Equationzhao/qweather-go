package minForecast

import (
	"net/http"

	"qweather/json"
	"qweather/util"
)

const EndPoint = "https://api.qweather.com/v7/minutely/5m"
const FreeEndPoint = "https://devapi.qweather.com/v7/minutely/5m"

func url(isFreePlan bool, u ...string) string {
	if isFreePlan {
		return util.Url(FreeEndPoint, u...)
	}
	return util.Url(EndPoint, u...)
}

// MinPrecipitation 分钟级降水
//
// 分钟级降水（临近预报）支持中国1公里精度的未来2小时每5分钟降雨预报数据。
//
// GET https://api.qweather.com/v7/minutely/5m?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)。例如 location=116.41,39.92
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// lang
//
//	多语言设置，本数据仅支持中文和英文，可选值是lang=zh 和 lang=en
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	isFreePlan 为是否是免费用户
func MinPrecipitation(para *Para, key string, isFreePlan bool) (*MinPrecipitationResponse, error) {
	get, err := util.Get(
		url(isFreePlan), func(r *http.Request) {
			q := r.URL.Query()
			q.Add("key", key)
			q.Add("location", para.Location)
			q.Add("lang", para.Lang)
			r.URL.RawQuery = q.Encode()
		},
	)
	if err != nil {
		return nil, err
	}

	var response MinPrecipitationResponse
	err = json.Unmarshal(get, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
