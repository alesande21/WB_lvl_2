package main

import (
	chainFilter2 "chain_pattern/internal/chainFilter"
	"fmt"
)

func main() {
	filters := chainFilter2.NewFilterChainImpl()
	filters.AddFilter(chainFilter2.NewFilterAge(18)).
		AddFilter(chainFilter2.NewFilterName("Ivan", "Andrei")).
		AddFilter(chainFilter2.NewFilterRiskLevel(0.3)).
		AddFilter(chainFilter2.NewFilterTransfer(1000000))

	from := chainFilter2.NewPerson("Ivan", 21, 700000, 0.29)
	to := chainFilter2.NewPerson("Andrei", 22, 200000, 0.30)

	newRequest := chainFilter2.NewRequest(from, to, 200000)

	filters.DoFilter(newRequest)

	fmt.Printf("Отправитель(From): %s, баланс: %.2f\n", from.GetName(), from.GetBalance())
	fmt.Printf("Отправитель(to): %s, баланс: %.2f\n", to.GetName(), to.GetBalance())
}
