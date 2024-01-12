package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetDistrictList = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/district/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
