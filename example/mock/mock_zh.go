package mock

import (
	"github.com/WGrape/apimock/apimock"
	"github.com/WGrape/apimock/example/model"
)

func GetUserInZh() *apimock.APIMock {
	request := model.QueryGetUserReq{}
	response := model.RespBase{
		Code: 0,
		Data: model.QueryGetUserResp{
			Uid:    88328876,
			Name:   "Peter",
			Age:    45,
			Gender: 1,
			City:   "New York",
		},
	}
	apiMock := apimock.APIMock{
		Title:       "获取用户接口",
		Description: "获取用户接口",
		Route:       "/api/get_user",
		Request:     request,
		Response:    response,
		ResponseExplain: []apimock.ResponseExplainItem{
			{
				Field: "data.uid",
				Type:  "int",
				Mark:  "用户ID",
			},
			{
				Field: "data.name",
				Type:  "string",
				Mark:  "用户昵称",
			},
			{
				Field: "data.age",
				Type:  "int",
				Mark:  "用户年龄",
			},
			{
				Field: "data.gender",
				Type:  "int",
				Mark:  "用户性别, 1女, 2男",
			},
			{
				Field: "data.city",
				Type:  "string",
				Mark:  "用户所在城市",
			},
		},
	}

	return &apiMock
}
