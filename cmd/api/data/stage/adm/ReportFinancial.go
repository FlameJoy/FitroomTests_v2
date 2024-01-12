package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var ReportFinancial = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/report/revenue",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(newRepFinancial),
}

type RepFinancial struct {
	Clubs    []int  `json:"clubs"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

var newRepFinancial = RepFinancial{
	Clubs:    GetClubIDList(),
	FromDate: settings.P_Date,
	ToDate:   settings.C_Date,
}
