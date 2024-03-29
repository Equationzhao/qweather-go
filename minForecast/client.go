// Package minForecast 分钟预报
package minForecast

import (
	"net/http"
	nurl "net/url"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	iutil "github.com/Equationzhao/qweather-go/internal/util"
	"github.com/Equationzhao/qweather-go/util"
)

const (
	EndPoint     = "https://api.qweather.com/v7/minutely/5m"
	FreeEndPoint = "https://devapi.qweather.com/v7/minutely/5m"
)

func url(isFreePlan bool, u ...string) string {
	if isFreePlan {
		return iutil.Url(FreeEndPoint, u...)
	}
	return iutil.Url(EndPoint, u...)
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
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func MinPrecipitation(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (
	*MinPrecipitationResponse, error,
) {
	request, err := MinPrecipitationRequest(para, key, isFreePlan)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	get, err := util.Get(request, client)
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

// MinPrecipitationWithRequiredParam 分钟级降水
// para 为其余参数，可以为 nil
// 详见 MinPrecipitation
func MinPrecipitationWithRequiredParam(location string, key qweather.Credential, para *Para, isFreePlan bool, client qweather.Client) (
	*MinPrecipitationResponse, error,
) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return MinPrecipitation(para, key, isFreePlan, client)
}

// MinPrecipitationRequest 分钟级降水
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
func MinPrecipitationRequest(para *Para, key qweather.Credential, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(
		url(isFreePlan), func(r *http.Request) {
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

// MinPrecipitationRequestWithRequiredParam 分钟级降水
// para 为其余参数，可以为 nil
// 详见 MinPrecipitationRequest
func MinPrecipitationRequestWithRequiredParam(location string, key qweather.Credential, para *Para, isFreePlan bool) (
	*http.Request, error,
) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return MinPrecipitationRequest(para, key, isFreePlan)
}
