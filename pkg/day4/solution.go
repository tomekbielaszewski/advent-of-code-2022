package day1

import (
	"advent-of-code-2022/pkg"
	"strings"
)

func solution(lines []string) int {
	sum := 0
	for _, line := range lines {
		if fullyContain(parsePair(line)) {
			sum++
		}
	}
	return sum
}

func solution2(lines []string) int {
	sum := 0
	for _, line := range lines {
		if partiallyContain(parsePair(line)) || fullyContain(parsePair(line)) {
			sum++
		}
	}
	return sum
}

func parsePair(line string) (int, int, int, int) {
	sections := strings.Split(line, ",")
	secAB := strings.Split(sections[0], "-")
	secCD := strings.Split(sections[1], "-")
	return pkg.StringToIntNoErr(secAB[0]),
		pkg.StringToIntNoErr(secAB[1]),
		pkg.StringToIntNoErr(secCD[0]),
		pkg.StringToIntNoErr(secCD[1])
}

func fullyContain(a, b, c, d int) bool {
	return a <= c && b >= d ||
		c <= a && d >= b
}

func partiallyContain(a, b, c, d int) bool {
	return isBetween(c, a, b) ||
		isBetween(d, a, b)
}

func isBetween(checked, a, b int) bool {
	return a <= checked && checked <= b
}
