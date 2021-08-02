package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func OpenCsvFile() *os.File {
	csv_file, err := os.Open("Person.csv")
	if err != nil {
		fmt.Println(err)
	}
	return csv_file
}

func ReadCsvFile(csv_file *os.File) [][]string {
	r := csv.NewReader(csv_file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return records

}
