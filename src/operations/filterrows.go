package operations

type FilterOperation struct {
	filterFunc func([]string) bool
}

func NewFilterOperation(filterFunc func([]string) bool) *FilterOperation {
	return &FilterOperation{filterFunc: filterFunc}
}

func (f *FilterOperation) Execute(record []string) ([]string, error) {
	if f.filterFunc(record) {
		return record, nil
	}
	return nil, nil
}
