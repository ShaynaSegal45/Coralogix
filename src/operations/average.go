package operations

import (
	"fmt"
	"strconv"
)

type AvgOperation struct {
	column int
	sum    float64
	count  int
}

func NewAvgOperation(column int) *AvgOperation {
	return &AvgOperation{column: column, sum: 0, count: 0}
}

func (a *AvgOperation) Execute(record []string) ([]string, error) {
	value, err := strconv.ParseFloat(record[a.column], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse value %s as float: %v", record[a.column], err)
	}
	a.sum += value
	a.count++
	return record, nil
}

func (a *AvgOperation) Result() float64 {
	if a.count == 0 {
		return 0
	}
	return a.sum / float64(a.count)
}
