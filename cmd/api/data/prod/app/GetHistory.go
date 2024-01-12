package app

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetHistory = models.Request{
	URL:     "https://back.fitroom.ru/api/lk/appointment/history",
	Method:  "POST",
	ReqBody: helper.PrepareReqBody(history),
	Auth:    &settings.TokenStage,
}

type HistoryReq struct {
	Paging Paging `json:"paging"`
}
type Paging struct {
	Page  uint `json:"page"`
	Limit uint `json:"limit"`
}

var history = HistoryReq{
	Paging: Paging{
		Page:  1,
		Limit: 20,
	},
}
