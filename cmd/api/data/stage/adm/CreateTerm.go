package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateTerm = models.Request{
	URL:      "https://back.fitroom.fun/api/admin/term/add",
	Method:   "POST",
	Auth:     &settings.TokenStage,
	ReqBody:  helper.PrepareReqBody(newTerm),
	RespRead: &resp,
	Actions:  []func(){Term_F1},
}

var EditTerm = models.Request{
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(editTerm),
}

var DeleteTerm = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type TermReq struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

var newTerm = TermReq{
	Title:       "AUTOTEST",
	Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
}

var editTerm = TermReq{
	Title:       "EDITTED AUTOTEST",
	Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
}

func Term_F1() {
	EditTerm.URL = "https://back.fitroom.fun/api/admin/term/edit/" + strconv.Itoa(resp.ID)
	DeleteTerm.URL = "https://back.fitroom.fun/api/admin/term/delete/" + strconv.Itoa(resp.ID)
}
