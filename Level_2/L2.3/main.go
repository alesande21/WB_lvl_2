package main

import (
	"fmt"
	"task_2.3/internal/solver"
)

func main() {
	var str string = "45"
	//_, err := fmt.Fscanf(os.Stdin, "%s", &str)
	//if err != nil {
	//	fmt.Printf("Ошибка при считывании %s", err)
	//	return
	//}

	var s solver.Solution
	res, err := s.StringUnpacking(str)
	if err != nil {
		fmt.Printf("Ошибка при распакеовке: %s", err)
	}
	fmt.Printf("Результат: %s", res)
}
