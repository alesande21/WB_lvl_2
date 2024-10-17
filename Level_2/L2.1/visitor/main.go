package main

import (
	"fmt"
	"visitor_pattern/internal/callTaxi"
)

func main() {
	//arr := []Transportable{&Person{
	//	fare: 200,
	//	numb: 2,
	//}, &Animal{
	//	fare: 100,
	//	numb: 1,
	//}, &Luggage{
	//	fare:   120,
	//	weight: 5,
	//}}

	taxi := callTaxi.NewTaxi()

	taxi.VisitPerson(callTaxi.NewPerson(200, 2))

	taxi.VisitAnimal(callTaxi.NewAnimal(150, 1))

	taxi.VisitLuggage(callTaxi.NewLuggage(120, 2))

	fmt.Println("Сумма заказа: ", taxi.GetTotalFare())
}
