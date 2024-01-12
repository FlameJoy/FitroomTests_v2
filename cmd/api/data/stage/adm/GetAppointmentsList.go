package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetAppointmentsList = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/appointment/get-list",
	Method:  "POST",
	ReqBody: helper.PrepareReqBody(appointmentList),
	Auth:    &settings.TokenStage,
}

type Appointments struct {
	Status   string `json:"status"`
	FromDate string `json:"fromDate"`
	ToDate   string `json:"toDate"`
}

var appointmentList = Appointments{
	Status:   "booked",
	FromDate: settings.C_Date,
	ToDate:   settings.F_Date,
}
