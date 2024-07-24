package operations

import (
	"fmt"
	"strconv"
)

type SumOperation struct {
	column int
	sum    float64
}

func NewSumOperation(column int) *SumOperation {
	return &SumOperation{column: column, sum: 0}
}

func (s *SumOperation) Execute(record []string) ([]string, error) {
	value, err := strconv.ParseFloat(record[s.column], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse value %s as float: %v", record[s.column], err)
	}
	s.sum += value
	return record, nil
}

func (s *SumOperation) Result() float64 {
	return s.sum
}
