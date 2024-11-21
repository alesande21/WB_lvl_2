package app

import (
	"fmt"
	"myCut/internal/cut"
	parser2 "myCut/internal/parser"
	"os"
)

func Run() int {
	parser := parser2.NewParser()
	cutter := cut.NewCut()

	flags, err := parser.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Printf("%v", err)
		return -1
	}

	cutter.Start(flags)

	return 0
}
