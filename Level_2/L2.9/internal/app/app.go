package app

import (
	"bufio"
	"fmt"
	"myShell/internal/shell"
	"os"
	"strings"
)

func Run() {
	myShell, err := shell.NewMyShell()
	if err != nil {
		fmt.Printf("main-> shell.NewMyShell%s", err)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	var command string
	fmt.Printf("Добро пожаловать в MyShall!\n")
	run := true
	for run {
		var commands []string
		fmt.Printf("MyShell :%s> ", myShell.CurrentLocation())
		in, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("->reader.ReadString: ошибка при считывания команды: %s", err)
			break
		}

		if strings.Contains(in, "|") {
			commands = strings.Split(in, " | ")
		} else {
			commands = append(commands, in)
		}

		for _, arg := range commands {
			command, arg, err = shell.GetCommand(arg)
			run = myShell.ExecCommand(command, arg)
		}

	}
}
