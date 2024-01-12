package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetCityList = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/city/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
