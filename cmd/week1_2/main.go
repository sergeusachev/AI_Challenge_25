package main

import(
	"aichallenge/internal/week1_1"
	"aichallenge/internal/common"
)

func main() {
	agentContext := common.ReadTextFromFile("../../task/week1_2/context.txt")
	week1_1.LaunchAgent(agentContext) //move to internal
}