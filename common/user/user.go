package user

import (
	"encoding/json"
	"log"
)

type User struct {
	Name  string
	Repo  string
	Image string
}

func ParseFromJson(content []byte) *User {
	var data map[string]interface{}

	err := json.Unmarshal(content, &data)
	if err != nil {
		log.Printf("Error parsing response: $s", err)
	}
	return &User{
		Name:  data["login"].(string),
		Repo:  data["html_url"].(string),
		Image: data["avatar_url"].(string),
	}
}
