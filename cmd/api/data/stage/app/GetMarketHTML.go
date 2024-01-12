package app

import "FitroomTests_v2/cmd/api/models"

var GetMarketHTML = models.Request{
	URL:    "https://app.fitroom.fun/market",
	Method: "GET",
	Auth:   new(string),
}
