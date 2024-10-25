package shell

import (
	"fmt"
	"os"
	"path/filepath"
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

func (ms *MyShell) Cd(path string) (string, error) {
	newPath := filepath.Join(ms.currentLocation, path)
	info, err := os.Stat(newPath)
	if err != nil || os.IsNotExist(err) || !info.IsDir() {
		return newPath, fmt.Errorf("-> Cd-> os.Stat: путь не найден %s", newPath)
	}

	ms.currentLocation = newPath
	return newPath, nil
}

func (ms *MyShell) Pwd(text string) (string, error) {
	if text != "" {
		return "", fmt.Errorf(": Не удается найти позиционный параметр, принимающий аргумент: %s", text)
	}
	return ms.CurrentLocation(), nil
}

func (ms *MyShell) Echo(text string) (string, error) {
	return "", nil
}
