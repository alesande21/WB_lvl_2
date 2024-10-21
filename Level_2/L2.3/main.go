package main

import (
	"fmt"
	"strings"
	stack2 "task_2.3/internal/stack"
	"unicode/utf8"
)

type Solution struct{}

func (s *Solution) isNum(r rune) bool {
	if r > 47 && r < 58 {
		return true
	}
	return false
}

func (s *Solution) Solving(str string) string {
	if len(str) == 0 {
		return ""
	}

	stack := make(stack2.Stack[rune], 0)
	var b strings.Builder
	//prevNum := false

	for i := 0; i < len(str); {
		runeValue, width := utf8.DecodeRuneInString(str[i:])
		if s.isNum(runeValue) {
			if stack.Len() == 0 {
				return "(некорректная строка)"
			}

			//val := int(runeValue) - 48
			prevRune := stack.Pop()
			for j := 0; j < int(runeValue)-48; j++ {
				b.WriteRune(prevRune)
			}
		} else {
			if stack.Len() != 0 {
				prevRune := stack.Pop()
				b.WriteRune(prevRune)
			}
			stack.Push(runeValue)
		}

		fmt.Printf("%d %#U starts at byte position %d\n", runeValue, runeValue, i+width)
		//b.WriteRune(runeValue)
		i += width
	}

	if stack.Len() != 0 {
		prevRune := stack.Pop()
		b.WriteRune(prevRune)
	}

	return b.String()
}

func main() {
	var str string = "45"
	//_, err := fmt.Fscanf(os.Stdin, "%s", &str)
	//if err != nil {
	//	fmt.Printf("Ошибка при считывании %s", err)
	//	return
	//}

	var s Solution
	fmt.Printf("Результат: %s", s.Solving(str))
}
