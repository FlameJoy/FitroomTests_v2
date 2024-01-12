package site

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

type History struct {
	FromDate   string `json:"fromDate"`
	ToDate     string `json:"toDate"`
	OnlyGroups int    `json:"onlyGroups"`
}

var history = History{
	FromDate:   settings.P_Date,
	ToDate:     settings.C_Date,
	OnlyGroups: 1,
}

var GetHistory = models.Request{
	URL:     "https://back.fitroom.fun/api/lk/appointment/history",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(history),
}
