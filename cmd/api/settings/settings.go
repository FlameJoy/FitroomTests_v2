package settings

import (
	"fmt"
	"time"
)

const (
	PhoneNum              = "79643920409"
	SendSMSCode_Stage_URL = "https://back.fitroom.fun/auth/send-code"
	Authorize_Stage_URL   = "https://back.fitroom.fun/auth/authorize/"
)

var TokenStage string
var TokenApp = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpYXQiOjE2ODM4Nzk1NzMsImV4cCI6MTY4NjQ3MTU3Mywicm9sZXMiOlsiUk9MRV9VU0VSIl0sInVzZXJuYW1lIjoiNzk5MjE3MTcxMzQifQ.Zz-o7-CZxZrhcqjXrdhWoihUjB1hgOq626wLnvNOobAfNKeh6St8C5LyO-6FoFgLgZAyOEBHqlaczBTz1tCd7kaviejk0FWzODvLSVJrdK5eHx5dZrIlnAxIgyRx0639YuwvfzzGzN7P1FdCJEsUz7CVcySvJZp3V1UKf98nkf8-iVDl4vneKIltVD5BiytC7QhC_Ez4zLTfkbALgxTPds2VsgA-KY7FbLuYXytOOc52Vh-PGQB4OW8X82IGBiBqn5EPlE-NLEsgj4FAiTlpOOvCTCJCdtjGLVVPWrN_qR8N6YWZUoiaSabzhfjYCiNdQvWFBix5rH2SgkFppyRkfg"                  // 79921717134
var TokenAdm = "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJpYXQiOjE2ODM4OTIwNzYsImV4cCI6MTY4NjQ4NDA3Niwicm9sZXMiOlsiUk9MRV9BRE1JTiIsIlJPTEVfVVNFUiJdLCJ1c2VybmFtZSI6Ijc5NjQzOTIwNDA5In0.XUR-fqwlvZ251ShEEwdAxoiBjnZoCollSKM01KC9UlQKQjH7KKHN_CV4sdbfQcmie6OYnMUrKJ_CJDgRW-rLhnVB0VhyOukQXf2I6Yk8NoshPIcpyaC_vaW4FQagj_uumIfjK-HAcKt-IrOcWnZJr5iO1aSCCWQ5B3RpgOpjY1kBYrGAOWVqxTMZbjU8C2nRqgblxj5F80fqrwaqnzjvVG6wnRDCigz6ijQLmr9BILqfypgJVfjwuPOjT9BK_hDF1MDqXepAuUSwg_tA4tmsBa4ntd1p5wvpEQUS_Hi1oCFupbhpzNJA8n_sJXPMROqJ_qm9sS6LnoGjsVsWGjJsMQ" // 79643920409
var DefaultLength = 90

var Year, Month, Day = time.Now().Date()
var PastDate = time.Date(Year, Month, Day-5, 0, 0, 0, 0, time.Local)
var CurrentDate = time.Date(Year, Month, Day, 0, 0, 0, 0, time.Local)
var FutureDate = time.Date(Year, Month, Day+5, 0, 0, 0, 0, time.Local)
var year, month, day = CurrentDate.Date()
var fYear, fMonth, fDay = FutureDate.Date()
var pYear, pMonth, pDay = PastDate.Date()

// Dates
var C_Date = fmt.Sprintf("%d-%02d-%02d", year, month, day)
var F_Date = fmt.Sprintf("%d-%02d-%02d", fYear, fMonth, fDay)
var P_Date = fmt.Sprintf("%d-%02d-%02d", pYear, pMonth, pDay)

// Rent
var StartDate1 = fmt.Sprintf("%d-%02d-%02d 00:00", fYear, fMonth, fDay)
var EndDate1 = fmt.Sprintf("%d-%02d-%02d 01:00", fYear, fMonth, fDay)

// Training
var StartDate2 = fmt.Sprintf("%d-%02d-%02d 01:00", fYear, fMonth, fDay)
var EndDate2 = fmt.Sprintf("%d-%02d-%02d 02:00", fYear, fMonth, fDay)

// Cardio
var StartDate3 = fmt.Sprintf("%d-%02d-%02d 02:00", fYear, fMonth, fDay)
var EndDate3 = fmt.Sprintf("%d-%02d-%02d 03:00", fYear, fMonth, fDay)

// Default 5 days
var StartDate = fmt.Sprintf("%d-%02d-%02dT00:00:00.760Z", year, month, day)
var EndDate = fmt.Sprintf("%d-%02d-%02dT00:00:00.760Z", fYear, fMonth, fDay)
