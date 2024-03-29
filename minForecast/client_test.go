package minForecast

import (
	"testing"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/util"
)

var key = *util.Credential("qweather_key", "qweather_public_id").SetEncrypt()

func TestMinPrecipitation(t *testing.T) {
	para := &Para{
		Location: "116.41,39.92",
		Lang:     "zh",
	}
	resp, err := MinPrecipitation(para, key, qweather.FreePlan, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("return code is not 200")
	}
}
