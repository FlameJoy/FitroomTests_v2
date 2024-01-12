package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateCountry = models.Request{
	URL:      "https://back.fitroom.fun/api/admin/country/add",
	Method:   "POST",
	Auth:     &settings.TokenStage,
	ReqBody:  helper.PrepareReqBody(newCountry),
	RespRead: &resp,
	Actions:  []func(){Country_F1},
}

var EditCountry = models.Request{
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(editCountry),
}

var DeleteCountry = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type CountryReq struct {
	Title string `json:"title"`
}

var newCountry = CountryReq{
	Title: "AUTOTEST",
}

var editCountry = CountryReq{
	Title: "EDITTED AUTOTEST",
}

func Country_F1() {
	EditCountry.URL = "https://back.fitroom.fun/api/admin/country/edit/" + strconv.Itoa(resp.ID)
	DeleteCountry.URL = "https://back.fitroom.fun/api/admin/country/delete/" + strconv.Itoa(resp.ID)
}
