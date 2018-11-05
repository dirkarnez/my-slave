package main

import (
	"fmt"
	"github.com/dirkarnez/my-slave/impl"
	"log"
)

func main() {
	inPath := "test.txt"
	outPath := "out.txt"

	err := impl.CsvToSqlInsert(inPath,  func(titles map[string]int, record []string) string {
		return fmt.Sprintf("INSERT INTO Customers (CustomerName, ContactName, Address, City, PostalCode, Country) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')", record[titles["xxx"]], record[titles["yyy"]], record[titles["zzz"]])
	}, outPath)

	//err := impl.ApiParamToCsField(inPath, outPath)
	//err := impl.UniqueLines(inPath, outPath)
	if err != nil {
		log.Fatal(err)
	}
}