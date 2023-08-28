package cityWeather

import (
	"io"
	"net/http"
	"os"
	"testing"

	qweathergo "qweather"
	"qweather/json"
)

var key = os.Getenv("qweather_key")
var publicID = os.Getenv("qweather_public_id")

func helper(t *testing.T, request *http.Request, m any) {
	t.Helper()
	qweathergo.ChangeRequest(publicID, key, request)
	response, err := http.DefaultClient.Do(request)
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
		Lang:     "zh",
		Unit:     METRIC,
	}
	resp, err := RealTime(para, key, true, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	if resp.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := RealTimeRequest(para, key, true)
	if err != nil {
		t.Fatal(err)
	}
	helper(t, request, &RealTimeResponse{})

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
		resp, err := Hourly(para, key, arg, true, nil)
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
		Lang:     "zh",
		Unit:     METRIC,
	}

	// 3,7
	args := []uint8{3, 7}
	for _, arg := range args {
		resp, err := Daily(para, key, arg, true, nil)
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
	helper(t, request, &DaysResponse{})

}
