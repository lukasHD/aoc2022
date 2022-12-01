package lib

import (
	"bufio"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func ReadLinesToString(fname string) []string {
	file, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var out []string
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return out
}

func GroupedIntLinesOnEmpty(fname string) [][]int {
	lines := ReadLinesToString(fname)
	var out [][]int
	log.Debug(lines)
	var group []int
	for i, line := range lines {
		log.Debugf("%d Line: %v", i, line)
		if line == "" {
			log.Debug("WE HAVE AN EMPTY LINE")
			log.Debug(group)
			out = append(out, group)
			group = nil
		} else {
			value, err := strconv.Atoi(line)
			if err != nil {
				log.WithError(err).Fatal("Cannot Convert to Int")
			}
			group = append(group, value)
		}
	}
	if len(group) > 0 {
		out = append(out, group)
	}
	return out
}
