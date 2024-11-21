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

	if flags.FValue()-1 > len(slcLine) {
		fmt.Println("")
	} else {
		fmt.Printf("%v\n", slcLine[flags.FValue()])
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if flags.D() {
			slcLine = strings.Split(line, flags.DValue())
		}

		if flags.D() && flags.S() {
			if flags.FValue()-1 > len(slcLine) {
				fmt.Println("")
			} else {
				fmt.Printf("%v\n", slcLine[flags.FValue()])
			}
		}

	}

}
