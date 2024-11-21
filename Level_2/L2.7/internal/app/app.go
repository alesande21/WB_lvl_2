package app

import (
	"fmt"
	parser2 "myCut/internal/parser"
	"os"
)

func Run() int {
	parser := parser2.NewParser()

	flags, line, err := parser.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Printf("%v", err)
		return -1
	}

	fmt.Printf("%v\n", flags)
	fmt.Printf("line: %v", line)

	return 0
}
