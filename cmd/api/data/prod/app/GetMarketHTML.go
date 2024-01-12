package app

import "FitroomTests_v2/cmd/api/models"

var GetMarketHTML = models.Request{
	URL:    "https://app.fitroom.ru/market",
	Method: "GET",
	Auth:   new(string),
}
