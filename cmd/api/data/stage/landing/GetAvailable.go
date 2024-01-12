package site

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

type Available struct {
	Workplace           int    `json:"workplace"`
	ServiceProductTypes []int  `json:"serviceProductTypes"`
	Trainers            []int  `json:"trainers"`
	StartDate           string `json:"startDate"`
	EndDate             string `json:"endDate"`
}

var available = Available{
	Workplace:           120,
	ServiceProductTypes: []int{10, 8, 11},
	StartDate:           settings.P_Date,
	EndDate:             settings.P_Date,
	Trainers:            []int{2058, 1625},
}

var GetAvailable = models.Request{
	URL:     "https://back.fitroom.fun/public/appointment/available",
	Method:  "POST",
	Auth:    new(string),
	ReqBody: helper.PrepareReqBody(available),
}
