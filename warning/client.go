package warning

import (
	"net/http"
	nurl "net/url"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	"github.com/Equationzhao/qweather-go/util"
)

const (
	EndPoint     = "https://api.qweather.com/v7/warning/"
	FreeEndPoint = "https://devapi.qweather.com/v7/warning/"
)

func url(isFreePlan bool, u ...string) string {
	if isFreePlan {
		return util.Url(FreeEndPoint, u...)
	}
	return util.Url(EndPoint, u...)
}

// RealTime 实时灾害预警
//
// GET https://api.qweather.com/v7/warning/now?[请求参数]
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
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	isFreePlan 为是否是免费用户
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func RealTime(para *Para, key qweather.Credential, isFreePlan bool, client *http.Client) (*RealTimeResponse, error) {
	request, err := RealTimeRequest(para, key, isFreePlan)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = http.DefaultClient
	}
	get, err := util.Get(request, client)
	if err != nil {
		return nil, err
	}
	var resp RealTimeResponse
	err = json.Unmarshal(get, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RealTimeWithRequiredParam 实时灾害预警
// para 为其余参数，可以为 nil
// 详见 RealTime
func RealTimeWithRequiredParam(location string, para *Para, key qweather.Credential, isFreePlan bool, client *http.Client) (*RealTimeResponse, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return RealTime(para, key, isFreePlan, client)
}

// RealTimeRequest 实时灾害预警
//
// GET https://api.qweather.com/v7/warning/now?[请求参数]
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
func RealTimeRequest(para *Para, key qweather.Credential, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(
		url(isFreePlan, "now"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("location", para.Location)
			q.Add("lang", para.Lang)
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

// RealTimeRequestWithRequiredParam 实时灾害预警
// para 为其余参数，可以为 nil
// 详见 RealTimeRequest
func RealTimeRequestWithRequiredParam(location string, para *Para, key qweather.Credential, isFreePlan bool) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return RealTimeRequest(para, key, isFreePlan)
}

// CityList 天气预警城市列表
// 获取指定国家或地区当前正在发生天气灾害预警的城市列表，根据这些城市列表再查询对应城市的天气灾害预警。
//
// > 注意：目前天气预警城市列表仅适用于获取中国（包括港澳台）城市列表。其他国家和地区，请使用天气灾害预警。
//
// > 提示：关于更多天气预警数据的说明，请参考实用资料-预警信息。
//
// GET https://api.qweather.com/v7/warning/list?[请求参数]
//
// 请求参数说明：
//
// range(必选)
//
//	选择指定的国家或地区，使用ISO 3166(https://dev.qweather.com/docs/resource/glossary/#iso-3166)格式。例如range=cn或range=hk。目前该功能仅支持中国（包括港澳台）地区的城市列表，其他国家和地区请使用请使用[天气灾害预警]单独获取。
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	isFreePlan 为是否是免费用户
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func CityList(para *Para, key qweather.Credential, isFreePlan bool, client *http.Client) (*CityListResponse, error) {
	request, err := CityListRequest(para, key, isFreePlan)
	if err != nil {
		return nil, err
	}
	if client == nil {
		client = http.DefaultClient
	}
	get, err := util.Get(request, client)
	if err != nil {
		return nil, err
	}
	var resp CityListResponse
	err = json.Unmarshal(get, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CityListWithRequiredParam 天气预警城市列表
// para 为其余参数，可以为 nil
// 详见 RealTimeRequest
func CityListWithRequiredParam(location string, para *Para, key qweather.Credential, isFreePlan bool, client *http.Client) (*CityListResponse, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return CityList(para, key, isFreePlan, client)
}

// CityListRequest 天气预警城市列表
// 获取指定国家或地区当前正在发生天气灾害预警的城市列表，根据这些城市列表再查询对应城市的天气灾害预警。
//
// > 注意：目前天气预警城市列表仅适用于获取中国（包括港澳台）城市列表。其他国家和地区，请使用天气灾害预警。
//
// > 提示：关于更多天气预警数据的说明，请参考实用资料-预警信息。
//
// GET https://api.qweather.com/v7/warning/list?[请求参数]
//
// 请求参数说明：
//
// range(必选)
//
//	选择指定的国家或地区，使用ISO 3166(https://dev.qweather.com/docs/resource/glossary/#iso-3166)格式。例如range=cn或range=hk。目前该功能仅支持中国（包括港澳台）地区的城市列表，其他国家和地区请使用请使用[天气灾害预警]单独获取。
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	isFreePlan 为是否是免费用户
func CityListRequest(para *Para, key qweather.Credential, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(
		url(isFreePlan, "list"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("range", para.Range)
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

// CityListRequestWithRequiredParam 天气预警城市列表
// para 为其余参数，可以为 nil
// 详见 CityListRequest
func CityListRequestWithRequiredParam(Range string, para *Para, key qweather.Credential, isFreePlan bool) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Range: Range,
		}
	} else {
		para.Range = Range
	}
	return CityListRequest(para, key, isFreePlan)
}
