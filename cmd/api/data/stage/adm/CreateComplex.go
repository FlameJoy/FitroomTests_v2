package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateComplex = models.Request{
	URL:      "https://back.fitroom.fun/api/admin/complex/add",
	Method:   "POST",
	Auth:     &settings.TokenStage,
	ReqBody:  helper.PrepareReqBody(newComplex),
	RespRead: &resp,
	Actions:  []func(){Complex_F1},
}

var EditComplex = models.Request{
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(editComplex),
}

var DeleteComplex = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type ComplexReq struct {
	Title string `json:"title"`
	City  int    `json:"city"`
}

var newComplex = ComplexReq{
	Title: "AUTOTEST",
	City:  1,
}

var editComplex = ComplexReq{
	Title: "EDITTED AUTOTEST",
	City:  1,
}

func Complex_F1() {
	EditComplex.URL = "https://back.fitroom.fun/api/admin/omplex/edit/" + strconv.Itoa(resp.ID)
	DeleteComplex.URL = "https://back.fitroom.fun/api/admin/complex/delete/" + strconv.Itoa(resp.ID)
}
