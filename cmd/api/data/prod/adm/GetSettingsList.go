package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetSettingsList = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/setting/get-list",
	Method:  "POST",
	ReqBody: helper.PrepareReqBody(nil),
	Auth:    &settings.TokenStage,
}
