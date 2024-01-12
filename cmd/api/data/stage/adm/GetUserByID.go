package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetUserByID = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/user/get-by-id/2058",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
