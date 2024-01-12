package app

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

type GroupTrainBook struct {
	User               uint   `json:"user"`
	Club               uint   `json:"club"`
	Workplace          uint   `json:"workplace"`
	ServiceProductType uint   `json:"serviceProductType"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
	UserPackage        uint   `json:"userPackage"`
	GroupTraining      int    `json:"groupTraining"`
}

var GroupTrain = models.Request{
	URL:      "https://back.fitroom.fun/api/lk/appointment/book",
	Method:   "POST",
	Auth:     &settings.TokenStage,
	ReqBody:  helper.PrepareReqBody(groupTrain),
	RespRead: &rentResp,
	Actions:  []func(){Rent_F1},
}

var groupTrain = GroupTrainBook{
	User:               1625,
	Club:               166,
	Workplace:          120,
	ServiceProductType: 8,
	StartDate:          settings.StartDate3,
	EndDate:            settings.EndDate3,
	UserPackage:        8064,
	GroupTraining:      16,
}
