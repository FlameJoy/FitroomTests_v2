package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetComplexList = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/complex/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
