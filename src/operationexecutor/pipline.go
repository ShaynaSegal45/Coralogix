package operationexecutor

import (
	"fmt"
	"io"

	"src/csvprovider"
	"src/operation"
)

const pageSize = 1

type Pipeline struct {
	operations            []operation.Operation
	aggregatingOperations []operation.AggregatingOperation
}

func NewPipeline() *Pipeline {
	return &Pipeline{}
}

func (p *Pipeline) AddOperation(op operation.Operation) {
	p.operations = append(p.operations, op)
	if aggOperation, ok := op.(operation.AggregatingOperation); ok {
		p.aggregatingOperations = append(p.aggregatingOperations, aggOperation)
	}
}

func (p *Pipeline) Execute(inputFilePath, outputFilePath string) error {
	reader, err := csvprovider.NewCSVReader(inputFilePath)
	if err != nil {
		return err
	}
	defer reader.Close()

	writer, err := csvprovider.NewCSVWriter(outputFilePath)
	if err != nil {
		return err
	}
	defer writer.Close()

	buffer := make([][]string, 0, pageSize)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			if len(buffer) > 0 {
				if err := p.processBuffer(buffer, writer); err != nil {
					return err
				}
			}
			break
		}
		if err != nil {
			return err
		}
		buffer = append(buffer, record)
		if len(buffer) == pageSize {
			if err := p.processBuffer(buffer, writer); err != nil {
				return err
			}
			buffer = buffer[:0]
		}
	}

	writer.Flush()

	for _, aggOperation := range p.aggregatingOperations {
		result := aggOperation.Result()
		err := writer.Write([]string{fmt.Sprintf("%f", result)})
		if err != nil {
			return err
		}
	}

	writer.Flush()

	return nil
}

func (p *Pipeline) processBuffer(buffer [][]string, writer *csvprovider.CSVWriter) error {
	for _, record := range buffer {
		for _, op := range p.operations {
			var err error
			record, err = op.Execute(record)
			if err != nil {
				return err
			}
			if record == nil {
				break
			}
		}

		if record != nil && len(p.aggregatingOperations) == 0 {
			if err := writer.Write(record); err != nil {
				return err
			}
		}
	}

	return nil
}
