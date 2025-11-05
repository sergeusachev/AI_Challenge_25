package main

import (
	"os"
	"strings"
)

func GetOauthToken() string {
	data, err := os.ReadFile("../../secret/oauth_gigachat_token.txt")
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}
