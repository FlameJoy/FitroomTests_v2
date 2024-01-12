package app

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetAppointments = models.Request{
	URL:     "https://back.fitroom.ru/api/lk/v2/appointment/available",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(available),
}

type Available struct {
	Workplace           int   `json:"workplace"`
	ServiceProductTypes []int `json:"serviceProductTypes"`
	// Trainers            []int  `json:"trainers"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

var available = Available{
	Workplace:           27,
	ServiceProductTypes: []int{1},
	StartDate:           settings.C_Date,
	EndDate:             settings.F_Date,
}

// var GetAppointmentsGroup = models.Request{
// 	URL:     "https://back.fitroom.ru/api/lk/v2/appointment/available",
// 	Method:  "POST",
// 	Auth:    &settings.TokenStage,
// 	ReqBody: helper.PrepareReqBody(available2),
// }

// var available2 = Available{
// 	Workplace:           27,
// 	ServiceProductTypes: []int{10, 11, 8},
// 	StartDate:           settings.C_Date,
// 	EndDate:             settings.F_Date,
// }
