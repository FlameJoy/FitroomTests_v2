package app

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var Train = models.Request{
	URL:      "https://back.fitroom.ru/api/lk/appointment/book",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(bookTrain),
	RespRead: &rentResp,
	Auth:     &settings.TokenStage,
	Actions:  []func(){Rent_F1},
}

var bookTrain = Book{
	User:               6158,
	Club:               13,
	Workplace:          13,
	ServiceProductType: 2,
	StartDate:          settings.StartDate2,
	EndDate:            settings.EndDate2,
	UserPackage:        8019,
}
