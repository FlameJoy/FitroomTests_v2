package adm

import (
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetNotification = models.Request{
	URL:    "https://back.fitroom.ru/api/admin/notification/get-list",
	Method: "POST",
	Auth:   &settings.TokenStage,
}
