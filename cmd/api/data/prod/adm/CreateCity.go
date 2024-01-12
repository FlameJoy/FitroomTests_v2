package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateCity = models.Request{
	URL:      "https://back.fitroom.ru/api/admin/city/add",
	Method:   "POST",
	Auth:     &settings.TokenStage,
	ReqBody:  helper.PrepareReqBody(newCity),
	RespRead: &resp,
	Actions:  []func(){City_F1},
}

var EditCity = models.Request{
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(editCity),
}

var DeleteCity = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type CityReq struct {
	Title   string `json:"title"`
	Country int    `json:"country"`
	Slug    string `json:"slug"`
	Active  bool   `json:"active"`
	Phone   string `json:"phone"`
}

var newCity = CityReq{
	Title:   "AUTOTEST",
	Country: 1,
	Slug:    "autotest",
	// Active:  false,
	Phone: "79999999999",
}

var editCity = CityReq{
	Title:   "EDITTED AUTOTEST",
	Country: 1,
	Slug:    "editted autotest",
	Active:  false,
	Phone:   "78888888888",
}

func City_F1() {
	EditCity.URL = "https://back.fitroom.ru/api/admin/city/edit/" + strconv.Itoa(resp.ID)
	DeleteCity.URL = "https://back.fitroom.ru/api/admin/city/delete/" + strconv.Itoa(resp.ID)
}
