package indices

import (
	"testing"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	itest "github.com/Equationzhao/qweather-go/internal/test"
	"github.com/Equationzhao/qweather-go/statusCode"
	"github.com/Equationzhao/qweather-go/util"
)

var key = *util.Credential("qweather_key", "qweather_public_id").SetEncrypt()

var (
	Location = "116.41,39.92"
	Type     = "0"
	para     = &Para{
		Location: Location,
		Type:     Type,
	}
)

func TestDays(t *testing.T) {
	fs := []func(para *Para, key qweather.Credential, plan qweather.Version, client qweather.Client) (*Response, error){Day1, Day3}
	for _, f := range fs {
		day, err := f(para, key, qweather.Free, &itest.NoProxyClient)
		if err != nil {
			t.Fatal(err)
		}
		if !statusCode.IsSuccess(day.Code) {
			t.Fatal(err)
		}

		t.Log(day)
	}
}

func TestIndicesWithRequiredParam(t *testing.T) {
	days := [...]uint8{1, 3}
	for _, day := range days {
		r, err := IndicesWithRequiredParam(Location, Type, para, key, day, qweather.Free, &itest.NoProxyClient)
		if err != nil {
			t.Fatal(err)
		}
		if !statusCode.IsSuccess(r.Code) {
			t.Fatal(err)
		}
		t.Log(r)
	}
}

func TestIndicesRequestWithRequiredParam(t *testing.T) {
	days := [...]uint8{1, 3}
	for _, day := range days {
		r, err := IndicesRequestWithRequiredParam(Location, Type, para, key, day, qweather.Free)
		if err != nil {
			t.Fatal(err)
		}
		get, err := util.Get(r, &itest.NoProxyClient)
		if err != nil {
			t.Fatal(err)
		}
		var resp Response
		err = json.Unmarshal(get, &resp)
		if err != nil {
			t.Fatal(err)
		}
		if !statusCode.IsSuccess(resp.Code) {
			t.Fatal(err)
		}
		t.Log(resp)
	}
}
