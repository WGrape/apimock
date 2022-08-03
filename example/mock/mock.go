package mock

import (
	"net/http"
)

const (
	App   = "example"
	Token = "1ZmRzZ3NhZXIzNHR3c2ZmZHNm"
)

func IsMockRequest(r *http.Request) bool {
	if r.URL.Query().Get("mock_app") != App {
		return false
	}
	if r.URL.Query().Get("mock_token") != Token {
		return false
	}
	return true
}
