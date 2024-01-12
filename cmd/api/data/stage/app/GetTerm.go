package app

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetTerm = models.Request{
	URL:     "https://back.fitroom.fun/public/term/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
