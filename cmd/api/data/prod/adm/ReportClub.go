package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var ReportClub = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/report/club",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(newRepClub),
}

type RepClub struct {
	Clubs         []int  `json:"clubs"`
	FromDate      string `json:"fromDate"`
	ToDate        string `json:"toDate"`
	SalaryPercent int    `json:"salaryPercent"`
}

var newRepClub = RepClub{
	Clubs:         GetClubIDList(),
	FromDate:      settings.P_Date,
	ToDate:        settings.C_Date,
	SalaryPercent: 50,
}
