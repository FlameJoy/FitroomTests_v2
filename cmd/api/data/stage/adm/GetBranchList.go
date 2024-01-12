package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetBranchList = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/branch/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
