package warning

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

func TestWarning(t *testing.T) {
	para := &Para{
		Location: "116.41,39.92",
		Lang:     lang.ZHCN,
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
	t.Log(request.URL)
	if err != nil {
		t.Fatal(err)
	}
	helper(t, request, &RealTimeResponse{})
}

func TestCityList(t *testing.T) {
	para := &Para{
		Range: "cn",
	}
	response, err := CityList(para, key, qweather.Free, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	response, err = CityList(para, key, qweather.Free, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := CityListRequest(para, key, qweather.Free)
	t.Log(request.URL)

	if err != nil {
		t.Fatal(err)
	}
	helper(t, request, &CityListResponse{})
}
