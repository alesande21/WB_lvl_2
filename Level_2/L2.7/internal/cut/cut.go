package cut

import (
	"bufio"
	"fmt"
	"myCut/internal/parser"
	"os"
	"strings"
)

type MyCut struct{}

func NewCut() *MyCut {
	return &MyCut{}
}

func (mc *MyCut) Start(flags *parser.Flags) error {
	var slcLine []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if !flags.D() && !flags.S() {
			fmt.Printf("%s", line)
			continue
		}

		if flags.D() {
			slcLine = strings.Split(line, flags.DValue())
		}

		if flags.S() {
			if len(slcLine) > 1 {
				if flags.FValue()-1 > len(slcLine) {
					fmt.Println("")
				} else {
					fmt.Printf("%v\n", slcLine[flags.FValue()-1])
				}
			}
		} else {
			if len(slcLine) < 2 {
				fmt.Printf("%v", line)
			} else if flags.FValue()-1 > len(slcLine) {
				fmt.Println("")
			} else {
				fmt.Printf("%v\n", slcLine[flags.FValue()-1])
			}
		}

	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтнения: %v", err)
	}

	return nil
}
