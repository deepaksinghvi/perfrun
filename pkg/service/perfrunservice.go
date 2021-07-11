package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/deepaksinghvi/perfrun/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type PerfRunRequest struct {
	Suitename      string `json:"suitename"`
	Testname       string `json:"testname"`
	Cloud          string `json:"cloud"`
	Durationinmins int    `json:"durationinmins"`
}

type PerfRunResult struct {
	Suitename      string `json:"suitename"`
	Testname       string `json:"testname"`
	Cloud          string `json:"cloud"`
	Durationinmins int    `json:"durationinmins"`
	ID             int    `json:"id"`
}

func GetPerfRun(targetUrl string, perfRunId int) (PerfRunResult, error) {
	perfRun := PerfRunResult{}

	//targetUrl = "https://be66f1a1-3b11-4cfc-a293-0195be93f244.mock.pstmn.io"

	requestUrl := fmt.Sprintf(`%s/perfrun/%s`,
		targetUrl, perfRunId,
	)
	log.Info(fmt.Sprintf("Requst URL %s", requestUrl))
	response, err := utils.GetRequestWithClient(requestUrl, http.Client{})
	if nil!=err {
		log.Error(err)
		return perfRun,err
	}

	json.Unmarshal(response, &perfRun)
	return perfRun,nil
}

func GetPerfRuns(targetUrl string) []PerfRunResult {
	perfRunResults := []PerfRunResult{}

	//targetUrl = "https://be66f1a1-3b11-4cfc-a293-0195be93f244.mock.pstmn.io"

	requestUrl := fmt.Sprintf(`%s/perfruns`, targetUrl,)
	log.Info(fmt.Sprintf("Requst URL %s", requestUrl))
	response, err := utils.GetRequestWithClient(requestUrl, http.Client{})
	if nil!=err {
		log.Error(err)
		return perfRunResults
	}

	json.Unmarshal(response, &perfRunResults)
	return perfRunResults
}


func CreatePerfRun(targetUrl string, request PerfRunRequest) PerfRunResult{

	payload, _ := json.Marshal(request)
	body := bytes.NewBuffer(payload)
	requestUrl := fmt.Sprintf(`%s/perfrun`,
		targetUrl,
	)
	response, err  :=utils.PostRequest(requestUrl, body, "application/json",http.Client{})
	if( nil != err) {
		fmt.Println(err)
	}
	bytesData, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	perfRunResult := PerfRunResult{}
	json.Unmarshal(bytesData, &perfRunResult)
	fmt.Println(perfRunResult)
	return perfRunResult
}