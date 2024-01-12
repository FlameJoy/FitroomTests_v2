package site

import "FitroomTests_v2/cmd/api/models"

type City struct {
	City uint `json:"city"`
}

type Cities struct {
	Cities     []int `json:"cities"`
	OnlyGroups int   `json:"onlyGroups"`
}

var city = City{
	City: 5,
}

var cities = Cities{
	Cities:     []int{2}, // Pushkin
	OnlyGroups: 1,
}

var RequestList = []models.Requester{
	&GetLanding,
	&GetCityList,
	&GetTagsList,
	&GetPackagesList,
	&GetClubsList,
	&GetTrainersList,
	&Get2GisMap,
	&Get2GisCatalog,
	&GetUserProfile,
	&GetAvailable,
	&GetHistory,
}
