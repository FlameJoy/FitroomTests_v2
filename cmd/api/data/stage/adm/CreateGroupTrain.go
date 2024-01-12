package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateGroupTrain = models.Request{
	URL:      "https://back.fitroom.fun/api/admin/group-training/add",
	Method:   "POST",
	Auth:     &settings.TokenStage,
	ReqBody:  helper.PrepareReqBody(&newGroupTrain),
	RespRead: &resp,
	Actions:  []func(){GroupTrain_F1},
}

var EditGroupTrain = models.Request{
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(&editGroupTrain),
}

var DeleteGroupTrain = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type GroupTrainReq struct {
	Title   string `json:"title"`
	Slots   int    `json:"slots"`
	Service int    `json:"service"`
	Club    int    `json:"club"`
	Trainer int    `json:"trainer"`
}

var newGroupTrain = GroupTrainReq{
	Title:   "AUTOTEST",
	Slots:   6,
	Service: 10,
	Club:    29,
	Trainer: 6289,
}

var editGroupTrain = GroupTrainReq{
	Title:   "EDITTED AUTOTEST",
	Slots:   10,
	Service: 10,
	Club:    29,
	Trainer: 6290,
}

// type GroupTrainResp struct {
// }

var GroupTrain_F1 = func() {
	EditGroupTrain.URL = "https://back.fitroom.fun/api/admin/group-training/edit/" + strconv.Itoa(resp.ID)
	DeleteGroupTrain.URL = "https://back.fitroom.fun/api/admin/group-training/delete/" + strconv.Itoa(resp.ID)
}
