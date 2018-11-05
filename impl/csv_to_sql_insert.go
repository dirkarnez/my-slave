package impl

import (
	"github.com/dirkarnez/my-slave/util"
)

func CsvToSqlInsert(inPath string, onEachRecord func(map[string]int, []string) string, outPath string) error {
	arr := make([]string, 0)

	util.ReadCsv(inPath, func(titles map[string]int, records []string) {
		arr = append(arr, onEachRecord(titles, records))
	})

	return util.WriteLines(outPath, arr)
}