package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetReview = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/review/get-list",
	Method:  "POST",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(&paging),
}
