package main

import(
	"aichallenge/internal/week1_1"
	"aichallenge/internal/common"
)

func main() {
	tuiConfig := week1_1.TUIConfig{
		Prompt: "Введите сообщение: ",
	}
	agentContext := common.ReadTextFromFile("../../task/week1_3/system_prompt.txt")
	week1_1.LaunchAgent(agentContext, tuiConfig) //move to internal
}