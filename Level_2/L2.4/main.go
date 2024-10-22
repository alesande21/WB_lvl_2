package main

import (
	"fmt"
	"sort_command/internal/parser"
	"sort_command/internal/sorter"
)

func main() {

	p := parser.NewParser()
	s := sorter.NewSorter()
	flags, filePath, err := p.ParseFlags()
	if err != nil {
		fmt.Printf("Ошибка: fl.ParseFlags %s\n", err)
	}

	s.Run(flags, filePath)
	fmt.Println(fl, filePath)
}
