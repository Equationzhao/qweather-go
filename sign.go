package qweather

import (
	"crypto/md5"
	"encoding/hex"
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
func getSignature(key string, param url.Values) (signature string) {
	sa := make([]string, 0, len(param)+2)
	for k, v := range param {
		if len(v) == 0 || v[0] == "" {
			continue
		}
		sa = append(sa, k+"="+strings.Join(v, ","))
	}

	sort.Strings(sa)
	return _MD5(strings.Join(sa, "&") + key)
}

func _MD5(s string) string {
	sum := md5.Sum([]byte(s))
	return hex.EncodeToString(sum[:])
}

func AddSignature(publicID string, key string, q url.Values) {
	q.Add("t", strconv.FormatInt(time.Now().Unix(), 10))
	q.Add("publicid", publicID)
	q.Add("sign", getSignature(key, q))
}
