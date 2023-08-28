package minForecast

import (
	"os"
	"testing"
)

var key = os.Getenv("qweather_key")

func TestMinPrecipitation(t *testing.T) {

	para := &Para{
		Location: "116.41,39.92",
		Lang:     "zh",
	}
	resp, err := MinPrecipitation(para, key, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("return code is not 200")
	}
}
