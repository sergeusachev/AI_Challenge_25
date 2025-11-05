package week1_1

import (
	"bufio"
	"fmt"
	"os"
)

func LaunchAgent(agentContext string) {
	agent, err := NewAgent(agentContext)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка инициализации агента: %v\n", err)
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Пожалуйста введите ваш вопрос: ")

		if !scanner.Scan() {
			break
		}

		question := scanner.Text()

		answer, err := agent.AskQuestion(question)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ошибка при получении ответа: %v\n", err)
			continue
		}

		fmt.Println()
		fmt.Println(answer)
		fmt.Println()
	}
}