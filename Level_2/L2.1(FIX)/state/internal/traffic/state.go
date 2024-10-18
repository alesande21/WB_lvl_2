package traffic

import "fmt"

type RedState struct{}

func NewRedState() *RedState {
	return &RedState{}
}

func (s *RedState) Handle() {
	fmt.Println("Светофор красный. Остановитесь!")
}

type YellowState struct{}

func NewYellowState() *YellowState {
	return &YellowState{}
}

func (s *YellowState) Handle() {
	fmt.Println("Светофор желтый. Внимание снизьте скорость!")
}

type GreenState struct{}

func NewGreenState() *GreenState {
	return &GreenState{}
}

func (s *GreenState) Handle() {
	fmt.Println("Светофор зеленый. Можно ехать!")
}
