package chainFilter

import "fmt"

// Filter интерфейс фильтра
type Filter interface {
	DoFilter(request *Request, chain FilterChain)
}

// FilterName Проверка имени
type FilterName struct {
	approvedNames map[string]struct{}
}

func NewFilterName(names ...string) *FilterName {
	m := make(map[string]struct{})
	for _, name := range names {
		m[name] = struct{}{}
	}
	return &FilterName{approvedNames: m}
}

func (n *FilterName) DoFilter(request *Request, chain FilterChain) {
	fmt.Println("FilterName обрабатывает запрос...")
	_, okFrom := n.approvedNames[request.GetFrom().GetName()]
	_, okTo := n.approvedNames[request.GetTo().GetName()]

	if okFrom && okTo {
		chain.DoFilter(request)
	}

	fmt.Println("FilterName завершил обработку...")
}

type FilterAge struct {
	min int
}

func (a *FilterAge) DoFilter(request *Request, chain FilterChain) {
	fmt.Println("FilterAge обрабатывает запрос...")

	if request.GetFrom().GetAge() >= a.min && request.GetTo().GetAge() >= a.min {
		chain.DoFilter(request)
	}

	fmt.Println("FilterAge завершил обработку...")
}

func NewFilterAge(min int) *FilterAge {
	return &FilterAge{min: min}
}

type FilterRiskLevel struct {
	average float32
}

func (f *FilterRiskLevel) DoFilter(request *Request, chain FilterChain) {
	fmt.Println("FilterRiskLevel обрабатывает запрос...")
	if request.GetFrom().GetClientRiskLevel() <= f.average && request.GetTo().GetClientRiskLevel() <= f.average {
		chain.DoFilter(request)
	}
	fmt.Println("FilterRiskLevel завершил обработку...")

}

func NewFilterRiskLevel(average float32) *FilterRiskLevel {
	return &FilterRiskLevel{average: average}
}

type FilterTransfer struct {
	materiality float64
}

func (f *FilterTransfer) DoFilter(request *Request, chain FilterChain) {
	fmt.Println("FilterTransfer обрабатывает запрос...")
	if request.GetTransferAmount() <= f.materiality && request.GetFrom().GetBalance() >= request.GetTransferAmount() {
		request.MakeTransfer()
		fmt.Printf("Запрос %s на сумму %f выполнен\n", request.GetId(), request.GetTransferAmount())
	}

	fmt.Println("FilterTransfer завершил обработку...")

}

func NewFilterTransfer(materiality float64) *FilterTransfer {
	return &FilterTransfer{materiality: materiality}
}
