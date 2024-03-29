package geo

import (
	"testing"

	itest "github.com/Equationzhao/qweather-go/internal/test"
	"github.com/Equationzhao/qweather-go/lang"
	"github.com/Equationzhao/qweather-go/util"
)

var key = *util.Credential("qweather_key", "qweather_public_id").SetEncrypt()

func TestGet(t *testing.T) {
	para := &Para{
		Location: "深圳",
		Adm:      "深圳",
		Lang:     lang.ZHCN,
		Number:   20,
		Range:    "cn",
	}
	{
		resp, err := SearchCity(para, key, &itest.NoProxyClient)
		if err != nil {
			t.Error(err)
		}
		t.Log(resp)
		if resp.Code != "200" {
			t.Error("return code is not 200")
		}
	}
	{
		resp, err := SearchCityWithRequiredParam("深圳", key, para, &itest.NoProxyClient)
		if err != nil {
			t.Error(err)
		}
		t.Log(resp)
		if resp.Code != "200" {
			t.Error("return code is not 200")
		}
	}
	// {
	// 	resp, err := SearchCityWithRequiredParam("深圳", key, nil, &itest.NoProxyClient)
	// 	if err != nil {
	// 		t.Error(err)
	// 	}
	// 	t.Log(resp)
	// 	if resp.Code != "200" {
	// 		t.Error("return code is not 200")
	// 	}
	// }
	{
		srwp, _ := SearchCityRequestWithRequiredParam("深圳", key, nil)
		_, err := itest.NoProxyClient.Do(srwp)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestHitCity(t *testing.T) {
	para := &Para{
		Number: 20,
		Range:  "cn",
	}
	resp, err := HitCity(para, key, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("return code is not 200")
	}
}

func TestPOI(t *testing.T) {
	para := &Para{
		Location: "北京",
		Lang:     lang.ZHCN,
		Number:   20,
		Type:     Scenic,
	}
	resp, err := POI(para, key, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("return code is not 200")
	}
}

func TestPOIRange(t *testing.T) {
	para := &Para{
		Location: "116.41,39.92",
		Lang:     lang.ZHCN,
		Number:   20,
		Type:     Scenic,
		Radius:   20,
	}
	resp, err := POIRange(para, key, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("return code is not 200")
	}
}
