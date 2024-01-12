package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateWorkplace = models.Request{
	URL:      "https://back.fitroom.fun/api/admin/workplace/add",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(newWorkplace), // Pushkin
	Auth:     &settings.TokenStage,
	RespRead: &workplaceResp,
	Actions:  []func(){WP_F1},
}

var EditWorkplace = models.Request{
	Method:  "PUT",
	ReqBody: helper.PrepareReqBody(editWorkplace), // Pushkin
	Auth:    &settings.TokenStage,
}

var DeleteWorkplace = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type WorkplaceReq struct {
	Title string `json:"title"`
	Club  uint   `json:"club"`
}

type WorkplaceResp struct {
	ID     int  `json:"id"`
	Actice bool `json:"actice"`
	Club   struct {
		ID       int    `json:"id"`
		Title    string `json:"title"`
		Street   string `json:"street"`
		House    string `json:"house"`
		Building string `json:"building"`
		City     struct {
			ID    int    `json:"id"`
			Title string `json:"title"`
		} `json:"club"`
		Complex  []int `json:"complex"`
		District []int `json:"district"`
	} `json:"club"`
	Title   string `json:"title"`
	Gallery struct {
		ID int `json:"id"`
	} `json:"gallery"`
}

var newWorkplace = WorkplaceReq{
	Title: "AUTOTEST",
	Club:  13,
}

var editWorkplace = WorkplaceReq{
	Title: "EDITTED AUTOTEST",
	Club:  13,
}

var workplaceResp WorkplaceResp

var WP_F1 = func() {
	EditWorkplace.URL = "https://back.fitroom.fun/api/admin/workplace/edit/" + strconv.Itoa(workplaceResp.ID)
	DeleteWorkplace.URL = "https://back.fitroom.fun/api/admin/workplace/delete/" + strconv.Itoa(workplaceResp.ID)
}
