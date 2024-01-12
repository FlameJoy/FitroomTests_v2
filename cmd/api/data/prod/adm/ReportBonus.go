package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var ReportBonus = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/report/bonus-flow",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(newRepBonus),
}

type RepBonus struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

var newRepBonus = RepBonus{
	FromDate: settings.C_Date,
	ToDate:   settings.F_Date,
}
