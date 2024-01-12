package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateBranch = models.Request{
	URL:      "https://back.fitroom.ru/api/admin/branch/add",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(newBranch),
	Auth:     &settings.TokenStage,
	RespRead: &branchResp,
	Actions:  []func(){Branch_F1},
}

var EditBranch = models.Request{
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(editBranch),
}

var DeleteBranch = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type BranchReq struct {
	Active      uint   `json:"active"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var newBranch = BranchReq{
	Active:      1,
	Title:       "AUTOTEST",
	Description: "AUTOTEST",
}

var editBranch = BranchReq{
	Active:      1,
	Title:       "EDITTED AUTOTEST",
	Description: "EDITTED AUTOTEST",
}

type BranchResp struct {
	ID          int    `json:"id"`
	Actice      bool   `json:"actice"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Gallery     []struct {
		ID int `json:"id"`
	} `json:"gallery"`
}

var branchResp BranchResp

var Branch_F1 = func() {
	EditBranch.URL = "https://back.fitroom.ru/api/admin/branch/edit/" + strconv.Itoa(branchResp.ID)
	DeleteBranch.URL = "https://back.fitroom.ru/api/admin/branch/delete/" + strconv.Itoa(branchResp.ID)
}
