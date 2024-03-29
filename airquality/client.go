// Package airqualityBeta 和风天气 空气质量(beta)
// 推荐阅读空气质量信息文档，以便了解指数类型、污染物、支持的国家等信息。
// 全新的实时空气质量目前处于beta阶段，标准订阅暂不收取费用。数据内容在正式发布后可能会有所不同。查看当前版本的发行公告。 https://blog.qweather.com/announce/aqi-v1-1-released/
// 推荐阅读空气质量信息文档，以便了解指数类型、污染物、支持的国家等信息。 https://dev.qweather.com/docs/resource/air-info/
// 全新的监测站数据目前处于beta阶段，标准订阅暂不收取费用。数据内容在正式发布后可能会有所不同。
// 监测站数据是实验性数据，仅供参考，可能受到各种因素的影响，我们无法确保该数据的可用性，请优先使用空气质量指数数据。
// 目前测试使用数字签名方式进行认证的情况下会遇到 401 认证失败的情况
package airqualityBeta

import (
	"net/http"
	nurl "net/url"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	iutil "github.com/Equationzhao/qweather-go/internal/util"
	"github.com/Equationzhao/qweather-go/util"
)

const (
	EndPoint     = "https://api.qweather.com/airquality/v1/"
	FreeEndPoint = "https://devapi.qweather.com/airquality/v1/"
)

func url(isFreePlan bool, u ...string) string {
	if isFreePlan {
		return iutil.Url(FreeEndPoint, u...)
	}
	return iutil.Url(EndPoint, u...)
}

func urlNow(isFreePlan bool, u ...string) string {
	us := []string{"now"}
	return url(isFreePlan, append(us, u...)...)
}

func urlStation(isFreePlan bool, u ...string) string {
	us := []string{"station"}
	return url(isFreePlan, append(us, u...)...)
}

func NowRequest(para *NowPara, key qweather.Credential, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(
		urlNow(isFreePlan, para.LocationID), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("lang", para.Lang)
			if para.Pollutant {
				q.Add("pollutant", "true")
			}
			if para.Station {
				q.Add("station", "true")
			}
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

func Now(para *NowPara, key qweather.Credential, isFreePlan bool, client qweather.Client) (*NowAirQualityResponse, error) {
	request, err := NowRequest(para, key, isFreePlan)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	get, err := util.Get(request, client)
	if err != nil {
		return nil, err
	}
	var res NowAirQualityResponse
	err = json.Unmarshal(get, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func NowWithRequiredParam(locationID string, key qweather.Credential, para *NowPara, isFreePlan bool, client qweather.Client) (*NowAirQualityResponse, error) {
	if para == nil {
		para = &NowPara{}
	}
	para.LocationID = locationID
	return Now(para, key, isFreePlan, client)
}

func StationRequest(para *StationPara, key qweather.Credential, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(
		urlStation(isFreePlan, para.LocationID), func(r *http.Request) {
			q := nurl.Values{}
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

func Station(para *StationPara, key qweather.Credential, isFreePlan bool, client qweather.Client) (*NowAirQualityResponse, error) {
	request, err := StationRequest(para, key, isFreePlan)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	get, err := util.Get(request, client)
	if err != nil {
		return nil, err
	}
	var res NowAirQualityResponse
	err = json.Unmarshal(get, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func StationWithRequiredParam(locationID string, key qweather.Credential, para *StationPara, isFreePlan bool, client qweather.Client) (*NowAirQualityResponse, error) {
	if para == nil {
		para = &StationPara{}
	}
	para.LocationID = locationID
	return Station(para, key, isFreePlan, client)
}
