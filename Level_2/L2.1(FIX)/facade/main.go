package main

import "facade_pattern/internal/perceptron"

func main() {
	facade := perceptron.GetPerceptronFacade("graph")
	facade.Run()
	facade = perceptron.GetPerceptronFacade("matrix")
	facade.Run()
}

/*
вывод
Графовая модель:  init
Графовая модель:  setInput
Графовая модель:  forwardFeed
Графовая модель:  backPropogation
Графовая модель:  weightUpdate
Матричная модель:  init
Матричная модель:  setInput
Матричная модель:  forwardFeed
Матричная модель:  backPropogation
Матричная модель:  weightUpdate

*/
