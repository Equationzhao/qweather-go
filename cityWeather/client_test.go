package cityWeather

import (
	"io"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	itest "github.com/Equationzhao/qweather-go/internal/test"
	"github.com/Equationzhao/qweather-go/lang"
	"github.com/Equationzhao/qweather-go/util"
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
			t.Fatal(err)
		}
		if resp.Code != "200" {
			t.Fatal("return code is not 200")
		}
	}
}
