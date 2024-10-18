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
	fmt.Printf("Получатель(to): %s, баланс: %.2f\n", to.GetName(), to.GetBalance())
}

/*
FilterAge обрабатывает запрос...
FilterName обрабатывает запрос...
FilterRiskLevel обрабатывает запрос...
FilterTransfer обрабатывает запрос...
Запрос fc66c444-34d5-45af-b1e5-3910365e3e3e на сумму 200000.00 выполнен
FilterTransfer завершил обработку...
FilterRiskLevel завершил обработку...
FilterName завершил обработку...
FilterAge завершил обработку...
Отправитель(From): Ivan, баланс: 500000.00
Получатель(to): Andrei, баланс: 400000.00
*/
