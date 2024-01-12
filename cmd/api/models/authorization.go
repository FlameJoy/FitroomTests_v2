package models

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/settings"
)

var authReq = struct {
	Code string `json:"code"`
}{
	Code: "2167",
}

type authResp struct {
	Token string `json:"token"`
}

var authData = authResp{}

var reqInfo = Request{
	URL:      "https://back.fitroom.fun/auth/authorize/79643920409",
	Method:   "POST",
	Auth:     new(string),
	ReqBody:  helper.PrepareReqBody(authReq),
	RespRead: &authData,
}

func Authorization() {
	reqInfo.POST()
	// log.Println(authData.Token)
	settings.TokenStage = "Bearer " + authData.Token
	// log.Println(settings.TokenStage)
}
