package qweather

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
		Client: http.DefaultClient,
	}
}

func (c *DefaultClient) SetConfig(config *Config) {
	c.Timeout = config.Timeout
	c.Transport = config.Proxy
	c.CheckRedirect = config.CheckRedirect
	c.Jar = config.Jar
}

type Credential struct {
	Key      string
	PublicID string
	Encrypt  bool
}

func (c *Credential) SetEncrypt() *Credential {
	c.Encrypt = true
	return c
}

func (c *Credential) UnsetEncrypt() *Credential {
	c.Encrypt = false
	return c
}

type Version uint8

const (
	_ Version = iota
	Free
	Standard
	Pro
)
