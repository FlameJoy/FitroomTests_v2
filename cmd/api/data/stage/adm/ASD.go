package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var ASD = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/access-system-device/add",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(&newASD),
}

type ASDReq struct {
	SerialNumber   string `json:"serialNumber"`
	Type           string `json:"type"`
	Code           string `json:"code"`
	OpenAfterTime  int    `json:"openAfterTime"`
	OpenBeforeTime int    `json:"openBeforeTime"`
	Workplace      int    `json:"workplace"`
}

var newASD = ASDReq{
	SerialNumber:   "010101",
	Type:           "AUTOTEST",
	Code:           "1111",
	OpenAfterTime:  1200,
	OpenBeforeTime: 600,
	Workplace:      3,
}
