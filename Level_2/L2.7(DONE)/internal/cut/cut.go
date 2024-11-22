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
	//var slcLine []string

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()

		if !flags.S() && !flags.D() {
			fmt.Printf("%v\n", line)
			continue
		}

		parts := strings.Split(line, flags.DValue())
		if flags.S() && len(parts) < 2 {
			continue
		}

		if flags.F() {
			fieldIndex := flags.FValue() - 1

			if len(parts) < 2 {
				fmt.Printf("%v\n", line)
			} else if fieldIndex >= 0 && fieldIndex < len(parts) {
				fmt.Println(parts[fieldIndex])
			} else {
				fmt.Println("")
			}
		} else {
			fmt.Println(line)
		}

	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("ошибка чтнения: %v", err)
	}

	return nil
}

/*
echo -e "field1\tfield2\tfield3\nfield4\tfield5" > input.txt
./myCut -f2 < input.txt
*/
