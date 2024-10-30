package shell

import (
	"fmt"
	"strings"
)

func GetCommand(text string) (string, string, error) {
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
