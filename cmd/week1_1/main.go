package main

import(
	"aichallenge/internal/week1_1"
	"aichallenge/internal/common"
)

func main() {
	tuiConfig := week1_1.TUIConfig{
		Prompt: "Введите Ваш вопрос: ",
	}
	agentContext := common.ReadTextFromFile("../../task/week1_1/context.txt")
	week1_1.LaunchAgent(agentContext, tuiConfig) //move to internal
}
