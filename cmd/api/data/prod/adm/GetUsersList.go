package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetUsersList = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/user/get-list",
	Method:  "POST",
	ReqBody: helper.PrepareReqBody(pageLimit),
	Auth:    &settings.TokenStage,
}
