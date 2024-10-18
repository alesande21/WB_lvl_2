package chainFilter

type FilterChain interface {
	getFilter(index int) *Filter
	DoFilter(request *Request)
	AddFilter(filter Filter) *FilterChainImpl
}

type FilterChainImpl struct {
	filters            []Filter
	currentIndexFilter int
}

func NewFilterChainImpl() *FilterChainImpl {
	return &FilterChainImpl{filters: make([]Filter, 0), currentIndexFilter: 0}
}

func (f *FilterChainImpl) SetFilters(filters []Filter) {
	f.filters = filters
}

func (f *FilterChainImpl) CurrentIndexFilter() int {
	return f.currentIndexFilter
}

func (f *FilterChainImpl) SetCurrentIndexFilter(currentIndexFilter int) {
	f.currentIndexFilter = currentIndexFilter
}

func (f *FilterChainImpl) getFilter(index int) *Filter {
	if index < 0 || index >= len(f.filters) {
		return nil
	}
	filter := f.filters[index]
	return &filter
}

func (f *FilterChainImpl) DoFilter(request *Request) {
	if f.currentIndexFilter < len(f.filters) {
		currFilter := f.getFilter(f.currentIndexFilter)
		f.currentIndexFilter++
		(*currFilter).DoFilter(request, f)
	}
}

func (f *FilterChainImpl) AddFilter(filter Filter) *FilterChainImpl {
	f.filters = append(f.filters, filter)
	return f
}
