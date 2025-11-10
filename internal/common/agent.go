package common

import (
	"fmt"
)

type Agent struct {
	model string
	temperature float64
	messages []Message
	networkService *NetworkService
}

func NewAgent(networkService *NetworkService) *Agent {
	return &Agent{
		model: "GigaChat-2",
		temperature: 0.0,
		messages: []Message{},
		networkService: networkService,
	}
}

func (a *Agent) SetContext(agentContext string) {
	agentContextMessage := Message{
		Role:    "system",
		Content: agentContext,
	}

	a.messages = append(a.messages, agentContextMessage)
}

func (a *Agent) SetTemperature(temperature float64) {
	a.temperature = temperature
}

func (a *Agent) AskQuestion(question string) (string, error) {
	a.messages = append(a.messages, Message{
		Role:    "user",
		Content: question,
	})
	fmt.Println("DEBUG, temperature =", a.temperature)
	answerMessage, err := a.networkService.GetCompletion(a.messages, a.model, a.temperature) //extract to data bean
	if err != nil {
		return "", fmt.Errorf("failed to get answer: %w", err)
	}
	a.messages = append(a.messages, *answerMessage)

	return (*answerMessage).Content, nil
}

func (a *Agent) GetContext() string {
	if len(a.messages) == 0 || a.messages[0].Role != "system" {
		return ""
	}

	return a.messages[0].Content
}
