package app

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetAvatar = models.Request{
	URL:     "https://back.fitroom.ru/public/image/get-by-id/3703",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
