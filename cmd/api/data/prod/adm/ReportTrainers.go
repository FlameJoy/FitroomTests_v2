package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var ReportTrainers = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/report/club-trainers",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(newRepTrainers),
}

type RepTrainers struct {
	SalaryPercent uint   `json:"salaryPercent"`
	FromDate      string `json:"fromDate"`
	ToDate        string `json:"toDate"`
}

var newRepTrainers = RepTrainers{
	FromDate:      settings.P_Date,
	ToDate:        settings.C_Date,
	SalaryPercent: 50,
}
