package csvprovider

import (
	"encoding/csv"
	"os"
)

type CSVReader struct {
	file   *os.File
	reader *csv.Reader
}

func NewCSVReader(filePath string) (*CSVReader, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	return &CSVReader{
		file:   file,
		reader: csv.NewReader(file),
	}, nil
}

func (r *CSVReader) Read() ([]string, error) {
	return r.reader.Read()
}

func (r *CSVReader) Close() error {
	return r.file.Close()
}
