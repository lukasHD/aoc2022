package main

import (
	"fmt"

	"github.com/lukashd/aoc2022/lib"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Warn("Welcom to day one")
	// fname := "day01/test_a.txt"
	fname := "day01/input_a.txt"
	groups := lib.GroupedIntLinesOnEmpty(fname)
	maxval := 0
	for _, group := range groups {
		val := lib.AddArray(group...)
		if val > maxval {
			maxval = val
		}
	}
	fmt.Printf("%+v \n", maxval)
}
