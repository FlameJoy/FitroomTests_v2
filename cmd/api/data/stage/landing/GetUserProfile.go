package site

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetUserProfile = models.Request{
	URL:    "https://back.fitroom.fun/api/lk/user/profile",
	Method: "GET",
	Auth:   &settings.TokenStage,
}
