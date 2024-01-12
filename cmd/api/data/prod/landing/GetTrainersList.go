package site

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
)

var GetTrainersList = models.Request{
	URL:     "https://back.fitroom.ru/public/user/get-trainers",
	Method:  "POST",
	Auth:    new(string),
	ReqBody: helper.PrepareReqBody(city),
}
