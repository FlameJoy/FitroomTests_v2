package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateClub = models.Request{
	URL:      "https://back.fitroom.ru/api/admin/club/add",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(newClub),
	Auth:     &settings.TokenStage,
	RespRead: &clubResp,
	Actions:  []func(){Club_F1},
}

var EditClub = models.Request{
	Method:  "PUT",
	ReqBody: helper.PrepareReqBody(editClub),
	Auth:    &settings.TokenStage,
}

var DeleteClub = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type ClubReq struct {
	Title     string `json:"title"`
	City      uint   `json:"city"`
	District  uint   `json:"district"`
	Complex   uint   `json:"complex"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Building  string `json:"building"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Cardio    uint   `json:"cardio"`
	OpenDate  string `json:"openDate"`
}

type ClubResp struct {
	ID        int    `json:"id"`
	Actice    bool   `json:"actice"`
	Title     string `json:"title"`
	Street    string `json:"street"`
	House     string `json:"house"`
	Building  string `json:"building"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	City      struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
	Complex struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"complex"`
	District struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"district"`
	Users   []int `json:"users"`
	Gallery struct {
		ID int `json:"id"`
	} `json:"gallery"`
	ServiceProductTypes []int  `json:"serviceProductTypes"`
	Cardio              bool   `json:"cardio"`
	OpenDate            string `json:"openDate"`
}

var newClub = ClubReq{
	Title:     "AUTOTEST",
	City:      1,
	District:  1,
	Complex:   1,
	Street:    "AUTOTEST",
	House:     "1",
	Building:  "1",
	Latitude:  "50.00000",
	Longitude: "50.00000",
	OpenDate:  settings.F_Date,
}

var editClub = ClubReq{
	Title:     "EDITTED AUTOTEST",
	City:      2,
	District:  2,
	Complex:   2,
	Street:    "EDITTED AUTOTEST",
	House:     "2",
	Building:  "2",
	Latitude:  "50.20000",
	Longitude: "50.20000",
	OpenDate:  settings.F_Date,
}

var clubResp ClubResp

var Club_F1 = func() {
	EditClub.URL = "https://back.fitroom.ru/api/admin/club/edit/" + strconv.Itoa(clubResp.ID)
	DeleteClub.URL = "https://back.fitroom.ru/api/admin/club/delete/" + strconv.Itoa(clubResp.ID)
}
