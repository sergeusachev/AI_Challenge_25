package main

import (
	"aichallenge/internal/common"
	"bufio"
	"fmt"
	"os"
)

func main() {
	networkService := getNetworkService()
	agent := common.NewAgent(networkService)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите текст для подсчета токенов: ")
		if !scanner.Scan() { //why do I need this check?
			break
		}
		input := scanner.Text()
		answer, err := agent.AskQuestion(input)
		if err != nil {
			fmt.Println("1 Answer error: ", err)
			continue
		}
		/*
		tokensCountInputResponse, err := networkService.GetTokensCount(input, model)
		if err != nil {
			fmt.Println("1 GetTokensCount error: ", err)
			continue
		}
			*/
		fmt.Println()
		fmt.Println("Вопрос: ", input)
		//fmt.Println("Символов в запросе: ", tokensCountInputResponse.Characters)
		//fmt.Println("Токенов в запросе: ", tokensCountInputResponse.Tokens)
		fmt.Println()
		fmt.Println("Ответ:", answer)
		/*tokensCountOutputResponse, err := networkService.GetTokensCount(answer, model)
		if err != nil {
			fmt.Println("2 GetTokensCount error: ", err)
			continue
		}
		fmt.Println("Символов в ответе: ", tokensCountOutputResponse.Characters)
		fmt.Println("Токенов в ответе: ", tokensCountOutputResponse.Tokens)*/
		fmt.Println()
		fmt.Println()
	}
}

func getNetworkService() *common.NetworkService {
	networkService, err := common.GetNetworkService()
	if err != nil {
		fmt.Println("Network service creation error: ", err)
		panic(err)
	}

	return networkService
}