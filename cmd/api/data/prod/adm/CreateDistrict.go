package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateDistrict = models.Request{
	URL:      "https://back.fitroom.ru/api/admin/district/add",
	Method:   "POST",
	Auth:     &settings.TokenStage,
	ReqBody:  helper.PrepareReqBody(newDistrict),
	RespRead: &resp,
	Actions:  []func(){District_F1},
}

var EditDistrict = models.Request{
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(editDistrict),
}

var DeleteDistrict = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type DistrictReq struct {
	Title string `json:"title"`
	City  int    `json:"city"`
}

var newDistrict = DistrictReq{
	Title: "AUTOTEST",
	City:  1,
}

var editDistrict = DistrictReq{
	Title: "EDITTED AUTOTEST",
	City:  1,
}

func District_F1() {
	EditDistrict.URL = "https://back.fitroom.ru/api/admin/district/edit/" + strconv.Itoa(resp.ID)
	DeleteDistrict.URL = "https://back.fitroom.ru/api/admin/district/delete/" + strconv.Itoa(resp.ID)
}
