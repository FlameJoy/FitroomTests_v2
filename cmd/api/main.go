package main

import (
	"FitroomTests_v2/cmd/api/data/stage/adm"
	"FitroomTests_v2/cmd/api/data/stage/app"
	siteFun "FitroomTests_v2/cmd/api/data/stage/landing"
	"FitroomTests_v2/cmd/api/models"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// var (
// 	regressLanding = flag.Bool("landing", false, "Regress landing")
// 	regressApp     = flag.Bool("app", false, "Regress app.Stage.fun")
// 	regressAdm     = flag.Bool("adm", false, "Regress adm.stage.fun")
// 	regressFailed  = flag.Bool("failed", false, "Resend failed requests")
// 	special        = flag.Bool("special", false, "Special requests")
// )

func main() {
	flag.Parse()
	// models.Authorization()
	// switch {
	// case *regressLanding:
	// 	models.Regress(landing.RequestList)
	// 	models.SaveFailedRequests()
	// case *regressApp:
	// 	models.Regress(app.RequestList)
	// 	models.SaveFailedRequests()
	// case *regressAdm:
	// 	models.Regress(adm.RequestList)
	// 	models.SaveFailedRequests()
	// case *regressFailed:
	// 	models.LoadFailedRequests()
	// 	models.Regress(models.FailedRequestArr)
	// 	models.SaveFailedRequests()
	// case *special:
	// 	adm.TriggerPipeline()
	// 	models.SaveFailedRequests()
	// default:
	// 	fmt.Println("Landing")
	// 	models.Regress(landing.RequestList)
	// 	fmt.Println("app.Stage")
	// 	models.Regress(app.RequestList)
	// 	fmt.Println("adm.Stage")
	// 	models.Regress(adm.RequestList)
	// 	models.SaveFailedRequests()
	// }
	fmt.Println("Server start in :8080")
	fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("Authorization"))
	models.Authorization()
	start := time.Now()
	go func(start time.Time) {
		for {
			fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("Landing regress"))
			models.Regress(siteFun.RequestList)
			fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("app.Stage regress"))
			models.Regress(app.RequestList)
			fmt.Printf("\033[103m%s\033[0m\n", strings.ToUpper("adm.Stage regress"))
			models.Regress(adm.RequestList)
			models.SaveFailedRequests()
			fmt.Println("-------------------------------------------------------------------------END-------------------------------------------------------------------------")
			fmt.Println(time.Since(start))
			models.SendReport()
			models.ResetCounters()
			time.Sleep(time.Hour * 1)
		}
	}(start)

	log.Panicln(http.ListenAndServe(":8080", nil))
}
