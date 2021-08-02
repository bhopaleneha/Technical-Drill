package main

import (
	"fmt"
)

func main() {
	file := OpenCsvFile()
	rec := ReadCsvFile(file)
	fmt.Println()
	fmt.Println("Given input csv file as a 2d array \n ", rec)
	fmt.Println()
	valid_users := Validate(rec)
	conToJson(rec, valid_users)

}
