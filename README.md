# qweather-go

> 和风天气 v7 sdk

## 使用

```shell
go get github.com/Equationzhao/qweather-go
```

## 获取 key && public id

参考 https://dev.qweather.com/docs/configuration/project-and-key/

不需要签名认证时，请获取 key 并设置
```go
key := qweather.Credential{
    Key      : "your key",
    Encrypt  : false,
}
```

需要签名认证时，请参考下文 [加密签名认证](#加密签名认证) 部分

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

	// 设置参数｜设置 credential｜设置订阅模式｜设置 client(nil 则使用 DefaultClient)
	r, err := warning.RealTime(&warning.Para{
		Range:    "cn",
		Location: "101230204",
	}, key, qweather.Free, &http.Client{
		Transport: &http.Transport{
			Proxy: nil,
		},
	})
	
	/*
	    r, err := warning.RealTime(&warning.Para{
			Range:    "cn",
			Location: "101230204",
		}, key, qweather.Free, nil)
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

## 加密签名认证

参考 https://dev.qweather.com/docs/api/signature-auth/

在 Credential 中设置 `SetEncrypt()` 即可, 若要取消，则使用 `UnsetEncrypt()`, 默认为不加密

可以使用 `util.Credential` 来从环境变量中获取获取 Credential

> 注意 该函数以及(Un)SetEncrypt 返回的都是 *Credential 变量，可以直接使用 如下方法链式调用

```go
key := *util.Credential("qweather_key", "qweather_public_id").SetEncrypt()
```

## 生成 request

可以使用 函数名为 对应方法名+Request 的函数来生成 `*http.Request`

例如:
```go
// 原方法
response, err := cityWeather.RealTime(para, key, qweather.Free, nil)
// 对应生成 Request 的方法如下
request, err := cityWeather.RealTimeRequest(para, key, qweather.Free)
```

## 默认参数

可以使用 函数名为 对应方法名+RequiredParam 的函数来调用要求默认参数的 对应方法

例如:
```go
response, err := cityWeather.RealTime(para, key, true, nil)
//  对应生产 Request 的方法如下
request, err := cityWeather.RealTimeWithRequiredParam("101010100",para, key, true，nil)
```

> **⚠ 注意**
> 
> > 注意 : 该方法将用要求的参数覆盖 para 中对应的参数，上述例子中 "101010100" 覆盖了 para.Location

## 设置 client

对于一般方法，最后一个参数为 `*http.Client`, 传入 `nil` ，则使用 `http.DefaultClient`

可以自定义 client 来设置代理等

```go
client := &http.Client{
    Transport: &http.Transport{
        Proxy: http.ProxyURL(&url.URL{
            Scheme: "http",
            Host:   "localhost:1080",
        })
    },
}

response, err := cityWeather.RealTime(para, key, true, client)
```

## 状态码

详见 [statusCode](statusCode/README.md)

可以通过 `statusCode.XXX` 来获取对应的状态码

可以通过 `statusCode.Translate(statusCode.XXX)` 或 `code.Translate()` 来获取对应的状态码含义

```go
slog.Error("failed", "code", response.Code, "meaning", statusCode.Translate(response.Code))
slog.Error("failed", "code", response.Code, "meaning", response.Code.Translate())
```

## 多语言

详见 [lang](lang/README.md)

使用 lang.XXX 来获取对应的语言代码

```go
para.Lang = lang.ZHCN
```

## plan 参数

参考 https://dev.qweather.com/docs/configuration/api-config/#了解api地址和参数

> **⚠ 注意**
> 
> >如果你使用的是免费订阅，必须设置为免费订阅
> 
>  免费订阅会将API host改为 devapi.qweather.com。
> 
>  但地理信息服务除外，无论免费订阅还是付费订阅，都使用geoapi.qweather.com。

##  自定义json库

内置了 [sonic](https://github.com/bytedance/sonic) 和 [jsoniter](https://github.com/json-iterator/go)

使用 tag 'sonic'/'jsoniter' 来指定使用的 json 库

```shell
go build -tags sonic
```

```shell
go build -tags jsoniter
```

或者设置 `qweather.SetJsonMarshal(xxx)` / `qweather.SetJsonUnmarshal(xxx)` 来指定使用其他 json 库

```go
qweather.SetJsonMarshal(MyMarshal)
qweather.SetJsonUnmarshal(MyUnmarshal)
```

## 注意事项

项目基于 [qweather v7](https://dev.qweather.com/docs/) 文档开发

请在使用前阅读 [和风天气使用条款](https://dev.qweather.com/docs/terms/)

使用和风天气服务，您需要同意和风天气的各项服务条款和限制。

您使用本项目或进行的任何二次开发均代表您的个人行为，不代表项目开发者对该等行为的任何观点或立场，也不代表对您的任何明示或暗示、担保或认可。

### LICENSE

```text
MIT License

Copyright (c) 2023 equationzhao

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```