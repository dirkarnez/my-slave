package impl

import (
	"github.com/dirkarnez/my-slave/container"
	"github.com/dirkarnez/my-slave/util"
	"sort"
)

func UniqueLines(inPath, outPath string) error {
	uniqueLinesSet := container.NewHashSet(false)

	err := util.ReadLines(inPath, func(line string) {
		uniqueLinesSet.Add(line)
	})

	if err != nil {
		return err
	}

	arr := uniqueLinesSet.AsArray()

	sort.Strings(arr)

	return util.WriteLines(outPath, arr)
}