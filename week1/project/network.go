package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	oauthURL       = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	completionsURL = "https://gigachat.devices.sberbank.ru/api/v1/chat/completions"
	rqUID          = "270fee8f-3594-4cb7-b9cb-d0690691f735"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type CompletionRequest struct {
	Model             string    `json:"model"`
	Messages          []Message `json:"messages"`
	RepetitionPenalty float64   `json:"repetition_penalty"`
}

type CompletionResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

func GetRequestToken(oauthToken string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	data := url.Values{}
	data.Set("scope", "GIGACHAT_API_PERS")

	req, err := http.NewRequest("POST", oauthURL, strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+oauthToken)
	req.Header.Set("RqUID", rqUID)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var tokenResp TokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return "", err
	}

	return tokenResp.AccessToken, nil
}

func GetCompletion(requestToken, userQuestion string) (string, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	reqData := CompletionRequest{
		Model: "GigaChat",
		Messages: []Message{
			{
				Role:    "user",
				Content: userQuestion,
			},
		},
		RepetitionPenalty: 1,
	}

	jsonData, err := json.Marshal(reqData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", completionsURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+requestToken)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var completionResp CompletionResponse
	err = json.Unmarshal(body, &completionResp)
	if err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if len(completionResp.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	return completionResp.Choices[0].Message.Content, nil
}
