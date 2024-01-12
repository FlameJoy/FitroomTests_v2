package app

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var Rent = models.Request{
	URL:      "https://back.fitroom.fun/api/lk/appointment/book",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(bookRent),
	RespRead: &rentResp,
	Auth:     &settings.TokenStage,
	Actions:  []func(){Rent_F1},
}

var bookRent = Book{
	User:               2058,
	Club:               13,
	Workplace:          13,
	ServiceProductType: 1,
	StartDate:          settings.StartDate1,
	EndDate:            settings.EndDate1,
	UserPackage:        7985,
}

type RentResp struct {
	ID int `json:"id"`
}

var rentResp RentResp

var Rent_F1 = func() {
	CancelRent.URL = "https://back.fitroom.fun/api/lk/appointment/cancel/" + strconv.Itoa(rentResp.ID)
	// fmt.Println(CancelRent.URL)
}

var CancelRent = models.Request{
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(new(Book)),
}
