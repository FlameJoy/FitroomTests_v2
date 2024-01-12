package site

import "FitroomTests_v2/cmd/api/models"

var GetLanding = models.Request{
	URL:    "https://fitroom.ru/",
	Method: "GET",
	Auth:   new(string),
}
