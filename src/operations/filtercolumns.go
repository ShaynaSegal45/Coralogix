package operations

type SelectColumnsOperation struct {
	columns []int
}

func NewSelectColumnsOperation(columns []int) *SelectColumnsOperation {
	return &SelectColumnsOperation{columns: columns}
}

func (s *SelectColumnsOperation) Execute(record []string) ([]string, error) {
	var selected []string
	for _, col := range s.columns {
		if col < len(record) {
			selected = append(selected, record[col])
		}
	}
	return selected, nil
}
