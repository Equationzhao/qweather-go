package geo

import (
	"os"
	"testing"
)

var key = os.Getenv("qweather_key")

func TestGet(t *testing.T) {

	para := &Para{
		Location: "深圳",
		Adm:      "深圳",
		Lang:     "zh",
		Number:   20,
		Range:    "cn",
	}
	resp, err := SearchCity(para, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("code not 200")
	}
}

func TestHitCity(t *testing.T) {

	para := &Para{
		Number: 20,
		Range:  "cn",
	}
	resp, err := HitCity(para, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("code not 200")
	}
}

func TestPOI(t *testing.T) {

	para := &Para{
		Location: "北京",
		Lang:     "zh",
		Number:   20,
		Type:     Scenic,
	}
	resp, err := POI(para, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("code not 200")
	}
}

func TestPOIRange(t *testing.T) {

	para := &Para{
		Location: "北京",
		Lang:     "zh",
		Number:   20,
		Type:     Scenic,
		Radius:   20,
	}
	resp, err := POIRange(para, key)
	if err != nil {
		t.Error(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("code not 200")
	}
}
