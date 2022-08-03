package model

type RespBase struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	IsMock bool        `json:"is_mock"`
}

type QueryGetUserReq struct {
	Uid int64 `json:"uid"`
}

type QueryGetUserResp struct {
	Uid    int64  `json:"uid"`
	Name   string `json:"name"`
	Age    uint8  `json:"age"`
	Gender uint8  `json:"gender"`
	City   string `json:"city"`
}
