package minForecast

import (
	"os"
	"testing"

	"github.com/Equationzhao/qweather-go"
)

var (
	k        = os.Getenv("qweather_key")
	publicID = os.Getenv("qweather_public_id")
	key      = qweather.Credential{
		Key:      k,
		PublicID: publicID,
		Encrypt:  false,
	}
)

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
