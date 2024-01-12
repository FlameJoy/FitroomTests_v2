package site

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
)

var GetPackagesList = models.Request{
	URL:     "https://back.fitroom.ru/public/package/get-packages",
	Method:  "POST",
	Auth:    new(string),
	ReqBody: helper.PrepareReqBody(city),
}
