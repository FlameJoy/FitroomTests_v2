package adm

import (
	"FitroomTests_v2/cmd/api/helper"
	"FitroomTests_v2/cmd/api/models"
	"FitroomTests_v2/cmd/api/settings"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var Pipeline = models.Request{
	URL:       "https://gitlab.fitroom.fun/api/v4/projects/11/trigger/pipeline",
	Method:    "POST",
	Auth:      &settings.TokenStage,
	Multipart: formData,
}

// func genBoundary() string {
// 	source := rand.NewSource(time.Now().Unix())
// 	rng := rand.New(source)
// 	return fmt.Sprintf("-----------------------------%d", rng.Int63())
// }

// var respData = struct{}{}
// var boundary = genBoundary()
var formData = fmt.Sprintf(`%s
Content-Disposition: form-data; name="token"

e32a451742d2c51d5262e0b36d2b54
%s
Content-Disposition: form-data; name="ref"

develop
%s--`, "-----------------------------392273126813228294301798566887", "-----------------------------392273126813228294301798566887", "-----------------------------392273126813228294301798566887")

func TriggerPipeline() {
	fmt.Println(formData)
	var resp, err = http.Post(Pipeline.URL, "multipart/form-data; boundary=-----------------------------392273126813228294301798566887", bytes.NewBufferString(Pipeline.Multipart))
	if err != nil {
		log.Printf("\033[96mError:\033[0m can't send a request: %s ", err.Error())
		Pipeline.AddToFailed()
		return
	}
	defer resp.Body.Close()
	if Pipeline.RespRead != nil {
		if err = json.NewDecoder(resp.Body).Decode(Pipeline.RespRead); err != nil {
			log.Printf("\033[96m%s\033[0m can't decode response body \n", helper.URLFormat(Pipeline.URL))
			Pipeline.AddToFailed()
			return
		}
	}
	// Check status
	if resp.StatusCode != http.StatusCreated {
		log.Printf("\033[96m%s\033[0m response code status: \033[33m%s\033[0m \n", helper.URLFormat(Pipeline.URL), resp.Status)
		Pipeline.AddToFailed()
		return
	}
	// End
	log.Printf("\033[96m%s\033[0m response code status: \033[32m%s\033[0m\n", helper.URLFormat(Pipeline.URL), resp.Status)
}
