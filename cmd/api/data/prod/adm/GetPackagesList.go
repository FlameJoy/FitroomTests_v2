package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetPackagesList = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/package/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
