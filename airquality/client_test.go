package airqualityBeta

import (
	"testing"

	"github.com/Equationzhao/qweather-go"
	itest "github.com/Equationzhao/qweather-go/internal/test"
	"github.com/Equationzhao/qweather-go/util"
)

var key = *util.Credential("qweather_key", "qweather_public_id")

func TestNow(t *testing.T) {
	param := &NowPara{
		LocationID: "101010100",
		Lang:       "zh",
		Pollutant:  true,
		Station:    true,
	}
	got, err := Now(param, key, qweather.FreePlan, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}

	got, err = Now(param, key, qweather.FreePlan, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}

	got, err = NowWithRequiredParam(param.LocationID, key, param, qweather.FreePlan, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}

	got, err = NowWithRequiredParam(param.LocationID, key, nil, qweather.FreePlan, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}
}

func TestStation(t *testing.T) {
	param := &StationPara{
		LocationID: "P58911",
		Lang:       "zh",
	}
	got, err := Station(param, key, qweather.FreePlan, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}

	got, err = Station(param, key, qweather.FreePlan, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}

	got, err = StationWithRequiredParam(param.LocationID, key, param, qweather.FreePlan, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}

	got, err = StationWithRequiredParam(param.LocationID, key, nil, qweather.FreePlan, &itest.NoProxyClient)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}
}
