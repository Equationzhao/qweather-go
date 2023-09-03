package test

import (
	"net/http"
	"net/url"
	"strconv"
)

var NoProxyClient = http.Client{
	Transport:     &http.Transport{Proxy: nil},
	CheckRedirect: nil,
	Jar:           nil,
	Timeout:       0,
}

func HttpProxyClient(port uint16) http.Client {
	return http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: "http",
				Host:   "localhost:" + strconv.Itoa(int(port)),
			}),
		},
	}
}
