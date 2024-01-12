package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreatePackage = models.Request{
	URL:      "https://back.fitroom.ru/api/admin/package/add",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(newPackage),
	Auth:     &settings.TokenStage,
	RespRead: &packageResp,
	Actions:  []func(){Package_F1},
}

var EditPackage = models.Request{
	Method:  "PUT",
	ReqBody: helper.PrepareReqBody(editPackage),
	Auth:    &settings.TokenStage,
}

var DeletePackage = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type PackageReq struct {
	Title               string `json:"title"`
	ServiceProductType  uint   `json:"serviceProductType"`
	TrainingAmount      uint   `json:"trainingAmount"`
	Cost                uint   `json:"cost"`
	Cashback            uint   `json:"cashback"`
	WritingOff          uint   `json:"writingOff"`
	DaysAmount          uint   `json:"daysAmount"`
	CashbackStartDate   string `json:"cashbackStartDate"`
	CashbackEndDate     string `json:"cashbackEndDate"`
	WritingOffStartDate string `json:"writingOffStartDate"`
	WritingOffEndDate   string `json:"writingOffEndDate"`
}

type PackageResp struct {
	ID                 int    `json:"id"`
	Title              string `json:"title"`
	DaysAmount         uint   `json:"daysAmount"`
	TrainingAmount     uint   `json:"trainingAmount"`
	Cost               string `json:"cost"`
	Clubs              []int  `json:"clubs"`
	ServiceProductType struct {
		ID                   int    `json:"id"`
		Title                string `json:"title"`
		VerificationRequired bool   `json:"verificationRequired"`
	} `json:"serviceProductType"`
	IsTrial            bool  `json:"isTrial"`
	PackageConstraints []int `json:"packageConstraints"`
}

var newPackage = PackageReq{
	Title:               "AUTOTEST",
	ServiceProductType:  1,
	TrainingAmount:      5,
	Cost:                1000,
	Cashback:            5,
	WritingOff:          10,
	DaysAmount:          30,
	CashbackStartDate:   "2023-03-01",
	CashbackEndDate:     "2023-03-03",
	WritingOffStartDate: "2023-03-01",
	WritingOffEndDate:   "2023-03-03",
}

var editPackage = PackageReq{
	Title:               "EDITTED AUTOTEST",
	ServiceProductType:  2,
	TrainingAmount:      10,
	Cost:                2000,
	Cashback:            10,
	WritingOff:          20,
	DaysAmount:          40,
	CashbackStartDate:   "2023-03-01",
	CashbackEndDate:     "2023-03-03",
	WritingOffStartDate: "2023-03-01",
	WritingOffEndDate:   "2023-03-03",
}

var packageResp PackageResp

var Package_F1 = func() {
	EditPackage.URL = "https://back.fitroom.ru/api/admin/package/edit/" + strconv.Itoa(packageResp.ID)
	DeletePackage.URL = "https://back.fitroom.ru/api/admin/package/delete/" + strconv.Itoa(packageResp.ID)
}
