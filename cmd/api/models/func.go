package models

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/settings"
	"bytes"
	"encoding/json"
	"log"
	"os"
	"sync"
	"time"
)

func Regress(reqList []Requester) {
	TotalReqCounter += len(reqList)
	for _, req := range reqList {
		switch req.(*Request).Method {
		case "GET":
			req.GET()
		case "POST":
			req.POST()
		case "PUT":
			req.PUT()
		case "DELETE":
			req.DELETE()
		default:
			log.Printf("\033[96m%s\033[0m method \033[33m%s\033[0m is not supported \n", helper.URLFormat(req.(*Request).URL), req.(*Request).Method)
		}
	}
}

func ParallelRegress(reqList []Requester) {
	var wg sync.WaitGroup
	wg.Add(len((reqList)))
	for _, req := range reqList {
		go func(req Requester) {
			defer wg.Done()
			switch req.(*Request).Method {
			case "GET":
				req.GET()
			case "POST":
				req.POST()
			case "PUT":
				req.PUT()
			case "DELETE":
				req.DELETE()
			default:
				log.Printf("\033[96m%s\033[0m method \033[33m%s\033[0m is not supported \n", helper.URLFormat(req.(*Request).URL), req.(*Request).Method)
			}
		}(req)
	}
	wg.Wait()
}

var TotalReqCounter int

func SendReport() {
	errTotalCounter := len(FailedRequestArr)
	log.Printf("Total: %d \t Failed: %d \n", TotalReqCounter, errTotalCounter)
}

func SaveFailedRequests() {
	file, err := json.MarshalIndent(FailedRequestArr, "", " ")
	if err != nil {
		log.Println("Error: can't MarshalIndent failed requests")
		return
	}
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	if err = os.WriteFile("failedRequests/"+formattedTime+".json", file, 0644); err != nil {
		log.Println("Error: can't save failed requests in file")
		return
	}
}

func ResetCounters() {
	TotalReqCounter = 0
	FailedRequestArr = []Requester{}
}

func LoadFailedRequests() {
	file, err := os.ReadFile("failedRequests/failedRequests.json")
	if err != nil {
		log.Println("Error: can't read file with failed requests")
		return
	}
	dataArr := []struct {
		URL     string
		Method  string
		ReqBody []byte
	}{}
	if err = json.NewDecoder(bytes.NewBuffer(file)).Decode(&dataArr); err != nil {
		log.Println("Error: can't Decode")
	}
	for _, failedReq := range dataArr {
		newReq := Request{
			URL:     failedReq.URL,
			Method:  failedReq.Method,
			ReqBody: failedReq.ReqBody,
			Auth:    &settings.TokenStage,
		}
		FailedRequestArr = append(FailedRequestArr, &newReq)
	}
}
