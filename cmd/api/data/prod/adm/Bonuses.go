package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var BonusRefferal = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/setting/edit/3",
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(value1),
}

var BonusMaxProcent = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/setting/edit/5",
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(value2),
}

var BonusMinProcent = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/setting/edit/4",
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(value3),
}

var BonusDayBeforeDel = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/setting/edit/7",
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(value4),
}

var BonusDelProcent = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/setting/edit/6",
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(value5),
}

var BonusAdd = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/setting/edit/1",
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(value6),
}

var BonusRemove = models.Request{
	URL:     "https://back.fitroom.ru/api/admin/setting/edit/2",
	Method:  "PUT",
	Auth:    &settings.TokenStage,
	ReqBody: helper.PrepareReqBody(value7),
}

type BonusReq struct {
	Value int `json:"value"`
}

var (
	value1 = BonusReq{Value: 500}
	value2 = BonusReq{Value: 50}
	value3 = BonusReq{Value: 20}
	value4 = BonusReq{Value: 7}
	value5 = BonusReq{Value: 30}
	value6 = BonusReq{Value: 5}
	value7 = BonusReq{Value: 4}
)
