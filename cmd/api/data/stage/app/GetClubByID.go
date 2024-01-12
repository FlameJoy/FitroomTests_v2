package app

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetClubByID = models.Request{
	URL:     "https://back.fitroom.fun/public/club/get-by-id/13", // Pushkin
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
