package statusCode

import "testing"

func TestTranslate(t *testing.T) {
	codes := []Code{Success, DataNotExists, ParamError, AuthorizeFailed, ExceedLimit, NoPermission, NotFound, TooManyRequests, Timeout, "unknown"}
	for _, code := range codes {
		t.Log(Translate(code))
		t.Log(code.Translate())
	}
}
