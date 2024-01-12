package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetTerms = models.Request{
	URL:     "https://back.fitroom.fun/api/admin/term/get",
	Method:  "GET",
	ReqBody: nil,
	Auth:    &settings.TokenStage,
}
