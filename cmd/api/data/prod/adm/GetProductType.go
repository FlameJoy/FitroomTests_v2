package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetProductType = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/service-product-type/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
