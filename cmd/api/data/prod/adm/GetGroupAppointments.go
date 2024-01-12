package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetGroupAppintments = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/appointment/available",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(groupAppointmentList),
}

type GroupAppointments struct {
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
	Workplace           int    `json:"workplace"`
	ServiceProductTypes []int  `json:"serviceProductTypes"`
}

var groupAppointmentList = GroupAppointments{
	StartDate:           settings.P_Date,
	EndDate:             settings.C_Date,
	Workplace:           120,
	ServiceProductTypes: []int{10, 8, 11},
}
