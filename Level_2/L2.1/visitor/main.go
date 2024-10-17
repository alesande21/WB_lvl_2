package main

import (
	"fmt"
	"visitor_pattern/internal/callTaxi"
)

func main() {

	dispatcher := callTaxi.NewCabCompanyDispatcher()

	dispatcher.AddElement(callTaxi.NewPerson(200, 2)).
		AddElement(callTaxi.NewAnimal(150, 1)).
		AddElement(callTaxi.NewLuggage(120, 2))

	taxi := callTaxi.NewTaxi()

	dispatcher.Accept(taxi)

	fmt.Println("Сумма заказа: ", taxi.GetTotalFare())
}
