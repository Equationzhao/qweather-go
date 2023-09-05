package gridWeather

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	itest "github.com/Equationzhao/qweather-go/internal/test"
)

var (
	k        = os.Getenv("qweather_key")
	publicID = os.Getenv("qweather_public_id")
	key      = qweather.Credential{
		Key:      k,
		PublicID: publicID,
		Encrypt:  true,
	}
)

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
		Lang:     "zh",
		Unit:     qweather.METRIC,
	}
	response, err := RealTime(para, key, true, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := RealTimeRequest(para, key, true)
	if err != nil {
		return
	}
	helper(t, request, &RealTimeResponse{})
}

func TestHourly(t *testing.T) {
	para := &Para{
		Location: "116.41,39.92",
		Lang:     "zh",
		Unit:     qweather.METRIC,
	}
	response, err := Hourly(para, key, 24, true, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := HourlyRequest(para, key, 24, true)
	if err != nil {
		return
	}
	helper(t, request, &HourlyResponse{})

	response, err = Hour24(para, key, true, &itest.NoProxyClient)
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
		Lang:     "zh",
		Unit:     qweather.METRIC,
	}
	response, err := Daily(para, key, 3, true, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := DailyRequest(para, key, 3, true)
	if err != nil {
		t.Fatal(err)
	}
	helper(t, request, &DailyResponse{})

	response, err = Day3(para, key, true, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}
}
