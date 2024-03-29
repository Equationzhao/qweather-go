// Package statusCode 状态码
package statusCode

type Code string

func (c Code) Translate() string {
	return Translate(c)
}

func (c Code) IsSuccess() bool {
	return c == Success
}

const (
	Success         Code = "200" // 请求成功
	DataNotExists   Code = "204" // 请求成功，但你查询的地区暂时没有你需要的数据。
	ParamError      Code = "400" // 请求错误，可能包含错误的请求参数或缺少必选的请求参数。
	AuthorizeFailed Code = "401" // 认证失败，可能使用了错误的KEY、数字签名(https://dev.qweather.com/docs/resource/signature-auth/)错误、KEY的类型错误（如使用SDK的KEY去访问Web API）。
	ExceedLimit     Code = "402" // 超过访问次数或余额不足以支持继续访问服务，你可以充值、升级访问量或等待访问量重置。
	NoPermission    Code = "403" // 无访问权限，可能是绑定的PackageName、BundleID、域名IP地址不一致，或者是需要额外付费的数据。
	NotFound        Code = "404" // 查询的数据或地区不存在。
	TooManyRequests Code = "429" // 超过限定的QPM（每分钟访问次数），请参考QPM说明(https://dev.qweather.com/docs/resource/glossary/#qpm)
	Timeout         Code = "500" // 无响应或超时，接口服务异常请联系我们(https://www.qweather.com/contact)
)

// Translate 翻译状态码
func Translate(code Code) string {
	switch code {
	case Success:
		return "Success"
	case DataNotExists:
		return "DataNotExists"
	case ParamError:
		return "ParamError"
	case AuthorizeFailed:
		return "AuthorizeFailed"
	case ExceedLimit:
		return "ExceedLimit"
	case NoPermission:
		return "NoPermission"
	case NotFound:
		return "NotFound"
	case TooManyRequests:
		return "TooManyRequests"
	case Timeout:
		return "Timeout"
	default:
		return "Unknown"
	}
}

func IsSuccess(code Code) bool {
	return code == Success
}
