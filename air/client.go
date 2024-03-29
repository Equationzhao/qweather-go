package air

import (
	"net/http"

	"github.com/Equationzhao/qweather-go"
	"github.com/Equationzhao/qweather-go/internal/json"
	"github.com/Equationzhao/qweather-go/util"
)

const (
	StandardEndPoint = "https://api.qweather.com/v7/air/"
	FreeEndPoint     = "https://devapi.qweather.com/v7/air/"
)

var ProEndPoint *string = nil

var url = util.UrlHelperBuilder(FreeEndPoint, StandardEndPoint, ProEndPoint)

func NowRequest(para *Para, key qweather.Credential, plan qweather.Version) (*http.Request, error) {
	r, err := util.Request(
		url(plan, "now"), func(r *http.Request) {
			q := r.URL.Query()
			q.Add("location", para.LocationID)
			q.Add("lang", para.Lang)
			if key.Encrypt {
				qweather.AddSignature(key.PublicID, key.Key, q)
			} else {
				q.Add("key", key.Key)
			}
			r.URL.RawQuery = q.Encode()
		},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func Now(para *Para, key qweather.Credential, plan qweather.Version, client qweather.Client) (*NowResponse, error) {
	request, err := NowRequest(para, key, plan)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	get, err := util.Get(request, client)
	if err != nil {
		return nil, err
	}
	var resp NowResponse
	err = json.Unmarshal(get, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func Day5Request(para *Para, key qweather.Credential, plan qweather.Version) (*http.Request, error) {
	r, err := util.Request(
		url(plan, "5d"), func(r *http.Request) {
			q := r.URL.Query()
			q.Add("location", para.LocationID)
			q.Add("lang", para.Lang)
			if key.Encrypt {
				qweather.AddSignature(key.PublicID, key.Key, q)
			} else {
				q.Add("key", key.Key)
			}
			r.URL.RawQuery = q.Encode()
		},
	)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func Day5(para *Para, key qweather.Credential, plan qweather.Version, client qweather.Client) (*Day5Response, error) {
	request, err := Day5Request(para, key, plan)
	if err != nil {
		return nil, err
	}
	client = util.CheckNilClient(client)
	get, err := util.Get(request, client)
	if err != nil {
		return nil, err
	}
	var resp Day5Response
	err = json.Unmarshal(get, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
