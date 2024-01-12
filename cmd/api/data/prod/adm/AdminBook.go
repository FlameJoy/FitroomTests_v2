package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var AdminBook = models.Request{
	URL:      "https://back.fitroom.ru/api/admin/appointment/add",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(adminBook),
	RespRead: &adminBookResp,
	Auth:     &settings.TokenStage,
	Actions:  []func(){AdminBook_F1},
}

type Book struct {
	Clients            []uint `json:"clients"`
	Club               uint   `json:"club"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
	Workplace          uint   `json:"workplace"`
	ServiceProductType uint   `json:"serviceProductType"`
	Status             string `json:"status"`
	User               uint   `json:"user"`
	UserPackage        uint   `json:"userPackage"`
}

var adminBook = Book{
	Clients:            []uint{2058},
	Club:               29,
	Workplace:          36,
	ServiceProductType: 1,
	StartDate:          settings.StartDate3,
	EndDate:            settings.EndDate3,
	Status:             "booked",
	User:               2058,
	UserPackage:        7985,
}

type RentResp struct {
	ID int `json:"id"`
}

var adminBookResp RentResp

var AdminBook_F1 = func() {
	CancelRent.URL = "https://back.fitroom.ru/api/admin/appointment/delete/" + strconv.Itoa(adminBookResp.ID)
	// fmt.Println(CancelRent.URL)
}

var CancelRent = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}
