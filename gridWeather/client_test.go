package gridWeather

import (
	"io"
	"net/http"
	"testing"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	itest "github.com/Equationzhao/qweather-go/internal/test"
	"github.com/Equationzhao/qweather-go/lang"
	"github.com/Equationzhao/qweather-go/util"
)

var key = *util.Credential("qweather_key", "qweather_public_id").SetEncrypt()

func helper(t *testing.T, request *http.Request, m any) {
	t.Helper()
	response, err := itest.NoProxyClient.Do(request)
	if err != nil {
		t.Fatal(err)
	}
	defer response.Body.Close()
	all, err := io.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(all, m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(m)
}

func TestRealTime(t *testing.T) {
	para := &Para{
		Location: "116.41,39.92",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}
	response, err := RealTime(para, key, qweather.Free, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := RealTimeRequest(para, key, qweather.Free)
	if err != nil {
		return
	}
	helper(t, request, &RealTimeResponse{})
}

func TestHourly(t *testing.T) {
	para := &Para{
		Location: "116.41,39.92",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}
	response, err := Hourly(para, key, 24, qweather.Free, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := HourlyRequest(para, key, 24, qweather.Free)
	if err != nil {
		return
	}
	helper(t, request, &HourlyResponse{})

	response, err = Hour24(para, key, qweather.Free, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}
}

func TestDaily(t *testing.T) {
	para := &Para{
		Location: "116.41,39.92",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}
	response, err := Daily(para, key, 3, qweather.Free, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := DailyRequest(para, key, 3, qweather.Free)
	if err != nil {
		t.Fatal(err)
	}
	helper(t, request, &DailyResponse{})

	response, err = Day3(para, key, qweather.Free, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}
}
