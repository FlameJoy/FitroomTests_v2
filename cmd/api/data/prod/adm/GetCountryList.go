package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetCountryList = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/country/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
