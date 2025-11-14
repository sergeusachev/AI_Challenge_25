package common

import (
	"fmt"
)

type Agent struct {
	Model string
	temperature float64
	messages []Message
	networkService *NetworkService
}

func NewAgent(networkService *NetworkService) *Agent {
	return &Agent{
		Model: "GigaChat-2", // extract to const
		temperature: 0.0,
		messages: []Message{},
		networkService: networkService,
	}
}

func (a *Agent) SetContext(agentContext string) {
	agentContextMessage := Message{
		Role:    "system", // extract to const
		Content: agentContext,
	}

	a.messages = append(a.messages, agentContextMessage)
}

func (a *Agent) GetHistorySize() int {
	return len(a.messages) - 1
}

func (a *Agent) ClearHistory() {
	a.messages = []Message{}
}

func (a *Agent) GetContext() string {
	if len(a.messages) == 0 || a.messages[0].Role != "system" {
		return ""
	}

	return a.messages[0].Content
}

func (a *Agent) SetTemperature(temperature float64) {
	a.temperature = temperature
}

func (a *Agent) SetModel(model string) {
	a.Model = model
}

func (a *Agent) AskQuestion(question string) (string, error) {
	a.messages = append(a.messages, Message{
		Role:    "user", // extract to const
		Content: question,
	})
	answerMessage, err := a.networkService.GetCompletion(a.messages, a.Model, a.temperature) //extract to data bean
	if err != nil {
		return "", fmt.Errorf("failed to get answer: %w", err) // figure out with error, how to orginize
	}
	a.messages = append(a.messages, *answerMessage)

	return (*answerMessage).Content, nil
}


