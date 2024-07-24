package main

import (
	"log"

	"src/operationexecutor"
	"src/operations"
)

func main() {
	pipeline := operationexecutor.NewPipeline()

	filterFunc := func(record []string) bool {
		return record[0] == "valid"
	}
	pipeline.AddOperation(operations.NewFilterOperation(filterFunc))
	pipeline.AddOperation(operations.NewSelectColumnsOperation([]int{1, 2}))

	// sumOp := operations.NewSumOperation(0)
	// pipeline.AddOperation(sumOp)

	pipeline.AddOperation(operations.NewAvgOperation(0))

	err := pipeline.Execute("input.csv", "output.csv")
	if err != nil {
		log.Fatalf("Error processing input.csv file: %v", err)
	}

	log.Println("CSV processing complete")
}
