package callTaxi

// Transportable Element
type Transportable interface {
	accept(taxi ITaxi)
}

// ITaxi Visitor
type ITaxi interface {
	VisitPerson(p Person)
	VisitAnimal(a Animal)
	VisitLuggage(l Luggage)
	GetTotalFare() float32
}

// Taxi ConcreteVistior
type Taxi struct {
	totalFare float32
}

func NewTaxi() *Taxi {
	return &Taxi{totalFare: 0}
}

func (t *Taxi) VisitPerson(p *Person) {
	t.totalFare += p.fare * float32(p.numb)
}

func (t *Taxi) VisitAnimal(a *Animal) {
	t.totalFare += a.fare * float32(a.numb)
}

func (t *Taxi) VisitLuggage(l *Luggage) {
	t.totalFare += l.fare * l.weight
}

func (t *Taxi) GetTotalFare() float32 {
	return t.totalFare
}

// CabCompanyDispatcher ObjectStructure
type CabCompanyDispatcher struct {
}

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
	taxi.VisitPerson(*p)
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
	taxi.VisitAnimal(*a)
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
	taxi.VisitLuggage(*l)
}
