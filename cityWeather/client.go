// Package cityWeather 城市天气
package cityWeather

import (
	"net/http"
	nurl "net/url"
	"strconv"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	iutil "github.com/Equationzhao/qweather-go/internal/util"

	"github.com/Equationzhao/qweather-go/util"
)

const (
	EndPoint     = "https://api.qweather.com/v7/weather/"
	FreeEndPoint = "https://devapi.qweather.com/v7/weather/"
)

func url(isFreePlan bool, u ...string) string {
	if isFreePlan {
		return iutil.Url(FreeEndPoint, u...)
	}
	return iutil.Url(EndPoint, u...)
}

// RealTime 实时天气
//
// 获取中国3000+市县区和海外20万个城市实时天气数据，包括实时温度、体感温度、风力风向、相对湿度、大气压强、降水量、能见度、露点温度、云量等。
// > 注意：实况数据均为近实时数据，相比真实的物理世界有5-20分钟的延迟，请根据实况数据中的obsTime确定数据对应的准确时间。
//
// GET https://api.qweather.com/v7/weather/now?[请求参数]
//
// 请求参数说明：
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// location(必选)
//
//	需要查询地区的LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)，LocationID可通过城市搜索服务(https://dev.qweather.com/docs/api/geoapi/)获取。例如 location=101010100 或 location=116.41,39.92
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码(https://dev.qweather.com/docs/resource/language/)。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// unit
//
//	数据单位设置，可选值包括unit=m（公制单位，默认）和unit=i（英制单位）。更多选项和说明参考度量衡单位(https://dev.qweather.com/docs/resource/unit)。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	isFreePlan 为是否为免费用户, 若是，则将上述API Host更改为devapi.qweather.com。参考免费订阅可用的数据(https://dev.qweather.com/docs/finance/subscription/#comparison)。
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func RealTime(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*RealTimeResponse, error) {
	req, err := RealTimeRequest(para, key, isFreePlan)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = http.DefaultClient
	}
	get, err := util.Get(req, client)
	if err != nil {
		return nil, err
	}
	var response RealTimeResponse
	err = json.Unmarshal(get, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// RealTimeWithRequiredParam 实时天气
// para 为其余参数，可以为 nil
// 详见 RealTime
func RealTimeWithRequiredParam(location string, key qweather.Credential, para *Para, isFreePlan bool, client qweather.Client) (*RealTimeResponse, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return RealTime(para, key, isFreePlan, client)
}

// RealTimeRequest 实时天气
//
// 获取中国3000+市县区和海外20万个城市实时天气数据，包括实时温度、体感温度、风力风向、相对湿度、大气压强、降水量、能见度、露点温度、云量等。
// > 注意：实况数据均为近实时数据，相比真实的物理世界有5-20分钟的延迟，请根据实况数据中的obsTime确定数据对应的准确时间。
//
// GET https://api.qweather.com/v7/weather/now?[请求参数]
//
// 请求参数说明：
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// location(必选)
//
//	需要查询地区的LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)，LocationID可通过城市搜索服务(https://dev.qweather.com/docs/api/geoapi/)获取。例如 location=101010100 或 location=116.41,39.92
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码(https://dev.qweather.com/docs/resource/language/)。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// unit
//
//	数据单位设置，可选值包括unit=m（公制单位，默认）和unit=i（英制单位）。更多选项和说明参考度量衡单位(https://dev.qweather.com/docs/resource/unit)。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	isFreePlan 为是否为免费用户, 若是，则将上述API Host更改为devapi.qweather.com。参考免费订阅可用的数据(https://dev.qweather.com/docs/finance/subscription/#comparison)。
func RealTimeRequest(para *Para, key qweather.Credential, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(
		url(isFreePlan, "now"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("location", para.Location)
			q.Add("lang", para.Lang)
			q.Add("unit", para.Unit.String())
			if key.Encrypt {
				qweather.AddSignature(key.PublicID, key.Key, q)
			} else {
				q.Add("key", key.Key)
			}
			r.URL.RawQuery = q.Encode()
		},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// RealTimeRequestWithRequiredParam 实时天气
// para 为其余参数，可以为 nil
// 详见 RealTimeRequest
func RealTimeRequestWithRequiredParam(location string, key qweather.Credential, para *Para, isFreePlan bool) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return RealTimeRequest(para, key, isFreePlan)
}

// Now
// alias for RealTime
var Now = RealTime

// Daily 每日天气预报
//
// 每日天气预报，提供全球城市未来3-30天天气预报，包括：日出日落、月升月落、最高最低温度、天气白天和夜间状况、风力、风速、风向、相对湿度、大气压强、降水量、露点温度、紫外线强度、能见度等。
//
// 3天预报
// GET https://api.qweather.com/v7/weather/3d?[请求参数]
//
// 7天预报
// GET https://api.qweather.com/v7/weather/7d?[请求参数]
//
// 10天预报 Paid plan only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/10d?[请求参数]
//
// 15天预报 Paid plan only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/15d?[请求参数]
//
// 30天预报 Paid plan only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/30d?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)，LocationID可通过城市搜索服务(https://dev.qweather.com/docs/api/geoapi/)获取。例如 location=101010100 或 location=116.41,39.92
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码(https://dev.qweather.com/docs/resource/language/)。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// unit
//
//	数据单位设置，可选值包括unit=m（公制单位，默认）和unit=i（英制单位）。更多选项和说明参考度量衡单位(https://dev.qweather.com/docs/resource/unit)。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	count 为天数
//	isFreePlan 为是否为免费用户, 若是，则将上述API Host更改为devapi.qweather.com。参考免费订阅可用的数据(https://dev.qweather.com/docs/finance/subscription/#comparison)。
//	client 为自定义的 Client, 若为nil, 则使用 http.DefaultClient
func Daily(para *Para, key qweather.Credential, count uint8, isFreePlan bool, client qweather.Client) (*DailyResponse, error) {
	req, err := DailyRequest(para, key, count, isFreePlan)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = http.DefaultClient
	}
	get, err := util.Get(req, client)
	if err != nil {
		return nil, err
	}
	var response DailyResponse
	err = json.Unmarshal(get, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Day3 三日天气预报
func Day3(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*DailyResponse, error) {
	return Daily(para, key, 3, isFreePlan, client)
}

// Day7 七日天气预报
func Day7(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*DailyResponse, error) {
	return Daily(para, key, 7, isFreePlan, client)
}

// Day10 十日天气预报
func Day10(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*DailyResponse, error) {
	return Daily(para, key, 10, isFreePlan, client)
}

// Day15 十五日天气预报
func Day15(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*DailyResponse, error) {
	return Daily(para, key, 15, isFreePlan, client)
}

// Day30 三十日天气预报
func Day30(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*DailyResponse, error) {
	return Daily(para, key, 30, isFreePlan, client)
}

// DailyWithRequiredParam 每日天气预报
// para 为其余参数，可以为 nil
// 详见 Daily
func DailyWithRequiredParam(location string, key qweather.Credential, count uint8, para *Para, isFreePlan bool, client qweather.Client) (*DailyResponse, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return Daily(para, key, count, isFreePlan, client)
}

// DailyRequest 每日天气预报
//
// 每日天气预报，提供全球城市未来3-30天天气预报，包括：日出日落、月升月落、最高最低温度、天气白天和夜间状况、风力、风速、风向、相对湿度、大气压强、降水量、露点温度、紫外线强度、能见度等。
//
// 3天预报
// GET https://api.qweather.com/v7/weather/3d?[请求参数]
//
// 7天预报
// GET https://api.qweather.com/v7/weather/7d?[请求参数]
//
// 10天预报 Paid plan only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/10d?[请求参数]
//
// 15天预报 Paid plan only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/15d?[请求参数]
//
// 30天预报 Paid plan only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/30d?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)，LocationID可通过城市搜索服务(https://dev.qweather.com/docs/api/geoapi/)获取。例如 location=101010100 或 location=116.41,39.92
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码(https://dev.qweather.com/docs/resource/language/)。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// unit
//
//	数据单位设置，可选值包括unit=m（公制单位，默认）和unit=i（英制单位）。更多选项和说明参考度量衡单位(https://dev.qweather.com/docs/resource/unit)。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	count 为天数
//	isFreePlan 为是否为免费用户, 若是，则将上述API Host更改为devapi.qweather.com。参考免费订阅可用的数据(https://dev.qweather.com/docs/finance/subscription/#comparison)。
func DailyRequest(para *Para, key qweather.Credential, count uint8, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(
		url(isFreePlan, strconv.Itoa(int(count))+"d"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("location", para.Location)
			q.Add("lang", para.Lang)
			q.Add("unit", para.Unit.String())
			if key.Encrypt {
				qweather.AddSignature(key.PublicID, key.Key, q)
			} else {
				q.Add("key", key.Key)
			}
			r.URL.RawQuery = q.Encode()
		},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// DailyRequestWithRequiredParam 每日天气预报
// para 为其余参数，可以为 nil
// 详见 DailyRequest
func DailyRequestWithRequiredParam(location string, key qweather.Credential, count uint8, para *Para, isFreePlan bool) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return DailyRequest(para, key, count, isFreePlan)
}

// Hourly 逐小时天气预报
//
// 逐小时天气预报，提供全球城市未来24-168小时逐小时天气预报，包括：温度、天气状况、风力、风速、风向、相对湿度、大气压强、降水概率、露点温度、云量。
//
// 逐小时预报（未来24小时）
// GET https://api.qweather.com/v7/weather/24h?[请求参数]
//
// 逐小时预报（未来72小时） Paid plan only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/72h?[请求参数]
//
// 逐小时预报（未来168小时） Paid plan only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/168h?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)，LocationID可通过城市搜索服务(https://dev.qweather.com/docs/api/geoapi/)获取。例如 location=101010100 或 location=116.41,39.92
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码(https://dev.qweather.com/docs/resource/language/)。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// unit
//
//	数据单位设置，可选值包括unit=m（公制单位，默认）和unit=i（英制单位）。更多选项和说明参考度量衡单位(https://dev.qweather.com/docs/resource/unit)。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	count 为小时数
//	isFreePlan 为是否为免费用户, 若是，则将上述API Host更改为devapi.qweather.com。参考免费订阅可用的数据(https://dev.qweather.com/docs/finance/subscription/#comparison)。
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func Hourly(para *Para, key qweather.Credential, count uint8, isFreePlan bool, client qweather.Client) (*HourlyResponse, error) {
	req, err := HourlyRequest(para, key, count, isFreePlan)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = http.DefaultClient
	}
	get, err := util.Get(req, client)
	if err != nil {
		return nil, err
	}

	var response HourlyResponse
	err = json.Unmarshal(get, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Hour24 24小时天气预报
func Hour24(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*HourlyResponse, error) {
	return Hourly(para, key, 24, isFreePlan, client)
}

// Hour72 72小时天气预报
// Paid plan only 付费订阅用户可用
func Hour72(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*HourlyResponse, error) {
	return Hourly(para, key, 72, isFreePlan, client)
}

// Hour168 168小时天气预报
// Paid plan only 付费订阅用户可用
func Hour168(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*HourlyResponse, error) {
	return Hourly(para, key, 168, isFreePlan, client)
}

// HourlyWithRequiredParam 逐小时天气预报
// para 为其余参数，可以为 nil
// 详见 Hourly
func HourlyWithRequiredParam(location string, key qweather.Credential, count uint8, para *Para, isFreePlan bool, client qweather.Client) (*HourlyResponse, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return Hourly(para, key, count, isFreePlan, client)
}

// HourlyRequest 逐小时天气预报
//
// 逐小时天气预报，提供全球城市未来24-168小时逐小时天气预报，包括：温度、天气状况、风力、风速、风向、相对湿度、大气压强、降水概率、露点温度、云量。
//
// 逐小时预报（未来24小时）
// GET https://api.qweather.com/v7/weather/24h?[请求参数]
//
// 逐小时预报（未来72小时） Paid plan only only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/72h?[请求参数]
//
// 逐小时预报（未来168小时） Paid plan only only 付费订阅用户可用
// GET https://api.qweather.com/v7/weather/168h?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)，LocationID可通过城市搜索服务(https://dev.qweather.com/docs/api/geoapi/)获取。例如 location=101010100 或 location=116.41,39.92
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码(https://dev.qweather.com/docs/resource/language/)。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// unit
//
//	数据单位设置，可选值包括unit=m（公制单位，默认）和unit=i（英制单位）。更多选项和说明参考度量衡单位(https://dev.qweather.com/docs/resource/unit)。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	count 为小时数
//	isFreePlan 为是否为免费用户, 若是，则将上述API Host更改为devapi.qweather.com。参考免费订阅可用的数据(https://dev.qweather.com/docs/finance/subscription/#comparison)。
func HourlyRequest(para *Para, key qweather.Credential, count uint8, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(
		url(isFreePlan, strconv.Itoa(int(count))+"h"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("location", para.Location)
			q.Add("lang", para.Lang)
			q.Add("unit", para.Unit.String())
			if key.Encrypt {
				qweather.AddSignature(key.PublicID, key.Key, q)
			} else {
				q.Add("key", key.Key)
			}
			r.URL.RawQuery = q.Encode()
		},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// HourlyRequestWithRequiredParam 逐小时天气预报
// para 为其余参数，可以为 nil`
// 详见 HourlyRequest
func HourlyRequestWithRequiredParam(location string, key qweather.Credential, count uint8, para *Para, isFreePlan bool) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return HourlyRequest(para, key, count, isFreePlan)
}
