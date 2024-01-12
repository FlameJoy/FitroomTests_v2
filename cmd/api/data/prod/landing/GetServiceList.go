package site

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
)

var GerSerivceList = models.Request{
	URL:     "https://back.fitroom.ru/public/service/get-list",
	Method:  "POST",
	Auth:    new(string),
	ReqBody: helper.PrepareReqBody(cities),
}
