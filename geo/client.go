// Package geo 和风天气GeoAPI
package geo

import (
	"net/http"
	nurl "net/url"
	"strconv"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	iutil "github.com/Equationzhao/qweather-go/internal/util"
	"github.com/Equationzhao/qweather-go/util"
)

const EndPoint = "https://geoapi.qweather.com/v2/"

func url(u ...string) string {
	return iutil.Url(EndPoint, u...)
}

// SearchCity 城市搜索
//
// 城市搜索API提供全球地理位位置、全球城市搜索服务，支持经纬度坐标反查、多语言、模糊搜索等功能。
//
// 天气数据是基于地理位置的数据，因此获取天气之前需要先知道具体的位置信息。使用城市搜索，可获取到该城市的基本信息，包括城市的Location ID（你需要这个ID去查询天气），多语言名称、经纬度、时区、海拔、Rank值、归属上级行政区域、所在行政区域等。
//
// 另外，城市搜索也可以帮助你在你的APP中实现模糊搜索，用户只需要输入1-2个字即可获得结果。
//
// GET https://geoapi.qweather.com/v2/city/lookup?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的名称，支持文字、以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)、LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或Adcode(https://dev.qweather.com/docs/resource/glossary/#adcode)（仅限中国城市）。例如 location=北京 或 location=116.41,39.92
//	> 模糊搜索，当location传递的为文字时，支持模糊搜索，即用户可以只输入城市名称一部分进行搜索，最少一个汉字或2个字符，结果将按照相关性和Rank值(https://dev.qweather.com/docs/resource/glossary/#rank)进行排列，便于开发或用户进行选择他们需要查看哪个城市的天气。例如location=bei，将返回与bei相关性最强的若干结果，包括黎巴嫩的贝鲁特和中国的北京市
//	> 重名，当location传递的为文字时，可能会出现重名的城市，例如陕西省西安市、吉林省辽源市下辖的西安区和黑龙江省牡丹江市下辖的西安区，此时会根据Rank值(https://dev.qweather.com/docs/resource/glossary/#rank)排序返回所有结果。在这种情况下，可以通过adm参数的方式进一步确定需要查询的城市或地区，例如location=西安&adm=黑龙江
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// adm
//
//	城市的上级行政区划，可设定只在某个行政区划范围内进行搜索，用于排除重名城市或对结果进行过滤。例如 adm=beijing
//	> 如请求参数为location=chaoyang&adm=beijing时，返回的结果只包括北京市的朝阳区，而不包括辽宁省的朝阳市
//
//	> 如请求参数仅为location=chaoyang时，返回的结果包括北京市的朝阳区、辽宁省的朝阳市以及长春市的朝阳区
//
// range
//
//	搜索范围，可设定只在某个国家或地区范围内进行搜索，国家和地区名称需使用ISO 3166 所定义的国家代码(https://dev.qweather.com/docs/resource/glossary/#iso-3166)。如果不设置此参数，搜索范围将在所有城市。例如 range=cn
//
// number
//
//	返回结果的数量，取值范围1-20，默认返回10个结果。
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func SearchCity(para *Para, key qweather.Credential, client qweather.Client) (*SearchResponse, error) {
	req, err := SearchCityRequest(para, key)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	get, err := util.Get(req, client)
	if err != nil {
		return nil, err
	}
	var response SearchResponse
	err = json.Unmarshal(get, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// SearchCityWithRequiredParam 城市搜索
// para 为其余参数，可以为 nil
// 详见 SearchCity
func SearchCityWithRequiredParam(location string, key qweather.Credential, para *Para, client qweather.Client) (*SearchResponse, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return SearchCity(para, key, client)
}

// SearchCityRequest 城市搜索
//
// 城市搜索提供全球地理位位置、全球城市搜索服务，支持经纬度坐标反查、多语言、模糊搜索等功能。
//
// 天气数据是基于地理位置的数据，因此获取天气之前需要先知道具体的位置信息。使用城市搜索，可获取到该城市的基本信息，包括城市的Location ID（你需要这个ID去查询天气），多语言名称、经纬度、时区、海拔、Rank值、归属上级行政区域、所在行政区域等。
//
// 另外，城市搜索也可以帮助你在你的APP中实现模糊搜索，用户只需要输入1-2个字即可获得结果。
//
// GET https://geoapi.qweather.com/v2/city/lookup?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的名称，支持文字、以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)、LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或Adcode(https://dev.qweather.com/docs/resource/glossary/#adcode)（仅限中国城市）。例如 location=北京 或 location=116.41,39.92
//	> 模糊搜索，当location传递的为文字时，支持模糊搜索，即用户可以只输入城市名称一部分进行搜索，最少一个汉字或2个字符，结果将按照相关性和Rank值(https://dev.qweather.com/docs/resource/glossary/#rank)进行排列，便于开发或用户进行选择他们需要查看哪个城市的天气。例如location=bei，将返回与bei相关性最强的若干结果，包括黎巴嫩的贝鲁特和中国的北京市
//	> 重名，当location传递的为文字时，可能会出现重名的城市，例如陕西省西安市、吉林省辽源市下辖的西安区和黑龙江省牡丹江市下辖的西安区，此时会根据Rank值(https://dev.qweather.com/docs/resource/glossary/#rank)排序返回所有结果。在这种情况下，可以通过adm参数的方式进一步确定需要查询的城市或地区，例如location=西安&adm=黑龙江
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// adm
//
//	城市的上级行政区划，可设定只在某个行政区划范围内进行搜索，用于排除重名城市或对结果进行过滤。例如 adm=beijing
//	> 如请求参数为location=chaoyang&adm=beijing时，返回的结果只包括北京市的朝阳区，而不包括辽宁省的朝阳市
//
//	> 如请求参数仅为location=chaoyang时，返回的结果包括北京市的朝阳区、辽宁省的朝阳市以及长春市的朝阳区
//
// range
//
//	搜索范围，可设定只在某个国家或地区范围内进行搜索，国家和地区名称需使用ISO 3166 所定义的国家代码(https://dev.qweather.com/docs/resource/glossary/#iso-3166)。如果不设置此参数，搜索范围将在所有城市。例如 range=cn
//
// number
//
//	返回结果的数量，取值范围1-20，默认返回10个结果。
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
func SearchCityRequest(para *Para, key qweather.Credential) (*http.Request, error) {
	r, err := util.Request(
		url("city", "lookup"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("location", para.Location)
			q.Add("adm", para.Adm)
			q.Add("lang", para.Lang)
			q.Add("number", strconv.Itoa(int(para.Number)))
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

// SearchCityRequestWithRequiredParam 城市搜索
// para 为其余参数，可以为 nil
// 详见 SearchCityRequest
func SearchCityRequestWithRequiredParam(location string, key qweather.Credential, para *Para) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Location: location,
		}
	} else {
		para.Location = location
	}
	return SearchCityRequest(para, key)
}

// HitCity 热门城市查询
//
// 获取全球各国热门城市列表。
//
// Get https://geoapi.qweather.com/v2/city/top?[请求参数]
//
// 请求参数说明：
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// range
//
//	搜索范围，可设定只在某个国家或地区范围内进行搜索，国家和地区名称需使用ISO 3166 所定义的国家代码(https://dev.qweather.com/docs/resource/glossary/#iso-3166)。如果不设置此参数，搜索范围将在所有城市。例如 range=cn
//
// number
//
//	返回结果的数量，取值范围1-20，默认返回10个结果。
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func HitCity(para *Para, key qweather.Credential, client qweather.Client) (*HitResponse, error) {
	req, err := HitCityRequest(para, key)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	get, err := util.Get(req, client)
	if err != nil {
		return nil, err
	}
	var response HitResponse
	err = json.Unmarshal(get, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// HitCityRequest 热门城市查询
//
// 获取全球各国热门城市列表。
//
// Get https://geoapi.qweather.com/v2/city/top?[请求参数]
//
// 请求参数说明：
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// range
//
//	搜索范围，可设定只在某个国家或地区范围内进行搜索，国家和地区名称需使用ISO 3166 所定义的国家代码(https://dev.qweather.com/docs/resource/glossary/#iso-3166)。如果不设置此参数，搜索范围将在所有城市。例如 range=cn
//
// number
//
//	返回结果的数量，取值范围1-20，默认返回10个结果。
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
func HitCityRequest(para *Para, key qweather.Credential) (*http.Request, error) {
	r, err := util.Request(
		url("city", "top"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("location", para.Location)
			q.Add("adm", para.Adm)
			q.Add("lang", para.Lang)
			q.Add("number", strconv.Itoa(int(para.Number)))
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

// POI POI搜索
//
// 使用关键字和坐标查询POI信息（景点、火车站、飞机场、港口等）
//
// GET https://geoapi.qweather.com/v2/poi/lookup?[请求参数]
//
// 请求参数说明：
//
// type(必选)
//
//	POI类型，可选择搜索某一类型的POI。
//	    scenic 景点
//	    CSTA 潮流站点
//	    TSTA 潮汐站点
//
// location(必选)
//
//	需要查询地区的名称，支持文字、以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)、LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或Adcode(https://dev.qweather.com/docs/resource/glossary/#adcode)（仅限中国城市）。例如 location=北京 或 location=116.41,39.92
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// city
//
//	选择POI所在城市，可设定只搜索在特定城市内的POI信息。城市名称可以是文字或城市的LocationID。**默认不限制特定城市**。
//	> 城市名称为精准匹配，建议使用LocaitonID，如文字无法匹配，则数据返回为空。
//
// number
//
//	返回结果的数量，取值范围1-20，默认返回10个结果。
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func POI(para *Para, key qweather.Credential, client qweather.Client) (*POIResponse, error) {
	req, err := POIRequest(para, key)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	var response POIResponse
	get, err := util.Get(req, client)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(get, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// POIWithRequiredParam POI搜索
// para 为其余参数，可以为 nil
// 详见 POI
func POIWithRequiredParam(location string, key qweather.Credential, t Type, para *Para, client qweather.Client) (*POIResponse, error) {
	if para == nil {
		para = &Para{
			Location: location,
			Type:     t,
		}
	} else {
		para.Location = location
	}
	return POI(para, key, client)
}

// POIRequest POI搜索
//
// 使用关键字和坐标查询POI信息（景点、火车站、飞机场、港口等）
//
// GET https://geoapi.qweather.com/v2/poi/lookup?[请求参数]
//
// 请求参数说明：
//
// type(必选)
//
//	POI类型，可选择搜索某一类型的POI。
//	    scenic 景点
//	    CSTA 潮流站点
//	    TSTA 潮汐站点
//
// location(必选)
//
//	需要查询地区的名称，支持文字、以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)、LocationID(https://dev.qweather.com/docs/resource/glossary/#locationid)或Adcode(https://dev.qweather.com/docs/resource/glossary/#adcode)（仅限中国城市）。例如 location=北京 或 location=116.41,39.92
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// city
//
//	选择POI所在城市，可设定只搜索在特定城市内的POI信息。城市名称可以是文字或城市的LocationID。**默认不限制特定城市**。
//	> 城市名称为精准匹配，建议使用LocaitonID，如文字无法匹配，则数据返回为空。
//
// number
//
//	返回结果的数量，取值范围1-20，默认返回10个结果。
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
func POIRequest(para *Para, key qweather.Credential) (*http.Request, error) {
	r, err := util.Request(
		url("poi", "lookup"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("type", para.Type.String())
			q.Add("location", para.Location)
			q.Add("city", para.City)
			q.Add("number", strconv.Itoa(int(para.Number)))
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

// POIRequestWithRequiredParam POI搜索
// para 为其余参数，可以为 nil
// 详见 POIRequest
func POIRequestWithRequiredParam(location string, key qweather.Credential, t Type, para *Para) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Location: location,
			Type:     t,
		}
	} else {
		para.Location = location
	}
	return POIRequest(para, key)
}

// POIRange POI范围搜索
//
// 提供指定区域范围内查询所有POI信息。
//
// GET https://geoapi.qweather.com/v2/poi/range?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)。例如 location=116.41,39.92
//
// type(必选)
//
//	POI类型，可选择搜索某一类型的POI。
//	    scenic 景点
//	    CSTA 潮流站点
//	    TSTA 潮汐站点
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// radius(必选)
//
//	搜索范围，可设置搜索半径，取值范围1-50，单位：公里。**默认5公里**。
//
// number
//
//	返回结果的数量，取值范围1-20，默认返回10个结果。
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
//	client 为自定义的 Client, 若为nil, 则使用http.DefaultClient
func POIRange(para *Para, key qweather.Credential, client qweather.Client) (*POIResponse, error) {
	req, err := POIRangeRequest(para, key)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	var response POIResponse
	get, err := util.Get(req, client)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(get, &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// POIRangeWithRequiredParam POI范围搜索
// para 为其余参数，可以为 nil
// 详见 POIRange
func POIRangeWithRequiredParam(location string, key qweather.Credential, t Type, radius uint16, para *Para, client qweather.Client) (*POIResponse, error) {
	if para == nil {
		para = &Para{
			Location: location,
			Type:     t,
			Radius:   radius,
		}
	} else {
		para.Location = location
	}
	return POIRange(para, key, client)
}

// POIRangeRequest POI范围搜索
//
// 提供指定区域范围内查询所有POI信息。
//
// GET https://geoapi.qweather.com/v2/poi/range?[请求参数]
//
// 请求参数说明：
//
// location(必选)
//
//	需要查询地区的以英文逗号分隔的经度,纬度坐标（十进制，最多支持小数点后两位）(https://dev.qweather.com/docs/resource/glossary/#coordinate)。例如 location=116.41,39.92
//
// type(必选)
//
//	POI类型，可选择搜索某一类型的POI。
//	    scenic 景点
//	    CSTA 潮流站点
//	    TSTA 潮汐站点
//
// key(必选)
//
//	用户认证key，请参考如何获取你的KEY(https://dev.qweather.com/docs/configuration/project-and-key/)。支持数字签名(https://dev.qweather.com/docs/resource/signature-auth/)方式进行认证。例如 key=123456789ABC
//
// radius(必选)
//
//	搜索范围，可设置搜索半径，取值范围1-50，单位：公里。**默认5公里**。
//
// number
//
//	返回结果的数量，取值范围1-20，默认返回10个结果。
//
// lang
//
//	多语言设置，更多语言可选值参考语言代码。当数据不匹配你设置的语言时，将返回英文或其本地语言结果。
//
// 函数参数说明
//
//	para 为请求参数
//	key 为用户认证key
func POIRangeRequest(para *Para, key qweather.Credential) (*http.Request, error) {
	r, err := util.Request(
		url("poi", "range"), func(r *http.Request) {
			q := nurl.Values{}
			q.Add("type", para.Type.String())
			q.Add("location", para.Location)
			q.Add("radius", strconv.Itoa(int(para.Radius)))
			q.Add("number", strconv.Itoa(int(para.Number)))
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

// POIRangeRequestWithRequiredParam POI范围搜索
// para 为其余参数，可以为 nil
// 详见 POIRangeRequest
func POIRangeRequestWithRequiredParam(location string, key qweather.Credential, t Type, radius uint16, para *Para) (*http.Request, error) {
	if para == nil {
		para = &Para{
			Location: location,
			Type:     t,
			Radius:   radius,
		}
	} else {
		para.Location = location
	}
	return POIRangeRequest(para, key)
}
