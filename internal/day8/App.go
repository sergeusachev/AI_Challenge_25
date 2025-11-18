package day8

import (
	"aichallenge/internal/common"
	"bufio"
	"fmt"
	"os"
)

func LaunchApp() {
	networkService := getNetworkService()
	agent := common.NewAgent(networkService)
	presenter := NewPresenter(agent, networkService)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введите сообщение: ")
		if !scanner.Scan() { //why do I need this check?
			break
		}
		input := scanner.Text()
		answer, err := presenter.AskQuestion(input)
		if err != nil {
			fmt.Println("Answer error: ", err)
			continue
		}
		fmt.Println(answer)
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