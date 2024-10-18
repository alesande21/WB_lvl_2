package sort

import "fmt"

type MergeSortStrategy struct{}

func NewMergeSortStrategy() *MergeSortStrategy {
	return &MergeSortStrategy{}
}

func (m *MergeSortStrategy) sort(numbers []int) {
	fmt.Println("Сортировка с использованием MergeSort:", numbers)
}

type QuickSortStrategy struct{}

func NewQuickSortStrategy() *QuickSortStrategy {
	return &QuickSortStrategy{}
}

func (q *QuickSortStrategy) sort(numbers []int) {
	fmt.Println("Сортировка с использованием QuickSort:", numbers)
}
