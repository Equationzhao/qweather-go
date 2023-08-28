package cityWeather

import (
	"os"
	"testing"
)

var key = os.Getenv("qweather_key")

func TestRealTime(t *testing.T) {

	para := &Para{
		Location: "101010100",
		Lang:     "zh",
		Unit:     METRIC,
	}
	resp, err := RealTime(para, key, true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Error("return code is not 200")
	}
}

func TestHourly(t *testing.T) {

	para := &Para{
		Location: "101010100",
		Lang:     "zh",
		Unit:     METRIC,
	}

	// 24
	args := []uint8{24}
	for _, arg := range args {
		resp, err := Hourly(para, key, arg, true)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
		if resp.Code != "200" {
			t.Error("return code is not 200")
		}
	}

}

func TestDaily(t *testing.T) {

	para := &Para{
		Location: "101010100",
		Lang:     "zh",
		Unit:     METRIC,
	}

	// 3,7
	args := []uint8{3, 7}
	for _, arg := range args {
		resp, err := Daily(para, key, arg, true)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
		if resp.Code != "200" {
			t.Error("return code is not 200")
		}
	}

}
