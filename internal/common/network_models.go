package common

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
	Temperature		  float64	`json:"temperature"`	
	RepetitionPenalty float64   `json:"repetition_penalty"`
}

type CompletionResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

type TokensCountRequest struct {
	Model	string		`json:"model"`
	Input	[]string	`json:"input"`
}

type TokensCountResponse struct {
	Object 		string `json:"object"`
	Tokens 		int `json:"tokens"`
	Characters 	int `json:"characters"`
}