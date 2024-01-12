package site

import "FitroomTests_v2/cmd/api/models"

var GetClubsList = models.Request{
	URL:    "https://back.fitroom.ru/public/club/get-list",
	Method: "POST",
	Auth:   new(string),
}
