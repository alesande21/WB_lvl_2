package chainFilter

import (
	"github.com/google/uuid"
)

type Person struct {
	name            string
	age             int
	clientRiskLevel float32
	balance         float64
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) GetAge() int {
	return p.age
}

func (p *Person) GetClientRiskLevel() float32 {
	return p.clientRiskLevel
}

func (p *Person) GetBalance() float64 {
	return p.balance
}

func (p *Person) AddToBalance(value float64) {
	p.balance += value
}

func (p *Person) WithdrawAmount(value float64) {
	p.balance -= value
}

func NewPerson(name string, age int, balance float64, perValueAtRisk float32) *Person {
	return &Person{
		name:            name,
		age:             age,
		clientRiskLevel: perValueAtRisk,
		balance:         balance,
	}
}

type Request struct {
	id             string
	from           *Person
	to             *Person
	transferAmount float64
}

func (r *Request) GetId() string {
	return r.id
}

func (r *Request) GetFrom() *Person {
	return r.from
}

func (r *Request) GetTo() *Person {
	return r.to
}

func (r *Request) GetTransferAmount() float64 {
	return r.transferAmount
}

func (r *Request) MakeTransfer() {
	r.GetFrom().WithdrawAmount(r.transferAmount)
	r.GetTo().AddToBalance(r.transferAmount)
}

func NewRequest(from *Person, to *Person, transferAmount float64) *Request {
	uuidStr := uuid.New().String()
	return &Request{id: uuidStr, from: from, to: to, transferAmount: transferAmount}
}
