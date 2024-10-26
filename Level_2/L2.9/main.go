package main

import (
	"bufio"
	"fmt"
	"myShell/internal/shell"
	"os"
	"strings"
)

func getCommand(text string) (string, string, error) {
	if text == "" {
		return "", "", fmt.Errorf("введена пуста строка")
	}
	var command string
	if strings.HasPrefix(text, "exit") {
		command = "exit"
		text = strings.TrimPrefix(text, "exit")
	} else if strings.HasPrefix(text, "cd") {
		command = "cd"
		text = strings.TrimPrefix(text, "cd")
	} else if strings.HasPrefix(text, "echo") {
		command = "echo"
		text = strings.TrimPrefix(text, "echo")
	} else if strings.HasPrefix(text, "kill") {
		command = "kill"
		text = strings.TrimPrefix(text, "kill")
	} else if strings.HasPrefix(text, "ps") {
		command = "ps"
		text = strings.TrimPrefix(text, "ps")
	} else if strings.HasPrefix(text, "pwd") {
		command = "pwd"
		text = strings.TrimPrefix(text, "pwd")
	} else {
		return "", "", fmt.Errorf("команда не распознана: %v", text)
	}

	text = strings.TrimSpace(text)

	return command, text, nil
}

func main() {
	myShell, err := shell.NewMyShell()
	if err != nil {
		fmt.Printf("main-> shell.NewMyShell%s", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	var command string
	fmt.Printf("Добро пожаловать в MyShall!\n")
	for {
		fmt.Printf("MyShell :%s>", myShell.CurrentLocation())
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("->reader.ReadString: ошибка при считывания команды: %s", err)
			break
		}

		command, text, err = getCommand(text)

		switch command {
		case "exit":
			fmt.Printf("Выходим из MyShell %s", text)
			return
		case "cd":
			newPath, err := myShell.Cd(text)
			if err != nil {
				fmt.Printf("cd : Не удается найти путь %s, так как он не существует.\n", newPath)
			}
		case "echo":
			if len(text) != 0 {
				fmt.Printf("%v\n", text)
			}
		case "kill":
			fmt.Printf("kill %s\n", text)
		case "ps":
			_, err := myShell.Ps(text)
			if err != nil {
				fmt.Printf("myShell.Ps:%s", err)
			}
		case "pwd":
			curPath, err := myShell.Pwd(text)
			if err != nil {
				fmt.Printf("Get-Location : %s", err)
			} else {
				fmt.Printf("\nPath\n---\n%s\n\n\n", curPath)
			}
		default:
			fmt.Printf("такая команда не обнаружена: %s", command)
		}
	}
	//select {}
}
