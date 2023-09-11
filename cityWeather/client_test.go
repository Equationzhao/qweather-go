package cityWeather

import (
	"io"
	"net/http"
	"testing"
	"time"

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
		Location: "101010100",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}
	resp, err := RealTime(para, key, true, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Fatal("return code is not 200")
	}
	key.Encrypt = false

	request, err := RealTimeRequest(para, key, true)
	if err != nil {
		t.Fatal(err)
	}
	key.Encrypt = true

	helper(t, request, &RealTimeResponse{})
}

func TestHourly(t *testing.T) {
	para := &Para{
		Location: "101010100",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}
	t.Log(key.Key, key.PublicID)

	// 24
	args := []uint8{24}
	for _, arg := range args {
		resp, err := Hourly(para, key, arg, true, &itest.NoProxyClient)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
		if resp.Code != "200" {
			t.Fatal("return code is not 200")
		}
	}

	request, err := HourlyRequest(para, key, 24, true)
	if err != nil {
		t.Fatal(err)
	}
	helper(t, request, &HourlyResponse{})
}

func TestDaily(t *testing.T) {
	para := &Para{
		Location: "101010100",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}

	// 3,7
	args := []uint8{3, 7}
	for _, arg := range args {
		resp, err := Daily(para, key, arg, true, &itest.NoProxyClient)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(resp)
		if resp.Code != "200" {
			t.Fatal("return code is not 200")
		}
	}

	request, err := DailyRequest(para, key, 3, true)
	if err != nil {
		t.Fatal(err)
	}
	helper(t, request, &DailyResponse{})
}

func TestWithClientRequest(t *testing.T) {
	// proxy http://localhost:1087
	client := itest.HttpProxyClient(1087)
	client.Timeout = 10 * time.Second

	para := &Para{
		Location: "101010100",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}

	// 3,7
	args := []uint8{3, 7}
	for _, arg := range args {
		req, err := DailyRequest(para, key, arg, true)
		if err != nil {
			t.Fatal(err)
		}

		resp, err := util.Get(req, &client)
		if err != nil {
			t.Fatal(err)
		}
		t.Log(string(resp))
	}
}

func TestWithClient(t *testing.T) {
	// proxy http://localhost:1087
	client := itest.HttpProxyClient(1087)
	client.Timeout = 10 * time.Second

	para := &Para{
		Location: "101010100",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}

	// 3,7
	args := []uint8{3, 7}
	for _, arg := range args {
		resp, err := Daily(para, key, arg, true, &client)
		if err != nil {
			t.Error(err)
			continue
		}
		if resp.Code != "200" {
			t.Error("return code is not 200", resp)
		}
	}
}

func TestDays(t *testing.T) {
	fs := [...]func(para *Para, key qweather.Credential, isFreePlan bool, client qweather.Client) (*DailyResponse, error){
		Day3,
		Day7,
		Day10,
		Day15,
		Day30,
	}

	para := &Para{
		Location: "101010100",
		Lang:     lang.ZHCN,
		Unit:     qweather.METRIC,
	}

	for i, f := range fs {
		resp, err := f(para, key, true, &itest.NoProxyClient)
		if err != nil {
			t.Error(err)
			continue
		}
		if resp.Code != "200" {
			t.Error("return code is not 200", resp, i)
		}
	}
}
