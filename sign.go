package qweather_go

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// MIT License
//
// Copyright (c) 2020 Ink33
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// GetSignature 和风天气签名生成算法-Golang版本
// https://github.com/Ink-33/go-heweather/blob/main/v7/client.go#L65C2-L65C2
// From https://github.com/Ink-33/go-heweather/blob/3695eeab1c0d1590ced1fb3d3fd4dadd3f014245/v7/client.go#L65C2-L65C2
func GetSignature(publicID, key string, para url.Values) (signature string) {
	var sa []string
	var escapesa []string

	for k, v := range para {
		if len(v) == 0 || k == "sign" {
			continue
		}
		sa = append(sa, k+"="+v[0])
		escapesa = append(escapesa, k+"="+v[0])
	}

	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	sa = append(sa, "t="+timestamp, "publicid="+publicID)
	escapesa = append(escapesa, "t="+timestamp, "publicid="+publicID)
	sort.Strings(sa)
	sort.Strings(escapesa)

	md5c := md5.New()
	md5c.Reset()
	_, _ = md5c.Write([]byte(strings.Join(sa, "&") + key))
	return fmt.Sprintf("%x", md5c.Sum(nil))
}

// ChangeRequest 用于修改请求参数, 将GetSignature得到的签名结果 作为参数添加至请求中，参数名为 sign, 即 sign=xxxxxxx
func ChangeRequest(publicID, key string, req *http.Request) {
	q := req.URL.Query()
	signature := GetSignature(publicID, key, q)
	q.Add("sign", signature)
	req.URL.RawQuery = q.Encode()
}
