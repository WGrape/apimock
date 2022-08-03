package mock

import (
	"github.com/WGrape/apimock/apimock"
	"github.com/WGrape/apimock/example/model"
)

func GetUserInEn() *apimock.APIMock {
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
		Title:       "Get user information",
		Description: "information",
		Route:       "/api/get_user",
		Request:     request,
		Response:    response,
		ResponseExplain: []apimock.ResponseExplainItem{
			{
				Field: "data.uid",
				Type:  "int",
				Mark:  "the id of user",
			},
			{
				Field: "data.name",
				Type:  "string",
				Mark:  "the name of user",
			},
			{
				Field: "data.age",
				Type:  "int",
				Mark:  "the age of user",
			},
			{
				Field: "data.gender",
				Type:  "int",
				Mark:  "the sex of user, 1(female), 2(male)",
			},
			{
				Field: "data.city",
				Type:  "string",
				Mark:  "the city of user",
			},
		},
	}

	return &apiMock
}
