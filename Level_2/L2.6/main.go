package main

import (
	"fmt"
	"grep/internal/parser"
)

func main() {
	p := parser.NewParser()

	flags, filePath, err := p.ParseFlags()
	if err != nil {
		fmt.Printf("Ошибка: fl.ParseFlags %s\n", err)
	}

	fmt.Printf("%v\n%v", flags, filePath)
	//err = s.Run(flags, filePath)
	if err != nil {
		fmt.Printf("s.Run%s", err)
		return
	}
}
