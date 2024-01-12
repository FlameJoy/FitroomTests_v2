package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreatePromotion = models.Request{
	URL:      "https://back.fitroom.ru/api/admin/promotion/add",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(newPromotion),
	Auth:     &settings.TokenStage,
	RespRead: &promotionResp,
	Actions:  []func(){Promotion_F1},
}

var EditPromotion = models.Request{
	Method:  "PUT",
	ReqBody: helper.PrepareReqBody(editPromotion),
	Auth:    &settings.TokenStage,
}

var DeletePromotion = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type PromotionReq struct {
	Title         string `json:"title"`
	PromotionType int    `json:"promotionType"`
	Segments      []int  `json:"segments"`
	Clubs         []int  `json:"clubs"`
	Packages      []int  `json:"packages"`
	Discount      int    `json:"discount"`
	Count         int    `json:"count"`
	PersonalCount int    `json:"personalCount"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
}

type PromotionResp struct {
	ID            int  `json:"id"`
	Actice        bool `json:"actice"`
	PromotionType struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
		Type  string `json:"type"`
	} `json:"promotionType"`
	Title         string `json:"title"`
	Discount      string `json:"discount"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	Count         int    `json:"count"`
	PersonalCount int    `json:"personalCount"`
	Packages      []int  `json:"packages"`
	Clubs         []int  `json:"clubs"`
	Segments      []struct {
		ID    int    `json:"id"`
		Title string `json:"title"`
	} `json:"segments"`
	RestrictUse bool `json:"restrictUse"`
}

var newPromotion = PromotionReq{
	Clubs:         []int{},
	Count:         1,
	Discount:      10,
	EndDate:       settings.EndDate,
	Packages:      []int{},
	PersonalCount: 1,
	PromotionType: 1,
	Segments:      []int{1},
	StartDate:     settings.StartDate,
	Title:         "AUTOTEST",
}

var editPromotion = PromotionReq{
	Clubs:         []int{1, 2},
	Count:         20,
	Discount:      20,
	EndDate:       settings.EndDate,
	Packages:      []int{},
	PersonalCount: 2,
	PromotionType: 2,
	Segments:      []int{1, 2},
	StartDate:     settings.StartDate,
	Title:         "EDITTED AUTOTEST",
}

var promotionResp PromotionResp

var Promotion_F1 = func() {
	EditPromotion.URL = "https://back.fitroom.ru/api/admin/promotion/edit/" + strconv.Itoa(promotionResp.ID)
	DeletePromotion.URL = "https://back.fitroom.ru/api/admin/promotion/delete/" + strconv.Itoa(promotionResp.ID)
}
