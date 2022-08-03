package server

import (
	"encoding/json"
	"github.com/WGrape/apimock/example/mock"
	"github.com/WGrape/apimock/example/model"
	"github.com/WGrape/apimock/example/service"
	"net/http"
	"strconv"
)

func ResponseError(w http.ResponseWriter, isMock bool, code int) {
	responseString, _ := json.Marshal(model.RespBase{Code: code, IsMock: isMock})
	_, _ = w.Write(responseString)
}

func ResponseOk(w http.ResponseWriter, isMock bool, data interface{}) {
	responseString, _ := json.Marshal(model.RespBase{Code: 0, Data: data, IsMock: isMock})
	_, _ = w.Write(responseString)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var isMock = mock.IsMockRequest(r)

	uidString := r.URL.Query().Get("uid")
	uidInt, err := strconv.ParseInt(uidString, 10, 64)
	if err != nil {
		ResponseError(w, isMock, 100)
		return
	}

	if mock.IsMockRequest(r) {
		ResponseOk(w, isMock, mock.GetUserInZh())
		return
	} else {
		ResponseOk(w, isMock, service.GetUser(uidInt))
		return
	}
}
