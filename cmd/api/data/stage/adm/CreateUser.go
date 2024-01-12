package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"strconv"
)

var CreateUser = models.Request{
	URL:      "https://back.fitroom.fun/api/admin/user/add",
	Method:   "POST",
	ReqBody:  helper.PrepareReqBody(newUser),
	Auth:     &settings.TokenStage,
	RespRead: &userResp,
	Actions:  []func(){User_F1},
}

var EditUser = models.Request{
	Method:  "PUT",
	ReqBody: helper.PrepareReqBody(editUser),
	Auth:    &settings.TokenStage,
}

var DeleteUser = models.Request{
	Method: "DELETE",
	Auth:   &settings.TokenStage,
}

type UserReq struct {
	Phone              string `json:"phone"`
	Email              string `json:"email"`
	FirstName          string `json:"firstName"`
	SecondName         string `json:"secondName"`
	LastName           string `json:"lastName"`
	BirthDate          string `json:"birthDate"`
	WorkExperience     string `json:"workExperience"`
	Gender             string `json:"gender"`
	VerificationStatus string `json:"verificationStatus"`
	Tags               []int  `json:"tags"`
}

type UserResp struct {
	ID                 int      `json:"id"`
	Email              string   `json:"email"`
	Phone              string   `json:"phone"`
	Roles              []string `json:"roles"`
	Actice             bool     `json:"actice"`
	FirstName          string   `json:"firstName"`
	SecondName         string   `json:"secondName"`
	LastName           string   `json:"lastName"`
	Gender             string   `json:"gender"`
	VerificationStatus string   `json:"verificationStatus"`
	BirthDate          string   `json:"birthDate"`
	CreatedAt          string   `json:"createdAt"`
	UpdatedAt          string   `json:"updatedAt"`
	Gallery            struct {
		ID int `json:"id"`
	} `json:"gallery"`
	Tags             []int  `json:"tags"`
	BonusBalance     string `json:"bonusBalance  "`
	ReferralNeedGift bool   `json:"referralNeedGift"`
}

var newUser = UserReq{
	Phone:              "79483274324",
	Email:              "autotest@gmail.com",
	FirstName:          "TEST",
	SecondName:         "TEST",
	LastName:           "TEST",
	BirthDate:          "1996-01-01",
	WorkExperience:     "1996-01-01",
	Gender:             "male",
	VerificationStatus: "not_verified",
	Tags:               []int{},
}

var editUser = UserReq{
	Phone:              "79483274324",
	Email:              "editted@gmail.com",
	FirstName:          "Editted",
	SecondName:         "Editted",
	LastName:           "Editted",
	BirthDate:          "1996-01-01",
	WorkExperience:     "1996-01-01",
	Gender:             "male",
	VerificationStatus: "not_verified",
	Tags:               []int{},
}

var userResp UserResp

var User_F1 = func() {
	EditUser.URL = "https://back.fitroom.fun/api/admin/user/edit/" + strconv.Itoa(userResp.ID)
	DeleteUser.URL = "https://back.fitroom.fun/api/admin/user/delete/" + strconv.Itoa(userResp.ID)
}
