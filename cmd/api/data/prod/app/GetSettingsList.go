package app

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetSettingsList = models.Request{
	URL:     "https://back.fitroom.ru/api/lk/setting/get-list",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
