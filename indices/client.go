// Package indices 天气指数
package indices

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
	EndPoint     = "https://api.qweather.com/v7/indices/"
	FreeEndPoint = "https://devapi.qweather.com/v7/indices/"
)

func url(isFreePlan bool, u ...string) string {
	if isFreePlan {
		return iutil.Url(FreeEndPoint, u...)
	}
	return iutil.Url(EndPoint, u...)
}

// Indices 天气指数预报
//
// 获取中国和全球城市天气生活指数预报数据。
//
// 中国天气生活指数：舒适度指数、洗车指数、穿衣指数、感冒指数、运动指数、旅游指数、紫外线指数、空气污染扩散条件指数、空调开启指数、过敏指数、太阳镜指数、化妆指数、晾晒指数、交通指数、钓鱼指数、防晒指数
// 海外天气生活指数：运动指数、洗车指数、紫外线指数、钓鱼指数
//
// 当天生活指数
//
// https://api.qweather.com/v7/indices/1d?[请求参数]
// 未来3天生活指数
//
// https://api.qweather.com/v7/indices/3d?[请求参数]
//
// 请求参数说明:
//
// location(必选)
//
//	需要查询地区的LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)，LocationID可通过城市搜索服务(https://dev.qweather.com/docs/api/geoapi/)获取。例如 location=101010100 或 location=116.41,39.92
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// type(必选)
//
//	生活指数的类型ID，包括洗车指数、穿衣指数、钓鱼指数等。可以一次性获取多个类型的生活指数，多个类型用英文,分割。例如type=3,5。具体生活指数的ID和等级参考天气指数信息 https://dev.qweather.com/docs/resource/indices-info/ 各项生活指数并非适用于所有城市。
//
// lang
//
//	多语言设置，本数据仅支持中文和英文，可选值是lang=zh 和 lang=en
//
// 函数参数说明:
//
//	para 为请求参数
//	key 为用户认证key
//	count 为天数
//	isFreePlan 为是否为免费用户, 若是，则将上述API Host更改为devapi.qweather.com。参考免费订阅可用的数据(https://dev.qweather.com/docs/finance/subscription/#comparison)。
func Indices(para *Para, key qweather.Credential, count uint8, isFreePlan bool, client qweather.Client) (*Response, error) {
	request, err := IndicesRequest(para, key, count, isFreePlan)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	get, err := util.Get(request, client)
	var resp Response
	err = json.Unmarshal(get, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// IndicesWithRequiredParam 天气指数预报
// para 为其余参数，可以为 nil
// 详见 Indices
func IndicesWithRequiredParam(location, Type string, para *Para, key qweather.Credential, count uint8, isFreePlan bool, client qweather.Client) (*Response, error) {
	if para == nil {
		para = &Para{
			Location: location,
			Type:     Type,
		}
	} else {
		para.Location = location
		para.Type = Type
	}
	return Indices(para, key, count, isFreePlan, client)
}

func IndicesRequest(para *Para, key qweather.Credential, count uint8, isFreePlan bool) (*http.Request, error) {
	r, err := util.Request(url(isFreePlan, strconv.Itoa(int(count))+"d"), func(req *http.Request) {
		q := nurl.Values{}
		q.Add("location", para.Location)
		q.Add("lang", para.Lang)
		q.Add("type", para.Type)
		if key.Encrypt {
			qweather.AddSignature(key.PublicID, key.Key, q)
		} else {
			q.Add("key", key.Key)
		}
		req.URL.RawQuery = q.Encode()
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}

// IndicesRequestWithRequiredParam 天气指数预报
// para 为其余参数，可以为 nil
// 详见 IndicesRequest
func IndicesRequestWithRequiredParam(location, Type string, para *Para, key qweather.Credential, count uint8, isFreePlan bool) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Location: location,
			Type:     Type,
		}
	} else {
		para.Location = location
		para.Type = Type
	}
	return IndicesRequest(para, key, count, isFreePlan)
}

func Day1(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*Response, error) {
	return Indices(para, key, 1, isFreePlan, client)
}

func Day3(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*Response, error) {
	return Indices(para, key, 3, isFreePlan, client)
}
