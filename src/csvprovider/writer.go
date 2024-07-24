package csvprovider

import (
	"encoding/csv"
	"os"
)

type CSVWriter struct {
	file   *os.File
	writer *csv.Writer
}

func NewCSVWriter(filePath string) (*CSVWriter, error) {
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return &CSVWriter{
		file:   file,
		writer: csv.NewWriter(file),
	}, nil
}

func (w *CSVWriter) Write(record []string) error {
	return w.writer.Write(record)
}

func (w *CSVWriter) Flush() {
	w.writer.Flush()
}

func (w *CSVWriter) Close() error {
	w.writer.Flush()
	return w.file.Close()
}
