package adm

import "FitroomTests_v2/cmd/api/models"

type Page struct {
	Limit string `json:"limit"`
	Page  string `json:"page"`
}

type Paging struct {
	Paging Page `json:"paging"`
}

var pageLimit = Page{
	Limit: "20",
	Page:  "1",
}

var paging = Paging{
	Paging: pageLimit,
}

type Resp struct {
	ID int `json:"id"`
}

var resp Resp

var RequestList = []models.Requester{
	&GetUserByID,
	&GetUsersList,
	&GetMenu,
	&GetPermission,
	&GetCityList,
	&GetDistrictList,
	&GetComplexList,
	&GetClubsList,
	&GetClubByID,
	&GetMetroList,
	&GetWorkplaceList,
	&GetBranchList,
	&GetAppointmentsList,
	&GetProductType,
	&GetPackagesList,
	&GetPromotionList,
	&GetPromotionType,
	&GetTagsList,
	&GetSettingsList,
	&GetCountryList,
	&GetTerms,
	&GetNotifications,
	&GetOrdersList,
	&GetOrders,
	&GetReview,
	&GetGroupAppintments,
	&GetNotification,
	// Book ---------------------------------------------------------
	// &AdminBook,
	// &CancelRent,
	// Create / Edit / Delete elements ------------------------------
	// &CreateBranch,
	// &EditBranch,
	// &DeleteBranch,
	// &CreateClub,
	// &EditClub,
	// &DeleteClub,
	// &CreatePackage,
	// &EditPackage,
	// &DeletePackage,
	// &CreatePromotion,
	// &EditPromotion,
	// &DeletePromotion,
	// &CreateTag,
	// &EditTag,
	// &DeleteTag,
	// &CreateWorkplace,
	// &EditWorkplace,
	// &DeleteWorkplace,
	// &CreateTerm,
	// &EditTerm,
	// &DeleteTerm,
	// &ASD,
	// &CreateUser,
	// &EditUser,
	// &DeleteUser,
	// &CreateGroupTrain,
	// &EditGroupTrain,
	// &DeleteGroupTrain,
	// Bonus --------------------------------------------------------
	// &BonusRefferal,
	// &BonusAdd,
	// &BonusRemove,
	// &BonusDayBeforeDel,
	// &BonusMaxProcent,
	// &BonusMinProcent,
	// &BonusDelProcent,
	// Reports ------------------------------------------------------
	&ReportFinancial,
	&ReportClub,
	&ReportTrainers,
	&ReportBonus,
	&ReportConversion,
}

// Don't use --------
// &CreateCountry,
// &EditCountry,
// &DeleteCountry,
// &CreateCity,
// &EditCity,
// &DeleteCity,
// &CreateDistrict,
// &EditDistrict,
// &DeleteDistrict,
// &CreateComplex,
// &EditComplex,
// &DeleteComplex,
// Don't use --------
