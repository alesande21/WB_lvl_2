package shell

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type MyShell struct {
	currentLocation string
}

func (ms *MyShell) CurrentLocation() string {
	return ms.currentLocation
}

func NewMyShell() (*MyShell, error) {
	curDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("-> os.Getwd: не удалось получить текущую дерикторию: %s", err)
	}
	return &MyShell{currentLocation: curDir}, nil
}

func (ms *MyShell) cd(path string) (string, error) {
	newPath := filepath.Join(ms.currentLocation, path)
	info, err := os.Stat(newPath)
	if err != nil || os.IsNotExist(err) || !info.IsDir() {
		return newPath, fmt.Errorf("-> Cd-> os.Stat: путь не найден %s", newPath)
	}

	ms.currentLocation = newPath
	return newPath, nil
}

func (ms *MyShell) pwd(text string) (string, error) {
	if text != "" {
		return "", fmt.Errorf(": Не удается найти позиционный параметр, принимающий аргумент: %s", text)
	}
	return ms.CurrentLocation(), nil
}

func (ms *MyShell) echo(text string) (string, error) {
	return "", nil
}

func (ms *MyShell) ps(arg string) (string, error) {
	var cmd *exec.Cmd
	if len(arg) != 0 {
		cmd = exec.Command("ps", arg)
	} else {
		cmd = exec.Command("ps")
	}
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("-> cmd.Run(): %s", err)
	}

	return "", nil
}

func (ms *MyShell) kill(arg string) (string, error) {
	if len(arg) == 0 {
		return "", fmt.Errorf("kill: недостаточно аргументов")
	}
	pid, err := strconv.Atoi(arg)

	if err != nil {
		return "", fmt.Errorf("kill: неверный аргумент: %s", arg)
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		return "", fmt.Errorf("kill: %s", err)
	}

	err = proc.Kill()
	if err != nil {
		return "", fmt.Errorf("kill: %s", err)
	}

	return "", nil
}

func (ms *MyShell) ExecCommand(command string, text string) bool {
	switch command {
	case "exit":
		fmt.Printf("Выходим из MyShell %s", text)
		return false
	case "cd":
		newPath, err := ms.cd(text)
		if err != nil {
			fmt.Printf("cd : Не удается найти путь %s, так как он не существует.\n", newPath)
		}
	case "echo":
		if len(text) != 0 {
			fmt.Printf("%v\n", text)
		}
	case "kill":
		_, err := ms.kill(text)
		if err != nil {
			fmt.Printf("%s", err)
		}
	case "ps":
		_, err := ms.ps(text)
		if err != nil {
			fmt.Printf("myShell.Ps:%s", err)
		}
	case "pwd":
		curPath, err := ms.pwd(text)
		if err != nil {
			fmt.Printf("Get-Location : %s", err)
		} else {
			fmt.Printf("\nPath\n---\n%s\n\n\n", curPath)
		}
	default:
		fmt.Printf("такая команда не обнаружена: %s", command)
	}
	return true
}
