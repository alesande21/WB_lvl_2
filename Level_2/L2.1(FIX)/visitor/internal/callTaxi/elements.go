package callTaxi

// Person ConcreteElement
type Person struct {
	fare float32
	numb int
}

func NewPerson(fare float32, numb int) *Person {
	return &Person{
		fare: fare,
		numb: numb,
	}
}

func (p *Person) accept(taxi ITaxi) {
	taxi.VisitPerson(p)
}

// Animal ConcreteElement
type Animal struct {
	fare float32
	numb int
}

func NewAnimal(fare float32, numb int) *Animal {
	return &Animal{
		fare: fare,
		numb: numb,
	}
}

func (a *Animal) accept(taxi ITaxi) {
	taxi.VisitAnimal(a)
}

// Luggage ConcreteElement
type Luggage struct {
	fare   float32
	weight float32
}

func NewLuggage(fare float32, weight float32) *Luggage {
	return &Luggage{
		fare:   fare,
		weight: weight,
	}
}

func (l *Luggage) accept(taxi ITaxi) {
	taxi.VisitLuggage(l)
}
