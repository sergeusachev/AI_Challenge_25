package week1_1

import "fmt"

const agentContext = "Ты специалист по новейшей истории России, специализируешься конкретно на истории России c 1991 г. по настоящее время."

type Agent struct {
	networkService *NetworkService
}

func NewAgent() (*Agent, error) {
	networkService, err := GetNetworkService()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize NetworkService: %w", err)
	}

	contextSetAnswer, err := networkService.SetContext(agentContext)
	if err != nil {
		return nil, fmt.Errorf("failed to set context: %w", err)
	}
	_ = contextSetAnswer

	return &Agent{
		networkService: networkService,
	}, nil
}

func (a *Agent) AskQuestion(question string) (string, error) {
	answer, err := a.networkService.GetCompletion(question)
	if err != nil {
		return "", fmt.Errorf("failed to get answer: %w", err)
	}

	return answer, nil
}
