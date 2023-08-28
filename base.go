package qweather_go

import (
	"net/http"
	"time"
)

type Config struct {
	Timeout       time.Duration
	Proxy         *http.Transport
	CheckRedirect func(req *http.Request, via []*http.Request) error
	Jar           http.CookieJar
}

func (c *Config) BuildClient() Client {
	return &DefaultClient{
		Client: &http.Client{
			Timeout:       c.Timeout,
			Transport:     c.Proxy,
			CheckRedirect: c.CheckRedirect,
			Jar:           c.Jar,
		},
	}
}

type doClient interface {
	Do(r *http.Request) (*http.Response, error)
}

type Client interface {
	doClient
}

type DefaultClient struct {
	*http.Client
}

func NewDefaultClient() *DefaultClient {
	return &DefaultClient{
		Client: &http.Client{},
	}
}

func (c *DefaultClient) SetConfig(config *Config) {
	c.Timeout = config.Timeout
	c.Transport = config.Proxy
	c.CheckRedirect = config.CheckRedirect
	c.Jar = config.Jar
}
