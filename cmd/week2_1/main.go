package main

import (
	"aichallenge/internal/common"
	"bufio"
	"fmt"
	"os"
	"strconv"
)
// Расскажи историю про Угрюмого бычка длиною 5 предложений
func main() {
	prompt := "Введите сообщение: "
	networkService := getNetworkService()
	agent := common.NewAgent(networkService)
	//agent.SetContext("") 
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print(prompt)
		if !scanner.Scan() { //why do I need this check?
			break
		}
		input := scanner.Text()
		if input == "TEMP" {
			checkInputTemperature(input, scanner, agent)
		} else {
			checkInputQuestion(input, agent)
		}
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println()
		fmt.Println("--------------------------------------------------------------------------------------")
	}
}

func checkInputQuestion(question string, agent *common.Agent) {
	answer, err := agent.AskQuestion(question)
	if err != nil {
		fmt.Printf("Ошибка при получении ответа: %v\n", err)
		panic(err)
	}
	fmt.Println()
	fmt.Println(answer)
}

func checkInputTemperature(input string, scanner *bufio.Scanner, agent *common.Agent) {
	fmt.Print("Установите температуру: ")
	scanner.Scan()
	temperature, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Float64 parse error: ", err)
		panic(err)
	}
	agent.SetTemperature(temperature)
	fmt.Println("Установлена температура:", temperature)
	
}

func getNetworkService() *common.NetworkService {
	networkService, err := common.GetNetworkService()
	if err != nil {
		fmt.Println("Network service creation error: ", err)
		panic(err)
	}

	return networkService
}
