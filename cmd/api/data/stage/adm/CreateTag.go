package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateTag = models.Request{
	URL:      "https://back.fitroom.fun/api/admin/tag/add",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(newTag),
	Auth:     &settings.TokenStage,
	RespRead: &tagResp,
	Actions:  []func(){Tag_F1},
}

var EditTag = models.Request{
	Method:  "PUT",
	ReqBody: helper.PrepareReqBody(editTag),
	Auth:    &settings.TokenStage,
}

var DeleteTag = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type TagReq struct {
	Title string `jsom:"title"`
	Color string `jsom:"color"`
}

type TagResp struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

var newTag = TagReq{
	Title: "AUTOTEST",
	Color: "yellow",
}

var editTag = TagReq{
	Title: "EDITTED AUTOTEST",
	Color: "red",
}

var tagResp TagResp

var Tag_F1 = func() {
	EditTag.URL = "https://back.fitroom.fun/api/admin/tag/edit/" + strconv.Itoa(tagResp.ID)
	DeleteTag.URL = "https://back.fitroom.fun/api/admin/tag/delete/" + strconv.Itoa(tagResp.ID)
}
