package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var ReportConversion = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/report/conversion",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(newRepConversion),
}

type RepConversion struct {
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

var newRepConversion = RepConversion{
	FromDate: "2023-01-01",
	ToDate:   settings.C_Date,
}
