package warning

import (
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	itest "github.com/Equationzhao/qweather-go/internal/test"
	"github.com/Equationzhao/qweather-go/lang"
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

func TestWarning(t *testing.T) {
	para := &Para{
		Location: "116.41,39.92",
		Lang:     lang.ZHCN,
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
	response, err := CityList(para, key, true, &itest.NoProxyClient)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(response)
	if response.Code != "200" {
		t.Fatal("return code is not 200")
	}

	request, err := CityListRequest(para, key, true)
	t.Log(request.URL)

	if err != nil {
		t.Fatal(err)
	}
	helper(t, request, &CityListResponse{})
}
