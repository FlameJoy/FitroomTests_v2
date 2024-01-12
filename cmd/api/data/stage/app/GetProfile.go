package app

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetProfile = models.Request{
	URL:     "https://back.fitroom.fun/api/lk/user/profile",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
