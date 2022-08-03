package service

import (
	"fmt"
	"github.com/WGrape/apimock/example/model"
)

func GetUser(uid int64) *model.QueryGetUserResp {
	return &model.QueryGetUserResp{
		Uid:    uid,
		Name:   fmt.Sprintf("Name-%d", uid),
		Age:    45,
		Gender: 1,
		City:   "Unknown",
	}
}
