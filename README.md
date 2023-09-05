# qweather-go

> 和风天气 sdk

## 使用

```shell
go get github.com/Equationzhao/qweather-go
```

## 获取 key && public id

参考 https://dev.qweather.com/docs/configuration/project-and-key/

## 样例

```go
package main

import (
	"log/slog"
	"net/http"

	"github.com/Equationzhao/qweather-go/statusCode"
	"github.com/Equationzhao/qweather-go/util"
	"github.com/Equationzhao/qweather-go/warning"
)

func main() {
	// 从环境变量中获取
	key := *util.Credential("qweather_key", "qweather_public_id").SetEncrypt()

	// 设置参数｜设置 credential｜设置是否为 free plan｜设置 client(nil 则使用 DefaultClient)
	r, err := warning.RealTime(&warning.Para{
		Range:    "cn",
		Location: "101230204",
	}, key, true, &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
		},
	})
	
	/*
	    r, err := warning.RealTime(&warning.Para{
			Range:    "cn",
			Location: "101230204",
		}, key, true, nil)
	 */

	if err != nil {
		slog.Error("failed to get city list", "err", err)
		return
	} else if r.Code != statusCode.Success {
		slog.Error("failed", "code", r.Code, "meaning", statusCode.Translate(r.Code))
		return
	}
	slog.Info("success", "code", r.Code, "update time", r.UpdateTime, "warning", r.Warning[0].Text)
}
```