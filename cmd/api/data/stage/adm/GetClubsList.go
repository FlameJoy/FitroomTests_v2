package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
)

var GetClubsList = models.Request{
	URL:      "https://back.fitroom.fun/api/admin/club/get-list",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(nil),
	RespRead: &clubList,
	Auth:     &settings.TokenStage,
}

var clubList = []struct {
	ID int `json:"id"`
}{}

var clubIDList = []int{}

func GetClubIDList() []int {
	models.Authorization()
	GetClubsList.POST()
	for _, club := range clubList {
		clubIDList = append(clubIDList, club.ID)
	}
	// log.Println(list)
	return clubIDList
}
