package site

import "FitroomTests_v2/cmd/api/models"

var GetLanding = models.Request{
	URL:    "https://fitroom.fun/",
	Method: "GET",
	Auth:   new(string),
}
