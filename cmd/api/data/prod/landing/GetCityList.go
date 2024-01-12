package site

import (
	"FitroomTests_v2/cmd/api/models"
)

var GetCityList = models.Request{
	URL:    "https://back.fitroom.ru/public/city/get-list",
	Method: "GET",
	Auth:   new(string),
}
