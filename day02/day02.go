package main

import (
	"github.com/lukashd/aoc2022/lib"
	log "github.com/sirupsen/logrus"
)

type HandShape int

const (
	Rock HandShape = iota
	Paper
	Scisors
)

func (hs HandShape) String() string {
	switch hs {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scisors:
		return "Scisors"
	}
	return "unknown"
}
func (hs HandShape) Value() int {
	switch hs {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scisors:
		return 3
	}
	panic("unknown value")
	return 0
}

type Winner int

const (
	P1   Winner = 0
	Draw Winner = 3
	P2   Winner = 6
)

func (w Winner) String() string {
	switch w {
	case P1:
		return "P1"
	case P2:
		return "P2"
	case Draw:
		return "Draw"
	}
	return "unknown"
}
func (w Winner) Value() int {
	switch w {
	case P1:
		return 0
	case P2:
		return 6
	case Draw:
		return 3
	}
	panic("Unknown Value")
	return 0
}

func play(p1, p2 HandShape) Winner {
	var w Winner
	if p1 == p2 {
		w = Draw
	} else if (p1 == Rock && p2 == Scisors) || (p1 == Paper && p2 == Rock) || (p1 == Scisors && p2 == Paper) {
		w = P1
	} else if (p1 == Rock && p2 == Paper) || (p1 == Paper && p2 == Scisors) || (p1 == Scisors && p2 == Rock) {
		w = P2
	} else {
		log.Panic("Somthing went really wrong")
	}

	log.Debugf("%v vs %v ==> Winner: %v", p1, p2, w)
	return w
}

func parse(s string) HandShape {
	if len(s) != 1 {
		log.Panic("Try to parse something wrong")
	}
	mapping := make(map[string]HandShape)
	mapping["A"] = Rock
	mapping["B"] = Paper
	mapping["C"] = Scisors

	mapping["X"] = Rock
	mapping["Y"] = Paper
	mapping["Z"] = Scisors
	return mapping[s]
}

func parse2(s string) (w Winner) {
	if len(s) != 1 {
		log.Panic("Try to parse something wrong")
	}
	if s == "X" {
		return P1
	}
	if s == "Y" {
		return Draw
	}
	if s == "Z" {
		return P2
	}
	log.Panic("cannot Parse")
	return w
}

func runPart1(test bool) int {
	log.Infof("Run Part One (test = %v)", test)
	var table [][]string
	if test {
		log.SetLevel(log.DebugLevel)
		table = lib.ReadTable("day02/test_a.txt")
	} else {
		log.SetLevel(log.InfoLevel)
		// log.SetLevel(log.DebugLevel)
		table = lib.ReadTable("day02/input_a.txt")
		// log.Panic("Not Implemented")
	}
	log.Debugf("%v \n", table)
	var score int64
	score = 0
	for _, line := range table {
		// log.Debug(len(line))
		p1 := parse(line[0])
		p2 := parse(line[1])
		winner := play(p1, p2)
		// if winner == P1 {
		round := int64(p2.Value() + winner.Value())
		score += round
		log.Debugf("This rounds Score: %d", round)
		log.Debugf("Total Score: %d", score)
		// }
	}
	log.Infof("Result of Day 1 (test = %v): %d", test, score)
	return 0
}

func runPart2(test bool) int {
	log.Infof("Run Part One (test = %v)", test)
	var table [][]string
	if test {
		log.SetLevel(log.DebugLevel)
		table = lib.ReadTable("day02/test_a.txt")
	} else {
		log.SetLevel(log.InfoLevel)
		// log.SetLevel(log.DebugLevel)
		table = lib.ReadTable("day02/input_a.txt")
		// log.Panic("Not Implemented")
	}
	log.Debugf("%v \n", table)
	var score int64
	score = 0
	for _, line := range table {
		// log.Debug(len(line))
		p1 := parse(line[0])
		desiredWinner := parse2(line[1])
		allShapes := []HandShape{Rock, Paper, Scisors}
		var p2 HandShape
		for _, testPlay := range allShapes {
			winner := play(p1, testPlay)
			if winner == desiredWinner {
				p2 = testPlay
			}
		}
		winner := play(p1, p2)
		// if winner == P1 {
		round := int64(p2.Value() + winner.Value())
		score += round
		log.Debugf("This rounds Score: %d", round)
		log.Debugf("Total Score: %d", score)
		// }
	}
	log.Infof("Result of Day 1 (test = %v): %d", test, score)
	return 0
}

func main() {
	log.Info("Welcome to day 02")
	// runPart1(true)
	// runPart1(false) // Wrong ones 13349 (too low), 12340
	runPart2(true)
	runPart2(false)
}
