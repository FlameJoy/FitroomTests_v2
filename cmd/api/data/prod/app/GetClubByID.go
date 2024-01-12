package app

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetClubByID = models.Request{
	URL:     "https://back.fitroom.ru/public/club/get-by-id/42", // Matisov kanal
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
