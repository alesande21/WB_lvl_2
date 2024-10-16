package perceptron

import "fmt"

type Matrix struct {
	name string
}

func GetMatrixModel() *Matrix {
	return &Matrix{name: "Матричная модель: "}
}

func (g *Matrix) init() {
	fmt.Println(g.name, "init")
}
func (g *Matrix) setInput() {
	fmt.Println(g.name, "setInput")
}

func (g *Matrix) forwardFeed() {
	fmt.Println(g.name, "forwardFeed")
}

func (g *Matrix) backPropogation() {
	fmt.Println(g.name, "backPropogation")
}

func (g *Matrix) weightUpdate() {
	fmt.Println(g.name, "weightUpdate")
}
