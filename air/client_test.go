package air

import (
	"testing"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/lang"
	"github.com/Equationzhao/qweather-go/util"
)

var key = *util.Credential("qweather_key", "qweather_public_id").SetEncrypt()

func TestNow(t *testing.T) {
	para := &Para{
		LocationID: "101010100",
		Lang:       lang.ZHCN,
	}
	got, err := Now(para, key, qweather.Free, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}

	_key := key
	_key.UnsetEncrypt()
	got, err = Now(para, _key, qweather.Free, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}
}

func TestDay5(t *testing.T) {
	para := &Para{
		LocationID: "101010100",
		Lang:       lang.ZHCN,
	}
	got, err := Day5(para, key, qweather.Free, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}

	_key := key
	_key.UnsetEncrypt()
	got, err = Day5(para, _key, qweather.Free, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(got)
	if !got.Code.IsSuccess() {
		t.Error("return code is not 200")
	}
}
