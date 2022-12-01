package main

import (
	"fmt"
	"sort"

	"github.com/lukashd/aoc2022/lib"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Warn("Welcom to day one")
	// fname := "day01/test_a.txt"
	fname := "day01/input_a.txt"
	groups := lib.GroupedIntLinesOnEmpty(fname)
	var sums []int
	for _, group := range groups {
		val := lib.AddArray(group...)
		sums = append(sums, val)
	}
	sort.Ints(sums)
	sumslen := len(sums)
	fmt.Printf("%+v \n", sums[sumslen-3:])
	fmt.Printf("%+v \n", lib.AddArray(sums[sumslen-3:]...))
}
