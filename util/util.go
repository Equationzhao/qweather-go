package util

import (
	"io"
	"net/http"
	"os"

	"github.com/Equationzhao/qweather-go"
	iutil "github.com/Equationzhao/qweather-go/internal/util"
)

func Get(req *http.Request, client qweather.Client) ([]byte, error) {
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func Request(url string, f func(r *http.Request)) (*http.Request, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	f(req)
	return req, nil
}

// Credential  从环境变量中获取 key 与 public id 并生成 *qweather.Credential
//
// 参数:
//
//	key, public id  环境变量名
//
// eg:
//
//	Credential("qweather_key","qweather_public_id")
func Credential(key, publicId string) *qweather.Credential {
	c := qweather.Credential{
		Key:      os.Getenv(key),
		PublicID: os.Getenv(publicId),
	}
	return &c
}

func CheckNilClient(client qweather.Client) qweather.Client {
	if client == nil {
		return qweather.NewDefaultClient()
	}
	return client
}

func UrlHelperBuilder(FreeEndPoint, StandardEndPoint string, ProEndPoint *string) func(isFreePlan qweather.Version, u ...string) string {
	return func(isFreePlan qweather.Version, u ...string) string {
		EndPoint := ""
		switch isFreePlan {
		case qweather.Free:
			EndPoint = FreeEndPoint
		case qweather.Standard:
			EndPoint = StandardEndPoint
		case qweather.Pro:
			if ProEndPoint == nil {
				return ""
			}
			EndPoint = *ProEndPoint
		default:
		}
		return iutil.Url(EndPoint, u...)
	}
}
