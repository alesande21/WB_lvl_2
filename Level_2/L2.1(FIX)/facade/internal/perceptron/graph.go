package perceptron

import "fmt"

type Graph struct {
	name string
}

func GetGraphModel() *Graph {
	return &Graph{name: "Графовая модель: "}
}

func (g *Graph) init() {
	fmt.Println(g.name, "init")
}
func (g *Graph) setInput() {
	fmt.Println(g.name, "setInput")
}

func (g *Graph) forwardFeed() {
	fmt.Println(g.name, "forwardFeed")
}

func (g *Graph) backPropogation() {
	fmt.Println(g.name, "backPropogation")
}

func (g *Graph) weightUpdate() {
	fmt.Println(g.name, "weightUpdate")
}
