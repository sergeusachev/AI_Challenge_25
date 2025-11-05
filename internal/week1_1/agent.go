package week1_1

import (
	"fmt"
	"aichallenge/internal/common"
)

type Agent struct {
	context string
	messages []common.Message
	networkService *common.NetworkService
}

func NewAgent(agentContext string) (*Agent, error) {
	networkService, err := common.GetNetworkService()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NetworkService: %w", err)
	}

	agentContextMessage := common.Message{
		Role:    "system",
		Content: agentContext,
	}

	return &Agent{
		context: agentContext, //why do I need to store this??
		messages: []common.Message{ agentContextMessage },
		networkService: networkService,
	}, nil
}

func (a *Agent) AskQuestion(question string) (string, error) {
	a.messages = append(a.messages, common.Message{
		Role:    "user",
		Content: question,
	})
	answerMessage, err := a.networkService.GetCompletion(a.messages)
	if err != nil {
		return "", fmt.Errorf("failed to get answer: %w", err)
	}
	a.messages = append(a.messages, *answerMessage)

	return (*answerMessage).Content, nil
}
