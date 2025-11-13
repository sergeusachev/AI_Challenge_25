package day8

import(
	"aichallenge/internal/common"
	"strings"
	"fmt"
)


type Presenter struct {
	NeedCountTokens bool
	IsCompactContextEnabled bool
	Agent *common.Agent
	NetworkService *common.NetworkService
}

func NewPresenter(agent *common.Agent, networkService *common.NetworkService) *Presenter {
	return &Presenter{
		NeedCountTokens: true,
		IsCompactContextEnabled: true,
		Agent: agent,
		NetworkService: networkService,
	}
}

func (p *Presenter) AskQuestion(question string) (string, error) {	
	var builder strings.Builder
	answer, err := p.Agent.AskQuestion(question)
	builder.WriteString("\n\n")
	builder.WriteString(answer)
	if err != nil {
		return "", err
	}

	if p.NeedCountTokens {
		countByLine, tokensErr := p.countTokens([]string{question, answer}) //shadowing err??
		if tokensErr != nil {
			return "", tokensErr
		}
		builder.WriteString(fmt.Sprintf("\nInput tokens: %d\nOutput tokens: %d\n\n", countByLine[0], countByLine[1]))
	}

	if p.IsCompactContextEnabled && p.NeedCompactContext() {
		compactOrder := "Резюмируй все, о чем мы общались. Перескажи это не более чем в пяти предложениях."
		compactContext, compactErr := p.Agent.AskQuestion(compactOrder)
		if compactErr != nil {
			return "", err
		}
		p.Agent.ClearHistory()
		p.Agent.SetContext(compactContext)
		builder.WriteString(fmt.Sprintf("History was cleared. New context is:\n %s\n\n", compactContext))
	}
	
	return builder.String(), nil
}

func (p *Presenter) NeedCompactContext() bool {
	return p.Agent.GetHistorySize() >= 5
}

func (p *Presenter) countTokens(lines []string) ([]int, error) {
	model := p.Agent.Model
	tokensCountResponse, err := p.NetworkService.GetTokensCount(lines, model) 
	if err != nil {
		return []int{}, err
	}

	return mapTokens(tokensCountResponse), nil
}

func mapTokens(responses []common.TokensCountResponse) []int {
	mapped := make([]int, len(responses))
	for i, v := range responses {
		mapped[i] = v.Tokens
	}
	return mapped
}