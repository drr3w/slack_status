package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type Profile struct {
	Status_text  string `json:"status_text"`
	Status_emoji string `json:"status_emoji"`
}

func setPresence(client *http.Client, params *url.Values) *http.Response {

	url := "https://slack.com/api/users.setPresence"
	//url := "https://httpbin.org/anything"
	req, _ := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF8")
	req.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))
	resp, _ := client.Do(req)
	req.Body.Close()
	return resp

}

func setProfile(client *http.Client, params *url.Values) *http.Response {

	url := "https://slack.com/api/users.profile.set"
	req, _ := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF8")
	req.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))
	resp, _ := client.Do(req)
	req.Body.Close()
	return resp
}

func main() {

	params := url.Values{}

	token := os.Getenv("SLACK_TOKEN")
	profile := Profile{"Hammerin the rocks", ":hammer:"}
	m, _ := json.Marshal(profile)
	params.Set("token", token)
	params.Set("profile", string(m))
	params.Set("presence", "away")
	//presenceUrl := "https://slack.com/api/users.setPresence"
	client := &http.Client{}
	// resp := setProfile(client, &params)
	resp := setPresence(client, &params)
	// req, _ := http.NewRequest("POST", url, strings.NewReader(params.Encode()))
	// req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF8")
	// req.Header.Add("Content-Length", strconv.Itoa(len(params.Encode())))
	// resp, _ := client.Do(req)
	// req.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
