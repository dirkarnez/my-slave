package impl

import (
	"fmt"
	"github.com/dirkarnez/my-slave/util"
)

func CsvToSqlInsert(inPath, outPath string) error {
	arr := make([]string, 0)

	util.ReadCsv(inPath, func(titles map[string]int, record []string) {
		arr = append(arr, fmt.Sprintf("INSERT INTO Customers (CustomerName, ContactName, Address, City, PostalCode, Country) VALUES ('%s', '%s', '%s', '%s', '%s', '%s')", record[titles["xxx"]], record[titles["yyy"]], record[titles["zzz"]]))
	})

	return util.WriteLines(outPath, arr)
}