package main

import (
	"fmt"
	"grep/internal/grep"
	"grep/internal/parser"
)

func main() {
	p := parser.NewParser()
	flags, _, err := p.ParseFlags()
	if err != nil {
		fmt.Printf("Ошибка: fl.ParseFlags %s\n", err)
	}

	g := grep.NewGrep(flags)
	//fmt.Printf("%v\n%v", flags, filePath)
	err = g.Run()
	if err != nil {
		fmt.Printf("s.Run%s", err)
		return
	}
}
