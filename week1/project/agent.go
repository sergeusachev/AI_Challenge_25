package main

import "fmt"

type Agent struct {
	requestToken string
}

func NewAgent() (*Agent, error) {
	oauthToken := GetOauthToken()
	requestToken, err := GetRequestToken(oauthToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get request token: %w", err)
	}

	return &Agent{
		requestToken: requestToken,
	}, nil
}

func (a *Agent) AskQuestion(question string) (string, error) {
	answer, err := GetCompletion(a.requestToken, question)
	if err != nil {
		return "", fmt.Errorf("failed to get answer: %w", err)
	}

	return answer, nil
}
