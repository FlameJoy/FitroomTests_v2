package app

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetAvatar = models.Request{
	URL:     "https://back.fitroom.fun/public/image/get-by-id/2076",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
