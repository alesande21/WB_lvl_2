package main

import (
	"fmt"
	parser2 "myCut/internal/parser"
	"os"
)

func main() {
	//var in string
	//
	//scanner := bufio.NewScanner(os.Stdin)
	//
	//if scanner.Scan() {
	//	in = scanner.Text()
	//}

	parser := parser2.NewParser()

	flags, line, err := parser.ParseFlags(os.Args[1:])
	if err != nil {
		fmt.Printf("%v", err)
	}

	fmt.Printf("%v", flags)
	fmt.Printf("line: %v", line)

}
