package sort

type SortingStrategy interface {
	sort(numbs []int)
}

type Sorter struct {
	sortingStrategy SortingStrategy
}

func NewSorter(sortingStrategy SortingStrategy) *Sorter {
	return &Sorter{sortingStrategy: sortingStrategy}
}

func (s *Sorter) SetStrategy(strategy SortingStrategy) {
	s.sortingStrategy = strategy
}

func (s *Sorter) SortNumbers(numbs []int) {
	s.sortingStrategy.sort(numbs)
}
