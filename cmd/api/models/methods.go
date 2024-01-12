package models

import (
	"FitroomTests_v2/cmd/api/helper"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func (reqInfo *Request) GET() io.ReadCloser {
	req, err := http.NewRequest("GET", reqInfo.URL, nil)
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't create new request: %s ", err.Error())
		reqInfo.AddToFailed()
		return nil
	}
	if *reqInfo.Auth != "" {
		req.Header.Set("Authorization", *reqInfo.Auth)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't send a request: %s ", err.Error())
		reqInfo.AddToFailed()
		return nil
	}
	reqInfo.RespBody = resp.Body
	defer resp.Body.Close()
	// Decode response
	if body, err := io.ReadAll(resp.Body); strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") && err == nil && len(body) != 0 {
		if err = json.NewDecoder(bytes.NewReader(body)).Decode(&reqInfo.RespRead); err != nil {
			log.Printf("\033[96m%s\033[0m can't decode response body \n", helper.URLFormat(reqInfo.URL))
			reqInfo.AddToFailed()
			return nil
		}
	}
	// Check status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Printf("\033[96m%s\033[0m response code status: \033[33m%s\033[0m \n", helper.URLFormat(reqInfo.URL), resp.Status)
		reqInfo.AddToFailed()
		return nil
	}
	// End
	log.Printf("\033[96m%s\033[0m response code status: \033[32m%s\033[0m\n", helper.URLFormat(reqInfo.URL), resp.Status)
	if len(reqInfo.Actions) > 0 {
		for _, f := range reqInfo.Actions {
			f()
		}
	}
	return reqInfo.RespBody
}

func (reqInfo *Request) POST() io.ReadCloser {
	req, err := http.NewRequest("POST", reqInfo.URL, bytes.NewBuffer(reqInfo.ReqBody))
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't create new request: %s ", err.Error())
		reqInfo.AddToFailed()
		return nil
	}
	if *reqInfo.Auth != "" {
		req.Header.Set("Authorization", *reqInfo.Auth)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't send a request: %s ", err.Error())
		reqInfo.AddToFailed()
		return nil
	}
	reqInfo.RespBody = resp.Body
	defer resp.Body.Close()
	// Decode response
	if body, err := io.ReadAll(resp.Body); strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") && err == nil && len(body) != 0 {
		if err = json.NewDecoder(bytes.NewReader(body)).Decode(&reqInfo.RespRead); err != nil {
			log.Printf("\033[96m%s\033[0m can't decode response body \n", helper.URLFormat(reqInfo.URL))
			reqInfo.AddToFailed()
			return nil
		}
	}
	// Check status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Printf("\033[96m%s\033[0m response code status: \033[33m%s\033[0m \n", helper.URLFormat(reqInfo.URL), resp.Status)
		reqInfo.AddToFailed()
		return nil
	}
	// End
	log.Printf("\033[96m%s\033[0m response code status: \033[32m%s\033[0m\n", helper.URLFormat(reqInfo.URL), resp.Status)
	if len(reqInfo.Actions) > 0 {
		for _, f := range reqInfo.Actions {
			f()
		}
	}
	return reqInfo.RespBody
}

func (reqInfo *Request) PUT() io.ReadCloser {
	req, err := http.NewRequest("PUT", reqInfo.URL, bytes.NewBuffer(reqInfo.ReqBody))
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't create new request: %s ", err.Error())
		reqInfo.AddToFailed()
		return nil
	}
	if *reqInfo.Auth != "" {
		req.Header.Set("Authorization", *reqInfo.Auth)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't send a request: %s ", err.Error())
		reqInfo.AddToFailed()
		return nil
	}
	reqInfo.RespBody = resp.Body
	defer resp.Body.Close()
	// Decode response
	if body, err := io.ReadAll(resp.Body); strings.HasPrefix(resp.Header.Get("Content-Type"), "application/json") && err == nil && len(body) != 0 {
		if err = json.NewDecoder(bytes.NewReader(body)).Decode(&reqInfo.RespRead); err != nil {
			log.Printf("\033[96m%s\033[0m can't decode response body \n", helper.URLFormat(reqInfo.URL))
			reqInfo.AddToFailed()
			return nil
		}
	}
	// Check status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Printf("\033[96m%s\033[0m response code status: \033[33m%s\033[0m \n", helper.URLFormat(reqInfo.URL), resp.Status)
		reqInfo.AddToFailed()
		return nil
	}
	// End
	log.Printf("\033[96m%s\033[0m response code status: \033[32m%s\033[0m\n", helper.URLFormat(reqInfo.URL), resp.Status)
	if len(reqInfo.Actions) > 0 {
		for _, f := range reqInfo.Actions {
			f()
		}
	}
	return reqInfo.RespBody
}

func (reqInfo *Request) DELETE() io.ReadCloser {
	req, err := http.NewRequest("DELETE", reqInfo.URL, nil)
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't create new request: %s ", err.Error())
		reqInfo.AddToFailed()
		return nil
	}
	if *reqInfo.Auth != "" {
		req.Header.Set("Authorization", *reqInfo.Auth)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't send a request: %s ", err.Error())
		reqInfo.AddToFailed()
		return nil
	}
	defer resp.Body.Close()
	// Check status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		log.Printf("\033[96m%s\033[0m response code status: \033[33m%s\033[0m \n", helper.URLFormat(reqInfo.URL), resp.Status)
		reqInfo.AddToFailed()
		return nil
	}
	// End
	log.Printf("\033[96m%s\033[0m response code status: \033[32m%s\033[0m\n", helper.URLFormat(reqInfo.URL), resp.Status)
	if len(reqInfo.Actions) > 0 {
		for _, f := range reqInfo.Actions {
			f()
		}
	}
	return reqInfo.RespBody
}

var FailedRequestArr []Requester

func (reqInfo *Request) AddToFailed() {
	FailedRequestArr = append(FailedRequestArr, reqInfo)
}
