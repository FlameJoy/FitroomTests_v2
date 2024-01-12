package app

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetPackageList = models.Request{
	URL:     "https://back.fitroom.fun/api/lk/package/get-list",
	Method:  "POST",
	ReqBody: helper.PrepareReqBody(club),
	Auth:    &settings.TokenStage,
}

var club = Club{
	Club: 13,
}
