package site

import (
	"FitroomTests_v2/cmd/api/models"
)

var GetTagsList = models.Request{
	URL:    "https://back.fitroom.ru/public/tag/get-list",
	Method: "POST",
	Auth:   new(string),
}
