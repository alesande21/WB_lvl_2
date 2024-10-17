package callTaxi

// CabCompanyDispatcher ObjectStructure
type CabCompanyDispatcher struct {
	elements []Transportable
}

// NewCabCompanyDispatcher новый диспетчер для обработки элементов
func NewCabCompanyDispatcher() *CabCompanyDispatcher {
	return &CabCompanyDispatcher{elements: make([]Transportable, 0)}
}

func (d *CabCompanyDispatcher) AddElement(element Transportable) *CabCompanyDispatcher {
	d.elements = append(d.elements, element)
	return d
}

func (d *CabCompanyDispatcher) Accept(visitor ITaxi) {
	for _, elem := range d.elements {
		elem.accept(visitor)
	}
}
