package util

import (
	"io"
	"net/http"
	"strings"

	"github.com/Equationzhao/qweather-go"
)

func Url(Endpoint string, u ...string) string {
	b := strings.Builder{}
	b.WriteString(Endpoint)
	l := len(u)
	if l == 0 {
		return b.String()
	}
	for i := 0; i < l-1; i++ {
		b.WriteString(u[i])
		b.WriteString("/")
	}
	b.WriteString(u[l-1])
	return b.String()
}

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
