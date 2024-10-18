package main

import "strategy_pattern/internal/sort"

func main() {
	numbs := []int{1, 2, 3, 4, 5, 6, 7}

	sorter := sort.NewSorter(sort.NewMergeSortStrategy())
	sorter.SortNumbers(numbs)

	sorter.SetStrategy(sort.NewQuickSortStrategy())
	sorter.SortNumbers(numbs)
}
