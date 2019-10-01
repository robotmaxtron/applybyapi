package api

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"bytes"
)

// API url
var url = "https://applybyapi.com/"

// POST /gentoken/ response struct
type tokenResponse struct {
	Token string `json:"token"`
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
	req, err := http.NewRequest("POST", url + "gentoken/", bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return "", err
	}
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
