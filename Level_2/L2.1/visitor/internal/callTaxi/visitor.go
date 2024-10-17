package callTaxi

// Transportable Element
type Transportable interface {
	accept(taxi ITaxi)
}

// ITaxi Visitor
type ITaxi interface {
	VisitPerson(p *Person)
	VisitAnimal(a *Animal)
	VisitLuggage(l *Luggage)
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
