package app

import "FitroomTests_v2/cmd/api/models"

type Book struct {
	User               uint   `json:"user"`
	Club               uint   `json:"club"`
	Workplace          uint   `json:"workplace"`
	ServiceProductType uint   `json:"serviceProductType"`
	StartDate          string `json:"startDate"`
	EndDate            string `json:"endDate"`
	UserPackage        uint   `json:"userPackage"`
}
type Club struct {
	Club uint `json:"club"`
}

type Cities struct {
	Cities     []int `json:"cities"`
	OnlyGroups int   `json:"onlyGroups"`
}

var cities = Cities{
	Cities:     []int{2}, // Pushkin
	OnlyGroups: 1,
}

var RequestList = []models.Requester{
	&GetProfile,
	&GetAppointments,
	&GetAppointmentsGroup,
	&GetTerm,
	&GetSettingsList,
	&GetClubByID,
	&GetClubList,
	&GetPackageList,
	&GetAvatar,
	&GetHistory,
	&GetMarketHTML,
	&Train,
	&CancelRent,
	&Rent,
	&CancelRent,
	&GroupTrain,
	&CancelRent,
}
