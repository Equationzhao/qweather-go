package util

import "strings"

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
