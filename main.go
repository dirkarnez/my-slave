package main

import (
	"github.com/dirkarnez/my-slave/impl"
	"log"
)

func main() {
	inPath := "test.txt"
	outPath := "out.txt"

	err := impl.UniqueLines(inPath, outPath)
	if err != nil {
		log.Fatal(err)
	}
}