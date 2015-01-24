package login

import (
	"encoding/json"
	"github.com/gophergala/GopherKombat/common/user"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

const (
	ACCESS_TOKEN_URL = "https://github.com/login/oauth/access_token"
	CLIENT_ID        = "fe6528d512e0697b7883"
	CLIENT_SECRET    = "cec5d9ce37eb963149e5ef9cdb8f445f0d891227"
	GITHUB_API       = "https://api.github.com/user"
)

func LoginCallback(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	code := params["code"][0]

	client := http.Client{}

	form := url.Values{}
	form.Add("client_id", CLIENT_ID)
	form.Add("client_secret", CLIENT_SECRET)
	form.Add("code", code)

	req, err := http.NewRequest("POST", ACCESS_TOKEN_URL, strings.NewReader(form.Encode()))
	if err != nil {
		log.Printf("Error creating request: $s", err)
	}
	req.PostForm = form
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error executing request: $s", err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: $s", err)
	}
	var data map[string]interface{}

	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Printf("Error parsing response: $s", err)
	}
	log.Printf("GitHub data: %s", resp)
	Login(FetchUser(data["access_token"].(string)))
}

func FetchUser(accessToken string) *user.User {
	client := http.Client{}
	url := GITHUB_API + "?access_token=" + accessToken
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: $s", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error executing request: $s", err)
	}
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response: $s", err)
	}
	return user.ParseFromJson(content)

}

func Login(user *user.User) {
	log.Printf("User logged in: %s", user.Name)
}
