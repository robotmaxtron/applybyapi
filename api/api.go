package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/bpiddubnyi/multipartstreamer"
	"github.com/rakshazi/applybyapi/tui"
)

// API url
var url = "https://applybyapi.com/"

// POST /gentoken/ response struct
type tokenResponse struct {
	Token string `json:"token"`
}

type applyResponse struct {
	ApplicationId int `json:"application_id"`
}

// Get token to later use
func GetToken(posting int) (string, error) {
	formdata := make(map[string]int)
	formdata["posting"] = posting
	data, err := json.Marshal(formdata)
	if err != nil {
		return "", err
	}
	// NOTE: don't forget about that weird trailing slash,
	// you will receive 405 "wrong verb" error without it!
	req, err := http.NewRequest("POST", url+"gentoken/", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	result := &tokenResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	return result.Token, nil
}

// Send apply with data
func Apply(surveyData *tui.Data) (int, error) {
	ms := multipartstreamer.New()
	data, resume := preprocess(surveyData)

	err := ms.WriteFields(data)
	if err != nil {
		return 0, err
	}

	err = ms.WriteFile("resume", resume)
	if err != nil {
		return 0, err
	}
	req, err := http.NewRequest("POST", url+"apply/", nil)
	if err != nil {
		return 0, err
	}
	ms.SetupRequest(req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	result := &applyResponse{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	return result.ApplicationId, nil
}

// Preprocess data from TUI for multipartstreamer
func preprocess(data *tui.Data) (map[string]string, string) {
	var mapData map[string]string
	var resume string
	marshaledData, _ := json.Marshal(data)
	json.Unmarshal(marshaledData, &mapData)

	if resumePath, ok := mapData["resume"]; ok {
		resume = resumePath
		delete(mapData, "resume")
	}

	return mapData, resume
}
