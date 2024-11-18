package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in string

	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		in = scanner.Text()
	}

	fmt.Printf("%s", in)

}
